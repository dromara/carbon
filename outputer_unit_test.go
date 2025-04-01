package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_String(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00", NewCarbon().String())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").String())
		assert.Empty(t, Parse("0").String())
		assert.Empty(t, Parse("xxx").String())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15", Parse("2020-08-05 13:14:15").String())
		assert.Equal(t, "2020-08-05", Parse("2020-08-05 13:14:15").SetLayout(DateLayout).String())
	})
}

func TestCarbon_GoString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)", NewCarbon().GoString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").GoString())
		assert.Empty(t, Parse("0").GoString())
		assert.Empty(t, Parse("xxx").GoString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "time.Date(2020, time.August, 5, 13, 14, 15, 0, time.UTC)", Parse("2020-08-05 13:14:15").GoString())
		assert.Equal(t, "time.Date(2020, time.August, 5, 13, 14, 15, 0, time.Location(\"PRC\"))", Parse("2020-08-05 13:14:15", PRC).GoString())
	})
}

func TestCarbon_ToString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", NewCarbon().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToString())
		assert.Empty(t, Parse("0").ToString())
		assert.Empty(t, Parse("xxx").ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", Parse("2020-08-05 13:14:15").ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", Parse("2020-08-05 13:14:15", PRC).ToString())
		assert.Equal(t, "2020-08-05 21:14:15 +0800 CST", Parse("2020-08-05 13:14:15").ToString(PRC))
	})
}

func TestCarbon_ToMonthString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, January, NewCarbon().ToMonthString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToMonthString())
		assert.Empty(t, Parse("0").ToMonthString())
		assert.Empty(t, Parse("xxx").ToMonthString())
	})

	t.Run("invalid resources", func(t *testing.T) {
		lang := NewLanguage()
		resources := map[string]string{
			"months": "xxx",
		}
		lang.SetResources(resources)
		c := Parse("2020-01-05").SetLanguage(lang)
		assert.Empty(t, c.ToMonthString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, January, Parse("2020-01-05").ToMonthString())
		assert.Equal(t, February, Parse("2020-02-05").ToMonthString())
		assert.Equal(t, March, Parse("2020-03-05").ToMonthString())
		assert.Equal(t, April, Parse("2020-04-05").ToMonthString())
		assert.Equal(t, May, Parse("2020-05-05").ToMonthString())
		assert.Equal(t, June, Parse("2020-06-05").ToMonthString())
		assert.Equal(t, July, Parse("2020-07-05").ToMonthString())
		assert.Equal(t, August, Parse("2020-08-05").ToMonthString())
		assert.Equal(t, September, Parse("2020-09-05").ToMonthString())
		assert.Equal(t, October, Parse("2020-10-05").ToMonthString())
		assert.Equal(t, November, Parse("2020-11-05").ToMonthString())
		assert.Equal(t, December, Parse("2020-12-05").ToMonthString(PRC))
	})
}

func TestCarbon_ToShortMonthString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "Jan", NewCarbon().ToShortMonthString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToShortMonthString())
		assert.Empty(t, Parse("0").ToShortMonthString())
		assert.Empty(t, Parse("xxx").ToShortMonthString())
	})

	t.Run("invalid resources", func(t *testing.T) {
		lang := NewLanguage()
		resources := map[string]string{
			"months": "",
		}
		lang.SetResources(resources)
		c := Parse("2020-01-05").SetLanguage(lang)
		assert.Empty(t, c.ToShortMonthString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "Jan", Parse("2020-01-05").ToShortMonthString())
		assert.Equal(t, "Feb", Parse("2020-02-05").ToShortMonthString())
		assert.Equal(t, "Mar", Parse("2020-03-05").ToShortMonthString())
		assert.Equal(t, "Apr", Parse("2020-04-05").ToShortMonthString())
		assert.Equal(t, "May", Parse("2020-05-05").ToShortMonthString())
		assert.Equal(t, "Jun", Parse("2020-06-05").ToShortMonthString())
		assert.Equal(t, "Jul", Parse("2020-07-05").ToShortMonthString())
		assert.Equal(t, "Aug", Parse("2020-08-05").ToShortMonthString())
		assert.Equal(t, "Sep", Parse("2020-09-05").ToShortMonthString())
		assert.Equal(t, "Oct", Parse("2020-10-05").ToShortMonthString())
		assert.Equal(t, "Nov", Parse("2020-11-05").ToShortMonthString())
		assert.Equal(t, "Dec", Parse("2020-12-05").ToShortMonthString(PRC))
	})
}

func TestCarbon_ToWeekString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, Monday, NewCarbon().ToWeekString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToWeekString())
		assert.Empty(t, Parse("0").ToWeekString())
		assert.Empty(t, Parse("xxx").ToWeekString())
	})

	t.Run("invalid resources", func(t *testing.T) {
		lang := NewLanguage()
		resources := map[string]string{
			"weeks": "xxx",
		}
		lang.SetResources(resources)
		c := Parse("2020-01-05").SetLanguage(lang)
		assert.Empty(t, c.ToWeekString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, Saturday, Parse("2020-08-01").ToWeekString())
		assert.Equal(t, Sunday, Parse("2020-08-02").ToWeekString())
		assert.Equal(t, Monday, Parse("2020-08-03").ToWeekString())
		assert.Equal(t, Tuesday, Parse("2020-08-04").ToWeekString())
		assert.Equal(t, Wednesday, Parse("2020-08-05").ToWeekString())
		assert.Equal(t, Thursday, Parse("2020-08-06").ToWeekString())
		assert.Equal(t, Friday, Parse("2020-08-07").ToWeekString(PRC))
	})
}

