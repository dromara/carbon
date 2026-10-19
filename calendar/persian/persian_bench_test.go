package persian

import (
	"testing"
	"time"
)

func BenchmarkFromStdTime(b *testing.B) {
	testDates := []time.Time{
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 6, 21, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 9, 22, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 12, 21, 0, 0, 0, 0, time.UTC),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		date := testDates[i%len(testDates)]
		FromStdTime(date)
	}
}

func BenchmarkToGregorian(b *testing.B) {
	testPersianDates := []*Persian{
		NewPersian(1403, 1, 1),
		NewPersian(1403, 6, 15),
		NewPersian(1403, 12, 29),
		NewPersian(1404, 1, 1),
		NewPersian(1404, 12, 30),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := testPersianDates[i%len(testPersianDates)]
		p.ToGregorian()
	}
}

func BenchmarkIsLeapYear(b *testing.B) {
	testYears := []int{1395, 1396, 1399, 1400, 1403, 1404, 1407, 1408, 1411, 1412}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year := testYears[i%len(testYears)]
		p := NewPersian(year, 1, 1)
		p.IsLeapYear()
	}
}

func BenchmarkIsValid(b *testing.B) {
	testDates := []*Persian{
		NewPersian(1395, 1, 1),
		NewPersian(1395, 12, 30), // 闰年
		NewPersian(1396, 12, 29), // 非闰年
		NewPersian(0, 1, 1),      // 无效
		NewPersian(1395, 13, 1),  // 无效
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := testDates[i%len(testDates)]
		p.IsValid()
	}
}

func BenchmarkString(b *testing.B) {
	p := NewPersian(1403, 6, 15)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.String()
	}
}

func BenchmarkToMonthString(b *testing.B) {
	p := NewPersian(1403, 6, 15)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.ToMonthString()
	}
}

func BenchmarkToWeekString(b *testing.B) {
	p := NewPersian(1403, 6, 15)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.ToWeekString()
	}
}
