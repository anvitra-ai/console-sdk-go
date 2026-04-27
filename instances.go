package anvitra

import (
	"fmt"
	"net/http"
)

// GetShilpStats returns the discovery statistics for the given account
func (c *Client) GetShilpStats(accountID string) (*ShilpStats, error) {
	var result ShilpStats
	err := c.doRequest(http.MethodGet, "/control/shilp/stats", nil, &result, map[string]string{"account_id": accountID})
	return &result, err
}

// UpdateShilpSyncStatus updates the sync state of a Shilp replica for the given account
func (c *Client) UpdateShilpSyncStatus(accountID, address string, status SyncStatus) (*GenericResponse, error) {
	req := UpdateSyncStatusRequest{
		AccountID: accountID,
		Address:   address,
		Status:    status,
	}
	var result GenericResponse
	err := c.doRequest(http.MethodPost, "/control/shilp/sync-status", req, &result, nil)
	return &result, err
}

// RegisterShilpService registers a Shilp service instance in the discovery registry.
// For SingleNode replicas both a read and a write registration are performed.
func (c *Client) RegisterShilpService(accountID, address, id string, replicaType ReplicaType) error {
	return c.applyShilpRegistration(accountID, address, id, replicaType, c.registerShilp)
}

// UnregisterShilpService removes a Shilp service instance from the discovery registry.
// For SingleNode replicas both the read and write registrations are removed.
func (c *Client) UnregisterShilpService(accountID, address, id string, replicaType ReplicaType) error {
	return c.applyShilpRegistration(accountID, address, id, replicaType, c.unregisterShilp)
}

// applyShilpRegistration calls fn for the appropriate replica role(s) derived from replicaType.
// SingleNode triggers two calls – one for read and one for write.
func (c *Client) applyShilpRegistration(
	accountID, address, id string,
	replicaType ReplicaType,
	fn func(RegisterInstanceRequest) (*GenericResponse, error),
) error {
	base := RegisterInstanceRequest{
		AccountID: accountID,
		Address:   address,
		Id:        id,
	}

	switch replicaType {
	case ReadReplica:
		base.IsRead = true
		_, err := fn(base)
		return err
	case WriteReplica:
		base.IsWrite = true
		_, err := fn(base)
		return err
	default: // SingleNode
		base.IsRead = true
		if _, err := fn(base); err != nil {
			return fmt.Errorf("failed to apply read registration: %w", err)
		}
		base.IsRead = false
		base.IsWrite = true
		if _, err := fn(base); err != nil {
			return fmt.Errorf("failed to apply write registration: %w", err)
		}
		return nil
	}
}

// RegisterTEIService registers a Text Embeddings Inference (TEI) service in the discovery registry
func (c *Client) RegisterTEIService(accountID, address, id string) error {
	payload := RegisterInstanceRequest{
		AccountID: accountID,
		Address:   address,
		Id:        id,
		IsRead:    true,
	}
	return c.doRequest(http.MethodPost, "/control/tei/register", payload, nil, nil)
}

// UnregisterTEIService removes a TEI service from the discovery registry
func (c *Client) UnregisterTEIService(accountID, address, id string) error {
	payload := RegisterInstanceRequest{
		AccountID: accountID,
		Address:   address,
		Id:        id,
		IsRead:    true,
	}
	return c.doRequest(http.MethodPost, "/control/tei/unregister", payload, nil, nil)
}

func (c *Client) registerShilp(payload RegisterInstanceRequest) (*GenericResponse, error) {
	var result GenericResponse
	err := c.doRequest(http.MethodPost, "/control/shilp/register", payload, &result, nil)
	return &result, err
}

func (c *Client) unregisterShilp(payload RegisterInstanceRequest) (*GenericResponse, error) {
	var result GenericResponse
	err := c.doRequest(http.MethodPost, "/control/shilp/unregister", payload, &result, nil)
	return &result, err
}
