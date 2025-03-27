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

// Carbon defines a Carbon struct.
// 定义 Carbon 结构体
type Carbon struct {
	time         time.Time
	layout       string
	weekStartsAt time.Weekday
	loc          *time.Location
	lang         *Language
	Error        error
}

// NewCarbon returns a new Carbon instance.
// 初始化 Carbon 结构体
func NewCarbon(time ...time.Time) Carbon {
	c := Carbon{lang: NewLanguage()}
	c.loc, c.Error = getLocationByTimezone(DefaultTimezone)
	if weekday, ok := weekdays[DefaultWeekStartsAt]; ok {
		c.weekStartsAt = weekday
	}
	if len(time) > 0 {
		c.time = time[0]
		c.loc = time[0].Location()
	}
	c.layout = DefaultLayout
	return c
}
