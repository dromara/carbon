package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_SetWeekStartsAt(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetWeekStartsAt(Sunday),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetWeekStartsAt(Sunday),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2021-06-13").SetWeekStartsAt(Sunday),
			want:   "2021-06-13",
		},
		{
			name:   "case4",
			carbon: Parse("2021-06-13").SetWeekStartsAt(Monday),
			want:   "2021-06-07",
		},
		{
			name:   "case5",
			carbon: Parse("2021-06-13").SetWeekStartsAt(Tuesday),
			want:   "2021-06-08",
		},
		{
			name:   "case6",
			carbon: Parse("2021-06-13").SetWeekStartsAt(Wednesday),
			want:   "2021-06-09",
		},
		{
			name:   "case7",
			carbon: Parse("2021-06-13").SetWeekStartsAt(Thursday),
			want:   "2021-06-10",
		},
		{
			name:   "case8",
			carbon: Parse("2021-06-13").SetWeekStartsAt(Friday),
			want:   "2021-06-11",
		},
		{
			name:   "case9",
			carbon: Parse("2021-06-13").SetWeekStartsAt(Saturday),
			want:   "2021-06-12",
		},
		{
			name:   "case10",
			carbon: SetWeekStartsAt(Saturday).Parse("2021-06-13"),
			want:   "2021-06-12",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.StartOfWeek().ToDateString(), "SetWeekStartsAt()")
		})
	}
}

func TestCarbon_SetTimezone(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetTimezone(PRC),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetTimezone(PRC),
			want:   "",
		},
		{
			name:   "case3",
			carbon: SetTimezone(PRC).Parse("2020-08-05 13:14:15"),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "case4",
			carbon: SetTimezone(Tokyo).Parse("2020-08-05 13:14:15"),
			want:   "2020-08-05 12:14:15",
		},
		{
			name:   "case5",
			carbon: SetTimezone(London).Parse("2020-08-05 13:14:15"),
			want:   "2020-08-05 20:14:15",
		},
		{
			name:   "case6",
			carbon: SetTimezone("").Parse("2020-08-05 13:14:15"),
			want:   "",
		},
		{
			name:   "case7",
			carbon: SetTimezone("xxx").Parse("2020-08-05 13:14:15"),
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateTimeString(PRC), "SetTimezone()")
		})
	}
}

func TestCarbon_SetLocation(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetLocation(time.UTC),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetLocation(time.UTC),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15").SetLocation(nil),
			want:   "",
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05 13:14:15").SetLocation(time.UTC),
			want:   "UTC",
		},
		{
			name:   "case5",
			carbon: Parse("2020-08-05 13:14:15").SetLocation(time.UTC),
			want:   "UTC",
		},
		{
			name:   "case6",
			carbon: SetLocation(time.UTC).Parse("2020-08-05 13:14:15"),
			want:   "UTC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Timezone(), "SetLocation()")
		})
	}
}

func TestCarbon_SetLocale(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetLocale("en"),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetLocale("en"),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15").SetLocale("en"),
			want:   "August",
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05 13:14:15").SetLocale("zh-CN"),
			want:   "八月",
		},
		{
			name:   "case5",
			carbon: Parse("2020-08-05 13:14:15").SetLocale("jp"),
			want:   "8月",
		},
		{
			name:   "case6",
			carbon: SetLocale("en").Parse("2020-08-05 13:14:15"),
			want:   "August",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToMonthString(), "SetLocale()")
		})
	}
}

func TestCarbon_SetDateTime(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetDateTime(2019, 02, 02, 13, 14, 15),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetDateTime(2019, 02, 02, 13, 14, 15),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetDateTime(2019, 02, 02, 13, 14, 15),
			want:   "2019-02-02 13:14:15",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01").SetDateTime(2019, 02, 31, 13, 14, 15),
			want:   "2019-03-03 13:14:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateTimeString(PRC), "SetDateTime()")
		})
	}
}

func TestCarbon_SetDateTimeMilli(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetDateTimeMilli(2019, 02, 02, 13, 14, 15, 999),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetDateTimeMilli(2019, 02, 02, 13, 14, 15, 999),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetDateTimeMilli(2019, 02, 02, 13, 14, 15, 999),
			want:   "2019-02-02 13:14:15.999 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01").SetDateTimeMilli(2019, 02, 31, 13, 14, 15, 999),
			want:   "2019-03-03 13:14:15.999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetDateTimeMilli()")
		})
	}
}

