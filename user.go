package anvitra

import "net/http"

// GetCurrentUser fetches the authenticated user.
// Creates the user and a default project if they don't exist.
func (c *Client) GetCurrentUser() (*UserWithProjects, error) {
	var result UserWithProjects
	err := c.doRequest(http.MethodGet, "/api/user/me", nil, &result, nil)
	return &result, err
}
