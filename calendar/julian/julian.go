// Package julian is part of the carbon package.
package julian

import (
	"math"
	"strconv"
	"time"

	"github.com/dromara/carbon/v2/calendar"
)

var (
	// julian day or modified julian day decimal precision
	decimalPrecision = 6

	// difference between Julian Day and Modified Julian Day
	diffJdFromMjd = 2400000.5
)

// Julian defines a Julian struct.
type Julian struct {
	jd, mjd float64
}

// NewJulian returns a new Lunar instance.
func NewJulian(f float64) (j Julian) {
	// get length of the integer part
	l := len(strconv.Itoa(int(math.Ceil(f))))
	switch l {
	// modified julian day
	case 5:
		j.mjd = f
		j.jd = f + diffJdFromMjd
	// julian day
	case 7:
		j.jd = f
		j.mjd = f - diffJdFromMjd
	default:
		j.jd = 0
		j.mjd = 0
	}
	return
}

// FromStdTime creates a Julian instance from standard time.Time.
func FromStdTime(t time.Time) (j Julian) {
	if t.IsZero() {
		j.jd = 1721423.5
		j.mjd = -678577
		return j
	}
	y := t.Year()
	m := int(t.Month())
	d := float64(t.Day()) + ((float64(t.Second())/60+float64(t.Minute()))/60+float64(t.Hour()))/24
	n := 0
	f := false
	if y*372+m*31+int(d) >= 588829 {
		f = true
	}
	if m <= 2 {
		m += 12
		y--
	}
	if f {
		n = y / 100
		n = 2 - n + n/4
	}
	jd := float64(int(365.25*(float64(y)+4716))) + float64(int(30.6001*(float64(m)+1))) + d + float64(n) - 1524.5
	return NewJulian(jd)
}

// ToGregorian converts Julian instance to Gregorian instance.
func (j Julian) ToGregorian(timezone ...string) (g calendar.Gregorian) {
	loc := time.UTC
	if len(timezone) > 0 {
		loc, g.Error = time.LoadLocation(timezone[0])
	}
	if g.Error != nil {
		return g
	}
	d := int(j.jd + 0.5)
	f := j.jd + 0.5 - float64(d)
	if d >= 2299161 {
		c := int((float64(d) - 1867216.25) / 36524.25)
		d += 1 + c - c/4
	}
	d += 1524
	year := int((float64(d) - 122.1) / 365.25)
	d -= int(365.25 * float64(year))
	month := int(float64(d) / 30.601)
	d -= int(30.601 * float64(month))
	day := d
	if month > 13 {
		month -= 13
		year -= 4715
	} else {
		month -= 1
		year -= 4716
	}
	f *= 24
	hour := int(f)

	f -= float64(hour)
	f *= 60
	minute := int(f)

	f -= float64(minute)
	f *= 60
	second := int(math.Round(f))
	g.Time = time.Date(year, time.Month(month), day, hour, minute, second, 0, loc)
	return g
}

// JD gets julian day like 2460332.5
func (j Julian) JD(precision ...int) float64 {
	p := decimalPrecision
	if len(precision) > 0 {
		p = precision[0]
	}
	return parseFloat64(j.jd, p)
}

// MJD gets modified julian day like 60332
func (j Julian) MJD(precision ...int) float64 {
	p := decimalPrecision
	if len(precision) > 0 {
		p = precision[0]
	}
	return parseFloat64(j.mjd, p)
}

// parseFloat64 round to n decimal places
func parseFloat64(f float64, n int) float64 {
	p10 := math.Pow10(n)
	return math.Round(f*p10) / p10
}
