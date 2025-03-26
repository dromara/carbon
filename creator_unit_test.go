package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_CreateFromStdTime(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: CreateFromStdTime(time.Now()),
			want:   time.Now().Format(DateTimeLayout),
		},
		{
			name:   "case2",
			carbon: CreateFromStdTime(time.Now(), PRC),
			want:   time.Now().Format(DateTimeLayout),
		},
		{
			name:   "case3",
			carbon: CreateFromStdTime(time.Now(), "xxx"),
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateTimeString(), "CreateFromStdTime()")
		})
	}
}

func TestCarbon_CreateFromTimestamp(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("xxx").CreateFromTimestamp(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: CreateFromTimestamp(-1),
			want:   "1970-01-01 07:59:59 +0800 CST",
		},
		{
			name:   "case3",
			carbon: CreateFromTimestamp(0),
			want:   "1970-01-01 08:00:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: CreateFromTimestamp(1),
			want:   "1970-01-01 08:00:01 +0800 CST",
		},
		{
			name:   "case5",
			carbon: CreateFromTimestamp(1649735755, PRC),
			want:   "2022-04-12 11:55:55 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "CreateFromTimestamp()")
		})
	}
}

func TestCarbon_CreateFromTimestampMilli(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("xxx").CreateFromTimestampMilli(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: CreateFromTimestampMilli(-1),
			want:   "1970-01-01 07:59:59.999 +0800 CST",
		},
		{
			name:   "case3",
			carbon: CreateFromTimestampMilli(0),
			want:   "1970-01-01 08:00:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: CreateFromTimestampMilli(1),
			want:   "1970-01-01 08:00:00.001 +0800 CST",
		},
		{
			name:   "case5",
			carbon: CreateFromTimestampMilli(1649735755981, PRC),
			want:   "2022-04-12 11:55:55.981 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "CreateFromTimestampMilli()")
		})
	}
}

func TestCarbon_CreateFromTimestampMicro(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("xxx").CreateFromTimestampMicro(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: CreateFromTimestampMicro(-1),
			want:   "1970-01-01 07:59:59.999999 +0800 CST",
		},
		{
			name:   "case3",
			carbon: CreateFromTimestampMicro(0),
			want:   "1970-01-01 08:00:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: CreateFromTimestampMicro(1),
			want:   "1970-01-01 08:00:00.000001 +0800 CST",
		},
		{
			name:   "case5",
			carbon: CreateFromTimestampMicro(1649735755981566, PRC),
			want:   "2022-04-12 11:55:55.981566 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "CreateFromTimestampMicro()")
		})
	}
}

func TestCarbon_CreateFromTimestampNano(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("xxx").CreateFromTimestampNano(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: CreateFromTimestampNano(-1),
			want:   "1970-01-01 07:59:59.999999999 +0800 CST",
		},
		{
			name:   "case3",
			carbon: CreateFromTimestampNano(0),
			want:   "1970-01-01 08:00:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: CreateFromTimestampNano(1, PRC),
			want:   "1970-01-01 08:00:00.000000001 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "CreateFromTimestampNano()")
		})
	}
}

func TestCarbon_CreateFromDateTime(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("xxx").CreateFromDateTime(0, 0, 0, 0, 0, 0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: CreateFromDateTime(2020, 1, 1, 13, 14, 15),
			want:   "2020-01-01 13:14:15",
		},
		{
			name:   "case3",
			carbon: CreateFromDateTime(2020, 1, 31, 13, 14, 15),
			want:   "2020-01-31 13:14:15",
		},
		{
			name:   "case4",
			carbon: CreateFromDateTime(2020, 2, 1, 13, 14, 15, PRC),
			want:   "2020-02-01 13:14:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateTimeString(), "CreateFromDateTime()")
		})
	}
}

