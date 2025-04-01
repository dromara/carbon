package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetLayout(t *testing.T) {
	defer SetLayout(DateTimeLayout)

	t.Run("zero time", func(t *testing.T) {
		SetLayout(DateLayout)
		assert.Equal(t, DateLayout, DefaultLayout)
		assert.Equal(t, DateLayout, NewCarbon().CurrentLayout())

		SetLayout(DateTimeLayout)
		assert.Equal(t, DateTimeLayout, DefaultLayout)
		assert.Equal(t, DateTimeLayout, NewCarbon().CurrentLayout())
	})

	t.Run("valid time", func(t *testing.T) {
		SetLayout(DateLayout)
		c1 := Parse("2020-08-05")
		assert.Equal(t, DateLayout, DefaultLayout)
		assert.Equal(t, DateLayout, c1.CurrentLayout())

		SetLayout(DateTimeLayout)
		c2 := Parse("2020-08-05 13:14:15")
		assert.Equal(t, DateTimeLayout, DefaultLayout)
		assert.Equal(t, DateTimeLayout, c2.CurrentLayout())
	})
}

func TestCarbon_SetLayout(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().SetLayout(DateLayout)
		assert.Equal(t, DateLayout, c.CurrentLayout())
	})

	t.Run("empty layout", func(t *testing.T) {
		assert.True(t, Now().SetLayout("").HasError())
		assert.Empty(t, Now().SetLayout("").CurrentLayout())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetLayout(DateLayout).CurrentLayout())
		assert.Empty(t, Parse("0").SetLayout(DateLayout).CurrentLayout())
		assert.Empty(t, Parse("xxx").SetLayout(DateLayout).CurrentLayout())
	})

	t.Run("valid time", func(t *testing.T) {
		c1 := Parse("2020-08-05").SetLayout(DateLayout)
		assert.Equal(t, DateLayout, c1.CurrentLayout())

		c2 := Parse("2020-08-05").SetLayout(DateTimeLayout)
		assert.Equal(t, DateTimeLayout, c2.CurrentLayout())
	})
}

func TestSetFormat(t *testing.T) {
	defer SetFormat(DateTimeFormat)

	t.Run("zero time", func(t *testing.T) {
		SetFormat(DateFormat)
		assert.Equal(t, DateLayout, DefaultLayout)
		assert.Equal(t, DateLayout, NewCarbon().CurrentLayout())

		SetFormat(DateTimeFormat)
		assert.Equal(t, DateTimeLayout, DefaultLayout)
		assert.Equal(t, DateTimeLayout, NewCarbon().CurrentLayout())
	})

	t.Run("valid time", func(t *testing.T) {
		SetFormat(DateFormat)
		c1 := Parse("2020-08-05")
		assert.Equal(t, DateLayout, DefaultLayout)
		assert.Equal(t, DateLayout, c1.CurrentLayout())

		SetFormat(DateTimeFormat)
		c2 := Parse("2020-08-05 13:14:15")
		assert.Equal(t, DateTimeLayout, DefaultLayout)
		assert.Equal(t, DateTimeLayout, c2.CurrentLayout())
	})
}

func TestCarbon_SetFormat(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().SetFormat(DateFormat)
		assert.Equal(t, DateLayout, c.CurrentLayout())
	})

	t.Run("empty layout", func(t *testing.T) {
		assert.True(t, Now().SetFormat("").HasError())
		assert.Empty(t, Now().SetFormat("").CurrentLayout())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetFormat(DateFormat).CurrentLayout())
		assert.Empty(t, Parse("0").SetFormat(DateFormat).CurrentLayout())
		assert.Empty(t, Parse("xxx").SetFormat(DateFormat).CurrentLayout())
	})

	t.Run("valid time", func(t *testing.T) {
		c1 := Parse("2020-08-05").SetFormat(DateFormat)
		assert.Equal(t, DateLayout, c1.CurrentLayout())

		c2 := Parse("2020-08-05").SetFormat(DateTimeFormat)
		assert.Equal(t, DateTimeLayout, c2.CurrentLayout())
	})
}

func TestSetWeekStartsAt(t *testing.T) {
	defer SetWeekStartsAt(Sunday)

	t.Run("zero time", func(t *testing.T) {
		SetWeekStartsAt(Sunday)
		assert.Equal(t, Sunday, DefaultWeekStartsAt)
		assert.Equal(t, Sunday, NewCarbon().WeekStartsAt())
		assert.Equal(t, "0000-12-31 00:00:00 +0000 UTC", NewCarbon().StartOfWeek().ToString())

		SetWeekStartsAt(Monday)
		assert.Equal(t, Monday, DefaultWeekStartsAt)
		assert.Equal(t, Monday, NewCarbon().WeekStartsAt())
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfWeek().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		SetWeekStartsAt(Monday)
		c1 := Parse("2020-08-05")
		assert.Equal(t, Monday, DefaultWeekStartsAt)
		assert.Equal(t, Monday, c1.WeekStartsAt())
		assert.Equal(t, "2020-08-03 00:00:00 +0000 UTC", c1.StartOfWeek().ToString())

		SetWeekStartsAt(Sunday)
		c2 := Parse("2020-08-05")
		assert.Equal(t, Sunday, DefaultWeekStartsAt)
		assert.Equal(t, Sunday, c2.WeekStartsAt())
		assert.Equal(t, "2020-08-02 00:00:00 +0000 UTC", c2.StartOfWeek().ToString())
	})
}

