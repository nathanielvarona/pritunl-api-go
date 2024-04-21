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
	organization := os.Getenv("PRITUNL_DATA_ORG")
	user := os.Getenv("PRITUNL_DATA_USER")

	// Create a context for the request
	ctx := context.Background()

	// Delete an existing user for the organization
	_, err = client.UserDelete(ctx, organization, user)
	if err != nil {
		log.Fatal(err)
	}

	// Handle the result
	fmt.Println("Successfully deleted user")
}
