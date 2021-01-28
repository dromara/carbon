package carbon

import (
	"fmt"
	"testing"
	"time"
)

var TimezoneTests = []struct {
	input    string // 输入值
	timezone string // 输入参数
	output   string // 期望输出值
}{
	{"2020-08-05 13:14:15", PRC, "2020-08-05 13:14:15"},
	{"2020-08-05", Tokyo, "2020-08-05 01:00:00"},
	{"2020-08-05", "Hangzhou", "panic"}, // 异常情况
}

func TestCarbon_SetTimezone1(t *testing.T) {
	for _, v := range TimezoneTests {
		output := SetTimezone(v.timezone).Parse(v.input)

		if output.Error != nil {
			fmt.Println("catch an exception in SetTimezone():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output.ToDateTimeString())
		}
	}
}

func TestCarbon_SetTimezone2(t *testing.T) {
	for _, v := range TimezoneTests {
		output := SetTimezone(PRC).SetTimezone(v.timezone).Parse(v.input)

		if output.Error != nil {
			fmt.Println("catch an exception in SetTimezone():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output.ToDateTimeString())
		}
	}
}

func TestCarbon_SetYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		year   int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-01-01", 2019, "2019-01-01"},
		{"2020-01-31", 2019, "2019-01-31"},
		{"2020-02-01", 2019, "2019-02-01"},
		{"2020-02-28", 2019, "2019-02-28"},
		{"2020-02-29", 2019, "2019-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SetYear(v.year).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_SetMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		month  int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-01-01", 2, "2020-02-01"},
		{"2020-01-30", 2, "2020-03-01"},
		{"2020-01-31", 2, "2020-03-02"},
		{"2020-08-05", 2, "2020-02-05"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SetMonth(v.month).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_SetDay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		day    int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-01-01", 31, "2020-01-31"},
		{"2020-02-01", 31, "2020-03-02"},
		{"2020-02-28", 31, "2020-03-02"},
		{"2020-02-29", 31, "2020-03-02"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SetDay(v.day).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_SetHour(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		hour   int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-08-05 13:14:15", 10, "2020-08-05 10:14:15"},
		{"2020-08-05 13:14:15", 24, "2020-08-06 00:14:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SetHour(v.hour).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_SetMinute(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		minute int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-08-05 13:14:15", 10, "2020-08-05 13:10:15"},
		{"2020-08-05 13:14:15", 60, "2020-08-05 14:00:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SetMinute(v.minute).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_SetSecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		second int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-08-05 13:14:15", 10, "2020-08-05 13:14:10"},
		{"2020-08-05 13:14:15", 60, "2020-08-05 13:15:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SetSecond(v.second).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_Now(t *testing.T) {
	expected := time.Now().Format(DateTimeFormat)

	output := Now().Time.Format(DateTimeFormat)

	if expected != output {
		t.Fatalf("Expected %s, but got %s", expected, output)
	}

	output = SetTimezone(PRC).Now().Time.Format(DateTimeFormat)

	if expected != output {
		t.Fatalf("Expected %s, but got %s", expected, output)
	}

}

func TestCarbon_Yesterday(t *testing.T) {
	expected := time.Now().AddDate(0, 0, -1).Format(DateTimeFormat)

	output := Yesterday().Time.Format(DateTimeFormat)

	if expected != output {
		t.Fatalf("Expected %s, but got %s", expected, output)
	}

	output = SetTimezone(PRC).Yesterday().Time.Format(DateTimeFormat)

	if expected != output {
		t.Fatalf("Expected %s, but got %s", expected, output)
	}
}

func TestCarbon_Tomorrow(t *testing.T) {
	expected := time.Now().AddDate(0, 0, 1).Format(DateTimeFormat)

	output := Tomorrow().Time.Format(DateTimeFormat)

	if expected != output {
		t.Fatalf("Expected %s, but got %s", expected, output)
	}

	output = SetTimezone(PRC).Tomorrow().Time.Format(DateTimeFormat)

	if expected != output {
		t.Fatalf("Expected %s, but got %s", expected, output)
	}
}

func TestCarbon_StartOfYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-01 00:00:00"},
		{"2020-02-01 13:14:15", "2020-01-01 00:00:00"},
		{"2020-02-28", "2020-01-01 00:00:00"},
		{"2020-02-29", "2020-01-01 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfYear().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-12-31 23:59:59"},
		{"2020-01-31 23:59:59", "2020-12-31 23:59:59"},
		{"2020-02-01 13:14:15", "2020-12-31 23:59:59"},
		{"2020-02-28", "2020-12-31 23:59:59"},
		{"2020-02-29", "2020-12-31 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfYear().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_StartOfMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-01 00:00:00"},
		{"2020-02-01 13:14:15", "2020-02-01 00:00:00"},
		{"2020-02-28", "2020-02-01 00:00:00"},
		{"2020-02-29", "2020-02-01 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfMonth().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-31 23:59:59"},
		{"2020-01-31 23:59:59", "2020-01-31 23:59:59"},
		{"2020-02-01 13:14:15", "2020-02-29 23:59:59"},
		{"2020-02-28", "2020-02-29 23:59:59"},
		{"2020-02-29", "2020-02-29 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfMonth().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_StartOfWeek(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2019-12-30 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-27 00:00:00"},
		{"2020-02-01 13:14:15", "2020-01-27 00:00:00"},
		{"2020-02-28", "2020-02-24 00:00:00"},
		{"2020-02-29", "2020-02-24 00:00:00"},
		{"2020-10-04", "2020-09-28 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfWeek().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfWeek(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-05 23:59:59"},
		{"2020-01-31 23:59:59", "2020-02-02 23:59:59"},
		{"2020-02-01 13:14:15", "2020-02-02 23:59:59"},
		{"2020-02-28", "2020-03-01 23:59:59"},
		{"2020-02-29", "2020-03-01 23:59:59"},
		{"2020-10-04", "2020-10-04 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfWeek().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_StartOfDay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-31 00:00:00"},
		{"2020-02-01 13:14:15", "2020-02-01 00:00:00"},
		{"2020-02-28", "2020-02-28 00:00:00"},
		{"2020-02-29", "2020-02-29 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfDay().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfDay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-01 23:59:59"},
		{"2020-01-31 23:59:59", "2020-01-31 23:59:59"},
		{"2020-02-01 13:14:15", "2020-02-01 23:59:59"},
		{"2020-02-28", "2020-02-28 23:59:59"},
		{"2020-02-29", "2020-02-29 23:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfDay().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_StartOfHour(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-31 23:00:00"},
		{"2020-02-01 13:14:15", "2020-02-01 13:00:00"},
		{"2020-02-28", "2020-02-28 00:00:00"},
		{"2020-02-29", "2020-02-29 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfHour().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfHour(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-01 00:59:59"},
		{"2020-01-31 23:59:59", "2020-01-31 23:59:59"},
		{"2020-02-01 13:14:15", "2020-02-01 13:59:59"},
		{"2020-02-28", "2020-02-28 00:59:59"},
		{"2020-02-29", "2020-02-29 00:59:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfHour().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_StartOfMinute(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-01 00:00:00"},
		{"2020-01-31 23:59:59", "2020-01-31 23:59:00"},
		{"2020-02-01 13:14:15", "2020-02-01 13:14:00"},
		{"2020-02-28", "2020-02-28 00:00:00"},
		{"2020-02-29", "2020-02-29 00:00:00"},
	}

	for _, v := range Tests {
		output := Parse(v.input).StartOfMinute().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_EndOfMinute(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 00:00:00", "2020-01-01 00:00:59"},
		{"2020-01-31 23:59:59", "2020-01-31 23:59:59"},
		{"2020-02-01 13:14:15", "2020-02-01 13:14:59"},
		{"2020-02-28", "2020-02-28 00:00:59"},
		{"2020-02-29", "2020-02-29 00:00:59"},
	}

	for _, v := range Tests {
		output := Parse(v.input).EndOfMinute().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s", v.input, v.output, output)
		}
	}
}

func TestCarbon_CreateFromTimestamp(t *testing.T) {
	Tests := []struct {
		timestamp int64  // 输入参数
		output    string // 期望输出值
	}{
		{123456, "1970-01-01 08:00:00"},
		{1577855655, "2020-01-01 13:14:15"},
		{1604074084682, "2020-10-31 00:08:04"},
		{1604074196366540, "2020-10-31 00:09:56"},
		{1604074298500312000, "2020-10-31 00:11:38"},
	}

	for _, v := range Tests {
		output := CreateFromTimestamp(v.timestamp).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %d, expected %s, but got %s", v.timestamp, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).CreateFromTimestamp(v.timestamp).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}
}

func TestCarbon_CreateFromDateTime(t *testing.T) {
	Tests := []struct {
		year, month, day, hour, minute, second int    // 输入参数
		output                                 string // 期望输出值
	}{
		{2020, 01, 01, 13, 14, 15, "2020-01-01 13:14:15"},
		{2020, 1, 31, 13, 14, 15, "2020-01-31 13:14:15"},
		{2020, 2, 1, 13, 14, 15, "2020-02-01 13:14:15"},
		{2020, 2, 28, 13, 14, 15, "2020-02-28 13:14:15"},
		{2020, 2, 29, 13, 14, 15, "2020-02-29 13:14:15"},
	}

	for _, v := range Tests {
		output := CreateFromDateTime(v.year, v.month, v.day, v.hour, v.minute, v.second).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).CreateFromDateTime(v.year, v.month, v.day, v.hour, v.minute, v.second).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}
}

func TestCarbon_CreateFromDate(t *testing.T) {
	clock := Now().ToTimeString()

	Tests := []struct {
		year, month, day int    // 输入参数
		output           string // 期望输出值
	}{
		{2020, 01, 01, "2020-01-01 " + clock},
		{2020, 1, 31, "2020-01-31 " + clock},
		{2020, 2, 1, "2020-02-01 " + clock},
		{2020, 2, 28, "2020-02-28 " + clock},
		{2020, 2, 29, "2020-02-29 " + clock},
	}

	for _, v := range Tests {
		output := CreateFromDate(v.year, v.month, v.day).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).CreateFromDate(v.year, v.month, v.day).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}
}

func TestCarbon_CreateFromTime(t *testing.T) {
	date := Now().ToDateString()

	Tests := []struct {
		hour, minute, second int    // 输入参数
		output               string // 期望输出值
	}{
		{0, 0, 0, date + " 00:00:00"},
		{00, 00, 15, date + " 00:00:15"},
		{00, 14, 15, date + " 00:14:15"},
		{13, 14, 15, date + " 13:14:15"},
	}

	for _, v := range Tests {
		output := CreateFromTime(v.hour, v.minute, v.second).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).CreateFromTime(v.hour, v.minute, v.second).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}
}

func TestCarbon_CreateFromGoTime(t *testing.T) {
	Tests := []struct {
		time   time.Time // // 输入参数
		output string    // 期望输出值
	}{
		{time.Now(), time.Now().Format(DateTimeFormat)},
	}

	for _, v := range Tests {
		output := CreateFromGoTime(v.time).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).CreateFromGoTime(v.time).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}
}

