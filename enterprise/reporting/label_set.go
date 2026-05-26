package reporting

import (
	"github.com/rudderlabs/rudder-server/utils/types"
)

type LabelSet struct {
	WorkspaceID             string
	SourceDefinitionID      string
	SourceCategory          string
	SourceID                string
	DestinationDefinitionID string
	DestinationID           string
	SourceTaskRunID         string
	SourceJobID             string
	SourceJobRunID          string
	TransformationID        string
	TransformationVersionID string
	TrackingPlanID          string
	TrackingPlanVersion     int
	InPU                    string
	PU                      string
	Status                  string
	TerminalState           bool
	InitialState            bool
	StatusCode              int
	EventName               string
	EventType               string
	ErrorType               string
	ErrorCode               string
	ErrorMessage            string
	Bucket                  int64
}

func NewLabelSet(metric types.PUReportedMetric, bucket int64) LabelSet {
	_ = "STUB: not implemented"
	return *new(LabelSet)
}

// NewLabelSetFromEDReportsDB creates a LabelSet from EDReportsDB for event sampling
func NewLabelSetFromEDReportsDB(metric types.EDReportsDB, bucket int64) LabelSet {
	_ = "STUB: not implemented"
	return *new(LabelSet)
}

// EDReportsDB doesn't have SourceCategory

// EDReportsDB doesn't have these fields

// EDReportsDB doesn't have InPU

// EDReportsDB doesn't have Status
// EDReportsDB doesn't have these fields

// EDReportsDB doesn't have ErrorType

func (labelSet LabelSet) generateHash() string { _ = "STUB: not implemented"; return "" }
