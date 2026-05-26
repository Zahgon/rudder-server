package backendconfig

type EventReplayConfigs map[string]*EventReplayConfig

// ApplyReplaySources reads the event replay configuration and adds replay sources to the config
// A replay source is a copy of the original source with a different ID and source definition
// This replay source contains as destinations replay destinations which are copies of the original destinations but with a different ID
func (c *ConfigT) ApplyReplaySources() { _ = "STUB: not implemented"; return }

// no event uploads for replay sources for now
// destinations are added later

// processor is always enabled for replay destinations

// add destinations to sources

// add replay sources to config, only the ones that have destinations

type EventReplayConfig struct {
	Sources      map[string]EventReplaySource      `json:"sources"`
	Destinations map[string]EventReplayDestination `json:"destinations"`
	Connections  []EventReplayConnection           `json:"connections"`
}

type EventReplaySource struct {
	OriginalSourceID string `json:"originalId"`
}

type EventReplayDestination struct {
	OriginalDestinationID string `json:"originalId"`
}

type EventReplayConnection struct {
	SourceID      string `json:"sourceId"`
	DestinationID string `json:"destinationId"`
}
