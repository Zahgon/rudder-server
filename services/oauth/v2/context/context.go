package context

import (
	"context"
	"encoding/json"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

type (
	destContextKey   struct{}
	secretContextKey struct{}
)

// CtxWithDestination returns a new context with the given destination.
func CtxWithDestination(ctx context.Context, dest *backendconfig.DestinationT) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// DestinationFromCtx returns the destination from the context, if present.
func DestinationFromCtx(ctx context.Context) (*backendconfig.DestinationT, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// CtxWithSecret returns a new context with the given secret.
func CtxWithSecret(ctx context.Context, secret json.RawMessage) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// SecretFromCtx returns the secret from the context, if present.
func SecretFromCtx(ctx context.Context) (json.RawMessage, bool) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), false
}
