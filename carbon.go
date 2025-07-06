// @Package carbon
// @Description a simple, semantic and developer-friendly time package for golang
// @Page github.com/dromara/carbon
// @Developer gouguoyin
// @Blog www.gouguoyin.com
// @Email 245629560@qq.com

// Package carbon is a simple, semantic and developer-friendly time package for golang.
package carbon

import (
	"time"
)

type StdTime = time.Time
type Weekday = time.Weekday
type Location = time.Location
type Duration = time.Duration

// Carbon defines a Carbon struct.
type Carbon struct {
	time          StdTime
	weekStartsAt  Weekday
	weekendDays   []Weekday
	loc           *Location
	lang          *Language
	currentLayout string
	isEmpty       bool
	Error         error
}

// NewCarbon returns a new Carbon instance.
func NewCarbon(stdTime ...time.Time) Carbon {
	c := Carbon{}
	c.lang = NewLanguage().SetLocale(DefaultLocale)
	c.weekStartsAt = DefaultWeekStartsAt
	c.weekendDays = DefaultWeekendDays
	c.currentLayout = DefaultLayout
	if len(stdTime) > 0 {
		c.time = stdTime[0]
		c.loc = stdTime[0].Location()
		return c
	}
	c.loc, c.Error = parseTimezone(DefaultTimezone)
	return c
}

// Sleep sleeps for the specified duration like time.Sleep.
func (c Carbon) Sleep(d time.Duration) {
	if IsTestNow() && d > 0 {
		frozenNow.testNow = frozenNow.testNow.AddDuration(d.String())
		return
	}
	time.Sleep(d)
}
