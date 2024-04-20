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
		host   string = os.Getenv("PRITUNL_DATA_HOST")
		server string = os.Getenv("PRITUNL_DATA_SERVER")
	)

	// Create a ServerRouteRequest object with desired data
	newServerHost := &pritunl.ServerHostRequest{
		Server: server,
		ID:     host,
	}

	// Create a context for the request
	ctx := context.Background()

	// Attach a Host to a Server
	serverhosts, err := client.ServerHostAttach(ctx, server, host, *newServerHost)
	if err != nil {
		log.Fatal(err)
	}

	// Struct Output
	for _, serverhost := range serverhosts {
		fmt.Println("Server Host ID:", serverhost.ID)
		fmt.Println("Server Host Server", serverhost.Server)
		fmt.Println("------")
	}
}
