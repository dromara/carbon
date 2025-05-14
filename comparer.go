package carbon

import (
	"time"
)

// HasError reports whether it has error.
func (c Carbon) HasError() bool {
	if c.IsEmpty() {
		return false
	}
	return c.Error != nil
}

// IsEmpty reports whether is a empty time.
func (c Carbon) IsEmpty() bool {
	if c.isEmpty == true {
		return true
	}
	return false
}

// IsZero reports whether is a zero time(0001-01-01 00:00:00 +0000 UTC).
func (c Carbon) IsZero() bool {
	if c.IsEmpty() || c.HasError() {
		return false
	}
	return c.time.IsZero()
}

// IsEpoch reports whether is a unix epoch time(1970-01-01 00:00:00 +0000 UTC).
func (c Carbon) IsEpoch() bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Equal(time.Date(EpochYear, 1, 1, 0, 0, 0, 0, time.UTC))
}

// IsValid reports whether is a valid time.
func (c Carbon) IsValid() bool {
	if !c.IsEmpty() && !c.HasError() {
		return true
	}
	return false
}

// IsInvalid reports whether is an invalid time.
func (c Carbon) IsInvalid() bool {
	return !c.IsValid()
}

// IsDST reports whether is a daylight saving time.
func (c Carbon) IsDST() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.IsDST()
}

// IsAM reports whether is before noon.
func (c Carbon) IsAM() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Format("a") == "am"
}

// IsPM reports whether is after noon.
func (c Carbon) IsPM() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Format("a") == "pm"
}

// IsLeapYear reports whether is a leap year.
func (c Carbon) IsLeapYear() bool {
	if c.IsInvalid() {
		return false
	}
	year := c.Year()
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

// IsLongYear reports whether is a long year, refer to  https://en.wikipedia.org/wiki/ISO_8601#Week_dates.
func (c Carbon) IsLongYear() bool {
	if c.IsInvalid() {
		return false
	}
	_, w := time.Date(c.Year(), 12, 31, 0, 0, 0, 0, c.loc).ISOWeek()
	return w == WeeksPerLongYear
}

// IsJanuary reports whether is January.
func (c Carbon) IsJanuary() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.January)
}

// IsFebruary reports whether is February.
func (c Carbon) IsFebruary() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.February)
}

// IsMarch reports whether is March.
func (c Carbon) IsMarch() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.March)
}

// IsApril reports whether is April.
func (c Carbon) IsApril() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.April)
}

// IsMay reports whether is May.
func (c Carbon) IsMay() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.May)
}

// IsJune reports whether is June.
func (c Carbon) IsJune() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.June)
}

// IsJuly reports whether is July.
func (c Carbon) IsJuly() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.July)
}

// IsAugust reports whether is August.
func (c Carbon) IsAugust() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.August)
}

// IsSeptember reports whether is September.
func (c Carbon) IsSeptember() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.September)
}

// IsOctober reports whether is October.
func (c Carbon) IsOctober() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.October)
}

// IsNovember reports whether is November.
func (c Carbon) IsNovember() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.November)
}

// IsDecember reports whether is December.
func (c Carbon) IsDecember() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.December)
}

// IsMonday reports whether is Monday.
func (c Carbon) IsMonday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Weekday() == time.Monday
}

// IsTuesday reports whether is Tuesday.
func (c Carbon) IsTuesday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Weekday() == time.Tuesday
}

// IsWednesday reports whether is Wednesday.
func (c Carbon) IsWednesday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Weekday() == time.Wednesday
}

// IsThursday reports whether is Thursday.
func (c Carbon) IsThursday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Weekday() == time.Thursday
}

// IsFriday reports whether is Friday.
func (c Carbon) IsFriday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Weekday() == time.Friday
}

// IsSaturday reports whether is Saturday.
func (c Carbon) IsSaturday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Weekday() == time.Saturday
}

// IsSunday reports whether is Sunday.
func (c Carbon) IsSunday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Weekday() == time.Sunday
}

// IsWeekday reports whether is weekday.
func (c Carbon) IsWeekday() bool {
	if c.IsInvalid() {
		return false
	}
	return !c.IsWeekend()
}

// IsWeekend reports whether is weekend.
func (c Carbon) IsWeekend() bool {
	if c.IsInvalid() {
		return false
	}
	d := c.StdTime().Weekday()
	for i := range c.weekendDays {
		if d == c.weekendDays[i] {
			return true
		}
	}
	return false
}

// IsNow reports whether is now time.
func (c Carbon) IsNow() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Timestamp() == Now().SetLocation(c.loc).Timestamp()
}

