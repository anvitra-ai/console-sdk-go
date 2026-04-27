package anvitra

import "net/http"

// GetCurrentUser fetches the authenticated user.
// Creates the user and a default project if they don't exist.
func (c *Client) GetCurrentUser() (*UserWithProjects, error) {
	var result UserResponse
	err := c.doRequest(http.MethodGet, "/api/user/me", nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &UserWithProjects{
		User:     result.User,
		Projects: result.Projects,
	}, nil
}
