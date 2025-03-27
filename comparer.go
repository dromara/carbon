package carbon

import (
	"time"
)

// HasError reports whether it has error.
// 是否有错误
func (c Carbon) HasError() bool {
	if c.IsNil() {
		return false
	}
	return c.Error != nil
}

// IsNil reports whether is a nil time.
// 是否是空时间
func (c Carbon) IsNil() bool {
	if c.isNil || &c.time == nil {
		return true
	}
	return false
}

// IsZero reports whether is zero time(0001-01-01 00:00:00 +0000 UTC).
// 是否是零值时间(0001-01-01 00:00:00 +0000 UTC)
func (c Carbon) IsZero() bool {
	if c.IsNil() || c.HasError() {
		return false
	}
	return c.time.IsZero()
}

// IsValid reports whether is valid time.
// 是否是有效时间
func (c Carbon) IsValid() bool {
	if !c.IsNil() && !c.HasError() {
		return true
	}
	return false
}

// IsInvalid reports whether is invalid time.
// 是否是无效时间
func (c Carbon) IsInvalid() bool {
	return !c.IsValid()
}

// IsDST reports whether is daylight saving time.
// 是否是夏令时
func (c Carbon) IsDST() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.IsDST()
}

// IsAM reports whether is before noon.
// 是否是上午
func (c Carbon) IsAM() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Format("a") == "am"
}

// IsPM reports whether is after noon.
// 是否是下午
func (c Carbon) IsPM() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Format("a") == "pm"
}

// IsNow reports whether is now time.
// 是否是当前时间
func (c Carbon) IsNow() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Timestamp() == c.Now().Timestamp()
}

// IsFuture reports whether is future time.
// 是否是未来时间
func (c Carbon) IsFuture() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Timestamp() > c.Now().Timestamp()
}

// IsPast reports whether is past time.
// 是否是过去时间
func (c Carbon) IsPast() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Timestamp() < c.Now().Timestamp()
}

// IsLeapYear reports whether is a leap year.
// 是否是闰年
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

// IsLongYear reports whether is a long year, see https://en.wikipedia.org/wiki/ISO_8601#Week_dates.
// 是否是长年
func (c Carbon) IsLongYear() bool {
	if c.IsInvalid() {
		return false
	}
	_, w := time.Date(c.Year(), 12, 31, 0, 0, 0, 0, c.loc).ISOWeek()
	return w == weeksPerLongYear
}

// IsJanuary reports whether is January.
// 是否是一月
func (c Carbon) IsJanuary() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.January)
}

// IsFebruary reports whether is February.
// 是否是二月
func (c Carbon) IsFebruary() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.February)
}

// IsMarch reports whether is March.
// 是否是三月
func (c Carbon) IsMarch() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.March)
}

// IsApril reports whether is April.
// 是否是四月
func (c Carbon) IsApril() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.April)
}

// IsMay reports whether is May.
// 是否是五月
func (c Carbon) IsMay() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.May)
}

// IsJune reports whether is June.
// 是否是六月
func (c Carbon) IsJune() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.June)
}

// IsJuly reports whether is July.
// 是否是七月
func (c Carbon) IsJuly() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.July)
}

// IsAugust reports whether is August.
// 是否是八月
func (c Carbon) IsAugust() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.August)
}

// IsSeptember reports whether is September.
// 是否是九月
func (c Carbon) IsSeptember() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.September)
}

// IsOctober reports whether is October.
// 是否是十月
func (c Carbon) IsOctober() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.October)
}

// IsNovember reports whether is November.
// 是否是十一月
func (c Carbon) IsNovember() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.November)
}

// IsDecember reports whether is December.
// 是否是十二月
func (c Carbon) IsDecember() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Month() == int(time.December)
}

// IsMonday reports whether is Monday.
// 是否是周一
func (c Carbon) IsMonday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Weekday() == time.Monday
}

// IsTuesday reports whether is Tuesday.
// 是否是周二
func (c Carbon) IsTuesday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Weekday() == time.Tuesday
}

// IsWednesday reports whether is Wednesday.
// 是否是周三
func (c Carbon) IsWednesday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Weekday() == time.Wednesday
}

// IsThursday reports whether is Thursday.
// 是否是周四
func (c Carbon) IsThursday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Weekday() == time.Thursday
}

// IsFriday reports whether is Friday.
// 是否是周五
func (c Carbon) IsFriday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Weekday() == time.Friday
}

// IsSaturday reports whether is Saturday.
// 是否是周六
func (c Carbon) IsSaturday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Weekday() == time.Saturday
}

// IsSunday reports whether is Sunday.
// 是否是周日
func (c Carbon) IsSunday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.StdTime().Weekday() == time.Sunday
}

// IsWeekday reports whether is weekday.
// 是否是工作日
func (c Carbon) IsWeekday() bool {
	if c.IsInvalid() {
		return false
	}
	return !c.IsSaturday() && !c.IsSunday()
}

