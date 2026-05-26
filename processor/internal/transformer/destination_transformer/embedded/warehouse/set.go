package warehouse

import (
	"github.com/rudderlabs/rudder-server/processor/internal/transformer/destination_transformer/embedded/warehouse/internal/rules"
)

func setDataAndMetadataFromInput(
	tec *transformEventContext,
	input any,
	data map[string]any, metadata map[string]string,
	pi *prefixInfo,
) error {
	_ = "STUB: not implemented"
	return nil
}

func shouldHandleStringLikeObject(inputMap map[string]any, pi *prefixInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func handleStringLikeObject(
	tec *transformEventContext,
	inputMap map[string]any,
	data map[string]any, metadata map[string]string,
	pi *prefixInfo,
) error {
	_ = "STUB: not implemented"
	return nil
}

func addDataAndMetadata(tec *transformEventContext, key string, val any, isJSONKey bool, data map[string]any, metadata map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func isValidJSONPath(tec *transformEventContext, key string, pi *prefixInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func handleValidJSONPath(
	tec *transformEventContext,
	key string, val any,
	data map[string]any, metadata map[string]string,
	pi *prefixInfo,
) error {
	_ = "STUB: not implemented"
	return nil
}

func shouldProcessNestedObject(tec *transformEventContext, val any, pi *prefixInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func processNestedObject(
	tec *transformEventContext,
	key string, val map[string]any,
	data map[string]any, metadata map[string]string,
	pi *prefixInfo,
) error {
	_ = "STUB: not implemented"
	return nil
}

func processNonNestedObject(
	tec *transformEventContext,
	key string, val any,
	data map[string]any, metadata map[string]string,
	pi *prefixInfo,
) error {
	_ = "STUB: not implemented"
	return nil
}

func setDataAndMetadataFromRules(
	tec *transformEventContext,
	data map[string]any, metadata map[string]string,
	rules map[string]rules.Rules,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transformer) storeRudderEvent(
	tec *transformEventContext,
	data map[string]any, metadata map[string]string,
) error {
	_ = "STUB: not implemented"
	return nil
}
