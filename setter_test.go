package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetLayout(t *testing.T) {
	defer SetLayout(DateTimeLayout)

	t.Run("zero carbon", func(t *testing.T) {
		SetLayout(DateLayout)
		assert.Equal(t, DateLayout, DefaultLayout)
		assert.Equal(t, DateLayout, NewCarbon().CurrentLayout())

		SetLayout(DateTimeLayout)
		assert.Equal(t, DateTimeLayout, DefaultLayout)
		assert.Equal(t, DateTimeLayout, NewCarbon().CurrentLayout())
	})

	t.Run("valid carbon", func(t *testing.T) {
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

func TestSetFormat(t *testing.T) {
	defer SetFormat(DateTimeFormat)

	t.Run("zero carbon", func(t *testing.T) {
		SetFormat(DateFormat)
		assert.Equal(t, DateLayout, DefaultLayout)
		assert.Equal(t, DateLayout, NewCarbon().CurrentLayout())

		SetFormat(DateTimeFormat)
		assert.Equal(t, DateTimeLayout, DefaultLayout)
		assert.Equal(t, DateTimeLayout, NewCarbon().CurrentLayout())
	})

	t.Run("valid carbon", func(t *testing.T) {
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

func TestSetTimezone(t *testing.T) {
	defer SetTimezone(UTC)

	t.Run("zero carbon", func(t *testing.T) {
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

	t.Run("valid carbon", func(t *testing.T) {
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

func TestSetLocation(t *testing.T) {
	defer SetLocation(time.UTC)

	t.Run("zero carbon", func(t *testing.T) {
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

	t.Run("valid carbon", func(t *testing.T) {
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

func TestSetLocale(t *testing.T) {
	defer SetLocale("en")

	t.Run("zero carbon", func(t *testing.T) {
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

	t.Run("valid carbon", func(t *testing.T) {
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

func TestSetWeekStartsAt(t *testing.T) {
	defer SetWeekStartsAt(Monday)

	t.Run("zero carbon", func(t *testing.T) {
		SetWeekStartsAt(Sunday)
		assert.Equal(t, Sunday, DefaultWeekStartsAt)
		assert.Equal(t, Sunday, NewCarbon().WeekStartsAt())
		assert.Equal(t, "0000-12-31 00:00:00 +0000 UTC", NewCarbon().StartOfWeek().ToString())

		SetWeekStartsAt(Monday)
		assert.Equal(t, Monday, DefaultWeekStartsAt)
		assert.Equal(t, Monday, NewCarbon().WeekStartsAt())
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfWeek().ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		SetWeekStartsAt(Monday)
		c1 := Parse("2020-08-05")
		assert.Equal(t, Monday, DefaultWeekStartsAt)
		assert.Equal(t, Monday, c1.WeekStartsAt())
		assert.Equal(t, "2020-08-03 00:00:00 +0000 UTC", c1.StartOfWeek().ToString())

		SetWeekStartsAt(Sunday)
		c2 := Parse("2020-08-05")
		assert.Equal(t, "Sunday", DefaultWeekStartsAt.String())
		assert.Equal(t, Sunday, c2.WeekStartsAt())
		assert.Equal(t, "2020-08-02 00:00:00 +0000 UTC", c2.StartOfWeek().ToString())
	})
}

func TestSetWeekendDays(t *testing.T) {
	defer SetWeekendDays([]Weekday{
		Saturday, Sunday,
	})

	t.Run("zero carbon", func(t *testing.T) {
		SetWeekendDays([]Weekday{
			Saturday, Sunday,
		})
		assert.True(t, NewCarbon().IsWeekday())
		assert.False(t, NewCarbon().IsWeekend())
	})

	t.Run("valid carbon", func(t *testing.T) {
		SetWeekendDays([]Weekday{
			Saturday,
		})
		assert.True(t, Parse("2025-04-12").IsWeekend())
		assert.False(t, Parse("2025-04-13").IsWeekend())

		SetWeekendDays([]Weekday{
			Sunday,
		})
		assert.False(t, Parse("2025-04-12").IsWeekend())
		assert.True(t, Parse("2025-04-13").IsWeekend())
	})
}

func TestCarbon_SetLayout(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetLayout(DateLayout)
		assert.False(t, c.HasError())
		assert.Equal(t, DateLayout, c.CurrentLayout())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetLayout(DateLayout)
		assert.False(t, c.HasError())
		assert.Empty(t, c.CurrentLayout())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetLayout(DateLayout)
		assert.True(t, c.HasError())
		assert.Empty(t, c.CurrentLayout())
	})

	t.Run("valid carbon", func(t *testing.T) {
		c1 := Parse("2020-08-05").SetLayout(DateLayout)
		assert.Equal(t, DateLayout, c1.CurrentLayout())

		c2 := Parse("2020-08-05").SetLayout(DateTimeLayout)
		assert.Equal(t, DateTimeLayout, c2.CurrentLayout())
	})

	t.Run("empty layout", func(t *testing.T) {
		c := Now().SetLayout("")
		assert.True(t, c.HasError())
		assert.Empty(t, c.CurrentLayout())
	})
}

func TestCarbon_SetFormat(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetFormat(DateFormat)
		assert.False(t, c.HasError())
		assert.Equal(t, DateLayout, c.CurrentLayout())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetFormat(DateFormat)
		assert.True(t, c.HasError())
		assert.Empty(t, c.CurrentLayout())
	})

	t.Run("valid carbon", func(t *testing.T) {
		c1 := Parse("2020-08-05").SetFormat(DateFormat)
		assert.Equal(t, DateLayout, c1.CurrentLayout())

		c2 := Parse("2020-08-05").SetFormat(DateTimeFormat)
		assert.Equal(t, DateTimeLayout, c2.CurrentLayout())
	})

	t.Run("empty layout", func(t *testing.T) {
		c := Now().SetFormat("")
		assert.True(t, c.HasError())
		assert.Empty(t, c.CurrentLayout())
	})
}

func TestCarbon_SetWeekStartsAt(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c1 := NewCarbon().SetWeekStartsAt(Sunday)
		assert.Equal(t, "0000-12-31 00:00:00 +0000 UTC", c1.StartOfWeek().ToString())

		c2 := NewCarbon().SetWeekStartsAt(Monday)
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", c2.StartOfWeek().ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetWeekStartsAt(Sunday)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetWeekStartsAt(Sunday)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		c1 := Parse("2020-08-05").SetWeekStartsAt(Monday)
		assert.Equal(t, Monday, c1.WeekStartsAt())
		assert.Equal(t, "2020-08-03 00:00:00 +0000 UTC", c1.StartOfWeek().ToString())

		c2 := Parse("2020-08-05").SetWeekStartsAt(Sunday)
		assert.Equal(t, Sunday, c2.WeekStartsAt())
		assert.Equal(t, "2020-08-02 00:00:00 +0000 UTC", c2.StartOfWeek().ToString())
	})
}

func TestCarbon_SetWeekendDays(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		wd := []Weekday{
			Saturday, Sunday,
		}
		c := NewCarbon().SetWeekendDays(wd)
		assert.True(t, c.IsWeekday())
		assert.False(t, c.IsWeekend())
	})

	t.Run("empty carbon", func(t *testing.T) {
		wd := []Weekday{
			Saturday, Sunday,
		}
		c := Parse("").SetWeekendDays(wd)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		wd := []Weekday{
			Saturday, Sunday,
		}
		c := Parse("xxx").SetWeekendDays(wd)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		wd1 := []Weekday{
			Saturday,
		}
		assert.True(t, Parse("2025-04-12").SetWeekendDays(wd1).IsWeekend())
		assert.False(t, Parse("2025-04-13").SetWeekendDays(wd1).IsWeekend())

		wd2 := []Weekday{
			Sunday,
		}
		assert.False(t, Parse("2025-04-12").SetWeekendDays(wd2).IsWeekend())
		assert.True(t, Parse("2025-04-13").SetWeekendDays(wd2).IsWeekend())
	})
}

func TestCarbon_SetLocale(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
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

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetLocale("en")
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetLocale("en")
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
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

	t.Run("empty locale", func(t *testing.T) {
		c := Now().SetLocale("")
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error locale", func(t *testing.T) {
		c := Now().SetLocale("xxx")
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})
}

func TestCarbon_SetTimezone(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
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

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetTimezone(UTC)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetTimezone(UTC)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
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

	t.Run("empty timezone", func(t *testing.T) {
		c := Parse("2020-08-05").SetTimezone("")
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error timezone", func(t *testing.T) {
		c := Parse("2020-08-05").SetTimezone("xxx")
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})
}

func TestCarbon_SetLocation(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
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

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetLocation(time.UTC)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetLocation(time.UTC)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
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

	t.Run("nil location", func(t *testing.T) {
		assert.Empty(t, Parse("2020-08-05").SetLocation(nil).Timezone())
		assert.Empty(t, Parse("2020-08-05").SetLocation(nil).ZoneName())
		assert.Empty(t, Parse("2020-08-05").SetLocation(nil).ZoneOffset())
		assert.Empty(t, Parse("2020-08-05").SetLocation(nil).ToString())
	})
}

func TestCarbon_SetLanguage(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetLocale("en")
		c := NewCarbon().SetLanguage(lang)
		assert.False(t, c.HasError())
		assert.Equal(t, "en", c.Locale())
	})

	t.Run("empty carbon", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetLocale("en")
		c := Parse("").SetLanguage(lang)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetLocale("en")
		c := Parse("xxx").SetLanguage(lang)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		lang := NewLanguage()

		lang.SetLocale("en")
		assert.Equal(t, "August", Parse("2020-08-05").SetLanguage(lang).ToMonthString())

		lang.SetLocale("zh-CN")
		assert.Equal(t, "八月", Parse("2020-08-05").SetLanguage(lang).ToMonthString())
	})

	t.Run("nil language", func(t *testing.T) {
		lang := NewLanguage()
		lang = nil
		c := NewCarbon().SetLanguage(lang)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error language", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetLocale("xxx")
		c := NewCarbon().SetLanguage(lang)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})
}

func TestCarbon_SetDateTime(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetDateTime(2020, 8, 5, 13, 14, 15)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetDateTime(2020, 8, 5, 13, 14, 15)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetDateTime(2020, 8, 5, 13, 14, 15)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		c := Parse("2020-08-05").SetDateTime(2020, 8, 5, 13, 14, 15)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", c.ToString())
	})
}

func TestCarbon_SetDateTimeMilli(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 13:14:15.999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		c := Parse("2020-08-05").SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 13:14:15.999 +0000 UTC", c.ToString())
	})
}

func TestCarbon_SetDateTimeMicro(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := Parse("2020-08-05").SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 13:14:15.999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		c := Parse("2020-08-05").SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 13:14:15.999999 +0000 UTC", c.ToString())
	})
}

func TestCarbon_SetDateTimeNano(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 13:14:15.999999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		c := Parse("2020-08-05").SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 13:14:15.999999999 +0000 UTC", c.ToString())
	})
}

func TestCarbon_SetDate(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := Parse("2020-08-05").SetDate(2020, 8, 5)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetDate(2020, 8, 5)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetDate(2020, 8, 5)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		c := Parse("2020-08-05").SetDate(2020, 8, 5)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", c.ToString())
	})
}