func TestCarbon_SetWeekStartsAt(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c1 := NewCarbon().SetWeekStartsAt(Sunday)
		assert.Equal(t, "0000-12-31 00:00:00 +0000 UTC", c1.StartOfWeek().ToString())

		c2 := NewCarbon().SetWeekStartsAt(Monday)
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", c2.StartOfWeek().ToString())
	})

	t.Run("invalid day", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05").SetWeekStartsAt("").HasError())
		assert.True(t, Parse("2020-08-05").SetWeekStartsAt("0").HasError())
		assert.True(t, Parse("2020-08-05").SetWeekStartsAt("xxx").HasError())

		assert.Empty(t, Parse("2020-08-05").SetWeekStartsAt("").ToString())
		assert.Empty(t, Parse("2020-08-05").SetWeekStartsAt("0").ToString())
		assert.Empty(t, Parse("2020-08-05").SetWeekStartsAt("xxx").ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetWeekStartsAt(Sunday).ToString())
		assert.Empty(t, Parse("xxx").SetWeekStartsAt(Sunday).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		c1 := Parse("2020-08-05").SetWeekStartsAt(Monday)
		assert.Equal(t, Monday, c1.WeekStartsAt())
		assert.Equal(t, "2020-08-03 00:00:00 +0000 UTC", c1.StartOfWeek().ToString())

		c2 := Parse("2020-08-05").SetWeekStartsAt(Sunday)
		assert.Equal(t, Sunday, c2.WeekStartsAt())
		assert.Equal(t, "2020-08-02 00:00:00 +0000 UTC", c2.StartOfWeek().ToString())
	})
}

func TestSetLocale(t *testing.T) {
	defer SetLocale("en")

	t.Run("zero time", func(t *testing.T) {
		SetLocale("zh-CN")

		assert.Equal(t, "zh-CN", DefaultLocale)
		assert.Equal(t, "zh-CN", NewCarbon().Locale())
		assert.Equal(t, "摩羯座", NewCarbon().Constellation())
		assert.Equal(t, "冬季", NewCarbon().Season())
		assert.Equal(t, "一月", NewCarbon().ToMonthString())
		assert.Equal(t, "1月", NewCarbon().ToShortMonthString())
		assert.Equal(t, "星期一", NewCarbon().ToWeekString())
		assert.Equal(t, "周一", NewCarbon().ToShortWeekString())

		SetLocale("en")

		assert.Equal(t, "en", DefaultLocale)
		assert.Equal(t, "en", NewCarbon().Locale())
		assert.Equal(t, "Capricorn", NewCarbon().Constellation())
		assert.Equal(t, "Winter", NewCarbon().Season())
		assert.Equal(t, "January", NewCarbon().ToMonthString())
		assert.Equal(t, "Jan", NewCarbon().ToShortMonthString())
		assert.Equal(t, "Monday", NewCarbon().ToWeekString())
		assert.Equal(t, "Mon", NewCarbon().ToShortWeekString())
	})

	t.Run("valid time", func(t *testing.T) {
		SetLocale("zh-CN")
		c1 := Parse("2020-08-05")
		assert.Equal(t, "zh-CN", DefaultLocale)
		assert.Equal(t, "zh-CN", c1.Locale())
		assert.Equal(t, "狮子座", c1.Constellation())
		assert.Equal(t, "夏季", c1.Season())
		assert.Equal(t, "八月", c1.ToMonthString())
		assert.Equal(t, "8月", c1.ToShortMonthString())
		assert.Equal(t, "星期三", c1.ToWeekString())
		assert.Equal(t, "周三", c1.ToShortWeekString())

		SetLocale("en")
		c2 := Parse("2020-08-05")
		assert.Equal(t, "en", DefaultLocale)
		assert.Equal(t, "en", c2.Locale())
		assert.Equal(t, "Leo", c2.Constellation())
		assert.Equal(t, "Summer", c2.Season())
		assert.Equal(t, "August", c2.ToMonthString())
		assert.Equal(t, "Aug", c2.ToShortMonthString())
		assert.Equal(t, "Wednesday", c2.ToWeekString())
		assert.Equal(t, "Wed", c2.ToShortWeekString())
	})
}