func TestCarbon_ToShortWeekString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "Mon", NewCarbon().ToShortWeekString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToShortWeekString())
		assert.Empty(t, Parse("0").ToShortWeekString())
		assert.Empty(t, Parse("xxx").ToShortWeekString())
	})

	t.Run("invalid resources", func(t *testing.T) {
		lang := NewLanguage()
		resources := map[string]string{
			"short_weeks": "xxx",
		}
		lang.SetResources(resources)
		c := Parse("2020-01-05").SetLanguage(lang)
		assert.Empty(t, c.ToShortWeekString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "Sat", Parse("2020-08-01").ToShortWeekString())
		assert.Equal(t, "Sun", Parse("2020-08-02").ToShortWeekString())
		assert.Equal(t, "Mon", Parse("2020-08-03").ToShortWeekString())
		assert.Equal(t, "Tue", Parse("2020-08-04").ToShortWeekString())
		assert.Equal(t, "Wed", Parse("2020-08-05").ToShortWeekString())
		assert.Equal(t, "Thu", Parse("2020-08-06").ToShortWeekString())
		assert.Equal(t, "Fri", Parse("2020-08-07").ToShortWeekString(PRC))
	})
}

func TestCarbon_ToDayDateTimeString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "Mon, Jan 1, 0001 12:00 AM", NewCarbon().ToDayDateTimeString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToDayDateTimeString())
		assert.Empty(t, Parse("0").ToDayDateTimeString())
		assert.Empty(t, Parse("xxx").ToDayDateTimeString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "Wed, Aug 5, 2020 1:14 PM", Parse("2020-08-05 13:14:15").ToDayDateTimeString())
		assert.Equal(t, "Wed, Aug 5, 2020 12:00 AM", Parse("2020-08-05", PRC).ToDayDateTimeString(PRC))
	})
}

func TestCarbon_ToDateTimeString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00", NewCarbon().ToDateTimeString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToDateTimeString())
		assert.Empty(t, Parse("0").ToDateTimeString())
		assert.Empty(t, Parse("xxx").ToDateTimeString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15", Parse("2020-08-05 13:14:15").ToDateTimeString())
		assert.Equal(t, "2020-08-05 13:14:15", Parse("2020-08-05T13:14:15.999999999+00:00").ToDateTimeString())
		assert.Equal(t, "2020-08-05 00:00:00", Parse("2020-08-05", PRC).ToDateTimeString(PRC))
	})
}

func TestCarbon_ToDateTimeMilliString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00", NewCarbon().ToDateTimeMilliString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToDateTimeMilliString())
		assert.Empty(t, Parse("0").ToDateTimeMilliString())
		assert.Empty(t, Parse("xxx").ToDateTimeMilliString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15", Parse("2020-08-05 13:14:15").ToDateTimeMilliString())
		assert.Equal(t, "2020-08-05 13:14:15.999", Parse("2020-08-05T13:14:15.999999999+00:00").ToDateTimeMilliString())
		assert.Equal(t, "2020-08-05 00:00:00", Parse("2020-08-05", PRC).ToDateTimeMilliString(PRC))
	})
}

func TestCarbon_ToDateTimeMicroString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00", NewCarbon().ToDateTimeMicroString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToDateTimeMicroString())
		assert.Empty(t, Parse("0").ToDateTimeMicroString())
		assert.Empty(t, Parse("xxx").ToDateTimeMicroString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15", Parse("2020-08-05 13:14:15").ToDateTimeMicroString())
		assert.Equal(t, "2020-08-05 13:14:15.999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToDateTimeMicroString())
		assert.Equal(t, "2020-08-05 00:00:00", Parse("2020-08-05", PRC).ToDateTimeMicroString(PRC))
	})
}

func TestCarbon_ToDateTimeNanoString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00", NewCarbon().ToDateTimeNanoString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToDateTimeNanoString())
		assert.Empty(t, Parse("0").ToDateTimeNanoString())
		assert.Empty(t, Parse("xxx").ToDateTimeNanoString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15", Parse("2020-08-05 13:14:15").ToDateTimeNanoString())
		assert.Equal(t, "2020-08-05 13:14:15.999999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToDateTimeNanoString())
		assert.Equal(t, "2020-08-05 00:00:00", Parse("2020-08-05", PRC).ToDateTimeNanoString(PRC))
	})
}

func TestCarbon_ToShortDateTimeString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "00010101000000", NewCarbon().ToShortDateTimeString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToShortDateTimeString())
		assert.Empty(t, Parse("0").ToShortDateTimeString())
		assert.Empty(t, Parse("xxx").ToShortDateTimeString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "20200805131415", Parse("2020-08-05 13:14:15").ToShortDateTimeString())
		assert.Equal(t, "20200805131415", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateTimeString())
		assert.Equal(t, "20200805000000", Parse("2020-08-05", PRC).ToShortDateTimeString(PRC))
	})
}

func TestCarbon_ToShortDateTimeMilliString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "00010101000000", NewCarbon().ToShortDateTimeMilliString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToShortDateTimeMilliString())
		assert.Empty(t, Parse("0").ToShortDateTimeMilliString())
		assert.Empty(t, Parse("xxx").ToShortDateTimeMilliString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "20200805131415", Parse("2020-08-05 13:14:15").ToShortDateTimeMilliString())
		assert.Equal(t, "20200805131415.999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateTimeMilliString())
		assert.Equal(t, "20200805000000", Parse("2020-08-05", PRC).ToShortDateTimeMilliString(PRC))
	})
}

