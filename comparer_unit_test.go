package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_IsDST(t *testing.T) {
	tzWithDST, tzWithoutDST := "Australia/Sydney", "Australia/Brisbane"

	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: NewCarbon(),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("0000-01-01 00:00:00 +0000 UTC", tzWithDST),
			want:   false,
		},
		{
			name:   "case4",
			carbon: Parse("0001-01-01 00:00:00 +0000 UTC", tzWithDST),
			want:   false,
		},
		{
			name:   "case5",
			carbon: Parse("2009-01-01", tzWithDST),
			want:   true,
		},
		{
			name:   "case6",
			carbon: Parse("2009-01-01", tzWithoutDST),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsDST(), "IsDST()")
		})
	}
}

func TestCarbon_IsZero(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   true,
		},
		{
			name:   "case2",
			carbon: NewCarbon(),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("0000-01-01 00:00:00 +0000 UTC"),
			want:   false,
		},
		{
			name:   "case4",
			carbon: Parse("0001-01-01 00:00:00 +0000 UTC"),
			want:   true,
		},
		{
			name:   "case5",
			carbon: Parse("2020-08-00 00:00:00"),
			want:   true,
		},
		{
			name:   "case6",
			carbon: Parse("2020-00-05 00:00:00"),
			want:   true,
		},
		{
			name:   "case7",
			carbon: Parse("2020-00-00 00:00:00"),
			want:   true,
		},
		{
			name:   "case8",
			carbon: Parse("2020-08-05 00:00:00"),
			want:   false,
		},
		{
			name:   "case9",
			carbon: Parse("0000-01-01 13:14:15"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsZero(), "IsZero()")
		})
	}
}

func TestCarbon_IsValid(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: NewCarbon(),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2024-00-00 00:00:00 +0000 UTC"),
			want:   false,
		},
		{
			name:   "case4",
			carbon: Parse("0000-01-01 00:00:00 +0000 UTC"),
			want:   true,
		},
		{
			name:   "case5",
			carbon: Parse("0001-01-01 13:14:15 +0000 UTC"),
			want:   true,
		},
		{
			name:   "case6",
			carbon: Parse("2020-08-05"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsValid(), "IsValid()")
		})
	}
}

func TestCarbon_IsInvalid(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   true,
		},
		{
			name:   "case2",
			carbon: NewCarbon(),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("0000-01-01 00:00:00 +0000 UTC"),
			want:   false,
		},
		{
			name:   "case4",
			carbon: Parse("0001-01-01 13:14:15 +0000 UTC"),
			want:   false,
		},
		{
			name:   "case5",
			carbon: Parse("2020-08-05"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsInvalid(), "IsInvalid()")
		})
	}
}

func TestCarbon_IsAM(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 00:00:00"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 08:00:00"),
			want:   true,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05 12:00:00"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsAM(), "IsAM()")
		})
	}
}

func TestCarbon_IsPM(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 00:00:00"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 08:00:00"),
			want:   false,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05 12:00:00"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsPM(), "IsPM()")
		})
	}
}

func TestCarbon_IsNow(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Tomorrow(),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Yesterday(),
			want:   false,
		},
		{
			name:   "case4",
			carbon: Now(),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsNow(), "IsNow()")
		})
	}
}

func TestCarbon_IsFuture(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Tomorrow(),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Yesterday(),
			want:   false,
		},
		{
			name:   "case4",
			carbon: Now(),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsFuture(), "IsFuture()")
		})
	}
}

func TestCarbon_IsPast(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Tomorrow(),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Yesterday(),
			want:   true,
		},
		{
			name:   "case4",
			carbon: Now(),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsPast(), "IsPast()")
		})
	}
}

func TestCarbon_IsLeapYear(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2015-01-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2016-01-01"),
			want:   true,
		},
		{
			name:   "case4",
			carbon: Parse("2017-01-01"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsLeapYear(), "IsLeapYear()")
		})
	}
}

func TestCarbon_IsLongYear(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2015-01-01"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2016-01-01"),
			want:   false,
		},
		{
			name:   "case4",
			carbon: Parse("2017-01-01"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsLongYear(), "IsLongYear()")
		})
	}
}

func TestCarbon_IsJanuary(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2020-02-01"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsJanuary(), "IsJanuary()")
		})
	}
}

func TestCarbon_IsFebruary(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-02-01"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsFebruary(), "IsFebruary()")
		})
	}
}

func TestCarbon_IsMarch(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-03-01"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsMarch(), "IsMarch()")
		})
	}
}

