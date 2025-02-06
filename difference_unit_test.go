package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_DiffInYears(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    int64
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    0,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    0,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-12-31 13:14:15"),
			carbon2: Parse("2021-01-01 13:14:15"),
			want:    0,
		},
		{
			name:    "case4",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2021-08-28 13:14:59"),
			want:    1,
		},
		{
			name:    "case5",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2018-08-28 13:14:59"),
			want:    -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffInYears(tt.carbon2), "DiffInYears()")
		})
	}
}

func TestCarbon_DiffAbsInYears(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    int64
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    0,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    0,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-12-31 13:14:15"),
			carbon2: Parse("2021-01-01 13:14:15"),
			want:    0,
		},
		{
			name:    "case4",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2021-08-28 13:14:59"),
			want:    1,
		},
		{
			name:    "case5",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2018-08-28 13:14:59"),
			want:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffAbsInYears(tt.carbon2), "DiffAbsInYears()")
		})
	}
}

func TestCarbon_DiffInMonths(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    int64
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    0,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-07-28 13:14:00"),
			want:    0,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-09-06 13:14:59"),
			want:    1,
		},
		{
			name:    "case4",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2018-08-28 13:14:59"),
			want:    -23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffInMonths(tt.carbon2), "DiffInMonths()")
		})
	}
}

func TestCarbon_DiffAbsInMonths(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    int64
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    0,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-07-28 13:14:00"),
			want:    0,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-09-06 13:14:59"),
			want:    1,
		},
		{
			name:    "case4",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2018-08-28 13:14:59"),
			want:    23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffAbsInMonths(tt.carbon2), "DiffAbsInMonths()")
		})
	}
}

func TestCarbon_DiffInWeeks(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    int64
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    0,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    0,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-07-28 13:14:00"),
			want:    -1,
		},
		{
			name:    "case4",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-12 13:14:15"),
			want:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffInWeeks(tt.carbon2), "DiffInWeeks()")
		})
	}
}

func TestCarbon_DiffAbsInWeeks(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    int64
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    0,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    0,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-07-28 13:14:00"),
			want:    1,
		},
		{
			name:    "case4",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-12 13:14:15"),
			want:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffAbsInWeeks(tt.carbon2), "DiffAbsInWeeks()")
		})
	}
}

func TestCarbon_DiffInDays(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    int64
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    0,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    0,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-04 13:14:59"),
			want:    0,
		},
		{
			name:    "case4",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-06 13:14:15"),
			want:    1,
		},
		{
			name:    "case5",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-04 13:00:00"),
			want:    -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffInDays(tt.carbon2), "DiffInDays()")
		})
	}
}

func TestCarbon_DiffAbsInDays(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    int64
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    0,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    0,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-04 13:14:59"),
			want:    0,
		},
		{
			name:    "case4",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-06 13:14:15"),
			want:    1,
		},
		{
			name:    "case5",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-04 13:00:00"),
			want:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffAbsInDays(tt.carbon2), "DiffAbsInDays()")
		})
	}
}

func TestCarbon_DiffInHours(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    int64
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    0,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    0,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 12:14:00"),
			want:    -1,
		},
		{
			name:    "case4",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 14:14:15"),
			want:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffInHours(tt.carbon2), "DiffInHours()")
		})
	}
}

func TestCarbon_DiffAbsInHours(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    int64
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    0,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    0,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 12:14:00"),
			want:    1,
		},
		{
			name:    "case4",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 14:14:15"),
			want:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffAbsInHours(tt.carbon2), "DiffAbsInHours()")
		})
	}
}

func TestCarbon_DiffInMinutes(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    int64
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    0,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:13:00"),
			want:    -1,
		},
		{
			name:    "case4",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:15:15"),
			want:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffInMinutes(tt.carbon2), "DiffInMinutes()")
		})
	}
}

