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
func (c *Carbon) Constellation() string {
	if c.IsInvalid() {
		return ""
	}

	lang := c.lang
	if lang == nil {
		return ""
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

	lang.rw.RLock()
	defer lang.rw.RUnlock()

	if resources, ok := lang.resources["constellations"]; ok {
		slice := strings.Split(resources, "|")
		if len(slice) == MonthsPerYear {
			return slice[index]
		}
	}
	return ""
}

// IsAries reports whether is Aries.
func (c *Carbon) IsAries() bool {
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
func (c *Carbon) IsTaurus() bool {
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
func (c *Carbon) IsGemini() bool {
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
func (c *Carbon) IsCancer() bool {
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
func (c *Carbon) IsLeo() bool {
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
func (c *Carbon) IsVirgo() bool {
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
func (c *Carbon) IsLibra() bool {
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
func (c *Carbon) IsScorpio() bool {
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
func (c *Carbon) IsSagittarius() bool {
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
func (c *Carbon) IsCapricorn() bool {
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
func (c *Carbon) IsAquarius() bool {
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
func (c *Carbon) IsPisces() bool {
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
