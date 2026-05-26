package tablearchiver

import (
	"database/sql"

	"github.com/rudderlabs/rudder-go-kit/filemanager"
	"github.com/rudderlabs/rudder-go-kit/logger"
)

const (
	PaginationAction = "{{.Pagination}}"
	OffsetAction     = "{{.Offset}}"
)

type TableJSONArchiver struct {
	DbHandle      *sql.DB
	Pagination    int
	Offset        int
	QueryTemplate string
	OutputPath    string
	FileManager   filemanager.FileManager
}

var pkgLogger logger.Logger

func init() {
	pkgLogger = logger.NewLogger().Child("tablearchiver")
}

func (jsonArchiver *TableJSONArchiver) Do() (location string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// break when json is null

// replacing ", \n " with "\n"
// stripping starting '[' and ending ']'
// appending '\n'
