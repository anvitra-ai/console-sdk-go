package anvitra

// GenericResponse represents the standard response structure
type GenericResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Version string `json:"version,omitempty"`
}

// SyncStatus represents the synchronisation state of a replica
type SyncStatus string

const (
	SyncStatusReady   SyncStatus = "ready"
	SyncStatusSyncing SyncStatus = "syncing"
)

// ReplicaType represents the role of a replica in the cluster
type ReplicaType int

const (
	ReadReplica  ReplicaType = iota // ReadReplica handles read traffic
	WriteReplica                    // WriteReplica handles write traffic
	SingleNode                      // SingleNode acts as both read and write
)

// IsRead reports whether the replica type handles read traffic
func (rt ReplicaType) IsRead() bool {
	return rt == ReadReplica || rt == SingleNode
}

// IsWrite reports whether the replica type handles write traffic
func (rt ReplicaType) IsWrite() bool {
	return rt == WriteReplica || rt == SingleNode
}

// Replica represents a single Shilp service node
type Replica struct {
	Id        string `json:"id"`
	Address   string `json:"address"`
	IsHealthy bool   `json:"is_healthy"`
	IsSyncing bool   `json:"is_syncing"` // Traffic gate - if true, no traffic is sent to this node
}

// RegistryStatus holds the state of the write replica and all read replicas
type RegistryStatus struct {
	WriteReplica Replica    `json:"write_replica"`
	ReadReplicas []*Replica `json:"read_replicas"`
	Available    int        `json:"available_count"`
	Total        int        `json:"total_count"`
}

// ProxyStats contains live proxy information for an account
type ProxyStats struct {
	ActiveProxies int      `json:"active_proxies"`
	Targets       []string `json:"targets"`
}

// ShilpStats is the aggregated discovery statistics for a single account
type ShilpStats struct {
	Registry RegistryStatus `json:"registry"`
	Proxy    ProxyStats     `json:"proxy"`
}

// RegisterInstanceRequest is the payload for registering a Shilp or TEI service
type RegisterInstanceRequest struct {
	AccountID string `json:"account_id"`
	Address   string `json:"address"`
	Id        string `json:"id"`
	IsRead    bool   `json:"is_read"`
	IsWrite   bool   `json:"is_write"`
}

// UpdateSyncStatusRequest is the payload for updating the sync state of a replica
type UpdateSyncStatusRequest struct {
	AccountID string     `json:"account_id"`
	Address   string     `json:"address"`
	Status    SyncStatus `json:"status"`
}

// Account represents an Anvitra platform account
type Account struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at,omitempty"`
}

// CreateAccountRequest is the payload for creating a new account
type CreateAccountRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

// ListAccountsResponse is the response for the list-accounts endpoint
type ListAccountsResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Account `json:"data"`
}

// GetAccountResponse is the response for the get-account endpoint
type GetAccountResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    *Account `json:"data,omitempty"`
}

// LoginRequest is the payload for the login endpoint
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse is the response for the login endpoint
type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}
