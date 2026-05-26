package kvstoremanager

//go:generate mockgen -destination=../../mocks/services/kvstoremanager/mock_kvstoremanager.go -package=mock_kvstoremanager github.com/rudderlabs/rudder-server/services/kvstoremanager KVStoreManager

import (
	"encoding/json"
)

type KVStoreManager interface {
	CreateClient()
	Close() error
	HMSet(key string, fields map[string]any) error
	HSet(key, field string, value any) error
	StatusCode(err error) int
	DeleteKey(key string) (err error)
	HMGet(key string, fields ...string) (result []any, err error)
	HGetAll(key string) (result map[string]string, err error)

	SendDataAsJSON(jsonData json.RawMessage, config map[string]any) (any, error)
	ShouldSendDataAsJSON(config map[string]any) bool
}

type SettingsT struct {
	Provider string
	Config   map[string]any
}

const (
	hashPath  = "message.hash"
	keyPath   = "message.key"
	valuePath = "message.value"
)

func New(provider string, config map[string]any) (m KVStoreManager) {
	_ = "STUB: not implemented"
	return *new(KVStoreManager)
}

func newManager(settings SettingsT) (m KVStoreManager) {
	_ = "STUB: not implemented"
	return *new(KVStoreManager)
}

func EventToKeyValue(jsonData json.RawMessage) (string, map[string]any) {
	_ = "STUB: not implemented"
	return "", nil
}

// IsHSETCompatibleEvent identifies if the event supports HSET operation
// To support HSET, the event must have the following fields:
// - message.key
// - message.value
// - message.hash
// It doesn't account for the value of the fields.
func IsHSETCompatibleEvent(jsonData json.RawMessage) bool { _ = "STUB: not implemented"; return false }

func ExtractHashKeyValueFromEvent(jsonData json.RawMessage) (hash, key, value string) {
	_ = "STUB: not implemented"
	return "", "", ""
}
