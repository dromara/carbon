package carbon

import (
	"strings"
	"time"
)

type Carbon struct {
	Time time.Time
	loc  *time.Location
}

// format转layout
func format2layout(format string) string {
	format = strings.Trim(format, " ")
	layout := strings.Replace(format, "Y", "2006", 1)
	layout = strings.Replace(layout, "y", "06", 1)
	layout = strings.Replace(layout, "m", "01", 1)
	layout = strings.Replace(layout, "n", "1", 1)
	layout = strings.Replace(layout, "d", "02", 1)
	layout = strings.Replace(layout, "j", "2", 1)
	layout = strings.Replace(layout, "H", "15", 1)
	layout = strings.Replace(layout, "h", "03", 1)
	layout = strings.Replace(layout, "g", "3", 1)
	layout = strings.Replace(layout, "i", "04", 1)
	layout = strings.Replace(layout, "s", "05", 1)
	return layout
}

// Now 当前
func Now() *Carbon {
	c := New()
	c.Time = time.Now()
	return c
}

// CreateFromTimestamp 时间戳转时间对象
func CreateFromTimestamp(timestamp int64) *Carbon {
	c := New()
	c.Time = time.Unix(timestamp, 0)
	return c
}

// CreateFromDateTime 年月日时分秒转时间对象
func CreateFromDateTime(year, month, day, hour, minute, second int) *Carbon {
	c := New()
	c.Time = time.Date(year, time.Month(month), day, hour, minute, second, 0, c.loc)
	return c
}

// CreateFromDate 年月日转时间对象
func CreateFromDate(year, month, day int) *Carbon {
	c := New()
	c.Time = time.Date(year, time.Month(month), day, 0, 0, 0, 0, c.loc)
	return c
}

// CreateFromTime 时分秒转时间对象
func CreateFromTime(hour, minute, second int) *Carbon {
	c := Now()
	return c.setTime(hour, minute, second, 0, c.loc)
}

// Today 今天
func Today() *Carbon {
	return Now().StartOfDay()
}

// Tomorrow 明天
func Tomorrow() *Carbon {
	return Now().AddDay().StartOfDay()
}

// Yesterday 昨天
func Yesterday() *Carbon {
	return Now().SubDay().StartOfDay()
}

// New 初始化,设置默认时区
func New() *Carbon {
	loc, _ := time.LoadLocation(Local)
	return &Carbon{loc: loc}
}

// Timezone 设置时区
func (c *Carbon) Timezone(name string) *Carbon {
	loc, _ := time.LoadLocation(name)
	c.loc = loc
	return c
}

// CreateFromTimestamp 时间戳转时间对象
func (c *Carbon) CreateFromTimestamp(timestamp int64) *Carbon {
	c.Time = time.Unix(timestamp, 0)
	return c
}

// CreateFromDateTime 年月日时分秒转时间对象
func (c *Carbon) CreateFromDateTime(year int, month time.Month, day int, hour int, minute int, second int) *Carbon {
	c.Time = time.Date(year, month, day, hour, minute, second, 0, c.loc)
	return c
}

// CreateFromDate 年月日转时间对象
func (c *Carbon) CreateFromDate(year int, month time.Month, day int) *Carbon {
	c.Time = time.Date(year, month, day, 0, 0, 0, 0, c.loc)
	return c
}

// CreateFromTime 时分秒转时间对象
func (c *Carbon) CreateFromTime(hour int, minute int, second int) *Carbon {
	c.Time = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hour, minute, second, 0, c.loc)
	return c
}

// Now 当前
func (c *Carbon) Now() *Carbon {
	c.Time = time.Now()
	return c
}

// AddYears N年后
func (c *Carbon) AddYears(years int) *Carbon {
	c.Time = c.Time.AddDate(years, 0, 0)
	return c
}

// AddYear 1年后
func (c *Carbon) AddYear() *Carbon {
	c.Time = c.Time.AddDate(1, 0, 0)
	return c
}

// SubYears N年前
func (c *Carbon) SubYears(years int) *Carbon {
	c.Time = c.Time.AddDate(-years, 0, 0)
	return c
}

// SubYear 1年前
func (c *Carbon) SubYear() *Carbon {
	c.Time = c.Time.AddDate(-1, 0, 0)
	return c
}

// AddMonths N月后
func (c *Carbon) AddMonths(months int) *Carbon {
	c.Time = c.Time.AddDate(0, months, 0)
	return c
}

// AddMonth 1月后
func (c *Carbon) AddMonth() *Carbon {
	c.Time = c.Time.AddDate(0, 1, 0)
	return c
}

// SubMonths N月前
func (c *Carbon) SubMonths(months int) *Carbon {
	c.Time = c.Time.AddDate(0, -months, 0)
	return c
}

// SubMonth 1月前
func (c *Carbon) SubMonth() *Carbon {
	c.Time = c.Time.AddDate(0, -1, 0)
	return c
}

// AddWeeks N周后
func (c *Carbon) AddWeeks(weeks int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, weeks*7)
	return c
}

