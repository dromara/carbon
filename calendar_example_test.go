package carbon_test

import (
	"fmt"

	"github.com/dromara/carbon/v2"
)

func ExampleCarbon_Lunar() {
	fmt.Println(carbon.Parse("2023-03-02", carbon.PRC).Lunar().String())
	fmt.Println(carbon.Parse("2023-04-01", carbon.PRC).Lunar().String())

	// Output:
	// 2023-02-11
	// 2023-02-11
}

func ExampleCreateFromLunar() {
	fmt.Println(carbon.CreateFromLunar(2023, 2, 11, false).ToString(carbon.PRC))
	fmt.Println(carbon.CreateFromLunar(2023, 2, 11, true).ToString(carbon.PRC))

	// Output:
	// 2023-03-02 00:00:00 +0800 CST
	// 2023-04-01 00:00:00 +0800 CST
}

func ExampleCarbon_Julian() {
	fmt.Println("JD(default precision 6):", carbon.Parse("2024-01-23 13:14:15").Julian().JD())
	fmt.Println("MJD(default precision 6):", carbon.Parse("2024-01-23 13:14:15").Julian().MJD())

	fmt.Println("JD(specify precision 4):", carbon.Parse("2024-01-23 13:14:15").Julian().JD(4))
	fmt.Println("MJD(specify precision 4):", carbon.Parse("2024-01-23 13:14:15").Julian().MJD(4))

	// Output:
	// JD(default precision 6): 2.460333051563e+06
	// MJD(default precision 6): 60332.551563
	// JD(specify precision 4): 2.4603330516e+06
	// MJD(specify precision 4): 60332.5516
}

func ExampleCreateFromJulian() {
	fmt.Println(carbon.CreateFromJulian(2.460333051563e+06).ToString())
	fmt.Println(carbon.CreateFromJulian(60332.551563).ToString())

	// Output:
	// 2024-01-23 13:14:15 +0000 UTC
	// 2024-01-23 13:14:15 +0000 UTC
}

func ExampleCarbon_Persian() {
	fmt.Println(carbon.Parse("1800-01-01 00:00:00").Persian().String())
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").Persian().String())
	fmt.Println(carbon.Parse("2024-01-01 00:00:00").Persian().String())

	// Output:
	// 1178-10-11
	// 1399-05-15
	// 1402-10-11
}

func ExampleCreateFromPersian() {
	fmt.Println(carbon.CreateFromPersian(1178, 10, 11).ToDateString())
	fmt.Println(carbon.CreateFromPersian(1402, 10, 11).ToDateString())
	fmt.Println(carbon.CreateFromPersian(1403, 5, 15).ToDateString())

	// Output:
	// 1800-01-01
	// 2024-01-01
	// 2024-08-05
}
