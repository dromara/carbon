package carbon

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// GoString implements "fmt.GoStringer" interface for Carbon struct.
func (c *Carbon) GoString() string {
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().GoString()
}

// ToString outputs a string in "2006-01-02 15:04:05.999999999 -0700 MST" layout.
func (c *Carbon) ToString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().String()
}

// ToMonthString outputs a string in month layout like "January", i18n is supported.
func (c *Carbon) ToMonthString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}

	lang := c.lang
	if lang == nil {
		return ""
	}

	lang.rw.RLock()
	defer lang.rw.RUnlock()

	if resources, ok := lang.resources["months"]; ok {
		slice := strings.Split(resources, "|")
		if len(slice) == MonthsPerYear {
			return slice[c.Month()-1]
		}
	}
	return ""
}

// ToShortMonthString outputs a string in short month layout like "Jan", i18n is supported.
func (c *Carbon) ToShortMonthString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}

	lang := c.lang
	if lang == nil {
		return ""
	}

	lang.rw.RLock()
	defer lang.rw.RUnlock()

	if resources, ok := lang.resources["short_months"]; ok {
		slice := strings.Split(resources, "|")
		if len(slice) == MonthsPerYear {
			return slice[c.Month()-1]
		}
	}
	return ""
}

// ToWeekString outputs a string in week layout like "Sunday", i18n is supported.
func (c *Carbon) ToWeekString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}

	lang := c.lang
	if lang == nil {
		return ""
	}

	lang.rw.RLock()
	defer lang.rw.RUnlock()

	if resources, ok := lang.resources["weeks"]; ok {
		slice := strings.Split(resources, "|")
		if len(slice) == DaysPerWeek {
			return slice[c.DayOfWeek()%DaysPerWeek]
		}
	}
	return ""
}

// ToShortWeekString outputs a string in short week layout like "Sun", i18n is supported.
func (c *Carbon) ToShortWeekString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}

	lang := c.lang
	if lang == nil {
		return ""
	}

	lang.rw.RLock()
	defer lang.rw.RUnlock()

	if resources, ok := lang.resources["short_weeks"]; ok {
		slice := strings.Split(resources, "|")
		if len(slice) == DaysPerWeek {
			return slice[c.DayOfWeek()%DaysPerWeek]
		}
	}
	return ""
}

// ToDayDateTimeString outputs a string in "Mon, Jan 2, 2006 3:04 PM" layout.
func (c *Carbon) ToDayDateTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(DayDateTimeLayout)
}

// ToDateTimeString outputs a string in "2006-01-02 15:04:05" layout.
func (c *Carbon) ToDateTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(DateTimeLayout)
}

// ToDateTimeMilliString outputs a string in "2006-01-02 15:04:05.999" layout.
func (c *Carbon) ToDateTimeMilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(DateTimeMilliLayout)
}

// ToDateTimeMicroString outputs a string in "2006-01-02 15:04:05.999999" layout.
func (c *Carbon) ToDateTimeMicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(DateTimeMicroLayout)
}

// ToDateTimeNanoString outputs a string in "2006-01-02 15:04:05.999999999" layout.
func (c *Carbon) ToDateTimeNanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(DateTimeNanoLayout)
}

// ToShortDateTimeString outputs a string in "20060102150405" layout.
func (c *Carbon) ToShortDateTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ShortDateTimeLayout)
}

// ToShortDateTimeMilliString outputs a string in "20060102150405.999" layout.
// 输出 "20060102150405.999" 格式字符串
func (c *Carbon) ToShortDateTimeMilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ShortDateTimeMilliLayout)
}

// ToShortDateTimeMicroString outputs a string in "20060102150405.999999" layout.
func (c *Carbon) ToShortDateTimeMicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ShortDateTimeMicroLayout)
}

// ToShortDateTimeNanoString outputs a string in "20060102150405.999999999" layout.
func (c *Carbon) ToShortDateTimeNanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ShortDateTimeNanoLayout)
}

// ToDateString outputs a string in "2006-01-02" layout.
func (c *Carbon) ToDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(DateLayout)
}

// ToDateMilliString outputs a string in "2006-01-02.999" layout.
func (c *Carbon) ToDateMilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(DateMilliLayout)
}

// ToDateMicroString outputs a string in "2006-01-02.999999" layout.
func (c *Carbon) ToDateMicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(DateMicroLayout)
}

// ToDateNanoString outputs a string in "2006-01-02.999999999" layout.
func (c *Carbon) ToDateNanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(DateNanoLayout)
}

// ToShortDateString outputs a string in "20060102" layout.
func (c *Carbon) ToShortDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ShortDateLayout)
}

// ToShortDateMilliString outputs a string in "20060102.999" layout.
func (c *Carbon) ToShortDateMilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ShortDateMilliLayout)
}

// ToShortDateMicroString outputs a string in "20060102.999999" layout.
func (c *Carbon) ToShortDateMicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ShortDateMicroLayout)
}

// ToShortDateNanoString outputs a string in "20060102.999999999" layout.
func (c *Carbon) ToShortDateNanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ShortDateNanoLayout)
}

