package pritunl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type KeyResponse struct {
	ID        string `json:"id"`
	KeyURL    string `json:"key_url"`
	KeyZipURL string `json:"key_zip_url"`
	KeyOncURL string `json:"key_onc_url"`
	ViewURL   string `json:"view_url"`
	URIURL    string `json:"uri_url"`
}

func handleUnmarshalKey(body io.Reader, keys *[]KeyResponse) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	// Attempt to unmarshal the entire response into a slice of KeyResponse
	if err := json.Unmarshal(bodyBytes, keys); err != nil {
		// If unmarshalling as a list fails, try unmarshalling as a single KeyResponse
		var singleOrg KeyResponse
		if unmarshalErr := json.Unmarshal(bodyBytes, &singleOrg); unmarshalErr == nil {
			*keys = append(*keys, singleOrg) // Add the single key to the slice
		} else {
			return fmt.Errorf("failed to unmarshal key response: %w", err) // Return original error
		}
	}
	return nil
}

// KeyGet retrieves a key or keys on the server
func (c *Client) KeyGet(ctx context.Context, orgId string, userId string) ([]KeyResponse, error) {
	var data []byte
	path := fmt.Sprintf("/key/%s/%s", orgId, userId)

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
	var keys []KeyResponse
	if err := handleUnmarshalKey(body, &keys); err != nil {
		return nil, err
	}

	// Return the slice of keys
	return keys, nil
}
