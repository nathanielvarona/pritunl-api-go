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

	// Create a ServerHostRequest object with desired data
	newServerHost := &pritunl.ServerHostRequest{
		Server: server,
		ID:     host,
	}

	// Create a context for the request
	ctx := context.Background()

	// Attach the specified host to the specified server
	serverHosts, err := client.ServerHostAttach(ctx, server, host, *newServerHost)
	if err != nil {
		log.Fatal(err)
	}

	// Print the attached server host details
	fmt.Println("Attached Server Host:")
	for _, serverHost := range serverHosts {
		fmt.Println("Server Host ID:", serverHost.ID)
		fmt.Println("Server Host Server", serverHost.Server)
		fmt.Println("------")
	}
}