func TestCarbon_SetDateTimeMicro(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetDateTimeMicro(2019, 02, 02, 13, 14, 15, 999999),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetDateTimeMicro(2019, 02, 02, 13, 14, 15, 999999),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetDateTimeMicro(2019, 02, 02, 13, 14, 15, 999999),
			want:   "2019-02-02 13:14:15.999999 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01").SetDateTimeMicro(2019, 02, 31, 13, 14, 15, 999999),
			want:   "2019-03-03 13:14:15.999999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetDateTimeMicro()")
		})
	}
}

func TestCarbon_SetDateTimeNano(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetDateTimeNano(2019, 02, 02, 13, 14, 15, 999999999),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetDateTimeNano(2019, 02, 02, 13, 14, 15, 999999999),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetDateTimeNano(2019, 02, 02, 13, 14, 15, 999999999),
			want:   "2019-02-02 13:14:15.999999999 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01").SetDateTimeNano(2019, 02, 31, 13, 14, 15, 999999999),
			want:   "2019-03-03 13:14:15.999999999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetDateTimeNano()")
		})
	}
}

func TestCarbon_SetDate(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetDate(2019, 02, 02),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetDate(2019, 02, 02),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetDate(2019, 02, 02),
			want:   "2019-02-02 00:00:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01").SetDate(2019, 02, 31),
			want:   "2019-03-03 00:00:00 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetDate()")
		})
	}
}

func TestCarbon_SetDateMilli(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetDateMilli(2019, 02, 02, 999),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetDateMilli(2019, 02, 02, 999),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetDateMilli(2019, 02, 02, 999),
			want:   "2019-02-02 00:00:00.999 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01").SetDateMilli(2019, 02, 31, 999),
			want:   "2019-03-03 00:00:00.999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetDateMilli()")
		})
	}
}

func TestCarbon_SetDateMicro(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetDateMicro(2019, 02, 02, 999999),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetDateMicro(2019, 02, 02, 999999),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetDateMicro(2019, 02, 02, 999999),
			want:   "2019-02-02 00:00:00.999999 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01").SetDateMicro(2019, 02, 31, 999999),
			want:   "2019-03-03 00:00:00.999999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetDateMicro()")
		})
	}
}

func TestCarbon_SetDateNano(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetDateNano(2019, 02, 02, 999999999),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetDateNano(2019, 02, 02, 999999999),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetDateNano(2019, 02, 02, 999999999),
			want:   "2019-02-02 00:00:00.999999999 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01").SetDateNano(2019, 02, 31, 999999999),
			want:   "2019-03-03 00:00:00.999999999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetDateNano()")
		})
	}
}

func TestCarbon_SetTime(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetTime(13, 14, 15),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetTime(13, 14, 15),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetTime(13, 14, 15),
			want:   "2020-01-01 13:14:15 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01").SetTime(13, 14, 90),
			want:   "2020-01-01 13:15:30 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetTime()")
		})
	}
}

func TestCarbon_SetTimeMilli(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetTimeMilli(13, 14, 15, 999),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetTimeMilli(13, 14, 15, 999),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetTimeMilli(13, 14, 15, 999),
			want:   "2020-01-01 13:14:15.999 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01").SetTimeMilli(13, 14, 90, 999),
			want:   "2020-01-01 13:15:30.999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetTimeMilli()")
		})
	}
}

func TestCarbon_SetTimeMicro(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetTimeMicro(13, 14, 15, 999999),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetTimeMicro(13, 14, 15, 999999),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetTimeMicro(13, 14, 15, 999999),
			want:   "2020-01-01 13:14:15.999999 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01").SetTimeMicro(13, 14, 90, 999999),
			want:   "2020-01-01 13:15:30.999999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetTimeMicro()")
		})
	}
}

func TestCarbon_SetTimeNano(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetTimeNano(13, 14, 15, 999999999),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetTimeNano(13, 14, 15, 999999999),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetTimeNano(13, 14, 15, 999999999),
			want:   "2020-01-01 13:14:15.999999999 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01").SetTimeNano(13, 14, 90, 999999999),
			want:   "2020-01-01 13:15:30.999999999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetTimeNano()")
		})
	}
}

