package batcher

import (
	"github.com/rudderlabs/rudder-server/jobsdb"
	proto "github.com/rudderlabs/rudder-server/proto/event-schema"
	"github.com/rudderlabs/rudder-server/schema-forwarder/internal/transformer"
)

// A batch of jobs that share the same schema.
type EventSchemaMessageBatch struct {
	Index   int
	Message *proto.EventSchemaMessage
	Jobs    []*jobsdb.JobT
}

// NewEventSchemaMessageBatcher creates a new batcher.
func NewEventSchemaMessageBatcher(transformer transformer.Transformer) *EventSchemaMessageBatcher {
	_ = "STUB: not implemented"
	return nil
}

// EventSchemaMessageBatcher batches jobs by their schema.
type EventSchemaMessageBatcher struct {
	transformer transformer.Transformer

	batchOrder []batchKey
	batchIndex map[batchKey]*EventSchemaMessageBatch
}

// Add adds a job to the batcher after transforming it to an [EventSchemaMessage].
// If the message is already in the batcher, the two messages will be merged to one.
func (sb *EventSchemaMessageBatcher) Add(job *jobsdb.JobT) error {
	_ = "STUB: not implemented"
	return nil
}

// GetMessageBatches returns the message batches in the order they were added.
func (sb *EventSchemaMessageBatcher) GetMessageBatches() []*EventSchemaMessageBatch {
	_ = "STUB: not implemented"
	return nil
}

// batchKey is the key used for batching.
type batchKey struct {
	writeKey        string
	eventType       string
	eventIdentifier string
	hash            string
}
