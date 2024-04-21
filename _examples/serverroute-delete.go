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

	// Retrieve the server and route IDs from environment variables
	server := os.Getenv("PRITUNL_DATA_SERVER")
	route := os.Getenv("PRITUNL_DATA_ROUTE")

	// Create a context for the request
	ctx := context.Background()

	// Delete the specified server route
	_, err = client.ServerRouteDelete(ctx, server, route)
	if err != nil {
		log.Fatal(err)
	}

	// Print a success message
	fmt.Println("Server Route Deleted Successfully")
}
