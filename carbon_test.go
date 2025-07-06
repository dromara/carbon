package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCarbon(t *testing.T) {
	loc, _ := time.LoadLocation(PRC)

	t1, _ := time.Parse(DateTimeLayout, "2020-08-05 13:14:15")
	t2, _ := time.ParseInLocation(DateTimeLayout, "2020-08-05 13:14:15", loc)

	t.Run("zero time", func(t *testing.T) {
		c1 := NewCarbon()
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", c1.ToString())
		assert.Equal(t, time.Time{}.String(), c1.ToString())

		c2 := NewCarbon().SetLocation(loc)
		assert.Equal(t, "0001-01-01 08:05:43 +0805 LMT", c2.ToString())
		assert.Equal(t, time.Time{}.In(loc).String(), c2.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		c1 := NewCarbon(t1)
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", c1.ToString())
		assert.Equal(t, t1.String(), c1.ToString())

		c2 := NewCarbon(t2)
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", c2.ToString())
		assert.Equal(t, t2.String(), c2.ToString())
	})
}

func TestCarbon_Sleep(t *testing.T) {
	t.Run("sleep in normal mode", func(t *testing.T) {
		ClearTestNow()
		assert.False(t, IsTestNow())

		c := NewCarbon()
		start := time.Now()

		c.Sleep(1 * time.Millisecond)

		duration := time.Since(start)
		assert.GreaterOrEqual(t, duration, 1*time.Millisecond)
	})

	t.Run("sleep in test mode", func(t *testing.T) {
		testNow := Parse("2020-08-05 13:14:15")
		SetTestNow(testNow)
		defer ClearTestNow()

		assert.True(t, IsTestNow())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", Now().ToString())

		c := NewCarbon()
		c.Sleep(1 * time.Hour)

		assert.Equal(t, "2020-08-05 14:14:15 +0000 UTC", Now().ToString())
	})

	t.Run("sleep zero duration", func(t *testing.T) {
		testNow := Parse("2020-08-05 13:14:15")
		SetTestNow(testNow)
		defer ClearTestNow()

		assert.True(t, IsTestNow())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", Now().ToString())

		c := NewCarbon()
		c.Sleep(0)

		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", Now().ToString())
	})

	t.Run("sleep negative duration", func(t *testing.T) {
		testNow := Parse("2020-08-05 13:14:15")
		SetTestNow(testNow)
		defer ClearTestNow()

		assert.True(t, IsTestNow())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", Now().ToString())

		c := NewCarbon()
		c.Sleep(-1 * time.Hour)

		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", Now().ToString())
	})

	t.Run("sleep multiple times", func(t *testing.T) {
		testNow := Parse("2020-08-05 13:14:15")
		SetTestNow(testNow)
		defer ClearTestNow()

		assert.True(t, IsTestNow())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", Now().ToString())

		c := NewCarbon()
		c.Sleep(30 * time.Minute)
		assert.Equal(t, "2020-08-05 13:44:15 +0000 UTC", Now().ToString())

		c.Sleep(15 * time.Minute)
		assert.Equal(t, "2020-08-05 13:59:15 +0000 UTC", Now().ToString())

		c.Sleep(45 * time.Second)
		assert.Equal(t, "2020-08-05 14:00:00 +0000 UTC", Now().ToString())
	})
}
