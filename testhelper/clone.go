package testhelper

import (
	"testing"
)

func Clone[T any](t testing.TB, v T) T { _ = "STUB: not implemented"; return *new(T) }
