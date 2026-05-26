package crash

import "net/http"

type NOOP struct{}

func (*NOOP) Notify(team string) func() { _ = "STUB: not implemented"; return nil }

func (*NOOP) Handler(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
