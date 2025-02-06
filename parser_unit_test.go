package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Parse(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("now"),
			want:   Now().ToDateTimeString(PRC),
		},
		{
			name:   "case3",
			carbon: Parse("yesterday"),
			want:   Yesterday().ToDateTimeString(PRC),
		},
		{
			name:   "case4",
			carbon: Parse("tomorrow"),
			want:   Tomorrow().ToDateTimeString(PRC),
		},
		{
			name:   "case5",
			carbon: Parse("2020-8-5"),
			want:   "2020-08-05 00:00:00",
		},
		{
			name:   "case6",
			carbon: Parse("2020-8-05"),
			want:   "2020-08-05 00:00:00",
		},
		{
			name:   "case7",
			carbon: Parse("2020-08-05"),
			want:   "2020-08-05 00:00:00",
		},
		{
			name:   "case8",
			carbon: Parse("2020-8-5 1:2:3"),
			want:   "2020-08-05 01:02:03",
		},
		{
			name:   "case9",
			carbon: Parse("2020-08-05 1:2:03"),
			want:   "2020-08-05 01:02:03",
		},
		{
			name:   "case10",
			carbon: Parse("2020-08-05 1:02:03"),
			want:   "2020-08-05 01:02:03",
		},
		{
			name:   "case11",
			carbon: Parse("2020-08-05 01:02:03"),
			want:   "2020-08-05 01:02:03",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateTimeString(PRC), "ParseByFormat()")
		})
	}
}

func TestCarbon_ParseByFormat(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: ParseByFormat("", ""),
			want:   "",
		},
		{
			name:   "case2",
			carbon: ParseByFormat("2020|08|05 13:14:15", "Y|m|d H:i:s"),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "case3",
			carbon: ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s"),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "case4",
			carbon: ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒"),
			want:   "2020-08-05 13:14:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateTimeString(PRC), "ParseByFormat()")
		})
	}
}

func TestCarbon_ParseByLayout(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: ParseByLayout("", ""),
			want:   "",
		},
		{
			name:   "case2",
			carbon: ParseByLayout("2020|08|05 13:14:15", "2006|01|02 15:04:05"),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "case3",
			carbon: ParseByLayout("2020|08|05 13:14:15", "2006|01|02 15:04:05"),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "case4",
			carbon: ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒"),
			want:   "2020-08-05 13:14:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateTimeString(PRC), "ParseByLayout()")
		})
	}
}

func TestError_ParseByLayout(t *testing.T) {
	assert.Error(t, ParseByLayout("2020-08-05", "2006-03-04", "xxx").Error, "It should catch an exception in ParseByLayout")
	assert.Error(t, ParseByLayout("xxx", "2006-03-04", PRC).Error, "It should catch an exception in ParseByLayout")
}

func TestError_ParseByFormat(t *testing.T) {
	assert.Error(t, ParseByFormat("2020-08-05", "Y-m-d", "xxx").Error, "It should catch an exception in ParseByFormat()")
	assert.Error(t, ParseByFormat("xxx", "Y-m-d", PRC).Error, "It should catch an exception in ParseByFormat")
}

// https://github.com/dromara/carbon/issues/202
func TestCarbon_Issue202(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("2023-01-08T09:02:48"),
			want:   "2023-01-08 09:02:48 +0800 CST",
		},
		{
			name:   "case2",
			carbon: Parse("2023-1-8T09:02:48"),
			want:   "2023-01-08 09:02:48 +0800 CST",
		},
		{
			name:   "case3",
			carbon: Parse("2023-01-08T9:2:48"),
			want:   "2023-01-08 09:02:48 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2023-01-8T9:2:48"),
			want:   "2023-01-08 09:02:48 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "Parse()")
		})
	}
}

// https://github.com/dromara/carbon/issues/206
func TestCarbon_Issue206(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: ParseByFormat("1699677240", "U"),
			want:   "2023-11-11 12:34:00 +0800 CST",
		},
		{
			name:   "case2",
			carbon: ParseByFormat("1699677240666", "V"),
			want:   "2023-11-11 12:34:00.666 +0800 CST",
		},
		{
			name:   "case3",
			carbon: ParseByFormat("1699677240666666", "X"),
			want:   "2023-11-11 12:34:00.666666 +0800 CST",
		},
		{
			name:   "case4",
			carbon: ParseByFormat("1699677240666666666", "Z"),
			want:   "2023-11-11 12:34:00.666666666 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "Parse()")
		})
	}
}

// https://github.com/dromara/carbon/issues/232
func TestCarbon_Issue232(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("0000-01-01 00:00:00"),
			want:   "0000-01-01 00:00:00",
		},
		{
			name:   "case2",
			carbon: Parse("0001-00-00 00:00:00"),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("0001-01-01 00:00:00"),
			want:   "0001-01-01 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateTimeString(PRC), "Parse()")
		})
	}
}
