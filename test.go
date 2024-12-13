package carbon

import (
	"os"
	"testing"
)

// SetTestNow sets a test Carbon instance (real or mock) to be returned when a "now" instance is created.
// 设置当前测试时间
func (c *Carbon) SetTestNow(carbon Carbon) {
	c.testNow, c.loc = carbon.TimestampNano(), carbon.loc
}

// UnSetTestNow clears a test Carbon instance (real or mock) to be returned when a "now" instance is created.
// 清除当前测试时间
func (c *Carbon) UnSetTestNow() {
	c.testNow = 0
}

// IsSetTestNow report whether there is testing time now.
// 是否设置过当前测试时间
func (c Carbon) IsSetTestNow() bool {
	return c.testNow > 0
}

// TestMain sets up the testing environment for all tests
// https://pkg.go.dev/testing#hdr-Main
func TestMain(m *testing.M) {
	// The whole tests were written for PRC timezone (China).
	// The codebase of test is too large to be changed.
	// Without this hack the tests will fail if you use a different timezone than PRC
	// This will affect the way Go compute the timezone when using time.Local
	_ = os.Setenv("TZ", "PRC")

	m.Run()
}

func prepareTest(tb testing.TB) {
	tb.Helper()

	// Store the current default
	savedDefault := Default{
		Layout:       defaultLayout,
		Timezone:     defaultTimezone,
		Locale:       defaultLocale,
		WeekStartsAt: defaultWeekStartsAt,
	}
	tb.Cleanup(func() {
		// restore the default when test is done
		SetDefault(savedDefault)
	})
}
