package pritunl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserResponse struct {
	ID               string        `json:"id"`
	Organization     string        `json:"organization"`
	OrganizationName string        `json:"organization_name"`
	Name             string        `json:"name"`
	Email            string        `json:"email"`
	Groups           []interface{} `json:"groups"`
	LastActive       int64         `json:"last_active"`
	Pin              bool          `json:"pin"`
	Type             string        `json:"type"`
	AuthType         string        `json:"auth_type"`
	YubicoID         interface{}   `json:"yubico_id"`
	OTPSecret        string        `json:"otp_secret"`
	Disabled         bool          `json:"disabled"`
	BypassSecondary  bool          `json:"bypass_secondary"`
	ClientToClient   bool          `json:"client_to_client"`
	MacAddresses     []interface{} `json:"mac_addresses"`
	DNSServers       []interface{} `json:"dns_servers"`
	DNSSuffix        interface{}   `json:"dns_suffix"`
	PortForwarding   []interface{} `json:"port_forwarding"`
	Devices          []interface{} `json:"devices"`
	Gravatar         bool          `json:"gravatar"`
	Audit            bool          `json:"audit"`
	Status           bool          `json:"status"`
	SSO              interface{}   `json:"sso"`
	AuthModes        []interface{} `json:"auth_modes"`
	DNSMapping       interface{}   `json:"dns_mapping"`
	NetworkLinks     []interface{} `json:"network_links"`
	Servers          []Server      `json:"servers"` // Nested struct for servers
}

type Server struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Status         bool        `json:"status"`
	ServerID       string      `json:"server_id"`
	DeviceName     interface{} `json:"device_name"`
	Platform       interface{} `json:"platform"`
	RealAddress    interface{} `json:"real_address"`
	VirtAddress    string      `json:"virt_address"`
	VirtAddress6   string      `json:"virt_address6"`
	ConnectedSince interface{} `json:"connected_since"`
}

// User retrieves a user or a list of users from the server and unmarshals the response to a UserResponse struct
func (c *Client) User(ctx context.Context, orgId string, userId ...string) ([]UserResponse, error) {
	var data []byte
	path := fmt.Sprintf("/user/%s", orgId)

	// Handle optional userId argument
	if len(userId) > 0 {
		path = fmt.Sprintf("%s/%s", path, userId[0]) // Use the first element if userId is provided
	}

	response, err := c.AuthRequest(ctx, http.MethodGet, path, data)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode user response: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	// Unmarshal the JSON data into a slice of UserResponse objects
	var users []UserResponse
	if err := json.Unmarshal(bodyBytes, &users); err != nil {
		// Check for single user response (may not be wrapped in an array)
		var singleUser UserResponse
		if unmarshalErr := json.Unmarshal(bodyBytes, &singleUser); unmarshalErr == nil {
			users = append(users, singleUser)
		} else {
			return nil, fmt.Errorf("failed to unmarshal user response: %w", err)
		}
	}

	// Return the slice of users
	return users, nil
}
