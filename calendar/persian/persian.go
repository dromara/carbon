// Package persian is part of the carbon package.
package persian

import (
	"fmt"
	"time"

	"github.com/dromara/carbon/v2/calendar"
	"github.com/dromara/carbon/v2/calendar/julian"
)

type Locale string

const (
	EnLocale      Locale = "en"
	FaLocale      Locale = "fa"
	defaultLocale        = EnLocale
	persianEpoch         = 1948320
)

var (
	EnMonths = []string{"Farvardin", "Ordibehesht", "Khordad", "Tir", "Mordad", "Shahrivar", "Mehr", "Aban", "Azar", "Dey", "Bahman", "Esfand"}
	FaMonths = []string{"فروردین", "اردیبهشت", "خرداد", "تیر", "مرداد", "شهریور", "مهر", "آبان", "آذر", "دی", "بهمن", "اسفند"}

	EnWeeks = []string{"Yekshanbeh", "Doshanbeh", "Seshanbeh", "Chaharshanbeh", "Panjshanbeh", "Jomeh", "Shanbeh"}
	FaWeeks = []string{"یکشنبه", "دوشنبه", "سه شنبه", "چهارشنبه", "پنجشنبه", "جمعه", "شنبه"}
)

// Persian defines a Persian struct.
type Persian struct {
	year, month, day int
	Error            error
}

// NewPersian returns a new Persian instance.
func NewPersian(year, month, day int) *Persian {
	p := &Persian{year: year, month: month, day: day}
	if !p.IsValid() {
		p.Error = fmt.Errorf("invalid persian date: %04d-%02d-%02d", year, month, day)
	}
	return p
}

// FromStdTime creates a Persian instance from standard time.Time.
func FromStdTime(t time.Time) (p *Persian) {
	if t.IsZero() {
		return nil
	}
	gjdn := int(julian.FromStdTime(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())).JD(0))
	year, month, day := jdn2persian(gjdn)
	return &Persian{year: year, month: month, day: day}
}

// ToGregorian converts Persian instance to Gregorian instance.
func (p *Persian) ToGregorian(timezone ...string) *calendar.Gregorian {
	g := new(calendar.Gregorian)
	if !p.IsValid() {
		return g
	}
	loc := time.UTC
	if len(timezone) > 0 {
		loc, g.Error = time.LoadLocation(timezone[0])
	}
	if g.Error != nil {
		return g
	}
	jdn := getPersianJdn(p.year, p.month, p.day)

	l := jdn + 68569
	n := 4 * l / 146097
	l = l - (146097*n+3)/4
	i := 4000 * (l + 1) / 1461001
	l = l - 1461*i/4 + 31
	j := 80 * l / 2447
	day := l - 2447*j/80
	l = j / 11
	month := j + 2 - 12*l
	year := 100*(n-49) + i + l

	g.Time = time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
	return g
}

// Year gets the Persian year like 2020.
func (p *Persian) Year() int {
	if !p.IsValid() {
		return 0
	}
	return p.year
}

// Month gets the Persian month like 8.
func (p *Persian) Month() int {
	if !p.IsValid() {
		return 0
	}
	return p.month
}

// Day gets the Persian day like 5.
func (p *Persian) Day() int {
	if !p.IsValid() {
		return 0
	}
	return p.day
}

// String implements the "Stringer" interface for Persian.
func (p *Persian) String() string {
	if !p.IsValid() {
		return ""
	}
	return fmt.Sprintf("%04d-%02d-%02d", p.year, p.month, p.day)
}

// ToMonthString outputs a string in Persian month format like "فروردین".
func (p *Persian) ToMonthString(locale ...Locale) (month string) {
	if !p.IsValid() {
		return ""
	}
	loc := defaultLocale
	if len(locale) > 0 {
		loc = locale[0]
	}
	switch loc {
	case EnLocale:
		return EnMonths[p.month-1]
	case FaLocale:
		return FaMonths[p.month-1]
	}
	return ""
}

// ToWeekString outputs a string in week layout like "چهارشنبه".
func (p *Persian) ToWeekString(locale ...Locale) (month string) {
	if !p.IsValid() {
		return ""
	}
	loc := defaultLocale
	if len(locale) > 0 {
		loc = locale[0]
	}
	week := p.ToGregorian().Time.Weekday()
	switch loc {
	case EnLocale:
		return EnWeeks[week]
	case FaLocale:
		return FaWeeks[week]
	}
	return ""
}

// IsValid reports whether the Persian date is valid.
func (p *Persian) IsValid() bool {
	if p == nil || p.Error != nil {
		return false
	}
	// Check year range validation (Persian calendar starts from 622 CE)
	if p.year < 1 || p.year > 9999 || p.month <= 0 || p.month > 12 || p.day <= 0 || p.day > 31 {
		return false
	}
	// Check month-specific day validation
	if p.month > 6 && p.month <= 11 && p.day > 30 {
		return false
	}
	if p.month == 12 {
		// Use IsLeapYear method
		if (!p.IsLeapYear() && p.day > 29) || (p.IsLeapYear() && p.day > 30) {
			return false
		}
	}
	return true
}

// IsLeapYear reports whether the Persian year is a leap year.
func (p *Persian) IsLeapYear() bool {
	if p == nil || p.Error != nil {
		return false
	}
	currentYearJdn := getPersianJdn(p.year, 1, 1)
	nextYearJdn := getPersianJdn(p.year+1, 1, 1)
	daysDiff := nextYearJdn - currentYearJdn
	return daysDiff > 365
}

// getPersianYear gets the Persian year from Julian Day Number.
func getPersianYear(jdn int) int {
	days := jdn - persianEpoch
	year := 474 + days/365
	if year < 1 || year > 9999 {
		return -1
	}
	low := 1
	high := 9999
	for low <= high {
		mid := (low + high) / 2
		yearStartJdn := getPersianJdn(mid, 1, 1)
		nextYearStartJdn := getPersianJdn(mid+1, 1, 1)
		if jdn >= yearStartJdn && jdn < nextYearStartJdn {
			return mid
		}
		if jdn < yearStartJdn {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// getPersianJdn gets the Julian day number in the Persian calendar.
func getPersianJdn(year, month, day int) int {
	yearOffset := year - 474
	if yearOffset < 0 {
		yearOffset--
	}
	cycleYear := 474 + (yearOffset % 2820)
	var monthDays int
	if month <= 7 {
		monthDays = (month - 1) * 31
	} else {
		monthDays = (month-1)*30 + 6
	}
	return day + monthDays + (cycleYear*682-110)/2816 + (cycleYear-1)*365 + yearOffset/2820*1029983 + persianEpoch
}

// jdn2persian converts Julian Day Number to Persian date (year, month, day).
func jdn2persian(jdn int) (year, month, day int) {
	year = getPersianYear(jdn)
	days := jdn - getPersianJdn(year, 1, 1) + 1
	if days <= 186 {
		month = (days-1)/31 + 1
	} else {
		month = (days-186-1)/30 + 7
	}
	day = jdn - getPersianJdn(year, month, 1) + 1
	return
}
