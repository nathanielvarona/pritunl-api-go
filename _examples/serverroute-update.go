package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nathanielvarona/pritunl-api-go"
)

func main() {
	// Provide authentication credentials as needed for client creation
	// Automaticlly sets from environment variables if present
	client, err := pritunl.NewClient()
	if err != nil {
		log.Fatal(err)
	}

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

	// Update a Server Route
	serverroutes, err := client.ServerRouteUpdate(ctx, "6621cfb1af8440ea3c661091", "302e302e302e302f30", *newServerRoute)
	if err != nil {
		log.Fatal(err)
	}

	// Struct Output
	for _, serverroute := range serverroutes {
		fmt.Println("Server Route ID:", serverroute.ID)
		fmt.Println("Server Route Network", serverroute.Network)
		fmt.Println("------")
	}
}
