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

	// Retrieve the organization and user IDs from environment variables
	organization := os.Getenv("PRITUNL_DATA_ORG")
	user := os.Getenv("PRITUNL_DATA_USER")

	// Create a context for the request
	ctx := context.Background()

	// Create a new user request object with desired data
	updateUser := &pritunl.UserRequest{
		Name:  "new_user",
		Email: "updateuser@domain.dev",
		// Set Disabled to false (default behavior) or any other desired value
		Disabled: false, // Or true if you want the user to be disabled
	}

	// Update an existing user for the organization
	users, err := client.UserUpdate(ctx, organization, user, *updateUser)
	if err != nil {
		log.Fatal(err)
	}

	// Handle the result
	if err != nil {
		fmt.Println("Error updating user:", err)
	} else {
		fmt.Println("Successfully updated users:", users)
	}
}