func TestCarbon_SetDateMilli(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetDateMilli(2020, 8, 5, 999)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 00:00:00.999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetDateMilli(2020, 8, 5, 999)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetDateMilli(2020, 8, 5, 999)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		c := Parse("2020-08-05").SetDateMilli(2020, 8, 5, 999)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 00:00:00.999 +0000 UTC", Parse("2020-08-05").SetDateMilli(2020, 8, 5, 999).ToString())
	})
}

func TestCarbon_SetDateMicro(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetDateMicro(2020, 8, 5, 999999)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 00:00:00.999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetDateMicro(2020, 8, 5, 999999)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetDateMicro(2020, 8, 5, 999999)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		c := Parse("2020-08-05").SetDateMicro(2020, 8, 5, 999999)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 00:00:00.999999 +0000 UTC", c.ToString())
	})
}

func TestCarbon_SetDateNano(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetDateNano(2020, 8, 5, 999999999)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 00:00:00.999999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetDateNano(2020, 8, 5, 999999999)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetDateNano(2020, 8, 5, 999999999)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		c := Parse("2020-08-05").SetDateNano(2020, 8, 5, 999999999)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 00:00:00.999999999 +0000 UTC", c.ToString())
	})
}

func TestCarbon_SetTime(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetTime(13, 14, 15)
		assert.False(t, c.HasError())
		assert.Equal(t, "0001-01-01 13:14:15 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetTime(13, 14, 15)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetTime(13, 14, 15)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		c := Parse("2020-08-05").SetTime(13, 14, 15)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", c.ToString())
	})
}

