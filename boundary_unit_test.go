package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_StartOfCentury(t *testing.T) {
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
			carbon: Parse("2020-01-01 00:00:00"),
			want:   "2000-01-01 00:00:00",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-31 23:59:59"),
			want:   "2000-01-01 00:00:00",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-01 13:14:15"),
			want:   "2000-01-01 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.StartOfCentury().ToDateTimeString(), "StartOfCentury()")
		})
	}
}

func TestCarbon_EndOfCentury(t *testing.T) {
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
			carbon: Parse("2020-01-01 00:00:00"),
			want:   "2099-12-31 23:59:59",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-31 23:59:59"),
			want:   "2099-12-31 23:59:59",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-01 13:14:15"),
			want:   "2099-12-31 23:59:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.EndOfCentury().ToDateTimeString(), "EndOfCentury()")
		})
	}
}

func TestCarbon_StartOfDecade(t *testing.T) {
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
			carbon: Parse("2020-01-31 23:59:59"),
			want:   "2020-01-01 00:00:00",
		},
		{
			name:   "case4",
			carbon: Parse("2021-01-01 00:00:00"),
			want:   "2020-01-01 00:00:00",
		},
		{
			name:   "case5",
			carbon: Parse("2029-01-31 23:59:59"),
			want:   "2020-01-01 00:00:00",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-01 13:14:15"),
			want:   "2020-01-01 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.StartOfDecade().ToDateTimeString(), "StartOfDecade()")
		})
	}
}

func TestCarbon_EndOfDecade(t *testing.T) {
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
			carbon: Parse("2020-01-31 23:59:59"),
			want:   "2029-12-31 23:59:59",
		},
		{
			name:   "case4",
			carbon: Parse("2021-01-01 00:00:00"),
			want:   "2029-12-31 23:59:59",
		},
		{
			name:   "case5",
			carbon: Parse("2029-01-31 23:59:59"),
			want:   "2029-12-31 23:59:59",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-01 13:14:15"),
			want:   "2029-12-31 23:59:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.EndOfDecade().ToDateTimeString(), "EndOfDecade()")
		})
	}
}

func TestCarbon_StartOfYear(t *testing.T) {
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
			carbon: Parse("2020-01-01 00:00:00"),
			want:   "2020-01-01 00:00:00",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-31 23:59:59"),
			want:   "2020-01-01 00:00:00",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-01 13:14:15"),
			want:   "2020-01-01 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.StartOfYear().ToDateTimeString(), "StartOfYear()")
		})
	}
}

func TestCarbon_EndOfYear(t *testing.T) {
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
			carbon: Parse("2020-01-01 00:00:00"),
			want:   "2020-12-31 23:59:59",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-31 23:59:59"),
			want:   "2020-12-31 23:59:59",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-01 13:14:15"),
			want:   "2020-12-31 23:59:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.EndOfYear().ToDateTimeString(), "EndOfYear()")
		})
	}
}

func TestCarbon_StartOfQuarter(t *testing.T) {
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
			carbon: Parse("2020-01-15 00:00:00"),
			want:   "2020-01-01 00:00:00",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-15 00:00:00"),
			want:   "2020-01-01 00:00:00",
		},
		{
			name:   "case5",
			carbon: Parse("2020-03-15 00:00:00"),
			want:   "2020-01-01 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.StartOfQuarter().ToDateTimeString(), "StartOfQuarter()")
		})
	}
}

func TestCarbon_EndOfQuarter(t *testing.T) {
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
			carbon: Parse("2020-01-15 00:00:00"),
			want:   "2020-03-31 23:59:59",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-15 00:00:00"),
			want:   "2020-03-31 23:59:59",
		},
		{
			name:   "case5",
			carbon: Parse("2020-03-15 00:00:00"),
			want:   "2020-03-31 23:59:59",
		},
		{
			name:   "case6",
			carbon: Parse("2020-04-15 23:59:59"),
			want:   "2020-06-30 23:59:59",
		},
		{
			name:   "case7",
			carbon: Parse("2020-05-15 23:59:59"),
			want:   "2020-06-30 23:59:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.EndOfQuarter().ToDateTimeString(), "EndOfQuarter()")
		})
	}
}

func TestCarbon_StartOfMonth(t *testing.T) {
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
			carbon: Parse("2020-01-15 00:00:00"),
			want:   "2020-01-01 00:00:00",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-15 00:00:00"),
			want:   "2020-02-01 00:00:00",
		},
		{
			name:   "case5",
			carbon: Parse("2020-03-15 00:00:00"),
			want:   "2020-03-01 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.StartOfMonth().ToDateTimeString(), "StartOfMonth()")
		})
	}
}

func TestCarbon_EndOfMonth(t *testing.T) {
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
			carbon: Parse("2020-01-15 00:00:00"),
			want:   "2020-01-31 23:59:59",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-15 00:00:00"),
			want:   "2020-02-29 23:59:59",
		},
		{
			name:   "case5",
			carbon: Parse("2020-03-15 00:00:00"),
			want:   "2020-03-31 23:59:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.EndOfMonth().ToDateTimeString(), "EndOfMonth()")
		})
	}
}

