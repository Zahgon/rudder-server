package filemanagerutil

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/filemanager"
)

func GetProviderConfigForBackupsFromEnv(ctx context.Context, config *config.Config) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func ProviderConfigOpts(ctx context.Context, provider string, config *config.Config) filemanager.ProviderConfigOpts {
	_ = "STUB: not implemented"
	return *new(filemanager.ProviderConfigOpts)
}
