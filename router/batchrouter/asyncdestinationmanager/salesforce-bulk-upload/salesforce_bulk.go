package salesforcebulkupload

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/common"
)

func NewManager(
	conf *config.Config,
	logger logger.Logger,
	statsFactory stats.Stats,
	destination *backendconfig.DestinationT,
	backendConfig backendconfig.BackendConfig,
) (common.AsyncDestinationManager, error) {
	_ = "STUB: not implemented"
	return *new(common.AsyncDestinationManager), nil
}

func NewUploader(
	conf *config.Config,
	logger logger.Logger,
	statsFactory stats.Stats,
	apiService APIServiceInterface,
	destination *backendconfig.DestinationT,
) *Uploader {
	_ = "STUB: not implemented"
	return nil
}

func (s *Uploader) Transform(job *jobsdb.JobT) (string, error) {
	_ = "STUB: not implemented"
	// Extract required fields from the input
	return "", nil
}

// Build the metadata object

// We are supporting only upsert operation

// Add externalId to metadata if it exists

// Build the message object

// Add all traits to message

// Add externalId field to message

// Create the final AsyncJob structure

// Marshal and return

func (s *Uploader) readJobsFromFile(filePath string) ([]common.AsyncJob, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Uploader) Upload(_ context.Context, asyncDestStruct *common.AsyncDestinationStruct) common.AsyncUploadOutput {
	_ = "STUB: not implemented"
	return *new(common.AsyncUploadOutput)
}

func (s *Uploader) Poll(_ context.Context, pollInput common.AsyncPoll) common.PollStatusResponse {
	_ = "STUB: not implemented"
	return *new(common.PollStatusResponse)
}

func (s *Uploader) GetUploadStats(input common.GetUploadStatsInput) common.GetUploadStatsResponse {
	_ = "STUB: not implemented"
	return *new(common.GetUploadStatsResponse)
}

func (s *Uploader) handlePollError(apiError *APIError) common.PollStatusResponse {
	_ = "STUB: not implemented"
	return *new(common.PollStatusResponse)
}

func (s *Uploader) matchRecordsToJobs(
	importingList []*jobsdb.JobT,
	failedRecords, successRecords []map[string]string,
	headers []string,
) common.EventStatMeta {
	_ = "STUB: not implemented"
	return *new(common.EventStatMeta)
}

func (s *Uploader) clearDataHashToJobID() { _ = "STUB: not implemented"; return }
