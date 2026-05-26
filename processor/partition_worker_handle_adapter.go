package processor

import (
	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/services/rsources"
)

// workerHandleAdapter is a wrapper around processor.Handle that implements the workerHandle interface
type workerHandleAdapter struct {
	*Handle
}

func (h *workerHandleAdapter) logger() logger.Logger {
	_ = "STUB: not implemented"
	return *new(logger.Logger)
}

func (h *workerHandleAdapter) config() workerHandleConfig {
	_ = "STUB: not implemented"
	return *new(workerHandleConfig)
}

func (h *workerHandleAdapter) rsourcesService() rsources.JobService {
	_ = "STUB: not implemented"
	return *new(rsources.JobService)
}

func (h *workerHandleAdapter) stats() *processorStats { _ = "STUB: not implemented"; return nil }
