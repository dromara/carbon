package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_StdTime(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, time.Time{}, NewCarbon().StdTime())
	})

	t.Run("nil location", func(t *testing.T) {
		c := Parse("2020-08-05").SetTimezone(PRC)
		assert.Equal(t, "2020-08-05 08:00:00 +0800 CST", c.StdTime().String())
		c.loc = nil
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", c.StdTime().String())
	})

	t.Run("without timezone", func(t *testing.T) {
		c := Parse("2020-08-05")
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", c.StdTime().String())
	})

	t.Run("with timezone", func(t *testing.T) {
		c := Parse("2020-08-05", PRC)
		assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", c.StdTime().String())
	})
}

func TestCarbon_DaysInYear(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 365, NewCarbon().DaysInYear())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DaysInYear())
		assert.Zero(t, Parse("0").DaysInYear())
		assert.Zero(t, Parse("xxx").DaysInYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 366, Parse("2020-08-05").DaysInYear())
		assert.Equal(t, 365, Parse("2021-08-05").DaysInYear())
	})
}

func TestCarbon_DaysInMonth(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 31, NewCarbon().DaysInMonth())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DaysInMonth())
		assert.Zero(t, Parse("0").DaysInMonth())
		assert.Zero(t, Parse("xxx").DaysInMonth())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 31, Parse("2020-01-05").DaysInMonth())
		assert.Equal(t, 29, Parse("2020-02-05").DaysInMonth())
		assert.Equal(t, 31, Parse("2020-03-05").DaysInMonth())
		assert.Equal(t, 30, Parse("2020-04-05").DaysInMonth())
	})
}

func TestCarbon_MonthOfYear(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 1, NewCarbon().MonthOfYear())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").MonthOfYear())
		assert.Zero(t, Parse("0").MonthOfYear())
		assert.Zero(t, Parse("xxx").MonthOfYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 1, Parse("2020-01-05").MonthOfYear())
		assert.Equal(t, 2, Parse("2020-02-05").MonthOfYear())
		assert.Equal(t, 3, Parse("2020-03-05").MonthOfYear())
		assert.Equal(t, 4, Parse("2020-04-05").MonthOfYear())
	})
}

func TestCarbon_DayOfYear(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 1, NewCarbon().DayOfYear())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DayOfYear())
		assert.Zero(t, Parse("0").DayOfYear())
		assert.Zero(t, Parse("xxx").DayOfYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 5, Parse("2020-01-05").DayOfYear())
		assert.Equal(t, 36, Parse("2020-02-05").DayOfYear())
		assert.Equal(t, 65, Parse("2020-03-05").DayOfYear())
		assert.Equal(t, 96, Parse("2020-04-05").DayOfYear())
	})
}

func TestCarbon_DayOfMonth(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 1, NewCarbon().DayOfMonth())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DayOfMonth())
		assert.Zero(t, Parse("0").DayOfMonth())
		assert.Zero(t, Parse("xxx").DayOfMonth())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 1, Parse("2020-01-01").DayOfMonth())
		assert.Equal(t, 5, Parse("2020-01-05").DayOfMonth())
		assert.Equal(t, 31, Parse("2020-01-31").DayOfMonth())
	})
}

func TestCarbon_DayOfWeek(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 1, NewCarbon().DayOfWeek())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").DayOfWeek())
		assert.Zero(t, Parse("0").DayOfWeek())
		assert.Zero(t, Parse("xxx").DayOfWeek())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 1, Parse("2020-08-03").DayOfWeek())
		assert.Equal(t, 2, Parse("2020-08-04").DayOfWeek())
		assert.Equal(t, 3, Parse("2020-08-05").DayOfWeek())
	})
}

