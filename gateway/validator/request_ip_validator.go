package validator

import (
	"github.com/rudderlabs/rudder-schemas/go/stream"
)

type requestIPValidator struct{}

func newRequestIPValidator() *requestIPValidator { _ = "STUB: not implemented"; return nil }

func (e *requestIPValidator) ValidatorName() string { _ = "STUB: not implemented"; return "" }

func (p *requestIPValidator) Validate(payload []byte, properties *stream.MessageProperties) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
