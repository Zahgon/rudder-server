package fileuploader

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/filemanager"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

type StorageSettings struct {
	Bucket      backendconfig.StorageBucket
	Preferences backendconfig.StoragePreferences
	updatedAt   time.Time
}

var (
	ErrNoStorageForWorkspace = fmt.Errorf("no storage settings found for workspace")
	ErrNotSubscribed         = fmt.Errorf("provider not subscribed to backend config")
)

// Provider is an interface that provides file managers and storage preferences for a given workspace.
type Provider interface {
	// GetFileManager gets a file manager for the given workspace.
	GetFileManager(ctx context.Context, workspaceID string) (filemanager.FileManager, error)
	// GetStoragePreferences gets the storage preferences for the given workspace.
	GetStoragePreferences(ctx context.Context, workspaceID string) (backendconfig.StoragePreferences, error)
}

// NewProvider creates a new provider that updates its storage settings while backend configuration gets updated.
func NewProvider(ctx context.Context, config backendconfig.BackendConfig) Provider {
	_ = "STUB: not implemented"
	return *new(Provider)
}

// NewStaticProvider creates a new provider that operates against a predefined storage settings.
// Useful for tests.
func NewStaticProvider(storageSettings map[string]StorageSettings) Provider {
	_ = "STUB: not implemented"
	return *new(Provider)
}

// NewDefaultProvider creates a new provider that operates against the default storage settings populated from the env.
// Useful for tests that populate settings from env.
func NewDefaultProvider() Provider { _ = "STUB: not implemented"; return *new(Provider) }

type provider struct {
	init          chan struct{}
	initOnce      sync.Once
	notSubscribed chan struct{}

	mu              sync.RWMutex
	storageSettings map[string]StorageSettings
	fileManagerMap  map[string]func() (filemanager.FileManager, error)
}

func (p *provider) GetFileManager(ctx context.Context, workspaceID string) (filemanager.FileManager, error) {
	_ = "STUB: not implemented"
	return *new(filemanager.FileManager), nil
}

func (p *provider) GetStoragePreferences(ctx context.Context, workspaceID string) (backendconfig.StoragePreferences, error) {
	_ = "STUB: not implemented"
	return *new(backendconfig.StoragePreferences), nil
}

// blockUntilInit blocks until:
// - the provider is initialized
// - the provider is not subscribed to the backend config anymore
// - the context is done
// If it has been initialized at least once, we still check if it's not subscribed to the backend config.
// If that were to happen there might be a small chance of serving stale data.
func (p *provider) blockUntilInit(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// updateLoop uses backend config to retrieve & keep up-to-date the storage settings of all workspaces.
func (p *provider) updateLoop(ctx context.Context, backendConfig backendconfig.BackendConfig) {
	_ = "STUB: not implemented"
	return
}

// no change in workspace config, don't process the same config again

// bucket type and configuration must not be empty

// if no change in storage configuration, don't create new Filemanager

// either newly polled workspace settings or updated storage config - update object storage Filemanager

type defaultProvider struct{}

func (*defaultProvider) GetFileManager(context.Context, string) (filemanager.FileManager, error) {
	_ = "STUB: not implemented"
	return *new(filemanager.FileManager), nil
}

func (*defaultProvider) GetStoragePreferences(context.Context, string) (backendconfig.StoragePreferences, error) {
	_ = "STUB: not implemented"
	return *new(backendconfig.StoragePreferences), nil
}

func getDefaultBucket(ctx context.Context, provider string) backendconfig.StorageBucket {
	_ = "STUB: not implemented"
	return *new(backendconfig.StorageBucket)
}

func overrideWithSettings(defaultConfig map[string]any, settings backendconfig.StorageBucket, workspaceID string) backendconfig.StorageBucket {
	_ = "STUB: not implemented"
	return *new(backendconfig.StorageBucket)
}

// By default, region is set to AWS_REGION by GetProviderConfigFromEnv,
// but we remove it here to allow customers to use their own bucket
// in a different region than the default AWS_REGION