func TestCarbon_ToShortDateTimeMicroString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "00010101000000", NewCarbon().ToShortDateTimeMicroString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToShortDateTimeMicroString())
		assert.Empty(t, Parse("0").ToShortDateTimeMicroString())
		assert.Empty(t, Parse("xxx").ToShortDateTimeMicroString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "20200805131415", Parse("2020-08-05 13:14:15").ToShortDateTimeMicroString())
		assert.Equal(t, "20200805131415.999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateTimeMicroString())
		assert.Equal(t, "20200805000000", Parse("2020-08-05", PRC).ToShortDateTimeMicroString(PRC))
	})
}

func TestCarbon_ToShortDateTimeNanoString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "00010101000000", NewCarbon().ToShortDateTimeNanoString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToShortDateTimeNanoString())
		assert.Empty(t, Parse("0").ToShortDateTimeNanoString())
		assert.Empty(t, Parse("xxx").ToShortDateTimeNanoString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "20200805131415", Parse("2020-08-05 13:14:15").ToShortDateTimeNanoString())
		assert.Equal(t, "20200805131415.999999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateTimeNanoString())
		assert.Equal(t, "20200805000000", Parse("2020-08-05", PRC).ToShortDateTimeNanoString(PRC))
	})
}

func TestCarbon_ToDateString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01", NewCarbon().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToDateString())
		assert.Empty(t, Parse("0").ToDateString())
		assert.Empty(t, Parse("xxx").ToDateString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05", Parse("2020-08-05 13:14:15").ToDateString())
		assert.Equal(t, "2020-08-05", Parse("2020-08-05T13:14:15.999999999+00:00").ToDateString())
		assert.Equal(t, "2020-08-05", Parse("2020-08-05", PRC).ToDateString(PRC))
	})
}

func TestCarbon_ToDateMilliString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01", NewCarbon().ToDateMilliString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToDateMilliString())
		assert.Empty(t, Parse("0").ToDateMilliString())
		assert.Empty(t, Parse("xxx").ToDateMilliString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05", Parse("2020-08-05 13:14:15").ToDateMilliString())
		assert.Equal(t, "2020-08-05.999", Parse("2020-08-05T13:14:15.999999999+00:00").ToDateMilliString())
		assert.Equal(t, "2020-08-05", Parse("2020-08-05", PRC).ToDateMilliString(PRC))
	})
}

func TestCarbon_ToDateMicroString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01", NewCarbon().ToDateMicroString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToDateMicroString())
		assert.Empty(t, Parse("0").ToDateMicroString())
		assert.Empty(t, Parse("xxx").ToDateMicroString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05", Parse("2020-08-05 13:14:15").ToDateMicroString())
		assert.Equal(t, "2020-08-05.999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToDateMicroString())
		assert.Equal(t, "2020-08-05", Parse("2020-08-05", PRC).ToDateMicroString(PRC))
	})
}

func TestCarbon_ToDateNanoString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01", NewCarbon().ToDateNanoString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToDateNanoString())
		assert.Empty(t, Parse("0").ToDateNanoString())
		assert.Empty(t, Parse("xxx").ToDateNanoString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05", Parse("2020-08-05 13:14:15").ToDateNanoString())
		assert.Equal(t, "2020-08-05.999999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToDateNanoString())
		assert.Equal(t, "2020-08-05", Parse("2020-08-05", PRC).ToDateNanoString(PRC))
	})
}

func TestCarbon_ToShortDateString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "00010101", NewCarbon().ToShortDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToShortDateString())
		assert.Empty(t, Parse("0").ToShortDateString())
		assert.Empty(t, Parse("xxx").ToShortDateString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "20200805", Parse("2020-08-05 13:14:15").ToShortDateString())
		assert.Equal(t, "20200805", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateString())
		assert.Equal(t, "20200805", Parse("2020-08-05", PRC).ToShortDateString(PRC))
	})
}

func TestCarbon_ToShortDateMilliString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "00010101", NewCarbon().ToShortDateMilliString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToShortDateMilliString())
		assert.Empty(t, Parse("0").ToShortDateMilliString())
		assert.Empty(t, Parse("xxx").ToShortDateMilliString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "20200805", Parse("2020-08-05 13:14:15").ToShortDateMilliString())
		assert.Equal(t, "20200805.999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateMilliString())
		assert.Equal(t, "20200805", Parse("2020-08-05", PRC).ToShortDateMilliString(PRC))
	})
}

func TestCarbon_ToShortDateMicroString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "00010101", NewCarbon().ToShortDateMicroString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToShortDateMicroString())
		assert.Empty(t, Parse("0").ToShortDateMicroString())
		assert.Empty(t, Parse("xxx").ToShortDateMicroString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "20200805", Parse("2020-08-05 13:14:15").ToShortDateMicroString())
		assert.Equal(t, "20200805.999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateMicroString())
		assert.Equal(t, "20200805", Parse("2020-08-05", PRC).ToShortDateMicroString(PRC))
	})
}

func TestCarbon_ToShortDateNanoString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "00010101", NewCarbon().ToShortDateNanoString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToShortDateNanoString())
		assert.Empty(t, Parse("0").ToShortDateNanoString())
		assert.Empty(t, Parse("xxx").ToShortDateNanoString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "20200805", Parse("2020-08-05 13:14:15").ToShortDateNanoString())
		assert.Equal(t, "20200805.999999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateNanoString())
		assert.Equal(t, "20200805", Parse("2020-08-05", PRC).ToShortDateNanoString(PRC))
	})
}

