package encoding

import (
	"os"

	"github.com/xitongsys/parquet-go/writer"

	"github.com/rudderlabs/rudder-server/utils/misc"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

const (
	parquetInt64           = "type=INT64, repetitiontype=OPTIONAL"
	parquetBoolean         = "type=BOOLEAN, repetitiontype=OPTIONAL"
	parquetDouble          = "type=DOUBLE, repetitiontype=OPTIONAL"
	parquetString          = "type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"
	parquetTimestampMicros = "type=INT64, convertedtype=TIMESTAMP_MICROS, repetitiontype=OPTIONAL"
)

var rudderDataTypeToParquetDataType = map[string]map[string]string{
	warehouseutils.RS: {
		"bigint":   parquetInt64,
		"int":      parquetInt64,
		"boolean":  parquetBoolean,
		"float":    parquetDouble,
		"string":   parquetString,
		"text":     parquetString,
		"datetime": parquetTimestampMicros,
	},
	warehouseutils.S3Datalake: {
		"bigint":   parquetInt64,
		"int":      parquetInt64,
		"boolean":  parquetBoolean,
		"float":    parquetDouble,
		"string":   parquetString,
		"text":     parquetString,
		"datetime": parquetTimestampMicros,
	},
	warehouseutils.GCSDatalake: {
		"int":      parquetInt64,
		"boolean":  parquetBoolean,
		"float":    parquetDouble,
		"string":   parquetString,
		"datetime": parquetTimestampMicros,
	},
	warehouseutils.AzureDatalake: {
		"int":      parquetInt64,
		"boolean":  parquetBoolean,
		"float":    parquetDouble,
		"string":   parquetString,
		"datetime": parquetTimestampMicros,
	},
	warehouseutils.DELTALAKE: {
		"int":      parquetInt64,
		"boolean":  parquetBoolean,
		"float":    parquetDouble,
		"string":   parquetString,
		"datetime": parquetTimestampMicros,
	},
}

type parquetWriter struct {
	writer     *writer.CSVWriter
	fileWriter misc.BufferedWriter
}

func createParquetWriter(outputFilePath string, schema model.TableSchema, destType string, maxParallelWriters int64, disableParquetColumnIndex bool) (LoadFileWriter, error) {
	_ = "STUB: not implemented"
	return *new(LoadFileWriter), nil
}

// Disable column index to avoid the column index being written to the parquet file.

func (p *parquetWriter) WriteRow(row []any) error { _ = "STUB: not implemented"; return nil }

func (p *parquetWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (*parquetWriter) WriteGZ(_ string) error { _ = "STUB: not implemented"; return nil }

func (*parquetWriter) Write(_ []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *parquetWriter) GetLoadFile() *os.File { _ = "STUB: not implemented"; return nil }

func sortedTableColumns(schema model.TableSchema) []string { _ = "STUB: not implemented"; return nil }

func parquetSchema(schema model.TableSchema, destType string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
