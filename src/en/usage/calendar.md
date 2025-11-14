---
head:
  - - meta
    - name: description
      content: "Supports Julian Day/Modified Julian Day, Chinese Lunar, Persian, and Hebrew calendars; provides conversion with Gregorian calendar, formatting (month/week/day strings), validation and leap year checks"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, calendar, gregorian calendar, julian day, modified julian day, chinese lunar, persian calendar, iranian calendar, hebrew calendar, jewish calendar"
---

# Calendar

The following calendars are currently supported

- [Julian Day/Modified Julian Day](#julian-day-modified-julian-day)
- [Chinese Lunar](#chinese-lunar)
- [Persian (Jalaali)](#persian-jalaali)
- [Hebrew (Jewish)](#hebrew-jewish)

## Julian Day/Modified Julian Day

### Convert `Gregorian` to `Julian Day`
```go
// By default, 6 decimal places are retained for precision
carbon.Parse("2024-01-24 12:00:00").Julian().JD() // 2460334
carbon.Parse("2024-01-24 13:14:15").Julian().JD() // 2460334.051563

// 4 decimal places are retained for precision
carbon.Parse("2024-01-24 12:00:00").Julian().JD(4) // 2460334
carbon.Parse("2024-01-24 13:14:15").Julian().JD(4) // 2460334.0516
```

### Convert `Gregorian` to `Modified Julian Day`
```go
// By default, 6 decimal places are retained for precision
carbon.Parse("2024-01-24 12:00:00").Julian().MJD() // 60333.5
carbon.Parse("2024-01-24 13:14:15").Julian().MJD() // 60333.551563

// 4 decimal places are retained for precision
carbon.Parse("2024-01-24 12:00:00").Julian().MJD(4) // 60333.5
carbon.Parse("2024-01-24 13:14:15").Julian().MJD(4) // 60333.5516
```

### Convert `Julian Day` to `Modified Julian Day`
```go
// By default, 6 decimal places are retained for precision
carbon.CreateFromJulian(2460334).Julian().MJD() // 60333.5
carbon.CreateFromJulian(2460334.051563).Julian().MJD() // 60332.551563

// 4 decimal places are retained for precision
carbon.CreateFromJulian(2460334).Julian().MJD(4) // 60333.5
carbon.CreateFromJulian(2460334.051563).Julian().MJD(4) // 60332.5516
```

### Convert `Modified Julian Day` to `Julian Day`
```go
// By default, 6 decimal places are retained for precision
carbon.CreateFromJulian(60333.5).Julian().JD() // 2460334
carbon.CreateFromJulian(60333.551563).Julian().JD() // 2460333.051563

// 4 decimal places are retained for precision
carbon.CreateFromJulian(60333.5).Julian().JD(4) // 2460334
carbon.CreateFromJulian(60333.551563).Julian().JD(4) // 2460333.0516
```

### Convert `Julian Day`/`Modified Julian Day` to `Gregorian`
```go
// Convert Julian Day to Gregorian
carbon.CreateFromJulian(2460334).ToDateTimeString() // 2024-01-24 12:00:00
carbon.CreateFromJulian(2460334.051563).ToDateTimeString() // 2024-01-24 13:14:15

// Convert Modified Julian Day to Gregorian
carbon.CreateFromJulian(60333.5).ToDateTimeString() // 2024-01-24 12:00:00
carbon.CreateFromJulian(60333.551563).ToDateTimeString() // 2024-01-24 13:14:15
```

## Chinese Lunar
> Currently only supports a 200-year span from `1900` to `2100`

### Convert `Gregorian` to `Lunar`
```go
// Get lunar zodiac animal
carbon.Parse("2020-08-05").Lunar().Animal() // 鼠
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

### Convert `Lunar` to `Gregorian`
```go
// Convert lunar December 11, 2023 to Gregorian
carbon.CreateFromLunar(2023, 12, 11, false).ToDateString() // 2024-01-21
// Convert lunar February 11, 2023 to Gregorian
carbon.CreateFromLunar(2023, 2, 11, false).ToDateString() // 2023-03-02
// Convert lunar leap February 11, 2023 to Gregorian
carbon.CreateFromLunar(2023, 2, 11, true).ToDateString() // 2023-04-01
```

### Lunar Date Validation
```go
// Whether is a valid lunar date
carbon.Parse("0000-00-00").Lunar().IsValid() // false
carbon.Parse("2020-08-05").Lunar().IsValid() // true

// Whether is a lunar leap year
carbon.Parse("2020-08-05").Lunar().IsLeapYear() // true
// Whether is a lunar leap month
carbon.Parse("2020-08-05").Lunar().IsLeapMonth() // false

// Whether is the year of the rat
carbon.Parse("2020-08-05").Lunar().IsRatYear() // true
// Whether is the year of the ox
carbon.Parse("2020-08-05").Lunar().IsOxYear() // false
// Whether is the year of the tiger
carbon.Parse("2020-08-05").Lunar().IsTigerYear() // false
// Whether is the year of the rabbit
carbon.Parse("2020-08-05").Lunar().IsRabbitYear() // false
// Whether is the year of the dragon
carbon.Parse("2020-08-05").Lunar().IsDragonYear() // false
// Whether is the year of the snake
carbon.Parse("2020-08-05").Lunar().IsSnakeYear() // false
// Whether is the year of the horse
carbon.Parse("2020-08-05").Lunar().IsHorseYear() // false
// Whether is the year of the goat
carbon.Parse("2020-08-05").Lunar().IsGoatYear() // false
// Whether is the year of the monkey
carbon.Parse("2020-08-05").Lunar().IsMonkeyYear() // false
// Whether is the year of the rooster
carbon.Parse("2020-08-05").Lunar().IsRoosterYear() // false
// Whether is the year of the dog
carbon.Parse("2020-08-05").Lunar().IsDogYear() // false
// Whether is the year of the pig
carbon.Parse("2020-08-05").Lunar().IsPigYear() // false
```

## Persian (Jalaali)

### Convert `Gregorian` to `Persian`
```go
// Get Persian year
carbon.Parse("2020-08-05").Persian().Year() // 1399
// Get Persian month
carbon.Parse("2020-08-05").Persian().Month() // 5
// Get Persian day
carbon.Parse("2020-08-05").Persian().Day() // 15

// Get Persian date and time string
carbon.Parse("2020-08-05").Persian().String() // 1399-05-15
fmt.Printf("%s", carbon.Parse("2020-08-05").Persian()) // 1399-05-15

// Get Persian month as string
carbon.Parse("2020-08-05").Persian().ToMonthString() // Mordad
carbon.Parse("2020-08-05").Persian().ToMonthString("en") // Mordad
carbon.Parse("2020-08-05").Persian().ToMonthString("fa") // مرداد

// Get Persian week as string
carbon.Parse("2020-08-05").Persian().ToWeekString() // Chaharshanbeh
carbon.Parse("2020-08-05").Persian().ToWeekString("en") // Chaharshanbeh
carbon.Parse("2020-08-05").Persian().ToWeekString("fa") // چهارشنبه
```

### Convert `Persian` to `Gregorian`
```go
carbon.CreateFromPersian(1, 1, 1).ToDateString() // 2016-03-20
carbon.CreateFromPersian(622, 1, 1).ToDateString() // 1243-03-21
carbon.CreateFromPersian(1395, 1, 1).ToDateString() // 2016-03-20
carbon.CreateFromPersian(9377, 1, 1).ToDateString() // 9998-03-19
```

### Persian Date Validation
```go
// Whether is a valid Persian date
carbon.CreateFromPersian(1, 1, 1).IsValid() // true
carbon.CreateFromPersian(622, 1, 1).IsValid() // true
carbon.CreateFromPersian(9377, 1, 1).IsValid() // true
carbon.CreateFromPersian(0, 0, 0, 0).IsValid() // false
carbon.CreateFromPersian(2024, 0, 1).IsValid() // false
carbon.CreateFromPersian(2024, 1, 0).IsValid() // false

// Whether is a Persian leap year
carbon.CreateFromPersian(1395, 1, 1).IsLeapYear() // true
carbon.CreateFromPersian(9377, 1, 1).IsLeapYear() // true
carbon.CreateFromPersian(622, 1, 1).IsLeapYear() // false
carbon.CreateFromPersian(9999, 1, 1).IsLeapYear() // false
```

## Hebrew (Jewish)

### Convert `Gregorian` to `Hebrew`
```go
// Get Hebrew year
carbon.Parse("2024-01-01").Hebrew().Year() // 5784
// Get Hebrew month
carbon.Parse("2024-01-01").Hebrew().Month() // 10
// Get Hebrew day
carbon.Parse("2024-01-01").Hebrew().Day() // 20

// Get Hebrew date and time string
carbon.Parse("2024-01-01").Hebrew().String() // 5784-10-20
fmt.Printf("%s", carbon.Parse("2024-01-01").Hebrew()) // 5784-10-20

// Get Hebrew month as string
carbon.Parse("2020-08-05").Hebrew().ToMonthString() // Av
carbon.Parse("2020-08-05").Hebrew().ToMonthString("en") // Av
carbon.Parse("2020-08-05").Hebrew().ToMonthString("he") // אב

// Get Hebrew week as string
carbon.Parse("2020-08-05").Hebrew().ToWeekString() // Wednesday
carbon.Parse("2020-08-05").Hebrew().ToWeekString("en") // Wednesday
carbon.Parse("2020-08-05").Hebrew().ToWeekString("he") // רביעי
```

### Convert `Hebrew` to `Gregorian`
```go
carbon.CreateFromHebrew(5784, 10, 20).ToDateString() // 2023-12-17
carbon.CreateFromHebrew(5784, 5, 1).ToDateString() // 2024-07-21
carbon.CreateFromHebrew(5786, 7, 10).ToDateString() // 2025-09-18
```

### Hebrew Date Validation
```go
// Whether is a valid Hebrew date
carbon.CreateFromHebrew(5780, 14, 1).IsValid() // false
carbon.CreateFromHebrew(5780, 10, 30).IsValid() // true

// Whether is a Hebrew leap year
carbon.CreateFromHebrew(5784, 1, 1).IsLeapYear() // true
carbon.CreateFromHebrew(5788, 1, 1).IsLeapYear() // false
```