// ToTimeString outputs a string in "15:04:05" layout.
func (c *Carbon) ToTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(TimeLayout)
}

// ToTimeMilliString outputs a string in "15:04:05.999" layout.
func (c *Carbon) ToTimeMilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(TimeMilliLayout)
}

// ToTimeMicroString outputs a string in "15:04:05.999999" layout.
func (c *Carbon) ToTimeMicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(TimeMicroLayout)
}

// ToTimeNanoString outputs a string in "15:04:05.999999999" layout.
func (c *Carbon) ToTimeNanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(TimeNanoLayout)
}

// ToShortTimeString outputs a string in "150405" layout.
func (c *Carbon) ToShortTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ShortTimeLayout)
}

// ToShortTimeMilliString outputs a string in "150405.999" layout.
func (c *Carbon) ToShortTimeMilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ShortTimeMilliLayout)
}

// ToShortTimeMicroString outputs a string in "150405.999999" layout.
func (c *Carbon) ToShortTimeMicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ShortTimeMicroLayout)
}

// ToShortTimeNanoString outputs a string in "150405.999999999" layout.
func (c *Carbon) ToShortTimeNanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ShortTimeNanoLayout)
}

// ToAtomString outputs a string in "2006-01-02T15:04:05Z07:00" layout.
func (c *Carbon) ToAtomString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(AtomLayout)
}

// ToAnsicString outputs a string in "Mon Jan _2 15:04:05 2006" layout.
func (c *Carbon) ToAnsicString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ANSICLayout)
}

// ToCookieString outputs a string in "Monday, 02-Jan-2006 15:04:05 MST" layout.
func (c *Carbon) ToCookieString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(CookieLayout)
}

// ToRssString outputs a string in "Mon, 02 Jan 2006 15:04:05 -0700" format.
func (c *Carbon) ToRssString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(RssLayout)
}

// ToW3cString outputs a string in "2006-01-02T15:04:05Z07:00" layout.
func (c *Carbon) ToW3cString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(W3cLayout)
}

// ToUnixDateString outputs a string in "Mon Jan _2 15:04:05 MST 2006" layout.
func (c *Carbon) ToUnixDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(UnixDateLayout)
}

// ToRubyDateString outputs a string in "Mon Jan 02 15:04:05 -0700 2006" layout.
func (c *Carbon) ToRubyDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(RubyDateLayout)
}

// ToKitchenString outputs a string in "3:04PM" layout.
func (c *Carbon) ToKitchenString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(KitchenLayout)
}

// ToHttpString outputs a string in "Mon, 02 Jan 2006 15:04:05 GMT" layout.
func (c *Carbon) ToHttpString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(HttpLayout)
}

// ToIso8601String outputs a string in "2006-01-02T15:04:05-07:00" layout.
func (c *Carbon) ToIso8601String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ISO8601Layout)
}

// ToIso8601MilliString outputs a string in "2006-01-02T15:04:05.999-07:00" layout.
func (c *Carbon) ToIso8601MilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ISO8601MilliLayout)
}

// ToIso8601MicroString outputs a string in "2006-01-02T15:04:05.999999-07:00" layout.
func (c *Carbon) ToIso8601MicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ISO8601MicroLayout)
}

// ToIso8601NanoString outputs a string in "2006-01-02T15:04:05.999999999-07:00" layout.
func (c *Carbon) ToIso8601NanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ISO8601NanoLayout)
}

// ToIso8601ZuluString outputs a string in "2006-01-02T15:04:05Z" layout.
func (c *Carbon) ToIso8601ZuluString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ISO8601ZuluLayout)
}

// ToIso8601ZuluMilliString outputs a string in "2006-01-02T15:04:05.999Z" layout.
func (c *Carbon) ToIso8601ZuluMilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ISO8601ZuluMilliLayout)
}

// ToIso8601ZuluMicroString outputs a string in "2006-01-02T15:04:05.999999Z" layout.
func (c *Carbon) ToIso8601ZuluMicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ISO8601ZuluMicroLayout)
}

// ToIso8601ZuluNanoString outputs a string in "2006-01-02T15:04:05.999999999Z" layout.
func (c *Carbon) ToIso8601ZuluNanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(ISO8601ZuluNanoLayout)
}

// ToRfc822String outputs a string in "02 Jan 06 15:04 MST" layout.
func (c *Carbon) ToRfc822String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(RFC822Layout)
}

// ToRfc822zString outputs a string in "02 Jan 06 15:04 -0700" layout.
func (c *Carbon) ToRfc822zString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(RFC822ZLayout)
}

// ToRfc850String outputs a string in "Monday, 02-Jan-06 15:04:05 MST" layout.
func (c *Carbon) ToRfc850String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(RFC850Layout)
}

// ToRfc1036String outputs a string in "Mon, 02 Jan 06 15:04:05 -0700" layout.
func (c *Carbon) ToRfc1036String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(RFC1036Layout)
}

