package hebrew

import (
	"encoding/json"
	"fmt"
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
		assert.Equal(t, "5784-10-20", FromStdTime(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)).String())
		assert.Equal(t, "5784-05-01", FromStdTime(time.Date(2024, 8, 5, 12, 0, 0, 0, time.UTC)).String())
		assert.Equal(t, "5786-07-11", FromStdTime(time.Date(2025, 10, 3, 12, 0, 0, 0, time.UTC)).String())
		assert.Equal(t, "5784-07-01", FromStdTime(time.Date(2023, 9, 16, 12, 0, 0, 0, time.UTC)).String())
	})

	t.Run("rosh hashanah anchors", func(t *testing.T) {
		// 1 Tishri (month 7) of each Hebrew year is a fixed historical date.
		assert.Equal(t, "5780-07-01", FromStdTime(time.Date(2019, 9, 30, 12, 0, 0, 0, time.UTC)).String())
		assert.Equal(t, "5784-07-01", FromStdTime(time.Date(2023, 9, 16, 12, 0, 0, 0, time.UTC)).String())
		assert.Equal(t, "5785-07-01", FromStdTime(time.Date(2024, 10, 3, 12, 0, 0, 0, time.UTC)).String())
		// 1 Nisan (month 1) 5780.
		assert.Equal(t, "5780-01-01", FromStdTime(time.Date(2020, 3, 26, 12, 0, 0, 0, time.UTC)).String())
	})
}

func TestHebrew_Gregorian(t *testing.T) {
	t.Run("invalid hebrew", func(t *testing.T) {
		assert.NotEmpty(t, new(Hebrew).ToGregorian().String())
		assert.NotEmpty(t, NewHebrew(10000, 1, 1).ToGregorian().String())
	})

	t.Run("invalid timezone", func(t *testing.T) {
		g := NewHebrew(5784, 1, 1).ToGregorian("xxx")
		assert.Error(t, g.Error)
		assert.Empty(t, g.String())
	})

	t.Run("without timezone", func(t *testing.T) {
		fmtDate := func(h *Hebrew) string { return h.ToGregorian("UTC").Time.Format("2006-01-02") }
		assert.Equal(t, "2024-04-09", fmtDate(NewHebrew(5784, 1, 1)))
		assert.Equal(t, "2024-07-21", fmtDate(NewHebrew(5784, 4, 15)))
		assert.Equal(t, "2024-01-13", fmtDate(NewHebrew(5784, 11, 3)))
		assert.Equal(t, "2025-07-11", fmtDate(NewHebrew(5785, 4, 15)))
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.NotEmpty(t, NewHebrew(5784, 1, 1).ToGregorian("PRC").String())
		assert.NotEmpty(t, NewHebrew(5784, 4, 15).ToGregorian("PRC").String())
		assert.NotEmpty(t, NewHebrew(5784, 11, 3).ToGregorian("PRC").String())
		assert.NotEmpty(t, NewHebrew(5785, 4, 15).ToGregorian("PRC").String())
	})
}

func TestHebrew_Year(t *testing.T) {
	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Hebrew).Year())
		assert.Equal(t, 0, NewHebrew(10000, 1, 1).Year())
	})

	t.Run("nil hebrew", func(t *testing.T) {
		var h *Hebrew
		assert.Equal(t, 0, h.Year())
	})

	t.Run("valid time", func(t *testing.T) {
		h := FromStdTime(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC))
		assert.NotEmpty(t, h.String())
		assert.True(t, h.Year() > 0)
	})
}

func TestHebrew_Month(t *testing.T) {
	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Hebrew).Month())
		assert.Equal(t, 0, NewHebrew(10000, 1, 1).Month())
	})

	t.Run("nil hebrew", func(t *testing.T) {
		var h *Hebrew
		assert.Equal(t, 0, h.Month())
	})

	t.Run("valid time", func(t *testing.T) {
		h := FromStdTime(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC))
		assert.NotEmpty(t, h.String())
		assert.True(t, h.Month() > 0 && h.Month() <= 13)
	})
}

