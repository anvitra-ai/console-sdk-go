package anvitra

import (
	"fmt"
	"net/http"
)

// ListAPITokens lists all API tokens for the authenticated user.
// Plain tokens are not included in the response.
func (c *Client) ListAPITokens() ([]APIToken, error) {
	var result APITokensResponse
	err := c.doRequest(http.MethodGet, "/api/tokens", nil, &result, nil)
	return result.APIKeys, err
}

// CreateAPIToken creates a new API token for a project.
// The plain token is returned only once and cannot be retrieved again.
func (c *Client) CreateAPIToken(req CreateAPITokenRequest) (*APIToken, error) {
	var result APITokenResponse
	err := c.doRequest(http.MethodPost, "/api/tokens", req, &result, nil)
	return result.APIKey, err
}

// DeleteAPIToken deletes an API token by ID.
// Users can only delete their own tokens.
func (c *Client) DeleteAPIToken(id string) error {
	var result BaseResponse
	err := c.doRequest(http.MethodDelete, fmt.Sprintf("/api/tokens/%s", id), nil, &result, nil)
	return err
}