// AddWeek 1周后
func (c *Carbon) AddWeek() *Carbon {
	c.Time = c.Time.AddDate(0, 0, 7)
	return c
}

// SubWeeks N周前
func (c *Carbon) SubWeeks(weeks int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, -weeks*7)
	return c
}

// SubMonth 1月前
func (c *Carbon) SubWeek() *Carbon {
	c.Time = c.Time.AddDate(0, 0, -7)
	return c
}

// AddDays N天后
func (c *Carbon) AddDays(days int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, days)
	return c
}

// AddDay 1天后
func (c *Carbon) AddDay() *Carbon {
	c.Time = c.Time.AddDate(0, 0, 1)
	return c
}

// SubDays N天前
func (c *Carbon) SubDays(days int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, -days)
	return c
}

// SubDay 1天前
func (c *Carbon) SubDay() *Carbon {
	c.Time = c.Time.AddDate(0, 0, -1)
	return c
}

// AddHours N小时后
func (c *Carbon) AddHours(hours int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, hours/24)
	return c
}

// AddHour 1小时后
func (c *Carbon) AddHour() *Carbon {
	c.Time = c.Time.AddDate(0, 0, 1/24)
	return c
}

// AddHours N小时前
func (c *Carbon) SubHours(hours int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, -hours/24)
	return c
}

// AddHour 1小时前
func (c *Carbon) SubHour() *Carbon {
	c.Time = c.Time.AddDate(0, 0, -1/24)
	return c
}

// AddMinutes N分钟后
func (c *Carbon) AddMinutes(minutes int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, minutes/1440)
	return c
}

// AddMinute 1分钟后
func (c *Carbon) AddMinute() *Carbon {
	c.Time = c.Time.AddDate(0, 0, 1/1440)
	return c
}

// SubMinutes N分钟前
func (c *Carbon) SubMinutes(minutes int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, -minutes/1440)
	return c
}

// SubMinute 1分钟前
func (c *Carbon) SubMinute() *Carbon {
	c.Time = c.Time.AddDate(0, 0, -1/1440)
	return c
}

// AddSeconds N秒钟后
func (c *Carbon) AddSeconds(second int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, second/86400)
	return c
}

// AddMinute 1秒钟后
func (c *Carbon) AddSecond() *Carbon {
	c.Time = c.Time.AddDate(0, 0, 1/86400)
	return c
}

// SubMinutes N秒钟前
func (c *Carbon) SubSeconds(second int) *Carbon {
	c.Time = c.Time.AddDate(0, 0, -second/86400)
	return c
}

// SubMinute 1秒钟前
func (c *Carbon) SubSecond() *Carbon {
	c.Time = c.Time.AddDate(0, 0, -1/86400)
	return c
}

// Parse 解析标准时间格式
func (c *Carbon) Parse(value string) *Carbon {
	value = strings.Trim(value, " ")
	layout := "2006-01-02 15:04:05"

	if len(value) == 10 && strings.Contains(value, "-") {
		layout = "2006-01-02"
	}

	if len(value) == 10 && strings.Contains(value, "/") {
		layout = "2006/01/02"
	}

	if len(value) == 19 && strings.Contains(value, "/") {
		layout = "2006/01/02 15:04:05"
	}

	if len(value) == 14 {
		layout = "20060102150405"
	}

	if len(value) == 8 {
		layout = "20060102"
	}

	t, _ := time.ParseInLocation(layout, value, c.loc)
	c.Time = t
	return c
}

// ParseByCustom 解析自定义时间格式
func (c *Carbon) ParseByCustom(value string, format string) *Carbon {
	value = strings.Trim(value, " ")
	layout := format2layout(format)
	t, _ := time.ParseInLocation(layout, value, c.loc)
	c.Time = t
	return c
}

// Format 格式化时间
func (c *Carbon) Format(format string) string {
	if c.Time.IsZero() {
		return ""
	}
	return c.Time.In(c.loc).Format(format2layout(format))
}

// Today 今天
func (c *Carbon) Today() string {
	return time.Now().In(c.loc).Format("2006-01-02 00:00:00")
}

// Tomorrow 明天
func (c *Carbon) Tomorrow() string {
	return time.Now().AddDate(0, 0, 1).In(c.loc).Format("2006-01-02 00:00:00")
}

// Yesterday 昨天
func (c *Carbon) Yesterday() string {
	return time.Now().AddDate(0, 0, -1).In(c.loc).Format("2006-01-02 00:00:00")
}

// FirstDayInYear 年初
func (c *Carbon) FirstDayInYear() string {
	return c.CreateFromDate(c.Time.Year(), 01, 01).ToDateTimeString()
}

// LastDayInYear 年末
func (c *Carbon) LastDayInYear() string {
	return c.CreateFromDate(c.Time.Year(), 12, 31).ToDateTimeString()
}

// FirstDayInMonth 月初
func (c *Carbon) FirstDayInMonth() string {
	return c.CreateFromDate(c.Time.Year(), c.Time.Month(), 1).ToDateTimeString()
}

