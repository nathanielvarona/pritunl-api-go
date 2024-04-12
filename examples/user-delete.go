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
	client, err := pritunl.NewClient( /* provide credentials here if environment variables is not present */ )
	if err != nil {
		log.Fatal(err)
	}

	// Create a context for the request
	ctx := context.Background()

	// Create a new UserRequest object with desired data
	deleteUser := &pritunl.UserRequest{
		Name:  "new_user",
		Email: "updateuser@domain.dev",
		// Set Disabled to false (default behavior) or any other desired value
		Disabled: false, // Or true if you want the user to be disabled
	}

	// Call UserUpdate to update existing user for organization `641351fee8f281432b807a50`
	users, err := client.UserDelete(ctx, "641351fee8f281432b807a50", "6618b060a7013fe771cae4ef", *deleteUser)
	if err != nil {
		log.Fatal(err)
	}

	/* PRESENTATION */
	if err != nil {
		// Handle error
		fmt.Println("Error updating user:", err)
	} else {
		fmt.Println("Successfully updating users:", users)
	}

}
