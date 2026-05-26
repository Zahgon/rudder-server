package suppression

import (
	"io"
	"sync"

	"github.com/rudderlabs/rudder-server/enterprise/suppress-user/model"
)

// RepoSwitcher is a repository that can be be used to switch repository at runtime
type RepoSwitcher struct {
	Repository
	mu sync.RWMutex
}

func (rh *RepoSwitcher) Stop() error { _ = "STUB: not implemented"; return nil }

func (rh *RepoSwitcher) GetToken() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (rh *RepoSwitcher) Add(suppressions []model.Suppression, token []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (rh *RepoSwitcher) Suppressed(workspaceID, userID, sourceID string) (*model.Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rh *RepoSwitcher) Backup(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (rh *RepoSwitcher) Restore(r io.Reader) error { _ = "STUB: not implemented"; return nil }

func (rh *RepoSwitcher) Switch(newRepo Repository) { _ = "STUB: not implemented"; return }
