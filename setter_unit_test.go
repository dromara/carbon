package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type SetterSuite struct {
	suite.Suite
}

func TestSetterSuite(t *testing.T) {
	suite.Run(t, new(SetterSuite))
}

func (s *SetterSuite) TearDownTest() {
	ResetDefault()
}

func (s *SetterSuite) TestSetLayout() {
	s.Run("zero carbon", func() {
		SetLayout(DateLayout)
		c1 := NewCarbon()
		s.Equal(DateLayout, DefaultLayout)
		s.Equal(DateLayout, c1.CurrentLayout())
		s.Equal("0001-01-01", c1.String())

		SetLayout(DateTimeLayout)
		c2 := NewCarbon()
		s.Equal(DateTimeLayout, DefaultLayout)
		s.Equal(DateTimeLayout, c2.CurrentLayout())
		s.Equal("0001-01-01 00:00:00", c2.String())
	})

	s.Run("valid carbon", func() {
		SetLayout(DateLayout)
		c1 := Parse("2020-08-05")
		s.Equal(DateLayout, DefaultLayout)
		s.Equal(DateLayout, c1.CurrentLayout())
		s.Equal("2020-08-05", c1.String())

		SetLayout(DateTimeLayout)
		c2 := Parse("2020-08-05 13:14:15")
		s.Equal(DateTimeLayout, DefaultLayout)
		s.Equal(DateTimeLayout, c2.CurrentLayout())
		s.Equal("2020-08-05 13:14:15", c2.String())
	})
}

func (s *SetterSuite) TestSetFormat() {
	s.Run("zero carbon", func() {
		SetFormat(DateFormat)
		c1 := NewCarbon()
		s.Equal(DateLayout, DefaultLayout)
		s.Equal(DateLayout, c1.CurrentLayout())
		s.Equal("0001-01-01", c1.String())

		SetFormat(DateTimeFormat)
		c2 := NewCarbon()
		s.Equal(DateTimeLayout, DefaultLayout)
		s.Equal(DateTimeLayout, c2.CurrentLayout())
		s.Equal("0001-01-01 00:00:00", c2.String())
	})

	s.Run("valid carbon", func() {
		SetFormat(DateFormat)
		c1 := Parse("2020-08-05")
		s.Equal(DateLayout, DefaultLayout)
		s.Equal(DateLayout, c1.CurrentLayout())
		s.Equal("2020-08-05", c1.String())

		SetFormat(DateTimeFormat)
		c2 := Parse("2020-08-05 13:14:15")
		s.Equal(DateTimeLayout, DefaultLayout)
		s.Equal(DateTimeLayout, c2.CurrentLayout())
		s.Equal("2020-08-05 13:14:15", c2.String())
	})
}

func (s *SetterSuite) TestSetTimezone() {
	s.Run("zero carbon", func() {
		SetTimezone(UTC)
		c1 := NewCarbon()
		s.Equal(UTC, DefaultTimezone)
		s.Equal(UTC, c1.Timezone())
		s.Equal(UTC, c1.ZoneName())
		s.Equal(0, c1.ZoneOffset())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c1.ToString())

		SetTimezone(PRC)
		c2 := NewCarbon()
		s.Equal(PRC, DefaultTimezone)
		s.Equal(PRC, c2.Timezone())
		s.Equal("LMT", c2.ZoneName())
		s.Equal(29143, c2.ZoneOffset())
		s.Equal("0001-01-01 08:05:43 +0805 LMT", c2.ToString())

		SetTimezone(Japan)
		c3 := NewCarbon()
		s.Equal(Japan, DefaultTimezone)
		s.Equal(Japan, c3.Timezone())
		s.Equal("LMT", c3.ZoneName())
		s.Equal(33539, c3.ZoneOffset())
		s.Equal("0001-01-01 09:18:59 +0918 LMT", c3.ToString())
	})

	s.Run("valid carbon", func() {
		SetTimezone(UTC)
		c1 := Parse("2020-08-05 13:14:15")
		s.Equal(UTC, DefaultTimezone)
		s.Equal(UTC, c1.Timezone())
		s.Equal(UTC, c1.ZoneName())
		s.Equal(0, c1.ZoneOffset())
		s.Equal("2020-08-05 13:14:15 +0000 UTC", c1.ToString())

		SetTimezone(PRC)
		c2 := Parse("2020-08-05 13:14:15")
		s.Equal(PRC, DefaultTimezone)
		s.Equal(PRC, c2.Timezone())
		s.Equal("CST", c2.ZoneName())
		s.Equal(28800, c2.ZoneOffset())
		s.Equal("2020-08-05 13:14:15 +0800 CST", c2.ToString())

		SetTimezone(Japan)
		c3 := Parse("2020-08-05 13:14:15")
		s.Equal(Japan, DefaultTimezone)
		s.Equal(Japan, c3.Timezone())
		s.Equal("JST", c3.ZoneName())
		s.Equal(32400, c3.ZoneOffset())
		s.Equal("2020-08-05 13:14:15 +0900 JST", c3.ToString())
	})
}

