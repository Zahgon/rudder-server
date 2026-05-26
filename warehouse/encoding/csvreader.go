package encoding

import (
	"encoding/csv"
	"io"
)

type csvReader struct {
	reader *csv.Reader
}

func (csv *csvReader) Read([]string) (record []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newCsvReader(r io.Reader) *csvReader { _ = "STUB: not implemented"; return nil }
