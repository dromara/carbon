package carbon

import (
	"math"
	"strings"
	"time"
)

// DiffInYears gets the difference in years.
func (c *Carbon) DiffInYears(carbon ...*Carbon) int64 {
	start := c
	if start.IsInvalid() {
		return 0
	}
	var end *Carbon
	if len(carbon) > 0 {
		end = carbon[0]
	} else {
		end = Now().SetLocation(c.loc)
	}
	if end.IsInvalid() {
		return 0
	}
	dy, dm, dd := end.Year()-start.Year(), end.Month()-start.Month(), end.Day()-start.Day()
	if dm < 0 || (dm == 0 && dd < 0) {
		dy--
	}
	if dy < 0 && (dd != 0 || dm != 0) {
		dy++
	}
	return int64(dy)
}

// DiffAbsInYears gets the difference in years with absolute value.
func (c *Carbon) DiffAbsInYears(carbon ...*Carbon) int64 {
	return getAbsValue(c.DiffInYears(carbon...))
}

// DiffInMonths gets the difference in months.
func (c *Carbon) DiffInMonths(carbon ...*Carbon) int64 {
	start := c
	if start.IsInvalid() {
		return 0
	}
	var end *Carbon
	if len(carbon) > 0 {
		end = carbon[0]
	} else {
		end = Now().SetLocation(c.loc)
	}
	if end.IsInvalid() {
		return 0
	}
	if start.Month() == end.Month() && start.Year() == end.Year() {
		return 0
	}
	dd := start.DiffInDays(end)
	sign := 1
	if dd <= 0 {
		start, end = end, start
		sign = -1
	}
	months := getDiffInMonths(start, end)
	return months * int64(sign)
}

// DiffAbsInMonths gets the difference in months with absolute value.
func (c *Carbon) DiffAbsInMonths(carbon ...*Carbon) int64 {
	return getAbsValue(c.DiffInMonths(carbon...))
}

// DiffInWeeks gets the difference in weeks.
func (c *Carbon) DiffInWeeks(carbon ...*Carbon) int64 {
	start := c
	if start.IsInvalid() {
		return 0
	}
	var end *Carbon
	if len(carbon) > 0 {
		end = carbon[0]
	} else {
		end = Now().SetLocation(c.loc)
	}
	if end.IsInvalid() {
		return 0
	}
	return int64(math.Floor(float64((end.Timestamp() - start.Timestamp()) / (DaysPerWeek * HoursPerDay * SecondsPerHour))))
}

// DiffAbsInWeeks gets the difference in weeks with absolute value.
func (c *Carbon) DiffAbsInWeeks(carbon ...*Carbon) int64 {
	return getAbsValue(c.DiffInWeeks(carbon...))
}

// DiffInDays gets the difference in days.
func (c *Carbon) DiffInDays(carbon ...*Carbon) int64 {
	start := c
	if start.IsInvalid() {
		return 0
	}
	var end *Carbon
	if len(carbon) > 0 {
		end = carbon[0]
	} else {
		end = Now().SetLocation(c.loc)
	}
	if end.IsInvalid() {
		return 0
	}
	return int64(math.Floor(float64((end.Timestamp() - start.Timestamp()) / (HoursPerDay * SecondsPerHour))))
}

// DiffAbsInDays gets the difference in days with absolute value.
func (c *Carbon) DiffAbsInDays(carbon ...*Carbon) int64 {
	return getAbsValue(c.DiffInDays(carbon...))
}

// DiffInHours gets the difference in hours.
func (c *Carbon) DiffInHours(carbon ...*Carbon) int64 {
	start := c
	if start.IsInvalid() {
		return 0
	}
	var end *Carbon
	if len(carbon) > 0 {
		end = carbon[0]
	} else {
		end = Now().SetLocation(c.loc)
	}
	if end.IsInvalid() {
		return 0
	}
	return start.DiffInSeconds(end) / SecondsPerHour
}

// DiffAbsInHours gets the difference in hours with absolute value.
func (c *Carbon) DiffAbsInHours(carbon ...*Carbon) int64 {
	return getAbsValue(c.DiffInHours(carbon...))
}

// DiffInMinutes gets the difference in minutes.
func (c *Carbon) DiffInMinutes(carbon ...*Carbon) int64 {
	start := c
	if start.IsInvalid() {
		return 0
	}
	var end *Carbon
	if len(carbon) > 0 {
		end = carbon[0]
	} else {
		end = Now().SetLocation(c.loc)
	}
	if end.IsInvalid() {
		return 0
	}
	return start.DiffInSeconds(end) / SecondsPerMinute
}

