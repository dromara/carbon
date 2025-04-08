package carbon

import (
	"fmt"
)

// returns a failed parse error.
// 失败的解析错误
var failedParseError = func(value string) error {
	return fmt.Errorf("failed to parse %q as carbon", value)
}

// returns a failed scan error.
// 失败的扫描错误
var failedScanError = func(value any) error {
	return fmt.Errorf("failed to scan value %v as carbon", value)
}

// returns an invalid timestamp error.
// 无效的时间戳错误
var invalidTimestampError = func(value string) error {
	return fmt.Errorf("timestamp %s is invalid", value)
}

// returns a nil location error.
// 空指针位置错误
var nilLocationError = func() error {
	return fmt.Errorf("location cannot be nil")
}

// returns a nil language error.
// 空指针语言错误
var nilLanguageError = func() error {
	return fmt.Errorf("language cannot be nil")
}

var (
	// returns an empty locale error.
	// 空的区域错误
	emptyLocaleError = func() error {
		return fmt.Errorf("locale cannot be empty")
	}
	// returns an invalid locale error.
	// 无效的区域错误
	invalidLocaleError = func(locale string) error {
		return fmt.Errorf("locale file %q doesn't exist or is invalid", locale)
	}
)

var (
	// returns an empty resources error.
	// 空的资源错误
	emptyResourcesError = func() error {
		return fmt.Errorf("resources cannot be empty")
	}
	// returns an invalid resources error.
	// 无效的资源错误
	invalidResourcesError = func() error {
		return fmt.Errorf("resources is invalid")
	}
)

var (
	// returns an empty timezone error.
	// 空的时区错误
	emptyTimezoneError = func() error {
		return fmt.Errorf("timezone cannot be empty")
	}
	// returns an invalid timezone error.
	// 无效的时区错误
	invalidTimezoneError = func(timezone string) error {
		return fmt.Errorf("timezone %q is invalid, please see the file %q for all valid timezones", timezone, "$GOROOT/lib/time/zoneinfo.zip")
	}
)

var (
	// returns an empty duration error.
	// 空的时长错误
	emptyDurationError = func() error {
		return fmt.Errorf("duration cannot be empty")
	}
	// returns an invalid duration error.
	// 无效的时长错误
	invalidDurationError = func(duration string) error {
		return fmt.Errorf("duration %q is invalid", duration)
	}
)

var (
	// returns an empty layout error.
	// 空的布局模板错误
	emptyLayoutError = func() error {
		return fmt.Errorf("layout cannot be empty")
	}
	// returns an invalid layout error.
	// 无效的布局模板错误
	invalidLayoutError = func(value, layout string) error {
		return fmt.Errorf("cannot parse string %q as carbon by layout %q, please make sure the value and layout match", value, layout)
	}
)

var (
	// returns an empty format error.
	// 空的格式模板错误
	emptyFormatError = func() error {
		return fmt.Errorf("format cannot be empty")
	}
	// returns an invalid format error.
	// 无效的格式模板错误
	invalidFormatError = func(value, format string) error {
		return fmt.Errorf("cannot parse string %q as carbon by format %q, please make sure the value and format match", value, format)
	}
)

var (
	// returns an empty week starts day error.
	// 空的周起始日期错误
	emptyWeekStartsDayError = func() error {
		return fmt.Errorf("week start day cannot be empty")
	}
	// returns an invalid week starts at day error.
	// 无效的周起始日期错误
	invalidWeekStartsAtError = func(day string) error {
		return fmt.Errorf("week starts at day %q is invalid", day)
	}
)
