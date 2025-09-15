package carbon

import (
	"testing"
)

// BenchmarkFormat2Layout benchmarks the format2layout function
func BenchmarkFormat2Layout(b *testing.B) {
	formats := []string{
		"Y-m-d H:i:s",
		"Y-m-d H:i:s.u",
		"Y-m-d H:i:s.v",
		"Y-m-d H:i:s.x",
		"D, M j, Y g:i A",
		"l, F jS, Y",
		"Y-m-d\\TH:i:sP",
		"Y-m-d\\TH:i:s.uP",
		"Y-m-d\\TH:i:s.vP",
		"Y-m-d\\TH:i:s.xP",
		"Y-m-d\\TH:i:s\\Z",
		"Y-m-d\\TH:i:s.u\\Z",
		"Y-m-d\\TH:i:s.v\\Z",
		"Y-m-d\\TH:i:s.x\\Z",
		"Y-m-d H:i:s O",
		"Y-m-d H:i:s P",
		"Y-m-d H:i:s R",
		"Y-m-d H:i:s Q",
		"Y-m-d H:i:s Z",
		"Y-m-d H:i:s T",
		"Y-m-d H:i:s c",
		"Y-m-d H:i:s C",
		"Y-m-d H:i:s r",
		"Y-m-d H:i:s U",
		"Y-m-d H:i:s V",
		"Y-m-d H:i:s X",
		"Y-m-d H:i:s S",
		"Y-m-d H:i:s W",
		"Y-m-d H:i:s N",
		"Y-m-d H:i:s w",
		"Y-m-d H:i:s z",
		"Y-m-d H:i:s t",
		"Y-m-d H:i:s L",
		"Y-m-d H:i:s o",
		"Y-m-d H:i:s B",
		"Y-m-d H:i:s g",
		"Y-m-d H:i:s h",
		"Y-m-d H:i:s G",
		"Y-m-d H:i:s H",
		"Y-m-d H:i:s i",
		"Y-m-d H:i:s s",
		"Y-m-d H:i:s a",
		"Y-m-d H:i:s A",
		"Y-m-d H:i:s K",
		"Y-m-d H:i:s q",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, format := range formats {
			_ = format2layout(format)
		}
	}
}

// BenchmarkFormat2Layout_Simple benchmarks simple format strings
func BenchmarkFormat2Layout_Simple(b *testing.B) {
	formats := []string{
		"Y-m-d",
		"H:i:s",
		"Y-m-d H:i:s",
		"D, M j, Y",
		"l, F jS, Y",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, format := range formats {
			_ = format2layout(format)
		}
	}
}

// BenchmarkFormat2Layout_Complex benchmarks complex format strings
func BenchmarkFormat2Layout_Complex(b *testing.B) {
	formats := []string{
		"Y-m-d\\TH:i:sP",
		"Y-m-d\\TH:i:s.uP",
		"Y-m-d\\TH:i:s.vP",
		"Y-m-d\\TH:i:s.xP",
		"Y-m-d\\TH:i:s\\Z",
		"Y-m-d\\TH:i:s.u\\Z",
		"Y-m-d\\TH:i:s.v\\Z",
		"Y-m-d\\TH:i:s.x\\Z",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, format := range formats {
			_ = format2layout(format)
		}
	}
}

// BenchmarkFormat2Layout_Memory benchmarks memory allocation
func BenchmarkFormat2Layout_Memory(b *testing.B) {
	format := "Y-m-d H:i:s.u.v.x P Q R Z T c C r U V X S W N w z t L o B g h G H i s a A K q"

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = format2layout(format)
	}
}

// BenchmarkParseTimezone benchmarks the parseTimezone function
func BenchmarkParseTimezone(b *testing.B) {
	timezones := []string{
		"UTC",
		"Local",
		"Asia/Shanghai",
		"America/New_York",
		"Europe/London",
		"Asia/Tokyo",
		"Australia/Sydney",
		"America/Los_Angeles",
		"Europe/Paris",
		"Asia/Kolkata",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tz := range timezones {
			_, _ = parseTimezone(tz)
		}
	}
}

// BenchmarkParseDuration benchmarks the parseDuration function
func BenchmarkParseDuration(b *testing.B) {
	durations := []string{
		"1s",
		"1m",
		"1h",
		"1d",
		"1w",
		"1M",
		"1y",
		"1.5s",
		"2.5m",
		"3.5h",
		"1h30m",
		"2h45m30s",
		"1d2h3m4s",
		"1w2d3h4m5s",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, dur := range durations {
			_, _ = parseDuration(dur)
		}
	}
}

// BenchmarkGetAbsValue benchmarks the getAbsValue function
func BenchmarkGetAbsValue(b *testing.B) {
	values := []int64{
		0, 1, -1, 100, -100, 1000, -1000,
		10000, -10000, 100000, -100000,
		1000000, -1000000, 10000000, -10000000,
		100000000, -100000000, 1000000000, -1000000000,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, val := range values {
			_ = getAbsValue(val)
		}
	}
}
