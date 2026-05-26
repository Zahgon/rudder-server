package snowpipestreaming

import (
	"context"
	"errors"

	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager/snowpipestreaming/internal/model"
	"github.com/rudderlabs/rudder-server/warehouse/integrations/manager"
	whutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

var errAbort = errors.New("abort error")

// initializeChannelWithSchema creates a new channel for the given table if it doesn't exist.
// If the channel already exists, it checks for new columns and adds them to the table.
// It returns the channel response after creating or recreating the channel.
func (m *Manager) initializeChannelWithSchema(
	ctx context.Context,
	destinationID string,
	destConf *destConfig,
	tableName string,
	eventSchema whutils.ModelTableSchema,
) (*model.ChannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findNewColumns(eventSchema, snowpipeSchema whutils.ModelTableSchema) []whutils.ColumnInfo {
	_ = "STUB: not implemented"
	return nil
}

// addColumns adds columns to a Snowflake table one at a time, as ALTER TABLE does not support IF NOT EXISTS with multiple columns.
func (m *Manager) addColumns(ctx context.Context, namespace, tableName string, columns []whutils.ColumnInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// createChannel creates a new channel for importing data to Snowpipe.
// If the channel already exists in the cache, it returns the cached response. Otherwise, it sends a request to create a new channel.
// It also handles errors related to missing schemas and tables.
func (m *Manager) createChannel(
	ctx context.Context,
	rudderIdentifier string,
	destConf *destConfig,
	tableName string,
	eventSchema whutils.ModelTableSchema,
) (*model.ChannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handleSchemaError handles errors related to missing schemas.
// It creates the necessary schema and table, then attempts to create the channel again.
func (m *Manager) handleSchemaError(
	ctx context.Context,
	channelReq *model.CreateChannelRequest,
	eventSchema whutils.ModelTableSchema,
) (*model.ChannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handleTableError handles errors related to missing tables.
// It creates the necessary table and then attempts to create the channel again.
func (m *Manager) handleTableError(
	ctx context.Context,
	channelReq *model.CreateChannelRequest,
	eventSchema whutils.ModelTableSchema,
) (*model.ChannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// recreateChannel deletes an existing channel and then creates a new one.
func (m *Manager) recreateChannel(
	ctx context.Context,
	destinationID string,
	destConf *destConfig,
	tableName string,
	eventSchema whutils.ModelTableSchema,
	existingChannelID string,
) (*model.ChannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// deleteChannel removes a channel from the cache and deletes it from the Snowpipe.
func (m *Manager) deleteChannel(ctx context.Context, tableName, channelID string) error {
	_ = "STUB: not implemented"
	return nil
}

// deleteChannelFromCache removes a channel from the cache
func (m *Manager) deleteChannelFromCache(tableName string) { _ = "STUB: not implemented"; return }

func (m *Manager) createSnowflakeManager(ctx context.Context, namespace string) (manager.Manager, error) {
	_ = "STUB: not implemented"
	return *new(manager.Manager), nil
}
