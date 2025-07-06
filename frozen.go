package carbon

import (
	"sync"
	"sync/atomic"
)

// FrozenNow defines a FrozenNow struct.
type FrozenNow struct {
	isFrozen int32
	testNow  Carbon
	rw       sync.RWMutex
}

var frozenNow = FrozenNow{}

// SetTestNow sets a test Carbon instance for now.
func SetTestNow(c Carbon) {
	frozenNow.rw.Lock()
	frozenNow.testNow = c
	frozenNow.rw.Unlock()

	atomic.StoreInt32(&frozenNow.isFrozen, 1)
}

// ClearTestNow clears the test Carbon instance for now.
func ClearTestNow() {
	atomic.StoreInt32(&frozenNow.isFrozen, 0)
}

// IsTestNow reports whether is testing time.
func IsTestNow() bool {
	return atomic.LoadInt32(&frozenNow.isFrozen) == 1
}
