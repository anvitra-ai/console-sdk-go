package anvitra

import "net/http"

// HealthCheck performs a health check on the Anvitra Console API
func (c *Client) HealthCheck() (*HealthResponse, error) {
	var result HealthResponse
	err := c.doRequest(http.MethodGet, "/health", nil, &result, nil)
	return &result, err
}