func TestCarbon_IsApril(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-04-01"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsApril(), "IsApril()")
		})
	}
}

func TestCarbon_IsMay(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-05-01"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsMay(), "IsMay()")
		})
	}
}

func TestCarbon_IsJune(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-06-01"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsJune(), "IsJune()")
		})
	}
}

func TestCarbon_IsJuly(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-07-01"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsJuly(), "IsJuly()")
		})
	}
}

func TestCarbon_IsAugust(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-01"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsAugust(), "IsAugust()")
		})
	}
}

func TestCarbon_IsSeptember(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-09-01"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsSeptember(), "IsSeptember()")
		})
	}
}

func TestCarbon_IsOctober(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-10-01"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsOctober(), "IsOctober()")
		})
	}
}

func TestCarbon_IsNovember(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-11-01"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsNovember(), "IsNovember()")
		})
	}
}

func TestCarbon_IsDecember(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-12-01"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsDecember(), "IsDecember()")
		})
	}
}

func TestCarbon_IsMonday(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-10-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-10-05"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsMonday(), "IsMonday()")
		})
	}
}

func TestCarbon_IsTuesday(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-10-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-10-06"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsTuesday(), "IsTuesday()")
		})
	}
}

func TestCarbon_IsWednesday(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-10-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-10-07"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsWednesday(), "IsWednesday()")
		})
	}
}

func TestCarbon_IsThursday(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-10-05"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-10-08"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsThursday(), "IsThursday()")
		})
	}
}

func TestCarbon_IsFriday(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-10-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-10-09"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsFriday(), "IsFriday()")
		})
	}
}

func TestCarbon_IsSaturday(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-10-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-10-10"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsSaturday(), "IsSaturday()")
		})
	}
}

func TestCarbon_IsSunday(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-10-01"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-10-11"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsSunday(), "IsSunday()")
		})
	}
}

func TestCarbon_IsWeekday(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-10-05"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2020-10-10"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsWeekday(), "IsWeekday()")
		})
	}
}

func TestCarbon_IsWeekend(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: Parse("2020-10-05"),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Parse("2020-10-10"),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsWeekend(), "IsWeekend()")
		})
	}
}

func TestCarbon_IsYesterday(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: NewCarbon(),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Yesterday(),
			want:   true,
		},
		{
			name:   "case4",
			carbon: Tomorrow(),
			want:   false,
		},
		{
			name:   "case5",
			carbon: Now(),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsYesterday(), "IsYesterday()")
		})
	}
}

func TestCarbon_IsToday(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: NewCarbon(),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Yesterday(),
			want:   false,
		},
		{
			name:   "case4",
			carbon: Tomorrow(),
			want:   false,
		},
		{
			name:   "case5",
			carbon: Now(),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsToday(), "IsToday()")
		})
	}
}

func TestCarbon_IsTomorrow(t *testing.T) {

	tests := []struct {
		name   string
		carbon *Carbon
		want   bool
	}{
		{
			name:   "case1",
			carbon: Parse(""),
			want:   false,
		},
		{
			name:   "case2",
			carbon: NewCarbon(),
			want:   false,
		},
		{
			name:   "case3",
			carbon: Yesterday(),
			want:   false,
		},
		{
			name:   "case4",
			carbon: Tomorrow(),
			want:   true,
		},
		{
			name:   "case5",
			carbon: Now(),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsTomorrow(), "IsTomorrow()")
		})
	}
}

func TestCarbon_IsSameCentury(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    bool
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    false,
		},
		{
			name:    "case2",
			carbon1: Parse("xxx"),
			carbon2: Parse("xxx"),
			want:    false,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05"),
			carbon2: Parse("3020-08-05"),
			want:    false,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05"),
			carbon2: Parse("2099-08-05"),
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.IsSameCentury(tt.carbon2), "IsSameCentury()")
		})
	}
}

func TestCarbon_IsSameDecade(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    bool
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    false,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05"),
			carbon2: Parse("2030-08-05"),
			want:    false,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05"),
			carbon2: Parse("2021-08-05"),
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.IsSameDecade(tt.carbon2), "IsSameDecade()")
		})
	}
}

func TestCarbon_IsSameYear(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    bool
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    false,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05"),
			carbon2: Parse("2021-08-05"),
			want:    false,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-01-01"),
			carbon2: Parse("2020-12-31"),
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.IsSameYear(tt.carbon2), "IsSameYear()")
		})
	}
}

