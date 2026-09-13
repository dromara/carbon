package carbon

import (
	"sync"
	"time"
)

// SetLayout sets globally default layout.
func SetLayout(layout string) *Carbon {
	c := NewCarbon().SetLayout(layout)
	if !c.HasError() {
		DefaultLayout = layout
	}
	return c
}

// SetFormat sets globally default format.
func SetFormat(format string) *Carbon {
	layout := format2layout(format)
	c := NewCarbon().SetLayout(layout)
	if !c.HasError() {
		DefaultLayout = layout
	}
	return c
}

// SetTimezone sets globally default timezone.
func SetTimezone(name string) *Carbon {
	c := NewCarbon().SetTimezone(name)
	if !c.HasError() {
		DefaultTimezone = name
	}
	return c
}

// SetLocation sets globally default location.
func SetLocation(loc *Location) *Carbon {
	c := NewCarbon().SetLocation(loc)
	if !c.HasError() {
		DefaultTimezone = loc.String()
	}
	return c
}

// SetLocale sets globally default locale.
func SetLocale(locale string) *Carbon {
	c := NewCarbon().SetLocale(locale)
	if !c.HasError() {
		DefaultLocale = locale
	}
	return c
}

// SetWeekStartsAt sets globally default start day of the week.
func SetWeekStartsAt(weekday Weekday) *Carbon {
	c := NewCarbon().SetWeekStartsAt(weekday)
	if !c.HasError() {
		DefaultWeekStartsAt = weekday
	}
	return c
}

// SetWeekendDays sets globally default weekend days of the week.
func SetWeekendDays(weekDays []Weekday) *Carbon {
	c := NewCarbon().SetWeekendDays(weekDays)
	if !c.HasError() {
		DefaultWeekendDays = weekDays
	}
	return c
}

// SetLayout sets layout.
func (c *Carbon) SetLayout(layout string) *Carbon {
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
func (c *Carbon) SetFormat(format string) *Carbon {
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
func (c *Carbon) SetTimezone(name string) *Carbon {
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
func (c *Carbon) SetLocation(loc *Location) *Carbon {
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
func (c *Carbon) SetLocale(locale string) *Carbon {
	if locale == "" {
		c.Error = ErrEmptyLocale()
		return c
	}
	if c.IsInvalid() {
		return c
	}
	c.lang = defaultLanguage(locale)
	c.Error = c.lang.Error
	return c
}

// SetWeekStartsAt sets start day of the week.
func (c *Carbon) SetWeekStartsAt(weekDay Weekday) *Carbon {
	if c.IsInvalid() {
		return c
	}
	c.weekStartsAt = weekDay
	return c
}

// SetWeekendDays sets weekend days of the week.
func (c *Carbon) SetWeekendDays(weekDays []Weekday) *Carbon {
	if c.IsInvalid() {
		return c
	}
	c.weekendDays = weekDays
	return c
}

// SetLanguage sets language.
func (c *Carbon) SetLanguage(lang *Language) *Carbon {
	if c.IsInvalid() || c.isEmpty {
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
	lang.rw.RLock()
	resources := lang.resources
	lang.rw.RUnlock()

	c.lang = &Language{
		dir:       lang.dir,
		locale:    lang.locale,
		resources: resources,
		shared:    true,
		Error:     lang.Error,
		rw:        new(sync.RWMutex),
	}
	return c
}

// SetDateTime sets year, month, day, hour, minute and second.
func (c *Carbon) SetDateTime(year, month, day, hour, minute, second int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, c.Nanosecond(), c.loc)
	return c
}

// SetDateTimeMilli sets year, month, day, hour, minute, second and millisecond.
func (c *Carbon) SetDateTimeMilli(year, month, day, hour, minute, second, millisecond int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, millisecond*1e6, c.loc)
	return c
}

// SetDateTimeMicro sets year, month, day, hour, minute, second and microsecond.
func (c *Carbon) SetDateTimeMicro(year, month, day, hour, minute, second, microsecond int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, microsecond*1e3, c.loc)
	return c
}

// SetDateTimeNano sets year, month, day, hour, minute, second and nanosecond.
func (c *Carbon) SetDateTimeNano(year, month, day, hour, minute, second, nanosecond int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, c.loc)
	return c
}

// SetDate sets year, month and day.
func (c *Carbon) SetDate(year, month, day int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	hour, minute, second := c.Time()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, c.Nanosecond(), c.loc)
	return c
}

// SetDateMilli sets year, month, day and millisecond.
func (c *Carbon) SetDateMilli(year, month, day, millisecond int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	hour, minute, second := c.Time()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, millisecond*1e6, c.loc)
	return c
}

// SetDateMicro sets year, month, day and microsecond.
func (c *Carbon) SetDateMicro(year, month, day, microsecond int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	hour, minute, second := c.Time()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, microsecond*1e3, c.loc)
	return c
}

// SetDateNano sets year, month, day and nanosecond.
func (c *Carbon) SetDateNano(year, month, day, nanosecond int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	hour, minute, second := c.Time()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, c.loc)
	return c
}

// SetTime sets hour, minute and second.
func (c *Carbon) SetTime(hour, minute, second int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, c.Nanosecond(), c.loc)
	return c
}

// SetTimeMilli sets hour, minute, second and millisecond.
func (c *Carbon) SetTimeMilli(hour, minute, second, millisecond int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, millisecond*1e6, c.loc)
	return c
}

// SetTimeMicro sets hour, minute, second and microsecond.
func (c *Carbon) SetTimeMicro(hour, minute, second, microsecond int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, microsecond*1e3, c.loc)
	return c
}

// SetTimeNano sets hour, minute, second and nanosecond.
func (c *Carbon) SetTimeNano(hour, minute, second, nanosecond int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, c.loc)
	return c
}

// SetYear sets year.
func (c *Carbon) SetYear(year int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	_, month, day, hour, minute, second := c.DateTime()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, c.Nanosecond(), c.loc)
	return c
}

// SetYearNoOverflow sets year without overflowing month.
func (c *Carbon) SetYearNoOverflow(year int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.AddYearsNoOverflow(year - c.Year())
}

// SetMonth sets month.
func (c *Carbon) SetMonth(month int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, _, day, hour, minute, second := c.DateTime()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, c.Nanosecond(), c.loc)
	return c
}

// SetMonthNoOverflow sets month without overflowing month.
func (c *Carbon) SetMonthNoOverflow(month int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.AddMonthsNoOverflow(month - c.Month())
}

// SetDay sets day.
func (c *Carbon) SetDay(day int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _, hour, minute, second := c.DateTime()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, c.Nanosecond(), c.loc)
	return c
}

// SetHour sets hour.
func (c *Carbon) SetHour(hour int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, _, minute, second := c.DateTime()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, c.Nanosecond(), c.loc)
	return c
}

// SetMinute sets minute.
func (c *Carbon) SetMinute(minute int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, _, second := c.DateTime()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, c.Nanosecond(), c.loc)
	return c
}

// SetSecond sets second.
func (c *Carbon) SetSecond(second int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, _ := c.DateTime()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, c.Nanosecond(), c.loc)
	return c
}

// SetMillisecond sets millisecond.
func (c *Carbon) SetMillisecond(millisecond int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, millisecond*1e6, c.loc)
	return c
}

// SetMicrosecond sets microsecond.
func (c *Carbon) SetMicrosecond(microsecond int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, microsecond*1e3, c.loc)
	return c
}

// SetNanosecond sets nanosecond.
func (c *Carbon) SetNanosecond(nanosecond int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, c.loc)
	return c
}
