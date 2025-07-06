package carbon

import "time"

const (
	minDuration Duration = -1 << 63
	maxDuration Duration = 1<<63 - 1
)

// ZeroValue returns the zero value of Carbon instance.
func ZeroValue() Carbon {
	return MinValue()
}

// EpochValue returns the unix epoch value of Carbon instance.
func EpochValue() Carbon {
	return NewCarbon(time.Date(EpochYear, MinMonth, MinDay, MinHour, MinMinute, MinSecond, MinNanosecond, time.UTC))
}

// MaxValue returns the maximum value of Carbon instance.
func MaxValue() Carbon {
	c := NewCarbon()
	c.time = time.Date(MaxYear, time.December, MaxDay, MaxHour, MaxMinute, MaxSecond, MaxNanosecond, time.UTC)
	return c
}

// MinValue returns the minimum value of Carbon instance.
func MinValue() Carbon {
	c := NewCarbon()
	c.time = time.Date(MinYear, time.January, MinDay, MinHour, MinMinute, MinSecond, MinNanosecond, time.UTC)
	return c
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
func Max(c1 Carbon, c2 ...Carbon) (c Carbon) {
	c = c1
	if c.IsInvalid() {
		return
	}
	if len(c2) == 0 {
		return
	}
	for i := range c2 {
		if c2[i].IsInvalid() {
			return c2[i]
		}
		if c2[i].Gte(c) {
			c = c2[i]
		}
	}
	return
}

// Min returns the minimum Carbon instance from some given Carbon instances.
func Min(c1 Carbon, c2 ...Carbon) (c Carbon) {
	c = c1
	if c.IsInvalid() {
		return
	}
	if len(c2) == 0 {
		return
	}
	for i := range c2 {
		if c2[i].IsInvalid() {
			return c2[i]
		}
		if c2[i].Lte(c) {
			c = c2[i]
		}
	}
	return
}

// Closest returns the closest Carbon instance from some given Carbon instances.
func (c Carbon) Closest(c1 Carbon, c2 ...Carbon) Carbon {
	if c.IsInvalid() {
		return c
	}
	if c1.IsInvalid() {
		return c1
	}
	if len(c2) == 0 {
		return c1
	}
	args := append([]Carbon{c1}, c2...)
	for i := range args {
		if args[i].IsInvalid() {
			return args[i]
		}
	}
	closest := args[0]
	minDiff := c.DiffAbsInSeconds(closest)
	for i := range args[1:] {
		arg := args[1:][i]
		diff := c.DiffAbsInSeconds(arg)
		if diff < minDiff {
			minDiff = diff
			closest = arg
		}
	}
	return closest
}

// Farthest returns the farthest Carbon instance from some given Carbon instances.
func (c Carbon) Farthest(c1 Carbon, c2 ...Carbon) Carbon {
	if c.IsInvalid() {
		return c
	}
	if c1.IsInvalid() {
		return c1
	}
	if len(c2) == 0 {
		return c1
	}
	args := append([]Carbon{c1}, c2...)
	for i := range args {
		if args[i].IsInvalid() {
			return args[i]
		}
	}
	farthest := args[0]
	maxDiff := c.DiffAbsInSeconds(farthest)
	for i := range args[1:] {
		arg := args[1:][i]
		diff := c.DiffAbsInSeconds(arg)
		if diff > maxDiff {
			maxDiff = diff
			farthest = arg
		}
	}
	return farthest
}
