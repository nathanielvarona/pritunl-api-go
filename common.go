package pritunl

import (
	"fmt"
	"io"
	"net/http"
)

func handleResponse(response *http.Response) (io.ReadCloser, error) {
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	return response.Body, nil
}
