package common

import (
	"context"

	"github.com/rudderlabs/rudder-server/jobsdb"
)

type InvalidManager struct {
	Error error
}

func (*InvalidManager) Transform(job *jobsdb.JobT) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (*InvalidManager) Upload(_ context.Context, asyncDestStruct *AsyncDestinationStruct) AsyncUploadOutput {
	_ = "STUB: not implemented"
	return *new(AsyncUploadOutput)
}

// AbortReason:   `{"error":"BingAds could not be initialized. Please check account settings."}`,

func (*InvalidManager) Poll(_ context.Context, _ AsyncPoll) PollStatusResponse {
	_ = "STUB: not implemented"
	return *new(PollStatusResponse)
}

func (*InvalidManager) GetUploadStats(_ GetUploadStatsInput) GetUploadStatsResponse {
	_ = "STUB: not implemented"
	return *new(GetUploadStatsResponse)
}
