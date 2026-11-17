package carbon

import "time"

// StartOfCentury returns a Carbon instance for start of the century.
func (c *Carbon) StartOfCentury() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.startBoundary(c.Year()/YearsPerCentury*YearsPerCentury, MinMonth, MinDay)
}

// EndOfCentury returns a Carbon instance for end of the century.
func (c *Carbon) EndOfCentury() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.endBoundary(c.Year()/YearsPerCentury*YearsPerCentury+99, MaxMonth, MaxDay)
}

// StartOfDecade returns a Carbon instance for start of the decade.
func (c *Carbon) StartOfDecade() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.startBoundary(c.Year()/YearsPerDecade*YearsPerDecade, MinMonth, MinDay)
}

// EndOfDecade returns a Carbon instance for end of the decade.
func (c *Carbon) EndOfDecade() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.endBoundary(c.Year()/YearsPerDecade*YearsPerDecade+9, MaxMonth, MaxDay)
}

// StartOfYear returns a Carbon instance for start of the year.
func (c *Carbon) StartOfYear() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.startBoundary(c.Year(), MinMonth, MinDay)
}

// EndOfYear returns a Carbon instance for end of the year.
func (c *Carbon) EndOfYear() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.endBoundary(c.Year(), MaxMonth, MaxDay)
}

// StartOfQuarter returns a Carbon instance for start of the quarter.
func (c *Carbon) StartOfQuarter() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, quarter := c.Year(), c.Quarter()
	return c.startBoundary(year, 3*quarter-2, MinDay)
}

// EndOfQuarter returns a Carbon instance for end of the quarter.
func (c *Carbon) EndOfQuarter() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, quarter := c.Year(), c.Quarter()
	var quarterDays = [4]int{31, 30, 30, 31} // Q1, Q2, Q3, Q4
	day := quarterDays[quarter-1]
	return c.endBoundary(year, 3*quarter, day)
}

// StartOfMonth returns a Carbon instance for start of the month.
func (c *Carbon) StartOfMonth() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _ := c.Date()
	return c.startBoundary(year, month, MinDay)
}

// EndOfMonth returns a Carbon instance for end of the month.
func (c *Carbon) EndOfMonth() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _ := c.Date()
	return c.endBoundary(year, month+1, 0)
}

// StartOfWeek returns a Carbon instance for start of the week.
func (c *Carbon) StartOfWeek() *Carbon {
	if c.IsInvalid() {
		return c
	}
	dayOfWeek, weekStartsAt := c.StdTime().Weekday(), c.WeekStartsAt()
	// step back on the calendar date, not the clock: going back whole days
	// from a clock time that does not exist on the target day lands a day early
	year, month, day := c.Date()
	return c.startBoundary(year, month, day-int(DaysPerWeek+dayOfWeek-weekStartsAt)%DaysPerWeek)
}

// EndOfWeek returns a Carbon instance for end of the week.
func (c *Carbon) EndOfWeek() *Carbon {
	if c.IsInvalid() {
		return c
	}
	dayOfWeek, weekEndsAt := c.StdTime().Weekday(), c.WeekEndsAt()
	year, month, day := c.Date()
	return c.endBoundary(year, month, day+int(DaysPerWeek-dayOfWeek+weekEndsAt)%DaysPerWeek)
}

// StartOfDay returns a Carbon instance for start of the day.
func (c *Carbon) StartOfDay() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.startBoundary(year, month, day)
}

// EndOfDay returns a Carbon instance for end of the day.
func (c *Carbon) EndOfDay() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.endBoundary(year, month, day)
}

