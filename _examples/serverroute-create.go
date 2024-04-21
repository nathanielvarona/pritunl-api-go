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

	// Create a ServerRouteRequest object with desired data
	newServerRoute := &pritunl.ServerRouteRequest{
		Network:    "0.0.0.0/0",
		Comment:    "Route Internet Traffic",
		Nat:        true,
		Advertise:  false,
		NetGateway: false,
	}

	// Create a context for the request
	ctx := context.Background()

	// Create a new server route for the specified server
	serverRoutes, err := client.ServerRouteCreate(ctx, server, *newServerRoute)
	if err != nil {
		log.Fatal(err)
	}

	// Print the created server route details
	fmt.Println("Created Server Route:")
	for _, serverRoute := range serverRoutes {
		fmt.Println("Server Route ID:", serverRoute.ID)
		fmt.Println("Server Route Network:", serverRoute.Network)
		fmt.Println("------")
	}
}
