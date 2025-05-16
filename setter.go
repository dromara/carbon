package carbon

import (
	"time"
)

// SetLayout sets globally default layout.
func SetLayout(layout string) Carbon {
	c := NewCarbon().SetLayout(layout)
	if !c.HasError() {
		DefaultLayout = layout
	}
	return c
}

// SetFormat sets globally default format.
func SetFormat(format string) Carbon {
	layout := format2layout(format)
	c := NewCarbon().SetLayout(layout)
	if !c.HasError() {
		DefaultLayout = layout
	}
	return c
}

// SetTimezone sets globally default timezone.
func SetTimezone(name string) Carbon {
	c := NewCarbon().SetTimezone(name)
	if !c.HasError() {
		DefaultTimezone = name
	}
	return c
}

// SetLocation sets globally default location.
func SetLocation(loc *time.Location) Carbon {
	c := NewCarbon().SetLocation(loc)
	if !c.HasError() {
		DefaultTimezone = loc.String()
	}
	return c
}

// SetLocale sets globally default locale.
func SetLocale(locale string) Carbon {
	c := NewCarbon().SetLocale(locale)
	if !c.HasError() {
		DefaultLocale = locale
	}
	return c
}

// SetWeekStartsAt sets globally default start day of the week.
func SetWeekStartsAt(weekday Weekday) Carbon {
	c := NewCarbon().SetWeekStartsAt(weekday)
	if !c.HasError() {
		DefaultWeekStartsAt = weekday
	}
	return c
}

// SetWeekendDays sets weekend days of the week.
func SetWeekendDays(weekDays []Weekday) Carbon {
	c := NewCarbon().SetWeekendDays(weekDays)
	if !c.HasError() {
		DefaultWeekendDays = weekDays
	}
	return c
}

// SetLayout sets layout.
func (c Carbon) SetLayout(layout string) Carbon {
	if layout == "" {
		c.Error = ErrEmptyLayout()
		return c
	}
	if c.IsInvalid() {
		return c
	}
	c.currentLayout = layout
	return c
}

// SetFormat sets format.
func (c Carbon) SetFormat(format string) Carbon {
	if format == "" {
		c.Error = ErrEmptyFormat()
		return c
	}
	if c.IsInvalid() {
		return c
	}
	c.currentLayout = format2layout(format)
	return c
}

// SetTimezone sets timezone.
func (c Carbon) SetTimezone(name string) Carbon {
	if name == "" {
		c.Error = ErrEmptyTimezone()
		return c
	}
	if c.IsInvalid() {
		return c
	}
	c.loc, c.Error = parseTimezone(name)
	return c
}

// SetLocation sets location.
func (c Carbon) SetLocation(loc *time.Location) Carbon {
	if loc == nil {
		c.Error = ErrNilLocation()
		return c
	}
	if c.IsInvalid() {
		return c
	}
	c.loc = loc
	return c
}

// SetLocale sets locale.
func (c Carbon) SetLocale(locale string) Carbon {
	if locale == "" {
		c.Error = ErrEmptyLocale()
		return c
	}
	if c.IsInvalid() {
		return c
	}
	c.lang = NewLanguage().SetLocale(locale)
	c.Error = c.lang.Error
	return c
}

// SetWeekStartsAt sets start day of the week.
func (c Carbon) SetWeekStartsAt(weekday Weekday) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.weekStartsAt = weekday
	return c
}

// SetWeekendDays sets weekend days of the week.
func (c Carbon) SetWeekendDays(weekDays []Weekday) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.weekendDays = weekDays
	return c
}

// SetLanguage sets language.
func (c Carbon) SetLanguage(lang *Language) Carbon {
	if c.IsInvalid() {
		return c
	}
	if lang == nil {
		c.Error = ErrNilLanguage()
		return c
	}
	if lang.Error != nil {
		c.Error = ErrInvalidLanguage(lang)
		return c
	}
	c.lang.dir = lang.dir
	c.lang.locale = lang.locale
	c.lang.resources = lang.resources
	c.lang.Error = lang.Error
	return c
}

