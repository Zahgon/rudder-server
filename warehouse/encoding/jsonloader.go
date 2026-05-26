package encoding

type jsonLoader struct {
	destType   string
	columnData map[string]any
	fileWriter LoadFileWriter
}

// newJSONLoader returns a new jsonLoader is only for BQ now. Treat this is as custom BQ loader.
// If more warehouses are added in the future, change this accordingly.
func newJSONLoader(writer LoadFileWriter, destType string) *jsonLoader {
	_ = "STUB: not implemented"
	return nil
}

func (loader *jsonLoader) IsLoadTimeColumn(columnName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (loader *jsonLoader) GetLoadTimeFormat(columnName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (loader *jsonLoader) AddColumn(columnName, _ string, val any) {
	_ = "STUB: not implemented"
	return
}

func (loader *jsonLoader) AddRow(columnNames, row []string) { _ = "STUB: not implemented"; return }

func (loader *jsonLoader) AddEmptyColumn(columnName string) { _ = "STUB: not implemented"; return }

func (loader *jsonLoader) WriteToString() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (loader *jsonLoader) Write() error { _ = "STUB: not implemented"; return nil }
