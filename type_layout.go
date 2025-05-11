package carbon

import (
	"bytes"
	"database/sql/driver"
)

// LayoutTyper defines a LayoutTyper interface
// 定义 LayoutTyper 接口
type LayoutTyper interface {
	~string
	Layout() string
}

// LayoutType defines a LayoutType generic struct
// 定义 LayoutType 泛型结构体
type LayoutType[T LayoutTyper] struct {
	Carbon
}

// NewLayoutType returns a new LayoutType generic instance.
// 返回 LayoutType 泛型实例
func NewLayoutType[T LayoutTyper](c Carbon) LayoutType[T] {
	return LayoutType[T]{
		Carbon: c,
	}
}

// Scan implements driver.Scanner interface for LayoutType generic struct.
// 实现 driver.Scanner 接口
func (t *LayoutType[T]) Scan(src any) error {
	var c Carbon
	switch v := src.(type) {
	case nil:
		return nil
	case []byte:
		c = Parse(string(v), DefaultTimezone)
	case string:
		c = Parse(v, DefaultTimezone)
	case int64:
		c = CreateFromTimestamp(v, DefaultTimezone)
	case StdTime:
		c = CreateFromStdTime(v, DefaultTimezone)
	case *StdTime:
		c = CreateFromStdTime(*v, DefaultTimezone)
	default:
		return ErrFailedScan(v)
	}
	*t = NewLayoutType[T](c)
	return t.Error
}

// Value implements driver.Valuer interface for LayoutType generic struct.
// 实现 driver.Valuer 接口
func (t LayoutType[T]) Value() (driver.Value, error) {
	if t.IsEmpty() || t.IsZero() {
		return nil, nil
	}
	if t.HasError() {
		return nil, t.Error
	}
	return t.StdTime(), nil
}

// MarshalJSON implements json.Marshal interface for LayoutType generic struct.
// 实现 json.Marshaler 接口
func (t LayoutType[T]) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`null`), nil
	}
	if t.HasError() {
		return []byte(`null`), t.Error
	}
	if t.IsEmpty() {
		return []byte(`""`), nil
	}
	v := t.Layout(t.getLayout())
	b := make([]byte, 0, len(v)+2)
	b = append(b, '"')
	b = append(b, v...)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implements json.Unmarshal interface for LayoutType generic struct.
// 实现 json.Unmarshaler 接口
func (t *LayoutType[T]) UnmarshalJSON(src []byte) error {
	v := string(bytes.Trim(src, `"`))
	if v == "" || v == "null" || v == "0" {
		return nil
	}
	*t = NewLayoutType[T](ParseByLayout(v, t.getLayout()))
	return t.Error
}

// String implements Stringer interface for LayoutType generic struct.
// 实现 Stringer 接口
func (t LayoutType[T]) String() string {
	if t.IsInvalid() || t.IsZero() {
		return ""
	}
	return t.Layout(t.getLayout())
}

// getLayout returns the set layout.
// 返回设置的布局模板
func (t LayoutType[T]) getLayout() string {
	var typer T
	return typer.Layout()
}
