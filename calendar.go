package carbon

import (
	"github.com/dromara/carbon/v2/calendar/julian"
	"github.com/dromara/carbon/v2/calendar/lunar"
	"github.com/dromara/carbon/v2/calendar/persian"
)

// Lunar converts Carbon instance to Lunar instance.
func (c Carbon) Lunar() (l lunar.Lunar) {
	if c.IsEmpty() {
		return l
	}
	if c.HasError() {
		l.Error = c.Error
		return l
	}
	return lunar.FromStdTime(c.StdTime())
}

// CreateFromLunar creates a Carbon instance from Lunar date and time.
func CreateFromLunar(year, month, day int, isLeapMonth bool) Carbon {
	l := lunar.NewLunar(year, month, day, isLeapMonth)
	if !l.IsValid() {
		c := NewCarbon()
		c.Error = l.Error
		return c
	}
	return NewCarbon(l.ToGregorian(DefaultTimezone).Time)
}

// Julian converts Carbon instance to Julian instance.
func (c Carbon) Julian() (j julian.Julian) {
	if c.IsEmpty() {
		return j
	}
	if c.HasError() {
		return j
	}
	return julian.FromStdTime(c.StdTime())
}

// CreateFromJulian creates a Carbon instance from Julian Day or Modified Julian Day.
func CreateFromJulian(f float64) Carbon {
	return NewCarbon(julian.NewJulian(f).ToGregorian(DefaultTimezone).Time)
}

// Persian converts Carbon instance to Persian instance.
func (c Carbon) Persian() (p persian.Persian) {
	if c.IsEmpty() {
		return p
	}
	if c.HasError() {
		p.Error = c.Error
		return p
	}
	return persian.FromStdTime(c.StdTime())
}

// CreateFromPersian creates a Carbon instance from Persian date and time.
func CreateFromPersian(year, month, day int) (c Carbon) {
	p := persian.NewPersian(year, month, day)
	if !p.IsValid() {
		c.Error = p.Error
		return c
	}
	return NewCarbon(p.ToGregorian(DefaultTimezone).Time)
}