// startBoundary returns the first instant of the given local date.
// time.Date may not pick it when a clock change happens at midnight:
//   - where clocks jump from 00:00 to 01:00 (Chile, Cuba and the Azores do
//     this), midnight does not exist and time.Date may resolve it to 23:00 of
//     the day before. The first instant is then the transition itself, which is
//     where the zone period holding that result ends;
//   - where clocks go back from 01:00 to 00:00, midnight happens twice and
//     time.Date may pick the second occurrence. Use the earlier one.
func (c *Carbon) startBoundary(year, month, day int) *Carbon {
	result := c.create(year, month, day, MinHour, MinMinute, MinSecond, MinNanosecond)
	if wallClock(result.time).Before(time.Date(year, time.Month(month), day, MinHour, MinMinute, MinSecond, MinNanosecond, time.UTC)) {
		if _, end := result.time.ZoneBounds(); !end.IsZero() {
			result.time = end
		}
		return result
	}
	if start, _ := result.time.ZoneBounds(); !start.IsZero() {
		_, offset := result.time.Zone()
		_, previousOffset := start.Add(-time.Nanosecond).Zone()
		alternative := result.time.Add(time.Duration(offset-previousOffset) * time.Second)
		if alternative.Before(result.time) && wallClock(alternative).Equal(wallClock(result.time)) {
			result.time = alternative
		}
	}
	return result
}

// endBoundary returns the last instant of the given local date.
// Where clocks go back from 24:00 to 23:00 (as in Chile), the last hour of the
// date happens twice and time.Date may pick the first occurrence, which is an
// hour before the date really ends. Use the later occurrence when there is one.
func (c *Carbon) endBoundary(year, month, day int) *Carbon {
	result := c.create(year, month, day, MaxHour, MaxMinute, MaxSecond, MaxNanosecond)
	if _, end := result.time.ZoneBounds(); !end.IsZero() {
		_, offset := result.time.Zone()
		_, nextOffset := end.Zone()
		alternative := result.time.Add(time.Duration(offset-nextOffset) * time.Second)
		if alternative.After(result.time) && wallClock(alternative).Equal(wallClock(result.time)) {
			result.time = alternative
		}
	}
	return result
}

// wallClock returns the local date and clock fields of t as a UTC time, so two
// instants can be compared by what a wall clock shows.
func wallClock(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
}

// StartOfHour returns a Carbon instance for start of the hour.
func (c *Carbon) StartOfHour() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.hourBoundary(year, month, day, MinMinute, MinSecond, MinNanosecond)
}

// EndOfHour returns a Carbon instance for end of the hour.
func (c *Carbon) EndOfHour() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.hourBoundary(year, month, day, MaxMinute, MaxSecond, MaxNanosecond)
}

// hourBoundary selects the receiver's occurrence when the boundary is ambiguous.
func (c *Carbon) hourBoundary(year, month, day, minute, second, nanosecond int) *Carbon {
	result := c.create(year, month, day, c.Hour(), minute, second, nanosecond)
	_, offset := c.StdTime().Zone()
	_, boundaryOffset := result.time.Zone()
	alternative := result.time.Add(time.Duration(boundaryOffset-offset) * time.Second)
	_, alternativeOffset := alternative.Zone()
	// An equal offset means the alternative has the same local date and clock
	// fields as the original boundary. Otherwise that occurrence does not exist
	// (for example, 01:00 after Lord Howe's half-hour rollback).
	if alternativeOffset == offset {
		result.time = alternative
	}
	return result
}

// StartOfMinute returns a Carbon instance for start of the minute.
func (c *Carbon) StartOfMinute() *Carbon {
	if c.IsInvalid() {
		return c
	}
	elapsed := time.Duration(c.Second())*time.Second + time.Duration(c.Nanosecond())
	result := c.Copy()
	result.time = result.time.Add(-elapsed)
	return result
}

// EndOfMinute returns a Carbon instance for end of the minute.
func (c *Carbon) EndOfMinute() *Carbon {
	if c.IsInvalid() {
		return c
	}
	elapsed := time.Duration(c.Second())*time.Second + time.Duration(c.Nanosecond())
	result := c.Copy()
	result.time = result.time.Add(time.Minute - time.Nanosecond - elapsed)
	return result
}

// StartOfSecond returns a Carbon instance for start of the second.
func (c *Carbon) StartOfSecond() *Carbon {
	if c.IsInvalid() {
		return c
	}
	result := c.Copy()
	result.time = result.time.Truncate(time.Second)
	return result
}

// EndOfSecond returns a Carbon instance for end of the second.
func (c *Carbon) EndOfSecond() *Carbon {
	if c.IsInvalid() {
		return c
	}
	result := c.Copy()
	result.time = result.time.Truncate(time.Second).Add(time.Second - time.Nanosecond)
	return result
}
