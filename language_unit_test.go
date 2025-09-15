package carbon

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LanguageSuite struct {
	suite.Suite
}

func TestLanguageSuite(t *testing.T) {
	suite.Run(t, new(LanguageSuite))
}

func (s *LanguageSuite) TestLanguage_Copy() {
	s.Run("copy nil language", func() {
		oldLang := NewLanguage()
		oldLang = nil
		newCarbon := oldLang.Copy()

		s.Nil(oldLang)
		s.Nil(newCarbon)
	})

	s.Run("copy nil resources", func() {
		oldLang := NewLanguage()
		oldLang.resources = nil
		newCarbon := oldLang.Copy()
		s.Nil(oldLang.resources)
		s.Nil(newCarbon.resources)
	})

	s.Run("copy dir", func() {
		oldLang := NewLanguage()
		oldLang.dir = "lang"
		newCarbon := oldLang.Copy()
		s.Equal(oldLang.dir, newCarbon.dir)
	})

	s.Run("copy locale", func() {
		oldLang := NewLanguage()
		oldLang.locale = "en"
		newCarbon := oldLang.Copy()
		s.Equal(oldLang.locale, newCarbon.locale)
	})

	s.Run("copy resources", func() {
		oldLang := NewLanguage()
		oldLang.SetLocale("en").SetResources(map[string]string{
			"months":       "Ⅰ月|Ⅱ月|Ⅲ月|Ⅳ月|Ⅴ月|Ⅵ月|Ⅶ月|Ⅷ月|Ⅸ月|Ⅹ月|Ⅺ月|Ⅻ月",
			"short_months": "Ⅰ|Ⅱ|Ⅲ|Ⅳ|Ⅴ|Ⅵ|Ⅶ|Ⅷ|Ⅸ|Ⅹ|Ⅺ|Ⅻ",
		})
		newCarbon := oldLang.Copy()
		s.Equal(oldLang.resources, newCarbon.resources)
	})
}

func (s *LanguageSuite) TestLanguage_SetLocale() {
	s.Run("nil language", func() {
		lang := NewLanguage()
		lang = nil
		lang.SetLocale("en")
		s.Empty(Parse("2020-08-05 13:14:15").SetLanguage(lang).ToMonthString())
	})

	s.Run("error locale", func() {
		lang := NewLanguage()
		lang.SetLocale("xxx")
		s.Error(lang.Error)
		s.Empty(Parse("2020-08-05 13:14:15").SetLanguage(lang).ToMonthString())
	})

	s.Run("empty locale", func() {
		lang := NewLanguage()
		lang.SetLocale("")
		s.Error(lang.Error)
		s.Empty(Parse("2020-08-05 13:14:15").SetLanguage(lang).ToMonthString())
	})

	s.Run("valid carbon", func() {
		lang := NewLanguage()

		lang.SetLocale("en")
		s.Nil(lang.Error)
		s.Equal("Leo", Parse("2020-08-05").SetLanguage(lang).Constellation())
		s.Equal("Summer", Parse("2020-08-05").SetLanguage(lang).Season())
		s.Equal("4 years before", Parse("2020-08-05").SetLanguage(lang).DiffForHumans(Parse("2024-08-05")))
		s.Equal("August", Parse("2020-08-05").SetLanguage(lang).ToMonthString())
		s.Equal("Aug", Parse("2020-08-05").SetLanguage(lang).ToShortMonthString())
		s.Equal("Wednesday", Parse("2020-08-05").SetLanguage(lang).ToWeekString())
		s.Equal("Wed", Parse("2020-08-05").SetLanguage(lang).ToShortWeekString())

		lang.SetLocale("zh-CN")
		s.Nil(lang.Error)
		s.Equal("狮子座", Parse("2020-08-05").SetLanguage(lang).Constellation())
		s.Equal("夏季", Parse("2020-08-05").SetLanguage(lang).Season())
		s.Equal("4 年前", Parse("2020-08-05").SetLanguage(lang).DiffForHumans(Parse("2024-08-05")))
		s.Equal("八月", Parse("2020-08-05").SetLanguage(lang).ToMonthString())
		s.Equal("8月", Parse("2020-08-05").SetLanguage(lang).ToShortMonthString())
		s.Equal("星期三", Parse("2020-08-05").SetLanguage(lang).ToWeekString())
		s.Equal("周三", Parse("2020-08-05").SetLanguage(lang).ToShortWeekString())
	})

	s.Run("error language", func() {
		lang := NewLanguage()
		lang.Error = fmt.Errorf("test error")
		lang.SetLocale("en")
		s.Error(lang.Error)
	})
}

