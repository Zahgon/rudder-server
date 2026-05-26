package multitenant

import (
	"context"
	"sync"

	"github.com/rudderlabs/rudder-go-kit/config"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

type Manager struct {
	backendConfig        backendconfig.BackendConfig
	degradedWorkspaceIDs []string

	sourceIDToWorkspaceID map[string]string
	excludeWorkspaceIDMap map[string]struct{}

	ready     chan struct{}
	sourceMu  sync.Mutex
	readyOnce sync.Once
	initOnce  sync.Once
}

func New(conf *config.Config, bcConfig backendconfig.BackendConfig) *Manager {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) init() {
	m.initOnce.Do(func() {
		m.sourceIDToWorkspaceID = make(map[string]string)
		m.excludeWorkspaceIDMap = make(map[string]struct{})

		for _, workspaceID := range m.degradedWorkspaceIDs {
			m.excludeWorkspaceIDMap[workspaceID] = struct{}{}
		}
		m.ready = make(chan struct{})
	})
}

// Run is a blocking function that executes manager background logic.
func (m *Manager) Run(ctx context.Context) { _ = "STUB: not implemented"; return }

// DegradedWorkspace returns true if the workspaceID is degraded.
func (m *Manager) DegradedWorkspace(workspaceID string) bool {
	_ = "STUB: not implemented"
	return false
}

// DegradedWorkspaces returns a list of degraded workspaceIDs.
func (m *Manager) DegradedWorkspaces() []string { _ = "STUB: not implemented"; return nil }

// SourceToWorkspace returns the workspaceID for a given sourceID, even if workspaceID is degraded.
// An error is returned if the sourceID is not found, or context is canceled.
//
//	NOTE: This function blocks until the backend config is loaded.
func (m *Manager) SourceToWorkspace(ctx context.Context, sourceID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// WatchConfig returns a backend config map that excludes degraded workspaces.
//
// NOTE: WatchConfig is responsible for closing the channel when context gets cancel.
func (m *Manager) WatchConfig(ctx context.Context) <-chan map[string]backendconfig.ConfigT {
	_ = "STUB: not implemented"
	return nil
}
