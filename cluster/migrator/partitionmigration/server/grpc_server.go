// GRPC server factory for partition migration
package server

import (
	"sync"

	"google.golang.org/grpc"

	"github.com/rudderlabs/rudder-go-kit/config"

	proto "github.com/rudderlabs/rudder-server/proto/cluster"
)

// NewGRPCServer creates a new GRPCServer instance
func NewGRPCServer(conf *config.Config, pms proto.PartitionMigrationServer) *GRPCServer {
	_ = "STUB: not implemented"
	return nil
}

type GRPCServer struct {
	conf *config.Config
	pms  proto.PartitionMigrationServer
	wg   sync.WaitGroup

	server *grpc.Server
}

// Start creates the gRPC server and starts listening for incoming connections
func (s *GRPCServer) Start() error { _ = "STUB: not implemented"; return nil }

// This shouldn't really happen, only in very exceptional cases.
// No error is returned during GracefulStop or Stop.

// Stop gracefully stops the gRPC server with a timeout
func (s *GRPCServer) Stop() { _ = "STUB: not implemented"; return }

// Graceful stop completed

// Timeout exceeded, force stop

// newGrpcServer creates and configures a new gRPC server instance
func (s *GRPCServer) newGrpcServer() *grpc.Server { _ = "STUB: not implemented"; return nil }

// Keepalive enforcement policy - controls what the server requires from clients

// If client pings more often than every MinTime, terminate the connection

// Allow client to ping even if there are no active streams

// Keepalive parameters - controls server's own keepalive behavior

// If a client is idle for MaxConnectionIdle, send a GOAWAY

// Ping the client if no data is received for Time duration

// Wait for Timeout for the ping ack before assuming the connection is dead
