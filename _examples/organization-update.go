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

	// Retrieve the organization ID from an environment variable
	organization := os.Getenv("PRITUNL_DATA_ORG")

	// Create an organization request with the updated name
	newOrganization := &pritunl.OrganizationRequest{
		Name: "pritunl.organization.updated",
	}

	// Create a context for the request
	ctx := context.Background()

	// Update the specified organization
	orgs, err := client.OrganizationUpdate(ctx, organization, *newOrganization)
	if err != nil {
		log.Fatal(err)
	}

	// Print the updated organization details
	fmt.Println("Updated Organization:")
	for _, org := range orgs {
		fmt.Println("Organization Name:", org.Name)
		fmt.Println("Organization ID:", org.ID)
		fmt.Println("User Count:", org.UserCount)
		fmt.Println("------")
	}
}
