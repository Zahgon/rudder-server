package eloqua

import (
	"context"

	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/common"
)

func (b *EloquaBulkUploader) createAsyncUploadErrorOutput(errorString string, err error, destinationId string, asyncDestStruct *common.AsyncDestinationStruct) common.AsyncUploadOutput {
	_ = "STUB: not implemented"
	return *new(common.AsyncUploadOutput)
}

func (*EloquaBulkUploader) Transform(job *jobsdb.JobT) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (b *EloquaBulkUploader) Upload(_ context.Context, asyncDestStruct *common.AsyncDestinationStruct) common.AsyncUploadOutput {
	_ = "STUB: not implemented"
	return *new(common.AsyncUploadOutput)
}

func (b *EloquaBulkUploader) Poll(_ context.Context, pollInput common.AsyncPoll) common.PollStatusResponse {
	_ = "STUB: not implemented"
	return *new(common.PollStatusResponse)
}

func (b *EloquaBulkUploader) GetUploadStats(UploadStatsInput common.GetUploadStatsInput) common.GetUploadStatsResponse {
	_ = "STUB: not implemented"
	return *new(common.GetUploadStatsResponse)
}

// Deletes import definition from Eloqua itself
func (b *EloquaBulkUploader) deleteImportDef(importDefId string) { _ = "STUB: not implemented"; return }

func (b *EloquaBulkUploader) clearJobToCsvMap() { _ = "STUB: not implemented"; return }
