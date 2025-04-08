package carbon

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

// TimestampNano defines a TimestampNano struct.
// 定义 TimestampNano 结构体
type TimestampNano struct {
	Carbon
}

// NewTimestampNano returns a new TimestampNano instance.
// 初始化 TimestampNano 结构体
func NewTimestampNano(carbon Carbon) TimestampNano {
	return TimestampNano{Carbon: carbon}
}

// Scan implements driver.Scanner interface.
// 实现 driver.Scanner 接口
func (t *TimestampNano) Scan(src any) (err error) {
	ts := int64(0)
	c := NewCarbon()
	switch v := src.(type) {
	case []byte:
		ts, err = strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return invalidTimestampError(string(v))
		}
	case string:
		ts, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			return invalidTimestampError(v)
		}
	case int64:
		ts = v
	case time.Time:
		c = CreateFromStdTime(v, DefaultTimezone)
		*t = NewTimestampNano(c)
		return t.Error
	default:
		return failedScanError(src)
	}
	c = CreateFromTimestampNano(ts, DefaultTimezone)
	*t = NewTimestampNano(c)
	return t.Error
}

// Value implements driver.Valuer interface.
// 实现 driver.Valuer 接口
func (t TimestampNano) Value() (driver.Value, error) {
	if t.IsNil() || t.IsZero() {
		return nil, nil
	}
	if t.HasError() {
		return nil, t.Error
	}
	return t.TimestampNano(), nil
}

// MarshalJSON implements json.Marshal interface.
// 实现 json.Marshaler 接口
func (t TimestampNano) MarshalJSON() ([]byte, error) {
	ts := int64(0)
	if t.IsNil() || t.IsZero() {
		return []byte(fmt.Sprintf(`%d`, ts)), nil
	}
	if t.HasError() {
		return []byte(fmt.Sprintf(`%d`, ts)), t.Error
	}
	ts = t.TimestampNano()
	return []byte(fmt.Sprintf(`%d`, ts)), nil
}

// UnmarshalJSON implements json.Unmarshal interface.
// 实现 json.Unmarshaler 接口
func (t *TimestampNano) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	c := NewCarbon()
	if value == "" || value == "null" || value == "0" {
		return nil
	}
	ts, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return invalidTimestampError(value)
	}
	c = CreateFromTimestampNano(ts, DefaultTimezone)
	*t = NewTimestampNano(c)
	return t.Error
}

// String implements Stringer interface.
// 实现 Stringer 接口
func (t TimestampNano) String() string {
	return strconv.FormatInt(t.Int64(), 10)
}

// Int64 returns the timestamp value.
// 返回时间戳
func (t TimestampNano) Int64() int64 {
	ts := int64(0)
	if t.IsZero() || t.IsInvalid() {
		return ts
	}
	return t.TimestampNano()
}

// GormDataType sets gorm data type.
// 设置 gorm 数据类型
func (t TimestampNano) GormDataType() string {
	return "time"
}
