package carbon_test

import (
	"fmt"

	"github.com/dromara/carbon/v2"
)

func ExampleCarbon_HasError() {
	fmt.Println(carbon.NewCarbon().HasError())
	fmt.Println(carbon.Parse("").HasError())
	fmt.Println(carbon.Now().HasError())
	fmt.Println(carbon.Parse("xxx").HasError())

	// Output:
	// false
	// false
	// false
	// true
}

func ExampleCarbon_IsNil() {
	fmt.Println(carbon.NewCarbon().IsNil())
	fmt.Println(carbon.Parse("xxx").IsNil())
	fmt.Println(carbon.Now().IsNil())
	fmt.Println(carbon.Parse("").IsNil())

	c := carbon.NewCarbon()
	c = nil
	fmt.Println(c.IsNil())

	// Output:
	// false
	// false
	// false
	// false
	// true
}

func ExampleCarbon_IsEmpty() {
	var c *carbon.Carbon
	c = nil
	fmt.Println(c.IsEmpty())

	fmt.Println(carbon.NewCarbon().IsEmpty())
	fmt.Println(carbon.Parse("xxx").IsEmpty())
	fmt.Println(carbon.Now().IsEmpty())
	fmt.Println(carbon.Parse("").IsEmpty())

	// Output:
	// false
	// false
	// false
	// false
	// true
}

func ExampleCarbon_IsZero() {
	fmt.Println(carbon.NewCarbon().IsZero())
	fmt.Println(carbon.ZeroValue().IsZero())
	fmt.Println(carbon.CreateFromDateTimeNano(1, 1, 1, 0, 0, 0, 0, carbon.UTC).IsZero())
	fmt.Println(carbon.CreateFromTimestamp(0).IsZero())

	fmt.Println(carbon.Parse("").IsZero())
	fmt.Println(carbon.Parse("xxx").IsZero())
	fmt.Println(carbon.Now().IsZero())

	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
	// false
}

func ExampleCarbon_IsEpoch() {
	fmt.Println(carbon.NewCarbon().IsEpoch())
	fmt.Println(carbon.EpochValue().IsEpoch())
	fmt.Println(carbon.CreateFromDateTimeNano(1970, 1, 1, 0, 0, 0, 0, carbon.UTC).IsEpoch())
	fmt.Println(carbon.CreateFromTimestamp(0).IsEpoch())

	fmt.Println(carbon.Parse("").IsEpoch())
	fmt.Println(carbon.Parse("xxx").IsEpoch())
	fmt.Println(carbon.Now().IsEpoch())

	// Output:
	// false
	// true
	// true
	// true
	// false
	// false
	// false
}
func ExampleCarbon_IsValid() {
	fmt.Println(carbon.NewCarbon().IsValid())
	fmt.Println(carbon.Now().IsValid())
	fmt.Println(carbon.Parse("").IsValid())
	fmt.Println(carbon.Parse("xxx").IsValid())

	// Output:
	// true
	// true
	// false
	// false
}

func ExampleCarbon_IsInvalid() {
	fmt.Println(carbon.NewCarbon().IsInvalid())
	fmt.Println(carbon.Parse("").IsInvalid())
	fmt.Println(carbon.Parse("xxx").IsInvalid())
	fmt.Println(carbon.Now().IsInvalid())

	// Output:
	// false
	// true
	// true
	// false
}