func TestCarbon_DiffAbsInMinutes(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    int64
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    0,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:13:00"),
			want:    1,
		},
		{
			name:    "case4",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:15:15"),
			want:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffAbsInMinutes(tt.carbon2), "DiffAbsInMinutes()")
		})
	}
}

func TestCarbon_DiffInSeconds(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    int64
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    0,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    0,
		},
		{
			name:    "case4",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:14:20"),
			want:    5,
		},
		{
			name:    "case5",
			carbon1: Parse("2020-08-05 13:14:20"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    -5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffInSeconds(tt.carbon2), "DiffInSeconds()")
		})
	}
}

func TestCarbon_DiffAbsInSeconds(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    int64
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    0,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    0,
		},
		{
			name:    "case4",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:14:20"),
			want:    5,
		},
		{
			name:    "case5",
			carbon1: Parse("2020-08-05 13:14:20"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffAbsInSeconds(tt.carbon2), "DiffAbsInSeconds()")
		})
	}
}

func TestCarbon_DiffInString(t *testing.T) {
	now := Now()
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    string
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    "",
		},
		{
			name:    "case2",
			carbon1: Parse("xxx"),
			carbon2: Parse("xxx"),
			want:    "",
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    "just now",
		},
		{
			name:    "case4",
			carbon1: now,
			carbon2: now,
			want:    "just now",
		},
		{
			name:    "case5",
			carbon1: now.AddYearsNoOverflow(1),
			carbon2: now,
			want:    "-1 year",
		},
		{
			name:    "case6",
			carbon1: now.SubYearsNoOverflow(1),
			carbon2: now,
			want:    "1 year",
		},
		{
			name:    "case7",
			carbon1: now.AddMonthsNoOverflow(1),
			carbon2: now,
			want:    "-1 month",
		},
		{
			name:    "case8",
			carbon1: now.SubMonthsNoOverflow(1),
			carbon2: now,
			want:    "1 month",
		},
		{
			name:    "case9",
			carbon1: now.AddDays(1),
			carbon2: now,
			want:    "-1 day",
		},
		{
			name:    "case10",
			carbon1: now.SubDays(1),
			carbon2: now,
			want:    "1 day",
		},
		{
			name:    "case11",
			carbon1: now.AddHours(1),
			carbon2: now,
			want:    "-1 hour",
		},
		{
			name:    "case12",
			carbon1: now.SubHours(1),
			carbon2: now,
			want:    "1 hour",
		},
		{
			name:    "case13",
			carbon1: now.AddHours(1),
			carbon2: now,
			want:    "-1 hour",
		},
		{
			name:    "case14",
			carbon1: now.SubMinutes(1),
			carbon2: now,
			want:    "1 minute",
		},
		{
			name:    "case15",
			carbon1: now.AddMinutes(1),
			carbon2: now,
			want:    "-1 minute",
		},
		{
			name:    "case16",
			carbon1: now.SubMinutes(1),
			carbon2: now,
			want:    "1 minute",
		},

		{
			name:    "case17",
			carbon1: now.AddSeconds(1),
			carbon2: now,
			want:    "-1 second",
		},
		{
			name:    "case18",
			carbon1: now.SubSeconds(1),
			carbon2: now,
			want:    "1 second",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffInString(tt.carbon2), "DiffInString()")
		})
	}
}

