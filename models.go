package anvitra

import "time"

// HealthResponse represents the health check response
type HealthResponse struct {
	Success bool   `json:"success"`
	Version string `json:"version"`
}

// User represents an Anvitra user
type User struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	Name      string     `json:"name,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// Project represents an Anvitra project
type Project struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	UserID    string     `json:"user_id,omitempty"`
	Enabled   bool       `json:"enabled"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
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
	ID                          string                 `json:"id"`
	Name                        string                 `json:"name"`
	Version                     string                 `json:"version"`
	Description                 string                 `json:"description,omitempty"`
	ProjectID                   string                 `json:"project_id"`
	ModelType                   string                 `json:"model_type,omitempty"`
	Mode                        string                 `json:"mode,omitempty"`
	FilePath                    string                 `json:"file_path,omitempty"`
	FileSize                    int64                  `json:"file_size,omitempty"`
	Collection                  string                 `json:"collection,omitempty"`
	EmbeddingDim                int                    `json:"embedding_dim,omitempty"`
	LabelField                  string                 `json:"label_field,omitempty"`
	Labels                      []string               `json:"labels,omitempty"`
	LabelGrouping               map[string]interface{} `json:"label_grouping,omitempty"`
	ClassifierSelectionStrategy map[string]interface{} `json:"classifier_selection_strategy,omitempty"`
	NumSamples                  int                    `json:"num_samples,omitempty"`
	Skipped                     int                    `json:"skipped,omitempty"`
	Status                      string                 `json:"status,omitempty"`
	SupportedVersion            string                 `json:"supported_version,omitempty"`
	Enabled                     bool                   `json:"enabled"`
	CreatedAt                   time.Time              `json:"created_at,omitempty"`
	UpdatedAt                   time.Time              `json:"updated_at,omitempty"`
	DeletedAt                   *time.Time             `json:"deleted_at,omitempty"`
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

// AttributeType represents the data type of a field
type AttributeType int

const (
	AttributeTypeNumerical AttributeType = 1
	AttributeTypeString    AttributeType = 2
)

// ComparisonOperator represents comparison operators for conditions
type ComparisonOperator string

const (
	ComparisonOperatorGT    ComparisonOperator = "GT"
	ComparisonOperatorGTE   ComparisonOperator = "GTE"
	ComparisonOperatorLT    ComparisonOperator = "LT"
	ComparisonOperatorLTE   ComparisonOperator = "LTE"
	ComparisonOperatorEQ    ComparisonOperator = "EQ"
	ComparisonOperatorNEQ   ComparisonOperator = "NEQ"
	ComparisonOperatorIN    ComparisonOperator = "IN"
	ComparisonOperatorNOTIN ComparisonOperator = "NOTIN"
)

// ConditionValue represents a value in a condition (either numerical or text)
type ConditionValue struct {
	Numerical *float64 `json:"numerical,omitempty"`
	Text      *string  `json:"text,omitempty"`
}

// GeneratedCondition represents a condition generated for phrase rules
type GeneratedCondition struct {
	Operator ComparisonOperator `json:"operator"`
	Value    ConditionValue     `json:"value"`
}

// FieldConfig represents canonical fields and their synonyms
type FieldConfig struct {
	CanonicalName string        `json:"canonical_name"`
	Synonyms      []string      `json:"synonyms"`
	IsMetadata    bool          `json:"is_metadata"`
	DataType      AttributeType `json:"data_type"`
}

// ValueSynonymConfig represents value-level synonyms
type ValueSynonymConfig struct {
	Field          string   `json:"field"`
	CanonicalValue string   `json:"canonical_value"`
	Synonyms       []string `json:"synonyms"`
}

// PhraseRuleConfig represents phrase rules with optional generated conditions
type PhraseRuleConfig struct {
	Phrase      string              `json:"phrase"`
	MappedField string              `json:"mapped_field"`
	Condition   *GeneratedCondition `json:"condition,omitempty"`
}

