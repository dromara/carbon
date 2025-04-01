package persian

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMaxValue(t *testing.T) {
	assert.Equal(t, "9377-12-31", MaxValue().String())
}

func TestMinValue(t *testing.T) {
	assert.Equal(t, "0001-01-01", MinValue().String())
}

func TestFromStdTime(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("zero time", func(t *testing.T) {
		assert.Empty(t, FromStdTime(time.Time{}).String())
		assert.Empty(t, FromStdTime(time.Time{}.In(loc)).String())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "1178-10-11", FromStdTime(time.Date(1800, 1, 1, 0, 0, 0, 0, time.UTC)).String())
		assert.Equal(t, "1402-10-11", FromStdTime(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)).String())
		assert.Equal(t, "1403-05-15", FromStdTime(time.Date(2024, 8, 5, 0, 0, 0, 0, time.UTC)).String())
		assert.Equal(t, "1403-10-11", FromStdTime(time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)).String())
		assert.Equal(t, "2024-01-01", FromStdTime(time.Date(2645, 3, 21, 0, 0, 0, 0, time.UTC)).String())
	})
}

func TestPersian_Gregorian(t *testing.T) {
	t.Run("invalid lunar", func(t *testing.T) {
		assert.Empty(t, new(Persian).ToGregorian().String())
		assert.Empty(t, NewPersian(9400, 1, 1).ToGregorian().String())
	})

	t.Run("invalid timezone", func(t *testing.T) {
		g := NewPersian(1395, 1, 1).ToGregorian("xxx")
		assert.Error(t, g.Error)
		assert.Empty(t, g.String())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "2016-03-20 00:00:00 +0000 UTC", NewPersian(1395, 1, 1).ToGregorian().String())
		assert.Equal(t, "2024-01-01 00:00:00 +0000 UTC", NewPersian(1402, 10, 11).ToGregorian().String())
		assert.Equal(t, "2024-08-05 00:00:00 +0000 UTC", NewPersian(1403, 5, 15).ToGregorian().String())
		assert.Equal(t, "2024-12-31 00:00:00 +0000 UTC", NewPersian(1403, 10, 11).ToGregorian().String())
		assert.Equal(t, "2645-03-21 00:00:00 +0000 UTC", NewPersian(2024, 1, 1).ToGregorian().String())
		assert.Equal(t, "2722-03-22 00:00:00 +0000 UTC", NewPersian(2100, 12, 31).ToGregorian().String())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "2016-03-20 00:00:00 +0800 CST", NewPersian(1395, 1, 1).ToGregorian("PRC").String())
		assert.Equal(t, "2024-01-01 00:00:00 +0800 CST", NewPersian(1402, 10, 11).ToGregorian("PRC").String())
		assert.Equal(t, "2024-08-05 00:00:00 +0800 CST", NewPersian(1403, 5, 15).ToGregorian("PRC").String())
		assert.Equal(t, "2024-12-31 00:00:00 +0800 CST", NewPersian(1403, 10, 11).ToGregorian("PRC").String())
		assert.Equal(t, "2645-03-21 00:00:00 +0800 CST", NewPersian(2024, 1, 1).ToGregorian("PRC").String())
		assert.Equal(t, "2722-03-22 00:00:00 +0800 CST", NewPersian(2100, 12, 31).ToGregorian("PRC").String())
	})
}

func TestPersian_Year(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Persian).Year())
		assert.Empty(t, NewPersian(9400, 1, 1).Year())
	})

	t.Run("valid time", func(t *testing.T) {
		p := FromStdTime(time.Date(2024, 1, 1, 0, 0, 0, 0, loc))
		assert.Equal(t, "1402-10-11", p.String())
		assert.Equal(t, 1402, p.Year())
	})
}

func TestPersian_Month(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Persian).Month())
		assert.Empty(t, NewPersian(9400, 1, 1).Month())
	})

	t.Run("valid time", func(t *testing.T) {
		p := FromStdTime(time.Date(2024, 1, 1, 0, 0, 0, 0, loc))
		assert.Equal(t, "1402-10-11", p.String())
		assert.Equal(t, 10, p.Month())
	})
}

