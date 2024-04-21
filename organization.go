package pritunl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// OrganizationRequest represents the structure of the organization request
type OrganizationRequest struct {
	Name       string `json:"name"`
	AuthApi    bool   `json:"auth_api"`
	AuthToken  bool   `json:"auth_token"`  // Addition for Put Method
	AuthSecret bool   `json:"auth_secret"` // Addition for Put Method
}

// OrganizationResponse represents the structure of the organization response
type OrganizationResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	AuthApi    bool   `json:"auth_api"`
	AuthToken  bool   `json:"auth_token"`
	AuthSecret bool   `json:"auth_secret"`
	UserCount  int    `json:"user_count"`
}

// OrganizationGet retrieves a organization or organizations on the server
func (c *Client) OrganizationGet(ctx context.Context, orgId ...string) ([]OrganizationResponse, error) {
	// The API path for the organization
	path := "/organization"

	// Handle optional orgId argument
	if len(orgId) > 0 {
		path = fmt.Sprintf("%s/%s", path, orgId[0]) // Use the first element if orgId is provided
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
	defer body.Close()

	// Unmarshal the JSON data into a slice of OrganizationResponse
	var organizations []OrganizationResponse
	if err := handleUnmarshal(body, &organizations); err != nil {
		return nil, err
	}

	// Return the slice of organizations
	return organizations, nil
}

// OrganizationCreate creates a new organization on the server
func (c *Client) OrganizationCreate(ctx context.Context, newOrganization OrganizationRequest) ([]OrganizationResponse, error) {
	// Marshal the OrganizationRequest struct into JSON data
	orgData, err := json.Marshal(newOrganization)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal organization data: %w", err)
	}

	// The API path for the organization
	path := "/organization"

	// Send an authenticated HTTP POST request to the API with the organization data
	response, err := c.AuthRequest(ctx, http.MethodPost, path, orgData)
	if err != nil {
		return nil, err
	}

	// Handle the HTTP response
	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data into a slice of OrganizationResponse
	var organizations []OrganizationResponse
	if err := handleUnmarshal(body, &organizations); err != nil {
		return nil, err
	}

	// Return the slice of organizations
	return organizations, nil
}

// OrganizationUpdate updates an existing organization on the server
func (c *Client) OrganizationUpdate(ctx context.Context, orgId string, updateOrganization OrganizationRequest) ([]OrganizationResponse, error) {
	// Marshal the OrganizationRequest struct into JSON data
	orgData, err := json.Marshal(updateOrganization)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal organization data: %w", err)
	}

	// Construct the API path for the organization
	path := fmt.Sprintf("/organization/%s", orgId)

	// Send an authenticated HTTP PUT request to the API with the organization data
	response, err := c.AuthRequest(ctx, http.MethodPut, path, orgData)
	if err != nil {
		return nil, err
	}

	// Handle the HTTP response
	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data into a slice of OrganizationResponse
	var organizations []OrganizationResponse
	if err := handleUnmarshal(body, &organizations); err != nil {
		return nil, err
	}

	// Return the slice of organizations
	return organizations, nil
}

// OrganizationDelete deletes an existing organization on the server
func (c *Client) OrganizationDelete(ctx context.Context, orgId string) ([]OrganizationResponse, error) {
	// The API path for the organization
	path := fmt.Sprintf("/organization/%s", orgId)

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

	// Unmarshal the JSON data using the helper function
	var organizations []OrganizationResponse
	if err := handleUnmarshal(body, &organizations); err != nil {
		return nil, err
	}

	// Return the slice of organizations
	return organizations, nil
}
