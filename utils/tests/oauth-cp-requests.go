package testutils

import (
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/rudderlabs/rudder-server/utils/types/deployment"
)

type CpResponseParams struct {
	Timeout  time.Duration
	Code     int
	Response string
}
type CpResponseProducer struct {
	Responses []CpResponseParams
	callCount int
}

func (cp *CpResponseProducer) GetNext() CpResponseParams {
	_ = "STUB: not implemented"
	return *new(CpResponseParams)
}

func (cp *CpResponseProducer) MockCpRequests() *chi.Mux { _ = "STUB: not implemented"; return nil }

// iterating over request parameters

// This case wouldn't occur I guess

// sleep is being used to mimic the waiting in actual transformer response

type BasicAuthMock struct{}

func (b *BasicAuthMock) BasicAuth() (string, string) { _ = "STUB: not implemented"; return "", "" }

func (b *BasicAuthMock) ID() string { _ = "STUB: not implemented"; return "" }

func (b *BasicAuthMock) Type() deployment.Type {
	_ = "STUB: not implemented"
	return *new(deployment.Type)
}