func TestCarbon_ToTimeString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "00:00:00", NewCarbon().ToTimeString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToTimeString())
		assert.Empty(t, Parse("0").ToTimeString())
		assert.Empty(t, Parse("xxx").ToTimeString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "13:14:15", Parse("2020-08-05 13:14:15").ToTimeString())
		assert.Equal(t, "13:14:15", Parse("2020-08-05T13:14:15.999999999+00:00").ToTimeString())
		assert.Equal(t, "00:00:00", Parse("2020-08-05", PRC).ToTimeString(PRC))
	})
}

func TestCarbon_ToTimeMilliString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "00:00:00", NewCarbon().ToTimeMilliString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToTimeMilliString())
		assert.Empty(t, Parse("0").ToTimeMilliString())
		assert.Empty(t, Parse("xxx").ToTimeMilliString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "13:14:15", Parse("2020-08-05 13:14:15").ToTimeMilliString())
		assert.Equal(t, "13:14:15.999", Parse("2020-08-05T13:14:15.999999999+00:00").ToTimeMilliString())
		assert.Equal(t, "00:00:00", Parse("2020-08-05", PRC).ToTimeMilliString(PRC))
	})
}

func TestCarbon_ToTimeMicroString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "00:00:00", NewCarbon().ToTimeMicroString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToTimeMicroString())
		assert.Empty(t, Parse("0").ToTimeMicroString())
		assert.Empty(t, Parse("xxx").ToTimeMicroString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "13:14:15", Parse("2020-08-05 13:14:15").ToTimeMicroString())
		assert.Equal(t, "13:14:15.999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToTimeMicroString())
		assert.Equal(t, "00:00:00", Parse("2020-08-05", PRC).ToTimeMicroString(PRC))
	})
}

func TestCarbon_ToTimeNanoString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "00:00:00", NewCarbon().ToTimeNanoString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToTimeNanoString())
		assert.Empty(t, Parse("0").ToTimeNanoString())
		assert.Empty(t, Parse("xxx").ToTimeNanoString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "13:14:15", Parse("2020-08-05 13:14:15").ToTimeNanoString())
		assert.Equal(t, "13:14:15.999999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToTimeNanoString())
		assert.Equal(t, "00:00:00", Parse("2020-08-05", PRC).ToTimeNanoString(PRC))
	})
}

func TestCarbon_ToShortTimeString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "000000", NewCarbon().ToShortTimeString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToShortTimeString())
		assert.Empty(t, Parse("0").ToShortTimeString())
		assert.Empty(t, Parse("xxx").ToShortTimeString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "131415", Parse("2020-08-05 13:14:15").ToShortTimeString())
		assert.Equal(t, "131415", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortTimeString())
		assert.Equal(t, "000000", Parse("2020-08-05", PRC).ToShortTimeString(PRC))
	})
}

func TestCarbon_ToShortTimeMilliString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "000000", NewCarbon().ToShortTimeMilliString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToShortTimeMilliString())
		assert.Empty(t, Parse("0").ToShortTimeMilliString())
		assert.Empty(t, Parse("xxx").ToShortTimeMilliString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "131415", Parse("2020-08-05 13:14:15").ToShortTimeMilliString())
		assert.Equal(t, "131415.999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortTimeMilliString())
		assert.Equal(t, "000000", Parse("2020-08-05", PRC).ToShortTimeMilliString(PRC))
	})
}

func TestCarbon_ToShortTimeMicroString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "000000", NewCarbon().ToShortTimeMicroString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToShortTimeMicroString())
		assert.Empty(t, Parse("0").ToShortTimeMicroString())
		assert.Empty(t, Parse("xxx").ToShortTimeMicroString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "131415", Parse("2020-08-05 13:14:15").ToShortTimeMicroString())
		assert.Equal(t, "131415.999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortTimeMicroString())
		assert.Equal(t, "000000", Parse("2020-08-05", PRC).ToShortTimeMicroString(PRC))
	})
}

func TestCarbon_ToShortTimeNanoString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "000000", NewCarbon().ToShortTimeNanoString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToShortTimeNanoString())
		assert.Empty(t, Parse("0").ToShortTimeNanoString())
		assert.Empty(t, Parse("xxx").ToShortTimeNanoString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "131415", Parse("2020-08-05 13:14:15").ToShortTimeNanoString())
		assert.Equal(t, "131415.999999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortTimeNanoString())
		assert.Equal(t, "000000", Parse("2020-08-05", PRC).ToShortTimeNanoString(PRC))
	})
}

func TestCarbon_ToAtomString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01T00:00:00Z", NewCarbon().ToAtomString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToAtomString())
		assert.Empty(t, Parse("0").ToAtomString())
		assert.Empty(t, Parse("xxx").ToAtomString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToAtomString())
		assert.Equal(t, "2020-08-05T13:14:15Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToAtomString())
		assert.Equal(t, "2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToAtomString(PRC))
	})
}

func TestCarbon_ToAnsicString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "Mon Jan  1 00:00:00 0001", NewCarbon().ToAnsicString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToAnsicString())
		assert.Empty(t, Parse("0").ToAnsicString())
		assert.Empty(t, Parse("xxx").ToAnsicString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "Wed Aug  5 13:14:15 2020", Parse("2020-08-05 13:14:15").ToAnsicString())
		assert.Equal(t, "Wed Aug  5 13:14:15 2020", Parse("2020-08-05T13:14:15.999999999+00:00").ToAnsicString())
		assert.Equal(t, "Wed Aug  5 00:00:00 2020", Parse("2020-08-05", PRC).ToAnsicString(PRC))
	})
}

