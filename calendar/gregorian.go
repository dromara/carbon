// Package calendar is part of the carbon package.
package calendar

import (
	"time"
)

// Gregorian defines a Gregorian struct.
// 定义 Gregorian 结构体
type Gregorian struct {
	Time  time.Time
	Error error
}

// String implements Stringer interface.
// 实现 Stringer 接口
func (g Gregorian) String() string {
	if g.Error != nil {
		return ""
	}
	if g.Time.IsZero() {
		return ""
	}
	return g.Time.String()
}
