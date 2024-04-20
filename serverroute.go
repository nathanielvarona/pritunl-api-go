package pritunl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ServerRouteRequest struct {
	ID                 string      `json:"id"`
	Server             string      `json:"server"`
	Network            string      `json:"network"`
	Comment            string      `json:"comment"`
	Metric             interface{} `json:"metric"`
	Nat                bool        `json:"nat"`
	NatInterface       string      `json:"nat_interface"`
	NatNetmap          string      `json:"nat_netmap"`
	Advertise          bool        `json:"advertise"`
	VpcRegion          interface{} `json:"vpc_region"`
	VpcID              interface{} `json:"vpc_id"`
	NetGateway         bool        `json:"net_gateway"`
	VirtualNetwork     bool        `json:"virtual_network"`
	NetworkLink        bool        `json:"network_link"`
	ServerLink         bool        `json:"server_link"` // Addition for Server Route Update
	LinkVirtualNetwork bool        `json:"link_virtual_network"`
	WgNetwork          interface{} `json:"wg_network"`
}

type ServerRouteResponse struct {
	ID           string      `json:"id"`
	Server       string      `json:"server"`
	Network      string      `json:"network"`
	Comment      interface{} `json:"comment"`
	Metric       interface{} `json:"metric"`
	Nat          bool        `json:"nat"`
	NatInterface interface{} `json:"nat_interface"`
	NatNetmap    interface{} `json:"nat_netmap"`
	Advertise    bool        `json:"advertise"`
	VpcRegion    interface{} `json:"vpc_region"`
	VpcID        interface{} `json:"vpc_id"`
	NetGateway   bool        `json:"net_gateway"`
}

func handleUnmarshalServerRoutes(body io.Reader, serverroutes *[]ServerRouteResponse) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	// Attempt to unmarshal the entire response into a slice of ServerRouteResponse
	if err := json.Unmarshal(bodyBytes, serverroutes); err != nil {
		// If unmarshalling as a list fails, try unmarshalling as a single ServerRouteResponse
		var singleServerRoute ServerRouteResponse
		if unmarshalErr := json.Unmarshal(bodyBytes, &singleServerRoute); unmarshalErr == nil {
			*serverroutes = append(*serverroutes, singleServerRoute) // Add the single server route to the slice
		} else {
			return fmt.Errorf("failed to unmarshal server response: %w", err) // Return original error
		}
	}
	return nil
}

// ServerRouteGet retrieves the server routes
func (c *Client) ServerRouteGet(ctx context.Context, srvId string) ([]ServerRouteResponse, error) {
	var serverRouteData []byte
	path := fmt.Sprintf("/server/%s/routes", srvId)

	response, err := c.AuthRequest(ctx, http.MethodGet, path, serverRouteData)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data using the helper function
	var serverroutes []ServerRouteResponse
	if err := handleUnmarshalServerRoutes(body, &serverroutes); err != nil {
		return nil, err
	}

	// Return the slice of serverroutes
	return serverroutes, nil
}

// ServerRouteCreate add a route to a network
func (c *Client) ServerRouteCreate(ctx context.Context, srvId string, newServerRoute ServerRouteRequest) ([]ServerRouteResponse, error) {
	serverRouteData, err := json.Marshal(newServerRoute)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal serverroute data: %w", err)
	}

	path := fmt.Sprintf("/server/%s/route", srvId)

	response, err := c.AuthRequest(ctx, http.MethodPost, path, serverRouteData)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data using the helper function
	var serverroutes []ServerRouteResponse
	if err := handleUnmarshalServerRoutes(body, &serverroutes); err != nil {
		return nil, err
	}

	// Return the slice of serverroutes
	return serverroutes, nil
}

// ServerRouteUpdate update a server route
func (c *Client) ServerRouteUpdate(ctx context.Context, srvId string, routeId string, newServerRoute ServerRouteRequest) ([]ServerRouteResponse, error) {
	serverRouteData, err := json.Marshal(newServerRoute)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal serverroute data: %w", err)
	}

	path := fmt.Sprintf("/server/%s/route/%s", srvId, routeId)

	response, err := c.AuthRequest(ctx, http.MethodPut, path, serverRouteData)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data using the helper function
	var serverroutes []ServerRouteResponse
	if err := handleUnmarshalServerRoutes(body, &serverroutes); err != nil {
		return nil, err
	}

	// Return the slice of serverroutes
	return serverroutes, nil
}
