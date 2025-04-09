package carbon

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

// TimestampMilli defines a TimestampMilli struct.
// 定义 TimestampMilli 结构体
type TimestampMilli struct {
	Carbon
}

// NewTimestampMilli returns a new TimestampMilli instance.
// 初始化 TimestampMilli 结构体
func NewTimestampMilli(carbon Carbon) TimestampMilli {
	return TimestampMilli{Carbon: carbon}
}

// Scan implements driver.Scanner interface.
// 实现 driver.Scanner 接口
func (t *TimestampMilli) Scan(src any) (err error) {
	ts := int64(0)
	c := NewCarbon()
	switch v := src.(type) {
	case []byte:
		ts, err = strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return ErrInvalidTimestamp(string(v))
		}
	case string:
		ts, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			return ErrInvalidTimestamp(v)
		}
	case int64:
		ts = v
	case time.Time:
		c = CreateFromStdTime(v, DefaultTimezone)
		*t = NewTimestampMilli(c)
		return t.Error
	default:
		return ErrFailedScan(src)
	}
	c = CreateFromTimestampMilli(ts, DefaultTimezone)
	*t = NewTimestampMilli(c)
	return t.Error
}

// Value implements driver.Valuer interface.
// 实现 driver.Valuer 接口
func (t TimestampMilli) Value() (driver.Value, error) {
	if t.IsNil() || t.IsZero() {
		return nil, nil
	}
	if t.HasError() {
		return nil, t.Error
	}
	return t.TimestampMilli(), nil
}

// MarshalJSON implements json.Marshal interface.
// 实现 json.Marshaler 接口
func (t TimestampMilli) MarshalJSON() ([]byte, error) {
	ts := int64(0)
	if t.IsNil() || t.IsZero() {
		return []byte(fmt.Sprintf(`%d`, ts)), nil
	}
	if t.HasError() {
		return []byte(fmt.Sprintf(`%d`, ts)), t.Error
	}
	ts = t.TimestampMilli()
	return []byte(fmt.Sprintf(`%d`, ts)), nil
}

// UnmarshalJSON implements json.Unmarshal interface.
// 实现 json.Unmarshaler 接口
func (t *TimestampMilli) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	c := NewCarbon()
	if value == "" || value == "null" || value == "0" {
		return nil
	}
	ts, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return ErrInvalidTimestamp(value)
	}
	c = CreateFromTimestampMilli(ts, DefaultTimezone)
	*t = NewTimestampMilli(c)
	return t.Error
}

// String implements Stringer interface.
// 实现 Stringer 接口
func (t TimestampMilli) String() string {
	return strconv.FormatInt(t.Int64(), 10)
}

// Int64 returns the timestamp value.
// 返回时间戳
func (t TimestampMilli) Int64() int64 {
	ts := int64(0)
	if t.IsZero() || t.IsInvalid() {
		return ts
	}
	return t.TimestampMilli()
}

// GormDataType sets gorm data type.
// 设置 gorm 数据类型
func (t TimestampMilli) GormDataType() string {
	return "time"
}
