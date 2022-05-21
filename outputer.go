package carbon

import (
	"bytes"
	"strconv"
	"strings"
)

// String implements the interface Stringer for Carbon struct.
// 实现 Stringer 接口
func (c Carbon) String() string {
	return c.ToDateTimeString()
}

// ToString outputs a string in "2006-01-02 15:04:05.999999999 -0700 MST" layout.
// 输出 "2006-01-02 15:04:05.999999999 -0700 MST" 格式字符串
func (c Carbon) ToString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().String()
}

// ToMonthString outputs a string in month layout like "January", i18n is supported.
// 输出完整月份字符串，支持i18n
func (c Carbon) ToMonthString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	if len(c.lang.resources) == 0 {
		c.lang.SetLocale(defaultLocale)
	}
	if months, ok := c.lang.resources["months"]; ok {
		slice := strings.Split(months, "|")
		if len(slice) == 12 {
			return slice[c.Month()-1]
		}
	}
	return ""
}

// ToShortMonthString outputs a string in short month layout like "Jan", i18n is supported.
// 输出缩写月份字符串，支持i18n
func (c Carbon) ToShortMonthString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	if len(c.lang.resources) == 0 {
		c.lang.SetLocale(defaultLocale)
	}
	if months, ok := c.lang.resources["short_months"]; ok {
		slice := strings.Split(months, "|")
		if len(slice) == 12 {
			return slice[c.Month()-1]
		}
	}
	return ""
}

// ToWeekString outputs a string in week layout like "Sunday", i18n is supported.
// 输出完整星期字符串，支持i18n
func (c Carbon) ToWeekString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	if len(c.lang.resources) == 0 {
		c.lang.SetLocale(defaultLocale)
	}
	if months, ok := c.lang.resources["weeks"]; ok {
		slice := strings.Split(months, "|")
		if len(slice) == 7 {
			return slice[c.Week()]
		}
	}
	return ""
}

// ToShortWeekString outputs a string in short week layout like "Sun", i18n is supported.
// 输出缩写星期字符串，支持i18n
func (c Carbon) ToShortWeekString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	if len(c.lang.resources) == 0 {
		c.lang.SetLocale(defaultLocale)
	}
	if months, ok := c.lang.resources["short_weeks"]; ok {
		slice := strings.Split(months, "|")
		if len(slice) == 7 {
			return slice[c.Week()]
		}
	}
	return ""
}

// ToDayDateTimeString outputs a string in "Mon, Jan 2, 2006 3:04 PM" layout.
// 输出 "Mon, Jan 2, 2006 3:04 PM" 格式字符串
func (c Carbon) ToDayDateTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(DayDateTimeLayout)
}

// ToDateTimeString outputs a string in "2006-01-02 15:04:05" layout.
// 输出 "2006-01-02 15:04:05" 格式字符串
func (c Carbon) ToDateTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(DateTimeLayout)
}

// ToDateTimeMilliString outputs a string in "2006-01-02 15:04:05.999" layout.
// 输出 "2006-01-02 15:04:05.999" 格式字符串
func (c Carbon) ToDateTimeMilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(DateTimeMilliLayout)
}

// ToDateTimeMicroString outputs a string in "2006-01-02 15:04:05.999999" layout.
// 输出 "2006-01-02 15:04:05.999999" 格式字符串
func (c Carbon) ToDateTimeMicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(DateTimeMicroLayout)
}

// ToDateTimeNanoString outputs a string in "2006-01-02 15:04:05.999999999" layout.
// 输出 "2006-01-02 15:04:05.999999999" 格式字符串
func (c Carbon) ToDateTimeNanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(DateTimeNanoLayout)
}

// ToShortDateTimeString outputs a string in "20060102150405" layout.
// 输出 "20060102150405" 格式字符串
func (c Carbon) ToShortDateTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortDateTimeLayout)
}

// ToShortDateTimeMilliString outputs a string in "20060102150405.999" layout.
// 输出 "20060102150405.999" 格式字符串
func (c Carbon) ToShortDateTimeMilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortDateTimeMilliLayout)
}

// ToShortDateTimeMicroString outputs a string in "20060102150405.999999" layout.
// 输出 "20060102150405.999999" 格式字符串
func (c Carbon) ToShortDateTimeMicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortDateTimeMicroLayout)
}

