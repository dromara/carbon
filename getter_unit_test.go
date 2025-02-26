package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_StdTime(t *testing.T) {
	expected := time.Now().Format(DateTimeLayout)
	actual := Now().StdTime().Format(DateTimeLayout)
	assert.Equal(t, expected, actual)
}

func TestCarbon_DaysInYear(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05"),
			want:   366,
		},
		{
			name:   "case3",
			carbon: Parse("2021-08-05"),
			want:   365,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.DaysInYear(), "DaysInYear()")
		})
	}
}

func TestCarbon_DaysInMonth(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-05"),
			want:   31,
		},
		{
			name:   "case3",
			carbon: Parse("2020-02-05"),
			want:   29,
		},
		{
			name:   "case4",
			carbon: Parse("2020-03-05"),
			want:   31,
		},
		{
			name:   "case5",
			carbon: Parse("2020-04-05"),
			want:   30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.DaysInMonth(), "DaysInMonth()")
		})
	}
}

func TestCarbon_MonthOfYear(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-05"),
			want:   1,
		},
		{
			name:   "case3",
			carbon: Parse("2020-02-05"),
			want:   2,
		},
		{
			name:   "case4",
			carbon: Parse("2020-03-05"),
			want:   3,
		},
		{
			name:   "case5",
			carbon: Parse("2020-04-05"),
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.MonthOfYear(), "MonthOfYear()")
		})
	}
}

func TestCarbon_DayOfYear(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-05"),
			want:   5,
		},
		{
			name:   "case3",
			carbon: Parse("2020-02-05"),
			want:   36,
		},
		{
			name:   "case4",
			carbon: Parse("2020-03-05"),
			want:   65,
		},
		{
			name:   "case5",
			carbon: Parse("2020-04-05"),
			want:   96,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.DayOfYear(), "DayOfYear()")
		})
	}
}

func TestCarbon_DayOfMonth(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01"),
			want:   1,
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-05"),
			want:   5,
		},
		{
			name:   "case5",
			carbon: Parse("2020-01-31"),
			want:   31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.DayOfMonth(), "DayOfMonth()")
		})
	}
}

func TestCarbon_DayOfWeek(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-03"),
			want:   1,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-04"),
			want:   2,
		},
		{
			name:   "case5",
			carbon: Parse("2020-08-05"),
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.DayOfWeek(), "DayOfWeek()")
		})
	}
}

func TestCarbon_WeekOfYear(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Parse("2021-01-01"),
			want:   53,
		},
		{
			name:   "case4",
			carbon: Parse("2021-02-01"),
			want:   5,
		},
		{
			name:   "case5",
			carbon: Parse("2021-03-01"),
			want:   9,
		},
		{
			name:   "case6",
			carbon: Parse("2021-04-01"),
			want:   13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.WeekOfYear(), "WeekOfYear()")
		})
	}
}

func TestCarbon_WeekOfMonth(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Parse("2021-07-01"),
			want:   1,
		},
		{
			name:   "case4",
			carbon: Parse("2021-07-02"),
			want:   1,
		},
		{
			name:   "case5",
			carbon: Parse("2021-07-03"),
			want:   1,
		},
		{
			name:   "case6",
			carbon: Parse("2021-07-04"),
			want:   1,
		},
		{
			name:   "case7",
			carbon: Parse("2021-07-05"),
			want:   2,
		},
		{
			name:   "case8",
			carbon: Parse("2021-07-06"),
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.WeekOfMonth(), "WeekOfMonth()")
		})
	}
}

func TestCarbon_DateTime(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   struct {
			year, month, day, hour, minute, second int
		}
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   struct{ year, month, day, hour, minute, second int }{year: 0, month: 0, day: 0},
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   struct{ year, month, day, hour, minute, second int }{year: 0, month: 0, day: 0},
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01"),
			want:   struct{ year, month, day, hour, minute, second int }{year: 2020, month: 1, day: 1},
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01 13:14:15.999999999"),
			want:   struct{ year, month, day, hour, minute, second int }{year: 2020, month: 1, day: 1, hour: 13, minute: 14, second: 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			year, month, day, hour, minute, second := tt.carbon.DateTime()
			assert.Equalf(t, tt.want.year, year, "DateTime()")
			assert.Equalf(t, tt.want.month, month, "DateTime()")
			assert.Equalf(t, tt.want.day, day, "DateTime()")
			assert.Equalf(t, tt.want.hour, hour, "DateTime()")
			assert.Equalf(t, tt.want.minute, minute, "DateTime()")
			assert.Equalf(t, tt.want.second, second, "DateTime()")
		})
	}
}

