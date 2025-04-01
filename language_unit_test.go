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