func TestCarbon_WeekOfYear(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 1, NewCarbon().WeekOfYear())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").WeekOfYear())
		assert.Zero(t, Parse("0").WeekOfYear())
		assert.Zero(t, Parse("xxx").WeekOfYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 53, Parse("2021-01-01").WeekOfYear())
		assert.Equal(t, 5, Parse("2021-02-01").WeekOfYear())
		assert.Equal(t, 9, Parse("2021-03-01").WeekOfYear())
		assert.Equal(t, 13, Parse("2021-04-01").WeekOfYear())
	})
}

func TestCarbon_WeekOfMonth(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 1, NewCarbon().WeekOfMonth())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").WeekOfMonth())
		assert.Zero(t, Parse("0").WeekOfMonth())
		assert.Zero(t, Parse("xxx").WeekOfMonth())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 1, Parse("2021-07-01").WeekOfMonth())
		assert.Equal(t, 1, Parse("2021-07-02").WeekOfMonth())
		assert.Equal(t, 1, Parse("2021-07-03").WeekOfMonth())
		assert.Equal(t, 1, Parse("2021-07-04").WeekOfMonth())
		assert.Equal(t, 2, Parse("2021-07-05").WeekOfMonth())
		assert.Equal(t, 2, Parse("2021-07-06").WeekOfMonth())
	})
}

func TestCarbon_DateTime(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		year, month, day, hour, minute, second := NewCarbon().DateTime()
		assert.Equal(t, 1, year)
		assert.Equal(t, 1, month)
		assert.Equal(t, 1, day)
		assert.Zero(t, hour)
		assert.Zero(t, minute)
		assert.Zero(t, second)
	})

	t.Run("invalid time", func(t *testing.T) {
		c1 := Parse("")
		year1, month1, day1, hour1, minute1, second1 := c1.DateTime()
		assert.Zero(t, year1)
		assert.Zero(t, month1)
		assert.Zero(t, day1)
		assert.Zero(t, hour1)
		assert.Zero(t, minute1)
		assert.Zero(t, second1)

		c2 := Parse("xxx")
		year2, month2, day2, hour2, minute2, second2 := c2.DateTime()
		assert.Zero(t, year2)
		assert.Zero(t, month2)
		assert.Zero(t, day2)
		assert.Zero(t, hour2)
		assert.Zero(t, minute2)
		assert.Zero(t, second2)
	})

	t.Run("valid time", func(t *testing.T) {
		year, month, day, hour, minute, second := Parse("2020-08-05 13:14:15.999999999").DateTime()
		assert.Equal(t, 2020, year)
		assert.Equal(t, 8, month)
		assert.Equal(t, 5, day)
		assert.Equal(t, 13, hour)
		assert.Equal(t, 14, minute)
		assert.Equal(t, 15, second)
	})
}

func TestCarbon_DateTimeMilli(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		year, month, day, hour, minute, second, millisecond := NewCarbon().DateTimeMilli()
		assert.Equal(t, 1, year)
		assert.Equal(t, 1, month)
		assert.Equal(t, 1, day)
		assert.Zero(t, hour)
		assert.Zero(t, minute)
		assert.Zero(t, second)
		assert.Zero(t, second)
		assert.Zero(t, millisecond)
	})

	t.Run("invalid time", func(t *testing.T) {
		c1 := Parse("")
		year1, month1, day1, hour1, minute1, second1, millisecond1 := c1.DateTimeMilli()
		assert.Zero(t, year1)
		assert.Zero(t, month1)
		assert.Zero(t, day1)
		assert.Zero(t, hour1)
		assert.Zero(t, minute1)
		assert.Zero(t, second1)
		assert.Zero(t, millisecond1)

		c2 := Parse("xxx")
		year2, month2, day2, hour2, minute2, second2, millisecond2 := c2.DateTimeMilli()
		assert.Zero(t, year2)
		assert.Zero(t, month2)
		assert.Zero(t, day2)
		assert.Zero(t, hour2)
		assert.Zero(t, minute2)
		assert.Zero(t, second2)
		assert.Zero(t, millisecond2)
	})

	t.Run("valid time", func(t *testing.T) {
		year, month, day, hour, minute, second, millisecond := Parse("2020-08-05 13:14:15.999999999").DateTimeMilli()
		assert.Equal(t, 2020, year)
		assert.Equal(t, 8, month)
		assert.Equal(t, 5, day)
		assert.Equal(t, 13, hour)
		assert.Equal(t, 14, minute)
		assert.Equal(t, 15, second)
		assert.Equal(t, 999, millisecond)
	})
}

