package hebrew

import (
	"fmt"
	"time"

	"github.com/dromara/carbon/v2/calendar"
)

type Locale string

const (
	EnLocale      Locale = "en"
	HeLocale      Locale = "he"
	defaultLocale        = EnLocale

	// hebrewEpochRD is the Rata Die (fixed day) number of 1 Tishri of Hebrew
	// year 1, and jdnRataDie is the Julian Day Number at noon of Rata Die day 1
	// (1 January 1 CE in the proleptic Gregorian calendar). Together they place
	// the Hebrew and Gregorian conversions on the same integer Julian Day axis.
	hebrewEpochRD = -1373427
	jdnRataDie    = 1721425
)

var (
	EnMonths = []string{"Nisan", "Iyyar", "Sivan", "Tammuz", "Av", "Elul", "Tishri", "Heshvan", "Kislev", "Teveth", "Shevat", "Adar", "Adar Bet"}
	HeMonths = []string{"ניסן", "אייר", "סיוון", "תמוז", "אב", "אלול", "תשרי", "חשוון", "כסלו", "טבת", "שבט", "אדר", "אדר ב"}
	EnWeeks  = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	HeWeeks  = []string{"ראשון", "שני", "שלישי", "רביעי", "חמישי", "שישי", "שבת"}
)

type Hebrew struct {
	year, month, day int
	Error            error
}

// NewHebrew creates a new Hebrew calendar instance with specified year, month, and day
func NewHebrew(year, month, day int) *Hebrew {
	h := &Hebrew{year: year, month: month, day: day}
	if !h.IsValid() {
		h.Error = fmt.Errorf("invalid Hebrew date: %d-%d-%d", year, month, day)
	}
	return h
}

// FromStdTime converts standard time to Hebrew calendar date
func FromStdTime(t time.Time) *Hebrew {
	if t.IsZero() {
		return nil
	}

	jdn := gregorian2jdn(t.Year(), int(t.Month()), t.Day())
	y, m, d := jdn2hebrew(jdn)
	return &Hebrew{year: y, month: m, day: d}
}

// ToGregorian converts Hebrew date to Gregorian date
func (h *Hebrew) ToGregorian(timezone ...string) *calendar.Gregorian {
	g := new(calendar.Gregorian)
	if h == nil {
		return g
	}
	loc := time.UTC
	if len(timezone) > 0 {
		loc, g.Error = time.LoadLocation(timezone[0])
	}
	if g.Error != nil {
		return g
	}
	year, month, day := jdn2gregorian(hebrew2jdn(h.year, h.month, h.day))
	g.Time = time.Date(year, time.Month(month), day, 12, 0, 0, 0, loc)
	return g
}

// IsValid checks if the Hebrew date is valid
func (h *Hebrew) IsValid() bool {
	if h == nil || h.Error != nil {
		return false
	}
	// Hebrew year range: 1-9999, including 3761 (corresponding to 1 CE)
	if h.year < 1 || h.year > 9999 || h.month < 1 || h.month > 13 || h.day < 1 || h.day > 31 {
		return false
	}
	// Check if month is within valid range for the year
	if h.month > getMonthsInYear(h.year) {
		return false
	}
	// Check if day is within valid range for the month
	if h.day > getDaysInMonth(h.year, h.month) {
		return false
	}
	return true
}

// IsLeapYear checks if the Hebrew year is a leap year
func (h *Hebrew) IsLeapYear() bool {
	if !h.IsValid() {
		return false
	}
	return ((7*h.year + 1) % 19) < 7
}

// Year returns the Hebrew year
func (h *Hebrew) Year() int {
	if !h.IsValid() {
		return 0
	}
	return h.year
}

// Month returns the Hebrew month (1-13, where 13 is Adar Bet in leap years)
func (h *Hebrew) Month() int {
	if !h.IsValid() {
		return 0
	}
	return h.month
}

// Day returns the day of the Hebrew month
func (h *Hebrew) Day() int {
	if !h.IsValid() {
		return 0
	}
	return h.day
}

// String returns the Hebrew date in "YYYY-MM-DD" format
func (h *Hebrew) String() string {
	if !h.IsValid() {
		return ""
	}
	return fmt.Sprintf("%04d-%02d-%02d", h.year, h.month, h.day)
}

// ToMonthString returns the Hebrew month name in the specified locale
func (h *Hebrew) ToMonthString(locale ...Locale) string {
	if !h.IsValid() {
		return ""
	}
	loc := defaultLocale
	if len(locale) > 0 {
		loc = locale[0]
	}
	idx := h.month - 1
	switch loc {
	case EnLocale:
		if idx >= 0 && idx < len(EnMonths) {
			return EnMonths[idx]
		}
	case HeLocale:
		if idx >= 0 && idx < len(HeMonths) {
			return HeMonths[idx]
		}
	}
	return ""
}

// ToWeekString returns the weekday name in the specified locale
func (h *Hebrew) ToWeekString(locale ...Locale) string {
	if !h.IsValid() {
		return ""
	}
	loc := defaultLocale
	if len(locale) > 0 {
		loc = locale[0]
	}
	// JDN at noon modulo 7: 0 = Monday, so +1 shifts to a Sunday-based index.
	weekday := (hebrew2jdn(h.year, h.month, h.day) + 1) % 7
	switch loc {
	case EnLocale:
		return EnWeeks[weekday]
	case HeLocale:
		return HeWeeks[weekday]
	}
	return ""
}

