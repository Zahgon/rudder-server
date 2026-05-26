package pubsub

import (
	"context"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	types "github.com/rudderlabs/rudder-server/processor/types"
)

var sourceKeys = []string{"properties", "traits", "context.traits"}

func Transform(_ context.Context, events []types.TransformerEvent) types.Response {
	_ = "STUB: not implemented"
	return *new(types.Response)
}

// TODO: Currently, it's getting ignored during JSON marshalling Remove this once we start using it.

func getAttributesMap(destination backendconfig.DestinationT) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func getTopic(event types.TransformerEvent, topicMap map[string]string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Capital "No" needed for mirroring/comparison; re-enable lint after mirroring ends.
//nolint:staticcheck

func getAttributeKeysFromEvent(event types.TransformerEvent, attributesMap map[string][]string) []string {
	_ = "STUB: not implemented"
	return nil
}

func getAttributesMapFromEvent(event types.TransformerEvent, attributesMap map[string][]string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// getAttributeValue searches for an attribute in the message and its nested structures
func getAttributeValue(message map[string]any, attribute string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
