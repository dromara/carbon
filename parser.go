package carbon

import (
	"fmt"
	"time"
)

// Parse parses a time string as a Carbon instance by default layouts.
//
// Note: it doesn't support parsing timestamp string.
func Parse(value string, timezone ...string) *Carbon {
	if value == "" {
		return &Carbon{isEmpty: true}
	}
	loc, err := parseTimezone(timezone...)
	if err != nil {
		return &Carbon{Error: err}
	}

	switch value {
	case "now":
		return Now().SetLocation(loc)
	case "yesterday":
		return Yesterday().SetLocation(loc)
	case "tomorrow":
		return Tomorrow().SetLocation(loc)
	}

	c := NewCarbon().SetLocation(loc)
	for _, layout := range defaultLayouts {
		if tt, err := time.ParseInLocation(layout, value, loc); err == nil {
			c.time = tt
			c.currentLayout = layout
			return c
		}
	}
	c.Error = ErrFailedParse(value)
	return c
}

// ParseByLayout parses a time string as a Carbon instance by a confirmed layout.
//
// Note: it doesn't support parsing timestamp string.
func ParseByLayout(value, layout string, timezone ...string) *Carbon {
	if value == "" {
		return &Carbon{isEmpty: true}
	}
	if layout == "" {
		return &Carbon{Error: ErrEmptyLayout()}
	}

	loc, err := parseTimezone(timezone...)
	if err != nil {
		return &Carbon{Error: err}
	}

	tt, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return &Carbon{Error: fmt.Errorf("%w: %w", ErrMismatchedLayout(value, layout), err)}
	}

	c := NewCarbon()
	c.loc = loc
	c.time = tt
	c.currentLayout = layout
	return c
}

// ParseByFormat parses a time string as a Carbon instance by a confirmed format.
//
// Note: If the letter used conflicts with the format sign, please use the escape character "\" to escape the letter
func ParseByFormat(value, format string, timezone ...string) *Carbon {
	if value == "" {
		return &Carbon{isEmpty: true}
	}
	if format == "" {
		return &Carbon{Error: ErrEmptyFormat()}
	}
	loc, err := parseTimezone(timezone...)
	if err != nil {
		return &Carbon{Error: err}
	}

	layout := format2layout(format)
	tt, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return &Carbon{Error: fmt.Errorf("%w: %w", ErrMismatchedFormat(value, format), err)}
	}

	c := NewCarbon()
	c.loc = loc
	c.time = tt
	c.currentLayout = layout
	return c
}

// ParseByLayouts parses a time string as a Carbon instance by multiple fuzzy layouts.
//
// Note: it doesn't support parsing timestamp string.
func ParseByLayouts(value string, layouts []string, timezone ...string) *Carbon {
	if value == "" {
		return &Carbon{isEmpty: true}
	}
	if len(layouts) == 0 {
		return &Carbon{Error: ErrEmptyLayout()}
	}

	loc, err := parseTimezone(timezone...)
	if err != nil {
		return &Carbon{Error: err}
	}

	c := NewCarbon().SetLocation(loc)
	for _, layout := range layouts {
		if tt, err := time.ParseInLocation(layout, value, loc); err == nil {
			c.time = tt
			c.currentLayout = layout
			return c
		}
	}
	c.Error = ErrFailedParse(value)
	return c
}

// ParseByFormats parses a time string as a Carbon instance by multiple fuzzy formats.
//
// Note: it doesn't support parsing timestamp string.
func ParseByFormats(value string, formats []string, timezone ...string) *Carbon {
	if value == "" {
		return &Carbon{isEmpty: true}
	}
	if len(formats) == 0 {
		return &Carbon{Error: ErrEmptyFormat()}
	}

	loc, err := parseTimezone(timezone...)
	if err != nil {
		return &Carbon{Error: err}
	}

	c := NewCarbon().SetLocation(loc)
	for _, format := range formats {
		layout := format2layout(format)
		if tt, err := time.ParseInLocation(layout, value, loc); err == nil {
			c.time = tt
			c.currentLayout = layout
			return c
		}
	}
	c.Error = ErrFailedParse(value)
	return c
}
