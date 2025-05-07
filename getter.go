package carbon

import (
	"time"
)

// StdTime gets standard time.Time.
// 获取标准 time.Time
func (c Carbon) StdTime() time.Time {
	if c.loc == nil {
		return c.time
	}
	return c.time.In(c.loc)
}

// DaysInYear gets total days in year like 365.
// 获取本年的总天数
func (c Carbon) DaysInYear() int {
	if c.IsInvalid() {
		return 0
	}
	if c.IsLeapYear() {
		return DaysPerLeapYear
	}
	return DaysPerNormalYear
}

// DaysInMonth gets total days in month like 30.
// 获取本月的总天数
func (c Carbon) DaysInMonth() int {
	if c.IsInvalid() {
		return 0
	}
	return c.EndOfMonth().StdTime().Day()
}

// MonthOfYear gets month of year like 12.
// 获取本年的第几月
func (c Carbon) MonthOfYear() int {
	if c.IsInvalid() {
		return 0
	}
	return int(c.StdTime().Month())
}

// DayOfYear gets day of year like 365.
// 获取本年的第几天
func (c Carbon) DayOfYear() int {
	if c.IsInvalid() {
		return 0
	}
	return c.StdTime().YearDay()
}

// DayOfMonth gets day of month like 30.
// 获取本月的第几天
func (c Carbon) DayOfMonth() int {
	if c.IsInvalid() {
		return 0
	}
	return c.StdTime().Day()
}

// DayOfWeek gets day of week like 6.
// 获取本周的第几天
func (c Carbon) DayOfWeek() int {
	if c.IsInvalid() {
		return 0
	}
	return (int(c.StdTime().Weekday())+DaysPerWeek-int(c.weekStartsAt))%DaysPerWeek + 1
}

// WeekOfYear gets week of year like 1, see https://en.wikipedia.org/wiki/ISO_8601#Week_dates.
// 获取本年的第几周
func (c Carbon) WeekOfYear() int {
	if c.IsInvalid() {
		return 0
	}
	_, week := c.StdTime().ISOWeek()
	return week
}

// WeekOfMonth gets week of month like 1.
// 获取本月的第几周
func (c Carbon) WeekOfMonth() int {
	if c.IsInvalid() {
		return 0
	}
	days := c.Day() + c.StartOfMonth().DayOfWeek() - 1
	if days%DaysPerWeek == 0 {
		return days / DaysPerWeek
	}
	return days/DaysPerWeek + 1
}

// DateTime gets current year, month, day, hour, minute, and second like 2020, 8, 5, 13, 14, 15.
// 获取当前年、月、日、时、分、秒
func (c Carbon) DateTime() (year, month, day, hour, minute, second int) {
	if c.IsInvalid() {
		return
	}
	year, month, day = c.Date()
	hour, minute, second = c.Time()
	return year, month, day, hour, minute, second
}

// DateTimeMilli gets current year, month, day, hour, minute, second and millisecond like 2020, 8, 5, 13, 14, 15, 999.
// 获取当前年、月、日、时、分、秒、毫秒
func (c Carbon) DateTimeMilli() (year, month, day, hour, minute, second, millisecond int) {
	if c.IsInvalid() {
		return
	}
	year, month, day, hour, minute, second = c.DateTime()
	return year, month, day, hour, minute, second, c.Millisecond()
}

// DateTimeMicro gets current year, month, day, hour, minute, second and microsecond like 2020, 8, 5, 13, 14, 15, 999999.
// 获取当前年、月、日、时、分、秒、微秒
func (c Carbon) DateTimeMicro() (year, month, day, hour, minute, second, microsecond int) {
	if c.IsInvalid() {
		return
	}
	year, month, day, hour, minute, second = c.DateTime()
	return year, month, day, hour, minute, second, c.Microsecond()
}

// DateTimeNano gets current year, month, day, hour, minute, second and nanosecond like 2020, 8, 5, 13, 14, 15, 999999999.
// 获取当前年、月、日、时、分、秒、纳秒
func (c Carbon) DateTimeNano() (year, month, day, hour, minute, second, nanosecond int) {
	if c.IsInvalid() {
		return
	}
	year, month, day, hour, minute, second = c.DateTime()
	return year, month, day, hour, minute, second, c.Nanosecond()
}

