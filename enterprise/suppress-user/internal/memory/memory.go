package memory

import (
	"io"
	"sync"

	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/enterprise/suppress-user/model"
)

// Repository is a repository backed by memory.

type Repository struct {
	log            logger.Logger
	token          []byte
	suppressionsMu sync.RWMutex
	suppressions   map[string]map[string]map[string]model.Metadata
}

// NewRepository returns a new repository backed by memory.
func NewRepository(log logger.Logger) *Repository { _ = "STUB: not implemented"; return nil }

// GetToken returns the current token
func (m *Repository) GetToken() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Suppressed returns true if the given user is suppressed, false otherwise
		nil
}

func (m *Repository) Suppressed(workspaceID, userID, sourceID string) (*model.Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add adds the given suppressions to the repository
func (m *Repository) Add(suppressions []model.Suppression, token []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop is a no-op for the memory repository.
func (*Repository) Stop() error {
	_ = "STUB: not implemented"

	// Backup is not supported for the memory repository.
	return nil
}

func (*Repository) Backup(_ io.Writer) error { _ = "STUB: not implemented"; return nil }

// Restore is not supported for the memory repository.
func (*Repository) Restore(_ io.Reader) error { _ = "STUB: not implemented"; return nil }
