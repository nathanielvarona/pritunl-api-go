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
		organization string = os.Getenv("PRITUNL_DATA_ORGANIZATION")
		server       string = os.Getenv("PRITUNL_DATA_SERVER")
	)

	// Create a context for the request
	ctx := context.Background()

	// Attach an Organization to a Server
	serverorgs, err := client.ServerOrgDetach(ctx, server, organization)
	if err != nil {
		log.Fatal(err)
	}

	// Struct Output
	for _, serverorg := range serverorgs {
		fmt.Println("Server Org ID:", serverorg.ID)
		fmt.Println("Server Org Server", serverorg.Server)
		fmt.Println("------")
	}
}
