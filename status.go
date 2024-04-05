package pritunl

import (
	"net/http"
)

func (c *Client) StatusGet() (*http.Response, error) {
	var headers = make(map[string]string)
	var data []byte

	headers = map[string]string{
		"Content-Type": "application/json",
	}

	response, error := c.AuthRequest(http.MethodGet, "/status", headers, data)

	if error != nil {
		return nil, error
	}

	return response, nil
}
