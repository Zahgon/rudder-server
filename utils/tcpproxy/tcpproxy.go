package tcpproxy

import (
	"io"
	"sync"
	"sync/atomic"
	"testing"
)

type Proxy struct {
	LocalAddr     string
	RemoteAddr    string
	BytesSent     atomic.Int64
	BytesReceived atomic.Int64
	Verbose       bool

	wg   sync.WaitGroup
	stop chan struct{}
}

func (p *Proxy) Start(t testing.TB) { _ = "STUB: not implemented"; return }

// error accepting connection

// cannot dial remote, return and listen for new connections

// one of the connections got terminated
// TCP proxy stopped

func (p *Proxy) Stop() { _ = "STUB: not implemented"; return }

func (p *Proxy) pipe(src io.Reader, dst io.Writer, bytesMetric *atomic.Int64, done chan struct{}) {
	_ = "STUB: not implemented"
	return
}

// this is a blocking call, it terminates when the connection is closed

// connection is closed, send signal to stop proxy