func TestHebrew_Day(t *testing.T) {
	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Hebrew).Day())
		assert.Equal(t, 0, NewHebrew(10000, 1, 1).Day())
	})

	t.Run("nil hebrew", func(t *testing.T) {
		var h *Hebrew
		assert.Equal(t, 0, h.Day())
	})

	t.Run("valid time", func(t *testing.T) {
		h := FromStdTime(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC))
		assert.NotEmpty(t, h.String())
		assert.True(t, h.Day() > 0 && h.Day() <= 30)
	})
}

func TestHebrew_ToMonthString(t *testing.T) {
	t.Run("nil hebrew", func(t *testing.T) {
		hebrew := new(Hebrew)
		hebrew = nil
		assert.Empty(t, hebrew.ToMonthString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Hebrew).ToMonthString())
		assert.Empty(t, NewHebrew(5780, 0, 1).ToMonthString())
		assert.Empty(t, NewHebrew(5780, 14, 1).ToMonthString())
	})

	t.Run("invalid locale", func(t *testing.T) {
		h := NewHebrew(5780, 11, 6)
		assert.Empty(t, h.ToMonthString("xxx"))
	})

	t.Run("valid time", func(t *testing.T) {
		h := NewHebrew(5780, 11, 6)
		assert.Equal(t, "5780-11-06", h.String())
		assert.Equal(t, "Shevat", h.ToMonthString(EnLocale))
		assert.Equal(t, "שבט", h.ToMonthString(HeLocale))
	})
}

func TestHebrew_ToWeekString(t *testing.T) {
	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, new(Hebrew).ToWeekString())
	})

	t.Run("nil hebrew", func(t *testing.T) {
		var h *Hebrew
		assert.Empty(t, h.ToWeekString())
		assert.Empty(t, h.ToWeekString(EnLocale))
		assert.Empty(t, h.ToWeekString(HeLocale))
	})

	t.Run("invalid locale", func(t *testing.T) {
		h := NewHebrew(5780, 10, 7)
		assert.Empty(t, h.ToWeekString("xxx"))
		assert.Empty(t, h.ToWeekString(Locale("invalid")))
		assert.Empty(t, h.ToWeekString(Locale("")))
		assert.Empty(t, h.ToWeekString(Locale("en-US")))
		assert.Empty(t, h.ToWeekString(Locale("he-IL")))
	})

	t.Run("valid time", func(t *testing.T) {
		h := NewHebrew(5780, 10, 7)
		assert.Equal(t, "5780-10-07", h.String())
		assert.Equal(t, "Saturday", h.ToWeekString(EnLocale))
		assert.Equal(t, "שבת", h.ToWeekString(HeLocale))
	})

	t.Run("all weekdays with EnLocale", func(t *testing.T) {
		expectedEnWeeks := []string{"Thursday", "Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday"}
		for i := 1; i <= 7; i++ {
			h := NewHebrew(5780, 1, i)
			assert.Equal(t, expectedEnWeeks[i-1], h.ToWeekString(EnLocale), "Failed for date 5780-1-%d", i)
		}
	})

	t.Run("all weekdays with HeLocale", func(t *testing.T) {
		expectedHeWeeks := []string{"חמישי", "שישי", "שבת", "ראשון", "שני", "שלישי", "רביעי"}
		for i := 1; i <= 7; i++ {
			h := NewHebrew(5780, 1, i)
			assert.Equal(t, expectedHeWeeks[i-1], h.ToWeekString(HeLocale), "Failed for date 5780-1-%d", i)
		}
	})

	t.Run("default locale", func(t *testing.T) {
		h := NewHebrew(5780, 10, 7)
		assert.Equal(t, "Saturday", h.ToWeekString()) // default is EnLocale
	})

	t.Run("verify actual dates", func(t *testing.T) {
		// Let's verify what the actual weekdays are for these dates
		for i := 1; i <= 7; i++ {
			h := NewHebrew(5780, 1, i)
			weekday := h.ToWeekString(EnLocale)
			t.Logf("5780-1-%d -> %s", i, weekday)
		}
	})
}

