package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nathanielvarona/pritunl-api-go"
)

func main() {
	// Initialize the Pritunl API client
	client, err := pritunl.NewClient()
	// Alternatively, you can initialize the client with manual arguments
	// client, err := pritunl.NewClient(&pritunl.Client{
	// 	BaseUrl:   "<PRITUNL BASE URL>",
	// 	ApiToken:  "<PRITUNL API TOKEN>",
	// 	ApiSecret: "<PRITUNL API SECRET>",
	// })
	if err != nil {
		log.Fatal(err)
	}

	// Create a context for the request
	ctx := context.Background()

	// Retrieve the server status
	status, err := client.StatusGet(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Print server status details
	fmt.Println("Server Status:")
	for _, stat := range status {
		fmt.Println("Server Version:", stat.ServerVersion)
		fmt.Println("Local Networks:", stat.LocalNetworks)
		fmt.Println("Host Online:", stat.HostsOnline)
		fmt.Println("------")
	}

	// Marshal server status to JSON
	statusBytes, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		log.Println("Error marshalling status:", err)
	} else {
		fmt.Println("Server Status in JSON:")
		fmt.Println(string(statusBytes))
	}
}
