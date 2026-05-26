package suppression

import (
	"io"

	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/enterprise/suppress-user/internal/badgerdb"
	"github.com/rudderlabs/rudder-server/enterprise/suppress-user/model"
)

// Repository provides a generic interface for managing user suppressions
type Repository interface {
	// Stop stops the repository
	Stop() error

	// GetToken returns the current token
	GetToken() ([]byte, error)

	// Add adds the given suppressions to the repository
	Add(suppressions []model.Suppression, token []byte) error

	// Suppressed returns true if the given user is suppressed, false otherwise
	Suppressed(workspaceID, userID, sourceID string) (*model.Metadata, error)

	// Backup writes a backup of the repository to the given writer
	Backup(w io.Writer) error

	// Restore restores the repository from the given reader
	Restore(r io.Reader) error
}

// NewMemoryRepository returns a new repository backed by memory.
func NewMemoryRepository(log logger.Logger) Repository {
	_ = "STUB: not implemented"
	return *new(Repository)
}

var (
	WithSeederSource = badgerdb.WithSeederSource
	WithMaxSeedWait  = badgerdb.WithMaxSeedWait
)

// NewBadgerRepository returns a new repository backed by badgerDB.
func NewBadgerRepository(path string, log logger.Logger, opts ...badgerdb.Opt) (Repository, error) {
	_ = "STUB: not implemented"
	return *new(Repository), nil
}
