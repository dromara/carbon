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
// 获取当前季节(以气象划分)，支持 i18n
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
// 本季节开始时间
func (c Carbon) StartOfSeason() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _ := c.Date()
	if month == 1 || month == 2 {
		return create(year-1, 12, 1, 0, 0, 0, 0, c.Timezone())
	}
	return create(year, month/3*3, 1, 0, 0, 0, 0, c.Timezone())
}

// EndOfSeason returns a Carbon instance for end of the season.
// 本季节结束时间
func (c Carbon) EndOfSeason() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _ := c.Date()
	if month == 1 || month == 2 {
		return create(year, 3, 0, 23, 59, 59, 999999999, c.Timezone())
	}
	if month == 12 {
		return create(year+1, 3, 0, 23, 59, 59, 999999999, c.Timezone())
	}
	return create(year, month/3*3+3, 0, 23, 59, 59, 999999999, c.Timezone())
}

// IsSpring reports whether is spring.
// 是否是春季
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
// 是否是夏季
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
// 是否是秋季
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
// 是否是冬季
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
