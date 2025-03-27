package carbon

import "sync"

type TestNow struct {
	isFrozen  bool
	frozenNow Carbon
	rw        *sync.RWMutex
}

var testNow = &TestNow{
	rw: new(sync.RWMutex),
}

// SetTestNow sets a test Carbon instance (real or mock) to be returned when a "now" instance is created.
// 设置当前测试时间
func SetTestNow(carbon Carbon) {
	testNow.rw.Lock()
	defer testNow.rw.Unlock()

	testNow.isFrozen = true
	testNow.frozenNow = carbon
}

// CleanTestNow clears a test Carbon instance (real or mock) to be returned when a "now" instance is created.
// 清除当前测试时间
func CleanTestNow() {
	testNow.rw.Lock()
	defer testNow.rw.Unlock()

	testNow.isFrozen = false
}

// Deprecated: it will be removed in the future, use CleanTestNow instead.
// UnSetTestNow clears a test Carbon instance (real or mock) to be returned when a "now" instance is created.
// 清除当前测试时间
func UnSetTestNow() {
	CleanTestNow()
}

// Deprecated: it will be removed in the future, use IsTestNow instead.
// IsSetTestNow report whether there is testing time now.
// 是否设置过当前测试时间
func IsSetTestNow() bool {
	testNow.rw.Lock()
	defer testNow.rw.Unlock()

	return testNow.isFrozen
}

// IsTestNow report whether there is testing time now.
// 是否设置过当前测试时间
func IsTestNow() bool {
	testNow.rw.Lock()
	defer testNow.rw.Unlock()

	return testNow.isFrozen
}
