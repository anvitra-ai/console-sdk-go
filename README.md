# console-sdk-go

Official Go SDK for the [Anvitra Console Platform](https://github.com/anvitra-ai/anvitra-platform).

## Installation

```bash
go get -u github.com/anvitra-ai/console-sdk-go
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	anvitra "github.com/anvitra-ai/console-sdk-go"
)

func main() {
	// Initialise the client (optionally pass a JWT with anvitra.WithAuth)
	client := anvitra.NewClient("https://console.anvitra.ai")

	// Check health
	health, err := client.HealthCheck()
	if err != nil {
		log.Fatalf("Health check failed: %v", err)
	}
	fmt.Printf("Healthy: %v\n", health.Success)

	// Authenticate and reuse the token for subsequent calls
	loginResp, err := client.Login(anvitra.LoginRequest{
		Email:    "user@example.com",
		Password: "secret",
	})
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	authed := anvitra.NewClient("https://console.anvitra.ai",
		anvitra.WithAuth(loginResp.Token),
	)

	// List all accounts
	accounts, err := authed.ListAccounts()
	if err != nil {
		log.Fatalf("Failed to list accounts: %v", err)
	}
	fmt.Printf("Accounts: %+v\n", accounts.Data)

	// Create a new account
	created, err := authed.CreateAccount(anvitra.CreateAccountRequest{
		Name:  "My Org",
		Email: "org@example.com",
	})
	if err != nil {
		log.Fatalf("Failed to create account: %v", err)
	}
	fmt.Printf("Created account: %s\n", created.Data.ID)
}
```

### Instance Management

```go
// Register a Shilp service with the discovery registry
err := client.RegisterShilpService("account-1", "http://shilp:3000", "node-1", anvitra.SingleNode)
if err != nil {
	log.Printf("Failed to register Shilp service: %v", err)
}

// Update sync status
_, err = client.UpdateShilpSyncStatus("account-1", "http://shilp:3000", anvitra.SyncStatusReady)
if err != nil {
	log.Printf("Failed to update sync status: %v", err)
}

// Get Shilp stats for an account
stats, err := client.GetShilpStats("account-1")
if err != nil {
	log.Printf("Failed to get Shilp stats: %v", err)
}
fmt.Printf("Available replicas: %d\n", stats.Registry.Available)

// Register a TEI (Text Embeddings Inference) service
err = client.RegisterTEIService("account-1", "http://tei:8080", "tei-1")
if err != nil {
	log.Printf("Failed to register TEI service: %v", err)
}

// Unregister services when no longer needed
err = client.UnregisterShilpService("account-1", "http://shilp:3000", "node-1", anvitra.SingleNode)
err = client.UnregisterTEIService("account-1", "http://tei:8080", "tei-1")
```

## Features

- **Authentication** – Login and token-based auth via `WithAuth`
- **Account Management** – List, get, create and delete platform accounts
- **Shilp Instance Management** – Register/unregister Shilp nodes in the discovery registry
- **TEI Instance Management** – Register/unregister Text Embeddings Inference services
- **Sync Status** – Update replica sync state for traffic gating
- **Shilp Stats** – Retrieve live registry and proxy statistics per account
- **Health Check** – Verify platform availability

## API Reference

| Method | Description |
|--------|-------------|
| `HealthCheck()` | Check platform health |
| `Login(req)` | Authenticate and obtain a JWT |
| `ListAccounts()` | List all accounts |
| `GetAccount(id)` | Get a single account by ID |
| `CreateAccount(req)` | Create a new account |
| `DeleteAccount(id)` | Delete an account |
| `GetShilpStats(accountID)` | Get Shilp registry and proxy stats |
| `UpdateShilpSyncStatus(accountID, address, status)` | Update replica sync state |
| `RegisterShilpService(accountID, address, id, replicaType)` | Register a Shilp node |
| `UnregisterShilpService(accountID, address, id, replicaType)` | Unregister a Shilp node |
| `RegisterTEIService(accountID, address, id)` | Register a TEI service |
| `UnregisterTEIService(accountID, address, id)` | Unregister a TEI service |
