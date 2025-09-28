// @Package carbon
// @Description a simple, semantic and developer-friendly time package for golang
// @Source github.com/dromara/carbon
// @Document carbon.go-pkg.com
// @Developer gouguoyin
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
func NewCarbon(stdTime ...StdTime) *Carbon {
	c := new(Carbon)
	c.lang = NewLanguage().SetLocale(DefaultLocale)
	c.weekStartsAt = DefaultWeekStartsAt
	c.weekendDays = DefaultWeekendDays
	c.currentLayout = DefaultLayout
	if len(stdTime) > 0 {
		c.time = stdTime[0]
		c.loc = c.time.Location()
		return c
	}
	c.loc, c.Error = parseTimezone(DefaultTimezone)
	return c
}

// Copy returns a copy of the Carbon instance.
func (c *Carbon) Copy() *Carbon {
	if c.IsNil() {
		return nil
	}

	// Create a deep copy of weekendDays slice to avoid shared reference
	weekendDays := make([]Weekday, len(c.weekendDays))
	copy(weekendDays, c.weekendDays)

	return &Carbon{
		time:          c.time,
		weekStartsAt:  c.weekStartsAt,
		weekendDays:   weekendDays,
		loc:           c.loc,
		lang:          c.lang,
		currentLayout: c.currentLayout,
		isEmpty:       c.isEmpty,
		Error:         c.Error,
	}
}

// Sleep sleeps for the specified duration like time.Sleep.
func Sleep(d time.Duration) {
	if IsTestNow() && d > 0 {
		frozenNow.testNow = frozenNow.testNow.AddDuration(d.String())
		return
	}
	time.Sleep(d)
}
