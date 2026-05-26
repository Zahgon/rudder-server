package processor

import (
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/processor/types"
	reportingtypes "github.com/rudderlabs/rudder-server/utils/types"
)

type TrackingPlanStatT struct {
	numEvents                   stats.Measurement
	numValidationSuccessEvents  stats.Measurement
	numValidationFailedEvents   stats.Measurement
	numValidationFilteredEvents stats.Measurement
	tpValidationTime            stats.Measurement
}

// reportViolations It is going add violationErrors in context depending upon certain criteria:
// 1. sourceSchemaConfig in Metadata.MergedTpConfig should be true
func reportViolations(validateEvent *types.TransformerResponse, trackingPlanID string, trackingPlanVersion int) {
	_ = "STUB: not implemented"
	return
}

// enhanceWithViolation It enhances extra information of ValidationErrors in context for:
// 1. response.Events
// 1. response.FailedEvents
func enhanceWithViolation(response types.Response, trackingPlanID string, trackingPlanVersion int) {
	_ = "STUB: not implemented"
	return
}

// validateEvents If the TrackingPlanId exist for a particular write key then we are going to Validate from the transformer.
// The Response will contain both the Events and FailedEvents
// 1. eventsToTransform gets added to validatedEventsBySourceId
func (proc *Handle) validateEvents(groupedEventsBySourceId map[SourceIDT][]types.TransformerEvent, eventsByMessageID map[string]types.SingularEventWithReceivedAt, srcHydrationEnabledMap map[SourceIDT]bool) (map[SourceIDT][]types.TransformerEvent, []*reportingtypes.PUReportedMetric, sourceIDPipelineSteps) {
	_ = "STUB: not implemented"
	return nil, nil, *new(sourceIDPipelineSteps)
}

// pass on the jobs for transformation(User, Dest)

// If transformerInput does not match with transformerOutput then we do not consider transformerOutput
// This is a safety check we are adding so that if something unexpected comes from transformer
// We are ignoring it.

// Set sourcePipelineSteps.trackingPlanValidation for the sourceID to true.
// This is being used to distinguish the flows in reporting service

// Note: Sending false for usertransformation enabled is safe because this stage is before user transformation.

// REPORTING - START

// There will be no diff metrics for tracking plan validation

// REPORTING - END

// newValidationStat Creates a new TrackingPlanStatT instance
func (proc *Handle) newValidationStat(metadata *types.Metadata) *TrackingPlanStatT {
	_ = "STUB: not implemented"
	return nil
}