func TestCarbon_SetTimeMilli(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetTimeMilli(13, 14, 15, 999)
		assert.False(t, c.HasError())
		assert.Equal(t, "0001-01-01 13:14:15.999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetTimeMilli(13, 14, 15, 999)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetTimeMilli(13, 14, 15, 999)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		c := Parse("2020-08-05").SetTimeMilli(13, 14, 15, 999)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 13:14:15.999 +0000 UTC", c.ToString())
	})
}

func TestCarbon_SetTimeMicro(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetTimeMicro(13, 14, 15, 999999)
		assert.False(t, c.HasError())
		assert.Equal(t, "0001-01-01 13:14:15.999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetTimeMicro(13, 14, 15, 999999)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetTimeMicro(13, 14, 15, 999999)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		c := Parse("2020-08-05").SetTimeMicro(13, 14, 15, 999999)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 13:14:15.999999 +0000 UTC", c.ToString())
	})
}

func TestCarbon_SetTimeNano(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetTimeNano(13, 14, 15, 999999999)
		assert.False(t, c.HasError())
		assert.Equal(t, "0001-01-01 13:14:15.999999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetTimeNano(13, 14, 15, 999999999)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetTimeNano(13, 14, 15, 999999999)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		c := Parse("2020-08-05").SetTimeNano(13, 14, 15, 999999999)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-08-05 13:14:15.999999999 +0000 UTC", c.ToString())
	})
}

