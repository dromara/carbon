package carbon

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

// Timestamp defines a Timestamp struct.
// 定义 Timestamp 结构体
type Timestamp struct {
	Carbon
}

// NewTimestamp returns a new Timestamp instance.
// 初始化 Timestamp 结构体
func NewTimestamp(carbon Carbon) Timestamp {
	return Timestamp{Carbon: carbon}
}

// Scan implements driver.Scanner interface.
// 实现 driver.Scanner 接口
func (t *Timestamp) Scan(src any) (err error) {
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
		*t = NewTimestamp(c)
		return t.Error
	default:
		return failedScanError(src)
	}
	c = CreateFromTimestamp(ts, DefaultTimezone)
	*t = NewTimestamp(c)
	return t.Error
}

// Value implements driver.Valuer interface.
// 实现 driver.Valuer 接口
func (t Timestamp) Value() (driver.Value, error) {
	if t.IsNil() || t.IsZero() {
		return nil, nil
	}
	if t.HasError() {
		return nil, t.Error
	}
	return t.Timestamp(), nil
}

// MarshalJSON implements json.Marshal interface.
// 实现 json.Marshaler 接口
func (t Timestamp) MarshalJSON() ([]byte, error) {
	ts := int64(0)
	if t.IsNil() || t.IsZero() {
		return []byte(fmt.Sprintf(`%d`, ts)), nil
	}
	if t.HasError() {
		return []byte(fmt.Sprintf(`%d`, ts)), t.Error
	}
	ts = t.Timestamp()
	return []byte(fmt.Sprintf(`%d`, ts)), nil
}

// UnmarshalJSON implements json.Unmarshal interface.
// 实现 json.Unmarshaler 接口
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	c := NewCarbon()
	if value == "" || value == "null" || value == "0" {
		return nil
	}
	ts, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return invalidTimestampError(value)
	}
	c = CreateFromTimestamp(ts, DefaultTimezone)
	*t = NewTimestamp(c)
	return t.Error
}

// String implements Stringer interface.
// 实现 Stringer 接口
func (t Timestamp) String() string {
	return strconv.FormatInt(t.Int64(), 10)
}

// Int64 returns the timestamp value.
// 返回时间戳
func (t Timestamp) Int64() int64 {
	ts := int64(0)
	if t.IsZero() || t.IsInvalid() {
		return ts
	}
	return t.Timestamp()
}

// GormDataType sets gorm data type.
// 设置 gorm 数据类型
func (t Timestamp) GormDataType() string {
	return "time"
}
