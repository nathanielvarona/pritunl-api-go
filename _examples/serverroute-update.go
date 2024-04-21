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

	// Create a ServerRouteRequest object with desired data
	newServerRoute := &pritunl.ServerRouteRequest{
		Network:    "0.0.0.0/0",
		Comment:    "Route Internet Traffic - Updated",
		Nat:        false,
		Advertise:  false,
		NetGateway: false,
	}

	// Create a context for the request
	ctx := context.Background()

	// Update the specified server route
	serverRoutes, err := client.ServerRouteUpdate(ctx, server, route, *newServerRoute)
	if err != nil {
		log.Fatal(err)
	}

	// Print the updated server route details
	fmt.Println("Updated Server Route:")
	for _, serverRoute := range serverRoutes {
		fmt.Println("Server Route ID:", serverRoute.ID)
		fmt.Println("Server Route Network:", serverRoute.Network)
		fmt.Println("------")
	}
}
