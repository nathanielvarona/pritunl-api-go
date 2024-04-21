package pritunl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ServerRequest represents a request to create or update a server
type ServerRequest struct {
	Name             string      `json:"name"`
	Network          string      `json:"network"`
	NetworkWg        string      `json:"network_wg"`
	NetworkMode      string      `json:"network_mode"`
	NetworkStart     string      `json:"network_start"`
	NetworkEnd       string      `json:"network_end"`
	RestrictRoutes   bool        `json:"restrict_routes"`
	Wg               bool        `json:"wg"`
	Ipv6             bool        `json:"ipv6"`
	Ipv6Firewall     bool        `json:"ipv6_firewall"`
	DynamicFirewall  bool        `json:"dynamic_firewall"`
	DeviceAuth       bool        `json:"device_auth"`
	BindAddress      string      `json:"bind_address"`
	Protocol         string      `json:"protocol"`
	Port             int         `json:"port"`
	PortWg           int         `json:"port_wg"`
	DhParamBits      int         `json:"dh_param_bits"`
	Groups           []string    `json:"groups"`
	MultiDevice      bool        `json:"multi_device"`
	DnsServers       []string    `json:"dns_servers"`
	SearchDomain     string      `json:"search_domain"`
	InterClient      bool        `json:"inter_client"`
	PingInterval     int         `json:"ping_interval"`
	PingTimeout      int         `json:"ping_timeout"`
	LinkPingInterval int         `json:"link_ping_interval"`
	LinkPingTimeout  int         `json:"link_ping_timeout"`
	InactiveTimeout  int         `json:"inactive_timeout"`
	SessionTimeout   int         `json:"session_timeout"`
	AllowedDevices   string      `json:"allowed_devices"`
	MaxClients       int         `json:"max_clients"`
	MaxDevices       int         `json:"max_devices"`
	ReplicaCount     int         `json:"replica_count"`
	Vxlan            bool        `json:"vxlan"`
	DnsMapping       bool        `json:"dns_mapping"`
	RouteDns         bool        `json:"route_dns"`
	Debug            bool        `json:"debug"`
	SsoAuth          bool        `json:"sso_auth"`
	OtpAuth          bool        `json:"otp_auth"`
	LzoCompression   bool        `json:"lzo_compression"`
	Cipher           string      `json:"cipher"`
	Hash             string      `json:"hash"`
	BlockOutsideDns  bool        `json:"block_outside_dns"`
	JumboFrames      bool        `json:"jumbo_frames"`
	PreConnectMsg    string      `json:"pre_connect_msg"`
	Policy           string      `json:"policy"`
	MssFix           interface{} `json:"mss_fix"`
	Multihome        bool        `json:"multihome"`
}

// ServerResponse represents a server response
type ServerResponse struct {
	ID               string      `json:"id"`
	Status           string      `json:"status"`
	Uptime           uint        `json:"uptime"`
	UsersOnline      int         `json:"users_online"`
	DevicesOnline    int         `json:"devices_online"`
	UserCount        int         `json:"user_count"`
	Name             string      `json:"name"` // Starting here is common to `ServerRequest` Struct
	Network          string      `json:"network"`
	NetworkWg        string      `json:"network_wg"`
	NetworkMode      string      `json:"network_mode"`
	NetworkStart     string      `json:"network_start"`
	NetworkEnd       string      `json:"network_end"`
	RestrictRoutes   bool        `json:"restrict_routes"`
	Wg               bool        `json:"wg"`
	Ipv6             bool        `json:"ipv6"`
	Ipv6Firewall     bool        `json:"ipv6_firewall"`
	DynamicFirewall  bool        `json:"dynamic_firewall"`
	DeviceAuth       bool        `json:"device_auth"`
	BindAddress      string      `json:"bind_address"`
	Protocol         string      `json:"protocol"`
	Port             int         `json:"port"`
	PortWg           int         `json:"port_wg"`
	DhParamBits      int         `json:"dh_param_bits"`
	Groups           []string    `json:"groups"`
	MultiDevice      bool        `json:"multi_device"`
	DnsServers       []string    `json:"dns_servers"`
	SearchDomain     string      `json:"search_domain"`
	InterClient      bool        `json:"inter_client"`
	PingInterval     int         `json:"ping_interval"`
	PingTimeout      int         `json:"ping_timeout"`
	LinkPingInterval int         `json:"link_ping_interval"`
	LinkPingTimeout  int         `json:"link_ping_timeout"`
	InactiveTimeout  int         `json:"inactive_timeout"`
	SessionTimeout   int         `json:"session_timeout"`
	AllowedDevices   string      `json:"allowed_devices"`
	MaxClients       int         `json:"max_clients"`
	MaxDevices       int         `json:"max_devices"`
	ReplicaCount     int         `json:"replica_count"`
	Vxlan            bool        `json:"vxlan"`
	DnsMapping       bool        `json:"dns_mapping"`
	RouteDns         bool        `json:"route_dns"`
	Debug            bool        `json:"debug"`
	SsoAuth          bool        `json:"sso_auth"`
	OtpAuth          bool        `json:"otp_auth"`
	LzoCompression   bool        `json:"lzo_compression"`
	Cipher           string      `json:"cipher"`
	Hash             string      `json:"hash"`
	BlockOutsideDns  bool        `json:"block_outside_dns"`
	JumboFrames      bool        `json:"jumbo_frames"`
	PreConnectMsg    string      `json:"pre_connect_msg"`
	Policy           string      `json:"policy"`
	MssFix           interface{} `json:"mss_fix"`
	Multihome        bool        `json:"multihome"`
}

