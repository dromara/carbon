// Package persian is part of the carbon package.
package persian

import (
	"fmt"
	"math"
	"time"

	"github.com/dromara/carbon/v2/calendar"
	"github.com/dromara/carbon/v2/calendar/julian"
)

type Locale string

const EnLocale Locale = "en"
const FaLocale Locale = "fa"

const defaultLocale = EnLocale

var (
	EnMonths      = []string{"Farvardin", "Ordibehesht", "Khordad", "Tir", "Mordad", "Shahrivar", "Mehr", "Aban", "Azar", "Dey", "Bahman", "Esfand"}
	ShortEnMonths = []string{"Far", "Ord", "Kho", "Tir", "Mor", "Sha", "Meh", "Aba", "Aza", "Dey", "Bah", "Esf"}

	FaMonths      = []string{"فروردین", "اردیبهشت", "خرداد", "تیر", "مرداد", "شهریور", "مهر", "آبان", "آذر", "دی", "بهمن", "اسفند"}
	ShortFaMonths = []string{"فرو", "ارد", "خرد", "تیر", "مرد", "شهر", "مهر", "آبا", "آذر", "دی", "بهم", "اسف"}

	EnWeeks      = []string{"Yekshanbeh", "Doshanbeh", "Seshanbeh", "Chaharshanbeh", "Panjshanbeh", "Jomeh", "Shanbeh"}
	ShortEnWeeks = []string{"Yek", "Dos", "Ses", "Cha", "Pan", "Jom", "Sha"}

	FaWeeks      = []string{"نجشنبه", "دوشنبه", "سه شنبه", "چهارشنبه", "پنجشنبه", "جمعه", "شنبه"}
	ShortFaWeeks = []string{"پ", "چ", "س", "د", "ی", "ش", "ج"}
)

// ErrInvalidPersian returns a invalid persian date.
var ErrInvalidPersian = func() error {
	return fmt.Errorf("invalid persian date, please make sure the persian date is valid")
}

// Persian defines a Persian struct.
type Persian struct {
	year, month, day int
	Error            error
}

// NewPersian returns a new Persian instance.
func NewPersian(year, month, day int) *Persian {
	p := new(Persian)
	p.year, p.month, p.day = year, month, day
	if !p.IsValid() {
		p.Error = ErrInvalidPersian()
	}
	return p
}

// MaxValue returns a Persian instance for the greatest supported date.
func MaxValue() *Persian {
	return &Persian{
		year:  9377,
		month: 12,
		day:   31,
	}
}

// MinValue returns a Persian instance for the lowest supported date.
func MinValue() *Persian {
	return &Persian{
		year:  1,
		month: 1,
		day:   1,
	}
}

// FromStdTime creates a Persian instance from standard time.Time.
func FromStdTime(t time.Time) (p Persian) {
	if t.IsZero() {
		return p
	}
	gjdn := int(julian.FromStdTime(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())).JD(0))
	pjdn := getPersianJdn(475, 1, 1)

	diff := gjdn - pjdn
	div := diff / 1029983
	mod := diff % 1029983
	p.year = (2134*mod/366+2816*(mod%366)+2815)/1028522 + mod/366 + 1 + 2820*div + 474
	pjdn = getPersianJdn(p.year, 1, 1)
	fjdn := float64(gjdn - pjdn + 1)
	if fjdn <= 186 {
		p.month = int(math.Ceil(fjdn / 31.0))
	} else {
		p.month = int(math.Ceil((fjdn - 6) / 30.0))
	}
	pjdn = getPersianJdn(p.year, p.month, 1)
	p.day = gjdn - pjdn + 1
	return p
}

// ToGregorian converts Persian instance to Gregorian instance.
func (p Persian) ToGregorian(timezone ...string) (g calendar.Gregorian) {
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
	d := l - 2447*j/80
	l = j / 11
	m := j + 2 - 12*l
	y := 100*(n-49) + i + l

	g.Time = time.Date(y, time.Month(m), d, 0, 0, 0, 0, loc)
	return g
}

// Year gets lunar year like 2020.
func (p Persian) Year() int {
	if !p.IsValid() {
		return 0
	}
	return p.year
}

// Month gets lunar month like 8.
func (p Persian) Month() int {
	if !p.IsValid() {
		return 0
	}
	return p.month
}

// Day gets lunar day like 5.
func (p Persian) Day() int {
	if !p.IsValid() {
		return 0
	}
	return p.day
}

// String implements "Stringer" interface for Persian.
func (p Persian) String() string {
	if !p.IsValid() {
		return ""
	}
	return fmt.Sprintf("%04d-%02d-%02d", p.year, p.month, p.day)
}

// ToMonthString outputs a string in persian month format like "فروردین".
func (p Persian) ToMonthString(locale ...Locale) (month string) {
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

// ToShortMonthString outputs a short string in persian month format like "فروردین".
func (p Persian) ToShortMonthString(locale ...Locale) (month string) {
	if !p.IsValid() {
		return ""
	}
	loc := defaultLocale
	if len(locale) > 0 {
		loc = locale[0]
	}
	switch loc {
	case EnLocale:
		return ShortEnMonths[p.month-1]
	case FaLocale:
		return ShortFaMonths[p.month-1]
	}
	return ""
}

// ToWeekString outputs a string in week layout like "چهارشنبه".
func (p Persian) ToWeekString(locale ...Locale) (month string) {
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

// ToShortWeekString outputs a short string in week layout like "چهارشنبه".
func (p Persian) ToShortWeekString(locale ...Locale) (month string) {
	if !p.IsValid() {
		return ""
	}
	loc := defaultLocale
	if len(locale) > 0 {
		loc = locale[0]
	}
	week := p.ToGregorian().Time.Weekday()
	switch loc {
	case "en":
		return ShortEnWeeks[week]
	case "fa":
		return ShortFaWeeks[week]
	}
	return ""
}

// IsValid reports whether is a valid persian date.
func (p Persian) IsValid() bool {
	if p.Error != nil {
		return false
	}
	if p.year >= MinValue().year && p.year <= MaxValue().year {
		return true
	}
	return false
}

// IsLeapYear reports whether is a persian leap year.
func (p Persian) IsLeapYear() bool {
	if !p.IsValid() {
		return false
	}
	return (25*p.year+11)%33 < 8
}

// gets Julian day number in Persian calendar
func getPersianJdn(year, month, day int) int {
	year = year - 473
	if year >= 0 {
		year--
	}
	epy := 474 + (year % 2820)
	var md int
	if month <= 7 {
		md = (month - 1) * 31
	} else {
		md = (month-1)*30 + 6
	}
	return day + md + (epy*682-110)/2816 + (epy-1)*365 + year/2820*1029983 + 1948320
}