// IsWeekend reports whether is weekend.
// 是否是周末
func (c Carbon) IsWeekend() bool {
	if c.IsInvalid() {
		return false
	}
	return c.IsSaturday() || c.IsSunday()
}

// IsYesterday reports whether is yesterday.
// 是否是昨天
func (c Carbon) IsYesterday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.ToDateString() == Yesterday().ToDateString()
}

// IsToday reports whether is today.
// 是否是今天
func (c Carbon) IsToday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.ToDateString() == Now().ToDateString()
}

// IsTomorrow reports whether is tomorrow.
// 是否是明天
func (c Carbon) IsTomorrow() bool {
	if c.IsInvalid() {
		return false
	}
	return c.ToDateString() == Tomorrow().ToDateString()
}

// IsSameCentury reports whether is same century.
// 是否是同一世纪
func (c Carbon) IsSameCentury(t Carbon) bool {
	if c.Error != nil || t.Error != nil {
		return false
	}
	return c.Century() == t.Century()
}

// IsSameDecade reports whether is same decade.
// 是否是同一年代
func (c Carbon) IsSameDecade(t Carbon) bool {
	if c.Error != nil || t.Error != nil {
		return false
	}
	return c.Decade() == t.Decade()
}

// IsSameYear reports whether is same year.
// 是否是同一年
func (c Carbon) IsSameYear(t Carbon) bool {
	if c.Error != nil || t.Error != nil {
		return false
	}
	return c.Year() == t.Year()
}

// IsSameQuarter reports whether is same quarter.
// 是否是同一季节
func (c Carbon) IsSameQuarter(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Year() == t.Year() && c.Quarter() == t.Quarter()
}

// IsSameMonth reports whether is same month.
// 是否是同一月
func (c Carbon) IsSameMonth(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Format("Ym") == t.Format("Ym")
}

// IsSameDay reports whether is same day.
// 是否是同一天
func (c Carbon) IsSameDay(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Format("Ymd") == t.Format("Ymd")
}

// IsSameHour reports whether is same hour.
// 是否是同一小时
func (c Carbon) IsSameHour(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Format("YmdH") == t.Format("YmdH")
}

// IsSameMinute reports whether is same minute.
// 是否是同一分钟
func (c Carbon) IsSameMinute(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Format("YmdHi") == t.Format("YmdHi")
}

// IsSameSecond reports whether is same second.
// 是否是同一秒
func (c Carbon) IsSameSecond(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Format("YmdHis") == t.Format("YmdHis")

}

// Compare compares by an operator.
// 时间比较
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
// 是否大于
func (c Carbon) Gt(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.time.After(t.time)
}

// Lt reports whether less than.
// 是否小于
func (c Carbon) Lt(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.time.Before(t.time)
}

// Eq reports whether equal.
// 是否等于
func (c Carbon) Eq(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.time.Equal(t.time)
}

// Ne reports whether not equal.
// 是否不等于
func (c Carbon) Ne(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return !c.Eq(t)
}

// Gte reports whether greater than or equal.
// 是否大于等于
func (c Carbon) Gte(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Gt(t) || c.Eq(t)
}

// Lte reports whether less than or equal.
// 是否小于等于
func (c Carbon) Lte(t Carbon) bool {
	if c.IsInvalid() || t.IsInvalid() {
		return false
	}
	return c.Lt(t) || c.Eq(t)
}

// Between reports whether between two times, excluded the start and end time.
// 是否在两个时间之间(不包括这两个时间)
func (c Carbon) Between(start Carbon, end Carbon) bool {
	if c.IsInvalid() || start.IsInvalid() || end.IsInvalid() {
		return false
	}
	if c.Gt(start) && c.Lt(end) {
		return true
	}
	return false
}

// BetweenIncludedStart reports whether between two times, included the start time.
// 是否在两个时间之间(包括开始时间)
func (c Carbon) BetweenIncludedStart(start Carbon, end Carbon) bool {
	if c.IsInvalid() || start.IsInvalid() || end.IsInvalid() {
		return false
	}
	if c.Gte(start) && c.Lt(end) {
		return true
	}
	return false
}

// BetweenIncludedEnd reports whether between two times, included the end time.
// 是否在两个时间之间(包括结束时间)
func (c Carbon) BetweenIncludedEnd(start Carbon, end Carbon) bool {
	if c.IsInvalid() || start.IsInvalid() || end.IsInvalid() {
		return false
	}
	if c.Gt(start) && c.Lte(end) {
		return true
	}
	return false
}

// BetweenIncludedBoth reports whether between two times, included the start and end time.
// 是否在两个时间之间(包括这两个时间)
func (c Carbon) BetweenIncludedBoth(start Carbon, end Carbon) bool {
	if c.IsInvalid() || start.IsInvalid() || end.IsInvalid() {
		return false
	}
	if c.Gte(start) && c.Lte(end) {
		return true
	}
	return false
}
