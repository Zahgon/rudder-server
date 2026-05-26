package exporter

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/logger"

	suppression "github.com/rudderlabs/rudder-server/enterprise/suppress-user"
	"github.com/rudderlabs/rudder-server/services/controlplane/identity"
	"github.com/rudderlabs/rudder-server/suppression-backup-service/model"
)

const TmpExportFilePrefix = "tmp-export"

// CleanupLingeringTmpExportFiles removes any lingering temporary export files from previous runs
func CleanupLingeringTmpExportFiles() error { _ = "STUB: not implemented"; return nil }

// find all files starting with tmp-export in the tmpDir and remove them

func Export(repo suppression.Repository, file model.File) error {
	_ = "STUB: not implemented"
	// export initially to a temp file
	return nil
}

// export to temp file

// move temp file to final destination

func newBadgerDBInstance(baseDir string, pkgLogger logger.Logger) (suppression.Repository, error) {
	_ = "STUB: not implemented"
	return *new(suppression.Repository), nil
}

type Exporter struct {
	Id   identity.Identifier
	File model.File
	Log  logger.Logger
}

func (e *Exporter) FullExporterLoop(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Exporter) LatestExporterLoop(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// manually add token to repo to make sure that we get suppression regulation corresponding to last one month and not older than that.

func latestToken() (string, error) { _ = "STUB: not implemented"; return "", nil }

type Token struct {
	SyncStartTime time.Time
	SyncSeqId     int
}
