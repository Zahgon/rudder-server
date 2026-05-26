package transformer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
	proto "github.com/rudderlabs/rudder-server/proto/event-schema"
)

type Transformer interface {
	Start()
	Transform(job *jobsdb.JobT) (*proto.EventSchemaMessage, error)
	Stop()
}

// New returns a new instance of Schema Transformer
func New(backendConfig backendconfig.BackendConfig, config *config.Config) Transformer {
	_ = "STUB: not implemented"
	return *new(Transformer)
}

// Start starts the schema transformer
func (st *transformer) Start() { _ = "STUB: not implemented"; return }

// Stop stops the schema transformer
func (st *transformer) Stop() { _ = "STUB: not implemented"; return }

// Transform transforms the job into a schema message and returns the schema message along with write key
func (st *transformer) Transform(job *jobsdb.JobT) (*proto.EventSchemaMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getSchemaKeyFromJob returns the schema key from the job based on the event type and event identifier
func (st *transformer) getSchemaKeyFromJob(eventPayload map[string]any, writeKey string) *proto.EventSchemaKey {
	_ = "STUB: not implemented"
	return nil
}

func (st *transformer) backendConfigSubscriber(ctx context.Context, loopFn func()) {
	_ = "STUB: not implemented"
	return
}

// getEventType returns the event type from the event
func (st *transformer) getEventType(event map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

// getEventIdentifier returns the event identifier from the event
func (st *transformer) getEventIdentifier(event map[string]any, eventType string) string {
	_ = "STUB: not implemented"
	return ""
}

// getSchemaMessage returns the schema message from the event by flattening the event and getting the schema
func (st *transformer) getSchemaMessage(key *proto.EventSchemaKey, event map[string]any, sample json.RawMessage, workspaceId string, observedAt time.Time) (*proto.EventSchemaMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// redact event

// getSchema returns the schema from the flattened event
func (st *transformer) getSchema(flattenedEvent map[string]any) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// flattenEvent flattens the event
func (st *transformer) flattenEvent(event map[string]any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// disablePIIReporting returns whether PII reporting is disabled for the write key
func (st *transformer) disablePIIReporting(writeKey string) bool {
	_ = "STUB: not implemented"
	return false
}

// getWriteKeyFromParams returns the write key from the job parameters
func (st *transformer) getWriteKeyFromParams(parameters json.RawMessage) string {
	_ = "STUB: not implemented"
	return ""
}
