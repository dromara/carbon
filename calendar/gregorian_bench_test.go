package calendar

import (
	"testing"
	"time"
)

func BenchmarkString(b *testing.B) {
	testCases := []*Gregorian{
		{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
		{Time: time.Date(2024, 2, 29, 12, 30, 45, 123456789, time.UTC)}, // 闰年闰日
		{Time: time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC)},
		{Time: time.Date(2020, 8, 5, 0, 0, 0, 0, time.UTC)}, // 闰年
		{Time: time.Date(2021, 6, 15, 15, 30, 0, 0, time.UTC)},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := testCases[i%len(testCases)]
		_ = g.String()
	}
}

func BenchmarkStringNil(b *testing.B) {
	var g *Gregorian

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.String()
	}
}

func BenchmarkStringZeroTime(b *testing.B) {
	g := &Gregorian{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.String()
	}
}

func BenchmarkIsLeapYear(b *testing.B) {
	testCases := []*Gregorian{
		{Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},   // 闰年
		{Time: time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)},  // 闰年闰日
		{Time: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)},   // 世纪闰年
		{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},   // 非闰年
		{Time: time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)}, // 非闰年
		{Time: time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC)},   // 世纪非闰年
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := testCases[i%len(testCases)]
		g.IsLeapYear()
	}
}

func BenchmarkIsLeapYearNil(b *testing.B) {
	var g *Gregorian

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.IsLeapYear()
	}
}

func BenchmarkIsLeapYearWithError(b *testing.B) {
	g := &Gregorian{Error: &time.ParseError{}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.IsLeapYear()
	}
}

func BenchmarkIsLeapYearZeroTime(b *testing.B) {
	g := &Gregorian{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.IsLeapYear()
	}
}

func BenchmarkLeapYearEdgeCases(b *testing.B) {
	testCases := []*Gregorian{
		{Time: time.Date(1600, 1, 1, 0, 0, 0, 0, time.UTC)}, // 世纪闰年
		{Time: time.Date(1700, 1, 1, 0, 0, 0, 0, time.UTC)}, // 世纪非闰年
		{Time: time.Date(1800, 1, 1, 0, 0, 0, 0, time.UTC)}, // 世纪非闰年
		{Time: time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)}, // 世纪非闰年
		{Time: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)}, // 世纪闰年
		{Time: time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC)}, // 世纪非闰年
		{Time: time.Date(2200, 1, 1, 0, 0, 0, 0, time.UTC)}, // 世纪非闰年
		{Time: time.Date(2300, 1, 1, 0, 0, 0, 0, time.UTC)}, // 世纪非闰年
		{Time: time.Date(2400, 1, 1, 0, 0, 0, 0, time.UTC)}, // 世纪闰年
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := testCases[i%len(testCases)]
		g.IsLeapYear()
	}
}

func BenchmarkLeapYearRegularYears(b *testing.B) {
	testCases := []*Gregorian{
		{Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}, // 闰年
		{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)}, // 非闰年
		{Time: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)}, // 非闰年
		{Time: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)}, // 非闰年
		{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}, // 闰年
		{Time: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)}, // 非闰年
		{Time: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)}, // 非闰年
		{Time: time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC)}, // 非闰年
		{Time: time.Date(2028, 1, 1, 0, 0, 0, 0, time.UTC)}, // 闰年
		{Time: time.Date(2029, 1, 1, 0, 0, 0, 0, time.UTC)}, // 非闰年
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := testCases[i%len(testCases)]
		g.IsLeapYear()
	}
}
