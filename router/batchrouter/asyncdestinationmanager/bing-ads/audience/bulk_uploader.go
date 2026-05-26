package audience

import (
	"context"

	"github.com/rudderlabs/bing-ads-go-sdk/bingads"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/common"
)

func NewBingAdsBulkUploader(logger logger.Logger, statsFactory stats.Stats, destName string, service bingads.BulkServiceI, client *Client) *BingAdsBulkUploader {
	_ = "STUB: not implemented"
	return nil
}

func (*BingAdsBulkUploader) Transform(job *jobsdb.JobT) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

/*
This function create at most 3 zip files from the text file created by the batchrouter
It takes the text file path as input and returns the zip file path
The maximum size of the zip file is 100MB, if the size of the zip file exceeds 100MB then the job is marked as failed
*/
func (b *BingAdsBulkUploader) Upload(_ context.Context, asyncDestStruct *common.AsyncDestinationStruct) common.AsyncUploadOutput {
	_ = "STUB: not implemented"
	return *new(common.AsyncUploadOutput)
}

// remove the file that could not be uploaded

func (b *BingAdsBulkUploader) pollSingleImport(requestId string) common.PollStatusResponse {
	_ = "STUB: not implemented"
	return *new(common.PollStatusResponse)
}

func (b *BingAdsBulkUploader) Poll(_ context.Context, pollInput common.AsyncPoll) common.PollStatusResponse {
	_ = "STUB: not implemented"
	return *new(common.PollStatusResponse)
}

/*
	Cumulative Response logic:
	1. If any of the request is in progress then the whole request is in progress and it should retry
	2. if all the requests are completed then the whole request is completed
	3. if any of the requests are completed with errors then the whole request is completed with errors
	4. if any of the requests are failed then the whole request is failed and retried
	5. if all the requests are failed then the whole request is failed
*/

// creating a comma separated string of all the result file urls

func (b *BingAdsBulkUploader) getUploadStatsOfSingleImport(filePath string) (common.GetUploadStatsResponse, error) {
	_ = "STUB: not implemented"
	return *new(common.GetUploadStatsResponse), nil
}

func (b *BingAdsBulkUploader) GetUploadStats(uploadStatsInput common.GetUploadStatsInput) common.GetUploadStatsResponse {
	_ = "STUB: not implemented"
	// considering importing jobs are the primary list of jobs sent
	// making an array of those jobIds
	return *new(common.GetUploadStatsResponse)
}

// only one file should be there
