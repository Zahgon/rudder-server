package validator

import (
	"github.com/rudderlabs/rudder-schemas/go/stream"
)

type messageIDValidator struct{}

func newMessageIDValidator() *messageIDValidator { _ = "STUB: not implemented"; return nil }

func (p *messageIDValidator) ValidatorName() string { _ = "STUB: not implemented"; return "" }

func (p *messageIDValidator) Validate(payload []byte, properties *stream.MessageProperties) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
