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
	// Create a new Pritunl client instance
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
	status, err := client.StatusGet(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Print the status information (consider using a marshaller for better formatting)
	// fmt.Printf("Pritunl Status:\n")
	// fmt.Printf("  Organizations: %d\n", status.OrgCount)
	// fmt.Printf("  Online Users:  %d\n", status.UsersOnline)
	// fmt.Printf("  Total Users:   %d\n", status.UserCount)
	// fmt.Printf("  Online Servers: %d\n", status.ServersOnline)
	// fmt.Printf("  Total Servers:  %d\n", status.ServerCount)
	// fmt.Printf("  Server Version: %s\n", status.ServerVersion)
	// ... print other relevant fields

	// Optional: Marshal the status struct for a more structured representation
	statusBytes, err := json.MarshalIndent(status, "", "  ")
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
User         | Get    | Get the Information of Existing User    | Not yet
User         | Post   | Create a New User                       | Not yet
User         | Put    | Update an Existing User                 | Not yet
User         | Delete | Delete an User                          | Not yet
Organization | Get    | Get the Information of Existing Org     | Not yet
Organization | Post   | Create a New Org                        | Not yet
Organization | Put    | Update an Existing Org                  | Not yet
Organization | Delete | Delete an Org                           | Not yet
