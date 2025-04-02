// Package lunar is part of the carbon package.
package lunar

import (
	"fmt"
	"strings"
	"time"

	"github.com/dromara/carbon/v2/calendar"
)

var (
	numbers = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	months  = []string{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "腊"}
	weeks   = []string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}
	animals = []string{"猴", "鸡", "狗", "猪", "鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊"}

	festivals = map[string]string{
		// "month-day": "name"
		"1-1":   "春节",
		"1-15":  "元宵节",
		"2-2":   "龙抬头",
		"3-3":   "上巳节",
		"5-5":   "端午节",
		"7-7":   "七夕节",
		"7-15":  "中元节",
		"8-15":  "中秋节",
		"9-9":   "重阳节",
		"10-1":  "寒衣节",
		"10-15": "下元节",
		"12-8":  "腊八节",
	}

	years = []int{
		0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2, // 1900-1909
		0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977, // 1910-1919
		0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970, // 1920-1929
		0x06566, 0x0d4a0, 0x0ea50, 0x16a95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950, // 1930-1939
		0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4, 0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557, // 1940-1949
		0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5d0, 0x14573, 0x052d0, 0x0a9a8, 0x0e950, 0x06aa0, // 1950-1959
		0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0, // 1960-1969
		0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b5a0, 0x195a6, // 1970-1979
		0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570, // 1980-1989
		0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50, 0x06b58, 0x05ac0, 0x0ab60, 0x096d5, 0x092e0, // 1990-1999
		0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5, // 2000-2009
		0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9, 0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930, // 2010-2019
		0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530, // 2020-2029
		0x05aa0, 0x076a3, 0x096d0, 0x04bd7, 0x04ad0, 0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45, // 2030-2039
		0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0, // 2040-2049
		0x14b63, 0x09370, 0x049f8, 0x04970, 0x064b0, 0x168a6, 0x0ea50, 0x06b20, 0x1a6c4, 0x0aae0, // 2050-2059
		0x0a2e0, 0x0d2e3, 0x0c960, 0x0d557, 0x0d4a0, 0x0da50, 0x05d55, 0x056a0, 0x0a6d0, 0x055d4, // 2060-2069
		0x052d0, 0x0a9b8, 0x0a950, 0x0b4a0, 0x0b6a6, 0x0ad50, 0x055a0, 0x0aba4, 0x0a5b0, 0x052b0, // 2070-2079
		0x0b273, 0x06930, 0x07337, 0x06aa0, 0x0ad50, 0x14b55, 0x04b60, 0x0a570, 0x054e4, 0x0d160, // 2080-2089
		0x0e968, 0x0d520, 0x0daa0, 0x16aa6, 0x056d0, 0x04ae0, 0x0a9d4, 0x0a2d0, 0x0d150, 0x0f252, // 2090-2099
		0x0d520, // 2100
	}
)

// returns a invalid lunar date.
// 无效的农历日期错误
var invalidLunarError = func() error {
	return fmt.Errorf("invalid lunar date, please make sure the lunar date is valid")
}

// Lunar defines a Lunar struct.
// 定义 Lunar 结构体
type Lunar struct {
	year, month, day int
	isLeapMonth      bool
	Error            error
}

// NewLunar returns a new Lunar instance.
// 返回 Lunar 实例
func NewLunar(year, month, day int, isLeapMonth bool) (l Lunar) {
	l.year, l.month, l.day, l.isLeapMonth = year, month, day, isLeapMonth
	if !l.IsValid() {
		l.Error = invalidLunarError()
	}
	return l
}

// MaxValue returns a Lunar instance for the greatest supported date.
// 返回 Carbon 的最大值
func MaxValue() Lunar {
	return Lunar{
		year:  2100,
		month: 12,
		day:   31,
	}
}

// MinValue returns a Lunar instance for the lowest supported date.
// 返回 Lunar 的最小值
func MinValue() Lunar {
	return Lunar{
		year:  1900,
		month: 1,
		day:   1,
	}
}

// FromStdTime creates a Lunar instance from standard time.Time.
// 从标准 time.Time 创建 Lunar 实例
func FromStdTime(t time.Time) (l Lunar) {
	if t.IsZero() {
		return l
	}
	daysInYear, daysInMonth, leapMonth := 365, 30, 0
	maxYear, minYear := MaxValue().year, MinValue().year

	offset := int(t.Truncate(time.Hour).Sub(time.Date(minYear, 1, 31, 0, 0, 0, 0, t.Location())).Hours() / 24)
	for l.year = minYear; l.year <= maxYear && offset > 0; l.year++ {
		daysInYear = l.getDaysInYear()
		offset -= daysInYear
	}
	if offset < 0 {
		offset += daysInYear
		l.year--
	}
	leapMonth = l.LeapMonth()
	for l.month = 1; l.month <= 12 && offset > 0; l.month++ {
		if leapMonth > 0 && l.month == (leapMonth+1) && !l.isLeapMonth {
			l.month--
			l.isLeapMonth = true
			daysInMonth = l.getDaysInLeapMonth()
		} else {
			daysInMonth = l.getDaysInMonth()
		}
		offset -= daysInMonth
		if l.isLeapMonth && l.month == (leapMonth+1) {
			l.isLeapMonth = false
		}
	}
	// offset为0时，并且月份是闰月，要校正
	if offset == 0 && leapMonth > 0 && l.month == leapMonth+1 {
		if l.isLeapMonth {
			l.isLeapMonth = false
		} else {
			l.isLeapMonth = true
			l.month--
		}
	}
	// offset小于0时，也要校正
	if offset < 0 {
		offset += daysInMonth
		l.month--
	}
	l.day = offset + 1
	return l
}