func TestCarbon_SetYear(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetYear(2019),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetYear(2019),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetYear(2019),
			want:   "2019-01-01 00:00:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-31").SetYear(2019),
			want:   "2019-01-31 00:00:00 +0800 CST",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-29").SetYear(2019),
			want:   "2019-03-01 00:00:00 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetYear()")
		})
	}
}

func TestCarbon_SetYearNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetYearNoOverflow(2019),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetYearNoOverflow(2019),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetYearNoOverflow(2019),
			want:   "2019-01-01 00:00:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-31").SetYearNoOverflow(2019),
			want:   "2019-01-31 00:00:00 +0800 CST",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-29").SetYearNoOverflow(2019),
			want:   "2019-02-28 00:00:00 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetYearNoOverflow()")
		})
	}
}

func TestCarbon_SetMonth(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetMonth(2),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetMonth(2),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetMonth(2),
			want:   "2020-02-01 00:00:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-30").SetMonth(2),
			want:   "2020-03-01 00:00:00 +0800 CST",
		},
		{
			name:   "case5",
			carbon: Parse("2020-01-31").SetMonth(2),
			want:   "2020-03-02 00:00:00 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetMonth()")
		})
	}
}

func TestCarbon_SetMonthNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetMonthNoOverflow(2),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetMonthNoOverflow(2),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetMonthNoOverflow(2),
			want:   "2020-02-01 00:00:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-30").SetMonthNoOverflow(2),
			want:   "2020-02-29 00:00:00 +0800 CST",
		},
		{
			name:   "case5",
			carbon: Parse("2020-01-31").SetMonthNoOverflow(2),
			want:   "2020-02-29 00:00:00 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetMonthNoOverflow()")
		})
	}
}

func TestCarbon_SetDay(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetDay(31),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetDay(31),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetDay(31),
			want:   "2020-01-31 00:00:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SetDay(31),
			want:   "2020-03-02 00:00:00 +0800 CST",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SetDay(31),
			want:   "2020-03-02 00:00:00 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetDay()")
		})
	}
}

func TestCarbon_SetHour(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetHour(10),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetHour(10),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetHour(10),
			want:   "2020-01-01 10:00:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SetHour(24),
			want:   "2020-02-02 00:00:00 +0800 CST",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SetHour(31),
			want:   "2020-02-29 07:00:00 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetHour()")
		})
	}
}

func TestCarbon_SetMinute(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetMinute(10),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetMinute(10),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetMinute(10),
			want:   "2020-01-01 00:10:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SetMinute(24),
			want:   "2020-02-01 00:24:00 +0800 CST",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SetMinute(60),
			want:   "2020-02-28 01:00:00 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetMinute()")
		})
	}
}

func TestCarbon_SetSecond(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetSecond(10),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetSecond(10),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetSecond(10),
			want:   "2020-01-01 00:00:10 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SetSecond(24),
			want:   "2020-02-01 00:00:24 +0800 CST",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SetSecond(60),
			want:   "2020-02-28 00:01:00 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetSecond()")
		})
	}
}

func TestCarbon_SetMillisecond(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetMillisecond(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetMillisecond(0),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetMillisecond(0),
			want:   "2020-01-01 00:00:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SetMillisecond(100),
			want:   "2020-02-01 00:00:00.1 +0800 CST",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SetMillisecond(999),
			want:   "2020-02-28 00:00:00.999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetMillisecond()")
		})
	}
}

func TestCarbon_SetMicrosecond(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetMicrosecond(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetMicrosecond(0),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetMicrosecond(0),
			want:   "2020-01-01 00:00:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SetMicrosecond(100000),
			want:   "2020-02-01 00:00:00.1 +0800 CST",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SetMicrosecond(999999),
			want:   "2020-02-28 00:00:00.999999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetMicrosecond()")
		})
	}
}

func TestCarbon_SetNanosecond(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SetNanosecond(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("xxx").SetNanosecond(0),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SetNanosecond(0),
			want:   "2020-01-01 00:00:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SetNanosecond(100000),
			want:   "2020-02-01 00:00:00.0001 +0800 CST",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SetNanosecond(999999),
			want:   "2020-02-28 00:00:00.000999999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(PRC), "SetNanosecond()")
		})
	}
}
