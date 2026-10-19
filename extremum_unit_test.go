package carbon

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExtremumSuite struct {
	suite.Suite
}

func TestExtremumSuite(t *testing.T) {
	suite.Run(t, new(ExtremumSuite))
}

func (s *ExtremumSuite) TestZeroValue() {
	s.Equal("0001-01-01 00:00:00 +0000 UTC", ZeroValue().ToString())
	s.Equal("January", ZeroValue().ToMonthString())
	s.Equal("一月", ZeroValue().SetLocale("zh-CN").ToMonthString())
}

func (s *ExtremumSuite) TestEpochValue() {
	s.Equal("1970-01-01 00:00:00 +0000 UTC", EpochValue().ToString())
	s.Equal("January", EpochValue().ToMonthString())
	s.Equal("一月", EpochValue().SetLocale("zh-CN").ToMonthString())
}

func (s *ExtremumSuite) TestMaxValue() {
	s.Equal("9999-12-31 23:59:59.999999999 +0000 UTC", MaxValue().ToString())
	s.Equal("December", MaxValue().ToMonthString())
	s.Equal("十二月", MaxValue().SetLocale("zh-CN").ToMonthString())
}

func (s *ExtremumSuite) TestMinValue() {
	s.Equal("0001-01-01 00:00:00 +0000 UTC", MinValue().ToString())
	s.Equal("January", MinValue().ToMonthString())
	s.Equal("一月", MinValue().SetLocale("zh-CN").ToMonthString())
}

func (s *ExtremumSuite) TestMaxDuration() {
	s.Equal(9.223372036854776e+09, MaxDuration().Seconds())
}

func (s *ExtremumSuite) TestMinDuration() {
	s.Equal(-9.223372036854776e+09, MinDuration().Seconds())
}

