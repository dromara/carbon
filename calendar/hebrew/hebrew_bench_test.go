package hebrew

import (
	"testing"
	"time"
)

func BenchmarkNewHebrew(b *testing.B) {
	b.Run("valid_date", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NewHebrew(5784, 10, 20)
		}
	})

	b.Run("invalid_date", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NewHebrew(10000, 13, 1)
		}
	})

	b.Run("leap_year", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NewHebrew(5784, 13, 1)
		}
	})
}

func BenchmarkFromStdTime(b *testing.B) {
	testTime := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)

	b.Run("normal_date", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FromStdTime(testTime)
		}
	})

	b.Run("zero_time", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FromStdTime(time.Time{})
		}
	})

	b.Run("year_1_ce", func(b *testing.B) {
		year1Time := time.Date(1, 1, 1, 12, 0, 0, 0, time.UTC)
		for i := 0; i < b.N; i++ {
			FromStdTime(year1Time)
		}
	})
}

func BenchmarkHebrew_ToGregorian(b *testing.B) {
	h := NewHebrew(5784, 10, 20)

	b.Run("without_timezone", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h.ToGregorian()
		}
	})

	b.Run("with_timezone", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h.ToGregorian("PRC")
		}
	})

	b.Run("invalid_timezone", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h.ToGregorian("xxx")
		}
	})
}

func BenchmarkHebrew_IsValid(b *testing.B) {
	validHebrew := NewHebrew(5784, 10, 20)
	invalidHebrew := NewHebrew(10000, 13, 1)

	b.Run("valid_date", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validHebrew.IsValid()
		}
	})

	b.Run("invalid_date", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			invalidHebrew.IsValid()
		}
	})

	b.Run("nil_hebrew", func(b *testing.B) {
		var h *Hebrew
		for i := 0; i < b.N; i++ {
			h.IsValid()
		}
	})
}

func BenchmarkHebrew_IsLeapYear(b *testing.B) {
	leapYear := NewHebrew(5784, 1, 1)
	nonLeapYear := NewHebrew(5785, 1, 1)

	b.Run("leap_year", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			leapYear.IsLeapYear()
		}
	})

	b.Run("non_leap_year", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nonLeapYear.IsLeapYear()
		}
	})
}

func BenchmarkHebrew_String(b *testing.B) {
	h := NewHebrew(5784, 10, 20)

	b.Run("valid_date", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h.String()
		}
	})
}

func BenchmarkHebrew_ToMonthString(b *testing.B) {
	h := NewHebrew(5784, 10, 20)

	b.Run("english_locale", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h.ToMonthString(EnLocale)
		}
	})

	b.Run("hebrew_locale", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h.ToMonthString(HeLocale)
		}
	})

	b.Run("default_locale", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h.ToMonthString()
		}
	})

	b.Run("invalid_locale", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h.ToMonthString("xxx")
		}
	})
}

func BenchmarkHebrew_ToWeekString(b *testing.B) {
	h := NewHebrew(5784, 10, 20)

	b.Run("english_locale", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h.ToWeekString(EnLocale)
		}
	})

	b.Run("hebrew_locale", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h.ToWeekString(HeLocale)
		}
	})

	b.Run("default_locale", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h.ToWeekString()
		}
	})
}

func BenchmarkGregorian2jdn(b *testing.B) {
	b.Run("normal_date", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gregorian2jdn(2024, 1, 1)
		}
	})

	b.Run("leap_year", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gregorian2jdn(2024, 2, 29)
		}
	})

	b.Run("early_date", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gregorian2jdn(1, 1, 1)
		}
	})

	b.Run("late_date", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gregorian2jdn(9999, 12, 31)
		}
	})
}

func BenchmarkJdn2gregorian(b *testing.B) {
	b.Run("normal_jdn", func(b *testing.B) {
		jdn := 2460333 // 2024-01-23
		for i := 0; i < b.N; i++ {
			jdn2gregorian(jdn)
		}
	})

	b.Run("early_jdn", func(b *testing.B) {
		jdn := 1721425 // 1-01-01
		for i := 0; i < b.N; i++ {
			jdn2gregorian(jdn)
		}
	})

	b.Run("late_jdn", func(b *testing.B) {
		jdn := 5373484 // 9999-12-31
		for i := 0; i < b.N; i++ {
			jdn2gregorian(jdn)
		}
	})
}

func BenchmarkJdn2hebrew(b *testing.B) {
	b.Run("normal_jdn", func(b *testing.B) {
		jdn := 2459580 // Hebrew 5782-10-27
		for i := 0; i < b.N; i++ {
			jdn2hebrew(jdn)
		}
	})

	b.Run("early_jdn", func(b *testing.B) {
		jdn := 1721426 // Hebrew 3761-10-18
		for i := 0; i < b.N; i++ {
			jdn2hebrew(jdn)
		}
	})

	b.Run("late_jdn", func(b *testing.B) {
		jdn := 5373484 // Hebrew 13760-08-28
		for i := 0; i < b.N; i++ {
			jdn2hebrew(jdn)
		}
	})
}

func BenchmarkHebrew2jdn(b *testing.B) {
	b.Run("normal_date", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hebrew2jdn(5784, 10, 20)
		}
	})

	b.Run("leap_year", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hebrew2jdn(5784, 13, 1)
		}
	})

	b.Run("early_date", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hebrew2jdn(1, 1, 1)
		}
	})

	b.Run("late_date", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hebrew2jdn(9999, 12, 30)
		}
	})
}

