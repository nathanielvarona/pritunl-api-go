// example/user.go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nathanielvarona/pritunl-api-go"
)

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

	// Create a new UserRequest object with desired data
	newUser := &pritunl.UserRequest{
		Name:  "new_user",
		Email: "newuser@domain.dev",
		// Set Disabled to false (default behavior) or any other desired value
		Disabled: false, // Or true if you want the user to be disabled
	}

	// Call UserCreate to create the user for organization `641351fee8f281432b807a50`
	users, err := client.UserCreate(ctx, "641351fee8f281432b807a50", *newUser)
	if err != nil {
		log.Fatal(err)
	}

	/* PRESENTATION */
	if err != nil {
		// Handle error
		fmt.Println("Error creating user:", err)
	} else {
		fmt.Println("Successfully created users:", users)
	}

}