func TestCarbon_DateTimeMilli(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   struct {
			year, month, day, hour, minute, second, millisecond int
		}
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   struct{ year, month, day, hour, minute, second, millisecond int }{year: 0, month: 0, day: 0, millisecond: 0},
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   struct{ year, month, day, hour, minute, second, millisecond int }{year: 0, month: 0, day: 0, millisecond: 0},
		},
		{
			name:   "case3",
			carbon: Parse("xxx"),
			want:   struct{ year, month, day, hour, minute, second, millisecond int }{year: 0, month: 0, day: 0, millisecond: 0},
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01"),
			want:   struct{ year, month, day, hour, minute, second, millisecond int }{year: 2020, month: 1, day: 1, millisecond: 0},
		},
		{
			name:   "case5",
			carbon: Parse("2020-01-01 13:14:15.999999999"),
			want:   struct{ year, month, day, hour, minute, second, millisecond int }{year: 2020, month: 1, day: 1, hour: 13, minute: 14, second: 15, millisecond: 999},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			year, month, day, hour, minute, second, millisecond := tt.carbon.DateTimeMilli()
			assert.Equalf(t, tt.want.year, year, "DateTimeMilli()")
			assert.Equalf(t, tt.want.month, month, "DateTimeMilli()")
			assert.Equalf(t, tt.want.day, day, "DateTimeMilli()")
			assert.Equalf(t, tt.want.hour, hour, "DateTimeMilli()")
			assert.Equalf(t, tt.want.minute, minute, "DateTimeMilli()")
			assert.Equalf(t, tt.want.second, second, "DateTimeMilli()")
			assert.Equalf(t, tt.want.millisecond, millisecond, "DateTimeMilli()")
		})
	}
}

func TestCarbon_DateTimeMicro(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   struct {
			year, month, day, hour, minute, second, microsecond int
		}
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   struct{ year, month, day, hour, minute, second, microsecond int }{year: 0, month: 0, day: 0, microsecond: 0},
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   struct{ year, month, day, hour, minute, second, microsecond int }{year: 0, month: 0, day: 0, microsecond: 0},
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01"),
			want:   struct{ year, month, day, hour, minute, second, microsecond int }{year: 2020, month: 1, day: 1, microsecond: 0},
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01 13:14:15.999999999"),
			want:   struct{ year, month, day, hour, minute, second, microsecond int }{year: 2020, month: 1, day: 1, hour: 13, minute: 14, second: 15, microsecond: 999999},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			year, month, day, hour, minute, second, microsecond := tt.carbon.DateTimeMicro()
			assert.Equalf(t, tt.want.year, year, "DateTimeMicro()")
			assert.Equalf(t, tt.want.month, month, "DateTimeMicro()")
			assert.Equalf(t, tt.want.day, day, "DateTimeMicro()")
			assert.Equalf(t, tt.want.hour, hour, "DateTimeMicro()")
			assert.Equalf(t, tt.want.minute, minute, "DateTimeMicro()")
			assert.Equalf(t, tt.want.second, second, "DateTimeMicro()")
			assert.Equalf(t, tt.want.microsecond, microsecond, "DateTimeMicro()")
		})
	}
}

func TestCarbon_DateTimeNano(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   struct {
			year, month, day, hour, minute, second, nanosecond int
		}
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   struct{ year, month, day, hour, minute, second, nanosecond int }{year: 0, month: 0, day: 0, nanosecond: 0},
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   struct{ year, month, day, hour, minute, second, nanosecond int }{year: 0, month: 0, day: 0, nanosecond: 0},
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01"),
			want:   struct{ year, month, day, hour, minute, second, nanosecond int }{year: 2020, month: 1, day: 1, nanosecond: 0},
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01 13:14:15.999999999"),
			want:   struct{ year, month, day, hour, minute, second, nanosecond int }{year: 2020, month: 1, day: 1, hour: 13, minute: 14, second: 15, nanosecond: 999999999},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			year, month, day, hour, minute, second, nanosecond := tt.carbon.DateTimeNano()
			assert.Equalf(t, tt.want.year, year, "DateTimeNano()")
			assert.Equalf(t, tt.want.month, month, "DateTimeNano()")
			assert.Equalf(t, tt.want.day, day, "DateTimeNano()")
			assert.Equalf(t, tt.want.hour, hour, "DateTimeNano()")
			assert.Equalf(t, tt.want.minute, minute, "DateTimeNano()")
			assert.Equalf(t, tt.want.second, second, "DateTimeNano()")
			assert.Equalf(t, tt.want.nanosecond, nanosecond, "DateTimeNano()")
		})
	}
}

