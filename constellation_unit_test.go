package carbon

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConstellationSuite struct {
	suite.Suite
}

func TestConstellationSuite(t *testing.T) {
	suite.Run(t, new(ConstellationSuite))
}

func (s *ConstellationSuite) TestCarbon_Constellation() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.Constellation())
	})

	s.Run("zero carbon", func() {
		s.Equal(Capricorn, NewCarbon().Constellation())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").Constellation())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").Constellation())
	})

	s.Run("nil lang", func() {
		c := Now()
		c.lang = nil
		s.Empty(c.Constellation())
	})

	s.Run("valid carbon", func() {
		s.Equal(Capricorn, Parse("2020-01-05").Constellation())
		s.Equal(Aquarius, Parse("2020-02-05").Constellation())
		s.Equal(Pisces, Parse("2020-03-05").Constellation())
		s.Equal(Aries, Parse("2020-04-05").Constellation())
		s.Equal(Taurus, Parse("2020-05-05").Constellation())
		s.Equal(Gemini, Parse("2020-06-05").Constellation())
		s.Equal(Cancer, Parse("2020-07-05").Constellation())
		s.Equal(Leo, Parse("2020-08-05").Constellation())
		s.Equal(Virgo, Parse("2020-09-05").Constellation())
		s.Equal(Libra, Parse("2020-10-05").Constellation())
		s.Equal(Scorpio, Parse("2020-11-05").Constellation())
		s.Equal(Sagittarius, Parse("2020-12-05").Constellation())

		s.Equal("摩羯座", Parse("2020-01-05").SetLocale("zh-CN").Constellation())
		s.Equal("水瓶座", Parse("2020-01-22").SetLocale("zh-CN").Constellation())
		s.Equal("水瓶座", Parse("2020-02-05").SetLocale("zh-CN").Constellation())
		s.Equal("双鱼座", Parse("2020-03-05").SetLocale("zh-CN").Constellation())
		s.Equal("白羊座", Parse("2020-04-05").SetLocale("zh-CN").Constellation())
		s.Equal("金牛座", Parse("2020-05-05").SetLocale("zh-CN").Constellation())
		s.Equal("双子座", Parse("2020-06-05").SetLocale("zh-CN").Constellation())
		s.Equal("巨蟹座", Parse("2020-07-05").SetLocale("zh-CN").Constellation())
		s.Equal("狮子座", Parse("2020-08-05").SetLocale("zh-CN").Constellation())
		s.Equal("处女座", Parse("2020-09-05").SetLocale("zh-CN").Constellation())
		s.Equal("天秤座", Parse("2020-10-05").SetLocale("zh-CN").Constellation())
		s.Equal("天蝎座", Parse("2020-11-05").SetLocale("zh-CN").Constellation())
		s.Equal("射手座", Parse("2020-12-05").SetLocale("zh-CN").Constellation())
	})

	s.Run("cross-year constellations", func() {
		// Capricorn: 12/22-1/19
		s.Equal(Capricorn, Parse("2023-12-22").Constellation())
		s.Equal(Capricorn, Parse("2023-12-31").Constellation())
		s.Equal(Capricorn, Parse("2024-01-01").Constellation())
		s.Equal(Capricorn, Parse("2024-01-19").Constellation())
		s.Equal(Aquarius, Parse("2024-01-20").Constellation())

		// Aquarius: 1/20-2/18
		s.Equal(Aquarius, Parse("2024-01-20").Constellation())
		s.Equal(Aquarius, Parse("2024-02-10").Constellation())
		s.Equal(Aquarius, Parse("2024-02-18").Constellation())
		s.Equal(Pisces, Parse("2024-02-19").Constellation())

		// Pisces: 2/19-3/20
		s.Equal(Pisces, Parse("2024-02-19").Constellation())
		s.Equal(Pisces, Parse("2024-03-10").Constellation())
		s.Equal(Pisces, Parse("2024-03-20").Constellation())
		s.Equal(Aries, Parse("2024-03-21").Constellation())

		// Chinese cross-year constellation tests
		s.Equal("摩羯座", Parse("2023-12-22").SetLocale("zh-CN").Constellation())
		s.Equal("摩羯座", Parse("2023-12-31").SetLocale("zh-CN").Constellation())
		s.Equal("摩羯座", Parse("2024-01-01").SetLocale("zh-CN").Constellation())
		s.Equal("摩羯座", Parse("2024-01-19").SetLocale("zh-CN").Constellation())
		s.Equal("水瓶座", Parse("2024-01-20").SetLocale("zh-CN").Constellation())

		s.Equal("水瓶座", Parse("2024-01-20").SetLocale("zh-CN").Constellation())
		s.Equal("水瓶座", Parse("2024-02-10").SetLocale("zh-CN").Constellation())
		s.Equal("水瓶座", Parse("2024-02-18").SetLocale("zh-CN").Constellation())
		s.Equal("双鱼座", Parse("2024-02-19").SetLocale("zh-CN").Constellation())

		s.Equal("双鱼座", Parse("2024-02-19").SetLocale("zh-CN").Constellation())
		s.Equal("双鱼座", Parse("2024-03-10").SetLocale("zh-CN").Constellation())
		s.Equal("双鱼座", Parse("2024-03-20").SetLocale("zh-CN").Constellation())
		s.Equal("白羊座", Parse("2024-03-21").SetLocale("zh-CN").Constellation())
	})

	s.Run("empty resources", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{})
		s.Error(lang.Error)
		c := Parse("2020-01-05").SetLanguage(lang)
		s.Empty(c.Constellation())
	})

	s.Run("error resources", func() {
		lang1 := NewLanguage()
		lang1.SetResources(map[string]string{
			"xxx": "xxx",
		})
		s.Empty(Parse("2020-08-05").SetLanguage(lang1).Constellation())

		lang2 := NewLanguage()
		lang2.SetResources(map[string]string{
			"constellations": "xxx",
		})
		c := Parse("2020-01-05").SetLanguage(lang2)
		s.Empty(c.Constellation())
	})
}

