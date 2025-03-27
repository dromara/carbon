package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Now(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Now("xxx"),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Now(),
			want:   time.Now().Format(DateLayout),
		},
		{
			name:   "case3",
			carbon: Now(Local),
			want:   time.Now().Format(DateLayout),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "Now()")
		})
	}
}

func TestCarbon_Yesterday(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Yesterday("xxx"),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Yesterday(),
			want:   time.Now().AddDate(0, 0, -1).Format(DateLayout),
		},
		{
			name:   "case3",
			carbon: Yesterday(Local),
			want:   time.Now().AddDate(0, 0, -1).Format(DateLayout),
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05").Yesterday(),
			want:   "2020-08-04",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "Yesterday()")
		})
	}
}

func TestCarbon_Tomorrow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Tomorrow("xxx"),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Tomorrow(),
			want:   time.Now().AddDate(0, 0, 1).Format(DateLayout),
		},
		{
			name:   "case3",
			carbon: Tomorrow(Local),
			want:   time.Now().AddDate(0, 0, 1).Format(DateLayout),
		},
		{
			name:   "case4",
			carbon: Parse("2020-08-05").Tomorrow(),
			want:   "2020-08-06",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "Tomorrow()")
		})
	}
}

func TestCarbon_AddDuration(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddDuration(""),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("now").AddDuration("xxx"),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01 13:14:15").AddDuration("10h"),
			want:   "2020-01-01 23:14:15",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01 13:14:15").AddDuration("10.5h"),
			want:   "2020-01-01 23:44:15",
		},
		{
			name:   "case5",
			carbon: Parse("2020-01-01 13:14:15").AddDuration("10m"),
			want:   "2020-01-01 13:24:15",
		},
		{
			name:   "case6",
			carbon: Parse("2020-01-01 13:14:15").AddDuration("10.5m"),
			want:   "2020-01-01 13:24:45",
		},
		{
			name:   "case7",
			carbon: Parse("2020-01-01 13:14:15").AddDuration(""),
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateTimeString(), "AddDuration()")
		})
	}
}

func TestCarbon_SubDuration(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubDuration(""),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("now").SubDuration("xxx"),
			want:   "",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01 13:14:15").SubDuration("10h"),
			want:   "2020-01-01 03:14:15",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01 13:14:15").SubDuration("10.5h"),
			want:   "2020-01-01 02:44:15",
		},
		{
			name:   "case5",
			carbon: Parse("2020-01-01 13:14:15").SubDuration("10m"),
			want:   "2020-01-01 13:04:15",
		},
		{
			name:   "case6",
			carbon: Parse("2020-01-01 13:14:15").SubDuration("10.5m"),
			want:   "2020-01-01 13:03:45",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateTimeString(), "SubDuration()")
		})
	}
}

func TestCarbon_AddCenturies(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddCenturies(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddCenturies(3),
			want:   "2320-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddCenturies(3),
			want:   "2320-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddCenturies(3),
			want:   "2320-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddCenturies(3),
			want:   "2320-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddCenturies(3),
			want:   "2320-02-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddCenturies()")
		})
	}
}

func TestCarbon_AddCenturiesNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddCenturiesNoOverflow(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddCenturiesNoOverflow(3),
			want:   "2320-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddCenturiesNoOverflow(3),
			want:   "2320-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddCenturiesNoOverflow(3),
			want:   "2320-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddCenturiesNoOverflow(3),
			want:   "2320-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddCenturiesNoOverflow(3),
			want:   "2320-02-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddCenturiesNoOverflow()")
		})
	}
}

func TestCarbon_SubCenturies(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubCenturies(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubCenturies(3),
			want:   "1720-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubCenturies(3),
			want:   "1720-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubCenturies(3),
			want:   "1720-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubCenturies(3),
			want:   "1720-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubCenturies(3),
			want:   "1720-02-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubCenturies()")
		})
	}
}