func TestCarbon_Parse(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"0000-00-00 00:00:00", ""},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15"},
		{"20200805131415", "2020-08-05 13:14:15"},
		{"20200805", "2020-08-05 00:00:00"},
		{"2020-08-05", "2020-08-05 00:00:00"},
		{"2020-08-05T13:14:15+08:00", "2020-08-05 13:14:15"},
		{"12345678", ""}, // 异常情况
	}

	for _, v := range Tests {
		output := Parse(v.input)

		if output.Error != nil {
			fmt.Println("catch an exception in Parse():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input)

		if output.Error != nil {
			fmt.Println("catch an exception in Parse():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}
}

var ParseByFormatTests = []struct {
	input  string // 输入值
	format string // 输入参数
	output string // 期望输出值
}{
	{"2020|08|05", "Y|m|d", "2020-08-05 00:00:00"},
	{"2020|08|05 13:14:15", "Y|m|d H:i:s", "2020-08-05 13:14:15"},
	{"12345678", "XXXX", ""}, // 异常情况
}

func TestCarbon_ParseByFormat1(t *testing.T) {
	for _, v := range ParseByFormatTests {
		output := ParseByFormat(v.input, v.format)

		if output.Error != nil {
			fmt.Println("catch an exception in ParseByFormat():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}
}

func TestCarbon_ParseByFormat2(t *testing.T) {
	for _, v := range ParseByFormatTests {
		output := SetTimezone("XXXX").ParseByFormat(v.input, v.format)

		if output.Error != nil {
			fmt.Println("catch an exception in ParseByFormat():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}
}

var ParseByDurationTests = []struct {
	input    string // 输入值
	duration string // 输入参数
	output   string // 期望输出值
}{
	{Now().ToDateTimeString(), "10h", ParseByDuration("10h").Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "-10h", ParseByDuration("-10h").Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "10.5h", ParseByDuration("10.5h").Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "-10.5h", ParseByDuration("-10.5h").Format(DateTimeFormat)},

	{Now().ToDateTimeString(), "10m", ParseByDuration("10m").Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "-10m", ParseByDuration("-10m").Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "10.5m", ParseByDuration("10.5m").Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "-10.5m", ParseByDuration("-10.5m").Format(DateTimeFormat)},

	{Now().ToDateTimeString(), "10s", ParseByDuration("10s").Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "-10s", ParseByDuration("-10s").Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "10.5s", ParseByDuration("10.5s").Format(DateTimeFormat)},
	{Now().ToDateTimeString(), "-10.5s", ParseByDuration("-10.5s").Format(DateTimeFormat)},

	{Now().ToDateTimeString(), "XXXX", ""}, // 异常情况
}

func TestCarbon_ParseByDuration1(t *testing.T) {
	for _, v := range ParseByDurationTests {
		output := ParseByDuration(v.duration)

		if output.Error != nil {
			fmt.Println("catch an exception in ParseByDuration():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}
}

func TestCarbon_ParseByDuration2(t *testing.T) {
	for _, v := range ParseByDurationTests {
		output := SetTimezone("XXXX").ParseByDuration(v.duration)

		if output.Error != nil {
			fmt.Println("catch an exception in ParseByDuration():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}
}

func TestCarbon_AddDuration(t *testing.T) {
	Tests := []struct {
		input    string // 输入值
		duration string // 输入参数
		output   string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "10h", "2020-01-01 23:14:15"},
		{"2020-01-01 13:14:15", "10.5h", "2020-01-01 23:44:15"},

		{"2020-01-01 13:14:15", "10m", "2020-01-01 13:24:15"},
		{"2020-01-01 13:14:15", "10.5m", "2020-01-01 13:24:45"},

		{"2020-01-01 13:14:15", "10s", "2020-01-01 13:14:25"},
		{"2020-01-01 13:14:15", "10.5s", "2020-01-01 13:14:25"},

		{"12345678", "10h", ""},             // 异常情况
		{"2020-01-01 13:14:15", "XXXX", ""}, // 异常情况
	}

	for _, v := range Tests {
		output := Parse(v.input).AddDuration(v.duration)

		if output.Error != nil {
			fmt.Println("catch an exception in AddDuration():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}

	for _, v := range Tests {
		output := SetTimezone("XXXX").Parse(v.input).AddDuration(v.duration)

		if output.Error != nil {
			fmt.Println("catch an exception in AddDuration():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}
}

func TestCarbon_SubDuration(t *testing.T) {
	Tests := []struct {
		input    string // 输入值
		duration string // 输入参数
		output   string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "10h", "2020-01-01 03:14:15"},
		{"2020-01-01 13:14:15", "10.5h", "2020-01-01 02:44:15"},

		{"2020-01-01 13:14:15", "10m", "2020-01-01 13:04:15"},
		{"2020-01-01 13:14:15", "10.5m", "2020-01-01 13:03:45"},

		{"2020-01-01 13:14:15", "10s", "2020-01-01 13:14:05"},
		{"2020-01-01 13:14:15", "10.5s", "2020-01-01 13:14:04"},

		{"2020-01-01 13:14:15 XXXX", "10h", ""}, // 异常情况
		{"2020-01-01 13:14:15", "10x", ""},      // 异常情况
	}

	for _, v := range Tests {
		output := Parse(v.input).SubDuration(v.duration)

		if output.Error != nil {
			fmt.Println("catch an exception in SubDuration():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}

	for _, v := range Tests {
		output := SetTimezone("XXXX").Parse(v.input).SubDuration(v.duration)

		if output.Error != nil {
			fmt.Println("catch an exception in SubDuration():", output.Error)
			return
		}

		if output.ToDateTimeString() != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}
}

func TestCarbon_AddCenturies(t *testing.T) {
	type Test struct {
		input     string // 输入值
		centuries int    // 输入参数
		output    string // 期望输出值
	}

	Tests := []Test{
		{"2020-01-01", 3, "2320-01-01"},
		{"2020-01-31", 3, "2320-01-31"},
		{"2020-02-01", 3, "2320-02-01"},
		{"2020-02-28", 3, "2320-02-28"},
		{"2020-02-29", 3, "2320-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddCenturies(v.centuries).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddCenturies(v.centuries).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_NextCenturies(t *testing.T) {
	Tests := []struct {
		input     string // 输入值
		centuries int    // 输入参数
		output    string // 期望输出值
	}{
		{"2020-01-01", 3, "2320-01-01"},
		{"2020-01-31", 3, "2320-01-31"},
		{"2020-02-01", 3, "2320-02-01"},
		{"2020-02-28", 3, "2320-02-28"},
		{"2020-02-29", 3, "2320-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).NextCenturies(v.centuries).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).NextCenturies(v.centuries).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubCenturies(t *testing.T) {
	type Test struct {
		input     string // 输入值
		centuries int    // 输入参数
		output    string // 期望输出值
	}

	Tests := []Test{
		{"2020-01-01", 3, "1720-01-01"},
		{"2020-01-31", 3, "1720-01-31"},
		{"2020-02-01", 3, "1720-02-01"},
		{"2020-02-28", 3, "1720-02-28"},
		{"2020-02-29", 3, "1720-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubCenturies(v.centuries).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubCenturies(v.centuries).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_PreCenturies(t *testing.T) {
	Tests := []struct {
		input     string // 输入值
		centuries int
		output    string // 期望输出值
	}{
		{"2020-01-01", 3, "1720-01-01"},
		{"2020-01-31", 3, "1720-01-31"},
		{"2020-02-01", 3, "1720-02-01"},
		{"2020-02-28", 3, "1720-02-28"},
		{"2020-02-29", 3, "1720-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).PreCenturies(v.centuries).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).PreCenturies(v.centuries).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddCentury(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2120-01-01"},
		{"2020-01-31", "2120-01-31"},
		{"2020-02-01", "2120-02-01"},
		{"2020-02-28", "2120-02-28"},
		{"2020-02-29", "2120-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddCentury().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddCentury().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_NextCentury(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2120-01-01"},
		{"2020-01-31", "2120-01-31"},
		{"2020-02-01", "2120-02-01"},
		{"2020-02-28", "2120-02-28"},
		{"2020-02-29", "2120-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).NextCentury().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).NextCentury().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubCentury(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "1920-01-01"},
		{"2020-01-31", "1920-01-31"},
		{"2020-02-01", "1920-02-01"},
		{"2020-02-28", "1920-02-28"},
		{"2020-02-29", "1920-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubCentury().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubCentury().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_PreCentury(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "1920-01-01"},
		{"2020-01-31", "1920-01-31"},
		{"2020-02-01", "1920-02-01"},
		{"2020-02-28", "1920-02-28"},
		{"2020-02-29", "1920-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).PreCentury().ToDateString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).PreCentury().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddYears(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		years  int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2023-01-01"},
		{"2020-01-31", 3, "2023-01-31"},
		{"2020-02-01", 3, "2023-02-01"},
		{"2020-02-28", 3, "2023-02-28"},
		{"2020-02-29", 3, "2023-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddYears(v.years).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddYears(v.years).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_NextYears(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		years  int    // 输入参数
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2023-01-01"},
		{"2020-01-31", 3, "2023-01-31"},
		{"2020-02-01", 3, "2023-02-01"},
		{"2020-02-28", 3, "2023-02-28"},
		{"2020-02-29", 3, "2023-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).NextYears(v.years).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).NextYears(v.years).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubYears(t *testing.T) {
	type Test struct {
		input  string // 输入值
		years  int    // 输入参数
		output string // 期望输出值
	}

	Tests := []Test{
		{"2020-01-01", 3, "2017-01-01"},
		{"2020-01-31", 3, "2017-01-31"},
		{"2020-02-01", 3, "2017-02-01"},
		{"2020-02-28", 3, "2017-02-28"},
		{"2020-02-29", 3, "2017-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubYears(v.years).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubYears(v.years).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_PreYears(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		years  int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2017-01-01"},
		{"2020-01-31", 3, "2017-01-31"},
		{"2020-02-01", 3, "2017-02-01"},
		{"2020-02-28", 3, "2017-02-28"},
		{"2020-02-29", 3, "2017-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).PreYears(v.years).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).PreYears(v.years).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2021-01-01"},
		{"2020-01-31", "2021-01-31"},
		{"2020-02-01", "2021-02-01"},
		{"2020-02-28", "2021-02-28"},
		{"2020-02-29", "2021-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddYear().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddYear().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_NextYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2021-01-01"},
		{"2020-01-31", "2021-01-31"},
		{"2020-02-01", "2021-02-01"},
		{"2020-02-28", "2021-02-28"},
		{"2020-02-29", "2021-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).NextYear().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).NextYear().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-01-01"},
		{"2020-01-31", "2019-01-31"},
		{"2020-02-01", "2019-02-01"},
		{"2020-02-28", "2019-02-28"},
		{"2020-02-29", "2019-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubYear().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubYear().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_PreYear(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-01-01"},
		{"2020-01-31", "2019-01-31"},
		{"2020-02-01", "2019-02-01"},
		{"2020-02-28", "2019-02-28"},
		{"2020-02-29", "2019-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).PreYear().ToDateString()

		if output != v.output {
			t.Fatalf("Expected %s, but got %s", v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).PreYear().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddQuarters(t *testing.T) {
	Tests := []struct {
		input    string // 输入值
		quarters int
		output   string // 期望输出值
	}{
		{"2019-08-01", 2, "2020-02-01"},
		{"2019-08-31", 2, "2020-03-02"},
		{"2020-01-01", 2, "2020-07-01"},
		{"2020-02-28", 2, "2020-08-28"},
		{"2020-02-29", 2, "2020-08-29"},
		{"2020-08-31", 2, "2021-03-03"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddQuarters(v.quarters).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddQuarters(v.quarters).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_NextQuarters(t *testing.T) {
	Tests := []struct {
		input    string // 输入值
		quarters int
		output   string // 期望输出值
	}{
		{"2019-08-01", 2, "2020-02-01"},
		{"2019-08-31", 2, "2020-02-29"},
		{"2020-01-01", 2, "2020-07-01"},
		{"2020-02-28", 2, "2020-08-28"},
		{"2020-02-29", 2, "2020-08-29"},
		{"2020-08-31", 2, "2021-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).NextQuarters(v.quarters).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).NextQuarters(v.quarters).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubQuarters(t *testing.T) {
	Tests := []struct {
		input    string // 输入值
		quarters int
		output   string // 期望输出值
	}{
		{"2019-08-01", 2, "2019-02-01"},
		{"2019-08-31", 2, "2019-03-03"},
		{"2020-01-01", 2, "2019-07-01"},
		{"2020-02-28", 2, "2019-08-28"},
		{"2020-02-29", 2, "2019-08-29"},
		{"2020-08-31", 2, "2020-03-02"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubQuarters(v.quarters).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubQuarters(v.quarters).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_PreQuarters(t *testing.T) {
	Tests := []struct {
		input    string // 输入值
		quarters int
		output   string // 期望输出值
	}{
		{"2019-08-01", 2, "2019-02-01"},
		{"2019-08-31", 2, "2019-02-28"},
		{"2020-01-01", 2, "2019-07-01"},
		{"2020-02-28", 2, "2019-08-28"},
		{"2020-02-29", 2, "2019-08-29"},
		{"2020-08-31", 2, "2020-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).PreQuarters(v.quarters).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).PreQuarters(v.quarters).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddQuarter(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2019-11-01", "2020-02-01"},
		{"2019-11-30", "2020-03-01"},
		{"2020-02-28", "2020-05-28"},
		{"2020-02-29", "2020-05-29"},
		{"2020-08-31", "2020-12-01"},
		{"2020-11-01", "2021-02-01"},
		{"2020-11-30", "2021-03-02"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddQuarter().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddQuarter().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_NextQuarter(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2019-11-01", "2020-02-01"},
		{"2019-11-30", "2020-02-29"},
		{"2020-02-28", "2020-05-28"},
		{"2020-02-29", "2020-05-29"},
		{"2020-08-31", "2020-11-30"},
		{"2020-11-01", "2021-02-01"},
		{"2020-11-30", "2021-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).NextQuarter().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).NextQuarter().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubQuarter(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2019-04-01", "2019-01-01"},
		{"2019-04-30", "2019-01-30"},
		{"2020-05-01", "2020-02-01"},
		{"2020-05-31", "2020-03-02"},
		{"2020-04-01", "2020-01-01"},
		{"2020-04-30", "2020-01-30"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubQuarter().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubQuarter().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_PreQuarter(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2019-04-01", "2019-01-01"},
		{"2019-04-30", "2019-01-30"},
		{"2020-05-01", "2020-02-01"},
		{"2020-05-31", "2020-02-29"},
		{"2020-04-01", "2020-01-01"},
		{"2020-04-30", "2020-01-30"},
	}

	for _, v := range Tests {
		output := Parse(v.input).PreQuarter().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).PreQuarter().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddMonths(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		months int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2020-04-01"},
		{"2020-01-31", 3, "2020-05-01"},
		{"2020-02-01", 3, "2020-05-01"},
		{"2020-02-28", 3, "2020-05-28"},
		{"2020-02-29", 3, "2020-05-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddMonths(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddMonths(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_NextMonths(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		months int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2020-04-01"},
		{"2020-01-31", 3, "2020-04-30"},
		{"2020-02-01", 3, "2020-05-01"},
		{"2020-02-28", 3, "2020-05-28"},
		{"2020-02-29", 3, "2020-05-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).NextMonths(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).NextMonths(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubMonths(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		months int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2019-10-01"},
		{"2020-01-31", 3, "2019-10-31"},
		{"2020-02-01", 3, "2019-11-01"},
		{"2020-02-28", 3, "2019-11-28"},
		{"2020-02-29", 3, "2019-11-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubMonths(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubMonths(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_PreMonths(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		months int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2019-10-01"},
		{"2020-01-31", 3, "2019-10-31"},
		{"2020-02-01", 3, "2019-11-01"},
		{"2020-02-28", 3, "2019-11-28"},
		{"2020-02-29", 3, "2019-11-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).PreMonths(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).PreMonths(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2020-02-01"},
		{"2020-01-31", "2020-03-02"},
		{"2020-02-01", "2020-03-01"},
		{"2020-02-28", "2020-03-28"},
		{"2020-02-29", "2020-03-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddMonth().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddMonth().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_NextMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2020-02-01"},
		{"2020-01-31", "2020-02-29"},
		{"2020-02-01", "2020-03-01"},
		{"2020-02-28", "2020-03-28"},
		{"2020-02-29", "2020-03-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).NextMonth().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).NextMonth().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-12-01"},
		{"2020-01-31", "2019-12-31"},
		{"2020-02-01", "2020-01-01"},
		{"2020-02-28", "2020-01-28"},
		{"2020-02-29", "2020-01-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubMonth().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubMonth().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_PreMonth(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-12-01"},
		{"2020-01-31", "2019-12-31"},
		{"2020-02-01", "2020-01-01"},
		{"2020-02-28", "2020-01-28"},
		{"2020-02-29", "2020-01-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).PreMonth().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).PreMonth().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddWeeks(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		weeks  int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2020-01-22"},
		{"2020-01-31", 3, "2020-02-21"},
		{"2020-02-01", 3, "2020-02-22"},
		{"2020-02-28", 3, "2020-03-20"},
		{"2020-02-29", 3, "2020-03-21"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddWeeks(v.weeks).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddWeeks(v.weeks).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubWeeks(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		weeks  int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2019-12-11"},
		{"2020-01-31", 3, "2020-01-10"},
		{"2020-02-01", 3, "2020-01-11"},
		{"2020-02-28", 3, "2020-02-07"},
		{"2020-02-29", 3, "2020-02-08"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubWeeks(v.weeks).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubWeeks(v.weeks).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddWeek(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2020-01-08"},
		{"2020-01-31", "2020-02-07"},
		{"2020-02-01", "2020-02-08"},
		{"2020-02-28", "2020-03-06"},
		{"2020-02-29", "2020-03-07"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddWeek().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddWeek().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubWeek(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-12-25"},
		{"2020-01-31", "2020-01-24"},
		{"2020-02-01", "2020-01-25"},
		{"2020-02-28", "2020-02-21"},
		{"2020-02-29", "2020-02-22"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubWeek().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubWeek().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddDays(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		days   int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2020-01-04"},
		{"2020-01-31", 3, "2020-02-03"},
		{"2020-02-01", 3, "2020-02-04"},
		{"2020-02-28", 3, "2020-03-02"},
		{"2020-02-29", 3, "2020-03-03"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddDays(v.days).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddDays(v.days).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubDays(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		days   int
		output string // 期望输出值
	}{
		{"2020-01-01", 3, "2019-12-29"},
		{"2020-01-31", 3, "2020-01-28"},
		{"2020-02-01", 3, "2020-01-29"},
		{"2020-02-28", 3, "2020-02-25"},
		{"2020-02-29", 3, "2020-02-26"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubDays(v.days).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubDays(v.days).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddDay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2020-01-02"},
		{"2020-01-31", "2020-02-01"},
		{"2020-02-01", "2020-02-02"},
		{"2020-02-28", "2020-02-29"},
		{"2020-02-29", "2020-03-01"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddDay().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddDay().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubDay(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01", "2019-12-31"},
		{"2020-01-31", "2020-01-30"},
		{"2020-02-01", "2020-01-31"},
		{"2020-02-28", "2020-02-27"},
		{"2020-02-29", "2020-02-28"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubDay().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubDay().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddHours(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		hours  int
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 16:14:15"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 16:14:15"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 16:14:15"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 16:14:15"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 16:14:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddHours(v.hours).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddHours(v.hours).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubHours(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		hours  int
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 10:14:15"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 10:14:15"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 10:14:15"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 10:14:15"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 10:14:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubHours(v.hours).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubHours(v.hours).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddHour(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 14:14:15"},
		{"2020-01-31 13:14:15", "2020-01-31 14:14:15"},
		{"2020-02-01 13:14:15", "2020-02-01 14:14:15"},
		{"2020-02-28 13:14:15", "2020-02-28 14:14:15"},
		{"2020-02-29 13:14:15", "2020-02-29 14:14:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddHour().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddHour().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubHour(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 12:14:15"},
		{"2020-01-31 13:14:15", "2020-01-31 12:14:15"},
		{"2020-02-01 13:14:15", "2020-02-01 12:14:15"},
		{"2020-02-28 13:14:15", "2020-02-28 12:14:15"},
		{"2020-02-29 13:14:15", "2020-02-29 12:14:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubHour().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubHour().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddMinutes(t *testing.T) {
	Tests := []struct {
		input   string // 输入值
		minutes int
		output  string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 13:17:15"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 13:17:15"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 13:17:15"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 13:17:15"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 13:17:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddMinutes(v.minutes).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddMinutes(v.minutes).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubMinutes(t *testing.T) {
	Tests := []struct {
		input   string // 输入值
		minutes int
		output  string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 13:11:15"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 13:11:15"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 13:11:15"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 13:11:15"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 13:11:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubMinutes(v.minutes).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubMinutes(v.minutes).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddMinute(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 13:15:15"},
		{"2020-01-31 13:14:15", "2020-01-31 13:15:15"},
		{"2020-02-01 13:14:15", "2020-02-01 13:15:15"},
		{"2020-02-28 13:14:15", "2020-02-28 13:15:15"},
		{"2020-02-29 13:14:15", "2020-02-29 13:15:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddMinute().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddMinute().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubMinute(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 13:13:15"},
		{"2020-01-31 13:14:15", "2020-01-31 13:13:15"},
		{"2020-02-01 13:14:15", "2020-02-01 13:13:15"},
		{"2020-02-28 13:14:15", "2020-02-28 13:13:15"},
		{"2020-02-29 13:14:15", "2020-02-29 13:13:15"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubMinute().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubMinute().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddSeconds(t *testing.T) {
	Tests := []struct {
		input   string // 输入值
		seconds int
		output  string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 13:14:18"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 13:14:18"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 13:14:18"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 13:14:18"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 13:14:18"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddSeconds(v.seconds).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddSeconds(v.seconds).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubSeconds(t *testing.T) {
	Tests := []struct {
		input   string // 输入值
		seconds int
		output  string // 期望输出值
	}{
		{"2020-01-01 13:14:15", 3, "2020-01-01 13:14:12"},
		{"2020-01-31 13:14:15", 3, "2020-01-31 13:14:12"},
		{"2020-02-01 13:14:15", 3, "2020-02-01 13:14:12"},
		{"2020-02-28 13:14:15", 3, "2020-02-28 13:14:12"},
		{"2020-02-29 13:14:15", 3, "2020-02-29 13:14:12"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubSeconds(v.seconds).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubSeconds(v.seconds).ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddSecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 13:14:16"},
		{"2020-01-31 13:14:15", "2020-01-31 13:14:16"},
		{"2020-02-01 13:14:15", "2020-02-01 13:14:16"},
		{"2020-02-28 13:14:15", "2020-02-28 13:14:16"},
		{"2020-02-29 13:14:15", "2020-02-29 13:14:16"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddSecond().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).AddSecond().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubSecond(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-01 13:14:15", "2020-01-01 13:14:14"},
		{"2020-01-31 13:14:15", "2020-01-31 13:14:14"},
		{"2020-02-01 13:14:15", "2020-02-01 13:14:14"},
		{"2020-02-28 13:14:15", "2020-02-28 13:14:14"},
		{"2020-02-29 13:14:15", "2020-02-29 13:14:14"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubSecond().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input).SubSecond().ToDateTimeString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddMonthsNoOverflow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		months int
		output string // 期望输出值
	}{
		{"2020-01-31", 1, "2020-02-29"},
		{"2019-12-30", 2, "2020-02-29"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddMonthsNoOverflow(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_AddMonthNoOverflow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-01-30", "2020-02-29"},
		{"2019-05-31", "2019-06-30"},
	}

	for _, v := range Tests {
		output := Parse(v.input).AddMonthNoOverflow().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubMonthsNoOverflow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		months int
		output string // 期望输出值
	}{
		{"2020-03-31", 1, "2020-02-29"},
		{"2020-10-31", 6, "2020-04-30"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubMonthsNoOverflow(v.months).ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}

func TestCarbon_SubMonthNoOverflow(t *testing.T) {
	Tests := []struct {
		input  string // 输入值
		output string // 期望输出值
	}{
		{"2020-03-30", "2020-02-29"},
		{"2020-05-31", "2020-04-30"},
	}

	for _, v := range Tests {
		output := Parse(v.input).SubMonthNoOverflow().ToDateString()

		if output != v.output {
			t.Fatalf("Input %s, expected %s, but got %s\n", v.input, v.output, output)
		}
	}
}
