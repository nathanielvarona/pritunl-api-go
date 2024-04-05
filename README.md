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
	"encoding/json"
	"fmt"
	"log"

	"github.com/nathanielvarona/pritunl-api-go"
)

func main() {

	client := pritunl.PritunlClient()

	// You can also initialize an instance by manually providing the arguments.
	//
	// client := pritunl.PritunlClient(&pritunl.Client{
	// 	BaseUrl:   "<PRITUNL API URL>",
	// 	ApiToken:  "<PRITUNL API TOKEN>",
	// 	ApiSecret: "<PRITUNL API SECRET>",
	// })

	resp, err := client.StatusGet()

	if err != nil {
		log.Fatalf("Error performing status request: %v", err)
	}

	defer resp.Body.Close()

	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		log.Fatalf("Error decoding response body: %v", err)
	}

	fmt.Println(responseBody)

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
