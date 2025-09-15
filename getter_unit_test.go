package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type GetterSuite struct {
	suite.Suite
}

func TestGetterSuite(t *testing.T) {
	suite.Run(t, new(GetterSuite))
}

func (s *GetterSuite) TestCarbon_StdTime() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Equal(time.Time{}, c.StdTime())
	})

	s.Run("zero carbon", func() {
		s.Equal(time.Time{}, NewCarbon().StdTime())
	})

	s.Run("empty carbon", func() {
		s.Equal(time.Time{}, Parse("").StdTime())
	})

	s.Run("error carbon", func() {
		s.Equal(time.Time{}, Parse("xxx").StdTime())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 00:00:00 +0000 UTC", Parse("2020-08-05").StdTime().String())
		s.Equal("2020-08-05 00:00:00 +0800 CST", Parse("2020-08-05", PRC).StdTime().String())
	})

	s.Run("nil location", func() {
		c := Parse("2020-08-05").SetTimezone(PRC)
		s.Equal("2020-08-05 08:00:00 +0800 CST", c.StdTime().String())
		c.loc = nil
		s.Equal("2020-08-05 00:00:00 +0000 UTC", c.StdTime().String())
	})

	// https://github.com/dromara/carbon/issues/294
	s.Run("issue294", func() {
		s.Equal("0001-01-01 00:00:00 +0000 UTC", Parse("").StdTime().String())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", Parse("0").StdTime().String())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", Parse("xxx").StdTime().String())
	})
}

func (s *GetterSuite) TestCarbon_DaysInYear() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DaysInYear())
	})

	s.Run("zero carbon", func() {
		s.Equal(365, NewCarbon().DaysInYear())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").DaysInYear())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").DaysInYear())
	})

	s.Run("valid carbon", func() {
		s.Equal(366, Parse("2020-08-05").DaysInYear())
		s.Equal(365, Parse("2021-08-05").DaysInYear())
	})
}

func (s *GetterSuite) TestCarbon_DaysInMonth() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DaysInMonth())
	})

	s.Run("zero carbon", func() {
		s.Equal(31, NewCarbon().DaysInMonth())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").DaysInMonth())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").DaysInMonth())
	})

	s.Run("valid carbon", func() {
		s.Equal(31, Parse("2020-01-05").DaysInMonth())
		s.Equal(29, Parse("2020-02-05").DaysInMonth())
		s.Equal(31, Parse("2020-03-05").DaysInMonth())
		s.Equal(30, Parse("2020-04-05").DaysInMonth())
	})
}

func (s *GetterSuite) TestCarbon_MonthOfYear() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.MonthOfYear())
	})

	s.Run("zero carbon", func() {
		s.Equal(1, NewCarbon().MonthOfYear())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").MonthOfYear())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").MonthOfYear())
	})

	s.Run("valid carbon", func() {
		s.Equal(1, Parse("2020-01-05").MonthOfYear())
		s.Equal(2, Parse("2020-02-05").MonthOfYear())
		s.Equal(3, Parse("2020-03-05").MonthOfYear())
		s.Equal(4, Parse("2020-04-05").MonthOfYear())
	})
}

func (s *GetterSuite) TestCarbon_DayOfYear() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DayOfYear())
	})

	s.Run("zero carbon", func() {
		s.Equal(1, NewCarbon().DayOfYear())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").DayOfYear())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").DayOfYear())
	})

	s.Run("valid carbon", func() {
		s.Equal(5, Parse("2020-01-05").DayOfYear())
		s.Equal(36, Parse("2020-02-05").DayOfYear())
		s.Equal(65, Parse("2020-03-05").DayOfYear())
		s.Equal(96, Parse("2020-04-05").DayOfYear())
	})
}

func (s *GetterSuite) TestCarbon_DayOfMonth() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DayOfMonth())
	})

	s.Run("zero carbon", func() {
		s.Equal(1, NewCarbon().DayOfMonth())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").DayOfMonth())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").DayOfMonth())
	})

	s.Run("valid carbon", func() {
		s.Equal(1, Parse("2020-01-01").DayOfMonth())
		s.Equal(5, Parse("2020-01-05").DayOfMonth())
		s.Equal(31, Parse("2020-01-31").DayOfMonth())
	})
}

