# pritunl-api-go

Pritunl API Client for Go

A [Go](https://go.dev/) client for the Pritunl API, allowing you to interact with [Pritunl](https://pritunl.com/) servers and perform various actions.

## Getting Started

### Environment Variables

Load your Pritunl API credentials as environment variables:

```bash
export PRITUNL_BASE_URL="https://vpn.domain.tld"
export PRITUNL_API_TOKEN="<PRITUNL API TOKEN>"
export PRITUNL_API_SECRET="<PRITUNL API SECRET>"
```

### Installation

Get the Pritunl API Client for Go package/library:

```bash
go get github.com/nathanielvarona/pritunl-api-go
```

### Usage

Initialize an API instance and call available feature functions:

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
	// Initialize the Pritunl API client
	client, err := pritunl.NewClient()
	// Alternatively, you can initialize the client with manual arguments
	// client, err := pritunl.NewClient(&pritunl.Client{
	// 	BaseUrl:   "<PRITUNL BASE URL>",
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

	// Print server status details
	fmt.Println("Server Status:")
	for _, stat := range status {
		fmt.Println("Server Version:", stat.ServerVersion)
		fmt.Println("Local Networks:", stat.LocalNetworks)
		fmt.Println("Host Online:", stat.HostsOnline)
		fmt.Println("------")
	}

	// Marshal server status to JSON
	statusBytes, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		log.Println("Error marshalling status:", err)
	} else {
		fmt.Println("Server Status in JSON:")
		fmt.Println(string(statusBytes))
	}
}

```

### Examples
Check the [_examples](./_examples) folder for code examples demonstrating how to use this package/library.

## Features

### Core Pritunl API Client

| Feature Function   | Description                             | Status                 |
|--------------------|-----------------------------------------|------------------------|
| StatusGet          | Status of Pritunl Server                | :white_check_mark: Yes |
| KeyGet             | Generate or Retrieve a Key for the User | :white_check_mark: Yes |
| UserGet            | Get the Information of Existing User    | :white_check_mark: Yes |
| UserCreate         | Create a New User                       | :white_check_mark: Yes |
| UserUpdate         | Update an Existing User                 | :white_check_mark: Yes |
| UserDelete         | Delete an User                          | :white_check_mark: Yes |
| OrganizationGet    | Get the Information of Existing Org     | :white_check_mark: Yes |
| OrganizationCreate | Create a New Org                        | :white_check_mark: Yes |
| OrganizationUpdate | Update an Existing Org                  | :white_check_mark: Yes |
| OrganizationDelete | Delete an Org                           | :white_check_mark: Yes |
| ServerGet          | Get the Information of Existing Server  | :white_check_mark: Yes |
| ServerCreate       | Create a New Server                     | :white_check_mark: Yes |
| ServerUpdate       | Update an existing Server               | :white_check_mark: Yes |
| ServerDelete       | Delete a Server                         | :white_check_mark: Yes |
| ServerRouteGet     | Get the Routes for a Server             | :white_check_mark: Yes |
| ServerRouteCreate  | Create/Add a Server Route               | :white_check_mark: Yes |
| ServerRouteUpdate  | Update a Server Route                   | :white_check_mark: Yes |
| ServerRouteDelete  | Remove/Delete a Server Route            | :white_check_mark: Yes |
| ServerOrgAttach    | Attach an Organization for a Server     | :white_check_mark: Yes |
| ServerOrgDetach    | Detach an Organization for a Server     | :white_check_mark: Yes |
| ServerHostAttach   | Attach a Host for a Server              | :white_check_mark: Yes |
| ServerHostDetach   | Detach a Host for a Server              | :white_check_mark: Yes |

### Future Enhancements (CLI)

1. **CLI Framework:** Consider using a popular framework like `spf13/cobra` (https://cobra.dev/), `urfave/cli` (https://cli.urfave.org/) or `alecthomas/kong` (https://github.com/alecthomas/kong) to simplify the command structure, argument parsing, and flag handling.
2. **Build Distribution Workflow:** Implement a CI/CD workflow (e.g., using GitHub Actions) to automate building and distributing the CLI tool across various platforms (Windows, macOS, Linux) and architectures (32-bit, 64-bit). This will streamline setup for users on different systems.

## Alternative API Clients
* Python - [Pritunl API Client for Python](https://github.com/nathanielvarona/pritunl-api-python) by [@nathanielvarona](https://github.com/nathanielvarona)
  - _(fork from [Pritunl API client for Python 3](https://github.com/ijat/pritunl-api-python) by [@ijat](https://github.com/ijat))_.
* Ruby - [Pritunl API Client](https://github.com/eterry1388/pritunl_api_client) by [@eterry1388](https://github.com/eterry1388)
