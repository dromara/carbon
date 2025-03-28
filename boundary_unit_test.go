package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_StartOfCentury(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfCentury().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").StartOfCentury().ToString())
		assert.Empty(t, Parse("0").StartOfCentury().ToString())
		assert.Empty(t, Parse("xxx").StartOfCentury().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2000-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfCentury().ToString())
		assert.Equal(t, "2000-01-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfCentury().ToString())
		assert.Equal(t, "2000-01-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfCentury().ToString())
	})
}

func TestCarbon_EndOfCentury(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0099-12-31 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfCentury().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").EndOfCentury().ToString())
		assert.Empty(t, Parse("0").EndOfCentury().ToString())
		assert.Empty(t, Parse("xxx").EndOfCentury().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2099-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfCentury().ToString())
		assert.Equal(t, "2099-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfCentury().ToString())
		assert.Equal(t, "2099-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfCentury().ToString())
	})
}

func TestCarbon_StartOfDecade(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfDecade().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").StartOfDecade().ToString())
		assert.Empty(t, Parse("0").StartOfDecade().ToString())
		assert.Empty(t, Parse("xxx").StartOfDecade().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfDecade().ToString())
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfDecade().ToString())
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfDecade().ToString())
	})
}

func TestCarbon_EndOfDecade(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0009-12-31 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfDecade().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").EndOfDecade().ToString())
		assert.Empty(t, Parse("xxx").EndOfDecade().ToString())
		assert.Empty(t, Parse("0").EndOfDecade().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2029-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfDecade().ToString())
		assert.Equal(t, "2029-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfDecade().ToString())
		assert.Equal(t, "2029-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfDecade().ToString())
	})
}

func TestCarbon_StartOfYear(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfYear().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").StartOfYear().ToString())
		assert.Empty(t, Parse("0").StartOfYear().ToString())
		assert.Empty(t, Parse("xxx").StartOfYear().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfYear().ToString())
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfYear().ToString())
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfYear().ToString())
	})
}

func TestCarbon_EndOfYear(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-12-31 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfYear().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").EndOfYear().ToString())
		assert.Empty(t, Parse("0").EndOfYear().ToString())
		assert.Empty(t, Parse("xxx").EndOfYear().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfYear().ToString())
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfYear().ToString())
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfYear().ToString())
	})
}

func TestCarbon_StartOfQuarter(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfQuarter().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").StartOfQuarter().ToString())
		assert.Empty(t, Parse("0").StartOfQuarter().ToString())
		assert.Empty(t, Parse("xxx").StartOfQuarter().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfQuarter().ToString())
		assert.Equal(t, "2020-07-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfQuarter().ToString())
		assert.Equal(t, "2020-10-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfQuarter().ToString())
	})
}

func TestCarbon_EndOfQuarter(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-03-31 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfQuarter().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").EndOfQuarter().ToString())
		assert.Empty(t, Parse("0").EndOfQuarter().ToString())
		assert.Empty(t, Parse("xxx").EndOfQuarter().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-03-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfQuarter().ToString())
		assert.Equal(t, "2020-09-30 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfQuarter().ToString())
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfQuarter().ToString())
	})
}

func TestCarbon_StartOfMonth(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfMonth().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").StartOfMonth().ToString())
		assert.Empty(t, Parse("0").StartOfMonth().ToString())
		assert.Empty(t, Parse("xxx").StartOfMonth().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfMonth().ToString())
		assert.Equal(t, "2020-08-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfMonth().ToString())
		assert.Equal(t, "2020-12-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfMonth().ToString())
	})
}

func TestCarbon_EndOfMonth(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-31 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfMonth().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").EndOfMonth().ToString())
		assert.Empty(t, Parse("0").EndOfMonth().ToString())
		assert.Empty(t, Parse("xxx").EndOfMonth().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfMonth().ToString())
		assert.Equal(t, "2020-08-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfMonth().ToString())
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfMonth().ToString())
	})
}

func TestCarbon_StartOfWeek(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-31 00:00:00 +0000 UTC", NewCarbon().StartOfWeek().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").StartOfWeek().ToString())
		assert.Empty(t, Parse("0").StartOfWeek().ToString())
		assert.Empty(t, Parse("xxx").StartOfWeek().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2019-12-29 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfWeek().ToString())
		assert.Equal(t, "2020-08-09 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfWeek().ToString())
		assert.Equal(t, "2020-12-27 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfWeek().ToString())
	})
}

