package carbon

import (
	"time"
)

// SetLayout sets globally default layout.
// 设置全局默认布局模板
func SetLayout(layout string) Carbon {
	c := NewCarbon().SetLayout(layout)
	if !c.HasError() {
		DefaultLayout = layout
	}
	return c
}

// SetFormat sets globally default format.
// 设置全局默认格式模板
func SetFormat(format string) Carbon {
	layout := format2layout(format)
	c := NewCarbon().SetLayout(layout)
	if !c.HasError() {
		DefaultLayout = layout
	}
	return c
}

// SetTimezone sets globally default timezone.
// 设置全局默认时区
func SetTimezone(name string) Carbon {
	c := NewCarbon().SetTimezone(name)
	if !c.HasError() {
		DefaultTimezone = name
	}
	return c
}

// SetLocation sets globally default location.
// 设置全局默认位置
func SetLocation(loc *time.Location) Carbon {
	c := NewCarbon().SetLocation(loc)
	if !c.HasError() {
		DefaultTimezone = loc.String()
	}
	return c
}

// SetLocale sets globally default locale.
// 设置全局默认语言区域
func SetLocale(locale string) Carbon {
	c := NewCarbon().SetLocale(locale)
	if !c.HasError() {
		DefaultLocale = locale
	}
	return c
}

// SetWeekStartsAt sets globally default start day of the week.
// 设置全局默认周起始日期
func SetWeekStartsAt(weekday Weekday) Carbon {
	c := NewCarbon().SetWeekStartsAt(weekday)
	if !c.HasError() {
		DefaultWeekStartsAt = weekday
	}
	return c
}

// SetWeekendDays sets weekend days of the week.
// 设置一周周末日期
func SetWeekendDays(weekDays []Weekday) Carbon {
	c := NewCarbon().SetWeekendDays(weekDays)
	if !c.HasError() {
		DefaultWeekendDays = weekDays
	}
	return c
}

// SetLayout sets layout.
// 设置布局模板
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
// 设置格式模板
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
// 设置时区
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
// 设置位置
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
// 设置语言区域
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
// 设置周起始日期
func (c Carbon) SetWeekStartsAt(weekday Weekday) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.weekStartsAt = weekday
	return c
}

// SetWeekendDays sets weekend days of the week.
// 设置一周周末日期
func (c Carbon) SetWeekendDays(weekDays []Weekday) Carbon {
	if c.IsInvalid() {
		return c
	}
	c.weekendDays = weekDays
	return c
}