func TestCarbon_SetLocale(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c1 := NewCarbon().SetLocale("zh-CN")
		assert.Equal(t, "zh-CN", c1.Locale())
		assert.Equal(t, "摩羯座", c1.Constellation())
		assert.Equal(t, "冬季", c1.Season())
		assert.Equal(t, "一月", c1.ToMonthString())
		assert.Equal(t, "1月", c1.ToShortMonthString())
		assert.Equal(t, "星期一", c1.ToWeekString())
		assert.Equal(t, "周一", c1.ToShortWeekString())

		c2 := NewCarbon().SetLocale("en")
		assert.Equal(t, "Capricorn", c2.Constellation())
		assert.Equal(t, "Winter", c2.Season())
		assert.Equal(t, "January", c2.ToMonthString())
		assert.Equal(t, "Jan", c2.ToShortMonthString())
		assert.Equal(t, "Monday", c2.ToWeekString())
		assert.Equal(t, "Mon", c2.ToShortWeekString())
	})

	t.Run("invalid locale", func(t *testing.T) {
		assert.True(t, Now().SetLocale("").HasError())
		assert.True(t, Now().SetLocale("xxx").HasError())

		assert.Empty(t, Now().SetLocale("").ToString())
		assert.Empty(t, Now().SetLocale("xxx").ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetLocale("en").ToString())
		assert.Empty(t, Parse("0").SetLocale("en").ToString())
		assert.Empty(t, Parse("xxx").SetLocale("en").ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		c1 := Parse("2020-08-05").SetLocale("zh-CN")
		assert.Equal(t, "zh-CN", c1.Locale())
		assert.Equal(t, "狮子座", c1.Constellation())
		assert.Equal(t, "夏季", c1.Season())
		assert.Equal(t, "八月", c1.ToMonthString())
		assert.Equal(t, "8月", c1.ToShortMonthString())
		assert.Equal(t, "星期三", c1.ToWeekString())
		assert.Equal(t, "周三", c1.ToShortWeekString())

		c2 := Parse("2020-08-05").SetLocale("en")
		assert.Equal(t, "en", c2.Locale())
		assert.Equal(t, "Leo", c2.Constellation())
		assert.Equal(t, "Summer", c2.Season())
		assert.Equal(t, "August", c2.ToMonthString())
		assert.Equal(t, "Aug", c2.ToShortMonthString())
		assert.Equal(t, "Wednesday", c2.ToWeekString())
		assert.Equal(t, "Wed", c2.ToShortWeekString())
	})
}

