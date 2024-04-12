// example/user.go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nathanielvarona/pritunl-api-go"
)

// Include UserResponse struct definition here or import from its file

func main() {
	/* INITIALIZATION */
	// Provide authentication credentials as needed for client creation
	// Automaticlly sets from environment variables if present
	client, err := pritunl.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Create a OrganizationRequest object with desired data
	newOrganization := &pritunl.OrganizationRequest{
		Name: "pritunl.orgnaization.new",
	}

	// Create a context for the request
	ctx := context.Background()

	// Create a New Organization
	orgs, err := client.OrganizationCreate(ctx, *newOrganization)
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
