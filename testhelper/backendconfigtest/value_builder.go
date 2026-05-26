package backendconfigtest

type valueBuilder[V any] struct {
	v *V
}

// Build builds the value
func (b *valueBuilder[V]) Build() V { _ = "STUB: not implemented"; return *new(V) }
