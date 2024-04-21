package pritunl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// handleResponse checks the HTTP response status code and returns the response body
func handleResponse(response *http.Response) (io.ReadCloser, error) {
	if response.StatusCode != http.StatusOK {
		// If the status code is not 200 OK, return an error
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	// Return the response body and no error
	return response.Body, nil
}

// handleUnmarshal reads the response body, checks if it's a single object, and unmarshals it into the result interface
func handleUnmarshal(body io.Reader, result interface{}) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		// If there's an error reading the response body, return it
		return fmt.Errorf("failed to read response body: %w", err)
	}
	// Check if the response is a single object ( `{}` )
	if bytes.HasPrefix(bodyBytes, []byte("{")) && bytes.HasSuffix(bodyBytes, []byte("}")) {
		// Convert the single object to a slice ( `[{}]` )
		bodyBytes = []byte("[" + string(bodyBytes) + "]")
	}
	// Attempt to unmarshal the entire response into the result interface
	if err := json.Unmarshal(bodyBytes, result); err != nil {
		// If there's an error unmarshaling, return it
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}
	// If everything is successful, return no error
	return nil
}
