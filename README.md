# pritunl-api-go
Pritunl API Client for Go

> [!WARNING]
> The project is currently in development, so some features may be limited or unavailable at this time.



## API Usage

Load your Pritunl API Credentials in our Environment Variables.

```bash
export PRITUNL_BASE_URL="https://vpn.domain.tld/"
export PRITUNL_API_TOKEN="XXXXXXXXXXXXXXXXXXXXX"
export PRITUNL_API_SECRET="XXXXXXXXXXXXXXXXXXXXX"
```

Get the Pritunl API Client for Go Package/Library

```bash
go get github.com/nathanielvarona/pritunl-api-go
```



```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nathanielvarona/pritunl-api-go"
)

func main() {
	/* INITIALIZATION AND FETCHING */

	// Provide authentication credentials as needed for client creation
	client, err := pritunl.NewClient()

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
	statuses, err := client.Status(ctx)
	if err != nil {
		log.Fatal(err)
	}

	/* PRESENTATION */

	// Per Object Representation
	for _, status := range statuses {
		fmt.Println("Server Version", status.ServerVersion)
		fmt.Println("Loacal Networks:", status.LocalNetworks)
	}

	// JSON Representation
	statusBytes, err := json.MarshalIndent(statuses, "", "  ")
	if err != nil {
		log.Println("Error marshalling status:", err)
	} else {
		fmt.Println(string(statusBytes))
	}

}
```

## Features Status

Core API Client Package/Library Features

Feature      | Method | Description                             | Status
-------------|--------|-----------------------------------------|-----------------------
Server       | Get    | Status of Pritunl Server                | :white_check_mark: Yes
Key          | Get    | Generate or Retrieve a Key for the User | Not yet
User         | Get    | Get the Information of Existing User    | :white_check_mark: Yes
User         | Post   | Create a New User                       | Not yet
User         | Put    | Update an Existing User                 | Not yet
User         | Delete | Delete an User                          | Not yet
Organization | Get    | Get the Information of Existing Org     | Not yet
Organization | Post   | Create a New Org                        | Not yet
Organization | Put    | Update an Existing Org                  | Not yet
Organization | Delete | Delete an Org                           | Not yet
