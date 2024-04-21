package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nathanielvarona/pritunl-api-go"
)

// UserResponse represents a user response from the Pritunl API
type UserResponse struct {
	Name         string `json:"name"`
	ID           string `json:"id"`
	Organization string `json:"organization"`
}

func main() {
	// Initialize the Pritunl API client
	client, err := pritunl.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve organization and user IDs from environment variables
	organization := os.Getenv("PRITUNL_DATA_ORG")
	user := os.Getenv("PRITUNL_DATA_USER")

	// Create a context for the request
	ctx := context.Background()

	// Retrieve a specific user under the organization
	user_org1, err := client.UserGet(ctx, organization, user)
	if err != nil {
		log.Fatal(err)
	}

	// Print user details
	fmt.Println("Specific User:")
	for _, user := range user_org1 {
		fmt.Println("User Name:", user.Name)
		fmt.Println("User ID:", user.ID)
		fmt.Println("Organization ID:", user.Organization)
		fmt.Println("------")
	}

	fmt.Println("####")

	// Retrieve all users under the organization
	users_org2, err := client.UserGet(ctx, organization)
	if err != nil {
		log.Fatal(err)
	}

	// Print all user details
	fmt.Println("All Users:")
	for _, user := range users_org2 {
		fmt.Println("User Name:", user.Name)
		fmt.Println("User ID:", user.ID)
		fmt.Println("Organization ID:", user.Organization)
		fmt.Println("------")
	}

	// Marshal users to JSON
	userBytes, err := json.MarshalIndent(users_org2, "", "  ")
	if err != nil {
		log.Println("Error marshalling user:", err)
	} else {
		fmt.Println(string(userBytes))
	}
}
