package api

import (
	"context"
	"database/sql"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/notifier"
	"github.com/rudderlabs/rudder-server/warehouse/bcm"
	sqlmw "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/internal/repo"
	"github.com/rudderlabs/rudder-server/warehouse/multitenant"
	"github.com/rudderlabs/rudder-server/warehouse/source"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

const triggerUploadQPName = "triggerUpload"

type pendingEventsRequest struct {
	SourceID  string `json:"source_id"`
	TaskRunID string `json:"task_run_id"`
}

type pendingEventsResponse struct {
	PendingEvents            bool  `json:"pending_events"`
	PendingStagingFilesCount int64 `json:"pending_staging_files"`
	PendingUploadCount       int64 `json:"pending_uploads"`
	AbortedEvents            bool  `json:"aborted_events"`
}

type fetchTablesRequest struct {
	Connections []warehouseutils.SourceIDDestinationID `json:"connections"`
}

type fetchTablesResponse struct {
	ConnectionsTables []warehouseutils.FetchTableInfo `json:"connections_tables"`
}

type triggerUploadRequest struct {
	SourceID      string `json:"source_id"`
	DestinationID string `json:"destination_id"`
}

type Api struct {
	mode          string
	conf          *config.Config
	logger        logger.Logger
	statsFactory  stats.Stats
	db            *sqlmw.DB
	notifier      *notifier.Notifier
	bcConfig      backendconfig.BackendConfig
	tenantManager *multitenant.Manager
	bcManager     *bcm.BackendConfigManager
	sourceManager *source.Manager
	stagingRepo   *repo.StagingFiles
	uploadRepo    *repo.Uploads
	schemaRepo    *repo.WHSchema
	triggerStore  *sync.Map

	config struct {
		healthTimeout       time.Duration
		readerHeaderTimeout time.Duration
		runningMode         string
		webPort             int
		mode                string
	}
}

func NewApi(
	mode string,
	conf *config.Config,
	log logger.Logger,
	statsFactory stats.Stats,
	bcConfig backendconfig.BackendConfig,
	db *sqlmw.DB,
	notifier *notifier.Notifier,
	tenantManager *multitenant.Manager,
	bcManager *bcm.BackendConfigManager,
	sourceManager *source.Manager,
	triggerStore *sync.Map,
) *Api {
	_ = "STUB: not implemented"
	return nil
}

func (a *Api) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *Api) addMasterEndpoints(ctx context.Context, r chi.Router) {
	_ = "STUB: not implemented"
	return
}

// TODO: add degraded mode
// TODO: add degraded mode

func (a *Api) healthHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func checkHealth(ctx context.Context, db *sql.DB) bool { _ = "STUB: not implemented"; return false }

// pendingEventsHandler check whether there are any pending staging files or uploads for the given source id
func (a *Api) pendingEventsHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (a *Api) triggerUploadHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (a *Api) fetchTablesHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (a *Api) logMiddleware(delegate http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
