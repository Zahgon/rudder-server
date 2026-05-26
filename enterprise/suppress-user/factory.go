package suppression

/*
The suppression package provides functionality to manage user suppression lists in a server application.
This package includes a rudder-server suppression component that has been modified to make use of the
suppression-backup-service in order to get the latest and complete user suppression list, instead of
syncing with the data-regulation-service.

This modification drastically reduces the suppression sync time at the gateway, ensuring that the gateway
starts only once the latest user suppressions are available. The package also includes functionality to
asynchronously retrieve the full suppression list in a separate badgerdb repository, as it might take some time.
Once the full list is available in the badgerdb, the package provides functionality to swap the old badgerdb (with latest users)
with the new badgerdb (with all users).
*/

import (
	"context"
	"io"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/services/controlplane/identity"
	"github.com/rudderlabs/rudder-server/utils/types"
)

type Factory struct {
	EnterpriseToken string
	Log             logger.Logger
}

// Setup initializes the user suppression feature
func (m *Factory) Setup(ctx context.Context, backendConfig backendconfig.BackendConfig) (types.UserSuppression, error) {
	_ = "STUB: not implemented"
	return *new(types.UserSuppression), nil
}

// First starting a repository seeded with the latest data which is faster to load

func alreadySynced(repoPath string) bool { _ = "STUB: not implemented"; return false }

func (m *Factory) retryIndefinitely(ctx context.Context, f func() error, wait time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (m *Factory) newSyncerWithBadgerRepo(repoPath string, seederSource func() (io.ReadCloser, error), maxSeedWaitTime time.Duration, identity identity.Identifier, pollInterval config.ValueLoader[time.Duration]) (*Syncer, Repository, error) {
	_ = "STUB: not implemented"
	return nil, *new(Repository), nil
}

func getRepoPath() (fullSuppressionPath, latestSuppressionPath string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func latestDataSeed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func fullDataSeed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func seederSource(endpoint string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// close body afterwards.
