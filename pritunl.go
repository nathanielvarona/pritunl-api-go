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

	if len(pritunl) > 0 && pritunl[0] != nil {
		baseURL = pritunl[0].BaseUrl
		apiToken = pritunl[0].ApiToken
		apiSecret = pritunl[0].ApiSecret
	} else {
		if baseURL == "" {
			baseURL = os.Getenv("PRITUNL_BASE_URL")
			if baseURL == "" {
				return nil, errors.New("missing Pritunl base URL")
			}
		}
		if apiToken == "" {
			apiToken = os.Getenv("PRITUNL_API_TOKEN")
			if apiToken == "" {
				return nil, errors.New("missing Pritunl API token")
			}
		}
		if apiSecret == "" {
			apiSecret = os.Getenv("PRITUNL_API_SECRET")
			if apiSecret == "" {
				return nil, errors.New("missing Pritunl API secret")
			}
		}
	}

	return &Client{
		BaseUrl:   baseURL,
		ApiToken:  apiToken,
		ApiSecret: apiSecret,
	}, nil
}

// AuthRequest performs an authenticated API request
func (c *Client) AuthRequest(ctx context.Context, method, path string, data []byte) (*http.Response, error) {
	authTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
	authNonce := strings.ReplaceAll(uuid.New().String(), "-", "")
	authString := fmt.Sprintf("%s&%s&%s&%s&%s", c.ApiToken, authTimestamp, authNonce, method, path)
	hash := hmac.New(sha256.New, []byte(c.ApiSecret))
	hash.Write([]byte(authString))
	authSignature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	headers := map[string]string{
		"Auth-Token":     c.ApiToken,
		"Auth-Timestamp": authTimestamp,
		"Auth-Nonce":     authNonce,
		"Auth-Signature": authSignature,
		"Content-Type":   "application/json", // Default content type
	}

	req, err := http.NewRequestWithContext(ctx, method, c.BaseUrl+path, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
