package snowpipestreaming

import (
	"context"
	"time"

	"github.com/hashicorp/go-retryablehttp"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/common"
	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/snowpipestreaming/internal/model"
	whutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

func New(
	conf *config.Config,
	log logger.Logger,
	statsFactory stats.Stats,
	destination *backendconfig.DestinationT,
) *Manager {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) retryableClient() *retryablehttp.Client { _ = "STUB: not implemented"; return nil }

func (m *Manager) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (m *Manager) Transform(job *jobsdb.JobT) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Upload uploads data to the Snowpipe streaming destination.
// It reads events from the file, groups them by table, and sends them to Snowpipe.
// It returns the IDs of the importing and failed jobs.
// In case of failure, it aborts the jobs and returns the aborted job IDs.
func (m *Manager) Upload(_ context.Context, asyncDest *common.AsyncDestinationStruct) common.AsyncUploadOutput {
	_ = "STUB: not implemented"
	return *new(common.AsyncUploadOutput)
}

// Don't use Manager context here: it may be cancelled while polling/upload is in flight.
// Cancelling in-flight requests can leave final state unknown and trigger retries, causing duplicate events.
// Use a background context; request-level timeouts on HTTP/SQL still bound execution.

// Pre-fetch channel statuses for all cached channels in a single bulk call to avoid duplicate requests.

// If failed reason is not set, set it to the default reason

func (m *Manager) eventsFromFile(fileName string, eventsCount int) ([]*event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// splitEventsExceedingMaxInsertRequestSize splits grouped events into included and overflowed events
// based on the configured max insert request size.
func splitEventsExceedingMaxInsertRequestSize(groupedEvents map[string][]*event, maxInsertRequestSizeBytes int64) (map[string][]*event, []*event) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only account for the insert rows payload size:
// rows JSON = '[' + row1 + ',' + row2 + ... + ']'
// '[' + ']'

// ',' between rows

// sendEventsToSnowpipe sends events to Snowpipe for the given table.
// It creates a channel for the table, inserts the events into the channel, and sends the discard events to the discards table.
// It returns the import info for the table and the discard import info if any.
// In case of failure, it deletes the channel.
func (m *Manager) sendEventsToSnowpipe(
	ctx context.Context,
	destinationID string,
	destConf *destConfig,
	info *uploadInfo,
	channelStatuses map[string]*model.StatusResponse,
) (*importInfo, *importInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

/*
	Discards table events are inserted before main table events.
	This is because if discards insert succeeds but main table insert fails,
	then on retry, duplicate records will occur in the discards table.
	But if events are inserted in the reverse order, then on retry,
	duplicate records will occur in the main table.
*/

func checkForDuplicateIDsInBatch(events []*event) (duplicateCount int) {
	_ = "STUB: not implemented"
	return 0
}

// checkForDuplicatesDueToOffset checks for duplicate ids due to offset.
func (m *Manager) checkForDuplicatesDueToOffset(
	ctx context.Context,
	channelID string,
	events []*event,
	channelStatuses map[string]*model.StatusResponse,
) (duplicateCount int) {
	_ = "STUB: not implemented"
	return 0
}

// If channel status is not found in the cache, we need to get it from the Snowpipe.

// Ignoring negative job ids due to migration

func (m *Manager) insert(ctx context.Context, destinationID string, destConf *destConfig, info *uploadInfo, insertReq *model.InsertRequest, channelID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

/*
	404 can happen in the case where the channel is cached in the rudder-server,
	but then the snowpipe service restarts making the channel id invalid.
	But we don't want to pre-emptively check for validity of the channel before every insert
	because such 404 errors are expected to be rare.
*/

// For all other errors, clean up and return

// schemaFromEvents builds a schema by iterating over events and merging their columns
// using a first-encountered type basis for each column.
func schemaFromEvents(events []*event) whutils.ModelTableSchema {
	_ = "STUB: not implemented"
	return *new(whutils.ModelTableSchema)
}

func (m *Manager) abortJobs(asyncDest *common.AsyncDestinationStruct, abortReason string) common.AsyncUploadOutput {
	_ = "STUB: not implemented"
	return *new(common.AsyncUploadOutput)
}

func (m *Manager) failedJobs(asyncDest *common.AsyncDestinationStruct, failedReason string) common.AsyncUploadOutput {
	_ = "STUB: not implemented"
	return *new(common.AsyncUploadOutput)
}

// Poll checks the status of multiple imports using the import ID from pollInput.
// For the once which have reached the terminal state (success or failure), it caches the import infos in polledImportInfoMap. Later if Poll is called again, it does not need to do the status check again.
// Once all the imports have reached the terminal state, if any imports have failed, it deletes the channels for those imports.
// It returns a PollStatusResponse indicating if any imports are still in progress or if any have failed or succeeded
func (m *Manager) Poll(_ context.Context, pollInput common.AsyncPoll) common.PollStatusResponse {
	_ = "STUB: not implemented"
	return *new(common.PollStatusResponse)
}

// Don't use Manager context here: it may be cancelled while polling/upload is in flight.
// Cancelling in-flight requests can leave final state unknown and trigger retries, causing duplicate events.
// Use a background context; request-level timeouts on HTTP/SQL still bound execution.

// Check for stuck pipeline batch

// Fail events which are still in progress after the threshold

// Reset batch polling start time since all imports completed

// provideInProgressImportInfos provides the import infos that are in progress.
// It also updates the polledImportInfoMap with the import infos when the import reaches the terminal state.
func (m *Manager) provideInProgressImportInfos(ctx context.Context, infos []*importInfo) []*importInfo {
	_ = "STUB: not implemented"
	return nil
}

func isInProgress(statusRes *model.StatusResponse, info *importInfo, log logger.Logger) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Case 1: All events have been flushed - proceed to next batch

// Case 2: Events lost - restart/error scenario

// Case 3: Flushing in progress - continue polling

// Unexpected case - should not reach here based on the logic

func (m *Manager) getBulkStatusForChannel(ctx context.Context, channelID string) (*model.StatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Manager) getBulkStatus(ctx context.Context, channelIDs []string) (map[string]*model.StatusResponse, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (m *Manager) getNonProcessedImportInfosWithBulkStatus(ctx context.Context, infos []*importInfo) []*importInfo {
	_ = "STUB: not implemented"
	return nil
}

// In case of failure, we return all the import infos to be polled again

func (m *Manager) resolveStatusResponsePostBulkStatus(
	ctx context.Context,
	info *importInfo,
	statuses map[string]*model.StatusResponse,
	notFoundMap map[string]struct{},
) (*model.StatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Manager) markImportInfoFailed(info *importInfo, reason string) {
	_ = "STUB: not implemented"
	return
}

func (m *Manager) handleChannelRecoveryPostBulkStatus(ctx context.Context, info *importInfo, deleteChannel bool) (*model.StatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete the channel from the cache and the Snowpipe as it is invalid

func convertToInt(a string) (int64, error) {
	_ = "STUB: not implemented"

	// Using math.MinInt64 instead of 0 to avoid skipping events
	// with negative job ids in case of downscaling of nodes
	return 0, nil
}

func jobIDRangeFields(jobIDs []int64) []logger.Field { _ = "STUB: not implemented"; return nil }

func (m *Manager) cleanupFailedImports(ctx context.Context, failedInfos []*importInfo) {
	_ = "STUB: not implemented"
	return
}

func (m *Manager) updateJobStatistics(importInfos []*importInfo) { _ = "STUB: not implemented"; return }

func (m *Manager) buildPollStatusResponse(importInfos, failedImports []*importInfo) common.PollStatusResponse {
	_ = "STUB: not implemented"
	return *new(common.PollStatusResponse)
}

// GetUploadStats returns the status of the uploads for the snowpipe streaming destination.
// It returns the status of the uploads for the given job IDs.
// If any of the uploads have failed, it returns the reason for the failure.
func (m *Manager) GetUploadStats(input common.GetUploadStatsInput) common.GetUploadStatsResponse {
	_ = "STUB: not implemented"
	return *new(common.GetUploadStatsResponse)
}

func buildCreateChannelRequest(destinationID, partition string, destConf *destConfig, tableName string) *model.CreateChannelRequest {
	_ = "STUB: not implemented"
	return nil
}