// LastDayInMonth 月末
func (c *Carbon) LastDayInMonth() string {
	return c.CreateFromDate(c.Time.Year(), c.Time.Month(), 1).Time.AddDate(0, 1, -1).Format("2006-01-02 00:00:00")
}

// ToDateTimeString 转日期时间字符串
func (c *Carbon) ToDateTimeString() string {
	return c.Time.In(c.loc).Format("2006-01-02 15:04:05")
}

// ToDateString 转日期字符串
func (c *Carbon) ToDateString() string {
	return c.Time.In(c.loc).Format("2006-01-02")
}

// ToTimeString 转时间字符串
func (c *Carbon) ToTimeString() string {
	return c.Time.In(c.loc).Format("15:04:05")
}

// ToTimestamp 转时间戳
func (c *Carbon) ToTimestamp() int64 {
	return c.Time.Unix()
}

// IsLeapYear 是否是闰年
func (c *Carbon) IsLeapYear() bool {
	year := c.Time.Year()
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

// IsJanuary 是否是一月
func (c *Carbon) IsJanuary() bool {
	return c.Time.In(c.loc).Month() == time.January
}

// IsMonday 是否是二月
func (c *Carbon) IsFebruary() bool {
	return c.Time.In(c.loc).Month() == time.February
}

// IsMarch 是否是三月
func (c *Carbon) IsMarch() bool {
	return c.Time.In(c.loc).Month() == time.March
}

// IsApril 是否是四月
func (c *Carbon) IsApril() bool {
	return c.Time.In(c.loc).Month() == time.April
}

// IsMay 是否是五月
func (c *Carbon) IsMay() bool {
	return c.Time.In(c.loc).Month() == time.May
}

// IsJune 是否是六月
func (c *Carbon) IsJune() bool {
	return c.Time.In(c.loc).Month() == time.June
}

// IsJuly 是否是七月
func (c *Carbon) IsJuly() bool {
	return c.Time.In(c.loc).Month() == time.July
}

// IsAugust 是否是八月
func (c *Carbon) IsAugust() bool {
	return c.Time.In(c.loc).Month() == time.August
}

// IsSeptember 是否是九月
func (c *Carbon) IsSeptember() bool {
	return c.Time.In(c.loc).Month() == time.September
}

// IsOctober 是否是十月
func (c *Carbon) IsOctober() bool {
	return c.Time.In(c.loc).Month() == time.October
}

// IsNovember 是否是十一月
func (c *Carbon) IsNovember() bool {
	return c.Time.In(c.loc).Month() == time.November
}

// IsDecember 是否是十二月
func (c *Carbon) IsDecember() bool {
	return c.Time.In(c.loc).Month() == time.December
}

// IsMonday 是否是周一
func (c *Carbon) IsMonday() bool {
	return c.Time.In(c.loc).Weekday() == time.Monday
}

// IsTuesday 是否是周二
func (c *Carbon) IsTuesday() bool {
	return c.Time.In(c.loc).Weekday() == time.Tuesday
}

// IsWednesday 是否是周三
func (c *Carbon) IsWednesday() bool {
	return c.Time.In(c.loc).Weekday() == time.Wednesday
}

// IsThursday 是否是周四
func (c *Carbon) IsThursday() bool {
	return c.Time.In(c.loc).Weekday() == time.Thursday
}

// IsFriday 是否是周五
func (c *Carbon) IsFriday() bool {
	return c.Time.In(c.loc).Weekday() == time.Friday
}

// IsSaturday 是否是周六
func (c *Carbon) IsSaturday() bool {
	return c.Time.In(c.loc).Weekday() == time.Saturday
}

// IsSunday 是否是周日
func (c *Carbon) IsSunday() bool {
	return c.Time.In(c.loc).Weekday() == time.Sunday
}

// IsFirstDayInYear 是否年初
func (c *Carbon) IsFirstDayInYear() bool {
	_, month, day := c.Time.Date()
	if month == time.January && day == 1 {
		return true
	}

	return false
}

// IsLastDayInYear 是否是年末
func (c *Carbon) IsLastDayInYear() bool {
	_, month, day := c.Time.Date()
	if month == time.December && day == 31 {
		return true
	}

	return false
}

// IsFirstDayInMonth 是否月初
func (c *Carbon) IsFirstDayInMonth() bool {
	return c.Time.In(c.loc).Day() == 1
}

// IsLastDayInMonth 是否是月末
func (c *Carbon) IsLastDayInMonth() bool {
	return c.Time.In(c.loc).Format("2006-01-02 00:00:00") == c.LastDayInMonth()
}

// StartOfDay 当日开始时间
func (c *Carbon) StartOfDay() *Carbon {
	return c.setTime(0, 0, 0, 0, c.loc)
}

func (c *Carbon) EndOfDay() *Carbon {
	return c.setTime(HoursPerDay-1, MinutesPerHour-1, SecondsPerMinute-1, MicrosecondsPerSecond-1, c.loc)
}

func (c *Carbon) setTime(hour, min, sec, nsec int, loc *time.Location) *Carbon {
	year, month, day := c.Time.Date()
	c.Time = time.Date(year, month, day, hour, min, sec, nsec, loc)
	return c
}
