package anvitra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client is the client for the Anvitra Console API
type Client struct {
	baseURL    string
	httpClient *http.Client
	authToken  string
}

// ClientOption is a function that configures the Client
type ClientOption func(*Client)

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = client
	}
}

// WithAuth sets the authentication token for the client.
// The token is the api key generated from anvitra console.
func WithAuth(token string) ClientOption {
	return func(c *Client) {
		c.authToken = token
	}
}

// NewClient creates a new Anvitra Console API client
func NewClient(baseURL string, opts ...ClientOption) *Client {
	c := &Client{
		baseURL: strings.TrimRight(baseURL, "/"),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// doRequest performs an HTTP request
func (c *Client) doRequest(method, path string, body interface{}, result interface{}, queryParams map[string]string) error {
	return c.doRequestWithHeaders(method, path, body, result, queryParams, nil)
}

// doRequestWithHeaders performs an HTTP request with additional headers
func (c *Client) doRequestWithHeaders(method, path string, body interface{}, result interface{}, queryParams map[string]string, headers map[string]string) error {
	req, err := c.prepareRequest(method, path, body, queryParams)
	if err != nil {
		return err
	}

	if c.authToken != "" {
		req.Header.Set("X-API-Key", c.authToken)
	}

	// Add custom headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error: %s (status: %d)", string(bodyBytes), resp.StatusCode)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

func (c *Client) prepareRequest(method, path string, body interface{}, queryParams map[string]string) (*http.Request, error) {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	if len(queryParams) > 0 {
		q := u.Query()
		for k, v := range queryParams {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, u.String(), reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
