package carbon

// SetTestNow sets a test Carbon instance (real or mock) to be returned when a "now" instance is created.
// 设置当前测试时间
func (c *Carbon) SetTestNow(carbon *Carbon) {
	c.testNow, c.loc = carbon.TimestampNano(), carbon.loc
}

// UnSetTestNow clears a test Carbon instance (real or mock) to be returned when a "now" instance is created.
// 清除当前测试时间
func (c *Carbon) UnSetTestNow() {
	c.testNow = 0
}

// IsSetTestNow report whether there is testing time now.
// 是否设置过当前测试时间
func (c *Carbon) IsSetTestNow() bool {
	return c.testNow > 0
}