func TestSetTimezone(t *testing.T) {
	defer SetTimezone(UTC)

	t.Run("zero time", func(t *testing.T) {
		SetTimezone(UTC)
		assert.Equal(t, UTC, DefaultTimezone)
		assert.Equal(t, UTC, NewCarbon().Timezone())
		assert.Equal(t, UTC, NewCarbon().ZoneName())
		assert.Equal(t, 0, NewCarbon().ZoneOffset())
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", NewCarbon().ToString())

		SetTimezone(PRC)
		assert.Equal(t, PRC, DefaultTimezone)
		assert.Equal(t, PRC, NewCarbon().Timezone())
		assert.Equal(t, "LMT", NewCarbon().ZoneName())
		assert.Equal(t, 29143, NewCarbon().ZoneOffset())
		assert.Equal(t, "0001-01-01 08:05:43 +0805 LMT", NewCarbon().ToString())

		SetTimezone(Japan)
		assert.Equal(t, Japan, DefaultTimezone)
		assert.Equal(t, Japan, NewCarbon().Timezone())
		assert.Equal(t, "LMT", NewCarbon().ZoneName())
		assert.Equal(t, 33539, NewCarbon().ZoneOffset())
		assert.Equal(t, "0001-01-01 09:18:59 +0918 LMT", NewCarbon().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		SetTimezone(UTC)
		c1 := Parse("2020-08-05 13:14:15")
		assert.Equal(t, UTC, DefaultTimezone)
		assert.Equal(t, UTC, c1.Timezone())
		assert.Equal(t, UTC, c1.ZoneName())
		assert.Equal(t, 0, c1.ZoneOffset())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", c1.ToString())

		SetTimezone(PRC)
		c2 := Parse("2020-08-05 13:14:15")
		assert.Equal(t, PRC, DefaultTimezone)
		assert.Equal(t, PRC, c2.Timezone())
		assert.Equal(t, "CST", c2.ZoneName())
		assert.Equal(t, 28800, c2.ZoneOffset())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", c2.ToString())

		SetTimezone(Japan)
		c3 := Parse("2020-08-05 13:14:15")
		assert.Equal(t, Japan, DefaultTimezone)
		assert.Equal(t, Japan, c3.Timezone())
		assert.Equal(t, "JST", c3.ZoneName())
		assert.Equal(t, 32400, c3.ZoneOffset())
		assert.Equal(t, "2020-08-05 13:14:15 +0900 JST", c3.ToString())
	})
}

func TestCarbon_SetTimezone(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		l1, _ := time.LoadLocation(UTC)
		t1 := time.Time{}.In(l1)
		n1, f1 := t1.Zone()
		assert.Equal(t, UTC, t1.Location().String())
		assert.Equal(t, UTC, n1)
		assert.Equal(t, 0, f1)
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", t1.String())

		c1 := NewCarbon().SetTimezone(UTC)
		assert.Equal(t, UTC, c1.Timezone())
		assert.Equal(t, UTC, c1.ZoneName())
		assert.Equal(t, 0, c1.ZoneOffset())
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", c1.ToString())

		l2, _ := time.LoadLocation(PRC)
		t2 := time.Time{}.In(l2)
		n2, f2 := t2.Zone()
		assert.Equal(t, PRC, t2.Location().String())
		assert.Equal(t, "LMT", n2)
		assert.Equal(t, 29143, f2)
		assert.Equal(t, "0001-01-01 08:05:43 +0805 LMT", t2.String())

		c2 := NewCarbon().SetTimezone(PRC)
		assert.Equal(t, PRC, c2.Timezone())
		assert.Equal(t, "LMT", c2.ZoneName())
		assert.Equal(t, 29143, c2.ZoneOffset())
		assert.Equal(t, "0001-01-01 08:05:43 +0805 LMT", c2.ToString())

		l3, _ := time.LoadLocation(Japan)
		t3 := time.Time{}.In(l3)
		n3, f3 := t3.Zone()
		assert.Equal(t, "LMT", n3)
		assert.Equal(t, 33539, f3)
		assert.Equal(t, "0001-01-01 09:18:59 +0918 LMT", t3.String())

		c3 := NewCarbon().SetTimezone(Japan)
		assert.Equal(t, Japan, c3.Timezone())
		assert.Equal(t, "LMT", c3.ZoneName())
		assert.Equal(t, 33539, c3.ZoneOffset())
		assert.Equal(t, "0001-01-01 09:18:59 +0918 LMT", c3.ToString())
	})

	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05").SetTimezone("").HasError())
		assert.True(t, Parse("2020-08-05").SetTimezone("0").HasError())
		assert.True(t, Parse("2020-08-05").SetTimezone("XXX").HasError())

		assert.Empty(t, Parse("2020-08-05").SetTimezone("").ToString())
		assert.Empty(t, Parse("2020-08-05").SetTimezone("0").ToString())
		assert.Empty(t, Parse("2020-08-05").SetTimezone("XXX").ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetTimezone(UTC).ToString())
		assert.Empty(t, Parse("0").SetTimezone(PRC).ToString())
		assert.Empty(t, Parse("xxx").SetTimezone(PRC).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		c := Parse("2020-08-05")

		c.SetTimezone(UTC)
		assert.Equal(t, UTC, c.Timezone())
		assert.Equal(t, UTC, c.ZoneName())
		assert.Equal(t, 0, c.ZoneOffset())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", c.ToString())

		c = c.SetTimezone(PRC)
		assert.Equal(t, PRC, c.Timezone())
		assert.Equal(t, "CST", c.ZoneName())
		assert.Equal(t, 28800, c.ZoneOffset())
		assert.Equal(t, "2020-08-05 08:00:00 +0800 CST", c.ToString())

		c = c.SetTimezone(Japan)
		assert.Equal(t, Japan, c.Timezone())
		assert.Equal(t, "JST", c.ZoneName())
		assert.Equal(t, 32400, c.ZoneOffset())
		assert.Equal(t, "2020-08-05 09:00:00 +0900 JST", c.ToString())
	})
}