func TestHebrew_IsLeapYear(t *testing.T) {
	t.Run("invalid hebrew", func(t *testing.T) {
		assert.False(t, new(Hebrew).IsLeapYear())
		assert.False(t, NewHebrew(10000, 1, 1).IsLeapYear())
	})

	t.Run("nil hebrew", func(t *testing.T) {
		var h *Hebrew
		assert.False(t, h.IsLeapYear())
	})

	t.Run("leap years", func(t *testing.T) {
		assert.True(t, NewHebrew(5784, 1, 1).IsLeapYear())
		assert.True(t, NewHebrew(5787, 1, 1).IsLeapYear())
		assert.True(t, NewHebrew(5790, 1, 1).IsLeapYear())
	})

	t.Run("non-leap years", func(t *testing.T) {
		assert.False(t, NewHebrew(5785, 1, 1).IsLeapYear())
		assert.False(t, NewHebrew(5786, 1, 1).IsLeapYear())
		assert.False(t, NewHebrew(5788, 1, 1).IsLeapYear())
	})
}

func TestHebrew_IsValid(t *testing.T) {
	t.Run("invalid_hebrew", func(t *testing.T) {
		// Test invalid year ranges
		assert.False(t, NewHebrew(0, 1, 1).IsValid())     // Year 0 is invalid
		assert.False(t, NewHebrew(10000, 1, 1).IsValid()) // Year 10000 is invalid

		// Test invalid month ranges
		assert.False(t, NewHebrew(5780, 0, 1).IsValid())  // Month 0 is invalid
		assert.False(t, NewHebrew(5780, 14, 1).IsValid()) // Month 14 is invalid

		// Test invalid day ranges
		assert.False(t, NewHebrew(5780, 1, 0).IsValid())  // Day 0 is invalid
		assert.False(t, NewHebrew(5780, 1, 32).IsValid()) // Day 32 is invalid

		// Test invalid day for specific months
		assert.False(t, NewHebrew(5780, 2, 30).IsValid())  // Month 2 has max 29 days
		assert.False(t, NewHebrew(5780, 4, 30).IsValid())  // Month 4 has max 29 days
		assert.False(t, NewHebrew(5780, 6, 30).IsValid())  // Month 6 has max 29 days
		assert.False(t, NewHebrew(5780, 10, 30).IsValid()) // Month 10 has max 29 days
		assert.False(t, NewHebrew(5780, 13, 30).IsValid()) // Month 13 has max 29 days
	})

	t.Run("nil_hebrew", func(t *testing.T) {
		var h *Hebrew
		assert.False(t, h.IsValid())
	})

	t.Run("valid_hebrew", func(t *testing.T) {
		// Test valid dates
		assert.True(t, NewHebrew(5780, 1, 1).IsValid())
		assert.True(t, NewHebrew(5780, 1, 30).IsValid())
		assert.True(t, NewHebrew(5780, 2, 29).IsValid())
		assert.True(t, NewHebrew(5780, 3, 30).IsValid())
		assert.True(t, NewHebrew(5780, 7, 1).IsValid())
		assert.True(t, NewHebrew(5780, 12, 29).IsValid()) // Month 12 has 29 days in non-leap year
	})

	t.Run("boundary_values", func(t *testing.T) {
		// Test boundary years
		assert.True(t, NewHebrew(1, 1, 1).IsValid())
		assert.True(t, NewHebrew(9999, 12, 29).IsValid())

		// Test boundary months
		assert.True(t, NewHebrew(5780, 1, 1).IsValid())
		assert.True(t, NewHebrew(5780, 12, 29).IsValid()) // Month 12 has 29 days in non-leap year

		// Test boundary days
		assert.True(t, NewHebrew(5780, 1, 1).IsValid())
		assert.True(t, NewHebrew(5780, 1, 30).IsValid())
	})

	t.Run("leap_year_month_13", func(t *testing.T) {
		// 5784 should have month 13 (leap year)
		assert.True(t, NewHebrew(5784, 13, 1).IsValid())
		// 5785 is not a leap year, so month 13 should be invalid
		assert.False(t, NewHebrew(5785, 13, 1).IsValid()) // 5785不是闰年，没有第13个月
		// 年份1 is not a leap year, so month 13 should be invalid
		assert.False(t, NewHebrew(1, 13, 1).IsValid()) // 年份1不是闰年，没有第13个月
		// 年份9999 is not a leap year, so month 13 should be invalid
		assert.False(t, NewHebrew(9999, 13, 1).IsValid()) // 年份9999不是闰年，没有第13个月
	})

	t.Run("from_std_time", func(t *testing.T) {
		// Test FromStdTime with valid time
		testTime := time.Date(2020, 3, 26, 12, 0, 0, 0, time.UTC)
		h := FromStdTime(testTime)
		assert.True(t, h.IsValid())

		// Test FromStdTime with zero time
		h2 := FromStdTime(time.Time{})
		assert.False(t, h2.IsValid())
	})

	t.Run("error_handling", func(t *testing.T) {
		// Test Hebrew with error
		h := &Hebrew{year: 5780, month: 1, day: 1, Error: fmt.Errorf("test error")}
		assert.False(t, h.IsValid())
	})

	t.Run("month_validation_edge_cases", func(t *testing.T) {
		// Test months that don't exist in the year
		assert.False(t, NewHebrew(5785, 13, 1).IsValid()) // Non-leap year, no month 13
		assert.True(t, NewHebrew(5784, 13, 1).IsValid())  // Leap year, has month 13

		// Test edge cases for month validation
		assert.False(t, NewHebrew(5780, 14, 1).IsValid()) // Month 14 doesn't exist
		assert.False(t, NewHebrew(5780, -1, 1).IsValid()) // Negative month
	})

	t.Run("day_validation_edge_cases", func(t *testing.T) {
		// Test days that don't exist in specific months
		assert.False(t, NewHebrew(5780, 2, 30).IsValid())  // Month 2 has max 29 days
		assert.False(t, NewHebrew(5780, 4, 30).IsValid())  // Month 4 has max 29 days
		assert.False(t, NewHebrew(5780, 6, 30).IsValid())  // Month 6 has max 29 days
		assert.False(t, NewHebrew(5780, 10, 30).IsValid()) // Month 10 has max 29 days
		assert.False(t, NewHebrew(5780, 12, 30).IsValid()) // Month 12 has 29 days in non-leap year
		assert.False(t, NewHebrew(5780, 13, 30).IsValid()) // Month 13 has max 29 days

		// Test negative days
		assert.False(t, NewHebrew(5780, 1, -1).IsValid())
		assert.False(t, NewHebrew(5780, 1, 0).IsValid())
	})

	t.Run("year_validation_edge_cases", func(t *testing.T) {
		// Test year boundaries
		assert.False(t, NewHebrew(0, 1, 1).IsValid())     // Year 0 is invalid
		assert.True(t, NewHebrew(1, 1, 1).IsValid())      // Year 1 is valid
		assert.True(t, NewHebrew(9999, 12, 29).IsValid()) // Year 9999 is valid
		assert.False(t, NewHebrew(10000, 1, 1).IsValid()) // Year 10000 is invalid

		// Test negative years
		assert.False(t, NewHebrew(-1, 1, 1).IsValid())
	})
}