func TestCarbon_CreateFromDateTimeMilli(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("xxx").CreateFromDateTimeMilli(0, 0, 0, 0, 0, 0, 0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: CreateFromDateTimeMilli(2020, 1, 1, 13, 14, 15, 999),
			want:   "2020-01-01 13:14:15.999 +0800 CST",
		},
		{
			name:   "case3",
			carbon: CreateFromDateTimeMilli(2020, 1, 31, 13, 14, 15, 999),
			want:   "2020-01-31 13:14:15.999 +0800 CST",
		},
		{
			name:   "case4",
			carbon: CreateFromDateTimeMilli(2020, 2, 1, 13, 14, 15, 999),
			want:   "2020-02-01 13:14:15.999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "CreateFromDateTimeMilli()")
		})
	}
}

func TestCarbon_CreateFromDateTimeMicro(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("xxx").CreateFromDateTimeMicro(0, 0, 0, 0, 0, 0, 0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: CreateFromDateTimeMicro(2020, 1, 1, 13, 14, 15, 999999),
			want:   "2020-01-01 13:14:15.999999 +0800 CST",
		},
		{
			name:   "case3",
			carbon: CreateFromDateTimeMicro(2020, 1, 31, 13, 14, 15, 999999),
			want:   "2020-01-31 13:14:15.999999 +0800 CST",
		},
		{
			name:   "case4",
			carbon: CreateFromDateTimeMicro(2020, 2, 1, 13, 14, 15, 999999),
			want:   "2020-02-01 13:14:15.999999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "CreateFromDateTimeMicro()")
		})
	}
}

func TestCarbon_CreateFromDateTimeNano(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("xxx").CreateFromDateTimeNano(0, 0, 0, 0, 0, 0, 0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: CreateFromDateTimeNano(2020, 1, 1, 13, 14, 15, 999999999),
			want:   "2020-01-01 13:14:15.999999999 +0800 CST",
		},
		{
			name:   "case3",
			carbon: CreateFromDateTimeNano(2020, 1, 31, 13, 14, 15, 999999999),
			want:   "2020-01-31 13:14:15.999999999 +0800 CST",
		},
		{
			name:   "case4",
			carbon: CreateFromDateTimeNano(2020, 2, 1, 13, 14, 15, 999999999),
			want:   "2020-02-01 13:14:15.999999999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "CreateFromDateTimeNano()")
		})
	}
}

func TestCarbon_CreateFromDate(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("xxx").CreateFromDate(0, 0, 0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: CreateFromDate(2020, 1, 1),
			want:   "2020-01-01 00:00:00 +0800 CST",
		},
		{
			name:   "case3",
			carbon: CreateFromDate(2020, 1, 31),
			want:   "2020-01-31 00:00:00 +0800 CST",
		},
		{
			name:   "case4",
			carbon: CreateFromDate(2020, 2, 1),
			want:   "2020-02-01 00:00:00 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "CreateFromDate()")
		})
	}
}

func TestCarbon_CreateFromDateMilli(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("xxx").CreateFromDateMilli(0, 0, 0, 0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: CreateFromDateMilli(2020, 1, 1, 999),
			want:   "2020-01-01 00:00:00.999 +0800 CST",
		},
		{
			name:   "case3",
			carbon: CreateFromDateMilli(2020, 1, 31, 999),
			want:   "2020-01-31 00:00:00.999 +0800 CST",
		},
		{
			name:   "case4",
			carbon: CreateFromDateMilli(2020, 2, 1, 999),
			want:   "2020-02-01 00:00:00.999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "CreateFromDateMilli()")
		})
	}
}

func TestCarbon_CreateFromDateMicro(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("xxx").CreateFromDateMicro(0, 0, 0, 0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: CreateFromDateMicro(2020, 1, 1, 999999),
			want:   "2020-01-01 00:00:00.999999 +0800 CST",
		},
		{
			name:   "case3",
			carbon: CreateFromDateMicro(2020, 1, 31, 999999),
			want:   "2020-01-31 00:00:00.999999 +0800 CST",
		},
		{
			name:   "case4",
			carbon: CreateFromDateMicro(2020, 2, 1, 999999),
			want:   "2020-02-01 00:00:00.999999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "CreateFromDateMicro()")
		})
	}
}

