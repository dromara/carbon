package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Lunar(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("xxx", PRC),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2024-01-18", PRC),
			want:   "2023-12-08 00:00:00",
		},
		{
			name:   "case3",
			carbon: Parse("2024-01-21", PRC),
			want:   "2023-12-11 00:00:00",
		},
		{
			name:   "case4",
			carbon: Parse("2024-01-24", PRC),
			want:   "2023-12-14 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Lunar().String(), "Lunar()")
		})
	}
}

func TestCreateFromLunar(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: CreateFromLunar(2023, 12, 11, 0, 0, 0, false),
			want:   "2024-01-21 00:00:00",
		},
		{
			name:   "case2",
			carbon: CreateFromLunar(2023, 12, 8, 0, 0, 0, false),
			want:   "2024-01-18 00:00:00",
		},
		{
			name:   "case3",
			carbon: CreateFromLunar(2023, 12, 14, 12, 0, 0, false),
			want:   "2024-01-24 12:00:00",
		},
		{
			name:   "case4",
			carbon: CreateFromLunar(2200, 12, 14, 12, 0, 0, false),
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateTimeString(), "CreateFromLunar()")
		})
	}
}

func TestCarbon_Julian(t *testing.T) {
	tests := []struct {
		name    string
		carbon  *Carbon
		wantJD  float64
		wantMJD float64
	}{
		{
			name:    "case1",
			carbon:  Parse("xxx"),
			wantJD:  0,
			wantMJD: 0,
		},
		{
			name:    "case2",
			carbon:  Parse("2024-01-24 12:00:00"),
			wantJD:  2460334,
			wantMJD: 60333.5,
		},
		{
			name:    "case3",
			carbon:  CreateFromJulian(2460334),
			wantJD:  2460334,
			wantMJD: 60333.5,
		},
		{
			name:    "case4",
			carbon:  CreateFromJulian(60333.5),
			wantJD:  2460334,
			wantMJD: 60333.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantJD, tt.carbon.Julian().JD(), "JD()")
			assert.Equalf(t, tt.wantMJD, tt.carbon.Julian().MJD(), "MJD()")
		})
	}
}

func TestCarbon_CreateFromJulian(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: CreateFromJulian(2460334),
			want:   "2024-01-24 12:00:00",
		},
		{
			name:   "case2",
			carbon: CreateFromJulian(60333.5),
			want:   "2024-01-24 12:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateTimeString(), "CreateFromJulian()")
		})
	}
}

func TestCarbon_Persian(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("xxx", PRC),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("1800-01-01 00:00:00", PRC),
			want:   "1178-10-11 00:00:00",
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15", PRC),
			want:   "1399-05-15 13:14:15",
		},
		{
			name:   "case4",
			carbon: Parse("2024-01-01 00:00:00", PRC),
			want:   "1402-10-11 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Persian().String(), "Persian()")
		})
	}
}

func TestCarbon_CreateFromPersian(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: CreateFromPersian(1178, 10, 11, 0, 0, 0),
			want:   "1800-01-01 00:00:00",
		},
		{
			name:   "case2",
			carbon: CreateFromPersian(1402, 10, 11, 0, 0, 0),
			want:   "2024-01-01 00:00:00",
		},
		{
			name:   "case3",
			carbon: CreateFromPersian(1403, 5, 15, 12, 0, 0),
			want:   "2024-08-05 12:00:00",
		},
		{
			name:   "case4",
			carbon: CreateFromPersian(9999, 12, 14, 12, 0, 0),
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateTimeString(), "CreateFromPersian()")
		})
	}
}

// https://github.com/dromara/carbon/issues/246
func TestCarbon_Issue246(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("2024-09-21 00:00:00", PRC),
			want:   "2024-08-19 00:00:00",
		},
		{
			name:   "case2",
			carbon: Parse("2024-09-21 23:50:00", PRC),
			want:   "2024-08-19 23:50:00",
		},
		{
			name:   "case3",
			carbon: Parse("2024-09-21 23:54:00", PRC),
			want:   "2024-08-19 23:54:00",
		},
		{
			name:   "case4",
			carbon: Parse("2024-09-21 23:55:00", PRC),
			want:   "2024-08-19 23:55:00",
		},
		{
			name:   "case5",
			carbon: Parse("2024-09-21 23:59:00", PRC),
			want:   "2024-08-19 23:59:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Lunar().String(), "Lunar()")
		})
	}
}
