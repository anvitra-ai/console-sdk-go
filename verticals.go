package anvitra

import (
	"fmt"
	"net/http"
)

// ListVerticals lists all verticals in a project.
// If name is provided, returns all versions of that vertical.
// Requires projectID.
func (c *Client) ListVerticals(projectID string, name string) ([]Vertical, error) {
	var result VerticalsResponse
	headers := map[string]string{
		"X-Project-ID": projectID,
	}
	queryParams := make(map[string]string)
	if name != "" {
		queryParams["name"] = name
	}
	err := c.doRequestWithHeaders(http.MethodGet, "/api/verticals", nil, &result, queryParams, headers)
	return result.Verticals, err
}

// GetVertical retrieves a vertical by its ID. Requires projectID.
func (c *Client) GetVertical(id string, projectID string) (*Vertical, error) {
	var result VerticalResponse
	headers := map[string]string{
		"X-Project-ID": projectID,
	}
	err := c.doRequestWithHeaders(http.MethodGet, fmt.Sprintf("/api/verticals/%s", id), nil, &result, nil, headers)
	return result.Vertical, err
}

// CreateVertical creates a new vertical configuration. Requires projectID.
func (c *Client) CreateVertical(projectID string, req CreateVerticalRequest) (*Vertical, error) {
	var result VerticalResponse
	headers := map[string]string{
		"X-Project-ID": projectID,
	}
	err := c.doRequestWithHeaders(http.MethodPost, "/api/verticals", req, &result, nil, headers)
	return result.Vertical, err
}

// UpdateVertical updates vertical metadata and configuration. Requires projectID.
func (c *Client) UpdateVertical(id string, projectID string, req UpdateVerticalRequest) (*Vertical, error) {
	var result VerticalResponse
	headers := map[string]string{
		"X-Project-ID": projectID,
	}
	err := c.doRequestWithHeaders(http.MethodPut, fmt.Sprintf("/api/verticals/%s", id), req, &result, nil, headers)
	return result.Vertical, err
}

// DeleteVertical soft deletes a vertical. The vertical will be marked as disabled
// and deleted_at timestamp will be set. Requires projectID.
func (c *Client) DeleteVertical(id string, projectID string) error {
	var result BaseResponse
	headers := map[string]string{
		"X-Project-ID": projectID,
	}
	err := c.doRequestWithHeaders(http.MethodDelete, fmt.Sprintf("/api/verticals/%s", id), nil, &result, nil, headers)
	return err
}
