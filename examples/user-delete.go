package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nathanielvarona/pritunl-api-go"
)

func main() {
	// Provide authentication credentials as needed for client creation
	// Automaticlly sets from environment variables if present
	client, err := pritunl.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Create a context for the request
	ctx := context.Background()

	// Call UserUpdate to update existing user for organization `641351fee8f281432b807a50`
	users, err := client.UserDelete(ctx, "641351fee8f281432b807a50", "6618b08ea7013fe771cae514")
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
