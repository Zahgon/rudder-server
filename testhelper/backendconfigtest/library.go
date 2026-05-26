package backendconfigtest

import (
	"context"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/controlplane/identity"
	"github.com/rudderlabs/rudder-server/utils/pubsub"
)

var _ backendconfig.BackendConfig = &StaticLibrary{}

func NewStaticLibrary(configs map[string]backendconfig.ConfigT) *StaticLibrary {
	_ = "STUB: not implemented"
	return nil
}

type StaticLibrary struct {
	configs map[string]backendconfig.ConfigT
}

// AccessToken returns the access token for the backend config
func (l *StaticLibrary) AccessToken() string { _ = "STUB: not implemented"; return "" }

// TODO: Implement

func (l *StaticLibrary) Get(context.Context) (map[string]backendconfig.ConfigT, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *StaticLibrary) Identity() identity.Identifier {
	_ = "STUB: not implemented"
	return *
	// TODO: Implement
	new(identity.Identifier)
}

func (l *StaticLibrary) Subscribe(ctx context.Context, topic backendconfig.Topic) pubsub.DataChannel {
	_ = "STUB: not implemented"
	return *new(pubsub.DataChannel)
}

// on Subscribe, emulate a single backend configuration event

func (l *StaticLibrary) SetUp() error { _ = "STUB: not implemented"; return nil }

func (l *StaticLibrary) WaitForConfig(context.Context) { _ = "STUB: not implemented"; return }

func (l *StaticLibrary) Stop() { _ = "STUB: not implemented"; return }

func (l *StaticLibrary) StartWithIDs(context.Context, string) { _ = "STUB: not implemented"; return }
