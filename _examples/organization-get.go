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

	// Create a context for the request
	ctx := context.Background()

	// Retrieve organization
	orgs, err := client.OrganizationGet(ctx)
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

	fmt.Println("####")

	// Retrieve all organizations
	orgs_specifc, err := client.OrganizationGet(ctx, organization)
	if err != nil {
		log.Fatal(err)
	}

	// Struct Output
	for _, org := range orgs_specifc {
		fmt.Println("Organization Name:", org.Name)
		fmt.Println("Organization ID:", org.ID)
		fmt.Println("User Count:", org.UserCount)
		fmt.Println("------")
	}

}
