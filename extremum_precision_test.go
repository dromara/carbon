package carbon

import (
	"testing"
	"time"
)

func TestExtremumPrecision(t *testing.T) {
	base := time.Date(2026, 1, 2, 3, 4, 5, 900000000, time.UTC)
	for _, tc := range []struct {
		name          string
		first, second time.Duration
	}{
		{"same second future", 90 * time.Millisecond, 10 * time.Millisecond},
		{"same second past", -90 * time.Millisecond, -10 * time.Millisecond},
		{"cross second future", 200 * time.Millisecond, 100 * time.Millisecond},
		{"opposite sides", -800 * time.Millisecond, 200 * time.Millisecond},
		{"whole second tie", 1900 * time.Millisecond, -1100 * time.Millisecond},
		{"nanosecond", 2 * time.Nanosecond, time.Nanosecond},
		{"zero distance", time.Nanosecond, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := NewCarbon(base)
			far := NewCarbon(base.Add(tc.first))
			near := NewCarbon(base.Add(tc.second))
			if got := c.Closest(far, near); got != near {
				t.Errorf("Closest chose %s instead of %s", got.StdTime(), near.StdTime())
			}
			if got := c.Farthest(near, far); got != far {
				t.Errorf("Farthest chose %s instead of %s", got.StdTime(), far.StdTime())
			}
		})
	}
}

func TestExtremumDistanceControls(t *testing.T) {
	base := time.Date(2000, 1, 1, 0, 0, 0, 900000000, time.UTC)
	c := NewCarbon(base)
	past := NewCarbon(base.Add(-200 * time.Millisecond))
	future := NewCarbon(base.Add(200 * time.Millisecond))
	for _, pair := range [][2]*Carbon{{past, future}, {future, past}} {
		if c.Closest(pair[0], pair[1]) != pair[0] || c.Farthest(pair[0], pair[1]) != pair[0] {
			t.Fatal("an exact distance tie must retain the first candidate")
		}
	}
	far := NewCarbon(time.Date(4000, 1, 1, 0, 0, 0, 0, time.UTC))
	near := NewCarbon(time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC))
	if c.Closest(far, near) != near || c.Farthest(near, far) != far {
		t.Fatal("distances beyond time.Duration must remain ordered")
	}
	shifted := NewCarbon(base.Add(200 * time.Millisecond).In(time.FixedZone("offset", 3723)))
	if c.Closest(shifted, past) != shifted || c.Farthest(past, shifted) != past {
		t.Fatal("different locations must not break exact distance ties")
	}
}