func (s *ExtremumSuite) TestMax() {
	c1 := Parse("2020-08-01")
	c2 := Parse("2020-08-03")
	c3 := Parse("2020-08-06")

	s.Run("nil carbon", func() {
		var c *Carbon
		s.Nil(Max(c, c1, c2, c3))
		s.Nil(Max(c1, c, c2, c3))
		s.Nil(Max(c2, c1, c, c3))
		s.Nil(Max(c3, c1, c2, c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(c3, Max(c, c1, c2, c3))
		s.Equal(c3, Max(c1, c, c2, c3))
		s.Equal(c3, Max(c2, c1, c, c3))
		s.Equal(c3, Max(c3, c1, c2, c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Empty(Max(c, c1, c2, c3).ToString())
		s.Empty(Max(c1, c, c2, c3).ToString())
		s.Empty(Max(c2, c1, c, c3).ToString())
		s.Empty(Max(c3, c1, c2, c).ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Error(Max(c, c1, c2, c3).Error)
		s.Error(Max(c1, c, c2, c3).Error)
		s.Error(Max(c2, c1, c, c3).Error)
		s.Error(Max(c3, c1, c2, c).Error)
	})

	s.Run("valid carbon", func() {
		s.Equal(c1, Max(c1))
		s.Equal(c2, Max(c1, c2))
		s.Equal(c3, Max(c1, c2, c3))
	})
}

func (s *ExtremumSuite) TestMin() {
	c1 := Parse("2020-08-01")
	c2 := Parse("2020-08-03")
	c3 := Parse("2020-08-06")

	s.Run("nil carbon", func() {
		var c *Carbon
		s.Nil(Min(c, c1, c2, c3))
		s.Nil(Min(c1, c, c2, c3))
		s.Nil(Min(c2, c1, c, c3))
		s.Nil(Min(c3, c1, c2, c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(c, Min(c, c1, c2, c3))
		s.Equal(c, Min(c1, c, c2, c3))
		s.Equal(c, Min(c2, c1, c, c3))
		s.Equal(c, Min(c3, c1, c2, c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Empty(Min(c, c1, c2, c3).ToString())
		s.Empty(Min(c1, c, c2, c3).ToString())
		s.Empty(Min(c2, c1, c, c3).ToString())
		s.Empty(Min(c3, c1, c2, c).ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Error(Min(c, c1, c2, c3).Error)
		s.Error(Min(c1, c, c2, c3).Error)
		s.Error(Min(c2, c1, c, c3).Error)
		s.Error(Min(c3, c1, c2, c).Error)
	})

	s.Run("valid carbon", func() {
		s.Equal(c1, Min(c1))
		s.Equal(c1, Min(c1, c2))
		s.Equal(c1, Min(c1, c2, c3))
	})
}

func (s *ExtremumSuite) TestCarbon_Closest() {
	c1 := Parse("2020-08-01")
	c2 := Parse("2020-08-03")
	c3 := Parse("2020-08-06")

	s.Run("nil carbon", func() {
		var c *Carbon
		s.Equal(c, c.Closest(c1, c2, c3))
		s.Equal(c, c1.Closest(c, c2, c3))
		s.Equal(c, c2.Closest(c1, c, c3))
		s.Equal(c, c3.Closest(c1, c2, c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(c1, c.Closest(c1, c2, c3))
		s.Equal(c2, c1.Closest(c, c2, c3))
		s.Equal(c1, c2.Closest(c1, c, c3))
		s.Equal(c2, c3.Closest(c1, c2, c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Equal(c, c.Closest(c1, c2, c3))
		s.Equal(c, c1.Closest(c, c2, c3))
		s.Equal(c, c2.Closest(c1, c, c3))
		s.Equal(c, c3.Closest(c1, c2, c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Equal(c, c.Closest(c1, c2, c3))
		s.Equal(c, c1.Closest(c, c2, c3))
		s.Equal(c, c2.Closest(c1, c, c3))
		s.Equal(c, c3.Closest(c1, c2, c))
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05")
		s.Equal(c, c.Closest(c))
		s.Equal(c1, c.Closest(c1))
		s.Equal(c2, c.Closest(c1, c2))
		s.Equal(c3, c.Closest(c1, c3))
		s.Equal(c3, c.Closest(c1, c2, c3))
		s.Equal(c2, c1.Closest(c, c2, c3))
		s.Equal(c1, c2.Closest(c1, c, c3))
		s.Equal(c, c3.Closest(c1, c2, c))
	})
}

func (s *ExtremumSuite) TestCarbon_Farthest() {
	c1 := Parse("2020-08-01")
	c2 := Parse("2020-08-03")
	c3 := Parse("2020-08-06")

	s.Run("nil carbon", func() {
		var c *Carbon
		s.Equal(c, c.Farthest(c1, c2, c3))
		s.Equal(c, c1.Farthest(c, c2, c3))
		s.Equal(c, c2.Farthest(c1, c, c3))
		s.Equal(c, c3.Farthest(c1, c2, c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(c3, c.Farthest(c1, c2, c3))
		s.Equal(c, c1.Farthest(c, c2, c3))
		s.Equal(c, c2.Farthest(c1, c, c3))
		s.Equal(c, c3.Farthest(c1, c2, c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Equal(c, c.Farthest(c1, c2, c3))
		s.Equal(c, c1.Farthest(c, c2, c3))
		s.Equal(c, c2.Farthest(c1, c, c3))
		s.Equal(c, c3.Farthest(c1, c2, c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Equal(c, c.Farthest(c1, c2, c3))
		s.Equal(c, c1.Farthest(c, c2, c3))
		s.Equal(c, c2.Farthest(c1, c, c3))
		s.Equal(c, c3.Farthest(c1, c2, c))
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05")
		s.Equal(c, c.Farthest(c))
		s.Equal(c1, c.Farthest(c1))
		s.Equal(c1, c.Farthest(c1, c2))
		s.Equal(c1, c.Farthest(c1, c3))
		s.Equal(c1, c.Farthest(c1, c2, c3))
		s.Equal(c3, c1.Farthest(c, c2, c3))
		s.Equal(c3, c2.Farthest(c1, c, c3))
		s.Equal(c1, c3.Farthest(c1, c2, c))
	})
}
