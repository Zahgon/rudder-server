package client

import (
	"context"
	"net/http"
	"time"

	"github.com/rudderlabs/rudder-go-kit/stats"
)

const defaultTimeout = 10 * time.Second

// StagingFile contains the require metadata to process a staging file.
type StagingFile struct {
	WorkspaceID   string
	SourceID      string
	DestinationID string
	Location      string

	Schema map[string]map[string]string

	FirstEventAt          string
	LastEventAt           string
	TotalEvents           int
	TotalBytes            int
	BytesPerTable         map[string]int64
	UseRudderStorage      bool
	DestinationRevisionID string
	// cloud sources specific info
	SourceTaskRunID  string
	SourceJobID      string
	SourceJobRunID   string
	TimeWindow       time.Time
	ServerInstanceID string
}

// legacyPayload is used to maintain backwards compatibility with the /v1 endpoint.
type legacyPayload struct {
	WorkspaceID      string
	Schema           map[string]map[string]string
	BatchDestination stagingFileBatchDestination

	Location              string
	FirstEventAt          string
	LastEventAt           string
	TotalEvents           int
	TotalBytes            int
	BytesPerTable         map[string]int64
	UseRudderStorage      bool
	DestinationRevisionID string
	// cloud sources specific info
	SourceTaskRunID  string
	SourceJobID      string
	SourceJobRunID   string
	TimeWindow       time.Time
	ServerInstanceID string
}

type stagingFileBatchDestination struct {
	Source      struct{ ID string }
	Destination struct{ ID string }
}

type Warehouse struct {
	baseURL      string
	client       *http.Client
	statsFactory stats.Stats
}

type WarehouseOpts func(*Warehouse)

func WithTimeout(timeout time.Duration) WarehouseOpts {
	_ = "STUB: not implemented"
	return *new(WarehouseOpts)
}

func NewWarehouse(baseURL string, statsFactory stats.Stats, opts ...WarehouseOpts) *Warehouse {
	_ = "STUB: not implemented"
	return nil
}

func (w *Warehouse) Process(ctx context.Context, stagingFile StagingFile) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Warehouse) recordAPICallStats(stagingFile StagingFile, status string, statusCode int) {
	_ = "STUB: not implemented"
	return
}
