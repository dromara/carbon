package carbon

import (
	"strings"
)

var seasons = map[int]int{
	// month: index
	1:  3, // winter
	2:  3, // winter
	3:  0, // spring
	4:  0, // spring
	5:  0, // spring
	6:  1, // summer
	7:  1, // summer
	8:  1, // summer
	9:  2, // autumn
	10: 2, // autumn
	11: 2, // autumn
	12: 3, // winter
}

// Season gets season name according to the meteorological division method like "Spring", i18n is supported.
func (c *Carbon) Season() string {
	if c.IsInvalid() {
		return ""
	}

	lang := c.lang
	if lang == nil {
		return ""
	}

	lang.rw.RLock()
	defer lang.rw.RUnlock()

	if resources, ok := lang.resources["seasons"]; ok {
		slice := strings.Split(resources, "|")
		if len(slice) == QuartersPerYear {
			return slice[seasons[c.Month()]]
		}
	}
	return ""
}

// StartOfSeason returns a Carbon instance for start of the season.
func (c *Carbon) StartOfSeason() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _ := c.Date()
	if month == 1 || month == 2 {
		return c.create(year-1, MaxMonth, MinDay, MinHour, MinMinute, MinSecond, MinNanosecond)
	}
	return c.create(year, month/3*3, MinDay, MinHour, MinMinute, MinSecond, MinNanosecond)
}

// EndOfSeason returns a Carbon instance for end of the season.
func (c *Carbon) EndOfSeason() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _ := c.Date()
	if month == 1 || month == 2 {
		return c.create(year, 3, 0, MaxHour, MaxMinute, MaxSecond, MaxNanosecond)
	}
	if month == 12 {
		return c.create(year+1, 3, 0, MaxHour, MaxMinute, MaxSecond, MaxNanosecond)
	}
	return c.create(year, month/3*3+3, 0, MaxHour, MaxMinute, MaxSecond, MaxNanosecond)
}

// IsSpring reports whether is spring.
func (c *Carbon) IsSpring() bool {
	if c.IsInvalid() {
		return false
	}
	month := c.Month()
	return month == 3 || month == 4 || month == 5
}

// IsSummer reports whether is summer.
func (c *Carbon) IsSummer() bool {
	if c.IsInvalid() {
		return false
	}
	month := c.Month()
	return month == 6 || month == 7 || month == 8
}

// IsAutumn reports whether is autumn.
func (c *Carbon) IsAutumn() bool {
	if c.IsInvalid() {
		return false
	}
	month := c.Month()
	return month == 9 || month == 10 || month == 11
}

// IsWinter reports whether is winter.
func (c *Carbon) IsWinter() bool {
	if c.IsInvalid() {
		return false
	}
	month := c.Month()
	return month == 1 || month == 2 || month == 12
}
