package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/nathanielvarona/pritunl-api-go"
)

func main() {
	// Initialize the Pritunl API client
	client, err := pritunl.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the server ID from an environment variable
	server := os.Getenv("PRITUNL_DATA_SERVER")

	// Create a context for the request
	ctx := context.Background()

	// Create a ServerRequest object with desired data
	newServer := &pritunl.ServerRequest{
		Name:             "pritunl.server.updated",
		DnsServers:       []string{"8.8.8.8"}, // Use only one DNS server here
		Port:             14789,
		Network:          "192.168.234.0/24",
		Groups:           []string{},
		NetworkMode:      "tunnel",
		NetworkStart:     nil,
		NetworkEnd:       nil,
		RestrictRoutes:   true,
		Ipv6:             false,
		Ipv6Firewall:     true,
		BindAddress:      nil,
		Protocol:         "tcp",
		DhParamBits:      2048,
		MultiDevice:      false,
		OtpAuth:          false,
		Cipher:           "aes128",
		Hash:             "sha1",
		JumboFrames:      false,
		LzoCompression:   false,
		InterClient:      true,
		PingInterval:     10,
		PingTimeout:      60,
		LinkPingInterval: 1,
		LinkPingTimeout:  5,
		MaxClients:       10,
		ReplicaCount:     1,
		Vxlan:            true,
		DnsMapping:       false,
		Debug:            false,
	}

	// Update the specified server
	servers, err := client.ServerUpdate(ctx, server, *newServer)
	if err != nil {
		log.Fatal(err)
	}

	// Print the updated server details
	fmt.Println("Updated Server:")
	for _, server := range servers {
		fmt.Println("Server Name:", server.Name)
		fmt.Println("Server ID:", server.ID)
		fmt.Println("Server Status:", server.Status)
		fmt.Println("Server Uptime:", server.Uptime)
		fmt.Println("------")
	}
}
