package pritunl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ServerOrgRequest struct {
	ID     string      `json:"id"`
	Server string      `json:"server"`
	Name   interface{} `json:"name"`
}

type ServerOrgResponse struct {
	ID     string `json:"id"`
	Server string `json:"server"`
	Name   string `json:"name"`
}

func handleUnmarshalServerOrgs(body io.Reader, serverorgs *[]ServerOrgResponse) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	// Attempt to unmarshal the entire response into a slice of ServerOrgResponse
	if err := json.Unmarshal(bodyBytes, serverorgs); err != nil {
		// If unmarshalling as a list fails, try unmarshalling as a single ServerOrgResponse
		var singleServerOrg ServerOrgResponse
		if unmarshalErr := json.Unmarshal(bodyBytes, &singleServerOrg); unmarshalErr == nil {
			*serverorgs = append(*serverorgs, singleServerOrg) // Add the single server route to the slice
		} else {
			return fmt.Errorf("failed to unmarshal server response: %w", err) // Return original error
		}
	}
	return nil
}

// ServerOrgAttach attach an organization to a server
func (c *Client) ServerOrgAttach(ctx context.Context, srvId string, orgId string, newServerOrg ServerOrgRequest) ([]ServerOrgResponse, error) {
	serverOrgData, err := json.Marshal(newServerOrg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal serverorg data: %w", err)
	}

	path := fmt.Sprintf("/server/%s/organization/%s", srvId, orgId)

	response, err := c.AuthRequest(ctx, http.MethodPut, path, serverOrgData)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data using the helper function
	var serverorgs []ServerOrgResponse
	if err := handleUnmarshalServerOrgs(body, &serverorgs); err != nil {
		return nil, err
	}

	// Return the slice of serverorgs
	return serverorgs, nil
}