func (s *ConstellationSuite) TestCarbon_IsAries() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsAries())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsAries())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsAries())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsAries())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-03-21").IsAries())
		s.True(Parse("2020-04-19").IsAries())
		s.False(Parse("2020-08-05").IsAries())
	})
}

func (s *ConstellationSuite) TestCarbon_IsTaurus() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsTaurus())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsTaurus())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsTaurus())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsTaurus())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-04-20").IsTaurus())
		s.True(Parse("2020-05-20").IsTaurus())
		s.False(Parse("2020-08-05").IsTaurus())
	})
}

func (s *ConstellationSuite) TestCarbon_IsGemini() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsGemini())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsGemini())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsGemini())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsGemini())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-05-21").IsGemini())
		s.True(Parse("2020-06-21").IsGemini())
		s.False(Parse("2020-08-05").IsGemini())
	})
}

func (s *ConstellationSuite) TestCarbon_IsCancer() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsCancer())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsCancer())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsCancer())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsCancer())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-06-22").IsCancer())
		s.True(Parse("2020-07-22").IsCancer())
		s.False(Parse("2020-08-05").IsCancer())
	})
}

func (s *ConstellationSuite) TestCarbon_IsLeo() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsLeo())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsLeo())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsLeo())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsLeo())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-07-23").IsLeo())
		s.True(Parse("2020-08-22").IsLeo())
		s.False(Parse("2020-09-01").IsLeo())
	})
}

func (s *ConstellationSuite) TestCarbon_IsVirgo() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsVirgo())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsVirgo())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsVirgo())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsVirgo())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-23").IsVirgo())
		s.True(Parse("2020-09-22").IsVirgo())
		s.False(Parse("2020-08-05").IsVirgo())
	})
}

func (s *ConstellationSuite) TestCarbon_IsLibra() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsLibra())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsLibra())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsLibra())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsLibra())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-09-23").IsLibra())
		s.True(Parse("2020-10-23").IsLibra())
		s.False(Parse("2020-08-05").IsLibra())
	})
}

func (s *ConstellationSuite) TestCarbon_IsScorpio() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsScorpio())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsScorpio())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsScorpio())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsScorpio())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-10-24").IsScorpio())
		s.True(Parse("2020-11-22").IsScorpio())
		s.False(Parse("2020-08-05").IsScorpio())
	})
}

func (s *ConstellationSuite) TestCarbon_IsSagittarius() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsSagittarius())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsSagittarius())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsSagittarius())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsSagittarius())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-11-23").IsSagittarius())
		s.True(Parse("2020-12-21").IsSagittarius())
		s.False(Parse("2020-08-05").IsSagittarius())
	})
}

func (s *ConstellationSuite) TestCarbon_IsCapricorn() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsCapricorn())
	})

	s.Run("zero carbon", func() {
		s.True(NewCarbon().IsCapricorn())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsCapricorn())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsCapricorn())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-12-22").IsCapricorn())
		s.True(Parse("2020-01-19").IsCapricorn())
		s.False(Parse("2020-08-05").IsCapricorn())
	})

	s.Run("cross-year capricorn", func() {
		// Capricorn: 12/22-1/19
		s.True(Parse("2023-12-22").IsCapricorn())
		s.True(Parse("2023-12-31").IsCapricorn())
		s.True(Parse("2024-01-01").IsCapricorn())
		s.True(Parse("2024-01-19").IsCapricorn())
		s.False(Parse("2024-01-20").IsCapricorn())
		s.False(Parse("2024-08-05").IsCapricorn())
	})
}

func (s *ConstellationSuite) TestCarbon_IsAquarius() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsAquarius())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsAquarius())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsAquarius())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsAquarius())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-01-20").IsAquarius())
		s.True(Parse("2020-02-18").IsAquarius())
		s.False(Parse("2020-08-05").IsAquarius())
	})

	s.Run("cross-year aquarius", func() {
		// Aquarius: 1/20-2/18
		s.True(Parse("2024-01-20").IsAquarius())
		s.True(Parse("2024-02-10").IsAquarius())
		s.True(Parse("2024-02-18").IsAquarius())
		s.False(Parse("2024-02-19").IsAquarius())
		s.False(Parse("2024-08-05").IsAquarius())
	})
}

func (s *ConstellationSuite) TestCarbon_IsPisces() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsPisces())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsPisces())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsPisces())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsPisces())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-02-19").IsPisces())
		s.True(Parse("2020-03-20").IsPisces())
		s.False(Parse("2020-08-05").IsPisces())
	})

	s.Run("cross-year pisces", func() {
		// Pisces: 2/19-3/20
		s.True(Parse("2024-02-19").IsPisces())
		s.True(Parse("2024-03-10").IsPisces())
		s.True(Parse("2024-03-20").IsPisces())
		s.False(Parse("2024-03-21").IsPisces())
		s.False(Parse("2024-08-05").IsPisces())
	})
}
