package carbon

import (
	"testing"
	"time"
)

// A zero receiver equal to an included bound must not bypass validation of
// the other bound. Keep this cross-product separate from the existing cases
// which reuse the same invalid value in every position.
func TestBetweenIncludedInvalidMixedWithZero(t *testing.T) {
	values := []struct {
		name string
		date *Carbon
	}{
		{"nil", nil},
		{"empty", Parse("")},
		{"error", Parse("invalid")},
		{"zero", NewCarbon()},
		{"past", Parse("1969-12-31", "UTC")},
		{"future", Parse("2020-08-05", "UTC")},
	}
	methods := []struct {
		name string
		call func(*Carbon, *Carbon, *Carbon) bool
	}{
		{"Between", (*Carbon).Between},
		{"BetweenIncludedStart", (*Carbon).BetweenIncludedStart},
		{"BetweenIncludedEnd", (*Carbon).BetweenIncludedEnd},
		{"BetweenIncludedBoth", (*Carbon).BetweenIncludedBoth},
	}
	for _, method := range methods {
		for _, c := range values {
			for _, start := range values {
				for _, end := range values {
					if !c.date.IsInvalid() && !start.date.IsInvalid() && !end.date.IsInvalid() {
						continue
					}
					t.Run(method.name+"/"+c.name+"/"+start.name+"/"+end.name, func(t *testing.T) {
						if method.call(c.date, start.date, end.date) {
							t.Fatal("an invalid receiver or bound must return false")
						}
					})
				}
			}
		}
	}
}

func TestBetweenIncludedValidControl(t *testing.T) {
	// Use equal instants with distinct locations and nonzero nanoseconds too.
	origin := time.Date(2020, 8, 5, 12, 0, 0, 123, time.UTC)
	values := []time.Time{
		{}, origin.Add(-time.Hour), origin,
		origin.In(time.FixedZone("offset", 7*3600)), origin.Add(time.Nanosecond),
	}
	for i, c := range values {
		for j, start := range values {
			for k, end := range values {
				cc, ss, ee := NewCarbon(c), NewCarbon(start), NewCarbon(end)
				ordered := !start.After(end)
				zeroStart := c.IsZero() && start.IsZero()
				zeroEnd := c.IsZero() && end.IsZero()
				checks := []struct {
					name string
					got  bool
					want bool
				}{
					{"start", cc.BetweenIncludedStart(ss, ee), ordered && (zeroStart || (!c.Before(start) && c.Before(end)))},
					{"end", cc.BetweenIncludedEnd(ss, ee), ordered && (zeroEnd || (c.After(start) && !c.After(end)))},
					{"both", cc.BetweenIncludedBoth(ss, ee), ordered && (zeroStart || zeroEnd || (!c.Before(start) && !c.After(end)))},
				}
				for _, check := range checks {
					if check.got != check.want {
						t.Errorf("%s (%d,%d,%d): got %t, want %t", check.name, i, j, k, check.got, check.want)
					}
				}
			}
		}
	}
}