func TestCarbon_Date(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   struct {
			year, month, day int
		}
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   struct{ year, month, day int }{year: 0, month: 0, day: 0},
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   struct{ year, month, day int }{year: 2020, month: 1, day: 1},
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01 13:14:15.999999999"),
			want:   struct{ year, month, day int }{year: 2020, month: 1, day: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			year, month, day := tt.carbon.Date()
			assert.Equalf(t, tt.want.year, year, "Date()")
			assert.Equalf(t, tt.want.month, month, "Date()")
			assert.Equalf(t, tt.want.day, day, "Date()")
		})
	}
}

func TestCarbon_DateMilli(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   struct {
			year, month, day, millisecond int
		}
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   struct{ year, month, day, millisecond int }{year: 0, month: 0, day: 0, millisecond: 0},
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   struct{ year, month, day, millisecond int }{year: 2020, month: 1, day: 1, millisecond: 0},
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01 13:14:15.999999999"),
			want:   struct{ year, month, day, millisecond int }{year: 2020, month: 1, day: 1, millisecond: 999},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			year, month, day, millisecond := tt.carbon.DateMilli()
			assert.Equalf(t, tt.want.year, year, "DateMilli()")
			assert.Equalf(t, tt.want.month, month, "DateMilli()")
			assert.Equalf(t, tt.want.day, day, "DateMilli()")
			assert.Equalf(t, tt.want.millisecond, millisecond, "DateMilli()")
		})
	}
}

func TestCarbon_DateMicro(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   struct {
			year, month, day, microsecond int
		}
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   struct{ year, month, day, microsecond int }{year: 0, month: 0, day: 0, microsecond: 0},
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   struct{ year, month, day, microsecond int }{year: 2020, month: 1, day: 1, microsecond: 0},
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01 13:14:15.999999999"),
			want:   struct{ year, month, day, microsecond int }{year: 2020, month: 1, day: 1, microsecond: 999999},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			year, month, day, microsecond := tt.carbon.DateMicro()
			assert.Equalf(t, tt.want.year, year, "DateMicro()")
			assert.Equalf(t, tt.want.month, month, "DateMicro()")
			assert.Equalf(t, tt.want.day, day, "DateMicro()")
			assert.Equalf(t, tt.want.microsecond, microsecond, "DateMicro()")
		})
	}
}

func TestCarbon_DateNano(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   struct {
			year, month, day, nanosecond int
		}
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   struct{ year, month, day, nanosecond int }{year: 0, month: 0, day: 0, nanosecond: 0},
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   struct{ year, month, day, nanosecond int }{year: 2020, month: 1, day: 1, nanosecond: 0},
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01 13:14:15.999999999"),
			want:   struct{ year, month, day, nanosecond int }{year: 2020, month: 1, day: 1, nanosecond: 999999999},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			year, month, day, nanosecond := tt.carbon.DateNano()
			assert.Equalf(t, tt.want.year, year, "DateNano()")
			assert.Equalf(t, tt.want.month, month, "DateNano()")
			assert.Equalf(t, tt.want.day, day, "DateNano()")
			assert.Equalf(t, tt.want.nanosecond, nanosecond, "DateNano()")
		})
	}
}

func TestCarbon_Time(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   struct {
			hour, minute, second int
		}
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   struct{ hour, minute, second int }{hour: 0, minute: 0, second: 0},
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   struct{ hour, minute, second int }{hour: 0, minute: 0, second: 0},
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01 13:14:15.999999999"),
			want:   struct{ hour, minute, second int }{hour: 13, minute: 14, second: 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hour, minute, second := tt.carbon.Time()
			assert.Equalf(t, tt.want.hour, hour, "Time()")
			assert.Equalf(t, tt.want.minute, minute, "Time()")
			assert.Equalf(t, tt.want.second, second, "Time()")
		})
	}
}

