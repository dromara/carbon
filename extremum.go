package carbon

import (
	"time"
)

const (
	minDuration Duration = -1 << 63
	maxDuration Duration = 1<<63 - 1
)

// ZeroValue returns the zero value of Carbon instance.
func ZeroValue() *Carbon {
	return MinValue()
}

// EpochValue returns the unix epoch value of Carbon instance.
func EpochValue() *Carbon {
	return NewCarbon(time.Date(EpochYear, time.January, MinDay, MinHour, MinMinute, MinSecond, MinNanosecond, time.UTC))
}

// MaxValue returns the maximum value of Carbon instance.
func MaxValue() *Carbon {
	return NewCarbon(time.Date(MaxYear, time.December, MaxDay, MaxHour, MaxMinute, MaxSecond, MaxNanosecond, time.UTC))
}

// MinValue returns the minimum value of Carbon instance.
func MinValue() *Carbon {
	return NewCarbon(time.Date(MinYear, time.January, MinDay, MinHour, MinMinute, MinSecond, MinNanosecond, time.UTC))
}

// MaxDuration returns the maximum value of duration instance.
func MaxDuration() Duration {
	return maxDuration
}

// MinDuration returns the minimum value of duration instance.
func MinDuration() Duration {
	return minDuration
}

// Max returns the maximum Carbon instance from some given Carbon instances.
func Max(c1 *Carbon, c2 ...*Carbon) (c *Carbon) {
	// If first carbon is invalid, return it immediately
	if c1.IsInvalid() {
		return c1
	}

	c = c1
	// If no additional arguments, return the first one
	if len(c2) == 0 {
		return
	}

	// Check all additional arguments
	for _, carbon := range c2 {
		// If any carbon is invalid, return it immediately
		if carbon.IsInvalid() {
			return carbon
		}
		// Update maximum if current carbon is greater or equal
		if carbon.Gte(c) {
			c = carbon
		}
	}
	return
}

// Min returns the minimum Carbon instance from some given Carbon instances.
func Min(c1 *Carbon, c2 ...*Carbon) (c *Carbon) {
	// If first carbon is invalid, return it immediately
	if c1.IsInvalid() {
		return c1
	}

	c = c1
	// If no additional arguments, return the first one
	if len(c2) == 0 {
		return
	}

	// Check all additional arguments
	for _, carbon := range c2 {
		// If any carbon is invalid, return it immediately
		if carbon.IsInvalid() {
			return carbon
		}
		// Update minimum if current carbon is less or equal
		if carbon.Lte(c) {
			c = carbon
		}
	}
	return
}

// Closest returns the closest Carbon instance from some given Carbon instances.
func (c *Carbon) Closest(c1 *Carbon, c2 ...*Carbon) *Carbon {
	// Validate the base carbon instance
	if c.IsInvalid() {
		return c
	}

	// Validate the first comparison instance
	if c1.IsInvalid() {
		return c1
	}

	// If no additional arguments, return the first one
	if len(c2) == 0 {
		return c1
	}

	// Find the closest among all instances
	closest := c1
	minDiff := c.DiffAbsInSeconds(closest)

	// Check all additional arguments
	for _, arg := range c2 {
		// Validate each argument
		if arg.IsInvalid() {
			return arg
		}

		// Calculate difference and update if closer
		if diff := c.DiffAbsInSeconds(arg); diff < minDiff {
			minDiff = diff
			closest = arg
		}
	}
	return closest
}

// Farthest returns the farthest Carbon instance from some given Carbon instances.
func (c *Carbon) Farthest(c1 *Carbon, c2 ...*Carbon) *Carbon {
	// Validate the base carbon instance
	if c.IsInvalid() {
		return c
	}

	// Validate the first comparison instance
	if c1.IsInvalid() {
		return c1
	}

	// If no additional arguments, return the first one
	if len(c2) == 0 {
		return c1
	}

	// Find the farthest among all instances
	farthest := c1
	maxDiff := c.DiffAbsInSeconds(farthest)

	// Check all additional arguments
	for _, arg := range c2 {
		// Validate each argument
		if arg.IsInvalid() {
			return arg
		}

		// Calculate difference and update if farther
		if diff := c.DiffAbsInSeconds(arg); diff > maxDiff {
			maxDiff = diff
			farthest = arg
		}
	}
	return farthest
}
