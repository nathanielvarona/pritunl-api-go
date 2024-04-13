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

	// Retrieve server
	servers, err := client.ServerGet(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Struct Output
	for _, server := range servers {
		fmt.Println("Server Name:", server.Name)
		fmt.Println("Server Status:", server.Status)
		fmt.Println("Server Uptime:", server.Uptime)
		fmt.Println("------")
	}

	fmt.Println("####")

	// Retrieve all servers
	servers_specifc, err := client.ServerGet(ctx, "641358a8e8f281432b807e62")
	if err != nil {
		log.Fatal(err)
	}

	// Struct Output
	for _, server := range servers_specifc {
		fmt.Println("Server Name:", server.Name)
		fmt.Println("Server Status:", server.Status)
		fmt.Println("Server Uptime:", server.Uptime)
		fmt.Println("------")
	}

}
