package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_StartOfCentury(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().StartOfCentury()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0000-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").StartOfCentury()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").StartOfCentury()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2000-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfCentury().ToString())
		assert.Equal(t, "2000-01-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfCentury().ToString())
		assert.Equal(t, "2000-01-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfCentury().ToString())
	})
}

func TestCarbon_EndOfCentury(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().EndOfCentury()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0099-12-31 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").EndOfCentury()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").EndOfCentury()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2099-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfCentury().ToString())
		assert.Equal(t, "2099-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfCentury().ToString())
		assert.Equal(t, "2099-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfCentury().ToString())
	})
}

func TestCarbon_StartOfDecade(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().StartOfDecade()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0000-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").StartOfDecade()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").StartOfDecade()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfDecade().ToString())
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfDecade().ToString())
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfDecade().ToString())
	})
}

func TestCarbon_EndOfDecade(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().EndOfDecade()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0009-12-31 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").EndOfDecade()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").EndOfDecade()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2029-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfDecade().ToString())
		assert.Equal(t, "2029-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfDecade().ToString())
		assert.Equal(t, "2029-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfDecade().ToString())
	})
}

func TestCarbon_StartOfYear(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().StartOfYear()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").StartOfYear()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").StartOfYear()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfYear().ToString())
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfYear().ToString())
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfYear().ToString())
	})
}

func TestCarbon_EndOfYear(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().EndOfYear()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0001-12-31 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").EndOfYear()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").EndOfYear()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfYear().ToString())
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfYear().ToString())
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfYear().ToString())
	})
}

func TestCarbon_StartOfQuarter(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().StartOfQuarter()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").StartOfQuarter()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").StartOfQuarter()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfQuarter().ToString())
		assert.Equal(t, "2020-07-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfQuarter().ToString())
		assert.Equal(t, "2020-10-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfQuarter().ToString())
	})
}

func TestCarbon_EndOfQuarter(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().EndOfQuarter()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0001-03-31 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").EndOfQuarter()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").EndOfQuarter()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-03-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfQuarter().ToString())
		assert.Equal(t, "2020-09-30 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfQuarter().ToString())
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfQuarter().ToString())
	})
}

func TestCarbon_StartOfMonth(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().StartOfMonth()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").StartOfMonth()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").StartOfMonth()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfMonth().ToString())
		assert.Equal(t, "2020-08-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfMonth().ToString())
		assert.Equal(t, "2020-12-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfMonth().ToString())
	})
}

func TestCarbon_EndOfMonth(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().EndOfMonth()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0001-01-31 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").EndOfMonth()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").EndOfMonth()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfMonth().ToString())
		assert.Equal(t, "2020-08-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfMonth().ToString())
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfMonth().ToString())
	})
}

func TestCarbon_StartOfWeek(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().StartOfWeek()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").StartOfWeek()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").StartOfWeek()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2019-12-30 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfWeek().ToString())
		assert.Equal(t, "2020-08-10 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfWeek().ToString())
		assert.Equal(t, "2020-12-28 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfWeek().ToString())
		assert.Equal(t, "2025-04-07 00:00:00 +0000 UTC", Parse("2025-04-07 00:00:00").StartOfWeek().ToString())
	})
}

func TestCarbon_EndOfWeek(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().EndOfWeek()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0001-01-07 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").EndOfWeek()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").EndOfWeek()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-05 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfWeek().ToString())
		assert.Equal(t, "2020-08-16 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfWeek().ToString())
		assert.Equal(t, "2021-01-03 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfWeek().ToString())
		assert.Equal(t, "2025-04-13 23:59:59.999999999 +0000 UTC", Parse("2025-04-13 00:00:00").EndOfWeek().ToString())
	})
}

func TestCarbon_StartOfDay(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().StartOfDay()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").StartOfDay()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").StartOfDay()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfDay().ToString())
		assert.Equal(t, "2020-08-15 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfDay().ToString())
		assert.Equal(t, "2020-12-31 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfDay().ToString())
	})
}

func TestCarbon_EndOfDay(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().EndOfDay()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0001-01-01 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").EndOfDay()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").EndOfDay()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfDay().ToString())
		assert.Equal(t, "2020-08-15 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfDay().ToString())
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfDay().ToString())
	})
}

func TestCarbon_StartOfHour(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().StartOfHour()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").StartOfHour()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").StartOfHour()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfHour().ToString())
		assert.Equal(t, "2020-08-15 12:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfHour().ToString())
		assert.Equal(t, "2020-12-31 23:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfHour().ToString())
	})
}

func TestCarbon_EndOfHour(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().EndOfHour()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0001-01-01 00:59:59.999999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").EndOfHour()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").EndOfHour()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfHour().ToString())
		assert.Equal(t, "2020-08-15 12:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfHour().ToString())
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfHour().ToString())
	})
}

func TestCarbon_StartOfMinute(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().StartOfMinute()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").StartOfMinute()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").StartOfMinute()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfMinute().ToString())
		assert.Equal(t, "2020-08-15 12:30:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfMinute().ToString())
		assert.Equal(t, "2020-12-31 23:59:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfMinute().ToString())
	})
}

func TestCarbon_EndOfMinute(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().EndOfMinute()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0001-01-01 00:00:59.999999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").EndOfMinute()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").EndOfMinute()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfMinute().ToString())
		assert.Equal(t, "2020-08-15 12:30:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfMinute().ToString())
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfMinute().ToString())
	})
}

func TestCarbon_StartOfSecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().StartOfSecond()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").StartOfSecond()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").StartOfSecond()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfSecond().ToString())
		assert.Equal(t, "2020-08-15 12:30:30 +0000 UTC", Parse("2020-08-15 12:30:30.66666").StartOfSecond().ToString())
		assert.Equal(t, "2020-12-31 23:59:59 +0000 UTC", Parse("2020-12-31 23:59:59.999999999").StartOfSecond().ToString())
	})
}

func TestCarbon_EndOfSecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().EndOfSecond()
		assert.Nil(t, c.Error)
		assert.Equal(t, "0001-01-01 00:00:00.999999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").EndOfSecond()
		assert.Nil(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("error time", func(t *testing.T) {
		c := Parse("xxx").EndOfSecond()
		assert.Error(t, c.Error)
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfSecond().ToString())
		assert.Equal(t, "2020-08-15 12:30:30.999999999 +0000 UTC", Parse("2020-08-15 12:30:30.66666").EndOfSecond().ToString())
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59.999999999").EndOfSecond().ToString())
	})
}
