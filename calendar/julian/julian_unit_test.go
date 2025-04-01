package julian

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFromStdTime(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 1721423.5, FromStdTime(time.Time{}).JD())
		assert.Equal(t, float64(-678577), FromStdTime(time.Time{}).MJD())
	})

	t.Run("valid time", func(t *testing.T) {
		j := FromStdTime(time.Date(1800, 1, 1, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, 2378496.5, j.JD())
		assert.Equal(t, float64(-21504), j.MJD())
	})
}

func TestJulian_ToGregorian(t *testing.T) {
	t.Run("zero julian", func(t *testing.T) {
		assert.Equal(t, "-4712-01-01 12:00:00 +0000 UTC", new(Julian).ToGregorian().String())
		assert.Equal(t, "-4712-01-01 12:00:00 +0000 UTC", NewJulian(0).ToGregorian().String())
	})

	t.Run("invalid timezone", func(t *testing.T) {
		g := NewJulian(2460332.5).ToGregorian("xxx")
		assert.Error(t, g.Error)
		assert.Empty(t, g.String())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "2024-01-23 00:00:00 +0000 UTC", NewJulian(2460332.5).ToGregorian().String())
		assert.Equal(t, "2024-01-23 00:00:00 +0000 UTC", NewJulian(60332).ToGregorian().String())

		assert.Equal(t, "2024-01-23 12:00:00 +0000 UTC", NewJulian(2460333).ToGregorian().String())
		assert.Equal(t, "2024-01-23 12:00:00 +0000 UTC", NewJulian(60332.5).ToGregorian().String())

		assert.Equal(t, "2024-01-23 13:14:15 +0000 UTC", NewJulian(2460333.051563).ToGregorian().String())
		assert.Equal(t, "2024-01-23 13:14:15 +0000 UTC", NewJulian(60332.551563).ToGregorian().String())

		assert.Equal(t, "2023-10-15 12:00:00 +0000 UTC", NewJulian(60232.5).ToGregorian().String())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "2024-01-23 00:00:00 +0800 CST", NewJulian(2460332.5).ToGregorian("PRC").String())
		assert.Equal(t, "2024-01-23 00:00:00 +0800 CST", NewJulian(60332).ToGregorian("PRC").String())

		assert.Equal(t, "2024-01-23 12:00:00 +0800 CST", NewJulian(60332.5).ToGregorian("PRC").String())
		assert.Equal(t, "2024-01-23 12:00:00 +0800 CST", NewJulian(2460333).ToGregorian("PRC").String())
	})
}

func TestJulian_JD(t *testing.T) {
	t.Run("zero julian", func(t *testing.T) {
		assert.Zero(t, new(Julian).JD())
		assert.Zero(t, NewJulian(0).JD())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, 2378496.5, FromStdTime(time.Date(1800, 1, 1, 0, 0, 0, 0, time.UTC)).JD(4))
		assert.Equal(t, 2378496.5, FromStdTime(time.Date(1800, 1, 1, 0, 0, 0, 0, time.UTC)).JD(6))

		assert.Equal(t, 2460333.0516, FromStdTime(time.Date(2024, 1, 23, 13, 14, 15, 0, time.UTC)).JD(4))
		assert.Equal(t, 2460333.051563, FromStdTime(time.Date(2024, 1, 23, 13, 14, 15, 0, time.UTC)).JD(6))
	})
}

func TestJulian_MJD(t *testing.T) {
	t.Run("zero julian", func(t *testing.T) {
		assert.Zero(t, new(Julian).MJD())
		assert.Zero(t, NewJulian(0).MJD())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, float64(-21504), FromStdTime(time.Date(1800, 1, 1, 0, 0, 0, 0, time.UTC)).MJD(4))
		assert.Equal(t, float64(-21504), FromStdTime(time.Date(1800, 1, 1, 0, 0, 0, 0, time.UTC)).MJD(6))

		assert.Equal(t, 60332.5516, FromStdTime(time.Date(2024, 1, 23, 13, 14, 15, 0, time.UTC)).MJD(4))
		assert.Equal(t, 60332.551563, FromStdTime(time.Date(2024, 1, 23, 13, 14, 15, 0, time.UTC)).MJD(6))
	})
}
