package pritunl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Status represents the structure of Pritunl's status response
type Status struct {
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
func (c *Client) StatusGet(ctx context.Context) (*Status, error) {
	var data []byte

	response, err := c.AuthRequest(ctx, http.MethodGet, "/status", data)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	// Read entire body into a byte slice
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Decode the read bytes
	var status Status
	if err := json.Unmarshal(bodyBytes, &status); err != nil {
		return nil, fmt.Errorf("failed to decode status response: %w", err)
	}

	return &status, nil
}