func TestCarbon_SubCenturiesNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubCenturiesNoOverflow(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubCenturiesNoOverflow(3),
			want:   "1720-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubCenturiesNoOverflow(3),
			want:   "1720-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubCenturiesNoOverflow(3),
			want:   "1720-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubCenturiesNoOverflow(3),
			want:   "1720-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubCenturiesNoOverflow(3),
			want:   "1720-02-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubCenturiesNoOverflow()")
		})
	}
}

func TestCarbon_AddCentury(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddCentury(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddCentury(),
			want:   "2120-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddCentury(),
			want:   "2120-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddCentury(),
			want:   "2120-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddCentury(),
			want:   "2120-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddCentury(),
			want:   "2120-02-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddCentury()")
		})
	}
}

func TestCarbon_AddCenturyNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddCenturyNoOverflow(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddCenturyNoOverflow(),
			want:   "2120-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddCenturyNoOverflow(),
			want:   "2120-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddCenturyNoOverflow(),
			want:   "2120-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddCenturyNoOverflow(),
			want:   "2120-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddCenturyNoOverflow(),
			want:   "2120-02-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddCenturyNoOverflow()")
		})
	}
}

func TestCarbon_SubCentury(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubCentury(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubCentury(),
			want:   "1920-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubCentury(),
			want:   "1920-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubCentury(),
			want:   "1920-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubCentury(),
			want:   "1920-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubCentury(),
			want:   "1920-02-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubCentury()")
		})
	}
}

func TestCarbon_SubCenturyNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubCenturyNoOverflow(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubCenturyNoOverflow(),
			want:   "1920-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubCenturyNoOverflow(),
			want:   "1920-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubCenturyNoOverflow(),
			want:   "1920-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubCenturyNoOverflow(),
			want:   "1920-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubCenturyNoOverflow(),
			want:   "1920-02-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubCenturyNoOverflow()")
		})
	}
}

func TestCarbon_AddDecades(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddDecades(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddDecades(3),
			want:   "2050-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddDecades(3),
			want:   "2050-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddDecades(3),
			want:   "2050-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddDecades(3),
			want:   "2050-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddDecades(3),
			want:   "2050-03-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddDecades()")
		})
	}
}

func TestCarbon_AddDecadesNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddDecadesNoOverflow(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddDecadesNoOverflow(3),
			want:   "2050-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddDecadesNoOverflow(3),
			want:   "2050-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddDecadesNoOverflow(3),
			want:   "2050-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddDecadesNoOverflow(3),
			want:   "2050-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddDecadesNoOverflow(3),
			want:   "2050-02-28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddDecadesNoOverflow()")
		})
	}
}

func TestCarbon_AddDecade(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddDecade(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddDecade(),
			want:   "2030-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddDecade(),
			want:   "2030-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddDecade(),
			want:   "2030-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddDecade(),
			want:   "2030-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddDecade(),
			want:   "2030-03-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddDecade()")
		})
	}
}

func TestCarbon_AddDecadeNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddDecadeNoOverflow(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddDecadeNoOverflow(),
			want:   "2030-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddDecadeNoOverflow(),
			want:   "2030-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddDecadeNoOverflow(),
			want:   "2030-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddDecadeNoOverflow(),
			want:   "2030-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddDecadeNoOverflow(),
			want:   "2030-02-28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddDecadeNoOverflow()")
		})
	}
}

func TestCarbon_SubDecades(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubDecades(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubDecades(3),
			want:   "1990-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubDecades(3),
			want:   "1990-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubDecades(3),
			want:   "1990-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubDecades(3),
			want:   "1990-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubDecades(3),
			want:   "1990-03-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubDecades()")
		})
	}
}

func TestCarbon_SubDecadesNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubDecadesNoOverflow(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubDecadesNoOverflow(3),
			want:   "1990-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubDecadesNoOverflow(3),
			want:   "1990-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubDecadesNoOverflow(3),
			want:   "1990-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubDecadesNoOverflow(3),
			want:   "1990-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubDecadesNoOverflow(3),
			want:   "1990-02-28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubDecadesNoOverflow()")
		})
	}
}

