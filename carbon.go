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

// Carbon defines a Carbon struct.
// 定义 Carbon 结构体
type Carbon struct {
	time         StdTime
	layout       string
	isNil        bool
	weekStartsAt Weekday
	weekendDays  []Weekday
	loc          *Location
	lang         *Language
	Error        error
}

// NewCarbon returns a new Carbon instance.
// 初始化 Carbon 结构体
func NewCarbon(time ...time.Time) Carbon {
	c := Carbon{}
	c.lang = NewLanguage().SetLocale(DefaultLocale)
	c.layout = DefaultLayout
	c.weekStartsAt = DefaultWeekStartsAt
	c.weekendDays = DefaultWeekendDays
	if len(time) > 0 {
		c.time = time[0]
		c.loc = time[0].Location()
		return c
	}
	c.loc, c.Error = getLocationByTimezone(DefaultTimezone)
	return c
}
