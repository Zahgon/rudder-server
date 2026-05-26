package backendconfig

import (
	"context"

	"github.com/rudderlabs/rudder-server/services/controlplane/identity"
	"github.com/rudderlabs/rudder-server/utils/pubsub"
)

var _ BackendConfig = &NOOP{}

type NOOP struct{}

func (*NOOP) AccessToken() string { _ = "STUB: not implemented"; return "" }

func (*NOOP) Identity() identity.Identifier {
	_ = "STUB: not implemented"
	return *new(identity.Identifier)
}

func (*NOOP) SetUp() error { _ = "STUB: not implemented"; return nil }

func (*NOOP) Get(_ context.Context) (map[string]ConfigT, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*NOOP) WaitForConfig(_ context.Context) { _ = "STUB: not implemented"; return }

func (*NOOP) Subscribe(ctx context.Context, _ Topic) pubsub.DataChannel {
	_ = "STUB: not implemented"
	return *new(pubsub.DataChannel)
}

func (*NOOP) StartWithIDs(_ context.Context, _ string) { _ = "STUB: not implemented"; return }

func (*NOOP) Stop() { _ = "STUB: not implemented"; return }
