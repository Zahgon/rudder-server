/*
Http Interface Use this instead of http package.

usage example

import "github.com/rudderlabs/rudder-server/utils/sysUtils"

var	Http sysUtils.HttpI = &sysUtils.Http{}
			or
var	Http sysUtils.HttpI = sysUtils.NewHttp()

...

Http.NewRequest(...)
*/

//go:generate mockgen -destination=../../mocks/utils/sysUtils/mock_http.go -package mock_sysUtils github.com/rudderlabs/rudder-server/utils/sysUtils HttpI
package sysUtils

import (
	"context"
	"io"
	"net/http"
)

type HttpI interface {
	NewRequest(method, url string, body io.Reader) (*http.Request, error)
	NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*http.Request, error)
}

type Http struct{}

// NewHttp returns a Http instance
func NewHttp() *Http {
	_ = "STUB: not implemented"

	// NewRequest wraps NewRequestWithContext using the background context.
	return nil
}

func (*Http) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*Http) NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
