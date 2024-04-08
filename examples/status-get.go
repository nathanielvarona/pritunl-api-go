package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nathanielvarona/pritunl-api-go"
)

func main() {
	/* INITIALIZATION AND FETCHING */

	// Provide authentication credentials as needed for client creation
	client, err := pritunl.NewClient( /* provide credentials here if environment variables is not present */ )
	if err != nil {
		log.Fatal(err)
	}

	// Create a context for the request
	ctx := context.Background()

	// Retrieve the server status
	statuses, err := client.Status(ctx)
	if err != nil {
		log.Fatal(err)
	}

	/* PRESENTATION */

	// Per Object Representation
	for _, status := range statuses {
		fmt.Println("Server Version", status.ServerVersion)
		fmt.Println("Loacal Networks:", status.LocalNetworks)
	}

	// JSON Representation
	statusBytes, err := json.MarshalIndent(statuses, "", "  ")
	if err != nil {
		log.Println("Error marshalling status:", err)
	} else {
		fmt.Println(string(statusBytes))
	}

}
