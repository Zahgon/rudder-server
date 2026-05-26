package encoding

// parquetLoader is used for generating parquet load files.
type parquetLoader struct {
	destType string
	Values   []any
	writer   LoadFileWriter
}

func newParquetLoader(w LoadFileWriter, destType string) *parquetLoader {
	_ = "STUB: not implemented"
	return nil
}

func (loader *parquetLoader) IsLoadTimeColumn(columnName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (*parquetLoader) GetLoadTimeFormat(_ string) string { _ = "STUB: not implemented"; return "" }

func (loader *parquetLoader) AddColumn(columnName, colType string, val any) {
	_ = "STUB: not implemented"
	return
}

func (*parquetLoader) AddRow(_, _ []string) {
	_ = "STUB: not implemented"
	// TODO : implement
	return
}

func (loader *parquetLoader) AddEmptyColumn(columnName string) { _ = "STUB: not implemented"; return }

func (*parquetLoader) WriteToString() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (loader *parquetLoader) Write() error { _ = "STUB: not implemented"; return nil }

func parquetValue(val any, colType string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func getInt64(val any) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func getBool(val any) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func getFloat64(val any) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func getUnixTimestamp(val any) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func getString(val any) (string, error) { _ = "STUB: not implemented"; return "", nil }
