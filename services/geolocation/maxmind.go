package geolocation

import (
	"errors"

	"github.com/oschwald/maxminddb-golang"
)

var (
	ErrInvalidDatabase = errors.New("invalid database file")
	ErrInvalidIP       = errors.New("ip for lookup cannot be empty or invalid")
)

type maxmindDBReader struct {
	*maxminddb.Reader
}

func NewMaxmindDBReader(dbLoc string) (*maxmindDBReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *maxmindDBReader) Locate(ip string) (GeoInfo, error) {
	_ = "STUB: not implemented"
	return *new(GeoInfo), nil
}

func (f *maxmindDBReader) Close() error { _ = "STUB: not implemented"; return nil }
