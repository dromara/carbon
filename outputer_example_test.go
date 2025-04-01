package carbon_test

import (
	"fmt"

	"github.com/dromara/carbon/v2"
)

func ExampleCarbon_String() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").String())
	fmt.Println(carbon.Parse("2020-08-05 13:14:15", carbon.PRC).String())

	// Output:
	// 2020-08-05 13:14:15
	// 2020-08-05 13:14:15
}

func ExampleCarbon_GoString() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").GoString())
	fmt.Println(carbon.Parse("2020-08-05 13:14:15", carbon.PRC).GoString())

	// Output:
	// time.Date(2020, time.August, 5, 13, 14, 15, 0, time.UTC)
	// time.Date(2020, time.August, 5, 13, 14, 15, 0, time.Location("PRC"))
}

func ExampleCarbon_ToString() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").ToString())
	fmt.Println(carbon.Parse("2020-08-05 13:14:15", carbon.PRC).ToString())

	// Output:
	// 2020-08-05 13:14:15 +0000 UTC
	// 2020-08-05 13:14:15 +0800 CST
}

func ExampleCarbon_ToMonthString() {
	fmt.Println(carbon.Parse("2020-01-05").ToMonthString())
	fmt.Println(carbon.Parse("2020-02-05").ToMonthString())
	fmt.Println(carbon.Parse("2020-03-05").ToMonthString())
	fmt.Println(carbon.Parse("2020-04-05").ToMonthString())
	fmt.Println(carbon.Parse("2020-05-05").ToMonthString())
	fmt.Println(carbon.Parse("2020-06-05").ToMonthString())
	fmt.Println(carbon.Parse("2020-07-05").ToMonthString())
	fmt.Println(carbon.Parse("2020-08-05").ToMonthString())
	fmt.Println(carbon.Parse("2020-09-05").ToMonthString())
	fmt.Println(carbon.Parse("2020-10-05").ToMonthString())
	fmt.Println(carbon.Parse("2020-11-05").ToMonthString())
	fmt.Println(carbon.Parse("2020-12-05").ToMonthString())

	// Output:
	// January
	// February
	// March
	// April
	// May
	// June
	// July
	// August
	// September
	// October
	// November
	// December
}

func ExampleCarbon_ToShortMonthString() {
	fmt.Println(carbon.Parse("2020-01-05").ToShortMonthString())
	fmt.Println(carbon.Parse("2020-02-05").ToShortMonthString())
	fmt.Println(carbon.Parse("2020-03-05").ToShortMonthString())
	fmt.Println(carbon.Parse("2020-04-05").ToShortMonthString())
	fmt.Println(carbon.Parse("2020-05-05").ToShortMonthString())
	fmt.Println(carbon.Parse("2020-06-05").ToShortMonthString())
	fmt.Println(carbon.Parse("2020-07-05").ToShortMonthString())
	fmt.Println(carbon.Parse("2020-08-05").ToShortMonthString())
	fmt.Println(carbon.Parse("2020-09-05").ToShortMonthString())
	fmt.Println(carbon.Parse("2020-10-05").ToShortMonthString())
	fmt.Println(carbon.Parse("2020-11-05").ToShortMonthString())
	fmt.Println(carbon.Parse("2020-12-05").ToShortMonthString())

	// Output:
	// Jan
	// Feb
	// Mar
	// Apr
	// May
	// Jun
	// Jul
	// Aug
	// Sep
	// Oct
	// Nov
	// Dec
}

func ExampleCarbon_ToWeekString() {
	fmt.Println(carbon.Parse("2020-08-01").ToWeekString())
	fmt.Println(carbon.Parse("2020-08-02").ToWeekString())
	fmt.Println(carbon.Parse("2020-08-03").ToWeekString())
	fmt.Println(carbon.Parse("2020-08-04").ToWeekString())
	fmt.Println(carbon.Parse("2020-08-05").ToWeekString())
	fmt.Println(carbon.Parse("2020-08-06").ToWeekString())
	fmt.Println(carbon.Parse("2020-08-07").ToWeekString())

	// Output:
	// Saturday
	// Sunday
	// Monday
	// Tuesday
	// Wednesday
	// Thursday
	// Friday
}

