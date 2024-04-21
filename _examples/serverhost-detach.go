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

	// Retrieve the host and server IDs from environment variables
	host := os.Getenv("PRITUNL_DATA_HOST")
	server := os.Getenv("PRITUNL_DATA_SERVER")

	// Create a context for the request
	ctx := context.Background()

	// Detach the specified host from the specified server
	serverHosts, err := client.ServerHostDetach(ctx, server, host)
	if err != nil {
		log.Fatal(err)
	}

	// Print a success message
	fmt.Println("Server Host Detached Successfully")
}
