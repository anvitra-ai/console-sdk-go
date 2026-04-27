package anvitra

import (
	"fmt"
	"net/http"
)

// Login authenticates with the Anvitra Console API and returns a JWT token
func (c *Client) Login(req LoginRequest) (*LoginResponse, error) {
	var result LoginResponse
	err := c.doRequest(http.MethodPost, "/api/v1/auth/login", req, &result, nil)
	return &result, err
}

// ListAccounts returns all accounts visible to the authenticated user
func (c *Client) ListAccounts() (*ListAccountsResponse, error) {
	var result ListAccountsResponse
	err := c.doRequest(http.MethodGet, "/api/v1/accounts", nil, &result, nil)
	return &result, err
}

// GetAccount returns the account with the given ID
func (c *Client) GetAccount(id string) (*GetAccountResponse, error) {
	var result GetAccountResponse
	path := fmt.Sprintf("/api/v1/accounts/%s", id)
	err := c.doRequest(http.MethodGet, path, nil, &result, nil)
	return &result, err
}

// CreateAccount creates a new account on the platform
func (c *Client) CreateAccount(req CreateAccountRequest) (*GetAccountResponse, error) {
	var result GetAccountResponse
	err := c.doRequest(http.MethodPost, "/api/v1/accounts", req, &result, nil)
	return &result, err
}

// DeleteAccount removes the account with the given ID
func (c *Client) DeleteAccount(id string) (*GenericResponse, error) {
	var result GenericResponse
	path := fmt.Sprintf("/api/v1/accounts/%s", id)
	err := c.doRequest(http.MethodDelete, path, nil, &result, nil)
	return &result, err
}
