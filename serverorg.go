package pritunl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ServerOrgRequest represents a request to attach or detach an organization to a server
type ServerOrgRequest struct {
	ID     string `json:"id"`
	Server string `json:"server"`
	Name   string `json:"name"`
}

// ServerOrgResponse represents a server organization response
type ServerOrgResponse struct {
	ID     string `json:"id"`
	Server string `json:"server"`
	Name   string `json:"name"`
}

// ServerOrgAttach attaches an organization to a server
func (c *Client) ServerOrgAttach(ctx context.Context, srvId string, orgId string, newServerOrg ServerOrgRequest) ([]ServerOrgResponse, error) {
	// Marshal the ServerOrgRequest into JSON data
	serverOrgData, err := json.Marshal(newServerOrg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal serverorg data: %w", err)
	}

	// Construct the API path for attaching an organization
	path := fmt.Sprintf("/server/%s/organization/%s", srvId, orgId)

	// Send an authenticated PUT request to attach the organization
	response, err := c.AuthRequest(ctx, http.MethodPut, path, serverOrgData)
	if err != nil {
		return nil, err
	}

	// Handle the response and unmarshal the JSON data
	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	var serverorgs []ServerOrgResponse
	if err := handleUnmarshal(body, &serverorgs); err != nil {
		return nil, err
	}

	// Return the slice of serverorgs
	return serverorgs, nil
}

// ServerOrgDetach detaches an organization from a server
func (c *Client) ServerOrgDetach(ctx context.Context, srvId string, orgId string) ([]ServerOrgResponse, error) {
	var serverOrgData []byte

	// Construct the API path for detaching an organization
	path := fmt.Sprintf("/server/%s/organization/%s", srvId, orgId)

	// Send an authenticated DELETE request to detach the organization
	response, err := c.AuthRequest(ctx, http.MethodDelete, path, serverOrgData)
	if err != nil {
		return nil, err
	}

	// Handle the response and unmarshal the JSON data
	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	var serverorgs []ServerOrgResponse
	if err := handleUnmarshal(body, &serverorgs); err != nil {
		return nil, err
	}

	// Return the slice of serverorgs
	return serverorgs, nil
}