func ExampleCarbon_ToShortWeekString() {
	fmt.Println(carbon.Parse("2020-08-01").ToShortWeekString())
	fmt.Println(carbon.Parse("2020-08-02").ToShortWeekString())
	fmt.Println(carbon.Parse("2020-08-03").ToShortWeekString())
	fmt.Println(carbon.Parse("2020-08-04").ToShortWeekString())
	fmt.Println(carbon.Parse("2020-08-05").ToShortWeekString())
	fmt.Println(carbon.Parse("2020-08-06").ToShortWeekString())
	fmt.Println(carbon.Parse("2020-08-07").ToShortWeekString())

	// Output:
	// Sat
	// Sun
	// Mon
	// Tue
	// Wed
	// Thu
	// Fri
}

func ExampleCarbon_ToDayDateTimeString() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToDayDateTimeString())

	// Output:
	// Wed, Aug 5, 2020 1:14 PM
	// Wed, Aug 5, 2020 12:00 AM
}

func ExampleCarbon_ToDateTimeString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToDateTimeString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToDateTimeString())

	// Output:
	// 2020-08-05 13:14:15
	// 2020-08-05 00:00:00
}

func ExampleCarbon_ToDateTimeMilliString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToDateTimeMilliString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToDateTimeMilliString())

	// Output:
	// 2020-08-05 13:14:15.999
	// 2020-08-05 00:00:00
}

func ExampleCarbon_ToDateTimeMicroString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToDateTimeMicroString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToDateTimeMicroString())

	// Output:
	// 2020-08-05 13:14:15.999999
	// 2020-08-05 00:00:00
}

func ExampleCarbon_ToDateTimeNanoString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToDateTimeNanoString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToDateTimeNanoString())

	// Output:
	// 2020-08-05 13:14:15.999999999
	// 2020-08-05 00:00:00
}

func ExampleCarbon_ToShortDateTimeString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateTimeString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToShortDateTimeString())

	// Output:
	// 20200805131415
	// 20200805000000
}

func ExampleCarbon_ToShortDateTimeMilliString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateTimeMilliString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToShortDateTimeMilliString())

	// Output:
	// 20200805131415.999
	// 20200805000000
}

func ExampleCarbon_ToShortDateTimeMicroString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateTimeMicroString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToShortDateTimeMicroString())

	// Output:
	// 20200805131415.999999
	// 20200805000000
}

func ExampleCarbon_ToShortDateTimeNanoString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateTimeNanoString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToShortDateTimeNanoString())

	// Output:
	// 20200805131415.999999999
	// 20200805000000
}

func ExampleCarbon_ToDateString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToDateString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToDateString())

	// Output:
	// 2020-08-05
	// 2020-08-05
}

func ExampleCarbon_ToDateMilliString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToDateMilliString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToDateMilliString())

	// Output:
	// 2020-08-05.999
	// 2020-08-05
}

func ExampleCarbon_ToDateMicroString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToDateMicroString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToDateMicroString())

	// Output:
	// 2020-08-05.999999
	// 2020-08-05
}

func ExampleCarbon_ToDateNanoString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToDateNanoString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToDateNanoString())

	// Output:
	// 2020-08-05.999999999
	// 2020-08-05
}

func ExampleCarbon_ToShortDateString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToShortDateString())

	// Output:
	// 20200805
	// 20200805
}

func ExampleCarbon_ToShortDateMilliString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateMilliString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToShortDateMilliString())

	// Output:
	// 20200805.999
	// 20200805
}

func ExampleCarbon_ToShortDateMicroString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateMicroString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToShortDateMicroString())

	// Output:
	// 20200805.999999
	// 20200805
}

func ExampleCarbon_ToShortDateNanoString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToShortDateNanoString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToShortDateNanoString())

	// Output:
	// 20200805.999999999
	// 20200805
}

func ExampleCarbon_ToTimeString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToTimeString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToTimeString())

	// Output:
	// 13:14:15
	// 00:00:00
}

