package pritunl

import (
	"context"
	"fmt"
	"net/http"
)

// KeyResponse represents a key response from the Pritunl API
type KeyResponse struct {
	ID        string `json:"id"`
	KeyURL    string `json:"key_url"`
	KeyZipURL string `json:"key_zip_url"`
	KeyOncURL string `json:"key_onc_url"`
	ViewURL   string `json:"view_url"`
	URIURL    string `json:"uri_url"`
}

// KeyGet retrieves a key or keys on the server
func (c *Client) KeyGet(ctx context.Context, orgId string, userId string) ([]KeyResponse, error) {
	// Construct the API path using the organization ID and user ID
	path := fmt.Sprintf("/key/%s/%s", orgId, userId)

	// Send an authenticated GET request to retrieve the key(s)
	response, err := c.AuthRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	// Get the response body and handle any errors
	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data into a slice of KeyResponse structs
	var keys []KeyResponse
	if err := handleUnmarshal(body, &keys); err != nil {
		return nil, err
	}

	// Return the slice of keys
	return keys, nil
}