// ToRfc1123String outputs a string in "Mon, 02 Jan 2006 15:04:05 MST" layout.
func (c *Carbon) ToRfc1123String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(RFC1123Layout)
}

// ToRfc1123zString outputs a string in "Mon, 02 Jan 2006 15:04:05 -0700" layout.
func (c *Carbon) ToRfc1123zString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(RFC1123ZLayout)
}

// ToRfc2822String outputs a string in "Mon, 02 Jan 2006 15:04:05 -0700" layout.
func (c *Carbon) ToRfc2822String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(RFC2822Layout)
}

// ToRfc3339String outputs a string in "2006-01-02T15:04:05Z07:00" layout.
func (c *Carbon) ToRfc3339String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(RFC3339Layout)
}

// ToRfc3339MilliString outputs a string in "2006-01-02T15:04:05.999Z07:00" layout.
func (c *Carbon) ToRfc3339MilliString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(RFC3339MilliLayout)
}

// ToRfc3339MicroString outputs a string in "2006-01-02T15:04:05.999999Z07:00" layout.
func (c *Carbon) ToRfc3339MicroString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(RFC3339MicroLayout)
}

// ToRfc3339NanoString outputs a string in "2006-01-02T15:04:05.999999999Z07:00" layout.
func (c *Carbon) ToRfc3339NanoString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(RFC3339NanoLayout)
}

// ToRfc7231String outputs a string in "Mon, 02 Jan 2006 15:04:05 GMT" layout.
func (c *Carbon) ToRfc7231String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(RFC7231Layout)
}

// ToFormattedDateString outputs a string in "Jan 2, 2006" layout.
func (c *Carbon) ToFormattedDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(FormattedDateLayout)
}

// ToFormattedDayDateString outputs a string in "Mon, Jan 2, 2006" layout.
func (c *Carbon) ToFormattedDayDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	return c.StdTime().Format(FormattedDayDateLayout)
}

// Layout outputs a string by layout.
func (c *Carbon) Layout(layout string, timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	switch layout {
	case TimestampLayout:
		return strconv.FormatInt(c.Timestamp(), 10)
	case TimestampMilliLayout:
		return strconv.FormatInt(c.TimestampMilli(), 10)
	case TimestampMicroLayout:
		return strconv.FormatInt(c.TimestampMicro(), 10)
	case TimestampNanoLayout:
		return strconv.FormatInt(c.TimestampNano(), 10)
	}
	return c.StdTime().Format(layout)
}

// Format outputs a string by format.
func (c *Carbon) Format(format string, timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone...)
	}
	if c.IsInvalid() {
		return ""
	}
	buffer := &bytes.Buffer{}
	for i := 0; i < len(format); i++ {
		if layout, ok := formatMap[format[i]]; ok {
			switch format[i] {
			case 'D': // short week, such as Mon
				buffer.WriteString(c.ToShortWeekString())
			case 'F': // month, such as January
				buffer.WriteString(c.ToMonthString())
			case 'M': // short month, such as Jan
				buffer.WriteString(c.ToShortMonthString())
			case 'S': // timestamp with second, such as 1596604455
				buffer.WriteString(strconv.FormatInt(c.Timestamp(), 10))
			case 'U': // timestamp with millisecond, such as 1596604455000
				buffer.WriteString(strconv.FormatInt(c.TimestampMilli(), 10))
			case 'V': // timestamp with microsecond, such as 1596604455000000
				buffer.WriteString(strconv.FormatInt(c.TimestampMicro(), 10))
			case 'X': // timestamp with nanoseconds, such as 1596604455000000000
				buffer.WriteString(strconv.FormatInt(c.TimestampNano(), 10))
			default: // common symbols
				buffer.WriteString(c.StdTime().Format(layout))
			}
		} else {
			switch format[i] {
			case '\\': // raw output, no parse
				buffer.WriteByte(format[i+1])
				i++
				continue
			case 'W': // week number of the year, ranging from 1-52
				week := fmt.Sprintf("%d", c.WeekOfYear())
				buffer.WriteString(week)
			case 'N': // day of the week as a number, ranging from 1-7
				week := fmt.Sprintf("%d", c.DayOfWeek())
				buffer.WriteString(week)
			case 'K': // abbreviated suffix for the day of the month, such as st, nd, rd, th
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
			case 'w': // day of the week represented by the number, ranging from 0-6
				buffer.WriteString(strconv.Itoa(c.DayOfWeek() - 1))
			case 't': // number of days in the month, ranging from 28-31
				buffer.WriteString(strconv.Itoa(c.DaysInMonth()))
			case 'z': // current zone location, such as Asia/Tokyo
				buffer.WriteString(c.Timezone())
			case 'o': // current zone offset, such as 28800
				buffer.WriteString(strconv.Itoa(c.ZoneOffset()))
			case 'q': // current quarter, ranging from 1-4
				buffer.WriteString(strconv.Itoa(c.Quarter()))
			case 'c': // current century, ranging from 0-99
				buffer.WriteString(strconv.Itoa(c.Century()))
			default:
				buffer.WriteByte(format[i])
			}
		}
	}
	return buffer.String()
}