// Vertical represents a vertical configuration
type Vertical struct {
	ID        string `json:"id"`
	ProjectID string `json:"project_id"`
	Name      string `json:"name"`
	Label     string `json:"label"`
	Version   string `json:"version"`

	// LLM configuration
	BaseInstructions   string `json:"base_instructions"`
	DomainInstructions string `json:"domain_instructions"`
	Examples           string `json:"examples"`

	// Schema enrichment
	Fields        []FieldConfig        `json:"fields"`
	ValueSynonyms []ValueSynonymConfig `json:"value_synonyms"`
	PhraseRules   []PhraseRuleConfig   `json:"phrase_rules"`
	WordsToRemove []string             `json:"words_to_remove"`

	// Classifier configuration
	ClassifierSelectionStrategy *ClassifierSelectionStrategy `json:"classifier_selection_strategy,omitempty"`
	DefaultMetricField          *string                      `json:"default_metric_field,omitempty"`

	// Standard fields
	Enabled   bool       `json:"enabled"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type ClassifierSelectionStrategy struct {
	Method    string   `json:"method"` // e.g., "confidence_gap", "cumulative_threshold"
	Threshold *float32 `json:"threshold,omitempty"`
	K         *int     `json:"k,omitempty"`
}

// CreateVerticalRequest represents the request to create a vertical
type CreateVerticalRequest struct {
	Name    string `json:"name"`
	Label   string `json:"label"`
	Version string `json:"version"`

	// Optional LLM configuration
	BaseInstructions   string `json:"base_instructions,omitempty"`
	DomainInstructions string `json:"domain_instructions,omitempty"`
	Examples           string `json:"examples,omitempty"`

	// Optional schema enrichment
	Fields        []FieldConfig        `json:"fields,omitempty"`
	ValueSynonyms []ValueSynonymConfig `json:"value_synonyms,omitempty"`
	PhraseRules   []PhraseRuleConfig   `json:"phrase_rules,omitempty"`
	WordsToRemove []string             `json:"words_to_remove,omitempty"`

	// Optional classifier configuration
	ClassifierSelectionStrategy *ClassifierSelectionStrategy `json:"classifier_selection_strategy,omitempty"`
	DefaultMetricField          *string                      `json:"default_metric_field,omitempty"`
}

// UpdateVerticalRequest represents the request to update a vertical
type UpdateVerticalRequest struct {
	Label              *string `json:"label,omitempty"`
	BaseInstructions   *string `json:"base_instructions,omitempty"`
	DomainInstructions *string `json:"domain_instructions,omitempty"`
	Examples           *string `json:"examples,omitempty"`

	// Schema enrichment (if provided, replaces existing)
	Fields        *[]FieldConfig        `json:"fields,omitempty"`
	ValueSynonyms *[]ValueSynonymConfig `json:"value_synonyms,omitempty"`
	PhraseRules   *[]PhraseRuleConfig   `json:"phrase_rules,omitempty"`
	WordsToRemove *[]string             `json:"words_to_remove,omitempty"`

	// Classifier configuration
	ClassifierSelectionStrategy *ClassifierSelectionStrategy `json:"classifier_selection_strategy,omitempty"`
	DefaultMetricField          *string                      `json:"default_metric_field,omitempty"`
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

// Response wrapper types

// BaseResponse represents a standard API response
type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// UserResponse represents the response for GetCurrentUser
type UserResponse struct {
	Success  bool      `json:"success"`
	Message  string    `json:"message,omitempty"`
	User     User      `json:"user"`
	Projects []Project `json:"projects"`
}

// DataReposResponse represents the response for ListDataRepos
type DataReposResponse struct {
	Success   bool       `json:"success"`
	Message   string     `json:"message,omitempty"`
	DataRepos []DataRepo `json:"data_repos"`
}

// DataRepoResponse represents the response for GetDataRepo
type DataRepoResponse struct {
	Success  bool      `json:"success"`
	Message  string    `json:"message,omitempty"`
	DataRepo *DataRepo `json:"data_repo"`
}

// DataRepoDownloadResponse represents the response for GetDataRepoDownloadURL
type DataRepoDownloadResponse struct {
	Success     bool   `json:"success"`
	Message     string `json:"message,omitempty"`
	DownloadURL string `json:"download_url"`
	ExpiresAt   string `json:"expires_at,omitempty"`
	ExpiresIn   string `json:"expires_in,omitempty"`
	FileName    string `json:"file_name,omitempty"`
	FileSize    int64  `json:"file_size,omitempty"`
	FileType    string `json:"file_type,omitempty"`
}

// ModelsResponse represents the response for ListModels
type ModelsResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message,omitempty"`
	Models  []Model `json:"models"`
}

// ModelResponse represents the response for GetModel, CreateModel, and UpdateModel
type ModelResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Model   *Model `json:"model"`
}

// ModelDownloadResponse represents the response for GetModelDownloadURL
type ModelDownloadResponse struct {
	Success     bool   `json:"success"`
	Message     string `json:"message,omitempty"`
	DownloadURL string `json:"download_url"`
	ExpiresIn   string `json:"expires_in,omitempty"`
	FileName    string `json:"file_name,omitempty"`
	FileSize    int64  `json:"file_size,omitempty"`
}

// APITokensResponse represents the response for ListAPITokens
type APITokensResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message,omitempty"`
	APIKeys []APIToken `json:"api_keys"`
}

// APITokenResponse represents the response for CreateAPIToken
type APITokenResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message,omitempty"`
	APIKey  *APIToken `json:"api_key"`
}

// VerticalsResponse represents the response for ListVerticals
type VerticalsResponse struct {
	Success   bool       `json:"success"`
	Message   string     `json:"message,omitempty"`
	Verticals []Vertical `json:"verticals"`
}

// VerticalResponse represents the response for GetVertical, CreateVertical, and UpdateVertical
type VerticalResponse struct {
	Success  bool      `json:"success"`
	Message  string    `json:"message,omitempty"`
	Vertical *Vertical `json:"vertical"`
}