func (s *GetterSuite) TestCarbon_DayOfWeek() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DayOfWeek())
	})

	s.Run("zero carbon", func() {
		s.Equal(1, NewCarbon().DayOfWeek())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").DayOfWeek())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").DayOfWeek())
	})

	s.Run("valid carbon", func() {
		s.Equal(1, Parse("2020-08-03").DayOfWeek())
		s.Equal(2, Parse("2020-08-04").DayOfWeek())
		s.Equal(3, Parse("2020-08-05").DayOfWeek())
		s.Equal(7, Parse("2020-08-09").DayOfWeek())
	})
}

func (s *GetterSuite) TestCarbon_WeekOfYear() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.WeekOfYear())
	})

	s.Run("zero carbon", func() {
		s.Equal(1, NewCarbon().WeekOfYear())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").WeekOfYear())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").WeekOfYear())
	})

	s.Run("valid carbon", func() {
		s.Equal(53, Parse("2021-01-01").WeekOfYear())
		s.Equal(5, Parse("2021-02-01").WeekOfYear())
		s.Equal(9, Parse("2021-03-01").WeekOfYear())
		s.Equal(13, Parse("2021-04-01").WeekOfYear())
	})
}

func (s *GetterSuite) TestCarbon_WeekOfMonth() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.WeekOfMonth())
	})

	s.Run("zero carbon", func() {
		s.Equal(1, NewCarbon().WeekOfMonth())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").WeekOfMonth())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").WeekOfMonth())
	})

	s.Run("valid carbon", func() {
		s.Equal(1, Parse("2021-07-01").WeekOfMonth())
		s.Equal(1, Parse("2021-07-02").WeekOfMonth())
		s.Equal(1, Parse("2021-07-03").WeekOfMonth())
		s.Equal(1, Parse("2021-07-04").WeekOfMonth())
		s.Equal(2, Parse("2021-07-05").WeekOfMonth())
		s.Equal(2, Parse("2021-07-06").WeekOfMonth())
	})
}

func (s *GetterSuite) TestCarbon_DateTime() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DateTime())
	})

	s.Run("zero carbon", func() {
		year, month, day, hour, minute, second := NewCarbon().DateTime()
		s.Equal(1, year, month, day)
		s.Zero(hour, minute, second)
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").DateTime())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").DateTime())
	})

	s.Run("valid carbon", func() {
		year, month, day, hour, minute, second := Parse("2020-08-05 13:14:15.999999999").DateTime()
		s.Equal(2020, year)
		s.Equal(8, month)
		s.Equal(5, day)
		s.Equal(13, hour)
		s.Equal(14, minute)
		s.Equal(15, second)
	})
}

func (s *GetterSuite) TestCarbon_DateTimeMilli() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DateTimeMilli())
	})

	s.Run("zero carbon", func() {
		year, month, day, hour, minute, second, millisecond := NewCarbon().DateTimeMilli()
		s.Equal(1, year, month, day)
		s.Zero(hour, minute, second, millisecond)
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").DateTimeMilli())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").DateTimeMilli())
	})

	s.Run("valid carbon", func() {
		year, month, day, hour, minute, second, millisecond := Parse("2020-08-05 13:14:15.999999999").DateTimeMilli()
		s.Equal(2020, year)
		s.Equal(8, month)
		s.Equal(5, day)
		s.Equal(13, hour)
		s.Equal(14, minute)
		s.Equal(15, second)
		s.Equal(999, millisecond)
	})
}

func (s *GetterSuite) TestCarbon_DateTimeMicro() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DateTimeMicro())
	})

	s.Run("zero carbon", func() {
		year, month, day, hour, minute, second, microsecond := NewCarbon().DateTimeMicro()
		s.Equal(1, year, month, day)
		s.Zero(hour, minute, second, microsecond)
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").DateTimeMicro())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").DateTimeMicro())
	})

	s.Run("valid carbon", func() {
		year, month, day, hour, minute, second, microsecond := Parse("2020-08-05 13:14:15.999999999").DateTimeMicro()
		s.Equal(2020, year)
		s.Equal(8, month)
		s.Equal(5, day)
		s.Equal(13, hour)
		s.Equal(14, minute)
		s.Equal(15, second)
		s.Equal(999999, microsecond)
	})
}

