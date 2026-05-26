package validator

import (
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-schemas/go/stream"
)

// payloadValidator defines an interface for validating payloads and retrieving the validator's name.
type payloadValidator interface {
	Validate(payload []byte, properties *stream.MessageProperties) (bool, error)
	ValidatorName() string
}

// Mediator centralizes the orchestration of multiple payload validator processes.
type Mediator struct {
	log        logger.Logger
	validators []payloadValidator
}

// NewValidateMediator creates a new ValidatorMediator with default validators.
func NewValidateMediator(log logger.Logger, validatorFn func(properties *stream.MessageProperties) error) *Mediator {
	_ = "STUB: not implemented"
	return nil
}

// Validate runs the payload through all registered validators.
func (m *Mediator) Validate(payload []byte, properties *stream.MessageProperties) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