func TestCarbon_StartOfWeek(t *testing.T) {
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
			carbon: Parse("2021-06-13"),
			want:   "2021-06-13 00:00:00",
		},
		{
			name:   "case4",
			carbon: Parse("2021-06-14"),
			want:   "2021-06-13 00:00:00",
		},
		{
			name:   "case5",
			carbon: Parse("2021-06-18"),
			want:   "2021-06-13 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.StartOfWeek().ToDateTimeString(), "StartOfWeek()")
		})
	}
}

func TestCarbon_EndOfWeek(t *testing.T) {
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
			carbon: Parse("2021-06-13"),
			want:   "2021-06-19 23:59:59",
		},
		{
			name:   "case4",
			carbon: Parse("2021-06-14"),
			want:   "2021-06-19 23:59:59",
		},
		{
			name:   "case5",
			carbon: Parse("2021-06-18"),
			want:   "2021-06-19 23:59:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.EndOfWeek().ToDateTimeString(), "EndOfWeek()")
		})
	}
}

func TestCarbon_StartOfDay(t *testing.T) {
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
			carbon: Parse("2020-01-15 00:00:00"),
			want:   "2020-01-15 00:00:00",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-15 00:00:00"),
			want:   "2020-02-15 00:00:00",
		},
		{
			name:   "case5",
			carbon: Parse("2020-03-15 00:00:00"),
			want:   "2020-03-15 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.StartOfDay().ToDateTimeString(), "StartOfDay()")
		})
	}
}

func TestCarbon_EndOfDay(t *testing.T) {
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
			carbon: Parse("2020-01-15 00:00:00"),
			want:   "2020-01-15 23:59:59",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-15 00:00:00"),
			want:   "2020-02-15 23:59:59",
		},
		{
			name:   "case5",
			carbon: Parse("2020-03-15 00:00:00"),
			want:   "2020-03-15 23:59:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.EndOfDay().ToDateTimeString(), "EndOfDay()")
		})
	}
}

func TestCarbon_StartOfHour(t *testing.T) {
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
			carbon: Parse("2020-01-15 00:00:00"),
			want:   "2020-01-15 00:00:00",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-15 00:00:00"),
			want:   "2020-02-15 00:00:00",
		},
		{
			name:   "case5",
			carbon: Parse("2020-03-15 00:00:00"),
			want:   "2020-03-15 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.StartOfHour().ToDateTimeString(), "StartOfHour()")
		})
	}
}

func TestCarbon_EndOfHour(t *testing.T) {
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
			carbon: Parse(""),
			want:   "",
		},
		{
			name:   "case4",
			carbon: Parse("xxx"),
			want:   "",
		},
		{
			name:   "case5",
			carbon: Parse("2020-01-15 00:00:00"),
			want:   "2020-01-15 00:59:59",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-15 00:00:00"),
			want:   "2020-02-15 00:59:59",
		},
		{
			name:   "case7",
			carbon: Parse("2020-03-15 00:00:00"),
			want:   "2020-03-15 00:59:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.EndOfHour().ToDateTimeString(), "EndOfHour()")
		})
	}
}

func TestCarbon_StartOfMinute(t *testing.T) {
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
			carbon: Parse(""),
			want:   "",
		},
		{
			name:   "case4",
			carbon: Parse("xxx"),
			want:   "",
		},
		{
			name:   "case5",
			carbon: Parse("2020-01-15 00:00:00"),
			want:   "2020-01-15 00:00:00",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-15 00:00:00"),
			want:   "2020-02-15 00:00:00",
		},
		{
			name:   "case7",
			carbon: Parse("2020-03-15 00:00:00"),
			want:   "2020-03-15 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.StartOfMinute().ToDateTimeString(), "StartOfMinute()")
		})
	}
}

func TestCarbon_EndOfMinute(t *testing.T) {
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
			carbon: Parse("2020-01-15 00:00:00"),
			want:   "2020-01-15 00:00:59",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-15 00:00:00"),
			want:   "2020-02-15 00:00:59",
		},
		{
			name:   "case5",
			carbon: Parse("2020-03-15 00:00:00"),
			want:   "2020-03-15 00:00:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.EndOfMinute().ToDateTimeString(), "EndOfMinute()")
		})
	}
}

func TestCarbon_StartOfSecond(t *testing.T) {
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
			carbon: Parse("2020-01-15 00:00:00.999999999", PRC),
			want:   "2020-01-15 00:00:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-15 00:00:00.999999999", PRC),
			want:   "2020-02-15 00:00:00 +0800 CST",
		},
		{
			name:   "case5",
			carbon: Parse("2020-03-15 00:00:00.999999999", PRC),
			want:   "2020-03-15 00:00:00 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.StartOfSecond().ToString(), "StartOfSecond()")
		})
	}
}

func TestCarbon_EndOfSecond(t *testing.T) {
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
			carbon: Parse("2020-01-15 00:00:00.123", PRC),
			want:   "2020-01-15 00:00:00.999999999 +0800 CST",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-15 00:00:00.123", PRC),
			want:   "2020-02-15 00:00:00.999999999 +0800 CST",
		},
		{
			name:   "case5",
			carbon: Parse("2020-03-15 00:00:00.123", PRC),
			want:   "2020-03-15 00:00:00.999999999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.EndOfSecond().ToString(), "EndOfSecond()")
		})
	}
}
