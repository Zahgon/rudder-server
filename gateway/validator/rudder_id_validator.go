package validator

import (
	"github.com/rudderlabs/rudder-schemas/go/stream"
)

type rudderIDValidator struct{}

func newRudderIDValidator() *rudderIDValidator { _ = "STUB: not implemented"; return nil }

func (p *rudderIDValidator) ValidatorName() string { _ = "STUB: not implemented"; return "" }

func (p *rudderIDValidator) Validate(payload []byte, properties *stream.MessageProperties) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
