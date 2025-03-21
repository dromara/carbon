package carbon

import (
	"database/sql/driver"
)

type (
	// sql.Scanner
	sqlScannerFunc = func(src any) error
	// sql/driver.Valuer
	sqlValuerFunc = func() (driver.Value, error)
	// json.Marshaler
	jsonMarshalerFunc = func(v any) ([]byte, error)
	// json.Unmarshaler
	jsonUnmarshalerFunc = func(data []byte, v any) error
)

type formatter struct {
	// sql.Scanner
	sqlScanner sqlScannerFunc
	// sql/driver.Valuer
	sqlValuer sqlValuerFunc
	// json.Marshaler
	jsonMarshaler jsonMarshalerFunc
	// json.Unmarshaler
	jsonUnmarshaler jsonUnmarshalerFunc
}

func (c *Carbon) lazyInitFormatter() {
	if c.formatter == nil {
		c.formatter = &formatter{}
	}
}

func (c *Carbon) SetSqlScannerFunc(f sqlScannerFunc) {
	c.lazyInitFormatter()
	c.formatter.sqlScanner = f
}

func (c *Carbon) SetSqlValuerFunc(f sqlValuerFunc) {
	c.lazyInitFormatter()
	c.formatter.sqlValuer = f
}

func (c *Carbon) SetJsonMarshalerFunc(f jsonMarshalerFunc) {
	c.lazyInitFormatter()
	c.formatter.jsonMarshaler = f
}
func (c *Carbon) SetJsonUnmarshalerFunc(f jsonUnmarshalerFunc) {
	c.lazyInitFormatter()
	c.formatter.jsonUnmarshaler = f
}
