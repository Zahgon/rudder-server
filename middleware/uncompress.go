package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
)

// UncompressMiddleware uncompresses gzipped HTTP requests carrying a 'Content-Encoding: gzip' header.
var UncompressMiddleware = func(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Encoding") == "gzip" {
			r.Body = &gzipReader{body: r.Body}
		}
		h.ServeHTTP(w, r)
	})
}

// gzipReader wraps a body so it can lazily
// call gzip.NewReader on the first call to Read
type gzipReader struct {
	body io.ReadCloser // underlying request body
	zr   *gzip.Reader  // lazily-initialized gzip reader
	zerr error         // any error from gzip.NewReader; sticky
}

func (gz *gzipReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (gz *gzipReader) Close() error { _ = "STUB: not implemented"; return nil }
