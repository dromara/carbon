package carbon

import (
	"bytes"
	"database/sql/driver"
	"strconv"
)

// timestamp precision constants
// 时间戳精度常量
const (
	precisionSecond      = 9
	precisionMillisecond = 999
	precisionMicrosecond = 999999
	precisionNanosecond  = 999999999
)

// TimestampTyper defines a TimestampTyper interface.
// 定义 TimestampTyper 接口
type TimestampTyper interface {
	~int64
	Precision() int64
}

// TimestampType defines a TimestampType generic struct.
// 定义 TimestampType 泛型结构体
type TimestampType[T TimestampTyper] struct {
	Carbon
}

// NewTimestampType returns a new TimestampType generic instance.
// 返回 TimestampType 泛型实例
func NewTimestampType[T TimestampTyper](c Carbon) TimestampType[T] {
	return TimestampType[T]{
		Carbon: c,
	}
}

// Scan implements driver.Scanner interface for TimestampType generic struct.
// 实现 driver.Scanner 接口
func (t *TimestampType[T]) Scan(src any) (err error) {
	var (
		ts int64
		c  Carbon
	)
	switch v := src.(type) {
	case nil:
		return nil
	case []byte:
		if ts, err = parseTimestamp(string(v)); err != nil {
			return err
		}
	case string:
		if ts, err = parseTimestamp(v); err != nil {
			return err
		}
	case int64:
		ts = v
	case StdTime:
		*t = NewTimestampType[T](CreateFromStdTime(v, DefaultTimezone))
		return t.Error
	case *StdTime:
		*t = NewTimestampType[T](CreateFromStdTime(*v, DefaultTimezone))
		return t.Error
	default:
		return ErrFailedScan(src)
	}
	switch t.getPrecision() {
	case precisionSecond:
		c = CreateFromTimestamp(ts, DefaultTimezone)
	case precisionMillisecond:
		c = CreateFromTimestampMilli(ts, DefaultTimezone)
	case precisionMicrosecond:
		c = CreateFromTimestampMicro(ts, DefaultTimezone)
	case precisionNanosecond:
		c = CreateFromTimestampNano(ts, DefaultTimezone)
	}
	*t = NewTimestampType[T](c)
	return t.Error
}

// Value implements driver.Valuer interface for TimestampType generic struct.
// 实现 driver.Valuer 接口
func (t TimestampType[T]) Value() (driver.Value, error) {
	if t.IsZero() || t.IsEmpty() {
		return nil, nil
	}
	if t.HasError() {
		return nil, t.Error
	}
	var ts int64
	switch t.getPrecision() {
	case precisionSecond:
		ts = t.Timestamp()
	case precisionMillisecond:
		ts = t.TimestampMilli()
	case precisionMicrosecond:
		ts = t.TimestampMicro()
	case precisionNanosecond:
		ts = t.TimestampNano()
	}
	return ts, nil
}

// MarshalJSON implements json.Marshal interface for TimestampType generic struct.
// 实现 json.Marshaler 接口
func (t *TimestampType[T]) MarshalJSON() ([]byte, error) {
	if t.IsZero() || t.IsEmpty() {
		return []byte(`null`), nil
	}
	if t.HasError() {
		return []byte(`null`), t.Error
	}
	var ts int64
	switch t.getPrecision() {
	case precisionSecond:
		ts = t.Timestamp()
	case precisionMillisecond:
		ts = t.TimestampMilli()
	case precisionMicrosecond:
		ts = t.TimestampMicro()
	case precisionNanosecond:
		ts = t.TimestampNano()
	}
	return []byte(strconv.FormatInt(ts, 10)), nil
}

// UnmarshalJSON implements json.Unmarshal interface for TimestampType generic struct.
// 实现 json.Unmarshaler 接口
func (t *TimestampType[T]) UnmarshalJSON(src []byte) error {
	v := string(bytes.Trim(src, `"`))
	if v == "" || v == "null" {
		return nil
	}
	ts, err := parseTimestamp(v)
	if err != nil {
		return err
	}
	var c Carbon
	switch t.getPrecision() {
	case precisionSecond:
		c = CreateFromTimestamp(ts, DefaultTimezone)
	case precisionMillisecond:
		c = CreateFromTimestampMilli(ts, DefaultTimezone)
	case precisionMicrosecond:
		c = CreateFromTimestampMicro(ts, DefaultTimezone)
	case precisionNanosecond:
		c = CreateFromTimestampNano(ts, DefaultTimezone)
	}
	*t = NewTimestampType[T](c)
	return t.Error
}

// String implements Stringer interface for TimestampType generic struct.
// 实现 Stringer 接口
func (t TimestampType[T]) String() string {
	if t.IsInvalid() || t.IsZero() {
		return "0"
	}
	return strconv.FormatInt(t.Int64(), 10)
}

// Int64 returns the timestamp value.
// 返回时间戳
func (t TimestampType[T]) Int64() (ts int64) {
	if t.IsInvalid() || t.IsZero() {
		return
	}
	switch t.getPrecision() {
	case precisionSecond:
		ts = t.Timestamp()
	case precisionMillisecond:
		ts = t.TimestampMilli()
	case precisionMicrosecond:
		ts = t.TimestampMicro()
	case precisionNanosecond:
		ts = t.TimestampNano()
	}
	return
}

// getPrecision returns the set timestamp precision.
// 返回设置的时间戳精度
func (t TimestampType[T]) getPrecision() int64 {
	var typer T
	return typer.Precision()
}
