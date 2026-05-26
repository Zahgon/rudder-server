package router

import (
	"context"
)

func (job *UploadJob) generateLoadFiles() error { _ = "STUB: not implemented"; return nil }

func (job *UploadJob) setLoadFileIDs(startLoadFileID, endLoadFileID int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (job *UploadJob) matchRowsInStagingAndLoadFiles(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (job *UploadJob) getTotalRowsInLoadFiles(ctx context.Context) int64 {
	_ = "STUB: not implemented"
	return 0
}
