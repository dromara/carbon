package carbon

import "sync"

// FrozenNow defines a FrozenNow struct.
type FrozenNow struct {
	isFrozen bool
	testNow  Carbon
	rw       *sync.RWMutex
}

var frozenNow = &FrozenNow{
	rw: new(sync.RWMutex),
}

// SetTestNow sets a test Carbon instance for now.
func SetTestNow(c Carbon) {
	frozenNow.rw.Lock()
	defer frozenNow.rw.Unlock()

	frozenNow.isFrozen = true
	frozenNow.testNow = c
}

// CleanTestNow clears the test Carbon instance for now.
//
// Deprecated: it will be removed in the future, use ClearTestNow instead.
func CleanTestNow() {
	ClearTestNow()
}

// ClearTestNow clears the test Carbon instance for now.
func ClearTestNow() {
	frozenNow.rw.Lock()
	defer frozenNow.rw.Unlock()

	frozenNow.isFrozen = false
}

// IsSetTestNow report whether there is testing time now.
//
// Deprecated: it will be removed in the future, use IsTestNow instead.
func IsSetTestNow() bool {
	return IsTestNow()
}

// IsTestNow report whether there is testing time now.
func IsTestNow() bool {
	frozenNow.rw.Lock()
	defer frozenNow.rw.Unlock()

	return frozenNow.isFrozen
}
