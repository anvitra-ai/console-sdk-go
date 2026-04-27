package anvitra

import (
	"fmt"
	"net/http"
)

// ListDataRepos lists data repos. Without projectID, returns only global data repos.
// With projectID, returns global + project-scoped data repos.
func (c *Client) ListDataRepos(projectID string) ([]DataRepo, error) {
	var result []DataRepo
	headers := make(map[string]string)
	if projectID != "" {
		headers["X-Project-ID"] = projectID
	}
	err := c.doRequestWithHeaders(http.MethodGet, "/api/datarepo", nil, &result, nil, headers)
	return result, err
}

// GetDataRepo retrieves a data repo by its ID.
// If the data repo is project-scoped, you must provide the matching projectID.
func (c *Client) GetDataRepo(id string, projectID string) (*DataRepo, error) {
	var result DataRepo
	headers := make(map[string]string)
	if projectID != "" {
		headers["X-Project-ID"] = projectID
	}
	err := c.doRequestWithHeaders(http.MethodGet, fmt.Sprintf("/api/datarepo/%s", id), nil, &result, nil, headers)
	return &result, err
}

// GetDataRepoDownloadURL generates a presigned URL to download the file associated with a data repo.
// If the data repo is project-scoped, you must provide the matching projectID.
func (c *Client) GetDataRepoDownloadURL(id string, projectID string) (string, error) {
	var result string
	headers := make(map[string]string)
	if projectID != "" {
		headers["X-Project-ID"] = projectID
	}
	err := c.doRequestWithHeaders(http.MethodGet, fmt.Sprintf("/api/datarepo/%s/download", id), nil, &result, nil, headers)
	return result, err
}
