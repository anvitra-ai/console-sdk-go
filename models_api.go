package anvitra

import (
	"fmt"
	"net/http"
)

// ListModels lists all models in a project.
// If name is provided, returns all versions of that model.
// Requires projectID.
func (c *Client) ListModels(projectID string, name string) ([]Model, error) {
	var result ModelsResponse
	headers := map[string]string{
		"X-Project-ID": projectID,
	}
	queryParams := make(map[string]string)
	if name != "" {
		queryParams["name"] = name
	}
	err := c.doRequestWithHeaders(http.MethodGet, "/api/models", nil, &result, queryParams, headers)
	return result.Models, err
}

// GetModel retrieves a model by its ID. Requires projectID.
func (c *Client) GetModel(id string, projectID string) (*Model, error) {
	var result ModelResponse
	headers := map[string]string{
		"X-Project-ID": projectID,
	}
	err := c.doRequestWithHeaders(http.MethodGet, fmt.Sprintf("/api/models/%s", id), nil, &result, nil, headers)
	return result.Model, err
}

// CreateModel creates a new ML model with metadata and file.
// Note: This endpoint requires multipart/form-data. This method is a placeholder
// and would need to be implemented with proper multipart support.
// Requires projectID.
func (c *Client) CreateModel(projectID string, req CreateModelRequest) (*Model, error) {
	var result ModelResponse
	headers := map[string]string{
		"X-Project-ID": projectID,
	}
	err := c.doRequestWithHeaders(http.MethodPost, "/api/models", req, &result, nil, headers)
	return result.Model, err
}

// UpdateModel updates model metadata (description, status, supported_version).
// Model file cannot be updated. Requires projectID.
func (c *Client) UpdateModel(id string, projectID string, req UpdateModelRequest) (*Model, error) {
	var result ModelResponse
	headers := map[string]string{
		"X-Project-ID": projectID,
	}
	err := c.doRequestWithHeaders(http.MethodPut, fmt.Sprintf("/api/models/%s", id), req, &result, nil, headers)
	return result.Model, err
}

// DeleteModel soft deletes a model. The model will be marked as disabled
// and deleted_at timestamp will be set. Requires projectID.
func (c *Client) DeleteModel(id string, projectID string) error {
	var result BaseResponse
	headers := map[string]string{
		"X-Project-ID": projectID,
	}
	err := c.doRequestWithHeaders(http.MethodDelete, fmt.Sprintf("/api/models/%s", id), nil, &result, nil, headers)
	return err
}

// GetModelDownloadURL generates a presigned URL to download the model file.
// Requires projectID.
func (c *Client) GetModelDownloadURL(id string, projectID string) (*ModelDownloadResponse, error) {
	var result ModelDownloadResponse
	headers := map[string]string{
		"X-Project-ID": projectID,
	}
	err := c.doRequestWithHeaders(http.MethodGet, fmt.Sprintf("/api/models/%s/download", id), nil, &result, nil, headers)
	return &result, err
}
