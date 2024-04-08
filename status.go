package pritunl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// StatusGet retrieves the Pritunl server status and returns the decoded data
func (c *Client) StatusGet(ctx context.Context) (map[string]interface{}, error) {
	var data []byte

	response, err := c.AuthRequest(ctx, http.MethodGet, "/status", data)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	// Read entire body into a byte slice
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode status response: %w", err)
	}

	// Decode the read bytes into a map
	var status map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &status); err != nil {
		return nil, fmt.Errorf("failed to decode status response: %w", err)
	}

	return status, nil
}