func TestCarbon_IsSameQuarter(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    bool
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    false,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05"),
			carbon2: Parse("2020-01-05"),
			want:    false,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-01-01"),
			carbon2: Parse("2020-01-31"),
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.IsSameQuarter(tt.carbon2), "IsSameQuarter()")
		})
	}
}

func TestCarbon_IsSameMonth(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    bool
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    false,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05"),
			carbon2: Parse("2021-08-05"),
			want:    false,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-01-01"),
			carbon2: Parse("2020-01-31"),
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.IsSameMonth(tt.carbon2), "IsSameMonth()")
		})
	}
}

func TestCarbon_IsSameDay(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    bool
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    false,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2021-08-05 13:14:15"),
			want:    false,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 00:00:00"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.IsSameDay(tt.carbon2), "IsSameDay()")
		})
	}
}

func TestCarbon_IsSameHour(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    bool
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    false,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2021-08-05 13:14:15"),
			want:    false,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:00:00"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.IsSameHour(tt.carbon2), "IsSameDay()")
		})
	}
}

func TestCarbon_IsSameMinute(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    bool
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    false,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2021-08-05 13:14:15"),
			want:    false,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:14:00"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.IsSameMinute(tt.carbon2), "IsSameMinute()")
		})
	}
}

func TestCarbon_IsSameSecond(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 *Carbon
		carbon2 *Carbon
		want    bool
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse(""),
			want:    false,
		},
		{
			name:    "case2",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2021-08-05 13:14:15"),
			want:    false,
		},
		{
			name:    "case3",
			carbon1: Parse("2020-08-05 13:14:15"),
			carbon2: Parse("2020-08-05 13:14:15"),
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.IsSameSecond(tt.carbon2), "IsSameSecond()")
		})
	}
}

func TestCarbon_Compare(t *testing.T) {
	tests := []struct {
		name   string
		actual bool
		want   bool
	}{
		{
			name:   "case1",
			actual: Parse("xxx").Compare(">", Parse("2020-08-04")),
			want:   false,
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05").Compare(">", Parse("2020-08-04")),
			want:   true,
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").Compare("<", Parse("2020-08-04")),
			want:   false,
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05").Compare("<", Parse("2020-08-06")),
			want:   true,
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").Compare("=", Parse("2020-08-05")),
			want:   true,
		},
		{
			name:   "case6",
			actual: Parse("2020-08-05").Compare(">=", Parse("2020-08-05")),
			want:   true,
		},
		{
			name:   "case7",
			actual: Parse("2020-08-05").Compare("<=", Parse("2020-08-05")),
			want:   true,
		},
		{
			name:   "case8",
			actual: Parse("2020-08-05").Compare("<>", Parse("2020-08-06")),
			want:   true,
		},
		{
			name:   "case9",
			actual: Parse("2020-08-05").Compare("+", Parse("2020-08-06")),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "Compare()")
		})
	}
}

func TestCarbon_Gt(t *testing.T) {
	tests := []struct {
		name   string
		actual bool
		want   bool
	}{
		{
			name:   "case1",
			actual: Parse("xxx").Gt(Parse("2020-08-04")),
			want:   false,
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05").Gt(Parse("2020-08-04")),
			want:   true,
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").Gt(Parse("2020-08-04")),
			want:   true,
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05").Gt(Parse("2020-08-06")),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "Gt()")
		})
	}
}

func TestCarbon_Lt(t *testing.T) {
	tests := []struct {
		name   string
		actual bool
		want   bool
	}{
		{
			name:   "case1",
			actual: Parse("xxx").Lt(Parse("2020-08-04")),
			want:   false,
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05").Lt(Parse("2020-08-04")),
			want:   false,
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").Lt(Parse("2020-08-04")),
			want:   false,
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05").Lt(Parse("2020-08-06")),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "Lt()")
		})
	}
}

func TestCarbon_Eq(t *testing.T) {
	tests := []struct {
		name   string
		actual bool
		want   bool
	}{
		{
			name:   "case1",
			actual: Parse("xxx").Eq(Parse("2020-08-05")),
			want:   false,
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05").Eq(Parse("2020-08-05")),
			want:   true,
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").Eq(Parse("2020-08-04")),
			want:   false,
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05").Eq(Parse("2020-08-06")),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "Eq()")
		})
	}
}

func TestCarbon_Ne(t *testing.T) {
	tests := []struct {
		name   string
		actual bool
		want   bool
	}{
		{
			name:   "case1",
			actual: Parse("xxx").Ne(Parse("2020-08-05")),
			want:   true,
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05").Ne(Parse("2020-08-05")),
			want:   false,
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").Ne(Parse("2020-08-04")),
			want:   true,
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05").Ne(Parse("2020-08-06")),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "Ne()")
		})
	}
}

