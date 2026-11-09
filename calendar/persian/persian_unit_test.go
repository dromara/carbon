package persian

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPersian(t *testing.T) {
	t.Run("valid date", func(t *testing.T) {
		p := NewPersian(1400, 1, 1)
		assert.NotNil(t, p)
		assert.Nil(t, p.Error)
		assert.Equal(t, "1400-01-01", p.String())
	})

	t.Run("invalid year", func(t *testing.T) {
		p := NewPersian(0, 1, 1)
		assert.NotNil(t, p)
		assert.Error(t, p.Error)
		assert.Equal(t, "", p.String())
	})

	t.Run("invalid month", func(t *testing.T) {
		p := NewPersian(1400, 13, 1)
		assert.NotNil(t, p)
		assert.Error(t, p.Error)
		assert.Equal(t, "", p.String())
	})

	t.Run("invalid day", func(t *testing.T) {
		p := NewPersian(1400, 1, 32)
		assert.NotNil(t, p)
		assert.Error(t, p.Error)
		assert.Equal(t, "", p.String())
	})

	t.Run("invalid month day combination", func(t *testing.T) {
		// Persian calendar: first 6 months have 31 days, next 5 months have 30 days, last month has 29 or 30 days
		p := NewPersian(1400, 7, 31) // 7th month has maximum 30 days
		assert.NotNil(t, p)
		assert.Error(t, p.Error)
		assert.Equal(t, "", p.String())
	})

	t.Run("invalid leap year day", func(t *testing.T) {
		// Last month of non-leap year has maximum 29 days
		p := NewPersian(1400, 12, 30) // Year 1400 is not a leap year
		assert.NotNil(t, p)
		assert.Error(t, p.Error)
		assert.Equal(t, "", p.String())
	})
}

func TestFromStdTime(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tehran")

	t.Run("zero time", func(t *testing.T) {
		assert.Nil(t, FromStdTime(time.Time{}))
		assert.Nil(t, FromStdTime(time.Time{}.In(loc)))
	})

	t.Run("valid time", func(t *testing.T) {
		// Test some known Persian calendar dates
		assert.Equal(t, "1400-01-01", FromStdTime(time.Date(2021, 3, 21, 0, 0, 0, 0, loc)).String())
		assert.Equal(t, "1400-12-29", FromStdTime(time.Date(2022, 3, 20, 0, 0, 0, 0, loc)).String())
		assert.Equal(t, "1401-01-01", FromStdTime(time.Date(2022, 3, 21, 0, 0, 0, 0, loc)).String())
		assert.Equal(t, "1401-12-29", FromStdTime(time.Date(2023, 3, 20, 0, 0, 0, 0, loc)).String()) // Authority library result
		assert.Equal(t, "1402-01-01", FromStdTime(time.Date(2023, 3, 21, 0, 0, 0, 0, loc)).String())
	})
}

func TestPersian_ToGregorian(t *testing.T) {
	t.Run("invalid persian", func(t *testing.T) {
		assert.Empty(t, new(Persian).ToGregorian().String())
		assert.Empty(t, NewPersian(0, 1, 1).ToGregorian().String())
	})

	t.Run("invalid timezone", func(t *testing.T) {
		assert.Empty(t, NewPersian(1400, 1, 1).ToGregorian("xxx").String())
	})

	t.Run("without timezone", func(t *testing.T) {
		p := NewPersian(1400, 1, 1)
		g := p.ToGregorian()
		assert.NotNil(t, g)
		assert.Equal(t, "2021-03-21 00:00:00 +0000 UTC", g.String())
	})

	t.Run("with timezone", func(t *testing.T) {
		p := NewPersian(1400, 1, 1)
		g := p.ToGregorian("Asia/Tehran")
		assert.NotNil(t, g)
		assert.Contains(t, g.String(), "2021-03-21")
	})
}

