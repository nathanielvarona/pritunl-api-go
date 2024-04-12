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

	// Create a context for the request
	ctx := context.Background()

	// Retrieve all organizations
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
	orgs_specifc, err := client.OrganizationGet(ctx, "641351fee8f281432b807a50")
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
