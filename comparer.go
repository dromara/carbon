package carbon

import (
	"time"
)

// HasError reports whether it has error.
func (c *Carbon) HasError() bool {
	if c.IsNil() {
		return false
	}
	return c.Error != nil
}

// IsNil reports whether it is nil pointer.
func (c *Carbon) IsNil() bool {
	return c == nil
}

// IsEmpty reports whether it is empty value.
func (c *Carbon) IsEmpty() bool {
	if c.IsNil() || c.HasError() {
		return false
	}
	return c.isEmpty
}

// IsZero reports whether it is a zero time(0001-01-01 00:00:00 +0000 UTC).
func (c *Carbon) IsZero() bool {
	if c.IsNil() || c.IsEmpty() || c.HasError() {
		return false
	}
	return c.StdTime().IsZero()
}

// IsEpoch reports whether it is a unix epoch time(1970-01-01 00:00:00 +0000 UTC).
func (c *Carbon) IsEpoch() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Eq(EpochValue())
}

// IsValid reports whether it is a valid time.
func (c *Carbon) IsValid() bool {
	if !c.IsNil() && !c.HasError() && !c.IsEmpty() {
		return true
	}
	return false
}

// IsInvalid reports whether it is an invalid time.
func (c *Carbon) IsInvalid() bool {
	return !c.IsValid()
}

// IsDST reports whether it is a daylight saving time.
func (c *Carbon) IsDST() bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().IsDST()
}

// IsAM reports whether it is before noon.
func (c *Carbon) IsAM() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Hour() < 12
}

// IsPM reports whether it is after noon.
func (c *Carbon) IsPM() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Hour() >= 12
}

// IsLeapYear reports whether it is a leap year.
func (c *Carbon) IsLeapYear() bool {
	if c.IsInvalid() {
		return false
	}
	year := c.Year()
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

// IsLongYear reports whether it is a long year,
//
// refer to https://en.wikipedia.org/wiki/ISO_8601#Week_dates.
func (c *Carbon) IsLongYear() bool {
	if c.IsInvalid() {
		return false
	}
	_, w := time.Date(c.Year(), MaxMonth, MaxDay, MinHour, MinMinute, MinSecond, MinNanosecond, c.loc).ISOWeek()
	return w == WeeksPerLongYear
}

// IsJanuary reports whether it is January.
func (c *Carbon) IsJanuary() bool {
	return c.isMonth(time.January)
}

// IsFebruary reports whether it is February.
func (c *Carbon) IsFebruary() bool {
	return c.isMonth(time.February)
}

// IsMarch reports whether it is March.
func (c *Carbon) IsMarch() bool {
	return c.isMonth(time.March)
}

// IsApril reports whether it is April.
func (c *Carbon) IsApril() bool {
	return c.isMonth(time.April)
}

// IsMay reports whether it is May.
func (c *Carbon) IsMay() bool {
	return c.isMonth(time.May)
}

// IsJune reports whether it is June.
func (c *Carbon) IsJune() bool {
	return c.isMonth(time.June)
}

// IsJuly reports whether it is July.
func (c *Carbon) IsJuly() bool {
	return c.isMonth(time.July)
}

// IsAugust reports whether it is August.
func (c *Carbon) IsAugust() bool {
	return c.isMonth(time.August)
}

// IsSeptember reports whether it is September.
func (c *Carbon) IsSeptember() bool {
	return c.isMonth(time.September)
}

// IsOctober reports whether it is October.
func (c *Carbon) IsOctober() bool {
	return c.isMonth(time.October)
}

// IsNovember reports whether it is November.
func (c *Carbon) IsNovember() bool {
	return c.isMonth(time.November)
}

// IsDecember reports whether it is December.
func (c *Carbon) IsDecember() bool {
	return c.isMonth(time.December)
}

// IsMonday reports whether it is Monday.
func (c *Carbon) IsMonday() bool {
	return c.isWeekday(time.Monday)
}

// IsTuesday reports whether it is Tuesday.
func (c *Carbon) IsTuesday() bool {
	return c.isWeekday(time.Tuesday)
}

// IsWednesday reports whether it is Wednesday.
func (c *Carbon) IsWednesday() bool {
	return c.isWeekday(time.Wednesday)
}

// IsThursday reports whether it is Thursday.
func (c *Carbon) IsThursday() bool {
	return c.isWeekday(time.Thursday)
}

// IsFriday reports whether it is Friday.
func (c *Carbon) IsFriday() bool {
	return c.isWeekday(time.Friday)
}

// IsSaturday reports whether it is Saturday.
func (c *Carbon) IsSaturday() bool {
	return c.isWeekday(time.Saturday)
}

// IsSunday reports whether it is Sunday.
func (c *Carbon) IsSunday() bool {
	return c.isWeekday(time.Sunday)
}

// IsWeekday reports whether it is weekday.
func (c *Carbon) IsWeekday() bool {
	if c.IsInvalid() {
		return false
	}
	return !c.IsWeekend()
}

// IsWeekend reports whether it is weekend.
func (c *Carbon) IsWeekend() bool {
	if c.IsInvalid() {
		return false
	}
	d := c.StdTime().Weekday()
	for _, wd := range c.weekendDays {
		if d == wd {
			return true
		}
	}
	return false
}

// IsNow reports whether it is now time.
func (c *Carbon) IsNow() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Timestamp() == Now().SetLocation(c.loc).Timestamp()
}