func TestCarbon_SetYear(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetYear(2020)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetYear(2020)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetYear(2020)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		assert.Equal(t, "2019-01-01 00:00:00 +0000 UTC", Parse("2020-01-01").SetYear(2019).ToString())
		assert.Equal(t, "2019-01-31 00:00:00 +0000 UTC", Parse("2020-01-31").SetYear(2019).ToString())
		assert.Equal(t, "2019-03-01 00:00:00 +0000 UTC", Parse("2020-02-29").SetYear(2019).ToString())
	})
}

func TestCarbon_SetYearNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().SetYearNoOverflow(2020)
		assert.False(t, c.HasError())
		assert.Equal(t, "2020-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetYearNoOverflow(2020)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetYearNoOverflow(2020)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
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

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetMonth(2)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetMonth(2)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-02-01 00:00:00 +0000 UTC", Parse("2020-01-01").SetMonth(2).ToString())
		assert.Equal(t, "2020-03-01 00:00:00 +0000 UTC", Parse("2020-01-30").SetMonth(2).ToString())
		assert.Equal(t, "2020-03-02 00:00:00 +0000 UTC", Parse("2020-01-31").SetMonth(2).ToString())
	})
}

func TestCarbon_SetMonthNoOverflow(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetMonthNoOverflow(2)
		assert.False(t, c.HasError())
		assert.Equal(t, "0001-02-01 00:00:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetMonthNoOverflow(2)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetMonthNoOverflow(2)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		assert.Equal(t, "2020-02-01 00:00:00 +0000 UTC", Parse("2020-01-01").SetMonthNoOverflow(2).ToString())
		assert.Equal(t, "2020-02-29 00:00:00 +0000 UTC", Parse("2020-01-30").SetMonthNoOverflow(2).ToString())
		assert.Equal(t, "2020-02-29 00:00:00 +0000 UTC", Parse("2020-01-31").SetMonthNoOverflow(2).ToString())
	})
}