func (s *SetterSuite) TestSetLocation() {
	s.Run("zero carbon", func() {
		l1, _ := time.LoadLocation(UTC)
		SetLocation(l1)
		c1 := NewCarbon()
		s.Equal(UTC, DefaultTimezone)
		s.Equal(UTC, c1.Timezone())
		s.Equal(UTC, c1.ZoneName())
		s.Equal(0, c1.ZoneOffset())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c1.ToString())

		l2, _ := time.LoadLocation(PRC)
		SetLocation(l2)
		c2 := NewCarbon()
		s.Equal(PRC, DefaultTimezone)
		s.Equal(PRC, c2.Timezone())
		s.Equal("LMT", c2.ZoneName())
		s.Equal(29143, c2.ZoneOffset())
		s.Equal("0001-01-01 08:05:43 +0805 LMT", c2.ToString())

		l3, _ := time.LoadLocation(Japan)
		SetLocation(l3)
		c3 := NewCarbon()
		s.Equal(Japan, DefaultTimezone)
		s.Equal(Japan, c3.Timezone())
		s.Equal("LMT", c3.ZoneName())
		s.Equal(33539, c3.ZoneOffset())
		s.Equal("0001-01-01 09:18:59 +0918 LMT", c3.ToString())
	})

	s.Run("valid carbon", func() {
		l1, _ := time.LoadLocation(UTC)
		SetLocation(l1)
		c1 := Parse("2020-08-05 13:14:15")
		s.Equal(UTC, DefaultTimezone)
		s.Equal(UTC, c1.Timezone())
		s.Equal(UTC, c1.ZoneName())
		s.Equal(0, c1.ZoneOffset())
		s.Equal("2020-08-05 13:14:15 +0000 UTC", c1.ToString())

		l2, _ := time.LoadLocation(PRC)
		SetLocation(l2)
		c2 := Parse("2020-08-05 13:14:15")
		s.Equal(PRC, DefaultTimezone)
		s.Equal(PRC, c2.Timezone())
		s.Equal("CST", c2.ZoneName())
		s.Equal(28800, c2.ZoneOffset())
		s.Equal("2020-08-05 13:14:15 +0800 CST", c2.ToString())

		l3, _ := time.LoadLocation(Japan)
		SetLocation(l3)
		c3 := Parse("2020-08-05 13:14:15")
		s.Equal(Japan, DefaultTimezone)
		s.Equal(Japan, c3.Timezone())
		s.Equal("JST", c3.ZoneName())
		s.Equal(32400, c3.ZoneOffset())
		s.Equal("2020-08-05 13:14:15 +0900 JST", c3.ToString())
	})
}

func (s *SetterSuite) TestSetLocale() {
	s.Run("zero carbon", func() {
		SetLocale("zh-CN")
		s.Equal("zh-CN", DefaultLocale)
		s.Equal("zh-CN", NewCarbon().Locale())
		s.Equal("摩羯座", NewCarbon().Constellation())
		s.Equal("冬季", NewCarbon().Season())
		s.Equal("一月", NewCarbon().ToMonthString())
		s.Equal("1月", NewCarbon().ToShortMonthString())
		s.Equal("星期一", NewCarbon().ToWeekString())
		s.Equal("周一", NewCarbon().ToShortWeekString())

		SetLocale("en")
		s.Equal("en", DefaultLocale)
		s.Equal("en", NewCarbon().Locale())
		s.Equal("Capricorn", NewCarbon().Constellation())
		s.Equal("Winter", NewCarbon().Season())
		s.Equal("January", NewCarbon().ToMonthString())
		s.Equal("Jan", NewCarbon().ToShortMonthString())
		s.Equal("Monday", NewCarbon().ToWeekString())
		s.Equal("Mon", NewCarbon().ToShortWeekString())
	})

	s.Run("valid carbon", func() {
		SetLocale("zh-CN")
		c1 := Parse("2020-08-05")
		s.Equal("zh-CN", DefaultLocale)
		s.Equal("zh-CN", c1.Locale())
		s.Equal("狮子座", c1.Constellation())
		s.Equal("夏季", c1.Season())
		s.Equal("八月", c1.ToMonthString())
		s.Equal("8月", c1.ToShortMonthString())
		s.Equal("星期三", c1.ToWeekString())
		s.Equal("周三", c1.ToShortWeekString())

		SetLocale("en")
		c2 := Parse("2020-08-05")
		s.Equal("en", DefaultLocale)
		s.Equal("en", c2.Locale())
		s.Equal("Leo", c2.Constellation())
		s.Equal("Summer", c2.Season())
		s.Equal("August", c2.ToMonthString())
		s.Equal("Aug", c2.ToShortMonthString())
		s.Equal("Wednesday", c2.ToWeekString())
		s.Equal("Wed", c2.ToShortWeekString())
	})
}

