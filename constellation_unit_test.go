package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Constellation(t *testing.T) {
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
			carbon: Parse("2020-01-05"),
			want:   "Capricorn",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-22"),
			want:   "Aquarius",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-05"),
			want:   "Aquarius",
		},
		{
			name:   "case5",
			carbon: Parse("2020-03-05"),
			want:   "Pisces",
		},
		{
			name:   "case6",
			carbon: Parse("2020-04-05"),
			want:   "Aries",
		},
		{
			name:   "case7",
			carbon: Parse("2020-05-05"),
			want:   "Taurus",
		},
		{
			name:   "case8",
			carbon: Parse("2020-06-05"),
			want:   "Gemini",
		},
		{
			name:   "case9",
			carbon: Parse("2020-07-05"),
			want:   "Cancer",
		},
		{
			name:   "case10",
			carbon: Parse("2020-08-05"),
			want:   "Leo",
		},
		{
			name:   "case11",
			carbon: Parse("2020-09-05"),
			want:   "Virgo",
		},
		{
			name:   "case12",
			carbon: Parse("2020-10-05"),
			want:   "Libra",
		},
		{
			name:   "case13",
			carbon: Parse("2020-11-05"),
			want:   "Scorpio",
		},
		{
			name:   "case14",
			carbon: Parse("2020-12-05"),
			want:   "Sagittarius",
		},
		{
			name:   "case15",
			carbon: Parse("2020-12-22"),
			want:   "Capricorn",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.Constellation(), "Constellation()")
		})
	}
}

func TestCarbon_IsAries(t *testing.T) {
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
			carbon: Parse("2020-03-21"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2020-04-19"),
			want:   true,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsAries(), "IsAries()")
		})
	}
}

func TestCarbon_IsTaurus(t *testing.T) {
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
			carbon: Parse("2020-04-20"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2020-05-20"),
			want:   true,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsTaurus(), "IsTaurus()")
		})
	}
}

func TestCarbon_IsGemini(t *testing.T) {
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
			carbon: Parse("2020-05-21"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2020-06-21"),
			want:   true,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsGemini(), "IsGemini()")
		})
	}
}

func TestCarbon_IsCancer(t *testing.T) {
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
			carbon: Parse("2020-06-22"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2020-07-22"),
			want:   true,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsCancer(), "IsCancer()")
		})
	}
}

func TestCarbon_IsLeo(t *testing.T) {
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
			carbon: Parse("2020-07-23"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-23"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsLeo(), "IsLeo()")
		})
	}
}

func TestCarbon_IsVirgo(t *testing.T) {
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
			carbon: Parse("2020-08-23"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2020-09-22"),
			want:   true,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsVirgo(), "IsVirgo()")
		})
	}
}

func TestCarbon_IsLibra(t *testing.T) {
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
			carbon: Parse("2020-09-23"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2020-10-23"),
			want:   true,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsLibra(), "IsLibra()")
		})
	}
}

func TestCarbon_IsScorpio(t *testing.T) {
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
			carbon: Parse("2020-10-24"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2020-11-22"),
			want:   true,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsScorpio(), "IsScorpio()")
		})
	}
}

func TestCarbon_IsSagittarius(t *testing.T) {
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
			carbon: Parse("2020-11-23"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2020-12-21"),
			want:   true,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsSagittarius(), "IsSagittarius()")
		})
	}
}

func TestCarbon_IsCapricorn(t *testing.T) {
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
			carbon: Parse("2020-12-22"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-19"),
			want:   true,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsCapricorn(), "IsCapricorn()")
		})
	}
}

func TestCarbon_IsAquarius(t *testing.T) {
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
			carbon: Parse("2020-01-20"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2020-02-18"),
			want:   true,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsAquarius(), "IsAquarius()")
		})
	}
}

func TestCarbon_IsPisces(t *testing.T) {
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
			carbon: Parse("2020-02-19"),
			want:   true,
		},
		{
			name:   "case3",
			carbon: Parse("2020-03-20"),
			want:   true,
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.IsPisces(), "IsPisces()")
		})
	}
}
