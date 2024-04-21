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

	// Retrieve the organization and server IDs from environment variables
	organization := os.Getenv("PRITUNL_DATA_ORGANIZATION")
	server := os.Getenv("PRITUNL_DATA_SERVER")

	// Create a context for the request
	ctx := context.Background()

	// Detach the specified organization from the specified server
	_, err = client.ServerOrgDetach(ctx, server, organization)
	if err != nil {
		log.Fatal(err)
	}

	// Print a success message
	fmt.Println("Server Organization Detached Successfully")
}
