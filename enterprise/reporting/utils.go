package reporting

import (
	"encoding/json"

	"github.com/rudderlabs/rudder-server/enterprise/reporting/event_sampler"
	"github.com/rudderlabs/rudder-server/utils/types"
)

func floorFactor(intervalMs int64) int64 { _ = "STUB: not implemented"; return 0 }

// Find the smallest index where factors[i] >= intervalMs

// If index is 0, intervalMs is smaller than the smallest factor

// If factors[index] == intervalMs, return it directly

// Otherwise, return the previous factor

func GetAggregationBucketMinute(timeMs, intervalMs int64) (int64, int64) {
	_ = "STUB: not implemented"
	// If interval is not a factor of 60, then the bucket start will not be aligned to hour start
	// For example, if intervalMs is 7, and timeMs is 28891085 (6:05) then the bucket start will be 28891079 (5:59)
	// and current bucket will contain the data of 2 different hourly buckets, which is should not have happened.
	// To avoid this, we round the intervalMs to the nearest factor of 60.
	return 0, 0
}

func getStringifiedSampleEvent(rawSampleEvent json.RawMessage) string {
	_ = "STUB: not implemented"
	return ""
}

// getSampleWithEventSamplingCore contains the common event sampling logic
func getSampleWithEventSamplingCore(sampleEvent json.RawMessage, sampleResponse string, eventSampler event_sampler.EventSampler, eventSamplingEnabled bool, hashGenerator func() string) (json.RawMessage, string, error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), "", nil
}

func getSampleWithEventSampling(metric types.PUReportedMetric, reportedAt int64, eventSampler event_sampler.EventSampler, eventSamplingEnabled bool, eventSamplingDuration int64) (sampleEvent json.RawMessage, sampleResponse string, err error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), "", nil
}

// getSampleWithEventSamplingForEDReportsDB applies event sampling to EDReportsDB metrics
// It reuses the common logic from getSampleWithEventSampling but works with EDReportsDB
func getSampleWithEventSamplingForEDReportsDB(metric types.EDReportsDB, reportedAt int64, eventSampler event_sampler.EventSampler, eventSamplingEnabled bool, eventSamplingDuration int64) (sampleEvent json.RawMessage, sampleResponse string, err error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), "", nil
}

func transformMetricForPII(metric types.PUReportedMetric, piiColumns []string) types.PUReportedMetric {
	_ = "STUB: not implemented"
	return *new(types.PUReportedMetric)
}

func getPIIColumnsToExclude() []string { _ = "STUB: not implemented"; return nil }
