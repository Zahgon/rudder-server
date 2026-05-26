package retltest

import (
	"sync"
	"testing"

	whUtil "github.com/rudderlabs/rudder-server/testhelper/webhook"
)

// webhook is a test helper for webhook destinations.
type webhook struct {
	Recorder *whUtil.Recorder

	once sync.Once
	id   string
	name string
}

func (w *webhook) ID() string { _ = "STUB: not implemented"; return "" }

func (w *webhook) Name() string { _ = "STUB: not implemented"; return "" }

func (w *webhook) TypeName() string { _ = "STUB: not implemented"; return "" }

func (w *webhook) Config() map[string]any { _ = "STUB: not implemented"; return nil }

func (w *webhook) Start(t *testing.T) { _ = "STUB: not implemented"; return }

func (w *webhook) Shutdown(*testing.T) { _ = "STUB: not implemented"; return }

func (w *webhook) Count() int { _ = "STUB: not implemented"; return 0 }