func TestCarbon_DateTimeMicro(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		year, month, day, hour, minute, second, microsecond := NewCarbon().DateTimeMicro()
		assert.Equal(t, 1, year)
		assert.Equal(t, 1, month)
		assert.Equal(t, 1, day)
		assert.Zero(t, hour)
		assert.Zero(t, minute)
		assert.Zero(t, second)
		assert.Zero(t, second)
		assert.Zero(t, microsecond)
	})

	t.Run("invalid time", func(t *testing.T) {
		c1 := Parse("")
		year1, month1, day1, hour1, minute1, second1, microsecond1 := c1.DateTimeMicro()
		assert.Zero(t, year1)
		assert.Zero(t, month1)
		assert.Zero(t, day1)
		assert.Zero(t, hour1)
		assert.Zero(t, minute1)
		assert.Zero(t, second1)
		assert.Zero(t, microsecond1)

		c2 := Parse("xxx")
		year2, month2, day2, hour2, minute2, second2, microsecond2 := c2.DateTimeMicro()
		assert.Zero(t, year2)
		assert.Zero(t, month2)
		assert.Zero(t, day2)
		assert.Zero(t, hour2)
		assert.Zero(t, minute2)
		assert.Zero(t, second2)
		assert.Zero(t, microsecond2)
	})

	t.Run("valid time", func(t *testing.T) {
		year, month, day, hour, minute, second, microsecond := Parse("2020-08-05 13:14:15.999999999").DateTimeMicro()
		assert.Equal(t, 2020, year)
		assert.Equal(t, 8, month)
		assert.Equal(t, 5, day)
		assert.Equal(t, 13, hour)
		assert.Equal(t, 14, minute)
		assert.Equal(t, 15, second)
		assert.Equal(t, 999999, microsecond)
	})
}

func TestCarbon_DateTimeNano(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		year, month, day, hour, minute, second, nanosecond := NewCarbon().DateTimeNano()
		assert.Equal(t, 1, year)
		assert.Equal(t, 1, month)
		assert.Equal(t, 1, day)
		assert.Zero(t, hour)
		assert.Zero(t, minute)
		assert.Zero(t, second)
		assert.Zero(t, second)
		assert.Zero(t, nanosecond)
	})

	t.Run("invalid time", func(t *testing.T) {
		c1 := Parse("")
		year1, month1, day1, hour1, minute1, second1, nanosecond1 := c1.DateTimeNano()
		assert.Zero(t, year1)
		assert.Zero(t, month1)
		assert.Zero(t, day1)
		assert.Zero(t, hour1)
		assert.Zero(t, minute1)
		assert.Zero(t, second1)
		assert.Zero(t, nanosecond1)

		c2 := Parse("xxx")
		year2, month2, day2, hour2, minute2, second2, nanosecond2 := c2.DateTimeNano()
		assert.Zero(t, year2)
		assert.Zero(t, month2)
		assert.Zero(t, day2)
		assert.Zero(t, hour2)
		assert.Zero(t, minute2)
		assert.Zero(t, second2)
		assert.Zero(t, nanosecond2)
	})

	t.Run("valid time", func(t *testing.T) {
		year, month, day, hour, minute, second, nanosecond := Parse("2020-08-05 13:14:15.999999999").DateTimeNano()
		assert.Equal(t, 2020, year)
		assert.Equal(t, 8, month)
		assert.Equal(t, 5, day)
		assert.Equal(t, 13, hour)
		assert.Equal(t, 14, minute)
		assert.Equal(t, 15, second)
		assert.Equal(t, 999999999, nanosecond)
	})
}