func TestCarbon_ToCookieString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "Monday, 01-Jan-0001 00:00:00 UTC", NewCarbon().ToCookieString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToCookieString())
		assert.Empty(t, Parse("0").ToCookieString())
		assert.Empty(t, Parse("xxx").ToCookieString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "Wednesday, 05-Aug-2020 13:14:15 UTC", Parse("2020-08-05 13:14:15").ToCookieString())
		assert.Equal(t, "Wednesday, 05-Aug-2020 13:14:15 UTC", Parse("2020-08-05T13:14:15.999999999+00:00").ToCookieString())
		assert.Equal(t, "Wednesday, 05-Aug-2020 00:00:00 CST", Parse("2020-08-05", PRC).ToCookieString(PRC))
	})
}

func TestCarbon_ToRssString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "Mon, 01 Jan 0001 00:00:00 +0000", NewCarbon().ToRssString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToRssString())
		assert.Empty(t, Parse("0").ToRssString())
		assert.Empty(t, Parse("xxx").ToRssString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 +0000", Parse("2020-08-05 13:14:15").ToRssString())
		assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 +0000", Parse("2020-08-05T13:14:15.999999999+00:00").ToRssString())
		assert.Equal(t, "Wed, 05 Aug 2020 00:00:00 +0800", Parse("2020-08-05", PRC).ToRssString(PRC))
	})
}

func TestCarbon_ToW3cString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01T00:00:00Z", NewCarbon().ToW3cString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToW3cString())
		assert.Empty(t, Parse("0").ToW3cString())
		assert.Empty(t, Parse("xxx").ToW3cString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToW3cString())
		assert.Equal(t, "2020-08-05T13:14:15Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToW3cString())
		assert.Equal(t, "2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToW3cString(PRC))
	})
}

func TestCarbon_ToUnixDateString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "Mon Jan  1 00:00:00 UTC 0001", NewCarbon().ToUnixDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToUnixDateString())
		assert.Empty(t, Parse("0").ToUnixDateString())
		assert.Empty(t, Parse("xxx").ToUnixDateString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "Wed Aug  5 13:14:15 UTC 2020", Parse("2020-08-05 13:14:15").ToUnixDateString())
		assert.Equal(t, "Wed Aug  5 13:14:15 UTC 2020", Parse("2020-08-05T13:14:15.999999999+00:00").ToUnixDateString())
		assert.Equal(t, "Wed Aug  5 00:00:00 CST 2020", Parse("2020-08-05", PRC).ToUnixDateString(PRC))
	})
}

func TestCarbon_ToRubyDateString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "Mon Jan 01 00:00:00 +0000 0001", NewCarbon().ToRubyDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToRubyDateString())
		assert.Empty(t, Parse("0").ToRubyDateString())
		assert.Empty(t, Parse("xxx").ToRubyDateString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "Wed Aug 05 13:14:15 +0000 2020", Parse("2020-08-05 13:14:15").ToRubyDateString())
		assert.Equal(t, "Wed Aug 05 13:14:15 +0000 2020", Parse("2020-08-05T13:14:15.999999999+00:00").ToRubyDateString())
		assert.Equal(t, "Wed Aug 05 00:00:00 +0800 2020", Parse("2020-08-05", PRC).ToRubyDateString(PRC))
	})
}

func TestCarbon_ToKitchenString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "12:00AM", NewCarbon().ToKitchenString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToKitchenString())
		assert.Empty(t, Parse("0").ToKitchenString())
		assert.Empty(t, Parse("xxx").ToKitchenString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "1:14PM", Parse("2020-08-05 13:14:15").ToKitchenString())
		assert.Equal(t, "1:14PM", Parse("2020-08-05T13:14:15.999999999+00:00").ToKitchenString())
		assert.Equal(t, "12:00AM", Parse("2020-08-05", PRC).ToKitchenString(PRC))
	})
}

func TestCarbon_ToIso8601String(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01T00:00:00+00:00", NewCarbon().ToIso8601String())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToIso8601String())
		assert.Empty(t, Parse("0").ToIso8601String())
		assert.Empty(t, Parse("xxx").ToIso8601String())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05T13:14:15+00:00", Parse("2020-08-05 13:14:15").ToIso8601String())
		assert.Equal(t, "2020-08-05T13:14:15+00:00", Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601String())
		assert.Equal(t, "2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToIso8601String(PRC))
	})
}

func TestCarbon_ToIso8601MilliString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01T00:00:00+00:00", NewCarbon().ToIso8601MilliString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToIso8601MilliString())
		assert.Empty(t, Parse("0").ToIso8601MilliString())
		assert.Empty(t, Parse("xxx").ToIso8601MilliString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05T13:14:15+00:00", Parse("2020-08-05 13:14:15").ToIso8601MilliString())
		assert.Equal(t, "2020-08-05T13:14:15.999+00:00", Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601MilliString())
		assert.Equal(t, "2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToIso8601MilliString(PRC))
	})
}

func TestCarbon_TToIso8601MicroString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01T00:00:00+00:00", NewCarbon().ToIso8601MicroString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToIso8601MicroString())
		assert.Empty(t, Parse("0").ToIso8601MicroString())
		assert.Empty(t, Parse("xxx").ToIso8601MicroString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05T13:14:15+00:00", Parse("2020-08-05 13:14:15").ToIso8601MicroString())
		assert.Equal(t, "2020-08-05T13:14:15.999999+00:00", Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601MicroString())
		assert.Equal(t, "2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToIso8601MicroString(PRC))
	})
}

