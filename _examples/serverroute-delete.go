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

	// Create a context for the request
	ctx := context.Background()

	// Update a Server Route
	serverroutes, err := client.ServerRouteDelete(ctx, "6621cfb1af8440ea3c661091", "302e302e302e302f30")
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
