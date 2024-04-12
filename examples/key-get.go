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

	// Genreate or Retrieve Key
	keys, err := client.KeyGet(ctx, "641351fee8f281432b807a50", "641351ffe8f281432b807a6e")
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