func ExampleCarbon_IsDST() {
	tzWithDST, tzWithoutDST := "Australia/Sydney", "Australia/Brisbane"
	fmt.Println(carbon.Parse("2009-01-01", tzWithDST).IsDST())
	fmt.Println(carbon.Parse("2009-01-01", tzWithoutDST).IsDST())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsAM() {
	fmt.Println(carbon.Parse("2020-08-08 00:00:00").IsAM())
	fmt.Println(carbon.Parse("2020-08-08 12:00:00").IsAM())
	fmt.Println(carbon.Parse("2020-08-08 23:59:59").IsAM())

	// Output:
	// true
	// false
	// false
}

func ExampleCarbon_IsPM() {
	fmt.Println(carbon.Parse("2020-08-08 00:00:00").IsPM())
	fmt.Println(carbon.Parse("2020-08-08 12:00:00").IsPM())
	fmt.Println(carbon.Parse("2020-08-08 23:59:59").IsPM())

	// Output:
	// false
	// true
	// true
}

func ExampleCarbon_IsLeapYear() {
	fmt.Println(carbon.Parse("2015-01-01").IsLeapYear())
	fmt.Println(carbon.Parse("2016-01-01").IsLeapYear())

	// Output:
	// false
	// true
}

func ExampleCarbon_IsLongYear() {
	fmt.Println(carbon.Parse("2015-01-01").IsLongYear())
	fmt.Println(carbon.Parse("2016-01-01").IsLongYear())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsJanuary() {
	fmt.Println(carbon.Parse("2020-01-01").IsJanuary())
	fmt.Println(carbon.Parse("2020-02-01").IsJanuary())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsFebruary() {
	fmt.Println(carbon.Parse("2020-02-01").IsFebruary())
	fmt.Println(carbon.Parse("2020-03-01").IsFebruary())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsMarch() {
	fmt.Println(carbon.Parse("2020-03-01").IsMarch())
	fmt.Println(carbon.Parse("2020-04-01").IsMarch())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsApril() {
	fmt.Println(carbon.Parse("2020-04-01").IsApril())
	fmt.Println(carbon.Parse("2020-05-01").IsApril())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsMay() {
	fmt.Println(carbon.Parse("2020-05-01").IsMay())
	fmt.Println(carbon.Parse("2020-06-01").IsMay())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsJune() {
	fmt.Println(carbon.Parse("2020-06-01").IsJune())
	fmt.Println(carbon.Parse("2020-07-01").IsJune())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsJuly() {
	fmt.Println(carbon.Parse("2020-07-01").IsJuly())
	fmt.Println(carbon.Parse("2020-08-01").IsJuly())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsAugust() {
	fmt.Println(carbon.Parse("2020-08-01").IsAugust())
	fmt.Println(carbon.Parse("2020-09-01").IsAugust())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsSeptember() {
	fmt.Println(carbon.Parse("2020-09-01").IsSeptember())
	fmt.Println(carbon.Parse("2020-10-01").IsSeptember())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsOctober() {
	fmt.Println(carbon.Parse("2020-10-01").IsOctober())
	fmt.Println(carbon.Parse("2020-11-01").IsOctober())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsNovember() {
	fmt.Println(carbon.Parse("2020-11-01").IsNovember())
	fmt.Println(carbon.Parse("2020-12-01").IsNovember())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsDecember() {
	fmt.Println(carbon.Parse("2020-12-01").IsDecember())
	fmt.Println(carbon.Parse("2020-01-01").IsDecember())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsMonday() {
	fmt.Println(carbon.Parse("2025-03-03").IsMonday())
	fmt.Println(carbon.Parse("2025-03-04").IsMonday())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsTuesday() {
	fmt.Println(carbon.Parse("2025-03-04").IsTuesday())
	fmt.Println(carbon.Parse("2025-03-05").IsTuesday())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsWednesday() {
	fmt.Println(carbon.Parse("2025-03-05").IsWednesday())
	fmt.Println(carbon.Parse("2025-03-06").IsWednesday())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsThursday() {
	fmt.Println(carbon.Parse("2025-03-06").IsThursday())
	fmt.Println(carbon.Parse("2025-03-07").IsThursday())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsFriday() {
	fmt.Println(carbon.Parse("2025-03-07").IsFriday())
	fmt.Println(carbon.Parse("2025-03-08").IsFriday())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsSaturday() {
	fmt.Println(carbon.Parse("2025-03-08").IsSaturday())
	fmt.Println(carbon.Parse("2025-03-09").IsSaturday())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsSunday() {
	fmt.Println(carbon.Parse("2025-03-09").IsSunday())
	fmt.Println(carbon.Parse("2025-03-10").IsSunday())

	// Output:
	// true
	// false
}

func ExampleCarbon_IsWeekday() {
	fmt.Println(carbon.Parse("2025-03-01").IsWeekday())
	fmt.Println(carbon.Parse("2025-03-02").IsWeekday())
	fmt.Println(carbon.Parse("2025-03-03").IsWeekday())

	// Output:
	// false
	// false
	// true
}

func ExampleCarbon_IsWeekend() {
	fmt.Println(carbon.Parse("2025-03-01").IsWeekend())
	fmt.Println(carbon.Parse("2025-03-02").IsWeekend())
	fmt.Println(carbon.Parse("2025-03-03").IsWeekend())

	// Output:
	// true
	// true
	// false
}

func ExampleCarbon_IsNow() {
	fmt.Println(carbon.Yesterday().IsNow())
	fmt.Println(carbon.Now().IsNow())
	fmt.Println(carbon.Tomorrow().IsNow())
	fmt.Println(carbon.Parse("2025-03-01").IsNow())

	// Output:
	// false
	// true
	// false
	// false
}

func ExampleCarbon_IsFuture() {
	fmt.Println(carbon.Yesterday().IsFuture())
	fmt.Println(carbon.Now().IsFuture())
	fmt.Println(carbon.Tomorrow().IsFuture())
	fmt.Println(carbon.Parse("2025-03-01").IsFuture())

	// Output:
	// false
	// false
	// true
	// false
}

func ExampleCarbon_IsPast() {
	fmt.Println(carbon.Yesterday().IsPast())
	fmt.Println(carbon.Now().IsPast())
	fmt.Println(carbon.Tomorrow().IsPast())
	fmt.Println(carbon.Parse("2025-03-01").IsPast())

	// Output:
	// true
	// false
	// false
	// true
}

func ExampleCarbon_IsYesterday() {
	fmt.Println(carbon.Yesterday().IsYesterday())
	fmt.Println(carbon.Now().IsYesterday())
	fmt.Println(carbon.Tomorrow().IsYesterday())
	fmt.Println(carbon.Parse("2025-03-01").IsYesterday())

	// Output:
	// true
	// false
	// false
	// false
}

func ExampleCarbon_IsToday() {
	fmt.Println(carbon.Yesterday().IsToday())
	fmt.Println(carbon.Now().IsToday())
	fmt.Println(carbon.Tomorrow().IsToday())
	fmt.Println(carbon.Parse("2025-03-01").IsToday())

	// Output:
	// false
	// true
	// false
	// false
}

func ExampleCarbon_IsTomorrow() {
	fmt.Println(carbon.Yesterday().IsTomorrow())
	fmt.Println(carbon.Now().IsTomorrow())
	fmt.Println(carbon.Tomorrow().IsTomorrow())
	fmt.Println(carbon.Parse("2025-03-01").IsTomorrow())

	// Output:
	// false
	// false
	// true
	// false
}

func ExampleCarbon_IsSameCentury() {
	fmt.Println(carbon.Parse("2020-08-05").IsSameCentury(carbon.Parse("2010-01-01")))
	fmt.Println(carbon.Parse("2020-08-05").IsSameCentury(carbon.Parse("2030-12-31")))
	fmt.Println(carbon.Parse("2020-08-05").IsSameCentury(carbon.Parse("2100-08-05")))

	// Output:
	// true
	// true
	// false
}

func ExampleCarbon_IsSameDecade() {
	fmt.Println(carbon.Parse("2020-08-05").IsSameDecade(carbon.Parse("2020-01-01")))
	fmt.Println(carbon.Parse("2020-08-05").IsSameDecade(carbon.Parse("2020-12-31")))
	fmt.Println(carbon.Parse("2020-08-05").IsSameDecade(carbon.Parse("2010-08-05")))

	// Output:
	// true
	// true
	// false
}

func ExampleCarbon_IsSameYear() {
	fmt.Println(carbon.Parse("2020-08-05").IsSameYear(carbon.Parse("2010-08-05")))
	fmt.Println(carbon.Parse("2020-08-05").IsSameYear(carbon.Parse("2020-01-01")))
	fmt.Println(carbon.Parse("2020-08-05").IsSameYear(carbon.Parse("2020-12-31")))

	// Output:
	// false
	// true
	// true
}

func ExampleCarbon_IsSameQuarter() {
	fmt.Println(carbon.Parse("2020-08-05").IsSameQuarter(carbon.Parse("2020-08-06")))
	fmt.Println(carbon.Parse("2020-08-05").IsSameQuarter(carbon.Parse("2010-08-05")))
	fmt.Println(carbon.Parse("2020-08-05").IsSameQuarter(carbon.Parse("2010-01-05")))

	// Output:
	// true
	// false
	// false
}

func ExampleCarbon_IsSameMonth() {
	fmt.Println(carbon.Parse("2020-08-05").IsSameMonth(carbon.Parse("2020-08-01")))
	fmt.Println(carbon.Parse("2020-08-05").IsSameMonth(carbon.Parse("2010-08-05")))
	fmt.Println(carbon.Parse("2020-08-05").IsSameMonth(carbon.Parse("2020-09-05")))

	// Output:
	// true
	// false
	// false
}

func ExampleCarbon_IsSameDay() {
	fmt.Println(carbon.Parse("2020-08-05 00:00:00").IsSameDay(carbon.Parse("2020-08-05 23:59:59")))
	fmt.Println(carbon.Parse("2020-08-05 00:00:00").IsSameDay(carbon.Parse("2021-08-05 00:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 00:00:00").IsSameDay(carbon.Parse("2020-09-05 00:00:00")))

	// Output:
	// true
	// false
	// false
}

func ExampleCarbon_IsSameHour() {
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").IsSameHour(carbon.Parse("2020-08-05 22:59:59")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").IsSameHour(carbon.Parse("2021-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").IsSameHour(carbon.Parse("2020-09-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").IsSameHour(carbon.Parse("2020-08-06 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").IsSameHour(carbon.Parse("2020-08-05 23:00:00")))

	// Output:
	// true
	// false
	// false
	// false
	// false
}

func ExampleCarbon_IsSameMinute() {
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").IsSameMinute(carbon.Parse("2020-08-05 22:00:59")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").IsSameMinute(carbon.Parse("2021-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").IsSameMinute(carbon.Parse("2020-08-06 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").IsSameMinute(carbon.Parse("2020-08-05 23:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").IsSameMinute(carbon.Parse("2020-08-05 22:01:00")))

	// Output:
	// true
	// false
	// false
	// false
	// false
}

func ExampleCarbon_IsSameSecond() {
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").IsSameSecond(carbon.Parse("2020-08-05 22:00:00.999999")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").IsSameSecond(carbon.Parse("2021-08-05 22:00:00")))

	// Output:
	// true
	// false
}

func ExampleCarbon_IsBirthday() {
	defer carbon.ClearTestNow()
	carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15"))

	fmt.Println(carbon.Parse("1990-08-05").IsBirthday())
	fmt.Println(carbon.Parse("1990-08-06").IsBirthday())
	fmt.Println(carbon.Parse("2014-04-23").IsBirthday(carbon.Parse("1987-04-23")))
	fmt.Println(carbon.Parse("2014-09-26").IsBirthday(carbon.Parse("1987-04-23")))
	fmt.Println(carbon.Parse("2000-02-29").IsBirthday(carbon.Parse("2023-02-28")))

	// Output:
	// true
	// false
	// true
	// false
	// false
}

func ExampleCarbon_Compare() {
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Compare(">=", carbon.Parse("2020-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Compare("<", carbon.Parse("2020-08-05 22:00:00")))

	// Output:
	// true
	// false
}

func ExampleCarbon_Gt() {
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Gt(carbon.Parse("2020-08-05 21:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Gt(carbon.Parse("2020-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Gt(carbon.Parse("2020-08-05 23:00:00")))

	// Output:
	// true
	// false
	// false
}

func ExampleCarbon_Lt() {
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Lt(carbon.Parse("2020-08-05 21:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Lt(carbon.Parse("2020-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Lt(carbon.Parse("2020-08-05 23:00:00")))

	// Output:
	// false
	// false
	// true
}

func ExampleCarbon_Eq() {
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Eq(carbon.Parse("2020-08-05 21:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Eq(carbon.Parse("2020-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Eq(carbon.Parse("2020-08-05 23:00:00")))

	// Output:
	// false
	// true
	// false
}

func ExampleCarbon_Ne() {
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Ne(carbon.Parse("2020-08-05 21:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Ne(carbon.Parse("2020-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Ne(carbon.Parse("2020-08-05 23:00:00")))

	// Output:
	// true
	// false
	// true
}

func ExampleCarbon_Gte() {
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Gte(carbon.Parse("2020-08-05 21:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Gte(carbon.Parse("2020-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Gte(carbon.Parse("2020-08-05 23:00:00")))

	// Output:
	// true
	// true
	// false
}

func ExampleCarbon_Lte() {
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Lte(carbon.Parse("2020-08-05 21:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Lte(carbon.Parse("2020-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Lte(carbon.Parse("2020-08-05 23:00:00")))

	// Output:
	// false
	// true
	// true
}

func ExampleCarbon_Between() {
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Between(carbon.Parse("2020-08-05 21:00:00"), carbon.Parse("2020-08-05 23:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Between(carbon.Parse("2020-08-05 22:00:00"), carbon.Parse("2020-08-05 23:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Between(carbon.Parse("2020-08-05 21:00:00"), carbon.Parse("2020-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Between(carbon.Parse("2020-08-05 22:00:00"), carbon.Parse("2020-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").Between(carbon.Parse("2021-08-05 22:00:00"), carbon.Parse("2019-08-05 22:00:00")))

	// Output:
	// true
	// false
	// false
	// false
	// false
}

func ExampleCarbon_BetweenIncludedStart() {
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").BetweenIncludedStart(carbon.Parse("2020-08-05 22:00:00"), carbon.Parse("2020-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").BetweenIncludedStart(carbon.Parse("2020-08-05 21:00:00"), carbon.Parse("2020-08-05 23:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").BetweenIncludedStart(carbon.Parse("2020-08-05 22:00:00"), carbon.Parse("2020-08-05 23:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").BetweenIncludedStart(carbon.Parse("2020-08-05 21:00:00"), carbon.Parse("2020-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").BetweenIncludedStart(carbon.Parse("2022-08-05 22:00:00"), carbon.Parse("2021-08-05 22:00:00")))

	// Output:
	// false
	// true
	// true
	// false
	// false
}

func ExampleCarbon_BetweenIncludedEnd() {
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").BetweenIncludedEnd(carbon.Parse("2020-08-05 22:00:00"), carbon.Parse("2020-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").BetweenIncludedEnd(carbon.Parse("2020-08-05 21:00:00"), carbon.Parse("2020-08-05 23:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").BetweenIncludedEnd(carbon.Parse("2020-08-05 22:00:00"), carbon.Parse("2020-08-05 23:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").BetweenIncludedEnd(carbon.Parse("2020-08-05 21:00:00"), carbon.Parse("2020-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").BetweenIncludedEnd(carbon.Parse("2022-08-05 22:00:00"), carbon.Parse("2021-08-05 22:00:00")))

	// Output:
	// false
	// true
	// false
	// true
	// false
}

func ExampleCarbon_BetweenIncludedBoth() {
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").BetweenIncludedBoth(carbon.Parse("2020-08-05 22:00:00"), carbon.Parse("2020-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").BetweenIncludedBoth(carbon.Parse("2020-08-05 21:00:00"), carbon.Parse("2020-08-05 23:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").BetweenIncludedBoth(carbon.Parse("2020-08-05 22:00:00"), carbon.Parse("2020-08-05 23:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").BetweenIncludedBoth(carbon.Parse("2020-08-05 21:00:00"), carbon.Parse("2020-08-05 22:00:00")))
	fmt.Println(carbon.Parse("2020-08-05 22:00:00").BetweenIncludedBoth(carbon.Parse("2022-08-05 22:00:00"), carbon.Parse("2021-08-05 22:00:00")))

	// Output:
	// true
	// true
	// true
	// true
	// false
}
