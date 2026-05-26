package klaviyobulkupload

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/common"
)

const (
	BATCHSIZE             = 10000
	MAXALLOWEDPROFILESIZE = 512000
	MAXPAYLOADSIZE        = 4600000
	IMPORT_ID_SEPARATOR   = ":"
)

func createFinalPayload(combinedProfiles []Profile, listId string) Payload {
	_ = "STUB: not implemented"
	return *new(Payload)
}

func NewManager(logger logger.Logger, StatsFactory stats.Stats, destination *backendconfig.DestinationT) (*KlaviyoBulkUploader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func chunkBySizeAndElements(combinedProfiles []Profile, jobIDs []int64, maxBytes, maxElements int) ([][]Profile, [][]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// +1 for comma character

func (kbu *KlaviyoBulkUploader) Poll(_ context.Context, pollInput common.AsyncPoll) common.PollStatusResponse {
	_ = "STUB: not implemented"
	return *new(common.PollStatusResponse)
}

// Update the status in the map

// If Failed_count > 0, add the importId to failedImports

func (kbu *KlaviyoBulkUploader) GetUploadStats(UploadStatsInput common.GetUploadStatsInput) common.GetUploadStatsResponse {
	_ = "STUB: not implemented"
	return *new(common.GetUploadStatsResponse)
}

// make a map of jobId to error reason

// Iterate over the Data array and get the jobId and error detail and store in jobIdToErrorMap

func (kbu *KlaviyoBulkUploader) generateKlaviyoErrorOutput(errorString string, err error, importingJobIds []int64, destinationID string) common.AsyncUploadOutput {
	_ = "STUB: not implemented"
	return *new(common.AsyncUploadOutput)
}

func (kbu *KlaviyoBulkUploader) ExtractProfile(Data Data) Profile {
	_ = "STUB: not implemented"
	return *new(Profile)
}

// delete jobIdentifier from the attributes map as it is not required in the final payload

func (kbu *KlaviyoBulkUploader) Upload(_ context.Context, asyncDestStruct *common.AsyncDestinationStruct) common.AsyncUploadOutput {
	_ = "STUB: not implemented"
	return *new(common.AsyncUploadOutput)
}

// Track job IDs for each profile

// if profileStructure length is more than 500 kB, throw an error

// Record the size in the histogram

// DelimitedImportIds is : separated importIds

// Mark all job IDs in this profile chunk as failed

func (kbu *KlaviyoBulkUploader) Transform(job *jobsdb.JobT) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
