package pritunl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ServerHostRequest struct {
	ID      string      `json:"id"`
	Server  string      `json:"server"`
	Name    interface{} `json:"name"`
	Address interface{} `json:"address"`
}

type ServerHostResponse struct {
	ID      string `json:"id"`
	Server  string `json:"server"`
	Status  string `json:"status"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func handleUnmarshalServerHosts(body io.Reader, serverhosts *[]ServerHostResponse) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	// Attempt to unmarshal the entire response into a slice of ServerHostResponse
	if err := json.Unmarshal(bodyBytes, serverhosts); err != nil {
		// If unmarshalling as a list fails, try unmarshalling as a single ServerHostResponse
		var singleServerHost ServerHostResponse
		if unmarshalErr := json.Unmarshal(bodyBytes, &singleServerHost); unmarshalErr == nil {
			*serverhosts = append(*serverhosts, singleServerHost) // Add the single server route to the slice
		} else {
			return fmt.Errorf("failed to unmarshal server response: %w", err) // Return original error
		}
	}
	return nil
}

// ServerHostAttach attach a host to a server
func (c *Client) ServerHostAttach(ctx context.Context, srvId string, hostId string, newServerHost ServerHostRequest) ([]ServerHostResponse, error) {
	serverHostData, err := json.Marshal(newServerHost)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal serverhost data: %w", err)
	}

	path := fmt.Sprintf("/server/%s/host/%s", srvId, hostId)

	response, err := c.AuthRequest(ctx, http.MethodPut, path, serverHostData)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data using the helper function
	var serverhosts []ServerHostResponse
	if err := handleUnmarshalServerHosts(body, &serverhosts); err != nil {
		return nil, err
	}

	// Return the slice of serverhosts
	return serverhosts, nil
}

// ServerHostDetach detach a host to a server
func (c *Client) ServerHostDetach(ctx context.Context, srvId string, hostId string) ([]ServerHostResponse, error) {
	var serverHostData []byte
	path := fmt.Sprintf("/server/%s/host/%s", srvId, hostId)

	response, err := c.AuthRequest(ctx, http.MethodDelete, path, serverHostData)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data using the helper function
	var serverhosts []ServerHostResponse
	if err := handleUnmarshalServerHosts(body, &serverhosts); err != nil {
		return nil, err
	}

	// Return the slice of serverhosts
	return serverhosts, nil
}
