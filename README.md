# pritunl-api-go

Pritunl API Client for Go

A [Go](https://go.dev/) client for the Pritunl API, allowing you to interact with [Pritunl](https://pritunl.com/) servers and perform various actions.

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=nathanielvarona_pritunl-api-go&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=nathanielvarona_pritunl-api-go)

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

## Contributing

We welcome your contributions to pritunl-api-go. This guide outlines the process for contributing effectively.

### Fork & Pull Requests

#### Workflow:

1. **Fork the repository:** Visit the [pritunl-api-go](https://github.com/nathanielvarona/pritunl-api-go) repository on GitHub and click "`Fork`". This creates your own copy.
2. **Clone your forked repository:** Use git clone to clone your forked copy to your local development environment.
3. **Create a Branch:** Use a descriptive branch name following the convention `<type>/<descriptive-name>`.
   - `<type>`: Choose from `breaking`, `feature`, `improvement`, `automation`, or `documentation`.
   - Refer to the [.github/labels.yml](./.github/labels.yml) and [.github/pr-labeler.yml](./.github/pr-labeler.yml) file for valid `<type>` options and label descriptions. (e.g., `improvement/start-stop-a-server`)
4. **Make your changes:** Implement your code modifications, ensuring they adhere to Go coding conventions [gofmt](https://go.dev/blog/gofmt) and consider adding ~~unit tests~~ <sup>ref*</sup> for new features.
5. **Commit your changes:** Stage and commit your changes with clear and concise commit messages.
6. **Push your branch and Create a Pull Request:** Push your local branch to your forked repository on GitHub and create a pull request with a detailed description of your changes.

#### Additional Tips:

* <sup>ref*</sup> Include `examples` where relevant to illustrate your changes.
* **Simplify your development workflow!** We recommend using a `Go workspace` when contributing to `pritunl-api-go`. Go workspaces provide a clean and efficient way to manage dependencies.
   - Refer to the official guide for setting up a workspace: https://go.dev/doc/tutorial/workspaces

#### Rebasing and Squashing Commits
Before submitting your pull request, we recommend rebasing your branch on top of the main branch and squashing your commits. This helps maintain a clean and linear commit history.
1. **Fetch the latest changes:** Run `git fetch origin` to fetch the latest changes from the main repository.
2. **Rebase your branch:** Run `git rebase origin/main` to rebase your branch on top of the main branch.
3. **Count your local commits:** Run `git log origin/main..HEAD --oneline | wc -l` to count the number of commits you've made.
4. **Squash your commits:** Run `git rebase -i HEAD~<N>` (replace <N> with the number of commits you counted) to interactively squash your commits. This will open an editor where you can merge your commits into a single, cohesive commit.
5. **Force-push your branch:** Run `git push -f origin <your-branch> --force`to update your forked repository with the rebased and squashed commits.

**We appreciate your contributions to the project!**

By following these guidelines, you'll help us maintain a high-quality codebase and make it easier for others to contribute. Thank you for taking the time to contribute to pritunl-api-go!

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
| ServerStart        | Start an existing Server                | :white_check_mark: Yes |
| ServerStop         | Start an existing Server                | :white_check_mark: Yes |
| ServerRouteGet     | Get the Routes for a Server             | :white_check_mark: Yes |
| ServerRouteCreate  | Create/Add a Server Route               | :white_check_mark: Yes |
| ServerRouteUpdate  | Update a Server Route                   | :white_check_mark: Yes |
| ServerRouteDelete  | Remove/Delete a Server Route            | :white_check_mark: Yes |
| ServerOrgAttach    | Attach an Organization for a Server     | :white_check_mark: Yes |
| ServerOrgDetach    | Detach an Organization for a Server     | :white_check_mark: Yes |
| ServerHostAttach   | Attach a Host for a Server              | :white_check_mark: Yes |
| ServerHostDetach   | Detach a Host for a Server              | :white_check_mark: Yes |

### Future Enhancements (CLI)

1. **CLI Framework:** Consider using a popular framework like [spf13/cobra](https://github.com/spf13/cobra), [urfave/cli](https://github.com/urfave/cli), or [alecthomas/kong](https://github.com/alecthomas/kong) to simplify the command structure, argument parsing, and flag handling.
2. **Build Distribution Workflow:** Implement a CI/CD workflow (e.g., using GitHub Actions) to automate building and distributing the CLI tool across various platforms (Windows, macOS, Linux) and architectures (32-bit, 64-bit). This will streamline setup for users on different systems.

## Alternative API Clients
* Python - [Pritunl API Client for Python](https://github.com/nathanielvarona/pritunl-api-python) by [@nathanielvarona](https://github.com/nathanielvarona)
  - _fork from [Pritunl API client for Python 3](https://github.com/ijat/pritunl-api-python) by [@ijat](https://github.com/ijat)_
* Ruby - [Pritunl API Client](https://github.com/eterry1388/pritunl_api_client) by [@eterry1388](https://github.com/eterry1388)
