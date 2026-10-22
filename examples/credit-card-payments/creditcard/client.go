// Package creditcard implements the Go participants of the credit card payments example: the
// client application (ClientApp) and the payments service (PPS). The programs in client-go and
// payments-service add the command line and the prompts, and the end-to-end test drives the
// same code.
package creditcard

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/SEArch-Service-Execution-Architecture/search/gen/go/search/v1"
)

// Order is what the client wants to buy.
type Order struct {
	Items           map[string]int // book title -> quantity
	ShippingAddress string
}

// Card holds the credit card details sent to the payments service.
type Card struct {
	Number         string
	ExpirationDate string
	SecurityCode   string
}

// Receipt is the outcome of a purchase.
type Receipt struct {
	Nonce    string
	Approved bool // PurchaseOK if true, PurchaseFail if false.
}

// Purchase is a purchase in progress: the channel is open and the backend has sent the total.
type Purchase struct {
	ChannelID   string
	TotalAmount float64

	ch    channel
	order Order
}

// StartPurchase runs the first half of ClientApp's side of the global contract (the
// contract.fsa file of the example) through the middleware behind stub: it opens a channel,
// sends the order and receives the total amount computed by the backend. Pay completes the
// purchase.
func StartPurchase(ctx context.Context, stub pb.PrivateMiddlewareServiceClient, contract []byte, order Order) (*Purchase, error) {
	reg, err := stub.RegisterChannel(ctx, &pb.RegisterChannelRequest{
		RequirementsContract: &pb.GlobalContract{
			Contract:      contract,
			Format:        pb.GlobalContractFormat_GLOBAL_CONTRACT_FORMAT_FSA,
			InitiatorName: "ClientApp",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("registering channel: %w", err)
	}
	p := &Purchase{
		ChannelID: reg.GetChannelId(),
		ch:        channel{ctx: ctx, stub: stub, id: reg.GetChannelId()},
		order:     order,
	}

	err = p.ch.send("Srv", "PurchaseRequest", map[string]any{
		"items":           order.Items,
		"shippingAddress": order.ShippingAddress,
	})
	if err != nil {
		return nil, err
	}
	msg, err := p.ch.recv("Srv", "TotalAmount")
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(msg.GetBody(), &p.TotalAmount); err != nil {
		return nil, fmt.Errorf("parsing TotalAmount: %w", err)
	}
	return p, nil
}

// Pay runs the rest of the contract: it sends the card details to the payments service,
// confirms the purchase with the backend and returns the outcome. It uses the context passed
// to StartPurchase.
func (p *Purchase) Pay(card Card) (Receipt, error) {
	var receipt Receipt
	err := p.ch.send("PPS", "CardDetailsWithTotalAmount", map[string]any{
		"card_number":         card.Number,
		"card_expirationDate": card.ExpirationDate,
		"card_cvv":            card.SecurityCode,
		"total_amount":        p.TotalAmount,
	})
	if err != nil {
		return receipt, err
	}
	msg, err := p.ch.recv("PPS", "PaymentNonce")
	if err != nil {
		return receipt, err
	}
	receipt.Nonce = string(msg.GetBody())

	err = p.ch.send("Srv", "PurchaseWithPaymentNonce", map[string]any{
		"items":           p.order.Items,
		"shippingAddress": p.order.ShippingAddress,
		"nonce":           receipt.Nonce,
	})
	if err != nil {
		return receipt, err
	}
	msg, err = p.ch.recv("Srv", "PurchaseOK", "PurchaseFail")
	if err != nil {
		return receipt, err
	}
	receipt.Approved = msg.GetType() == "PurchaseOK"
	return receipt, nil
}

// channel sends and receives the messages of one channel, checking the results.
type channel struct {
	ctx  context.Context
	stub pb.PrivateMiddlewareServiceClient
	id   string
}

// send sends a message whose body is the JSON encoding of body. A []byte body is sent as is.
func (c channel) send(recipient, msgType string, body any) error {
	raw, ok := body.([]byte)
	if !ok {
		var err error
		if raw, err = json.Marshal(body); err != nil {
			return fmt.Errorf("encoding %s: %w", msgType, err)
		}
	}
	resp, err := c.stub.AppSend(c.ctx, &pb.AppSendRequest{
		ChannelId: c.id,
		Recipient: recipient,
		Message:   &pb.AppMessage{Type: msgType, Body: raw},
	})
	if err != nil {
		return fmt.Errorf("sending %s to %s: %w", msgType, recipient, err)
	}
	if resp.GetResult() != pb.AppSendResponse_RESULT_OK {
		return fmt.Errorf("sending %s to %s: %v", msgType, recipient, resp.GetResult())
	}
	return nil
}

// recv receives the next message from participant and checks that it has one of the given types.
func (c channel) recv(participant string, msgTypes ...string) (*pb.AppMessage, error) {
	resp, err := c.stub.AppRecv(c.ctx, &pb.AppRecvRequest{ChannelId: c.id, Participant: participant})
	if err != nil {
		return nil, fmt.Errorf("receiving %v from %s: %w", msgTypes, participant, err)
	}
	msg := resp.GetMessage()
	for _, t := range msgTypes {
		if msg.GetType() == t {
			return msg, nil
		}
	}
	return nil, fmt.Errorf("received %q from %s, expected one of %v", msg.GetType(), participant, msgTypes)
}
