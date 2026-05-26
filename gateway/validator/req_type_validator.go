package validator

import (
	"github.com/rudderlabs/rudder-schemas/go/stream"
)

type reqTypeValidator struct{}

func newReqTypeValidator() *reqTypeValidator { _ = "STUB: not implemented"; return nil }

func (p *reqTypeValidator) ValidatorName() string { _ = "STUB: not implemented"; return "" }

func (p *reqTypeValidator) Validate(payload []byte, properties *stream.MessageProperties) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
