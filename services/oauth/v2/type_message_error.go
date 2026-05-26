package v2

// TypeMessageError is an error that has both a type and a message.
type TypeMessageError struct {
	Type    string
	Message string
}

func (e *TypeMessageError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *TypeMessageError) Unwrap() error { _ = "STUB: not implemented"; return nil }
