package carbon_test

import (
	"fmt"

	"github.com/dromara/carbon/v2"
)

func ExampleNow() {
	defer carbon.ClearTestNow()

	carbon.SetTestNow(carbon.Parse("2020-08-05"))

	fmt.Println(carbon.Now().ToString())
	fmt.Println(carbon.Now(carbon.PRC).ToString())

	// Output:
	// 2020-08-05 00:00:00 +0000 UTC
	// 2020-08-05 08:00:00 +0800 CST
}

func ExampleTomorrow() {
	defer carbon.ClearTestNow()

	carbon.SetTestNow(carbon.Parse("2020-08-05"))

	fmt.Println(carbon.Tomorrow().ToString())
	fmt.Println(carbon.Tomorrow(carbon.PRC).ToString())

	// Output:
	// 2020-08-06 00:00:00 +0000 UTC
	// 2020-08-06 08:00:00 +0800 CST
}

func ExampleYesterday() {
	defer carbon.ClearTestNow()

	carbon.SetTestNow(carbon.Parse("2020-08-05"))

	fmt.Println(carbon.Yesterday().ToString())
	fmt.Println(carbon.Yesterday(carbon.PRC).ToString())

	// Output:
	// 2020-08-04 00:00:00 +0000 UTC
	// 2020-08-04 08:00:00 +0800 CST
}

func ExampleCarbon_AddDuration() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddDuration("10h").ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddDuration("10.5h").ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddDuration("10m").ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddDuration("10.5m").ToString())

	// Output:
	// 2020-01-01 23:14:15 +0000 UTC
	// 2020-01-01 23:44:15 +0000 UTC
	// 2020-01-01 13:24:15 +0000 UTC
	// 2020-01-01 13:24:45 +0000 UTC
}

func ExampleCarbon_SubDuration() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubDuration("10h").ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubDuration("10.5h").ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubDuration("10m").ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubDuration("10.5m").ToString())

	// Output:
	// 2020-01-01 03:14:15 +0000 UTC
	// 2020-01-01 02:44:15 +0000 UTC
	// 2020-01-01 13:04:15 +0000 UTC
	// 2020-01-01 13:03:45 +0000 UTC
}

