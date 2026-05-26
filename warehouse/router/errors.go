package router

import (
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
)

type errorMapper interface {
	ErrorMappings() []model.JobError
}

type ErrorHandler struct {
	Mapper errorMapper
}

// MatchUploadJobErrorType matches the error with the error mappings defined in the integrations
// and returns the corresponding matched error type else returns UncategorizedError
func (e *ErrorHandler) MatchUploadJobErrorType(err error) model.JobErrorType {
	_ = "STUB: not implemented"
	return *new(model.JobErrorType)
}
