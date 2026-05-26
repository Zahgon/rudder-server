package integrations

import (
	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/processor/types"
)

// PostParametersT is a struct for holding all the values from transformerResponse and use them to publish an event to a destination
type PostParametersT struct {
	Type          string `json:"type"`
	URL           string `json:"endpoint"`
	EndpointPath  string `json:"endpointPath,omitempty"`
	RequestMethod string `json:"method"`
	// Invalid tag used in struct. skipcq: SCC-SA5008
	UserID      string            `json:"userId"`
	Headers     map[string]string `json:"headers"`
	QueryParams map[string]any    `json:"params"`
	Body        map[string]any    `json:"body"`
	Files       map[string]any    `json:"files"`
}

type TransStatsT struct {
	StatTags map[string]string `json:"statTags"`
}

func CollectDestErrorStats(input []byte) { _ = "STUB: not implemented"; return }

func CollectIntgTransformErrorStats(input []byte) { _ = "STUB: not implemented"; return }

// FilterClientIntegrations parses the destination names from the
// input JSON, matches them with enabled destinations from controle plane and returns the IDSs
func FilterClientIntegrations(clientEvent types.SingularEventT, destNameIDMap map[string]backendconfig.DestinationDefinitionT) (retVal []string) {
	_ = "STUB: not implemented"
	return nil
}

// All is by default true, if not present make it true

// if dest is bool and is present in clientIntgretaion list, check if true/false

// Always add for syntax dest:{...}

// if dest  not present in clientIntgretaion list, add based on All flag
