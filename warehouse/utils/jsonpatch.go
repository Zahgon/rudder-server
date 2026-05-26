package warehouseutils

import (
	"encoding/json"
)

// GenerateJSONPatch generates a JSON patch (RFC 6902) that transforms original into modified.
// Returns the patch as json.RawMessage.
func GenerateJSONPatch(original, modified json.RawMessage) (json.RawMessage, error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), nil
}

// ApplyPatchToJSON applies a JSON patch (RFC 6902) to a document and returns the result as json.RawMessage.
func ApplyPatchToJSON(original, patch json.RawMessage) (json.RawMessage, error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), nil
}
