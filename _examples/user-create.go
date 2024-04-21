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

	// Create a new user request object with desired data
	newUser := &pritunl.UserRequest{
		Name:  "new.user",
		Email: "newuser@domain.dev",
		// Set Disabled to false (default behavior) or any other desired value
		Disabled: false, // Or true if you want the user to be disabled
	}

	// Create a new user for the organization
	users, err := client.UserCreate(ctx, organization, *newUser)
	if err != nil {
		log.Fatal(err)
	}

	// Handle the result
	if err != nil {
		fmt.Println("Error creating user:", err)
	} else {
		fmt.Println("Successfully created users:", users)
	}
}