func TestCarbon_Date(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		year, month, day := NewCarbon().Date()
		assert.Equal(t, 1, year)
		assert.Equal(t, 1, month)
		assert.Equal(t, 1, day)
	})

	t.Run("invalid time", func(t *testing.T) {
		c1 := Parse("")
		year1, month1, day1 := c1.Date()
		assert.Zero(t, year1)
		assert.Zero(t, month1)
		assert.Zero(t, day1)

		c2 := Parse("xxx")
		year2, month2, day2 := c2.Date()
		assert.Zero(t, year2)
		assert.Zero(t, month2)
		assert.Zero(t, day2)
	})

	t.Run("valid time", func(t *testing.T) {
		year, month, day := Parse("2020-08-05 13:14:15.999999999").Date()
		assert.Equal(t, 2020, year)
		assert.Equal(t, 8, month)
		assert.Equal(t, 5, day)
	})
}

func TestCarbon_DateMilli(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		year, month, day, millisecond := NewCarbon().DateMilli()
		assert.Equal(t, 1, year)
		assert.Equal(t, 1, month)
		assert.Equal(t, 1, day)
		assert.Zero(t, millisecond)
	})

	t.Run("invalid time", func(t *testing.T) {
		c1 := Parse("")
		year1, month1, day1, millisecond1 := c1.DateMilli()
		assert.Zero(t, year1)
		assert.Zero(t, month1)
		assert.Zero(t, day1)
		assert.Zero(t, millisecond1)

		c2 := Parse("xxx")
		year2, month2, day2, millisecond2 := c2.DateMilli()
		assert.Zero(t, year2)
		assert.Zero(t, month2)
		assert.Zero(t, day2)
		assert.Zero(t, millisecond2)
	})

	t.Run("valid time", func(t *testing.T) {
		year, month, day, millisecond := Parse("2020-08-05 13:14:15.999999999").DateMilli()
		assert.Equal(t, 2020, year)
		assert.Equal(t, 8, month)
		assert.Equal(t, 5, day)
		assert.Equal(t, 999, millisecond)
	})
}

func TestCarbon_DateMicro(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		year, month, day, microsecond := NewCarbon().DateMicro()
		assert.Equal(t, 1, year)
		assert.Equal(t, 1, month)
		assert.Equal(t, 1, day)
		assert.Zero(t, microsecond)
	})

	t.Run("invalid time", func(t *testing.T) {
		c1 := Parse("")
		year1, month1, day1, microsecond1 := c1.DateMicro()
		assert.Zero(t, year1)
		assert.Zero(t, month1)
		assert.Zero(t, day1)
		assert.Zero(t, microsecond1)

		c2 := Parse("xxx")
		year2, month2, day2, microsecond2 := c2.DateMicro()
		assert.Zero(t, year2)
		assert.Zero(t, month2)
		assert.Zero(t, day2)
		assert.Zero(t, microsecond2)
	})

	t.Run("valid time", func(t *testing.T) {
		year, month, day, microsecond := Parse("2020-08-05 13:14:15.999999999").DateMicro()
		assert.Equal(t, 2020, year)
		assert.Equal(t, 8, month)
		assert.Equal(t, 5, day)
		assert.Equal(t, 999999, microsecond)
	})
}

func TestCarbon_DateNano(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		year, month, day, nanosecond := NewCarbon().DateNano()
		assert.Equal(t, 1, year)
		assert.Equal(t, 1, month)
		assert.Equal(t, 1, day)
		assert.Zero(t, nanosecond)
	})

	t.Run("invalid time", func(t *testing.T) {
		c1 := Parse("")
		year1, month1, day1, nanosecond1 := c1.DateNano()
		assert.Zero(t, year1)
		assert.Zero(t, month1)
		assert.Zero(t, day1)
		assert.Zero(t, nanosecond1)

		c2 := Parse("xxx")
		year2, month2, day2, nanosecond2 := c2.DateNano()
		assert.Zero(t, year2)
		assert.Zero(t, month2)
		assert.Zero(t, day2)
		assert.Zero(t, nanosecond2)
	})

	t.Run("valid time", func(t *testing.T) {
		year, month, day, nanosecond := Parse("2020-08-05 13:14:15.999999999").DateNano()
		assert.Equal(t, 2020, year)
		assert.Equal(t, 8, month)
		assert.Equal(t, 5, day)
		assert.Equal(t, 999999999, nanosecond)
	})
}

