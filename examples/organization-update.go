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

	// Create an OrganizationRequest object with desired data
	newOrganization := &pritunl.OrganizationRequest{
		Name: "pritunl.orgnaization.updated",
	}

	// Create a context for the request
	ctx := context.Background()

	// Update an Existing Organization
	orgs, err := client.OrganizationUpdate(ctx, "65290fdcec07ec9111bd741e", *newOrganization)
	if err != nil {
		log.Fatal(err)
	}

	// Struct Output
	for _, org := range orgs {
		fmt.Println("Organization Name:", org.Name)
		fmt.Println("Organization ID:", org.ID)
		fmt.Println("User Count:", org.UserCount)
		fmt.Println("------")
	}

}
