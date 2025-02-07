package carbon

import (
	"time"
)

// SetWeekStartsAt sets start day of the week.
// 设置一周的开始日期
func (c *Carbon) SetWeekStartsAt(day string) *Carbon {
	if c.Error != nil {
		return c
	}
	if weekday, ok := weekdays[day]; ok {
		c.weekStartsAt = weekday
	}
	return c
}

// SetWeekStartsAt sets start day of the week.
// 设置一周的开始日期
func SetWeekStartsAt(day string) *Carbon {
	return NewCarbon().SetWeekStartsAt(day)
}

// SetTimezone sets timezone.
// 设置时区
func (c *Carbon) SetTimezone(name string) *Carbon {
	if c.Error != nil {
		return c
	}
	c.loc, c.Error = getLocationByTimezone(name)
	return c
}

// SetTimezone sets timezone.
// 设置时区
func SetTimezone(name string) *Carbon {
	return NewCarbon().SetTimezone(name)
}

// SetLocation sets location.
// 设置地区
func (c *Carbon) SetLocation(loc *time.Location) *Carbon {
	if c.Error != nil {
		return c
	}
	if loc == nil {
		c.Error = invalidLocationError()
	}
	c.loc = loc
	return c
}

// SetLocation sets location.
// 设置地区
func SetLocation(loc *time.Location) *Carbon {
	return NewCarbon().SetLocation(loc)
}

// SetLocale sets locale.
// 设置语言区域
func (c *Carbon) SetLocale(locale string) *Carbon {
	if c.Error != nil {
		return c
	}
	c.lang.SetLocale(locale)
	c.Error = c.lang.Error
	return c
}

// SetLocale sets locale.
// 设置语言区域
func SetLocale(locale string) *Carbon {
	c := NewCarbon()
	c.lang.SetLocale(locale)
	c.Error = c.lang.Error
	return c
}

// SetDateTime sets year, month, day, hour, minute and second.
// 设置年、月、日、时、分、秒
func (c *Carbon) SetDateTime(year, month, day, hour, minute, second int) *Carbon {
	if c.Error != nil {
		return c
	}
	return c.create(year, month, day, hour, minute, second, c.Nanosecond())
}

// SetDateTimeMilli sets year, month, day, hour, minute, second and millisecond.
// 设置年、月、日、时、分、秒、毫秒
func (c *Carbon) SetDateTimeMilli(year, month, day, hour, minute, second, millisecond int) *Carbon {
	if c.Error != nil {
		return c
	}
	return c.create(year, month, day, hour, minute, second, millisecond*1e6)
}

// SetDateTimeMicro sets year, month, day, hour, minute, second and microsecond.
// 设置年、月、日、时、分、秒、微秒
func (c *Carbon) SetDateTimeMicro(year, month, day, hour, minute, second, microsecond int) *Carbon {
	if c.Error != nil {
		return c
	}
	return c.create(year, month, day, hour, minute, second, microsecond*1e3)
}

// SetDateTimeNano sets year, month, day, hour, minute, second and nanosecond.
// 设置年、月、日、时、分、秒、纳秒
func (c *Carbon) SetDateTimeNano(year, month, day, hour, minute, second, nanosecond int) *Carbon {
	if c.Error != nil {
		return c
	}
	return c.create(year, month, day, hour, minute, second, nanosecond)
}

// SetDate sets year, month and day.
// 设置年、月、日
func (c *Carbon) SetDate(year, month, day int) *Carbon {
	if c.Error != nil {
		return c
	}
	hour, minute, second := c.Time()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond())
}

// SetDateMilli sets year, month, day and millisecond.
// 设置年、月、日、毫秒
func (c *Carbon) SetDateMilli(year, month, day, millisecond int) *Carbon {
	if c.Error != nil {
		return c
	}
	hour, minute, second := c.Time()
	return c.create(year, month, day, hour, minute, second, millisecond*1e6)
}

// SetDateMicro sets year, month, day and microsecond.
// 设置年、月、日、微秒
func (c *Carbon) SetDateMicro(year, month, day, microsecond int) *Carbon {
	if c.Error != nil {
		return c
	}
	hour, minute, second := c.Time()
	return c.create(year, month, day, hour, minute, second, microsecond*1e3)
}

