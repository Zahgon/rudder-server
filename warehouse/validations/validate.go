package validations

import (
	"context"
	"encoding/json"

	"github.com/rudderlabs/rudder-go-kit/filemanager"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/warehouse/integrations/manager"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

type DestinationValidationResponse struct {
	Success bool          `json:"success"`
	Error   string        `json:"error"`
	Steps   []*model.Step `json:"steps"`
}

type Validator interface {
	Validate(ctx context.Context) error
}

type objectStorage struct {
	destination *backendconfig.DestinationT
}

type connections struct {
	manager     manager.WarehouseOperations
	destination *backendconfig.DestinationT
}

type createSchema struct {
	manager manager.WarehouseOperations
}

type createAlterTable struct {
	manager manager.WarehouseOperations
	table   string
}

type fetchSchema struct {
	manager     manager.WarehouseOperations
	destination *backendconfig.DestinationT
}

type loadTable struct {
	manager     manager.WarehouseOperations
	destination *backendconfig.DestinationT
	table       string
}

type DestinationValidator interface {
	Validate(ctx context.Context, dest *backendconfig.DestinationT) *DestinationValidationResponse
}

type destinationValidationImpl struct{}

func NewDestinationValidator() DestinationValidator {
	_ = "STUB: not implemented"
	return *new(DestinationValidator)
}

func (*destinationValidationImpl) Validate(ctx context.Context, dest *backendconfig.DestinationT) *DestinationValidationResponse {
	_ = "STUB: not implemented"
	return nil
}

func validateDestinationFunc(ctx context.Context, dest *backendconfig.DestinationT, stepToValidate string) (json.RawMessage, error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), nil
}

func validateDestination(ctx context.Context, dest *backendconfig.DestinationT, stepToValidate string) *DestinationValidationResponse {
	_ = "STUB: not implemented"
	return nil
}

// check if req has specified a step in query params

// get validation step

// Iterate over all selected steps and validate

// if any of steps fails, the whole validation fails

func NewValidator(ctx context.Context, step string, dest *backendconfig.DestinationT) (Validator, error) {
	_ = "STUB: not implemented"
	return *new(Validator), nil
}

func (os *objectStorage) Validate(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *connections) Validate(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (cs *createSchema) Validate(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (cat *createAlterTable) Validate(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (fs *fetchSchema) Validate(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (lt *loadTable) Validate(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// CreateTempLoadFile creates a temporary load file
func CreateTempLoadFile(dest *backendconfig.DestinationT) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func uploadFile(ctx context.Context, dest *backendconfig.DestinationT, filePath string) (filemanager.UploadedFile, error) {
	_ = "STUB: not implemented"
	return *new(filemanager.UploadedFile), nil
}

// cleanup

func deleteFile(ctx context.Context, dest *backendconfig.DestinationT, location string) error {
	_ = "STUB: not implemented"
	return nil
}

func downloadFile(ctx context.Context, dest *backendconfig.DestinationT, location string) error {
	_ = "STUB: not implemented"
	return nil
}

// cleanup

func createFileManager(dest *backendconfig.DestinationT) (filemanager.FileManager, error) {
	_ = "STUB: not implemented"
	return *new(filemanager.FileManager), nil
}

func createManager(ctx context.Context, dest *backendconfig.DestinationT) (manager.WarehouseOperations, error) {
	_ = "STUB: not implemented"
	return *new(manager.WarehouseOperations), nil
}

func createDummyWarehouse(dest *backendconfig.DestinationT) model.Warehouse {
	_ = "STUB: not implemented"
	return *new(model.Warehouse)
}

func configuredNamespaceInDestination(dest *backendconfig.DestinationT) string {
	_ = "STUB: not implemented"
	return ""
}

func getTable(dest *backendconfig.DestinationT) string { _ = "STUB: not implemented"; return "" }

func tableWithUUID() string { _ = "STUB: not implemented"; return "" }

type dummyUploader struct {
	dest *backendconfig.DestinationT
}

func (*dummyUploader) IsWarehouseSchemaEmpty() bool { _ = "STUB: not implemented"; return false }
func (*dummyUploader) GetLocalSchema(context.Context) (model.Schema, error) {
	_ = "STUB: not implemented"
	return *new(model.Schema), nil
}

func (*dummyUploader) UpdateLocalSchema(context.Context, model.Schema) error {
	_ = "STUB: not implemented"
	return nil
}
func (*dummyUploader) ShouldOnDedupUseNewRecord() bool { _ = "STUB: not implemented"; return false }
func (*dummyUploader) GetTableSchemaInWarehouse(string) model.TableSchema {
	_ = "STUB: not implemented"
	return *new(model.TableSchema)
}
func (*dummyUploader) GetTableSchemaInUpload(string) model.TableSchema {
	_ = "STUB: not implemented"
	return *new(model.TableSchema)
}
func (*dummyUploader) CanAppend() bool { _ = "STUB: not implemented"; return false }
func (*dummyUploader) GetSampleLoadFileLocation(context.Context, string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (*dummyUploader) GetLoadFilesMetadata(context.Context, warehouseutils.GetLoadFilesOptions) ([]warehouseutils.LoadFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*dummyUploader) GetSingleLoadFile(context.Context, string) (warehouseutils.LoadFile, error) {
	_ = "STUB: not implemented"
	return *new(warehouseutils.LoadFile), nil
}

func (m *dummyUploader) GetLoadFileType() string { _ = "STUB: not implemented"; return "" }

func (m *dummyUploader) UseRudderStorage() bool { _ = "STUB: not implemented"; return false }
