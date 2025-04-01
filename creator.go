package carbon

import (
	"time"
)

// CreateFromStdTime creates a Carbon instance from standard time.Time.
// 从标准的 time.Time 创建 Carbon 实例
func CreateFromStdTime(tt time.Time, timezone ...string) Carbon {
	c := NewCarbon(tt)
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[0])
	}
	return c
}

// CreateFromTimestamp creates a Carbon instance from a given timestamp with second.
// 从给定的秒级时间戳创建 Carbon 实例
func CreateFromTimestamp(timestamp int64, timezone ...string) Carbon {
	c := NewCarbon()
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[0])
	}
	if c.HasError() {
		return c
	}
	c.time = time.Unix(timestamp, 0)
	return c
}

// CreateFromTimestampMilli creates a Carbon instance from a given timestamp with millisecond.
// 从给定的毫秒级时间戳创建 Carbon 实例
func CreateFromTimestampMilli(timestampMilli int64, timezone ...string) Carbon {
	c := NewCarbon()
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[0])
	}
	if c.HasError() {
		return c
	}
	c.time = time.Unix(timestampMilli/1e3, (timestampMilli%1e3)*1e6)
	return c
}

// CreateFromTimestampMicro creates a Carbon instance from a given timestamp with microsecond.
// 从给定的微秒级时间戳创建 Carbon 实例
func CreateFromTimestampMicro(timestampMicro int64, timezone ...string) Carbon {
	c := NewCarbon()
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[0])
	}
	if c.HasError() {
		return c
	}
	c.time = time.Unix(timestampMicro/1e6, (timestampMicro%1e6)*1e3)
	return c
}

// CreateFromTimestampNano creates a Carbon instance from a given timestamp with nanosecond.
// 从给定的纳秒级时间戳创建 Carbon 实例
func CreateFromTimestampNano(timestampNano int64, timezone ...string) Carbon {
	c := NewCarbon()
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[0])
	}
	if c.HasError() {
		return c
	}
	c.time = time.Unix(timestampNano/1e9, timestampNano%1e9)
	return c
}

// CreateFromDateTime creates a Carbon instance from a given date and time.
// 从给定的年、月、日、时、分、秒创建 Carbon 实例
func CreateFromDateTime(year, month, day, hour, minute, second int, timezone ...string) Carbon {
	return create(year, month, day, hour, minute, second, 0, timezone...)
}

// CreateFromDateTimeMilli creates a Carbon instance from a given date, time and millisecond.
// 从给定的年、月、日、时、分、秒、毫秒创建 Carbon 实例
func CreateFromDateTimeMilli(year, month, day, hour, minute, second, millisecond int, timezone ...string) Carbon {
	return create(year, month, day, hour, minute, second, millisecond*1e6, timezone...)
}

// CreateFromDateTimeMicro creates a Carbon instance from a given date, time and microsecond.
// 从给定的年、月、日、时、分、秒、微秒创建 Carbon 实例
func CreateFromDateTimeMicro(year, month, day, hour, minute, second, microsecond int, timezone ...string) Carbon {
	return create(year, month, day, hour, minute, second, microsecond*1e3, timezone...)
}

// CreateFromDateTimeNano creates a Carbon instance from a given date, time and nanosecond.
// 从给定的年、月、日、时、分、秒、纳秒创建 Carbon 实例
func CreateFromDateTimeNano(year, month, day, hour, minute, second, nanosecond int, timezone ...string) Carbon {
	return create(year, month, day, hour, minute, second, nanosecond, timezone...)
}

// CreateFromDate creates a Carbon instance from a given date.
// 从给定的年、月、日创建 Carbon 实例
func CreateFromDate(year, month, day int, timezone ...string) Carbon {
	return create(year, month, day, 0, 0, 0, 0, timezone...)
}

// CreateFromDateMilli creates a Carbon instance from a given date and millisecond.
// 从给定的年、月、日、毫秒创建 Carbon 实例
func CreateFromDateMilli(year, month, day, millisecond int, timezone ...string) Carbon {
	return create(year, month, day, 0, 0, 0, millisecond*1e6, timezone...)
}

// CreateFromDateMicro creates a Carbon instance from a given date and microsecond.
// 从给定的年、月、日、微秒创建 Carbon 实例
func CreateFromDateMicro(year, month, day, microsecond int, timezone ...string) Carbon {
	return create(year, month, day, 0, 0, 0, microsecond*1e3, timezone...)
}

// CreateFromDateNano creates a Carbon instance from a given date and nanosecond.
// 从给定的年、月、日、纳秒创建 Carbon 实例
func CreateFromDateNano(year, month, day, nanosecond int, timezone ...string) Carbon {
	return create(year, month, day, 0, 0, 0, nanosecond, timezone...)
}

// CreateFromTime creates a Carbon instance from a given time(year, month and day are taken from the current time).
// 从给定的时、分、秒创建 Carbon 实例(年、月、日取自当前时间)
func CreateFromTime(hour, minute, second int, timezone ...string) Carbon {
	year, month, day := Now(timezone...).Date()
	return create(year, month, day, hour, minute, second, 0, timezone...)
}

// CreateFromTimeMilli creates a Carbon instance from a given time and millisecond(year, month and day are taken from the current time).
// 从给定的时、分、秒、毫秒创建 Carbon 实例(年、月、日取自当前时间)
func CreateFromTimeMilli(hour, minute, second, millisecond int, timezone ...string) Carbon {
	year, month, day := Now(timezone...).Date()
	return create(year, month, day, hour, minute, second, millisecond*1e6, timezone...)
}

// CreateFromTimeMicro creates a Carbon instance from a given time and microsecond(year, month and day are taken from the current time).
// 从给定的时、分、秒、微秒创建 Carbon 实例(年、月、日取自当前时间)
func CreateFromTimeMicro(hour, minute, second, microsecond int, timezone ...string) Carbon {
	year, month, day := Now(timezone...).Date()
	return create(year, month, day, hour, minute, second, microsecond*1e3, timezone...)
}

// CreateFromTimeNano creates a Carbon instance from a given time and nanosecond(year, month and day are taken from the current time).
// 从给定的时、分、秒、纳秒创建 Carbon 实例(年、月、日取自当前时间)
func CreateFromTimeNano(hour, minute, second, nanosecond int, timezone ...string) Carbon {
	year, month, day := Now(timezone...).Date()
	return create(year, month, day, hour, minute, second, nanosecond, timezone...)
}

// creates a Carbon instance from a given date, time and nanosecond.
// 从给定的年、月、日、时、分、秒、纳秒创建 Carbon 实例
func create(year, month, day, hour, minute, second, nanosecond int, timezone ...string) Carbon {
	c := NewCarbon()
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[0])
	}
	if c.HasError() {
		return c
	}
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, c.loc)
	return c
}