// ToShortDateTimeNanoString outputs a string in "20060102150405.999999999" layout.
// 输出 "20060102150405.999999999" 格式字符串
func (c Carbon) ToShortDateTimeNanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortDateTimeNanoLayout)
}

// ToDateString outputs a string in "2006-01-02" layout.
// 输出 "2006-01-02" 格式字符串
func (c Carbon) ToDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(DateLayout)
}

// ToDateMilliString outputs a string in "2006-01-02.999" layout.
// 输出 "2006-01-02.999" 格式字符串
func (c Carbon) ToDateMilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(DateMilliLayout)
}

// ToDateMicroString outputs a string in "2006-01-02.999999" layout.
// 输出 "2006-01-02.999999" 格式字符串
func (c Carbon) ToDateMicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(DateMicroLayout)
}

// ToDateNanoString outputs a string in "2006-01-02.999999999" layout.
// 输出 "2006-01-02.999999999" 格式字符串
func (c Carbon) ToDateNanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(DateNanoLayout)
}

// ToShortDateString outputs a string in "20060102" layout.
// 输出 "20060102" 格式字符串
func (c Carbon) ToShortDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortDateLayout)
}

// ToShortDateMilliString outputs a string in "20060102.999" layout.
// 输出 "20060102.999" 格式字符串
func (c Carbon) ToShortDateMilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortDateMilliLayout)
}

// ToShortDateMicroString outputs a string in "20060102.999999" layout.
// 输出 "20060102.999999" 格式字符串
func (c Carbon) ToShortDateMicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortDateMicroLayout)
}

// ToShortDateNanoString outputs a string in "20060102.999999999" layout.
// 输出 "20060102.999999999" 格式字符串
func (c Carbon) ToShortDateNanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortDateNanoLayout)
}

// ToTimeString outputs a string in "15:04:05" layout.
// 输出 "15:04:05" 格式字符串
func (c Carbon) ToTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(TimeLayout)
}

// ToTimeMilliString outputs a string in "15:04:05.999" layout.
// 输出 "15:04:05.999" 格式字符串
func (c Carbon) ToTimeMilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(TimeMilliLayout)
}

// ToTimeMicroString outputs a string in "15:04:05.999999" layout.
// 输出 "15:04:05.999999" 格式字符串
func (c Carbon) ToTimeMicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(TimeMicroLayout)
}

// ToTimeNanoString outputs a string in "15:04:05.999999999" layout.
// 输出 "15:04:05.999999999" 格式字符串
func (c Carbon) ToTimeNanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(TimeNanoLayout)
}

// ToShortTimeString outputs a string in "150405" layout.
// 输出 "150405" 格式字符串
func (c Carbon) ToShortTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortTimeLayout)
}

// ToShortTimeMilliString outputs a string in "150405.999" layout.
// 输出 "150405.999" 格式字符串
func (c Carbon) ToShortTimeMilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortTimeMilliLayout)
}

// ToShortTimeMicroString outputs a string in "150405.999999" layout.
// 输出 "150405.999999" 格式字符串
func (c Carbon) ToShortTimeMicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortTimeMicroLayout)
}

// ToShortTimeNanoString outputs a string in "150405.999999999" layout.
// 输出 "150405.999999999" 格式字符串
func (c Carbon) ToShortTimeNanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ShortTimeNanoLayout)
}

// ToAtomString outputs a string in "2006-01-02T15:04:05Z07:00" layout.
// 输出 "2006-01-02T15:04:05Z07:00" 格式字符串
func (c Carbon) ToAtomString(timezone ...string) string {
	return c.ToRfc3339String(timezone...)
}

// ToANSICString outputs a string in "Mon Jan _2 15:04:05 2006" layout.
// 输出 "Mon Jan _2 15:04:05 2006" 格式字符串
func (c Carbon) ToANSICString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ANSICLayout)
}

// ToCookieString outputs a string in "Monday, 02-Jan-2006 15:04:05 MST" layout.
// 输出 "Monday, 02-Jan-2006 15:04:05 MST" 格式字符串
func (c Carbon) ToCookieString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(CookieLayout)
}