func TestCarbon_SubDecade(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubDecade(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubDecade(),
			want:   "2010-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubDecade(),
			want:   "2010-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubDecade(),
			want:   "2010-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubDecade(),
			want:   "2010-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubDecade(),
			want:   "2010-03-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubDecade()")
		})
	}
}

func TestCarbon_SubDecadeNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubDecadeNoOverflow(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubDecadeNoOverflow(),
			want:   "2010-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubDecadeNoOverflow(),
			want:   "2010-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubDecadeNoOverflow(),
			want:   "2010-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubDecadeNoOverflow(),
			want:   "2010-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubDecadeNoOverflow(),
			want:   "2010-02-28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubDecadeNoOverflow()")
		})
	}
}

func TestCarbon_AddYears(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddYears(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddYears(3),
			want:   "2023-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddYears(3),
			want:   "2023-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddYears(3),
			want:   "2023-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddYears(3),
			want:   "2023-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddYears(3),
			want:   "2023-03-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddYears()")
		})
	}
}

func TestCarbon_AddYearsNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddYearsNoOverflow(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddYearsNoOverflow(3),
			want:   "2023-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddYearsNoOverflow(3),
			want:   "2023-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddYearsNoOverflow(3),
			want:   "2023-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddYearsNoOverflow(3),
			want:   "2023-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddYearsNoOverflow(3),
			want:   "2023-02-28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddYearsNoOverflow()")
		})
	}
}

func TestCarbon_SubYears(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubYears(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubYears(3),
			want:   "2017-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubYears(3),
			want:   "2017-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubYears(3),
			want:   "2017-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubYears(3),
			want:   "2017-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubYears(3),
			want:   "2017-03-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubYears()")
		})
	}
}

func TestCarbon_SubYearsNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubYearsNoOverflow(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubYearsNoOverflow(3),
			want:   "2017-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubYearsNoOverflow(3),
			want:   "2017-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubYearsNoOverflow(3),
			want:   "2017-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubYearsNoOverflow(3),
			want:   "2017-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubYearsNoOverflow(3),
			want:   "2017-02-28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubYearsNoOverflow()")
		})
	}
}

func TestCarbon_AddYear(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddYear(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddYear(),
			want:   "2021-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddYear(),
			want:   "2021-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddYear(),
			want:   "2021-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddYear(),
			want:   "2021-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddYear(),
			want:   "2021-03-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddYear()")
		})
	}
}

func TestCarbon_AddYearNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddYearNoOverflow(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddYearNoOverflow(),
			want:   "2021-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddYearNoOverflow(),
			want:   "2021-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddYearNoOverflow(),
			want:   "2021-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddYearNoOverflow(),
			want:   "2021-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddYearNoOverflow(),
			want:   "2021-02-28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddYearNoOverflow()")
		})
	}
}

func TestCarbon_SubYear(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubYear(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubYear(),
			want:   "2019-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubYear(),
			want:   "2019-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubYear(),
			want:   "2019-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubYear(),
			want:   "2019-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubYear(),
			want:   "2019-03-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubYear()")
		})
	}
}

func TestCarbon_SubYearNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubYearNoOverflow(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubYearNoOverflow(),
			want:   "2019-01-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubYearNoOverflow(),
			want:   "2019-01-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubYearNoOverflow(),
			want:   "2019-02-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubYearNoOverflow(),
			want:   "2019-02-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubYearNoOverflow(),
			want:   "2019-02-28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubYearNoOverflow()")
		})
	}
}

func TestCarbon_AddQuarters(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddQuarters(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2019-08-01").AddQuarters(2),
			want:   "2020-02-01",
		},
		{
			name:   "case3",
			carbon: Parse("2019-08-31").AddQuarters(2),
			want:   "2020-03-02",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01").AddQuarters(2),
			want:   "2020-07-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddQuarters(2),
			want:   "2020-08-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddQuarters(2),
			want:   "2020-08-29",
		},
		{
			name:   "case7",
			carbon: Parse("2020-08-05").AddQuarters(2),
			want:   "2021-02-05",
		},
		{
			name:   "case8",
			carbon: Parse("2020-08-31").AddQuarters(2),
			want:   "2021-03-03",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddQuarters()")
		})
	}
}

