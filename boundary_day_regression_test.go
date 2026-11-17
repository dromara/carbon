package carbon_test

import (
	"testing"
	"time"
	_ "time/tzdata"

	"github.com/dromara/carbon/v2"
)

// Some zones change their clocks at midnight, so 00:00 can be skipped or
// repeated, and in Chile 23:00-24:00 is repeated. Boundaries must still be the
// first and last instants of the receiver's local date.
func TestDayBoundaryAtMidnightClockChange(t *testing.T) {
	for _, tt := range []struct {
		name, zone, input, method, want string
	}{
		{"utc-start", "UTC", "2020-08-05T13:14:15Z", "StartOfDay", "2020-08-05T00:00:00Z"},
		{"utc-end", "UTC", "2020-08-05T13:14:15Z", "EndOfDay", "2020-08-05T23:59:59.999999999Z"},
		// 00:00 is skipped: the day starts at 01:00, not at 23:00 the day before
		{"santiago-spring-start", "America/Santiago", "2024-09-08T10:00:00-03:00", "StartOfDay", "2024-09-08T01:00:00-03:00"},
		{"santiago-spring-end", "America/Santiago", "2024-09-08T10:00:00-03:00", "EndOfDay", "2024-09-08T23:59:59.999999999-03:00"},
		{"havana-spring-start", "America/Havana", "2024-03-10T10:00:00-04:00", "StartOfDay", "2024-03-10T01:00:00-04:00"},
		{"havana-spring-month", "America/Havana", "2012-04-15T10:00:00-04:00", "StartOfMonth", "2012-04-01T01:00:00-04:00"},
		{"santiago-spring-week", "America/Santiago", "2024-09-10T10:00:00-03:00", "StartOfWeek", "2024-09-08T01:00:00-03:00"},
		// already right before: time.Date resolves the gap forward east of UTC
		{"cairo-spring-start", "Africa/Cairo", "2024-04-26T10:00:00+03:00", "StartOfDay", "2024-04-26T01:00:00+03:00"},
		// 23:00-24:00 happens twice: the day ends after the second occurrence
		{"santiago-fall-start", "America/Santiago", "2024-04-06T23:30:00-04:00", "StartOfDay", "2024-04-06T00:00:00-03:00"},
		{"santiago-fall-end", "America/Santiago", "2024-04-06T12:00:00-03:00", "EndOfDay", "2024-04-06T23:59:59.999999999-04:00"},
		{"santiago-fall-week", "America/Santiago", "2024-04-03T12:00:00-03:00", "EndOfWeek", "2024-04-06T23:59:59.999999999-04:00"},
		// 00:00-01:00 happens twice: the day starts at the first occurrence
		{"azores-fall-start", "Atlantic/Azores", "2024-10-27T12:00:00-01:00", "StartOfDay", "2024-10-27T00:00:00Z"},
		{"havana-fall-start", "America/Havana", "2024-11-03T12:00:00-05:00", "StartOfDay", "2024-11-03T00:00:00-04:00"},
		{"gaza-fall-start", "Asia/Gaza", "1996-09-20T12:00:00+02:00", "StartOfDay", "1996-09-20T00:00:00+03:00"},
		{"gaza-fall-month", "Asia/Gaza", "2004-10-15T12:00:00+02:00", "StartOfMonth", "2004-10-01T00:00:00+03:00"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			instant, err := time.Parse(time.RFC3339Nano, tt.input)
			if err != nil {
				t.Fatal(err)
			}
			c := carbon.CreateFromStdTime(instant, tt.zone).SetWeekStartsAt(carbon.Sunday)
			if c.Error != nil {
				t.Fatal(c.Error)
			}
			var got *carbon.Carbon
			switch tt.method {
			case "StartOfDay":
				got = c.StartOfDay()
			case "EndOfDay":
				got = c.EndOfDay()
			case "StartOfMonth":
				got = c.StartOfMonth()
			case "StartOfWeek":
				got = c.StartOfWeek()
			case "EndOfWeek":
				got = c.EndOfWeek()
			}
			if s := got.ToRfc3339NanoString(); s != tt.want {
				t.Errorf("%s: got %s want %s", tt.method, s, tt.want)
			}
			if got.Timezone() != tt.zone {
				t.Errorf("timezone: got %s want %s", got.Timezone(), tt.zone)
			}
			if got.WeekStartsAt() != carbon.Sunday {
				t.Errorf("week start: got %s want Sunday", got.WeekStartsAt())
			}
		})
	}
}
