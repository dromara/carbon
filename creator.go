package carbon

import (
	"strconv"
	"time"
)

// CreateFromTimestamp creates a Carbon instance from a given timestamp, second, millisecond, microsecond and nanosecond are supported.
// 从给定的时间戳创建 Carbon 实例，支持秒、毫秒、微秒和纳秒
func (c Carbon) CreateFromTimestamp(timestamp int64, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	ts, count := timestamp, len(strconv.FormatInt(timestamp, 10))
	if timestamp < 0 {
		count -= 1
	}
	switch count {
	case 10:
		ts = timestamp
	case 13:
		ts = timestamp / 1e3
	case 16:
		ts = timestamp / 1e6
	case 19:
		ts = timestamp / 1e9
	}
	c.time = time.Unix(ts, 0)
	return c
}

// CreateFromTimestamp creates a Carbon instance from a given timestamp.
// 从给定的时间戳创建 Carbon 实例
func CreateFromTimestamp(timestamp int64, timezone ...string) Carbon {
	return NewCarbon().CreateFromTimestamp(timestamp, timezone...)
}

// CreateFromDateTime creates a Carbon instance from a given date and time.
// 从给定的年月日时分秒创建 Carbon 实例
func (c Carbon) CreateFromDateTime(year int, month int, day int, hour int, minute int, second int, nsecond int, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, nsecond, c.loc)
	return c
}

// CreateFromDateTime creates a Carbon instance from a given date and time.
// 从给定的年月日时分秒创建 Carbon 实例
func CreateFromDateTime(year int, month int, day int, hour int, minute int, second int, nsecond int, timezone ...string) Carbon {
	return NewCarbon().CreateFromDateTime(year, month, day, hour, minute, second, nsecond, timezone...)
}

// CreateFromDate creates a Carbon instance from a given date.
// 从给定的年月日创建 Carbon 实例
func (c Carbon) CreateFromDate(year int, month int, day int, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	hour, minute, second := time.Now().In(c.loc).Clock()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, time.Now().Nanosecond(), c.loc)
	return c
}

// CreateFromDate creates a Carbon instance from a given date.
// 从给定的年月日创建 Carbon 实例
func CreateFromDate(year int, month int, day int, timezone ...string) Carbon {
	return NewCarbon().CreateFromDate(year, month, day, timezone...)
}

// CreateFromTime creates a Carbon instance from a given time.
// 从给定的时分秒创建 Carbon 实例
func (c Carbon) CreateFromTime(hour int, minute int, second int, nsecond int, timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	year, month, day := time.Now().In(c.loc).Date()
	c.time = time.Date(year, month, day, hour, minute, second, nsecond, c.loc)
	return c
}

// CreateFromTime creates a Carbon instance from a given time.
// 从给定的时分秒创建 Carbon 实例
func CreateFromTime(hour int, minute int, second int, nsecond int, timezone ...string) Carbon {
	return NewCarbon().CreateFromTime(hour, minute, second, nsecond, timezone...)
}