func TestCarbon_DiffAbsInString(t *testing.T) {
	now := Now()
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    string
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    "",
		},
		{
			name:    "case2",
			carbon1: Parse("xxx"),
			carbon2: Parse("xxx"),
			want:    "",
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    "just now",
		},
		{
			name:    "case4",
			carbon1: now,
			carbon2: now,
			want:    "just now",
		},
		{
			name:    "case5",
			carbon1: now.AddYearsNoOverflow(1),
			carbon2: now,
			want:    "1 year",
		},
		{
			name:    "case6",
			carbon1: now.SubYearsNoOverflow(1),
			carbon2: now,
			want:    "1 year",
		},
		{
			name:    "case7",
			carbon1: now.AddMonthsNoOverflow(1),
			carbon2: now,
			want:    "1 month",
		},
		{
			name:    "case8",
			carbon1: now.SubMonthsNoOverflow(1),
			carbon2: now,
			want:    "1 month",
		},
		{
			name:    "case9",
			carbon1: now.AddDays(1),
			carbon2: now,
			want:    "1 day",
		},
		{
			name:    "case10",
			carbon1: now.SubDays(1),
			carbon2: now,
			want:    "1 day",
		},
		{
			name:    "case11",
			carbon1: now.AddHours(1),
			carbon2: now,
			want:    "1 hour",
		},
		{
			name:    "case12",
			carbon1: now.SubHours(1),
			carbon2: now,
			want:    "1 hour",
		},
		{
			name:    "case13",
			carbon1: now.AddHours(1),
			carbon2: now,
			want:    "1 hour",
		},
		{
			name:    "case14",
			carbon1: now.SubMinutes(1),
			carbon2: now,
			want:    "1 minute",
		},
		{
			name:    "case15",
			carbon1: now.AddMinutes(1),
			carbon2: now,
			want:    "1 minute",
		},
		{
			name:    "case16",
			carbon1: now.SubMinutes(1),
			carbon2: now,
			want:    "1 minute",
		},

		{
			name:    "case17",
			carbon1: now.AddSeconds(1),
			carbon2: now,
			want:    "1 second",
		},
		{
			name:    "case18",
			carbon1: now.SubSeconds(1),
			carbon2: now,
			want:    "1 second",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffAbsInString(tt.carbon2), "DiffAbsInString()")
		})
	}
}

func TestCarbon_DiffInDuration(t *testing.T) {
	now := Parse("2020-08-05 13:14:15")
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    string
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    "0s",
		},
		{
			name:    "case2",
			carbon1: Parse("xxx"),
			carbon2: Parse("xxx"),
			want:    "0s",
		},
		{
			name:    "case3",
			carbon1: now,
			carbon2: now,
			want:    "0s",
		},
		{
			name:    "case4",
			carbon1: now.AddYearsNoOverflow(1),
			carbon2: now,
			want:    "-8760h0m0s",
		},
		{
			name:    "case5",
			carbon1: now.SubYearsNoOverflow(1),
			carbon2: now,
			want:    "8784h0m0s",
		},
		{
			name:    "case6",
			carbon1: now.AddMonthsNoOverflow(1),
			carbon2: now,
			want:    "-744h0m0s",
		},
		{
			name:    "case7",
			carbon1: now.SubMonthsNoOverflow(1),
			carbon2: now,
			want:    "744h0m0s",
		},

		{
			name:    "case8",
			carbon1: now.AddDays(1),
			carbon2: now,
			want:    "-24h0m0s",
		},
		{
			name:    "case9",
			carbon1: now.SubDays(1),
			carbon2: now,
			want:    "24h0m0s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffInDuration(tt.carbon2).String(), "DiffInDuration()")
		})
	}
}

func TestCarbon_DiffAbsInDuration(t *testing.T) {
	now := Parse("2020-08-05 13:14:15")
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    string
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    "0s",
		},
		{
			name:    "case2",
			carbon1: Parse("xxx"),
			carbon2: Parse("xxx"),
			want:    "0s",
		},
		{
			name:    "case3",
			carbon1: now,
			carbon2: now,
			want:    "0s",
		},
		{
			name:    "case4",
			carbon1: now.AddYearsNoOverflow(1),
			carbon2: now,
			want:    "8760h0m0s",
		},
		{
			name:    "case5",
			carbon1: now.SubYearsNoOverflow(1),
			carbon2: now,
			want:    "8784h0m0s",
		},
		{
			name:    "case6",
			carbon1: now.AddMonthsNoOverflow(1),
			carbon2: now,
			want:    "744h0m0s",
		},
		{
			name:    "case7",
			carbon1: now.SubMonthsNoOverflow(1),
			carbon2: now,
			want:    "744h0m0s",
		},
		{
			name:    "case8",
			carbon1: now.AddDays(1),
			carbon2: now,
			want:    "24h0m0s",
		},
		{
			name:    "case9",
			carbon1: now.SubDays(1),
			carbon2: now,
			want:    "24h0m0s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffAbsInDuration(tt.carbon2).String(), "DiffAbsInDuration()")
		})
	}
}

