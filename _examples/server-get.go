package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/nathanielvarona/pritunl-api-go"
)

func main() {
	// Provide authentication credentials as needed for client creation
	// Automaticlly sets from environment variables if present
	client, err := pritunl.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	var (
		server_id string = os.Getenv("PRITUNL_DATA_SERVER")
	)

	// Create a context for the request
	ctx := context.Background()

	// Retrieve all Server
	servers, err := client.ServerGet(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Struct Output
	for _, server := range servers {
		fmt.Println("Server Name:", server.Name)
		fmt.Println("Server ID:", server.ID)
		fmt.Println("Server Status:", server.Status)
		fmt.Println("Server Uptime:", server.Uptime)
		fmt.Println("------")
	}

	fmt.Println("####")

	// Retrieve a server
	servers_specifc, err := client.ServerGet(ctx, server_id)
	if err != nil {
		log.Fatal(err)
	}

	// Struct Output
	for _, server := range servers_specifc {
		fmt.Println("Server Name:", server.Name)
		fmt.Println("Server ID:", server.ID)
		fmt.Println("Server Status:", server.Status)
		fmt.Println("Server Uptime:", server.Uptime)
		fmt.Println("------")
	}

}