// gregorian2jdn converts a Gregorian date to its Julian Day Number (at noon).
func gregorian2jdn(year, month, day int) int {
	a := (14 - month) / 12
	y := year + 4800 - a
	m := month + 12*a - 3
	return day + (153*m+2)/5 + 365*y + y/4 - y/100 + y/400 - 32045
}

// jdn2gregorian converts a Julian Day Number to a Gregorian date.
func jdn2gregorian(jdn int) (year, month, day int) {
	a := jdn + 32044
	b := (4*a + 3) / 146097
	c := a - 146097*b/4
	d := (4*c + 3) / 1461
	e := c - 1461*d/4
	m := (5*e + 2) / 153
	day = e - (153*m+2)/5 + 1
	month = m + 3 - 12*(m/10)
	year = 100*b + d - 4800 + m/10
	return
}

// jdn2hebrew converts a Julian Day Number to a Hebrew date.
func jdn2hebrew(jdn int) (year, month, day int) {
	// Estimate the year, then correct it against the true Rosh Hashanah bounds.
	year = (jdn - jdnRataDie - hebrewEpochRD) / 365
	if year < 1 {
		year = 1
	}
	for getJDNInYear(year+1) <= jdn {
		year++
	}
	for getJDNInYear(year) > jdn {
		year--
	}

	// Months run Tishri (7) .. Adar/Adar Bet, then Nisan (1) .. Elul (6).
	maxMonth := getMonthsInYear(year)
	month = 7
	for month != 6 && jdn >= hebrew2jdn(year, month, 1)+getDaysInMonth(year, month) {
		if month == maxMonth {
			month = 1
		} else {
			month++
		}
	}

	day = jdn - hebrew2jdn(year, month, 1) + 1
	return year, month, day
}

// hebrew2jdn converts a Hebrew date to its Julian Day Number (at noon).
func hebrew2jdn(year, month, day int) int {
	jdn := getJDNInYear(year)

	monthOffset := 0
	if month < 7 {
		for m := 7; m <= getMonthsInYear(year); m++ {
			monthOffset += getDaysInMonth(year, m)
		}
		for m := 1; m < month; m++ {
			monthOffset += getDaysInMonth(year, m)
		}
	} else {
		for m := 7; m < month; m++ {
			monthOffset += getDaysInMonth(year, m)
		}
	}

	return jdn + monthOffset + (day - 1)
}

// isLeapYear checks if the Hebrew year is a leap year
func isLeapYear(year int) bool {
	return ((7*year + 1) % 19) < 7
}

// getMonthsFromEpoch calculates the number of months elapsed since the Hebrew epoch
func getMonthsFromEpoch(year int) int {
	cycles := (year - 1) / 19
	yearInCycle := (year - 1) % 19
	return 235*cycles + 12*yearInCycle + (7*yearInCycle+1)/19
}

// getElapsedDays returns the number of days from the Hebrew epoch to Tishri 1
// of the given year, applying the molad-zaken and ADU postponements.
func getElapsedDays(year int) int {
	months := getMonthsFromEpoch(year)
	parts := 12084 + 13753*months
	day := 29*months + parts/25920
	// If the molad falls at or after the relevant hour, Rosh Hashanah is
	// postponed a day (combines molad zaken with the lo ADU rosh rule).
	if (3*(day+1))%7 < 3 {
		day++
	}
	return day
}

// getJDNInYear returns the Julian Day Number (at noon) of Hebrew New Year
// (Tishri 1) for the given year.
func getJDNInYear(year int) int {
	elapsed := getElapsedDays(year)
	// Length-based postponements: a 356-day span forces a two-day delay and a
	// 382-day span a one-day delay, keeping every year a legal length.
	correction := 0
	switch {
	case getElapsedDays(year+1)-elapsed == 356:
		correction = 2
	case elapsed-getElapsedDays(year-1) == 382:
		correction = 1
	}
	return hebrewEpochRD + elapsed + correction + jdnRataDie
}

// getMonthsInYear calculates the number of months in a year
func getMonthsInYear(year int) int {
	if isLeapYear(year) {
		return 13
	}
	return 12
}

// getDaysInMonth calculates the number of days in a month
func getDaysInMonth(year, month int) int {
	// Fixed 29-day months
	if month == 2 || month == 4 || month == 6 || month == 10 || month == 13 {
		return 29
	}

	// Adar in non-leap years is 29 days
	if month == 12 && !isLeapYear(year) {
		return 29
	}

	// Calculate total days in the year
	yearDays := getJDNInYear(year+1) - getJDNInYear(year)

	// Heshvan (month 8) is long only in a complete year (355 or 385 days).
	if month == 8 {
		if yearDays == 355 || yearDays == 385 {
			return 30
		}
		return 29
	}

	// Kislev (month 9) is short only in a deficient year (353 or 383 days).
	if month == 9 {
		if yearDays == 353 || yearDays == 383 {
			return 29
		}
		return 30
	}

	// Other months are 30 days
	return 30
}