func TestHebrew_String(t *testing.T) {
	t.Run("invalid hebrew", func(t *testing.T) {
		assert.Empty(t, new(Hebrew).String())
		assert.Empty(t, NewHebrew(10000, 1, 1).String())
	})

	t.Run("nil hebrew", func(t *testing.T) {
		var h *Hebrew
		assert.Equal(t, "", h.String())
	})

	t.Run("valid hebrew", func(t *testing.T) {
		assert.Equal(t, "5784-01-01", NewHebrew(5784, 1, 1).String())
		assert.Equal(t, "5784-12-30", NewHebrew(5784, 12, 30).String())
		assert.Equal(t, "0001-01-01", NewHebrew(1, 1, 1).String())
	})
}

func TestHebrew_NewHebrew(t *testing.T) {
	t.Run("valid cases", func(t *testing.T) {
		assert.NotNil(t, NewHebrew(1, 1, 1))
		assert.NotNil(t, NewHebrew(9999, 12, 30))
		assert.NotNil(t, NewHebrew(0, 1, 1))
		assert.NotNil(t, NewHebrew(10000, 1, 1))
	})
}

func TestHebrew_ToGregorian(t *testing.T) {
	t.Run("nil hebrew", func(t *testing.T) {
		var h *Hebrew
		g := h.ToGregorian()
		assert.NotNil(t, g)
		assert.True(t, g.Time.IsZero())
	})

	t.Run("valid cases", func(t *testing.T) {
		h := NewHebrew(5780, 1, 1)
		g := h.ToGregorian()
		assert.NotNil(t, g)
		assert.False(t, g.Time.IsZero())

		g = h.ToGregorian("UTC")
		assert.NotNil(t, g)
		assert.False(t, g.Time.IsZero())
	})

	t.Run("invalid timezone", func(t *testing.T) {
		h := NewHebrew(5780, 1, 1)
		g := h.ToGregorian("Invalid/Timezone")
		assert.NotNil(t, g)
		assert.NotNil(t, g.Error)
	})

	t.Run("empty timezone", func(t *testing.T) {
		h := NewHebrew(5780, 1, 1)
		g := h.ToGregorian()
		assert.NotNil(t, g)
		assert.False(t, g.Time.IsZero())
		assert.Nil(t, g.Error)
	})
}

