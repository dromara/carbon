package carbon

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

// returns a failed scan error.
// 失败的扫描错误
var failedScanError = func(src interface{}) error {
	return fmt.Errorf("failed to scan value: %v", src)
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (c *Carbon) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*c = Parse(string(v))
		}
	case string:
		if len(v) > 0 {
			*c = Parse(v)
		}
	case time.Time:
		*c = CreateFromStdTime(v)
	}
	if c.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (c Carbon) Value() (driver.Value, error) {
	if c.IsZero() {
		return nil, nil
	}
	return c.StdTime(), nil
}

// MarshalJSON implements the interface json.Marshal for Carbon struct.
// 实现 json.Marshaler 接口
func (c Carbon) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, c.Layout(c.layout))), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for Carbon struct.
// 实现 json.Unmarshaler 接口
func (c *Carbon) UnmarshalJSON(b []byte) error {
	if c.layout == "" {
		c.layout = DefaultLayout
	}
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	*c = ParseByLayout(value, c.layout)
	return c.Error
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (t *DateTime) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*t = NewDateTime(Parse(string(v)))
		}
	case string:
		if len(v) > 0 {
			*t = NewDateTime(Parse(v))
		}
	case time.Time:
		*t = NewDateTime(CreateFromStdTime(v))
	}
	if t.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (t DateTime) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil
}

// MarshalJSON implements the interface json.Marshal for DateTime struct.
// 实现 MarshalJSON 接口
func (t DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateTimeString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateTime struct.
// 实现 UnmarshalJSON 接口
func (t *DateTime) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	c := ParseByLayout(value, DateTimeLayout)
	if c.Error == nil {
		*t = NewDateTime(c)
	}
	return c.Error
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (t *DateTimeMilli) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*t = NewDateTimeMilli(Parse(string(v)))
		}
	case string:
		if len(v) > 0 {
			*t = NewDateTimeMilli(Parse(v))
		}
	case time.Time:
		*t = NewDateTimeMilli(CreateFromStdTime(v))
	}
	if t.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (t DateTimeMilli) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil
}

// MarshalJSON implements the interface json.Marshal for DateTimeMilli struct.
// 实现 MarshalJSON 接口
func (t DateTimeMilli) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateTimeMilliString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateTimeMilli struct.
// 实现 UnmarshalJSON 接口
func (t *DateTimeMilli) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	c := ParseByLayout(value, DateTimeMilliLayout)
	if c.Error == nil {
		*t = NewDateTimeMilli(c)
	}
	return c.Error
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (t *DateTimeMicro) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*t = NewDateTimeMicro(Parse(string(v)))
		}
	case string:
		if len(v) > 0 {
			*t = NewDateTimeMicro(Parse(v))
		}
	case time.Time:
		*t = NewDateTimeMicro(CreateFromStdTime(v))
	}
	if t.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (t DateTimeMicro) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil
}

// MarshalJSON implements the interface json.Marshal for DateTimeMicro struct.
// 实现 MarshalJSON 接口
func (t DateTimeMicro) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateTimeMicroString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateTimeMicro struct.
// 实现 UnmarshalJSON 接口
func (t *DateTimeMicro) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	c := ParseByLayout(value, DateTimeMicroLayout)
	if c.Error == nil {
		*t = NewDateTimeMicro(c)
	}
	return c.Error
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (t *DateTimeNano) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*t = NewDateTimeNano(Parse(string(v)))
		}
	case string:
		if len(v) > 0 {
			*t = NewDateTimeNano(Parse(v))
		}
	case time.Time:
		*t = NewDateTimeNano(CreateFromStdTime(v))
	}
	if t.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (t DateTimeNano) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil
}

// MarshalJSON implements the interface json.Marshal for DateTimeNano struct.
// 实现 MarshalJSON 接口
func (t DateTimeNano) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateTimeNanoString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateTimeNano struct.
// 实现 UnmarshalJSON 接口
func (t *DateTimeNano) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	c := ParseByLayout(value, DateTimeNanoLayout)
	if c.Error == nil {
		*t = NewDateTimeNano(c)
	}
	return c.Error
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (t *Date) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*t = NewDate(Parse(string(v)))
		}
	case string:
		if len(v) > 0 {
			*t = NewDate(Parse(v))
		}
	case time.Time:
		*t = NewDate(CreateFromStdTime(v))
	}
	if t.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (t Date) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil
}

