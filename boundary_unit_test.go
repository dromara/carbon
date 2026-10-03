package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type BoundarySuite struct {
	suite.Suite
}

func TestBoundarySuite(t *testing.T) {
	suite.Run(t, new(BoundarySuite))
}

func (s *BoundarySuite) TestCarbon_SecondBoundariesAtOffsetTransitions() {
	for _, test := range []struct {
		name     string
		zone     string
		inputUTC time.Time
	}{
		{"New York first occurrence", "America/New_York", time.Date(2026, 11, 1, 5, 34, 49, 123456789, time.UTC)},
		{"New York second occurrence", "America/New_York", time.Date(2026, 11, 1, 6, 34, 49, 123456789, time.UTC)},
		{"Lord Howe first occurrence", "Australia/Lord_Howe", time.Date(2026, 4, 4, 14, 45, 49, 123456789, time.UTC)},
		{"Lord Howe second occurrence", "Australia/Lord_Howe", time.Date(2026, 4, 4, 15, 15, 49, 123456789, time.UTC)},
		{"Algiers second-offset transition", "Africa/Algiers", time.Date(1891, 3, 15, 23, 48, 22, 123456789, time.UTC)},
		{"Monrovia forward transition", "Africa/Monrovia", time.Date(1972, 1, 7, 0, 44, 30, 123456789, time.UTC)},
		{"UTC control", "UTC", time.Date(2026, 1, 2, 12, 34, 49, 123456789, time.UTC)},
	} {
		s.Run(test.name, func() {
			loc, err := time.LoadLocation(test.zone)
			s.Require().NoError(err)
			input := test.inputUTC.In(loc)
			c := NewCarbon(input).SetWeekStartsAt(Sunday)
			start, end := c.StartOfSecond(), c.EndOfSecond()
			s.Equal(input.Unix(), start.StdTime().Unix())
			s.Equal(input.Unix(), end.StdTime().Unix())
			s.Equal(0, start.Nanosecond())
			s.Equal(MaxNanosecond, end.Nanosecond())
			s.Same(loc, start.StdTime().Location())
			s.Same(loc, end.StdTime().Location())
			s.Equal(Sunday, start.WeekStartsAt())
			s.Equal(Sunday, end.WeekStartsAt())
			s.Equal(input, c.StdTime())
		})
	}
}