func (s *GetterSuite) TestCarbon_DateTimeNano() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DateTimeNano())
	})

	s.Run("zero carbon", func() {
		year, month, day, hour, minute, second, nanosecond := NewCarbon().DateTimeNano()
		s.Equal(1, year, month, day)
		s.Zero(hour, minute, second, nanosecond)
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").DateTimeNano())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").DateTimeNano())
	})

	s.Run("valid carbon", func() {
		year, month, day, hour, minute, second, nanosecond := Parse("2020-08-05 13:14:15.999999999").DateTimeNano()
		s.Equal(2020, year)
		s.Equal(8, month)
		s.Equal(5, day)
		s.Equal(13, hour)
		s.Equal(14, minute)
		s.Equal(15, second)
		s.Equal(999999999, nanosecond)
	})
}

func (s *GetterSuite) TestCarbon_Date() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.Date())
	})

	s.Run("zero carbon", func() {
		year, month, day := NewCarbon().Date()
		s.Equal(1, year, month, day)
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").Date())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").Date())
	})

	s.Run("valid carbon", func() {
		year, month, day := Parse("2020-08-05 13:14:15.999999999").Date()
		s.Equal(2020, year)
		s.Equal(8, month)
		s.Equal(5, day)
	})
}

func (s *GetterSuite) TestCarbon_DateMilli() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DateMilli())
	})

	s.Run("zero carbon", func() {
		year, month, day, millisecond := NewCarbon().DateMilli()
		s.Equal(1, year, month, day)
		s.Zero(millisecond)
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").DateMilli())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").DateMilli())
	})

	s.Run("valid carbon", func() {
		year, month, day, millisecond := Parse("2020-08-05 13:14:15.999999999").DateMilli()
		s.Equal(2020, year)
		s.Equal(8, month)
		s.Equal(5, day)
		s.Equal(999, millisecond)
	})
}

func (s *GetterSuite) TestCarbon_DateMicro() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DateMicro())
	})

	s.Run("zero carbon", func() {
		year, month, day, microsecond := NewCarbon().DateMicro()
		s.Equal(1, year, month, day)
		s.Zero(microsecond)
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").DateMicro())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").DateMicro())
	})

	s.Run("valid carbon", func() {
		year, month, day, microsecond := Parse("2020-08-05 13:14:15.999999999").DateMicro()
		s.Equal(2020, year)
		s.Equal(8, month)
		s.Equal(5, day)
		s.Equal(999999, microsecond)
	})
}

func (s *GetterSuite) TestCarbon_DateNano() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DateNano())
	})

	s.Run("zero carbon", func() {
		year, month, day, nanosecond := NewCarbon().DateNano()
		s.Equal(1, year, month, day)
		s.Zero(nanosecond)
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").DateNano())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").DateNano())
	})

	s.Run("valid carbon", func() {
		year, month, day, nanosecond := Parse("2020-08-05 13:14:15.999999999").DateNano()
		s.Equal(2020, year)
		s.Equal(8, month)
		s.Equal(5, day)
		s.Equal(999999999, nanosecond)
	})
}

func (s *GetterSuite) TestCarbon_Time() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.Time())
	})

	s.Run("zero carbon", func() {
		s.Zero(NewCarbon().Time())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").Time())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").Time())
	})

	s.Run("valid carbon", func() {
		hour, minute, second := Parse("2020-08-05 13:14:15.999999999").Time()
		s.Equal(13, hour)
		s.Equal(14, minute)
		s.Equal(15, second)
	})
}

func (s *GetterSuite) TestCarbon_TimeMilli() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.TimeMilli())
	})

	s.Run("zero carbon", func() {
		s.Zero(NewCarbon().TimeMilli())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").TimeMilli())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").TimeMilli())
	})

	s.Run("valid carbon", func() {
		hour, minute, second, millisecond := Parse("2020-08-05 13:14:15.999999999").TimeMilli()
		s.Equal(13, hour)
		s.Equal(14, minute)
		s.Equal(15, second)
		s.Equal(999, millisecond)
	})
}

func (s *GetterSuite) TestCarbon_TimeMicro() {
	var c *Carbon
	c = nil
	s.Run("nil carbon", func() {
		s.Zero(c.TimeMicro())
	})

	s.Run("zero carbon", func() {
		s.Zero(NewCarbon().TimeMicro())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").TimeMicro())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").TimeMicro())
	})

	s.Run("valid carbon", func() {
		hour, minute, second, microsecond := Parse("2020-08-05 13:14:15.999999999").TimeMicro()
		s.Equal(13, hour)
		s.Equal(14, minute)
		s.Equal(15, second)
		s.Equal(999999, microsecond)
	})
}

