package filehandler

import (
	"context"
	"reflect"

	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/source"

	"github.com/rudderlabs/rudder-server/regulation-worker/internal/model"
)

type ParquetLocalFileHandler struct {
	records []any
	schema  []*parquet.SchemaElement
}

func NewParquetLocalFileHandler() *ParquetLocalFileHandler { _ = "STUB: not implemented"; return nil }

func (h *ParquetLocalFileHandler) Read(_ context.Context, path string) error {
	_ = "STUB: not implemented"
	return nil
}

// these entries will be filtered on in the `remove` stage.

// schema elements will be used to create writer instance for writing entries

func (h *ParquetLocalFileHandler) Write(_ context.Context, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *ParquetLocalFileHandler) RemoveIdentity(_ context.Context, attributes []model.User) error {
	_ = "STUB: not implemented"
	return nil
}

// As the records are struct and not pointer to struct
// , we need to create a copy of them using the below functions.
// In order for us to get value of the field, we need ptr to
// struct being passed.

// identity matched in the record, so this
// needs to be filtered out

func (*ParquetLocalFileHandler) identityMatched(recordValue reflect.Value, attribute *model.User) bool {
	_ = "STUB: not implemented"
	return false
}

// Only *string and string types are expected for the userId field
// In case anything else is found, return with an error to be used for warnings.

// pkgLogger.Debugf("unexpected data type for userId field: %v", reflect.ValueOf(userIdField).Kind())

func NewParquetReader(pFile source.ParquetFile, obj any, np int64) (*reader.ParquetReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// res.RenameSchema() // Stop from renaming the schema
