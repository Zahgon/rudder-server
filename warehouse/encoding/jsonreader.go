package encoding

import (
	"bufio"
	"io"
)

type jsonReader struct {
	scanner *bufio.Scanner
}

func (js *jsonReader) Read(columnNames []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newJSONReader returns a new JSON reader
// default scanner buffer maxCapacity is 64K
// set it to higher value to avoid read stop on read size error
func newJSONReader(r io.Reader, bufferCapacityInK int) *jsonReader {
	_ = "STUB: not implemented"
	return nil
}