func ExampleCarbon_ToTimeMilliString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToTimeMilliString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToTimeMilliString())

	// Output:
	// 13:14:15.999
	// 00:00:00
}

func ExampleCarbon_ToTimeMicroString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToTimeMicroString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToTimeMicroString())

	// Output:
	// 13:14:15.999999
	// 00:00:00
}

func ExampleCarbon_ToTimeNanoString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToTimeNanoString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToTimeNanoString())

	// Output:
	// 13:14:15.999999999
	// 00:00:00
}

func ExampleCarbon_ToShortTimeString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToShortTimeString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToShortTimeString())

	// Output:
	// 131415
	// 000000
}

func ExampleCarbon_ToShortTimeMilliString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToShortTimeMilliString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToShortTimeMilliString())

	// Output:
	// 131415.999
	// 000000
}

func ExampleCarbon_ToShortTimeMicroString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToShortTimeMicroString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToShortTimeMicroString())

	// Output:
	// 131415.999999
	// 000000
}

func ExampleCarbon_ToShortTimeNanoString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToShortTimeNanoString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToShortTimeNanoString())

	// Output:
	// 131415.999999999
	// 000000
}

func ExampleCarbon_ToAtomString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToAtomString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToAtomString())

	// Output:
	// 2020-08-05T13:14:15Z
	// 2020-08-05T00:00:00+08:00
}

func ExampleCarbon_ToAnsicString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToAnsicString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToAnsicString())

	// Output:
	// Wed Aug  5 13:14:15 2020
	// Wed Aug  5 00:00:00 2020
}

func ExampleCarbon_ToCookieString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToCookieString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToCookieString())

	// Output:
	// Wednesday, 05-Aug-2020 13:14:15 UTC
	// Wednesday, 05-Aug-2020 00:00:00 CST
}

func ExampleCarbon_ToRssString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToRssString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToRssString())

	// Output:
	// Wed, 05 Aug 2020 13:14:15 +0000
	// Wed, 05 Aug 2020 00:00:00 +0800
}

func ExampleCarbon_ToW3cString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToW3cString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToW3cString())

	// Output:
	// 2020-08-05T13:14:15Z
	// 2020-08-05T00:00:00+08:00
}

func ExampleCarbon_ToUnixDateString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToUnixDateString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToUnixDateString())

	// Output:
	// Wed Aug  5 13:14:15 UTC 2020
	// Wed Aug  5 00:00:00 CST 2020
}

func ExampleCarbon_ToRubyDateString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToRubyDateString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToRubyDateString())

	// Output:
	// Wed Aug 05 13:14:15 +0000 2020
	// Wed Aug 05 00:00:00 +0800 2020
}

func ExampleCarbon_ToKitchenString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToKitchenString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToKitchenString())

	// Output:
	// 1:14PM
	// 12:00AM
}

func ExampleCarbon_ToIso8601String() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601String())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToIso8601String())

	// Output:
	// 2020-08-05T13:14:15+00:00
	// 2020-08-05T00:00:00+08:00
}

func ExampleCarbon_ToIso8601MilliString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601MilliString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToIso8601MilliString())

	// Output:
	// 2020-08-05T13:14:15.999+00:00
	// 2020-08-05T00:00:00+08:00
}

func ExampleCarbon_ToIso8601MicroString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601MicroString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToIso8601MicroString())

	// Output:
	// 2020-08-05T13:14:15.999999+00:00
	// 2020-08-05T00:00:00+08:00
}

func ExampleCarbon_ToIso8601NanoString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601NanoString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToIso8601NanoString())

	// Output:
	// 2020-08-05T13:14:15.999999999+00:00
	// 2020-08-05T00:00:00+08:00
}

func ExampleCarbon_ToIso8601ZuluString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601ZuluString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToIso8601ZuluString())

	// Output:
	// 2020-08-05T13:14:15Z
	// 2020-08-05T00:00:00Z
}

