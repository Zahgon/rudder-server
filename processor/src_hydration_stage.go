package processor

import (
	"context"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/processor/types"
	reportingtypes "github.com/rudderlabs/rudder-server/utils/types"
)

type srcHydrationMessage struct {
	partition                     string
	subJobs                       subJob
	archivalJobs                  []*jobsdb.JobT
	eventSchemaJobsBySourceId     map[SourceIDT][]*jobsdb.JobT
	connectionDetailsMap          map[string]*reportingtypes.ConnectionDetails
	statusDetailsMap              map[string]map[string]*reportingtypes.StatusDetail
	enricherStatusDetailsMap      map[string]map[string]*reportingtypes.StatusDetail
	botManagementStatusDetailsMap map[string]map[string]*reportingtypes.StatusDetail
	eventBlockingStatusDetailsMap map[string]map[string]*reportingtypes.StatusDetail
	destFilterStatusDetailMap     map[string]map[string]*reportingtypes.StatusDetail
	reportMetrics                 []*reportingtypes.PUReportedMetric
	totalEvents                   int
	groupedEventsBySourceId       map[SourceIDT][]types.TransformerEvent
	eventsByMessageID             map[string]types.SingularEventWithReceivedAt
	jobIDToSpecificDestMapOnly    map[int64]string
	statusList                    []*jobsdb.JobStatusT
	jobList                       []*jobsdb.JobT
	sourceDupStats                map[dupStatKey]int
	dedupKeys                     map[string]struct{}
}

func (proc *Handle) srcHydrationStage(partition string, message *srcHydrationMessage) (*preTransformationMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Process sources in parallel using errgroup

// Mutex to protect shared maps

// report metrics for failed hydration

// Update shared maps with mutex protection

// Append hydration failed reports if any

// If no hydrated jobs, remove entry from groupedEventsBySourceId

// Update eventsByMessageID map

// Update the groupedEventsBySourceId map

// Update eventSchemaJobsBySourceId map if needed

// Wait for all goroutines to complete and check for errors

func (proc *Handle) hydrate(ctx context.Context, source *backendconfig.SourceT, events []types.TransformerEvent) ([]types.TransformerEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (proc *Handle) getHydrationFailedReports(source *backendconfig.SourceT, jobs []types.TransformerEvent, err error) []*reportingtypes.PUReportedMetric {
	_ = "STUB: not implemented"
	return nil
}
