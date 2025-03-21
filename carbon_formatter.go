package carbon

import (
	"database/sql/driver"
)

type (
	sqlScannerFunc      = func(c *Carbon, src interface{}) error
	sqlValuerFunc       = func(c *Carbon) (driver.Value, error)
	jsonMarshalerFunc   = func(c *Carbon) ([]byte, error)
	jsonUnmarshalerFunc = func(data []byte, c *Carbon) error
)

type Formatter struct {
	sqlScanner      sqlScannerFunc
	sqlValuer       sqlValuerFunc
	jsonMarshaler   jsonMarshalerFunc
	jsonUnmarshaler jsonUnmarshalerFunc
}

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (f *Formatter) isDefaultFormatter() bool {
	return f == defaultFormatter
}

func (f *Formatter) clone() Formatter {
	return Formatter{
		sqlScanner:      f.sqlScanner,
		sqlValuer:       f.sqlValuer,
		jsonMarshaler:   f.jsonMarshaler,
		jsonUnmarshaler: f.jsonUnmarshaler,
	}
}

func (f *Formatter) init() {
	// If it is the default formatter, then it needs to be cloned to
	// avoid affecting the global when setting functions
	if f.isDefaultFormatter() {
		*f = defaultFormatter.clone()
	}
}

func (f *Formatter) SetSqlScannerFunc(fn sqlScannerFunc) {
	if f == nil {
		return
	}
	f.init()
	f.sqlScanner = fn
}

func (f *Formatter) SetSqlValuerFunc(fn sqlValuerFunc) {
	if f == nil {
		return
	}
	f.init()
	f.sqlValuer = fn
}

func (f *Formatter) SetJsonMarshalerFunc(fn jsonMarshalerFunc) {
	if f == nil {
		return
	}
	f.init()
	f.jsonMarshaler = fn
}
func (f *Formatter) SetJsonUnmarshalerFunc(fn jsonUnmarshalerFunc) {
	if f == nil {
		return
	}
	f.init()
	f.jsonUnmarshaler = fn
}

func (f *Formatter) sqlScannerIsNil() bool {
	return f == nil || f.sqlScanner == nil
}

func (f *Formatter) sqlValuerIsNil() bool {
	return f == nil || f.sqlValuer == nil
}

func (f *Formatter) jsonMarshalerIsNil() bool {
	return f == nil || f.jsonMarshaler == nil
}

func (f *Formatter) jsonUnmarshalerIsNil() bool {
	return f == nil || f.jsonUnmarshaler == nil
}