func TestCarbon_Time(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		hour, minute, second := NewCarbon().Time()
		assert.Zero(t, hour)
		assert.Zero(t, minute)
		assert.Zero(t, second)
	})

	t.Run("invalid time", func(t *testing.T) {
		c1 := Parse("")
		hour1, minute1, second1 := c1.Time()
		assert.Zero(t, hour1)
		assert.Zero(t, minute1)
		assert.Zero(t, second1)

		c2 := Parse("xxx")
		hour2, minute2, second2 := c2.Time()
		assert.Zero(t, hour2)
		assert.Zero(t, minute2)
		assert.Zero(t, second2)
	})

	t.Run("valid time", func(t *testing.T) {
		hour, minute, second := Parse("2020-08-05 13:14:15.999999999").Time()
		assert.Equal(t, 13, hour)
		assert.Equal(t, 14, minute)
		assert.Equal(t, 15, second)
	})
}

func TestCarbon_TimeMilli(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		hour, minute, second, millisecond := NewCarbon().TimeMilli()
		assert.Zero(t, hour)
		assert.Zero(t, minute)
		assert.Zero(t, second)
		assert.Zero(t, millisecond)
	})

	t.Run("invalid time", func(t *testing.T) {
		c1 := Parse("")
		hour1, minute1, second1, millisecond1 := c1.TimeMilli()
		assert.Zero(t, hour1)
		assert.Zero(t, minute1)
		assert.Zero(t, second1)
		assert.Zero(t, millisecond1)

		c2 := Parse("xxx")
		hour2, minute2, second2, millisecond2 := c2.TimeMilli()
		assert.Zero(t, hour2)
		assert.Zero(t, minute2)
		assert.Zero(t, second2)
		assert.Zero(t, millisecond2)
	})

	t.Run("valid time", func(t *testing.T) {
		hour, minute, second, millisecond := Parse("2020-08-05 13:14:15.999999999").TimeMilli()
		assert.Equal(t, 13, hour)
		assert.Equal(t, 14, minute)
		assert.Equal(t, 15, second)
		assert.Equal(t, 999, millisecond)
	})
}

func TestCarbon_TimeMicro(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		hour, minute, second, microsecond := NewCarbon().TimeMicro()
		assert.Zero(t, hour)
		assert.Zero(t, minute)
		assert.Zero(t, second)
		assert.Zero(t, microsecond)
	})

	t.Run("invalid time", func(t *testing.T) {
		c1 := Parse("")
		hour1, minute1, second1, microsecond1 := c1.TimeMicro()
		assert.Zero(t, hour1)
		assert.Zero(t, minute1)
		assert.Zero(t, second1)
		assert.Zero(t, microsecond1)

		c2 := Parse("xxx")
		hour2, minute2, second2, microsecond2 := c2.TimeMicro()
		assert.Zero(t, hour2)
		assert.Zero(t, minute2)
		assert.Zero(t, second2)
		assert.Zero(t, microsecond2)
	})

	t.Run("valid time", func(t *testing.T) {
		hour, minute, second, microsecond := Parse("2020-08-05 13:14:15.999999999").TimeMicro()
		assert.Equal(t, 13, hour)
		assert.Equal(t, 14, minute)
		assert.Equal(t, 15, second)
		assert.Equal(t, 999999, microsecond)
	})
}

