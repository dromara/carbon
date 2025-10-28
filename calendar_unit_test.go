package carbon

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CalendarSuite struct {
	suite.Suite
}

func TestCalendarSuite(t *testing.T) {
	suite.Run(t, new(CalendarSuite))
}

func (s *CalendarSuite) TestCarbon_Julian() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.Julian())
	})

	s.Run("zero carbon", func() {
		j := NewCarbon().Julian()
		s.Equal(1.7214235e+06, j.JD())
		s.Equal(float64(-678577), j.MJD())
	})

	s.Run("empty carbon", func() {
		j := Parse("").Julian()
		s.Zero(j.JD())
		s.Zero(j.MJD())
	})

	s.Run("error carbon", func() {
		j := Parse("xxx").Julian()
		s.Zero(j.JD())
		s.Zero(j.MJD())
	})

	s.Run("valid carbon", func() {
		j := Parse("2024-01-23 13:14:15").Julian()

		s.Equal(2460333.051563, j.JD())
		s.Equal(60332.551563, j.MJD())

		s.Equal(2460333.0516, j.JD(4))
		s.Equal(60332.5516, j.MJD(4))

		s.Equal(2460333.0515625, j.JD(7))
		s.Equal(60332.5515625, j.MJD(7))
	})
}

func (s *CalendarSuite) TestCreateFromJulian() {
	s.Run("error julian", func() {
		s.Equal("-4712-01-01 12:00:00", CreateFromJulian(0).String())
		s.Equal("-4712-01-01 12:00:00", CreateFromJulian(-1).String())
	})

	s.Run("valid julian", func() {
		s.Equal("2024-01-23 13:14:15 +0000 UTC", CreateFromJulian(2460333.051563).ToString())
		s.Equal("2024-01-23 13:14:15 +0000 UTC", CreateFromJulian(60332.551563).ToString())

		s.Equal("2024-01-23 13:14:18 +0000 UTC", CreateFromJulian(2460333.0516).ToString())
		s.Equal("2024-01-23 13:14:18 +0000 UTC", CreateFromJulian(60332.5516).ToString())

		s.Equal("2024-01-23 13:14:15 +0000 UTC", CreateFromJulian(2460333.0515625).ToString())
		s.Equal("2024-01-23 13:14:15 +0000 UTC", CreateFromJulian(60332.5515625).ToString())
	})
}

func (s *CalendarSuite) TestCarbon_Lunar() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.Lunar())
	})

	s.Run("zero carbon", func() {
		l := NewCarbon().Lunar()
		s.Nil(l.Error)
		s.Empty(l.String())
	})

	s.Run("empty carbon", func() {
		l := Parse("").Lunar()
		s.Nil(l.Error)
		s.Empty(l.String())
	})

	s.Run("error carbon", func() {
		l := Parse("xxx").Lunar()
		s.Error(l.Error)
		s.Empty(l.String())
	})

	s.Run("valid carbon", func() {
		s.Equal("2023-12-08", Parse("2024-01-18", PRC).Lunar().String())
		s.Equal("2023-12-11", Parse("2024-01-21", PRC).Lunar().String())
		s.Equal("2023-12-14", Parse("2024-01-24", PRC).Lunar().String())
	})
}

func (s *CalendarSuite) TestCreateFromLunar() {
	s.Run("error lunar", func() {
		s.Empty(CreateFromLunar(2200, 12, 14, false).ToString())
	})

	s.Run("valid lunar", func() {
		s.Equal("2024-01-21 00:00:00 +0800 CST", CreateFromLunar(2023, 12, 11, false).ToString(PRC))
		s.Equal("2024-01-18 00:00:00 +0800 CST", CreateFromLunar(2023, 12, 8, false).ToString(PRC))
		s.Equal("2024-01-24 00:00:00 +0800 CST", CreateFromLunar(2023, 12, 14, false).ToString(PRC))
	})
}

