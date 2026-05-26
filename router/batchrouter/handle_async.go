package batchrouter

import (
	"context"
	stdjson "encoding/json"
	"time"

	"github.com/rudderlabs/rudder-server/jobsdb"
	common "github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/common"
	utilTypes "github.com/rudderlabs/rudder-server/utils/types"
)

func (brt *Handle) getImportingJobs(ctx context.Context, augmentQueryParams func(*jobsdb.GetQueryParams), limit int) (jobsdb.JobsResult, error) {
	_ = "STUB: not implemented"
	return *new(jobsdb.JobsResult), nil
}

// we need to get all importing jobs based on limit, overcoming dsLimits and payload size limits

// stop when we have enough jobs
// or we are confident there are no more jobs to fetch
// or we have reached max iterations

func (brt *Handle) updateJobStatuses(ctx context.Context, allJobs, completedJobs []*jobsdb.JobT, statusList []*jobsdb.JobStatusT) error {
	_ = "STUB: not implemented"
	return nil
}

// rsources stats

func getPollInput(job *jobsdb.JobT) common.AsyncPoll {
	_ = "STUB: not implemented"
	return *new(common.AsyncPoll)
}

func enhanceResponseWithFirstAttemptedAt(msg stdjson.RawMessage, resp []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func getFirstAttemptAtFromErrorResponse(msg stdjson.RawMessage) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (brt *Handle) prepareJobStatusList(importingList []*jobsdb.JobT, defaultStatus jobsdb.JobStatusT, sourceID, destinationID string) ([]*jobsdb.JobStatusT, []*jobsdb.JobT, map[int64]jobsdb.ConnectionDetails) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (brt *Handle) getParamertsFromJobs(jobs []*jobsdb.JobT) map[int64]stdjson.RawMessage {
	_ = "STUB: not implemented"
	return nil
}

func (brt *Handle) updatePollStatusToDB(ctx context.Context, destinationID, sourceID string, importingJob *jobsdb.JobT, importingCount int, pollResp common.PollStatusResponse) ([]*jobsdb.JobStatusT, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (brt *Handle) pollAsyncStatus(ctx context.Context) { _ = "STUB: not implemented"; return }

// TODO: Add metrics

// fallback to maxEventsInABatch if import count is not set

// Log a warning if there is a mismatch in the lengths

func (brt *Handle) asyncUploadWorker(ctx context.Context) { _ = "STUB: not implemented"; return }

func (brt *Handle) asyncStructSetup(sourceID, destinationID string, jobsList []*jobsdb.JobT) {
	_ = "STUB: not implemented"
	return
}

func (brt *Handle) asyncStructCleanUp(destinationID string) { _ = "STUB: not implemented"; return }

func (brt *Handle) sendJobsToStorage(batchJobs BatchedJobs) error {
	_ = "STUB: not implemented"
	return nil
}

// skipcq: GO-R4002

// Waiting for previous upload to complete, mark all jobs as failed

// mark overflown jobs as failed

// turn on CanUpload flag to true

func (brt *Handle) createFakeJob(jobID int64, parameters stdjson.RawMessage) *jobsdb.JobT {
	_ = "STUB: not implemented"
	return nil
}

func (brt *Handle) getReportMetrics(params getReportMetricsParams) []*utilTypes.PUReportedMetric {
	_ = "STUB: not implemented"
	return nil
}

func (brt *Handle) setMultipleJobStatus(params setMultipleJobStatusParams) {
	_ = "STUB: not implemented"
	return
}

// synthesize job parameters with the destination ID in case of missing job parameters

// Debugging negative pending events count issue

// Mark the status of the jobs

// rsources stats

func (brt *Handle) GetWorkspaceIDForDestID(destID string) string {
	_ = "STUB: not implemented"
	return ""
}
