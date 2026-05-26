package client

import (
	"context"
	"net/http"

	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/regulation-worker/internal/model"
	"github.com/rudderlabs/rudder-server/services/controlplane/identity"
)

var pkgLogger = logger.NewLogger().Child("client")

type JobAPI struct {
	Client    *http.Client
	URLPrefix string
	Identity  identity.Identifier
}

func (j *JobAPI) URL() string { _ = "STUB: not implemented"; return "" }

// Get sends http request with ID in the url and receives a json payload
// which is decoded using schema and then mapped from schema to internal model.Job struct,
// which is actually returned.
func (j *JobAPI) Get(ctx context.Context) (model.Job, error) {
	_ = "STUB: not implemented"
	return *new(model.Job), nil
}

// if successful

// UpdateStatus marshals status into appropriate status schema, and sent as payload
// checked for returned status code.
func (j *JobAPI) UpdateStatus(ctx context.Context, status model.JobStatus, jobID int) error {
	_ = "STUB: not implemented"
	return nil
}

func mapPayloadToJob(wjs jobSchema) (model.Job, error) {
	_ = "STUB: not implemented"
	return *new(model.Job), nil
}