func TestCarbon_TimeNano(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		hour, minute, second, nanosecond := NewCarbon().TimeNano()
		assert.Zero(t, hour)
		assert.Zero(t, minute)
		assert.Zero(t, second)
		assert.Zero(t, nanosecond)
	})

	t.Run("invalid time", func(t *testing.T) {
		c1 := Parse("")
		hour1, minute1, second1, nanosecond1 := c1.TimeNano()
		assert.Zero(t, hour1)
		assert.Zero(t, minute1)
		assert.Zero(t, second1)
		assert.Zero(t, nanosecond1)

		c2 := Parse("xxx")
		hour2, minute2, second2, nanosecond2 := c2.TimeNano()
		assert.Zero(t, hour2)
		assert.Zero(t, minute2)
		assert.Zero(t, second2)
		assert.Zero(t, nanosecond2)
	})

	t.Run("valid time", func(t *testing.T) {
		hour, minute, second, nanosecond := Parse("2020-08-05 13:14:15.999999999").TimeNano()
		assert.Equal(t, 13, hour)
		assert.Equal(t, 14, minute)
		assert.Equal(t, 15, second)
		assert.Equal(t, 999999999, nanosecond)
	})
}

func TestCarbon_Century(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 1, NewCarbon().Century())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").Century())
		assert.Zero(t, Parse("0").Century())
		assert.Zero(t, Parse("xxx").Century())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 20, Parse("1990-08-05").Century())
		assert.Equal(t, 21, Parse("2021-08-05").Century())
	})
}

func TestCarbon_Decade(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Zero(t, NewCarbon().Decade())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").Decade())
		assert.Zero(t, Parse("0").Decade())
		assert.Zero(t, Parse("xxx").Decade())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 10, Parse("2010-08-05").Decade())
		assert.Equal(t, 10, Parse("2011-08-05").Decade())
		assert.Equal(t, 20, Parse("2020-08-05").Decade())
		assert.Equal(t, 20, Parse("2021-08-05").Decade())
	})
}

func TestCarbon_Year(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 1, NewCarbon().Year())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").Year())
		assert.Zero(t, Parse("0").Year())
		assert.Zero(t, Parse("xxx").Year())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 2010, Parse("2010-08-05").Year())
		assert.Equal(t, 2011, Parse("2011-08-05").Year())
		assert.Equal(t, 2020, Parse("2020-08-05").Year())
		assert.Equal(t, 2021, Parse("2021-08-05").Year())
	})
}

func TestCarbon_Quarter(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 1, NewCarbon().Quarter())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").Quarter())
		assert.Zero(t, Parse("0").Quarter())
		assert.Zero(t, Parse("xxx").Quarter())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 1, Parse("2020-01-05").Quarter())
		assert.Equal(t, 2, Parse("2020-04-05").Quarter())
		assert.Equal(t, 3, Parse("2020-08-05").Quarter())
		assert.Equal(t, 4, Parse("2020-11-05").Quarter())
	})
}

func TestCarbon_Month(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 1, NewCarbon().Month())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").Month())
		assert.Zero(t, Parse("0").Month())
		assert.Zero(t, Parse("xxx").Month())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 1, Parse("2020-01-05").Month())
		assert.Equal(t, 4, Parse("2020-04-05").Month())
		assert.Equal(t, 8, Parse("2020-08-05").Month())
		assert.Equal(t, 11, Parse("2020-11-05").Month())
	})
}

func TestCarbon_Week(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 0, NewCarbon().Week())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Equal(t, -1, Parse("").Week())
		assert.Equal(t, -1, Parse("0").Week())
		assert.Equal(t, -1, Parse("xxx").Week())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 0, Parse("2020-08-03").Week())
		assert.Equal(t, 1, Parse("2020-08-04").Week())
		assert.Equal(t, 2, Parse("2020-08-05").Week())
		assert.Equal(t, 6, Parse("2020-08-09").Week())
	})
}