func TestCarbon_SetDay(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetDay(31)
		assert.False(t, c.HasError())
		assert.Equal(t, "0001-01-31 00:00:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetDay(31)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetDay(31)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		assert.Equal(t, "2020-01-31 00:00:00 +0000 UTC", Parse("2020-01-01").SetDay(31).ToString())
		assert.Equal(t, "2020-03-02 00:00:00 +0000 UTC", Parse("2020-02-01").SetDay(31).ToString())
		assert.Equal(t, "2020-03-02 00:00:00 +0000 UTC", Parse("2020-02-28").SetDay(31).ToString())
	})
}

func TestCarbon_SetHour(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetHour(10)
		assert.False(t, c.HasError())
		assert.Equal(t, "0001-01-01 10:00:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetHour(10)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetHour(10)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 10:00:00 +0000 UTC", Parse("2020-01-01").SetHour(10).ToString())
		assert.Equal(t, "2020-02-02 00:00:00 +0000 UTC", Parse("2020-02-01").SetHour(24).ToString())
		assert.Equal(t, "2020-02-29 07:00:00 +0000 UTC", Parse("2020-02-28").SetHour(31).ToString())
	})
}

func TestCarbon_SetMinute(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetMinute(10)
		assert.False(t, c.HasError())
		assert.Equal(t, "0001-01-01 00:10:00 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetMinute(31)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetMinute(31)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:10:00 +0000 UTC", Parse("2020-01-01").SetMinute(10).ToString())
		assert.Equal(t, "2020-02-01 00:24:00 +0000 UTC", Parse("2020-02-01").SetMinute(24).ToString())
		assert.Equal(t, "2020-02-28 01:00:00 +0000 UTC", Parse("2020-02-28").SetMinute(60).ToString())
	})
}

func TestCarbon_SetSecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().SetSecond(10)
		assert.False(t, c.HasError())
		assert.Equal(t, "0001-01-01 00:00:10 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetSecond(10)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetSecond(10)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:10 +0000 UTC", Parse("2020-01-01").SetSecond(10).ToString())
		assert.Equal(t, "2020-02-01 00:00:24 +0000 UTC", Parse("2020-02-01").SetSecond(24).ToString())
		assert.Equal(t, "2020-02-28 00:01:00 +0000 UTC", Parse("2020-02-28").SetSecond(60).ToString())
	})
}

func TestCarbon_SetMillisecond(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetMillisecond(999)
		assert.False(t, c.HasError())
		assert.Equal(t, "0001-01-01 00:00:00.999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetMillisecond(999)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetMillisecond(999)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00.1 +0000 UTC", Parse("2020-01-01").SetMillisecond(100).ToString())
		assert.Equal(t, "2020-01-01 00:00:00.999 +0000 UTC", Parse("2020-01-01").SetMillisecond(999).ToString())
	})
}

func TestCarbon_SetMicrosecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon().SetMicrosecond(999999)
		assert.False(t, c.HasError())
		assert.Equal(t, "0001-01-01 00:00:00.999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetMicrosecond(999999)
		assert.False(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetMicrosecond(999999)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00.1 +0000 UTC", Parse("2020-01-01").SetMicrosecond(100000).ToString())
		assert.Equal(t, "2020-01-01 00:00:00.999999 +0000 UTC", Parse("2020-01-01").SetMicrosecond(999999).ToString())
	})
}

func TestCarbon_SetNanosecond(t *testing.T) {
	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon().SetNanosecond(999999999)
		assert.False(t, c.HasError())
		assert.Equal(t, "0001-01-01 00:00:00.999999999 +0000 UTC", c.ToString())
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("").SetNanosecond(999999999)
		assert.Empty(t, c.ToString())
		assert.False(t, c.HasError())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx").SetNanosecond(999999999)
		assert.True(t, c.HasError())
		assert.Empty(t, c.ToString())
	})

	t.Run("valid carbon", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 00:00:00.1 +0000 UTC", Parse("2020-01-01").SetNanosecond(100000000).ToString())
		assert.Equal(t, "2020-01-01 00:00:00.999999999 +0000 UTC", Parse("2020-01-01").SetNanosecond(999999999).ToString())
	})
}
