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

type formatter struct {
	sqlScanner      sqlScannerFunc
	sqlValuer       sqlValuerFunc
	jsonMarshaler   jsonMarshalerFunc
	jsonUnmarshaler jsonUnmarshalerFunc
}

func (c *Carbon) lazyInitFormatter() {
	if c.formatter == nil {
		c.formatter = &formatter{}
	}
}

func (c *Carbon) SetSqlScannerFunc(f sqlScannerFunc) {
	if f == nil {
		return
	}
	c.lazyInitFormatter()
	c.formatter.sqlScanner = f
}

func (c *Carbon) SetSqlValuerFunc(f sqlValuerFunc) {
	if f == nil {
		return
	}
	c.lazyInitFormatter()
	c.formatter.sqlValuer = f
}

func (c *Carbon) SetJsonMarshalerFunc(f jsonMarshalerFunc) {
	if f == nil {
		return
	}
	c.lazyInitFormatter()
	c.formatter.jsonMarshaler = f
}
func (c *Carbon) SetJsonUnmarshalerFunc(f jsonUnmarshalerFunc) {
	if f == nil {
		return
	}
	c.lazyInitFormatter()
	c.formatter.jsonUnmarshaler = f
}
