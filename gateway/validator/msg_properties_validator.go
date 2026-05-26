package validator

import (
	"github.com/rudderlabs/rudder-schemas/go/stream"
)

type msgPropertiesValidator struct {
	validateFn func(*stream.MessageProperties) error
}

func newMsgPropertiesValidator(validateFn func(*stream.MessageProperties) error) *msgPropertiesValidator {
	_ = "STUB: not implemented"
	return nil
}

func (p *msgPropertiesValidator) ValidatorName() string { _ = "STUB: not implemented"; return "" }

func (p *msgPropertiesValidator) Validate(payload []byte, properties *stream.MessageProperties) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