// IsFuture reports whether is future time.
func (c Carbon) IsFuture() bool {
	if c.IsInvalid() {
		return false
	}
	if c.IsZero() {
		return false
	}
	return c.Timestamp() > Now().SetLocation(c.loc).Timestamp()
}

// IsPast reports whether is past time.
func (c Carbon) IsPast() bool {
	if c.IsInvalid() {
		return false
	}
	if c.IsZero() {
		return true
	}
	return c.Timestamp() < Now().SetLocation(c.loc).Timestamp()
}

// IsYesterday reports whether is yesterday.
func (c Carbon) IsYesterday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.ToDateString() == Yesterday().SetLocation(c.loc).ToDateString()
}

// IsToday reports whether is today.
func (c Carbon) IsToday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.ToDateString() == Now().SetLocation(c.loc).ToDateString()
}

// IsTomorrow reports whether is tomorrow.
func (c Carbon) IsTomorrow() bool {
	if c.IsInvalid() {
		return false
	}
	return c.ToDateString() == Tomorrow().SetLocation(c.loc).ToDateString()
}

// IsSameCentury reports whether is same century.
func (c Carbon) IsSameCentury(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Century() == t.Century()
}

// IsSameDecade reports whether is same decade.
func (c Carbon) IsSameDecade(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Decade() == t.Decade()
}

// IsSameYear reports whether is same year.
func (c Carbon) IsSameYear(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Year() == t.Year()
}

// IsSameQuarter reports whether is same quarter.
func (c Carbon) IsSameQuarter(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Year() == t.Year() && c.Quarter() == t.Quarter()
}

// IsSameMonth reports whether is same month.
func (c Carbon) IsSameMonth(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Format("Ym") == t.Format("Ym")
}

// IsSameDay reports whether is same day.
func (c Carbon) IsSameDay(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Format("Ymd") == t.Format("Ymd")
}

// IsSameHour reports whether is same hour.
func (c Carbon) IsSameHour(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Format("YmdH") == t.Format("YmdH")
}

// IsSameMinute reports whether is same minute.
func (c Carbon) IsSameMinute(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Format("YmdHi") == t.Format("YmdHi")
}

// IsSameSecond reports whether is same second.
func (c Carbon) IsSameSecond(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Format("YmdHis") == t.Format("YmdHis")

}

// Compare compares by an operator.
func (c Carbon) Compare(operator string, t Carbon) bool {
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
func (c Carbon) Gt(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.time.After(t.time)
}

// Lt reports whether less than.
func (c Carbon) Lt(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.time.Before(t.time)
}

// Eq reports whether equal.
func (c Carbon) Eq(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.time.Equal(t.time)
}

// Ne reports whether not equal.
func (c Carbon) Ne(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return !c.Eq(t)
}

// Gte reports whether greater than or equal.
func (c Carbon) Gte(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Gt(t) || c.Eq(t)
}

// Lte reports whether less than or equal.
func (c Carbon) Lte(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Lt(t) || c.Eq(t)
}

// Between reports whether between two times, excluded the start and end time.
func (c Carbon) Between(start Carbon, end Carbon) bool {
	if c.IsInvalid() || start.IsInvalid() || end.IsInvalid() {
		return false
	}
	if start.Gt(end) {
		return false
	}
	if c.Gt(start) && c.Lt(end) {
		return true
	}
	return false
}

// BetweenIncludedStart reports whether between two times, included the start time.
func (c Carbon) BetweenIncludedStart(start Carbon, end Carbon) bool {
	if c.IsInvalid() || start.IsInvalid() || end.IsInvalid() {
		return false
	}
	if start.Gt(end) {
		return false
	}
	if c.Gte(start) && c.Lt(end) {
		return true
	}
	return false
}

// BetweenIncludedEnd reports whether between two times, included the end time.
func (c Carbon) BetweenIncludedEnd(start Carbon, end Carbon) bool {
	if c.IsInvalid() || start.IsInvalid() || end.IsInvalid() {
		return false
	}
	if start.Gt(end) {
		return false
	}
	if c.Gt(start) && c.Lte(end) {
		return true
	}
	return false
}

// BetweenIncludedBoth reports whether between two times, included the start and end time.
func (c Carbon) BetweenIncludedBoth(start Carbon, end Carbon) bool {
	if c.IsInvalid() || start.IsInvalid() || end.IsInvalid() {
		return false
	}
	if start.Gt(end) {
		return false
	}
	if c.Gte(start) && c.Lte(end) {
		return true
	}
	return false
}