// ToRssString outputs a string in "Mon, 02 Jan 2006 15:04:05 -0700" format.
// 输出 "Mon, 02 Jan 2006 15:04:05 -0700" 格式字符串
func (c Carbon) ToRssString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RssLayout)
}

// ToW3cString outputs a string in "2006-01-02T15:04:05Z07:00" layout.
// 输出 "2006-01-02T15:04:05Z07:00" 格式字符串
func (c Carbon) ToW3cString(timezone ...string) string {
	return c.ToRfc3339String(timezone...)
}

// ToUnixDateString outputs a string in "Mon Jan _2 15:04:05 MST 2006" layout.
// 输出 "Mon Jan _2 15:04:05 MST 2006" 格式字符串
func (c Carbon) ToUnixDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(UnixDateLayout)
}

// ToRubyDateString outputs a string in "Mon Jan 02 15:04:05 -0700 2006" layout.
// 输出 "Mon Jan 02 15:04:05 -0700 2006" 格式字符串
func (c Carbon) ToRubyDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RubyDateLayout)
}

// ToKitchenString outputs a string in "3:04PM" layout.
// 输出 "3:04PM" 格式字符串
func (c Carbon) ToKitchenString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(KitchenLayout)
}

// ToIso8601String outputs a string in "2006-01-02T15:04:05-07:00" layout.
// 输出 "2006-01-02T15:04:05-07:00" 格式字符串
func (c Carbon) ToIso8601String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ISO8601Layout)
}

// ToIso8601MilliString outputs a string in "2006-01-02T15:04:05.999-07:00" layout.
// 输出 "2006-01-02T15:04:05.999-07:00" 格式字符串
func (c Carbon) ToIso8601MilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ISO8601MilliLayout)
}

// ToIso8601MicroString outputs a string in "2006-01-02T15:04:05.999999-07:00" layout.
// 输出 "2006-01-02T15:04:05.999999-07:00" 格式字符串
func (c Carbon) ToIso8601MicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ISO8601MicroLayout)
}

// ToIso8601NanoString outputs a string in "2006-01-02T15:04:05.999999999-07:00" layout.
// 输出 "2006-01-02T15:04:05.999999999-07:00" 格式字符串
func (c Carbon) ToIso8601NanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(ISO8601NanoLayout)
}

// ToRfc822String outputs a string in "02 Jan 06 15:04 MST" layout.
// 输出 "02 Jan 06 15:04 MST" 格式字符串
func (c Carbon) ToRfc822String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC822Layout)
}

// ToRfc822zString outputs a string in "02 Jan 06 15:04 -0700" layout.
// 输出 "02 Jan 06 15:04 -0700" 格式字符串
func (c Carbon) ToRfc822zString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC822ZLayout)
}

// ToRfc850String outputs a string in "Monday, 02-Jan-06 15:04:05 MST" layout.
// 输出 "Monday, 02-Jan-06 15:04:05 MST" 格式字符串
func (c Carbon) ToRfc850String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC850Layout)
}

// ToRfc1036String outputs a string in "Mon, 02 Jan 06 15:04:05 -0700" layout.
// 输出 "Mon, 02 Jan 06 15:04:05 -0700" 格式字符串
func (c Carbon) ToRfc1036String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC1036Layout)
}

// ToRfc1123String outputs a string in "Mon, 02 Jan 2006 15:04:05 MST" layout.
// 输出 "Mon, 02 Jan 2006 15:04:05 MST" 格式字符串
func (c Carbon) ToRfc1123String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC1123Layout)
}

// ToRfc1123zString outputs a string in "Mon, 02 Jan 2006 15:04:05 -0700" layout.
// 输出 "Mon, 02 Jan 2006 15:04:05 -0700" 格式字符串
func (c Carbon) ToRfc1123zString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC1123ZLayout)
}

// ToRfc2822String outputs a string in "Mon, 02 Jan 2006 15:04:05 -0700" layout.
// 输出 "Mon, 02 Jan 2006 15:04:05 -0700" 格式字符串
func (c Carbon) ToRfc2822String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC2822Layout)
}

// ToRfc3339String outputs a string in "2006-01-02T15:04:05Z07:00" layout.
// 输出 "2006-01-02T15:04:05Z07:00" 格式字符串
func (c Carbon) ToRfc3339String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC3339Layout)
}