func ExampleCarbon_ToIso8601ZuluMilliString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601ZuluMilliString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToIso8601ZuluMilliString())

	// Output:
	// 2020-08-05T13:14:15.999Z
	// 2020-08-05T00:00:00Z
}

func ExampleCarbon_ToIso8601ZuluMicroString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601ZuluMicroString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToIso8601ZuluMicroString())

	// Output:
	// 2020-08-05T13:14:15.999999Z
	// 2020-08-05T00:00:00Z
}

func ExampleCarbon_ToIso8601ZuluNanoString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToIso8601ZuluNanoString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToIso8601ZuluNanoString())

	// Output:
	// 2020-08-05T13:14:15.999999999Z
	// 2020-08-05T00:00:00Z
}

func ExampleCarbon_ToRfc822String() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc822String())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToRfc822String())

	// Output:
	// 05 Aug 20 13:14 UTC
	// 05 Aug 20 00:00 CST
}

func ExampleCarbon_ToRfc822zString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc822zString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToRfc822zString())

	// Output:
	// 05 Aug 20 13:14 +0000
	// 05 Aug 20 00:00 +0800
}

func ExampleCarbon_ToRfc850String() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc850String())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToRfc850String())

	// Output:
	// Wednesday, 05-Aug-20 13:14:15 UTC
	// Wednesday, 05-Aug-20 00:00:00 CST
}

func ExampleCarbon_ToRfc1036String() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc1036String())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToRfc1036String())

	// Output:
	// Wed, 05 Aug 20 13:14:15 +0000
	// Wed, 05 Aug 20 00:00:00 +0800
}

func ExampleCarbon_ToRfc1123String() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc1123String())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToRfc1123String())

	// Output:
	// Wed, 05 Aug 2020 13:14:15 UTC
	// Wed, 05 Aug 2020 00:00:00 CST
}

func ExampleCarbon_ToRfc1123zString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc1123zString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToRfc1123zString())

	// Output:
	// Wed, 05 Aug 2020 13:14:15 +0000
	// Wed, 05 Aug 2020 00:00:00 +0800
}

func ExampleCarbon_ToRfc2822String() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc2822String())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToRfc2822String())

	// Output:
	// Wed, 05 Aug 2020 13:14:15 +0000
	// Wed, 05 Aug 2020 00:00:00 +0800
}

func ExampleCarbon_ToRfc3339String() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc3339String())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToRfc3339String())

	// Output:
	// 2020-08-05T13:14:15Z
	// 2020-08-05T00:00:00+08:00
}

func ExampleCarbon_ToRfc3339MilliString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc3339MilliString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToRfc3339MilliString())

	// Output:
	// 2020-08-05T13:14:15.999Z
	// 2020-08-05T00:00:00+08:00
}

func ExampleCarbon_ToRfc3339MicroString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc3339MicroString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToRfc3339MicroString())

	// Output:
	// 2020-08-05T13:14:15.999999Z
	// 2020-08-05T00:00:00+08:00
}

func ExampleCarbon_ToRfc3339NanoString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc3339NanoString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToRfc3339NanoString())

	// Output:
	// 2020-08-05T13:14:15.999999999Z
	// 2020-08-05T00:00:00+08:00
}

func ExampleCarbon_ToRfc7231String() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToRfc7231String())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToRfc7231String())

	// Output:
	// Wed, 05 Aug 2020 13:14:15 UTC
	// Wed, 05 Aug 2020 00:00:00 CST
}

func ExampleCarbon_ToFormattedDateString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToFormattedDateString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToFormattedDateString())

	// Output:
	// Aug 5, 2020
	// Aug 5, 2020
}

func ExampleCarbon_ToFormattedDayDateString() {
	fmt.Println(carbon.Parse("2020-08-05T13:14:15.999999999+00:00").ToFormattedDayDateString())
	fmt.Println(carbon.Parse("2020-08-05", carbon.PRC).ToFormattedDayDateString())

	// Output:
	// Wed, Aug 5, 2020
	// Wed, Aug 5, 2020
}

