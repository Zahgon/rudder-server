package downloader

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/filemanager"

	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

type Downloader interface {
	Download(ctx context.Context, tableName string) ([]string, error)
}

type downloaderImpl struct {
	warehouse  *model.Warehouse
	uploader   warehouseutils.Uploader
	numWorkers int
}

func NewDownloader(
	warehouse *model.Warehouse,
	uploader warehouseutils.Uploader,
	numWorkers int,
) Downloader {
	_ = "STUB: not implemented"
	return *new(Downloader)
}

func (l *downloaderImpl) Download(ctx context.Context, tableName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *downloaderImpl) downloadSingleObject(ctx context.Context, fileManager filemanager.FileManager, object warehouseutils.LoadFile) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
