package lyticsBulkUpload

import (
	"context"

	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/common"
)

type LyticsServiceImpl struct {
	BulkApi string
}

func (u *LyticsServiceImpl) getBulkApi(destConfig DestinationConfig) *LyticsServiceImpl {
	_ = "STUB: not implemented"
	return nil
}

func (*LyticsBulkUploader) Transform(job *jobsdb.JobT) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (u *LyticsServiceImpl) MakeHTTPRequest(data *HttpRequestData) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (u *LyticsServiceImpl) UploadBulkFile(data *HttpRequestData, filePath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *LyticsBulkUploader) Upload(_ context.Context, asyncDestStruct *common.AsyncDestinationStruct) common.AsyncUploadOutput {
	_ = "STUB: not implemented"
	return *new(common.AsyncUploadOutput)
}

// remove the file that could not be uploaded
