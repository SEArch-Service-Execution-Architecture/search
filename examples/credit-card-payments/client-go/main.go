package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/SEArch-Service-Execution-Architecture/search/examples/credit-card-payments/creditcard"
	searchv1 "github.com/SEArch-Service-Execution-Architecture/search/gen/go/search/v1"
)

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
	contract, err := os.ReadFile("contract.fsa")
	if err != nil {
		log.Fatalf("Failed to read contract file: %v", err)
	}

	// Connect to middleware
	conn, err := grpc.NewClient("dns:///"+middlewareURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	order := creditcard.Order{Items: make(map[string]int), ShippingAddress: shippingAddress}
	for _, selectedBook := range selectedBooks {
		order.Items[bookTitles[selectedBook-1]]++
	}

	purchase, err := creditcard.StartPurchase(context.Background(), searchv1.NewPrivateMiddlewareServiceClient(conn), contract, order)
	if err != nil {
		log.Fatalf("Purchase failed: %v", err)
	}
	fmt.Printf("Total amount: %.2f\n", purchase.TotalAmount)

	// Get credit card details
	var card creditcard.Card
	fmt.Print("Enter the credit card number: ")
	fmt.Scanln(&card.Number)
	fmt.Print("Enter the credit card expiration date: ")
	fmt.Scanln(&card.ExpirationDate)
	fmt.Print("Enter the credit card security code: ")
	fmt.Scanln(&card.SecurityCode)

	receipt, err := purchase.Pay(card)
	if err != nil {
		log.Fatalf("Purchase failed: %v", err)
	}
	fmt.Printf("Payment nonce: %s\n", receipt.Nonce)

	if receipt.Approved {
		fmt.Println("Purchase successful!")
	} else {
		fmt.Println("Purchase failed!")
	}
}