func (s *BoundarySuite) TestCarbon_StartOfCentury() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.StartOfCentury()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().StartOfCentury()
		s.False(c.HasError())
		s.Equal("0000-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfCentury()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfCentury()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2000-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfCentury().ToString())
		s.Equal("2000-01-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfCentury().ToString())
		s.Equal("2000-01-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfCentury().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfCentury() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.EndOfCentury()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().EndOfCentury()
		s.False(c.HasError())
		s.Equal("0099-12-31 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfCentury()
		s.False(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfCentury()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2099-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfCentury().ToString())
		s.Equal("2099-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfCentury().ToString())
		s.Equal("2099-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfCentury().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfDecade() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.StartOfDecade()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().StartOfDecade()
		s.False(c.HasError())
		s.Equal("0000-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfDecade()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfDecade()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfDecade().ToString())
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfDecade().ToString())
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfDecade().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfDecade() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.EndOfDecade()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().EndOfDecade()
		s.False(c.HasError())
		s.Equal("0009-12-31 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfDecade()
		s.False(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfDecade()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2029-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfDecade().ToString())
		s.Equal("2029-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfDecade().ToString())
		s.Equal("2029-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfDecade().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfYear() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.StartOfYear()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().StartOfYear()
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfYear()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfYear()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfYear().ToString())
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfYear().ToString())
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfYear().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfYear() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.EndOfYear()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().EndOfYear()
		s.False(c.HasError())
		s.Equal("0001-12-31 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfYear()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfYear()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfYear().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfYear().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfYear().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfQuarter() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.StartOfQuarter()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().StartOfQuarter()
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfQuarter()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfQuarter()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfQuarter().ToString())
		s.Equal("2020-07-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfQuarter().ToString())
		s.Equal("2020-10-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfQuarter().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfQuarter() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.EndOfQuarter()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().EndOfQuarter()
		s.False(c.HasError())
		s.Equal("0001-03-31 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfQuarter()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfQuarter()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-03-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfQuarter().ToString())
		s.Equal("2020-09-30 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfQuarter().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfQuarter().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfMonth() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.StartOfMonth()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().StartOfMonth()
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfMonth()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfMonth()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfMonth().ToString())
		s.Equal("2020-08-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfMonth().ToString())
		s.Equal("2020-12-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfMonth().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfMonth() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.EndOfMonth()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().EndOfMonth()
		s.False(c.HasError())
		s.Equal("0001-01-31 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfMonth()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfMonth()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfMonth().ToString())
		s.Equal("2020-08-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfMonth().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfMonth().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfWeek() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.StartOfWeek()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().StartOfWeek()
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfWeek()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfWeek()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-12-30 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfWeek().ToString())
		s.Equal("2020-08-10 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfWeek().ToString())
		s.Equal("2020-12-28 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfWeek().ToString())
		s.Equal("2025-04-07 00:00:00 +0000 UTC", Parse("2025-04-07 00:00:00").StartOfWeek().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfWeek() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.EndOfWeek()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().EndOfWeek()
		s.False(c.HasError())
		s.Equal("0001-01-07 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfWeek()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfWeek()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-05 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfWeek().ToString())
		s.Equal("2020-08-16 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfWeek().ToString())
		s.Equal("2021-01-03 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfWeek().ToString())
		s.Equal("2025-04-13 23:59:59.999999999 +0000 UTC", Parse("2025-04-13 00:00:00").EndOfWeek().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfDay() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.StartOfDay()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().StartOfDay()
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfDay()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfDay()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfDay().ToString())
		s.Equal("2020-08-15 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfDay().ToString())
		s.Equal("2020-12-31 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfDay().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfDay() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.EndOfDay()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().EndOfDay()
		s.False(c.HasError())
		s.Equal("0001-01-01 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfDay()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfDay()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfDay().ToString())
		s.Equal("2020-08-15 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfDay().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfDay().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfHour() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.StartOfHour()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().StartOfHour()
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfHour()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfHour()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfHour().ToString())
		s.Equal("2020-08-15 12:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfHour().ToString())
		s.Equal("2020-12-31 23:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfHour().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfHour() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.EndOfHour()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().EndOfHour()
		s.False(c.HasError())
		s.Equal("0001-01-01 00:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfHour()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfHour()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfHour().ToString())
		s.Equal("2020-08-15 12:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfHour().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfHour().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfMinute() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.StartOfMinute()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().StartOfMinute()
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfMinute()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfMinute()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfMinute().ToString())
		s.Equal("2020-08-15 12:30:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfMinute().ToString())
		s.Equal("2020-12-31 23:59:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfMinute().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfMinute() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.EndOfMinute()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().EndOfMinute()
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfMinute()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfMinute()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfMinute().ToString())
		s.Equal("2020-08-15 12:30:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfMinute().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfMinute().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfSecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.StartOfSecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().StartOfSecond()
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfSecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfSecond()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfSecond().ToString())
		s.Equal("2020-08-15 12:30:30 +0000 UTC", Parse("2020-08-15 12:30:30.66666").StartOfSecond().ToString())
		s.Equal("2020-12-31 23:59:59 +0000 UTC", Parse("2020-12-31 23:59:59.999999999").StartOfSecond().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfSecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.EndOfSecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().EndOfSecond()
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfSecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfSecond()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfSecond().ToString())
		s.Equal("2020-08-15 12:30:30.999999999 +0000 UTC", Parse("2020-08-15 12:30:30.66666").EndOfSecond().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59.999999999").EndOfSecond().ToString())
	})
}
