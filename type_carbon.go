package carbon

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"time"
)

// Scan implements driver.Scanner interface.
// 实现 driver.Scanner 接口
func (c *Carbon) Scan(src any) error {
	switch v := src.(type) {
	case []byte:
		*c = Parse(string(v), DefaultTimezone)
	case string:
		*c = Parse(v, DefaultTimezone)
	case time.Time:
		*c = CreateFromStdTime(v, DefaultTimezone)
	case int64:
		*c = CreateFromTimestamp(v, DefaultTimezone)
	default:
		return ErrFailedScan(v)
	}
	return c.Error
}

// Value implements driver.Valuer interface.
// 实现 driver.Valuer 接口
func (c Carbon) Value() (driver.Value, error) {
	if c.IsNil() || c.IsZero() {
		return nil, nil
	}
	if c.HasError() {
		return nil, c.Error
	}
	return c.StdTime(), nil
}

// MarshalJSON implements json.Marshal interface.
// 实现 json.Marshaler 接口
func (c Carbon) MarshalJSON() ([]byte, error) {
	emptyBytes := []byte(`""`)
	if c.IsNil() || c.IsZero() {
		return emptyBytes, nil
	}
	if c.HasError() {
		return emptyBytes, c.Error
	}
	return []byte(fmt.Sprintf(`"%s"`, c.ToDateTimeString())), nil
}

// UnmarshalJSON implements json.Unmarshal interface.
// 实现 json.Unmarshaler 接口
func (c *Carbon) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" || value == "0" {
		return nil
	}
	*c = ParseByLayout(value, DefaultLayout)
	return c.Error
}

// String implements the interface Stringer for Carbon struct.
// 实现 Stringer 接口
func (c Carbon) String() string {
	if c.IsZero() || c.IsInvalid() {
		return ""
	}
	return c.Layout(DefaultLayout, c.Timezone())
}

// GormDataType sets gorm data type.
// 设置 gorm 数据类型
func (c Carbon) GormDataType() string {
	return "time"
}