// SetLanguage sets language.
// 设置语言对象
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
// 设置年、月、日、时、分、秒
func (c Carbon) SetDateTime(year, month, day, hour, minute, second int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetDateTimeMilli sets year, month, day, hour, minute, second and millisecond.
// 设置年、月、日、时、分、秒、毫秒
func (c Carbon) SetDateTimeMilli(year, month, day, hour, minute, second, millisecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return create(year, month, day, hour, minute, second, millisecond*1e6).SetLocation(c.loc)
}

// SetDateTimeMicro sets year, month, day, hour, minute, second and microsecond.
// 设置年、月、日、时、分、秒、微秒
func (c Carbon) SetDateTimeMicro(year, month, day, hour, minute, second, microsecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return create(year, month, day, hour, minute, second, microsecond*1e3).SetLocation(c.loc)
}

// SetDateTimeNano sets year, month, day, hour, minute, second and nanosecond.
// 设置年、月、日、时、分、秒、纳秒
func (c Carbon) SetDateTimeNano(year, month, day, hour, minute, second, nanosecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return create(year, month, day, hour, minute, second, nanosecond).SetLocation(c.loc)
}

// SetDate sets year, month and day.
// 设置年、月、日
func (c Carbon) SetDate(year, month, day int) Carbon {
	if c.IsInvalid() {
		return c
	}
	hour, minute, second := c.Time()
	return create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetDateMilli sets year, month, day and millisecond.
// 设置年、月、日、毫秒
func (c Carbon) SetDateMilli(year, month, day, millisecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	hour, minute, second := c.Time()
	return create(year, month, day, hour, minute, second, millisecond*1e6).SetLocation(c.loc)
}

// SetDateMicro sets year, month, day and microsecond.
// 设置年、月、日、微秒
func (c Carbon) SetDateMicro(year, month, day, microsecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	hour, minute, second := c.Time()
	return create(year, month, day, hour, minute, second, microsecond*1e3).SetLocation(c.loc)
}

// SetDateNano sets year, month, day and nanosecond.
// 设置年、月、日、纳秒
func (c Carbon) SetDateNano(year, month, day, nanosecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	hour, minute, second := c.Time()
	return create(year, month, day, hour, minute, second, nanosecond).SetLocation(c.loc)
}

// SetTime sets hour, minute and second.
// 设置时、分、秒
func (c Carbon) SetTime(hour, minute, second int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetTimeMilli sets hour, minute, second and millisecond.
// 设置时、分、秒、毫秒
func (c Carbon) SetTimeMilli(hour, minute, second, millisecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return create(year, month, day, hour, minute, second, millisecond*1e6).SetLocation(c.loc)
}

// SetTimeMicro sets hour, minute, second and microsecond.
// 设置时、分、秒、微秒
func (c Carbon) SetTimeMicro(hour, minute, second, microsecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return create(year, month, day, hour, minute, second, microsecond*1e3).SetLocation(c.loc)
}

// SetTimeNano sets hour, minute, second and nanosecond.
// 设置、时、分、秒、纳秒
func (c Carbon) SetTimeNano(hour, minute, second, nanosecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return create(year, month, day, hour, minute, second, nanosecond).SetLocation(c.loc)
}

// SetYear sets year.
// 设置年份
func (c Carbon) SetYear(year int) Carbon {
	if c.IsInvalid() {
		return c
	}
	_, month, day, hour, minute, second := c.DateTime()
	return create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetYearNoOverflow sets year without overflowing month.
// 设置年份(月份不溢出)
func (c Carbon) SetYearNoOverflow(year int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.AddYearsNoOverflow(year - c.Year())
}

// SetMonth sets month.
// 设置月份
func (c Carbon) SetMonth(month int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, _, day, hour, minute, second := c.DateTime()
	return create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetMonthNoOverflow sets month without overflowing month.
// 设置月份(月份不溢出)
func (c Carbon) SetMonthNoOverflow(month int) Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.AddMonthsNoOverflow(month - c.Month())
}

// SetDay sets day.
// 设置日期
func (c Carbon) SetDay(day int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _, hour, minute, second := c.DateTime()
	return create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetHour sets hour.
// 设置小时
func (c Carbon) SetHour(hour int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, _, minute, second := c.DateTime()
	return create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetMinute sets minute.
// 设置分钟
func (c Carbon) SetMinute(minute int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, _, second := c.DateTime()
	return create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetSecond sets second.
// 设置秒数
func (c Carbon) SetSecond(second int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, _ := c.DateTime()
	return create(year, month, day, hour, minute, second, c.Nanosecond()).SetLocation(c.loc)
}

// SetMillisecond sets millisecond.
// 设置毫秒
func (c Carbon) SetMillisecond(millisecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return create(year, month, day, hour, minute, second, millisecond*1e6).SetLocation(c.loc)
}

// SetMicrosecond sets microsecond.
// 设置微秒
func (c Carbon) SetMicrosecond(microsecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return create(year, month, day, hour, minute, second, microsecond*1e3).SetLocation(c.loc)
}

// SetNanosecond sets nanosecond.
// 设置纳秒
func (c Carbon) SetNanosecond(nanosecond int) Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return create(year, month, day, hour, minute, second, nanosecond).SetLocation(c.loc)
}