// ToGregorian converts Lunar instance to Gregorian instance.
// 将 Lunar 实例转化为 Gregorian 实例
func (l Lunar) ToGregorian(timezone ...string) (g calendar.Gregorian) {
	if !l.IsValid() {
		return g
	}
	loc := time.UTC
	if len(timezone) > 0 {
		loc, g.Error = time.LoadLocation(timezone[0])
	}
	if g.Error != nil {
		return g
	}
	days := l.getDaysInMonth()
	offset := l.getOffsetInYear()
	offset += l.getOffsetInMonth()

	// 补充闰月的前一个月的时差
	if l.isLeapMonth {
		offset += days
	}
	// https://github.com/golang-module/carbon/issues/219
	ts := int64(offset+l.day)*86400 - int64(2206512000)
	g.Time = time.Unix(ts, 0).In(loc)
	return g
}

// Animal gets lunar animal name like "猴".
// 获取农历生肖
func (l Lunar) Animal() string {
	if !l.IsValid() {
		return ""
	}
	return animals[l.year%12]
}

// Festival gets lunar festival name like "春节".
// 获取农历节日
func (l Lunar) Festival() string {
	if !l.IsValid() {
		return ""
	}
	return festivals[fmt.Sprintf("%d-%d", l.month, l.day)]
}

// Year gets lunar year like 2020.
// 获取农历年份
func (l Lunar) Year() int {
	if !l.IsValid() {
		return 0
	}
	return l.year
}

// Month gets lunar month like 8.
// 获取农历月份
func (l Lunar) Month() int {
	if !l.IsValid() {
		return 0
	}
	return l.month
}

// Day gets lunar day like 5.
// 获取农历日，如 5
func (l Lunar) Day() int {
	if !l.IsValid() {
		return 0
	}
	return l.day
}

// LeapMonth gets lunar leap month like 2.
// 获取农历闰月月份，如 2
func (l Lunar) LeapMonth() int {
	if !l.IsValid() {
		return 0
	}
	minYear := MinValue().year
	return years[l.year-minYear] & 0xf
}

// String implements Stringer interface for Lunar.
// 实现 Stringer 接口
func (l Lunar) String() string {
	if !l.IsValid() {
		return ""
	}
	return fmt.Sprintf("%04d-%02d-%02d", l.year, l.month, l.day)
}

// ToYearString outputs a string in lunar year format like "二零二零".
// 获取农历年份字符串，如 "二零二零"
func (l Lunar) ToYearString() (year string) {
	if !l.IsValid() {
		return ""
	}
	year = fmt.Sprintf("%d", l.year)
	for index, number := range numbers {
		year = strings.Replace(year, fmt.Sprintf("%d", index), number, -1)
	}
	return year
}

// ToMonthString outputs a string in lunar month format like "正月".
// 获取农历月份字符串，如 "正月"
func (l Lunar) ToMonthString() (month string) {
	if !l.IsValid() {
		return ""
	}
	month = months[l.month-1] + "月"
	if l.IsLeapMonth() {
		return "闰" + month
	}
	return
}

// ToWeekString outputs a string in week layout like "周一".
// 输出完整农历星期字符串，如 "周一"
func (l Lunar) ToWeekString() (month string) {
	if !l.IsValid() {
		return ""
	}
	return weeks[l.ToGregorian().Time.Weekday()]
}

// ToDayString outputs a string in lunar day format like "廿一".
// 获取农历日字符串，如 "廿一"
func (l Lunar) ToDayString() (day string) {
	if !l.IsValid() {
		return ""
	}
	num := numbers[l.day%10]
	switch {
	case l.day == 30:
		day = "三十"
	case l.day > 20:
		day = "廿" + num
	case l.day == 20:
		day = "二十"
	case l.day > 10:
		day = "十" + num
	case l.day == 10:
		day = "初十"
	case l.day < 10:
		day = "初" + num
	}
	return
}

// ToDateString outputs a string in lunar date format like "二零二零年腊月初五".
// 获取农历日期字符串，如 "二零二零年腊月初五"
func (l Lunar) ToDateString() string {
	if !l.IsValid() {
		return ""
	}
	return l.ToYearString() + "年" + l.ToMonthString() + l.ToDayString()
}

