package service

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/logger"
)

var pkgLogger = logger.NewLogger().Child("service")

type Looper struct {
	Svc JobSvc
}

func (l *Looper) Loop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// this is to make sure that we don't panic when any of the API call fails with deadline exceeded error.

func getenvInt(key string, fallback int) (int, error) { _ = "STUB: not implemented"; return 0, nil }