func (s *GetterSuite) TestCarbon_TimeNano() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.TimeNano())
	})

	s.Run("zero carbon", func() {
		s.Zero(NewCarbon().TimeNano())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").TimeNano())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").TimeNano())
	})

	s.Run("valid carbon", func() {
		hour, minute, second, nanosecond := Parse("2020-08-05 13:14:15.999999999").TimeNano()
		s.Equal(13, hour)
		s.Equal(14, minute)
		s.Equal(15, second)
		s.Equal(999999999, nanosecond)
	})
}

func (s *GetterSuite) TestCarbon_Century() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.Century())
	})

	s.Run("zero carbon", func() {
		s.Equal(1, NewCarbon().Century())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").Century())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").Century())
	})

	s.Run("valid carbon", func() {
		s.Equal(20, Parse("1990-08-05").Century())
		s.Equal(21, Parse("2021-08-05").Century())
	})
}

func (s *GetterSuite) TestCarbon_Decade() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.Decade())
	})

	s.Run("zero carbon", func() {
		s.Zero(NewCarbon().Decade())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").Decade())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").Decade())
	})

	s.Run("valid carbon", func() {
		s.Equal(10, Parse("2010-08-05").Decade())
		s.Equal(10, Parse("2011-08-05").Decade())
		s.Equal(20, Parse("2020-08-05").Decade())
		s.Equal(20, Parse("2021-08-05").Decade())
	})
}

func (s *GetterSuite) TestCarbon_Year() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.Year())
	})

	s.Run("zero carbon", func() {
		s.Equal(1, NewCarbon().Year())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").Year())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").Year())
	})

	s.Run("valid carbon", func() {
		s.Equal(2010, Parse("2010-08-05").Year())
		s.Equal(2011, Parse("2011-08-05").Year())
		s.Equal(2020, Parse("2020-08-05").Year())
		s.Equal(2021, Parse("2021-08-05").Year())
	})
}

func (s *GetterSuite) TestCarbon_Quarter() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.Quarter())
	})

	s.Run("zero carbon", func() {
		s.Equal(1, NewCarbon().Quarter())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").Quarter())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").Quarter())
	})

	s.Run("valid carbon", func() {
		s.Equal(1, Parse("2020-01-05").Quarter())
		s.Equal(2, Parse("2020-04-05").Quarter())
		s.Equal(3, Parse("2020-08-05").Quarter())
		s.Equal(4, Parse("2020-11-05").Quarter())
	})
}

func (s *GetterSuite) TestCarbon_Month() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.Month())
	})

	s.Run("zero carbon", func() {
		s.Equal(1, NewCarbon().Month())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").Month())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").Month())
	})

	s.Run("valid carbon", func() {
		s.Equal(1, Parse("2020-01-05").Month())
		s.Equal(4, Parse("2020-04-05").Month())
		s.Equal(8, Parse("2020-08-05").Month())
		s.Equal(11, Parse("2020-11-05").Month())
	})
}

func (s *GetterSuite) TestCarbon_Week() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Equal(-1, c.Week())
	})

	s.Run("zero carbon", func() {
		s.Equal(0, NewCarbon().Week())
	})

	s.Run("empty carbon", func() {
		s.Equal(-1, Parse("").Week())
	})

	s.Run("error carbon", func() {
		s.Equal(-1, Parse("xxx").Week())
	})

	s.Run("valid carbon", func() {
		s.Equal(0, Parse("2020-08-03").Week())
		s.Equal(1, Parse("2020-08-04").Week())
		s.Equal(2, Parse("2020-08-05").Week())
		s.Equal(6, Parse("2020-08-09").Week())
	})
}

func (s *GetterSuite) TestCarbon_Day() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.Day())
	})

	s.Run("zero carbon", func() {
		s.Equal(1, NewCarbon().Day())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").Day())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").Day())
	})

	s.Run("valid carbon", func() {
		s.Equal(3, Parse("2020-08-03").Day())
		s.Equal(4, Parse("2020-08-04").Day())
		s.Equal(5, Parse("2020-08-05").Day())
		s.Equal(9, Parse("2020-08-09").Day())
	})
}

func (s *GetterSuite) TestCarbon_Hour() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.Hour())
	})

	s.Run("zero carbon", func() {
		s.Zero(NewCarbon().Hour())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").Hour())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").Hour())
	})

	s.Run("valid carbon", func() {
		s.Equal(13, Parse("2020-08-05 13:14:15.999999999").Hour())
	})
}

