package auth

import (
	"errors"
	"net/http"

	gwtypes "github.com/rudderlabs/rudder-server/gateway/types"
)

var ErrSourceNotFound = errors.New("source not found")

type WebhookAuth struct {
	onFailure             func(w http.ResponseWriter, r *http.Request, errorMessage string, authCtx *gwtypes.AuthRequestContext)
	authReqCtxForWriteKey func(writeKey string) (*gwtypes.AuthRequestContext, error)
}

func NewWebhookAuth(
	onFailure func(w http.ResponseWriter, r *http.Request, errorMessage string, authCtx *gwtypes.AuthRequestContext),
	authReqCtxForWriteKey func(writeKey string) (*gwtypes.AuthRequestContext, error),
) *WebhookAuth {
	_ = "STUB: not implemented"
	return nil
}

func (wa *WebhookAuth) AuthHandler(next http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
