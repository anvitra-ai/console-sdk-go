package anvitra

import "time"

// HealthResponse represents the health check response
type HealthResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// User represents an Anvitra user
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// Project represents an Anvitra project
type Project struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	UserID    string    `json:"user_id,omitempty"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// UserWithProjects represents user details with associated projects
type UserWithProjects struct {
	User     User      `json:"user"`
	Projects []Project `json:"projects"`
}

// DataRepo represents a data repository
type DataRepo struct {
	ID          string                   `json:"id"`
	Name        string                   `json:"name"`
	Description string                   `json:"description,omitempty"`
	ProjectID   string                   `json:"project_id,omitempty"`
	FilePath    string                   `json:"file_path,omitempty"`
	FileType    string                   `json:"file_type,omitempty"`
	FileSize    int64                    `json:"file_size,omitempty"`
	Enabled     bool                     `json:"enabled"`
	Tags        []string                 `json:"tags,omitempty"`
	Schema      []map[string]interface{} `json:"schema,omitempty"`
	PreviewData []map[string]interface{} `json:"preview_data,omitempty"`
	CreatedAt   time.Time                `json:"created_at,omitempty"`
	UpdatedAt   time.Time                `json:"updated_at,omitempty"`
	DeletedAt   *time.Time               `json:"deleted_at,omitempty"`
}

// Model represents an ML model
type Model struct {
	ID                           string                 `json:"id"`
	Name                         string                 `json:"name"`
	Version                      string                 `json:"version"`
	Description                  string                 `json:"description,omitempty"`
	ProjectID                    string                 `json:"project_id"`
	ModelType                    string                 `json:"model_type,omitempty"`
	Mode                         string                 `json:"mode,omitempty"`
	FilePath                     string                 `json:"file_path,omitempty"`
	FileSize                     int64                  `json:"file_size,omitempty"`
	Collection                   string                 `json:"collection,omitempty"`
	EmbeddingDim                 int                    `json:"embedding_dim,omitempty"`
	LabelField                   string                 `json:"label_field,omitempty"`
	Labels                       []string               `json:"labels,omitempty"`
	LabelGrouping                map[string]interface{} `json:"label_grouping,omitempty"`
	ClassifierSelectionStrategy  map[string]interface{} `json:"classifier_selection_strategy,omitempty"`
	NumSamples                   int                    `json:"num_samples,omitempty"`
	Skipped                      int                    `json:"skipped,omitempty"`
	Status                       string                 `json:"status,omitempty"`
	SupportedVersion             string                 `json:"supported_version,omitempty"`
	Enabled                      bool                   `json:"enabled"`
	CreatedAt                    time.Time              `json:"created_at,omitempty"`
	UpdatedAt                    time.Time              `json:"updated_at,omitempty"`
	DeletedAt                    *time.Time             `json:"deleted_at,omitempty"`
}

// CreateModelRequest represents the request to create a model
type CreateModelRequest struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description,omitempty"`
	ModelType   string `json:"model_type,omitempty"`
	Mode        string `json:"mode,omitempty"`
}

// UpdateModelRequest represents the request to update a model
type UpdateModelRequest struct {
	Description      string `json:"description,omitempty"`
	Status           string `json:"status,omitempty"`
	SupportedVersion string `json:"supported_version,omitempty"`
}

// Vertical represents a vertical configuration
type Vertical struct {
	ID                           string                 `json:"id"`
	Name                         string                 `json:"name"`
	Version                      string                 `json:"version"`
	Description                  string                 `json:"description,omitempty"`
	ProjectID                    string                 `json:"project_id"`
	BaseInstructions             string                 `json:"base_instructions,omitempty"`
	DefaultMetricField           *string                `json:"default_metric_field,omitempty"`
	ClassifierSelectionStrategy  map[string]interface{} `json:"classifier_selection_strategy,omitempty"`
	Enabled                      bool                   `json:"enabled"`
	CreatedAt                    time.Time              `json:"created_at,omitempty"`
	UpdatedAt                    time.Time              `json:"updated_at,omitempty"`
	DeletedAt                    *time.Time             `json:"deleted_at,omitempty"`
}

// CreateVerticalRequest represents the request to create a vertical
type CreateVerticalRequest struct {
	Name                        string                 `json:"name"`
	Version                     string                 `json:"version"`
	Description                 string                 `json:"description,omitempty"`
	BaseInstructions            string                 `json:"base_instructions,omitempty"`
	DefaultMetricField          *string                `json:"default_metric_field,omitempty"`
	ClassifierSelectionStrategy map[string]interface{} `json:"classifier_selection_strategy,omitempty"`
}

// UpdateVerticalRequest represents the request to update a vertical
type UpdateVerticalRequest struct {
	Description                 string                 `json:"description,omitempty"`
	BaseInstructions            string                 `json:"base_instructions,omitempty"`
	DefaultMetricField          *string                `json:"default_metric_field,omitempty"`
	ClassifierSelectionStrategy map[string]interface{} `json:"classifier_selection_strategy,omitempty"`
}

// APIToken represents an API token
type APIToken struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	ProjectID  string     `json:"project_id"`
	UserID     string     `json:"user_id"`
	HashedKey  string     `json:"hashed_key,omitempty"`
	LookupHash string     `json:"lookup_hash,omitempty"`
	PlainKey   string     `json:"plain_key,omitempty"` // Only returned on creation
	Enabled    bool       `json:"enabled"`
	ExpiresAt  *time.Time `json:"expires_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at,omitempty"`
	UpdatedAt  time.Time  `json:"updated_at,omitempty"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

// CreateAPITokenRequest represents the request to create an API token
type CreateAPITokenRequest struct {
	Name      string     `json:"name"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}
