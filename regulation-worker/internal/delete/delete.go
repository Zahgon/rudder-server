package delete

import (
	"context"
	"sync"

	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/regulation-worker/internal/model"
)

var pkgLogger = logger.NewLogger().Child("client")

//go:generate mockgen -source=delete.go -destination=mock_delete_test.go -package=delete github.com/rudderlabs/rudder-server/regulation-worker/internal/delete
type deleteManager interface {
	Delete(ctx context.Context, job model.Job, destination *backendconfig.DestinationT) model.JobStatus
	GetSupportedDestinations() []string
}

type Router struct {
	Managers []deleteManager
	router   map[string]deleteManager
	once     sync.Once
}

func NewRouter(managers ...deleteManager) *Router { _ = "STUB: not implemented"; return nil }

func (r *Router) Delete(ctx context.Context, job model.Job, dest *backendconfig.DestinationT) model.JobStatus {
	_ = "STUB: not implemented"
	return *new(model.JobStatus)
}