func TestCarbon_CreateFromDateNano(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("xxx").CreateFromDateNano(0, 0, 0, 0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: CreateFromDateNano(2020, 1, 1, 999999999),
			want:   "2020-01-01 00:00:00.999999999 +0800 CST",
		},
		{
			name:   "case3",
			carbon: CreateFromDateNano(2020, 1, 31, 999999999),
			want:   "2020-01-31 00:00:00.999999999 +0800 CST",
		},
		{
			name:   "case4",
			carbon: CreateFromDateNano(2020, 2, 1, 999999999),
			want:   "2020-02-01 00:00:00.999999999 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "CreateFromDateNano()")
		})
	}
}

func TestCarbon_CreateFromTime(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   Carbon
	}{
		{
			name:   "case1",
			carbon: CreateFromTime(0, 0, 0),
			want:   Now().SetTime(0, 0, 0),
		},
		{
			name:   "case2",
			carbon: CreateFromTime(0, 0, 15),
			want:   Now().SetTime(0, 0, 15),
		},
		{
			name:   "case3",
			carbon: CreateFromTime(0, 14, 15),
			want:   Now().SetTime(0, 14, 15),
		},
		{
			name:   "case4",
			carbon: CreateFromTime(13, 14, 15),
			want:   Now().SetTime(13, 14, 15),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want.ToTimeString(), tt.carbon.ToTimeString(), "CreateFromTime()")
		})
	}
}

func TestCarbon_CreateFromTimeMilli(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   Carbon
	}{
		{
			name:   "case1",
			carbon: CreateFromTimeMilli(0, 0, 0, 999),
			want:   Now().SetTimeMilli(0, 0, 0, 999),
		},
		{
			name:   "case2",
			carbon: CreateFromTimeMilli(0, 0, 15, 999),
			want:   Now().SetTimeMilli(0, 0, 15, 999),
		},
		{
			name:   "case3",
			carbon: CreateFromTimeMilli(0, 14, 15, 999),
			want:   Now().SetTimeMilli(0, 14, 15, 999),
		},
		{
			name:   "case4",
			carbon: CreateFromTimeMilli(13, 14, 15, 999),
			want:   Now().SetTimeMilli(13, 14, 15, 999),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want.ToTimeMilliString(), tt.carbon.ToTimeMilliString(), "CreateFromTimeMilli()")
		})
	}
}

func TestCarbon_CreateFromTimeMicro(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   Carbon
	}{
		{
			name:   "case1",
			carbon: CreateFromTimeMicro(0, 0, 0, 999999),
			want:   Now().SetTimeMicro(0, 0, 0, 999999),
		},
		{
			name:   "case2",
			carbon: CreateFromTimeMicro(0, 0, 15, 999999),
			want:   Now().SetTimeMicro(0, 0, 15, 999999),
		},
		{
			name:   "case3",
			carbon: CreateFromTimeMicro(0, 14, 15, 999999),
			want:   Now().SetTimeMicro(0, 14, 15, 999999),
		},
		{
			name:   "case4",
			carbon: CreateFromTimeMicro(13, 14, 15, 999999),
			want:   Now().SetTimeMicro(13, 14, 15, 999999),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want.ToTimeMilliString(), tt.carbon.ToTimeMilliString(), "CreateFromTimeMicro()")
		})
	}
}

func TestCarbon_CreateFromTimeNano(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   Carbon
	}{
		{
			name:   "case1",
			carbon: CreateFromTimeNano(0, 0, 0, 999999999),
			want:   Now().SetTimeNano(0, 0, 0, 999999999),
		},
		{
			name:   "case2",
			carbon: CreateFromTimeNano(0, 0, 15, 999999999),
			want:   Now().SetTimeNano(0, 0, 15, 999999999),
		},
		{
			name:   "case3",
			carbon: CreateFromTimeNano(0, 14, 15, 999999999),
			want:   Now().SetTimeNano(0, 14, 15, 999999999),
		},
		{
			name:   "case4",
			carbon: CreateFromTimeNano(13, 14, 15, 999999999),
			want:   Now().SetTimeNano(13, 14, 15, 999999999),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want.ToTimeNanoString(), tt.carbon.ToTimeNanoString(), "CreateFromTimeMicro()")
		})
	}
}
