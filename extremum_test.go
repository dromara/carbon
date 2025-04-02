package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxValue(t *testing.T) {
	assert.Equal(t, "9999-12-31 23:59:59.999999999 +0000 UTC", MaxValue().ToString())
}

func TestMinValue(t *testing.T) {
	assert.Equal(t, "-9998-01-01 00:00:00 +0000 UTC", MinValue().ToString())
}

func TestMaxDuration(t *testing.T) {
	assert.Equal(t, 9.223372036854776e+09, MaxDuration().Seconds())
}

func TestMinDuration(t *testing.T) {
	assert.Equal(t, -9.223372036854776e+09, MinDuration().Seconds())
}

func TestMax(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", Max(NewCarbon(), NewCarbon()).ToString())
		assert.Equal(t, "2021-08-05 00:00:00 +0000 UTC", Max(NewCarbon(), Parse("2021-08-05")).ToString())
		assert.Equal(t, "2021-08-05 00:00:00 +0000 UTC", Max(Parse("2021-08-05"), NewCarbon()).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Max(Parse(""), Tomorrow()).ToString())
		assert.Empty(t, Max(Parse("0"), Tomorrow()).ToString())
		assert.Empty(t, Max(Now(), Parse("xxx")).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2021-08-05 00:00:00 +0000 UTC", Max(Parse("2020-08-06"), Parse("2021-08-05")).ToString())
	})
}

func TestMin(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", Min(NewCarbon(), NewCarbon()).ToString())
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", Min(NewCarbon(), Parse("2021-08-05")).ToString())
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", Min(Parse("2021-08-05"), NewCarbon()).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Min(Parse(""), Tomorrow()).ToString())
		assert.Empty(t, Min(Parse("0"), Tomorrow()).ToString())
		assert.Empty(t, Min(Now(), Parse("xxx")).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-06 00:00:00 +0000 UTC", Min(Parse("2020-08-06"), Parse("2021-08-05")).ToString())
	})
}

func TestCarbon_Closest(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", NewCarbon().Closest(NewCarbon(), NewCarbon()).ToString())
		assert.Equal(t, "2021-08-05 00:00:00 +0000 UTC", Parse("2020-08-05").Closest(NewCarbon(), Parse("2021-08-05")).ToString())
		assert.Equal(t, "2021-08-05 00:00:00 +0000 UTC", Parse("2020-08-05").Closest(Parse("2021-08-05"), NewCarbon()).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").Closest(Now(), Tomorrow()).ToString())
		assert.Empty(t, Parse("0").Closest(Now(), Tomorrow()).ToString())
		assert.Empty(t, Parse("xxx").Closest(Now(), Tomorrow()).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-06 00:00:00 +0000 UTC", Parse("2020-08-05").Closest(Parse("2020-08-06"), Parse("2021-08-05")).ToString())
	})
}

func TestCarbon_Farthest(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", NewCarbon().Farthest(NewCarbon(), NewCarbon()).ToString())
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", Parse("2020-08-05").Farthest(NewCarbon(), Parse("2021-08-05")).ToString())
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", Parse("2020-08-05").Farthest(Parse("2021-08-05"), NewCarbon()).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").Farthest(Now(), Tomorrow()).ToString())
		assert.Empty(t, Parse("0").Farthest(Now(), Tomorrow()).ToString())
		assert.Empty(t, Parse("xxx").Farthest(Now(), Tomorrow()).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2021-08-05 00:00:00 +0000 UTC", Parse("2020-08-05").Farthest(Parse("2020-08-06"), Parse("2021-08-05")).ToString())
	})
}
