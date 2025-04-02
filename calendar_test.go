package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Lunar(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Empty(t, NewCarbon().Lunar().String())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").Lunar().String())
		assert.Empty(t, Parse("0").Lunar().String())
		assert.Empty(t, Parse("xxx").Lunar().String())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2023-12-08", Parse("2024-01-18", PRC).Lunar().String())
		assert.Equal(t, "2023-12-11", Parse("2024-01-21", PRC).Lunar().String())
		assert.Equal(t, "2023-12-14", Parse("2024-01-24", PRC).Lunar().String())
	})
}

func TestCreateFromLunar(t *testing.T) {
	t.Run("invalid lunar", func(t *testing.T) {
		assert.Empty(t, CreateFromLunar(2200, 12, 14, false).ToString())
	})

	t.Run("valid lunar", func(t *testing.T) {
		assert.Equal(t, "2024-01-21 00:00:00 +0800 CST", CreateFromLunar(2023, 12, 11, false).ToString(PRC))
		assert.Equal(t, "2024-01-18 00:00:00 +0800 CST", CreateFromLunar(2023, 12, 8, false).ToString(PRC))
		assert.Equal(t, "2024-01-24 00:00:00 +0800 CST", CreateFromLunar(2023, 12, 14, false).ToString(PRC))
	})
}

func TestCarbon_Julian(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, 1.7214235e+06, NewCarbon().Julian().JD())
		assert.Equal(t, float64(-678577), NewCarbon().Julian().MJD())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Zero(t, Parse("").Julian().JD())
		assert.Zero(t, Parse("0").Julian().JD())
		assert.Zero(t, Parse("xxx").Julian().JD())

		assert.Zero(t, Parse("").Julian().MJD())
		assert.Zero(t, Parse("0").Julian().MJD())
		assert.Zero(t, Parse("xxx").Julian().MJD())
	})

	t.Run("valid time", func(t *testing.T) {
		j := Parse("2024-01-23 13:14:15").Julian()

		assert.Equal(t, 2460333.051563, j.JD())
		assert.Equal(t, 60332.551563, j.MJD())

		assert.Equal(t, 2460333.0516, j.JD(4))
		assert.Equal(t, 60332.5516, j.MJD(4))

		assert.Equal(t, 2460333.0515625, j.JD(7))
		assert.Equal(t, 60332.5515625, j.MJD(7))
	})
}

func TestCreateFromJulian(t *testing.T) {
	t.Run("invalid julian", func(t *testing.T) {
		assert.Equal(t, "-4712-01-01 12:00:00", CreateFromJulian(0).String())
		assert.Equal(t, "-4712-01-01 12:00:00", CreateFromJulian(-1).String())
	})

	t.Run("valid julian", func(t *testing.T) {
		assert.Equal(t, "2024-01-23 13:14:15 +0000 UTC", CreateFromJulian(2460333.051563).ToString())
		assert.Equal(t, "2024-01-23 13:14:15 +0000 UTC", CreateFromJulian(60332.551563).ToString())

		assert.Equal(t, "2024-01-23 13:14:18 +0000 UTC", CreateFromJulian(2460333.0516).ToString())
		assert.Equal(t, "2024-01-23 13:14:18 +0000 UTC", CreateFromJulian(60332.5516).ToString())

		assert.Equal(t, "2024-01-23 13:14:15 +0000 UTC", CreateFromJulian(2460333.0515625).ToString())
		assert.Equal(t, "2024-01-23 13:14:15 +0000 UTC", CreateFromJulian(60332.5515625).ToString())
	})
}

func TestCarbon_Persian(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Empty(t, NewCarbon().Persian().String())
		assert.Empty(t, NewCarbon().Persian().String())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").Persian().String())
		assert.Empty(t, Parse("0").Persian().String())
		assert.Empty(t, Parse("xxx").Persian().String())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "1178-10-11", Parse("1800-01-01 00:00:00").Persian().String())
		assert.Equal(t, "1399-05-15", Parse("2020-08-05 13:14:15").Persian().String())
		assert.Equal(t, "1402-10-11", Parse("2024-01-01 00:00:00").Persian().String())
	})
}

func TestCreateFromPersian(t *testing.T) {
	t.Run("invalid persian", func(t *testing.T) {
		assert.Empty(t, CreateFromPersian(9999, 12, 14).ToDateTimeString())
	})

	t.Run("valid persian", func(t *testing.T) {
		assert.Equal(t, "1800-01-01 00:00:00", CreateFromPersian(1178, 10, 11).ToDateTimeString())
		assert.Equal(t, "2024-01-01 00:00:00", CreateFromPersian(1402, 10, 11).ToDateTimeString())
		assert.Equal(t, "2024-08-05 00:00:00", CreateFromPersian(1403, 5, 15).ToDateTimeString())
	})
}