// ServerGet retrieves a server or servers
func (c *Client) ServerGet(ctx context.Context, srvId ...string) ([]ServerResponse, error) {
	var serverData []byte
	path := "/server"

	// Handle optional srvId argument
	if len(srvId) > 0 {
		path = fmt.Sprintf("%s/%s", path, srvId[0]) // Use the first element if srvId is provided
	}

	// Send an authenticated GET request to retrieve the server(s)
	response, err := c.AuthRequest(ctx, http.MethodGet, path, serverData)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data into a slice of ServerResponse
	var servers []ServerResponse
	if err := handleUnmarshal(body, &servers); err != nil {
		return nil, err
	}

	// Return the slice of servers
	return servers, nil
}

// ServerCreate creates a new server
func (c *Client) ServerCreate(ctx context.Context, newServer ServerRequest) ([]ServerResponse, error) {
	// Marshal the ServerRequest into JSON data
	serverData, err := json.Marshal(newServer)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal server data: %w", err)
	}

	path := "/server"

	// Send an authenticated POST request to create a new server
	response, err := c.AuthRequest(ctx, http.MethodPost, path, serverData)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data into a slice of ServerResponse
	var servers []ServerResponse
	if err := handleUnmarshal(body, &servers); err != nil {
		return nil, err
	}

	// Return the slice of servers
	return servers, nil
}

// ServerUpdate updates an existing server
func (c *Client) ServerUpdate(ctx context.Context, srvId string, newServer ServerRequest) ([]ServerResponse, error) {
	// Marshal the ServerRequest into JSON data
	serverData, err := json.Marshal(newServer)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal server data: %w", err)
	}

	path := fmt.Sprintf("/server/%s", srvId)

	// Send an authenticated PUT request to update the server
	response, err := c.AuthRequest(ctx, http.MethodPut, path, serverData)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data into a slice of ServerResponse
	var servers []ServerResponse
	if err := handleUnmarshal(body, &servers); err != nil {
		return nil, err
	}

	// Return the slice of servers
	return servers, nil
}

// ServerDelete deletes an existing server
func (c *Client) ServerDelete(ctx context.Context, srvId string) ([]ServerResponse, error) {
	var serverData []byte

	path := fmt.Sprintf("/server/%s", srvId)

	// Send an authenticated DELETE request to delete the server
	response, err := c.AuthRequest(ctx, http.MethodDelete, path, serverData)
	if err != nil {
		return nil, err
	}

	body, err := handleResponse(response)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	// Unmarshal the JSON data into a slice of ServerResponse
	var servers []ServerResponse
	if err := handleUnmarshal(body, &servers); err != nil {
		return nil, err
	}

	// Return the slice of servers
	return servers, nil
}