func TestCarbon_ToIso8601NanoString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01T00:00:00+00:00", NewCarbon().ToIso8601NanoString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToIso8601NanoString())
		assert.Empty(t, Parse("0").ToIso8601NanoString())
		assert.Empty(t, Parse("xxx").ToIso8601NanoString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05T13:14:15+00:00", Parse("2020-08-05 13:14:15").ToIso8601NanoString())
		assert.Equal(t, "2020-08-05T13:14:15.999999999+00:00", Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601NanoString())
		assert.Equal(t, "2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToIso8601NanoString(PRC))
	})
}

func TestCarbon_ToIso8601ZuluString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01T00:00:00Z", NewCarbon().ToIso8601ZuluString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToIso8601ZuluString())
		assert.Empty(t, Parse("0").ToIso8601ZuluString())
		assert.Empty(t, Parse("xxx").ToIso8601ZuluString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToIso8601ZuluString())
		assert.Equal(t, "2020-08-05T13:14:15Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601ZuluString())
		assert.Equal(t, "2020-08-05T00:00:00Z", Parse("2020-08-05", PRC).ToIso8601ZuluString(PRC))
	})
}

func TestCarbon_ToIso8601ZuluMilliString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01T00:00:00Z", NewCarbon().ToIso8601ZuluMilliString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToIso8601ZuluMilliString())
		assert.Empty(t, Parse("0").ToIso8601ZuluMilliString())
		assert.Empty(t, Parse("xxx").ToIso8601ZuluMilliString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToIso8601ZuluMilliString())
		assert.Equal(t, "2020-08-05T13:14:15.999Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601ZuluMilliString())
		assert.Equal(t, "2020-08-05T00:00:00Z", Parse("2020-08-05", PRC).ToIso8601ZuluMilliString(PRC))
	})
}

func TestCarbon_ToIso8601ZuluMicroString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01T00:00:00Z", NewCarbon().ToIso8601ZuluMicroString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToIso8601ZuluMicroString())
		assert.Empty(t, Parse("0").ToIso8601ZuluMicroString())
		assert.Empty(t, Parse("xxx").ToIso8601ZuluMicroString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToIso8601ZuluMicroString())
		assert.Equal(t, "2020-08-05T13:14:15.999999Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601ZuluMicroString())
		assert.Equal(t, "2020-08-05T00:00:00Z", Parse("2020-08-05", PRC).ToIso8601ZuluMicroString(PRC))
	})
}

func TestCarbon_ToIso8601ZuluNanoString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01T00:00:00Z", NewCarbon().ToIso8601ZuluNanoString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToIso8601ZuluNanoString())
		assert.Empty(t, Parse("0").ToIso8601ZuluNanoString())
		assert.Empty(t, Parse("xxx").ToIso8601ZuluNanoString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToIso8601ZuluNanoString())
		assert.Equal(t, "2020-08-05T13:14:15.999999999Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601ZuluNanoString())
		assert.Equal(t, "2020-08-05T00:00:00Z", Parse("2020-08-05", PRC).ToIso8601ZuluNanoString(PRC))
	})
}

func TestCarbon_ToRfc822String(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "01 Jan 01 00:00 UTC", NewCarbon().ToRfc822String())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToRfc822String())
		assert.Empty(t, Parse("0").ToRfc822String())
		assert.Empty(t, Parse("xxx").ToRfc822String())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "05 Aug 20 13:14 UTC", Parse("2020-08-05 13:14:15").ToRfc822String())
		assert.Equal(t, "05 Aug 20 13:14 UTC", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc822String())
		assert.Equal(t, "05 Aug 20 00:00 CST", Parse("2020-08-05", PRC).ToRfc822String(PRC))
	})
}

func TestCarbon_ToRfc822zString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "01 Jan 01 00:00 +0000", NewCarbon().ToRfc822zString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToRfc822zString())
		assert.Empty(t, Parse("0").ToRfc822zString())
		assert.Empty(t, Parse("xxx").ToRfc822zString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "05 Aug 20 13:14 +0000", Parse("2020-08-05 13:14:15").ToRfc822zString())
		assert.Equal(t, "05 Aug 20 13:14 +0000", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc822zString())
		assert.Equal(t, "05 Aug 20 00:00 +0800", Parse("2020-08-05", PRC).ToRfc822zString(PRC))
	})
}

func TestCarbon_ToRfc850String(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "Monday, 01-Jan-01 00:00:00 UTC", NewCarbon().ToRfc850String())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToRfc850String())
		assert.Empty(t, Parse("0").ToRfc850String())
		assert.Empty(t, Parse("xxx").ToRfc850String())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "Wednesday, 05-Aug-20 13:14:15 UTC", Parse("2020-08-05 13:14:15").ToRfc850String())
		assert.Equal(t, "Wednesday, 05-Aug-20 13:14:15 UTC", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc850String())
		assert.Equal(t, "Wednesday, 05-Aug-20 00:00:00 CST", Parse("2020-08-05", PRC).ToRfc850String(PRC))
	})
}

func TestCarbon_ToRfc1036String(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "Mon, 01 Jan 01 00:00:00 +0000", NewCarbon().ToRfc1036String())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToRfc1036String())
		assert.Empty(t, Parse("0").ToRfc1036String())
		assert.Empty(t, Parse("xxx").ToRfc1036String())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "Wed, 05 Aug 20 13:14:15 +0000", Parse("2020-08-05 13:14:15").ToRfc1036String())
		assert.Equal(t, "Wed, 05 Aug 20 13:14:15 +0000", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc1036String())
		assert.Equal(t, "Wed, 05 Aug 20 00:00:00 +0800", Parse("2020-08-05", PRC).ToRfc1036String(PRC))
	})
}

