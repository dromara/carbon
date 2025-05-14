package carbon

// StartOfCentury returns a Carbon instance for start of the century.
func (c Carbon) StartOfCentury() Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(c.Year()/YearsPerCentury*YearsPerCentury, 1, 1, 0, 0, 0, 0)
}

// EndOfCentury returns a Carbon instance for end of the century.
func (c Carbon) EndOfCentury() Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(c.Year()/YearsPerCentury*YearsPerCentury+99, 12, 31, 23, 59, 59, 999999999)
}

// StartOfDecade returns a Carbon instance for start of the decade.
func (c Carbon) StartOfDecade() Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(c.Year()/YearsPerDecade*YearsPerDecade, 1, 1, 0, 0, 0, 0)
}

// EndOfDecade returns a Carbon instance for end of the decade.
func (c Carbon) EndOfDecade() Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(c.Year()/YearsPerDecade*YearsPerDecade+9, 12, 31, 23, 59, 59, 999999999)
}

// StartOfYear returns a Carbon instance for start of the year.
func (c Carbon) StartOfYear() Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(c.Year(), 1, 1, 0, 0, 0, 0)
}

// EndOfYear returns a Carbon instance for end of the year.
func (c Carbon) EndOfYear() Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(c.Year(), 12, 31, 23, 59, 59, 999999999)
}

// StartOfQuarter returns a Carbon instance for start of the quarter.
func (c Carbon) StartOfQuarter() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, quarter, day := c.Year(), c.Quarter(), 1
	return c.create(year, 3*quarter-2, day, 0, 0, 0, 0)
}

// EndOfQuarter returns a Carbon instance for end of the quarter.
func (c Carbon) EndOfQuarter() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, quarter, day := c.Year(), c.Quarter(), 30
	switch quarter {
	case 1, 4:
		day = 31
	case 2, 3:
		day = 30
	}
	return c.create(year, 3*quarter, day, 23, 59, 59, 999999999)
}

// StartOfMonth returns a Carbon instance for start of the month.
func (c Carbon) StartOfMonth() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _ := c.Date()
	return c.create(year, month, 1, 0, 0, 0, 0)
}

// EndOfMonth returns a Carbon instance for end of the month.
func (c Carbon) EndOfMonth() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _ := c.Date()
	return c.create(year, month+1, 0, 23, 59, 59, 999999999)
}

// StartOfWeek returns a Carbon instance for start of the week.
func (c Carbon) StartOfWeek() Carbon {
	if c.IsInvalid() {
		return c
	}
	dayOfWeek, weekStartsAt := c.StdTime().Weekday(), c.WeekStartsAt()
	if dayOfWeek == weekStartsAt {
		return c.StartOfDay()
	}
	return c.SubDays(int(DaysPerWeek+dayOfWeek-weekStartsAt) % DaysPerWeek).StartOfDay()
}

// EndOfWeek returns a Carbon instance for end of the week.
func (c Carbon) EndOfWeek() Carbon {
	if c.IsInvalid() {
		return c
	}
	dayOfWeek, weekEndsAt := c.StdTime().Weekday(), c.WeekEndsAt()
	if dayOfWeek == weekEndsAt {
		return c.EndOfDay()
	}
	return c.AddDays(int(DaysPerWeek-dayOfWeek+weekEndsAt) % DaysPerWeek).EndOfDay()
}

// StartOfDay returns a Carbon instance for start of the day.
func (c Carbon) StartOfDay() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, 0, 0, 0, 0)
}

// EndOfDay returns a Carbon instance for end of the day.
func (c Carbon) EndOfDay() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, 23, 59, 59, 999999999)
}

// StartOfHour returns a Carbon instance for start of the hour.
func (c Carbon) StartOfHour() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, c.Hour(), 0, 0, 0)
}

// EndOfHour returns a Carbon instance for end of the hour.
func (c Carbon) EndOfHour() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, c.Hour(), 59, 59, 999999999)
}

// StartOfMinute returns a Carbon instance for start of the minute.
func (c Carbon) StartOfMinute() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, _ := c.DateTime()
	return c.create(year, month, day, hour, minute, 0, 0)
}

// EndOfMinute returns a Carbon instance for end of the minute.
func (c Carbon) EndOfMinute() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, _ := c.DateTime()
	return c.create(year, month, day, hour, minute, 59, 999999999)
}

// StartOfSecond returns a Carbon instance for start of the second.
func (c Carbon) StartOfSecond() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, 0)
}

// EndOfSecond returns a Carbon instance for end of the second.
func (c Carbon) EndOfSecond() Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, 999999999)
}
