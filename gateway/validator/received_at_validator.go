package validator

import (
	"github.com/rudderlabs/rudder-schemas/go/stream"
)

type receivedAtValidator struct{}

func newReceivedAtValidator() *receivedAtValidator { _ = "STUB: not implemented"; return nil }

func (e *receivedAtValidator) ValidatorName() string { _ = "STUB: not implemented"; return "" }

func (e *receivedAtValidator) Validate(payload []byte, properties *stream.MessageProperties) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
