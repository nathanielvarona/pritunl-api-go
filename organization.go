package pritunl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OrganizationRequest struct {
	Name       string `json:"name"`
	AuthApi    bool   `json:"auth_api"`
	AuthToken  bool   `json:"auth_token"`  // Addition for Put Method
	AuthSecret bool   `json:"auth_secret"` // Addition for Put Method
}

type OrganizationResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	AuthApi    bool   `json:"auth_api"`
	AuthToken  bool   `json:"auth_token"`
	AuthSecret bool   `json:"auth_secret"`
	UserCount  int    `json:"user_count"`
}

func handleUnmarshalOrganizations(body io.Reader, organizations *[]OrganizationResponse) error {
	return handleUnmarshal(body, organizations)
}

// OrganizationGet retrieves a organization or organizations on the server
func (c *Client) OrganizationGet(ctx context.Context, orgId ...string) ([]OrganizationResponse, error) {
	// The API path for the organization
	path := "/organization"

	// Handle optional orgId argument
	if len(orgId) > 0 {
		path = fmt.Sprintf("%s/%s", path, orgId[0]) // Use the first element if orgId is provided
	}

	response, err := c.AuthRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data using the helper function
	var organizations []OrganizationResponse
	if err := handleUnmarshalOrganizations(body, &organizations); err != nil {
		return nil, err
	}

	// Return the slice of organizations
	return organizations, nil
}

// OrganizationCreate create a new organization on the server
func (c *Client) OrganizationCreate(ctx context.Context, newOrganization OrganizationRequest) ([]OrganizationResponse, error) {
	// Marshal the OrganizationRequest struct into JSON data
	orgData, err := json.Marshal(newOrganization)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal organization data: %w", err)
	}

	// The API path for the organization
	path := "/organization"

	response, err := c.AuthRequest(ctx, http.MethodPost, path, orgData)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data using the helper function
	var organizations []OrganizationResponse
	if err := handleUnmarshalOrganizations(body, &organizations); err != nil {
		return nil, err
	}

	// Return the slice of organizations
	return organizations, nil
}

// OrganizationUpdate update an existing organization on the server
func (c *Client) OrganizationUpdate(ctx context.Context, orgId string, updateOrganization OrganizationRequest) ([]OrganizationResponse, error) {
	// Marshal the OrganizationRequest struct into JSON data
	orgData, err := json.Marshal(updateOrganization)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal organization data: %w", err)
	}

	// Construct the API path for the organization
	path := fmt.Sprintf("/organization/%s", orgId)

	response, err := c.AuthRequest(ctx, http.MethodPut, path, orgData)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data using the helper function
	var organizations []OrganizationResponse
	if err := handleUnmarshalOrganizations(body, &organizations); err != nil {
		return nil, err
	}

	// Return the slice of organizations
	return organizations, nil
}

// OrganizationDelete delete an existing organization on the server
func (c *Client) OrganizationDelete(ctx context.Context, orgId string) ([]OrganizationResponse, error) {
	// The API path for the organization
	path := fmt.Sprintf("/organization/%s", orgId)

	response, err := c.AuthRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data using the helper function
	var organizations []OrganizationResponse
	if err := handleUnmarshalOrganizations(body, &organizations); err != nil {
		return nil, err
	}

	// Return the slice of organizations
	return organizations, nil
}
