package carbon

import (
	"time"
)

// Version current version
const Version = "2.6.15"

// timezone constants
const (
	Local = "Local"
	UTC   = "UTC"

	CET  = "CET"
	EET  = "EET"
	EST  = "EST"
	GMT  = "GMT"
	MET  = "MET"
	MST  = "MST"
	UCT  = "MST"
	WET  = "WET"
	Zulu = "Zulu"

	Cuba      = "Cuba"
	Egypt     = "Egypt"
	Eire      = "Eire"
	Greenwich = "Greenwich"
	Iceland   = "Iceland"
	Iran      = "Iran"
	Israel    = "Israel"
	Jamaica   = "Jamaica"
	Japan     = "Japan"
	Libya     = "Libya"
	Poland    = "Poland"
	Portugal  = "Portugal"
	PRC       = "PRC"
	Singapore = "Singapore"
	Turkey    = "Turkey"

	Shanghai   = "Asia/Shanghai"
	Chongqing  = "Asia/Chongqing"
	Harbin     = "Asia/Harbin"
	Urumqi     = "Asia/Urumqi"
	HongKong   = "Asia/Hong_Kong"
	Macao      = "Asia/Macao"
	Taipei     = "Asia/Taipei"
	Tokyo      = "Asia/Tokyo"
	HoChiMinh  = "Asia/Ho_Chi_Minh"
	Hanoi      = "Asia/Hanoi"
	Saigon     = "Asia/Saigon"
	Seoul      = "Asia/Seoul"
	Pyongyang  = "Asia/Pyongyang"
	Bangkok    = "Asia/Bangkok"
	Dubai      = "Asia/Dubai"
	Qatar      = "Asia/Qatar"
	Bangalore  = "Asia/Bangalore"
	Kolkata    = "Asia/Kolkata"
	Mumbai     = "Asia/Mumbai"
	MexicoCity = "America/Mexico_City"
	NewYork    = "America/New_York"
	LosAngeles = "America/Los_Angeles"
	Chicago    = "America/Chicago"
	SaoPaulo   = "America/Sao_Paulo"
	Moscow     = "Europe/Moscow"
	London     = "Europe/London"
	Berlin     = "Europe/Berlin"
	Paris      = "Europe/Paris"
	Rome       = "Europe/Rome"
	Sydney     = "Australia/Sydney"
	Melbourne  = "Australia/Melbourne"
	Darwin     = "Australia/Darwin"
)

// month constants
const (
	January   = time.January
	February  = time.February
	March     = time.March
	April     = time.April
	May       = time.May
	June      = time.June
	July      = time.July
	August    = time.August
	September = time.September
	October   = time.October
	November  = time.November
	December  = time.December
)

// season constants
const (
	Spring = "Spring"
	Summer = "Summer"
	Autumn = "Autumn"
	Winter = "Winter"
)

// constellation constants
const (
	Aries       = "Aries"
	Taurus      = "Taurus"
	Gemini      = "Gemini"
	Cancer      = "Cancer"
	Leo         = "Leo"
	Virgo       = "Virgo"
	Libra       = "Libra"
	Scorpio     = "Scorpio"
	Sagittarius = "Sagittarius"
	Capricorn   = "Capricorn"
	Aquarius    = "Aquarius"
	Pisces      = "Pisces"
)

// week constants
const (
	Monday    = time.Monday
	Tuesday   = time.Tuesday
	Wednesday = time.Wednesday
	Thursday  = time.Thursday
	Friday    = time.Friday
	Saturday  = time.Saturday
	Sunday    = time.Sunday
)

// number constants
const (
	EpochYear          = 1970
	YearsPerMillennium = 1000
	YearsPerCentury    = 100
	YearsPerDecade     = 10
	QuartersPerYear    = 4
	MonthsPerYear      = 12
	MonthsPerQuarter   = 3
	WeeksPerNormalYear = 52
	WeeksPerLongYear   = 53
	WeeksPerMonth      = 4
	DaysPerLeapYear    = 366
	DaysPerNormalYear  = 365
	DaysPerWeek        = 7
	HoursPerWeek       = 168
	HoursPerDay        = 24
	MinutesPerDay      = 1440
	MinutesPerHour     = 60
	SecondsPerWeek     = 604800
	SecondsPerDay      = 86400
	SecondsPerHour     = 3600
	SecondsPerMinute   = 60
)

