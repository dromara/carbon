package carbon

import (
	"fmt"
	"time"
)

type ToDateTimeString time.Time
type ToDateString time.Time
type ToTimeString time.Time
type ToTimestamp time.Time

func (t ToDateTimeString) MarshalJSON() ([]byte, error) {
	format := ""
	if !time.Time(t).IsZero() {
		format = time.Time(t).Format("2006-01-02 15:04:05")
	}
	formatted := fmt.Sprintf("\"%s\"", format)
	return []byte(formatted), nil
}

func (t ToDateString) MarshalJSON() ([]byte, error) {
	format := ""
	if !time.Time(t).IsZero() {
		format = time.Time(t).Format("2006-01-02")
	}
	formatted := fmt.Sprintf("\"%s\"", format)
	return []byte(formatted), nil
}

func (t ToTimeString) MarshalJSON() ([]byte, error) {
	format := ""
	if !time.Time(t).IsZero() {
		format = time.Time(t).Format("15:04:05")
	}
	formatted := fmt.Sprintf("\"%s\"", format)
	return []byte(formatted), nil
}

func (t ToTimestamp) MarshalJSON() ([]byte, error) {
	timestamp := int64(0)
	if !time.Time(t).IsZero() {
		timestamp = time.Time(t).Unix()
	}
	formatted := fmt.Sprintf("%d", timestamp)
	return []byte(formatted), nil
}
