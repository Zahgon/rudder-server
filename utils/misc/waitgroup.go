package misc

import (
	"sync"
)

type WaitGroup struct {
	wg       sync.WaitGroup
	errChan  chan error
	doneChan chan bool
}

func NewWaitGroup() *WaitGroup { _ = "STUB: not implemented"; return nil }

func (wg *WaitGroup) Add(delta int) { _ = "STUB: not implemented"; return }

func (wg *WaitGroup) Done() { _ = "STUB: not implemented"; return }

func (wg *WaitGroup) Err(err error) { _ = "STUB: not implemented"; return }

func (wg *WaitGroup) Wait() error { _ = "STUB: not implemented"; return nil }

func (wg *WaitGroup) WaitForAll() []error { _ = "STUB: not implemented"; return nil }
