// @Package carbon
// @Description a simple, semantic and developer-friendly golang package for datetime
// @Page github.com/golang-module/carbon
// @Developer gouguoyin
// @Blog www.gouguoyin.cn
// @Email contact@gouguoyin.cn

// Package carbon is a simple, semantic and developer-friendly golang package for datetime.
package carbon

import (
	"time"
)

// Version current version
// 当前版本号
const Version = "2.2.3"

// timezone constants
// 时区常量
const (
	Local = "Local" // 本地时间
	UTC   = "UTC"   // 世界协调时间
	GMT   = "GMT"   // 格林尼治标准时间
	CST   = "CST"   // 中国标准时间
	EET   = "EET"   // 欧洲东部标准时间
	WET   = "WET"   // 欧洲西部标准时间
	CET   = "CET"   // 欧洲中部标准时间
	EST   = "EST"   // 美国东部标准时间
	MST   = "MST"   // 美国山地标准时间

	Cuba      = "Cuba"      // 古巴
	Egypt     = "Egypt"     // 埃及
	Eire      = "Eire"      // 爱尔兰
	Greenwich = "Greenwich" // 格林尼治
	Iceland   = "Iceland"   // 冰岛
	Iran      = "Iran"      // 伊朗
	Israel    = "Israel"    // 以色列
	Jamaica   = "Jamaica"   // 牙买加
	Japan     = "Japan"     // 日本
	Libya     = "Libya"     // 利比亚
	Poland    = "Poland"    // 波兰
	Portugal  = "Portugal"  // 葡萄牙
	PRC       = "PRC"       // 中国
	Singapore = "Singapore" // 新加坡
	Turkey    = "Turkey"    // 土耳其

	Shanghai   = "Asia/Shanghai"       // 上海
	Chongqing  = "Asia/Chongqing"      // 重庆
	Harbin     = "Asia/Harbin"         // 哈尔滨
	Urumqi     = "Asia/Urumqi"         // 乌鲁木齐
	HongKong   = "Asia/Hong_Kong"      // 香港
	Macao      = "Asia/Macao"          // 澳门
	Taipei     = "Asia/Taipei"         // 台北
	Tokyo      = "Asia/Tokyo"          // 东京
	Saigon     = "Asia/Saigon"         // 西贡
	Seoul      = "Asia/Seoul"          // 首尔
	Bangkok    = "Asia/Bangkok"        // 曼谷
	Dubai      = "Asia/Dubai"          // 迪拜
	NewYork    = "America/New_York"    // 纽约
	LosAngeles = "America/Los_Angeles" // 洛杉矶
	Chicago    = "America/Chicago"     // 芝加哥
	Moscow     = "Europe/Moscow"       // 莫斯科
	London     = "Europe/London"       // 伦敦
	Berlin     = "Europe/Berlin"       // 柏林
	Paris      = "Europe/Paris"        // 巴黎
	Rome       = "Europe/Rome"         // 罗马
	Sydney     = "Australia/Sydney"    // 悉尼
	Melbourne  = "Australia/Melbourne" // 墨尔本
	Darwin     = "Australia/Darwin"    // 达尔文
)

// month constants
// 月份常量
const (
	January   = "January"   // 一月
	February  = "February"  // 二月
	March     = "March"     // 三月
	April     = "April"     // 四月
	May       = "May"       // 五月
	June      = "June"      // 六月
	July      = "July"      // 七月
	August    = "August"    // 八月
	September = "September" // 九月
	October   = "October"   // 十月
	November  = "November"  // 十一月
	December  = "December"  // 十二月
)

// week constants
// 星期常量
const (
	Monday    = "Monday"    // 周一
	Tuesday   = "Tuesday"   // 周二
	Wednesday = "Wednesday" // 周三
	Thursday  = "Thursday"  // 周四
	Friday    = "Friday"    // 周五
	Saturday  = "Saturday"  // 周六
	Sunday    = "Sunday"    // 周日
)

// number constants
// 数字常量
const (
	YearsPerMillennium = 1000   // 每千年1000年
	YearsPerCentury    = 100    // 每世纪100年
	YearsPerDecade     = 10     // 每十年10年
	QuartersPerYear    = 4      // 每年4个季度
	MonthsPerYear      = 12     // 每年12月
	MonthsPerQuarter   = 3      // 每季度3月
	WeeksPerNormalYear = 52     // 每常规年52周
	weeksPerLongYear   = 53     // 每长年53周
	WeeksPerMonth      = 4      // 每月4周
	DaysPerLeapYear    = 366    // 每闰年366天
	DaysPerNormalYear  = 365    // 每常规年365天
	DaysPerWeek        = 7      // 每周7天
	HoursPerWeek       = 168    // 每周168小时
	HoursPerDay        = 24     // 每天24小时
	MinutesPerDay      = 1440   // 每天1440分钟
	MinutesPerHour     = 60     // 每小时60分钟
	SecondsPerWeek     = 604800 // 每周604800秒
	SecondsPerDay      = 86400  // 每天86400秒
	SecondsPerHour     = 3600   // 每小时3600秒
	SecondsPerMinute   = 60     // 每分钟60秒
)