func TestCarbon_AddQuartersNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddQuartersNoOverflow(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2019-08-01").AddQuartersNoOverflow(2),
			want:   "2020-02-01",
		},
		{
			name:   "case3",
			carbon: Parse("2019-08-31").AddQuartersNoOverflow(2),
			want:   "2020-02-29",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01").AddQuartersNoOverflow(2),
			want:   "2020-07-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddQuartersNoOverflow(2),
			want:   "2020-08-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddQuartersNoOverflow(2),
			want:   "2020-08-29",
		},
		{
			name:   "case7",
			carbon: Parse("2020-08-05").AddQuartersNoOverflow(2),
			want:   "2021-02-05",
		},
		{
			name:   "case8",
			carbon: Parse("2020-08-31").AddQuartersNoOverflow(2),
			want:   "2021-02-28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddQuartersNoOverflow()")
		})
	}
}

func TestCarbon_SubQuarters(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubQuarters(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2019-08-01").SubQuarters(2),
			want:   "2019-02-01",
		},
		{
			name:   "case3",
			carbon: Parse("2019-08-31").SubQuarters(2),
			want:   "2019-03-03",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01").SubQuarters(2),
			want:   "2019-07-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubQuarters(2),
			want:   "2019-08-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubQuarters(2),
			want:   "2019-08-29",
		},
		{
			name:   "case7",
			carbon: Parse("2020-08-05").SubQuarters(2),
			want:   "2020-02-05",
		},
		{
			name:   "case8",
			carbon: Parse("2020-08-31").SubQuarters(2),
			want:   "2020-03-02",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubQuarters()")
		})
	}
}

func TestCarbon_SubQuartersNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubQuartersNoOverflow(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2019-08-01").SubQuartersNoOverflow(2),
			want:   "2019-02-01",
		},
		{
			name:   "case3",
			carbon: Parse("2019-08-31").SubQuartersNoOverflow(2),
			want:   "2019-02-28",
		},
		{
			name:   "case4",
			carbon: Parse("2020-01-01").SubQuartersNoOverflow(2),
			want:   "2019-07-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubQuartersNoOverflow(2),
			want:   "2019-08-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubQuartersNoOverflow(2),
			want:   "2019-08-29",
		},
		{
			name:   "case7",
			carbon: Parse("2020-08-05").SubQuartersNoOverflow(2),
			want:   "2020-02-05",
		},
		{
			name:   "case8",
			carbon: Parse("2020-08-31").SubQuartersNoOverflow(2),
			want:   "2020-02-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubQuartersNoOverflow()")
		})
	}
}

func TestCarbon_AddQuarter(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddQuarter(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2019-11-01").AddQuarter(),
			want:   "2020-02-01",
		},
		{
			name:   "case3",
			carbon: Parse("2019-11-30").AddQuarter(),
			want:   "2020-03-01",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-28").AddQuarter(),
			want:   "2020-05-28",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-29").AddQuarter(),
			want:   "2020-05-29",
		},
		{
			name:   "case6",
			carbon: Parse("2020-11-01").AddQuarter(),
			want:   "2021-02-01",
		},
		{
			name:   "case7",
			carbon: Parse("2020-11-30").AddQuarter(),
			want:   "2021-03-02",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddQuarter()")
		})
	}
}

func TestCarbon_AddQuarterNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddQuarterNoOverflow(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2019-11-01").AddQuarterNoOverflow(),
			want:   "2020-02-01",
		},
		{
			name:   "case3",
			carbon: Parse("2019-11-30").AddQuarterNoOverflow(),
			want:   "2020-02-29",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-28").AddQuarterNoOverflow(),
			want:   "2020-05-28",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-29").AddQuarterNoOverflow(),
			want:   "2020-05-29",
		},
		{
			name:   "case6",
			carbon: Parse("2020-11-01").AddQuarterNoOverflow(),
			want:   "2021-02-01",
		},
		{
			name:   "case7",
			carbon: Parse("2020-11-30").AddQuarterNoOverflow(),
			want:   "2021-02-28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddQuarterNoOverflow()")
		})
	}
}

