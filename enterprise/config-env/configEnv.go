package configenv

import (
	"github.com/rudderlabs/rudder-go-kit/logger"
)

type HandleT struct {
	Log logger.Logger
}

var configEnvReplacer string

func loadConfig() { _ = "STUB: not implemented"; return }

// ReplaceConfigWithEnvVariables : Replaces all env variables in the config
func (h *HandleT) ReplaceConfigWithEnvVariables(workspaceConfig []byte) (updatedConfig []byte) {
	_ = "STUB: not implemented"
	return nil
}
