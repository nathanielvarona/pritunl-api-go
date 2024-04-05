package pritunl

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	BaseUrl   string
	ApiToken  string
	ApiSecret string
}

func PritunlClient(pritunl ...*Client) *Client {
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
		}
		if apiToken == "" {
			apiToken = os.Getenv("PRITUNL_API_TOKEN")
		}
		if apiSecret == "" {
			apiSecret = os.Getenv("PRITUNL_API_SECRET")
		}
	}

	return &Client{
		BaseUrl:   baseURL,
		ApiToken:  apiToken,
		ApiSecret: apiSecret,
	}
}

func (c *Client) AuthRequest(method, path string, headers map[string]string, data []byte) (*http.Response, error) {
	authTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
	authNonce := strings.ReplaceAll(uuid.New().String(), "-", "")
	authString := fmt.Sprintf("%s&%s&%s&%s&%s", c.ApiToken, authTimestamp, authNonce, method, path)
	hash := hmac.New(sha256.New, []byte(c.ApiSecret))
	hash.Write([]byte(authString))
	authSignature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	authHeaders := map[string]string{
		"Auth-Token":     c.ApiToken,
		"Auth-Timestamp": authTimestamp,
		"Auth-Nonce":     authNonce,
		"Auth-Signature": authSignature,
	}

	for k, v := range headers {
		authHeaders[k] = v
	}

	client := &http.Client{}
	req, error := http.NewRequest(method, c.BaseUrl+path, bytes.NewBuffer(data))
	if error != nil {
		return nil, error
	}

	for k, v := range authHeaders {
		req.Header.Set(k, v)
	}

	return client.Do(req)
}