func TestCarbon_ToRfc1123String(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "Mon, 01 Jan 0001 00:00:00 UTC", NewCarbon().ToRfc1123String())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToRfc1123String())
		assert.Empty(t, Parse("0").ToRfc1123String())
		assert.Empty(t, Parse("xxx").ToRfc1123String())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 UTC", Parse("2020-08-05 13:14:15").ToRfc1123String())
		assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 UTC", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc1123String())
		assert.Equal(t, "Wed, 05 Aug 2020 00:00:00 CST", Parse("2020-08-05", PRC).ToRfc1123String(PRC))
	})
}

func TestCarbon_ToRfc1123zString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "Mon, 01 Jan 0001 00:00:00 +0000", NewCarbon().ToRfc1123zString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToRfc1123zString())
		assert.Empty(t, Parse("0").ToRfc1123zString())
		assert.Empty(t, Parse("xxx").ToRfc1123zString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 +0000", Parse("2020-08-05 13:14:15").ToRfc1123zString())
		assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 +0000", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc1123zString())
		assert.Equal(t, "Wed, 05 Aug 2020 00:00:00 +0800", Parse("2020-08-05", PRC).ToRfc1123zString(PRC))
	})
}

func TestCarbon_ToRfc2822String(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "Mon, 01 Jan 0001 00:00:00 +0000", NewCarbon().ToRfc2822String())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToRfc2822String())
		assert.Empty(t, Parse("0").ToRfc2822String())
		assert.Empty(t, Parse("xxx").ToRfc2822String())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 +0000", Parse("2020-08-05 13:14:15").ToRfc2822String())
		assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 +0000", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc2822String())
		assert.Equal(t, "Wed, 05 Aug 2020 00:00:00 +0800", Parse("2020-08-05", PRC).ToRfc2822String(PRC))
	})
}

func TestCarbon_ToRfc3339String(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01T00:00:00Z", NewCarbon().ToRfc3339String())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToRfc3339String())
		assert.Empty(t, Parse("0").ToRfc3339String())
		assert.Empty(t, Parse("xxx").ToRfc3339String())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToRfc3339String())
		assert.Equal(t, "2020-08-05T13:14:15Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc3339String())
		assert.Equal(t, "2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToRfc3339String(PRC))
	})
}

func TestCarbon_ToRfc3339MilliString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01T00:00:00Z", NewCarbon().ToRfc3339MilliString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToRfc3339MilliString())
		assert.Empty(t, Parse("0").ToRfc3339MilliString())
		assert.Empty(t, Parse("xxx").ToRfc3339MilliString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToRfc3339MilliString())
		assert.Equal(t, "2020-08-05T13:14:15.999Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc3339MilliString())
		assert.Equal(t, "2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToRfc3339MilliString(PRC))
	})
}

func TestCarbon_ToRfc3339MicroString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01T00:00:00Z", NewCarbon().ToRfc3339MicroString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToRfc3339MicroString())
		assert.Empty(t, Parse("0").ToRfc3339MicroString())
		assert.Empty(t, Parse("xxx").ToRfc3339MicroString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToRfc3339MicroString())
		assert.Equal(t, "2020-08-05T13:14:15.999999Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc3339MicroString())
		assert.Equal(t, "2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToRfc3339MicroString(PRC))
	})
}

func TestCarbon_ToRfc3339NanoString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01T00:00:00Z", NewCarbon().ToRfc3339NanoString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToRfc3339NanoString())
		assert.Empty(t, Parse("0").ToRfc3339NanoString())
		assert.Empty(t, Parse("xxx").ToRfc3339NanoString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToRfc3339NanoString())
		assert.Equal(t, "2020-08-05T13:14:15.999999999Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc3339NanoString())
		assert.Equal(t, "2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToRfc3339NanoString(PRC))
	})
}

func TestCarbon_ToRfc7231String(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "Mon, 01 Jan 0001 00:00:00 UTC", NewCarbon().ToRfc7231String())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToRfc7231String())
		assert.Empty(t, Parse("0").ToRfc7231String())
		assert.Empty(t, Parse("xxx").ToRfc7231String())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 UTC", Parse("2020-08-05 13:14:15").ToRfc7231String())
		assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 UTC", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc7231String())
		assert.Equal(t, "Wed, 05 Aug 2020 00:00:00 CST", Parse("2020-08-05", PRC).ToRfc7231String(PRC))
	})
}

func TestCarbon_ToFormattedDateString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "Jan 1, 0001", NewCarbon().ToFormattedDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToFormattedDateString())
		assert.Empty(t, Parse("0").ToFormattedDateString())
		assert.Empty(t, Parse("xxx").ToFormattedDateString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "Aug 5, 2020", Parse("2020-08-05 13:14:15").ToFormattedDateString())
		assert.Equal(t, "Aug 5, 2020", Parse("2020-08-05T13:14:15.999999999+00:00").ToFormattedDateString())
		assert.Equal(t, "Aug 5, 2020", Parse("2020-08-05", PRC).ToFormattedDateString(PRC))
	})
}

func TestCarbon_ToFormattedDayDateString(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "Mon, Jan 1, 0001", NewCarbon().ToFormattedDayDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").ToFormattedDayDateString())
		assert.Empty(t, Parse("0").ToFormattedDayDateString())
		assert.Empty(t, Parse("xxx").ToFormattedDayDateString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "Wed, Aug 5, 2020", Parse("2020-08-05 13:14:15").ToFormattedDayDateString())
		assert.Equal(t, "Wed, Aug 5, 2020", Parse("2020-08-05T13:14:15.999999999+00:00").ToFormattedDayDateString())
		assert.Equal(t, "Wed, Aug 5, 2020", Parse("2020-08-05", PRC).ToFormattedDayDateString(PRC))
	})
}

