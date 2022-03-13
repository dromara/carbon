package carbon

import (
	"strings"
)

// Constellation gets constellation name, i18n is supported.
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
	switch {
	case month == 3 && day >= 21, month == 4 && day <= 19:
		index = 0 // 白羊座
	case month == 4 && day >= 20, month == 5 && day <= 20:
		index = 1 // 金牛座
	case month == 5 && day >= 21, month == 6 && day <= 21:
		index = 2 // 双子座
	case month == 6 && day >= 22, month == 7 && day <= 22:
		index = 3 // 巨蟹座
	case month == 7 && day >= 23, month == 8 && day <= 22:
		index = 4 // 狮子座
	case month == 8 && day >= 23, month == 9 && day <= 22:
		index = 5 // 处女座
	case month == 9 && day >= 23, month == 10 && day <= 23:
		index = 6 // 天秤座
	case month == 10 && day >= 24, month == 11 && day <= 22:
		index = 7 // 天蝎座
	case month == 11 && day >= 23, month == 12 && day <= 21:
		index = 8 // 射手座
	case month == 12 && day >= 22, month == 1 && day <= 19:
		index = 9 // 摩羯座
	case month == 1 && day >= 20, month == 2 && day <= 18:
		index = 10 // 水瓶座
	case month == 2 && day >= 19, month == 3 && day <= 20:
		index = 11 // 双鱼座
	}
	if constellations, ok := c.lang.resources["constellations"]; ok {
		slice := strings.Split(constellations, "|")
		if len(slice) == 12 {
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
