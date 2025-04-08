package carbon

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"time"
)

// TimeNano defines a TimeNano struct.
// 定义 TimeNano 结构体
type TimeNano struct {
	Carbon
}

// NewTimeNano returns a new TimeNano instance.
// 初始化 TimeNano 结构体
func NewTimeNano(carbon Carbon) TimeNano {
	return TimeNano{Carbon: carbon}
}

// Scan implements driver.Scanner interface.
// 实现 driver.Scanner 接口
func (t *TimeNano) Scan(src any) error {
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
		*t = NewTimeNano(c)
	}
	return t.Error
}

// Value implements driver.Valuer interface.
// 实现 driver.Valuer 接口
func (t TimeNano) Value() (driver.Value, error) {
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
func (t TimeNano) MarshalJSON() ([]byte, error) {
	emptyBytes := []byte(`""`)
	if t.IsNil() || t.IsZero() {
		return emptyBytes, nil
	}
	if t.HasError() {
		return emptyBytes, t.Error
	}
	return []byte(fmt.Sprintf(`"%s"`, t.ToTimeNanoString())), nil
}

// UnmarshalJSON implements json.Unmarshal interface.
// 实现 json.Unmarshaler 接口
func (t *TimeNano) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" || value == "0" {
		return nil
	}
	c := ParseByLayout(value, TimeLayout)
	if c.Error == nil {
		*t = NewTimeNano(c)
	}
	return c.Error
}

// String implements Stringer interface.
// 实现 Stringer 接口
func (t TimeNano) String() string {
	if t.IsZero() || t.IsInvalid() {
		return ""
	}
	return t.ToTimeNanoString()
}

// GormDataType sets gorm data type.
// 设置 gorm 数据类型
func (t TimeNano) GormDataType() string {
	return "time"
}