func TestCarbon_DiffForHumans(t *testing.T) {
	now := Now()
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    string
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    "",
		},
		{
			name:    "case2",
			carbon1: Parse("xxx"),
			carbon2: Parse("xxx"),
			want:    "",
		},
		{
			name:    "case3",
			carbon1: now,
			carbon2: now,
			want:    "just now",
		},
		{
			name:    "case4",
			carbon1: now.AddYearsNoOverflow(1),
			carbon2: now,
			want:    "1 year after",
		},
		{
			name:    "case5",
			carbon1: now.SubYearsNoOverflow(1),
			carbon2: now,
			want:    "1 year before",
		},

		{
			name:    "case6",
			carbon1: now.AddMonthsNoOverflow(1),
			carbon2: now,
			want:    "1 month after",
		},
		{
			name:    "case7",
			carbon1: now.SubMonthsNoOverflow(1),
			carbon2: now,
			want:    "1 month before",
		},

		{
			name:    "case8",
			carbon1: now.AddDays(1),
			carbon2: now,
			want:    "1 day after",
		},
		{
			name:    "case9",
			carbon1: now.SubDays(1),
			carbon2: now,
			want:    "1 day before",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.DiffForHumans(tt.carbon2), "DiffForHumans()")
		})
	}

}

// https://github.com/dromara/carbon/issues/255
func TestCarbon_Issue255(t *testing.T) {
	tests := []struct {
		name  string
		start *Carbon
		end   *Carbon
		want  int64
	}{
		{
			name:  "case1",
			start: Parse("2024-10-11"),
			end:   Parse("2024-11-10"),
			want:  0,
		},
		{
			name:  "case2",
			start: Parse("2024-11-10"),
			end:   Parse("2024-10-11"),
			want:  0,
		},

		{
			name:  "case3",
			start: Parse("2024-10-11"),
			end:   Parse("2024-11-11"),
			want:  1,
		},
		{
			name:  "case4",
			start: Parse("2024-11-11"),
			end:   Parse("2024-10-11"),
			want:  -1,
		},
		{
			name:  "case5",
			start: Parse("2024-10-11 23:59:00"),
			end:   Parse("2024-11-11 00:00:00"),
			want:  0,
		},
		{
			name:  "case6",
			start: Parse("2024-11-11 00:00:00"),
			end:   Parse("2024-10-11 23:59:00"),
			want:  0,
		},
		{
			name:  "case7",
			start: Parse("2024-10-11 00:00:00"),
			end:   Parse("2024-11-11 23:59:59"),
			want:  1,
		},
		{
			name:  "case8",
			start: Parse("2024-11-11 23:59:59"),
			end:   Parse("2024-10-11 00:00:00"),
			want:  -1,
		},
		{
			name:  "case9",
			start: Parse("2020-08-05 13:14:15"),
			end:   Parse("2018-08-28 13:14:59"),
			want:  -23,
		},
		{
			name:  "case10",
			start: Parse("2018-08-28 13:14:59"),
			end:   Parse("2020-08-05 13:14:15"),
			want:  23,
		},
		{
			name:  "case11",
			start: Parse("1024-12-25 13:14:20"),
			end:   Parse("2024-12-25 13:14:20"),
			want:  12000,
		},
		{
			name:  "case12",
			start: Parse("1024-12-25 13:14:20"),
			end:   Parse("2024-12-25 13:14:19"),
			want:  11999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.start.DiffInMonths(tt.end), "DiffInMonths()")
		})
	}
}