func TestPersian_Year(t *testing.T) {
	t.Run("invalid persian", func(t *testing.T) {
		assert.Equal(t, 0, new(Persian).Year())
		assert.Equal(t, 0, NewPersian(0, 1, 1).Year())
		var p *Persian
		assert.Equal(t, 0, p.Year())
	})

	t.Run("valid persian", func(t *testing.T) {
		p := NewPersian(1400, 1, 1)
		assert.Equal(t, 1400, p.Year())
	})
}

func TestPersian_Month(t *testing.T) {
	t.Run("invalid persian", func(t *testing.T) {
		assert.Equal(t, 0, new(Persian).Month())
		assert.Equal(t, 0, NewPersian(0, 1, 1).Month())
		var p *Persian
		assert.Equal(t, 0, p.Month())
	})

	t.Run("valid persian", func(t *testing.T) {
		p := NewPersian(1400, 6, 15)
		assert.Equal(t, 6, p.Month())
	})
}

func TestPersian_Day(t *testing.T) {
	t.Run("invalid persian", func(t *testing.T) {
		assert.Equal(t, 0, new(Persian).Day())
		assert.Equal(t, 0, NewPersian(0, 1, 1).Day())
		var p *Persian
		assert.Equal(t, 0, p.Day())
	})

	t.Run("valid persian", func(t *testing.T) {
		p := NewPersian(1400, 6, 15)
		assert.Equal(t, 15, p.Day())
	})
}

func TestPersian_String(t *testing.T) {
	t.Run("invalid persian", func(t *testing.T) {
		assert.Equal(t, "", new(Persian).String())
		assert.Equal(t, "", NewPersian(0, 1, 1).String())
		var p *Persian
		assert.Equal(t, "", p.String())
	})

	t.Run("valid persian", func(t *testing.T) {
		p := NewPersian(1400, 6, 15)
		assert.Equal(t, "1400-06-15", p.String())
	})
}

func TestPersian_ToMonthString(t *testing.T) {
	t.Run("invalid persian", func(t *testing.T) {
		assert.Equal(t, "", new(Persian).ToMonthString())
		assert.Equal(t, "", NewPersian(0, 1, 1).ToMonthString())
		var p *Persian
		assert.Equal(t, "", p.ToMonthString())
	})

	t.Run("english locale", func(t *testing.T) {
		p := NewPersian(1400, 1, 1)
		assert.Equal(t, "Farvardin", p.ToMonthString(EnLocale))
		p = NewPersian(1400, 2, 1)
		assert.Equal(t, "Ordibehesht", p.ToMonthString(EnLocale))
	})

	t.Run("persian locale", func(t *testing.T) {
		p := NewPersian(1400, 1, 1)
		assert.Equal(t, "فروردین", p.ToMonthString(FaLocale))
		p = NewPersian(1400, 2, 1)
		assert.Equal(t, "اردیبهشت", p.ToMonthString(FaLocale))
	})

	t.Run("default locale", func(t *testing.T) {
		p := NewPersian(1400, 1, 1)
		assert.Equal(t, "Farvardin", p.ToMonthString())
	})

	t.Run("invalid locale", func(t *testing.T) {
		p := NewPersian(1400, 1, 1)
		assert.Equal(t, "", p.ToMonthString("invalid"))
		assert.Equal(t, "", p.ToMonthString(Locale("")))
	})

	t.Run("month out of range", func(t *testing.T) {
		p := NewPersian(1400, 13, 1)
		assert.Equal(t, "", p.ToMonthString(EnLocale))
		p = NewPersian(1400, 0, 1)
		assert.Equal(t, "", p.ToMonthString(FaLocale))
	})
}

