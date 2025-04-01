package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Season(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, Winter, NewCarbon().Season())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").Season())
		assert.Empty(t, Parse("0").Season())
		assert.Empty(t, Parse("xxx").Season())
	})

	t.Run("invalid resources", func(t *testing.T) {
		lang := NewLanguage()
		resources := map[string]string{
			"seasons": "xxx",
		}
		lang.SetResources(resources)
		c := Parse("2020-01-05").SetLanguage(lang)
		assert.Empty(t, c.Season())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, Winter, Parse("2020-01-05").Season())
		assert.Equal(t, Winter, Parse("2020-02-05").Season())
		assert.Equal(t, Spring, Parse("2020-03-05").Season())
		assert.Equal(t, Spring, Parse("2020-04-05").Season())
		assert.Equal(t, Spring, Parse("2020-05-05").Season())
		assert.Equal(t, Summer, Parse("2020-06-05").Season())
		assert.Equal(t, Summer, Parse("2020-07-05").Season())
		assert.Equal(t, Summer, Parse("2020-08-05").Season())
		assert.Equal(t, Autumn, Parse("2020-09-05").Season())
		assert.Equal(t, Autumn, Parse("2020-10-05").Season())
		assert.Equal(t, Autumn, Parse("2020-11-05").Season())
		assert.Equal(t, Winter, Parse("2020-12-05").Season())
	})
}

func TestCarbon_StartOfSeason(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-01 00:00:00 +0000 UTC", NewCarbon().StartOfSeason().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").StartOfSeason().ToString())
		assert.Empty(t, Parse("0").StartOfSeason().ToString())
		assert.Empty(t, Parse("xxx").StartOfSeason().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2019-12-01 00:00:00 +0000 UTC", Parse("2020-01-15").StartOfSeason().ToString())
		assert.Equal(t, "2019-12-01 00:00:00 +0000 UTC", Parse("2020-02-15").StartOfSeason().ToString())
		assert.Equal(t, "2020-03-01 00:00:00 +0000 UTC", Parse("2020-03-15").StartOfSeason().ToString())
		assert.Equal(t, "2020-03-01 00:00:00 +0000 UTC", Parse("2020-04-15").StartOfSeason().ToString())
		assert.Equal(t, "2020-03-01 00:00:00 +0000 UTC", Parse("2020-05-15").StartOfSeason().ToString())
		assert.Equal(t, "2020-06-01 00:00:00 +0000 UTC", Parse("2020-06-15").StartOfSeason().ToString())
		assert.Equal(t, "2020-06-01 00:00:00 +0000 UTC", Parse("2020-07-15").StartOfSeason().ToString())
		assert.Equal(t, "2020-06-01 00:00:00 +0000 UTC", Parse("2020-08-15").StartOfSeason().ToString())
		assert.Equal(t, "2020-09-01 00:00:00 +0000 UTC", Parse("2020-09-15").StartOfSeason().ToString())
		assert.Equal(t, "2020-09-01 00:00:00 +0000 UTC", Parse("2020-10-15").StartOfSeason().ToString())
		assert.Equal(t, "2020-09-01 00:00:00 +0000 UTC", Parse("2020-11-15").StartOfSeason().ToString())
		assert.Equal(t, "2020-12-01 00:00:00 +0000 UTC", Parse("2020-12-15").StartOfSeason().ToString())
	})
}

func TestCarbon_EndOfSeason(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-02-28 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfSeason().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").EndOfSeason().ToString())
		assert.Empty(t, Parse("0").EndOfSeason().ToString())
		assert.Empty(t, Parse("xxx").EndOfSeason().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-02-29 23:59:59.999999999 +0000 UTC", Parse("2020-01-15").EndOfSeason().ToString())
		assert.Equal(t, "2020-02-29 23:59:59.999999999 +0000 UTC", Parse("2020-02-15").EndOfSeason().ToString())
		assert.Equal(t, "2020-05-31 23:59:59.999999999 +0000 UTC", Parse("2020-03-15").EndOfSeason().ToString())
		assert.Equal(t, "2020-05-31 23:59:59.999999999 +0000 UTC", Parse("2020-04-15").EndOfSeason().ToString())
		assert.Equal(t, "2020-05-31 23:59:59.999999999 +0000 UTC", Parse("2020-05-15").EndOfSeason().ToString())
		assert.Equal(t, "2020-08-31 23:59:59.999999999 +0000 UTC", Parse("2020-06-15").EndOfSeason().ToString())
		assert.Equal(t, "2020-08-31 23:59:59.999999999 +0000 UTC", Parse("2020-07-15").EndOfSeason().ToString())
		assert.Equal(t, "2020-08-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15").EndOfSeason().ToString())
		assert.Equal(t, "2020-11-30 23:59:59.999999999 +0000 UTC", Parse("2020-09-15").EndOfSeason().ToString())
		assert.Equal(t, "2020-11-30 23:59:59.999999999 +0000 UTC", Parse("2020-10-15").EndOfSeason().ToString())
		assert.Equal(t, "2020-11-30 23:59:59.999999999 +0000 UTC", Parse("2020-11-15").EndOfSeason().ToString())
		assert.Equal(t, "2021-02-28 23:59:59.999999999 +0000 UTC", Parse("2020-12-15").EndOfSeason().ToString())
	})
}

func TestCarbon_IsSpring(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsSpring())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsSpring())
		assert.False(t, Parse("0").IsSpring())
		assert.False(t, Parse("xxx").IsSpring())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Parse("2020-01-01").IsSpring())
		assert.True(t, Parse("2020-03-01").IsSpring())
	})
}

func TestCarbon_IsSummer(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsSummer())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsSummer())
		assert.False(t, Parse("0").IsSummer())
		assert.False(t, Parse("xxx").IsSummer())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Parse("2020-01-01").IsSummer())
		assert.True(t, Parse("2020-06-01").IsSummer())
	})
}

func TestCarbon_IsAutumn(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsAutumn())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsAutumn())
		assert.False(t, Parse("0").IsAutumn())
		assert.False(t, Parse("xxx").IsAutumn())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Parse("2020-01-01").IsAutumn())
		assert.True(t, Parse("2020-09-01").IsAutumn())
	})
}

func TestCarbon_IsWinter(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsWinter())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsWinter())
		assert.False(t, Parse("0").IsWinter())
		assert.False(t, Parse("xxx").IsWinter())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-01-01").IsWinter())
		assert.False(t, Parse("2020-05-01").IsWinter())
	})
}
