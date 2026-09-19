package lunar

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFromStdTime(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("zero time", func(t *testing.T) {
		assert.Empty(t, FromStdTime(time.Time{}).String())
		assert.Empty(t, FromStdTime(time.Time{}.In(loc)).String())
	})

	t.Run("valid time", func(t *testing.T) {
		// Special boundary cases
		assert.Equal(t, "2020-04-01", FromStdTime(time.Date(2020, 5, 23, 0, 0, 0, 0, loc)).String())
		assert.Equal(t, "2020-05-01", FromStdTime(time.Date(2020, 6, 21, 0, 0, 0, 0, loc)).String())

		assert.Equal(t, "2020-06-16", FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc)).String())
		assert.Equal(t, "2023-02-11", FromStdTime(time.Date(2023, 3, 2, 0, 0, 0, 0, loc)).String())
		assert.Equal(t, "2023-02-11", FromStdTime(time.Date(2023, 4, 1, 0, 0, 0, 0, loc)).String())
	})
}

func TestLunar_Gregorian(t *testing.T) {
	t.Run("invalid lunar", func(t *testing.T) {
		assert.Empty(t, new(Lunar).ToGregorian().String())
		assert.Error(t, NewLunar(1800, 1, 1, false).Error)
		assert.Empty(t, NewLunar(1800, 1, 1, false).ToGregorian().String())
	})

	t.Run("invalid timezone", func(t *testing.T) {
		assert.Empty(t, NewLunar(2023, 2, 11, false).ToGregorian("xxx").String())
		assert.Empty(t, NewLunar(3200, 1, 1, true).ToGregorian("xxx").String())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "2023-03-01 16:00:00 +0000 UTC", NewLunar(2023, 2, 11, false).ToGregorian().String())
		assert.Equal(t, "2023-03-31 16:00:00 +0000 UTC", NewLunar(2023, 2, 11, true).ToGregorian().String())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "2023-03-02 00:00:00 +0800 CST", NewLunar(2023, 2, 11, false).ToGregorian("PRC").String())
		assert.Equal(t, "2023-04-01 00:00:00 +0800 CST", NewLunar(2023, 2, 11, true).ToGregorian("PRC").String())
	})
}

func TestLunar_Animal(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Lunar).Animal())
		assert.Empty(t, NewLunar(1800, 1, 1, false).Animal())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "虎", FromStdTime(time.Date(2010, 8, 5, 0, 0, 0, 0, loc)).Animal())
		assert.Equal(t, "兔", FromStdTime(time.Date(2011, 8, 5, 0, 0, 0, 0, loc)).Animal())
		assert.Equal(t, "龙", FromStdTime(time.Date(2012, 8, 5, 0, 0, 0, 0, loc)).Animal())
		assert.Equal(t, "蛇", FromStdTime(time.Date(2013, 8, 5, 0, 0, 0, 0, loc)).Animal())
		assert.Equal(t, "马", FromStdTime(time.Date(2014, 8, 5, 0, 0, 0, 0, loc)).Animal())
		assert.Equal(t, "羊", FromStdTime(time.Date(2015, 8, 5, 0, 0, 0, 0, loc)).Animal())
		assert.Equal(t, "猴", FromStdTime(time.Date(2016, 8, 5, 0, 0, 0, 0, loc)).Animal())
		assert.Equal(t, "鸡", FromStdTime(time.Date(2017, 8, 5, 0, 0, 0, 0, loc)).Animal())
		assert.Equal(t, "狗", FromStdTime(time.Date(2018, 8, 5, 0, 0, 0, 0, loc)).Animal())
		assert.Equal(t, "猪", FromStdTime(time.Date(2019, 8, 5, 0, 0, 0, 0, loc)).Animal())
		assert.Equal(t, "鼠", FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc)).Animal())
		assert.Equal(t, "牛", FromStdTime(time.Date(2021, 8, 5, 0, 0, 0, 0, loc)).Animal())
	})
}

