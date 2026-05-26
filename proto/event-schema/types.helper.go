package proto

func (esk *EventSchemaKey) MustMarshal() []byte { _ = "STUB: not implemented"; return nil }

func (esm *EventSchemaMessage) MustMarshal() []byte { _ = "STUB: not implemented"; return nil }

// UnmarshalEventSchemaMessage creates a new event schema message from the provided protobuf bytes.
func UnmarshalEventSchemaMessage(raw []byte) (*EventSchemaMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SchemaHash returns a hash of the schema. Keys are sorted lexicographically during hashing.
func SchemaHash(schema map[string]string) string { _ = "STUB: not implemented"; return "" }

// Merge merges the other event schema message into this one.
func (sm *EventSchemaMessage) Merge(other *EventSchemaMessage) { _ = "STUB: not implemented"; return }

// keep the smallest sample

// keep the latest observed time
