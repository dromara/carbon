package carbon

import "time"

const (
	minDuration time.Duration = -1 << 63
	maxDuration time.Duration = 1<<63 - 1
)

// MaxValue returns a Carbon instance for the greatest supported date.
// 返回 Carbon 的最大值
func MaxValue() Carbon {
	c := NewCarbon()
	c.time = time.Date(9999, time.December, 31, 23, 59, 59, 999999999, time.UTC)
	return c
}

// MinValue returns a Carbon instance for the lowest supported date.
// 返回 Carbon 的最小值
func MinValue() Carbon {
	c := NewCarbon()
	c.time = time.Date(-9998, time.January, 1, 0, 0, 0, 0, time.UTC)
	return c
}

// MaxDuration returns the maximum duration value.
// 返回 Duration 的最大值
func MaxDuration() time.Duration {
	return maxDuration
}

// MinDuration returns the minimum duration value.
// 返回 Duration 的最小值
func MinDuration() time.Duration {
	return minDuration
}

// Max returns the maximum Carbon instance from the given Carbon instance (second-precision).
// 返回最大的 Carbon 实例
func Max(c1 Carbon, c2 ...Carbon) (c Carbon) {
	c = c1
	for i := range c2 {
		if c.IsInvalid() || c2[i].IsInvalid() {
			c.isNil = true
			return c
		}
		if c2[i].Gte(c) {
			c = c2[i]
		}
	}
	return
}

// Min returns the minimum Carbon instance from the given Carbon instance (second-precision).
// 返回最小的 Carbon 实例
func Min(c1 Carbon, c2 ...Carbon) (c Carbon) {
	c = c1
	for i := range c2 {
		if c.IsInvalid() || c2[i].IsInvalid() {
			c.isNil = true
			return c
		}
		if c2[i].Lte(c) {
			c = c2[i]
		}
	}
	return
}

// Closest returns the closest Carbon instance from the given Carbon instance.
// 返回离给定 carbon 实例最近的 Carbon 实例
func (c Carbon) Closest(c1 Carbon, c2 Carbon) Carbon {
	if c.IsInvalid() || c1.IsInvalid() || c2.IsInvalid() {
		c.isNil = true
		return c
	}
	if c.DiffAbsInSeconds(c1) < c.DiffAbsInSeconds(c2) {
		return c1
	}
	return c2
}

// Farthest returns the farthest Carbon instance from the given Carbon instance.
// 返回离给定 carbon 实例最远的 Carbon 实例
func (c Carbon) Farthest(c1 Carbon, c2 Carbon) Carbon {
	if c.IsInvalid() || c1.IsInvalid() || c2.IsInvalid() {
		c.isNil = true
		return c
	}
	if c.DiffAbsInSeconds(c1) > c.DiffAbsInSeconds(c2) {
		return c1
	}
	return c2
}
