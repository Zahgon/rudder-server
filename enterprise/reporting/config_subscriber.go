package reporting

import (
	"context"
	"sync"

	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

type destDetail struct {
	destinationDefinitionID string
	destType                string // destination definition name
}

type configSubscriber struct {
	init     chan struct{}
	onceInit sync.Once

	log logger.Logger

	backendConfigMu           sync.RWMutex // protects the following
	workspaceID               string
	workspaceIDForSourceIDMap map[string]string
	destinationIDMap          map[string]destDetail
	piiReportingSettings      map[string]bool
}

func newConfigSubscriber(log logger.Logger) *configSubscriber {
	_ = "STUB: not implemented"
	return nil
}

func (cs *configSubscriber) Subscribe(
	ctx context.Context,
	bcConfig backendconfig.BackendConfig,
) {
	_ = "STUB: not implemented"
	return
}

// Reduce to destination detail based on destinationID

// Wait blocks until the config subscriber is initialized
func (cs *configSubscriber) Wait() {
	_ = "STUB: not implemented"

	// WorkspaceID returns the workspace ID
	return
}

func (cs *configSubscriber) WorkspaceID() string { _ = "STUB: not implemented"; return "" }

// WorkspaceIDFromSource returns the workspace ID for the given source ID
func (cs *configSubscriber) WorkspaceIDFromSource(sourceID string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetDestDetail returns the destination detail for the given destination ID
func (cs *configSubscriber) GetDestDetail(destID string) destDetail {
	_ = "STUB: not implemented"
	return *new(destDetail)
}

// IsPIIReportingDisabled returns true if PII reporting is disabled for the given workspace
func (cs *configSubscriber) IsPIIReportingDisabled(workspaceID string) bool {
	_ = "STUB: not implemented"
	return false
}