func ExampleCarbon_AddCenturies() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddCenturies(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddCenturies(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddCenturies(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2120-01-01 13:14:15 +0000 UTC
	// 2220-01-01 13:14:15 +0000 UTC
}

func ExampleCarbon_AddCenturiesNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddCenturiesNoOverflow(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddCenturiesNoOverflow(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddCenturiesNoOverflow(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2120-01-01 13:14:15 +0000 UTC
	// 2220-01-01 13:14:15 +0000 UTC
}

func ExampleCarbon_AddCentury() {
	defer carbon.ClearTestNow()

	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddCentury().ToString())

	// Output:
	// 2120-01-01 13:14:15 +0000 UTC
}

func ExampleCarbon_AddCenturyNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddCenturyNoOverflow().ToString())

	// Output:
	// 2120-01-01 13:14:15 +0000 UTC
}

func ExampleCarbon_SubCenturies() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubCenturies(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubCenturies(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubCenturies(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 1920-01-01 13:14:15 +0000 UTC
	// 1820-01-01 13:14:15 +0000 UTC
}

func ExampleCarbon_SubCenturiesNoOverflow() {
	fmt.Println(carbon.Parse("2020-02-28 13:14:15").SubCenturiesNoOverflow(0).ToString())
	fmt.Println(carbon.Parse("2020-02-28 13:14:15").SubCenturiesNoOverflow(1).ToString())
	fmt.Println(carbon.Parse("2020-02-28 13:14:15").SubCenturiesNoOverflow(2).ToString())

	// Output:
	// 2020-02-28 13:14:15 +0000 UTC
	// 1920-02-28 13:14:15 +0000 UTC
	// 1820-02-28 13:14:15 +0000 UTC
}

func ExampleCarbon_SubCentury() {
	fmt.Println(carbon.Parse("2020-02-28 13:14:15").SubCentury().ToString())

	// Output:
	// 1920-02-28 13:14:15 +0000 UTC
}

func ExampleCarbon_SubCenturyNoOverflow() {
	fmt.Println(carbon.Parse("2020-02-28 13:14:15").SubCenturyNoOverflow().ToString())

	// Output:
	// 1920-02-28 13:14:15 +0000 UTC
}

func ExampleCarbon_AddDecades() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddDecades(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddDecades(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddDecades(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2030-01-01 13:14:15 +0000 UTC
	// 2040-01-01 13:14:15 +0000 UTC
}

func ExampleCarbon_AddDecadesNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddDecadesNoOverflow(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddDecadesNoOverflow(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddDecadesNoOverflow(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2030-01-01 13:14:15 +0000 UTC
	// 2040-01-01 13:14:15 +0000 UTC
}

func ExampleCarbon_AddDecade() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddDecade().ToString())

	// Output:
	// 2030-01-01 13:14:15 +0000 UTC
}

func ExampleCarbon_AddDecadeNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddDecadeNoOverflow().ToString())

	// Output:
	// 2030-01-01 13:14:15 +0000 UTC
}

func ExampleCarbon_SubDecades() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubDecades(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubDecades(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubDecades(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2010-01-01 13:14:15 +0000 UTC
	// 2000-01-01 13:14:15 +0000 UTC
}

func ExampleCarbon_SubDecadesNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubDecadesNoOverflow(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubDecadesNoOverflow(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubDecadesNoOverflow(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2010-01-01 13:14:15 +0000 UTC
	// 2000-01-01 13:14:15 +0000 UTC
}

func ExampleCarbon_SubDecade() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubDecade().ToString())

	// Output:
	// 2010-01-01 13:14:15 +0000 UTC
}

func ExampleCarbon_SubDecadeNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubDecadeNoOverflow().ToString())

	// Output:
	// 2010-01-01 13:14:15 +0000 UTC
}

func ExampleCarbon_AddYears() {
	fmt.Println(carbon.Parse("2020-01-01").AddYears(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddYears(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddYears(2).ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").AddYears(3).ToDateString())

	// Output:
	// 2020-01-01
	// 2021-01-01
	// 2022-01-01
	// 2023-03-01
}

func ExampleCarbon_AddYearsNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01").AddYearsNoOverflow(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddYearsNoOverflow(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddYearsNoOverflow(2).ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").AddYearsNoOverflow(3).ToDateString())

	// Output:
	// 2020-01-01
	// 2021-01-01
	// 2022-01-01
	// 2023-02-28
}

func ExampleCarbon_AddYear() {
	fmt.Println(carbon.Parse("2020-01-01").AddYear().ToDateString())
	fmt.Println(carbon.Parse("2020-02-28").AddYear().ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").AddYear().ToDateString())

	// Output:
	// 2021-01-01
	// 2021-02-28
	// 2021-03-01
}

func ExampleCarbon_AddYearNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01").AddYearNoOverflow().ToDateString())
	fmt.Println(carbon.Parse("2020-02-28").AddYearNoOverflow().ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").AddYearNoOverflow().ToDateString())

	// Output:
	// 2021-01-01
	// 2021-02-28
	// 2021-02-28
}

func ExampleCarbon_SubYears() {
	fmt.Println(carbon.Parse("2020-01-01").SubYears(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").SubYears(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").SubYears(2).ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").SubYears(3).ToDateString())

	// Output:
	// 2020-01-01
	// 2019-01-01
	// 2018-01-01
	// 2017-03-01
}

func ExampleCarbon_SubYearsNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01").SubYearsNoOverflow(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").SubYearsNoOverflow(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").SubYearsNoOverflow(2).ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").SubYearsNoOverflow(3).ToDateString())

	// Output:
	// 2020-01-01
	// 2019-01-01
	// 2018-01-01
	// 2017-02-28
}

func ExampleCarbon_SubYear() {
	fmt.Println(carbon.Parse("2020-01-01").SubYear().ToDateString())
	fmt.Println(carbon.Parse("2020-02-28").SubYear().ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").SubYear().ToDateString())

	// Output:
	// 2019-01-01
	// 2019-02-28
	// 2019-03-01
}

func ExampleCarbon_SubYearNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01").SubYearNoOverflow().ToDateString())
	fmt.Println(carbon.Parse("2020-02-28").SubYearNoOverflow().ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").SubYearNoOverflow().ToDateString())

	// Output:
	// 2019-01-01
	// 2019-02-28
	// 2019-02-28
}

func ExampleCarbon_AddQuarters() {
	fmt.Println(carbon.Parse("2020-01-01").AddQuarters(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddQuarters(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddQuarters(2).ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").AddQuarters(3).ToDateString())
	fmt.Println(carbon.Parse("2020-08-31").AddQuarters(2).ToDateString())

	// Output:
	// 2020-01-01
	// 2020-04-01
	// 2020-07-01
	// 2020-11-29
	// 2021-03-03
}

func ExampleCarbon_AddQuartersNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01").AddQuartersNoOverflow(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddQuartersNoOverflow(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddQuartersNoOverflow(2).ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").AddQuartersNoOverflow(3).ToDateString())
	fmt.Println(carbon.Parse("2020-08-31").AddQuartersNoOverflow(2).ToDateString())

	// Output:
	// 2020-01-01
	// 2020-04-01
	// 2020-07-01
	// 2020-11-29
	// 2021-02-28
}

func ExampleCarbon_AddQuarter() {
	fmt.Println(carbon.Parse("2020-01-01").AddQuarter().ToDateString())
	fmt.Println(carbon.Parse("2020-02-28").AddQuarter().ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").AddQuarter().ToDateString())
	fmt.Println(carbon.Parse("2020-11-30").AddQuarter().ToDateString())

	// Output:
	// 2020-04-01
	// 2020-05-28
	// 2020-05-29
	// 2021-03-02
}

func ExampleCarbon_AddQuarterNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01").AddQuarterNoOverflow().ToDateString())
	fmt.Println(carbon.Parse("2020-02-28").AddQuarterNoOverflow().ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").AddQuarterNoOverflow().ToDateString())
	fmt.Println(carbon.Parse("2020-11-30").AddQuarterNoOverflow().ToDateString())

	// Output:
	// 2020-04-01
	// 2020-05-28
	// 2020-05-29
	// 2021-02-28
}

func ExampleCarbon_SubQuarters() {
	fmt.Println(carbon.Parse("2020-01-01").SubQuarters(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").SubQuarters(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").SubQuarters(2).ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").SubQuarters(3).ToDateString())
	fmt.Println(carbon.Parse("2020-08-31").SubQuarters(2).ToDateString())

	// Output:
	// 2020-01-01
	// 2019-10-01
	// 2019-07-01
	// 2019-05-29
	// 2020-03-02
}

func ExampleCarbon_SubQuartersNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01").SubQuartersNoOverflow(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").SubQuartersNoOverflow(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").SubQuartersNoOverflow(2).ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").SubQuartersNoOverflow(3).ToDateString())
	fmt.Println(carbon.Parse("2020-08-31").SubQuartersNoOverflow(2).ToDateString())

	// Output:
	// 2020-01-01
	// 2019-10-01
	// 2019-07-01
	// 2019-05-29
	// 2020-02-29
}

func ExampleCarbon_SubQuarter() {
	fmt.Println(carbon.Parse("2020-01-01").SubQuarter().ToDateString())
	fmt.Println(carbon.Parse("2020-02-28").SubQuarter().ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").SubQuarter().ToDateString())
	fmt.Println(carbon.Parse("2020-11-30").SubQuarter().ToDateString())

	// Output:
	// 2019-10-01
	// 2019-11-28
	// 2019-11-29
	// 2020-08-30
}

func ExampleCarbon_SubQuarterNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01").SubQuarterNoOverflow().ToDateString())
	fmt.Println(carbon.Parse("2020-02-28").SubQuarterNoOverflow().ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").SubQuarterNoOverflow().ToDateString())
	fmt.Println(carbon.Parse("2020-11-30").SubQuarterNoOverflow().ToDateString())

	// Output:
	// 2019-10-01
	// 2019-11-28
	// 2019-11-29
	// 2020-08-30
}

func ExampleCarbon_AddMonths() {
	fmt.Println(carbon.Parse("2020-01-01").AddMonths(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddMonths(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddMonths(2).ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").AddMonths(3).ToDateString())
	fmt.Println(carbon.Parse("2020-08-31").AddMonths(2).ToDateString())

	// Output:
	// 2020-01-01
	// 2020-02-01
	// 2020-03-01
	// 2020-05-29
	// 2020-10-31
}

func ExampleCarbon_AddMonthsNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01").AddMonthsNoOverflow(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddMonthsNoOverflow(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddMonthsNoOverflow(2).ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").AddMonthsNoOverflow(3).ToDateString())
	fmt.Println(carbon.Parse("2020-08-31").AddMonthsNoOverflow(2).ToDateString())

	// Output:
	// 2020-01-01
	// 2020-02-01
	// 2020-03-01
	// 2020-05-29
	// 2020-10-31
}

func ExampleCarbon_AddMonth() {
	fmt.Println(carbon.Parse("2020-01-01").AddMonth().ToDateString())
	fmt.Println(carbon.Parse("2020-02-28").AddMonth().ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").AddMonth().ToDateString())
	fmt.Println(carbon.Parse("2020-11-30").AddMonth().ToDateString())

	// Output:
	// 2020-02-01
	// 2020-03-28
	// 2020-03-29
	// 2020-12-30
}

func ExampleCarbon_AddMonthNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01").AddMonthNoOverflow().ToDateString())
	fmt.Println(carbon.Parse("2020-02-28").AddMonthNoOverflow().ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").AddMonthNoOverflow().ToDateString())
	fmt.Println(carbon.Parse("2020-11-30").AddMonthNoOverflow().ToDateString())

	// Output:
	// 2020-02-01
	// 2020-03-28
	// 2020-03-29
	// 2020-12-30
}

func ExampleCarbon_SubMonths() {
	fmt.Println(carbon.Parse("2020-01-01").SubMonths(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").SubMonths(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").SubMonths(2).ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").SubMonths(3).ToDateString())
	fmt.Println(carbon.Parse("2020-08-31").SubMonths(2).ToDateString())

	// Output:
	// 2020-01-01
	// 2019-12-01
	// 2019-11-01
	// 2019-11-29
	// 2020-07-01
}

func ExampleCarbon_SubMonthsNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01").SubMonthsNoOverflow(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").SubMonthsNoOverflow(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").SubMonthsNoOverflow(2).ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").SubMonthsNoOverflow(3).ToDateString())
	fmt.Println(carbon.Parse("2020-08-31").SubMonthsNoOverflow(2).ToDateString())

	// Output:
	// 2020-01-01
	// 2019-12-01
	// 2019-11-01
	// 2019-11-29
	// 2020-06-30
}

func ExampleCarbon_SubMonth() {
	fmt.Println(carbon.Parse("2020-01-01").SubMonth().ToDateString())
	fmt.Println(carbon.Parse("2020-02-28").SubMonth().ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").SubMonth().ToDateString())
	fmt.Println(carbon.Parse("2020-11-30").SubMonth().ToDateString())

	// Output:
	// 2019-12-01
	// 2020-01-28
	// 2020-01-29
	// 2020-10-30
}

func ExampleCarbon_SubMonthNoOverflow() {
	fmt.Println(carbon.Parse("2020-01-01").SubMonthNoOverflow().ToDateString())
	fmt.Println(carbon.Parse("2020-02-28").SubMonthNoOverflow().ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").SubMonthNoOverflow().ToDateString())
	fmt.Println(carbon.Parse("2020-11-30").SubMonthNoOverflow().ToDateString())

	// Output:
	// 2019-12-01
	// 2020-01-28
	// 2020-01-29
	// 2020-10-30
}

func ExampleCarbon_AddWeeks() {
	fmt.Println(carbon.Parse("2020-01-01").AddWeeks(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddWeeks(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddWeeks(2).ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").AddWeeks(3).ToDateString())
	fmt.Println(carbon.Parse("2020-08-31").AddWeeks(2).ToDateString())

	// Output:
	// 2020-01-01
	// 2020-01-08
	// 2020-01-15
	// 2020-03-21
	// 2020-09-14
}

func ExampleCarbon_AddWeek() {
	fmt.Println(carbon.Parse("2020-01-01").AddWeek().ToDateString())
	fmt.Println(carbon.Parse("2020-02-28").AddWeek().ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").AddWeek().ToDateString())
	fmt.Println(carbon.Parse("2020-11-30").AddWeek().ToDateString())

	// Output:
	// 2020-01-08
	// 2020-03-06
	// 2020-03-07
	// 2020-12-07
}

func ExampleCarbon_SubWeeks() {
	fmt.Println(carbon.Parse("2020-01-01").SubWeeks(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").SubWeeks(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").SubWeeks(2).ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").SubWeeks(3).ToDateString())
	fmt.Println(carbon.Parse("2020-08-31").SubWeeks(2).ToDateString())

	// Output:
	// 2020-01-01
	// 2019-12-25
	// 2019-12-18
	// 2020-02-08
	// 2020-08-17
}

func ExampleCarbon_SubWeek() {
	fmt.Println(carbon.Parse("2020-01-01").SubWeek().ToDateString())
	fmt.Println(carbon.Parse("2020-02-28").SubWeek().ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").SubWeek().ToDateString())
	fmt.Println(carbon.Parse("2020-11-30").SubWeek().ToDateString())

	// Output:
	// 2019-12-25
	// 2020-02-21
	// 2020-02-22
	// 2020-11-23
}

func ExampleCarbon_AddDays() {
	fmt.Println(carbon.Parse("2020-01-01").AddDays(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddDays(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddDays(2).ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").AddDays(3).ToDateString())
	fmt.Println(carbon.Parse("2020-08-31").AddDays(2).ToDateString())

	// Output:
	// 2020-01-01
	// 2020-01-02
	// 2020-01-03
	// 2020-03-03
	// 2020-09-02
}

func ExampleCarbon_AddDay() {
	fmt.Println(carbon.Parse("2020-01-01").AddDay().ToDateString())
	fmt.Println(carbon.Parse("2020-02-28").AddDay().ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").AddDay().ToDateString())
	fmt.Println(carbon.Parse("2020-11-30").AddDay().ToDateString())

	// Output:
	// 2020-01-02
	// 2020-02-29
	// 2020-03-01
	// 2020-12-01
}

func ExampleCarbon_SubDays() {
	fmt.Println(carbon.Parse("2020-01-01").SubDays(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").SubDays(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").SubDays(2).ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").SubDays(3).ToDateString())
	fmt.Println(carbon.Parse("2020-08-31").SubDays(2).ToDateString())

	// Output:
	// 2020-01-01
	// 2019-12-31
	// 2019-12-30
	// 2020-02-26
	// 2020-08-29
}

func ExampleCarbon_SubDay() {
	fmt.Println(carbon.Parse("2020-01-01").SubDay().ToDateString())
	fmt.Println(carbon.Parse("2020-02-28").SubDay().ToDateString())
	fmt.Println(carbon.Parse("2020-02-29").SubDay().ToDateString())
	fmt.Println(carbon.Parse("2020-11-30").SubDay().ToDateString())

	// Output:
	// 2019-12-31
	// 2020-02-27
	// 2020-02-28
	// 2020-11-29
}

func ExampleCarbon_AddWeekdays() {
	fmt.Println(carbon.Parse("2020-01-01").AddWeekdays(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddWeekdays(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddWeekdays(3).ToDateString())
	fmt.Println(carbon.Parse("2020-01-01").AddWeekdays(5).ToDateString())
	fmt.Println(carbon.Parse("2020-01-04").AddWeekdays(1).ToDateString())

	// Output:
	// 2020-01-01
	// 2020-01-02
	// 2020-01-06
	// 2020-01-08
	// 2020-01-06
}

func ExampleCarbon_AddWeekday() {
	fmt.Println(carbon.Parse("2020-01-01").AddWeekday().ToDateString())
	fmt.Println(carbon.Parse("2020-01-03").AddWeekday().ToDateString())
	fmt.Println(carbon.Parse("2020-01-04").AddWeekday().ToDateString())

	// Output:
	// 2020-01-02
	// 2020-01-06
	// 2020-01-06
}

func ExampleCarbon_SubWeekdays() {
	fmt.Println(carbon.Parse("2020-01-06").SubWeekdays(0).ToDateString())
	fmt.Println(carbon.Parse("2020-01-06").SubWeekdays(1).ToDateString())
	fmt.Println(carbon.Parse("2020-01-06").SubWeekdays(3).ToDateString())
	fmt.Println(carbon.Parse("2020-01-06").SubWeekdays(5).ToDateString())
	fmt.Println(carbon.Parse("2020-01-04").SubWeekdays(1).ToDateString())

	// Output:
	// 2020-01-06
	// 2020-01-03
	// 2020-01-01
	// 2019-12-30
	// 2020-01-03
}

func ExampleCarbon_SubWeekday() {
	fmt.Println(carbon.Parse("2020-01-06").SubWeekday().ToDateString())
	fmt.Println(carbon.Parse("2020-01-04").SubWeekday().ToDateString())

	// Output:
	// 2020-01-03
	// 2020-01-03
}

func ExampleCarbon_AddHours() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddHours(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddHours(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddHours(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2020-01-01 14:14:15 +0000 UTC
	// 2020-01-01 15:14:15 +0000 UTC
}

func ExampleCarbon_AddHour() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").AddHour().ToString())

	// Output:
	// 2020-08-05 14:14:15 +0000 UTC
}

func ExampleCarbon_SubHours() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubHours(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubHours(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubHours(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2020-01-01 12:14:15 +0000 UTC
	// 2020-01-01 11:14:15 +0000 UTC
}

func ExampleCarbon_SubHour() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").SubHour().ToString())

	// Output:
	// 2020-08-05 12:14:15 +0000 UTC
}

func ExampleCarbon_AddMinutes() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddMinutes(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddMinutes(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddMinutes(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2020-01-01 13:15:15 +0000 UTC
	// 2020-01-01 13:16:15 +0000 UTC
}

func ExampleCarbon_AddMinute() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").AddMinute().ToString())

	// Output:
	// 2020-08-05 13:15:15 +0000 UTC
}

func ExampleCarbon_SubMinutes() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubMinutes(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubMinutes(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubMinutes(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2020-01-01 13:13:15 +0000 UTC
	// 2020-01-01 13:12:15 +0000 UTC
}

func ExampleCarbon_SubMinute() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").SubMinute().ToString())

	// Output:
	// 2020-08-05 13:13:15 +0000 UTC
}

func ExampleCarbon_AddSeconds() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddSeconds(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddSeconds(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddSeconds(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2020-01-01 13:14:16 +0000 UTC
	// 2020-01-01 13:14:17 +0000 UTC
}

func ExampleCarbon_AddSecond() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").AddSecond().ToString())

	// Output:
	// 2020-08-05 13:14:16 +0000 UTC
}

func ExampleCarbon_SubSeconds() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubSeconds(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubSeconds(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubSeconds(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2020-01-01 13:14:14 +0000 UTC
	// 2020-01-01 13:14:13 +0000 UTC
}

func ExampleCarbon_SubSecond() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").SubSecond().ToString())

	// Output:
	// 2020-08-05 13:14:14 +0000 UTC
}

func ExampleCarbon_AddMilliseconds() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddMilliseconds(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddMilliseconds(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddMilliseconds(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2020-01-01 13:14:15.001 +0000 UTC
	// 2020-01-01 13:14:15.002 +0000 UTC
}

func ExampleCarbon_AddMillisecond() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").AddMillisecond().ToString())

	// Output:
	// 2020-08-05 13:14:15.001 +0000 UTC
}

func ExampleCarbon_SubMilliseconds() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubMilliseconds(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubMilliseconds(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubMilliseconds(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2020-01-01 13:14:14.999 +0000 UTC
	// 2020-01-01 13:14:14.998 +0000 UTC
}

func ExampleCarbon_SubMillisecond() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").SubMillisecond().ToString())

	// Output:
	// 2020-08-05 13:14:14.999 +0000 UTC
}

func ExampleCarbon_AddMicroseconds() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddMicroseconds(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddMicroseconds(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddMicroseconds(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2020-01-01 13:14:15.000001 +0000 UTC
	// 2020-01-01 13:14:15.000002 +0000 UTC
}

func ExampleCarbon_AddMicrosecond() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").AddMicrosecond().ToString())

	// Output:
	// 2020-08-05 13:14:15.000001 +0000 UTC
}

func ExampleCarbon_SubMicroseconds() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubMicroseconds(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubMicroseconds(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubMicroseconds(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2020-01-01 13:14:14.999999 +0000 UTC
	// 2020-01-01 13:14:14.999998 +0000 UTC
}

func ExampleCarbon_SubMicrosecond() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").SubMicrosecond().ToString())

	// Output:
	// 2020-08-05 13:14:14.999999 +0000 UTC
}

func ExampleCarbon_AddNanoseconds() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddNanoseconds(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddNanoseconds(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").AddNanoseconds(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2020-01-01 13:14:15.000000001 +0000 UTC
	// 2020-01-01 13:14:15.000000002 +0000 UTC
}

func ExampleCarbon_AddNanosecond() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").AddNanosecond().ToString())

	// Output:
	// 2020-08-05 13:14:15.000000001 +0000 UTC
}

func ExampleCarbon_SubNanoseconds() {
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubNanoseconds(0).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubNanoseconds(1).ToString())
	fmt.Println(carbon.Parse("2020-01-01 13:14:15").SubNanoseconds(2).ToString())

	// Output:
	// 2020-01-01 13:14:15 +0000 UTC
	// 2020-01-01 13:14:14.999999999 +0000 UTC
	// 2020-01-01 13:14:14.999999998 +0000 UTC
}

func ExampleCarbon_SubNanosecond() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").SubNanosecond().ToString())

	// Output:
	// 2020-08-05 13:14:14.999999999 +0000 UTC
}