// SetDateTime sets year, month, day, hour, minute and second.
func (c Carbon) SetDateTime(year, month, day, hour, minute, second int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetDateTimeMilli sets year, month, day, hour, minute, second and millisecond.
func (c Carbon) SetDateTimeMilli(year, month, day, hour, minute, second, millisecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(year, month, day, hour, minute, second, millisecond*1e6).SetLocation(c.loc)
}

// SetDateTimeMicro sets year, month, day, hour, minute, second and microsecond.
func (c Carbon) SetDateTimeMicro(year, month, day, hour, minute, second, microsecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(year, month, day, hour, minute, second, microsecond*1e3).SetLocation(c.loc)
}

// SetDateTimeNano sets year, month, day, hour, minute, second and nanosecond.
func (c Carbon) SetDateTimeNano(year, month, day, hour, minute, second, nanosecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.create(year, month, day, hour, minute, second, nanosecond).SetLocation(c.loc)
}

// SetDate sets year, month and day.
func (c Carbon) SetDate(year, month, day int) Carbon {
	if c.IsInvalid() {
		return c
	}
	hour, minute, second := c.Time()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetDateMilli sets year, month, day and millisecond.
func (c Carbon) SetDateMilli(year, month, day, millisecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	hour, minute, second := c.Time()
	return c.create(year, month, day, hour, minute, second, millisecond*1e6).SetLocation(c.loc)
}

// SetDateMicro sets year, month, day and microsecond.
func (c Carbon) SetDateMicro(year, month, day, microsecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	hour, minute, second := c.Time()
	return c.create(year, month, day, hour, minute, second, microsecond*1e3).SetLocation(c.loc)
}

// SetDateNano sets year, month, day and nanosecond.
func (c Carbon) SetDateNano(year, month, day, nanosecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	hour, minute, second := c.Time()
	return c.create(year, month, day, hour, minute, second, nanosecond).SetLocation(c.loc)
}

// SetTime sets hour, minute and second.
func (c Carbon) SetTime(hour, minute, second int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetTimeMilli sets hour, minute, second and millisecond.
func (c Carbon) SetTimeMilli(hour, minute, second, millisecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, hour, minute, second, millisecond*1e6).SetLocation(c.loc)
}

// SetTimeMicro sets hour, minute, second and microsecond.
func (c Carbon) SetTimeMicro(hour, minute, second, microsecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, hour, minute, second, microsecond*1e3).SetLocation(c.loc)
}

// SetTimeNano sets hour, minute, second and nanosecond.
func (c Carbon) SetTimeNano(hour, minute, second, nanosecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, hour, minute, second, nanosecond).SetLocation(c.loc)
}

// SetYear sets year.
func (c Carbon) SetYear(year int) Carbon {
	if c.IsInvalid() {
		return c
	}
	_, month, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetYearNoOverflow sets year without overflowing month.
func (c Carbon) SetYearNoOverflow(year int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.AddYearsNoOverflow(year - c.Year())
}

// SetMonth sets month.
func (c Carbon) SetMonth(month int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, _, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetMonthNoOverflow sets month without overflowing month.
func (c Carbon) SetMonthNoOverflow(month int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.AddMonthsNoOverflow(month - c.Month())
}

// SetDay sets day.
func (c Carbon) SetDay(day int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetHour sets hour.
func (c Carbon) SetHour(hour int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, _, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetMinute sets minute.
func (c Carbon) SetMinute(minute int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, _, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetSecond sets second.
func (c Carbon) SetSecond(second int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, _ := c.DateTime()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetMillisecond sets millisecond.
func (c Carbon) SetMillisecond(millisecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, millisecond*1e6).SetLocation(c.loc)
}

// SetMicrosecond sets microsecond.
func (c Carbon) SetMicrosecond(microsecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, microsecond*1e3).SetLocation(c.loc)
}

// SetNanosecond sets nanosecond.
func (c Carbon) SetNanosecond(nanosecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, nanosecond).SetLocation(c.loc)
}
