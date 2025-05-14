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

// Deprecated: CleanTestNow will be removed in the future, use ClearTestNow instead.
// CleanTestNow clears the test Carbon instance for now.
func CleanTestNow() {
	ClearTestNow()
}

// ClearTestNow clears the test Carbon instance for now.
func ClearTestNow() {
	frozenNow.rw.Lock()
	defer frozenNow.rw.Unlock()

	frozenNow.isFrozen = false
}

// Deprecated: IsSetTestNow will be removed in the future, use IsTestNow instead.
// IsSetTestNow report whether there is testing time now.
func IsSetTestNow() bool {
	return IsTestNow()
}

// IsTestNow report whether there is testing time now.
func IsTestNow() bool {
	frozenNow.rw.Lock()
	defer frozenNow.rw.Unlock()

	return frozenNow.isFrozen
}
