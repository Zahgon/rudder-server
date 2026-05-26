package log

import (
	"github.com/rudderlabs/rudder-go-kit/logger"
)

var GinkgoLogger logger.Logger = &ginkgoLogger{logger.NOP}

type ginkgoLogger struct {
	logger.Logger
}

func (ginkgoLogger) Debug(args ...any)                 { _ = "STUB: not implemented"; return }
func (ginkgoLogger) Info(args ...any)                  { _ = "STUB: not implemented"; return }
func (ginkgoLogger) Warn(args ...any)                  { _ = "STUB: not implemented"; return }
func (ginkgoLogger) Error(args ...any)                 { _ = "STUB: not implemented"; return }
func (ginkgoLogger) Fatal(args ...any)                 { _ = "STUB: not implemented"; return }
func (ginkgoLogger) Debugf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (ginkgoLogger) Infof(format string, args ...any)  { _ = "STUB: not implemented"; return }
func (ginkgoLogger) Warnf(format string, args ...any)  { _ = "STUB: not implemented"; return }
func (ginkgoLogger) Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (ginkgoLogger) Fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (ginkgoLogger) Debugw(format string, args ...any) { _ = "STUB: not implemented"; return }

func (ginkgoLogger) Infow(msg string, keysAndValues ...any) { _ = "STUB: not implemented"; return }

func (ginkgoLogger) Warnw(msg string, keysAndValues ...any) { _ = "STUB: not implemented"; return }

func (ginkgoLogger) Errorw(msg string, keysAndValues ...any) { _ = "STUB: not implemented"; return }

func (ginkgoLogger) Fatalw(msg string, keysAndValues ...any) { _ = "STUB: not implemented"; return }

func (ginkgoLogger) With(_ ...any) logger.Logger {
	_ = "STUB: not implemented"
	return *new(logger.Logger)
}
func (ginkgoLogger) Child(_ string) logger.Logger {
	_ = "STUB: not implemented"
	return *new(logger.Logger)
}
func (ginkgoLogger) IsDebugLevel() bool { _ = "STUB: not implemented"; return false }
