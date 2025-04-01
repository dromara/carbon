package carbon

import (
	"github.com/dromara/carbon/v2/calendar/julian"
	"github.com/dromara/carbon/v2/calendar/lunar"
	"github.com/dromara/carbon/v2/calendar/persian"
)

// Lunar converts Carbon instance to Lunar instance.
// 将 Carbon 实例转化为 Lunar 实例
func (c Carbon) Lunar() (l lunar.Lunar) {
	if c.IsNil() {
		return l
	}
	if c.HasError() {
		l.Error = c.Error
		return l
	}
	return lunar.FromStdTime(c.StdTime())
}

// CreateFromLunar creates a Carbon instance from Lunar date and time.
// 从 农历日期 创建 Carbon 实例
func CreateFromLunar(year, month, day int, isLeapMonth bool) Carbon {
	c := NewCarbon()
	l := lunar.NewLunar(year, month, day, isLeapMonth)
	if !l.IsValid() {
		c.Error = l.Error
		return c
	}
	return NewCarbon(l.ToGregorian(DefaultTimezone).Time)
}

// Julian converts Carbon instance to Julian instance.
// 将 Carbon 实例转化为 Julian 实例
func (c Carbon) Julian() (j julian.Julian) {
	if c.IsNil() {
		return j
	}
	if c.HasError() {
		return j
	}
	return julian.FromStdTime(c.StdTime())
}

// CreateFromJulian creates a Carbon instance from Julian Day or Modified Julian Day.
// 从 儒略日/简化儒略日 创建 Carbon 实例
func CreateFromJulian(f float64) Carbon {
	j := julian.NewJulian(f)
	return NewCarbon(j.ToGregorian(DefaultTimezone).Time)
}

// Persian converts Carbon instance to Persian instance.
// 将 Carbon 实例转化为 Persian 实例
func (c Carbon) Persian() (p persian.Persian) {
	if c.IsNil() {
		return p
	}
	if c.HasError() {
		p.Error = c.Error
		return p
	}
	return persian.FromStdTime(c.StdTime())
}

// CreateFromPersian creates a Carbon instance from Persian date and time.
// 从 波斯日期 创建 Carbon 实例
func CreateFromPersian(year, month, day int) (c Carbon) {
	p := persian.NewPersian(year, month, day)
	if !p.IsValid() {
		c.Error = p.Error
		return c
	}
	return NewCarbon(p.ToGregorian(DefaultTimezone).Time)
}
