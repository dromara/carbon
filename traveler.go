package carbon

import (
	"time"
)

// Now returns a Carbon instance for now.
func Now(timezone ...string) *Carbon {
	var (
		tz  string
		loc *Location
		err error
	)
	if len(timezone) > 0 {
		tz = timezone[0]
	} else {
		tz = DefaultTimezone
	}
	if loc, err = parseTimezone(tz); err != nil {
		return &Carbon{Error: err}
	}
	if IsTestNow() {
		frozenNow.rw.RLock()
		defer frozenNow.rw.RUnlock()
		return frozenNow.testNow.Copy().SetLocation(loc)
	}
	return CreateFromStdTime(time.Now().In(loc))
}

// Tomorrow returns a Carbon instance for tomorrow.
func Tomorrow(timezone ...string) *Carbon {
	now := Now(timezone...)
	if now.IsInvalid() {
		return now
	}
	return now.AddDay()
}

// Yesterday returns a Carbon instance for yesterday.
func Yesterday(timezone ...string) *Carbon {
	now := Now(timezone...)
	if now.IsInvalid() {
		return now
	}
	return now.SubDay()
}

// AddDuration adds duration.
func (c *Carbon) AddDuration(duration string) *Carbon {
	if c.IsInvalid() {
		return c
	}
	var (
		td  Duration
		err error
	)
	if td, err = parseDuration(duration); err != nil {
		c.Error = err
		return c
	}
	result := c.Copy()
	result.time = c.StdTime().Add(td)
	return result
}

// SubDuration subtracts duration.
func (c *Carbon) SubDuration(duration string) *Carbon {
	if c.IsInvalid() {
		return c
	}
	var (
		td  Duration
		err error
	)
	if td, err = parseDuration(duration); err != nil {
		c.Error = err
		return c
	}
	result := c.Copy()
	result.time = c.StdTime().Add(-td)
	return result
}

// AddCenturies adds some centuries.
func (c *Carbon) AddCenturies(centuries int) *Carbon {
	return c.AddYears(centuries * YearsPerCentury)
}

// AddCenturiesNoOverflow adds some centuries without overflowing month.
func (c *Carbon) AddCenturiesNoOverflow(centuries int) *Carbon {
	return c.AddYearsNoOverflow(centuries * YearsPerCentury)
}

// AddCentury adds one century.
func (c *Carbon) AddCentury() *Carbon {
	return c.AddCenturies(1)
}

// AddCenturyNoOverflow adds one century without overflowing month.
func (c *Carbon) AddCenturyNoOverflow() *Carbon {
	return c.AddCenturiesNoOverflow(1)
}

// SubCenturies subtracts some centuries.
func (c *Carbon) SubCenturies(centuries int) *Carbon {
	return c.SubYears(centuries * YearsPerCentury)
}

// SubCenturiesNoOverflow subtracts some centuries without overflowing month.
func (c *Carbon) SubCenturiesNoOverflow(centuries int) *Carbon {
	return c.SubYearsNoOverflow(centuries * YearsPerCentury)
}

// SubCentury subtracts one century.
func (c *Carbon) SubCentury() *Carbon {
	return c.SubCenturies(1)
}

// SubCenturyNoOverflow subtracts one century without overflowing month.
func (c *Carbon) SubCenturyNoOverflow() *Carbon {
	return c.SubCenturiesNoOverflow(1)
}

// AddDecades adds some decades.
func (c *Carbon) AddDecades(decades int) *Carbon {
	return c.AddYears(decades * YearsPerDecade)
}

// AddDecadesNoOverflow adds some decades without overflowing month.
func (c *Carbon) AddDecadesNoOverflow(decades int) *Carbon {
	return c.AddYearsNoOverflow(decades * YearsPerDecade)
}

// AddDecade adds one decade.
func (c *Carbon) AddDecade() *Carbon {
	return c.AddDecades(1)
}

// AddDecadeNoOverflow adds one decade without overflowing month.
func (c *Carbon) AddDecadeNoOverflow() *Carbon {
	return c.AddDecadesNoOverflow(1)
}

// SubDecades subtracts some decades.
func (c *Carbon) SubDecades(decades int) *Carbon {
	return c.SubYears(decades * YearsPerDecade)
}

// SubDecadesNoOverflow subtracts some decades without overflowing month.
func (c *Carbon) SubDecadesNoOverflow(decades int) *Carbon {
	return c.SubYearsNoOverflow(decades * YearsPerDecade)
}

// SubDecade subtracts one decade.
func (c *Carbon) SubDecade() *Carbon {
	return c.SubDecades(1)
}

// SubDecadeNoOverflow subtracts one decade without overflowing month.
func (c *Carbon) SubDecadeNoOverflow() *Carbon {
	return c.SubDecadesNoOverflow(1)
}

// AddYears adds some years.
func (c *Carbon) AddYears(years int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	result := c.Copy()
	result.time = c.StdTime().AddDate(years, 0, 0)
	return result
}

