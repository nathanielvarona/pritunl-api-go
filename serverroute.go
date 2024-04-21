package pritunl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ServerRouteRequest represents a request to create or update a server route
type ServerRouteRequest struct {
	ID                 string `json:"id"`
	Server             string `json:"server"`
	Network            string `json:"network"`
	Comment            string `json:"comment"`
	Metric             int    `json:"metric"`
	Nat                bool   `json:"nat"`
	NatInterface       string `json:"nat_interface"`
	NatNetmap          string `json:"nat_netmap"`
	Advertise          bool   `json:"advertise"`
	VpcRegion          string `json:"vpc_region"`
	VpcId              string `json:"vpc_id"`
	NetGateway         bool   `json:"net_gateway"`
	VirtualNetwork     bool   `json:"virtual_network"`
	NetworkLink        bool   `json:"network_link"`
	ServerLink         bool   `json:"server_link"` // Addition for Put Method
	LinkVirtualNetwork bool   `json:"link_virtual_network"`
	WgNetwork          bool   `json:"wg_network"`
}

// ServerRouteResponse represents a server route response
type ServerRouteResponse struct {
	ID           string `json:"id"`
	Server       string `json:"server"`
	Network      string `json:"network"`
	Comment      string `json:"comment"`
	Metric       int    `json:"metric"`
	Nat          bool   `json:"nat"`
	NatInterface string `json:"nat_interface"`
	NatNetmap    string `json:"nat_netmap"`
	Advertise    bool   `json:"advertise"`
	VpcRegion    string `json:"vpc_region"`
	VpcId        string `json:"vpc_id"`
	NetGateway   bool   `json:"net_gateway"`
}

// ServerRouteGet retrieves the server routes
func (c *Client) ServerRouteGet(ctx context.Context, srvId string) ([]ServerRouteResponse, error) {
	// Construct the API path using the server ID
	path := fmt.Sprintf("/server/%s/route", srvId)

	// Send an authenticated GET request to retrieve server routes
	response, err := c.AuthRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data into a slice of ServerRouteResponse
	var serverroutes []ServerRouteResponse
	if err := handleUnmarshal(body, &serverroutes); err != nil {
		return nil, err
	}

	// Return the slice of serverroutes
	return serverroutes, nil
}

// ServerRouteCreate adds a route to a network
func (c *Client) ServerRouteCreate(ctx context.Context, srvId string, newServerRoute ServerRouteRequest) ([]ServerRouteResponse, error) {
	// Marshal the ServerRouteRequest struct into JSON data
	serverRouteData, err := json.Marshal(newServerRoute)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal serverroute data: %w", err)
	}

	// Construct the API path using the server ID
	path := fmt.Sprintf("/server/%s/route", srvId)

	// Send an authenticated POST request to create a new server route
	response, err := c.AuthRequest(ctx, http.MethodPost, path, serverRouteData)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data into a slice of ServerRouteResponse
	var serverroutes []ServerRouteResponse
	if err := handleUnmarshal(body, &serverroutes); err != nil {
		return nil, err
	}

	// Return the slice of serverroutes
	return serverroutes, nil
}

// ServerRouteUpdate updates a server route
func (c *Client) ServerRouteUpdate(ctx context.Context, srvId string, routeId string, newServerRoute ServerRouteRequest) ([]ServerRouteResponse, error) {
	// Marshal the ServerRouteRequest struct into JSON data
	serverRouteData, err := json.Marshal(newServerRoute)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal serverroute data: %w", err)
	}

	// Construct the API path using the server ID and route ID
	path := fmt.Sprintf("/server/%s/route/%s", srvId, routeId)

	// Send an authenticated PUT request to update an existing server route
	response, err := c.AuthRequest(ctx, http.MethodPut, path, serverRouteData)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data into a slice of ServerRouteResponse
	var serverroutes []ServerRouteResponse
	if err := handleUnmarshal(body, &serverroutes); err != nil {
		return nil, err
	}

	// Return the slice of serverroutes
	return serverroutes, nil
}

// ServerRouteDelete removes a server route
func (c *Client) ServerRouteDelete(ctx context.Context, srvId string, routeId string) ([]ServerRouteResponse, error) {
	// Construct the API path using the server ID and route ID
	path := fmt.Sprintf("/server/%s/route/%s", srvId, routeId)

	// Send an authenticated DELETE request to remove a server route
	response, err := c.AuthRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data into a slice of ServerRouteResponse
	var serverroutes []ServerRouteResponse
	if err := handleUnmarshal(body, &serverroutes); err != nil {
		return nil, err
	}

	// Return the slice of serverroutes
	return serverroutes, nil
}
