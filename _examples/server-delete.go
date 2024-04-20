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
		server string = os.Getenv("PRITUNL_DATA_SERVER")
	)

	// Create a context for the request
	ctx := context.Background()

	// Delete Existing Server
	servers, err := client.ServerDelete(ctx, server)
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
}
