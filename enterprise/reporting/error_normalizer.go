package reporting

import (
	"context"
	"sync"
	"time"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/utils/types"
)

//go:generate mockgen -destination=../../../mocks/enterprise/reporting/mock_error_normalizer.go -package=mocks github.com/rudderlabs/rudder-server/enterprise/reporting ErrorNormalizer

// ErrorNormalizer defines the interface for normalizing error messages
// By normalizing errors, we can reduce the number of unique errors that are reported,
// This interface allows for easier testing by enabling mock implementations
type ErrorNormalizer interface {
	// NormalizeError returns the normalized error message for the given error
	// Returns "RedactedError" if rate limited, or the normalized error message if should be reported
	NormalizeError(ctx context.Context, errorDetailGroupKey types.ErrorDetailGroupKey, errorMessage string) string

	// StartCleanup starts the periodic cleanup routine
	StartCleanup(ctx context.Context) error
}

const (
	RedactedError = "RedactedError"
)

type errorEntry struct {
	message string
	time    time.Time
}

type boundedErrorSet struct {
	entries []*errorEntry          // oldest → newest
	index   map[string]*errorEntry // message → entry
}

func newBoundedErrorSet(max int) *boundedErrorSet { _ = "STUB: not implemented"; return nil }

// GetExact looks up exact match
func (s *boundedErrorSet) GetExact(msg string, now time.Time) (*errorEntry, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetSimilar scans for a near match (caller supplies similarity fn)
func (s *boundedErrorSet) GetSimilar(msg string, now time.Time, similar func(a, b string) bool) (*errorEntry, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Add inserts a new entry, evicting oldest if full
func (s *boundedErrorSet) Add(msg string, now time.Time) *errorEntry {
	_ = "STUB: not implemented"
	return nil
}

// DropStale removes messages older than cutoff
func (s *boundedErrorSet) DropStale(cutoff time.Time) { _ = "STUB: not implemented"; return }

// --- internal helpers ---

func (s *boundedErrorSet) moveToEnd(entry *errorEntry) { _ = "STUB: not implemented"; return }

func (s *boundedErrorSet) moveToEndAt(i int) { _ = "STUB: not implemented"; return }

type errorGroup struct {
	errors boundedErrorSet
	// lastUpdated is the time when the group's error set last changed (new error added or old error removed)
	lastUpdated time.Time
	// lastBlocked is the time when the error group was last blocked due to rate limiting
	lastBlocked time.Time
}

type errorNormalizer struct {
	mu     sync.Mutex
	log    logger.Logger
	groups map[types.ErrorDetailGroupKey]*errorGroup
	stats  stats.Stats

	similarityThreshold config.ValueLoader[float64]
	maxErrorsPerGroup   config.ValueLoader[int]
	maxGroups           config.ValueLoader[int]
	cleanupInterval     config.ValueLoader[time.Duration]
	staleTime           config.ValueLoader[time.Duration]

	// Stats manager reference
	statsManager *ErrorReportingStats
}

func NewErrorNormalizer(log logger.Logger, conf *config.Config, statsInstance stats.Stats, statsManager *ErrorReportingStats) ErrorNormalizer {
	_ = "STUB: not implemented"
	return *new(ErrorNormalizer)
}

// NormalizeError normalizes the error message for the given connection key
func (e *errorNormalizer) NormalizeError(ctx context.Context, errorDetailGroupKey types.ErrorDetailGroupKey, msg string) string {
	_ = "STUB: not implemented"
	return ""
}

// Exact match

// Similarity match

// Rate limit the new error message

// Insert the new error message

func (e *errorNormalizer) cleanup() { _ = "STUB: not implemented"; return }

// Drop if no messages left OR no new unique error for >staleTime

func (e *errorNormalizer) shouldDropCounter(currentGroup *errorGroup, now time.Time, maxErrorsPerGroup int) bool {
	_ = "STUB: not implemented"
	// Drop if no messages left
	return false
}

// Drop if counter has reached maxErrorsPerGroup and no changes to errors for >staleTime and
// an error has been blocked in the last staleTime, means this counter is starving other errors

func (e *errorNormalizer) StartCleanup(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