// IsFuture reports whether it is future time.
func (c *Carbon) IsFuture() bool {
	if c.IsInvalid() {
		return false
	}
	if c.IsZero() {
		return false
	}
	return c.Timestamp() > Now().SetLocation(c.loc).Timestamp()
}

// IsPast reports whether it is past time.
func (c *Carbon) IsPast() bool {
	if c.IsInvalid() {
		return false
	}
	if c.IsZero() {
		return true
	}
	return c.Timestamp() < Now().SetLocation(c.loc).Timestamp()
}

// IsYesterday reports whether it is yesterday.
func (c *Carbon) IsYesterday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.ToDateString() == Yesterday().SetLocation(c.loc).ToDateString()
}

// IsToday reports whether it is today.
func (c *Carbon) IsToday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.ToDateString() == Now().SetLocation(c.loc).ToDateString()
}

// IsTomorrow reports whether it is tomorrow.
func (c *Carbon) IsTomorrow() bool {
	if c.IsInvalid() {
		return false
	}
	return c.ToDateString() == Tomorrow().SetLocation(c.loc).ToDateString()
}

// IsSameCentury reports whether it is same century.
func (c *Carbon) IsSameCentury(t *Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Century() == t.Century()
}

// IsSameDecade reports whether it is same decade.
func (c *Carbon) IsSameDecade(t *Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Decade() == t.Decade()
}

// IsSameYear reports whether it is same year.
func (c *Carbon) IsSameYear(t *Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Year() == t.Year()
}

// IsSameQuarter reports whether it is same quarter.
func (c *Carbon) IsSameQuarter(t *Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Year() == t.Year() && c.Quarter() == t.Quarter()
}

// IsSameMonth reports whether it is same month.
func (c *Carbon) IsSameMonth(t *Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	cTime := c.StdTime()
	tTime := t.StdTime()
	return cTime.Year() == tTime.Year() && cTime.Month() == tTime.Month()
}

// IsSameDay reports whether it is same day.
func (c *Carbon) IsSameDay(t *Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	cTime := c.StdTime()
	tTime := t.StdTime()
	return cTime.Year() == tTime.Year() &&
		cTime.Month() == tTime.Month() &&
		cTime.Day() == tTime.Day()
}

// IsSameHour reports whether it is same hour.
func (c *Carbon) IsSameHour(t *Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	cTime := c.StdTime()
	tTime := t.StdTime()
	return cTime.Year() == tTime.Year() &&
		cTime.Month() == tTime.Month() &&
		cTime.Day() == tTime.Day() &&
		cTime.Hour() == tTime.Hour()
}

