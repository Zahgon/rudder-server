package redis

import (
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/tidwall/gjson"

	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/utils/types"
)

var abortableErrors = []string{}

type RedisManager struct {
	logger        logger.Logger
	clusterMode   bool
	config        types.ConfigT
	client        *redis.Client
	clusterClient *redis.ClusterClient
}

func init() {
	abortableErrors = []string{"connection refused", "invalid password"}
}

func NewRedisManager(config types.ConfigT) *RedisManager { _ = "STUB: not implemented"; return nil }

func (m *RedisManager) GetClient() redis.Cmdable {
	_ = "STUB: not implemented"
	return *new(redis.Cmdable)
}

func (m *RedisManager) CreateClient() { _ = "STUB: not implemented"; return }

// setting redis to cluster mode by default if setting missing in config

func (m *RedisManager) Close() error { _ = "STUB: not implemented"; return nil }

func (m *RedisManager) HMSet(key string, fields map[string]any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (*RedisManager) StatusCode(err error) int { _ = "STUB: not implemented"; return 0 }

func (m *RedisManager) DeleteKey(key string) (err error) { _ = "STUB: not implemented"; return nil }

func (m *RedisManager) HMGet(key string, fields ...string) (result []any, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *RedisManager) HGetAll(key string) (result map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *RedisManager) HSet(hash, key string, value any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type jsonSetCmdArgs struct {
	key   string
	path  string
	value string
}

func (m *RedisManager) setArgsForMergeStrategy(inputArgs setArguments) (*jsonSetCmdArgs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// value to which the transformed value should be merged which will be inserted into Redis

// transformed value

// merge jsons

type setArguments struct {
	key     string
	path    string
	jsonVal gjson.Result
}

// nolint:unparam
func (m *RedisManager) extractJSONSetArgs(transformedData json.RawMessage, config map[string]any) (*jsonSetCmdArgs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *RedisManager) SendDataAsJSON(jsonData json.RawMessage, config map[string]any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (*RedisManager) ShouldSendDataAsJSON(config map[string]any) bool {
	_ = "STUB: not implemented"
	return false
}
