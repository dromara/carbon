package carbon

import (
	"testing"
)

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
			t.Logf("catch an exception in Parse():%v", output.Error)
		} else if output.ToDateTimeString() != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}

	for _, v := range Tests {
		output := SetTimezone(PRC).Parse(v.input)

		if output.Error != nil {
			t.Logf("catch an exception in Parse():%v", output.Error)
		} else if output.ToDateTimeString() != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}
}

var ParseByFormatTests = []struct {
	input  string // 输入值
	format string // 输入参数
	output string // 期望输出值
}{
	{"2020|08|05 13:14:15", "Y|m|d H:i:s", "2020-08-05 13:14:15"},
	{"It is 2020|08|05 13:14:15", "It is Y|m|d H:i:s", "2020-08-05 13:14:15"},
	{"今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒", "2020-08-05 13:14:15"},
	{"12345678", "XXXX", ""}, // 异常情况
}

func TestCarbon_ParseByFormat1(t *testing.T) {
	for _, v := range ParseByFormatTests {
		output := ParseByFormat(v.input, v.format)

		if output.Error != nil {
			t.Logf("catch an exception in ParseByFormat():%v", output.Error)
		} else if output.ToDateTimeString() != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}
}

func TestCarbon_ParseByFormat2(t *testing.T) {
	for _, v := range ParseByFormatTests {
		output := SetTimezone("XXXX").ParseByFormat(v.input, v.format)

		if output.Error != nil {
			t.Logf("catch an exception in ParseByFormat():%v", output.Error)
		} else if output.ToDateTimeString() != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}
}

var ParseByLayoutTests = []struct {
	input  string // 输入值
	format string // 输入参数
	output string // 期望输出值
}{
	{"2020|08|05 13:14:15", "2006|01|02 15:04:05", "2020-08-05 13:14:15"},
	{"It is 2020|08|05 13:14:15", "It is 2006|01|02 15:04:05", "2020-08-05 13:14:15"},
	{"今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒", "2020-08-05 13:14:15"},
	{"12345678", "XXXX", ""}, // 异常情况
}

func TestCarbon_ParseByLayout1(t *testing.T) {
	for _, v := range ParseByLayoutTests {
		output := ParseByLayout(v.input, v.format)

		if output.Error != nil {
			t.Logf("catch an exception in ParseByLayout():%v", output.Error)
		} else if output.ToDateTimeString() != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}
}

func TestCarbon_ParseByLayout2(t *testing.T) {
	for _, v := range ParseByLayoutTests {
		output := SetTimezone("XXXX").ParseByLayout(v.input, v.format)

		if output.Error != nil {
			t.Logf("catch an exception in ParseByLayout():%v", output.Error)
		} else if output.ToDateTimeString() != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}
}

var ParseByInputTests = []struct {
	input  string
	format string
	output string
}{
	{"2021年 5月 1日 星期六 15时24分25秒 CST", "Y年 m月 d日 星期六 H时i分s秒 CST", "2021-05-01 15:24:25"},
	{"2021年5月1日 星期六 农历三月二十 15:26:30", "Y年m月d日 星期六 农历三月二十 H:i:s", "2021-05-01 15:26:30"},
}

func TestCarbon_ParseByInput(t *testing.T) {
	for _, v := range ParseByInputTests {
		output := ParseByInput(v.input, v.format)

		if output.Error != nil {
			t.Logf("catch an exception in ParseByInput():%v", output.Error)
		} else if output.ToDateTimeString() != v.output {
			t.Errorf("Input %s, expected %s, but got %s\n", v.input, v.output, output.ToDateTimeString())
		}
	}
}