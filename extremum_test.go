package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroValue(t *testing.T) {
	assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", ZeroValue().ToString())
}

func TestEpochValue(t *testing.T) {
	assert.Equal(t, "1970-01-01 00:00:00 +0000 UTC", EpochValue().ToString())
}

func TestMaxValue(t *testing.T) {
	assert.Equal(t, "9999-12-31 23:59:59.999999999 +0000 UTC", MaxValue().ToString())
}

func TestMinValue(t *testing.T) {
	assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", MinValue().ToString())
}

func TestMaxDuration(t *testing.T) {
	assert.Equal(t, 9.223372036854776e+09, MaxDuration().Seconds())
}

func TestMinDuration(t *testing.T) {
	assert.Equal(t, -9.223372036854776e+09, MinDuration().Seconds())
}

func TestMax(t *testing.T) {
	c1 := Parse("2020-08-01")
	c2 := Parse("2020-08-03")
	c3 := Parse("2020-08-06")

	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon()
		assert.Equal(t, c3, Max(c, c1, c2, c3))
		assert.Equal(t, c3, Max(c1, c, c2, c3))
		assert.Equal(t, c3, Max(c2, c1, c, c3))
		assert.Equal(t, c3, Max(c3, c1, c2, c))
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("")
		assert.Empty(t, Max(c, c1, c2, c3).ToString())
		assert.Empty(t, Max(c1, c, c2, c3).ToString())
		assert.Empty(t, Max(c2, c1, c, c3).ToString())
		assert.Empty(t, Max(c3, c1, c2, c).ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx")
		assert.Error(t, Max(c, c1, c2, c3).Error)
		assert.Error(t, Max(c1, c, c2, c3).Error)
		assert.Error(t, Max(c2, c1, c, c3).Error)
		assert.Error(t, Max(c3, c1, c2, c).Error)
	})

	t.Run("valid carbon", func(t *testing.T) {
		assert.Equal(t, c1, Max(c1))
		assert.Equal(t, c2, Max(c1, c2))
		assert.Equal(t, c3, Max(c1, c2, c3))
	})
}

func TestMin(t *testing.T) {
	c1 := Parse("2020-08-01")
	c2 := Parse("2020-08-03")
	c3 := Parse("2020-08-06")

	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon()
		assert.Equal(t, c, Min(c, c1, c2, c3))
		assert.Equal(t, c, Min(c1, c, c2, c3))
		assert.Equal(t, c, Min(c2, c1, c, c3))
		assert.Equal(t, c, Min(c3, c1, c2, c))
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("")
		assert.Empty(t, Min(c, c1, c2, c3).ToString())
		assert.Empty(t, Min(c1, c, c2, c3).ToString())
		assert.Empty(t, Min(c2, c1, c, c3).ToString())
		assert.Empty(t, Min(c3, c1, c2, c).ToString())
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx")
		assert.Error(t, Min(c, c1, c2, c3).Error)
		assert.Error(t, Min(c1, c, c2, c3).Error)
		assert.Error(t, Min(c2, c1, c, c3).Error)
		assert.Error(t, Min(c3, c1, c2, c).Error)
	})

	t.Run("valid carbon", func(t *testing.T) {
		assert.Equal(t, c1, Min(c1))
		assert.Equal(t, c1, Min(c1, c2))
		assert.Equal(t, c1, Min(c1, c2, c3))
	})
}

func TestCarbon_Closest(t *testing.T) {
	c1 := Parse("2020-08-01")
	c2 := Parse("2020-08-03")
	c3 := Parse("2020-08-06")

	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon()
		assert.Equal(t, c1, c.Closest(c1, c2, c3))
		assert.Equal(t, c2, c1.Closest(c, c2, c3))
		assert.Equal(t, c1, c2.Closest(c1, c, c3))
		assert.Equal(t, c2, c3.Closest(c1, c2, c))
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("")
		assert.Equal(t, c, c.Closest(c1, c2, c3))
		assert.Equal(t, c, c1.Closest(c, c2, c3))
		assert.Equal(t, c, c2.Closest(c1, c, c3))
		assert.Equal(t, c, c3.Closest(c1, c2, c))
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx")
		assert.Equal(t, c, c.Closest(c1, c2, c3))
		assert.Equal(t, c, c1.Closest(c, c2, c3))
		assert.Equal(t, c, c2.Closest(c1, c, c3))
		assert.Equal(t, c, c3.Closest(c1, c2, c))
	})

	t.Run("valid carbon", func(t *testing.T) {
		c := Parse("2020-08-05")
		assert.Equal(t, c, c.Closest(c))
		assert.Equal(t, c1, c.Closest(c1))
		assert.Equal(t, c2, c.Closest(c1, c2))
		assert.Equal(t, c3, c.Closest(c1, c3))
		assert.Equal(t, c3, c.Closest(c1, c2, c3))
		assert.Equal(t, c2, c1.Closest(c, c2, c3))
		assert.Equal(t, c1, c2.Closest(c1, c, c3))
		assert.Equal(t, c, c3.Closest(c1, c2, c))
	})
}

func TestCarbon_Farthest(t *testing.T) {
	c1 := Parse("2020-08-01")
	c2 := Parse("2020-08-03")
	c3 := Parse("2020-08-06")

	t.Run("zero carbon", func(t *testing.T) {
		c := NewCarbon()
		assert.Equal(t, c3, c.Farthest(c1, c2, c3))
		assert.Equal(t, c, c1.Farthest(c, c2, c3))
		assert.Equal(t, c, c2.Farthest(c1, c, c3))
		assert.Equal(t, c, c3.Farthest(c1, c2, c))
	})

	t.Run("empty carbon", func(t *testing.T) {
		c := Parse("")
		assert.Equal(t, c, c.Farthest(c1, c2, c3))
		assert.Equal(t, c, c1.Farthest(c, c2, c3))
		assert.Equal(t, c, c2.Farthest(c1, c, c3))
		assert.Equal(t, c, c3.Farthest(c1, c2, c))
	})

	t.Run("error carbon", func(t *testing.T) {
		c := Parse("xxx")
		assert.Equal(t, c, c.Farthest(c1, c2, c3))
		assert.Equal(t, c, c1.Farthest(c, c2, c3))
		assert.Equal(t, c, c2.Farthest(c1, c, c3))
		assert.Equal(t, c, c3.Farthest(c1, c2, c))
	})

	t.Run("valid carbon", func(t *testing.T) {
		c := Parse("2020-08-05")
		assert.Equal(t, c, c.Farthest(c))
		assert.Equal(t, c1, c.Farthest(c1))
		assert.Equal(t, c1, c.Farthest(c1, c2))
		assert.Equal(t, c1, c.Farthest(c1, c3))
		assert.Equal(t, c1, c.Farthest(c1, c2, c3))
		assert.Equal(t, c3, c1.Farthest(c, c2, c3))
		assert.Equal(t, c3, c2.Farthest(c1, c, c3))
		assert.Equal(t, c1, c3.Farthest(c1, c2, c))
	})
}
