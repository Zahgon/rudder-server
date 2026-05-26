package sftp

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/sftp"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/common"
)

func (*defaultManager) Transform(job *jobsdb.JobT) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Upload uploads the data to the destination and marks all jobs to be completed
func (d *defaultManager) Upload(_ context.Context, asyncDestStruct *common.AsyncDestinationStruct) common.AsyncUploadOutput {
	_ = "STUB: not implemented"
	return *new(common.AsyncUploadOutput)
}

// Use same file path prefix for each file per sync

// Generate initial file path for file number 1 per sync

// Generate temporary file based on the destination's file format

// Upload file

func newDefaultManager(logger logger.Logger, statsFactory stats.Stats, fileManager sftp.FileManager, config destConfig) *defaultManager {
	_ = "STUB: not implemented"
	return nil
}

func newInternalManager(logger logger.Logger, statsFactory stats.Stats, destination *backendconfig.DestinationT) (common.AsyncUploadAndTransformManager, error) {
	_ = "STUB: not implemented"
	return *new(common.AsyncUploadAndTransformManager), nil
}

func NewManager(logger logger.Logger, statsFactory stats.Stats, destination *backendconfig.DestinationT) (common.AsyncDestinationManager, error) {
	_ = "STUB: not implemented"
	return *new(common.AsyncDestinationManager), nil
}