// AddYearsNoOverflow adds some years without overflowing month.
func (c *Carbon) AddYearsNoOverflow(years int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	result := c.Copy()
	nanosecond := result.Nanosecond()
	year, month, day, hour, minute, second := result.DateTime()
	// get the last day of this month after some years
	lastYear, lastMonth, lastDay := time.Date(year+years, time.Month(month+1), 0, hour, minute, second, nanosecond, c.loc).Date()
	if day > lastDay {
		day = lastDay
	}
	result.time = time.Date(lastYear, lastMonth, day, hour, minute, second, nanosecond, c.loc)
	return result
}

// AddYear adds one year.
func (c *Carbon) AddYear() *Carbon {
	return c.AddYears(1)
}

// AddYearNoOverflow adds one year without overflowing month.
func (c *Carbon) AddYearNoOverflow() *Carbon {
	return c.AddYearsNoOverflow(1)
}

// SubYears subtracts some years.
func (c *Carbon) SubYears(years int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.AddYears(-years)
}

// SubYearsNoOverflow subtracts some years without overflowing month.
func (c *Carbon) SubYearsNoOverflow(years int) *Carbon {
	return c.AddYearsNoOverflow(-years)
}

// SubYear subtracts one year.
func (c *Carbon) SubYear() *Carbon {
	return c.SubYears(1)
}

// SubYearNoOverflow subtracts one year without overflowing month.
func (c *Carbon) SubYearNoOverflow() *Carbon {
	return c.SubYearsNoOverflow(1)
}

// AddQuarters adds some quarters
func (c *Carbon) AddQuarters(quarters int) *Carbon {
	return c.AddMonths(quarters * MonthsPerQuarter)
}

// AddQuartersNoOverflow adds quarters without overflowing month.
func (c *Carbon) AddQuartersNoOverflow(quarters int) *Carbon {
	return c.AddMonthsNoOverflow(quarters * MonthsPerQuarter)
}

// AddQuarter adds one quarter
func (c *Carbon) AddQuarter() *Carbon {
	return c.AddQuarters(1)
}

// AddQuarterNoOverflow adds one quarter without overflowing month.
func (c *Carbon) AddQuarterNoOverflow() *Carbon {
	return c.AddQuartersNoOverflow(1)
}

// SubQuarters subtracts some quarters.
func (c *Carbon) SubQuarters(quarters int) *Carbon {
	return c.AddQuarters(-quarters)
}

// SubQuartersNoOverflow subtracts some quarters without overflowing month.
func (c *Carbon) SubQuartersNoOverflow(quarters int) *Carbon {
	return c.AddQuartersNoOverflow(-quarters)
}

// SubQuarter subtracts one quarter.
func (c *Carbon) SubQuarter() *Carbon {
	return c.SubQuarters(1)
}

// SubQuarterNoOverflow subtracts one quarter without overflowing month.
func (c *Carbon) SubQuarterNoOverflow() *Carbon {
	return c.SubQuartersNoOverflow(1)
}

// AddMonths adds some months.
func (c *Carbon) AddMonths(months int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	result := c.Copy()
	result.time = c.StdTime().AddDate(0, months, 0)
	return result
}

// AddMonthsNoOverflow adds some months without overflowing month.
func (c *Carbon) AddMonthsNoOverflow(months int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	result := c.Copy()
	nanosecond := result.Nanosecond()
	year, month, day, hour, minute, second := result.DateTime()
	// get the last day of this month after some months
	lastYear, lastMonth, lastDay := time.Date(year, time.Month(month+months+1), 0, hour, minute, second, nanosecond, c.loc).Date()
	if day > lastDay {
		day = lastDay
	}
	result.time = time.Date(lastYear, lastMonth, day, hour, minute, second, nanosecond, c.loc)
	return result
}

// AddMonth adds one month.
func (c *Carbon) AddMonth() *Carbon {
	return c.AddMonths(1)
}

// AddMonthNoOverflow adds one month without overflowing month.
func (c *Carbon) AddMonthNoOverflow() *Carbon {
	return c.AddMonthsNoOverflow(1)
}

// SubMonths subtracts some months.
func (c *Carbon) SubMonths(months int) *Carbon {
	return c.AddMonths(-months)
}

// SubMonthsNoOverflow subtracts some months without overflowing month.
func (c *Carbon) SubMonthsNoOverflow(months int) *Carbon {
	return c.AddMonthsNoOverflow(-months)
}

// SubMonth subtracts one month.
func (c *Carbon) SubMonth() *Carbon {
	return c.SubMonths(1)
}

// SubMonthNoOverflow subtracts one month without overflowing month.
func (c *Carbon) SubMonthNoOverflow() *Carbon {
	return c.SubMonthsNoOverflow(1)
}

// AddWeeks adds some weeks.
func (c *Carbon) AddWeeks(weeks int) *Carbon {
	return c.AddDays(weeks * DaysPerWeek)
}

