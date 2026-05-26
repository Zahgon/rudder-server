package backendconfig

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/backend-config/dynamicconfig"
	"github.com/rudderlabs/rudder-server/services/controlplane/identity"
	"github.com/rudderlabs/rudder-server/utils/types"
)

var (
	updatedAfterTimeFormat     = "2006-01-02T15:04:05.000Z"
	ErrIncrementalUpdateFailed = errors.New("incremental update failed")
)

type namespaceConfig struct {
	configEnvHandler types.ConfigEnvI
	cpRouterURL      string

	config *config.Config
	logger logger.Logger
	client *http.Client

	hostedServiceSecret string

	namespace                string
	configBackendURL         *url.URL
	region                   string
	incrementalConfigUpdates bool
	lastUpdatedAt            time.Time
	workspacesConfig         map[string]ConfigT
	dynamicConfigCache       dynamicconfig.Cache

	httpCallsStat        stats.Counter
	httpResponseSizeStat stats.Histogram
}

func (nc *namespaceConfig) SetUp() (err error) { _ = "STUB: not implemented"; return nil }

// Get returns sources from the workspace
func (nc *namespaceConfig) Get(ctx context.Context) (map[string]ConfigT, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// reset state here
// this triggers a full update

// getFromApi gets the workspace config from api
func (nc *namespaceConfig) getFromAPI(ctx context.Context) (map[string]ConfigT, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this workspace was not updated, populate it with the previous config

// Process dynamic config with the instance cache

// always set connection flags to true for hosted and multi-tenant warehouse service

func (nc *namespaceConfig) prepareHTTPRequest(ctx context.Context, url string) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (nc *namespaceConfig) makeHTTPRequest(req *http.Request) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (nc *namespaceConfig) AccessToken() string { _ = "STUB: not implemented"; return "" }

func (nc *namespaceConfig) Identity() identity.Identifier {
	_ = "STUB: not implemented"
	return *new(identity.Identifier)
}
