package carbon

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

// format map
// 格式符号映射表
var formatMap = map[byte]string{
	'd': "02",      // Day:    Day of the month, 2 digits with leading zeros. Eg: 01 to 31.
	'D': "Mon",     // Day:    A textual representation of a day, three letters. Eg: Mon through Sun.
	'j': "2",       // Day:    Day of the month without leading zeros. Eg: 1 to 31.
	'l': "Monday",  // Day:    A full textual representation of the day of the week. Eg: Sunday through Saturday.
	'F': "January", // Month:  A full textual representation of a month, such as January or March. Eg: January through December.
	'm': "01",      // Month:  Numeric representation of a month, with leading zeros. Eg: 01 through 12.
	'M': "Jan",     // Month:  A short textual representation of a month, three letters. Eg: Jan through Dec.
	'n': "1",       // Month:  Numeric representation of a month, without leading zeros. Eg: 1 through 12.
	'Y': "2006",    // Year:   A full numeric representation of a year, 4 digits. Eg: 1999 or 2003.
	'y': "06",      // Year:   A two digit representation of a year. Eg: 99 or 03.
	'a': "pm",      // Time:   Lowercase morning or afternoon sign. Eg: am or pm.
	'A': "PM",      // Time:   Uppercase morning or afternoon sign. Eg: AM or PM.
	'g': "3",       // Time:   12-hour format of an hour without leading zeros. Eg: 1 through 12.
	'h': "03",      // Time:   12-hour format of an hour with leading zeros. Eg: 01 through 12.
	'H': "15",      // Time:   24-hour format of an hour with leading zeros. Eg: 00 through 23.
	'i': "04",      // Time:   Minutes with leading zeros. Eg: 00 to 59.
	's': "05",      // Time:   Seconds with leading zeros. Eg: 00 through 59.
	'O': "-0700",   // Zone:   Difference to Greenwich time (GMT) in hours. Eg: +0200.
	'P': "-07:00",  // Zone:   Difference to Greenwich time (GMT) with colon between hours and minutes. Eg: +02:00.
	'Q': "Z0700",   // Zone:   ISO8601 timezone. Eg: Z, +0200.
	'R': "Z07:00",  // Zone:   ISO8601 colon timezone. Eg: Z, +02:00.
	'Z': "MST",     // Zone:   Zone name. Eg: UTC, EST, MDT ...

	'u': "999",       // Second: Millisecond. Eg: 999.
	'v': "999999",    // Second: Microsecond. Eg: 999999.
	'x': "999999999", // Second: Nanosecond. Eg: 999999999.

	'S': TimestampLayout,      // Timestamp: Timestamp with second precision. Eg: 1699677240.
	'U': TimestampMilliLayout, // Timestamp: Timestamp with millisecond precision. Eg: 1596604455666.
	'V': TimestampMicroLayout, // Timestamp: Timestamp with microsecond precision. Eg: 1596604455666666.
	'X': TimestampNanoLayout,  // Timestamp: Timestamp with nanosecond precision. Eg: 1596604455666666666.
}