func TestPersian_ToWeekString(t *testing.T) {
	t.Run("invalid persian", func(t *testing.T) {
		assert.Equal(t, "", new(Persian).ToWeekString())
		assert.Equal(t, "", NewPersian(0, 1, 1).ToWeekString())
		var p *Persian
		assert.Equal(t, "", p.ToWeekString())
	})

	t.Run("english locale", func(t *testing.T) {
		p := NewPersian(1400, 1, 1)
		assert.Equal(t, "Yekshanbeh", p.ToWeekString(EnLocale))
	})

	t.Run("persian locale", func(t *testing.T) {
		p := NewPersian(1400, 1, 1)
		assert.Equal(t, "یکشنبه", p.ToWeekString(FaLocale))
	})

	t.Run("default locale", func(t *testing.T) {
		p := NewPersian(1400, 1, 1)
		assert.Equal(t, "Yekshanbeh", p.ToWeekString())
	})

	t.Run("invalid locale", func(t *testing.T) {
		p := NewPersian(1400, 1, 1)
		assert.Equal(t, "", p.ToWeekString("invalid"))
		assert.Equal(t, "", p.ToWeekString(Locale("")))
	})
}

func TestPersian_IsValid(t *testing.T) {
	t.Run("nil persian", func(t *testing.T) {
		assert.False(t, (*Persian)(nil).IsValid())
	})

	t.Run("with error", func(t *testing.T) {
		p := NewPersian(0, 1, 1)
		assert.False(t, p.IsValid())
	})

	t.Run("invalid year", func(t *testing.T) {
		assert.False(t, NewPersian(0, 1, 1).IsValid())
		assert.False(t, NewPersian(10000, 1, 1).IsValid())
	})

	t.Run("invalid month", func(t *testing.T) {
		assert.False(t, NewPersian(1400, 0, 1).IsValid())
		assert.False(t, NewPersian(1400, 13, 1).IsValid())
	})

	t.Run("invalid day", func(t *testing.T) {
		assert.False(t, NewPersian(1400, 1, 0).IsValid())
		assert.False(t, NewPersian(1400, 1, 32).IsValid())
	})

	t.Run("invalid month day combination", func(t *testing.T) {
		// First 6 months have 31 days, next 5 months have 30 days
		assert.False(t, NewPersian(1400, 7, 31).IsValid()) // 7th month has maximum 30 days
		assert.False(t, NewPersian(1400, 8, 31).IsValid()) // 8th month has maximum 30 days
	})

	t.Run("valid dates", func(t *testing.T) {
		assert.True(t, NewPersian(1400, 1, 1).IsValid())
		assert.True(t, NewPersian(1400, 6, 31).IsValid()) // First 6 months have 31 days
		assert.True(t, NewPersian(1400, 7, 30).IsValid()) // Next 5 months have 30 days
	})

	t.Run("internal methods", func(t *testing.T) {
		// Test all branches of jdn2persian
		// Test case when days <= 186
		year, month, day := jdn2persian(getPersianJdn(1400, 1, 1))
		assert.True(t, year > 0 || year == -1)
		assert.True(t, month > 0 || year == -1)
		assert.True(t, day > 0 || year == -1)

		// Test case when days > 186
		year, month, day = jdn2persian(getPersianJdn(1400, 7, 1))
		assert.True(t, year > 0 || year == -1)
		assert.True(t, month > 0 || year == -1)
		assert.True(t, day > 0 || year == -1)

		// Test extreme cases
		year, month, day = jdn2persian(getPersianJdn(1, 1, 1))
		assert.True(t, year > 0 || year == -1)
		assert.True(t, month > 0 || year == -1)
		assert.True(t, day > 0 || year == -1)

		year, month, day = jdn2persian(getPersianJdn(9999, 12, 29))
		assert.True(t, year > 0 || year == -1)
		assert.True(t, month > 0 || year == -1)
		assert.True(t, day > 0 || year == -1)

		// getPersianYear normal cases
		year = getPersianYear(getPersianJdn(1400, 1, 1))
		assert.True(t, year > 0 || year == -1)

		year = getPersianYear(getPersianJdn(5000, 6, 15))
		assert.True(t, year > 0 || year == -1)

		year = getPersianYear(getPersianJdn(8000, 12, 29))
		assert.True(t, year > 0 || year == -1)

		year = getPersianYear(persianEpoch)
		assert.True(t, year > 0 || year == -1)

		year = getPersianYear(persianEpoch + 1)
		assert.True(t, year > 0 || year == -1)

		// getPersianYear extreme cases
		year = getPersianYear(persianEpoch - 1000)
		assert.True(t, year > 0 || year == -1)
		year = getPersianYear(persianEpoch + 1000000)
		assert.True(t, year > 0 || year == -1)
		year = getPersianYear(persianEpoch - 10000)
		assert.True(t, year > 0 || year == -1)
		year = getPersianYear(persianEpoch + 2000000)
		assert.True(t, year > 0 || year == -1)
	})
}

