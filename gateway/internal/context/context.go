package context

import (
	"context"

	gwtypes "github.com/rudderlabs/rudder-server/gateway/types"
)

// GetRequestTypeFromCtx : get request type from context
func GetRequestTypeFromCtx(ctx context.Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// GetAuthRequestFromCtx : get auth request from context
func GetAuthRequestFromCtx(ctx context.Context) (*gwtypes.AuthRequestContext, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
