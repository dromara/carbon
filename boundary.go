package carbon

import (
	"time"
)

// StartOfCentury returns a Carbon instance for start of the century.
// 本世纪开始时间
func (c Carbon) StartOfCentury() Carbon {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year()/YearsPerCentury*YearsPerCentury, 1, 1, 0, 0, 0, 0, c.loc)
	return c
}

// EndOfCentury returns a Carbon instance for end of the century.
// 本世纪结束时间
func (c Carbon) EndOfCentury() Carbon {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year()/YearsPerCentury*YearsPerCentury+99, 12, 31, 23, 59, 59, 999999999, c.loc)
	return c
}

// StartOfDecade returns a Carbon instance for start of the decade.
// 本年代开始时间
func (c Carbon) StartOfDecade() Carbon {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year()/YearsPerDecade*YearsPerDecade, 1, 1, 0, 0, 0, 0, c.loc)
	return c
}

// EndOfDecade returns a Carbon instance for end of the decade.
// 本年代结束时间
func (c Carbon) EndOfDecade() Carbon {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year()/YearsPerDecade*YearsPerDecade+9, 12, 31, 23, 59, 59, 999999999, c.loc)
	return c
}

// StartOfYear returns a Carbon instance for start of the year.
// 本年开始时间
func (c Carbon) StartOfYear() Carbon {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), 1, 1, 0, 0, 0, 0, c.loc)
	return c
}

// EndOfYear returns a Carbon instance for end of the year.
// 本年结束时间
func (c Carbon) EndOfYear() Carbon {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), 12, 31, 23, 59, 59, 999999999, c.loc)
	return c
}

// StartOfQuarter returns a Carbon instance for start of the quarter.
// 本季度开始时间
func (c Carbon) StartOfQuarter() Carbon {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(3*c.Quarter()-2), 1, 0, 0, 0, 0, c.loc)
	return c
}

// EndOfQuarter returns a Carbon instance for end of the quarter.
// 本季度结束时间
func (c Carbon) EndOfQuarter() Carbon {
	if c.IsInvalid() {
		return c
	}
	quarter, day := c.Quarter(), 30
	switch quarter {
	case 1, 4:
		day = 31
	case 2, 3:
		day = 30
	}
	c.time = time.Date(c.Year(), time.Month(3*quarter), day, 23, 59, 59, 999999999, c.loc)
	return c
}

// StartOfMonth returns a Carbon instance for start of the month.
// 本月开始时间
func (c Carbon) StartOfMonth() Carbon {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), 1, 0, 0, 0, 0, c.loc)
	return c
}

// EndOfMonth returns a Carbon instance for end of the month.
// 本月结束时间
func (c Carbon) EndOfMonth() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _ := c.Date()
	c.time = time.Date(year, month+time.Month(1), 0, 23, 59, 59, 999999999, c.loc)
	return c
}

// StartOfWeek returns a Carbon instance for start of the week.
// 本周开始时间
func (c Carbon) StartOfWeek() Carbon {
	if c.IsInvalid() {
		return c
	}
	dayOfWeek, weekStartsAt := c.DayOfWeek(), int(c.weekStartsAt)
	return c.SubDays((DaysPerWeek + dayOfWeek - weekStartsAt) % DaysPerWeek).StartOfDay()
}

// EndOfWeek returns a Carbon instance for end of the week.
// 本周结束时间
func (c Carbon) EndOfWeek() Carbon {
	if c.IsInvalid() {
		return c
	}
	dayOfWeek, weekEndsAt := c.DayOfWeek(), int(c.weekStartsAt)+DaysPerWeek-1
	return c.AddDays((DaysPerWeek - dayOfWeek + weekEndsAt) % DaysPerWeek).EndOfDay()
}

// StartOfDay returns a Carbon instance for start of the day.
// 本日开始时间
func (c Carbon) StartOfDay() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	c.time = time.Date(year, month, day, 0, 0, 0, 0, c.loc)
	return c
}

// EndOfDay returns a Carbon instance for end of the day.
// 本日结束时间
func (c Carbon) EndOfDay() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	c.time = time.Date(year, month, day, 23, 59, 59, 999999999, c.loc)
	return c
}

// StartOfHour returns a Carbon instance for start of the hour.
// 小时开始时间
func (c Carbon) StartOfHour() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	c.time = time.Date(year, month, day, c.Hour(), 0, 0, 0, c.loc)
	return c
}

// EndOfHour returns a Carbon instance for end of the hour.
// 小时结束时间
func (c Carbon) EndOfHour() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	c.time = time.Date(year, month, day, c.Hour(), 59, 59, 999999999, c.loc)
	return c
}

// StartOfMinute returns a Carbon instance for start of the minute.
// 分钟开始时间
func (c Carbon) StartOfMinute() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	hour, minute, _ := c.Clock()
	c.time = time.Date(year, month, day, hour, minute, 0, 0, c.loc)
	return c
}

// EndOfMinute returns a Carbon instance for end of the minute.
// 分钟结束时间
func (c Carbon) EndOfMinute() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	hour, minute, _ := c.Clock()
	c.time = time.Date(year, month, day, hour, minute, 59, 999999999, c.loc)
	return c
}

// StartOfSecond returns a Carbon instance for start of the second.
// 秒开始时间
func (c Carbon) StartOfSecond() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	hour, minute, second := c.Clock()
	c.time = time.Date(year, month, day, hour, minute, second, 0, c.loc)
	return c
}

// EndOfSecond returns a Carbon instance for end of the second.
// 秒结束时间
func (c Carbon) EndOfSecond() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	hour, minute, second := c.Clock()
	c.time = time.Date(year, month, day, hour, minute, second, 999999999, c.loc)
	return c
}
