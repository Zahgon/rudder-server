package backendconfigtest

import (
	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

// NewDestinationBuilder returns a new DestinationBuilder
func NewDestinationBuilder(destType string) *DestinationBuilder {
	_ = "STUB: not implemented"
	return nil
}

// DestinationBuilder is a builder for a destination
type DestinationBuilder struct {
	valueBuilder[backendconfig.DestinationT]
}

// WithID sets the ID of the destination
func (b *DestinationBuilder) WithID(id string) *DestinationBuilder {
	_ = "STUB: not implemented"
	return nil

	// WithRevisionID sets the revision ID of the destination
}

func (b *DestinationBuilder) WithRevisionID(revisionID string) *DestinationBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithConfigOption sets a config option for the destination
func (b *DestinationBuilder) WithConfigOption(key string, value any) *DestinationBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithUserTransformation adds a user transformation to the destination
func (b *DestinationBuilder) WithUserTransformation(id, version string) *DestinationBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithDefinitionConfigOption adds a config option to the destination definition
func (b *DestinationBuilder) WithDefinitionConfigOption(key string, value any) *DestinationBuilder {
	_ = "STUB: not implemented"
	return nil
}