func TestCarbon_Day(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 1, NewCarbon().Day())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").Day())
		assert.Zero(t, Parse("0").Day())
		assert.Zero(t, Parse("xxx").Day())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 3, Parse("2020-08-03").Day())
		assert.Equal(t, 4, Parse("2020-08-04").Day())
		assert.Equal(t, 5, Parse("2020-08-05").Day())
		assert.Equal(t, 9, Parse("2020-08-09").Day())
	})
}

func TestCarbon_Hour(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Zero(t, NewCarbon().Hour())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").Hour())
		assert.Zero(t, Parse("0").Hour())
		assert.Zero(t, Parse("xxx").Hour())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 13, Parse("2020-08-05 13:14:15.999999999").Hour())
	})
}

func TestCarbon_Minute(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Zero(t, NewCarbon().Minute())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").Minute())
		assert.Zero(t, Parse("0").Minute())
		assert.Zero(t, Parse("xxx").Minute())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 14, Parse("2020-08-05 13:14:15.999999999").Minute())
	})
}

func TestCarbon_Second(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Zero(t, NewCarbon().Second())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").Second())
		assert.Zero(t, Parse("0").Second())
		assert.Zero(t, Parse("xxx").Second())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 15, Parse("2020-08-05 13:14:15.999999999").Second())
	})
}

func TestCarbon_Millisecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Zero(t, NewCarbon().Millisecond())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").Millisecond())
		assert.Zero(t, Parse("0").Millisecond())
		assert.Zero(t, Parse("xxx").Millisecond())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 999, Parse("2020-08-05 13:14:15.999999999").Millisecond())
	})
}

func TestCarbon_Microsecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Zero(t, NewCarbon().Microsecond())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").Microsecond())
		assert.Zero(t, Parse("0").Microsecond())
		assert.Zero(t, Parse("xxx").Microsecond())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 999999, Parse("2020-08-05 13:14:15.999999999").Microsecond())
	})
}

func TestCarbon_Nanosecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Zero(t, NewCarbon().Nanosecond())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").Nanosecond())
		assert.Zero(t, Parse("0").Nanosecond())
		assert.Zero(t, Parse("xxx").Nanosecond())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 999999999, Parse("2020-08-05 13:14:15.999999999").Nanosecond())
	})
}

func TestCarbon_Timestamp(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, int64(-62135596800), NewCarbon().Timestamp())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("").Timestamp())
		assert.Equal(t, int64(0), Parse("0").Timestamp())
		assert.Equal(t, int64(0), Parse("xxx").Timestamp())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(1577855655), Parse("2020-01-01 13:14:15", PRC).Timestamp())
		assert.Equal(t, int64(1596633255), Parse("2020-08-05 13:14:15.999999999").Timestamp())
	})
}

func TestCarbon_TimestampMilli(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, int64(-62135596800000), NewCarbon().TimestampMilli())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("").TimestampMilli())
		assert.Equal(t, int64(0), Parse("0").TimestampMilli())
		assert.Equal(t, int64(0), Parse("xxx").TimestampMilli())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(1577855655000), Parse("2020-01-01 13:14:15", PRC).TimestampMilli())
		assert.Equal(t, int64(1596633255999), Parse("2020-08-05 13:14:15.999999999").TimestampMilli())
	})
}

func TestCarbon_TimestampMicro(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, int64(-62135596800000000), NewCarbon().TimestampMicro())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("").TimestampMicro())
		assert.Equal(t, int64(0), Parse("0").TimestampMicro())
		assert.Equal(t, int64(0), Parse("xxx").TimestampMicro())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(1577855655000000), Parse("2020-01-01 13:14:15", PRC).TimestampMicro())
		assert.Equal(t, int64(1596633255999999), Parse("2020-08-05 13:14:15.999999999").TimestampMicro())
	})
}

