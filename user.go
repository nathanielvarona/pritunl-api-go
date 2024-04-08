package pritunl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// UserGet retrieves a user or a list of users from the server
func (c *Client) UserGet(ctx context.Context, orgId string, userId ...string) ([]map[string]interface{}, error) {
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

	// Decode the response into a slice of maps to handle both single and multiple users
	var users []map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &users); err != nil {
		return nil, fmt.Errorf("failed to decode user response: %w", err)
	}

	return users, nil
}
