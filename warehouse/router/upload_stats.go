package router

import (
	"github.com/rudderlabs/rudder-go-kit/stats"

	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

const moduleName = "warehouse"

func warehouseTagName(destID, sourceName, destName, sourceID string) string {
	_ = "STUB: not implemented"
	return ""
}

func (job *UploadJob) buildTags(extraTags ...warehouseutils.Tag) stats.Tags {
	_ = "STUB: not implemented"
	return *new(stats.Tags)
}

func (job *UploadJob) timerStat(name string, extraTags ...warehouseutils.Tag) stats.Timer {
	_ = "STUB: not implemented"
	return *new(stats.Timer)
}

func (job *UploadJob) counterStat(name string, extraTags ...warehouseutils.Tag) stats.Counter {
	_ = "STUB: not implemented"
	return *new(stats.Counter)
}

func (job *UploadJob) gaugeStat(name string, extraTags ...warehouseutils.Tag) stats.Gauge {
	_ = "STUB: not implemented"
	return *new(stats.Gauge)
}

func (job *UploadJob) histogramStat(name string, extraTags ...warehouseutils.Tag) stats.Histogram {
	_ = "STUB: not implemented"
	return *new(stats.Histogram)
}

func (job *UploadJob) generateUploadSuccessMetrics() { _ = "STUB: not implemented"; return }

func (job *UploadJob) generateUploadAbortedMetrics() { _ = "STUB: not implemented"; return }

func (job *UploadJob) recordTableLoad(tableName string, numEvents int64) {
	_ = "STUB: not implemented"
	return
}

func (job *UploadJob) recordLoadFileGenerationTimeStat(startID, endID int64) error {
	_ = "STUB: not implemented"
	return nil
}
