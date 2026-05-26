package rudderserver

import (
	"context"
	"testing"

	"golang.org/x/sync/errgroup"
)

// BuildRudderServerBinary builds the rudder-server binary and returns its path.
func BuildRudderServerBinary(t *testing.T, mainPath, binaryPath string) {
	_ = "STUB: not implemented"
	return
}

// StartRudderServer starts a rudder-server process with the given environment configuration in a separate goroutine managed by the provided errgroup.Group.
// It sends SIGTERM when the context is cancelled to allow graceful shutdown and proper coverage data collection.
func StartRudderServer(t *testing.T, ctx context.Context, g *errgroup.Group, name, binaryPath string, configs map[string]string, otherEnv ...string) {
	_ = "STUB: not implemented"
	return
}

// Don't use exec.CommandContext - it sends SIGKILL on context cancellation,
// which prevents the process from writing coverage data.

// Wait for context cancellation, then send SIGTERM for graceful shutdown

func convertCoverageData(t *testing.T, coverDir, outputFile string) {
	_ = "STUB: not implemented"
	return
}
