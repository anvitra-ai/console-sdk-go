package anvitra

import (
	"fmt"
	"net/http"
)

// ListAPITokens lists all API tokens for the authenticated user.
// Plain tokens are not included in the response.
func (c *Client) ListAPITokens() ([]APIToken, error) {
	var result []APIToken
	err := c.doRequest(http.MethodGet, "/api/tokens", nil, &result, nil)
	return result, err
}

// CreateAPIToken creates a new API token for a project.
// The plain token is returned only once and cannot be retrieved again.
func (c *Client) CreateAPIToken(req CreateAPITokenRequest) (*APIToken, error) {
	var result APIToken
	err := c.doRequest(http.MethodPost, "/api/tokens", req, &result, nil)
	return &result, err
}

// DeleteAPIToken deletes an API token by ID.
// Users can only delete their own tokens.
func (c *Client) DeleteAPIToken(id string) error {
	err := c.doRequest(http.MethodDelete, fmt.Sprintf("/api/tokens/%s", id), nil, nil, nil)
	return err
}
