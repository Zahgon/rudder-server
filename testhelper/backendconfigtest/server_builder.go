package backendconfigtest

import (
	"net/http"

	"github.com/rudderlabs/rudder-go-kit/testhelper/httptest"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

// NewBuilder returns a new ServerBuilder
func NewBuilder() *ServerBuilder { _ = "STUB: not implemented"; return nil }

// ServerBuilder is a builder for a test server that returns backend configs
type ServerBuilder struct {
	namespace       string
	configs         map[string]backendconfig.ConfigT
	settingsHandler http.HandlerFunc
}

// WithNamespace sets the namespace for the server along with the configs for that namespace
func (b *ServerBuilder) WithNamespace(namespace string, configs ...backendconfig.ConfigT) *ServerBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithWorkspaceConfig sets the workspace config for the server
func (b *ServerBuilder) WithWorkspaceConfig(config backendconfig.ConfigT) *ServerBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Build builds the test server
func (b *ServerBuilder) Build() *httptest.Server { _ = "STUB: not implemented"; return nil }
