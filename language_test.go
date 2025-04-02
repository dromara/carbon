package carbon

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguage_SetLocale(t *testing.T) {
	t.Run("nil language", func(t *testing.T) {
		lang := NewLanguage()
		lang = nil
		lang.SetLocale("en")
		assert.Empty(t, Parse("2020-08-05 13:14:15").SetLanguage(lang).ToMonthString())
	})

	t.Run("invalid locale", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetLocale("xxx")
		assert.Empty(t, Parse("2020-08-05 13:14:15").SetLanguage(lang).ToMonthString())
	})

	t.Run("empty locale", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetLocale("")
		fmt.Println("lang", lang.locale)
		assert.Empty(t, Parse("2020-08-05 13:14:15").SetLanguage(lang).ToMonthString())
	})

	t.Run("valid time", func(t *testing.T) {
		lang := NewLanguage()

		lang.SetLocale("en")
		assert.Equal(t, "Leo", Parse("2020-08-05").SetLanguage(lang).Constellation())
		assert.Equal(t, "Summer", Parse("2020-08-05").SetLanguage(lang).Season())
		assert.Equal(t, "4 years before", Parse("2020-08-05").SetLanguage(lang).DiffForHumans(Parse("2024-08-05")))
		assert.Equal(t, "August", Parse("2020-08-05").SetLanguage(lang).ToMonthString())
		assert.Equal(t, "Aug", Parse("2020-08-05").SetLanguage(lang).ToShortMonthString())
		assert.Equal(t, "Wednesday", Parse("2020-08-05").SetLanguage(lang).ToWeekString())
		assert.Equal(t, "Wed", Parse("2020-08-05").SetLanguage(lang).ToShortWeekString())

		lang.SetLocale("zh-CN")
		assert.Equal(t, "狮子座", Parse("2020-08-05").SetLanguage(lang).Constellation())
		assert.Equal(t, "夏季", Parse("2020-08-05").SetLanguage(lang).Season())
		assert.Equal(t, "4 年前", Parse("2020-08-05").SetLanguage(lang).DiffForHumans(Parse("2024-08-05")))
		assert.Equal(t, "八月", Parse("2020-08-05").SetLanguage(lang).ToMonthString())
		assert.Equal(t, "8月", Parse("2020-08-05").SetLanguage(lang).ToShortMonthString())
		assert.Equal(t, "星期三", Parse("2020-08-05").SetLanguage(lang).ToWeekString())
		assert.Equal(t, "周三", Parse("2020-08-05").SetLanguage(lang).ToShortWeekString())
	})
}