// max constants
const (
	MaxYear       = 9999
	MaxMonth      = 12
	MaxDay        = 31
	MaxHour       = 23
	MaxMinute     = 59
	MaxSecond     = 59
	MaxNanosecond = 999999999
)

// min constants
const (
	MinYear       = 1
	MinMonth      = 1
	MinDay        = 1
	MinHour       = 0
	MinMinute     = 0
	MinSecond     = 0
	MinNanosecond = 0
)

// layout constants
const (
	AtomLayout     = RFC3339Layout
	ANSICLayout    = time.ANSIC
	CookieLayout   = "Monday, 02-Jan-2006 15:04:05 MST"
	KitchenLayout  = time.Kitchen
	RssLayout      = time.RFC1123Z
	RubyDateLayout = time.RubyDate
	UnixDateLayout = time.UnixDate
	W3cLayout      = RFC3339Layout
	HttpLayout     = "Mon, 02 Jan 2006 15:04:05 GMT"

	RFC1036Layout      = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC1123Layout      = time.RFC1123
	RFC1123ZLayout     = time.RFC1123Z
	RFC2822Layout      = time.RFC1123Z
	RFC3339Layout      = "2006-01-02T15:04:05Z07:00"
	RFC3339MilliLayout = "2006-01-02T15:04:05.999Z07:00"
	RFC3339MicroLayout = "2006-01-02T15:04:05.999999Z07:00"
	RFC3339NanoLayout  = "2006-01-02T15:04:05.999999999Z07:00"
	RFC7231Layout      = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC822Layout       = time.RFC822
	RFC822ZLayout      = time.RFC822Z
	RFC850Layout       = time.RFC850

	ISO8601Layout      = "2006-01-02T15:04:05-07:00"
	ISO8601MilliLayout = "2006-01-02T15:04:05.999-07:00"
	ISO8601MicroLayout = "2006-01-02T15:04:05.999999-07:00"
	ISO8601NanoLayout  = "2006-01-02T15:04:05.999999999-07:00"

	ISO8601ZuluLayout      = "2006-01-02T15:04:05Z"
	ISO8601ZuluMilliLayout = "2006-01-02T15:04:05.999Z"
	ISO8601ZuluMicroLayout = "2006-01-02T15:04:05.999999Z"
	ISO8601ZuluNanoLayout  = "2006-01-02T15:04:05.999999999Z"

	FormattedDateLayout    = "Jan 2, 2006"
	FormattedDayDateLayout = "Mon, Jan 2, 2006"

	DayDateTimeLayout        = "Mon, Jan 2, 2006 3:04 PM"
	DateTimeLayout           = "2006-01-02 15:04:05"
	DateTimeMilliLayout      = "2006-01-02 15:04:05.999"
	DateTimeMicroLayout      = "2006-01-02 15:04:05.999999"
	DateTimeNanoLayout       = "2006-01-02 15:04:05.999999999"
	ShortDateTimeLayout      = "20060102150405"
	ShortDateTimeMilliLayout = "20060102150405.999"
	ShortDateTimeMicroLayout = "20060102150405.999999"
	ShortDateTimeNanoLayout  = "20060102150405.999999999"

	DateLayout           = "2006-01-02"
	DateMilliLayout      = "2006-01-02.999"
	DateMicroLayout      = "2006-01-02.999999"
	DateNanoLayout       = "2006-01-02.999999999"
	ShortDateLayout      = "20060102"
	ShortDateMilliLayout = "20060102.999"
	ShortDateMicroLayout = "20060102.999999"
	ShortDateNanoLayout  = "20060102.999999999"

	TimeLayout           = "15:04:05"
	TimeMilliLayout      = "15:04:05.999"
	TimeMicroLayout      = "15:04:05.999999"
	TimeNanoLayout       = "15:04:05.999999999"
	ShortTimeLayout      = "150405"
	ShortTimeMilliLayout = "150405.999"
	ShortTimeMicroLayout = "150405.999999"
	ShortTimeNanoLayout  = "150405.999999999"

	TimestampLayout      = "unix"
	TimestampMilliLayout = "unixMilli"
	TimestampMicroLayout = "unixMicro"
	TimestampNanoLayout  = "unixNano"
)

