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

	// Retrieve the server ID from an environment variable
	server := os.Getenv("PRITUNL_DATA_SERVER")

	// Create a context for the request
	ctx := context.Background()

	// Retrieve server routes for the specified server
	serverRoutes, err := client.ServerRouteGet(ctx, server)
	if err != nil {
		log.Fatal(err)
	}

	// Print server route details
	fmt.Println("Server Routes:")
	for _, serverRoute := range serverRoutes {
		fmt.Println("Server Route ID:", serverRoute.ID)
		fmt.Println("Server Route Network:", serverRoute.Network)
		fmt.Println("------")
	}
}
