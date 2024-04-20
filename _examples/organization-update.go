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
		organization string = os.Getenv("PRITUNL_DATA_ORG")
	)

	// Create an OrganizationRequest object with desired data
	newOrganization := &pritunl.OrganizationRequest{
		Name: "pritunl.orgnaization.updated",
	}

	// Create a context for the request
	ctx := context.Background()

	// Update an Existing Organization
	orgs, err := client.OrganizationUpdate(ctx, organization, *newOrganization)
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
