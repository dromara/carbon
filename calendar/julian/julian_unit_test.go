package julian

import (
	"encoding/json"
	"os"
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

	t.Run("nil julian", func(t *testing.T) {
		j := new(Julian)
		j = nil
		assert.Nil(t, j.ToGregorian())
		assert.Empty(t, j.ToGregorian().String())
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
	t.Run("nil julian", func(t *testing.T) {
		j := new(Julian)
		j = nil
		assert.Zero(t, j.JD())
	})

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
	t.Run("nil julian", func(t *testing.T) {
		j := new(Julian)
		j = nil
		assert.Zero(t, j.MJD())
	})

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

// TestJulian_AuthorityData tests against authoritative data from Python convertdate library
func TestJulian_AuthorityData(t *testing.T) {
	file, err := os.Open("julian_test_data.json")
	if err != nil {
		t.Fatalf("failed to open test data file: %v", err)
	}
	defer func() { _ = file.Close() }()

	type testCase struct {
		Description string  `json:"description"`
		Julian      float64 `json:"julian"`
		Gregorian   struct {
			Year  int `json:"year"`
			Month int `json:"month"`
			Day   int `json:"day"`
		} `json:"gregorian"`
	}

	var cases []testCase
	dec := json.NewDecoder(file)
	if err := dec.Decode(&cases); err != nil {
		t.Fatalf("failed to decode test data: %v", err)
	}

	for _, c := range cases {
		// Julian to Gregorian conversion
		j := NewJulian(c.Julian)
		g := j.ToGregorian()
		if g.Time.IsZero() {
			t.Errorf("[%s] Julian->Gregorian failed: JDN %.1f", c.Description, c.Julian)
		} else {
			gy, gm, gd := g.Time.Date()
			if gy != c.Gregorian.Year || int(gm) != c.Gregorian.Month || gd != c.Gregorian.Day {
				t.Errorf("[%s] Julian->Gregorian error: expected %04d-%02d-%02d, got %04d-%02d-%02d",
					c.Description, c.Gregorian.Year, c.Gregorian.Month, c.Gregorian.Day, gy, int(gm), gd)
			}
		}

		// Gregorian to Julian conversion
		gt := time.Date(c.Gregorian.Year, time.Month(c.Gregorian.Month), c.Gregorian.Day, 0, 0, 0, 0, time.UTC)
		j2 := FromStdTime(gt)
		if j2 == nil {
			t.Errorf("[%s] Gregorian->Julian failed: %04d-%02d-%02d",
				c.Description, c.Gregorian.Year, c.Gregorian.Month, c.Gregorian.Day)
		} else {
			// Allow small floating point differences (within 0.1)
			expectedJDN := c.Julian
			actualJDN := j2.JD()
			if abs(actualJDN-expectedJDN) > 0.1 {
				t.Errorf("[%s] Gregorian->Julian error: expected JDN %.1f, got %.1f",
					c.Description, expectedJDN, actualJDN)
			}
		}
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