func TestLunar_Festival(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Lunar).Festival())
		assert.Empty(t, NewLunar(1800, 1, 1, false).Festival())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "春节", FromStdTime(time.Date(2021, 2, 12, 0, 0, 0, 0, loc)).Festival())
		assert.Equal(t, "元宵节", FromStdTime(time.Date(2021, 2, 26, 0, 0, 0, 0, loc)).Festival())
		assert.Equal(t, "端午节", FromStdTime(time.Date(2021, 6, 14, 0, 0, 0, 0, loc)).Festival())
		assert.Equal(t, "七夕节", FromStdTime(time.Date(2021, 8, 14, 0, 0, 0, 0, loc)).Festival())
		assert.Equal(t, "中元节", FromStdTime(time.Date(2021, 8, 22, 0, 0, 0, 0, loc)).Festival())
		assert.Equal(t, "中秋节", FromStdTime(time.Date(2021, 9, 21, 0, 0, 0, 0, loc)).Festival())
		assert.Equal(t, "重阳节", FromStdTime(time.Date(2021, 10, 14, 0, 0, 0, 0, loc)).Festival())
		assert.Equal(t, "寒衣节", FromStdTime(time.Date(2021, 11, 5, 0, 0, 0, 0, loc)).Festival())
		assert.Equal(t, "下元节", FromStdTime(time.Date(2021, 11, 19, 0, 0, 0, 0, loc)).Festival())
		assert.Equal(t, "腊八节", FromStdTime(time.Date(2022, 1, 10, 0, 0, 0, 0, loc)).Festival())
	})
}

func TestLunar_Year(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Lunar).Year())
		assert.Empty(t, NewLunar(1800, 1, 1, false).Year())
	})

	t.Run("valid time", func(t *testing.T) {
		l := FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc))
		assert.Equal(t, "2020-06-16", l.String())
		assert.Equal(t, 2020, l.Year())
	})
}

func TestLunar_Month(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Lunar).Month())
		assert.Empty(t, NewLunar(1800, 1, 1, false).Month())
	})

	t.Run("valid time", func(t *testing.T) {
		l := FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc))
		assert.Equal(t, "2020-06-16", l.String())
		assert.Equal(t, 6, l.Month())
	})
}

func TestLunar_LeapMonth(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Lunar).LeapMonth())
		assert.Empty(t, NewLunar(1800, 1, 1, false).LeapMonth())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 4, FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc)).LeapMonth())
		assert.Equal(t, 2, FromStdTime(time.Date(2023, 3, 2, 0, 0, 0, 0, loc)).LeapMonth())
	})
}

func TestLunar_Day(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Lunar).Day())
		assert.Empty(t, NewLunar(1800, 1, 1, false).Day())
	})

	t.Run("valid time", func(t *testing.T) {
		l := FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc))
		assert.Equal(t, "2020-06-16", l.String())
		assert.Equal(t, 16, l.Day())
	})
}

func TestLunar_ToYearString(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Lunar).ToYearString())
		assert.Empty(t, NewLunar(1800, 1, 1, false).ToYearString())
	})

	t.Run("valid time", func(t *testing.T) {
		l := FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc))
		assert.Equal(t, "2020-06-16", l.String())
		assert.Equal(t, "二零二零", l.ToYearString())
	})
}

func TestLunar_ToMonthString(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Lunar).ToMonthString())
		assert.Empty(t, NewLunar(1800, 1, 1, false).ToMonthString())
	})

	t.Run("valid time", func(t *testing.T) {
		l1 := FromStdTime(time.Date(2023, 3, 2, 0, 0, 0, 0, loc))
		assert.Equal(t, "2023-02-11", l1.String())
		assert.Equal(t, "二月", l1.ToMonthString())

		l2 := FromStdTime(time.Date(2023, 4, 1, 0, 0, 0, 0, loc))
		assert.Equal(t, "2023-02-11", l2.String())
		assert.Equal(t, "闰二月", l2.ToMonthString())
	})
}

