package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/nathanielvarona/pritunl-api-go"
)

func main() {
	// Initialize the Pritunl API client
	client, err := pritunl.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the organization and user IDs from environment variables
	org := os.Getenv("PRITUNL_DATA_ORG")
	user := os.Getenv("PRITUNL_DATA_USER")

	// Create a context for the request
	ctx := context.Background()

	// Retrieve or generate a key for the specified user in the organization
	keys, err := client.KeyGet(ctx, org, user)
	if err != nil {
		log.Fatal(err)
	}

	// Print the key details
	fmt.Println("Key Details:")
	for _, key := range keys {
		fmt.Println("Key View URL:", key.ViewURL)
		fmt.Println("Key URI URL:", key.URIURL)
		fmt.Println("Key URL:", key.KeyURL)
		fmt.Println("------")
	}
}
