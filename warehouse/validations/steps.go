package validations

import (
	"context"
	"encoding/json"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
)

func validateStepFunc(_ context.Context, destination *backendconfig.DestinationT, _ string) (json.RawMessage, error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), nil
}

func StepsToValidate(dest *backendconfig.DestinationT) *model.StepsResponse {
	_ = "STUB: not implemented"
	return nil
}

// No additional steps