func TestCarbon_EndOfWeek(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-06 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfWeek().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").EndOfWeek().ToString())
		assert.Empty(t, Parse("0").EndOfWeek().ToString())
		assert.Empty(t, Parse("xxx").EndOfWeek().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-04 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfWeek().ToString())
		assert.Equal(t, "2020-08-15 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfWeek().ToString())
		assert.Equal(t, "2021-01-02 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfWeek().ToString())
	})
}

func TestCarbon_StartOfDay(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfDay().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").StartOfDay().ToString())
		assert.Empty(t, Parse("0").StartOfDay().ToString())
		assert.Empty(t, Parse("xxx").StartOfDay().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfDay().ToString())
		assert.Equal(t, "2020-08-15 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfDay().ToString())
		assert.Equal(t, "2020-12-31 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfDay().ToString())
	})
}

func TestCarbon_EndOfDay(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfDay().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").EndOfDay().ToString())
		assert.Empty(t, Parse("0").EndOfDay().ToString())
		assert.Empty(t, Parse("xxx").EndOfDay().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfDay().ToString())
		assert.Equal(t, "2020-08-15 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfDay().ToString())
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfDay().ToString())
	})
}

func TestCarbon_StartOfHour(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfHour().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").StartOfHour().ToString())
		assert.Empty(t, Parse("0").StartOfHour().ToString())
		assert.Empty(t, Parse("xxx").StartOfHour().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfHour().ToString())
		assert.Equal(t, "2020-08-15 12:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfHour().ToString())
		assert.Equal(t, "2020-12-31 23:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfHour().ToString())
	})
}

func TestCarbon_EndOfHour(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:59:59.999999999 +0000 UTC", NewCarbon().EndOfHour().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").EndOfHour().ToString())
		assert.Empty(t, Parse("0").EndOfHour().ToString())
		assert.Empty(t, Parse("xxx").EndOfHour().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfHour().ToString())
		assert.Equal(t, "2020-08-15 12:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfHour().ToString())
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfHour().ToString())
	})
}

func TestCarbon_StartOfMinute(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfMinute().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").StartOfMinute().ToString())
		assert.Empty(t, Parse("0").StartOfMinute().ToString())
		assert.Empty(t, Parse("xxx").StartOfMinute().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfMinute().ToString())
		assert.Equal(t, "2020-08-15 12:30:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfMinute().ToString())
		assert.Equal(t, "2020-12-31 23:59:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfMinute().ToString())
	})
}

func TestCarbon_EndOfMinute(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:59.999999999 +0000 UTC", NewCarbon().EndOfMinute().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").EndOfMinute().ToString())
		assert.Empty(t, Parse("0").EndOfMinute().ToString())
		assert.Empty(t, Parse("xxx").EndOfMinute().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfMinute().ToString())
		assert.Equal(t, "2020-08-15 12:30:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfMinute().ToString())
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfMinute().ToString())
	})
}

func TestCarbon_StartOfSecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfSecond().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").StartOfSecond().ToString())
		assert.Empty(t, Parse("0").StartOfSecond().ToString())
		assert.Empty(t, Parse("xxx").StartOfSecond().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfSecond().ToString())
		assert.Equal(t, "2020-08-15 12:30:30 +0000 UTC", Parse("2020-08-15 12:30:30.66666").StartOfSecond().ToString())
		assert.Equal(t, "2020-12-31 23:59:59 +0000 UTC", Parse("2020-12-31 23:59:59.999999999").StartOfSecond().ToString())
	})
}

func TestCarbon_EndOfSecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00.999999999 +0000 UTC", NewCarbon().EndOfSecond().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").EndOfSecond().ToString())
		assert.Empty(t, Parse("0").EndOfSecond().ToString())
		assert.Empty(t, Parse("xxx").EndOfSecond().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfSecond().ToString())
		assert.Equal(t, "2020-08-15 12:30:30.999999999 +0000 UTC", Parse("2020-08-15 12:30:30.66666").EndOfSecond().ToString())
		assert.Equal(t, "2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59.999999999").EndOfSecond().ToString())
	})
}