func (s *GetterSuite) TestCarbon_Minute() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.Minute())
	})

	s.Run("zero carbon", func() {
		s.Zero(NewCarbon().Minute())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").Minute())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").Minute())
	})

	s.Run("valid carbon", func() {
		s.Equal(14, Parse("2020-08-05 13:14:15.999999999").Minute())
	})
}

func (s *GetterSuite) TestCarbon_Second() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.Second())
	})

	s.Run("zero carbon", func() {
		s.Zero(NewCarbon().Second())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").Second())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").Second())
	})

	s.Run("valid carbon", func() {
		s.Equal(15, Parse("2020-08-05 13:14:15.999999999").Second())
	})
}

func (s *GetterSuite) TestCarbon_Millisecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.Millisecond())
	})

	s.Run("zero carbon", func() {
		s.Zero(NewCarbon().Millisecond())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").Millisecond())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").Millisecond())
	})

	s.Run("valid carbon", func() {
		s.Equal(999, Parse("2020-08-05 13:14:15.999999999").Millisecond())
	})
}

func (s *GetterSuite) TestCarbon_Microsecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.Microsecond())
	})

	s.Run("zero carbon", func() {
		s.Zero(NewCarbon().Microsecond())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").Microsecond())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").Microsecond())
	})

	s.Run("valid carbon", func() {
		s.Equal(999999, Parse("2020-08-05 13:14:15.999999999").Microsecond())
	})
}

func (s *GetterSuite) TestCarbon_Nanosecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.Nanosecond())
	})

	s.Run("zero carbon", func() {
		s.Zero(NewCarbon().Nanosecond())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").Nanosecond())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").Nanosecond())
	})

	s.Run("valid carbon", func() {
		s.Equal(999999999, Parse("2020-08-05 13:14:15.999999999").Nanosecond())
	})
}

func (s *GetterSuite) TestCarbon_Timestamp() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.Timestamp())
	})

	s.Run("zero carbon", func() {
		s.Equal(int64(-62135596800), NewCarbon().Timestamp())
	})

	s.Run("empty carbon", func() {
		s.Equal(int64(0), Parse("").Timestamp())
	})

	s.Run("error carbon", func() {
		s.Equal(int64(0), Parse("xxx").Timestamp())
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(1577855655), Parse("2020-01-01 13:14:15", PRC).Timestamp())
		s.Equal(int64(1596633255), Parse("2020-08-05 13:14:15.999999999").Timestamp())
	})
}

func (s *GetterSuite) TestCarbon_TimestampMilli() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.TimestampMilli())
	})

	s.Run("zero carbon", func() {
		s.Equal(int64(-62135596800000), NewCarbon().TimestampMilli())
	})

	s.Run("empty carbon", func() {
		s.Equal(int64(0), Parse("").TimestampMilli())
	})

	s.Run("error carbon", func() {
		s.Equal(int64(0), Parse("xxx").TimestampMilli())
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(1577855655000), Parse("2020-01-01 13:14:15", PRC).TimestampMilli())
		s.Equal(int64(1596633255999), Parse("2020-08-05 13:14:15.999999999").TimestampMilli())
	})
}

func (s *GetterSuite) TestCarbon_TimestampMicro() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.TimestampMicro())
	})

	s.Run("zero carbon", func() {
		s.Equal(int64(-62135596800000000), NewCarbon().TimestampMicro())
	})

	s.Run("empty carbon", func() {
		s.Equal(int64(0), Parse("").TimestampMicro())
	})

	s.Run("error carbon", func() {
		s.Equal(int64(0), Parse("xxx").TimestampMicro())
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(1577855655000000), Parse("2020-01-01 13:14:15", PRC).TimestampMicro())
		s.Equal(int64(1596633255999999), Parse("2020-08-05 13:14:15.999999999").TimestampMicro())
	})
}

func (s *GetterSuite) TestCarbon_TimestampNano() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.TimestampNano())
	})

	s.Run("zero carbon", func() {
		s.Equal(int64(-6795364578871345152), NewCarbon().TimestampNano())
	})

	s.Run("empty carbon", func() {
		s.Equal(int64(0), Parse("").TimestampNano())
	})

	s.Run("error carbon", func() {
		s.Equal(int64(0), Parse("xxx").TimestampNano())
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(1577855655000000000), Parse("2020-01-01 13:14:15", PRC).TimestampNano())
		s.Equal(int64(1596633255999999999), Parse("2020-08-05 13:14:15.999999999").TimestampNano())
	})
}

