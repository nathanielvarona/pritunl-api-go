package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nathanielvarona/pritunl-api-go"
)

func main() {
	/* INITIALIZATION */
	// Provide authentication credentials as needed for client creation
	client, err := pritunl.NewClient( /* provide credentials here if environment variables is not present */ )

	// You can also initialize an instance by manually providing the arguments.
	//
	// client := pritunl.NewClient(&pritunl.Client{
	// 	BaseUrl:   "<PRITUNL API URL>",
	// 	ApiToken:  "<PRITUNL API TOKEN>",
	// 	ApiSecret: "<PRITUNL API SECRET>",
	// })

	if err != nil {
		log.Fatal(err)
	}

	// Create a context for the request
	ctx := context.Background()

	// Retrieve the server status
	status, err := client.StatusGet(ctx)
	if err != nil {
		log.Fatal(err)
	}

	/* PRESENTATION */
	// Struct Output
	for _, stat := range status {
		fmt.Println("Server Version", stat.ServerVersion)
		fmt.Println("Loacal Networks:", stat.LocalNetworks)
	}

	// JSON Output
	statusBytes, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		log.Println("Error marshalling status:", err)
	} else {
		fmt.Println(string(statusBytes))
	}

}
