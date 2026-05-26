package state

import (
	"context"

	"github.com/rudderlabs/rudder-server/app/cluster"
	"github.com/rudderlabs/rudder-server/cluster/migrator/etcdclient"
	"github.com/rudderlabs/rudder-server/utils/types/servermode"
)

var _ cluster.ChangeEventProvider = &StaticProvider{}

type StaticProvider struct {
	mode servermode.Mode
}

func NewStaticProvider(Mode servermode.Mode) *StaticProvider { _ = "STUB: not implemented"; return nil }

// ServerMode returns a channel with a single message containing this static provider's mode.
func (s *StaticProvider) ServerMode(ctx context.Context) <-chan servermode.ChangeEvent {
	_ = "STUB: not implemented"
	return nil
}

func (s *StaticProvider) EtcdClient() (etcdclient.Client, error) {
	_ = "STUB: not implemented"
	return *new(etcdclient.Client), nil
}
