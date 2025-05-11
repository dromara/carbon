package carbon

import "sync"

// FrozenNow defines a FrozenNow struct.
// 定义 FrozenNow 结构体
type FrozenNow struct {
	isFrozen bool
	testNow  Carbon
	rw       *sync.RWMutex
}

var frozenNow = &FrozenNow{
	rw: new(sync.RWMutex),
}

// SetTestNow sets a test Carbon instance (real or mock) to be returned when a "now" instance is created.
// 设置当前测试时间
func SetTestNow(c Carbon) {
	frozenNow.rw.Lock()
	defer frozenNow.rw.Unlock()

	frozenNow.isFrozen = true
	frozenNow.testNow = c
}

// Deprecated: it will be removed in the future, use ClearTestNow instead.
// CleanTestNow clears the test Carbon instance for now.
// 清除当前测试时间
func CleanTestNow() {
	ClearTestNow()
}

// ClearTestNow clears the test Carbon instance for now.
// 清除当前测试时间
func ClearTestNow() {
	frozenNow.rw.Lock()
	defer frozenNow.rw.Unlock()

	frozenNow.isFrozen = false
}

// Deprecated: it will be removed in the future, use IsTestNow instead.
// IsSetTestNow report whether there is testing time now.
// 是否设置过当前测试时间
func IsSetTestNow() bool {
	return IsTestNow()
}

// IsTestNow report whether there is testing time now.
// 是否设置过当前测试时间
func IsTestNow() bool {
	frozenNow.rw.Lock()
	defer frozenNow.rw.Unlock()

	return frozenNow.isFrozen
}