func TestCarbon_SubQuarter(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubQuarter(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubQuarter(),
			want:   "2019-10-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubQuarter(),
			want:   "2019-10-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubQuarter(),
			want:   "2019-11-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubQuarter(),
			want:   "2019-11-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubQuarter(),
			want:   "2019-11-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubQuarter()")
		})
	}
}

func TestCarbon_SubQuarterNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubQuarterNoOverflow(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubQuarterNoOverflow(),
			want:   "2019-10-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubQuarterNoOverflow(),
			want:   "2019-10-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubQuarterNoOverflow(),
			want:   "2019-11-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubQuarterNoOverflow(),
			want:   "2019-11-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubQuarterNoOverflow(),
			want:   "2019-11-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubQuarterNoOverflow()")
		})
	}
}

func TestCarbon_AddMonths(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddMonthsNoOverflow(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddMonths(3),
			want:   "2020-04-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddMonths(3),
			want:   "2020-05-01",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddMonths(3),
			want:   "2020-05-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddMonths(3),
			want:   "2020-05-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddMonths(3),
			want:   "2020-05-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddMonths()")
		})
	}
}

func TestCarbon_AddMonthsNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddMonthsNoOverflow(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddMonthsNoOverflow(3),
			want:   "2020-04-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddMonthsNoOverflow(3),
			want:   "2020-04-30",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddMonthsNoOverflow(3),
			want:   "2020-05-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddMonthsNoOverflow(3),
			want:   "2020-05-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddMonthsNoOverflow(3),
			want:   "2020-05-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddMonthsNoOverflow()")
		})
	}
}

func TestCarbon_SubMonths(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubMonths(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubMonths(1),
			want:   "2019-12-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SubMonths(2),
			want:   "2019-11-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubMonths()")
		})
	}
}

func TestCarbon_SubMonthsNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubMonthsNoOverflow(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubMonthsNoOverflow(1),
			want:   "2019-12-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SubMonthsNoOverflow(2),
			want:   "2019-11-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubMonthsNoOverflow()")
		})
	}
}

func TestCarbon_AddMonth(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddMonth(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddMonth(),
			want:   "2020-02-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddMonth(),
			want:   "2020-03-02",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddMonth(),
			want:   "2020-03-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddMonth(),
			want:   "2020-03-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddMonth(),
			want:   "2020-03-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddMonth()")
		})
	}
}

func TestCarbon_AddMonthNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddMonthNoOverflow(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddMonthNoOverflow(),
			want:   "2020-02-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddMonthNoOverflow(),
			want:   "2020-02-29",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddMonthNoOverflow(),
			want:   "2020-03-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddMonthNoOverflow(),
			want:   "2020-03-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddMonthNoOverflow(),
			want:   "2020-03-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddMonthNoOverflow()")
		})
	}
}

func TestCarbon_SubMonth(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubMonth(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubMonth(),
			want:   "2019-12-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubMonth(),
			want:   "2019-12-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubMonth(),
			want:   "2020-01-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubMonth(),
			want:   "2020-01-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubMonth(),
			want:   "2020-01-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubMonth()")
		})
	}
}

func TestCarbon_SubMonthNoOverflow(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubMonthNoOverflow(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubMonthNoOverflow(),
			want:   "2019-12-01",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubMonthNoOverflow(),
			want:   "2019-12-31",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubMonthNoOverflow(),
			want:   "2020-01-01",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubMonthNoOverflow(),
			want:   "2020-01-28",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubMonthNoOverflow(),
			want:   "2020-01-29",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubMonthNoOverflow()")
		})
	}
}

func TestCarbon_AddWeeks(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddWeeks(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddWeeks(1),
			want:   "2020-01-08",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").AddWeeks(2),
			want:   "2020-01-15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddWeeks()")
		})
	}
}

func TestCarbon_SubWeeks(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubWeeks(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubWeeks(1),
			want:   "2019-12-25",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SubWeeks(2),
			want:   "2019-12-18",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubWeeks()")
		})
	}
}