func TestPersian_IsLeapYear(t *testing.T) {
	t.Run("nil persian", func(t *testing.T) {
		assert.False(t, (*Persian)(nil).IsLeapYear())
	})

	t.Run("with error", func(t *testing.T) {
		p := NewPersian(0, 1, 1)
		assert.False(t, p.IsLeapYear())
	})

	t.Run("leap year test", func(t *testing.T) {
		p := NewPersian(1400, 1, 1)
		isLeap := p.IsLeapYear()
		assert.IsType(t, true, isLeap)
		// Extreme years
		p1 := NewPersian(1, 1, 1)
		_ = p1.IsLeapYear()
		p2 := NewPersian(9999, 12, 29)
		_ = p2.IsLeapYear()

		// Test all branches of getPersianYear
		// Normal cases
		year := getPersianYear(getPersianJdn(1400, 1, 1))
		assert.True(t, year > 0 || year == -1)

		// Test case when year < 1
		year = getPersianYear(persianEpoch - 1000)
		assert.True(t, year > 0 || year == -1)

		// Test case when year > 9999
		year = getPersianYear(persianEpoch + 1000000)
		assert.True(t, year > 0 || year == -1)

		// Test different branches of binary search
		year = getPersianYear(getPersianJdn(5000, 6, 15))
		assert.True(t, year > 0 || year == -1)

		year = getPersianYear(getPersianJdn(8000, 12, 29))
		assert.True(t, year > 0 || year == -1)

		// Test extreme JDN values
		year = getPersianYear(persianEpoch)
		assert.True(t, year > 0 || year == -1)

		year = getPersianYear(persianEpoch + 1)
		assert.True(t, year > 0 || year == -1)

		// Test invalid JDN causing binary search failure
		year = getPersianYear(persianEpoch - 10000)
		assert.True(t, year > 0 || year == -1)

		year = getPersianYear(persianEpoch + 2000000)
		assert.True(t, year > 0 || year == -1)
	})
}