func ExampleCarbon_Layout() {
	fmt.Printf("2006-01-02 15:04:05 layout: %s\n", carbon.Parse("2020-08-05 13:14:15").Layout(carbon.DateTimeLayout))
	fmt.Printf("2006-01-02 layout: %s\n", carbon.Parse("2020-08-05T13:14:15.999999999+00:00").Layout(carbon.DateLayout))
	fmt.Printf("2006-01-02 15:04:05 with PRC timezone layout: %s\n", carbon.Parse("2020-08-05", carbon.PRC).Layout(carbon.DateTimeLayout))
	fmt.Printf("2006-01-02 15:04:05 with PRC timezone layout: %s\n", carbon.Parse("2020-08-05").Layout(carbon.DateTimeLayout, carbon.PRC))
	fmt.Printf("2006年01月02日 layout: %s\n", carbon.Parse("2020-08-05 13:14:15").Layout("2006年01月02日"))
	fmt.Printf("Mon, 02 Jan 2006 15:04:05 GMT layout: %s\n", carbon.Parse("2020-08-05 13:14:15").Layout("Mon, 02 Jan 2006 15:04:05 GMT"))

	// Output:
	// 2006-01-02 15:04:05 layout: 2020-08-05 13:14:15
	// 2006-01-02 layout: 2020-08-05
	// 2006-01-02 15:04:05 with PRC timezone layout: 2020-08-05 00:00:00
	// 2006-01-02 15:04:05 with PRC timezone layout: 2020-08-05 08:00:00
	// 2006年01月02日 layout: 2020年08月05日
	// Mon, 02 Jan 2006 15:04:05 GMT layout: Wed, 05 Aug 2020 13:14:15 GMT
}

