package identity

import (
	"context"

	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/utils/misc"
	"github.com/rudderlabs/rudder-server/warehouse/encoding"
	sqlmiddleware "github.com/rudderlabs/rudder-server/warehouse/integrations/middleware/sqlquerywrapper"
	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
	"github.com/rudderlabs/rudder-server/warehouse/internal/service/loadfiles/downloader"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

var pkgLogger logger.Logger

func init() {
	pkgLogger = logger.NewLogger().Child("warehouse").Child("identity")
}

type WarehouseManager interface {
	DownloadIdentityRules(context.Context, *misc.GZipWriter) error
}

type Identity struct {
	warehouse        model.Warehouse
	db               *sqlmiddleware.DB
	uploader         warehouseutils.Uploader
	uploadID         int64
	warehouseManager WarehouseManager
	downloader       downloader.Downloader
	encodingFactory  *encoding.Factory
}

func New(warehouse model.Warehouse, db *sqlmiddleware.DB, uploader warehouseutils.Uploader, uploadID int64, warehouseManager WarehouseManager, loadFileDownloader downloader.Downloader, encodingFactory *encoding.Factory) *Identity {
	_ = "STUB: not implemented"
	return nil
}

func (idr *Identity) mergeRulesTable() string { _ = "STUB: not implemented"; return "" }

func (idr *Identity) mappingsTable() string { _ = "STUB: not implemented"; return "" }

func (idr *Identity) whMergeRulesTable() string { _ = "STUB: not implemented"; return "" }

func (idr *Identity) whMappingsTable() string { _ = "STUB: not implemented"; return "" }

func (idr *Identity) applyRule(txn *sqlmiddleware.Tx, ruleID int64, gzWriter *misc.GZipWriter) (totalRowsModified int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// if no rudder_id is found with properties in merge_rule, create a new one
// else if only one rudder_id is found with properties in merge_rule, use that rudder_id
// else create a new rudder_id and assign it to all properties found with properties in the merge_rule

// generate new one and assign to these two

// generate new one and update all

// TODO : support add row for parquet loader

func (idr *Identity) addRules(txn *sqlmiddleware.Tx, loadFileNames []string, gzWriter *misc.GZipWriter) (ids []int64, err error) {
	_ = "STUB: not implemented"
	// add rules from load files into temp table
	// use original table to delete redundant ones from temp table
	// insert from temp table into original table
	return nil, nil
}

// add rowID which allows us to insert in same order from staging to original merge _rules table

// write merge rules to file to be uploaded to warehouse in later steps

// select and insert distinct combination of merge rules and sort them by order in which they were added

func (idr *Identity) writeTableToFile(tableName string, txn *sqlmiddleware.Tx, gzWriter *misc.GZipWriter) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// TODO : use proper column type here

func (idr *Identity) uploadFile(ctx context.Context, filePath string, txn *sqlmiddleware.Tx, tableName string, totalRecords int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (idr *Identity) createTempGzFile(dirName string) (gzWriter misc.GZipWriter, path string) {
	_ = "STUB: not implemented"
	return *new(misc.GZipWriter), ""
}

func (idr *Identity) processMergeRules(ctx context.Context, fileNames []string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// START: Add new merge rules to local pg table and also to file

// END: Add new merge rules to local pg table and also to file

// START: Add new/changed identity mappings to local pg table and also to file

// END: Add new/changed identity mappings to local pg table and also to file

// upload new merge rules to object storage

// upload new/changed identity mappings to object storage

// Resolve does the below things in a single pg txn
// 1. Fetch all new merge rules added in the upload
// 2. Append to local identity merge rules table
// 3. Apply each merge rule and update local identity mapping table
// 4. Upload the diff of each table to load files for both tables
func (idr *Identity) Resolve(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (idr *Identity) ResolveHistoricIdentities(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}