// format constants
const (
	AtomFormat     = "Y-m-d\\TH:i:sR"
	ANSICFormat    = "D M  j H:i:s Y"
	CookieFormat   = "l, d-M-Y H:i:s Z"
	KitchenFormat  = "g:iA"
	RssFormat      = "D, d M Y H:i:s O"
	RubyDateFormat = "D M d H:i:s O Y"
	UnixDateFormat = "D M  j H:i:s Z Y"
	W3cFormat      = RFC3339Format
	HttpFormat     = "D, d M Y H:i:s \\G\\M\\T"

	RFC1036Format      = "D, d M y H:i:s O"
	RFC1123Format      = "D, d M Y H:i:s Z"
	RFC1123ZFormat     = "D, d M Y H:i:s O"
	RFC2822Format      = "D, d M Y H:i:s O"
	RFC3339Format      = "Y-m-d\\TH:i:sR"
	RFC3339MilliFormat = "Y-m-d\\TH:i:s.uR"
	RFC3339MicroFormat = "Y-m-d\\TH:i:s.vR"
	RFC3339NanoFormat  = "Y-m-d\\TH:i:s.xR"
	RFC7231Format      = "D, d M Y H:i:s Z"
	RFC822Format       = "d M y H:i Z"
	RFC822ZFormat      = "d M y H:i O"
	RFC850Format       = "l, d-M-y H:i:s Z"

	ISO8601Format      = "Y-m-d\\TH:i:sP"
	ISO8601MilliFormat = "Y-m-d\\TH:i:s.uP"
	ISO8601MicroFormat = "Y-m-d\\TH:i:s.vP"
	ISO8601NanoFormat  = "Y-m-d\\TH:i:s.xP"

	ISO8601ZuluFormat      = "Y-m-d\\TH:i:s\\Z"
	ISO8601ZuluMilliFormat = "Y-m-d\\TH:i:s.u\\Z"
	ISO8601ZuluMicroFormat = "Y-m-d\\TH:i:s.v\\Z"
	ISO8601ZuluNanoFormat  = "Y-m-d\\TH:i:s.x\\Z"

	FormattedDateFormat    = "M j, Y"
	FormattedDayDateFormat = "D, M j, Y"

	DayDateTimeFormat        = "D, M j, Y g:i A"
	DateTimeFormat           = "Y-m-d H:i:s"
	DateTimeMilliFormat      = "Y-m-d H:i:s.u"
	DateTimeMicroFormat      = "Y-m-d H:i:s.v"
	DateTimeNanoFormat       = "Y-m-d H:i:s.x"
	ShortDateTimeFormat      = "YmdHis"
	ShortDateTimeMilliFormat = "YmdHis.u"
	ShortDateTimeMicroFormat = "YmdHis.v"
	ShortDateTimeNanoFormat  = "YmdHis.x"

	DateFormat           = "Y-m-d"
	DateMilliFormat      = "Y-m-d.u"
	DateMicroFormat      = "Y-m-d.v"
	DateNanoFormat       = "Y-m-d.x"
	ShortDateFormat      = "Ymd"
	ShortDateMilliFormat = "Ymd.u"
	ShortDateMicroFormat = "Ymd.v"
	ShortDateNanoFormat  = "Ymd.x"

	TimeFormat           = "H:i:s"
	TimeMilliFormat      = "H:i:s.u"
	TimeMicroFormat      = "H:i:s.v"
	TimeNanoFormat       = "H:i:s.x"
	ShortTimeFormat      = "His"
	ShortTimeMilliFormat = "His.u"
	ShortTimeMicroFormat = "His.v"
	ShortTimeNanoFormat  = "His.x"

	TimestampFormat      = "S"
	TimestampMilliFormat = "U"
	TimestampMicroFormat = "V"
	TimestampNanoFormat  = "X"
)
