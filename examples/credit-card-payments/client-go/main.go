package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	searchv1 "github.com/SEArch-Service-Execution-Architecture/search/gen/go/search/v1"
)

type Client struct {
	client searchv1.PrivateMiddlewareServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		client: searchv1.NewPrivateMiddlewareServiceClient(conn),
	}
}

func (c *Client) RegisterChannel(ctx context.Context, contract *searchv1.GlobalContract) (string, error) {
	req := &searchv1.RegisterChannelRequest{
		RequirementsContract: contract,
	}
	resp, err := c.client.RegisterChannel(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.GetChannelId(), nil
}

func (c *Client) SendMessage(ctx context.Context, channelID, recipient string, message *searchv1.AppMessage) (*searchv1.AppSendResponse, error) {
	req := &searchv1.AppSendRequest{
		ChannelId: channelID,
		Recipient: recipient,
		Message:   message,
	}
	return c.client.AppSend(ctx, req)
}

func (c *Client) AppRecv(ctx context.Context, channelID, participant string) (*searchv1.AppRecvResponse, error) {
	req := &searchv1.AppRecvRequest{
		ChannelId:   channelID,
		Participant: participant,
	}
	return c.client.AppRecv(ctx, req)
}

func main() {
	middlewareURL := "middleware-client:11000"
	if len(os.Args) > 1 {
		middlewareURL = os.Args[1]
	}

	// List of book titles
	bookTitles := []string{
		"The Great Gatsby",
		"To Kill a Mockingbird",
		"1984",
		"Pride and Prejudice",
		"Animal Farm",
		"The Catcher in the Rye",
		"Lord of the Flies",
		"Brave New World",
		"The Hobbit",
		"The Fellowship of the Ring",
	}

	// Prompt user to select book titles
	fmt.Println("Please select at least one book to purchase:")
	for i, title := range bookTitles {
		fmt.Printf("%d. %s\n", i+1, title)
	}

	// Get user input for selected book titles
	fmt.Print("Enter the number(s) of the book(s) you want to purchase (comma-separated): ")
	var selectedBooksInput string
	fmt.Scanln(&selectedBooksInput)
	selectedBooksInputArray := strings.Split(selectedBooksInput, ",")
	selectedBooks := make([]int, len(selectedBooksInputArray))
	for i, s := range selectedBooksInputArray {
		fmt.Sscanf(strings.TrimSpace(s), "%d", &selectedBooks[i])
	}

	// Prompt user for shipping address
	fmt.Print("Enter your shipping address: ")
	var shippingAddress string
	fmt.Scanln(&shippingAddress)

	// Print purchase information
	fmt.Println("Purchase Information:")
	fmt.Println("Selected Books:")
	for _, selectedBook := range selectedBooks {
		fmt.Printf("- %s\n", bookTitles[selectedBook-1])
	}
	fmt.Printf("Shipping Address: %s\n", shippingAddress)

	// Load contract.fsa
	contractBytes, err := ioutil.ReadFile("contract.fsa")
	if err != nil {
		log.Fatalf("Failed to read contract file: %v", err)
	}

	contract := &searchv1.GlobalContract{
		Contract:      contractBytes,
		Format:        searchv1.GlobalContractFormat_GLOBAL_CONTRACT_FORMAT_FSA,
		InitiatorName: "ClientApp",
	}

	// Connect to middleware
	conn, err := grpc.Dial("dns:///"+middlewareURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := NewClient(conn)
	ctx := context.Background()

	channelID, err := client.RegisterChannel(ctx, contract)
	if err != nil {
		log.Fatalf("Failed to register channel: %v", err)
	}

	// Send PurchaseRequest
	items := make(map[string]int)
	for _, selectedBook := range selectedBooks {
		bookTitle := bookTitles[selectedBook-1]
		items[bookTitle]++
	}

	itemsJSON, err := json.Marshal(items)
	if err != nil {
		log.Fatalf("Failed to marshal items: %v", err)
	}

	body := fmt.Sprintf(`{"items": %s, "shippingAddress": "%s"}`, string(itemsJSON), shippingAddress)
	msg := &searchv1.AppMessage{
		Type: "PurchaseRequest",
		Body: []byte(body),
	}

	sendResp, err := client.SendMessage(ctx, channelID, "Srv", msg)
	if err != nil || sendResp.GetResult() != searchv1.AppSendResponse_RESULT_OK {
		log.Fatal("Error sending PurchaseRequest")
	}

	// Receive TotalAmount
	recvResp, err := client.AppRecv(ctx, channelID, "Srv")
	if err != nil || recvResp.GetMessage().GetType() != "TotalAmount" {
		log.Fatal("Error receiving TotalAmount")
	}

	var totalAmount float64
	if err := json.Unmarshal(recvResp.GetMessage().GetBody(), &totalAmount); err != nil {
		log.Fatalf("Failed to unmarshal total amount: %v", err)
	}
	fmt.Printf("Total amount: %.2f\n", totalAmount)

	// Get credit card details
	fmt.Print("Enter the credit card number: ")
	var cardNumber string
	fmt.Scanln(&cardNumber)

	fmt.Print("Enter the credit card expiration date: ")
	var cardExpirationDate string
	fmt.Scanln(&cardExpirationDate)

	fmt.Print("Enter the credit card security code: ")
	var cardSecurityCode string
	fmt.Scanln(&cardSecurityCode)

	// Send CardDetailsWithTotalAmount
	body = fmt.Sprintf(`{"card_number": "%s", "card_expirationDate": "%s", "card_cvv": "%s", "total_amount": %.2f}`,
		cardNumber, cardExpirationDate, cardSecurityCode, totalAmount)
	msg = &searchv1.AppMessage{
		Type: "CardDetailsWithTotalAmount",
		Body: []byte(body),
	}

	sendResp, err = client.SendMessage(ctx, channelID, "PPS", msg)
	if err != nil || sendResp.GetResult() != searchv1.AppSendResponse_RESULT_OK {
		log.Fatal("Error sending CardDetailsWithTotalAmount")
	}

	// Receive PaymentNonce
	recvResp, err = client.AppRecv(ctx, channelID, "PPS")
	if err != nil || recvResp.GetMessage().GetType() != "PaymentNonce" {
		log.Fatal("Error receiving PaymentNonce")
	}

	rawJSON := recvResp.GetMessage().GetBody()
	fmt.Printf("Raw JSON: %s\n", rawJSON) // crucial for debugging
	var paymentNonce = string(rawJSON)
	fmt.Printf("Payment nonce: %s\n", paymentNonce)

	// Send PurchaseWithPaymentNonce
	body = fmt.Sprintf(`{"items": %s, "shippingAddress": "%s", "nonce": "%s"}`,
		string(itemsJSON), shippingAddress, paymentNonce)
	msg = &searchv1.AppMessage{
		Type: "PurchaseWithPaymentNonce",
		Body: []byte(body),
	}

	sendResp, err = client.SendMessage(ctx, channelID, "Srv", msg)
	if err != nil || sendResp.GetResult() != searchv1.AppSendResponse_RESULT_OK {
		log.Fatal("Error sending PurchaseWithPaymentNonce")
	}

	// Receive final response
	recvResp, err = client.AppRecv(ctx, channelID, "Srv")
	if err != nil {
		log.Fatalf("Error receiving final response: %v", err)
	}

	switch recvResp.GetMessage().GetType() {
	case "PurchaseOK":
		fmt.Println("Purchase successful!")
	case "PurchaseFail":
		fmt.Println("Purchase failed!")
	default:
		log.Fatal("Error: Unexpected message type received")
	}
}
