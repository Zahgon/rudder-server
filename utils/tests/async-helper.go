package testutils

import (
	"sync"
	"time"

	"go.uber.org/mock/gomock"
)

// AsyncTestHelper provides synchronization methods to test goroutines.
// Example:
//
//	var (
//		asyncHelper testutils.AsyncTestHelper
//	)
//
//	BeforeEach(func() {
//		mockMyInterface.EXPECT().MyMethodInGoroutine().Do(asyncHelper.ExpectAndNotifyCallback())
//	})
//
//	AfterEach(func() {
//		asyncHelper.WaitWithTimeout(time.Second)
//	})
type AsyncTestHelper struct {
	wg             sync.WaitGroup
	waitingMap     map[string]int
	waitingMapLock sync.RWMutex
}

// ExpectAndNotifyCallback Adds one to this helper's WaitGroup, and provides a callback that calls Done on it.
// Should be used for gomock Do calls that trigger via mocked functions executed in a goroutine.
func (helper *AsyncTestHelper) ExpectAndNotifyCallback() func(...any) {
	_ = "STUB: not implemented"
	return nil
}

// ExpectAndNotifyCallback Adds one to this helper's WaitGroup, and provides a callback that calls Done on it.
// Should be used for gomock Do calls that trigger via mocked functions executed in a goroutine.
func (helper *AsyncTestHelper) ExpectAndNotifyCallbackWithName(name string) func(...any) {
	_ = "STUB: not implemented"
	return nil
}

// ExpectAndNotifyCallback Adds one to this helper's WaitGroup, and provides a callback that calls Done on it.
// Should be used for gomock Do calls that trigger via mocked functions executed in a goroutine.
func (helper *AsyncTestHelper) ExpectAndNotifyCallbackWithNameOnce(name string) func() {
	_ = "STUB: not implemented"
	return nil
}

// WaitWithTimeout waits for this helper's WaitGroup until provided timeout.
// Should wait for all ExpectAndNotifyCallback callbacks, registered in asynchronous mocks calls
func (helper *AsyncTestHelper) WaitWithTimeout(d time.Duration) { _ = "STUB: not implemented"; return }

// RegisterCalls registers a number of calls to this async helper, so that they are waited for.
func (helper *AsyncTestHelper) RegisterCalls(calls ...*gomock.Call) {
	_ = "STUB: not implemented"
	return
}

// RunTestWithTimeout runs function f until provided timeout.
// If the function times out, it will cause the ginkgo test to fail.
func (helper *AsyncTestHelper) RunTestWithTimeout(f func(), d time.Duration) {
	_ = "STUB: not implemented"
	return
}

// RunTestWithTimeout runs function f until provided timeout.
// If the function times out, it will cause the ginkgo test to fail.
func RunTestWithTimeout(f func(), d time.Duration) { _ = "STUB: not implemented"; return }

func (helper *AsyncTestHelper) Setup() { _ = "STUB: not implemented"; return }
