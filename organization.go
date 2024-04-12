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
	AuthToken  bool   `json:"auth_token"`
	AuthSecret bool   `json:"auth_secret"`
}

type OrganizationResponse struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	AuthAPI    bool    `json:"auth_api"`
	AuthToken  *string `json:"auth_token"`  // Using a pointer for optional field
	AuthSecret *string `json:"auth_secret"` // Using a pointer for optional field
	UserCount  int     `json:"user_count"`
}

func handleUnmarshalOrganizations(body io.Reader, organizations *[]OrganizationResponse) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	// Attempt to unmarshal the entire response into a slice of OrganizationResponse
	if err := json.Unmarshal(bodyBytes, organizations); err != nil {
		// If unmarshalling as a list fails, try unmarshalling as a single OrganizationResponse
		var singleOrg OrganizationResponse
		if unmarshalErr := json.Unmarshal(bodyBytes, &singleOrg); unmarshalErr == nil {
			*organizations = append(*organizations, singleOrg) // Add the single organization to the slice
		} else {
			return fmt.Errorf("failed to unmarshal organization response: %w", err) // Return original error
		}
	}
	return nil
}

// OrganizationGet retrieves a organization or organizations on the server
func (c *Client) OrganizationGet(ctx context.Context, orgId ...string) ([]OrganizationResponse, error) {
	var data []byte
	path := "/organization"

	// Handle optional orgId argument
	if len(orgId) > 0 {
		path = fmt.Sprintf("%s/%s", path, orgId[0]) // Use the first element if orgId is provided
	}

	response, err := c.AuthRequest(ctx, http.MethodGet, path, data)
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
	orgData, err := json.Marshal(newOrganization)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal organization data: %w", err)
	}

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
	orgData, err := json.Marshal(updateOrganization)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal organization data: %w", err)
	}

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
	var orgData []byte

	path := fmt.Sprintf("/organization/%s", orgId)

	response, err := c.AuthRequest(ctx, http.MethodDelete, path, orgData)
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
