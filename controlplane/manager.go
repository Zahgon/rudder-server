package controlplane

import (
	"sync"
	"time"

	"google.golang.org/grpc"
)

type ConnectionManager struct {
	AuthInfo        AuthInfo
	RegisterService func(*grpc.Server)
	RetryInterval   time.Duration
	UseTLS          bool
	Logger          LoggerI
	mu              sync.Mutex
	active          bool
	url             string
	connHandler     *ConnHandler
	Options         []grpc.ServerOption
}

type LoggerI interface {
	Warn(a ...any)
	Warnf(format string, a ...any)
	Info(a ...any)
	Infof(format string, a ...any)
	Error(a ...any)
	Errorf(format string, a ...any)
}

const defaultRetryInterval time.Duration = time.Second

func (cm *ConnectionManager) Apply(url string, active bool) { _ = "STUB: not implemented"; return }

func (cm *ConnectionManager) connect() error { _ = "STUB: not implemented"; return nil }

func (cm *ConnectionManager) maintainConnection() { _ = "STUB: not implemented"; return }

func (cm *ConnectionManager) closeConnection() error { _ = "STUB: not implemented"; return nil }

func (cm *ConnectionManager) retryInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
