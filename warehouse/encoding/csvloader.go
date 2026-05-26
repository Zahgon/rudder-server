package encoding

import (
	"bytes"
	"encoding/csv"
)

// csvLoader is common for non-BQ warehouses.
// If you need any custom logic, either extend this or use destType and if/else/switch.
type csvLoader struct {
	destType   string
	csvRow     []string
	buff       bytes.Buffer
	csvWriter  *csv.Writer
	fileWriter LoadFileWriter
}

func newCSVLoader(writer LoadFileWriter, destType string) *csvLoader {
	_ = "STUB: not implemented"
	return nil
}

func (loader *csvLoader) IsLoadTimeColumn(columnName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (*csvLoader) GetLoadTimeFormat(string) string { _ = "STUB: not implemented"; return "" }

func (loader *csvLoader) AddColumn(_, _ string, val any) { _ = "STUB: not implemented"; return }

func (loader *csvLoader) AddRow(_, row []string) { _ = "STUB: not implemented"; return }

func (loader *csvLoader) AddEmptyColumn(columnName string) { _ = "STUB: not implemented"; return }

func (loader *csvLoader) WriteToString() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (loader *csvLoader) Write() error { _ = "STUB: not implemented"; return nil }