func (s *SetterSuite) TestSetWeekStartsAt() {
	s.Run("zero carbon", func() {
		SetWeekStartsAt(Sunday)
		s.Equal(Sunday, DefaultWeekStartsAt)
		s.Equal(Sunday, NewCarbon().WeekStartsAt())
		s.Equal("0000-12-31 00:00:00 +0000 UTC", NewCarbon().StartOfWeek().ToString())

		SetWeekStartsAt(Monday)
		s.Equal(Monday, DefaultWeekStartsAt)
		s.Equal(Monday, NewCarbon().WeekStartsAt())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfWeek().ToString())
	})

	s.Run("valid carbon", func() {
		SetWeekStartsAt(Monday)
		c1 := Parse("2020-08-05")
		s.Equal(Monday, DefaultWeekStartsAt)
		s.Equal(Monday, c1.WeekStartsAt())
		s.Equal("2020-08-03 00:00:00 +0000 UTC", c1.StartOfWeek().ToString())

		SetWeekStartsAt(Sunday)
		c2 := Parse("2020-08-05")
		s.Equal(Sunday, DefaultWeekStartsAt)
		s.Equal(Sunday, c2.WeekStartsAt())
		s.Equal("2020-08-02 00:00:00 +0000 UTC", c2.StartOfWeek().ToString())
	})
}

func (s *SetterSuite) TestSetWeekendDays() {
	s.Run("zero carbon", func() {
		SetWeekendDays([]Weekday{
			Saturday, Sunday,
		})
		s.True(NewCarbon().IsWeekday())
		s.False(NewCarbon().IsWeekend())
	})

	s.Run("empty carbon", func() {
		SetWeekendDays([]Weekday{
			Saturday, Sunday,
		})
		s.False(Parse("").IsWeekday())
		s.False(Parse("").IsWeekend())
	})

	s.Run("error carbon", func() {
		SetWeekendDays([]Weekday{
			Saturday, Sunday,
		})
		s.False(Parse("xxx").IsWeekday())
		s.False(Parse("xxx").IsWeekend())
	})

	s.Run("valid carbon", func() {
		SetWeekendDays([]Weekday{
			Saturday,
		})
		s.True(Parse("2025-04-12").IsWeekend())
		s.False(Parse("2025-04-13").IsWeekend())

		SetWeekendDays([]Weekday{
			Sunday,
		})
		s.False(Parse("2025-04-12").IsWeekend())
		s.True(Parse("2025-04-13").IsWeekend())
	})
}

func (s *SetterSuite) TestCarbon_SetLayout() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetLayout(DateLayout)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetLayout(DateLayout)
		s.False(c.HasError())
		s.Equal(DateLayout, c.CurrentLayout())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetLayout(DateLayout)
		s.False(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetLayout(DateLayout)
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15.999999 +0000 UTC")

		s.Equal("2020-08-05 13:14:15", c.SetLayout(DateTimeLayout).String())
		s.Equal("1596633255", c.SetLayout(TimestampLayout).String())
		s.Equal("1596633255999", c.SetLayout(TimestampMilliLayout).String())
		s.Equal("1596633255999999", c.SetLayout(TimestampMicroLayout).String())
		s.Equal("1596633255999999000", c.SetLayout(TimestampNanoLayout).String())
	})

	s.Run("empty layout", func() {
		c := Now().SetLayout("")
		s.True(c.HasError())
		s.Empty(c.CurrentLayout())
		s.Empty(c.String())
	})
}

