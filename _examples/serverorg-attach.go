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

	// Retrieve the organization and server IDs from environment variables
	organization := os.Getenv("PRITUNL_DATA_ORGANIZATION")
	server := os.Getenv("PRITUNL_DATA_SERVER")

	// Create a ServerOrgRequest object with desired data
	newServerOrg := &pritunl.ServerOrgRequest{
		Server: server,
		ID:     organization,
	}

	// Create a context for the request
	ctx := context.Background()

	// Attach the specified organization to the specified server
	serverOrgs, err := client.ServerOrgAttach(ctx, server, organization, *newServerOrg)
	if err != nil {
		log.Fatal(err)
	}

	// Print the attached server organization details
	fmt.Println("Attached Server Organization:")
	for _, serverOrg := range serverOrgs {
		fmt.Println("Server Org ID:", serverOrg.ID)
		fmt.Println("Server Org Server", serverOrg.Server)
		fmt.Println("------")
	}
}