func (s *LanguageSuite) TestLanguage_SetResources() {
	s.Run("nil language", func() {
		lang := NewLanguage()
		lang = nil
		lang.SetResources(map[string]string{
			"months":       "Ⅰ月|Ⅱ月|Ⅲ月|Ⅳ月|Ⅴ月|Ⅵ月|Ⅶ月|Ⅷ月|Ⅸ月|Ⅹ月|Ⅺ月|Ⅻ月",
			"short_months": "Ⅰ|Ⅱ|Ⅲ|Ⅳ|Ⅴ|Ⅵ|Ⅶ|Ⅷ|Ⅸ|Ⅹ|Ⅺ|Ⅻ",
		})
		s.Empty(Parse("2020-08-05 13:14:15").SetLanguage(lang).ToMonthString())
	})

	s.Run("nil resources", func() {
		lang := NewLanguage()
		lang.SetResources(nil)
		s.Error(lang.Error)
		s.Empty(Parse("2020-08-05 13:14:15").SetLanguage(lang).ToMonthString())
	})

	s.Run("empty resources", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{})
		s.Error(lang.Error)
		s.Empty(Parse("2020-08-05 13:14:15").SetLanguage(lang).ToMonthString())
	})

	s.Run("set some resources", func() {
		lang := NewLanguage()
		lang.SetLocale("en").SetResources(map[string]string{
			"months":       "Ⅰ月|Ⅱ月|Ⅲ月|Ⅳ月|Ⅴ月|Ⅵ月|Ⅶ月|Ⅷ月|Ⅸ月|Ⅹ月|Ⅺ月|Ⅻ月",
			"short_months": "Ⅰ|Ⅱ|Ⅲ|Ⅳ|Ⅴ|Ⅵ|Ⅶ|Ⅷ|Ⅸ|Ⅹ|Ⅺ|Ⅻ",
		})
		s.Nil(lang.Error)

		s.Equal("Leo", Parse("2020-08-05").SetLanguage(lang).Constellation())
		s.Equal("Summer", Parse("2020-08-05").SetLanguage(lang).Season())
		s.Equal("4 years before", Parse("2020-08-05").SetLanguage(lang).DiffForHumans(Parse("2024-08-05")))
		s.Equal("Ⅷ月", Parse("2020-08-05").SetLanguage(lang).ToMonthString())
		s.Equal("Ⅷ", Parse("2020-08-05").SetLanguage(lang).ToShortMonthString())
		s.Equal("Wednesday", Parse("2020-08-05").SetLanguage(lang).ToWeekString())
		s.Equal("Wed", Parse("2020-08-05").SetLanguage(lang).ToShortWeekString())
	})

	s.Run("set all resources", func() {
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
		s.Nil(lang.Error)

		s.Equal("Leo", Parse("2020-08-05").SetLanguage(lang).Constellation())
		s.Equal("summer", Parse("2020-08-05").SetLanguage(lang).Season())
		s.Equal("4 yrs before", Parse("2020-08-05").SetLanguage(lang).DiffForHumans(Parse("2024-08-05")))
		s.Equal("August", Parse("2020-08-05").SetLanguage(lang).ToMonthString())
		s.Equal("Aug", Parse("2020-08-05").SetLanguage(lang).ToShortMonthString())
		s.Equal("Wednesday", Parse("2020-08-05").SetLanguage(lang).ToWeekString())
		s.Equal("Wed", Parse("2020-08-05").SetLanguage(lang).ToShortWeekString())
	})

	s.Run("error language", func() {
		lang := NewLanguage()
		lang.Error = fmt.Errorf("test error")
		lang.SetResources(map[string]string{"month": "1 month"})
		s.Error(lang.Error)
	})
}

func (s *LanguageSuite) TestLanguage_translate() {
	s.Run("nil language", func() {
		lang := NewLanguage()
		lang = nil
		s.Empty(lang.translate("month", 1))
	})

	s.Run("nil resources", func() {
		lang := NewLanguage()
		lang.resources = nil
		s.Empty(lang.translate("month", 1))
	})

	s.Run("empty resources", func() {
		lang := NewLanguage()
		lang.resources = make(map[string]string)
		// Empty resources will trigger SetLocale(DefaultLocale) and retry
		result := lang.translate("month", 1)
		s.NotEmpty(result) // Should get default locale result
	})

	s.Run("unit not exists", func() {
		lang := NewLanguage()
		lang.SetLocale("en")
		s.Empty(lang.translate("nonexistent", 1))
	})

	s.Run("single slice item", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{
			"month": "1 month",
		})
		s.Equal("1 month", lang.translate("month", 1))
		s.Equal("1 month", lang.translate("month", 5)) // Single slice always returns the same
	})

	s.Run("single slice item with placeholder", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{
			"month": "%d months",
		})
		s.Equal("1 months", lang.translate("month", 1))
		s.Equal("5 months", lang.translate("month", 5))
	})

	s.Run("multiple slice items - normal case", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{
			"month": "1 month|%d months",
		})
		s.Equal("1 month", lang.translate("month", 1))
		s.Equal("5 months", lang.translate("month", 5))
	})

	s.Run("multiple slice items - index out of range", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{
			"month": "1 month|%d months",
		})
		// getAbsValue(10) = 10, len(slice) = 2, so use slice[len(slice)-1]
		s.Equal("10 months", lang.translate("month", 10))
	})

	s.Run("negative value without placeholder", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{
			"month": "1 month|months",
		})
		// getAbsValue(-5) = 5, slice[5-1] = slice[4] doesn't exist, use slice[len(slice)-1]
		s.Equal("months", lang.translate("month", -5))
	})

	s.Run("negative value with placeholder", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{
			"month": "1 month|%d months",
		})
		s.Equal("-5 months", lang.translate("month", -5))
	})

	s.Run("large positive value", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{
			"month": "1 month|%d months",
		})
		s.Equal("1000 months", lang.translate("month", 1000))
	})

	s.Run("large negative value", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{
			"month": "1 month|%d months",
		})
		s.Equal("-1000 months", lang.translate("month", -1000))
	})

	s.Run("negative value in range without placeholder", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{
			"month": "1 month|months|many months",
		})
		s.Equal("-months", lang.translate("month", -2))
	})
}