func TestHebrew_YearMonthDay(t *testing.T) {
	t.Run("valid cases", func(t *testing.T) {
		h := NewHebrew(5780, 1, 1)
		assert.Equal(t, 5780, h.Year())
		assert.Equal(t, 1, h.Month())
		assert.Equal(t, 1, h.Day())

		h = NewHebrew(0, 1, 1)
		assert.Equal(t, 0, h.Year())
		assert.Equal(t, 0, h.Month())
		assert.Equal(t, 0, h.Day())

		h = NewHebrew(10000, 1, 1)
		assert.Equal(t, 0, h.Year())
		assert.Equal(t, 0, h.Month())
		assert.Equal(t, 0, h.Day())
	})
}

func TestJdn2gregorian(t *testing.T) {
	t.Run("authoritative JDN to Gregorian comparison", func(t *testing.T) {
		cases := []struct {
			jdn   int
			year  int
			month int
			day   int
		}{
			{1721426, 1, 1, 1},    // Rata Die day 1
			{2451545, 2000, 1, 1}, // J2000.0 epoch (noon)
			{2459580, 2021, 12, 31},
			{2459581, 2022, 1, 1},
			{2460100, 2023, 6, 4},
			{2460141, 2023, 7, 15},
			{2488434, 2100, 12, 31},
		}
		for _, c := range cases {
			y, m, d := jdn2gregorian(c.jdn)
			assert.Equal(t, c.year, y, "JDN %d year", c.jdn)
			assert.Equal(t, c.month, m, "JDN %d month", c.jdn)
			assert.Equal(t, c.day, d, "JDN %d day", c.jdn)
			// Gregorian -> JDN must invert JDN -> Gregorian.
			assert.Equal(t, c.jdn, gregorian2jdn(y, m, d), "round-trip JDN %d", c.jdn)
		}
	})
}

