package misc

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// ExponentialNumber is a simple exponentially increasing number.
type ExponentialNumber[T Number] struct {
	value T
}

// Reset resets the number to zero.
func (expo *ExponentialNumber[T]) Reset() {
	_ = "STUB: not implemented"

	// Next returns the next number, which is the previous one multiplied by 2, always abiding by the min and max provided.
	return
}

func (expo *ExponentialNumber[T]) Next(min, max T) T { _ = "STUB: not implemented"; return *new(T) }
