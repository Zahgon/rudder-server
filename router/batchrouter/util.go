package batchrouter

import (
	stdjson "encoding/json"
	"sync"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/filemanager"
	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

func IsObjectStorageDestination(destType string) bool { _ = "STUB: not implemented"; return false }

func IsWarehouseDestination(destType string) bool { _ = "STUB: not implemented"; return false }

func IsBatchRouterDestination(destination string) bool { _ = "STUB: not implemented"; return false }

func connectionIdentifier(batchDestination Connection) string { _ = "STUB: not implemented"; return "" }

func getNamespace(config any, source backendconfig.SourceT, destType string) string {
	_ = "STUB: not implemented"
	return ""
}

// TODO: Handle if configMap["database"] is nil

func warehouseConnectionIdentifier(destType, connIdentifier string, source backendconfig.SourceT, destination backendconfig.DestinationT) string {
	_ = "STUB: not implemented"
	return ""
}

func getBRTErrorCode(state string) int { _ = "STUB: not implemented"; return 0 }

func getReloadableBatchRouterConfigInt(key, destType string, defaultValue int) config.ValueLoader[int] {
	_ = "STUB: not implemented"
	return nil
}

type storageDateFormatProvider struct {
	dateFormatsCacheMu sync.RWMutex
	dateFormatsCache   map[string]string // (sourceId:destinationId) -> dateFormat
}

func (sdfp *storageDateFormatProvider) GetFormat(log logger.Logger, manager filemanager.FileManager, destination *Connection, folderName string) (dateFormat string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Returning the earlier default as we might not able to fetch the list.
// because "*:GetObject" and "*:ListBucket" permissions are not available.

func IsAsyncDestinationLimitNotReached(brt *Handle, destinationID string) bool {
	_ = "STUB: not implemented"
	return false
}

func getFirstSourceJobRunID(params map[int64]stdjson.RawMessage) string {
	_ = "STUB: not implemented"
	return ""
}
