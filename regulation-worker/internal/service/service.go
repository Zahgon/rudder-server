package service

import (
	"context"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/regulation-worker/internal/model"
)

//go:generate mockgen -source=service.go -destination=mock_service.go -package=service github.com/rudderlabs/rudder-server/regulation-worker/internal/service
type APIClient interface {
	Get(ctx context.Context) (model.Job, error)
	UpdateStatus(ctx context.Context, status model.JobStatus, jobID int) error
}

type dest interface {
	GetDestination(destID string) (*backendconfig.DestinationT, error)
}
type deleter interface {
	Delete(ctx context.Context, job model.Job, dest *backendconfig.DestinationT) model.JobStatus
}

type JobSvc struct {
	API               APIClient
	Deleter           deleter
	Destination       dest
	MaxFailedAttempts int
}

// JobSvc called by looper
// calls api-client.getJob(workspaceID)
// calls api-client to get new job with workspaceID, which returns jobID.
func (js *JobSvc) JobSvc(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// API request to get new job

// once job is successfully received, calling update-status API to update the status of job to running.

// executing deletion

func (js *JobSvc) updateStatus(ctx context.Context, status model.JobStatus, jobID int) error {
	_ = "STUB: not implemented"
	return nil
}
