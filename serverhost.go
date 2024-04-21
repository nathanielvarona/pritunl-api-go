package pritunl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ServerHostRequest represents a request to attach or detach a host to a server
type ServerHostRequest struct {
	ID      string `json:"id"`
	Server  string `json:"server"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// ServerHostResponse represents a server host response
type ServerHostResponse struct {
	ID      string `json:"id"`
	Server  string `json:"server"`
	Status  string `json:"status"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// ServerHostAttach attaches a host to a server
func (c *Client) ServerHostAttach(ctx context.Context, srvId string, hostId string, newServerHost ServerHostRequest) ([]ServerHostResponse, error) {
	// Marshal the ServerHostRequest into JSON data
	serverHostData, err := json.Marshal(newServerHost)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal serverhost data: %w", err)
	}

	// Construct the API path for attaching a host
	path := fmt.Sprintf("/server/%s/host/%s", srvId, hostId)

	// Send an authenticated PUT request to attach the host
	response, err := c.AuthRequest(ctx, http.MethodPut, path, serverHostData)
	if err != nil {
		return nil, err
	}

	// Handle the response and unmarshal the JSON data
	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	var serverhosts []ServerHostResponse
	if err := handleUnmarshal(body, &serverhosts); err != nil {
		return nil, err
	}

	// Return the slice of serverhosts
	return serverhosts, nil
}

// ServerHostDetach detaches a host from a server
func (c *Client) ServerHostDetach(ctx context.Context, srvId string, hostId string) ([]ServerHostResponse, error) {
	var serverHostData []byte

	// Construct the API path for detaching a host
	path := fmt.Sprintf("/server/%s/host/%s", srvId, hostId)

	// Send an authenticated DELETE request to detach the host
	response, err := c.AuthRequest(ctx, http.MethodDelete, path, serverHostData)
	if err != nil {
		return nil, err
	}

	// Handle the response and unmarshal the JSON data
	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	var serverhosts []ServerHostResponse
	if err := handleUnmarshal(body, &serverhosts); err != nil {
		return nil, err
	}

	// Return the slice of serverhosts
	return serverhosts, nil
}
