package backendconfigtest

import (
	"encoding/json"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

// NewSourceBuilder returns a new SourceBuilder
func NewSourceBuilder() *SourceBuilder { _ = "STUB: not implemented"; return nil }

// SourceBuilder is a builder for a source
type SourceBuilder struct {
	valueBuilder[backendconfig.SourceT]
}

// WithID sets the ID of the source
func (b *SourceBuilder) WithID(id string) *SourceBuilder { _ = "STUB: not implemented"; return nil }

// WithWriteKey sets the write key of the source
func (b *SourceBuilder) WithWriteKey(writeKey string) *SourceBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithConfigOption sets a config option for the source
func (b *SourceBuilder) WithConfigOption(key string, value any) (*SourceBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WithConnection adds a destination to the source
func (b *SourceBuilder) WithConnection(destination backendconfig.DestinationT) *SourceBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Disabled disables the source
func (b *SourceBuilder) Disabled() *SourceBuilder { _ = "STUB: not implemented"; return nil }

// WithTrackingPlan adds a tracking plan to the source
func (b *SourceBuilder) WithTrackingPlan(id string, version int) *SourceBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithGeoenrichmentEnabled enables geoenrichment for the source
func (b *SourceBuilder) WithGeoenrichmentEnabled(enabled bool) *SourceBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithSourceCategory sets the source definition category
func (b *SourceBuilder) WithSourceCategory(category string) *SourceBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithSourceDefOptions sets the source definition options
func (b *SourceBuilder) WithSourceDefOptions(opts backendconfig.SourceDefinitionOptions) *SourceBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithSourceType sets the source type
func (b *SourceBuilder) WithSourceType(sourceType string) *SourceBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithSourceType sets the source type
func (b *SourceBuilder) WithWorkspaceID(workspaceID string) *SourceBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithInternalSecrets sets the internal secrets for the source
func (b *SourceBuilder) WithInternalSecrets(secrets json.RawMessage) *SourceBuilder {
	_ = "STUB: not implemented"
	return nil
}
