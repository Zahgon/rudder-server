package suppression

import (
	"context"
	"net/http"
	"time"

	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/enterprise/suppress-user/model"
	"github.com/rudderlabs/rudder-server/services/controlplane/identity"
)

// SyncerOpt represents a configuration option for the syncer
type SyncerOpt func(*Syncer)

// WithHttpClient sets the http client to use
func WithHttpClient(client *http.Client) SyncerOpt {
	_ = "STUB: not implemented"
	return *new(SyncerOpt)
}

// WithPageSize sets the page size for each sync request
func WithPageSize(pageSize int) SyncerOpt { _ = "STUB: not implemented"; return *new(SyncerOpt) }

// WithPollIntervalFn sets the interval at which the syncer will poll the backend
func WithPollIntervalFn(pollIntervalFn func() time.Duration) SyncerOpt {
	_ = "STUB: not implemented"
	return *new(SyncerOpt)
}

// WithLogger sets the logger to use in the syncer
func WithLogger(log logger.Logger) SyncerOpt { _ = "STUB: not implemented"; return *new(SyncerOpt) }

// MustNewSyncer creates a new syncer, panics if an error occurs
func MustNewSyncer(baseURL string, identifier identity.Identifier, r Repository, opts ...SyncerOpt) *Syncer {
	_ = "STUB: not implemented"
	return nil
}

// NewSyncer creates a new syncer
func NewSyncer(baseURL string, identifier identity.Identifier, r Repository, opts ...SyncerOpt) (*Syncer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Syncer is responsible for syncing suppressions from the backend to the repository
type Syncer struct {
	url string
	id  identity.Identifier
	r   Repository

	client             *http.Client
	log                logger.Logger
	pageSize           int
	pollIntervalFn     func() time.Duration
	defaultWorkspaceID string
}

// SyncLoop runs the sync loop until the provided context is done
func (s *Syncer) SyncLoop(ctx context.Context) { _ = "STUB: not implemented"; return }

// Sync synchronises suppressions from the data regulation service in batches, until
// it completes, or an error occurs. Synchronisation completes when the service responds
// with suppressions whose number is less than the page size.
func (s *Syncer) Sync(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// sync fetches suppressions from the backend
func (s *Syncer) sync(token []byte) ([]model.Suppression, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// If statusCode is not 2xx, then returning empty regulations

type suppressionsResponse struct {
	Items []model.Suppression `json:"items"`
	Token string              `json:"token"`
}
