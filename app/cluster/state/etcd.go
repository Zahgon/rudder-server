package state

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/app/cluster"
	"github.com/rudderlabs/rudder-server/cluster/migrator/etcdclient"
	"github.com/rudderlabs/rudder-server/utils/types/servermode"
)

var (
	keepaliveTime    time.Duration
	keepaliveTimeout time.Duration
	dialTimeout      time.Duration
	envConfigOnce    sync.Once
)

const (
	modeRequestKeyPattern = `/%s/SERVER/%s/MODE` // /<releaseName>/server/<serverIndex>/mode

	defaultACKTimeout = 15 * time.Second
)

var _ cluster.ChangeEventProvider = &ETCDManager{}

type ETCDConfig struct {
	ReleaseName          string
	ServerIndex          string
	Endpoints            []string
	dialKeepAliveTime    time.Duration
	dialKeepAliveTimeout time.Duration
	ACKTimeout           time.Duration
	dialTimeout          time.Duration
}

type modeRequestValue struct {
	Mode   servermode.Mode `json:"mode"`
	AckKey string          `json:"ack_key"`
}

type modeAckValue struct {
	Status servermode.Mode `json:"status"`
}

func EnvETCDConfig() *ETCDConfig { _ = "STUB: not implemented"; return nil }

type ETCDManager struct {
	Config     *ETCDConfig
	Client     *clientv3.Client
	once       sync.Once
	initErr    error
	logger     logger.Logger
	ackTimeout time.Duration
}

func (manager *ETCDManager) init() error {
	manager.once.Do(func() {
		cli, err := clientv3.New(clientv3.Config{
			Endpoints:            manager.Config.Endpoints,
			DialTimeout:          manager.Config.dialTimeout,
			DialKeepAliveTime:    manager.Config.dialKeepAliveTime,
			DialKeepAliveTimeout: manager.Config.dialKeepAliveTimeout,
		})
		if err != nil {
			endpoints := strings.Join(manager.Config.Endpoints, `,`)
			manager.initErr = fmt.Errorf("etcd client connect (%q): %w", endpoints, err)
			return
		}
		manager.Client = cli
		if manager.logger == nil {
			manager.logger = logger.NewLogger().Child("etcd")
		}

		manager.ackTimeout = manager.Config.ACKTimeout
		if manager.ackTimeout == 0 {
			manager.ackTimeout = defaultACKTimeout
		}
	})

	return manager.initErr
}

// Ping ensures the connection to etcd is alive
func (manager *ETCDManager) Ping() error { _ = "STUB: not implemented"; return nil }

func (manager *ETCDManager) unmarshalMode(raw []byte) servermode.ChangeEvent {
	_ = "STUB: not implemented"
	return *new(servermode.ChangeEvent)
}

func errChModeRequest(err error) <-chan servermode.ChangeEvent {
	_ = "STUB: not implemented"
	return nil
}

func (manager *ETCDManager) ServerMode(ctx context.Context) <-chan servermode.ChangeEvent {
	_ = "STUB: not implemented"
	return nil
}

func (manager *ETCDManager) EtcdClient() (etcdclient.Client, error) {
	_ = "STUB: not implemented"
	return *new(etcdclient.Client), nil
}

func (manager *ETCDManager) Close() { _ = "STUB: not implemented"; return }

func NewETCDDynamicProvider() *ETCDManager { _ = "STUB: not implemented"; return nil }