// IsSameMinute reports whether it is same minute.
func (c *Carbon) IsSameMinute(t *Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	cTime := c.StdTime()
	tTime := t.StdTime()
	return cTime.Year() == tTime.Year() &&
		cTime.Month() == tTime.Month() &&
		cTime.Day() == tTime.Day() &&
		cTime.Hour() == tTime.Hour() &&
		cTime.Minute() == tTime.Minute()
}

// IsSameSecond reports whether it is same second.
func (c *Carbon) IsSameSecond(t *Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	cTime := c.StdTime()
	tTime := t.StdTime()
	return cTime.Year() == tTime.Year() &&
		cTime.Month() == tTime.Month() &&
		cTime.Day() == tTime.Day() &&
		cTime.Hour() == tTime.Hour() &&
		cTime.Minute() == tTime.Minute() &&
		cTime.Second() == tTime.Second()
}

// Compare compares by an operator.
func (c *Carbon) Compare(operator string, t *Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	switch operator {
	case "=":
		return c.Eq(t)
	case "<>", "!=":
		return !c.Eq(t)
	case ">":
		return c.Gt(t)
	case ">=":
		return c.Gte(t)
	case "<":
		return c.Lt(t)
	case "<=":
		return c.Lte(t)
	}
	return false
}

// Gt reports whether greater than.
func (c *Carbon) Gt(t *Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.time.After(t.time)
}

// Lt reports whether less than.
func (c *Carbon) Lt(t *Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.time.Before(t.time)
}

// Eq reports whether equal.
func (c *Carbon) Eq(t *Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.time.Equal(t.time)
}

// Ne reports whether not equal.
func (c *Carbon) Ne(t *Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return !c.Eq(t)
}

// Gte reports whether greater than or equal.
func (c *Carbon) Gte(t *Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Gt(t) || c.Eq(t)
}

// Lte reports whether less than or equal.
func (c *Carbon) Lte(t *Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Lt(t) || c.Eq(t)
}

// Between reports whether between two times, including the start and end time.
func (c *Carbon) Between(start *Carbon, end *Carbon) bool {
	if start.Gt(end) {
		return false
	}
	if c.IsInvalid() || start.IsInvalid() || end.IsInvalid() {
		return false
	}
	if c.Gt(start) && c.Lt(end) {
		return true
	}
	return false
}

// BetweenIncludedStart reports whether between two times, including the start time.
func (c *Carbon) BetweenIncludedStart(start *Carbon, end *Carbon) bool {
	if start.Gt(end) {
		return false
	}
	if c.IsZero() && start.IsZero() {
		return true
	}
	if c.IsInvalid() || start.IsInvalid() || end.IsInvalid() {
		return false
	}
	if c.Gte(start) && c.Lt(end) {
		return true
	}
	return false
}

// BetweenIncludedEnd reports whether between two times, including the end time.
func (c *Carbon) BetweenIncludedEnd(start *Carbon, end *Carbon) bool {
	if start.Gt(end) {
		return false
	}
	if c.IsZero() && end.IsZero() {
		return true
	}
	if c.IsInvalid() || start.IsInvalid() || end.IsInvalid() {
		return false
	}
	if c.Gt(start) && c.Lte(end) {
		return true
	}
	return false
}

// BetweenIncludedBoth reports whether between two times, including the start and end time.
func (c *Carbon) BetweenIncludedBoth(start *Carbon, end *Carbon) bool {
	if start.Gt(end) {
		return false
	}
	if (c.IsZero() && start.IsZero()) || (c.IsZero() && end.IsZero()) {
		return true
	}
	if c.IsInvalid() || start.IsInvalid() || end.IsInvalid() {
		return false
	}
	if c.Gte(start) && c.Lte(end) {
		return true
	}
	return false
}

// isMonth reports whether the current month matches the given month.
// It returns false if the Carbon instance is invalid.
func (c *Carbon) isMonth(month time.Month) bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Month() == month
}

// isWeekday reports whether the current weekday matches the given weekday.
// It returns false if the Carbon instance is invalid.
func (c *Carbon) isWeekday(weekday time.Weekday) bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Weekday() == weekday
}
