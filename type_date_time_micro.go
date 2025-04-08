package carbon

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"time"
)

// DateTimeMicro defines a DateTimeMicro struct.
// 定义 DateTimeMicro 结构体
type DateTimeMicro struct {
	Carbon
}

// NewDateTimeMicro returns a new DateTimeMicro instance.
// 初始化 DateTimeMicro 结构体
func NewDateTimeMicro(carbon Carbon) DateTimeMicro {
	return DateTimeMicro{Carbon: carbon}
}

// Scan implements driver.Scanner interface.
// 实现 driver.Scanner 接口
func (t *DateTimeMicro) Scan(src any) error {
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
		*t = NewDateTimeMicro(c)
	}
	return t.Error
}

// Value implements driver.Valuer interface.
// 实现 driver.Valuer 接口
func (t DateTimeMicro) Value() (driver.Value, error) {
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
func (t DateTimeMicro) MarshalJSON() ([]byte, error) {
	emptyBytes := []byte(`""`)
	if t.IsNil() || t.IsZero() {
		return emptyBytes, nil
	}
	if t.HasError() {
		return emptyBytes, t.Error
	}
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateTimeMicroString())), nil
}

// UnmarshalJSON implements json.Unmarshal interface.
// 实现 json.Unmarshaler 接口
func (t *DateTimeMicro) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" || value == "0" {
		return nil
	}
	c := ParseByLayout(value, DateTimeMicroLayout)
	if c.Error == nil {
		*t = NewDateTimeMicro(c)
	}
	return c.Error
}

// String implements Stringer interface.
// 实现 Stringer 接口
func (t DateTimeMicro) String() string {
	if t.IsZero() || t.IsInvalid() {
		return ""
	}
	return t.ToDateTimeMicroString()
}

// GormDataType sets gorm data type.
// 设置 gorm 数据类型
func (t DateTimeMicro) GormDataType() string {
	return "time"
}
