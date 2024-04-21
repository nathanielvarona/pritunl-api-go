package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nathanielvarona/pritunl-api-go"
)

func main() {
	// Initialize the Pritunl API client
	client, err := pritunl.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new organization request
	newOrganization := &pritunl.OrganizationRequest{
		Name: "pritunl.organization.new",
	}

	// Create a context for the request
	ctx := context.Background()

	// Create a new organization
	orgs, err := client.OrganizationCreate(ctx, *newOrganization)
	if err != nil {
		log.Fatal(err)
	}

	// Print the created organization details
	fmt.Println("Created Organization:")
	for _, org := range orgs {
		fmt.Println("Organization Name:", org.Name)
		fmt.Println("Organization ID:", org.ID)
		fmt.Println("User Count:", org.UserCount)
		fmt.Println("------")
	}
}