func TestLunar_ToWeekString(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Lunar).ToWeekString())
		assert.Empty(t, NewLunar(1800, 1, 1, false).ToWeekString())
	})

	t.Run("valid time", func(t *testing.T) {
		l1 := FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc))
		assert.Equal(t, "2020-06-16", l1.String())
		assert.Equal(t, "周二", l1.ToWeekString())
	})
}

func TestLunar_ToDayString(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Lunar).ToDayString())
		assert.Empty(t, NewLunar(1800, 1, 1, false).ToDayString())
	})

	t.Run("valid time", func(t *testing.T) {
		l1 := FromStdTime(time.Date(2020, 5, 1, 0, 0, 0, 0, loc))
		assert.Equal(t, "2020-04-09", l1.String())
		assert.Equal(t, "初九", l1.ToDayString())

		l2 := FromStdTime(time.Date(2020, 6, 1, 0, 0, 0, 0, loc))
		assert.Equal(t, "2020-04-10", l2.String())
		assert.Equal(t, "初十", l2.ToDayString())

		l3 := FromStdTime(time.Date(2020, 8, 1, 0, 0, 0, 0, loc))
		assert.Equal(t, "2020-06-12", l3.String())
		assert.Equal(t, "十二", l3.ToDayString())

		l4 := FromStdTime(time.Date(2021, 1, 3, 0, 0, 0, 0, loc))
		assert.Equal(t, "2020-11-20", l4.String())
		assert.Equal(t, "二十", l4.ToDayString())

		l5 := FromStdTime(time.Date(2021, 1, 5, 0, 0, 0, 0, loc))
		assert.Equal(t, "2020-11-22", l5.String())
		assert.Equal(t, "廿二", l5.ToDayString())

		l6 := FromStdTime(time.Date(2021, 4, 11, 0, 0, 0, 0, loc))
		assert.Equal(t, "2021-02-30", l6.String())
		assert.Equal(t, "三十", l6.ToDayString())
	})
}

func TestLunar_ToDateString(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Lunar).ToDateString())
		assert.Empty(t, NewLunar(1800, 1, 1, false).ToDateString())
	})

	t.Run("valid time", func(t *testing.T) {
		l1 := FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc))
		assert.Equal(t, "2020-06-16", l1.String())
		assert.Equal(t, "二零二零年六月十六", l1.ToDateString())
	})
}

func TestLunar_IsValid(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, new(Lunar).IsValid())
		assert.False(t, NewLunar(1800, 1, 1, true).IsValid())
	})

	t.Run("valid time", func(t *testing.T) {
		l1 := FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc))
		assert.Equal(t, "2020-06-16", l1.String())
		assert.True(t, l1.IsValid())
	})
}

func TestLunar_IsLeapYear(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, new(Lunar).IsLeapYear())
		assert.False(t, NewLunar(1800, 1, 1, true).IsLeapYear())
	})

	t.Run("valid time", func(t *testing.T) {
		l1 := FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc))
		assert.Equal(t, "2020-06-16", l1.String())
		assert.True(t, l1.IsLeapYear())

		l2 := FromStdTime(time.Date(2021, 7, 7, 0, 0, 0, 0, loc))
		assert.Equal(t, "2021-05-28", l2.String())
		assert.False(t, l2.IsLeapYear())
	})
}

func TestLunar_IsLeapMonth(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, new(Lunar).IsLeapMonth())
		assert.False(t, NewLunar(1800, 1, 1, true).IsLeapMonth())
	})

	t.Run("valid time", func(t *testing.T) {
		l1 := FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc))
		assert.Equal(t, 4, l1.LeapMonth())
		assert.False(t, l1.IsLeapMonth())

		l2 := FromStdTime(time.Date(2023, 4, 1, 0, 0, 0, 0, loc))
		assert.Equal(t, 2, l2.LeapMonth())
		assert.True(t, l2.IsLeapMonth())
	})
}

