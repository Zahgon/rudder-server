package servermode

import "context"

type Mode string

const (
	NormalMode   Mode = "NORMAL"
	DegradedMode Mode = "DEGRADED"
)

type ChangeEvent struct {
	err  error
	ack  func(context.Context) error
	mode Mode
}

func NewChangeEvent(mode Mode, ack func(context.Context) error) ChangeEvent {
	_ = "STUB: not implemented"
	return *new(ChangeEvent)
}

func ChangeEventError(err error) ChangeEvent { _ = "STUB: not implemented"; return *new(ChangeEvent) }

func (m ChangeEvent) Ack(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (m ChangeEvent) Mode() Mode { _ = "STUB: not implemented"; return *new(Mode) }

func (m ChangeEvent) Err() error { _ = "STUB: not implemented"; return nil }

func (mode Mode) Valid() bool { _ = "STUB: not implemented"; return false }
