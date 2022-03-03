package carbon

import (
	"fmt"
	"strings"
)

var (
	minYear, maxYear   = 1900, 2100
	chineseNumbers     = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	chineseMonths      = []string{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "腊"}
	chineseDayPrefixes = []string{"初", "十", "廿", "卅"}
	animals            = []string{"猴", "鸡", "狗", "猪", "鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊"}
	festivals          = []string{"春节", "元宵节", "端午节", "七夕节", "中元节", "中秋节", "重阳节", "寒衣节", "下元节", "腊八节", "小年"}

	lunarTerms = []int{
		0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2, //1900-1909
		0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977, //1910-1919
		0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970, //1920-1929
		0x06566, 0x0d4a0, 0x0ea50, 0x16a95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950, //1930-1939
		0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4, 0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557, //1940-1949
		0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5d0, 0x14573, 0x052d0, 0x0a9a8, 0x0e950, 0x06aa0, //1950-1959
		0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0, //1960-1969
		0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b5a0, 0x195a6, //1970-1979
		0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570, //1980-1989
		0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50, 0x06b58, 0x05ac0, 0x0ab60, 0x096d5, 0x092e0, //1990-1999
		0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5, //2000-2009
		0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9, 0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930, //2010-2019
		0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530, //2020-2029
		0x05aa0, 0x076a3, 0x096d0, 0x04bd7, 0x04ad0, 0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45, //2030-2039
		0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0, //2040-2049
		0x14b63, 0x09370, 0x049f8, 0x04970, 0x064b0, 0x168a6, 0x0ea50, 0x06b20, 0x1a6c4, 0x0aae0, //2050-2059
		0x0a2e0, 0x0d2e3, 0x0c960, 0x0d557, 0x0d4a0, 0x0da50, 0x05d55, 0x056a0, 0x0a6d0, 0x055d4, //2060-2069
		0x052d0, 0x0a9b8, 0x0a950, 0x0b4a0, 0x0b6a6, 0x0ad50, 0x055a0, 0x0aba4, 0x0a5b0, 0x052b0, //2070-2079
		0x0b273, 0x06930, 0x07337, 0x06aa0, 0x0ad50, 0x14b55, 0x04b60, 0x0a570, 0x054e4, 0x0d160, //2080-2089
		0x0e968, 0x0d520, 0x0daa0, 0x16aa6, 0x056d0, 0x04ae0, 0x0a9d4, 0x0a2d0, 0x0d150, 0x0f252, //2090-2099
		0x0d520, //2100
	}

	invalidYearError = func(year int) error {
		return fmt.Errorf("invalid year %d, currently only 200 years from 1900 to 2100 are supported", year)
	}
)

// lunar defines a lunar struct.
// 定义 lunar 结构体
type lunar struct {
	year, month, day int  // 农历年、月、日
	isLeapMonth      bool // 是否是闰月
	Error            error
}

// Lunar converts the gregorian calendar to the lunar calendar.
// 将公历转为农历
func (c Carbon) Lunar() (l lunar) {
	if c.IsInvalid() {
		l.Error = c.Error
		return
	}
	// leapMonths:闰月总数，daysOfYear:年天数，daysOfMonth:月天数，leapMonth:闰月月份
	daysInYear, daysInMonth, leapMonth := 365, 30, 0
	// 有效范围检验
	if c.Year() < minYear || c.Year() > maxYear {
		l.Error = invalidYearError(c.Year())
		return
	}
	offset := int(c.DiffInDaysWithAbs(c.CreateFromDateTime(minYear, 1, 31, 0, 0, 0, 0)))
	for l.year = minYear; l.year <= maxYear && offset > 0; l.year++ {
		daysInYear = l.getDaysInYear()
		offset -= daysInYear
	}
	if offset < 0 {
		offset += daysInYear
		l.year--
	}
	l.isLeapMonth = false
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
	// offset为0时，并且刚才计算的月份是闰月，要校正
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
	return
}

// getDaysInYear gets total days in lunar year.
// 获取该年总天数
func (l lunar) getDaysInYear() int {
	var sum = 348
	for i := 0x8000; i > 0x8; i >>= 1 {
		if (lunarTerms[l.year-minYear] & i) != 0 {
			sum++
		}
	}
	return sum + l.getDaysInLeapMonth()
}

// getDaysInMonth gets total days in lunar month.
// 获取该月总天数
func (l lunar) getDaysInMonth() int {
	if (lunarTerms[l.year-minYear] & (0x10000 >> uint(l.month))) == 0 {
		return 29
	}
	return 30
}

// getDaysInLeapMonth gets total days in lunar leap month.
// 获取闰月总天数
func (l lunar) getDaysInLeapMonth() int {
	if l.LeapMonth() == 0 {
		return 0
	}
	if (lunarTerms[l.year-minYear] & 0x10000) != 0 {
		return 30
	}
	return 29
}

// Animal gets lunar animal name.
// 获取生肖
func (l lunar) Animal() string {
	if l.year == 0 {
		return ""
	}
	return animals[l.year%MonthsPerYear]
}

// Festival gets lunar festival name.
// 获取农历节日
func (l lunar) Festival() string {
	if l.year == 0 {
		return ""
	}
	switch {
	case l.month == 1 && l.day == 1:
		return festivals[0]
	case l.month == 1 && l.day == 15:
		return festivals[1]
	case l.month == 5 && l.day == 5:
		return festivals[2]
	case l.month == 7 && l.day == 7:
		return festivals[3]
	case l.month == 7 && l.day == 15:
		return festivals[4]
	case l.month == 8 && l.day == 15:
		return festivals[5]
	case l.month == 9 && l.day == 9:
		return festivals[6]
	case l.month == 10 && l.day == 1:
		return festivals[7]
	case l.month == 10 && l.day == 15:
		return festivals[8]
	case l.month == 12 && l.day == 8:
		return festivals[9]
	case l.month == 12 && l.day == 23:
		return festivals[10]
	}
	return ""
}