// MarshalJSON implements the interface json.Marshal for Date struct.
// 实现 MarshalJSON 接口
func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for Date struct.
// 实现 UnmarshalJSON 接口
func (t *Date) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	c := ParseByLayout(value, DateLayout)
	if c.Error == nil {
		*t = NewDate(c)
	}
	return c.Error
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (t *DateMilli) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*t = NewDateMilli(Parse(string(v)))
		}
	case string:
		if len(v) > 0 {
			*t = NewDateMilli(Parse(v))
		}
	case time.Time:
		*t = NewDateMilli(CreateFromStdTime(v))
	}
	if t.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (t DateMilli) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil
}

// MarshalJSON implements the interface json.Marshal for DateMilli struct.
// 实现 MarshalJSON 接口
func (t DateMilli) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateMilliString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateMilli struct.
// 实现 UnmarshalJSON 接口
func (t *DateMilli) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	c := ParseByLayout(value, DateMilliLayout)
	if c.Error == nil {
		*t = NewDateMilli(c)
	}
	return c.Error
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (t *DateMicro) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*t = NewDateMicro(Parse(string(v)))
		}
	case string:
		if len(v) > 0 {
			*t = NewDateMicro(Parse(v))
		}
	case time.Time:
		*t = NewDateMicro(CreateFromStdTime(v))
	}
	if t.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (t DateMicro) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil
}

// MarshalJSON implements the interface json.Marshal for DateMicro struct.
// 实现 MarshalJSON 接口
func (t DateMicro) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateMicroString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateMicro struct.
// 实现 UnmarshalJSON 接口
func (t *DateMicro) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	c := ParseByLayout(value, DateMicroLayout)
	if c.Error == nil {
		*t = NewDateMicro(c)
	}
	return c.Error
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (t *DateNano) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*t = NewDateNano(Parse(string(v)))
		}
	case string:
		if len(v) > 0 {
			*t = NewDateNano(Parse(v))
		}
	case time.Time:
		*t = NewDateNano(CreateFromStdTime(v))
	}
	if t.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (t DateNano) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil
}

// MarshalJSON implements the interface json.Marshal for DateNano struct.
// 实现 MarshalJSON 接口
func (t DateNano) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToDateNanoString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for DateNano struct.
// 实现 UnmarshalJSON 接口
func (t *DateNano) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	c := ParseByLayout(value, DateNanoLayout)
	if c.Error == nil {
		*t = NewDateNano(c)
	}
	return c.Error
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (t *Time) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*t = NewTime(Parse(string(v)))
		}
	case string:
		if len(v) > 0 {
			*t = NewTime(Parse(v))
		}
	case time.Time:
		*t = NewTime(CreateFromStdTime(v))
	}
	if t.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (t Time) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil
}

// MarshalJSON implements the interface json.Marshal for Time struct.
// 实现 MarshalJSON 接口
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToTimeString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for Time struct.
// 实现 UnmarshalJSON 接口
func (t *Time) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	year, month, day := Now().Date()
	c := ParseByLayout(fmt.Sprintf("%04d-%02d-%02d %s", year, month, day, value), DateTimeLayout)
	fmt.Println("c", c)
	if c.Error == nil {
		*t = NewTime(c)
	}
	return c.Error
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (t *TimeMilli) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*t = NewTimeMilli(Parse(string(v)))
		}
	case string:
		if len(v) > 0 {
			*t = NewTimeMilli(Parse(v))
		}
	case time.Time:
		*t = NewTimeMilli(CreateFromStdTime(v))
	}
	if t.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (t TimeMilli) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil
}

// MarshalJSON implements the interface json.Marshal for TimeMilli struct.
// 实现 MarshalJSON 接口
func (t TimeMilli) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToTimeMilliString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimeMilli struct.
// 实现 UnmarshalJSON 接口
func (t *TimeMilli) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	year, month, day := Now().Date()
	c := ParseByLayout(fmt.Sprintf("%04d-%02d-%02d %s", year, month, day, value), DateTimeMilliLayout)
	if c.Error == nil {
		*t = NewTimeMilli(c)
	}
	return c.Error
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (t *TimeMicro) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*t = NewTimeMicro(Parse(string(v)))
		}
	case string:
		if len(v) > 0 {
			*t = NewTimeMicro(Parse(v))
		}
	case time.Time:
		*t = NewTimeMicro(CreateFromStdTime(v))
	}
	if t.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (t TimeMicro) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil
}

