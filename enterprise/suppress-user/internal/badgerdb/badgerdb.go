package badgerdb

import (
	"io"
	"sync"
	"time"

	badger "github.com/dgraph-io/badger/v4"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/enterprise/suppress-user/model"
)

// the key used in badgerdb to store the current token
const tokenKey = "__token__"

// Opt is a function that configures a badgerdb repository
type Opt func(*Repository)

// WithSeederSource sets the source of the seed data
func WithSeederSource(seederSource func() (io.ReadCloser, error)) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithMaxSeedWait sets the maximum time to wait for the seed to complete.
// If the seed takes longer than this, the repository will be started in restoring state and all
// repository methods will return [ErrRestoring] until the seed completes. The default wait time is 10 seconds.
func WithMaxSeedWait(maxSeedWait time.Duration) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// Repository is a repository backed by badgerdb
type Repository struct {
	// logger to use
	log logger.Logger
	// path to the badger db directory
	path string
	// max number of goroutines to use (badger config)
	maxGoroutines int

	maxSeedWait  time.Duration
	seederSource func() (io.ReadCloser, error)

	db *badger.DB

	// lock to prevent concurrent access to db during restore
	restoringLock sync.RWMutex
	restoring     bool
	closeOnce     sync.Once
	closed        chan struct{}
	stats         stats.Stats
}

// NewRepository returns a new repository backed by badgerdb.
func NewRepository(basePath string, log logger.Logger, stats stats.Stats, opts ...Opt) (*Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetToken returns the current token
func (b *Repository) GetToken() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// release the read lock at the end of the operation

// Suppressed returns true if the given user is suppressed, false otherwise
func (b *Repository) Suppressed(workspaceID, userID, sourceID string) (*model.Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add adds the given suppressions to the repository
func (b *Repository) Add(suppressions []model.Suppression, token []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// start the repository
func (b *Repository) start() (startErr error) { _ = "STUB: not implemented"; return nil }

// 16mb

// see https://dgraph.io/docs/badger/get-started/#garbage-collection

// Stop stops the repository
func (b *Repository) Stop() error { _ = "STUB: not implemented"; return nil }

// Backup writes a backup of the repository to the given writer
func (b *Repository) Backup(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Restore restores the repository from the given reader
func (b *Repository) Restore(r io.Reader) (err error) { _ = "STUB: not implemented"; return nil }

func (b *Repository) setRestoring(restoring bool) { _ = "STUB: not implemented"; return }

func (b *Repository) isRestoring() bool { _ = "STUB: not implemented"; return false }

type blogger struct {
	logger.Logger
}

func (l blogger) Warningf(fmt string, args ...any) { _ = "STUB: not implemented"; return }

func keyPrefix(workspaceID, userID string) string { _ = "STUB: not implemented"; return "" }

func getMetadataFromBadgerItem(item *badger.Item) (*model.Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// backwards compatibility
