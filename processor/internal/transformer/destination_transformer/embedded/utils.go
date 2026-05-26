package utils

import (
	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/processor/types"
)

var (
	genericTimestampFieldMap = []string{"timestamp", "originalTimestamp"}
	timestampValsMap         = map[string][]string{
		"identify": append([]string{"context.timestamp", "context.traits.timestamp", "traits.timestamp"}, genericTimestampFieldMap...),
		"track":    append([]string{"properties.timestamp"}, genericTimestampFieldMap...),
	}
)

func GetTopicMap(destination backendconfig.DestinationT, key string, convertKeyToLower bool) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func GetValidationErrorStatTags(destination backendconfig.DestinationT) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

/* TODO: remove this once we stop comparing response from embedded and legacy
 * we need this because response from legacy is a map[string]interface{} and not a types.SingularEventT
 * and we need to convert it to types.SingularEventT to compare with response from embedded
 */
func GetMessageAsMap(message types.SingularEventT) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func UpdateTimestampFieldForRETLEvent(eventMessage types.SingularEventT) types.SingularEventT {
	_ = "STUB: not implemented"
	return *new(types.SingularEventT)
}
