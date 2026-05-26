package webhook

import (
	"net/http"
	"net/http/httptest"
	"sync"
)

type Recorder struct {
	Server     *httptest.Server
	requests   [][]byte
	requestsMu sync.RWMutex
}

func NewRecorder() *Recorder { _ = "STUB: not implemented"; return nil }

func (whr *Recorder) RequestsCount() int { _ = "STUB: not implemented"; return 0 }

func (whr *Recorder) Requests() []*http.Request { _ = "STUB: not implemented"; return nil }

func (whr *Recorder) Close() { _ = "STUB: not implemented"; return }