func TestLanguage_SetResources(t *testing.T) {
	t.Run("nil language", func(t *testing.T) {
		lang := NewLanguage()
		lang = nil
		lang.SetResources(nil)
		assert.Empty(t, Parse("2020-08-05 13:14:15").SetLanguage(lang).ToMonthString())
	})

	t.Run("nil resources", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetResources(nil)
		assert.Empty(t, Parse("2020-08-05 13:14:15").SetLanguage(lang).ToMonthString())
	})

	t.Run("empty resources", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetResources(map[string]string{})
		assert.Equal(t, August, Parse("2020-08-05 13:14:15").SetLanguage(lang).ToMonthString())
	})

	t.Run("invalid resources", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetResources(map[string]string{
			"xxx": "xxx",
		})
		fmt.Println("lang", lang.locale)
		assert.Empty(t, Parse("2020-08-05 13:14:15").SetLanguage(lang).ToMonthString())
	})

	t.Run("set some resources", func(t *testing.T) {
		resources := map[string]string{
			"months":       "Ⅰ月|Ⅱ月|Ⅲ月|Ⅳ月|Ⅴ月|Ⅵ月|Ⅶ月|Ⅷ月|Ⅸ月|Ⅹ月|Ⅺ月|Ⅻ月",
			"short_months": "Ⅰ|Ⅱ|Ⅲ|Ⅳ|Ⅴ|Ⅵ|Ⅶ|Ⅷ|Ⅸ|Ⅹ|Ⅺ|Ⅻ",
		}

		lang := NewLanguage()
		lang.SetLocale("en").SetResources(resources)

		assert.Equal(t, "Leo", Parse("2020-08-05").SetLanguage(lang).Constellation())
		assert.Equal(t, "Summer", Parse("2020-08-05").SetLanguage(lang).Season())
		assert.Equal(t, "4 years before", Parse("2020-08-05").SetLanguage(lang).DiffForHumans(Parse("2024-08-05")))
		assert.Equal(t, "Ⅷ月", Parse("2020-08-05").SetLanguage(lang).ToMonthString())
		assert.Equal(t, "Ⅷ", Parse("2020-08-05").SetLanguage(lang).ToShortMonthString())
		assert.Equal(t, "Wednesday", Parse("2020-08-05").SetLanguage(lang).ToWeekString())
		assert.Equal(t, "Wed", Parse("2020-08-05").SetLanguage(lang).ToShortWeekString())
	})

	t.Run("set all resources", func(t *testing.T) {
		resources := map[string]string{
			"constellations": "Aries|Taurus|Gemini|Cancer|Leo|Virgo|Libra|Scorpio|Sagittarius|Capricorn|Aquarius|Pisces",
			"seasons":        "spring|summer|autumn|winter",
			"months":         "January|February|March|April|May|June|July|August|September|October|November|December",
			"short_months":   "Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec",
			"weeks":          "Sunday|Monday|Tuesday|Wednesday|Thursday|Friday|Saturday",
			"short_weeks":    "Sun|Mon|Tue|Wed|Thu|Fri|Sat",
			"year":           "1 yr|%d yrs",
			"month":          "1 mo|%d mos",
			"week":           "%dw",
			"day":            "%dd",
			"hour":           "%dh",
			"minute":         "%dm",
			"second":         "%ds",
			"now":            "just now",
			"ago":            "%s ago",
			"from_now":       "in %s",
			"before":         "%s before",
			"after":          "%s after",
		}

		lang := NewLanguage()
		lang.SetResources(resources)

		assert.Equal(t, "Leo", Parse("2020-08-05").SetLanguage(lang).Constellation())
		assert.Equal(t, "summer", Parse("2020-08-05").SetLanguage(lang).Season())
		assert.Equal(t, "4 yrs before", Parse("2020-08-05").SetLanguage(lang).DiffForHumans(Parse("2024-08-05")))
		assert.Equal(t, "August", Parse("2020-08-05").SetLanguage(lang).ToMonthString())
		assert.Equal(t, "Aug", Parse("2020-08-05").SetLanguage(lang).ToShortMonthString())
		assert.Equal(t, "Wednesday", Parse("2020-08-05").SetLanguage(lang).ToWeekString())
		assert.Equal(t, "Wed", Parse("2020-08-05").SetLanguage(lang).ToShortWeekString())
	})
}

func TestLanguage_translate(t *testing.T) {
	t.Run("nil language", func(t *testing.T) {
		lang := NewLanguage()
		lang = nil
		assert.Empty(t, lang.translate("month", 1))
	})

	t.Run("nil resources", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetResources(nil)
		assert.Empty(t, lang.translate("month", 1))
	})

	t.Run("empty resources", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetResources(map[string]string{})
		fmt.Println("lang", lang.locale)
		assert.Equal(t, "1 month", lang.translate("month", 1))
	})

	t.Run("invalid resources", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetResources(map[string]string{
			"xxx": "xxx",
		})
		fmt.Println("lang", lang.locale)
		assert.Empty(t, lang.translate("month", 1))
	})

	t.Run("valid resources", func(t *testing.T) {
		lang := NewLanguage()
		lang.SetLocale("en")
		fmt.Println("lang", lang.locale)
		assert.Equal(t, "1 month", lang.translate("month", 1))
		assert.Equal(t, "-1 month", lang.translate("month", -1))
	})
}
