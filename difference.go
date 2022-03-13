package carbon

import (
	"strings"
)

// DiffInYears gets the difference in years.
// 相差多少年
func (c Carbon) DiffInYears(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return c.DiffInMonths(end) / MonthsPerYear
}

// DiffInYearsWithAbs gets the difference in years with absolute value.
// 相差多少年(绝对值)
func (c Carbon) DiffInYearsWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInYears(end))
}

// DiffInMonths gets the difference in months.
// 相差多少月
func (c Carbon) DiffInMonths(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	year, month, day := c.Date()
	endYear, endMonth, endDay := end.Date()
	dy, dm, dd := endYear-year, int(endMonth-month), endDay-day
	if dd < 0 {
		dm = dm - 1
	}
	if dy == 0 && dm == 0 {
		return int64(0)
	}
	if dy == 0 && dm != 0 && dd != 0 {
		if int(end.DiffInHoursWithAbs(c)) < c.DaysInMonth()*HoursPerDay {
			return int64(0)
		}
		return int64(dm)
	}
	return int64(dy*MonthsPerYear + dm)
}

// DiffInMonthsWithAbs gets the difference in months with absolute value.
// 相差多少月(绝对值)
func (c Carbon) DiffInMonthsWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInMonths(end))
}

// DiffInWeeks gets the difference in weeks.
// 相差多少周
func (c Carbon) DiffInWeeks(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return c.DiffInDays(end) / DaysPerWeek
}

// DiffInWeeksWithAbs gets the difference in weeks with absolute value.
// 相差多少周(绝对值)
func (c Carbon) DiffInWeeksWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInWeeks(end))
}

// DiffInDays gets the difference in days.
// 相差多少天
func (c Carbon) DiffInDays(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return c.DiffInSeconds(end) / SecondsPerDay
}

// DiffInDaysWithAbs gets the difference in days with absolute value.
// 相差多少天(绝对值)
func (c Carbon) DiffInDaysWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInDays(end))
}

// DiffInHours gets the difference in hours.
// 相差多少小时
func (c Carbon) DiffInHours(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return c.DiffInSeconds(end) / SecondsPerHour
}

// DiffInHoursWithAbs gets the difference in hours with absolute value.
// 相差多少小时(绝对值)
func (c Carbon) DiffInHoursWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInHours(end))
}

// DiffInMinutes gets the difference in minutes.
// 相差多少分钟
func (c Carbon) DiffInMinutes(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return c.DiffInSeconds(end) / SecondsPerMinute
}

// DiffInMinutesWithAbs gets the difference in minutes with absolute value.
// 相差多少分钟(绝对值)
func (c Carbon) DiffInMinutesWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInMinutes(end))
}

// DiffInSeconds gets the difference in seconds.
// 相差多少秒
func (c Carbon) DiffInSeconds(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return end.Timestamp() - c.Timestamp()
}

// DiffInSecondsWithAbs gets the difference in seconds with absolute value.
// 相差多少秒(绝对值)
func (c Carbon) DiffInSecondsWithAbs(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInSeconds(end))
}

// DiffInString gets the difference in string, i18n is supported.
// 相差字符串，支持i18n
func (c Carbon) DiffInString(carbon ...Carbon) string {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	if c.Error != nil || end.Error != nil {
		return ""
	}
	unit, value := c.diff(end)
	return c.lang.translate(unit, value)
}

// DiffInStringWithAbs gets the difference in string with absolute value, i18n is supported.
// 相差字符串，支持i18n(绝对值)
func (c Carbon) DiffInStringWithAbs(carbon ...Carbon) string {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	if c.Error != nil || end.Error != nil {
		return ""
	}
	unit, value := c.diff(end)
	return c.lang.translate(unit, getAbsValue(value))
}

// DiffForHumans gets the difference in a human-readable format, i18n is supported.
// 获取对人类友好的可读格式时间差，支持i18n
func (c Carbon) DiffForHumans(carbon ...Carbon) string {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	if c.Error != nil || end.Error != nil {
		return ""
	}
	unit, value := c.diff(end)
	translation := c.lang.translate(unit, getAbsValue(value))
	if unit == "now" {
		return translation
	}
	if c.Lt(end) && len(carbon) == 0 {
		return strings.Replace(c.lang.resources["ago"], "%s", translation, 1)
	}
	if c.Lt(end) && len(carbon) > 0 {
		return strings.Replace(c.lang.resources["before"], "%s", translation, 1)
	}
	if c.Gt(end) && len(carbon) == 0 {
		return strings.Replace(c.lang.resources["from_now"], "%s", translation, 1)
	}
	return strings.Replace(c.lang.resources["after"], "%s", translation, 1)
}

// diff gets the difference for unit and value.
// 获取相差单位和差值
func (c Carbon) diff(end Carbon) (unit string, value int64) {
	switch true {
	case c.DiffInYearsWithAbs(end) > 0:
		unit = "year"
		value = c.DiffInYears(end)
		break
	case c.DiffInMonthsWithAbs(end) > 0:
		unit = "month"
		value = c.DiffInMonths(end)
		break
	case c.DiffInWeeksWithAbs(end) > 0:
		unit = "week"
		value = c.DiffInWeeks(end)
		break
	case c.DiffInDaysWithAbs(end) > 0:
		unit = "day"
		value = c.DiffInDays(end)
		break
	case c.DiffInHoursWithAbs(end) > 0:
		unit = "hour"
		value = c.DiffInHours(end)
		break
	case c.DiffInMinutesWithAbs(end) > 0:
		unit = "minute"
		value = c.DiffInMinutes(end)
	case c.DiffInSecondsWithAbs(end) > 0:
		unit = "second"
		value = c.DiffInSeconds(end)
	case c.DiffInSecondsWithAbs(end) == 0:
		unit = "now"
		value = 0
	}
	return
}