func TestHebrew_AuthorityData(t *testing.T) {
	// Load test data from JSON file
	data, err := os.ReadFile("hebrew_test_data.json")
	if err != nil {
		t.Skipf("Test data file not found: %v", err)
	}

	var testCases []struct {
		Description string `json:"description"`
		Hebrew      struct {
			Year  int `json:"year"`
			Month int `json:"month"`
			Day   int `json:"day"`
		} `json:"hebrew"`
		Gregorian struct {
			Year   int `json:"year"`
			Month  int `json:"month"`
			Day    int `json:"day"`
			Hour   int `json:"hour"`
			Minute int `json:"minute"`
			Second int `json:"second"`
		} `json:"gregorian"`
	}

	if err := json.Unmarshal(data, &testCases); err != nil {
		t.Fatalf("Failed to parse test data: %v", err)
	}

	t.Logf("Loaded %d test cases from authority data", len(testCases))

	// Every record carries the authoritative Gregorian date for a Hebrew date;
	// assert both conversion directions match it exactly.
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Case_%d_%s", idx+1, tc.Description), func(t *testing.T) {
			wantGreg := fmt.Sprintf("%04d-%02d-%02d", tc.Gregorian.Year, tc.Gregorian.Month, tc.Gregorian.Day)

			// Hebrew -> Gregorian.
			g := NewHebrew(tc.Hebrew.Year, tc.Hebrew.Month, tc.Hebrew.Day).ToGregorian("UTC")
			assert.Equal(t, wantGreg, g.Time.Format("2006-01-02"),
				"Hebrew %d-%d-%d -> Gregorian", tc.Hebrew.Year, tc.Hebrew.Month, tc.Hebrew.Day)

			// Gregorian -> Hebrew.
			h2 := FromStdTime(time.Date(tc.Gregorian.Year, time.Month(tc.Gregorian.Month), tc.Gregorian.Day, tc.Gregorian.Hour, tc.Gregorian.Minute, tc.Gregorian.Second, 0, time.UTC))
			assert.Equal(t, tc.Hebrew.Year, h2.Year(), "Gregorian %s -> Hebrew year", wantGreg)
			assert.Equal(t, tc.Hebrew.Month, h2.Month(), "Gregorian %s -> Hebrew month", wantGreg)
			assert.Equal(t, tc.Hebrew.Day, h2.Day(), "Gregorian %s -> Hebrew day", wantGreg)
		})
	}
}

func TestHebrew_jdn2hebrew(t *testing.T) {

	t.Run("known_jdn_values", func(t *testing.T) {
		// Julian Day Number (at noon) -> Hebrew date, verified against pyluach
		// and convertdate.hebrew.
		cases := []struct {
			jdn   int
			year  int
			month int
			day   int
		}{
			{1721426, 3761, 10, 18}, // 0001-01-01
			{2458935, 5780, 1, 1},   // 2020-03-26, 1 Nisan
			{2459580, 5782, 10, 27}, // 2021-12-31
			{2460204, 5784, 7, 1},   // 2023-09-16, Rosh Hashanah
		}
		for _, c := range cases {
			year, month, day := jdn2hebrew(c.jdn)
			assert.Equal(t, c.year, year, "JDN %d year", c.jdn)
			assert.Equal(t, c.month, month, "JDN %d month", c.jdn)
			assert.Equal(t, c.day, day, "JDN %d day", c.jdn)
		}
	})

	t.Run("first_day_of_every_month", func(t *testing.T) {
		// jdn2hebrew must invert hebrew2jdn for the first day of every month of
		// both a common (5785) and a leap (5784) year.
		for _, y := range []int{5784, 5785} {
			for m := 1; m <= getMonthsInYear(y); m++ {
				year, month, day := jdn2hebrew(hebrew2jdn(y, m, 1))
				assert.Equal(t, [3]int{y, m, 1}, [3]int{year, month, day}, "first day of %d-%d", y, m)
			}
		}
	})
}

func TestHebrew_getJDNInYear(t *testing.T) {
	// Rosh Hashanah (Tishri 1) Julian Day Number at noon, verified against
	// convertdate.hebrew for years exercising each postponement rule.
	cases := []struct {
		year int
		jdn  int
	}{
		{1, 347998},
		{5765, 2453265},
		{5766, 2453648},
		{5767, 2454002},
		{5768, 2454357},
		{5769, 2454740},
		{5770, 2455094},
		{5780, 2458757},
		{5784, 2460204},
		{9999, 3999723},
	}
	for _, c := range cases {
		assert.Equal(t, c.jdn, getJDNInYear(c.year), "Rosh Hashanah JDN for year %d", c.year)
	}
}
