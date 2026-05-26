package reporting

import (
	"github.com/rudderlabs/rudder-server/utils/types"
)

func (edr *ErrorDetailReporter) mergeMetricGroupsByErrorMessage(metricGroups map[types.ErrorDetailGroupKey][]*types.EDReportsDB) map[types.ErrorDetailGroupKey][]*types.EDReportsDB {
	_ = "STUB: not implemented"
	return nil
}

// groupMetricsByConnection groups metrics by source, destination, PU, and event type
func (edr *ErrorDetailReporter) groupByConnection(metrics []*types.EDReportsDB) map[types.ErrorDetailGroupKey][]*types.EDReportsDB {
	_ = "STUB: not implemented"
	return nil
}

// generateMetricGroupKey creates a unique key for connection details
func (edr *ErrorDetailReporter) generateMetricGroupKey(metric *types.EDReportsDB) types.ErrorDetailGroupKey {
	_ = "STUB: not implemented"
	return *new(types.ErrorDetailGroupKey)
}
