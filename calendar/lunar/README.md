# Chinese Lunar

English | [简体中文](README.cn.md) | [日本語](README.jp.md)

#### Usage and example

> Currently only `200` years from `1900` to `2100` are supported

##### Convert `Gregorian` calendar to `Lunar` calendar

```go
// Get Lunar year of animal
carbon.Parse("2020-08-05 13:14:15").Lunar().Animal() // 鼠
// Get lunar festival
carbon.Parse("2021-02-12").Lunar().Festival() // 春节

// Get lunar year
carbon.Parse("2020-08-05").Lunar().Year() // 2020
// Get lunar month
carbon.Parse("2020-08-05").Lunar().Month() // 6
// Get lunar leap month
carbon.Parse("2020-08-05").Lunar().LeapMonth() // 4
// Get lunar day
carbon.Parse("2020-08-05").Lunar().Day() // 16
// Get lunar hour
carbon.Parse("2020-08-05").Lunar().Hour() // 13
// Get lunar minute
carbon.Parse("2020-08-05").Lunar().Minute() // 14
// Get lunar second
carbon.Parse("2020-08-05").Lunar().Second() // 15

// Get lunar date and time string
carbon.Parse("2020-08-05").Lunar().String() // 2020-06-16
fmt.Printf("%s", carbon.Parse("2020-08-05").Lunar()) // 2020-06-16
// Get lunar year as string
carbon.Parse("2020-08-05").Lunar().ToYearString() // 二零二零
// Get lunar month as string
carbon.Parse("2020-08-05").Lunar().ToMonthString() // 六月
// Get lunar leap month as string
carbon.Parse("2020-04-23").Lunar().ToMonthString() // 闰四月
// Get lunar week as string
carbon.Parse("2020-04-23").Lunar().ToWeekString() // 周四
// Get lunar day as string
carbon.Parse("2020-08-05").Lunar().ToDayString() // 十六
// Get lunar date as string
carbon.Parse("2020-08-05").Lunar().ToDateString() // 二零二零年六月十六

```

##### Convert `Lunar` calendar to `Gregorian` calendar

```go
// Convert the Lunar Calendar December 11, 2023 to the gregorian calendar
carbon.CreateFromLunar(2023, 12, 11, 0, 0, 0, false).ToDateTimeString() // 2024-01-21 00:00:00
// Convert lunar calendar February 11, 2023 to gregorian calendar
carbon.CreateFromLunar(2023, 2, 11, 0, 0, 0, false).ToDateTimeString() // 2024-03-02 00:00:00
// Convert the Lunar Calendar Leap February 11, 2024 to the gregorian calendar
carbon.CreateFromLunar(2023, 2, 11, 0, 0, 0, true).ToDateTimeString() // 2023-04-01 00:00:00
```

##### Comparison
```go
// Whether is a valid lunar date
carbon.Parse("0000-00-00").Lunar().IsValid() // false
carbon.Parse("2020-08-05").Lunar().IsValid() // true

// Whether is a lunar leap year
carbon.Parse("2020-08-05").Lunar().IsLeapYear() // true
// Whether is a lunar leap month
carbon.Parse("2020-08-05").Lunar().IsLeapMonth() // false

// Whether is a lunar year of the rat
carbon.Parse("2020-08-05").Lunar().IsRatYear() // true
// Whether is a lunar year of the ox
carbon.Parse("2020-08-05").Lunar().IsOxYear() // false
// Whether is a lunar year of the tiger
carbon.Parse("2020-08-05").Lunar().IsTigerYear() // false
// Whether is a lunar year of the rabbit
carbon.Parse("2020-08-05").Lunar().IsRabbitYear() // false
// Whether is a lunar year of the dragon
carbon.Parse("2020-08-05").Lunar().IsDragonYear() // false
// Whether is a lunar year of the snake
carbon.Parse("2020-08-05").Lunar().IsSnakeYear() // false
// Whether is a lunar year of the horse
carbon.Parse("2020-08-05").Lunar().IsHorseYear() // false
// Whether is a lunar year of the goat
carbon.Parse("2020-08-05").Lunar().IsGoatYear() // false
// Whether is a lunar year of the monkey
carbon.Parse("2020-08-05").Lunar().IsMonkeyYear() // false
// Whether is a lunar year of the rooster
carbon.Parse("2020-08-05").Lunar().IsRoosterYear() // false
// Whether is a lunar year of the dog
carbon.Parse("2020-08-05").Lunar().IsDogYear() // false
// Whether is a lunar year of the dig
carbon.Parse("2020-08-05").Lunar().IsPigYear() // false

```