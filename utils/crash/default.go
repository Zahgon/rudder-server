package crash

import (
	"net/http"
	"sync/atomic"

	"github.com/rudderlabs/rudder-go-kit/logger"
)

var defaultHandler atomic.Pointer[panicHandler]

func init() {
	var noop panicHandler = &NOOP{}
	defaultHandler.Store(&noop)
}

func getDefault() panicHandler { _ = "STUB: not implemented"; return *new(panicHandler) }

type panicHandler interface {
	Notify(team string) func()
	Handler(h http.Handler) http.Handler
}

type PanicWrapperOpts struct {
	AppVersion   string
	ReleaseStage string
	AppType      string
}

func Configure(logger logger.Logger, opts PanicWrapperOpts) { _ = "STUB: not implemented"; return }

func NotifyWarehouse(fn func() error) func() error { _ = "STUB: not implemented"; return nil }

func Wrapper(fn func() error) func() error { _ = "STUB: not implemented"; return nil }

func WrapperNoError(fn func()) func() { _ = "STUB: not implemented"; return nil }

func Notify(team string) func() { _ = "STUB: not implemented"; return nil }

func Handler(h http.Handler) http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }
