package encoding

import (
	"io"
	"os"

	"github.com/rudderlabs/rudder-go-kit/config"

	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
)

const (
	UUIDTsColumn     = "uuid_ts"
	LoadedAtColumn   = "loaded_at"
	BQLoadedAtFormat = "2006-01-02 15:04:05.999999 Z"
	BQUuidTSFormat   = "2006-01-02 15:04:05 Z"
)

type Factory struct {
	config struct {
		maxStagingFileReadBufferCapacityInK int
		parquetParallelWriters              config.ValueLoader[int64]
		disableParquetColumnIndex           config.ValueLoader[bool]
	}
}

func NewFactory(conf *config.Config) *Factory { _ = "STUB: not implemented"; return nil }

// LoadFileWriter is an interface for writing events to a load file
type LoadFileWriter interface {
	WriteGZ(s string) error
	Write(p []byte) (int, error)
	WriteRow(r []any) error
	Close() error
	GetLoadFile() *os.File
}

func (m *Factory) NewLoadFileWriter(loadFileType, outputFilePath string, schema model.TableSchema, destType string) (LoadFileWriter, error) {
	_ = "STUB: not implemented"
	return *new(LoadFileWriter), nil
}

// EventLoader is an interface for loading events into a load file
// It's used to load singular BatchRouterEvent events into a load file
type EventLoader interface {
	IsLoadTimeColumn(columnName string) bool
	GetLoadTimeFormat(columnName string) string
	AddColumn(columnName, columnType string, val any)
	AddRow(columnNames, values []string)
	AddEmptyColumn(columnName string)
	WriteToString() (string, error)
	Write() error
}

func (m *Factory) NewEventLoader(w LoadFileWriter, loadFileType, destinationType string) EventLoader {
	_ = "STUB: not implemented"
	return *new(EventLoader)
}

// EventReader is an interface for reading events from a load file
type EventReader interface {
	Read(columnNames []string) (record []string, err error)
}

func (m *Factory) NewEventReader(r io.Reader, destType string) EventReader {
	_ = "STUB: not implemented"
	return *new(EventReader)
}
