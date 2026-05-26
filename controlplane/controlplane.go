package controlplane

import (
	"github.com/hashicorp/yamux"
	"google.golang.org/grpc"
)

type ConnHandler struct {
	GRPCServer *grpc.Server
	YamuxSess  *yamux.Session
	logger     LoggerI
}

func (cm *ConnectionManager) establishConnection() (*ConnHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ConnHandler) ServeOnConnection() error { _ = "STUB: not implemented"; return nil }

func (c *ConnHandler) Close() error { _ = "STUB: not implemented"; return nil }