// ToRfc3339MilliString outputs a string in "2006-01-02T15:04:05.999Z07:00" layout.
// 输出 "2006-01-02T15:04:05.999Z07:00" 格式字符串
func (c Carbon) ToRfc3339MilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC3339MilliLayout)
}

// ToRfc3339MicroString outputs a string in "2006-01-02T15:04:05.999999Z07:00" layout.
// 输出 "2006-01-02T15:04:05.999999Z07:00" 格式字符串
func (c Carbon) ToRfc3339MicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC3339MicroLayout)
}

// ToRfc3339NanoString outputs a string in "2006-01-02T15:04:05.999999999Z07:00" layout.
// 输出 "2006-01-02T15:04:05.999999999Z07:00" 格式字符串
func (c Carbon) ToRfc3339NanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC3339NanoLayout)
}

// ToRfc7231String outputs a string in "Mon, 02 Jan 2006 15:04:05 GMT" layout.
// 输出 "Mon, 02 Jan 2006 15:04:05 GMT" 格式字符串
func (c Carbon) ToRfc7231String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Carbon2Time().Format(RFC7231Layout)
}

// ToLayoutString outputs a string by layout.
// 输出指定布局模板的时间字符串
func (c Carbon) ToLayoutString(layout string, timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}

	// if it has \\ before any word/char, then it should be escaped!
	// we will not enumerate all the patterns
	// we will memorize the position of each escape by doing:
	// replace them in the layout with OTHER SPECIAL IDENTIFIABLE WORDS, and then put them back after formatting!

	type escapeWord struct {
		ID           string
		OriginalWord string
	}

	buffer := bytes.NewBuffer(nil)
	var escapedWords []escapeWord

	escapePattern := "\\"

	layoutRunes := []rune(layout)

	foundEscapedWord := false
	escWord := ""
	escapeID := ""
	isLastRow := false

	nrOfChars := len(layoutRunes)

	addEscapedWord := func(id, word string) {
		// reset
		foundEscapedWord = false
		escWord = ""
		escapeID = ""
		// add
		escapedWords = append(escapedWords, escapeWord{
			ID:           id,
			OriginalWord: word,
		})
	}

	for i := 0; i < nrOfChars; i++ {
		char := string(layoutRunes[i])

		if nrOfChars == i+1 {
			isLastRow = true
		}

		switch char {
		case escapePattern:
			if foundEscapedWord {
				// if we had already before an escape pattern, then copy here the word
				addEscapedWord(escapeID, escWord)
			}

			// Generate a unique id for later replacement with original content/word
			escapeID = "#" + randStringRunes(5) + "#"
			foundEscapedWord = true

			// We have found an escape pattern
			// also check if multiple times the pattern has been written

			// Write the escape ID to the buffer layout
			buffer.WriteString(escapeID)
		case " ":
			if foundEscapedWord {
				// there is a space, end of word...
				addEscapedWord(escapeID, escWord)
			}
			// write the space to the buffer
			buffer.WriteString(char)
		default:
			//

			if foundEscapedWord {
				// add to the word
				escWord = escWord + char
				if isLastRow {
					// last row...end of word...
					addEscapedWord(escapeID, escWord)
				}
			} else {
				// Write the current char!
				buffer.WriteString(char)
			}
		}
	}

	nrOfEscapedWords := len(escapedWords)
	if nrOfEscapedWords > 0 {
		layout = buffer.String()
	}

	// go doesn't support at the moment i18n (internalization) for time.Time -> it's in a roadmap
	// So, a temporary solution would be: "Check some specific patterns, and try replacing them!?"

	// Usual cases are Month,Short Month, Week, Short Week, Season, Short Season
	// TODO: these patterns should be moved from here
	layoutConvPatterns := []struct {
		pattern     string
		convHandler func() string
	}{
		// The patterns should be in this specific order!
		{
			pattern: "January",
			convHandler: func() string {
				return c.ToMonthString()
			},
		},
		{
			pattern: "Jan",
			convHandler: func() string {
				return c.ToShortMonthString()
			},
		},
		{
			pattern: "Monday",
			convHandler: func() string {
				return c.ToWeekString()
			},
		},
		{
			pattern: "Mon",
			convHandler: func() string {
				return c.ToShortWeekString()
			},
		},
	}

	for _, patternItem := range layoutConvPatterns {
		if strings.Contains(layout, patternItem.pattern) {
			layout = strings.ReplaceAll(layout, patternItem.pattern, patternItem.convHandler())
		}
	}

	formatted := c.Carbon2Time().Format(layout)

	// now let's put back the escaped words if there are any!
	if nrOfEscapedWords > 0 {
		for _, escapedItem := range escapedWords {
			formatted = strings.ReplaceAll(formatted, escapedItem.ID, escapedItem.OriginalWord)
		}
	}

	return formatted
}

