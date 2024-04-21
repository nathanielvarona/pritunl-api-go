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

	// Delete the specified organization
	orgs, err := client.OrganizationDelete(ctx, organization)
	if err != nil {
		log.Fatal(err)
	}

	// Print a success message
	fmt.Println("Organization Deleted Successfully")
}