// default layouts
// 默认布局模板
var defaultLayouts = []string{
	DateTimeLayout, DateLayout, TimeLayout, DayDateTimeLayout,

	"2006-01-02 15:04:05 -0700 MST", "2006-01-02T15:04:05Z07:00", "2006-01-02T15:04:05-07:00", "2006-01-02T15:04:05-0700", "2006-01-02T15:04:05",

	ISO8601Layout, RFC1036Layout, RFC822Layout, RFC822ZLayout, RFC850Layout, RFC1123Layout, RFC1123ZLayout, RFC3339Layout, RFC7231Layout,
	KitchenLayout, CookieLayout, ANSICLayout, UnixDateLayout, RubyDateLayout,

	ShortDateTimeLayout, ShortDateLayout, ShortTimeLayout,

	DateTimeMilliLayout, DateTimeMicroLayout, DateTimeNanoLayout,
	DateMilliLayout, DateMicroLayout, DateNanoLayout,
	TimeMilliLayout, TimeMicroLayout, TimeNanoLayout,

	ShortDateTimeMilliLayout, ShortDateTimeMicroLayout, ShortDateTimeNanoLayout,
	ShortDateMilliLayout, ShortDateMicroLayout, ShortDateNanoLayout,
	ShortTimeMilliLayout, ShortTimeMicroLayout, ShortTimeNanoLayout,

	ISO8601MilliLayout, ISO8601MicroLayout, ISO8601NanoLayout,
	RFC3339MilliLayout, RFC3339MicroLayout, RFC3339NanoLayout,

	"2006",
	"2006-1-2 15:4:5 -0700 MST", "2006-1-2 3:4:5 -0700 MST",
	"2006-1", "2006-1-2", "2006-1-2 15", "2006-1-2 15:4", "2006-1-2 15:4:5", "2006-1-2 15:4:5.999999999",
	"2006.1", "2006.1.2", "2006.1.2 15", "2006.1.2 15:4", "2006.1.2 15:4:5", "2006.1.2 15:4:5.999999999",
	"2006/1", "2006/1/2", "2006/1/2 15", "2006/1/2 15:4", "2006/1/2 15:4:5", "2006/1/2 15:4:5.999999999",
	"2006-01-02 15:04:05PM MST", "2006-01-02 15:04:05.999999999PM MST", "2006-1-2 15:4:5PM MST", "2006-1-2 15:4:5.999999999PM MST",
	"2006-01-02 15:04:05 PM MST", "2006-01-02 15:04:05.999999999 PM MST", "2006-1-2 15:4:5 PM MST", "2006-1-2 15:4:5.999999999 PM MST",
	"1/2/2006", "1/2/2006 15", "1/2/2006 15:4", "1/2/2006 15:4:5", "1/2/2006 15:4:5.999999999",
	"2006-1-2 15:4:5.999999999 -0700 MST", "2006-1-2 15:04:05 -0700 MST", "2006-1-2 15:04:05.999999999 -0700 MST",
	"2006-01-02T15:04:05.999999999", "2006-1-2T3:4:5", "2006-1-2T3:4:5.999999999",
	"2006-01-02T15:04:05.999999999Z07", "2006-1-2T15:4:5Z07", "2006-1-2T15:4:5.999999999Z07",
	"2006-01-02T15:04:05.999999999Z07:00", "2006-1-2T15:4:5Z07:00", "2006-1-2T15:4:5.999999999Z07:00",
	"2006-01-02T15:04:05.999999999-07:00", "2006-1-2T15:4:5-07:00", "2006-1-2T15:4:5.999999999-07:00",
	"2006-01-02T15:04:05.999999999-0700", "2006-1-2T3:4:5-0700", "2006-1-2T3:4:5.999999999-0700",
	"20060102150405-07:00", "20060102150405.999999999-07:00",
	"20060102150405Z07:00", "20060102150405.999999999Z07:00",
}

// converts format to layout.
// format 转 layout
func format2layout(format string) string {
	buffer := &bytes.Buffer{}
	for i := 0; i < len(format); i++ {
		if layout, ok := formatMap[format[i]]; ok {
			buffer.WriteString(layout)
		} else {
			switch format[i] {
			case '\\': // raw output, no parse
				buffer.WriteByte(format[i+1])
				i++
				continue
			default:
				buffer.WriteByte(format[i])
			}
		}
	}
	return buffer.String()
}

// parses a timezone string as a time.Location instance.
// 将 时区字符串 解析成 time.Location 实例
func parseTimezone(timezone string) (loc *Location, err error) {
	if timezone == "" {
		return nil, ErrEmptyTimezone()
	}
	loc, err = time.LoadLocation(timezone)
	if err != nil {
		err = fmt.Errorf("%w: %w", ErrInvalidTimezone(timezone), err)
	}
	return
}

// parses a duration string as a time.Duration instance.
// 将 时长字符串 解析成 time.Duration 实例
func parseDuration(duration string) (dur Duration, err error) {
	if duration == "" {
		return 0, ErrEmptyDuration()
	}
	dur, err = time.ParseDuration(duration)
	if err != nil {
		err = fmt.Errorf("%w: %w", ErrInvalidDuration(duration), err)
	}
	return
}

// parses a timestamp string as a int64 format timestamp.
// 将 时间戳字符串 解析成 int64 格式时间戳
func parseTimestamp(timestamp string) (ts int64, err error) {
	ts, err = strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrInvalidTimestamp(timestamp), err)
	}
	return
}

// gets absolute value.
// 获取绝对值
func getAbsValue(value int64) int64 {
	return (value ^ value>>31) - value>>31
}
