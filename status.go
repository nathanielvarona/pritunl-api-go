package pritunl

import (
	"context"
	"net/http"
)

// StatusResponse represents the structure of Pritunl's status response
type StatusResponse struct {
	OrgCount      int      `json:"org_count"`
	UsersOnline   int      `json:"users_online"`
	UserCount     int      `json:"user_count"`
	ServersOnline int      `json:"servers_online"`
	ServerCount   int      `json:"server_count"`
	HostsOnline   int      `json:"hosts_online"`
	HostCount     int      `json:"host_count"`
	ServerVersion string   `json:"server_version"`
	CurrentHost   string   `json:"current_host"`
	PublicIP      string   `json:"public_ip"`
	LocalNetworks []string `json:"local_networks"`
	Notification  string   `json:"notification"`
}

// StatusGet retrieves the Pritunl server status
func (c *Client) StatusGet(ctx context.Context) ([]StatusResponse, error) {
	// Construct the API path
	path := "/status"

	// Send an authenticated GET request to retrieve the status
	response, err := c.AuthRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	// Get the response body and handle any errors
	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close() // Close the response body when done

	// Unmarshal the JSON data into a slice of StatusResponse
	var status []StatusResponse
	if err := handleUnmarshal(body, &status); err != nil {
		return nil, err
	}

	// Return the slice of status
	return status, nil
}