// layout constants
// 布局模板常量
const (
	ANSICLayout              = time.ANSIC
	UnixDateLayout           = time.UnixDate
	RubyDateLayout           = time.RubyDate
	RFC822Layout             = time.RFC822
	RFC822ZLayout            = time.RFC822Z
	RFC850Layout             = time.RFC850
	RFC1123Layout            = time.RFC1123
	RFC1123ZLayout           = time.RFC1123Z
	RssLayout                = time.RFC1123Z
	KitchenLayout            = time.Kitchen
	RFC2822Layout            = time.RFC1123Z
	CookieLayout             = "Monday, 02-Jan-2006 15:04:05 MST"
	RFC3339Layout            = "2006-01-02T15:04:05Z07:00"
	RFC3339MilliLayout       = "2006-01-02T15:04:05.999Z07:00"
	RFC3339MicroLayout       = "2006-01-02T15:04:05.999999Z07:00"
	RFC3339NanoLayout        = "2006-01-02T15:04:05.999999999Z07:00"
	ISO8601Layout            = "2006-01-02T15:04:05-07:00"
	ISO8601MilliLayout       = "2006-01-02T15:04:05.999-07:00"
	ISO8601MicroLayout       = "2006-01-02T15:04:05.999999-07:00"
	ISO8601NanoLayout        = "2006-01-02T15:04:05.999999999-07:00"
	RFC1036Layout            = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC7231Layout            = "Mon, 02 Jan 2006 15:04:05 MST"
	DayDateTimeLayout        = "Mon, Jan 2, 2006 3:04 PM"
	DateTimeLayout           = "2006-01-02 15:04:05"
	DateTimeMilliLayout      = "2006-01-02 15:04:05.999"
	DateTimeMicroLayout      = "2006-01-02 15:04:05.999999"
	DateTimeNanoLayout       = "2006-01-02 15:04:05.999999999"
	ShortDateTimeLayout      = "20060102150405"
	ShortDateTimeMilliLayout = "20060102150405.999"
	ShortDateTimeMicroLayout = "20060102150405.999999"
	ShortDateTimeNanoLayout  = "20060102150405.999999999"
	DateLayout               = "2006-01-02"
	DateMilliLayout          = "2006-01-02.999"
	DateMicroLayout          = "2006-01-02.999999"
	DateNanoLayout           = "2006-01-02.999999999"
	ShortDateLayout          = "20060102"
	ShortDateMilliLayout     = "20060102.999"
	ShortDateMicroLayout     = "20060102.999999"
	ShortDateNanoLayout      = "20060102.999999999"
	TimeLayout               = "15:04:05"
	TimeMilliLayout          = "15:04:05.999"
	TimeMicroLayout          = "15:04:05.999999"
	TimeNanoLayout           = "15:04:05.999999999"
	ShortTimeLayout          = "150405"
	ShortTimeMilliLayout     = "150405.999"
	ShortTimeMicroLayout     = "150405.999999"
	ShortTimeNanoLayout      = "150405.999999999"
)

// Carbon defines a Carbon struct.
// 定义 Carbon 结构体
type Carbon struct {
	time         time.Time
	weekStartsAt time.Weekday
	loc          *time.Location
	lang         *Language
	Error        error
}

// NewCarbonPointer returns a new Carbon instance.
// 初始化 Carbon 结构体
func NewCarbonPointer() *Carbon {
	c := NewCarbon()
	return &c
}

// NewCarbon returns a new Carbon instance.
// 初始化 Carbon 结构体
func NewCarbon() Carbon {
	return Carbon{weekStartsAt: time.Sunday, loc: time.Local, lang: NewLanguage()}
}

// FromStdTime converts standard time.Time to Carbon.
// 将标准 time.Time 转换成 Carbon
func FromStdTime(tt time.Time) Carbon {
	c := NewCarbon()
	c.time = tt
	c.loc = tt.Location()
	return c
}

// ToStdTime converts Carbon to standard time.Time.
// 将 Carbon 转换成标准 time.Time
func (c Carbon) ToStdTime() time.Time {
	return c.time.In(c.loc)
}

// Time2Carbon converts standard time.Time to Carbon, will be removed in the future, recommended use FromStdTime
// 将标准 time.Time 转换成 Carbon，未来将移除，推荐使用 FromStdTime
func Time2Carbon(tt time.Time) Carbon {
	return FromStdTime(tt)
}

// Carbon2Time converts Carbon to standard time.Time, will be removed in the future, recommended use ToStdTime
// 将 Carbon 转换成标准 time.Time，未来将移除，推荐使用 ToStdTime
func (c Carbon) Carbon2Time() time.Time {
	return c.ToStdTime()
}