func TestCarbon_Gte(t *testing.T) {
	tests := []struct {
		name   string
		actual bool
		want   bool
	}{
		{
			name:   "case1",
			actual: Parse("xxx").Gte(Parse("2020-08-05")),
			want:   false,
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05").Gte(Parse("2020-08-05")),
			want:   true,
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").Gte(Parse("2020-08-04")),
			want:   true,
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05").Gte(Parse("2020-08-06")),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "Gte()")
		})
	}
}

func TestCarbon_Lte(t *testing.T) {
	tests := []struct {
		name   string
		actual bool
		want   bool
	}{
		{
			name:   "case1",
			actual: Parse("xxx").Lte(Parse("2020-08-05")),
			want:   false,
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05").Lte(Parse("2020-08-05")),
			want:   true,
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").Lte(Parse("2020-08-04")),
			want:   false,
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05").Lte(Parse("2020-08-06")),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "Lte()")
		})
	}
}

func TestCarbon_Between(t *testing.T) {
	tests := []struct {
		name   string
		actual bool
		want   bool
	}{
		{
			name:   "case1",
			actual: Parse("xxx").Between(Parse("2020-08-05"), Parse("2020-08-05")),
			want:   false,
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05").Between(Parse("2020-08-05"), Parse("2020-08-05")),
			want:   false,
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").Between(Parse("2020-08-05"), Parse("2020-08-06")),
			want:   false,
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05").Between(Parse("2020-08-05"), Parse("2020-08-05")),
			want:   false,
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").Between(Parse("2020-08-04"), Parse("2020-08-06")),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "Between()")
		})
	}
}

func TestCarbon_BetweenIncludedStart(t *testing.T) {
	tests := []struct {
		name   string
		actual bool
		want   bool
	}{
		{
			name:   "case1",
			actual: Parse("xxx").BetweenIncludedStart(Parse("2020-08-05"), Parse("2020-08-05")),
			want:   false,
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05").BetweenIncludedStart(Parse("2020-08-05"), Parse("2020-08-05")),
			want:   false,
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").BetweenIncludedStart(Parse("2020-08-05"), Parse("2020-08-06")),
			want:   true,
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05").BetweenIncludedStart(Parse("2020-08-04"), Parse("2020-08-05")),
			want:   false,
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").BetweenIncludedStart(Parse("2020-08-04"), Parse("2020-08-06")),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "BetweenIncludedStart()")
		})
	}
}

func TestCarbon_BetweenIncludedEnd(t *testing.T) {
	tests := []struct {
		name   string
		actual bool
		want   bool
	}{
		{
			name:   "case1",
			actual: Parse("xxx").BetweenIncludedEnd(Parse("2020-08-05"), Parse("2020-08-05")),
			want:   false,
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05").BetweenIncludedEnd(Parse("2020-08-05"), Parse("2020-08-05")),
			want:   false,
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").BetweenIncludedEnd(Parse("2020-08-05"), Parse("2020-08-06")),
			want:   false,
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05").BetweenIncludedEnd(Parse("2020-08-04"), Parse("2020-08-05")),
			want:   true,
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").BetweenIncludedEnd(Parse("2020-08-04"), Parse("2020-08-06")),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "BetweenIncludedEnd()")
		})
	}
}

func TestCarbon_BetweenIncludedBoth(t *testing.T) {
	tests := []struct {
		name   string
		actual bool
		want   bool
	}{
		{
			name:   "case1",
			actual: Parse("xxx").BetweenIncludedBoth(Parse("2020-08-05"), Parse("2020-08-05")),
			want:   false,
		},
		{
			name:   "case2",
			actual: Parse("2020-08-05").BetweenIncludedBoth(Parse("2020-08-05"), Parse("2020-08-05")),
			want:   true,
		},
		{
			name:   "case3",
			actual: Parse("2020-08-05").BetweenIncludedBoth(Parse("2020-08-05"), Parse("2020-08-06")),
			want:   true,
		},
		{
			name:   "case4",
			actual: Parse("2020-08-05").BetweenIncludedBoth(Parse("2020-08-04"), Parse("2020-08-05")),
			want:   true,
		},
		{
			name:   "case5",
			actual: Parse("2020-08-05").BetweenIncludedBoth(Parse("2020-08-06"), Parse("2020-08-06")),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "BetweenIncludedBoth()")
		})
	}
}