// DiffAbsInMinutes gets the difference in minutes with absolute value.
func (c *Carbon) DiffAbsInMinutes(carbon ...*Carbon) int64 {
	return getAbsValue(c.DiffInMinutes(carbon...))
}

// DiffInSeconds gets the difference in seconds.
func (c *Carbon) DiffInSeconds(carbon ...*Carbon) int64 {
	start := c
	if start.IsInvalid() {
		return 0
	}
	var end *Carbon
	if len(carbon) > 0 {
		end = carbon[0]
	} else {
		end = Now().SetLocation(c.loc)
	}
	if end.IsInvalid() {
		return 0
	}
	return end.Timestamp() - start.Timestamp()
}

// DiffAbsInSeconds gets the difference in seconds with absolute value.
func (c *Carbon) DiffAbsInSeconds(carbon ...*Carbon) int64 {
	return getAbsValue(c.DiffInSeconds(carbon...))
}

// DiffInString gets the difference in string, i18n is supported.
func (c *Carbon) DiffInString(carbon ...*Carbon) string {
	start := c
	if start.IsInvalid() {
		return ""
	}
	var end *Carbon
	if len(carbon) > 0 {
		end = carbon[0]
	} else {
		end = Now().SetLocation(c.loc)
	}
	if end.IsInvalid() {
		return ""
	}
	unit, value := start.diff(end)
	return c.lang.translate(unit, value)
}

// DiffAbsInString gets the difference in string with absolute value, i18n is supported.
func (c *Carbon) DiffAbsInString(carbon ...*Carbon) string {
	start := c
	if start.IsInvalid() {
		return ""
	}
	var end *Carbon
	if len(carbon) > 0 {
		end = carbon[0]
	} else {
		end = Now().SetLocation(c.loc)
	}
	if end.IsInvalid() {
		return ""
	}
	unit, value := start.diff(end)
	return c.lang.translate(unit, getAbsValue(value))
}

// DiffInDuration gets the difference in duration.
func (c *Carbon) DiffInDuration(carbon ...*Carbon) Duration {
	start := c
	if start.IsInvalid() {
		return 0
	}
	var end *Carbon
	if len(carbon) > 0 {
		end = carbon[0]
	} else {
		end = Now().SetLocation(c.loc)
	}
	if end.IsInvalid() {
		return 0
	}
	return end.StdTime().Sub(start.StdTime())
}

// DiffAbsInDuration gets the difference in duration with absolute value.
func (c *Carbon) DiffAbsInDuration(carbon ...*Carbon) Duration {
	return c.DiffInDuration(carbon...).Abs()
}

// DiffForHumans gets the difference in a human-readable format, i18n is supported.
func (c *Carbon) DiffForHumans(carbon ...*Carbon) string {
	start := c
	if start.IsInvalid() {
		return ""
	}
	var end *Carbon
	if len(carbon) > 0 {
		end = carbon[0]
	} else {
		end = Now().SetLocation(c.loc)
	}
	if end.IsInvalid() {
		return ""
	}
	unit, value := start.diff(end)
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

// gets the difference for unit and value.
func (c *Carbon) diff(end *Carbon) (unit string, value int64) {
	switch true {
	case c.DiffAbsInYears(end) > 0:
		unit = "year"
		value = c.DiffInYears(end)
	case c.DiffAbsInMonths(end) > 0:
		unit = "month"
		value = c.DiffInMonths(end)
	case c.DiffAbsInWeeks(end) > 0:
		unit = "week"
		value = c.DiffInWeeks(end)
	case c.DiffAbsInDays(end) > 0:
		unit = "day"
		value = c.DiffInDays(end)
	case c.DiffAbsInHours(end) > 0:
		unit = "hour"
		value = c.DiffInHours(end)
	case c.DiffAbsInMinutes(end) > 0:
		unit = "minute"
		value = c.DiffInMinutes(end)
	case c.DiffAbsInSeconds(end) > 0:
		unit = "second"
		value = c.DiffInSeconds(end)
	case c.DiffAbsInSeconds(end) == 0:
		unit = "now"
		value = 0
	}
	return
}

func getDiffInMonths(start, end *Carbon) int64 {
	if start.IsInvalid() || end.IsInvalid() {
		return 0
	}
	sy, sm, d, h, i, s, ns := start.DateTimeNano()
	ey, em, _ := end.Date()
	dm := (ey-sy)*12 + (em - sm)
	if time.Date(sy, time.Month(sm+dm), d, h, i, s, ns, start.StdTime().Location()).After(end.StdTime()) {
		return int64(dm - 1)
	}
	return int64(dm)
}
