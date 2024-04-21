package pritunl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// UserRequest represents the structure of User Get/Post/Put request
type UserRequest struct {
	Name            string                   `json:"name"`
	Email           string                   `json:"email"`
	AuthType        string                   `json:"auth_type"`
	YubicoId        string                   `json:"yubico_id"`
	Groups          []string                 `json:"groups"`
	Pin             string                   `json:"pin"`
	Disabled        bool                     `json:"disabled"`
	NetworkLinks    []string                 `json:"network_links"`
	BypassSecondary bool                     `json:"bypass_secondary"`
	ClientToClient  bool                     `json:"client_to_client"`
	MacAddresses    []string                 `json:"mac_addresses"`
	DnsServers      []string                 `json:"dns_servers"`
	DnsSuffix       string                   `json:"dns_suffix"`
	PortForwarding  []userPortForwardingData `json:"port_forwarding"`
	SendKeyEmail    bool                     `json:"send_key_email"` // Addition for Put Method
}

// UserResponse represents the structure of User response
type UserResponse struct {
	ID               string                   `json:"id"`
	Organization     string                   `json:"organization"`
	OrganizationName string                   `json:"organization_name"`
	Name             string                   `json:"name"`
	Email            string                   `json:"email"`
	Groups           []string                 `json:"groups"`
	LastActive       int64                    `json:"last_active"`
	Pin              bool                     `json:"pin"`
	Type             string                   `json:"type"`
	AuthType         string                   `json:"auth_type"`
	YubicoId         string                   `json:"yubico_id"`
	OTPSecret        string                   `json:"otp_secret"`
	Disabled         bool                     `json:"disabled"`
	BypassSecondary  bool                     `json:"bypass_secondary"`
	ClientToClient   bool                     `json:"client_to_client"`
	MacAddresses     []string                 `json:"mac_addresses"`
	DnsServers       []string                 `json:"dns_servers"`
	DnsSuffix        string                   `json:"dns_suffix"`
	PortForwarding   []userPortForwardingData `json:"port_forwarding"`
	Devices          []interface{}            `json:"devices"`
	Gravatar         bool                     `json:"gravatar"`
	Audit            bool                     `json:"audit"`
	Status           bool                     `json:"status"`
	SSO              interface{}              `json:"sso"`
	AuthModes        []interface{}            `json:"auth_modes"`
	DNSMapping       interface{}              `json:"dns_mapping"`
	NetworkLinks     []interface{}            `json:"network_links"`
	Servers          []serverData             `json:"servers"` // Nested struct for servers
}

// Substructure of `UserRequest` and `UserResponse` structs `PortForwarding` field
type userPortForwardingData struct {
	Protocol string `json:"protocol"`
	Port     string `json:"port"`
	Dport    string `json:"dport"`
}

// Substructure of `UserResponse` struct `Servers` field
type serverData struct {
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

// UserGet retrieves a user or users on the server
func (c *Client) UserGet(ctx context.Context, orgId string, userId ...string) ([]UserResponse, error) {
	// Construct the API path based on the orgId and optional userId
	path := fmt.Sprintf("/user/%s", orgId)

	// Handle optional userId argument
	if len(userId) > 0 {
		// If userId is provided, append it to the path
		path = fmt.Sprintf("%s/%s", path, userId[0])
	}

	// Send an authenticated HTTP GET request to the API
	response, err := c.AuthRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	// Handle the HTTP response
	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close() // Close the response body when done

	// Unmarshal the JSON data into a slice of UserResponse
	var users []UserResponse
	if err := handleUnmarshal(body, &users); err != nil {
		return nil, err
	}

	// Return the slice of users
	return users, nil
}

// UserCreate creates a new user on the server
func (c *Client) UserCreate(ctx context.Context, orgId string, newUser UserRequest) ([]UserResponse, error) {
	// Marshal the UserRequest struct into JSON data
	userData, err := json.Marshal(newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user data: %w", err)
	}

	// Construct the API path based on the orgId
	path := fmt.Sprintf("/user/%s", orgId)

	// Send an authenticated HTTP POST request to the API
	response, err := c.AuthRequest(ctx, http.MethodPost, path, userData)
	if err != nil {
		return nil, err
	}

	// Handle the HTTP response
	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data into a slice of UserResponse
	var users []UserResponse
	if err := handleUnmarshal(body, &users); err != nil {
		return nil, err
	}

	// Return the slice of users
	return users, nil
}

// UserUpdate updates an existing user on the server
func (c *Client) UserUpdate(ctx context.Context, orgId string, userId string, updateUser UserRequest) ([]UserResponse, error) {
	// Marshal the UserRequest struct into JSON data
	userData, err := json.Marshal(updateUser)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user data: %w", err)
	}

	// Construct the API path using the orgId and userId
	path := fmt.Sprintf("/user/%s/%s", orgId, userId)

	// Send an authenticated HTTP PUT request to the API
	response, err := c.AuthRequest(ctx, http.MethodPut, path, userData)
	if err != nil {
		return nil, err
	}

	// Handle the HTTP response
	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data into a slice of UserResponse
	var users []UserResponse
	if err := handleUnmarshal(body, &users); err != nil {
		return nil, err
	}

	// Return the slice of users
	return users, nil
}

// UserDelete deletes an existing user on the server
func (c *Client) UserDelete(ctx context.Context, orgId string, userId string) ([]UserResponse, error) {
	// Construct the API path using the organization ID and user ID
	path := fmt.Sprintf("/user/%s/%s", orgId, userId)

	// Send an authenticated HTTP DELETE request to the API
	response, err := c.AuthRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	// Handle the HTTP response
	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data into a slice of UserResponse
	var users []UserResponse
	if err := handleUnmarshal(body, &users); err != nil {
		return nil, err
	}

	return users, nil
}
