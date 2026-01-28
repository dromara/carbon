package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type OutputerSuite struct {
	suite.Suite
}

func TestOutputerSuite(t *testing.T) {
	suite.Run(t, new(OutputerSuite))
}

func (s *OutputerSuite) TestCarbon_GoString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.GoString())
	})

	s.Run("zero carbon", func() {
		s.Equal("time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)", NewCarbon().GoString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").GoString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").GoString())
	})

	s.Run("valid carbon", func() {
		s.Equal("time.Date(2020, time.August, 5, 13, 14, 15, 0, time.UTC)", Parse("2020-08-05 13:14:15").GoString())
		s.Equal("time.Date(2020, time.August, 5, 13, 14, 15, 0, time.Location(\"PRC\"))", Parse("2020-08-05 13:14:15", PRC).GoString())
	})
}

func (s *OutputerSuite) TestCarbon_ToString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:14:15 +0000 UTC", Parse("2020-08-05 13:14:15").ToString())
		s.Equal("2020-08-05 13:14:15 +0800 CST", Parse("2020-08-05 13:14:15", PRC).ToString())
		s.Equal("2020-08-05 21:14:15 +0800 CST", Parse("2020-08-05 13:14:15").ToString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToMonthString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToMonthString())
	})

	s.Run("zero carbon", func() {
		s.Equal("January", NewCarbon().ToMonthString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToMonthString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToMonthString())
	})

	s.Run("nil lang", func() {
		c := Now()
		c.lang = nil
		s.Empty(c.ToMonthString())
	})

	s.Run("valid carbon", func() {
		s.Equal(January.String(), Parse("2020-01-05").ToMonthString())
		s.Equal(February.String(), Parse("2020-02-05").ToMonthString())
		s.Equal(March.String(), Parse("2020-03-05").ToMonthString())
		s.Equal(April.String(), Parse("2020-04-05").ToMonthString())
		s.Equal(May.String(), Parse("2020-05-05").ToMonthString())
		s.Equal(June.String(), Parse("2020-06-05").ToMonthString())
		s.Equal(July.String(), Parse("2020-07-05").ToMonthString())
		s.Equal(August.String(), Parse("2020-08-05").ToMonthString())
		s.Equal(September.String(), Parse("2020-09-05").ToMonthString())
		s.Equal(October.String(), Parse("2020-10-05").ToMonthString())
		s.Equal(November.String(), Parse("2020-11-05").ToMonthString())
		s.Equal(December.String(), Parse("2020-12-05").ToMonthString(PRC))
	})

	s.Run("empty resources", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{})
		c := Parse("2020-01-05").SetLanguage(lang)
		s.Empty(c.ToMonthString())
	})

	s.Run("nil lang", func() {
		c := Now()
		c.lang = nil
		s.Empty(c.ToMonthString())
	})

	s.Run("error resources", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{
			"months": "xxx",
		})
		c := Parse("2020-01-05").SetLanguage(lang)
		s.Empty(c.ToMonthString())
	})
}

func (s *OutputerSuite) TestCarbon_ToShortMonthString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToShortMonthString())
	})

	s.Run("zero carbon", func() {
		s.Equal("Jan", NewCarbon().ToShortMonthString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToShortMonthString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToShortMonthString())
	})

	s.Run("nil lang", func() {
		c := Now()
		c.lang = nil
		s.Empty(c.ToShortMonthString())
	})

	s.Run("valid carbon", func() {
		s.Equal("Jan", Parse("2020-01-05").ToShortMonthString())
		s.Equal("Feb", Parse("2020-02-05").ToShortMonthString())
		s.Equal("Mar", Parse("2020-03-05").ToShortMonthString())
		s.Equal("Apr", Parse("2020-04-05").ToShortMonthString())
		s.Equal("May", Parse("2020-05-05").ToShortMonthString())
		s.Equal("Jun", Parse("2020-06-05").ToShortMonthString())
		s.Equal("Jul", Parse("2020-07-05").ToShortMonthString())
		s.Equal("Aug", Parse("2020-08-05").ToShortMonthString())
		s.Equal("Sep", Parse("2020-09-05").ToShortMonthString())
		s.Equal("Oct", Parse("2020-10-05").ToShortMonthString())
		s.Equal("Nov", Parse("2020-11-05").ToShortMonthString())
		s.Equal("Dec", Parse("2020-12-05").ToShortMonthString(PRC))
	})

	s.Run("empty resources", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{})
		c := Parse("2020-01-05").SetLanguage(lang)
		s.Empty(c.ToShortMonthString())
	})

	s.Run("error resources", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{
			"months": "xxx",
		})
		c := Parse("2020-01-05").SetLanguage(lang)
		s.Empty(c.ToShortMonthString())
	})
}

func (s *OutputerSuite) TestCarbon_ToWeekString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToWeekString())
	})

	s.Run("zero carbon", func() {
		s.Equal("Monday", NewCarbon().ToWeekString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToWeekString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToWeekString())
	})

	s.Run("nil lang", func() {
		c := Now()
		c.lang = nil
		s.Empty(c.ToWeekString())
	})

	s.Run("valid carbon", func() {
		s.Equal(Saturday.String(), Parse("2020-08-01").ToWeekString())
		s.Equal(Sunday.String(), Parse("2020-08-02").ToWeekString())
		s.Equal(Monday.String(), Parse("2020-08-03").ToWeekString())
		s.Equal(Tuesday.String(), Parse("2020-08-04").ToWeekString())
		s.Equal(Wednesday.String(), Parse("2020-08-05").ToWeekString())
		s.Equal(Thursday.String(), Parse("2020-08-06").ToWeekString())
		s.Equal(Friday.String(), Parse("2020-08-07").ToWeekString(PRC))
	})

	s.Run("empty resources", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{})
		c := Parse("2020-01-05").SetLanguage(lang)
		s.Empty(c.ToWeekString())
	})

	s.Run("error resources", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{
			"weeks": "xxx",
		})
		c := Parse("2020-01-05").SetLanguage(lang)
		s.Empty(c.ToWeekString())
	})
}

func (s *OutputerSuite) TestCarbon_ToShortWeekString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToShortWeekString())
	})

	s.Run("zero carbon", func() {
		s.Equal("Mon", NewCarbon().ToShortWeekString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToShortWeekString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToShortWeekString())
	})

	s.Run("nil lang", func() {
		c := Now()
		c.lang = nil
		s.Empty(c.ToShortWeekString())
	})

	s.Run("valid carbon", func() {
		s.Equal("Sat", Parse("2020-08-01").ToShortWeekString())
		s.Equal("Sun", Parse("2020-08-02").ToShortWeekString())
		s.Equal("Mon", Parse("2020-08-03").ToShortWeekString())
		s.Equal("Tue", Parse("2020-08-04").ToShortWeekString())
		s.Equal("Wed", Parse("2020-08-05").ToShortWeekString())
		s.Equal("Thu", Parse("2020-08-06").ToShortWeekString())
		s.Equal("Fri", Parse("2020-08-07").ToShortWeekString(PRC))
	})

	s.Run("empty resources", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{})
		c := Parse("2020-01-05").SetLanguage(lang)
		s.Empty(c.ToShortWeekString())
	})

	s.Run("error resources", func() {
		lang := NewLanguage()
		lang.SetResources(map[string]string{
			"short_weeks": "xxx",
		})
		c := Parse("2020-01-05").SetLanguage(lang)
		s.Empty(c.ToShortWeekString())
	})
}

