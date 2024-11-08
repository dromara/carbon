package carbon

import (
	"strings"
)

var constellations = []struct {
	startMonth, startDay int
	endMonth, endDay     int
}{
	{3, 21, 4, 19},   // Aries
	{4, 20, 5, 20},   // Taurus
	{5, 21, 6, 21},   // Gemini
	{6, 22, 7, 22},   // Cancer
	{7, 23, 8, 22},   // Leo
	{8, 23, 9, 22},   // Virgo
	{9, 23, 10, 23},  // Libra
	{10, 24, 11, 22}, // Scorpio
	{11, 23, 12, 21}, // Sagittarius
	{12, 22, 1, 19},  // Capricorn
	{1, 20, 2, 18},   // Aquarius
	{2, 19, 3, 20},   // Pisces
}

// Constellation gets constellation name like "Aries", i18n is supported.
// 获取星座，支持i18n
func (c Carbon) Constellation() string {
	if c.IsInvalid() {
		return ""
	}
	if len(c.lang.resources) == 0 {
		c.lang.SetLocale(defaultLocale)
	}
	index := -1
	_, month, day := c.Date()
	for i := 0; i < len(constellations); i++ {
		constellation := constellations[i]
		if month == constellation.startMonth && day >= constellation.startDay {
			index = i
		}
		if month == constellation.endMonth && day <= constellation.endDay {
			index = i
		}
	}
	c.lang.rw.RLock()
	defer c.lang.rw.RUnlock()
	if resources, ok := c.lang.resources["constellations"]; ok {
		slice := strings.Split(resources, "|")
		if len(slice) == MonthsPerYear {
			return slice[index]
		}
	}
	return ""
}

// IsAries reports whether is Aries.
// 是否是白羊座
func (c Carbon) IsAries() bool {
	if c.IsInvalid() {
		return false
	}
	_, month, day := c.Date()
	if month == 3 && day >= 21 {
		return true
	}
	if month == 4 && day <= 19 {
		return true
	}
	return false
}

// IsTaurus reports whether is Taurus.
// 是否是金牛座
func (c Carbon) IsTaurus() bool {
	if c.IsInvalid() {
		return false
	}
	_, month, day := c.Date()
	if month == 4 && day >= 20 {
		return true
	}
	if month == 5 && day <= 20 {
		return true
	}
	return false
}

// IsGemini reports whether is Gemini.
// 是否是双子座
func (c Carbon) IsGemini() bool {
	if c.IsInvalid() {
		return false
	}
	_, month, day := c.Date()
	if month == 5 && day >= 21 {
		return true
	}
	if month == 6 && day <= 21 {
		return true
	}
	return false
}

// IsCancer reports whether is Cancer.
// 是否是巨蟹座
func (c Carbon) IsCancer() bool {
	if c.IsInvalid() {
		return false
	}
	_, month, day := c.Date()
	if month == 6 && day >= 22 {
		return true
	}
	if month == 7 && day <= 22 {
		return true
	}
	return false
}

// IsLeo reports whether is Leo.
// 是否是狮子座
func (c Carbon) IsLeo() bool {
	if c.IsInvalid() {
		return false
	}
	_, month, day := c.Date()
	if month == 7 && day >= 23 {
		return true
	}
	if month == 8 && day <= 22 {
		return true
	}
	return false
}

// IsVirgo reports whether is Virgo.
// 是否是处女座
func (c Carbon) IsVirgo() bool {
	if c.IsInvalid() {
		return false
	}
	_, month, day := c.Date()
	if month == 8 && day >= 23 {
		return true
	}
	if month == 9 && day <= 22 {
		return true
	}
	return false
}

// IsLibra reports whether is Libra.
// 是否是天秤座
func (c Carbon) IsLibra() bool {
	if c.IsInvalid() {
		return false
	}
	_, month, day := c.Date()
	if month == 9 && day >= 23 {
		return true
	}
	if month == 10 && day <= 23 {
		return true
	}
	return false
}

// IsScorpio reports whether is Scorpio.
// 是否是天蝎座
func (c Carbon) IsScorpio() bool {
	if c.IsInvalid() {
		return false
	}
	_, month, day := c.Date()
	if month == 10 && day >= 24 {
		return true
	}
	if month == 11 && day <= 22 {
		return true
	}
	return false
}

// IsSagittarius reports whether is Sagittarius.
// 是否是射手座
func (c Carbon) IsSagittarius() bool {
	if c.IsInvalid() {
		return false
	}
	_, month, day := c.Date()
	if month == 11 && day >= 22 {
		return true
	}
	if month == 12 && day <= 21 {
		return true
	}
	return false
}

// IsCapricorn reports whether is Capricorn.
// 是否是摩羯座
func (c Carbon) IsCapricorn() bool {
	if c.IsInvalid() {
		return false
	}
	_, month, day := c.Date()
	if month == 12 && day >= 22 {
		return true
	}
	if month == 1 && day <= 19 {
		return true
	}
	return false
}

// IsAquarius reports whether is Aquarius.
// 是否是水瓶座
func (c Carbon) IsAquarius() bool {
	if c.IsInvalid() {
		return false
	}
	_, month, day := c.Date()
	if month == 1 && day >= 20 {
		return true
	}
	if month == 2 && day <= 18 {
		return true
	}
	return false
}

// IsPisces reports whether is Pisces.
// 是否是双鱼座
func (c Carbon) IsPisces() bool {
	if c.IsInvalid() {
		return false
	}
	_, month, day := c.Date()
	if month == 2 && day >= 19 {
		return true
	}
	if month == 3 && day <= 20 {
		return true
	}
	return false
}
