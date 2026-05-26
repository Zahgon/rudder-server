package drain_config

import (
	"context"
	"net/http"
)

func (d *drainConfigManager) DrainConfigHttpHandler() http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (d *drainConfigManager) drainJob(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (d *drainConfigManager) insert(ctx context.Context, key, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func ErrorResponder(errMsg string) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
