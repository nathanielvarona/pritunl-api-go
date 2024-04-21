package pritunl

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Client represents a Pritunl API client
type Client struct {
	BaseUrl   string
	ApiToken  string
	ApiSecret string
}

// NewClient creates a new Pritunl client instance
func NewClient(pritunl ...*Client) (*Client, error) {
	var baseURL string
	var apiToken string
	var apiSecret string

	// Check environment variables first
	baseURL = os.Getenv("PRITUNL_BASE_URL")
	apiToken = os.Getenv("PRITUNL_API_TOKEN")
	apiSecret = os.Getenv("PRITUNL_API_SECRET")

	// Then check arguments (if provided)
	if len(pritunl) > 0 && pritunl[0] != nil {
		// Override environment variables with arguments (optional)
		if pritunl[0].BaseUrl != "" {
			baseURL = pritunl[0].BaseUrl
		}
		if pritunl[0].ApiToken != "" {
			apiToken = pritunl[0].ApiToken
		}
		if pritunl[0].ApiSecret != "" {
			apiSecret = pritunl[0].ApiSecret
		}
	}

	// Check for missing credentials and return error if necessary
	if baseURL == "" || apiToken == "" || apiSecret == "" {
		return nil, errors.New("missing Pritunl API Access configuration: Base URL, API Token, or API Secret")
	}

	// Return a new Client instance with the configured credentials
	return &Client{
		BaseUrl:   baseURL,
		ApiToken:  apiToken,
		ApiSecret: apiSecret,
	}, nil
}

// AuthRequest performs an authenticated API request
func (c *Client) AuthRequest(ctx context.Context, method, path string, data []byte) (*http.Response, error) {
	// Generate a timestamp and nonce for authentication
	authTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
	authNonce := strings.ReplaceAll(uuid.New().String(), "-", "")

	// Construct the authentication string
	authString := fmt.Sprintf("%s&%s&%s&%s&%s", c.ApiToken, authTimestamp, authNonce, method, path)

	// Generate the authentication signature using HMAC SHA256
	hash := hmac.New(sha256.New, []byte(c.ApiSecret))
	hash.Write([]byte(authString))
	authSignature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	// Set the authentication headers
	headers := map[string]string{
		"Auth-Token":     c.ApiToken,
		"Auth-Timestamp": authTimestamp,
		"Auth-Nonce":     authNonce,
		"Auth-Signature": authSignature,
		"Content-Type":   "application/json", // Default content type
	}

	// Create a new HTTP request
	req, err := http.NewRequestWithContext(ctx, method, c.BaseUrl+path, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	// Set the request headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Send the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
