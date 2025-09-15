package carbon

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SeasonSuite struct {
	suite.Suite
}

func TestSeasonSuite(t *testing.T) {
	suite.Run(t, new(SeasonSuite))
}

func (s *SeasonSuite) TestSeason() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.Season())
	})

	s.Run("zero carbon", func() {
		s.Equal(Winter, NewCarbon().Season())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").Season())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").Season())
	})

	s.Run("nil lang", func() {
		c := Now()
		c.lang = nil
		s.Empty(c.Season())
	})

	s.Run("error resources", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{
			"seasons": "xxx",
		})
		c := Now().SetLanguage(lang)
		s.Empty(c.Season())
	})

	s.Run("valid carbon", func() {
		s.Equal(Winter, Parse("2020-01-05").Season())
		s.Equal(Winter, Parse("2020-02-05").Season())
		s.Equal(Spring, Parse("2020-03-05").Season())
		s.Equal(Spring, Parse("2020-04-05").Season())
		s.Equal(Spring, Parse("2020-05-05").Season())
		s.Equal(Summer, Parse("2020-06-05").Season())
		s.Equal(Summer, Parse("2020-07-05").Season())
		s.Equal(Summer, Parse("2020-08-05").Season())
		s.Equal(Autumn, Parse("2020-09-05").Season())
		s.Equal(Autumn, Parse("2020-10-05").Season())
		s.Equal(Autumn, Parse("2020-11-05").Season())
		s.Equal(Winter, Parse("2020-12-05").Season())
	})
}

func (s *SeasonSuite) TestStartOfSeason() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.StartOfSeason())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-01 00:00:00 +0000 UTC", NewCarbon().StartOfSeason().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").StartOfSeason().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfSeason().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-12-01 00:00:00 +0000 UTC", Parse("2020-01-15").StartOfSeason().ToString())
		s.Equal("2019-12-01 00:00:00 +0000 UTC", Parse("2020-02-15").StartOfSeason().ToString())
		s.Equal("2020-03-01 00:00:00 +0000 UTC", Parse("2020-03-15").StartOfSeason().ToString())
		s.Equal("2020-03-01 00:00:00 +0000 UTC", Parse("2020-04-15").StartOfSeason().ToString())
		s.Equal("2020-03-01 00:00:00 +0000 UTC", Parse("2020-05-15").StartOfSeason().ToString())
		s.Equal("2020-06-01 00:00:00 +0000 UTC", Parse("2020-06-15").StartOfSeason().ToString())
		s.Equal("2020-06-01 00:00:00 +0000 UTC", Parse("2020-07-15").StartOfSeason().ToString())
		s.Equal("2020-06-01 00:00:00 +0000 UTC", Parse("2020-08-15").StartOfSeason().ToString())
		s.Equal("2020-09-01 00:00:00 +0000 UTC", Parse("2020-09-15").StartOfSeason().ToString())
		s.Equal("2020-09-01 00:00:00 +0000 UTC", Parse("2020-10-15").StartOfSeason().ToString())
		s.Equal("2020-09-01 00:00:00 +0000 UTC", Parse("2020-11-15").StartOfSeason().ToString())
		s.Equal("2020-12-01 00:00:00 +0000 UTC", Parse("2020-12-15").StartOfSeason().ToString())
	})
}

func (s *SeasonSuite) TestEndOfSeason() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.EndOfSeason())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-02-28 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfSeason().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").EndOfSeason().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfSeason().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-02-29 23:59:59.999999999 +0000 UTC", Parse("2020-01-15").EndOfSeason().ToString())
		s.Equal("2020-02-29 23:59:59.999999999 +0000 UTC", Parse("2020-02-15").EndOfSeason().ToString())
		s.Equal("2020-05-31 23:59:59.999999999 +0000 UTC", Parse("2020-03-15").EndOfSeason().ToString())
		s.Equal("2020-05-31 23:59:59.999999999 +0000 UTC", Parse("2020-04-15").EndOfSeason().ToString())
		s.Equal("2020-05-31 23:59:59.999999999 +0000 UTC", Parse("2020-05-15").EndOfSeason().ToString())
		s.Equal("2020-08-31 23:59:59.999999999 +0000 UTC", Parse("2020-06-15").EndOfSeason().ToString())
		s.Equal("2020-08-31 23:59:59.999999999 +0000 UTC", Parse("2020-07-15").EndOfSeason().ToString())
		s.Equal("2020-08-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15").EndOfSeason().ToString())
		s.Equal("2020-11-30 23:59:59.999999999 +0000 UTC", Parse("2020-09-15").EndOfSeason().ToString())
		s.Equal("2020-11-30 23:59:59.999999999 +0000 UTC", Parse("2020-10-15").EndOfSeason().ToString())
		s.Equal("2020-11-30 23:59:59.999999999 +0000 UTC", Parse("2020-11-15").EndOfSeason().ToString())
		s.Equal("2021-02-28 23:59:59.999999999 +0000 UTC", Parse("2020-12-15").EndOfSeason().ToString())
	})
}

func (s *SeasonSuite) TestIsSpring() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsSpring())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsSpring())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsSpring())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsSpring())
	})

	s.Run("valid carbon", func() {
		s.False(Parse("2020-01-01").IsSpring())
		s.True(Parse("2020-03-01").IsSpring())
	})
}

func (s *SeasonSuite) TestIsSummer() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsSummer())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsSummer())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsSummer())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsSummer())
	})

	s.Run("valid carbon", func() {
		s.False(Parse("2020-01-01").IsSummer())
		s.True(Parse("2020-06-01").IsSummer())
	})
}

func (s *SeasonSuite) TestIsAutumn() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsAutumn())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsAutumn())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsAutumn())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsAutumn())
	})

	s.Run("valid carbon", func() {
		s.False(Parse("2020-01-01").IsAutumn())
		s.True(Parse("2020-09-01").IsAutumn())
	})
}

func (s *SeasonSuite) TestIsWinter() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsWinter())
	})

	s.Run("zero carbon", func() {
		s.True(NewCarbon().IsWinter())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsWinter())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsWinter())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-01-01").IsWinter())
		s.False(Parse("2020-05-01").IsWinter())
	})
}