// Date gets current year, month, and day like 2020, 8, 5.
// 获取当前年、月、日
func (c Carbon) Date() (year, month, day int) {
	if c.IsInvalid() {
		return
	}
	var tm time.Month
	year, tm, day = c.StdTime().Date()
	return year, int(tm), day
}

// DateMilli gets current year, month, day and millisecond like 2020, 8, 5, 999.
// 获取当前年、月、日、毫秒
func (c Carbon) DateMilli() (year, month, day, millisecond int) {
	if c.IsInvalid() {
		return
	}
	year, month, day = c.Date()
	return year, month, day, c.Millisecond()
}

// DateMicro gets current year, month, day and microsecond like 2020, 8, 5, 999999.
// 获取当前年、月、日、微秒
func (c Carbon) DateMicro() (year, month, day, microsecond int) {
	if c.IsInvalid() {
		return
	}
	year, month, day = c.Date()
	return year, month, day, c.Microsecond()
}

// DateNano gets current year, month, day and nanosecond like 2020, 8, 5, 999999999.
// 获取当前年、月、日、纳秒
func (c Carbon) DateNano() (year, month, day, nanosecond int) {
	if c.IsInvalid() {
		return
	}
	year, month, day = c.Date()
	return year, month, day, c.Nanosecond()
}

// Time gets current hour, minute, and second like 13, 14, 15.
// 获取当前时、分、秒
func (c Carbon) Time() (hour, minute, second int) {
	if c.IsInvalid() {
		return
	}
	return c.StdTime().Clock()
}

// TimeMilli gets current hour, minute, second and millisecond like 13, 14, 15, 999.
// 获取当前时、分、秒、毫秒
func (c Carbon) TimeMilli() (hour, minute, second, millisecond int) {
	if c.IsInvalid() {
		return
	}
	hour, minute, second = c.Time()
	return hour, minute, second, c.Millisecond()
}

// TimeMicro gets current hour, minute, second and microsecond like 13, 14, 15, 999999.
// 获取当前时、分、秒、微秒
func (c Carbon) TimeMicro() (hour, minute, second, microsecond int) {
	if c.IsInvalid() {
		return
	}
	hour, minute, second = c.Time()
	return hour, minute, second, c.Microsecond()
}

// TimeNano gets current hour, minute, second and nanosecond like 13, 14, 15, 999999999.
// 获取当前时、分、秒、纳秒
func (c Carbon) TimeNano() (hour, minute, second, nanosecond int) {
	if c.IsInvalid() {
		return
	}
	hour, minute, second = c.Time()
	return hour, minute, second, c.Nanosecond()
}

// Century gets current century like 21.
// 获取当前世纪
func (c Carbon) Century() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Year()/YearsPerCentury + 1
}

// Decade gets current decade like 20.
// 获取当前年代
func (c Carbon) Decade() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Year() % YearsPerCentury / YearsPerDecade * YearsPerDecade
}

// Year gets current year like 2020.
// 获取当前年
func (c Carbon) Year() int {
	if c.IsInvalid() {
		return 0
	}
	return c.StdTime().Year()
}

// Quarter gets current quarter like 3.
// 获取当前季度
func (c Carbon) Quarter() (quarter int) {
	if c.IsInvalid() {
		return
	}
	month := c.Month()
	switch {
	case month >= 10:
		quarter = 4
	case month >= 7:
		quarter = 3
	case month >= 4:
		quarter = 2
	case month >= 1:
		quarter = 1
	}
	return
}

// Month gets current month like 8.
// 获取当前月
func (c Carbon) Month() int {
	return c.MonthOfYear()
}

// Week gets current week like 6, start from 0.
// 获取当前周(从0开始)
func (c Carbon) Week() int {
	if c.IsInvalid() {
		return -1
	}
	return c.DayOfWeek() - 1
}

