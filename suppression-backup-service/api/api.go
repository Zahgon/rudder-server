package api

import (
	"context"
	"net/http"

	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/suppression-backup-service/model"
)

type API struct {
	log          logger.Logger
	fullBackup   model.File
	latestBackup model.File
}

func NewAPI(logger logger.Logger, fullBackup, latestBackup model.File) *API {
	_ = "STUB: not implemented"
	return nil
}

func (api *API) Handler(ctx context.Context) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func ServeFile(file model.File) func(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return nil
}
