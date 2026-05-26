package http

import (
	"net/http"

	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/services/rsources"
)

// NewV1Handler is a legacy handler for job status
//
//   - GET /v1/job-status/{job_run_id} - returns job status
//   - DELETE /v1/job-status/{job_run_id} - deletes job status AND failed records
//   - GET /v1/job-status/{job_run_id}/failed-records - returns failed records
//
// TODO: delete this handler once we remove support for the v1 api
func NewV1Handler(service rsources.JobService, logger logger.Logger) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// NewFailedKeysHandler creates a handler for failed keys
//
//   - GET /v2/job-status/{job_run_id}
//   - DELETE /v2/job-status/{job_run_id}
//   - GET /v2/failed-keys/{job_run_id}
//   - DELETE /v2/failed-keys/{job_run_id}
func NewV2Handler(service rsources.JobService, logger logger.Logger) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type handler struct {
	logger  logger.Logger
	service rsources.JobService
}

func (h *handler) delete(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func (h *handler) getStatus(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) deleteJobStatus(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) failedRecordsV1(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) failedRecords(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) deleteFailedRecords(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func getQueryParams(r *http.Request) (jobRunID string, taskRunID, sourceID []string) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func marshalAndWriteResponse(w http.ResponseWriter, response any) (err error) {
	_ = "STUB: not implemented"
	return nil
}
