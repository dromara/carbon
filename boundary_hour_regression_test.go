package carbon_test

import (
	"testing"
	"time"
	_ "time/tzdata"

	"github.com/dromara/carbon/v2"
)

func TestHourBoundaryOccurrence(t *testing.T) {
	for _, tt := range []struct {
		name, zone, input, start, end string
	}{
		{"utc", "UTC", "2020-08-05T13:14:15.123456789Z", "2020-08-05T13:00:00Z", "2020-08-05T13:59:59.999999999Z"},
		{"kathmandu", "Asia/Kathmandu", "2024-01-15T13:14:15.123456789+05:45", "2024-01-15T13:00:00+05:45", "2024-01-15T13:59:59.999999999+05:45"},
		{"new-york-first", "America/New_York", "2024-11-03T01:30:45.123456789-04:00", "2024-11-03T01:00:00-04:00", "2024-11-03T01:59:59.999999999-04:00"},
		{"new-york-second", "America/New_York", "2024-11-03T01:30:45.123456789-05:00", "2024-11-03T01:00:00-05:00", "2024-11-03T01:59:59.999999999-05:00"},
		{"berlin-first", "Europe/Berlin", "2024-10-27T02:30:45.123456789+02:00", "2024-10-27T02:00:00+02:00", "2024-10-27T02:59:59.999999999+02:00"},
		{"berlin-second", "Europe/Berlin", "2024-10-27T02:30:45.123456789+01:00", "2024-10-27T02:00:00+01:00", "2024-10-27T02:59:59.999999999+01:00"},
		{"lord-howe-first", "Australia/Lord_Howe", "2024-04-07T01:45:15.123456789+11:00", "2024-04-07T01:00:00+11:00", "2024-04-07T01:59:59.999999999+11:00"},
		// 01:00 +10:30 never occurs in this rollback. Keep the existing start,
		// rather than subtracting 45 elapsed minutes and returning 01:30 +11:00.
		{"lord-howe-second", "Australia/Lord_Howe", "2024-04-07T01:45:15.123456789+10:30", "2024-04-07T01:00:00+11:00", "2024-04-07T01:59:59.999999999+10:30"},
		{"lord-howe-spring", "Australia/Lord_Howe", "2024-10-06T02:45:15.123456789+11:00", "2024-10-06T02:30:00+11:00", "2024-10-06T02:59:59.999999999+11:00"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			instant, err := time.Parse(time.RFC3339Nano, tt.input)
			if err != nil {
				t.Fatal(err)
			}
			c := carbon.CreateFromStdTime(instant, tt.zone).SetLocale("zh-CN").SetWeekStartsAt(carbon.Sunday)
			if c.Error != nil {
				t.Fatal(c.Error)
			}
			start, end := c.StartOfHour(), c.EndOfHour()
			if got := start.ToRfc3339NanoString(); got != tt.start {
				t.Errorf("start: got %s want %s", got, tt.start)
			}
			if got := end.ToRfc3339NanoString(); got != tt.end {
				t.Errorf("end: got %s want %s", got, tt.end)
			}
			if start.StdTime().After(c.StdTime()) || end.StdTime().Before(c.StdTime()) {
				t.Error("hour bounds exclude input instant")
			}
			if got := c.ToRfc3339NanoString(); got != tt.input {
				t.Errorf("receiver changed: %s", got)
			}
			for _, boundary := range []*carbon.Carbon{start, end} {
				if boundary.Timezone() != tt.zone || boundary.WeekStartsAt() != carbon.Sunday || boundary.ToMonthString() != c.ToMonthString() {
					t.Error("boundary lost location or instance configuration")
				}
			}
			if !start.StartOfHour().StdTime().Equal(start.StdTime()) || !end.EndOfHour().StdTime().Equal(end.StdTime()) {
				t.Error("hour boundaries are not idempotent")
			}
		})
	}
}
