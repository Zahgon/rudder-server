package filehandler

import (
	"context"

	"github.com/rudderlabs/rudder-server/regulation-worker/internal/model"
)

type GZIPLocalFileHandler struct {
	records     []byte
	idFieldPath []string
}

func NewGZIPLocalFileHandler(idFieldPath []string) *GZIPLocalFileHandler {
	_ = "STUB: not implemented"
	return nil
}

func (h *GZIPLocalFileHandler) Read(_ context.Context, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *GZIPLocalFileHandler) Write(_ context.Context, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *GZIPLocalFileHandler) RemoveIdentity(_ context.Context, attributes []model.User) error {
	_ = "STUB: not implemented"
	return nil
}