func (s *CalendarSuite) TestCarbon_Persian() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.Persian())
	})

	s.Run("zero carbon", func() {
		p := NewCarbon().Persian()
		s.Nil(p.Error)
		s.Empty(p.String())
	})

	s.Run("empty carbon", func() {
		p := Parse("").Persian()
		s.Nil(p.Error)
		s.Empty(p.String())
	})

	s.Run("error carbon", func() {
		p := Parse("xxx").Persian()
		s.Error(p.Error)
		s.Empty(p.String())
	})

	s.Run("valid carbon", func() {
		s.Equal("1178-10-11", Parse("1800-01-01 00:00:00").Persian().String())
		s.Equal("1399-05-15", Parse("2020-08-05 13:14:15").Persian().String())
		s.Equal("1402-10-11", Parse("2024-01-01 00:00:00").Persian().String())
	})
}

func (s *CalendarSuite) TestCreateFromPersian() {
	s.Run("error persian", func() {
		s.Empty(CreateFromPersian(9999, 12, 31).ToDateTimeString())
	})

	s.Run("valid persian", func() {
		s.Equal("1800-01-01 00:00:00", CreateFromPersian(1178, 10, 11).ToDateTimeString())
		s.Equal("2024-01-01 00:00:00", CreateFromPersian(1402, 10, 11).ToDateTimeString())
		s.Equal("2024-08-05 00:00:00", CreateFromPersian(1403, 5, 15).ToDateTimeString())
	})
}

func (s *CalendarSuite) TestCarbon_Hebrew() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.Hebrew())
	})

	s.Run("zero carbon", func() {
		h := NewCarbon().Hebrew()
		s.Nil(h.Error)
		s.Empty(h.String())
	})

	s.Run("empty carbon", func() {
		h := Parse("").Hebrew()
		s.Nil(h.Error)
		s.Empty(h.String())
	})

	s.Run("error carbon", func() {
		h := Parse("xxx").Hebrew()
		s.Error(h.Error)
		s.Empty(h.String())
	})

	s.Run("valid carbon", func() {
		s.Equal("5784-10-20", Parse("2024-01-01 00:00:00").Hebrew().String())
		s.Equal("5784-05-01", Parse("2024-08-05 00:00:00").Hebrew().String())
		s.Equal("5786-07-10", Parse("2025-10-03 00:00:00").Hebrew().String())
		s.Empty(Parse("0001-01-01 00:00:00").Hebrew().String())
	})
}

func (s *CalendarSuite) TestCreateFromHebrew() {
	s.Run("error hebrew", func() {
		// Test invalid Hebrew date that would cause error
		c := CreateFromHebrew(10000, 13, 1)
		s.Error(c.Error)
	})

	s.Run("valid hebrew", func() {
		s.Equal("2023-12-17 12:00:00 +0000 UTC", CreateFromHebrew(5784, 10, 20).ToString())
		s.Equal("2024-07-21 12:00:00 +0000 UTC", CreateFromHebrew(5784, 5, 1).ToString())
		s.Equal("2025-09-18 12:00:00 +0000 UTC", CreateFromHebrew(5786, 7, 10).ToString())
		s.Equal("0001-01-01 12:00:00 +0000 UTC", CreateFromHebrew(3761, 10, 18).ToString())
	})

	s.Run("leap year hebrew", func() {
		// Test leap year with Adar Bet (month 13)
		s.Equal("2024-02-25 12:00:00 +0000 UTC", CreateFromHebrew(5784, 13, 1).ToString())
	})

	s.Run("boundary hebrew", func() {
		// Test boundary values
		s.Equal("-3759-04-01 12:00:00 +0000 UTC", CreateFromHebrew(1, 1, 1).ToString())
		s.Equal("2024-03-25 12:00:00 +0000 UTC", CreateFromHebrew(5784, 1, 1).ToString())
	})
}