func TestCarbon_Layout(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00", NewCarbon().Layout(DateTimeLayout))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").Layout(DateTimeLayout))
		assert.Empty(t, Parse("0").Layout(DateTimeLayout))
		assert.Empty(t, Parse("xxx").Layout(DateTimeLayout))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15", Parse("2020-08-05 13:14:15").Layout(DateTimeLayout))
		assert.Equal(t, "2020-08-05 13:14:15", Parse("2020-08-05T13:14:15.999999999+00:00").Layout(DateTimeLayout))
		assert.Equal(t, "2020-08-05 00:00:00", Parse("2020-08-05", PRC).Layout(DateTimeLayout, PRC))
		assert.Equal(t, "2020年08月05日", Parse("2020-08-05 13:14:15").Layout("2006年01月02日"))
		assert.Equal(t, "Wed, 05 Aug 2020 13:14:15 GMT", Parse("2020-08-05 13:14:15").Layout("Mon, 02 Jan 2006 15:04:05 GMT"))
	})
}

func TestCarbon_Format(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00", NewCarbon().Format(DateTimeFormat))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").Format(DateTimeFormat))
		assert.Empty(t, Parse("0").Format(DateTimeFormat))
		assert.Empty(t, Parse("xxx").Format(DateTimeFormat))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15", Parse("2020-08-05 13:14:15").Format(DateTimeFormat))
		assert.Equal(t, "2020-08-05", Parse("2020-08-05T13:14:15.999999999+00:00").Format(DateFormat))
		assert.Equal(t, "2020-08-05 00:00:00", Parse("2020-08-05", PRC).Format(DateTimeFormat, PRC))
		assert.Equal(t, "2020年08月05日", Parse("2020-08-05 13:14:15").Format("Y年m月d日"))
		assert.Equal(t, "Wed", Parse("2020-08-05 13:14:15").Format("D"))
		assert.Equal(t, "Wednesday", Parse("2020-08-05 13:14:15").Format("l"))
		assert.Equal(t, "August", Parse("2020-08-05 13:14:15").Format("F"))
		assert.Equal(t, "Aug", Parse("2020-08-05 13:14:15").Format("M"))
		assert.Equal(t, "5", Parse("2020-08-05 13:14:15").Format("j"))
		assert.Equal(t, "32", Parse("2020-08-05 13:14:15").Format("W"))
		assert.Equal(t, "August", Parse("2020-08-05 13:14:15").Format("F"))
		assert.Equal(t, "03", Parse("2020-08-05 13:14:15").Format("N"))
		assert.Equal(t, "1", Parse("2020-08-05 13:14:15").Format("L"))
		assert.Equal(t, "0", Parse("2021-08-05 13:14:15").Format("L"))
		assert.Equal(t, "13", Parse("2020-08-05 13:14:15").Format("G"))
		assert.Equal(t, "1596633255", Parse("2020-08-05 13:14:15").Format("S"))
		assert.Equal(t, "1596633255000", Parse("2020-08-05 13:14:15").Format("U"))
		assert.Equal(t, "1596633255000000", Parse("2020-08-05 13:14:15").Format("V"))
		assert.Equal(t, "1596633255000000000", Parse("2020-08-05 13:14:15").Format("X"))
		assert.Equal(t, "999", Parse("2020-08-05 13:14:15.999999999").Format("u"))
		assert.Equal(t, "999999", Parse("2020-08-05 13:14:15.999999999").Format("v"))
		assert.Equal(t, "999999999", Parse("2020-08-05 13:14:15.999999999").Format("x"))
		assert.Equal(t, "2", Parse("2020-08-05 13:14:15.999999999").Format("w"))
		assert.Equal(t, "31", Parse("2020-08-05 13:14:15.999999999").Format("t"))
		assert.Equal(t, "PRC", Parse("2020-08-05 13:14:15.999999999", PRC).Format("z"))
		assert.Equal(t, "28800", Parse("2020-08-05 13:14:15.999999999", PRC).Format("o"))
		assert.Equal(t, "3", Parse("2020-08-05 13:14:15.999999999").Format("q"))
		assert.Equal(t, "21", Parse("2020-08-05 13:14:15.999999999").Format("c"))
		assert.Equal(t, "Z", Parse("2020-08-05 13:14:15.999999999", UTC).Format("R"))
		assert.Equal(t, "+08:00", Parse("2020-08-05 13:14:15.999999999", PRC).Format("R"))
		assert.Equal(t, "Z", Parse("2020-08-05 13:14:15.999999999", UTC).Format("Q"))
		assert.Equal(t, "+0800", Parse("2020-08-05 13:14:15.999999999", PRC).Format("Q"))
		assert.Equal(t, "CST", Parse("2020-08-05 13:14:15.999999999", PRC).Format("Z"))
		assert.Equal(t, "5th", Parse("2020-08-05 13:14:15").Format("jK"))
		assert.Equal(t, "22nd", Parse("2020-08-22 13:14:15").Format("jK"))
		assert.Equal(t, "23rd", Parse("2020-08-23 13:14:15").Format("jK"))
		assert.Equal(t, "31st", Parse("2020-08-31 13:14:15").Format("jK"))
		assert.Equal(t, "It is 2020-08-31 13:14:15", Parse("2020-08-31 13:14:15").Format("I\\t \\i\\s Y-m-d H:i:s"))
		assert.Equal(t, "上次打卡时间:2020-08-05 13:14:15，请每日按时打卡", Parse("2020-08-05 13:14:15").Format("上次打卡时间:Y-m-d H:i:s，请每日按时打卡"))
	})
}
