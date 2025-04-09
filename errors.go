package carbon

import (
	"fmt"
)

var (
	// ErrFailedParse failed to parse error.
	// 失败的解析错误
	ErrFailedParse = func(value any) error {
		return fmt.Errorf("failed to parse %v as carbon", value)
	}

	// ErrFailedScan failed to scan error.
	// 失败的扫描错误
	ErrFailedScan = func(value any) error {
		return fmt.Errorf("failed to scan %v as carbon", value)
	}

	// ErrInvalidTimestamp invalid timestamp error.
	// 无效的时间戳错误
	ErrInvalidTimestamp = func(value any) error {
		return fmt.Errorf("invalid timestamp %v", value)
	}

	// ErrNilLocation nil location error.
	// 空指针位置错误
	ErrNilLocation = func() error {
		return fmt.Errorf("location cannot be nil")
	}

	// ErrNilLanguage nil language error.
	// 空指针语言错误
	ErrNilLanguage = func() error {
		return fmt.Errorf("language cannot be nil")
	}

	// ErrEmptyLocale empty locale error.
	// 空的语言环境错误
	ErrEmptyLocale = func() error {
		return fmt.Errorf("locale cannot be empty")
	}

	// ErrNotExistLocale not exist locale error.
	// 不存在的语言环境错误
	ErrNotExistLocale = func(locale string) error {
		return fmt.Errorf("locale %q doesn't exist", locale)
	}

	// ErrEmptyResources empty resources error.
	// 空的资源错误
	ErrEmptyResources = func() error {
		return fmt.Errorf("resources cannot be empty")
	}

	// ErrInvalidResourcesError invalid resources error.
	// 无效的资源错误
	ErrInvalidResourcesError = func() error {
		return fmt.Errorf("invalid resources")
	}

	// ErrEmptyTimezone empty timezone error.
	// 空的时区错误
	ErrEmptyTimezone = func() error {
		return fmt.Errorf("timezone cannot be empty")
	}

	// ErrInvalidTimezone invalid timezone error.
	// 无效的时区错误
	ErrInvalidTimezone = func(timezone string) error {
		return fmt.Errorf("invalid timezone %q, please see the file %q for all valid timezones", timezone, "$GOROOT/lib/time/zoneinfo.zip")
	}

	// ErrEmptyDuration empty duration error.
	// 空的时长错误
	ErrEmptyDuration = func() error {
		return fmt.Errorf("duration cannot be empty")
	}

	// ErrInvalidDuration invalid duration error.
	// 无效的时长错误
	ErrInvalidDuration = func(duration string) error {
		return fmt.Errorf("invalid duration %q", duration)
	}

	// ErrEmptyLayout empty layout error.
	// 空的布局模板错误
	ErrEmptyLayout = func() error {
		return fmt.Errorf("layout cannot be empty")
	}

	// ErrMismatchedLayout mismatched layout error.
	// 不匹配的布局模板错误
	ErrMismatchedLayout = func(value, layout string) error {
		return fmt.Errorf("value %q and layout %q are mismatched", value, layout)
	}

	// ErrEmptyFormat empty format error.
	// 空的格式模板错误
	ErrEmptyFormat = func() error {
		return fmt.Errorf("format cannot be empty")
	}

	// ErrMismatchedFormat mismatched format error.
	// 不匹配的格式模板错误
	ErrMismatchedFormat = func(value, format string) error {
		return fmt.Errorf("value %q and format %q are mismatched", value, format)
	}

	// ErrEmptyWeekStartDay empty week start day error.
	// 空的周起始日期错误
	ErrEmptyWeekStartDay = func() error {
		return fmt.Errorf("week start day cannot be empty")
	}

	// ErrInvalidWeekStartDay invalid week start day error.
	// 无效的周起始日期错误
	ErrInvalidWeekStartDay = func(day string) error {
		return fmt.Errorf("invalid week start day %q", day)
	}
)