// Layout outputs a string by layout, it is shorthand for ToLayoutString.
// 输出指定布局模板的时间字符串, 是 ToLayoutString 的简写
func (c Carbon) Layout(layout string, timezone ...string) string {
	return c.ToLayoutString(layout, timezone...)
}

// ToFormatString outputs a string by format.
// 输出指定格式模板的时间字符串
func (c Carbon) ToFormatString(format string, timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	buffer := bytes.NewBuffer(nil)
	for i := 0; i < len(format); i++ {
		if layout, ok := formats[format[i]]; ok {
			// support for i18n specific symbols
			switch format[i] {
			case 'l': // week, such as Monday
				buffer.WriteString(c.ToWeekString())
			case 'D': // short week, such as Mon
				buffer.WriteString(c.ToShortWeekString())
			case 'F': // month, such as January
				buffer.WriteString(c.ToMonthString())
			case 'M': // short month, such as Jan
				buffer.WriteString(c.ToShortMonthString())
			default: // common symbols
				buffer.WriteString(c.Carbon2Time().Format(layout))
			}
		} else {
			switch format[i] {
			case '\\': // raw output, no parse
				buffer.WriteByte(format[i+1])
				i++
				continue
			case 'W': // week number of the year in ISO-8601 format, ranging from 1-52
				buffer.WriteString(strconv.Itoa(c.WeekOfYear()))
			case 'N': // day of the week as a number in ISO-8601 format, ranging from 1-7
				buffer.WriteString(strconv.Itoa(c.DayOfWeek()))
			case 'S': // abbreviated suffix for the day of the month, such as st, nd, rd, th
				suffix := "th"
				switch c.Day() {
				case 1, 21, 31:
					suffix = "st"
				case 2, 22:
					suffix = "nd"
				case 3, 23:
					suffix = "rd"
				}
				buffer.WriteString(suffix)
			case 'L': // whether it is a leap year, if it is a leap year, it is 1, otherwise it is 0
				if c.IsLeapYear() {
					buffer.WriteString("1")
				} else {
					buffer.WriteString("0")
				}
			case 'G': // 24-hour format, no padding, ranging from 0-23
				buffer.WriteString(strconv.Itoa(c.Hour()))
			case 'U': // timestamp with second, such as 1611818268
				buffer.WriteString(strconv.FormatInt(c.Timestamp(), 10))
			case 'u': // current millisecond, such as 999
				buffer.WriteString(strconv.Itoa(c.Millisecond()))
			case 'w': // day of the week represented by the number, ranging from 0-6
				buffer.WriteString(strconv.Itoa(c.DayOfWeek() - 1))
			case 't': // number of days in the month, ranging from 28-31
				buffer.WriteString(strconv.Itoa(c.DaysInMonth()))
			case 'z': // day of the year, ranging from 0-365
				buffer.WriteString(strconv.Itoa(c.DayOfYear() - 1))
			case 'e': // current location, such as UTC，GMT，Atlantic/Azores
				buffer.WriteString(c.Location())
			case 'Q': // current quarter, ranging from 1-4
				buffer.WriteString(strconv.Itoa(c.Quarter()))
			case 'C': // current century, ranging from 0-99
				buffer.WriteString(strconv.Itoa(c.Century()))
			default:
				buffer.WriteByte(format[i])
			}
		}
	}
	return buffer.String()
}

// Format outputs a string by format, it is shorthand for ToFormatString.
// 输出指定格式模板的时间字符串, 是 ToFormatString 的简写
func (c Carbon) Format(format string, timezone ...string) string {
	return c.ToFormatString(format, timezone...)
}
