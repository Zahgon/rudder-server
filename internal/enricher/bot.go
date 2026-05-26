package enricher

import (
	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/processor/types"
)

type botDetails struct {
	Name             string `json:"name,omitempty"`
	URL              string `json:"url,omitempty"`
	IsInvalidBrowser bool   `json:"isInvalidBrowser,omitempty"`
}

type botEnricher struct{}

func NewBotEnricher() (PipelineEnricher, error) {
	_ = "STUB: not implemented"
	return *new(PipelineEnricher), nil
}

func (e *botEnricher) Enrich(_ *backendconfig.SourceT, request *types.GatewayBatchRequest, eventParams *types.EventParams) error {
	_ = "STUB: not implemented"
	return nil
}

// if the event is not a bot, we don't need to enrich it

// if the context section is missing on the event
// set it with default as map[string]any

// if the context is other than map[string]any, add error and continue

func (e *botEnricher) Close() error { _ = "STUB: not implemented"; return nil }