func ExampleCarbon_Format() {
	fmt.Printf("Y-m-d H:i:s format: %s\n", carbon.Parse("2020-08-05 13:14:15").Format(carbon.DateTimeFormat))
	fmt.Printf("Y-m-d format: %s\n", carbon.Parse("2020-08-05T13:14:15.999999999+00:00").Format(carbon.DateFormat))
	fmt.Printf("Y-m-d H:i:s with PRC timezone format: %s\n", carbon.Parse("2020-08-05", carbon.PRC).Format(carbon.DateTimeFormat, carbon.PRC))
	fmt.Printf("Y年m月d日 format: %s\n", carbon.Parse("2020-08-05 13:14:15").Format("Y年m月d日"))
	fmt.Printf("D format: %s\n", carbon.Parse("2020-08-05 01:14:15").Format("D"))
	fmt.Printf("l format: %s\n", carbon.Parse("2020-08-05 01:14:15").Format("l"))
	fmt.Printf("F format: %s\n", carbon.Parse("2020-08-05 01:14:15").Format("F"))
	fmt.Printf("M format: %s\n", carbon.Parse("2020-08-05 01:14:15").Format("M"))
	fmt.Printf("j format: %s\n", carbon.Parse("2020-08-05 01:14:15").Format("j"))
	fmt.Printf("W format: %s\n", carbon.Parse("2020-08-05 01:14:15").Format("W"))
	fmt.Printf("F format: %s\n", carbon.Parse("2020-08-05 01:14:15").Format("F"))
	fmt.Printf("N format: %s\n", carbon.Parse("2020-08-05 01:14:15").Format("N"))
	fmt.Printf("L format: %s\n", carbon.Parse("2020-08-05 01:14:15").Format("L"))
	fmt.Printf("L format: %s\n", carbon.Parse("2021-08-05 13:14:15").Format("L"))
	fmt.Printf("G format: %s\n", carbon.Parse("2020-08-05 13:14:15").Format("G"))
	fmt.Printf("S format: %s\n", carbon.Parse("2020-08-05 13:14:15").Format("S"))
	fmt.Printf("U format: %s\n", carbon.Parse("2020-08-05 13:14:15").Format("U"))
	fmt.Printf("V format: %s\n", carbon.Parse("2020-08-05 13:14:15").Format("V"))
	fmt.Printf("X format: %s\n", carbon.Parse("2020-08-05 13:14:15").Format("X"))
	fmt.Printf("u format: %s\n", carbon.Parse("2020-08-05 13:14:15.999999999").Format("u"))
	fmt.Printf("v format: %s\n", carbon.Parse("2020-08-05 13:14:15.999999999").Format("v"))
	fmt.Printf("x format: %s\n", carbon.Parse("2020-08-05 13:14:15.999999999").Format("x"))

	fmt.Printf("w format: %s\n", carbon.Parse("2020-08-05 13:14:15.999999999").Format("w"))
	fmt.Printf("t format: %s\n", carbon.Parse("2020-08-05 13:14:15.999999999").Format("t"))
	fmt.Printf("z format: %s\n", carbon.Parse("2020-08-05 13:14:15.999999999", carbon.PRC).Format("z"))
	fmt.Printf("o format: %s\n", carbon.Parse("2020-08-05 13:14:15.999999999", carbon.PRC).Format("o"))
	fmt.Printf("q format: %s\n", carbon.Parse("2020-08-05 13:14:15.999999999").Format("q"))
	fmt.Printf("c format: %s\n", carbon.Parse("2020-08-05 13:14:15.999999999").Format("c"))
	fmt.Printf("R format: %s\n", carbon.Parse("2020-08-05 13:14:15.999999999", carbon.UTC).Format("R"))
	fmt.Printf("R format: %s\n", carbon.Parse("2020-08-05 13:14:15.999999999", carbon.PRC).Format("R"))
	fmt.Printf("Q format: %s\n", carbon.Parse("2020-08-05 13:14:15.999999999", carbon.UTC).Format("Q"))
	fmt.Printf("Q format: %s\n", carbon.Parse("2020-08-05 13:14:15.999999999", carbon.PRC).Format("Q"))
	fmt.Printf("Z format: %s\n", carbon.Parse("2020-08-05 13:14:15.999999999", carbon.PRC).Format("Z"))
	fmt.Printf("jK format: %s\n", carbon.Parse("2020-08-05 13:14:15").Format("jK"))
	fmt.Printf("jK format: %s\n", carbon.Parse("2020-08-22 13:14:15").Format("jK"))
	fmt.Printf("jK format: %s\n", carbon.Parse("2020-08-23 13:14:15").Format("jK"))
	fmt.Printf("jK format: %s\n", carbon.Parse("2020-08-31 13:14:15").Format("jK"))
	fmt.Printf("I\\t \\i\\s Y-m-d H:i:s format: %s\n", carbon.Parse("2020-08-31 13:14:15").Format("I\\t \\i\\s Y-m-d H:i:s"))
	fmt.Printf("上次打卡时间:Y-m-d H:i:s，请每日按时打卡 format: %s\n", carbon.Parse("2020-08-31 13:14:15").Format("上次打卡时间:Y-m-d H:i:s，请每日按时打卡"))

	// Output:
	// Y-m-d H:i:s format: 2020-08-05 13:14:15
	// Y-m-d format: 2020-08-05
	// Y-m-d H:i:s with PRC timezone format: 2020-08-05 00:00:00
	// Y年m月d日 format: 2020年08月05日
	// D format: Wed
	// l format: Wednesday
	// F format: August
	// M format: Aug
	// j format: 5
	// W format: 32
	// F format: August
	// N format: 03
	// L format: 1
	// L format: 0
	// G format: 13
	// S format: 1596633255
	// U format: 1596633255000
	// V format: 1596633255000000
	// X format: 1596633255000000000
	// u format: 999
	// v format: 999999
	// x format: 999999999
	// w format: 2
	// t format: 31
	// z format: PRC
	// o format: 28800
	// q format: 3
	// c format: 21
	// R format: Z
	// R format: +08:00
	// Q format: Z
	// Q format: +0800
	// Z format: CST
	// jK format: 5th
	// jK format: 22nd
	// jK format: 23rd
	// jK format: 31st
	// I\t \i\s Y-m-d H:i:s format: It is 2020-08-31 13:14:15
	// 上次打卡时间:Y-m-d H:i:s，请每日按时打卡 format: 上次打卡时间:2020-08-31 13:14:15，请每日按时打卡
}