// MarshalJSON implements the interface json.Marshal for TimeMicro struct.
// 实现 MarshalJSON 接口
func (t TimeMicro) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToTimeMicroString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimeMicro struct.
// 实现 UnmarshalJSON 接口
func (t *TimeMicro) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	year, month, day := Now().Date()
	c := ParseByLayout(fmt.Sprintf("%04d-%02d-%02d %s", year, month, day, value), DateTimeMicroLayout)
	if c.Error == nil {
		*t = NewTimeMicro(c)
	}
	return c.Error
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (t *TimeNano) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*t = NewTimeNano(Parse(string(v)))
		}
	case string:
		if len(v) > 0 {
			*t = NewTimeNano(Parse(v))
		}
	case time.Time:
		*t = NewTimeNano(CreateFromStdTime(v))
	}
	if t.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (t TimeNano) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil
}

// MarshalJSON implements the interface json.Marshal for TimeNano struct.
// 实现 MarshalJSON 接口
func (t TimeNano) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToTimeNanoString())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimeNano struct.
// 实现 UnmarshalJSON 接口
func (t *TimeNano) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	year, month, day := Now().Date()
	c := ParseByLayout(fmt.Sprintf("%04d-%02d-%02d %s", year, month, day, value), DateTimeNanoLayout)
	if c.Error == nil {
		*t = NewTimeNano(c)
	}
	return c.Error
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (t *Timestamp) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*t = NewTimestamp(Parse(string(v)))
		}
	case string:
		if len(v) > 0 {
			*t = NewTimestamp(Parse(v))
		}
	case time.Time:
		*t = NewTimestamp(CreateFromStdTime(v))
	}
	if t.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (t Timestamp) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil
}

// MarshalJSON implements the interface json.Marshal for Timestamp struct.
// 实现 MarshalJSON 接口
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.Timestamp())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for Timestamp struct.
// 实现 UnmarshalJSON 接口
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	ts, _ := strconv.ParseInt(value, 10, 64)
	c := CreateFromTimestamp(ts)
	if c.Error == nil {
		*t = NewTimestamp(c)
	}
	return c.Error
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (t *TimestampMilli) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*t = NewTimestampMilli(Parse(string(v)))
		}
	case string:
		if len(v) > 0 {
			*t = NewTimestampMilli(Parse(v))
		}
	case time.Time:
		*t = NewTimestampMilli(CreateFromStdTime(v))
	}
	if t.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (t TimestampMilli) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil
}

// MarshalJSON implements the interface json.Marshal for TimestampMilli struct.
// 实现 MarshalJSON 接口
func (t TimestampMilli) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampMilli())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimestampMilli struct.
// 实现 UnmarshalJSON 接口
func (t *TimestampMilli) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	ts, _ := strconv.ParseInt(value, 10, 64)
	c := CreateFromTimestampMilli(ts)
	if c.Error == nil {
		*t = NewTimestampMilli(c)
	}
	return c.Error
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (t *TimestampMicro) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*t = NewTimestampMicro(Parse(string(v)))
		}
	case string:
		if len(v) > 0 {
			*t = NewTimestampMicro(Parse(v))
		}
	case time.Time:
		*t = NewTimestampMicro(CreateFromStdTime(v))
	}
	if t.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (t TimestampMicro) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil
}

// MarshalJSON implements the interface MarshalJSON for TimestampMicro struct.
// 实现 MarshalJSON 接口
func (t TimestampMicro) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampMicro())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimestampMicro struct.
// 实现 UnmarshalJSON 接口
func (t *TimestampMicro) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	ts, _ := strconv.ParseInt(value, 10, 64)
	c := CreateFromTimestampMicro(ts)
	if c.Error == nil {
		*t = NewTimestampMicro(c)
	}
	return c.Error
}

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (t *TimestampNano) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		if len(v) > 0 {
			*t = NewTimestampNano(Parse(string(v)))
		}
	case string:
		if len(v) > 0 {
			*t = NewTimestampNano(Parse(v))
		}
	case time.Time:
		*t = NewTimestampNano(CreateFromStdTime(v))
	}
	if t.Error == nil {
		return nil
	}
	return failedScanError(src)
}

