package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Season(t *testing.T) {
	tests := []struct {
		name   string
		carbon *Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("2020-01-05"),
			want:   "Winter",
		},
		{
			name:   "case2",
			carbon: Parse("2020-02-05"),
			want:   "Winter",
		},
		{
			name:   "case3",
			carbon: Parse("2020-03-05"),
			want:   "Spring",
		},
		{
			name:   "case4",
			carbon: Parse("2020-04-05"),
			want:   "Spring",
		},
		{
			name:   "case5",
			carbon: Parse("2020-05-05"),
			want:   "Spring",
		},
		{
			name:   "case6",
			carbon: Parse("2020-06-05"),
			want:   "Summer",
		},
		{
			name:   "case7",
			carbon: Parse("2020-07-05"),
			want:   "Summer",
		},
		{
			name:   "case8",
			carbon: Parse("2020-08-05"),
			want:   "Summer",
		},
		{
			name:   "case9",
			carbon: Parse("2020-09-05"),
			want:   "Autumn",
		},
		{
			name:   "case10",
			carbon: Parse("2020-10-05"),
			want:   "Autumn",
		},
		{
			name:   "case11",
			carbon: Parse("2020-11-05"),
			want:   "Autumn",
		},
		{
			name:   "case12",
			carbon: Parse("2020-12-05"),
			want:   "Winter",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Season(), "Season()")
		})
	}
}

func TestCarbon_StartOfSeason(t *testing.T) {
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
			carbon: Parse("2020-01-15"),
			want:   "2019-12-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-02-15"),
			want:   "2019-12-01",
		},
		{
			name:   "case4",
			carbon: Parse("2020-03-15"),
			want:   "2020-03-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-04-15"),
			want:   "2020-03-01",
		},
		{
			name:   "case6",
			carbon: Parse("2020-05-15"),
			want:   "2020-03-01",
		},
		{
			name:   "case7",
			carbon: Parse("2020-06-15"),
			want:   "2020-06-01",
		},
		{
			name:   "case8",
			carbon: Parse("2020-07-15"),
			want:   "2020-06-01",
		},
		{
			name:   "case9",
			carbon: Parse("2020-08-15"),
			want:   "2020-06-01",
		},
		{
			name:   "case10",
			carbon: Parse("2020-09-15"),
			want:   "2020-09-01",
		},
		{
			name:   "case11",
			carbon: Parse("2020-10-15"),
			want:   "2020-09-01",
		},
		{
			name:   "case12",
			carbon: Parse("2020-11-15"),
			want:   "2020-09-01",
		},
		{
			name:   "case13",
			carbon: Parse("2020-12-15"),
			want:   "2020-12-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.StartOfSeason().ToDateString(), "StartOfSeason()")
		})
	}
}

func TestCarbon_EndOfSeason(t *testing.T) {
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
			carbon: Parse("2020-01-15"),
			want:   "2020-02-29",
		},
		{
			name:   "case3",
			carbon: Parse("2020-02-15"),
			want:   "2020-02-29",
		},
		{
			name:   "case4",
			carbon: Parse("2020-03-15"),
			want:   "2020-05-31",
		},
		{
			name:   "case5",
			carbon: Parse("2020-04-15"),
			want:   "2020-05-31",
		},
		{
			name:   "case6",
			carbon: Parse("2020-05-15"),
			want:   "2020-05-31",
		},
		{
			name:   "case7",
			carbon: Parse("2020-06-15"),
			want:   "2020-08-31",
		},
		{
			name:   "case8",
			carbon: Parse("2020-07-15"),
			want:   "2020-08-31",
		},
		{
			name:   "case9",
			carbon: Parse("2020-08-15"),
			want:   "2020-08-31",
		},
		{
			name:   "case10",
			carbon: Parse("2020-09-15"),
			want:   "2020-11-30",
		},
		{
			name:   "case11",
			carbon: Parse("2020-10-15"),
			want:   "2020-11-30",
		},
		{
			name:   "case12",
			carbon: Parse("2020-11-15"),
			want:   "2020-11-30",
		},
		{
			name:   "case13",
			carbon: Parse("2020-12-15"),
			want:   "2021-02-28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.EndOfSeason().ToDateString(), "EndOfSeason()")
		})
	}
}

func TestCarbon_IsSpring(t *testing.T) {
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
			assert.Equalf(t, tt.want, tt.carbon.IsSpring(), "IsSpring()")
		})
	}
}

func TestCarbon_IsSummer(t *testing.T) {
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
			assert.Equalf(t, tt.want, tt.carbon.IsSummer(), "IsSummer()")
		})
	}
}

func TestCarbon_IsAutumn(t *testing.T) {
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
			assert.Equalf(t, tt.want, tt.carbon.IsAutumn(), "IsAutumn()")
		})
	}
}

func TestCarbon_IsWinter(t *testing.T) {
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
			carbon: Parse("2020-05-01"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsWinter(), "IsWinter()")
		})
	}
}
