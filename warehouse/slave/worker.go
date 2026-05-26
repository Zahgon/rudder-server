package slave

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/services/notifier"
	"github.com/rudderlabs/rudder-server/warehouse/bcm"
	"github.com/rudderlabs/rudder-server/warehouse/constraints"
	"github.com/rudderlabs/rudder-server/warehouse/encoding"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	"github.com/rudderlabs/rudder-server/warehouse/source"
)

type uploadProcessingResult struct {
	result uploadResult
	err    error
}

type uploadResult struct {
	TableName             string
	Location              string
	TotalRows             int
	ContentLength         int64
	UploadID              int64
	DestinationRevisionID string
	UseRudderStorage      bool
}

type worker struct {
	conf               *config.Config
	log                logger.Logger
	statsFactory       stats.Stats
	notifier           slaveNotifier
	bcManager          *bcm.BackendConfigManager
	constraintsManager *constraints.Manager
	encodingFactory    *encoding.Factory
	workerIdx          int
	activeJobId        atomic.Int64
	refreshClaimJitter time.Duration

	config struct {
		maxStagingFileReadBufferCapacityInK config.ValueLoader[int]
		maxConcurrentStagingFiles           config.ValueLoader[int]
		claimRefreshInterval                config.ValueLoader[time.Duration]
		enableNotifierHeartbeat             config.ValueLoader[bool]
	}
	stats struct {
		workerIdleTime                 stats.Timer
		workerClaimProcessingSucceeded stats.Counter
		workerClaimProcessingFailed    stats.Counter
		workerClaimProcessingTime      stats.Timer
	}
}

type dedupKey struct {
	tableName string
	idValue   string
}

func newWorker(
	conf *config.Config,
	logger logger.Logger,
	statsFactory stats.Stats,
	notifier slaveNotifier,
	bcManager *bcm.BackendConfigManager,
	constraintsManager *constraints.Manager,
	encodingFactory *encoding.Factory,
	workerIdx int,
) *worker {
	_ = "STUB: not implemented"
	return nil
}

// Increasing maxConcurrentStagingFiles config would also require increasing the memory requests for the slave pods

// Random jitter between [0-5) seconds

func (w *worker) start(ctx context.Context, notificationChan <-chan *notifier.ClaimJob, slaveID string) {
	_ = "STUB: not implemented"
	return
}

// Set active job ID

// Clear active job ID after processing

// run claimRefresh periodically to make sure that job is not orphaned
func (w *worker) runClaimRefresh(ctx context.Context) { _ = "STUB: not implemented"; return }

// Distribute claim refresh requests across workers by adding random delay
// This will prevent database load spikes from synchronized refresh attempts

// processClaimedUploadJob processes a claimed upload job from the notifier.
// It supports both regular upload jobs (one staging file per job) and upload_v2 jobs (multiple staging files per job).
func (w *worker) processClaimedUploadJob(ctx context.Context, claimedJob *notifier.ClaimJob) {
	_ = "STUB: not implemented"
	return
}

// processSingleStagingFile processes a single staging file and writes its data to the appropriate load files.
// It handles downloading, reading, and processing the file contents.
func (w *worker) processSingleStagingFile(
	ctx context.Context,
	jr *jobRun,
	job *basePayload,
	stagingFile stagingFileInfo,
	loadFileNamePrefix string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// default scanner buffer maxCapacity is 64K
// set it to higher value to avoid read stop on read size error

// Create separate load file for each table

// Duplicate detection by id column
// skip duplicate detection for users table as multiple identifies events can be present for same user	id

// Special handling for JSON arrays
// TODO: Will this work for both BQ and RS?

// After processing all lines, increment the metric for duplicates

func (w *worker) processMultiStagingFiles(ctx context.Context, job *payloadV2) ([]uploadResult, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Create a jobRun for all staging files
}

// calculate load file name prefix using md5 hash of all staging file locations

// Initialize Discards Table

func (w *worker) processClaimedSourceJob(ctx context.Context, claimedJob *notifier.ClaimJob) {
	_ = "STUB: not implemented"
	return
}

func (w *worker) runSourceJob(ctx context.Context, sourceJob source.NotifierRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *worker) destinationFromSlaveConnectionMap(destinationId, sourceId string) (model.Warehouse, error) {
	_ = "STUB: not implemented"
	return *new(model.Warehouse), nil
}

// HandleSchemaChange checks if the existing data type (from warehouse schema) is compatible with the inferred data type (from event data)
func HandleSchemaChange(log logger.Logger, existingDataType, inferredDataType model.SchemaType, value any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// only stringify if the previous type is non-string/text/json

// All warehouse destinations currently support int to float coercion.
// value should be float (go json unmarshals all numbers as float64), unless explicitly converted to int. Keeping logic agnostic in this function.
