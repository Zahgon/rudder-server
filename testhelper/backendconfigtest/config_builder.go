package backendconfigtest

import (
	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

// NewConfigBuilder returns a new ConfigBuilder
func NewConfigBuilder() *ConfigBuilder { _ = "STUB: not implemented"; return nil }

// ConfigBuilder is a builder for a backend config
type ConfigBuilder struct {
	valueBuilder[backendconfig.ConfigT]
}

// WithSource adds a source to the config
func (b *ConfigBuilder) WithSource(source backendconfig.SourceT) *ConfigBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithWorkspaceID adds a workspaceID to the config
func (b *ConfigBuilder) WithWorkspaceID(workspaceID string) *ConfigBuilder {
	_ = "STUB: not implemented"
	return nil
}
