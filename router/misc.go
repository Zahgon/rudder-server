package router

import (
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	kitsync "github.com/rudderlabs/rudder-go-kit/sync"

	"github.com/rudderlabs/rudder-server/processor/integrations"
	"github.com/rudderlabs/rudder-server/router/internal/eventorder"
	"github.com/rudderlabs/rudder-server/router/isolation"
)

func isSuccessStatus(status int) bool { _ = "STUB: not implemented"; return false }

func isJobTerminated(status int) bool { _ = "STUB: not implemented"; return false }

func nextAttemptAfter(attempt int, minRetryBackoff, maxRetryBackoff time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func getIterableStruct(payload []byte, transformAt string) ([]integrations.PostParametersT, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getWorkerPartition(key eventorder.BarrierKey, noOfWorkers int) int {
	_ = "STUB: not implemented"
	return 0
}

func isolationMode(destType string, config *config.Config) isolation.Mode {
	_ = "STUB: not implemented"
	return *new(isolation.Mode)
}

func LimiterPriorityValueFrom(v, max int) kitsync.LimiterPriorityValue {
	_ = "STUB: not implemented"
	return *new(kitsync.LimiterPriorityValue)
}