func (s *OutputerSuite) TestCarbon_ToDayDateTimeString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToDayDateTimeString())
	})

	s.Run("zero carbon", func() {
		s.Equal("Mon, Jan 1, 0001 12:00 AM", NewCarbon().ToDayDateTimeString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToDayDateTimeString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToDayDateTimeString())
	})

	s.Run("valid carbon", func() {
		s.Equal("Wed, Aug 5, 2020 1:14 PM", Parse("2020-08-05 13:14:15").ToDayDateTimeString())
		s.Equal("Wed, Aug 5, 2020 12:00 AM", Parse("2020-08-05", PRC).ToDayDateTimeString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToDateTimeString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToDateTimeString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00", NewCarbon().ToDateTimeString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToDateTimeString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToDateTimeString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:14:15", Parse("2020-08-05 13:14:15").ToDateTimeString())
		s.Equal("2020-08-05 13:14:15", Parse("2020-08-05T13:14:15.999999999+00:00").ToDateTimeString())
		s.Equal("2020-08-05 00:00:00", Parse("2020-08-05", PRC).ToDateTimeString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToDateTimeMilliString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToDateTimeMilliString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00", NewCarbon().ToDateTimeMilliString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToDateTimeMilliString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToDateTimeMilliString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:14:15", Parse("2020-08-05 13:14:15").ToDateTimeMilliString())
		s.Equal("2020-08-05 13:14:15.999", Parse("2020-08-05T13:14:15.999999999+00:00").ToDateTimeMilliString())
		s.Equal("2020-08-05 00:00:00", Parse("2020-08-05", PRC).ToDateTimeMilliString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToDateTimeMicroString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToDateTimeMicroString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00", NewCarbon().ToDateTimeMicroString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToDateTimeMicroString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToDateTimeMicroString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:14:15", Parse("2020-08-05 13:14:15").ToDateTimeMicroString())
		s.Equal("2020-08-05 13:14:15.999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToDateTimeMicroString())
		s.Equal("2020-08-05 00:00:00", Parse("2020-08-05", PRC).ToDateTimeMicroString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToDateTimeNanoString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToDateTimeNanoString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00", NewCarbon().ToDateTimeNanoString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToDateTimeNanoString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToDateTimeNanoString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:14:15", Parse("2020-08-05 13:14:15").ToDateTimeNanoString())
		s.Equal("2020-08-05 13:14:15.999999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToDateTimeNanoString())
		s.Equal("2020-08-05 00:00:00", Parse("2020-08-05", PRC).ToDateTimeNanoString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToShortDateTimeString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToShortDateTimeString())
	})

	s.Run("zero carbon", func() {
		s.Equal("00010101000000", NewCarbon().ToShortDateTimeString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToShortDateTimeString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToShortDateTimeString())
	})

	s.Run("valid carbon", func() {
		s.Equal("20200805131415", Parse("2020-08-05 13:14:15").ToShortDateTimeString())
		s.Equal("20200805131415", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateTimeString())
		s.Equal("20200805000000", Parse("2020-08-05", PRC).ToShortDateTimeString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToShortDateTimeMilliString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToShortDateTimeMilliString())
	})

	s.Run("zero carbon", func() {
		s.Equal("00010101000000", NewCarbon().ToShortDateTimeMilliString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToShortDateTimeMilliString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToShortDateTimeMilliString())
	})

	s.Run("valid carbon", func() {
		s.Equal("20200805131415", Parse("2020-08-05 13:14:15").ToShortDateTimeMilliString())
		s.Equal("20200805131415.999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateTimeMilliString())
		s.Equal("20200805000000", Parse("2020-08-05", PRC).ToShortDateTimeMilliString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToShortDateTimeMicroString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToShortDateTimeMicroString())
	})

	s.Run("zero carbon", func() {
		s.Equal("00010101000000", NewCarbon().ToShortDateTimeMicroString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToShortDateTimeMicroString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToShortDateTimeMicroString())
	})

	s.Run("valid carbon", func() {
		s.Equal("20200805131415", Parse("2020-08-05 13:14:15").ToShortDateTimeMicroString())
		s.Equal("20200805131415.999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateTimeMicroString())
		s.Equal("20200805000000", Parse("2020-08-05", PRC).ToShortDateTimeMicroString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToShortDateTimeNanoString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToShortDateTimeNanoString())
	})

	s.Run("zero carbon", func() {
		s.Equal("00010101000000", NewCarbon().ToShortDateTimeNanoString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToShortDateTimeNanoString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToShortDateTimeNanoString())
	})

	s.Run("valid carbon", func() {
		s.Equal("20200805131415", Parse("2020-08-05 13:14:15").ToShortDateTimeNanoString())
		s.Equal("20200805131415.999999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateTimeNanoString())
		s.Equal("20200805000000", Parse("2020-08-05", PRC).ToShortDateTimeNanoString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToDateString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToDateString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01", NewCarbon().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToDateString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToDateString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05", Parse("2020-08-05 13:14:15").ToDateString())
		s.Equal("2020-08-05", Parse("2020-08-05T13:14:15.999999999+00:00").ToDateString())
		s.Equal("2020-08-05", Parse("2020-08-05", PRC).ToDateString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToDateMilliString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToDateMilliString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01", NewCarbon().ToDateMilliString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToDateMilliString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToDateMilliString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05", Parse("2020-08-05 13:14:15").ToDateMilliString())
		s.Equal("2020-08-05.999", Parse("2020-08-05T13:14:15.999999999+00:00").ToDateMilliString())
		s.Equal("2020-08-05", Parse("2020-08-05", PRC).ToDateMilliString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToDateMicroString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToDateMicroString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01", NewCarbon().ToDateMicroString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToDateMicroString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToDateMicroString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05", Parse("2020-08-05 13:14:15").ToDateMicroString())
		s.Equal("2020-08-05.999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToDateMicroString())
		s.Equal("2020-08-05", Parse("2020-08-05", PRC).ToDateMicroString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToDateNanoString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToDateNanoString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01", NewCarbon().ToDateNanoString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToDateNanoString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToDateNanoString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05", Parse("2020-08-05 13:14:15").ToDateNanoString())
		s.Equal("2020-08-05.999999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToDateNanoString())
		s.Equal("2020-08-05", Parse("2020-08-05", PRC).ToDateNanoString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToShortDateString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToShortDateString())
	})

	s.Run("zero carbon", func() {
		s.Equal("00010101", NewCarbon().ToShortDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToShortDateString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToShortDateString())
	})

	s.Run("valid carbon", func() {
		s.Equal("20200805", Parse("2020-08-05 13:14:15").ToShortDateString())
		s.Equal("20200805", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateString())
		s.Equal("20200805", Parse("2020-08-05", PRC).ToShortDateString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToShortDateMilliString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToShortDateMilliString())
	})

	s.Run("zero carbon", func() {
		s.Equal("00010101", NewCarbon().ToShortDateMilliString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToShortDateMilliString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToShortDateMilliString())
	})

	s.Run("valid carbon", func() {
		s.Equal("20200805", Parse("2020-08-05 13:14:15").ToShortDateMilliString())
		s.Equal("20200805.999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateMilliString())
		s.Equal("20200805", Parse("2020-08-05", PRC).ToShortDateMilliString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToShortDateMicroString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToShortDateMicroString())
	})

	s.Run("zero carbon", func() {
		s.Equal("00010101", NewCarbon().ToShortDateMicroString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToShortDateMicroString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToShortDateMicroString())
	})

	s.Run("valid carbon", func() {
		s.Equal("20200805", Parse("2020-08-05 13:14:15").ToShortDateMicroString())
		s.Equal("20200805.999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateMicroString())
		s.Equal("20200805", Parse("2020-08-05", PRC).ToShortDateMicroString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToShortDateNanoString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToShortDateNanoString())
	})

	s.Run("zero carbon", func() {
		s.Equal("00010101", NewCarbon().ToShortDateNanoString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToShortDateNanoString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToShortDateNanoString())
	})

	s.Run("valid carbon", func() {
		s.Equal("20200805", Parse("2020-08-05 13:14:15").ToShortDateNanoString())
		s.Equal("20200805.999999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateNanoString())
		s.Equal("20200805", Parse("2020-08-05", PRC).ToShortDateNanoString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToTimeString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToTimeString())
	})

	s.Run("zero carbon", func() {
		s.Equal("00:00:00", NewCarbon().ToTimeString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToTimeString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToTimeString())
	})

	s.Run("valid carbon", func() {
		s.Equal("13:14:15", Parse("2020-08-05 13:14:15").ToTimeString())
		s.Equal("13:14:15", Parse("2020-08-05T13:14:15.999999999+00:00").ToTimeString())
		s.Equal("00:00:00", Parse("2020-08-05", PRC).ToTimeString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToTimeMilliString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToTimeMilliString())
	})

	s.Run("zero carbon", func() {
		s.Equal("00:00:00", NewCarbon().ToTimeMilliString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToTimeMilliString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToTimeMilliString())
	})

	s.Run("valid carbon", func() {
		s.Equal("13:14:15", Parse("2020-08-05 13:14:15").ToTimeMilliString())
		s.Equal("13:14:15.999", Parse("2020-08-05T13:14:15.999999999+00:00").ToTimeMilliString())
		s.Equal("00:00:00", Parse("2020-08-05", PRC).ToTimeMilliString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToTimeMicroString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToTimeMicroString())
	})

	s.Run("zero carbon", func() {
		s.Equal("00:00:00", NewCarbon().ToTimeMicroString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToTimeMicroString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToTimeMicroString())
	})

	s.Run("valid carbon", func() {
		s.Equal("13:14:15", Parse("2020-08-05 13:14:15").ToTimeMicroString())
		s.Equal("13:14:15.999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToTimeMicroString())
		s.Equal("00:00:00", Parse("2020-08-05", PRC).ToTimeMicroString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToTimeNanoString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToTimeNanoString())
	})

	s.Run("zero carbon", func() {
		s.Equal("00:00:00", NewCarbon().ToTimeNanoString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToTimeNanoString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToTimeNanoString())
	})

	s.Run("valid carbon", func() {
		s.Equal("13:14:15", Parse("2020-08-05 13:14:15").ToTimeNanoString())
		s.Equal("13:14:15.999999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToTimeNanoString())
		s.Equal("00:00:00", Parse("2020-08-05", PRC).ToTimeNanoString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToShortTimeString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToShortTimeString())
	})

	s.Run("zero carbon", func() {
		s.Equal("000000", NewCarbon().ToShortTimeString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToShortTimeString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToShortTimeString())
	})

	s.Run("valid carbon", func() {
		s.Equal("131415", Parse("2020-08-05 13:14:15").ToShortTimeString())
		s.Equal("131415", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortTimeString())
		s.Equal("000000", Parse("2020-08-05", PRC).ToShortTimeString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToShortTimeMilliString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToShortTimeMilliString())
	})

	s.Run("zero carbon", func() {
		s.Equal("000000", NewCarbon().ToShortTimeMilliString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToShortTimeMilliString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToShortTimeMilliString())
	})

	s.Run("valid carbon", func() {
		s.Equal("131415", Parse("2020-08-05 13:14:15").ToShortTimeMilliString())
		s.Equal("131415.999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortTimeMilliString())
		s.Equal("000000", Parse("2020-08-05", PRC).ToShortTimeMilliString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToShortTimeMicroString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToShortTimeMicroString())
	})

	s.Run("zero carbon", func() {
		s.Equal("000000", NewCarbon().ToShortTimeMicroString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToShortTimeMicroString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToShortTimeMicroString())
	})

	s.Run("valid carbon", func() {
		s.Equal("131415", Parse("2020-08-05 13:14:15").ToShortTimeMicroString())
		s.Equal("131415.999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortTimeMicroString())
		s.Equal("000000", Parse("2020-08-05", PRC).ToShortTimeMicroString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToShortTimeNanoString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToShortTimeNanoString())
	})

	s.Run("zero carbon", func() {
		s.Equal("000000", NewCarbon().ToShortTimeNanoString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToShortTimeNanoString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToShortTimeNanoString())
	})

	s.Run("valid carbon", func() {
		s.Equal("131415", Parse("2020-08-05 13:14:15").ToShortTimeNanoString())
		s.Equal("131415.999999999", Parse("2020-08-05T13:14:15.999999999+00:00").ToShortTimeNanoString())
		s.Equal("000000", Parse("2020-08-05", PRC).ToShortTimeNanoString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToAtomString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToAtomString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01T00:00:00Z", NewCarbon().ToAtomString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToAtomString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToAtomString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToAtomString())
		s.Equal("2020-08-05T13:14:15Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToAtomString())
		s.Equal("2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToAtomString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToAnsicString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToAnsicString())
	})

	s.Run("zero carbon", func() {
		s.Equal("Mon Jan  1 00:00:00 0001", NewCarbon().ToAnsicString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToAnsicString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToAnsicString())
	})

	s.Run("valid carbon", func() {
		s.Equal("Wed Aug  5 13:14:15 2020", Parse("2020-08-05 13:14:15").ToAnsicString())
		s.Equal("Wed Aug  5 13:14:15 2020", Parse("2020-08-05T13:14:15.999999999+00:00").ToAnsicString())
		s.Equal("Wed Aug  5 00:00:00 2020", Parse("2020-08-05", PRC).ToAnsicString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToCookieString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToCookieString())
	})

	s.Run("zero carbon", func() {
		s.Equal("Monday, 01-Jan-0001 00:00:00 UTC", NewCarbon().ToCookieString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToCookieString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToCookieString())
	})

	s.Run("valid carbon", func() {
		s.Equal("Wednesday, 05-Aug-2020 13:14:15 UTC", Parse("2020-08-05 13:14:15").ToCookieString())
		s.Equal("Wednesday, 05-Aug-2020 13:14:15 UTC", Parse("2020-08-05T13:14:15.999999999+00:00").ToCookieString())
		s.Equal("Wednesday, 05-Aug-2020 00:00:00 CST", Parse("2020-08-05", PRC).ToCookieString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToRssString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToRssString())
	})

	s.Run("zero carbon", func() {
		s.Equal("Mon, 01 Jan 0001 00:00:00 +0000", NewCarbon().ToRssString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToRssString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToRssString())
	})

	s.Run("valid carbon", func() {
		s.Equal("Wed, 05 Aug 2020 13:14:15 +0000", Parse("2020-08-05 13:14:15").ToRssString())
		s.Equal("Wed, 05 Aug 2020 13:14:15 +0000", Parse("2020-08-05T13:14:15.999999999+00:00").ToRssString())
		s.Equal("Wed, 05 Aug 2020 00:00:00 +0800", Parse("2020-08-05", PRC).ToRssString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToW3cString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToW3cString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01T00:00:00Z", NewCarbon().ToW3cString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToW3cString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToW3cString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToW3cString())
		s.Equal("2020-08-05T13:14:15Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToW3cString())
		s.Equal("2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToW3cString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToUnixDateString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToUnixDateString())
	})

	s.Run("zero carbon", func() {
		s.Equal("Mon Jan  1 00:00:00 UTC 0001", NewCarbon().ToUnixDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToUnixDateString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToUnixDateString())
	})

	s.Run("valid carbon", func() {
		s.Equal("Wed Aug  5 13:14:15 UTC 2020", Parse("2020-08-05 13:14:15").ToUnixDateString())
		s.Equal("Wed Aug  5 13:14:15 UTC 2020", Parse("2020-08-05T13:14:15.999999999+00:00").ToUnixDateString())
		s.Equal("Wed Aug  5 00:00:00 CST 2020", Parse("2020-08-05", PRC).ToUnixDateString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToRubyDateString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToRubyDateString())
	})

	s.Run("zero carbon", func() {
		s.Equal("Mon Jan 01 00:00:00 +0000 0001", NewCarbon().ToRubyDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToRubyDateString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToRubyDateString())
	})

	s.Run("valid carbon", func() {
		s.Equal("Wed Aug 05 13:14:15 +0000 2020", Parse("2020-08-05 13:14:15").ToRubyDateString())
		s.Equal("Wed Aug 05 13:14:15 +0000 2020", Parse("2020-08-05T13:14:15.999999999+00:00").ToRubyDateString())
		s.Equal("Wed Aug 05 00:00:00 +0800 2020", Parse("2020-08-05", PRC).ToRubyDateString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToKitchenString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToKitchenString())
	})

	s.Run("zero carbon", func() {
		s.Equal("12:00AM", NewCarbon().ToKitchenString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToKitchenString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToKitchenString())
	})

	s.Run("valid carbon", func() {
		s.Equal("1:14PM", Parse("2020-08-05 13:14:15").ToKitchenString())
		s.Equal("1:14PM", Parse("2020-08-05T13:14:15.999999999+00:00").ToKitchenString())
		s.Equal("12:00AM", Parse("2020-08-05", PRC).ToKitchenString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToHttpString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToHttpString())
	})

	s.Run("zero carbon", func() {
		s.Equal("Mon, 01 Jan 0001 00:00:00 GMT", NewCarbon().ToHttpString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToHttpString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToHttpString())
	})

	s.Run("valid carbon", func() {
		s.Equal("Wed, 05 Aug 2020 13:14:15 GMT", Parse("2020-08-05 13:14:15").ToHttpString())
		s.Equal("Wed, 05 Aug 2020 13:14:15 GMT", Parse("2020-08-05T13:14:15.999999999+00:00").ToHttpString())
		s.Equal("Wed, 05 Aug 2020 00:00:00 GMT", Parse("2020-08-05", PRC).ToHttpString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToIso8601String() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToIso8601String())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01T00:00:00+00:00", NewCarbon().ToIso8601String())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToIso8601String())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToIso8601String())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05T13:14:15+00:00", Parse("2020-08-05 13:14:15").ToIso8601String())
		s.Equal("2020-08-05T13:14:15+00:00", Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601String())
		s.Equal("2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToIso8601String(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToIso8601MilliString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToIso8601MilliString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01T00:00:00+00:00", NewCarbon().ToIso8601MilliString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToIso8601MilliString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToIso8601MilliString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05T13:14:15+00:00", Parse("2020-08-05 13:14:15").ToIso8601MilliString())
		s.Equal("2020-08-05T13:14:15.999+00:00", Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601MilliString())
		s.Equal("2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToIso8601MilliString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_TToIso8601MicroString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToIso8601MicroString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01T00:00:00+00:00", NewCarbon().ToIso8601MicroString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToIso8601MicroString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToIso8601MicroString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05T13:14:15+00:00", Parse("2020-08-05 13:14:15").ToIso8601MicroString())
		s.Equal("2020-08-05T13:14:15.999999+00:00", Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601MicroString())
		s.Equal("2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToIso8601MicroString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToIso8601NanoString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToIso8601NanoString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01T00:00:00+00:00", NewCarbon().ToIso8601NanoString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToIso8601NanoString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToIso8601NanoString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05T13:14:15+00:00", Parse("2020-08-05 13:14:15").ToIso8601NanoString())
		s.Equal("2020-08-05T13:14:15.999999999+00:00", Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601NanoString())
		s.Equal("2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToIso8601NanoString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToIso8601ZuluString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToIso8601ZuluString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01T00:00:00Z", NewCarbon().ToIso8601ZuluString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToIso8601ZuluString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToIso8601ZuluString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToIso8601ZuluString())
		s.Equal("2020-08-05T13:14:15Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601ZuluString())
		s.Equal("2020-08-05T00:00:00Z", Parse("2020-08-05", PRC).ToIso8601ZuluString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToIso8601ZuluMilliString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToIso8601ZuluMilliString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01T00:00:00Z", NewCarbon().ToIso8601ZuluMilliString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToIso8601ZuluMilliString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToIso8601ZuluMilliString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToIso8601ZuluMilliString())
		s.Equal("2020-08-05T13:14:15.999Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601ZuluMilliString())
		s.Equal("2020-08-05T00:00:00Z", Parse("2020-08-05", PRC).ToIso8601ZuluMilliString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToIso8601ZuluMicroString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToIso8601ZuluMicroString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01T00:00:00Z", NewCarbon().ToIso8601ZuluMicroString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToIso8601ZuluMicroString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToIso8601ZuluMicroString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToIso8601ZuluMicroString())
		s.Equal("2020-08-05T13:14:15.999999Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601ZuluMicroString())
		s.Equal("2020-08-05T00:00:00Z", Parse("2020-08-05", PRC).ToIso8601ZuluMicroString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToIso8601ZuluNanoString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToIso8601ZuluNanoString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01T00:00:00Z", NewCarbon().ToIso8601ZuluNanoString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToIso8601ZuluNanoString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToIso8601ZuluNanoString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToIso8601ZuluNanoString())
		s.Equal("2020-08-05T13:14:15.999999999Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601ZuluNanoString())
		s.Equal("2020-08-05T00:00:00Z", Parse("2020-08-05", PRC).ToIso8601ZuluNanoString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToRfc822String() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToRfc822String())
	})

	s.Run("zero carbon", func() {
		s.Equal("01 Jan 01 00:00 UTC", NewCarbon().ToRfc822String())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToRfc822String())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToRfc822String())
	})

	s.Run("valid carbon", func() {
		s.Equal("05 Aug 20 13:14 UTC", Parse("2020-08-05 13:14:15").ToRfc822String())
		s.Equal("05 Aug 20 13:14 UTC", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc822String())
		s.Equal("05 Aug 20 00:00 CST", Parse("2020-08-05", PRC).ToRfc822String(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToRfc822zString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToRfc822zString())
	})

	s.Run("zero carbon", func() {
		s.Equal("01 Jan 01 00:00 +0000", NewCarbon().ToRfc822zString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToRfc822zString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToRfc822zString())
	})

	s.Run("valid carbon", func() {
		s.Equal("05 Aug 20 13:14 +0000", Parse("2020-08-05 13:14:15").ToRfc822zString())
		s.Equal("05 Aug 20 13:14 +0000", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc822zString())
		s.Equal("05 Aug 20 00:00 +0800", Parse("2020-08-05", PRC).ToRfc822zString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToRfc850String() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToRfc850String())
	})

	s.Run("zero carbon", func() {
		s.Equal("Monday, 01-Jan-01 00:00:00 UTC", NewCarbon().ToRfc850String())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToRfc850String())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToRfc850String())
	})

	s.Run("valid carbon", func() {
		s.Equal("Wednesday, 05-Aug-20 13:14:15 UTC", Parse("2020-08-05 13:14:15").ToRfc850String())
		s.Equal("Wednesday, 05-Aug-20 13:14:15 UTC", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc850String())
		s.Equal("Wednesday, 05-Aug-20 00:00:00 CST", Parse("2020-08-05", PRC).ToRfc850String(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToRfc1036String() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToRfc1036String())
	})

	s.Run("zero carbon", func() {
		s.Equal("Mon, 01 Jan 01 00:00:00 +0000", NewCarbon().ToRfc1036String())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToRfc1036String())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToRfc1036String())
	})

	s.Run("valid carbon", func() {
		s.Equal("Wed, 05 Aug 20 13:14:15 +0000", Parse("2020-08-05 13:14:15").ToRfc1036String())
		s.Equal("Wed, 05 Aug 20 13:14:15 +0000", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc1036String())
		s.Equal("Wed, 05 Aug 20 00:00:00 +0800", Parse("2020-08-05", PRC).ToRfc1036String(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToRfc1123String() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToRfc1123String())
	})

	s.Run("zero carbon", func() {
		s.Equal("Mon, 01 Jan 0001 00:00:00 UTC", NewCarbon().ToRfc1123String())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToRfc1123String())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToRfc1123String())
	})

	s.Run("valid carbon", func() {
		s.Equal("Wed, 05 Aug 2020 13:14:15 UTC", Parse("2020-08-05 13:14:15").ToRfc1123String())
		s.Equal("Wed, 05 Aug 2020 13:14:15 UTC", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc1123String())
		s.Equal("Wed, 05 Aug 2020 00:00:00 CST", Parse("2020-08-05", PRC).ToRfc1123String(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToRfc1123zString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToRfc1123zString())
	})

	s.Run("zero carbon", func() {
		s.Equal("Mon, 01 Jan 0001 00:00:00 +0000", NewCarbon().ToRfc1123zString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToRfc1123zString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToRfc1123zString())
	})

	s.Run("valid carbon", func() {
		s.Equal("Wed, 05 Aug 2020 13:14:15 +0000", Parse("2020-08-05 13:14:15").ToRfc1123zString())
		s.Equal("Wed, 05 Aug 2020 13:14:15 +0000", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc1123zString())
		s.Equal("Wed, 05 Aug 2020 00:00:00 +0800", Parse("2020-08-05", PRC).ToRfc1123zString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToRfc2822String() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToRfc2822String())
	})

	s.Run("zero carbon", func() {
		s.Equal("Mon, 01 Jan 0001 00:00:00 +0000", NewCarbon().ToRfc2822String())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToRfc2822String())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToRfc2822String())
	})

	s.Run("valid carbon", func() {
		s.Equal("Wed, 05 Aug 2020 13:14:15 +0000", Parse("2020-08-05 13:14:15").ToRfc2822String())
		s.Equal("Wed, 05 Aug 2020 13:14:15 +0000", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc2822String())
		s.Equal("Wed, 05 Aug 2020 00:00:00 +0800", Parse("2020-08-05", PRC).ToRfc2822String(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToRfc3339String() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToRfc3339String())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01T00:00:00Z", NewCarbon().ToRfc3339String())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToRfc3339String())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToRfc3339String())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToRfc3339String())
		s.Equal("2020-08-05T13:14:15Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc3339String())
		s.Equal("2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToRfc3339String(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToRfc3339MilliString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToRfc3339MilliString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01T00:00:00Z", NewCarbon().ToRfc3339MilliString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToRfc3339MilliString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToRfc3339MilliString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToRfc3339MilliString())
		s.Equal("2020-08-05T13:14:15.999Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc3339MilliString())
		s.Equal("2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToRfc3339MilliString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToRfc3339MicroString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToRfc3339MicroString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01T00:00:00Z", NewCarbon().ToRfc3339MicroString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToRfc3339MicroString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToRfc3339MicroString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToRfc3339MicroString())
		s.Equal("2020-08-05T13:14:15.999999Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc3339MicroString())
		s.Equal("2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToRfc3339MicroString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToRfc3339NanoString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToRfc3339NanoString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01T00:00:00Z", NewCarbon().ToRfc3339NanoString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToRfc3339NanoString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToRfc3339NanoString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05T13:14:15Z", Parse("2020-08-05 13:14:15").ToRfc3339NanoString())
		s.Equal("2020-08-05T13:14:15.999999999Z", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc3339NanoString())
		s.Equal("2020-08-05T00:00:00+08:00", Parse("2020-08-05", PRC).ToRfc3339NanoString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToRfc7231String() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToRfc7231String())
	})

	s.Run("zero carbon", func() {
		s.Equal("Mon, 01 Jan 0001 00:00:00 UTC", NewCarbon().ToRfc7231String())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToRfc7231String())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToRfc7231String())
	})

	s.Run("valid carbon", func() {
		s.Equal("Wed, 05 Aug 2020 13:14:15 UTC", Parse("2020-08-05 13:14:15").ToRfc7231String())
		s.Equal("Wed, 05 Aug 2020 13:14:15 UTC", Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc7231String())
		s.Equal("Wed, 05 Aug 2020 00:00:00 CST", Parse("2020-08-05", PRC).ToRfc7231String(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToFormattedDateString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToFormattedDateString())
	})

	s.Run("zero carbon", func() {
		s.Equal("Jan 1, 0001", NewCarbon().ToFormattedDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToFormattedDateString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToFormattedDateString())
	})

	s.Run("valid carbon", func() {
		s.Equal("Aug 5, 2020", Parse("2020-08-05 13:14:15").ToFormattedDateString())
		s.Equal("Aug 5, 2020", Parse("2020-08-05T13:14:15.999999999+00:00").ToFormattedDateString())
		s.Equal("Aug 5, 2020", Parse("2020-08-05", PRC).ToFormattedDateString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_ToFormattedDayDateString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.ToFormattedDayDateString())
	})

	s.Run("zero carbon", func() {
		s.Equal("Mon, Jan 1, 0001", NewCarbon().ToFormattedDayDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").ToFormattedDayDateString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").ToFormattedDayDateString())
	})

	s.Run("valid carbon", func() {
		s.Equal("Wed, Aug 5, 2020", Parse("2020-08-05 13:14:15").ToFormattedDayDateString())
		s.Equal("Wed, Aug 5, 2020", Parse("2020-08-05T13:14:15.999999999+00:00").ToFormattedDayDateString())
		s.Equal("Wed, Aug 5, 2020", Parse("2020-08-05", PRC).ToFormattedDayDateString(PRC))
	})
}

func (s *OutputerSuite) TestCarbon_Layout() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.Layout(DateTimeLayout))
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00", NewCarbon().Layout(DateTimeLayout))
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").Layout(DateTimeLayout))
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").Layout(DateTimeLayout))
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:14:15", Parse("2020-08-05 13:14:15").Layout(DateTimeLayout))
		s.Equal("2020-08-05 13:14:15", Parse("2020-08-05T13:14:15.999999999+00:00").Layout(DateTimeLayout))
		s.Equal("2020-08-05 00:00:00", Parse("2020-08-05", PRC).Layout(DateTimeLayout, PRC))
		s.Equal("2020年08月05日", Parse("2020-08-05 13:14:15").Layout("2006年01月02日"))
		s.Equal("Wed, 05 Aug 2020 13:14:15 GMT", Parse("2020-08-05 13:14:15").Layout("Mon, 02 Jan 2006 15:04:05 GMT"))

		s.Equal("1596633255", Parse("2020-08-05 13:14:15.999999999").Layout(TimestampLayout))
		s.Equal("1596633255999", Parse("2020-08-05 13:14:15.999999999").Layout(TimestampMilliLayout))
		s.Equal("1596633255999999", Parse("2020-08-05 13:14:15.999999999").Layout(TimestampMicroLayout))
		s.Equal("1596633255999999999", Parse("2020-08-05 13:14:15.999999999").Layout(TimestampNanoLayout))
	})
}

func (s *OutputerSuite) TestCarbon_Format() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.Format(DateTimeFormat))
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00", NewCarbon().Format(DateTimeFormat))
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").Format(DateTimeFormat))
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").Format(DateTimeFormat))
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:14:15", Parse("2020-08-05 13:14:15").Format(DateTimeFormat))
		s.Equal("2020-08-05", Parse("2020-08-05T13:14:15.999999999+00:00").Format(DateFormat))
		s.Equal("2020-08-05 00:00:00", Parse("2020-08-05", PRC).Format(DateTimeFormat, PRC))
		s.Equal("2020年08月05日", Parse("2020-08-05 13:14:15").Format("Y年m月d日"))
		s.Equal("Wed", Parse("2020-08-05 13:14:15").Format("D"))
		s.Equal("Wednesday", Parse("2020-08-05 13:14:15").Format("l"))
		s.Equal("August", Parse("2020-08-05 13:14:15").Format("F"))
		s.Equal("Aug", Parse("2020-08-05 13:14:15").Format("M"))
		s.Equal("5", Parse("2020-08-05 13:14:15").Format("j"))
		s.Equal("32", Parse("2020-08-05 13:14:15").Format("W"))
		s.Equal("August", Parse("2020-08-05 13:14:15").Format("F"))
		s.Equal("3", Parse("2020-08-05 13:14:15").Format("N"))
		s.Equal("1", Parse("2020-08-05 13:14:15").Format("L"))
		s.Equal("0", Parse("2021-08-05 13:14:15").Format("L"))
		s.Equal("13", Parse("2020-08-05 13:14:15").Format("G"))
		s.Equal("1596633255", Parse("2020-08-05 13:14:15").Format("S"))
		s.Equal("1596633255000", Parse("2020-08-05 13:14:15").Format("U"))
		s.Equal("1596633255000000", Parse("2020-08-05 13:14:15").Format("V"))
		s.Equal("1596633255000000000", Parse("2020-08-05 13:14:15").Format("X"))
		s.Equal("999", Parse("2020-08-05 13:14:15.999999999").Format("u"))
		s.Equal("999999", Parse("2020-08-05 13:14:15.999999999").Format("v"))
		s.Equal("999999999", Parse("2020-08-05 13:14:15.999999999").Format("x"))
		s.Equal("2", Parse("2020-08-05 13:14:15.999999999").Format("w"))
		s.Equal("31", Parse("2020-08-05 13:14:15.999999999").Format("t"))
		s.Equal("PRC", Parse("2020-08-05 13:14:15.999999999", PRC).Format("z"))
		s.Equal("28800", Parse("2020-08-05 13:14:15.999999999", PRC).Format("o"))
		s.Equal("3", Parse("2020-08-05 13:14:15.999999999").Format("q"))
		s.Equal("21", Parse("2020-08-05 13:14:15.999999999").Format("c"))
		s.Equal("Z", Parse("2020-08-05 13:14:15.999999999", UTC).Format("R"))
		s.Equal("+08:00", Parse("2020-08-05 13:14:15.999999999", PRC).Format("R"))
		s.Equal("Z", Parse("2020-08-05 13:14:15.999999999", UTC).Format("Q"))
		s.Equal("+0800", Parse("2020-08-05 13:14:15.999999999", PRC).Format("Q"))
		s.Equal("CST", Parse("2020-08-05 13:14:15.999999999", PRC).Format("Z"))
		s.Equal("5th", Parse("2020-08-05 13:14:15").Format("jK"))
		s.Equal("22nd", Parse("2020-08-22 13:14:15").Format("jK"))
		s.Equal("23rd", Parse("2020-08-23 13:14:15").Format("jK"))
		s.Equal("31st", Parse("2020-08-31 13:14:15").Format("jK"))
		s.Equal("It is 2020-08-31 13:14:15", Parse("2020-08-31 13:14:15").Format("I\\t \\i\\s Y-m-d H:i:s"))
		s.Equal("上次打卡时间:2020-08-05 13:14:15，请每日按时打卡", Parse("2020-08-05 13:14:15").Format("上次打卡时间:Y-m-d H:i:s，请每日按时打卡"))

		s.Equal("1596633255", Parse("2020-08-05 13:14:15.999999999").Format(TimestampFormat))
		s.Equal("1596633255999", Parse("2020-08-05 13:14:15.999999999").Format(TimestampMilliFormat))
		s.Equal("1596633255999999", Parse("2020-08-05 13:14:15.999999999").Format(TimestampMicroFormat))
		s.Equal("1596633255999999999", Parse("2020-08-05 13:14:15.999999999").Format(TimestampNanoFormat))
	})

	// https://github.com/dromara/carbon/issues/328
	s.Run("issue328", func() {
		inputTime := time.Date(2026, 1, 22, 11, 30, 52, 624_225_666, time.UTC)
		c := CreateFromStdTime(inputTime)

		s.Equal(inputTime.Format("2006-01-02 15:04:05.999"), c.Format("Y-m-d H:i:s.u"))
		s.Equal(inputTime.Format("2006-01-02 15:04:05.999999"), c.Format("Y-m-d H:i:s.v"))
		s.Equal(inputTime.Format("2006-01-02 15:04:05.999999999"), c.Format("Y-m-d H:i:s.x"))
	})
}

func (s *OutputerSuite) TestFormat2Layout() {
	c := Parse("2020-08-05 13:14:15.999999999")

	s.Run("without timezone", func() {
		s.Equal(c.Layout(AtomLayout), c.Format(AtomFormat))
		s.Equal(c.Layout(ANSICLayout), c.Format(ANSICFormat))
		s.Equal(c.Layout(CookieLayout), c.Format(CookieFormat))
		s.Equal(c.Layout(KitchenLayout), c.Format(KitchenFormat))
		s.Equal(c.Layout(RssLayout), c.Format(RssFormat))
		s.Equal(c.Layout(RubyDateLayout), c.Format(RubyDateFormat))
		s.Equal(c.Layout(UnixDateLayout), c.Format(UnixDateFormat))
		s.Equal(c.Layout(W3cLayout), c.Format(W3cFormat))
		s.Equal(c.Layout(HttpLayout), c.Format(HttpFormat))

		s.Equal(c.Layout(RFC1036Layout), c.Format(RFC1036Format))
		s.Equal(c.Layout(RFC1123Layout), c.Format(RFC1123Format))
		s.Equal(c.Layout(RFC1123ZLayout), c.Format(RFC1123ZFormat))
		s.Equal(c.Layout(RFC2822Layout), c.Format(RFC2822Format))
		s.Equal(c.Layout(RFC3339Layout), c.Format(RFC3339Format))
		s.Equal(c.Layout(RFC3339MilliLayout), c.Format(RFC3339MilliFormat))
		s.Equal(c.Layout(RFC3339MicroLayout), c.Format(RFC3339MicroFormat))
		s.Equal(c.Layout(RFC3339NanoLayout), c.Format(RFC3339NanoFormat))
		s.Equal(c.Layout(RFC7231Layout), c.Format(RFC7231Format))
		s.Equal(c.Layout(RFC822Layout), c.Format(RFC822Format))
		s.Equal(c.Layout(RFC822ZLayout), c.Format(RFC822ZFormat))
		s.Equal(c.Layout(RFC850Layout), c.Format(RFC850Format))

		s.Equal(c.Layout(ISO8601Layout), c.Format(ISO8601Format))
		s.Equal(c.Layout(ISO8601MilliLayout), c.Format(ISO8601MilliFormat))
		s.Equal(c.Layout(ISO8601MicroLayout), c.Format(ISO8601MicroFormat))
		s.Equal(c.Layout(ISO8601NanoLayout), c.Format(ISO8601NanoFormat))

		s.Equal(c.Layout(ISO8601ZuluLayout), c.Format(ISO8601ZuluFormat))
		s.Equal(c.Layout(ISO8601ZuluMilliLayout), c.Format(ISO8601ZuluMilliFormat))
		s.Equal(c.Layout(ISO8601ZuluMicroLayout), c.Format(ISO8601ZuluMicroFormat))
		s.Equal(c.Layout(ISO8601ZuluNanoLayout), c.Format(ISO8601ZuluNanoFormat))

		s.Equal(c.Layout(FormattedDateLayout), c.Format(FormattedDateFormat))
		s.Equal(c.Layout(FormattedDayDateLayout), c.Format(FormattedDayDateFormat))

		s.Equal(c.Layout(DayDateTimeLayout), c.Format(DayDateTimeFormat))
		s.Equal(c.Layout(DateTimeLayout), c.Format(DateTimeFormat))
		s.Equal(c.Layout(DateTimeMilliLayout), c.Format(DateTimeMilliFormat))
		s.Equal(c.Layout(DateTimeMicroLayout), c.Format(DateTimeMicroFormat))
		s.Equal(c.Layout(DateTimeNanoLayout), c.Format(DateTimeNanoFormat))
		s.Equal(c.Layout(ShortDateTimeLayout), c.Format(ShortDateTimeFormat))
		s.Equal(c.Layout(ShortDateTimeMilliLayout), c.Format(ShortDateTimeMilliFormat))
		s.Equal(c.Layout(ShortDateTimeMicroLayout), c.Format(ShortDateTimeMicroFormat))
		s.Equal(c.Layout(ShortDateTimeNanoLayout), c.Format(ShortDateTimeNanoFormat))

		s.Equal(c.Layout(DateLayout), c.Format(DateFormat))
		s.Equal(c.Layout(DateMilliLayout), c.Format(DateMilliFormat))
		s.Equal(c.Layout(DateMicroLayout), c.Format(DateMicroFormat))
		s.Equal(c.Layout(DateNanoLayout), c.Format(DateNanoFormat))
		s.Equal(c.Layout(ShortDateLayout), c.Format(ShortDateFormat))
		s.Equal(c.Layout(ShortDateMilliLayout), c.Format(ShortDateMilliFormat))
		s.Equal(c.Layout(ShortDateMicroLayout), c.Format(ShortDateMicroFormat))
		s.Equal(c.Layout(ShortDateNanoLayout), c.Format(ShortDateNanoFormat))

		s.Equal(c.Layout(TimeLayout), c.Format(TimeFormat))
		s.Equal(c.Layout(TimeMilliLayout), c.Format(TimeMilliFormat))
		s.Equal(c.Layout(TimeMicroLayout), c.Format(TimeMicroFormat))
		s.Equal(c.Layout(TimeNanoLayout), c.Format(TimeNanoFormat))
		s.Equal(c.Layout(ShortTimeLayout), c.Format(ShortTimeFormat))
		s.Equal(c.Layout(ShortTimeMilliLayout), c.Format(ShortTimeMilliFormat))
		s.Equal(c.Layout(ShortTimeMicroLayout), c.Format(ShortTimeMicroFormat))
		s.Equal(c.Layout(ShortTimeNanoLayout), c.Format(ShortTimeNanoFormat))
	})

	s.Run("with timezone", func() {
		s.Equal(c.Layout(AtomLayout, PRC), c.Format(AtomFormat, PRC))
		s.Equal(c.Layout(ANSICLayout, PRC), c.Format(ANSICFormat, PRC))
		s.Equal(c.Layout(CookieLayout, PRC), c.Format(CookieFormat, PRC))
		s.Equal(c.Layout(KitchenLayout, PRC), c.Format(KitchenFormat, PRC))
		s.Equal(c.Layout(RssLayout, PRC), c.Format(RssFormat, PRC))
		s.Equal(c.Layout(RubyDateLayout, PRC), c.Format(RubyDateFormat, PRC))
		s.Equal(c.Layout(UnixDateLayout, PRC), c.Format(UnixDateFormat, PRC))
		s.Equal(c.Layout(W3cLayout), c.Format(W3cFormat))
		s.Equal(c.Layout(HttpLayout), c.Format(HttpFormat))

		s.Equal(c.Layout(RFC1036Layout, PRC), c.Format(RFC1036Format, PRC))
		s.Equal(c.Layout(RFC1123Layout, PRC), c.Format(RFC1123Format, PRC))
		s.Equal(c.Layout(RFC1123ZLayout, PRC), c.Format(RFC1123ZFormat, PRC))
		s.Equal(c.Layout(RFC2822Layout, PRC), c.Format(RFC2822Format, PRC))
		s.Equal(c.Layout(RFC3339Layout, PRC), c.Format(RFC3339Format, PRC))
		s.Equal(c.Layout(RFC3339MilliLayout, PRC), c.Format(RFC3339MilliFormat, PRC))
		s.Equal(c.Layout(RFC3339MicroLayout, PRC), c.Format(RFC3339MicroFormat, PRC))
		s.Equal(c.Layout(RFC3339NanoLayout, PRC), c.Format(RFC3339NanoFormat, PRC))
		s.Equal(c.Layout(RFC7231Layout, PRC), c.Format(RFC7231Format, PRC))
		s.Equal(c.Layout(RFC822Layout, PRC), c.Format(RFC822Format, PRC))
		s.Equal(c.Layout(RFC822ZLayout, PRC), c.Format(RFC822ZFormat, PRC))
		s.Equal(c.Layout(RFC850Layout, PRC), c.Format(RFC850Format, PRC))

		s.Equal(c.Layout(ISO8601Layout, PRC), c.Format(ISO8601Format, PRC))
		s.Equal(c.Layout(ISO8601MilliLayout, PRC), c.Format(ISO8601MilliFormat, PRC))
		s.Equal(c.Layout(ISO8601MicroLayout, PRC), c.Format(ISO8601MicroFormat, PRC))
		s.Equal(c.Layout(ISO8601NanoLayout, PRC), c.Format(ISO8601NanoFormat, PRC))

		s.Equal(c.Layout(ISO8601ZuluLayout, PRC), c.Format(ISO8601ZuluFormat, PRC))
		s.Equal(c.Layout(ISO8601ZuluMilliLayout, PRC), c.Format(ISO8601ZuluMilliFormat, PRC))
		s.Equal(c.Layout(ISO8601ZuluMicroLayout, PRC), c.Format(ISO8601ZuluMicroFormat, PRC))
		s.Equal(c.Layout(ISO8601ZuluNanoLayout, PRC), c.Format(ISO8601ZuluNanoFormat, PRC))

		s.Equal(c.Layout(FormattedDateLayout, PRC), c.Format(FormattedDateFormat, PRC))
		s.Equal(c.Layout(FormattedDayDateLayout, PRC), c.Format(FormattedDayDateFormat, PRC))

		s.Equal(c.Layout(DayDateTimeLayout, PRC), c.Format(DayDateTimeFormat, PRC))
		s.Equal(c.Layout(DateTimeLayout, PRC), c.Format(DateTimeFormat, PRC))
		s.Equal(c.Layout(DateTimeMilliLayout, PRC), c.Format(DateTimeMilliFormat, PRC))
		s.Equal(c.Layout(DateTimeMicroLayout, PRC), c.Format(DateTimeMicroFormat, PRC))
		s.Equal(c.Layout(DateTimeNanoLayout, PRC), c.Format(DateTimeNanoFormat, PRC))
		s.Equal(c.Layout(ShortDateTimeLayout, PRC), c.Format(ShortDateTimeFormat, PRC))
		s.Equal(c.Layout(ShortDateTimeMilliLayout, PRC), c.Format(ShortDateTimeMilliFormat, PRC))
		s.Equal(c.Layout(ShortDateTimeMicroLayout, PRC), c.Format(ShortDateTimeMicroFormat, PRC))
		s.Equal(c.Layout(ShortDateTimeNanoLayout, PRC), c.Format(ShortDateTimeNanoFormat, PRC))

		s.Equal(c.Layout(DateLayout, PRC), c.Format(DateFormat, PRC))
		s.Equal(c.Layout(DateMilliLayout, PRC), c.Format(DateMilliFormat, PRC))
		s.Equal(c.Layout(DateMicroLayout, PRC), c.Format(DateMicroFormat, PRC))
		s.Equal(c.Layout(DateNanoLayout, PRC), c.Format(DateNanoFormat, PRC))
		s.Equal(c.Layout(ShortDateLayout, PRC), c.Format(ShortDateFormat, PRC))
		s.Equal(c.Layout(ShortDateMilliLayout, PRC), c.Format(ShortDateMilliFormat, PRC))
		s.Equal(c.Layout(ShortDateMicroLayout, PRC), c.Format(ShortDateMicroFormat, PRC))
		s.Equal(c.Layout(ShortDateNanoLayout, PRC), c.Format(ShortDateNanoFormat, PRC))

		s.Equal(c.Layout(TimeLayout, PRC), c.Format(TimeFormat, PRC))
		s.Equal(c.Layout(TimeMilliLayout, PRC), c.Format(TimeMilliFormat, PRC))
		s.Equal(c.Layout(TimeMicroLayout, PRC), c.Format(TimeMicroFormat, PRC))
		s.Equal(c.Layout(TimeNanoLayout, PRC), c.Format(TimeNanoFormat, PRC))
		s.Equal(c.Layout(ShortTimeLayout, PRC), c.Format(ShortTimeFormat, PRC))
		s.Equal(c.Layout(ShortTimeMilliLayout, PRC), c.Format(ShortTimeMilliFormat, PRC))
		s.Equal(c.Layout(ShortTimeMicroLayout, PRC), c.Format(ShortTimeMicroFormat, PRC))
		s.Equal(c.Layout(ShortTimeNanoLayout, PRC), c.Format(ShortTimeNanoFormat, PRC))
	})
}