func TestCarbon_AddWeek(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddWeek(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddWeek(),
			want:   "2020-01-08",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddWeek(),
			want:   "2020-02-07",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddWeek(),
			want:   "2020-02-08",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddWeek(),
			want:   "2020-03-06",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddWeek(),
			want:   "2020-03-07",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddWeek()")
		})
	}
}

func TestCarbon_SubWeek(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubWeek(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubWeek(),
			want:   "2019-12-25",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubWeek(),
			want:   "2020-01-24",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubWeek(),
			want:   "2020-01-25",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubWeek(),
			want:   "2020-02-21",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubWeek(),
			want:   "2020-02-22",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubWeek()")
		})
	}
}

func TestCarbon_AddDays(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddDays(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddDays(1),
			want:   "2020-01-02",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").AddDays(2),
			want:   "2020-01-03",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddDays()")
		})
	}
}

func TestCarbon_SubDays(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddHours(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubDays(1),
			want:   "2019-12-31",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-01").SubDays(2),
			want:   "2019-12-30",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubDays()")
		})
	}
}

func TestCarbon_AddDay(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddDay(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").AddDay(),
			want:   "2020-01-02",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").AddDay(),
			want:   "2020-02-01",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").AddDay(),
			want:   "2020-02-02",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").AddDay(),
			want:   "2020-02-29",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").AddDay(),
			want:   "2020-03-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "AddDay()")
		})
	}
}

func TestCarbon_SubDay(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubDay(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-01-01").SubDay(),
			want:   "2019-12-31",
		},
		{
			name:   "case3",
			carbon: Parse("2020-01-31").SubDay(),
			want:   "2020-01-30",
		},
		{
			name:   "case4",
			carbon: Parse("2020-02-01").SubDay(),
			want:   "2020-01-31",
		},
		{
			name:   "case5",
			carbon: Parse("2020-02-28").SubDay(),
			want:   "2020-02-27",
		},
		{
			name:   "case6",
			carbon: Parse("2020-02-29").SubDay(),
			want:   "2020-02-28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToDateString(), "SubDay()")
		})
	}
}

func TestCarbon_AddHours(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddHours(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddHours(1),
			want:   "2020-08-05 14:14:15.222222222 +0800 CST",
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddHours(2),
			want:   "2020-08-05 15:14:15.222222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "AddHours()")
		})
	}
}

func TestCarbon_SubHours(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubHours(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubHours(1),
			want:   "2020-08-05 12:14:15.222222222 +0800 CST",
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubHours(2),
			want:   "2020-08-05 11:14:15.222222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "SubHours()")
		})
	}
}

func TestCarbon_AddHour(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddHour(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddHour(),
			want:   "2020-08-05 14:14:15.222222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "AddHour()")
		})
	}
}

func TestCarbon_SubHour(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubHour(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubHour(),
			want:   "2020-08-05 12:14:15.222222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "SubHour()")
		})
	}
}

func TestCarbon_AddMinutes(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddMinutes(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddMinutes(1),
			want:   "2020-08-05 13:15:15.222222222 +0800 CST",
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddMinutes(2),
			want:   "2020-08-05 13:16:15.222222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "AddMinutes()")
		})
	}
}

func TestCarbon_SubMinutes(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubMinutes(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubMinutes(1),
			want:   "2020-08-05 13:13:15.222222222 +0800 CST",
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubMinutes(2),
			want:   "2020-08-05 13:12:15.222222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "SubMinutes()")
		})
	}
}

func TestCarbon_AddMinute(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddMinute(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddMinute(),
			want:   "2020-08-05 13:15:15.222222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "AddMinute()")
		})
	}
}

func TestCarbon_SubMinute(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubMinute(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubMinute(),
			want:   "2020-08-05 13:13:15.222222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "SubMinute()")
		})
	}
}

func TestCarbon_AddSeconds(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddSeconds(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddSeconds(1),
			want:   "2020-08-05 13:14:16.222222222 +0800 CST",
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddSeconds(2),
			want:   "2020-08-05 13:14:17.222222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "AddSeconds()")
		})
	}
}