// Year gets lunar year.
// 获取农历年
func (l lunar) Year() int {
	return l.year
}

// Month gets lunar month.
// 获取农历月
func (l lunar) Month() int {
	return l.month
}

// LeapMonth gets lunar leap month.
//获取农历闰月月份
func (l lunar) LeapMonth() int {
	if l.year == 0 {
		return 0
	}
	return lunarTerms[l.year-minYear] & 0xf
}

// Day gets lunar day.
// 获取农历日
func (l lunar) Day() int {
	return l.day
}

// ToYearString outputs a string in lunar year format.
// 获取农历年字符串
func (l lunar) ToYearString() string {
	if l.year == 0 {
		return ""
	}
	year := fmt.Sprintf("%d", l.year)
	for i, replace := range chineseNumbers {
		year = strings.Replace(year, fmt.Sprintf("%d", i), replace, -1)
	}
	return year
}

// ToMonthString outputs a string in lunar month format.
// 获取农历月字符串
func (l lunar) ToMonthString() string {
	if l.month == 0 {
		return ""
	}
	return chineseMonths[l.month-1]
}

// ToDayString outputs a string in lunar day format.
// 获取农历日字符串
func (l lunar) ToDayString() string {
	if l.day == 0 {
		return ""
	}
	day := ""
	switch l.day {
	case 10:
		day = "初十"
	case 20:
		day = "二十"
	case 30:
		day = "三十"
	default:
		day = chineseDayPrefixes[(l.day/10)] + chineseNumbers[l.day%10]
	}
	return day
}

// ToString outputs a string in lunar date format.
// 获取农历日期字符串
func (l lunar) ToDateString() string {
	if l.year == 0 {
		return ""
	}
	return l.ToYearString() + "年" + l.ToMonthString() + "月" + l.ToDayString()
}

// String outputs a string in YYYY-MM-DD format, implement Stringer interface.
// 输出 YYYY-MM-DD 格式字符串， 实现 Stringer 接口
func (l lunar) String() string {
	if l.year == 0 {
		return ""
	}
	return fmt.Sprintf("%d-%02d-%02d", l.year, l.month, l.day)
}

// IsLeapYear reports whether is leap year.
// 是否是闰年
func (l lunar) IsLeapYear() bool {
	if l.year == 0 {
		return false
	}
	return l.LeapMonth() != 0
}

// IsLeapMonth reports whether is leap month.
// 是否是闰月
func (l lunar) IsLeapMonth() bool {
	if l.month == 0 {
		return false
	}
	return l.month == l.LeapMonth()
}

// IsRatYear reports whether is year of Rat.
// 是否是鼠年
func (l lunar) IsRatYear() bool {
	if l.year == 0 {
		return false
	}
	if l.year%MonthsPerYear == 4 {
		return true
	}
	return false
}

// IsOxYear reports whether is year of Ox.
// 是否是牛年
func (l lunar) IsOxYear() bool {
	if l.year == 0 {
		return false
	}
	if l.year%MonthsPerYear == 5 {
		return true
	}
	return false
}

// IsTigerYear reports whether is year of Tiger.
// 是否是虎年
func (l lunar) IsTigerYear() bool {
	if l.year == 0 {
		return false
	}
	if l.year%MonthsPerYear == 6 {
		return true
	}
	return false
}

// IsRabbitYear reports whether is year of Rabbit.
// 是否是兔年
func (l lunar) IsRabbitYear() bool {
	if l.year == 0 {
		return false
	}
	if l.year%MonthsPerYear == 7 {
		return true
	}
	return false
}

// IsDragonYear reports whether is year of Dragon.
// 是否是龙年
func (l lunar) IsDragonYear() bool {
	if l.year == 0 {
		return false
	}
	if l.year%MonthsPerYear == 8 {
		return true
	}
	return false
}

// IsSnakeYear reports whether is year of Snake.
// 是否是蛇年
func (l lunar) IsSnakeYear() bool {
	if l.year == 0 {
		return false
	}
	if l.year%MonthsPerYear == 9 {
		return true
	}
	return false
}

// IsHorseYear reports whether is year of Horse.
// 是否是马年
func (l lunar) IsHorseYear() bool {
	if l.year == 0 {
		return false
	}
	if l.year%MonthsPerYear == 10 {
		return true
	}
	return false
}

// IsGoatYear reports whether is year of Goat.
// 是否是羊年
func (l lunar) IsGoatYear() bool {
	if l.year == 0 {
		return false
	}
	if l.year%MonthsPerYear == 11 {
		return true
	}
	return false
}

// IsMonkeyYear reports whether is year of Monkey.
// 是否是猴年
func (l lunar) IsMonkeyYear() bool {
	if l.year == 0 {
		return false
	}
	if l.year%MonthsPerYear == 0 {
		return true
	}
	return false
}

// IsRoosterYear reports whether is year of Rooster.
// 是否是鸡年
func (l lunar) IsRoosterYear() bool {
	if l.year == 0 {
		return false
	}
	if l.year%MonthsPerYear == 1 {
		return true
	}
	return false
}

// IsDogYear reports whether is year of Dog.
// 是否是狗年
func (l lunar) IsDogYear() bool {
	if l.year == 0 {
		return false
	}
	if l.year%MonthsPerYear == 2 {
		return true
	}
	return false
}

// IsPigYear reports whether is year of Pig.
// 是否是猪年
func (l lunar) IsPigYear() bool {
	if l.year == 0 {
		return false
	}
	if l.year%MonthsPerYear == 3 {
		return true
	}
	return false
}
