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
		org  string = os.Getenv("PRITUNL_DATA_ORG")
		user string = os.Getenv("PRITUNL_DATA_USER")
	)

	// Create a context for the request
	ctx := context.Background()

	// Genreate or Retrieve Key
	keys, err := client.KeyGet(ctx, org, user)
	if err != nil {
		log.Fatal(err)
	}

	// Struct Output
	for _, key := range keys {
		fmt.Println("Key View URL:", key.ViewURL)
		fmt.Println("Key URI URI:", key.URIURL)
		fmt.Println("Key URL:", key.KeyURL)
		fmt.Println("------")
	}
}
