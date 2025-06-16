---
head:
  - - meta
    - name: description
      content: Calendar | A simple, semantic and developer-friendly time package for golang
  - - meta
    - name: keywords
      content: carbon, go-carbon, gregorian, julian day, modified julian day, chinese lunar, persian, jalaali
---

# Calendar

The following calendars are supported
- Gregorian
- Julian Day/Modified Julian Day
- Chinese Lunar
- Persian(Jalaali)

## Julian Day/Modified Julian Day

### Convert `gregorian` to `julian day`
```go
// By default, 6 decimal places are retained for precision
carbon.Parse("2024-01-24 12:00:00").Julian().JD() // 2460334
carbon.Parse("2024-01-24 13:14:15").Julian().JD() // 2460334.051563

// 4 decimal places are retained for precision
carbon.Parse("2024-01-24 12:00:00").Julian().JD(4) // 2460334
carbon.Parse("2024-01-24 13:14:15").Julian().JD(4) // 2460334.0516
```

### Convert `gregorian` to `modified julian day`
```go
// By default, 6 decimal places are retained for precision
carbon.Parse("2024-01-24 12:00:00").Julian().MJD() // 60333.5
carbon.Parse("2024-01-24 13:14:15").Julian().MJD() // 60333.551563

// 4 decimal places are retained for precision
carbon.Parse("2024-01-24 12:00:00").Julian().MJD(4) // 60333.5
carbon.Parse("2024-01-24 13:14:15").Julian().MJD(4) // 60333.5516
```

### Convert `julian day` to `modified dulian Day`
```go
// By default, 6 decimal places are retained for precision
carbon.CreateFromJulian(2460334).Julian().MJD() // 60333.5
carbon.CreateFromJulian(2460334.051563).Julian().MJD() // 60332.551563

// 4 decimal places are retained for precision
carbon.CreateFromJulian(2460334).Julian().MJD(4) // 60333.5
carbon.CreateFromJulian(2460334.051563).Julian().MJD(4) // 60332.5516
```

### Convert `modified julian day` to `julian day`
```go
// By default, 6 decimal places are retained for precision
carbon.CreateFromJulian(60333.5).Julian().JD()() // 2460334
carbon.CreateFromJulian(60333.551563).Julian().JD()() // 2460333.051563

// 4 decimal places are retained for precision
carbon.CreateFromJulian(60333.5).Julian().JD(4) // 2460334
carbon.CreateFromJulian(60333.551563).Julian().JD(4) // 2460333.0516
```

### Convert `modified julian day` to `julian day`
```go
// By default, 6 decimal places are retained for precision
carbon.CreateFromJulian(60333.5).Julian().JD()() // 2460334
carbon.CreateFromJulian(60333.551563).Julian().JD()() // 2460333.051563

// 4 decimal places are retained for precision
carbon.CreateFromJulian(60333.5).Julian().JD(4) // 2460334
carbon.CreateFromJulian(60333.551563).Julian().JD(4) // 2460333.0516
```

### Convert `julian day` to `gregorian`
```go
carbon.CreateFromJulian(2460334).ToDateTimeString() // 2024-01-24 12:00:00
carbon.CreateFromJulian(2460334.051563).ToDateTimeString() // 2024-01-24 13:14:15
```

### Convert `modified julian day` to `gregorian`
```go
carbon.CreateFromJulian(60333.551563).ToDateTimeString() // 2024-01-24 13:14:15
carbon.CreateFromJulian(60333.5).ToDateTimeString() // 2024-01-24 12:00:00
```

## Chinese Lunar
> Currently only `200` years from `1900` to `2100` are supported

### Convert `gregorian` to `lunar`

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

### Convert `lunar` to `gregorian`

```go
// Convert the Lunar Calendar December 11, 2023 to the gregorian calendar
carbon.CreateFromLunar(2023, 12, 11, 0, 0, 0, false).ToDateTimeString() // 2024-01-21 00:00:00
// Convert lunar calendar February 11, 2023 to gregorian calendar
carbon.CreateFromLunar(2023, 2, 11, 0, 0, 0, false).ToDateTimeString() // 2024-03-02 00:00:00
// Convert the Lunar Calendar Leap February 11, 2024 to the gregorian calendar
carbon.CreateFromLunar(2023, 2, 11, 0, 0, 0, true).ToDateTimeString() // 2023-04-01 00:00:00
```

### Lunar compare
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

## Persian(Jalaali)

### Convert `gregorian` to `persian`
```go
// Get persian year
carbon.Parse("2020-08-05").Persian().Year() // 1399
// Get persian month
carbon.Parse("2020-08-05").Persian().Month() // 5
// Get persian day
carbon.Parse("2020-08-05").Persian().Day() // 15
// Get persian hour
carbon.Parse("2020-08-05").Persian().Hour() // 13
// Get persian minute
carbon.Parse("2020-08-05").Persian().Minute() // 14
// Get persian second
carbon.Parse("2020-08-05").Persian().Second() // 15

// Get persian date and time string
carbon.Parse("2020-08-05").Persian().String() // 1399-05-15
fmt.Printf("%s", carbon.Parse("2020-08-05").Persian()) // 1399-05-15

// Get persian month as string
carbon.Parse("2020-08-05").Persian().ToMonthString() // Mordad
carbon.Parse("2020-08-05").Persian().ToMonthString("en") // Mordad
carbon.Parse("2020-08-05").Persian().ToMonthString("fa") // مرداد

// Get persian short month as string
carbon.Parse("2020-08-05").Persian().ToShortMonthString() // Mor
carbon.Parse("2020-08-05").Persian().ToShortMonthString("en") // Mor
carbon.Parse("2020-08-05").Persian().ToShortMonthString("fa") // مرد

// Get persian week as string
carbon.Parse("2020-08-05").Persian().ToWeekString() // Chaharshanbeh
carbon.Parse("2020-08-05").Persian().ToWeekString("en") // Chaharshanbeh
carbon.Parse("2020-08-05").Persian().ToWeekString("fa") // چهارشنبه

// Get persian short week as string
carbon.Parse("2020-08-05").Persian().ToShortWeekString() // Cha
carbon.Parse("2020-08-05").Persian().ToShortWeekString("en") // Cha
carbon.Parse("2020-08-05").Persian().ToShortWeekString("fa") // د
```

### Convert `persian` to `gregorian`
```go
carbon.CreateFromPersian(1, 1, 1).ToDateTimeString() // 2016-03-20 00:00:00
carbon.CreateFromPersian(622, 1, 1).ToDateTimeString() // 1243-03-21 00:00:00
carbon.CreateFromPersian(1395, 1, 1).ToDateTimeString() // 2016-03-20 00:00:00
carbon.CreateFromPersian(9377, 1, 1).ToDateTimeString() // 9998-03-19 00:00:00
```

### Persian compare
```go
// Whether is a valid persian date
carbon.CreateFromPersian(1, 1, 1).IsValid() // true
carbon.CreateFromPersian(622, 1, 1).IsValid() // true
carbon.CreateFromPersian(9377, 1, 1).IsValid() // true
carbon.CreateFromPersian(0, 0, 0, 0).IsValid() // false
carbon.CreateFromPersian(2024, 0, 1).IsValid() // false
carbon.CreateFromPersian(2024, 1, 0).IsValid() // false

// Whether is a persian leap year
carbon.CreateFromPersian(1395, 1, 1).IsLeapYear() // true
carbon.CreateFromPersian(9377, 1, 1).IsLeapYear() // true
carbon.CreateFromPersian(622, 1, 1).IsLeapYear() // false
carbon.CreateFromPersian(9999, 1, 1).IsLeapYear() // false
```