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

	// Create a context for the request
	ctx := context.Background()

	// Retrieve all organizations
	orgs, err := client.OrganizationGet(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Print all organization details
	fmt.Println("All Organizations:")
	for _, org := range orgs {
		fmt.Println("Organization Name:", org.Name)
		fmt.Println("Organization ID:", org.ID)
		fmt.Println("User Count:", org.UserCount)
		fmt.Println("------")
	}

	fmt.Println("####")

	// Retrieve the specified organization
	orgsSpecific, err := client.OrganizationGet(ctx, organization)
	if err != nil {
		log.Fatal(err)
	}

	// Print the specified organization details
	fmt.Println("Specified Organization:")
	for _, org := range orgsSpecific {
		fmt.Println("Organization Name:", org.Name)
		fmt.Println("Organization ID:", org.ID)
		fmt.Println("User Count:", org.UserCount)
		fmt.Println("------")
	}

	fmt.Println("####")

	// Print the first organization details
	org := orgsSpecific[0]
	fmt.Println("First Organization:")
	fmt.Println("Organization Name:", org.Name)
	fmt.Println("Organization ID:", org.ID)
	fmt.Println("User Count:", org.UserCount)
	fmt.Println("------")
}
