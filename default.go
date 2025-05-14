package carbon

var (
	// DefaultLayout default layout
	DefaultLayout = DateTimeLayout

	// DefaultTimezone default timezone
	DefaultTimezone = UTC

	// DefaultWeekStartsAt default week start date
	DefaultWeekStartsAt = Monday

	// DefaultWeekendDays Default weekend days
	DefaultWeekendDays = []Weekday{
		Saturday, Sunday,
	}

	// DefaultLocale default language locale
	DefaultLocale = "en"
)

// Default defines a Default struct.
type Default struct {
	Layout       string
	Timezone     string
	WeekStartsAt Weekday
	WeekendDays  []Weekday
	Locale       string
}

// SetDefault sets default.
func SetDefault(d Default) {
	if d.Layout != "" {
		DefaultLayout = d.Layout
	}
	if d.Timezone != "" {
		DefaultTimezone = d.Timezone
	}
	if d.WeekStartsAt.String() != "" {
		DefaultWeekStartsAt = d.WeekStartsAt
	}
	if len(d.WeekendDays) > 0 {
		DefaultWeekendDays = d.WeekendDays
	}
	if d.Locale != "" {
		DefaultLocale = d.Locale
	}
}

// ResetDefault resets default.
func ResetDefault() {
	DefaultLayout = DateTimeLayout
	DefaultTimezone = UTC
	DefaultWeekStartsAt = Monday
	DefaultWeekendDays = []Weekday{
		Saturday, Sunday,
	}
	DefaultLocale = "en"
}