func TestCarbon_SubSeconds(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubSeconds(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubSeconds(1),
			want:   "2020-08-05 13:14:14.222222222 +0800 CST",
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubSeconds(2),
			want:   "2020-08-05 13:14:13.222222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "SubSeconds()")
		})
	}
}

func TestCarbon_AddSecond(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddSecond(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddSecond(),
			want:   "2020-08-05 13:14:16.222222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "AddSecond()")
		})
	}
}

func TestCarbon_SubSecond(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubSecond(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubSecond(),
			want:   "2020-08-05 13:14:14.222222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "SubSecond()")
		})
	}
}

func TestCarbon_AddMilliseconds(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddMilliseconds(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddMilliseconds(1),
			want:   "2020-08-05 13:14:15.223222222 +0800 CST",
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddMilliseconds(2),
			want:   "2020-08-05 13:14:15.224222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "AddMilliseconds()")
		})
	}
}

func TestCarbon_SubMilliseconds(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubMilliseconds(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubMilliseconds(1),
			want:   "2020-08-05 13:14:15.221222222 +0800 CST",
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubMilliseconds(2),
			want:   "2020-08-05 13:14:15.220222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "SubMilliseconds()")
		})
	}
}

func TestCarbon_AddMillisecond(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddMillisecond(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddMillisecond(),
			want:   "2020-08-05 13:14:15.223222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "AddMillisecond()")
		})
	}
}

func TestCarbon_SubMillisecond(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubMillisecond(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubMillisecond(),
			want:   "2020-08-05 13:14:15.221222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "SubMillisecond()")
		})
	}
}

func TestCarbon_AddMicroseconds(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddMicroseconds(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddMicroseconds(1),
			want:   "2020-08-05 13:14:15.222223222 +0800 CST",
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddMicroseconds(2),
			want:   "2020-08-05 13:14:15.222224222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "SubMicroseconds()")
		})
	}
}

func TestCarbon_SubMicroseconds(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubMicroseconds(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubMicroseconds(1),
			want:   "2020-08-05 13:14:15.222221222 +0800 CST",
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubMicroseconds(2),
			want:   "2020-08-05 13:14:15.222220222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "SubMicroseconds()")
		})
	}
}

func TestCarbon_AddMicrosecond(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddMicrosecond(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddMicrosecond(),
			want:   "2020-08-05 13:14:15.222223222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "AddMicrosecond()")
		})
	}
}

func TestCarbon_SubMicrosecond(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubMicrosecond(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubMicrosecond(),
			want:   "2020-08-05 13:14:15.222221222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "SubMicrosecond()")
		})
	}
}

func TestCarbon_AddNanoseconds(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddNanoseconds(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddNanoseconds(1),
			want:   "2020-08-05 13:14:15.222222223 +0800 CST",
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddNanoseconds(2),
			want:   "2020-08-05 13:14:15.222222224 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "AddNanoseconds()")
		})
	}
}

func TestCarbon_SubNanoseconds(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubNanoseconds(0),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubNanoseconds(1),
			want:   "2020-08-05 13:14:15.222222221 +0800 CST",
		},
		{
			name:   "case3",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubNanoseconds(2),
			want:   "2020-08-05 13:14:15.22222222 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "SubNanoseconds()")
		})
	}
}

func TestCarbon_AddNanosecond(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").AddNanosecond(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").AddNanosecond(),
			want:   "2020-08-05 13:14:15.222222223 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "AddNanosecond()")
		})
	}
}

func TestCarbon_SubNanosecond(t *testing.T) {
	tests := []struct {
		name   string
		carbon Carbon
		want   string
	}{
		{
			name:   "case1",
			carbon: Parse("").SubNanosecond(),
			want:   "",
		},
		{
			name:   "case2",
			carbon: Parse("2020-08-05 13:14:15.222222222").SubNanosecond(),
			want:   "2020-08-05 13:14:15.222222221 +0800 CST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.carbon.ToString(), "SubNanosecond()")
		})
	}
}