func TestSetLocation(t *testing.T) {
	defer SetLocation(time.UTC)

	t.Run("zero time", func(t *testing.T) {
		l1, _ := time.LoadLocation(UTC)
		SetLocation(l1)
		assert.Equal(t, UTC, DefaultTimezone)
		assert.Equal(t, UTC, NewCarbon().Timezone())
		assert.Equal(t, UTC, NewCarbon().ZoneName())
		assert.Equal(t, 0, NewCarbon().ZoneOffset())
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", NewCarbon().ToString())

		l2, _ := time.LoadLocation(PRC)
		SetLocation(l2)
		assert.Equal(t, PRC, DefaultTimezone)
		assert.Equal(t, PRC, NewCarbon().Timezone())
		assert.Equal(t, "LMT", NewCarbon().ZoneName())
		assert.Equal(t, 29143, NewCarbon().ZoneOffset())
		assert.Equal(t, "0001-01-01 08:05:43 +0805 LMT", NewCarbon().ToString())

		l3, _ := time.LoadLocation(Japan)
		SetLocation(l3)
		assert.Equal(t, Japan, DefaultTimezone)
		assert.Equal(t, Japan, NewCarbon().Timezone())
		assert.Equal(t, "LMT", NewCarbon().ZoneName())
		assert.Equal(t, 33539, NewCarbon().ZoneOffset())
		assert.Equal(t, "0001-01-01 09:18:59 +0918 LMT", NewCarbon().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		l1, _ := time.LoadLocation(UTC)
		SetLocation(l1)
		c1 := Parse("2020-08-05 13:14:15")
		assert.Equal(t, UTC, DefaultTimezone)
		assert.Equal(t, UTC, c1.Timezone())
		assert.Equal(t, UTC, c1.ZoneName())
		assert.Equal(t, 0, c1.ZoneOffset())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", c1.ToString())

		l2, _ := time.LoadLocation(PRC)
		SetLocation(l2)
		c2 := Parse("2020-08-05 13:14:15")
		assert.Equal(t, PRC, DefaultTimezone)
		assert.Equal(t, PRC, c2.Timezone())
		assert.Equal(t, "CST", c2.ZoneName())
		assert.Equal(t, 28800, c2.ZoneOffset())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", c2.ToString())

		l3, _ := time.LoadLocation(Japan)
		SetLocation(l3)
		c3 := Parse("2020-08-05 13:14:15")
		assert.Equal(t, Japan, DefaultTimezone)
		assert.Equal(t, Japan, c3.Timezone())
		assert.Equal(t, "JST", c3.ZoneName())
		assert.Equal(t, 32400, c3.ZoneOffset())
		assert.Equal(t, "2020-08-05 13:14:15 +0900 JST", c3.ToString())
	})
}

func TestCarbon_SetLocation(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		l1, _ := time.LoadLocation(UTC)
		t1 := time.Time{}.In(l1)
		n1, f1 := t1.Zone()
		assert.Equal(t, UTC, t1.Location().String())
		assert.Equal(t, UTC, n1)
		assert.Equal(t, 0, f1)
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", t1.String())

		c1 := NewCarbon().SetLocation(l1)
		assert.Equal(t, UTC, c1.Timezone())
		assert.Equal(t, UTC, c1.ZoneName())
		assert.Equal(t, 0, c1.ZoneOffset())
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", c1.ToString())

		l2, _ := time.LoadLocation(PRC)
		t2 := time.Time{}.In(l2)
		n2, f2 := t2.Zone()
		assert.Equal(t, PRC, t2.Location().String())
		assert.Equal(t, "LMT", n2)
		assert.Equal(t, 29143, f2)
		assert.Equal(t, "0001-01-01 08:05:43 +0805 LMT", t2.String())

		c2 := NewCarbon().SetLocation(l2)
		assert.Equal(t, PRC, c2.Timezone())
		assert.Equal(t, "LMT", c2.ZoneName())
		assert.Equal(t, 29143, c2.ZoneOffset())
		assert.Equal(t, "0001-01-01 08:05:43 +0805 LMT", c2.ToString())

		l3, _ := time.LoadLocation(Japan)
		t3 := time.Time{}.In(l3)
		n3, f3 := t3.Zone()
		assert.Equal(t, "LMT", n3)
		assert.Equal(t, 33539, f3)
		assert.Equal(t, "0001-01-01 09:18:59 +0918 LMT", t3.String())

		c3 := NewCarbon().SetLocation(l3)
		assert.Equal(t, Japan, c3.Timezone())
		assert.Equal(t, "LMT", c3.ZoneName())
		assert.Equal(t, 33539, c3.ZoneOffset())
		assert.Equal(t, "0001-01-01 09:18:59 +0918 LMT", c3.ToString())
	})

	t.Run("nil location", func(t *testing.T) {
		assert.Empty(t, Parse("2020-08-05").SetLocation(nil).Timezone())
		assert.Empty(t, Parse("2020-08-05").SetLocation(nil).ZoneName())
		assert.Empty(t, Parse("2020-08-05").SetLocation(nil).ZoneOffset())
		assert.Empty(t, Parse("2020-08-05").SetLocation(nil).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetLocation(time.UTC).ToString())
		assert.Empty(t, Parse("0").SetLocation(time.UTC).ToString())
		assert.Empty(t, Parse("xxx").SetLocation(time.UTC).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		c := Parse("2020-08-05")

		l1, _ := time.LoadLocation(UTC)
		c.SetLocation(l1)
		assert.Equal(t, UTC, c.Timezone())
		assert.Equal(t, UTC, c.ZoneName())
		assert.Equal(t, 0, c.ZoneOffset())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", c.ToString())

		l2, _ := time.LoadLocation(PRC)
		c = c.SetLocation(l2)
		assert.Equal(t, PRC, c.Timezone())
		assert.Equal(t, "CST", c.ZoneName())
		assert.Equal(t, 28800, c.ZoneOffset())
		assert.Equal(t, "2020-08-05 08:00:00 +0800 CST", c.ToString())

		l3, _ := time.LoadLocation(Japan)
		c = c.SetLocation(l3)
		assert.Equal(t, Japan, c.Timezone())
		assert.Equal(t, "JST", c.ZoneName())
		assert.Equal(t, 32400, c.ZoneOffset())
		assert.Equal(t, "2020-08-05 09:00:00 +0900 JST", c.ToString())
	})
}

