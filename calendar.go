package carbon

import (
	"github.com/dromara/carbon/v2/calendar/hebrew"
	"github.com/dromara/carbon/v2/calendar/julian"
	"github.com/dromara/carbon/v2/calendar/lunar"
	"github.com/dromara/carbon/v2/calendar/persian"
)

// Lunar converts Carbon instance to Lunar instance.
func (c *Carbon) Lunar() *lunar.Lunar {
	if c.IsNil() {
		return nil
	}
	if c.IsZero() || c.IsEmpty() {
		return &lunar.Lunar{}
	}
	if c.HasError() {
		return &lunar.Lunar{Error: c.Error}
	}
	return lunar.FromStdTime(c.StdTime())
}

// CreateFromLunar creates a Carbon instance from Lunar date.
func CreateFromLunar(year, month, day int, isLeapMonth bool) *Carbon {
	l := lunar.NewLunar(year, month, day, isLeapMonth)
	if l.Error != nil {
		return &Carbon{Error: l.Error}
	}
	return NewCarbon(l.ToGregorian(DefaultTimezone).Time)
}

// Julian converts Carbon instance to Julian instance.
func (c *Carbon) Julian() *julian.Julian {
	if c.IsNil() {
		return nil
	}
	if c.IsEmpty() {
		return &julian.Julian{}
	}
	if c.HasError() {
		return &julian.Julian{}
	}
	return julian.FromStdTime(c.StdTime())
}

// CreateFromJulian creates a Carbon instance from Julian Day or Modified Julian Day.
func CreateFromJulian(f float64) *Carbon {
	return NewCarbon(julian.NewJulian(f).ToGregorian(DefaultTimezone).Time)
}

// Persian converts Carbon instance to Persian instance.
func (c *Carbon) Persian() *persian.Persian {
	if c.IsNil() {
		return nil
	}
	if c.IsZero() || c.IsEmpty() {
		return &persian.Persian{}
	}
	if c.HasError() {
		return &persian.Persian{Error: c.Error}
	}
	return persian.FromStdTime(c.StdTime())
}

// CreateFromPersian creates a Carbon instance from Persian date.
func CreateFromPersian(year, month, day int) *Carbon {
	p := persian.NewPersian(year, month, day)
	if p.Error != nil {
		return &Carbon{Error: p.Error}
	}
	return NewCarbon(p.ToGregorian(DefaultTimezone).Time)
}

// Hebrew converts Carbon instance to Hebrew instance.
func (c *Carbon) Hebrew() *hebrew.Hebrew {
	if c.IsNil() {
		return nil
	}
	if c.IsZero() || c.IsEmpty() {
		return &hebrew.Hebrew{}
	}
	if c.HasError() {
		return &hebrew.Hebrew{Error: c.Error}
	}
	return hebrew.FromStdTime(c.StdTime())
}

// CreateFromHebrew creates a Carbon instance from Hebrew date.
func CreateFromHebrew(year, month, day int) *Carbon {
	h := hebrew.NewHebrew(year, month, day)
	if h.Error != nil {
		return &Carbon{Error: h.Error}
	}
	return NewCarbon(h.ToGregorian(DefaultTimezone).Time)
}
