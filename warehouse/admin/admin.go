package admin

import (
	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

type QueryInput struct {
	DestID       string
	SourceID     string
	SQLStatement string
}

type ConfigurationTestInput struct {
	DestID string
}

type ConfigurationTestOutput struct {
	Valid bool
	Error string
}

type Admin struct {
	connectionSources  connectionSourcesFetcher
	createUploadAlways createUploadAlwaysSetter
	logger             logger.Logger
}

type connectionSourcesFetcher interface {
	ConnectionSourcesMap(destID string) (map[string]model.Warehouse, bool)
}

type createUploadAlwaysSetter interface {
	Store(bool)
}

func New(
	connectionSources connectionSourcesFetcher,
	createUploadAlways createUploadAlwaysSetter,
	logger logger.Logger,
) *Admin {
	_ = "STUB: not implemented"
	return nil
}

// TriggerUpload sets uploads to start without delay
func (a *Admin) TriggerUpload(off bool, reply *string) error { _ = "STUB: not implemented"; return nil }

// Query the underlying warehouse
func (a *Admin) Query(s QueryInput, reply *warehouseutils.QueryResult) error {
	_ = "STUB: not implemented"
	return nil
}

// use the sourceID-destID connection if sourceID is not empty

// use any source connected to the given destination otherwise

// ConfigurationTest test the underlying warehouse destination
func (a *Admin) ConfigurationTest(s ConfigurationTestInput, reply *ConfigurationTestOutput) error {
	_ = "STUB: not implemented"
	return nil
}
