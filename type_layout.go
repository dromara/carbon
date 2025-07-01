package carbon

import (
	"bytes"
	"database/sql/driver"
)

// LayoutType defines a LayoutType generic struct
type LayoutType[T LayoutTyper] struct {
	Carbon
}

// NewLayoutType returns a new LayoutType generic instance.
func NewLayoutType[T LayoutTyper](c Carbon) LayoutType[T] {
	return LayoutType[T]{
		Carbon: c,
	}
}

// Scan implements "driver.Scanner" interface for LayoutType generic struct.
func (t *LayoutType[T]) Scan(src any) error {
	var c Carbon
	switch v := src.(type) {
	case nil:
		return nil
	case []byte:
		c = Parse(string(v))
	case string:
		c = Parse(v)
	case StdTime:
		c = CreateFromStdTime(v)
	case *StdTime:
		c = CreateFromStdTime(*v)
	default:
		return ErrFailedScan(v)
	}
	*t = NewLayoutType[T](c)
	return t.Error
}

// Value implements "driver.Valuer" interface for LayoutType generic struct.
func (t LayoutType[T]) Value() (driver.Value, error) {
	if t.IsZero() || t.IsEmpty() {
		return nil, nil
	}
	if t.HasError() {
		return nil, t.Error
	}
	return t.StdTime(), nil
}

// MarshalJSON implements "json.Marshaler" interface for LayoutType generic struct.
func (t LayoutType[T]) MarshalJSON() ([]byte, error) {
	if t.IsZero() || t.IsEmpty() {
		return []byte(`null`), nil
	}
	if t.HasError() {
		return []byte(`null`), t.Error
	}
	v := t.Layout(t.getLayout())
	b := make([]byte, 0, len(v)+2)
	b = append(b, '"')
	b = append(b, v...)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implements "json.Unmarshaler" interface for LayoutType generic struct.
func (t *LayoutType[T]) UnmarshalJSON(src []byte) error {
	v := string(bytes.Trim(src, `"`))
	if v == "" {
		t.isEmpty = true
		return nil
	}
	if v == "null" {
		return nil
	}
	*t = NewLayoutType[T](ParseByLayout(v, t.getLayout()))
	return t.Error
}

// String implements "Stringer" interface for LayoutType generic struct.
func (t LayoutType[T]) String() string {
	if t.IsInvalid() {
		return ""
	}
	return t.Layout(t.getLayout())
}

// getLayout returns the layout of LayoutType generic struct.
func (t LayoutType[T]) getLayout() string {
	var typer T
	return typer.Layout()
}