func BenchmarkIsLeapYear(b *testing.B) {
	b.Run("leap_years", func(b *testing.B) {
		leapYears := []int{5784, 5787, 5790, 5793, 5796}
		for i := 0; i < b.N; i++ {
			for _, year := range leapYears {
				isLeapYear(year)
			}
		}
	})

	b.Run("non_leap_years", func(b *testing.B) {
		nonLeapYears := []int{5785, 5786, 5788, 5789, 5791}
		for i := 0; i < b.N; i++ {
			for _, year := range nonLeapYears {
				isLeapYear(year)
			}
		}
	})
}

func BenchmarkGetMonthsFromEpoch(b *testing.B) {
	b.Run("early_years", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for year := 1; year <= 100; year++ {
				getMonthsFromEpoch(year)
			}
		}
	})

	b.Run("middle_years", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for year := 5000; year <= 5100; year++ {
				getMonthsFromEpoch(year)
			}
		}
	})

	b.Run("late_years", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for year := 9900; year <= 9999; year++ {
				getMonthsFromEpoch(year)
			}
		}
	})
}

func BenchmarkGetJDNInYear(b *testing.B) {
	b.Run("early_years", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for year := 1; year <= 100; year++ {
				getJDNInYear(year)
			}
		}
	})

	b.Run("middle_years", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for year := 5000; year <= 5100; year++ {
				getJDNInYear(year)
			}
		}
	})

	b.Run("late_years", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for year := 9900; year <= 9999; year++ {
				getJDNInYear(year)
			}
		}
	})

	b.Run("dehiyyot_years", func(b *testing.B) {
		dehiyyotYears := []int{5765, 5766, 5767, 5768, 5769, 5770}
		for i := 0; i < b.N; i++ {
			for _, year := range dehiyyotYears {
				getJDNInYear(year)
			}
		}
	})
}

func BenchmarkGetMonthsInYear(b *testing.B) {
	b.Run("leap_years", func(b *testing.B) {
		leapYears := []int{5784, 5787, 5790, 5793, 5796}
		for i := 0; i < b.N; i++ {
			for _, year := range leapYears {
				getMonthsInYear(year)
			}
		}
	})

	b.Run("non_leap_years", func(b *testing.B) {
		nonLeapYears := []int{5785, 5786, 5788, 5789, 5791}
		for i := 0; i < b.N; i++ {
			for _, year := range nonLeapYears {
				getMonthsInYear(year)
			}
		}
	})
}

func BenchmarkGetDaysInMonth(b *testing.B) {
	b.Run("fixed_29_day_months", func(b *testing.B) {
		fixedMonths := []int{2, 4, 6, 10, 13}
		for i := 0; i < b.N; i++ {
			for _, month := range fixedMonths {
				getDaysInMonth(5784, month)
			}
		}
	})

	b.Run("heshvan_kislev", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			getDaysInMonth(5784, 8) // Heshvan
			getDaysInMonth(5784, 9) // Kislev
		}
	})

	b.Run("adar_variations", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			getDaysInMonth(5784, 12) // Adar in leap year
			getDaysInMonth(5785, 12) // Adar in non-leap year
		}
	})

	b.Run("regular_30_day_months", func(b *testing.B) {
		regularMonths := []int{1, 3, 5, 7, 11}
		for i := 0; i < b.N; i++ {
			for _, month := range regularMonths {
				getDaysInMonth(5784, month)
			}
		}
	})
}

func BenchmarkRoundTripConversion(b *testing.B) {
	b.Run("gregorian_to_hebrew_to_gregorian", func(b *testing.B) {
		testTime := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
		for i := 0; i < b.N; i++ {
			h := FromStdTime(testTime)
			g := h.ToGregorian()
			_ = g
		}
	})

	b.Run("hebrew_to_gregorian_to_hebrew", func(b *testing.B) {
		h := NewHebrew(5784, 10, 20)
		for i := 0; i < b.N; i++ {
			g := h.ToGregorian()
			h2 := FromStdTime(g.Time)
			_ = h2
		}
	})
}

func BenchmarkMultipleConversions(b *testing.B) {
	b.Run("batch_gregorian_to_hebrew", func(b *testing.B) {
		testDates := []time.Time{
			time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
			time.Date(2024, 8, 5, 12, 0, 0, 0, time.UTC),
			time.Date(2025, 10, 3, 12, 0, 0, 0, time.UTC),
			time.Date(2023, 9, 16, 12, 0, 0, 0, time.UTC),
			time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
		}
		for i := 0; i < b.N; i++ {
			for _, date := range testDates {
				FromStdTime(date)
			}
		}
	})

	b.Run("batch_hebrew_to_gregorian", func(b *testing.B) {
		testHebrewDates := []struct {
			year, month, day int
		}{
			{5784, 10, 20},
			{5784, 5, 1},
			{5786, 7, 10},
			{5784, 13, 1},
			{1, 1, 1},
		}
		for i := 0; i < b.N; i++ {
			for _, date := range testHebrewDates {
				h := NewHebrew(date.year, date.month, date.day)
				h.ToGregorian()
			}
		}
	})
}
