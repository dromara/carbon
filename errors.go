package carbon

import (
	"fmt"
)

var (
	// ErrFailedParse failed to parse error.
	ErrFailedParse = func(value any) error {
		return fmt.Errorf("failed to parse %v as carbon", value)
	}

	// ErrFailedScan failed to scan error.
	ErrFailedScan = func(value any) error {
		return fmt.Errorf("failed to scan %v as carbon", value)
	}

	// ErrInvalidTimestamp invalid timestamp error.
	ErrInvalidTimestamp = func(value string) error {
		return fmt.Errorf("invalid timestamp %v", value)
	}

	// ErrNilLocation nil location error.
	ErrNilLocation = func() error {
		return fmt.Errorf("location cannot be nil")
	}

	// ErrNilLanguage nil language error.
	ErrNilLanguage = func() error {
		return fmt.Errorf("language cannot be nil")
	}

	// ErrInvalidLanguage invalid language error.
	ErrInvalidLanguage = func(lang *Language) error {
		return fmt.Errorf("invalid Language %v", lang)
	}

	// ErrEmptyLocale empty locale error.
	ErrEmptyLocale = func() error {
		return fmt.Errorf("locale cannot be empty")
	}

	// ErrNotExistLocale not exist locale error.
	ErrNotExistLocale = func(locale string) error {
		return fmt.Errorf("locale %q doesn't exist", locale)
	}

	// ErrEmptyResources empty resources error.
	ErrEmptyResources = func() error {
		return fmt.Errorf("resources cannot be empty")
	}

	// ErrEmptyTimezone empty timezone error.
	ErrEmptyTimezone = func() error {
		return fmt.Errorf("timezone cannot be empty")
	}

	// ErrInvalidTimezone invalid timezone error.
	ErrInvalidTimezone = func(timezone string) error {
		return fmt.Errorf("invalid timezone %q, please see the file %q for all valid timezones", timezone, "$GOROOT/lib/time/zoneinfo.zip")
	}

	// ErrEmptyDuration empty duration error.
	ErrEmptyDuration = func() error {
		return fmt.Errorf("duration cannot be empty")
	}

	// ErrInvalidDuration invalid duration error.
	ErrInvalidDuration = func(duration string) error {
		return fmt.Errorf("invalid duration %q", duration)
	}

	// ErrEmptyLayout empty layout error.
	ErrEmptyLayout = func() error {
		return fmt.Errorf("layout cannot be empty")
	}

	// ErrMismatchedLayout mismatched layout error.
	ErrMismatchedLayout = func(value, layout string) error {
		return fmt.Errorf("value %q and layout %q are mismatched", value, layout)
	}

	// ErrEmptyFormat empty format error.
	ErrEmptyFormat = func() error {
		return fmt.Errorf("format cannot be empty")
	}

	// ErrMismatchedFormat mismatched format error.
	ErrMismatchedFormat = func(value, format string) error {
		return fmt.Errorf("value %q and format %q are mismatched", value, format)
	}
)
