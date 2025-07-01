package carbon

import (
	"bytes"
	"database/sql/driver"
)

// FormatType defines a FormatType generic struct.
type FormatType[T FormatTyper] struct {
	Carbon
}

// NewFormatType returns a new FormatType generic instance.
func NewFormatType[T FormatTyper](c Carbon) FormatType[T] {
	return FormatType[T]{
		Carbon: c,
	}
}

// Scan implements "driver.Scanner" interface for FormatType generic struct.
func (t *FormatType[T]) Scan(src any) error {
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
	*t = NewFormatType[T](c)
	return t.Error
}

// Value implements "driver.Valuer" interface for FormatType generic struct.
func (t FormatType[T]) Value() (driver.Value, error) {
	if t.IsZero() || t.IsEmpty() {
		return nil, nil
	}
	if t.HasError() {
		return nil, t.Error
	}
	return t.StdTime(), nil
}

// MarshalJSON implements "json.Marshaler" interface for FormatType generic struct.
func (t FormatType[T]) MarshalJSON() ([]byte, error) {
	if t.IsZero() || t.IsEmpty() {
		return []byte(`null`), nil
	}
	if t.HasError() {
		return []byte(`null`), t.Error
	}
	v := t.Format(t.getFormat())
	b := make([]byte, 0, len(v)+2)
	b = append(b, '"')
	b = append(b, v...)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implements "json.Unmarshaler" interface for FormatType generic struct.
func (t *FormatType[T]) UnmarshalJSON(src []byte) error {
	v := string(bytes.Trim(src, `"`))
	if v == "" {
		t.isEmpty = true
		return nil
	}
	if v == "null" {
		return nil
	}
	*t = NewFormatType[T](ParseByFormat(v, t.getFormat()))
	return t.Error
}

// String implements "Stringer" interface for FormatType generic struct.
func (t FormatType[T]) String() string {
	if t.IsInvalid() {
		return ""
	}
	return t.Format(t.getFormat())
}

// getFormat returns the format of FormatType generic struct.
func (t FormatType[T]) getFormat() string {
	var typer T
	return typer.Format()
}