func TestLunar_IsRatYear(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, new(Lunar).IsRatYear())
		assert.False(t, NewLunar(1800, 1, 1, true).IsRatYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc)).IsRatYear())
		assert.False(t, FromStdTime(time.Date(2022, 8, 5, 0, 0, 0, 0, loc)).IsRatYear())
	})
}

func TestLunar_IsOxYear(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, new(Lunar).IsOxYear())
		assert.False(t, NewLunar(1800, 1, 1, true).IsOxYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, FromStdTime(time.Date(2021, 7, 7, 0, 0, 0, 0, loc)).IsOxYear())
		assert.False(t, FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc)).IsOxYear())
	})
}

func TestLunar_IsTigerYear(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, new(Lunar).IsTigerYear())
		assert.False(t, NewLunar(1800, 1, 1, true).IsTigerYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, FromStdTime(time.Date(2022, 8, 5, 0, 0, 0, 0, loc)).IsTigerYear())
		assert.False(t, FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc)).IsTigerYear())
	})
}

func TestLunar_IsRabbitYear(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, new(Lunar).IsRabbitYear())
		assert.False(t, NewLunar(1800, 1, 1, true).IsRabbitYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, FromStdTime(time.Date(2023, 8, 5, 0, 0, 0, 0, loc)).IsRabbitYear())
		assert.False(t, FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc)).IsRabbitYear())
	})
}

func TestLunar_IsDragonYear(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, new(Lunar).IsDragonYear())
		assert.False(t, NewLunar(1800, 1, 1, true).IsDragonYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, FromStdTime(time.Date(2024, 8, 5, 0, 0, 0, 0, loc)).IsDragonYear())
		assert.False(t, FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc)).IsDragonYear())
	})
}

func TestLunar_IsSnakeYear(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, new(Lunar).IsSnakeYear())
		assert.False(t, NewLunar(1800, 1, 1, true).IsSnakeYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, FromStdTime(time.Date(2025, 8, 5, 0, 0, 0, 0, loc)).IsSnakeYear())
		assert.False(t, FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc)).IsSnakeYear())
	})
}

func TestLunar_IsHorseYear(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, new(Lunar).IsHorseYear())
		assert.False(t, NewLunar(1800, 1, 1, true).IsHorseYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, FromStdTime(time.Date(2026, 8, 5, 0, 0, 0, 0, loc)).IsHorseYear())
		assert.False(t, FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc)).IsHorseYear())
	})
}

func TestLunar_IsGoatYear(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, new(Lunar).IsGoatYear())
		assert.False(t, NewLunar(1800, 1, 1, true).IsGoatYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, FromStdTime(time.Date(2027, 8, 5, 0, 0, 0, 0, loc)).IsGoatYear())
		assert.False(t, FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc)).IsGoatYear())
	})
}

func TestLunar_IsMonkeyYear(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, new(Lunar).IsMonkeyYear())
		assert.False(t, NewLunar(1800, 1, 1, true).IsMonkeyYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, FromStdTime(time.Date(2028, 8, 5, 0, 0, 0, 0, loc)).IsMonkeyYear())
		assert.False(t, FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc)).IsMonkeyYear())
	})
}

func TestLunar_IsRoosterYear(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, new(Lunar).IsRoosterYear())
		assert.False(t, NewLunar(1800, 1, 1, true).IsRoosterYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, FromStdTime(time.Date(2029, 8, 5, 0, 0, 0, 0, loc)).IsRoosterYear())
		assert.False(t, FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc)).IsRoosterYear())
	})
}

func TestLunar_IsDogYear(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, new(Lunar).IsDogYear())
		assert.False(t, NewLunar(1800, 1, 1, true).IsDogYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, FromStdTime(time.Date(2030, 8, 5, 0, 0, 0, 0, loc)).IsDogYear())
		assert.False(t, FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc)).IsDogYear())
	})
}

