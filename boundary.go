package carbon

// StartOfCentury returns a Carbon instance for start of the century.
func (c *Carbon) StartOfCentury() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(c.Year()/YearsPerCentury*YearsPerCentury, MinMonth, MinDay, MinHour, MinMinute, MinSecond, MinNanosecond)
}

// EndOfCentury returns a Carbon instance for end of the century.
func (c *Carbon) EndOfCentury() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(c.Year()/YearsPerCentury*YearsPerCentury+99, MaxMonth, MaxDay, MaxHour, MaxMinute, MaxSecond, MaxNanosecond)
}

// StartOfDecade returns a Carbon instance for start of the decade.
func (c *Carbon) StartOfDecade() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(c.Year()/YearsPerDecade*YearsPerDecade, MinMonth, MinDay, MinHour, MinMinute, MinSecond, MinNanosecond)
}

// EndOfDecade returns a Carbon instance for end of the decade.
func (c *Carbon) EndOfDecade() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(c.Year()/YearsPerDecade*YearsPerDecade+9, MaxMonth, MaxDay, MaxHour, MaxMinute, MaxSecond, MaxNanosecond)
}

// StartOfYear returns a Carbon instance for start of the year.
func (c *Carbon) StartOfYear() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(c.Year(), MinMonth, MinDay, MinHour, MinMinute, MinSecond, MinNanosecond)
}

// EndOfYear returns a Carbon instance for end of the year.
func (c *Carbon) EndOfYear() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(c.Year(), MaxMonth, MaxDay, MaxHour, MaxMinute, MaxSecond, MaxNanosecond)
}

// StartOfQuarter returns a Carbon instance for start of the quarter.
func (c *Carbon) StartOfQuarter() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, quarter := c.Year(), c.Quarter()
	return c.create(year, 3*quarter-2, MinDay, MinHour, MinMinute, MinSecond, MinNanosecond)
}

// EndOfQuarter returns a Carbon instance for end of the quarter.
func (c *Carbon) EndOfQuarter() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, quarter := c.Year(), c.Quarter()
	var quarterDays = [4]int{31, 30, 30, 31} // Q1, Q2, Q3, Q4
	day := quarterDays[quarter-1]
	return c.create(year, 3*quarter, day, MaxHour, MaxMinute, MaxSecond, MaxNanosecond)
}

// StartOfMonth returns a Carbon instance for start of the month.
func (c *Carbon) StartOfMonth() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _ := c.Date()
	return c.create(year, month, MinDay, MinHour, MinMinute, MinSecond, MinNanosecond)
}

// EndOfMonth returns a Carbon instance for end of the month.
func (c *Carbon) EndOfMonth() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _ := c.Date()
	return c.create(year, month+1, 0, MaxHour, MaxMinute, MaxSecond, MaxNanosecond)
}

// StartOfWeek returns a Carbon instance for start of the week.
func (c *Carbon) StartOfWeek() *Carbon {
	if c.IsInvalid() {
		return c
	}
	dayOfWeek, weekStartsAt := c.StdTime().Weekday(), c.WeekStartsAt()
	if dayOfWeek == weekStartsAt {
		return c.StartOfDay()
	}
	return c.Copy().SubDays(int(DaysPerWeek+dayOfWeek-weekStartsAt) % DaysPerWeek).StartOfDay()
}

// EndOfWeek returns a Carbon instance for end of the week.
func (c *Carbon) EndOfWeek() *Carbon {
	if c.IsInvalid() {
		return c
	}
	dayOfWeek, weekEndsAt := c.StdTime().Weekday(), c.WeekEndsAt()
	if dayOfWeek == weekEndsAt {
		return c.EndOfDay()
	}
	return c.Copy().AddDays(int(DaysPerWeek-dayOfWeek+weekEndsAt) % DaysPerWeek).EndOfDay()
}

// StartOfDay returns a Carbon instance for start of the day.
func (c *Carbon) StartOfDay() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, MinHour, MinMinute, MinSecond, MinNanosecond)
}

// EndOfDay returns a Carbon instance for end of the day.
func (c *Carbon) EndOfDay() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, MaxHour, MaxMinute, MaxSecond, MaxNanosecond)
}

// StartOfHour returns a Carbon instance for start of the hour.
func (c *Carbon) StartOfHour() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, c.Hour(), MinMinute, MinSecond, MinNanosecond)
}

// EndOfHour returns a Carbon instance for end of the hour.
func (c *Carbon) EndOfHour() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, c.Hour(), MaxMinute, MaxSecond, MaxNanosecond)
}

// StartOfMinute returns a Carbon instance for start of the minute.
func (c *Carbon) StartOfMinute() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, _ := c.DateTime()
	return c.create(year, month, day, hour, minute, MinSecond, MinNanosecond)
}

// EndOfMinute returns a Carbon instance for end of the minute.
func (c *Carbon) EndOfMinute() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, _ := c.DateTime()
	return c.create(year, month, day, hour, minute, MaxSecond, MaxNanosecond)
}

// StartOfSecond returns a Carbon instance for start of the second.
func (c *Carbon) StartOfSecond() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, MinNanosecond)
}

// EndOfSecond returns a Carbon instance for end of the second.
func (c *Carbon) EndOfSecond() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, MaxNanosecond)
}
