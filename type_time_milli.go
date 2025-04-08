package carbon

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"time"
)

// TimeMilli defines a TimeMilli struct.
// 定义 TimeMilli 结构体
type TimeMilli struct {
	Carbon
}

// NewTimeMilli returns a new TimeMilli instance.
// 初始化 TimeMilli 结构体
func NewTimeMilli(carbon Carbon) TimeMilli {
	return TimeMilli{Carbon: carbon}
}

// Scan implements driver.Scanner interface.
// 实现 driver.Scanner 接口
func (t *TimeMilli) Scan(src any) error {
	c := NewCarbon()
	switch v := src.(type) {
	case []byte:
		c = Parse(string(v), DefaultTimezone)
	case string:
		c = Parse(v, DefaultTimezone)
	case time.Time:
		c = CreateFromStdTime(v, DefaultTimezone)
	case int64:
		c = CreateFromTimestamp(v, DefaultTimezone)
	default:
		return failedScanError(v)
	}
	if c.Error == nil {
		*t = NewTimeMilli(c)
	}
	return t.Error
}

// Value implements driver.Valuer interface.
// 实现 driver.Valuer 接口
func (t TimeMilli) Value() (driver.Value, error) {
	if t.IsNil() || t.IsZero() {
		return nil, nil
	}
	if t.HasError() {
		return nil, t.Error
	}
	return t.StdTime(), nil
}

// MarshalJSON implements json.Marshal interface.
// 实现 json.Marshaler 接口
func (t TimeMilli) MarshalJSON() ([]byte, error) {
	emptyBytes := []byte(`""`)
	if t.IsNil() || t.IsZero() {
		return emptyBytes, nil
	}
	if t.HasError() {
		return emptyBytes, t.Error
	}
	return []byte(fmt.Sprintf(`"%s"`, t.ToTimeMilliString())), nil
}

// UnmarshalJSON implements json.Unmarshal interface.
// 实现 json.Unmarshaler 接口
func (t *TimeMilli) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" || value == "0" {
		return nil
	}
	c := ParseByLayout(value, TimeLayout)
	if c.Error == nil {
		*t = NewTimeMilli(c)
	}
	return c.Error
}

// String implements Stringer interface.
// 实现 Stringer 接口
func (t TimeMilli) String() string {
	if t.IsZero() || t.IsInvalid() {
		return ""
	}
	return t.ToTimeMilliString()
}

// GormDataType sets gorm data type.
// 设置 gorm 数据类型
func (t TimeMilli) GormDataType() string {
	return "time"
}
