package carbon

import (
	"strings"
	"time"
)

// DiffInYears gets the difference in years.
func (c *Carbon) DiffInYears(carbon ...*Carbon) int64 {
	if c.IsInvalid() {
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
	sign := int64(1)
	s, e := c, end
	if e.Lt(s) {
		s, e = e, s
		sign = -1
	}

	years := e.Year() - s.Year()
	if s.StdTime().AddDate(years, 0, 0).After(e.StdTime()) {
		years--
	}
	return int64(years) * sign
}

// DiffAbsInYears gets the difference in years with absolute value.
func (c *Carbon) DiffAbsInYears(carbon ...*Carbon) int64 {
	return getAbsValue(c.DiffInYears(carbon...))
}

// DiffInMonths gets the difference in months.
func (c *Carbon) DiffInMonths(carbon ...*Carbon) int64 {
	if c.IsInvalid() {
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
	sign := int64(1)
	s, e := c, end
	if e.Lt(s) {
		s, e = e, s
		sign = -1
	}
	months := (e.Year()-s.Year())*12 + (e.Month() - s.Month())
	if s.StdTime().AddDate(0, months, 0).After(e.StdTime()) {
		months--
	}
	return int64(months) * sign
}

// DiffAbsInMonths gets the difference in months with absolute value.
func (c *Carbon) DiffAbsInMonths(carbon ...*Carbon) int64 {
	return getAbsValue(c.DiffInMonths(carbon...))
}

// DiffInWeeks gets the difference in weeks.
func (c *Carbon) DiffInWeeks(carbon ...*Carbon) int64 {
	if c.IsInvalid() {
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
	return (end.Timestamp() - c.Timestamp()) / SecondsPerWeek
}

// DiffAbsInWeeks gets the difference in weeks with absolute value.
func (c *Carbon) DiffAbsInWeeks(carbon ...*Carbon) int64 {
	return getAbsValue(c.DiffInWeeks(carbon...))
}

// DiffInDays gets the difference in days.
func (c *Carbon) DiffInDays(carbon ...*Carbon) int64 {
	if c.IsInvalid() {
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
	return (end.Timestamp() - c.Timestamp()) / SecondsPerDay
}

// DiffAbsInDays gets the difference in days with absolute value.
func (c *Carbon) DiffAbsInDays(carbon ...*Carbon) int64 {
	return getAbsValue(c.DiffInDays(carbon...))
}

// DiffInHours gets the difference in hours.
func (c *Carbon) DiffInHours(carbon ...*Carbon) int64 {
	if c.IsInvalid() {
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
	return c.DiffInSeconds(end) / SecondsPerHour
}

// DiffAbsInHours gets the difference in hours with absolute value.
func (c *Carbon) DiffAbsInHours(carbon ...*Carbon) int64 {
	return getAbsValue(c.DiffInHours(carbon...))
}

// DiffInMinutes gets the difference in minutes.
func (c *Carbon) DiffInMinutes(carbon ...*Carbon) int64 {
	if c.IsInvalid() {
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
	return c.DiffInSeconds(end) / SecondsPerMinute
}

// DiffAbsInMinutes gets the difference in minutes with absolute value.
func (c *Carbon) DiffAbsInMinutes(carbon ...*Carbon) int64 {
	return getAbsValue(c.DiffInMinutes(carbon...))
}

// DiffInSeconds gets the difference in seconds.
func (c *Carbon) DiffInSeconds(carbon ...*Carbon) int64 {
	if c.IsInvalid() {
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
	return end.Timestamp() - c.Timestamp()
}

// DiffAbsInSeconds gets the difference in seconds with absolute value.
func (c *Carbon) DiffAbsInSeconds(carbon ...*Carbon) int64 {
	return getAbsValue(c.DiffInSeconds(carbon...))
}

// DiffInString gets the difference in string, i18n is supported.
func (c *Carbon) DiffInString(carbon ...*Carbon) string {
	if c.IsInvalid() || c.lang == nil {
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
	unit, value := c.diff(end)
	return c.lang.translate(unit, value)
}

// DiffAbsInString gets the difference in string with absolute value, i18n is supported.
func (c *Carbon) DiffAbsInString(carbon ...*Carbon) string {
	if c.IsInvalid() || c.lang == nil {
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
	unit, value := c.diff(end)
	return c.lang.translate(unit, getAbsValue(value))
}

// DiffInDuration gets the difference in duration.
func (c *Carbon) DiffInDuration(carbon ...*Carbon) Duration {
	if c.IsInvalid() {
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
	return end.StdTime().Sub(c.StdTime())
}

// DiffAbsInDuration gets the difference in duration with absolute value.
func (c *Carbon) DiffAbsInDuration(carbon ...*Carbon) Duration {
	d := c.DiffInDuration(carbon...)
	switch {
	case d >= 0:
		return d
	case d == minDuration:
		return maxDuration
	default:
		return -d
	}
}

// DiffForHumans gets the difference in a human-readable format, i18n is supported.
func (c *Carbon) DiffForHumans(carbon ...*Carbon) string {
	if c.IsInvalid() || c.lang == nil {
		return ""
	}
	end := func() *Carbon {
		if len(carbon) > 0 {
			return carbon[0]
		}
		return Now().SetLocation(c.loc)
	}()
	if end.IsInvalid() {
		return ""
	}

	unit, value := c.diff(end)
	translation := c.lang.translate(unit, getAbsValue(value))
	if unit == "now" {
		return translation
	}

	// Concurrent-safe access to language resources
	c.lang.rw.RLock()
	resources := c.lang.resources
	ago := resources["ago"]
	before := resources["before"]
	fromNow := resources["from_now"]
	after := resources["after"]
	c.lang.rw.RUnlock()

	isBefore := value > 0
	if isBefore && len(carbon) == 0 {
		return strings.Replace(ago, "%s", translation, 1)
	}
	if isBefore && len(carbon) > 0 {
		return strings.Replace(before, "%s", translation, 1)
	}
	if !isBefore && len(carbon) == 0 {
		return strings.Replace(fromNow, "%s", translation, 1)
	}
	return strings.Replace(after, "%s", translation, 1)
}

// gets the difference for unit and value.
func (c *Carbon) diff(end *Carbon) (unit string, value int64) {
	// Years
	diffYears := c.DiffInYears(end)
	if getAbsValue(diffYears) > 0 {
		return "year", diffYears
	}

	// Months
	diffMonths := c.DiffInMonths(end)
	if getAbsValue(diffMonths) > 0 {
		return "month", diffMonths
	}

	// Weeks
	diffWeeks := c.DiffInWeeks(end)
	if getAbsValue(diffWeeks) > 0 {
		return "week", diffWeeks
	}

	// Days
	diffDays := c.DiffInDays(end)
	if getAbsValue(diffDays) > 0 {
		return "day", diffDays
	}

	// Hours
	diffHours := c.DiffInHours(end)
	if getAbsValue(diffHours) > 0 {
		return "hour", diffHours
	}

	// Minutes
	diffMinutes := c.DiffInMinutes(end)
	if getAbsValue(diffMinutes) > 0 {
		return "minute", diffMinutes
	}

	// Seconds
	diffSeconds := c.DiffInSeconds(end)
	if getAbsValue(diffSeconds) > 0 {
		return "second", diffSeconds
	}

	return "now", 0
}

// gets the difference in months.
// Nil and invalid inputs return 0 to match existing tests.
func getDiffInMonths(start, end *Carbon) int64 {
	if start == nil || end == nil {
		return 0
	}
	if start.IsInvalid() || end.IsInvalid() {
		return 0
	}
	sy, sm, d, h, i, s, ns := start.DateTimeNano()
	ey, em, _ := end.Date()
	dm := (ey-sy)*12 + (em - sm)
	loc := start.StdTime().Location()
	if time.Date(sy, time.Month(sm+dm), d, h, i, s, ns, loc).After(end.StdTime()) {
		return int64(dm - 1)
	}
	return int64(dm)
}
