package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nathanielvarona/pritunl-api-go"
)

// Include UserResponse struct definition here or import from its file

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

	// UserGet retrieve specific user under the organization
	user_org1, err := client.UserGet(ctx, organization, user) // Only provide organization ID and desired user ID
	if err != nil {
		log.Fatal(err)
	}

	// Struct Output
	for _, user := range user_org1 {
		fmt.Println("User Name:", user.Name)
		fmt.Println("User ID:", user.ID)
		fmt.Println("Organization ID:", user.Organization)
		fmt.Println("------")
	}

	fmt.Println("####")

	// UserGet retreive all users under the organization
	users_org2, err := client.UserGet(ctx, organization)
	if err != nil {
		log.Fatal(err)
	}

	// Struct Output
	for _, user := range users_org2 {
		fmt.Println("User Name:", user.Name)
		fmt.Println("User ID:", user.ID)
		fmt.Println("Organization ID:", user.Organization)
		fmt.Println("------")
	}

	// JSON Output
	userBytes, err := json.MarshalIndent(users_org2, "", "  ")
	if err != nil {
		log.Println("Error marshalling user:", err)
	} else {
		fmt.Println(string(userBytes))
	}
}
