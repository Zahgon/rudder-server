package processor

import (
	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/processor/types"
)

func (proc *Handle) getDroppedJobs(response types.Response, eventsToTransform []types.TransformerEvent) []*jobsdb.JobT {
	_ = "STUB: not implemented"
	// each messageID is one event when sending to the transformer
	return nil
}

// in transformer response, multiple messageIDs could be batched together

// for failed as well

// the remainder of the messageIDs are those that are dropped
// we get jobs for those dropped messageIDs - for rsources_stats_collector
