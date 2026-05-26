package marketobulkupload

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/common"
)

type MarketoBulkUploader struct {
	destName          string
	destinationConfig MarketoConfig
	logger            logger.Logger
	statsFactory      stats.Stats
	csvHeaders        []string
	dataHashToJobId   map[string]int64
	hasFailures       bool
	hasWarning        bool
	apiService        MarketoAPIServiceInterface
}

type MarketoAsyncFailedInput struct {
	Message  map[string]any
	Metadata struct {
		JobID int64
	}
}

type MarketoAsyncFailedPayload struct {
	Config   map[string]any
	Input    []MarketoAsyncFailedInput
	DestType string
	ImportId string
	MetaData common.MetaDataT
}

const (
	MARKETO_WARNING_HEADER = "Import Warning Reason"
	MARKETO_FAILED_HEADER  = "Import Failure Reason"
)

func (b *MarketoBulkUploader) Upload(_ context.Context, asyncDestStruct *common.AsyncDestinationStruct) common.AsyncUploadOutput {
	_ = "STUB: not implemented"
	return *new(common.AsyncUploadOutput)
}

// Clean up the temporary file

// Check file size

// return the response

func (b *MarketoBulkUploader) Poll(_ context.Context, pollInput common.AsyncPoll) common.PollStatusResponse {
	_ = "STUB: not implemented"
	return *new(common.PollStatusResponse)
}

// reflection on marketoResponse might not be a cheap operation

// nolint:forbidigo

// Check if the response is empty

// Set State

// in case of success, clear the hashToJobId map

func (b *MarketoBulkUploader) GetUploadStats(input common.GetUploadStatsInput) common.GetUploadStatsResponse {
	_ = "STUB: not implemented"
	// Extract importId from parameters
	return *new(common.GetUploadStatsResponse)
}

// Fetch and parse failed jobs

// Fetch and parse warning jobs

func (*MarketoBulkUploader) Transform(job *jobsdb.JobT) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (b *MarketoBulkUploader) updateJobStatus(importingList []*jobsdb.JobT, failedJobs, warningJobs []map[string]string) common.EventStatMeta {
	_ = "STUB: not implemented"
	return *new(common.EventStatMeta)
}

// get failedJob data

// get jobID from jobToDataHash

// get warningJob data

// get jobID from jobToDataHash

// Even if a job has warning, it is considered as a failure

// calculate succeeded keys

func (b *MarketoBulkUploader) clearHashToJobId() { _ = "STUB: not implemented"; return }
