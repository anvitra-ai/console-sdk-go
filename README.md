# console-sdk-go

Official Go SDK for the [Anvitra Console Platform](https://github.com/anvitra-ai/anvitra-platform).

## Installation

```bash
go get -u github.com/anvitra-ai/console-sdk-go
```

## Usage

### Basic Setup

```go
package main

import (
	"fmt"
	"log"

	anvitra "github.com/anvitra-ai/console-sdk-go"
)

func main() {
	// Initialize the client with your API key
	client := anvitra.NewClient(
		"https://console.anvitra.ai",
		anvitra.WithAuth("your-api-key"),
	)

	// Check health
	health, err := client.HealthCheck()
	if err != nil {
		log.Fatalf("Health check failed: %v", err)
	}
	fmt.Printf("Healthy: %v\n", health.Success)
}
```

### User Management

```go
// Get current user and their projects
userWithProjects, err := client.GetCurrentUser()
if err != nil {
	log.Fatalf("Failed to get user: %v", err)
}
fmt.Printf("User: %s (%s)\n", userWithProjects.User.Name, userWithProjects.User.Email)
fmt.Printf("Projects: %d\n", len(userWithProjects.Projects))
```

### API Token Management

```go
// List all API tokens
tokens, err := client.ListAPITokens()
if err != nil {
	log.Fatalf("Failed to list tokens: %v", err)
}

// Create a new API token
token, err := client.CreateAPIToken(anvitra.CreateAPITokenRequest{
	Name: "My API Token",
	ExpiresAt: nil, // Optional: set expiration time
})
if err != nil {
	log.Fatalf("Failed to create token: %v", err)
}
fmt.Printf("Token created! Keep this safe: %s\n", token.PlainKey)

// Delete an API token
err = client.DeleteAPIToken("token-id")
if err != nil {
	log.Fatalf("Failed to delete token: %v", err)
}
```

### Data Repository Management

```go
// List all data repos (global only)
repos, err := client.ListDataRepos("")
if err != nil {
	log.Fatalf("Failed to list data repos: %v", err)
}

// List data repos for a specific project
repos, err = client.ListDataRepos("project-id")
if err != nil {
	log.Fatalf("Failed to list data repos: %v", err)
}

// Get a specific data repo
repo, err := client.GetDataRepo("repo-id", "project-id")
if err != nil {
	log.Fatalf("Failed to get data repo: %v", err)
}
fmt.Printf("Repo: %s (%s)\n", repo.Name, repo.Description)

// Get download URL for data repo file
url, err := client.GetDataRepoDownloadURL("repo-id", "project-id")
if err != nil {
	log.Fatalf("Failed to get download URL: %v", err)
}
fmt.Printf("Download URL: %s\n", url)
```

### ML Model Management

```go
// List all models in a project
models, err := client.ListModels("project-id", "")
if err != nil {
	log.Fatalf("Failed to list models: %v", err)
}

// List all versions of a specific model
models, err = client.ListModels("project-id", "model-name")
if err != nil {
	log.Fatalf("Failed to list model versions: %v", err)
}

// Get a specific model
model, err := client.GetModel("model-id", "project-id")
if err != nil {
	log.Fatalf("Failed to get model: %v", err)
}
fmt.Printf("Model: %s v%s\n", model.Name, model.Version)

// Create a new model
newModel, err := client.CreateModel("project-id", anvitra.CreateModelRequest{
	Name:        "my-classifier",
	Version:     "1.0.0",
	Description: "Product classification model",
	ModelType:   "classifier",
	Mode:        "inference",
})
if err != nil {
	log.Fatalf("Failed to create model: %v", err)
}

// Update model metadata
updated, err := client.UpdateModel("model-id", "project-id", anvitra.UpdateModelRequest{
	Description:      "Updated description",
	Status:           "active",
	SupportedVersion: "2.0",
})
if err != nil {
	log.Fatalf("Failed to update model: %v", err)
}

// Get model download URL
url, err := client.GetModelDownloadURL("model-id", "project-id")
if err != nil {
	log.Fatalf("Failed to get download URL: %v", err)
}

// Delete a model
err = client.DeleteModel("model-id", "project-id")
if err != nil {
	log.Fatalf("Failed to delete model: %v", err)
}
```

### Vertical Management

```go
// List all verticals in a project
verticals, err := client.ListVerticals("project-id", "")
if err != nil {
	log.Fatalf("Failed to list verticals: %v", err)
}

// List all versions of a specific vertical
verticals, err = client.ListVerticals("project-id", "vertical-name")
if err != nil {
	log.Fatalf("Failed to list vertical versions: %v", err)
}

// Get a specific vertical
vertical, err := client.GetVertical("vertical-id", "project-id")
if err != nil {
	log.Fatalf("Failed to get vertical: %v", err)
}

// Create a new vertical
newVertical, err := client.CreateVertical("project-id", anvitra.CreateVerticalRequest{
	Name:             "ecommerce",
	Version:          "1.0.0",
	Description:      "E-commerce vertical configuration",
	BaseInstructions: "Handle e-commerce queries",
})
if err != nil {
	log.Fatalf("Failed to create vertical: %v", err)
}

// Update vertical configuration
updated, err := client.UpdateVertical("vertical-id", "project-id", anvitra.UpdateVerticalRequest{
	Description:      "Updated e-commerce vertical",
	BaseInstructions: "Enhanced e-commerce query handling",
})
if err != nil {
	log.Fatalf("Failed to update vertical: %v", err)
}

// Delete a vertical
err = client.DeleteVertical("vertical-id", "project-id")
if err != nil {
	log.Fatalf("Failed to delete vertical: %v", err)
}
```

## Features

- **User Management**: Get current user details and associated projects
- **API Token Management**: Create, list, and delete API tokens
- **Data Repository Management**: List, retrieve, and download data repositories
- **ML Model Management**: Full CRUD operations for ML models with versioning
- **Vertical Management**: Configure and manage vertical-specific settings
- **Authentication**: Bearer token authentication with API keys
- **Project Scoping**: Support for X-Project-ID header for project-specific operations

## API Documentation

For complete API documentation, refer to the [Anvitra Console API specification](https://console.anvitra.ai/docs).

## License

MIT

- **Authentication** – Login and token-based auth via `WithAuth`
- **Account Management** – List, get, create and delete platform accounts
- **Shilp Instance Management** – Register/unregister Shilp nodes in the discovery registry
- **TEI Instance Management** – Register/unregister Text Embeddings Inference services
- **Sync Status** – Update replica sync state for traffic gating
- **Shilp Stats** – Retrieve live registry and proxy statistics per account
- **Health Check** – Verify platform availability
