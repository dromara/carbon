package carbon

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"time"
)

// Date defines a Date struct.
// 定义 Date 结构体
type Date struct {
	Carbon
}

// NewDate returns a new Date instance.
// 初始化 Date 结构体
func NewDate(carbon Carbon) Date {
	return Date{Carbon: carbon}
}

// Scan implements driver.Scanner interface.
// 实现 driver.Scanner 接口
func (t *Date) Scan(src interface{}) error {
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
	*t = NewDate(c)
	return t.Error
}

// Value implements driver.Valuer interface.
// 实现 driver.Valuer 接口
func (t Date) Value() (driver.Value, error) {
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
func (t Date) MarshalJSON() ([]byte, error) {
	emptyBytes := []byte(`""`)
	if t.IsNil() || t.IsZero() {
		return emptyBytes, nil
	}
	if t.HasError() {
		return emptyBytes, t.Error
	}
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateString())), nil
}

// UnmarshalJSON implements json.Unmarshal interface.
// 实现 json.Unmarshaler 接口
func (t *Date) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	c := ParseByLayout(value, DateTimeLayout)
	if c.Error == nil {
		*t = NewDate(c)
	}
	return c.Error
}

// String implements Stringer interface.
// 实现 Stringer 接口
func (t Date) String() string {
	if t.IsZero() || t.IsInvalid() {
		return ""
	}
	return t.ToDateString()
}

// GormDataType sets gorm data type.
// 设置 gorm 数据类型
func (t Date) GormDataType() string {
	return "time"
}
