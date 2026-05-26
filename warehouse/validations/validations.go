package validations

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rudderlabs/rudder-go-kit/filemanager"
	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
)

const (
	namespace = "rudderstack_setup_test"
	table     = "setup_test_staging"
)

var (
	connectionTestingFolder string
	pkgLogger               logger.Logger
	fileManagerFactory      filemanager.Factory
	objectStorageTimeout    time.Duration
	queryTimeout            time.Duration
)

var (
	tableSchemaMap = model.TableSchema{
		"id":  "int",
		"val": "string",
	}
	payloadMap = map[string]any{
		"id":  1,
		"val": "RudderStack",
	}
	alterColumnMap = model.TableSchema{
		"val_alter": "string",
	}
)

type validationFunc struct {
	Func func(context.Context, *backendconfig.DestinationT, string) (json.RawMessage, error)
}

func Init() { _ = "STUB: not implemented"; return }

// Since we have a cp-router default timeout of 30 seconds, keeping the query timeout to 25 seconds

// Validate the destination by running all the validation steps
func Validate(ctx context.Context, req *model.ValidationRequest) (*model.ValidationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validationFunctions() map[string]*validationFunc { _ = "STUB: not implemented"; return nil }