func (s *SetterSuite) TestCarbon_SetFormat() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetFormat(DateFormat)
		s.False(c.HasError())
		s.Empty(c.CurrentLayout())
		s.Empty(c.String())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetFormat(DateFormat)
		s.False(c.HasError())
		s.Equal(DateLayout, c.CurrentLayout())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetFormat(DateFormat)
		s.False(c.HasError())
		s.Empty(c.CurrentLayout())
		s.Empty(c.String())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetFormat(DateFormat)
		s.True(c.HasError())
		s.Empty(c.CurrentLayout())
		s.Empty(c.String())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15.999999 +0000 UTC")

		s.Equal("2020-08-05 13:14:15", c.SetFormat(DateTimeFormat).String())
		s.Equal("1596633255", c.SetFormat(TimestampFormat).String())
		s.Equal("1596633255999", c.SetFormat(TimestampMilliFormat).String())
		s.Equal("1596633255999999", c.SetFormat(TimestampMicroFormat).String())
		s.Equal("1596633255999999000", c.SetFormat(TimestampNanoFormat).String())
	})

	s.Run("empty layout", func() {
		c := Now().SetFormat("")
		s.True(c.HasError())
		s.Empty(c.CurrentLayout())
		s.Empty(c.String())
	})
}

func (s *SetterSuite) TestCarbon_SetLocale() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetLocale("en")
		s.False(c.HasError())
		s.Empty(c.String())
	})

	s.Run("zero carbon", func() {
		c1 := NewCarbon().SetLocale("zh-CN")
		s.Equal("zh-CN", c1.Locale())
		s.Equal("摩羯座", c1.Constellation())
		s.Equal("冬季", c1.Season())
		s.Equal("一月", c1.ToMonthString())
		s.Equal("1月", c1.ToShortMonthString())
		s.Equal("星期一", c1.ToWeekString())
		s.Equal("周一", c1.ToShortWeekString())

		c2 := NewCarbon().SetLocale("en")
		s.Equal("Capricorn", c2.Constellation())
		s.Equal("Winter", c2.Season())
		s.Equal("January", c2.ToMonthString())
		s.Equal("Jan", c2.ToShortMonthString())
		s.Equal("Monday", c2.ToWeekString())
		s.Equal("Mon", c2.ToShortWeekString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetLocale("en")
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetLocale("en")
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c1 := Parse("2020-08-05").SetLocale("zh-CN")
		s.Equal("zh-CN", c1.Locale())
		s.Equal("狮子座", c1.Constellation())
		s.Equal("夏季", c1.Season())
		s.Equal("八月", c1.ToMonthString())
		s.Equal("8月", c1.ToShortMonthString())
		s.Equal("星期三", c1.ToWeekString())
		s.Equal("周三", c1.ToShortWeekString())

		c2 := Parse("2020-08-05").SetLocale("en")
		s.Equal("en", c2.Locale())
		s.Equal("Leo", c2.Constellation())
		s.Equal("Summer", c2.Season())
		s.Equal("August", c2.ToMonthString())
		s.Equal("Aug", c2.ToShortMonthString())
		s.Equal("Wednesday", c2.ToWeekString())
		s.Equal("Wed", c2.ToShortWeekString())
	})

	s.Run("empty locale", func() {
		c := Now().SetLocale("")
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error locale", func() {
		c := Now().SetLocale("xxx")
		s.True(c.HasError())
		s.Empty(c.ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetTimezone() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetTimezone(UTC)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		l1, _ := time.LoadLocation(UTC)
		t1 := time.Time{}.In(l1)
		n1, f1 := t1.Zone()
		s.Equal(UTC, t1.Location().String())
		s.Equal(UTC, n1)
		s.Equal(0, f1)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", t1.String())

		c1 := NewCarbon().SetTimezone(UTC)
		s.Equal(UTC, c1.Timezone())
		s.Equal(UTC, c1.ZoneName())
		s.Equal(0, c1.ZoneOffset())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c1.ToString())

		l2, _ := time.LoadLocation(PRC)
		t2 := time.Time{}.In(l2)
		n2, f2 := t2.Zone()
		s.Equal(PRC, t2.Location().String())
		s.Equal("LMT", n2)
		s.Equal(29143, f2)
		s.Equal("0001-01-01 08:05:43 +0805 LMT", t2.String())

		c2 := NewCarbon().SetTimezone(PRC)
		s.Equal(PRC, c2.Timezone())
		s.Equal("LMT", c2.ZoneName())
		s.Equal(29143, c2.ZoneOffset())
		s.Equal("0001-01-01 08:05:43 +0805 LMT", c2.ToString())

		l3, _ := time.LoadLocation(Japan)
		t3 := time.Time{}.In(l3)
		n3, f3 := t3.Zone()
		s.Equal("LMT", n3)
		s.Equal(33539, f3)
		s.Equal("0001-01-01 09:18:59 +0918 LMT", t3.String())

		c3 := NewCarbon().SetTimezone(Japan)
		s.Equal(Japan, c3.Timezone())
		s.Equal("LMT", c3.ZoneName())
		s.Equal(33539, c3.ZoneOffset())
		s.Equal("0001-01-01 09:18:59 +0918 LMT", c3.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetTimezone(PRC)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetTimezone(PRC)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05")

		c.SetTimezone(UTC)
		s.Equal(UTC, c.Timezone())
		s.Equal(UTC, c.ZoneName())
		s.Equal(0, c.ZoneOffset())
		s.Equal("2020-08-05 00:00:00 +0000 UTC", c.ToString())

		c.SetTimezone(PRC)
		s.Equal(PRC, c.Timezone())
		s.Equal("CST", c.ZoneName())
		s.Equal(28800, c.ZoneOffset())
		s.Equal("2020-08-05 08:00:00 +0800 CST", c.ToString())

		c.SetTimezone(Japan)
		s.Equal(Japan, c.Timezone())
		s.Equal("JST", c.ZoneName())
		s.Equal(32400, c.ZoneOffset())
		s.Equal("2020-08-05 09:00:00 +0900 JST", c.ToString())
	})

	s.Run("empty timezone", func() {
		c := Parse("2020-08-05").SetTimezone("")
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error timezone", func() {
		c := Parse("2020-08-05").SetTimezone("XXX")
		s.True(c.HasError())
		s.Empty(c.ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetLocation() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetLocation(time.UTC)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		l1, _ := time.LoadLocation(UTC)
		t1 := time.Time{}.In(l1)
		n1, f1 := t1.Zone()
		s.Equal(UTC, t1.Location().String())
		s.Equal(UTC, n1)
		s.Equal(0, f1)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", t1.String())

		c1 := NewCarbon().SetLocation(l1)
		s.Equal(UTC, c1.Timezone())
		s.Equal(UTC, c1.ZoneName())
		s.Equal(0, c1.ZoneOffset())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c1.ToString())

		l2, _ := time.LoadLocation(PRC)
		t2 := time.Time{}.In(l2)
		n2, f2 := t2.Zone()
		s.Equal(PRC, t2.Location().String())
		s.Equal("LMT", n2)
		s.Equal(29143, f2)
		s.Equal("0001-01-01 08:05:43 +0805 LMT", t2.String())

		c2 := NewCarbon().SetLocation(l2)
		s.Equal(PRC, c2.Timezone())
		s.Equal("LMT", c2.ZoneName())
		s.Equal(29143, c2.ZoneOffset())
		s.Equal("0001-01-01 08:05:43 +0805 LMT", c2.ToString())

		l3, _ := time.LoadLocation(Japan)
		t3 := time.Time{}.In(l3)
		n3, f3 := t3.Zone()
		s.Equal("LMT", n3)
		s.Equal(33539, f3)
		s.Equal("0001-01-01 09:18:59 +0918 LMT", t3.String())

		c3 := NewCarbon().SetLocation(l3)
		s.Equal(Japan, c3.Timezone())
		s.Equal("LMT", c3.ZoneName())
		s.Equal(33539, c3.ZoneOffset())
		s.Equal("0001-01-01 09:18:59 +0918 LMT", c3.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetLocation(time.UTC)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetLocation(time.UTC)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05")

		l1, _ := time.LoadLocation(UTC)
		c.SetLocation(l1)
		s.Equal(UTC, c.Timezone())
		s.Equal(UTC, c.ZoneName())
		s.Equal(0, c.ZoneOffset())
		s.Equal("2020-08-05 00:00:00 +0000 UTC", c.ToString())

		l2, _ := time.LoadLocation(PRC)
		c.SetLocation(l2)
		s.Equal(PRC, c.Timezone())
		s.Equal("CST", c.ZoneName())
		s.Equal(28800, c.ZoneOffset())
		s.Equal("2020-08-05 08:00:00 +0800 CST", c.ToString())

		l3, _ := time.LoadLocation(Japan)
		c.SetLocation(l3)
		s.Equal(Japan, c.Timezone())
		s.Equal("JST", c.ZoneName())
		s.Equal(32400, c.ZoneOffset())
		s.Equal("2020-08-05 09:00:00 +0900 JST", c.ToString())
	})

	s.Run("nil location", func() {
		c := Parse("2020-08-05").SetLocation(nil)
		s.True(c.HasError())
		s.Empty(c.Timezone())
		s.Empty(c.ZoneName())
		s.Empty(c.ZoneOffset())
		s.Empty(c.ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetWeekStartsAt() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetWeekStartsAt(Sunday)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c1 := NewCarbon().SetWeekStartsAt(Sunday)
		s.False(c1.HasError())
		s.Equal("0000-12-31 00:00:00 +0000 UTC", c1.StartOfWeek().ToString())

		c2 := NewCarbon().SetWeekStartsAt(Monday)
		s.False(c2.HasError())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c2.StartOfWeek().ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetWeekStartsAt(Sunday)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetWeekStartsAt(Sunday)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c1 := Parse("2020-08-05").SetWeekStartsAt(Monday)
		s.Equal(Monday, c1.WeekStartsAt())
		s.Equal("2020-08-03 00:00:00 +0000 UTC", c1.StartOfWeek().ToString())

		c2 := Parse("2020-08-05").SetWeekStartsAt(Sunday)
		s.Equal(Sunday, c2.WeekStartsAt())
		s.Equal("2020-08-02 00:00:00 +0000 UTC", c2.StartOfWeek().ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetWeekendDays() {
	s.Run("nil carbon", func() {
		wd := []Weekday{
			Saturday, Sunday,
		}
		var c *Carbon
		c = nil
		c = c.SetWeekendDays(wd)
		s.False(c.HasError())
		s.False(c.IsWeekday())
		s.False(c.IsWeekend())
	})

	s.Run("zero carbon", func() {
		wd := []Weekday{
			Saturday, Sunday,
		}
		c := NewCarbon().SetWeekendDays(wd)
		s.False(c.HasError())
		s.True(c.IsWeekday())
		s.False(c.IsWeekend())
	})

	s.Run("empty carbon", func() {
		wd := []Weekday{
			Saturday, Sunday,
		}
		c := Parse("").SetWeekendDays(wd)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		wd := []Weekday{
			Saturday, Sunday,
		}
		c := Parse("xxx").SetWeekendDays(wd)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		wd1 := []Weekday{
			Saturday,
		}
		s.True(Parse("2025-04-12").SetWeekendDays(wd1).IsWeekend())
		s.False(Parse("2025-04-13").SetWeekendDays(wd1).IsWeekend())

		wd2 := []Weekday{
			Sunday,
		}
		s.False(Parse("2025-04-12").SetWeekendDays(wd2).IsWeekend())
		s.True(Parse("2025-04-13").SetWeekendDays(wd2).IsWeekend())
	})
}

func (s *SetterSuite) TestCarbon_SetLanguage() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetLanguage(NewLanguage())
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		lang := NewLanguage()
		lang.SetLocale("en")
		c := NewCarbon().SetLanguage(lang)
		s.False(c.HasError())
		s.Equal("en", c.Locale())
	})

	s.Run("empty carbon", func() {
		lang := NewLanguage()
		lang.SetLocale("en")
		c := Parse("").SetLanguage(lang)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		lang := NewLanguage()
		lang.SetLocale("en")
		s.Empty(Parse("xxx").SetLanguage(lang).ToString())
	})

	s.Run("valid carbon", func() {
		lang := NewLanguage()

		lang.SetLocale("en")
		s.Equal("August", Parse("2020-08-05").SetLanguage(lang).ToMonthString())

		lang.SetLocale("zh-CN")
		s.Equal("八月", Parse("2020-08-05").SetLanguage(lang).ToMonthString())
	})

	s.Run("nil language", func() {
		lang := NewLanguage()
		lang = nil
		c := NewCarbon().SetLanguage(lang)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error language", func() {
		lang := NewLanguage()
		lang.SetLocale("xxx")
		c := NewCarbon().SetLanguage(lang)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	// https://github.com/dromara/carbon/issues/353
	s.Run("customizing the source language afterwards", func() {
		lang := NewLanguage()
		lang.SetLocale("en")
		c := Parse("2020-08-05").SetLanguage(lang)
		s.Equal("August", c.ToMonthString())

		lang.SetResources(map[string]string{
			"months": "Ⅰ|Ⅱ|Ⅲ|Ⅳ|Ⅴ|Ⅵ|Ⅶ|Ⅷ|Ⅸ|Ⅹ|Ⅺ|Ⅻ",
		})
		s.Equal("August", c.ToMonthString())
		s.Equal("August", Parse("2020-08-05").ToMonthString())
	})
}

func (s *SetterSuite) TestCarbon_SetDateTime() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetDateTime(2020, 8, 5, 13, 14, 15)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetDateTime(2020, 8, 5, 13, 14, 15)
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:15 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetDateTime(2020, 8, 5, 13, 14, 15)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetDateTime(2020, 8, 5, 13, 14, 15)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05").SetDateTime(2020, 8, 5, 13, 14, 15)
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:15 +0000 UTC", c.ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetDateTimeMilli() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999)
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:15.999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05").SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999)
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:15.999 +0000 UTC", c.ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetDateTimeMicro() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999)
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:15.999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05").SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999)
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:15.999999 +0000 UTC", c.ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetDateTimeNano() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999)
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:15.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05").SetDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999)
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:15.999999999 +0000 UTC", c.ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetDate() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetDate(2020, 8, 5)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetDate(2020, 8, 5)
		s.False(c.HasError())
		s.Equal("2020-08-05 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetDate(2020, 8, 5)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SetDate(2020, 8, 5).ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05").SetDate(2020, 8, 5)
		s.False(c.HasError())
		s.Equal("2020-08-05 00:00:00 +0000 UTC", c.ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetDateMilli() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetDateMilli(2020, 8, 5, 999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetDateMilli(2020, 8, 5, 999)
		s.False(c.HasError())
		s.Equal("2020-08-05 00:00:00.999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetDateMilli(2020, 8, 5, 999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SetDateMilli(2020, 8, 5, 999).ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05").SetDateMilli(2020, 8, 5, 999)
		s.False(c.HasError())
		s.Equal("2020-08-05 00:00:00.999 +0000 UTC", c.ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetDateMicro() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetDateMicro(2020, 8, 5, 999999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetDateMicro(2020, 8, 5, 999999)
		s.False(c.HasError())
		s.Equal("2020-08-05 00:00:00.999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetDateMicro(2020, 8, 5, 999999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetDateMicro(2020, 8, 5, 999999)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05").SetDateMicro(2020, 8, 5, 999999)
		s.False(c.HasError())
		s.Equal("2020-08-05 00:00:00.999999 +0000 UTC", c.ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetDateNano() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetDateNano(2020, 8, 5, 999999999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetDateNano(2020, 8, 5, 999999999)
		s.False(c.HasError())
		s.Equal("2020-08-05 00:00:00.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetDateNano(2020, 8, 5, 999999999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetDateNano(2020, 8, 5, 999999999)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05").SetDateNano(2020, 8, 5, 999999999)
		s.False(c.HasError())
		s.Equal("2020-08-05 00:00:00.999999999 +0000 UTC", c.ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetTime() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetTime(13, 14, 15)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetTime(13, 14, 15)
		s.False(c.HasError())
		s.Equal("0001-01-01 13:14:15 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetTime(13, 14, 15)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetTime(13, 14, 15)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05").SetTime(13, 14, 15)
		s.Equal("2020-08-05 13:14:15 +0000 UTC", c.ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetTimeMilli() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetTimeMilli(13, 14, 15, 999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetTimeMilli(13, 14, 15, 999)
		s.False(c.HasError())
		s.Equal("0001-01-01 13:14:15.999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetTimeMilli(13, 14, 15, 999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetTimeMilli(13, 14, 15, 999)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05").SetTimeMilli(13, 14, 15, 999)
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:15.999 +0000 UTC", c.ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetTimeMicro() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetTimeMicro(13, 14, 15, 999999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetTimeMicro(13, 14, 15, 999999)
		s.False(c.HasError())
		s.Equal("0001-01-01 13:14:15.999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetTimeMicro(13, 14, 15, 9999999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetTimeMicro(13, 14, 15, 9999999)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05").SetTimeMicro(13, 14, 15, 999999)
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:15.999999 +0000 UTC", c.ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetTimeNano() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetTimeNano(13, 14, 15, 999999999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetTimeNano(13, 14, 15, 999999999)
		s.False(c.HasError())
		s.Equal("0001-01-01 13:14:15.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetTimeNano(13, 14, 15, 999999999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetTimeNano(13, 14, 15, 999999999)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05").SetTimeNano(13, 14, 15, 999999999)
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:15.999999999 +0000 UTC", c.ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetYear() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetYear(2020)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetYear(2020)
		s.False(c.HasError())
		s.Equal("2020-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetYear(2020)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetYear(2020)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-01-01 00:00:00 +0000 UTC", Parse("2020-01-01").SetYear(2019).ToString())
		s.Equal("2019-01-31 00:00:00 +0000 UTC", Parse("2020-01-31").SetYear(2019).ToString())
		s.Equal("2019-03-01 00:00:00 +0000 UTC", Parse("2020-02-29").SetYear(2019).ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetYearNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetYearNoOverflow(2020)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetYearNoOverflow(2020)
		s.False(c.HasError())
		s.Equal("2020-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetYearNoOverflow(2020)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetYearNoOverflow(2020)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-01-01 00:00:00 +0000 UTC", Parse("2020-01-01").SetYearNoOverflow(2019).ToString())
		s.Equal("2019-01-31 00:00:00 +0000 UTC", Parse("2020-01-31").SetYearNoOverflow(2019).ToString())
		s.Equal("2019-02-28 00:00:00 +0000 UTC", Parse("2020-02-29").SetYearNoOverflow(2019).ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetMonth() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetMonth(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetMonth(2)
		s.False(c.HasError())
		s.Equal("0001-02-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetMonth(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetMonth(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-02-01 00:00:00 +0000 UTC", Parse("2020-01-01").SetMonth(2).ToString())
		s.Equal("2020-03-01 00:00:00 +0000 UTC", Parse("2020-01-30").SetMonth(2).ToString())
		s.Equal("2020-03-02 00:00:00 +0000 UTC", Parse("2020-01-31").SetMonth(2).ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetMonthNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetMonthNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetMonthNoOverflow(2)
		s.False(c.HasError())
		s.Equal("0001-02-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetMonthNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetMonthNoOverflow(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-02-01 00:00:00 +0000 UTC", Parse("2020-01-01").SetMonthNoOverflow(2).ToString())
		s.Equal("2020-02-29 00:00:00 +0000 UTC", Parse("2020-01-30").SetMonthNoOverflow(2).ToString())
		s.Equal("2020-02-29 00:00:00 +0000 UTC", Parse("2020-01-31").SetMonthNoOverflow(2).ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetDay() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetDay(31)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetDay(31)
		s.False(c.HasError())
		s.Equal("0001-01-31 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetDay(31)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetDay(31)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-31 00:00:00 +0000 UTC", Parse("2020-01-01").SetDay(31).ToString())
		s.Equal("2020-03-02 00:00:00 +0000 UTC", Parse("2020-02-01").SetDay(31).ToString())
		s.Equal("2020-03-02 00:00:00 +0000 UTC", Parse("2020-02-28").SetDay(31).ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetHour() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetHour(10)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetHour(10)
		s.False(c.HasError())
		s.Equal("0001-01-01 10:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetHour(31)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetHour(31)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 10:00:00 +0000 UTC", Parse("2020-01-01").SetHour(10).ToString())
		s.Equal("2020-02-02 00:00:00 +0000 UTC", Parse("2020-02-01").SetHour(24).ToString())
		s.Equal("2020-02-29 07:00:00 +0000 UTC", Parse("2020-02-28").SetHour(31).ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetMinute() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.HasError())
		s.Empty(c.ToString())
		s.Nil(c.SetMinute(10))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetMinute(10)
		s.False(c.HasError())
		s.Equal("0001-01-01 00:10:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetMinute(31)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetMinute(31)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:10:00 +0000 UTC", Parse("2020-01-01").SetMinute(10).ToString())
		s.Equal("2020-02-01 00:24:00 +0000 UTC", Parse("2020-02-01").SetMinute(24).ToString())
		s.Equal("2020-02-28 01:00:00 +0000 UTC", Parse("2020-02-28").SetMinute(60).ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetSecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetSecond(10)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetSecond(10)
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:10 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetSecond(31)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetSecond(31)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:10 +0000 UTC", Parse("2020-01-01").SetSecond(10).ToString())
		s.Equal("2020-02-01 00:00:24 +0000 UTC", Parse("2020-02-01").SetSecond(24).ToString())
		s.Equal("2020-02-28 00:01:00 +0000 UTC", Parse("2020-02-28").SetSecond(60).ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetMillisecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetMillisecond(999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetMillisecond(999)
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00.999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetMillisecond(999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetMillisecond(999)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00.1 +0000 UTC", Parse("2020-01-01").SetMillisecond(100).ToString())
		s.Equal("2020-01-01 00:00:00.999 +0000 UTC", Parse("2020-01-01").SetMillisecond(999).ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetMicrosecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetMicrosecond(999999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetMicrosecond(999999)
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00.999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetMicrosecond(999999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetMicrosecond(999999)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00.1 +0000 UTC", Parse("2020-01-01").SetMicrosecond(100000).ToString())
		s.Equal("2020-01-01 00:00:00.999999 +0000 UTC", Parse("2020-01-01").SetMicrosecond(999999).ToString())
	})
}

func (s *SetterSuite) TestCarbon_SetNanosecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SetNanosecond(999999999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SetNanosecond(999999999)
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SetNanosecond(999999999)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SetNanosecond(999999999)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00.1 +0000 UTC", Parse("2020-01-01").SetNanosecond(100000000).ToString())
		s.Equal("2020-01-01 00:00:00.999999999 +0000 UTC", Parse("2020-01-01").SetNanosecond(999999999).ToString())
	})
}
