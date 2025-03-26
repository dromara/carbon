package carbon

import (
	"strconv"
	"time"
)

// Parse parses a standard time string as a Carbon instance.
// 将标准格式时间字符串解析成 Carbon 实例
func (c Carbon) Parse(value string, timezone ...string) Carbon {
	if len(value) == 0 {
		c.Error = invalidValueError(value)
		return c
	}
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[0])
	}
	switch value {
	case "now":
		return c.Now(timezone...)
	case "yesterday":
		return c.Yesterday(timezone...)
	case "tomorrow":
		return c.Tomorrow(timezone...)
	}
	for _, layout := range layouts {
		t, err := time.ParseInLocation(layout, value, c.loc)
		if err == nil {
			c.time = t
			c.layout = layout
			return c
		}
	}
	c.Error = invalidValueError(value)
	return c
}

// Parse parses a standard time string as a Carbon instance.
// 将标准时间字符串解析成 Carbon 实例
func Parse(value string, timezone ...string) Carbon {
	return NewCarbon().Parse(value, timezone...)
}

// ParseByFormat parses a time string as a Carbon instance by format.
// 通过格式模板将时间字符串解析成 Carbon 实例
func (c Carbon) ParseByFormat(value, format string, timezone ...string) Carbon {
	carbon := c.ParseByLayout(value, format2layout(format), timezone...)
	if carbon.Error != nil {
		carbon.Error = invalidFormatError(value, format)
	}
	return carbon
}

// ParseByFormat parses a time string as a Carbon instance by format.
// 通过格式模板将时间字符串解析成 Carbon 实例
func ParseByFormat(value, format string, timezone ...string) Carbon {
	return NewCarbon().ParseByFormat(value, format, timezone...)
}

// ParseByLayout parses a time string as a Carbon instance by layout.
// 通过布局模板将时间字符串解析成 Carbon 实例
func (c Carbon) ParseByLayout(value, layout string, timezone ...string) Carbon {
	if len(value) == 0 {
		c.Error = invalidValueError(value)
		return c
	}
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[0])
	}
	if c.Error != nil {
		return c
	}
	if len(layout) == 0 {
		layout = DefaultLayout
	}
	if layout == "timestamp" {
		timestamp, _ := strconv.ParseInt(value, 10, 64)
		return c.CreateFromTimestamp(timestamp)
	}
	if layout == "timestampMilli" {
		timestamp, _ := strconv.ParseInt(value, 10, 64)
		return c.CreateFromTimestampMilli(timestamp)
	}
	if layout == "timestampMicro" {
		timestamp, _ := strconv.ParseInt(value, 10, 64)
		return c.CreateFromTimestampMicro(timestamp)
	}
	if layout == "timestampNano" {
		timestamp, _ := strconv.ParseInt(value, 10, 64)
		return c.CreateFromTimestampNano(timestamp)
	}
	tt, err := time.ParseInLocation(layout, value, c.loc)
	if err != nil {
		c.Error = invalidLayoutError(value, layout)
		return c
	}
	c.time = tt
	c.layout = layout
	return c
}

// ParseByLayout parses a time string as a Carbon instance by layout.
// 通过布局模板将时间字符串解析成 Carbon 实例
func ParseByLayout(value, layout string, timezone ...string) Carbon {
	return NewCarbon().ParseByLayout(value, layout, timezone...)
}
