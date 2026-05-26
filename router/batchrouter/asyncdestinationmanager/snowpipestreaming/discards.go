package snowpipestreaming

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/snowpipestreaming/internal/model"
	whutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

// sendDiscardEventsToSnowpipe uploads discarded records to the Snowpipe discards table.
// In case of failure, it deletes the channel.
func (m *Manager) sendDiscardEventsToSnowpipe(
	ctx context.Context,
	offset string,
	discardsChannelID string,
	discardInfos []discardInfo,
) (*importInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func discardsTable() string { _ = "STUB: not implemented"; return "" }

func discardsSchema() whutils.ModelTableSchema {
	_ = "STUB: not implemented"
	return *new(whutils.ModelTableSchema)
}

// getDiscardedRecordsFromEvent returns the records that were discarded due to schema mismatch
// It also updates the event data with the converted values
// If the conversion fails, the value is discarded
// If the value is a slice, it is marshalled to a string
func getDiscardedRecordsFromEvent(
	log logger.Logger,
	event *event,
	snowpipeSchema whutils.ModelTableSchema,
	tableName string,
	formattedTS string,
) (discardedRecords []discardInfo) {
	_ = "STUB: not implemented"
	return nil
}

// Discard value if conversion fails

// Update value if conversion succeeds

// Discard value if marshalling fails

// convertDiscardedInfosToRows converts discardInfo to model.Row
func convertDiscardedInfosToRows(discardInfos []discardInfo) []model.Row {
	_ = "STUB: not implemented"
	return nil
}
