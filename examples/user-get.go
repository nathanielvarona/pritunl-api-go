// example/user.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nathanielvarona/pritunl-api-go"
)

// Include UserResponse struct definition here or import from its file

func main() {
	/* INITIALIZATION AND FETCHING */

	// Provide authentication credentials as needed for client creation
	client, err := pritunl.NewClient( /* provide credentials here if environment variables is not present */ )
	if err != nil {
		log.Fatal(err)
	}

	// Create a context for the request
	ctx := context.Background()

	// Retrieve all users for organization "641351fee8f281432b807a50"
	users, err := client.User(ctx, "641351fee8f281432b807a50")
	if err != nil {
		log.Fatal(err)
	}

	/* PRESENTATION */

	// Per Object Representation
	for _, user := range users {
		fmt.Println("User ID:", user.ID)
		fmt.Println("User Name:", user.Name)

	}

	// Retrieve specific user "644b2ba8cc3f875be1b7658d" under the organization "64131e880654550010d30c54"
	user, err := client.User(ctx, "64131e880654550010d30c54", "644b2ba8cc3f875be1b7658d") // Only provide organization ID and desired user ID
	if err != nil {
		log.Fatal(err)
	}

	// Per Object Representation
	for _, user := range user {
		fmt.Println("User ID:", user.ID)
		fmt.Println("User Name:", user.Name)
	}

	// JSON Representation
	userBytes, err := json.MarshalIndent(users, "", "  ")
	// userBytes, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Println("Error marshalling user:", err)
	} else {
		fmt.Println(string(userBytes))
	}

}