// Day gets current day like 5.
// 获取当前日
func (c Carbon) Day() int {
	return c.DayOfMonth()
}

// Hour gets current hour like 13.
// 获取当前小时
func (c Carbon) Hour() int {
	if c.IsInvalid() {
		return 0
	}
	return c.StdTime().Hour()
}

// Minute gets current minute like 14.
// 获取当前分钟数
func (c Carbon) Minute() int {
	if c.IsInvalid() {
		return 0
	}
	return c.StdTime().Minute()
}

// Second gets current second like 9.
// 获取当前秒数
func (c Carbon) Second() int {
	if c.IsInvalid() {
		return 0
	}
	return c.StdTime().Second()
}

// Millisecond gets current millisecond like 999.
// 获取当前毫秒数
func (c Carbon) Millisecond() int {
	if c.IsInvalid() {
		return 0
	}
	return c.StdTime().Nanosecond() / 1e6
}

// Microsecond gets current microsecond like 999999.
// 获取当前微秒数
func (c Carbon) Microsecond() int {
	if c.IsInvalid() {
		return 0
	}
	return c.StdTime().Nanosecond() / 1e3
}

// Nanosecond gets current nanosecond like 999999999.
// 获取当前纳秒数
func (c Carbon) Nanosecond() int {
	if c.IsInvalid() {
		return 0
	}
	return c.StdTime().Nanosecond()
}

// Timestamp gets timestamp with second precision like 1596604455.
// 输出秒精度时间戳
func (c Carbon) Timestamp() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.StdTime().Unix()
}

// TimestampMilli gets timestamp with millisecond precision like 1596604455000.
// 获取毫秒精度时间戳
func (c Carbon) TimestampMilli() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.StdTime().UnixMilli()
}

// TimestampMicro gets timestamp with microsecond precision like 1596604455000000.
// 获取微秒精度时间戳
func (c Carbon) TimestampMicro() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.StdTime().UnixMicro()
}

// TimestampNano gets timestamp with nanosecond precision like 1596604455000000000.
// 获取纳秒精度时间戳
func (c Carbon) TimestampNano() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.StdTime().UnixNano()
}

// Timezone gets timezone location like "Asia/Shanghai".
// 获取时区位置
func (c Carbon) Timezone() string {
	if c.IsInvalid() {
		return ""
	}
	return c.loc.String()
}

// ZoneName gets timezone name like "CST".
// 获取时区名称
func (c Carbon) ZoneName() string {
	if c.IsInvalid() {
		return ""
	}
	name, _ := c.StdTime().Zone()
	return name
}

// ZoneOffset gets timezone offset seconds from the UTC timezone like 28800.
// 获取时区偏移量，单位秒
func (c Carbon) ZoneOffset() int {
	if c.IsInvalid() {
		return 0
	}
	_, offset := c.StdTime().Zone()
	return offset
}

// Locale gets locale name like "zh-CN".
// 获取语言区域
func (c Carbon) Locale() string {
	if c.IsInvalid() {
		return ""
	}
	return c.lang.locale
}

// WeekStartsAt returns start day of the week.
// 获取一周的开始日期
func (c Carbon) WeekStartsAt() Weekday {
	if c.IsInvalid() {
		return 0
	}
	return c.weekStartsAt
}

// WeekEndsAt returns end day of the week.
// 获取一周的结束日期
func (c Carbon) WeekEndsAt() Weekday {
	if c.IsInvalid() {
		return 0
	}
	return Weekday((int(c.weekStartsAt) + DaysPerWeek - 1) % 7)
}

// CurrentLayout returns the layout used for parsing the time string.
// 获取当前布局模板
func (c Carbon) CurrentLayout() string {
	if c.IsInvalid() {
		return ""
	}
	return c.currentLayout
}

// Age gets age like 18.
// 获取年龄
func (c Carbon) Age() int {
	if c.IsInvalid() {
		return 0
	}
	now := Now().SetLocation(c.loc)
	if c.Gte(now) {
		return 0
	}
	return int(c.DiffInYears(now))
}
