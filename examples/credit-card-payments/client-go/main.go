package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
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

	// Read whole lines: fmt.Scanln stops at the first space and leaves the rest of the line
	// to be consumed by the following prompts.
	stdin := bufio.NewScanner(os.Stdin)

	// Get user input for selected book titles
	var selectedBooks []int
	for selectedBooks == nil {
		input := prompt(stdin, "Enter the number(s) of the book(s) you want to purchase (comma-separated): ")
		var err error
		selectedBooks, err = parseSelection(input, len(bookTitles))
		if err != nil {
			fmt.Println(err)
		}
	}

	// Prompt user for shipping address
	shippingAddress := prompt(stdin, "Enter your shipping address: ")

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
	receipt, err := purchase.Pay(creditcard.Card{
		Number:         prompt(stdin, "Enter the credit card number: "),
		ExpirationDate: prompt(stdin, "Enter the credit card expiration date: "),
		SecurityCode:   prompt(stdin, "Enter the credit card security code: "),
	})
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

// prompt prints msg and returns the next line of input, without surrounding whitespace. It
// exits if the input ends.
func prompt(in *bufio.Scanner, msg string) string {
	fmt.Print(msg)
	if !in.Scan() {
		if err := in.Err(); err != nil {
			log.Fatalf("Failed to read input: %v", err)
		}
		log.Fatal("Unexpected end of input")
	}
	return strings.TrimSpace(in.Text())
}

// parseSelection parses a comma-separated list of book numbers between 1 and numBooks.
func parseSelection(input string, numBooks int) ([]int, error) {
	var selected []int
	for s := range strings.SplitSeq(input, ",") {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil || n < 1 || n > numBooks {
			return nil, fmt.Errorf("%q is not a book number between 1 and %d", strings.TrimSpace(s), numBooks)
		}
		selected = append(selected, n)
	}
	return selected, nil
}