func TestPersian_Day(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Persian).Day())
		assert.Empty(t, NewPersian(9400, 1, 1).Day())
	})

	t.Run("valid time", func(t *testing.T) {
		p := FromStdTime(time.Date(2024, 1, 1, 0, 0, 0, 0, loc))
		assert.Equal(t, "1402-10-11", p.String())
		assert.Equal(t, 11, p.Day())
	})
}

func TestPersian_ToMonthString(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Persian).ToMonthString())
		assert.Empty(t, NewPersian(9400, 1, 1).ToMonthString())
	})

	t.Run("invalid locale", func(t *testing.T) {
		p := FromStdTime(time.Date(2020, 2, 1, 0, 0, 0, 0, loc))
		assert.Empty(t, p.ToMonthString("xxx"))
	})

	t.Run("valid time", func(t *testing.T) {
		p := FromStdTime(time.Date(2020, 2, 1, 0, 0, 0, 0, loc))

		assert.Equal(t, "1398-11-12", p.String())
		assert.Equal(t, "Bahman", p.ToMonthString("en"))

		assert.Equal(t, "1398-11-12", p.String())
		assert.Equal(t, "بهمن", p.ToMonthString("fa"))
	})
}

func TestPersian_ToShortMonthString(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Persian).ToShortMonthString())
		assert.Empty(t, NewPersian(9400, 1, 1).ToShortMonthString())
	})

	t.Run("invalid locale", func(t *testing.T) {
		p := FromStdTime(time.Date(2020, 2, 1, 0, 0, 0, 0, loc))
		assert.Empty(t, p.ToShortMonthString("xxx"))
	})

	t.Run("valid time", func(t *testing.T) {
		p := FromStdTime(time.Date(2020, 2, 1, 0, 0, 0, 0, loc))

		assert.Equal(t, "1398-11-12", p.String())
		assert.Equal(t, "Bah", p.ToShortMonthString("en"))

		assert.Equal(t, "1398-11-12", p.String())
		assert.Equal(t, "بهم", p.ToShortMonthString("fa"))
	})
}

func TestPersian_ToWeekString(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Persian).ToWeekString())
		assert.Empty(t, NewPersian(9400, 1, 1).ToWeekString())
	})

	t.Run("invalid locale", func(t *testing.T) {
		p := FromStdTime(time.Date(2020, 2, 1, 0, 0, 0, 0, loc))
		assert.Empty(t, p.ToWeekString("xxx"))
	})

	t.Run("valid time", func(t *testing.T) {
		p := FromStdTime(time.Date(2020, 1, 1, 0, 0, 0, 0, loc))

		assert.Equal(t, "1398-10-11", p.String())
		assert.Equal(t, "Chaharshanbeh", p.ToWeekString("en"))

		assert.Equal(t, "1398-10-11", p.String())
		assert.Equal(t, "چهارشنبه", p.ToWeekString("fa"))
	})
}

func TestPersian_ToShortWeekString(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Persian).ToShortWeekString())
		assert.Empty(t, NewPersian(9400, 1, 1).ToShortWeekString())
	})

	t.Run("invalid locale", func(t *testing.T) {
		p := FromStdTime(time.Date(2020, 2, 1, 0, 0, 0, 0, loc))
		assert.Empty(t, p.ToShortWeekString("xxx"))
	})

	t.Run("valid time", func(t *testing.T) {
		p := FromStdTime(time.Date(2020, 1, 1, 0, 0, 0, 0, loc))

		assert.Equal(t, "1398-10-11", p.String())
		assert.Equal(t, "Cha", p.ToShortWeekString("en"))

		assert.Equal(t, "1398-10-11", p.String())
		assert.Equal(t, "د", p.ToShortWeekString("fa"))
	})
}

func TestPersian_IsLeapYear(t *testing.T) {
	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, new(Persian).IsLeapYear())
		assert.False(t, NewPersian(9400, 1, 1).IsLeapYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, NewPersian(1394, 1, 1).IsLeapYear())
		assert.True(t, NewPersian(1395, 1, 1).IsLeapYear())
	})
}
