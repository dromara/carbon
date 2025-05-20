package carbon

import (
	"bytes"
	"database/sql/driver"
)

// FormatTyper defines a FormatTyper interface.
type FormatTyper interface {
	~string
	DataType() string
	Format() string
}

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

// Scan implements driver.Scanner interface for FormatType generic struct.
func (t *FormatType[T]) Scan(src any) error {
	var c Carbon
	switch v := src.(type) {
	case nil:
		return nil
	case []byte:
		c = Parse(string(v), DefaultTimezone)
	case string:
		c = Parse(v, DefaultTimezone)
	case StdTime:
		c = CreateFromStdTime(v, DefaultTimezone)
	case *StdTime:
		c = CreateFromStdTime(*v, DefaultTimezone)
	default:
		return ErrFailedScan(v)
	}
	*t = NewFormatType[T](c)
	return t.Error
}

// Value implements driver.Valuer interface for FormatType generic struct.
func (t FormatType[T]) Value() (driver.Value, error) {
	if t.IsZero() || t.IsEmpty() {
		return nil, nil
	}
	if t.HasError() {
		return nil, t.Error
	}
	return t.StdTime(), nil
}

// MarshalJSON implements json.Marshal interface for FormatType generic struct.
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

// UnmarshalJSON implements json.Unmarshal interface for FormatType generic struct.
func (t *FormatType[T]) UnmarshalJSON(src []byte) error {
	v := string(bytes.Trim(src, `"`))
	if v == "" || v == "null" {
		return nil
	}
	*t = NewFormatType[T](ParseByFormat(v, t.getFormat()))
	return t.Error
}

// String implements Stringer interface for FormatType generic struct.
func (t FormatType[T]) String() string {
	if t.IsInvalid() || t.IsZero() {
		return ""
	}
	return t.Format(t.getFormat())
}

// GormDataType implements GormDataType interface for FormatType generic struct.
func (t FormatType[T]) GormDataType() string {
	return t.getDataType()
}

// getDataType returns the data type of FormatType generic struct.
func (t FormatType[T]) getDataType() string {
	var typer T
	return typer.DataType()
}

// getFormat returns the format of FormatType generic struct.
func (t FormatType[T]) getFormat() string {
	var typer T
	return typer.Format()
}
