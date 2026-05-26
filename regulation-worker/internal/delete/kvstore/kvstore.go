package kvstore

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/regulation-worker/internal/model"
)

var (
	supportedDestinations = []string{"REDIS"}
	pkgLogger             = logger.NewLogger().Child("kvstore")
)

type KVDeleteManager struct{}

func (*KVDeleteManager) GetSupportedDestinations() []string { _ = "STUB: not implemented"; return nil }

func (*KVDeleteManager) Delete(_ context.Context, job model.Job, destination *backendconfig.DestinationT) model.JobStatus {
	_ = "STUB: not implemented"
	return *new(model.JobStatus)
}
