package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarbon_Closest(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 Carbon
		carbon2 Carbon
		carbon3 Carbon
		want    string
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse("2023-03-28"),
			carbon3: Parse("2023-04-16"),
			want:    "",
		},
		{
			name:    "case2",
			carbon1: Parse("2023-04-01"),
			carbon2: Parse(""),
			carbon3: Parse("2023-04-16"),
			want:    "",
		},
		{
			name:    "case3",
			carbon1: Parse("2023-04-01"),
			carbon2: Parse("2023-03-28"),
			carbon3: Parse(""),
			want:    "",
		},
		{
			name:    "case4",
			carbon1: Parse("2023-04-01"),
			carbon2: Parse(""),
			carbon3: Parse(""),
			want:    "",
		},
		{
			name:    "case5",
			carbon1: Parse("2023-04-01"),
			carbon2: Parse("2023-03-28"),
			carbon3: Parse("2023-03-28"),
			want:    "2023-03-28",
		},
		{
			name:    "case6",
			carbon1: Parse("2023-04-01"),
			carbon2: Parse("2023-03-28"),
			carbon3: Parse("2023-04-16"),
			want:    "2023-03-28",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.Closest(tt.carbon2, tt.carbon3).ToDateString(), "Closest()")
		})
	}
}

func TestCarbon_Farthest(t *testing.T) {
	tests := []struct {
		name    string
		carbon1 Carbon
		carbon2 Carbon
		carbon3 Carbon
		want    string
	}{
		{
			name:    "case1",
			carbon1: Parse(""),
			carbon2: Parse("2023-03-28"),
			carbon3: Parse("2023-04-16"),
			want:    "",
		},
		{
			name:    "case2",
			carbon1: Parse("2023-04-01"),
			carbon2: Parse(""),
			carbon3: Parse("2023-04-16"),
			want:    "",
		},
		{
			name:    "case3",
			carbon1: Parse("2023-04-01"),
			carbon2: Parse("2023-03-28"),
			carbon3: Parse(""),
			want:    "",
		},
		{
			name:    "case4",
			carbon1: Parse("2023-04-01"),
			carbon2: Parse(""),
			carbon3: Parse(""),
			want:    "",
		},
		{
			name:    "case5",
			carbon1: Parse("2023-04-01"),
			carbon2: Parse("2023-03-28"),
			carbon3: Parse("2023-03-28"),
			want:    "2023-03-28",
		},
		{
			name:    "case6",
			carbon1: Parse("2023-04-01"),
			carbon2: Parse("2023-03-28"),
			carbon3: Parse("2023-03-28"),
			want:    "2023-03-28",
		},
		{
			name:    "case7",
			carbon1: Parse("2023-04-01"),
			carbon2: Parse("2023-04-05"),
			carbon3: Parse("2023-04-02"),
			want:    "2023-04-05",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon1.Farthest(tt.carbon2, tt.carbon3).ToDateString(), "Farthest()")
		})
	}
}

func TestCarbon_Max(t *testing.T) {
	now := Now()
	max := Max(now.SubDay(), now, now.AddDay())
	assert.Equal(t, now.AddDay().Timestamp(), max.Timestamp())
}

func TestCarbon_Min(t *testing.T) {
	now := Now()
	min := Min(now, now.SubDay(), now.AddDay())
	assert.Equal(t, now.SubDay().Timestamp(), min.Timestamp())
}