func TestCarbon_SetLanguage(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetLocale("en")
		c := NewCarbon().SetLanguage(lang)
		assert.Equal(t, "en", c.Locale())
	})

	t.Run("nil language", func(t *testing.T) {
		lang := NewLanguage()
		lang = nil
		assert.Empty(t, NewCarbon().SetLanguage(lang).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetLocale("en")
		assert.Empty(t, Parse("").SetLanguage(lang).ToString())
		assert.Empty(t, Parse("0").SetLanguage(lang).ToString())
		assert.Empty(t, Parse("xxx").SetLanguage(lang).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		lang := NewLanguage()

		lang.SetLocale("en")
		assert.Equal(t, "August", Parse("2020-08-05").SetLanguage(lang).ToMonthString())

		lang.SetLocale("zh-CN")
		assert.Equal(t, "八月", Parse("2020-08-05").SetLanguage(lang).ToMonthString())
	})
}

func TestCarbon_SetDateTime(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", NewCarbon().SetDateTime(2020, 8, 5, 13, 14, 15).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDateTime(2020, 8, 5, 13, 14, 15).ToString())
		assert.Empty(t, Parse("0").SetDateTime(2020, 8, 5, 13, 14, 15).ToString())
		assert.Empty(t, Parse("xxx").SetDateTime(2020, 8, 5, 13, 14, 15).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", Parse("2020-08-05").SetDateTime(2020, 8, 5, 13, 14, 15).ToString())
	})
}

func TestCarbon_SetDateTimeMilli(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999 +0000 UTC", NewCarbon().SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString())
		assert.Empty(t, Parse("0").SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString())
		assert.Empty(t, Parse("xxx").SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999 +0000 UTC", Parse("2020-08-05").SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString())
	})
}

func TestCarbon_SetDateTimeMicro(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999999 +0000 UTC", NewCarbon().SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString())
		assert.Empty(t, Parse("0").SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString())
		assert.Empty(t, Parse("xxx").SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999999 +0000 UTC", Parse("2020-08-05").SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString())
	})
}

func TestCarbon_SetDateTimeNano(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999999999 +0000 UTC", NewCarbon().SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString())
		assert.Empty(t, Parse("0").SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString())
		assert.Empty(t, Parse("xxx").SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999999999 +0000 UTC", Parse("2020-08-05").SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString())
	})
}

func TestCarbon_SetDate(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", NewCarbon().SetDate(2020, 8, 5).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDate(2020, 8, 5).ToString())
		assert.Empty(t, Parse("0").SetDate(2020, 8, 5).ToString())
		assert.Empty(t, Parse("xxx").SetDate(2020, 8, 5).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", Parse("2020-08-05").SetDate(2020, 8, 5).ToString())
	})
}

func TestCarbon_SetDateMilli(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00.999 +0000 UTC", NewCarbon().SetDateMilli(2020, 8, 5, 999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDateMilli(2020, 8, 5, 999).ToString())
		assert.Empty(t, Parse("0").SetDateMilli(2020, 8, 5, 999).ToString())
		assert.Empty(t, Parse("xxx").SetDateMilli(2020, 8, 5, 999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00.999 +0000 UTC", Parse("2020-08-05").SetDateMilli(2020, 8, 5, 999).ToString())
	})
}

func TestCarbon_SetDateMicro(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00.999999 +0000 UTC", NewCarbon().SetDateMicro(2020, 8, 5, 999999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDateMicro(2020, 8, 5, 999999).ToString())
		assert.Empty(t, Parse("0").SetDateMicro(2020, 8, 5, 999999).ToString())
		assert.Empty(t, Parse("xxx").SetDateMicro(2020, 8, 5, 999999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00.999999 +0000 UTC", Parse("2020-08-05").SetDateMicro(2020, 8, 5, 999999).ToString())
	})
}

func TestCarbon_SetDateNano(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00.999999999 +0000 UTC", NewCarbon().SetDateNano(2020, 8, 5, 999999999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDateNano(2020, 8, 5, 999999999).ToString())
		assert.Empty(t, Parse("0").SetDateNano(2020, 8, 5, 999999999).ToString())
		assert.Empty(t, Parse("xxx").SetDateNano(2020, 8, 5, 999999999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00.999999999 +0000 UTC", Parse("2020-08-05").SetDateNano(2020, 8, 5, 999999999).ToString())
	})
}