// SetDateNano sets year, month, day and nanosecond.
// 设置年、月、日、纳秒
func (c *Carbon) SetDateNano(year, month, day, nanosecond int) *Carbon {
	if c.Error != nil {
		return c
	}
	hour, minute, second := c.Time()
	return c.create(year, month, day, hour, minute, second, nanosecond)
}

// SetTime sets hour, minute and second.
// 设置时、分、秒
func (c *Carbon) SetTime(hour, minute, second int) *Carbon {
	if c.Error != nil {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond())
}

// SetTimeMilli sets hour, minute, second and millisecond.
// 设置时、分、秒、毫秒
func (c *Carbon) SetTimeMilli(hour, minute, second, millisecond int) *Carbon {
	if c.Error != nil {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, hour, minute, second, millisecond*1e6)
}

// SetTimeMicro sets hour, minute, second and microsecond.
// 设置时、分、秒、微秒
func (c *Carbon) SetTimeMicro(hour, minute, second, microsecond int) *Carbon {
	if c.Error != nil {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, hour, minute, second, microsecond*1e3)
}

// SetTimeNano sets hour, minute, second and nanosecond.
// 设置时、分、秒、纳秒
func (c *Carbon) SetTimeNano(hour, minute, second, nanosecond int) *Carbon {
	if c.Error != nil {
		return c
	}
	year, month, day := c.Date()
	return c.create(year, month, day, hour, minute, second, nanosecond)
}

// SetYear sets year.
// 设置年份
func (c *Carbon) SetYear(year int) *Carbon {
	if c.Error != nil {
		return c
	}
	_, month, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, c.Nanosecond())
}

// SetYearNoOverflow sets year without overflowing month.
// 设置年份(月份不溢出)
func (c *Carbon) SetYearNoOverflow(year int) *Carbon {
	if c.Error != nil {
		return c
	}
	return c.AddYearsNoOverflow(year - c.Year())
}

// SetMonth sets month.
// 设置月份
func (c *Carbon) SetMonth(month int) *Carbon {
	if c.Error != nil {
		return c
	}
	year, _, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, 0)
}

// SetMonthNoOverflow sets month without overflowing month.
// 设置月份(月份不溢出)
func (c *Carbon) SetMonthNoOverflow(month int) *Carbon {
	if c.Error != nil {
		return c
	}
	return c.AddMonthsNoOverflow(month - c.Month())
}

// SetDay sets day.
// 设置日期
func (c *Carbon) SetDay(day int) *Carbon {
	if c.Error != nil {
		return c
	}
	year, month, _, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, 0)
}

// SetHour sets hour.
// 设置小时
func (c *Carbon) SetHour(hour int) *Carbon {
	if c.Error != nil {
		return c
	}
	year, month, day, _, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, 0)
}

// SetMinute sets minute.
// 设置分钟
func (c *Carbon) SetMinute(minute int) *Carbon {
	if c.Error != nil {
		return c
	}
	year, month, day, hour, _, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, 0)
}

// SetSecond sets second.
// 设置秒数
func (c *Carbon) SetSecond(second int) *Carbon {
	if c.Error != nil {
		return c
	}
	year, month, day, hour, minute, _ := c.DateTime()
	return c.create(year, month, day, hour, minute, second, 0)
}

// SetMillisecond sets millisecond.
// 设置毫秒
func (c *Carbon) SetMillisecond(millisecond int) *Carbon {
	if c.Error != nil {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, millisecond*1e6)
}

// SetMicrosecond sets microsecond.
// 设置微秒
func (c *Carbon) SetMicrosecond(microsecond int) *Carbon {
	if c.Error != nil {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, microsecond*1e3)
}

// SetNanosecond sets nanosecond.
// 设置纳秒
func (c *Carbon) SetNanosecond(nanosecond int) *Carbon {
	if c.Error != nil {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return c.create(year, month, day, hour, minute, second, nanosecond)
}
