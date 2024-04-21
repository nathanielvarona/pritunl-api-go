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

	// Retrieve all servers
	servers, err := client.ServerGet(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Print all server details
	fmt.Println("All Servers:")
	for _, server := range servers {
		fmt.Println("Server Name:", server.Name)
		fmt.Println("Server ID:", server.ID)
		fmt.Println("Server Status:", server.Status)
		fmt.Println("Server Uptime:", server.Uptime)
		fmt.Println("------")
	}

	fmt.Println("####")

	// Retrieve the specified server
	serversSpecific, err := client.ServerGet(ctx, server)
	if err != nil {
		log.Fatal(err)
	}

	// Print the specified server details
	fmt.Println("Specified Server:")
	for _, server := range serversSpecific {
		fmt.Println("Server Name:", server.Name)
		fmt.Println("Server ID:", server.ID)
		fmt.Println("Server Status:", server.Status)
		fmt.Println("Server Uptime:", server.Uptime)
		fmt.Println("------")
	}
}
