package archive

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/filemanager"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	sqlmw "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/multitenant"
)

type backupRecordsArgs struct {
	tableName      string
	tableFilterSQL string
	sourceID       string
	destID         string
	uploadID       int64
}

type uploadRecord struct {
	sourceID           string
	destID             string
	uploadID           int64
	startStagingFileId int64
	endStagingFileId   int64
	startLoadFileID    int64
	endLoadFileID      int64
	uploadMetadata     json.RawMessage
	workspaceID        string
}

type Archiver struct {
	db            *sqlmw.DB
	stats         stats.Stats
	log           logger.Logger
	conf          *config.Config
	fileManager   filemanager.Factory
	tenantManager *multitenant.Manager

	config struct {
		archiveUploadRelatedRecords config.ValueLoader[bool]
		canDeleteUploads            config.ValueLoader[bool]
		uploadsArchivalTimeInDays   config.ValueLoader[int]
		uploadRetentionTimeInDays   config.ValueLoader[int]
		archiverTickerTime          config.ValueLoader[time.Duration]
		backupRowsBatchSize         config.ValueLoader[int]
		maxLimit                    config.ValueLoader[int]
	}

	archiveFailedStat stats.Counter
}

func New(
	conf *config.Config,
	log logger.Logger,
	stat stats.Stats,
	db *sqlmw.DB,
	fileManager filemanager.Factory,
	tenantManager *multitenant.Manager,
) *Archiver {
	_ = "STUB: not implemented"
	return nil
}

// default 6 hours

func (a *Archiver) backupRecords(ctx context.Context, args backupRecordsArgs) (backupLocation string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (a *Archiver) deleteFilesInStorage(ctx context.Context, locations []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Archiver) usedRudderStorage(metadata []byte) bool { _ = "STUB: not implemented"; return false }

func (a *Archiver) Do(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *Archiver) countUploadsToArchive(ctx context.Context) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (a *Archiver) archiveUploads(ctx context.Context, maxArchiveLimit int) error {
	_ = "STUB: not implemented"
	return nil
}

// empty workspace id should be excluded as a safety measure

// archive staging files

// delete staging file records

// delete load file records

// update upload metadata

func (a *Archiver) getStagingFilesData(
	ctx context.Context,
	txn *sqlmw.Tx,
	u *uploadRecord,
) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *Archiver) deleteLoadFileRecords(
	ctx context.Context,
	txn *sqlmw.Tx,
	uploadID int64,
	hasUsedRudderStorage bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// no need to delete files in rudder storage

func (a *Archiver) Delete(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *Archiver) deleteUploads(ctx context.Context, limit int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