func (s *GetterSuite) TestCarbon_Timezone() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.Timezone())
	})

	s.Run("zero carbon", func() {
		s.Equal(UTC, NewCarbon().Timezone())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").Timezone())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").Timezone())
	})

	s.Run("valid carbon", func() {
		s.Equal(UTC, Now().Timezone())
		s.Equal(Tokyo, Now(Tokyo).Timezone())
		s.Equal(PRC, Now(PRC).Timezone())
	})
}

func (s *GetterSuite) TestCarbon_ZoneName() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.ZoneName())
	})

	s.Run("zero carbon", func() {
		s.Equal(UTC, NewCarbon().ZoneName())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ZoneName())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ZoneName())
	})

	s.Run("valid carbon", func() {
		s.Equal("UTC", Now().ZoneName())
		s.Equal("JST", Now(Tokyo).ZoneName())
		s.Equal("CST", Now(PRC).ZoneName())
	})
}

func (s *GetterSuite) TestCarbon_ZoneOffset() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.ZoneOffset())
	})

	s.Run("zero carbon", func() {
		s.Zero(NewCarbon().ZoneOffset())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").ZoneOffset())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").ZoneOffset())
	})

	s.Run("valid carbon", func() {
		s.Zero(Parse("2020-08-05").ZoneOffset())
		s.Equal(32400, Parse("2020-08-05", Tokyo).ZoneOffset())
		s.Equal(28800, Parse("2020-08-05", PRC).ZoneOffset())
	})
}

func (s *GetterSuite) TestCarbon_Locale() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.Locale())
	})

	s.Run("zero carbon", func() {
		s.Equal("en", NewCarbon().Locale())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").Locale())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").Locale())
	})

	s.Run("nil lang", func() {
		c := Now()
		c.lang = nil
		s.Empty(c.Locale())
	})

	s.Run("valid carbon", func() {
		s.Empty(Now().SetLocale("").Locale())
		s.Equal("en", Now().SetLocale("en").Locale())
		s.Equal("zh-CN", Now().SetLocale("zh-CN").Locale())
	})
}

func (s *GetterSuite) TestCarbon_WeekStartsAt() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.WeekStartsAt())
	})

	s.Run("zero carbon", func() {
		s.Equal(DefaultWeekStartsAt, NewCarbon().WeekStartsAt())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").WeekStartsAt())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").WeekStartsAt())
	})

	s.Run("valid carbon", func() {
		s.Equal(Sunday, Now().SetWeekStartsAt(Sunday).WeekStartsAt())
		s.Equal(Monday, Now().SetWeekStartsAt(Monday).WeekStartsAt())
	})
}

func (s *GetterSuite) TestCarbon_WeekEndsAt() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.WeekEndsAt())
	})

	s.Run("zero carbon", func() {
		s.Equal(Sunday, NewCarbon().WeekEndsAt())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").WeekEndsAt())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").WeekEndsAt())
	})

	s.Run("valid carbon", func() {
		s.Equal(Saturday, Now().SetWeekStartsAt(Sunday).WeekEndsAt())
		s.Equal(Sunday, Now().SetWeekStartsAt(Monday).WeekEndsAt())
	})
}

func (s *GetterSuite) TestCarbon_CurrentLayout() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.CurrentLayout())
	})

	s.Run("zero carbon", func() {
		s.Equal(DefaultLayout, NewCarbon().CurrentLayout())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").CurrentLayout())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").CurrentLayout())
	})

	s.Run("valid carbon", func() {
		s.Equal(DateTimeLayout, Parse("now").CurrentLayout())
		s.Equal(DateTimeLayout, ParseByLayout("2020-08-05 13:14:15", DateTimeLayout).CurrentLayout())
		s.Equal(DateLayout, ParseByLayout("2020-08-05", DateLayout).CurrentLayout())
	})
}

func (s *GetterSuite) TestCarbon_Age() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.Age())
	})

	s.Run("zero carbon", func() {
		s.Equal(2024, NewCarbon().Age())
	})

	s.Run("empty carbon", func() {
		s.Zero(Parse("").Age())
	})

	s.Run("error carbon", func() {
		s.Zero(Parse("xxx").Age())
	})

	s.Run("valid carbon", func() {
		s.Zero(Now().AddYears(18).Age())
		s.Equal(18, Now().SubYears(18).Age())
	})
}
