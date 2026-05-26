package backendconfig

import (
	"context"
	"net/url"
	"sync"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/backend-config/dynamicconfig"
	"github.com/rudderlabs/rudder-server/services/controlplane/identity"
	"github.com/rudderlabs/rudder-server/utils/types"
)

type singleWorkspaceConfig struct {
	token            string
	configBackendURL *url.URL
	configJSONPath   string
	configEnvHandler types.ConfigEnvI
	region           string

	workspaceIDOnce sync.Once
	workspaceID     string

	dynamicConfigCache dynamicconfig.Cache

	logger               logger.Logger
	httpCallsStat        stats.Counter
	httpResponseSizeStat stats.Histogram
}

func (wc *singleWorkspaceConfig) SetUp() error { _ = "STUB: not implemented"; return nil }

func (wc *singleWorkspaceConfig) AccessToken() string {
	_ = "STUB: not implemented"

	// Get returns sources from the workspace
	return ""
}

func (wc *singleWorkspaceConfig) Get(ctx context.Context) (map[string]ConfigT, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getFromApi gets the workspace config from api
func (wc *singleWorkspaceConfig) getFromAPI(ctx context.Context) (map[string]ConfigT, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Process dynamic config with the instance cache

// getFromFile reads the workspace config from JSON file
func (wc *singleWorkspaceConfig) getFromFile() (map[string]ConfigT, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Process dynamic config with the instance cache

func (wc *singleWorkspaceConfig) makeHTTPRequest(ctx context.Context, url string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (wc *singleWorkspaceConfig) Identity() identity.Identifier {
	_ = "STUB: not implemented"
	return *new(identity.Identifier)
}
