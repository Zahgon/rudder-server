package types

import (
	"time"

	proctypes "github.com/rudderlabs/rudder-server/processor/types"
)

type Metadata struct {
	MessageID         any            `json:"messageId"`
	ReceivedAt        string         `json:"receivedAt"`
	SourceID          string         `json:"sourceId"`
	SourceType        string         `json:"sourceType"`
	DestinationID     string         `json:"destinationId"`
	DestinationType   string         `json:"destinationType"`
	SourceCategory    string         `json:"sourceCategory"`
	EventType         string         `json:"eventType,omitempty"`
	RecordID          any            `json:"recordId,omitempty"`
	DestinationConfig map[string]any `json:"destinationConfig"`
}

type TransformerEvent struct {
	Message  proctypes.SingularEventT `json:"message"`
	Metadata Metadata                 `json:"metadata"`
}

var destConfigFields = []string{
	"skipTracksTable", "skipUsersTable", "underscoreDivideNumbers", "allowUsersContextTraits",
	"storeFullEvent", "jsonPaths",
}

func New(
	event *proctypes.TransformerEvent,
	uuidGenerator func() string,
	now func() time.Time,
) *TransformerEvent {
	_ = "STUB: not implemented"
	return nil
}
