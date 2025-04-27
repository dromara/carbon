package carbon

import (
	"fmt"
	"time"
)

// Parse parses a standard time string as a Carbon instance.
// 将标准格式时间字符串解析成 Carbon 实例
func Parse(value string, timezone ...string) (c Carbon) {
	if value == "" {
		c.isNil = true
		return c
	}
	c = NewCarbon()
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone[0])
	}
	if c.HasError() {
		return c
	}
	switch value {
	case "now":
		return Now(c.Timezone())
	case "yesterday":
		return Yesterday(c.Timezone())
	case "tomorrow":
		return Tomorrow(c.Timezone())
	}
	for i := range defaultLayouts {
		if tt, err := time.ParseInLocation(defaultLayouts[i], value, c.loc); err == nil {
			c.time = tt
			c.layout = defaultLayouts[i]
			return c
		}
	}
	c.Error = ErrFailedParse(value)
	return c
}

// ParseByLayout parses a time string as a `Carbon` instance by a confirmed layout
// 通过一个确认的 布局模板 将时间字符串解析成 Carbon 实例
func ParseByLayout(value, layout string, timezone ...string) (c Carbon) {
	if value == "" {
		c.isNil = true
		return c
	}
	if layout == "" {
		c.Error = ErrEmptyLayout()
		return c
	}
	c = NewCarbon()
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone[0])
	}
	if c.HasError() {
		return c
	}
	if layout == TimestampLayout {
		ts, err := parseTimestamp(value)
		if err != nil {
			c.Error = err
			return c
		}
		return CreateFromTimestamp(ts, c.Timezone())
	}
	if layout == TimestampMilliLayout {
		ts, err := parseTimestamp(value)
		if err != nil {
			c.Error = err
			return c
		}
		return CreateFromTimestampMilli(ts, c.Timezone())
	}
	if layout == TimestampMicroLayout {
		ts, err := parseTimestamp(value)
		if err != nil {
			c.Error = err
			return c
		}
		return CreateFromTimestampMicro(ts, c.Timezone())
	}
	if layout == TimestampNanoLayout {
		ts, err := parseTimestamp(value)
		if err != nil {
			c.Error = err
			return c
		}
		return CreateFromTimestampNano(ts, c.Timezone())
	}
	tt, err := time.ParseInLocation(layout, value, c.loc)
	if err != nil {
		c.Error = fmt.Errorf("%w: %w", ErrMismatchedLayout(value, layout), c.Error)
		return c
	}
	c.time = tt
	c.layout = layout
	return c
}

// ParseByFormat parses a time string as a `Carbon` instance by a confirmed format.
// 通过一个确认的 格式模板 将时间字符串解析成 Carbon 实例
func ParseByFormat(value, format string, timezone ...string) (c Carbon) {
	if value == "" {
		c.isNil = true
	}
	if format == "" {
		c.Error = ErrEmptyFormat()
		return c
	}
	c = ParseByLayout(value, format2layout(format), timezone...)
	if c.HasError() {
		c.Error = fmt.Errorf("%w: %w", ErrMismatchedFormat(value, format), c.Error)
	}
	return c
}

// ParseWithLayouts parses a time string as a Carbon instance with multiple fuzzy layouts.
// 通过多个模糊的 布局模板 将时间字符串解析成 Carbon 实例
func ParseWithLayouts(value string, layouts []string, timezone ...string) (c Carbon) {
	if value == "" {
		c.isNil = true
		return c
	}
	c = NewCarbon()
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone[0])
	}
	if c.HasError() {
		return c
	}
	if len(layouts) == 0 {
		return Parse(value, timezone...)
	}
	for i := range layouts {
		if tt, err := time.ParseInLocation(layouts[i], value, c.loc); err == nil {
			c.time = tt
			c.layout = layouts[i]
			return c
		}
	}
	c.Error = ErrFailedParse(value)
	return c
}

// ParseWithFormats parses a time string as a Carbon instance with multiple fuzzy formats.
// 通过多个模糊的 格式模板 将时间字符串解析成 Carbon 实例
func ParseWithFormats(value string, formats []string, timezone ...string) (c Carbon) {
	if value == "" {
		c.isNil = true
		return c
	}
	c = NewCarbon()
	if len(timezone) > 0 {
		c.loc, c.Error = parseTimezone(timezone[0])
	}
	if c.HasError() {
		return c
	}
	if len(formats) == 0 {
		return Parse(value, timezone...)
	}
	var layouts []string
	for _, format := range formats {
		layouts = append(layouts, format2layout(format))
	}
	return ParseWithLayouts(value, layouts, timezone...)
}
