package testhelper

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"

	"github.com/rudderlabs/rudder-go-kit/filemanager"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	warehouseclient "github.com/rudderlabs/rudder-server/warehouse/client"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
)

const (
	WaitFor2Minute                 = 2 * time.Minute
	WaitFor10Minute                = 10 * time.Minute
	DefaultQueryFrequency          = 100 * time.Millisecond
	DefaultWarehouseQueryFrequency = 500 * time.Millisecond
	SourceJobQueryFrequency        = 1000 * time.Millisecond
)

const (
	jobsDBHost     = "localhost"
	jobsDBDatabase = "jobsdb"
	jobsDBUser     = "rudder"
	jobsDBPassword = "password"
)

type TestConfig struct {
	WriteKey                     string
	Schema                       string
	UserID                       string
	WorkspaceID                  string
	JobRunID                     string
	TaskRunID                    string
	SourceID                     string
	Destination                  backendconfig.DestinationT
	DestinationID                string
	DestinationType              string
	Tables                       []string
	Client                       *warehouseclient.Client
	TimestampBeforeSendingEvents time.Time
	Config                       map[string]any
	StagingFilePath              string
	EventsFilePath               string
	StagingFilesEventsMap        EventsCountMap
	TableUploadsEventsMap        EventsCountMap
	WarehouseEventsMap           EventsCountMap
	JobsDB                       *sql.DB
	SourceJob                    bool
	SkipWarehouse                bool
	HTTPPort                     int
	TransformerURL               string
}

func (w *TestConfig) VerifyEvents(t testing.TB) { _ = "STUB: not implemented"; return }

func (w *TestConfig) reset() { _ = "STUB: not implemented"; return }

func GetUserId(provider string) string { _ = "STUB: not implemented"; return "" }

func RandSchema(provider string) string { _ = "STUB: not implemented"; return "" }

func JobsDB(t testing.TB, port int) *sql.DB { _ = "STUB: not implemented"; return nil }

func WithConstantRetries(operation func() error) error { _ = "STUB: not implemented"; return nil }

func UploadSampleTestRecordsTemplateLoadFile(
	t testing.TB,
	fm filemanager.FileManager,
	prefixes []string,
	recordSetIndex int,
) filemanager.UploadedFile {
	_ = "STUB: not implemented"
	return *new(filemanager.UploadedFile)
}

func UploadLoadFile(
	t testing.TB,
	fm filemanager.FileManager,
	fileName string,
	tableName string,
) filemanager.UploadedFile {
	_ = "STUB: not implemented"
	return *new(filemanager.UploadedFile)
}

func UploadLoad(
	t testing.TB,
	fm filemanager.FileManager,
	tableName string,
	content [][]string,
) filemanager.UploadedFile {
	_ = "STUB: not implemented"
	return *new(filemanager.UploadedFile)
}

// Ensure all data is written and compressed

// RetrieveRecordsFromWarehouse retrieves records from the warehouse based on the given query.
// It returns a slice of slices, where each inner slice represents a record's values.
func RetrieveRecordsFromWarehouse(
	t testing.TB,
	db *sql.DB,
	query string,
) [][]string {
	_ = "STUB: not implemented"
	return nil
}

func ConvertRecordsToSchema(input [][]string) model.Schema {
	_ = "STUB: not implemented"
	return *new(model.Schema)
}