func TestCarbon_TimestampNano(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, int64(-6795364578871345152), NewCarbon().TimestampNano())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("").TimestampNano())
		assert.Equal(t, int64(0), Parse("0").TimestampNano())
		assert.Equal(t, int64(0), Parse("xxx").TimestampNano())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, int64(1577855655000000000), Parse("2020-01-01 13:14:15", PRC).TimestampNano())
		assert.Equal(t, int64(1596633255999999999), Parse("2020-08-05 13:14:15.999999999").TimestampNano())
	})
}

func TestCarbon_Timezone(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, UTC, NewCarbon().Timezone())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").Timezone())
		assert.Empty(t, Parse("0").Timezone())
		assert.Empty(t, Parse("xxx").Timezone())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, UTC, Now().Timezone())
		assert.Equal(t, Tokyo, Now(Tokyo).Timezone())
		assert.Equal(t, PRC, Now(PRC).Timezone())
	})
}

func TestCarbon_ZoneName(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, UTC, NewCarbon().ZoneName())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ZoneName())
		assert.Empty(t, Parse("0").ZoneName())
		assert.Empty(t, Parse("xxx").ZoneName())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "UTC", Now().ZoneName())
		assert.Equal(t, "JST", Now(Tokyo).ZoneName())
		assert.Equal(t, "CST", Now(PRC).ZoneName())
	})
}

func TestCarbon_ZoneOffset(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Zero(t, NewCarbon().ZoneOffset())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").ZoneOffset())
		assert.Zero(t, Parse("0").ZoneOffset())
		assert.Zero(t, Parse("xxx").ZoneOffset())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Zero(t, Parse("2020-08-05").ZoneOffset())
		assert.Equal(t, 32400, Parse("2020-08-05", Tokyo).ZoneOffset())
		assert.Equal(t, 28800, Parse("2020-08-05", PRC).ZoneOffset())
	})
}

func TestCarbon_Locale(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "en", NewCarbon().Locale())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").Locale())
		assert.Empty(t, Parse("0").Locale())
		assert.Empty(t, Parse("xxx").Locale())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Empty(t, Now().SetLocale("").Locale())
		assert.Equal(t, "en", Now().SetLocale("en").Locale())
		assert.Equal(t, "zh-CN", Now().SetLocale("zh-CN").Locale())
	})
}

func TestCarbon_WeekStartsAt(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, DefaultWeekStartsAt, NewCarbon().WeekStartsAt())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").WeekStartsAt())
		assert.Empty(t, Parse("0").WeekStartsAt())
		assert.Empty(t, Parse("xxx").WeekStartsAt())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, Sunday, Now().SetWeekStartsAt(Sunday).WeekStartsAt())
		assert.Equal(t, Monday, Now().SetWeekStartsAt(Monday).WeekStartsAt())
	})
}

func TestCarbon_WeekEndsAt(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, Sunday, NewCarbon().WeekEndsAt())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").WeekEndsAt())
		assert.Empty(t, Parse("0").WeekEndsAt())
		assert.Empty(t, Parse("xxx").WeekEndsAt())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, Saturday, Now().SetWeekStartsAt(Sunday).WeekEndsAt())
		assert.Equal(t, Sunday, Now().SetWeekStartsAt(Monday).WeekEndsAt())
	})
}

func TestCarbon_CurrentLayout(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, DefaultLayout, NewCarbon().CurrentLayout())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").CurrentLayout())
		assert.Empty(t, Parse("0").CurrentLayout())
		assert.Empty(t, Parse("xxx").CurrentLayout())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, DateTimeLayout, Parse("now").CurrentLayout())
		assert.Equal(t, DateTimeLayout, ParseByLayout("2020-08-05 13:14:15", DateTimeLayout).CurrentLayout())
		assert.Equal(t, DateLayout, ParseByLayout("2020-08-05", DateLayout).CurrentLayout())
	})
}

func TestCarbon_Age(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 2024, NewCarbon().Age())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").Age())
		assert.Zero(t, Parse("0").Age())
		assert.Zero(t, Parse("xxx").Age())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Zero(t, Now().AddYears(18).Age())
		assert.Equal(t, 18, Now().SubYears(18).Age())
	})
}
