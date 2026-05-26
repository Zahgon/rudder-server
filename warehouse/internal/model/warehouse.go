package model

import (
	"github.com/rudderlabs/rudder-go-kit/config"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

type Warehouse struct {
	WorkspaceID string
	Source      backendconfig.SourceT
	Destination backendconfig.DestinationT
	Namespace   string
	Type        string
	Identifier  string
}

func (w *Warehouse) GetBoolDestinationConfig(key DestinationConfigSetting) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *Warehouse) GetStringDestinationConfig(conf *config.Config, key DestinationConfigSetting) string {
	_ = "STUB: not implemented"
	return ""
}

func (w *Warehouse) GetMapDestinationConfig(key DestinationConfigSetting) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func (w *Warehouse) GetPreferAppendSetting() bool { _ = "STUB: not implemented"; return false }

// when the value is not defined it is important we choose the right default
// in order to maintain backward compatibility for existing destinations
// defaulting to true for BQ, false for other destination types

func (w *Warehouse) IsEnabled() bool { _ = "STUB: not implemented"; return false }