// IsValid reports whether is a valid lunar date.
// 是否是有效的年份
func (l Lunar) IsValid() bool {
	if l.Error != nil {
		return false
	}
	if l.year >= MinValue().year && l.year <= MaxValue().year {
		return true
	}
	return false
}

// IsLeapYear reports whether is a lunar leap year.
// 是否是农历闰年
func (l Lunar) IsLeapYear() bool {
	if !l.IsValid() {
		return false
	}
	return l.LeapMonth() != 0
}

// IsLeapMonth reports whether is a lunar leap month.
// 是否是农历闰月
func (l Lunar) IsLeapMonth() bool {
	if !l.IsValid() {
		return false
	}
	return l.isLeapMonth
}

// IsRatYear reports whether is lunar year of Rat.
// 是否是鼠年
func (l Lunar) IsRatYear() bool {
	if !l.IsValid() {
		return false
	}
	if l.year%12 == 4 {
		return true
	}
	return false
}

// IsOxYear reports whether is lunar year of Ox.
// 是否是牛年
func (l Lunar) IsOxYear() bool {
	if !l.IsValid() {
		return false
	}
	if l.year%12 == 5 {
		return true
	}
	return false
}

// IsTigerYear reports whether is lunar year of Tiger.
// 是否是虎年
func (l Lunar) IsTigerYear() bool {
	if !l.IsValid() {
		return false
	}
	if l.year%12 == 6 {
		return true
	}
	return false
}

// IsRabbitYear reports whether is lunar year of Rabbit.
// 是否是兔年
func (l Lunar) IsRabbitYear() bool {
	if !l.IsValid() {
		return false
	}
	if l.year%12 == 7 {
		return true
	}
	return false
}

// IsDragonYear reports whether is lunar year of Dragon.
// 是否是龙年
func (l Lunar) IsDragonYear() bool {
	if !l.IsValid() {
		return false
	}
	if l.year%12 == 8 {
		return true
	}
	return false
}

// IsSnakeYear reports whether is lunar year of Snake.
// 是否是蛇年
func (l Lunar) IsSnakeYear() bool {
	if !l.IsValid() {
		return false
	}
	if l.year%12 == 9 {
		return true
	}
	return false
}

// IsHorseYear reports whether is lunar year of Horse.
// 是否是马年
func (l Lunar) IsHorseYear() bool {
	if !l.IsValid() {
		return false
	}
	if l.year%12 == 10 {
		return true
	}
	return false
}

// IsGoatYear reports whether is lunar year of Goat.
// 是否是羊年
func (l Lunar) IsGoatYear() bool {
	if !l.IsValid() {
		return false
	}
	if l.year%12 == 11 {
		return true
	}
	return false
}

// IsMonkeyYear reports whether is lunar year of Monkey.
// 是否是猴年
func (l Lunar) IsMonkeyYear() bool {
	if !l.IsValid() {
		return false
	}
	if l.year%12 == 0 {
		return true
	}
	return false
}

// IsRoosterYear reports whether is lunar year of Rooster.
// 是否是鸡年
func (l Lunar) IsRoosterYear() bool {
	if !l.IsValid() {
		return false
	}
	if l.year%12 == 1 {
		return true
	}
	return false
}

// IsDogYear reports whether is lunar year of Dog.
// 是否是狗年
func (l Lunar) IsDogYear() bool {
	if !l.IsValid() {
		return false
	}
	if l.year%12 == 2 {
		return true
	}
	return false
}

// IsPigYear reports whether is lunar year of Pig.
// 是否是猪年
func (l Lunar) IsPigYear() bool {
	if !l.IsValid() {
		return false
	}
	if l.year%12 == 3 {
		return true
	}
	return false
}

func (l Lunar) getOffsetInYear() int {
	flag := false
	clone, month, offset := l, 0, 0
	for month = 1; month < l.month; month++ {
		leapMonth := l.LeapMonth()
		if !flag {
			// 处理闰月
			if leapMonth <= month && leapMonth > 0 {
				offset += l.getDaysInLeapMonth()
				flag = true
			}
		}
		clone.month = month
		offset += clone.getDaysInMonth()
	}
	return offset
}

func (l Lunar) getOffsetInMonth() int {
	clone, year, offset := l, 0, 0
	for year = MinValue().year; year < l.year; year++ {
		clone.year = year
		offset += clone.getDaysInYear()
	}
	return offset
}

func (l Lunar) getDaysInYear() int {
	var days = 348
	for i := 0x8000; i > 0x8; i >>= 1 {
		if (years[l.year-MinValue().year] & i) != 0 {
			days++
		}
	}
	return days + l.getDaysInLeapMonth()
}

func (l Lunar) getDaysInMonth() int {
	if (years[l.year-MinValue().year] & (0x10000 >> uint(l.month))) != 0 {
		return 30
	}
	return 29
}

func (l Lunar) getDaysInLeapMonth() int {
	if l.LeapMonth() == 0 {
		return 0
	}
	if years[l.year-MinValue().year]&0x10000 != 0 {
		return 30
	}
	return 29
}