// Value the interface providing the Value method for package database/sql/driver.
func (t TimestampNano) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.StdTime(), nil
}

// MarshalJSON implements the interface json.Marshal for TimestampNano struct.
// 实现 MarshalJSON 接口
func (t TimestampNano) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%d`, t.TimestampNano())), nil
}

// UnmarshalJSON implements the interface json.Unmarshal for TimestampNano struct.
// 实现 UnmarshalJSON 接口
func (t *TimestampNano) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	if value == "" || value == "null" {
		return nil
	}
	ts, _ := strconv.ParseInt(value, 10, 64)
	c := CreateFromTimestampNano(ts)
	if c.Error == nil {
		*t = NewTimestampNano(c)
	}
	return c.Error
}

// Int64 outputs timestamp with second.
// 输出秒级时间戳
func (t Timestamp) Int64() int64 {
	return t.Timestamp()
}

// Int64 outputs timestamp with millisecond.
// 输出豪秒级时间戳
func (t TimestampMilli) Int64() int64 {
	return t.TimestampMilli()
}

// Int64 outputs timestamp with microsecond.
// 输出微秒级时间戳
func (t TimestampMicro) Int64() int64 {
	return t.TimestampMicro()
}

// Int64 outputs timestamp with nanosecond.
// 输出纳秒级时间戳
func (t TimestampNano) Int64() int64 {
	return t.TimestampNano()
}

// String implements the interface Stringer for DateTime struct.
// 实现 Stringer 接口
func (t DateTime) String() string {
	return t.ToDateTimeString()
}

// String implements the interface Stringer for DateTimeMilli struct.
// 实现 Stringer 接口
func (t DateTimeMilli) String() string {
	return t.ToDateTimeMilliString()
}

// String implements the interface Stringer for DateTimeMicro struct.
// 实现 Stringer 接口
func (t DateTimeMicro) String() string {
	return t.ToDateTimeMicroString()
}

// String implements the interface Stringer for DateTimeNano struct.
// 实现 Stringer 接口
func (t DateTimeNano) String() string {
	return t.ToDateTimeNanoString()
}

// String implements the interface Stringer for Date struct.
// 实现 Stringer 接口
func (t Date) String() string {
	return t.ToDateString()
}

// String implements the interface Stringer for DateMilli struct.
// 实现 Stringer 接口
func (t DateMilli) String() string {
	return t.ToDateMilliString()
}

// String implements the interface Stringer for DateMicro struct.
// 实现 Stringer 接口
func (t DateMicro) String() string {
	return t.ToDateMicroString()
}

// String implements the interface Stringer for DateNano struct.
// 实现 Stringer 接口
func (t DateNano) String() string {
	return t.ToDateNanoString()
}

// String implements the interface Stringer for Time struct.
// 实现 Stringer 接口
func (t Time) String() string {
	return t.ToTimeString()
}

// String implements the interface Stringer for TimeMilli struct.
// 实现 Stringer 接口
func (t TimeMilli) String() string {
	return t.ToTimeMilliString()
}

// String implements the interface Stringer for TimeMicro struct.
// 实现 Stringer 接口
func (t TimeMicro) String() string {
	return t.ToTimeMicroString()
}

// String implements the interface Stringer for TimeNano struct.
// 实现 Stringer 接口
func (t TimeNano) String() string {
	return t.ToTimeNanoString()
}

// String implements the interface Stringer for Timestamp struct.
// 实现 Stringer 接口
func (t Timestamp) String() string {
	return strconv.FormatInt(t.Timestamp(), 10)
}

// String implements the interface Stringer for TimestampMilli struct.
// 实现 Stringer 接口
func (t TimestampMilli) String() string {
	return strconv.FormatInt(t.TimestampMilli(), 10)
}

// String implements the interface Stringer for TimestampMicro struct.
// 实现 Stringer 接口
func (t TimestampMicro) String() string {
	return strconv.FormatInt(t.TimestampMicro(), 10)
}

// String implements the interface Stringer for TimestampNano struct.
// 实现 Stringer 接口
func (t TimestampNano) String() string {
	return strconv.FormatInt(t.TimestampNano(), 10)
}
