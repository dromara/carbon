package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Constellation(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, Capricorn, NewCarbon().Constellation())
	})

	t.Run("invalid resources", func(t *testing.T) {
		lang := NewLanguage()
		resources := map[string]string{
			"constellations": "xxx",
		}
		lang.SetResources(resources)
		c := Parse("2020-01-05").SetLanguage(lang)
		assert.Empty(t, c.Constellation())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").Constellation())
		assert.Empty(t, Parse("0").Constellation())
		assert.Empty(t, Parse("xxx").Constellation())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, Capricorn, Parse("2020-01-05").Constellation())
		assert.Equal(t, Aquarius, Parse("2020-02-05").Constellation())
		assert.Equal(t, Pisces, Parse("2020-03-05").Constellation())
		assert.Equal(t, Aries, Parse("2020-04-05").Constellation())
		assert.Equal(t, Taurus, Parse("2020-05-05").Constellation())
		assert.Equal(t, Gemini, Parse("2020-06-05").Constellation())
		assert.Equal(t, Cancer, Parse("2020-07-05").Constellation())
		assert.Equal(t, Leo, Parse("2020-08-05").Constellation())
		assert.Equal(t, Virgo, Parse("2020-09-05").Constellation())
		assert.Equal(t, Libra, Parse("2020-10-05").Constellation())
		assert.Equal(t, Scorpio, Parse("2020-11-05").Constellation())
		assert.Equal(t, Sagittarius, Parse("2020-12-05").Constellation())

		assert.Equal(t, "摩羯座", Parse("2020-01-05").SetLocale("zh-CN").Constellation())
		assert.Equal(t, "水瓶座", Parse("2020-01-22").SetLocale("zh-CN").Constellation())
		assert.Equal(t, "水瓶座", Parse("2020-02-05").SetLocale("zh-CN").Constellation())
		assert.Equal(t, "双鱼座", Parse("2020-03-05").SetLocale("zh-CN").Constellation())
		assert.Equal(t, "白羊座", Parse("2020-04-05").SetLocale("zh-CN").Constellation())
		assert.Equal(t, "金牛座", Parse("2020-05-05").SetLocale("zh-CN").Constellation())
		assert.Equal(t, "双子座", Parse("2020-06-05").SetLocale("zh-CN").Constellation())
		assert.Equal(t, "巨蟹座", Parse("2020-07-05").SetLocale("zh-CN").Constellation())
		assert.Equal(t, "狮子座", Parse("2020-08-05").SetLocale("zh-CN").Constellation())
		assert.Equal(t, "处女座", Parse("2020-09-05").SetLocale("zh-CN").Constellation())
		assert.Equal(t, "天秤座", Parse("2020-10-05").SetLocale("zh-CN").Constellation())
		assert.Equal(t, "天蝎座", Parse("2020-11-05").SetLocale("zh-CN").Constellation())
		assert.Equal(t, "射手座", Parse("2020-12-05").SetLocale("zh-CN").Constellation())
	})
}

func TestCarbon_IsAries(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsAries())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsAries())
		assert.False(t, Parse("0").IsAries())
		assert.False(t, Parse("xxx").IsAries())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-03-21").IsAries())
		assert.True(t, Parse("2020-04-19").IsAries())
		assert.False(t, Parse("2020-08-05").IsAries())
	})
}

func TestCarbon_IsTaurus(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsTaurus())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsTaurus())
		assert.False(t, Parse("0").IsTaurus())
		assert.False(t, Parse("xxx").IsTaurus())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-04-20").IsTaurus())
		assert.True(t, Parse("2020-05-20").IsTaurus())
		assert.False(t, Parse("2020-08-05").IsTaurus())
	})
}

func TestCarbon_IsGemini(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsGemini())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsGemini())
		assert.False(t, Parse("0").IsGemini())
		assert.False(t, Parse("xxx").IsGemini())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-05-21").IsGemini())
		assert.True(t, Parse("2020-06-21").IsGemini())
		assert.False(t, Parse("2020-08-05").IsGemini())
	})
}

func TestCarbon_IsCancer(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsCancer())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsCancer())
		assert.False(t, Parse("0").IsCancer())
		assert.False(t, Parse("xxx").IsCancer())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-06-22").IsCancer())
		assert.True(t, Parse("2020-07-22").IsCancer())
		assert.False(t, Parse("2020-08-05").IsCancer())
	})
}

func TestCarbon_IsLeo(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsLeo())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsLeo())
		assert.False(t, Parse("0").IsLeo())
		assert.False(t, Parse("xxx").IsLeo())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-07-23").IsLeo())
		assert.True(t, Parse("2020-08-22").IsLeo())
		assert.False(t, Parse("2020-09-01").IsLeo())
	})
}

func TestCarbon_IsVirgo(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsVirgo())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsVirgo())
		assert.False(t, Parse("0").IsVirgo())
		assert.False(t, Parse("xxx").IsVirgo())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-23").IsVirgo())
		assert.True(t, Parse("2020-09-22").IsVirgo())
		assert.False(t, Parse("2020-08-05").IsVirgo())
	})
}

func TestCarbon_IsLibra(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsLibra())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsLibra())
		assert.False(t, Parse("0").IsLibra())
		assert.False(t, Parse("xxx").IsLibra())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-09-23").IsLibra())
		assert.True(t, Parse("2020-10-23").IsLibra())
		assert.False(t, Parse("2020-08-05").IsLibra())
	})
}

func TestCarbon_IsScorpio(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsScorpio())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsScorpio())
		assert.False(t, Parse("0").IsScorpio())
		assert.False(t, Parse("xxx").IsScorpio())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-10-24").IsScorpio())
		assert.True(t, Parse("2020-11-22").IsScorpio())
		assert.False(t, Parse("2020-08-05").IsScorpio())
	})
}

func TestCarbon_IsSagittarius(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsSagittarius())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsSagittarius())
		assert.False(t, Parse("0").IsSagittarius())
		assert.False(t, Parse("xxx").IsSagittarius())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-11-23").IsSagittarius())
		assert.True(t, Parse("2020-12-21").IsSagittarius())
		assert.False(t, Parse("2020-08-05").IsSagittarius())
	})
}

func TestCarbon_IsCapricorn(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsCapricorn())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsCapricorn())
		assert.False(t, Parse("0").IsCapricorn())
		assert.False(t, Parse("xxx").IsCapricorn())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-12-22").IsCapricorn())
		assert.True(t, Parse("2020-01-19").IsCapricorn())
		assert.False(t, Parse("2020-08-05").IsCapricorn())
	})
}

func TestCarbon_IsAquarius(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsAquarius())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsAquarius())
		assert.False(t, Parse("0").IsAquarius())
		assert.False(t, Parse("xxx").IsAquarius())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-01-20").IsAquarius())
		assert.True(t, Parse("2020-02-18").IsAquarius())
		assert.False(t, Parse("2020-08-05").IsAquarius())
	})
}

func TestCarbon_IsPisces(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsPisces())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsPisces())
		assert.False(t, Parse("0").IsPisces())
		assert.False(t, Parse("xxx").IsPisces())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-02-19").IsPisces())
		assert.True(t, Parse("2020-03-20").IsPisces())
		assert.False(t, Parse("2020-08-05").IsPisces())
	})
}
