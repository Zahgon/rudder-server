package archiver

import (
	"context"
	"time"

	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"
	kitsync "github.com/rudderlabs/rudder-go-kit/sync"

	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/services/fileuploader"
	"github.com/rudderlabs/rudder-server/utils/payload"
)

type worker struct {
	log              logger.Logger
	sourceID         string
	archiveFrom      string
	jobsDB           jobsdb.JobsDB
	payloadLimitFunc payload.AdaptiveLimiterFunc
	storageProvider  fileuploader.Provider
	stats            stats.Stats
	lifecycle        struct {
		ctx    context.Context
		cancel context.CancelFunc
	}

	fetchLimiter, uploadLimiter, updateLimiter kitsync.Limiter

	config struct {
		payloadLimit     func() int64
		jobsdbMaxRetries func() int
		instanceID       string
		eventsLimit      func() int
		minSleep         time.Duration
		uploadFrequency  time.Duration
	}
	lastUploadTime time.Time
	queryParams    jobsdb.GetQueryParams
}

func (w *worker) Work() bool { _ = "STUB: not implemented"; return false }

// respect the upload frequency

func (w *worker) SleepDurations() (min, max time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Duration)
}

func (w *worker) Stop() { _ = "STUB: not implemented"; return }

func (w *worker) uploadJobs(ctx context.Context, jobs []*jobsdb.JobT) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (w *worker) getJobs() ([]*jobsdb.JobT, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func marshalJob(job *jobsdb.JobT) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (w *worker) markStatus(jobs []*jobsdb.JobT, state string, response []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func errJSON(err error) []byte { _ = "STUB: not implemented"; return nil }

func locationJSON(location string) []byte { _ = "STUB: not implemented"; return nil }
