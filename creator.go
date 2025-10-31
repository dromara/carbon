package carbon

import (
	"time"
)

// CreateFromStdTime creates a Carbon instance from standard time.Time.
func CreateFromStdTime(stdTime StdTime, timezone ...string) *Carbon {
	if len(timezone) == 0 {
		return NewCarbon(stdTime)
	}
	var (
		loc *Location
		err error
	)
	if loc, err = parseTimezone(timezone...); err != nil {
		return &Carbon{Error: err}
	}
	return NewCarbon(stdTime.In(loc))
}

// CreateFromTimestamp creates a Carbon instance from a given timestamp with second precision.
func CreateFromTimestamp(timestamp int64, timezone ...string) *Carbon {
	var (
		loc *Location
		err error
	)
	if loc, err = parseTimezone(timezone...); err != nil {
		return &Carbon{Error: err}
	}
	return NewCarbon(time.Unix(timestamp, MinNanosecond).In(loc))
}

// CreateFromTimestampMilli creates a Carbon instance from a given timestamp with millisecond precision.
func CreateFromTimestampMilli(timestampMilli int64, timezone ...string) *Carbon {
	var (
		loc *Location
		err error
	)
	if loc, err = parseTimezone(timezone...); err != nil {
		return &Carbon{Error: err}
	}
	return NewCarbon(time.Unix(timestampMilli/1e3, (timestampMilli%1e3)*1e6).In(loc))
}

// CreateFromTimestampMicro creates a Carbon instance from a given timestamp with microsecond precision.
func CreateFromTimestampMicro(timestampMicro int64, timezone ...string) *Carbon {
	var (
		loc *Location
		err error
	)
	if loc, err = parseTimezone(timezone...); err != nil {
		return &Carbon{Error: err}
	}
	return NewCarbon(time.Unix(timestampMicro/1e6, (timestampMicro%1e6)*1e3).In(loc))
}

// CreateFromTimestampNano creates a Carbon instance from a given timestamp with nanosecond precision.
func CreateFromTimestampNano(timestampNano int64, timezone ...string) *Carbon {
	var (
		loc *Location
		err error
	)
	if loc, err = parseTimezone(timezone...); err != nil {
		return &Carbon{Error: err}
	}
	return NewCarbon(time.Unix(timestampNano/1e9, timestampNano%1e9).In(loc))
}

// CreateFromDateTime creates a Carbon instance from a given date and time.
func CreateFromDateTime(year, month, day, hour, minute, second int, timezone ...string) *Carbon {
	return create(year, month, day, hour, minute, second, MinNanosecond, timezone...)
}

// CreateFromDateTimeMilli creates a Carbon instance from a given date, time and millisecond.
func CreateFromDateTimeMilli(year, month, day, hour, minute, second, millisecond int, timezone ...string) *Carbon {
	return create(year, month, day, hour, minute, second, millisecond*1e6, timezone...)
}

// CreateFromDateTimeMicro creates a Carbon instance from a given date, time and microsecond.
func CreateFromDateTimeMicro(year, month, day, hour, minute, second, microsecond int, timezone ...string) *Carbon {
	return create(year, month, day, hour, minute, second, microsecond*1e3, timezone...)
}

// CreateFromDateTimeNano creates a Carbon instance from a given date, time and nanosecond.
func CreateFromDateTimeNano(year, month, day, hour, minute, second, nanosecond int, timezone ...string) *Carbon {
	return create(year, month, day, hour, minute, second, nanosecond, timezone...)
}

// CreateFromDate creates a Carbon instance from a given date.
func CreateFromDate(year, month, day int, timezone ...string) *Carbon {
	return create(year, month, day, MinHour, MinMinute, MinSecond, MinNanosecond, timezone...)
}

// CreateFromDateMilli creates a Carbon instance from a given date and millisecond.
func CreateFromDateMilli(year, month, day, millisecond int, timezone ...string) *Carbon {
	return create(year, month, day, MinHour, MinMinute, MinSecond, millisecond*1e6, timezone...)
}

// CreateFromDateMicro creates a Carbon instance from a given date and microsecond.
func CreateFromDateMicro(year, month, day, microsecond int, timezone ...string) *Carbon {
	return create(year, month, day, MinHour, MinMinute, MinSecond, microsecond*1e3, timezone...)
}

// CreateFromDateNano creates a Carbon instance from a given date and nanosecond.
func CreateFromDateNano(year, month, day, nanosecond int, timezone ...string) *Carbon {
	return create(year, month, day, MinHour, MinMinute, MinSecond, nanosecond, timezone...)
}

// CreateFromTime creates a Carbon instance from a given time(year, month and day are taken from the current time).
func CreateFromTime(hour, minute, second int, timezone ...string) *Carbon {
	year, month, day := Now(timezone...).Date()
	return create(year, month, day, hour, minute, second, MinNanosecond, timezone...)
}

// CreateFromTimeMilli creates a Carbon instance from a given time and millisecond(year, month and day are taken from the current time).
func CreateFromTimeMilli(hour, minute, second, millisecond int, timezone ...string) *Carbon {
	year, month, day := Now(timezone...).Date()
	return create(year, month, day, hour, minute, second, millisecond*1e6, timezone...)
}

// CreateFromTimeMicro creates a Carbon instance from a given time and microsecond(year, month and day are taken from the current time).
func CreateFromTimeMicro(hour, minute, second, microsecond int, timezone ...string) *Carbon {
	year, month, day := Now(timezone...).Date()
	return create(year, month, day, hour, minute, second, microsecond*1e3, timezone...)
}

// CreateFromTimeNano creates a Carbon instance from a given time and nanosecond(year, month and day are taken from the current time).
func CreateFromTimeNano(hour, minute, second, nanosecond int, timezone ...string) *Carbon {
	year, month, day := Now(timezone...).Date()
	return create(year, month, day, hour, minute, second, nanosecond, timezone...)
}

// creates a new Carbon instance from a given date, time and nanosecond.
func create(year, month, day, hour, minute, second, nanosecond int, timezone ...string) *Carbon {
	var (
		loc *Location
		err error
	)
	if loc, err = parseTimezone(timezone...); err != nil {
		return &Carbon{Error: err}
	}
	return NewCarbon(time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, loc))
}

// creates a new Carbon instance from a given date, time and nanosecond based on the existing Carbon.
func (c *Carbon) create(year, month, day, hour, minute, second, nanosecond int) *Carbon {
	return &Carbon{
		time:          time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, c.loc),
		weekStartsAt:  c.weekStartsAt,
		weekendDays:   c.weekendDays,
		loc:           c.loc,
		lang:          c.lang.Copy(),
		currentLayout: c.currentLayout,
		isEmpty:       c.isEmpty,
		Error:         c.Error,
	}
}
