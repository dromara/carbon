package carbon

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ParserSuite struct {
	suite.Suite
}

func TestParserSuite(t *testing.T) {
	suite.Run(t, new(ParserSuite))
}

func (s *ParserSuite) TestParse() {
	s.Run("empty value", func() {
		c := Parse("")
		s.False(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error value", func() {
		c := Parse("xxx")
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("empty timezone", func() {
		c := Parse("2020-08-05", "")
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error timezone", func() {
		c := Parse("2020-08-05", "xxx")
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("without timezone", func() {
		s.Equal(Now().Timestamp(), Parse("now").Timestamp())
		s.Equal(Yesterday().Timestamp(), Parse("yesterday").Timestamp())
		s.Equal(Tomorrow().Timestamp(), Parse("tomorrow").Timestamp())
		s.Equal("2020-08-05 00:00:00 +0000 UTC", Parse("2020-8-5").ToString())
		s.Equal("2020-08-05 00:00:00 +0000 UTC", Parse("2020-8-05").ToString())
		s.Equal("2020-08-05 00:00:00 +0000 UTC", Parse("2020-08-05").ToString())
		s.Equal("2020-08-05 01:02:03 +0000 UTC", Parse("2020-8-5 1:2:3").ToString())
		s.Equal("2020-08-05 01:02:03 +0000 UTC", Parse("2020-08-05 1:2:03").ToString())
		s.Equal("2020-08-05 01:02:03 +0000 UTC", Parse("2020-08-05 1:02:03").ToString())
		s.Equal("2020-08-05 01:02:03 +0000 UTC", Parse("2020-08-05 01:02:03").ToString())
		s.Equal(Parse("2022-03-08T03:01:14-07:00").ToString(), Parse("2022-03-08T10:01:14Z").ToString())
	})

	s.Run("with timezone", func() {
		s.Equal(Now().Timestamp(), Parse("now", PRC).Timestamp())
		s.Equal(Yesterday().Timestamp(), Parse("yesterday", PRC).Timestamp())
		s.Equal(Tomorrow().Timestamp(), Parse("tomorrow", PRC).Timestamp())
		s.Equal("2020-08-05 00:00:00 +0800 CST", Parse("2020-8-5", PRC).ToString())
		s.Equal("2020-08-05 00:00:00 +0800 CST", Parse("2020-8-05", PRC).ToString())
		s.Equal("2020-08-05 00:00:00 +0800 CST", Parse("2020-08-05", PRC).ToString())
		s.Equal("2020-08-05 01:02:03 +0800 CST", Parse("2020-8-5 1:2:3", PRC).ToString())
		s.Equal("2020-08-05 01:02:03 +0800 CST", Parse("2020-08-05 1:2:03", PRC).ToString())
		s.Equal("2020-08-05 01:02:03 +0800 CST", Parse("2020-08-05 1:02:03", PRC).ToString())
		s.Equal("2020-08-05 01:02:03 +0800 CST", Parse("2020-08-05 01:02:03", PRC).ToString())
		s.Equal(Parse("2022-03-08T03:01:14-07:00", PRC).ToString(), Parse("2022-03-08T10:01:14Z", PRC).ToString())
	})

	s.Run("postgres time type", func() {
		// date type
		s.Equal("2020-08-05 00:00:00 +0000 UTC", Parse("2020-08-05").ToString())

		// time
		s.Equal("0000-01-01 13:14:15 +0000 UTC", Parse("13:14:15").ToString())
		// timetz
		s.Equal("0000-01-01 05:14:15 +0000 UTC", Parse("13:14:15+08").ToString())

		// timestamp
		s.Equal("2020-08-05 13:14:15 +0000 UTC", Parse("2020-08-05 13:14:15").ToString())
		// timestamptz
		s.Equal("2020-08-05 05:14:15 +0000 UTC", Parse("2020-08-05 13:14:15+08").ToString())
	})

	s.Run("sqlserver time type", func() {
		// date type
		s.Equal("2020-08-05 00:00:00 +0000 UTC", Parse("2020-08-05").ToString())

		// time type
		s.Equal("0000-01-01 13:14:15 +0000 UTC", Parse("13:14:15.0000000").ToString())
		s.Equal("0000-01-01 13:14:15.9999999 +0000 UTC", Parse("13:14:15.9999999").ToString())

		// smalldatetime type
		s.Equal("2020-08-05 13:14:15 +0000 UTC", Parse("2020-08-05 13:14:15").ToString())

		// datetime type
		s.Equal("2020-08-05 13:14:15 +0000 UTC", Parse("2020-08-05 13:14:15.000").ToString())
		s.Equal("2020-08-05 13:14:15.999 +0000 UTC", Parse("2020-08-05 13:14:15.999").ToString())

		// datetime2 type
		s.Equal("2020-08-05 13:14:15 +0000 UTC", Parse("2020-08-05 13:14:15.0000000").ToString())
		s.Equal("2020-08-05 13:14:15.9999999 +0000 UTC", Parse("2020-08-05 13:14:15.9999999").ToString())

		// datetimeoffset type
		s.Equal("2020-08-05 13:14:15 +0000 UTC", Parse("2020-08-05 13:14:15.0000000 +00:00").ToString())
		s.Equal("2020-08-05 05:14:15.9999999 +0000 UTC", Parse("2020-08-05 13:14:15.9999999 +08:00").ToString())
	})

	s.Run("diverse separators and timezone variants", func() {
		cases := [][2]string{
			{"2020-08-05 13:14:15", "2020/8/5 13:14:15"},          // slash
			{"2020-08-05 13:14:15", "2020.8.5 13:14:15"},          // dot
			{"2020-08-05 13:14:15", "2020-08-05T13:14:15"},        // with T
			{"2020-08-05T13:14:15Z", "2020-08-05T13:14:15+00:00"}, // Z vs +00:00
			{"2020-08-05T13:14:15+02:00", "2020-08-05T11:14:15Z"}, // with offset
			{"2020-08-05T13:14:15-0700", "2020-08-05T20:14:15Z"},  // numeric offset
		}
		for _, pair := range cases {
			left := Parse(pair[0])
			right := Parse(pair[1])
			s.False(left.HasError())
			s.False(right.HasError())
			s.Equal(left.Timestamp(), right.Timestamp())
		}

		withTz := [][2]string{
			{"2020-08-05 13:14:15", "2020/8/5 13:14:15"},
			{"2020-08-05T13:14:15Z", "2020-08-05T21:14:15+08:00"},
			{"2020-08-05T13:14:15+02:00", "2020-08-05T19:14:15+08:00"},
		}
		for _, pair := range withTz {
			left := Parse(pair[0], PRC)
			right := Parse(pair[1], PRC)
			s.False(left.HasError())
			s.False(right.HasError())
			s.Equal(left.Timestamp(), right.Timestamp())
		}
	})

	// https://github.com/dromara/carbon/issues/202
	s.Run("issue202", func() {
		s.Equal("2023-01-08 09:02:48 +0000 UTC", Parse("2023-01-08T09:02:48").ToString())
		s.Equal("2023-01-08 09:02:48 +0000 UTC", Parse("2023-1-8T09:02:48").ToString())
		s.Equal("2023-01-08 09:02:48 +0000 UTC", Parse("2023-01-08T9:2:48").ToString())
		s.Equal("2023-01-08 09:02:48 +0000 UTC", Parse("2023-01-8T9:2:48").ToString())
	})

	// https://github.com/dromara/carbon/issues/232
	s.Run("issue232", func() {
		s.Equal("0000-01-01 00:00:00 +0000 UTC", Parse("0000-01-01 00:00:00").ToString())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", Parse("0001-01-01 00:00:00").ToString())
		s.Equal("", Parse("0001-00-00 00:00:00").ToString())
	})
}

func (s *ParserSuite) TestParseByLayout() {
	s.Run("empty value", func() {
		c := ParseByLayout("", DateFormat)
		s.False(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error value", func() {
		c := ParseByLayout("xxx", DateFormat)
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("empty layout", func() {
		c := ParseByLayout("2020-08-05", "")
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error layout", func() {
		c := ParseByLayout("2020-08-05", "xxx")
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("empty timezone", func() {
		c := ParseByLayout("2020-08-05", DateLayout, "")
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error timezone", func() {
		c := ParseByLayout("2020-08-05", DateLayout, "xxx")
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error timestamp", func() {
		s.Error(ParseByLayout("2020-08-05", TimestampLayout).Error)
		s.Error(ParseByLayout("2020-08-05", TimestampMilliLayout).Error)
		s.Error(ParseByLayout("2020-08-05", TimestampMicroLayout).Error)
		s.Error(ParseByLayout("2020-08-05", TimestampNanoLayout).Error)
	})

	s.Run("without timezone", func() {
		s.Equal("2020-08-05 00:00:00 +0000 UTC", ParseByLayout("2020-08-05", DateLayout).ToString())
		s.Equal("0000-01-01 13:14:15 +0000 UTC", ParseByLayout("13:14:15", TimeLayout).ToString())
		s.Equal("2020-08-05 13:14:15 +0000 UTC", ParseByLayout("2020-08-05 13:14:15", DateTimeLayout).ToString())

		s.Equal("2020-08-05 13:14:15 +0000 UTC", ParseByLayout("2020|08|05 13:14:15", "2006|01|02 15:04:05").ToString())
		s.Equal("2020-08-05 13:14:15 +0000 UTC", ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToString())
		s.Equal("2020-08-05 13:14:15 +0000 UTC", ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToString())
	})

	s.Run("with timezone", func() {
		s.Equal("2020-08-05 00:00:00 +0800 CST", ParseByLayout("2020-08-05", DateLayout, PRC).ToString())
		s.Equal("2020-08-05 13:14:15 +0800 CST", ParseByLayout("2020-08-05 13:14:15", DateTimeLayout, PRC).ToString())
		s.Equal("2020-08-05 13:14:15 +0800 CST", ParseByLayout("2020-08-05 13:14:15", DateTimeLayout, PRC).ToString())

		s.Equal("2020-08-05 13:14:15 +0800 CST", ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s 2006-01-02 15:04:05", PRC).ToString())
		s.Equal("2020-08-05 13:14:15 +0800 CST", ParseByLayout("2020|08|05 13:14:15", "2006|01|02 15:04:05", PRC).ToString())
		s.Equal("2020-08-05 13:14:15 +0800 CST", ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒", PRC).ToString())
	})
}

func (s *ParserSuite) TestParseByLayouts() {
	s.Run("empty value", func() {
		c := ParseByLayouts("", []string{DateTimeLayout})
		s.False(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error value", func() {
		c := ParseByLayouts("xxx", []string{DateTimeLayout})
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("empty timezone", func() {
		c := ParseByLayouts("2020-08-05 13:14:15", []string{DateLayout}, "")
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error timezone", func() {
		c := ParseByLayouts("2020-08-05 13:14:15", []string{DateLayout}, "xxx")
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("empty layouts", func() {
		c := ParseByLayouts("2020-08-05 13:14:15", []string{})
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("without timezone", func() {
		c := ParseByLayouts("2020|8|5 1|2|3", []string{"2006|01|02 15|04|05", "2006|1|2 3|4|5"})
		s.Equal("2020-08-05 01:02:03 +0000 UTC", c.ToString())
		s.Equal("2006|1|2 3|4|5", c.CurrentLayout())
	})

	s.Run("with timezone", func() {
		c := ParseByLayouts("2020|8|5 1|2|3", []string{"2006|01|02 15|04|05", "2006|1|2 3|4|5"}, PRC)
		s.Equal("2020-08-05 01:02:03 +0800 CST", c.ToString())
		s.Equal("2006|1|2 3|4|5", c.CurrentLayout())
	})
}

func (s *ParserSuite) TestParseByFormat() {
	s.Run("empty value", func() {
		c := ParseByFormat("", DateFormat)
		s.False(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error value", func() {
		c := ParseByFormat("xxx", DateFormat)
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("empty format", func() {
		c := ParseByFormat("2020-08-05", "")
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error format", func() {
		c := ParseByFormat("2020-08-05", "xxx")
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("empty timezone", func() {
		c := ParseByFormat("2020-08-05", DateFormat, "")
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error timezone", func() {
		c := ParseByFormat("2020-08-05", DateFormat, "xxx")
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error timestamp", func() {
		s.Error(ParseByFormat("2020-08-05", TimestampFormat).Error)
		s.Error(ParseByFormat("2020-08-05", TimestampMilliFormat).Error)
		s.Error(ParseByFormat("2020-08-05", TimestampMicroFormat).Error)
		s.Error(ParseByFormat("2020-08-05", TimestampNanoFormat).Error)
	})

	s.Run("without timezone", func() {
		s.Equal("2020-08-05 00:00:00 +0000 UTC", ParseByFormat("2020-08-05", DateFormat).ToString())
		s.Equal("0000-01-01 13:14:15 +0000 UTC", ParseByFormat("13:14:15", TimeFormat).ToString())
		s.Equal("2020-08-05 13:14:15 +0000 UTC", ParseByFormat("2020-08-05 13:14:15", DateTimeFormat).ToString())

		s.Equal("2020-08-05 13:14:15 +0000 UTC", ParseByFormat("2020|08|05 13:14:15", "Y|m|d H:i:s").ToString())
		s.Equal("2020-08-05 13:14:15 +0000 UTC", ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToString())
		s.Equal("2020-08-05 13:14:15 +0000 UTC", ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToString())
	})

	s.Run("with timezone", func() {
		s.Equal("2020-08-05 00:00:00 +0800 CST", ParseByFormat("2020-08-05", DateFormat, PRC).ToString())
		s.Equal("0000-01-01 13:14:15 +0805 LMT", ParseByFormat("13:14:15", TimeFormat, PRC).ToString())
		s.Equal("2020-08-05 13:14:15 +0800 CST", ParseByFormat("2020-08-05 13:14:15", DateTimeFormat, PRC).ToString())

		s.Equal("2020-08-05 13:14:15 +0800 CST", ParseByFormat("2020|08|05 13:14:15", "Y|m|d H:i:s", PRC).ToString())
		s.Equal("2020-08-05 13:14:15 +0800 CST", ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s", PRC).ToString())
		s.Equal("2020-08-05 13:14:15 +0800 CST", ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒", PRC).ToString())
	})
}

func (s *ParserSuite) TestParseByFormats() {
	s.Run("empty value", func() {
		c := ParseByFormats("", []string{DateTimeLayout})
		s.False(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error value", func() {
		c := ParseByFormats("xxx", []string{DateTimeLayout})
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("empty timezone", func() {
		c := ParseByFormats("2020-08-05 13:14:15", []string{DateFormat}, "")
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("error timezone", func() {
		c := ParseByFormats("2020-08-05 13:14:15", []string{DateFormat, DateTimeLayout}, "xxx")
		s.True(c.HasError())
		s.Empty(c.String())
	})

	s.Run("empty formats", func() {
		c := ParseByFormats("2020-08-05 13:14:15", []string{})
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("without timezone", func() {
		c := ParseByFormats("2020|8|5 01|02|03", []string{"Y|m|d H|i|s", "Y|n|j h|i|s"})
		s.Equal("2020-08-05 01:02:03 +0000 UTC", c.ToString())
		s.Equal("2006|1|2 03|04|05", c.CurrentLayout())
	})

	s.Run("with timezone", func() {
		c := ParseByFormats("2020|8|5 01|02|03", []string{"Y|m|d H|i|s", "Y|n|j h|i|s"}, PRC)
		s.Equal("2020-08-05 01:02:03 +0800 CST", c.ToString())
		s.Equal("2006|1|2 03|04|05", c.CurrentLayout())
	})
}

func (s *ParserSuite) TestFormat2Layout() {
	c := Parse("2020-08-05 13:14:15.999999999")

	s.Run("without timezone", func() {
		s.Equal(c.Layout(AtomLayout), c.Format(AtomFormat))
		s.Equal(c.Layout(ANSICLayout), c.Format(ANSICFormat))
		s.Equal(c.Layout(CookieLayout), c.Format(CookieFormat))
		s.Equal(c.Layout(KitchenLayout), c.Format(KitchenFormat))
		s.Equal(c.Layout(RssLayout), c.Format(RssFormat))
		s.Equal(c.Layout(RubyDateLayout), c.Format(RubyDateFormat))
		s.Equal(c.Layout(UnixDateLayout), c.Format(UnixDateFormat))

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