func TestCarbon_SetTime(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 13:14:15 +0000 UTC", NewCarbon().SetTime(13, 14, 15).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetTime(13, 14, 15).ToString())
		assert.Empty(t, Parse("0").SetTime(13, 14, 15).ToString())
		assert.Empty(t, Parse("xxx").SetTime(13, 14, 15).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", Parse("2020-08-05").SetTime(13, 14, 15).ToString())
	})
}

func TestCarbon_SetTimeMilli(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 13:14:15.999 +0000 UTC", NewCarbon().SetTimeMilli(13, 14, 15, 999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetTimeMilli(13, 14, 15, 999).ToString())
		assert.Empty(t, Parse("0").SetTimeMilli(13, 14, 15, 999).ToString())
		assert.Empty(t, Parse("xxx").SetTimeMilli(13, 14, 15, 999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999 +0000 UTC", Parse("2020-08-05").SetTimeMilli(13, 14, 15, 999).ToString())
	})
}

func TestCarbon_SetTimeMicro(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 13:14:15.999999 +0000 UTC", NewCarbon().SetTimeMicro(13, 14, 15, 999999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetTimeMicro(13, 14, 15, 9999999).ToString())
		assert.Empty(t, Parse("0").SetTimeMicro(13, 14, 15, 9999999).ToString())
		assert.Empty(t, Parse("xxx").SetTimeMicro(13, 14, 15, 9999999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999999 +0000 UTC", Parse("2020-08-05").SetTimeMicro(13, 14, 15, 999999).ToString())
	})
}

func TestCarbon_SetTimeNano(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 13:14:15.999999999 +0000 UTC", NewCarbon().SetTimeNano(13, 14, 15, 999999999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetTimeNano(13, 14, 15, 999999999).ToString())
		assert.Empty(t, Parse("0").SetTimeNano(13, 14, 15, 999999999).ToString())
		assert.Empty(t, Parse("xxx").SetTimeNano(13, 14, 15, 999999999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.999999999 +0000 UTC", Parse("2020-08-05").SetTimeNano(13, 14, 15, 999999999).ToString())
	})
}

func TestCarbon_SetYear(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", NewCarbon().SetYear(2020).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetYear(2020).ToString())
		assert.Empty(t, Parse("0").SetYear(2020).ToString())
		assert.Empty(t, Parse("xxx").SetYear(2020).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2019-01-01 00:00:00 +0000 UTC", Parse("2020-01-01").SetYear(2019).ToString())
		assert.Equal(t, "2019-01-31 00:00:00 +0000 UTC", Parse("2020-01-31").SetYear(2019).ToString())
		assert.Equal(t, "2019-03-01 00:00:00 +0000 UTC", Parse("2020-02-29").SetYear(2019).ToString())
	})
}

func TestCarbon_SetYearNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", NewCarbon().SetYearNoOverflow(2020).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetYearNoOverflow(2020).ToString())
		assert.Empty(t, Parse("0").SetYearNoOverflow(2020).ToString())
		assert.Empty(t, Parse("xxx").SetYearNoOverflow(2020).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2019-01-01 00:00:00 +0000 UTC", Parse("2020-01-01").SetYearNoOverflow(2019).ToString())
		assert.Equal(t, "2019-01-31 00:00:00 +0000 UTC", Parse("2020-01-31").SetYearNoOverflow(2019).ToString())
		assert.Equal(t, "2019-02-28 00:00:00 +0000 UTC", Parse("2020-02-29").SetYearNoOverflow(2019).ToString())
	})
}

func TestCarbon_SetMonth(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-02-01 00:00:00 +0000 UTC", NewCarbon().SetMonth(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetMonth(2).ToString())
		assert.Empty(t, Parse("0").SetMonth(2).ToString())
		assert.Empty(t, Parse("xxx").SetMonth(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-02-01 00:00:00 +0000 UTC", Parse("2020-01-01").SetMonth(2).ToString())
		assert.Equal(t, "2020-03-01 00:00:00 +0000 UTC", Parse("2020-01-30").SetMonth(2).ToString())
		assert.Equal(t, "2020-03-02 00:00:00 +0000 UTC", Parse("2020-01-31").SetMonth(2).ToString())
	})
}