func TestLunar_IsPigYear(t *testing.T) {
	loc, _ := time.LoadLocation("PRC")

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, new(Lunar).IsPigYear())
		assert.False(t, NewLunar(1800, 1, 1, true).IsPigYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, FromStdTime(time.Date(2031, 8, 5, 0, 0, 0, 0, loc)).IsPigYear())
		assert.False(t, FromStdTime(time.Date(2020, 8, 5, 0, 0, 0, 0, loc)).IsPigYear())
	})
}

func TestLunar_AuthorityData(t *testing.T) {
	file, err := os.Open("lunar_test_data.json")
	if err != nil {
		t.Fatalf("failed to open test data file: %v", err)
	}
	defer file.Close()

	type lunarData struct {
		Year        int  `json:"year"`
		Month       int  `json:"month"`
		Day         int  `json:"day"`
		IsLeapMonth bool `json:"isLeapMonth"`
	}
	type gregorianData struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	}
	type testCase struct {
		Description string        `json:"description"`
		Lunar       lunarData     `json:"lunar"`
		Gregorian   gregorianData `json:"gregorian"`
	}

	var cases []testCase
	dec := json.NewDecoder(file)
	if err := dec.Decode(&cases); err != nil {
		t.Fatalf("failed to decode test data: %v", err)
	}

	loc, _ := time.LoadLocation("PRC")

	for _, c := range cases {
		// Lunar to Gregorian conversion
		l := NewLunar(c.Lunar.Year, c.Lunar.Month, c.Lunar.Day, c.Lunar.IsLeapMonth)
		g := l.ToGregorian("PRC")
		if g.Time.IsZero() {
			t.Errorf("[%s] Lunar->Gregorian failed: %+v", c.Description, c.Lunar)
		} else {
			gy, gm, gd := g.Time.In(loc).Date()
			if gy != c.Gregorian.Year || int(gm) != c.Gregorian.Month || gd != c.Gregorian.Day {
				t.Errorf("[%s] Lunar->Gregorian error: expected %04d-%02d-%02d, got %04d-%02d-%02d", c.Description, c.Gregorian.Year, c.Gregorian.Month, c.Gregorian.Day, gy, int(gm), gd)
			}
		}

		// Gregorian to Lunar conversion
		gt := time.Date(c.Gregorian.Year, time.Month(c.Gregorian.Month), c.Gregorian.Day, 0, 0, 0, 0, loc)
		l2 := FromStdTime(gt)
		if l2 == nil {
			t.Errorf("[%s] Gregorian->Lunar failed: %+v", c.Description, c.Gregorian)
		} else {
			if l2.year != c.Lunar.Year || l2.month != c.Lunar.Month || l2.day != c.Lunar.Day || l2.isLeapMonth != c.Lunar.IsLeapMonth {
				t.Errorf("[%s] Gregorian->Lunar error: expected %04d-%02d-%02d leap=%v, got %04d-%02d-%02d leap=%v", c.Description, c.Lunar.Year, c.Lunar.Month, c.Lunar.Day, c.Lunar.IsLeapMonth, l2.year, l2.month, l2.day, l2.isLeapMonth)
			}
		}
	}
}

func TestDayTables(t *testing.T) {
	t.Run("days in year matches the lunar calendar data", func(t *testing.T) {
		for year := minYear; year <= maxYear; year++ {
			assert.Equal(t, countDaysInYear(year), getDaysInYear(year), "year %d", year)
		}
	})

	t.Run("days before year is the running total", func(t *testing.T) {
		total := 0
		for year := minYear; year <= maxYear; year++ {
			assert.Equal(t, total, getOffsetInMonth(year), "year %d", year)
			total += countDaysInYear(year)
		}
	})

	t.Run("tables cover the whole lunar calendar data", func(t *testing.T) {
		assert.Len(t, daysInYearTable, len(years))
		assert.Len(t, daysBeforeYearTable, len(years))
	})
}
