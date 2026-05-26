package kafka

import (
	"context"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/processor/types"
)

var canonicalNames = []string{"KAFKA", "kafka", "Kafka"}

func Transform(_ context.Context, events []types.TransformerEvent) types.Response {
	_ = "STUB: not implemented"
	return *new(types.Response)
}

// TODO: Currently, it's getting ignored during JSON marshalling Remove this once we start using it.

func getTopic(event types.TransformerEvent, integrationsObj map[string]any, eventTypeToTopicMap, eventToTopicMap map[string]string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func filterConfigTopics(message types.SingularEventT, destination backendconfig.DestinationT, eventTypeToTopicMap, eventToTopicMap map[string]string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
