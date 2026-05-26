package warehouseutils

var statsSupportedTableNames = map[string]struct{}{
	"tracks":     {},
	"identifies": {},
	"users":      {},
	"pages":      {},
	"screens":    {},
	"aliases":    {},
	"groups":     {},
}

func TableNameForStats(tableName string) string { _ = "STUB: not implemented"; return "" }

// making all other tableName as other, to reduce cardinality