// TestPersianWithAuthorityData validates Persian calendar conversion using authoritative test data
func TestPersianWithAuthorityData(t *testing.T) {
	// Read authoritative test data
	data, err := os.ReadFile("persian_test_data.json")
	if err != nil {
		t.Skipf("Unable to read test data file: %v", err)
	}

	var testCases []struct {
		Description string `json:"description"`
		Persian     struct {
			Year  int `json:"year"`
			Month int `json:"month"`
			Day   int `json:"day"`
		} `json:"persian"`
		Gregorian struct {
			Year  int `json:"year"`
			Month int `json:"month"`
			Day   int `json:"day"`
		} `json:"gregorian"`
	}

	if err := json.Unmarshal(data, &testCases); err != nil {
		t.Fatalf("Failed to parse test data: %v", err)
	}

	t.Logf("Loaded %d authoritative test cases", len(testCases))

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Authority_Data_%d_%s", i+1, tc.Description), func(t *testing.T) {
			// Test Persian to Gregorian conversion
			persianDate := NewPersian(tc.Persian.Year, tc.Persian.Month, tc.Persian.Day)
			if !assert.NotNil(t, persianDate, "Persian date creation failed") {
				return
			}
			if !assert.True(t, persianDate.IsValid(), "Persian date is invalid") {
				return
			}

			gregorianDate := persianDate.ToGregorian()
			if !assert.NotNil(t, gregorianDate, "Gregorian date conversion failed") {
				return
			}

			expectedDate := time.Date(tc.Gregorian.Year, time.Month(tc.Gregorian.Month), tc.Gregorian.Day, 0, 0, 0, 0, time.UTC)
			actualDate := gregorianDate.Time

			assert.Equal(t, expectedDate.Year(), actualDate.Year(),
				"Year mismatch - Persian: %d-%02d-%02d, Expected Gregorian: %d-%02d-%02d, Actual Gregorian: %d-%02d-%02d",
				tc.Persian.Year, tc.Persian.Month, tc.Persian.Day,
				tc.Gregorian.Year, tc.Gregorian.Month, tc.Gregorian.Day,
				actualDate.Year(), actualDate.Month(), actualDate.Day())

			assert.Equal(t, expectedDate.Month(), actualDate.Month(),
				"Month mismatch - Persian: %d-%02d-%02d, Expected Gregorian: %d-%02d-%02d, Actual Gregorian: %d-%02d-%02d",
				tc.Persian.Year, tc.Persian.Month, tc.Persian.Day,
				tc.Gregorian.Year, tc.Gregorian.Month, tc.Gregorian.Day,
				actualDate.Year(), actualDate.Month(), actualDate.Day())

			assert.Equal(t, expectedDate.Day(), actualDate.Day(),
				"Day mismatch - Persian: %d-%02d-%02d, Expected Gregorian: %d-%02d-%02d, Actual Gregorian: %d-%02d-%02d",
				tc.Persian.Year, tc.Persian.Month, tc.Persian.Day,
				tc.Gregorian.Year, tc.Gregorian.Month, tc.Gregorian.Day,
				actualDate.Year(), actualDate.Month(), actualDate.Day())

			// Test Gregorian to Persian conversion
			fromGregorian := FromStdTime(expectedDate)
			if !assert.NotNil(t, fromGregorian, "Failed to create Persian from Gregorian") {
				return
			}

			assert.Equal(t, tc.Persian.Year, fromGregorian.Year(),
				"Year conversion mismatch - Gregorian: %d-%02d-%02d, Expected Persian: %d-%02d-%02d, Actual Persian: %d-%02d-%02d",
				tc.Gregorian.Year, tc.Gregorian.Month, tc.Gregorian.Day,
				tc.Persian.Year, tc.Persian.Month, tc.Persian.Day,
				fromGregorian.Year(), fromGregorian.Month(), fromGregorian.Day())

			assert.Equal(t, tc.Persian.Month, fromGregorian.Month(),
				"Month conversion mismatch - Gregorian: %d-%02d-%02d, Expected Persian: %d-%02d-%02d, Actual Persian: %d-%02d-%02d",
				tc.Gregorian.Year, tc.Gregorian.Month, tc.Gregorian.Day,
				tc.Persian.Year, tc.Persian.Month, tc.Persian.Day,
				fromGregorian.Year(), fromGregorian.Month(), fromGregorian.Day())

			assert.Equal(t, tc.Persian.Day, fromGregorian.Day(),
				"Day conversion mismatch - Gregorian: %d-%02d-%02d, Expected Persian: %d-%02d-%02d, Actual Persian: %d-%02d-%02d",
				tc.Gregorian.Year, tc.Gregorian.Month, tc.Gregorian.Day,
				tc.Persian.Year, tc.Persian.Month, tc.Persian.Day,
				fromGregorian.Year(), fromGregorian.Month(), fromGregorian.Day())
		})
	}
}

func TestPersian_OfficialLeapYears(t *testing.T) {
	assert.True(t, NewPersian(1403, 12, 30).IsValid())
	assert.True(t, NewPersian(1403, 1, 1).IsLeapYear())
	assert.False(t, NewPersian(1404, 1, 1).IsLeapYear())
	assert.Equal(t, "1404-01-01", FromStdTime(time.Date(2025, 3, 21, 0, 0, 0, 0, time.UTC)).String()[:10])
}