func TestCarbon_TimeMilli(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   struct {
			hour, minute, second, millisecond int
		}
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   struct{ hour, minute, second, millisecond int }{hour: 0, minute: 0, second: 0, millisecond: 0},
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   struct{ hour, minute, second, millisecond int }{hour: 0, minute: 0, second: 0, millisecond: 0},
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01 13:14:15.999999999"),
			want:   struct{ hour, minute, second, millisecond int }{hour: 13, minute: 14, second: 15, millisecond: 999},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hour, minute, second, millisecond := tt.carbon.TimeMilli()
			assert.Equalf(t, tt.want.hour, hour, "TimeMilli()")
			assert.Equalf(t, tt.want.minute, minute, "TimeMilli()")
			assert.Equalf(t, tt.want.second, second, "TimeMilli()")
			assert.Equalf(t, tt.want.millisecond, millisecond, "TimeMilli()")
		})
	}
}

func TestCarbon_TimeMicro(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   struct {
			hour, minute, second, microsecond int
		}
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   struct{ hour, minute, second, microsecond int }{hour: 0, minute: 0, second: 0, microsecond: 0},
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   struct{ hour, minute, second, microsecond int }{hour: 0, minute: 0, second: 0, microsecond: 0},
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01 13:14:15.999999999"),
			want:   struct{ hour, minute, second, microsecond int }{hour: 13, minute: 14, second: 15, microsecond: 999999},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hour, minute, second, microsecond := tt.carbon.TimeMicro()
			assert.Equalf(t, tt.want.hour, hour, "TimeMicro()")
			assert.Equalf(t, tt.want.minute, minute, "TimeMicro()")
			assert.Equalf(t, tt.want.second, second, "TimeMicro()")
			assert.Equalf(t, tt.want.microsecond, microsecond, "TimeMicro()")
		})
	}
}

func TestCarbon_TimeNano(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   struct {
			hour, minute, second, nanosecond int
		}
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   struct{ hour, minute, second, nanosecond int }{hour: 0, minute: 0, second: 0, nanosecond: 0},
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   struct{ hour, minute, second, nanosecond int }{hour: 0, minute: 0, second: 0, nanosecond: 0},
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01 13:14:15.999999999"),
			want:   struct{ hour, minute, second, nanosecond int }{hour: 13, minute: 14, second: 15, nanosecond: 999999999},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hour, minute, second, nanosecond := tt.carbon.TimeNano()
			assert.Equalf(t, tt.want.hour, hour, "TimeNano()")
			assert.Equalf(t, tt.want.minute, minute, "TimeNano()")
			assert.Equalf(t, tt.want.second, second, "TimeNano()")
			assert.Equalf(t, tt.want.nanosecond, nanosecond, "TimeNano()")
		})
	}
}

func TestCarbon_Century(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("1990-08-05"),
			want:   20,
		},
		{
			name:   "case3",
			carbon: Parse("2021-08-05"),
			want:   21,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Century(), "Century()")
		})
	}
}

func TestCarbon_Decade(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("2010-08-05"),
			want:   10,
		},
		{
			name:   "case3",
			carbon: Parse("2011-08-05"),
			want:   10,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05"),
			want:   20,
		},
		{
			name:   "case5",
			carbon: Parse("2021-08-05"),
			want:   20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Decade(), "Decade()")
		})
	}
}

func TestCarbon_Year(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05"),
			want:   2020,
		},
		{
			name:   "case3",
			carbon: Parse("2021-08-05"),
			want:   2021,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Year(), "Year()")
		})
	}
}

func TestCarbon_Quarter(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-05"),
			want:   1,
		},
		{
			name:   "case3",
			carbon: Parse("2020-04-05"),
			want:   2,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05"),
			want:   3,
		},
		{
			name:   "case5",
			carbon: Parse("2020-11-05"),
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Quarter(), "Quarter()")
		})
	}
}

func TestCarbon_Month(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05"),
			want:   8,
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15"),
			want:   8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Month(), "Month()")
		})
	}
}

func TestCarbon_Week(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   -1,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   -1,
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-03"),
			want:   1,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-04"),
			want:   2,
		},
		{
			name:   "case5",
			carbon: Parse("2020-08-05"),
			want:   3,
		},
		{
			name:   "case6",
			carbon: Parse("2020-08-09"),
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Week(), "Week()")
		})
	}
}

func TestCarbon_Day(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05"),
			want:   5,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05 13:14:15"),
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Day(), "Day()")
		})
	}
}

func TestCarbon_Hour(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05"),
			want:   0,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05 13:14:15"),
			want:   13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Hour(), "Hour()")
		})
	}
}

func TestCarbon_Minute(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05"),
			want:   0,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05 13:14:15"),
			want:   14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Minute(), "Minute()")
		})
	}
}

func TestCarbon_Second(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05"),
			want:   0,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05 13:14:15"),
			want:   15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Second(), "Second()")
		})
	}
}