// AddWeek adds one week.
func (c *Carbon) AddWeek() *Carbon {
	return c.AddWeeks(1)
}

// SubWeeks subtracts some weeks.
func (c *Carbon) SubWeeks(weeks int) *Carbon {
	return c.SubDays(weeks * DaysPerWeek)
}

// SubWeek subtracts one week.
func (c *Carbon) SubWeek() *Carbon {
	return c.SubWeeks(1)
}

// AddDays adds some days.
func (c *Carbon) AddDays(days int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	result := c.Copy()
	result.time = c.StdTime().AddDate(0, 0, days)
	return result
}

// AddDay adds one day.
func (c *Carbon) AddDay() *Carbon {
	return c.AddDays(1)
}

// SubDays subtracts some days.
func (c *Carbon) SubDays(days int) *Carbon {
	return c.AddDays(-days)
}

// SubDay subtracts one day.
func (c *Carbon) SubDay() *Carbon {
	return c.SubDays(1)
}

// AddHours adds some hours.
func (c *Carbon) AddHours(hours int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	result := c.Copy()
	result.time = c.StdTime().Add(Duration(hours) * time.Hour)
	return result
}

// AddHour adds one hour.
func (c *Carbon) AddHour() *Carbon {
	return c.AddHours(1)
}

// SubHours subtracts some hours.
func (c *Carbon) SubHours(hours int) *Carbon {
	return c.AddHours(-hours)
}

// SubHour subtracts one hour.
func (c *Carbon) SubHour() *Carbon {
	return c.SubHours(1)
}

// AddMinutes adds some minutes.
func (c *Carbon) AddMinutes(minutes int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	result := c.Copy()
	result.time = c.StdTime().Add(Duration(minutes) * time.Minute)
	return result
}

// AddMinute adds one minute.
func (c *Carbon) AddMinute() *Carbon {
	return c.AddMinutes(1)
}

// SubMinutes subtracts some minutes.
func (c *Carbon) SubMinutes(minutes int) *Carbon {
	return c.AddMinutes(-minutes)
}

// SubMinute subtracts one minute.
func (c *Carbon) SubMinute() *Carbon {
	return c.SubMinutes(1)
}

// AddSeconds adds some seconds.
func (c *Carbon) AddSeconds(seconds int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	result := c.Copy()
	result.time = c.StdTime().Add(Duration(seconds) * time.Second)
	return result
}

// AddSecond adds one second.
func (c *Carbon) AddSecond() *Carbon {
	return c.AddSeconds(1)
}

// SubSeconds subtracts some seconds.
func (c *Carbon) SubSeconds(seconds int) *Carbon {
	return c.AddSeconds(-seconds)
}

// SubSecond subtracts one second.
func (c *Carbon) SubSecond() *Carbon {
	return c.SubSeconds(1)
}

// AddMilliseconds adds some milliseconds.
func (c *Carbon) AddMilliseconds(milliseconds int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	result := c.Copy()
	result.time = c.StdTime().Add(Duration(milliseconds) * time.Millisecond)
	return result
}

// AddMillisecond adds one millisecond.
func (c *Carbon) AddMillisecond() *Carbon {
	return c.AddMilliseconds(1)
}

// SubMilliseconds subtracts some milliseconds.
func (c *Carbon) SubMilliseconds(milliseconds int) *Carbon {
	return c.AddMilliseconds(-milliseconds)
}

// SubMillisecond subtracts one millisecond.
func (c *Carbon) SubMillisecond() *Carbon {
	return c.SubMilliseconds(1)
}

// AddMicroseconds adds some microseconds.
func (c *Carbon) AddMicroseconds(microseconds int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	result := c.Copy()
	result.time = c.StdTime().Add(Duration(microseconds) * time.Microsecond)
	return result
}

// AddMicrosecond adds one microsecond.
func (c *Carbon) AddMicrosecond() *Carbon {
	return c.AddMicroseconds(1)
}

// SubMicroseconds subtracts some microseconds.
func (c *Carbon) SubMicroseconds(microseconds int) *Carbon {
	return c.AddMicroseconds(-microseconds)
}

// SubMicrosecond subtracts one microsecond.
func (c *Carbon) SubMicrosecond() *Carbon {
	return c.SubMicroseconds(1)
}

// AddNanoseconds adds some nanoseconds.
func (c *Carbon) AddNanoseconds(nanoseconds int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	result := c.Copy()
	result.time = c.StdTime().Add(Duration(nanoseconds) * time.Nanosecond)
	return result
}

// AddNanosecond adds one nanosecond.
func (c *Carbon) AddNanosecond() *Carbon {
	return c.AddNanoseconds(1)
}

// SubNanoseconds subtracts some nanoseconds.
func (c *Carbon) SubNanoseconds(nanoseconds int) *Carbon {
	return c.AddNanoseconds(-nanoseconds)
}

// SubNanosecond subtracts one nanosecond.
func (c *Carbon) SubNanosecond() *Carbon {
	return c.SubNanoseconds(1)
}
