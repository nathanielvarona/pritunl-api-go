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
		user         string = os.Getenv("PRITUNL_DATA_USER")
	)

	// Create a context for the request
	ctx := context.Background()

	// Create a new UserRequest object with desired data
	updateUser := &pritunl.UserRequest{
		Name:  "new_user",
		Email: "updateuser@domain.dev",
		// Set Disabled to false (default behavior) or any other desired value
		Disabled: false, // Or true if you want the user to be disabled
	}

	// UserUpdate update existing user for organization
	users, err := client.UserUpdate(ctx, organization, user, *updateUser)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		// Handle error
		fmt.Println("Error updating user:", err)
	} else {
		fmt.Println("Successfully updating users:", users)
	}

}