func TestCarbon_Millisecond(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15"),
			want:   0,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05 13:14:15.999"),
			want:   999,
		},
		{
			name:   "case5",
			carbon: Parse("2020-08-05 13:14:15.999999"),
			want:   999,
		},
		{
			name:   "case6",
			carbon: Parse("2020-08-05 13:14:15.999999999"),
			want:   999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Millisecond(), "Millisecond()")
		})
	}
}

func TestCarbon_Microsecond(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15"),
			want:   0,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05 13:14:15.999"),
			want:   999000,
		},
		{
			name:   "case5",
			carbon: Parse("2020-08-05 13:14:15.999999"),
			want:   999999,
		},
		{
			name:   "case6",
			carbon: Parse("2020-08-05 13:14:15.999999999"),
			want:   999999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Microsecond(), "Microsecond()")
		})
	}
}

func TestCarbon_Nanosecond(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15"),
			want:   0,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05 13:14:15.999"),
			want:   999000000,
		},
		{
			name:   "case5",
			carbon: Parse("2020-08-05 13:14:15.999999"),
			want:   999999000,
		},
		{
			name:   "case6",
			carbon: Parse("2020-08-05 13:14:15.999999999"),
			want:   999999999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Nanosecond(), "Nanosecond()")
		})
	}
}

func TestCarbon_Timestamp(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int64
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01 13:14:15"),
			want:   1577855655,
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-31 13:14:15"),
			want:   1580447655,
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-01 13:14:15"),
			want:   1580534055,
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-28 13:14:15"),
			want:   1582866855,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Timestamp(), "Timestamp()")
		})
	}
}

func TestCarbon_TimestampMilli(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int64
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01 13:14:15"),
			want:   1577855655000,
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-31 13:14:15"),
			want:   1580447655000,
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-01 13:14:15"),
			want:   1580534055000,
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-28 13:14:15"),
			want:   1582866855000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.TimestampMilli(), "TimestampMilli()")
		})
	}
}

func TestCarbon_TimestampMicro(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int64
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01 13:14:15"),
			want:   1577855655000000,
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-31 13:14:15"),
			want:   1580447655000000,
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-01 13:14:15"),
			want:   1580534055000000,
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-28 13:14:15"),
			want:   1582866855000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.TimestampMicro(), "TimestampMicro()")
		})
	}
}

func TestCarbon_TimestampNano(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int64
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01 13:14:15"),
			want:   1577855655000000000,
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-31 13:14:15"),
			want:   1580447655000000000,
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-01 13:14:15"),
			want:   1580534055000000000,
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-28 13:14:15"),
			want:   1582866855000000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.TimestampNano(), "TimestampNano()")
		})
	}
}

func TestCarbon_Timezone(t *testing.T) {
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
			carbon: Parse("xxx"),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Now(PRC),
			want:   "CST",
		},
		{
			name:   "case4",
			carbon: Now(Tokyo),
			want:   "JST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Timezone(), "Timezone()")
		})
	}
}

func TestCarbon_Location(t *testing.T) {
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
			carbon: Parse("xxx"),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Now(PRC),
			want:   PRC,
		},
		{
			name:   "case4",
			carbon: Now(Tokyo),
			want:   Tokyo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Location(), "Location()")
		})
	}
}

func TestCarbon_Offset(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Now(PRC),
			want:   28800,
		},
		{
			name:   "case4",
			carbon: Now(Tokyo),
			want:   32400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Offset(), "Offset()")
		})
	}
}

func TestCarbon_Locale(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Now().SetLocale(""),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Now().SetLocale("en"),
			want:   "en",
		},
		{
			name:   "case3",
			carbon: Now().SetLocale("zh-CN"),
			want:   "zh-CN",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Locale(), "Locale()")
		})
	}
}

func TestCarbon_Age(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   int
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   0,
		},
		{
			name:   "case2",
			carbon: Parse("xxx"),
			want:   0,
		},
		{
			name:   "case3",
			carbon: Now().AddYears(18),
			want:   0,
		},
		{
			name:   "case4",
			carbon: Now().SubYears(18),
			want:   18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Age(), "Age()")
		})
	}
}

func TestCarbon_CurrentLayout(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: ParseByLayout("now", DateTimeLayout),
			want:   DateTimeLayout,
		},
		{
			name:   "case2",
			carbon: ParseByLayout("2025-02-26", DateLayout),
			want:   DateLayout,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.CurrentLayout(), "CurrentLayout()")
		})
	}
}