func TestCarbon_SetMonthNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-02-01 00:00:00 +0000 UTC", NewCarbon().SetMonthNoOverflow(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetMonthNoOverflow(2).ToString())
		assert.Empty(t, Parse("0").SetMonthNoOverflow(2).ToString())
		assert.Empty(t, Parse("xxx").SetMonthNoOverflow(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-02-01 00:00:00 +0000 UTC", Parse("2020-01-01").SetMonthNoOverflow(2).ToString())
		assert.Equal(t, "2020-02-29 00:00:00 +0000 UTC", Parse("2020-01-30").SetMonthNoOverflow(2).ToString())
		assert.Equal(t, "2020-02-29 00:00:00 +0000 UTC", Parse("2020-01-31").SetMonthNoOverflow(2).ToString())
	})
}

func TestCarbon_SetDay(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-31 00:00:00 +0000 UTC", NewCarbon().SetDay(31).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetDay(31).ToString())
		assert.Empty(t, Parse("0").SetDay(31).ToString())
		assert.Empty(t, Parse("xxx").SetDay(31).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-31 00:00:00 +0000 UTC", Parse("2020-01-01").SetDay(31).ToString())
		assert.Equal(t, "2020-03-02 00:00:00 +0000 UTC", Parse("2020-02-01").SetDay(31).ToString())
		assert.Equal(t, "2020-03-02 00:00:00 +0000 UTC", Parse("2020-02-28").SetDay(31).ToString())
	})
}

func TestCarbon_SetHour(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 10:00:00 +0000 UTC", NewCarbon().SetHour(10).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetHour(31).ToString())
		assert.Empty(t, Parse("0").SetHour(31).ToString())
		assert.Empty(t, Parse("xxx").SetHour(31).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 10:00:00 +0000 UTC", Parse("2020-01-01").SetHour(10).ToString())
		assert.Equal(t, "2020-02-02 00:00:00 +0000 UTC", Parse("2020-02-01").SetHour(24).ToString())
		assert.Equal(t, "2020-02-29 07:00:00 +0000 UTC", Parse("2020-02-28").SetHour(31).ToString())
	})
}

func TestCarbon_SetMinute(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:10:00 +0000 UTC", NewCarbon().SetMinute(10).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetMinute(31).ToString())
		assert.Empty(t, Parse("0").SetMinute(31).ToString())
		assert.Empty(t, Parse("xxx").SetMinute(31).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:10:00 +0000 UTC", Parse("2020-01-01").SetMinute(10).ToString())
		assert.Equal(t, "2020-02-01 00:24:00 +0000 UTC", Parse("2020-02-01").SetMinute(24).ToString())
		assert.Equal(t, "2020-02-28 01:00:00 +0000 UTC", Parse("2020-02-28").SetMinute(60).ToString())
	})
}

func TestCarbon_SetSecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:10 +0000 UTC", NewCarbon().SetSecond(10).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetSecond(31).ToString())
		assert.Empty(t, Parse("0").SetSecond(31).ToString())
		assert.Empty(t, Parse("xxx").SetSecond(31).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:10 +0000 UTC", Parse("2020-01-01").SetSecond(10).ToString())
		assert.Equal(t, "2020-02-01 00:00:24 +0000 UTC", Parse("2020-02-01").SetSecond(24).ToString())
		assert.Equal(t, "2020-02-28 00:01:00 +0000 UTC", Parse("2020-02-28").SetSecond(60).ToString())
	})
}

func TestCarbon_SetMillisecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00.999 +0000 UTC", NewCarbon().SetMillisecond(999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetMillisecond(999).ToString())
		assert.Empty(t, Parse("0").SetMillisecond(999).ToString())
		assert.Empty(t, Parse("xxx").SetMillisecond(999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00.1 +0000 UTC", Parse("2020-01-01").SetMillisecond(100).ToString())
		assert.Equal(t, "2020-01-01 00:00:00.999 +0000 UTC", Parse("2020-01-01").SetMillisecond(999).ToString())
	})
}

func TestCarbon_SetMicrosecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00.999999 +0000 UTC", NewCarbon().SetMicrosecond(999999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetMicrosecond(999999).ToString())
		assert.Empty(t, Parse("0").SetMicrosecond(999999).ToString())
		assert.Empty(t, Parse("xxx").SetMicrosecond(999999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00.1 +0000 UTC", Parse("2020-01-01").SetMicrosecond(100000).ToString())
		assert.Equal(t, "2020-01-01 00:00:00.999999 +0000 UTC", Parse("2020-01-01").SetMicrosecond(999999).ToString())
	})
}

func TestCarbon_SetNanosecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00.999999999 +0000 UTC", NewCarbon().SetNanosecond(999999999).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SetNanosecond(999999999).ToString())
		assert.Empty(t, Parse("0").SetNanosecond(999999999).ToString())
		assert.Empty(t, Parse("xxx").SetNanosecond(999999999).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00.1 +0000 UTC", Parse("2020-01-01").SetNanosecond(100000000).ToString())
		assert.Equal(t, "2020-01-01 00:00:00.999999999 +0000 UTC", Parse("2020-01-01").SetNanosecond(999999999).ToString())
	})
}
