package creditcard

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	pb "github.com/SEArch-Service-Execution-Architecture/search/gen/go/search/v1"
)

// PPSContract is the provider contract of the payments service.
const PPSContract = `
.outputs PPS
.state graph
q0 ClientApp ? CardDetailsWithTotalAmount q1
q1 ClientApp ! PaymentNonce q2
q2 Srv ? RequestChargeWithNonce q3
q3 Srv ! ChargeOK q4
q3 Srv ! ChargeFail q5
.marking q0
.end
`

// paymentNonce is the nonce the demo payments service hands out for every payment.
const paymentNonce = "1234567890"

// ServePPS registers the payments service with the middleware behind stub and serves every
// channel it is notified about, until ctx is cancelled (returns nil) or the registration
// stream fails. registered, if not nil, is called with the app ID once the middleware
// acknowledges the registration.
func ServePPS(ctx context.Context, stub pb.PrivateMiddlewareServiceClient, logger *log.Logger, registered func(appID string)) error {
	stream, err := stub.RegisterApp(ctx, &pb.RegisterAppRequest{
		ProviderContract: &pb.LocalContract{
			Contract: []byte(PPSContract),
			Format:   pb.LocalContractFormat_LOCAL_CONTRACT_FORMAT_FSA,
		},
	})
	if err != nil {
		return fmt.Errorf("registering app: %w", err)
	}
	ack, err := stream.Recv()
	if err != nil {
		return fmt.Errorf("receiving registration ack: %w", err)
	}
	logger.Printf("Finished registration. Got AppId %s", ack.GetAppId())
	if registered != nil {
		registered(ack.GetAppId())
	}

	var sessions sync.WaitGroup
	defer sessions.Wait()
	for {
		resp, err := stream.Recv()
		if ctx.Err() != nil {
			return nil
		}
		if err != nil {
			return fmt.Errorf("receiving channel notification: %w", err)
		}
		channelID := resp.GetNotification().GetChannelId()
		logger.Printf("Received Notification. ChannelID: %s", channelID)
		sessions.Go(func() {
			ch := channel{ctx: ctx, stub: stub, id: channelID}
			if err := servePPSSession(ch, logger); err != nil {
				logger.Printf("Session for channel %s failed: %v", channelID, err)
				return
			}
			logger.Printf("Finished PPS program for channel %s", channelID)
		})
	}
}

// servePPSSession runs PPS's side of the contract on one channel.
func servePPSSession(ch channel, logger *log.Logger) error {
	msg, err := ch.recv("ClientApp", "CardDetailsWithTotalAmount")
	if err != nil {
		return err
	}
	// We don't validate the card details, this is just a demo.
	var cardDetails struct {
		CardNumber         string  `json:"card_number"`
		CardExpirationDate string  `json:"card_expirationDate"`
		CardCVV            string  `json:"card_cvv"`
		TotalAmount        float64 `json:"total_amount"`
	}
	if err := json.Unmarshal(msg.GetBody(), &cardDetails); err != nil {
		return fmt.Errorf("parsing CardDetailsWithTotalAmount: %w", err)
	}
	logger.Printf("Received CardDetailsWithTotalAmount: %+v", cardDetails)

	if err := ch.send("ClientApp", "PaymentNonce", []byte(paymentNonce)); err != nil {
		return err
	}

	msg, err = ch.recv("Srv", "RequestChargeWithNonce")
	if err != nil {
		return err
	}
	var charge struct {
		Nonce  string  `json:"nonce"`
		Amount float64 `json:"amount"`
	}
	if err := json.Unmarshal(msg.GetBody(), &charge); err != nil {
		return fmt.Errorf("parsing RequestChargeWithNonce: %w", err)
	}
	logger.Printf("Received RequestChargeWithNonce: %+v", charge)

	var rejection string
	switch {
	case charge.Nonce != paymentNonce:
		rejection = "Invalid nonce"
	case charge.Amount != cardDetails.TotalAmount:
		rejection = "Invalid amount"
	}
	if rejection != "" {
		logger.Printf("Rejecting charge: %s", rejection)
		return ch.send("Srv", "ChargeFail", []byte(rejection))
	}
	return ch.send("Srv", "ChargeOK", []byte(""))
}
