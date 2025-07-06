package carbon

import (
	"strings"
)

var seasons = []struct {
	month, index int
}{
	{3, 0},  // spring
	{4, 0},  // spring
	{5, 0},  // spring
	{6, 1},  // summer
	{7, 1},  // summer
	{8, 1},  // summer
	{9, 2},  // autumn
	{10, 2}, // autumn
	{11, 2}, // autumn
	{12, 3}, // winter
	{1, 3},  // winter
	{2, 3},  // winter
}

// Season gets season name according to the meteorological division method like "Spring", i18n is supported.
func (c Carbon) Season() string {
	if c.IsInvalid() {
		return ""
	}
	index := -1
	month := c.Month()
	for i := 0; i < len(seasons); i++ {
		season := seasons[i]
		if month == season.month {
			index = season.index
		}
	}

	c.lang.rw.RLock()
	defer c.lang.rw.RUnlock()

	if resources, ok := c.lang.resources["seasons"]; ok {
		slice := strings.Split(resources, "|")
		if len(slice) == QuartersPerYear {
			return slice[index]
		}
	}
	return ""
}

// StartOfSeason returns a Carbon instance for start of the season.
func (c Carbon) StartOfSeason() Carbon {
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
func (c Carbon) EndOfSeason() Carbon {
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
func (c Carbon) IsSpring() bool {
	if c.IsInvalid() {
		return false
	}
	month := c.Month()
	if month == 3 || month == 4 || month == 5 {
		return true
	}
	return false
}

// IsSummer reports whether is summer.
func (c Carbon) IsSummer() bool {
	if c.IsInvalid() {
		return false
	}
	month := c.Month()
	if month == 6 || month == 7 || month == 8 {
		return true
	}
	return false
}

// IsAutumn reports whether is autumn.
func (c Carbon) IsAutumn() bool {
	if c.IsInvalid() {
		return false
	}
	month := c.Month()
	if month == 9 || month == 10 || month == 11 {
		return true
	}
	return false
}

// IsWinter reports whether is winter.
func (c Carbon) IsWinter() bool {
	if c.IsInvalid() {
		return false
	}
	month := c.Month()
	if month == 12 || month == 1 || month == 2 {
		return true
	}
	return false
}
