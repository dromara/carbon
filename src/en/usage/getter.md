# Getter

```go
// Get total days of the year
carbon.Parse("2019-08-05 13:14:15").DaysInYear() // 365
carbon.Parse("2020-08-05 13:14:15").DaysInYear() // 366
// Get total days of the month
carbon.Parse("2020-02-01 13:14:15").DaysInMonth() // 29
carbon.Parse("2020-04-01 13:14:15").DaysInMonth() // 30
carbon.Parse("2020-08-01 13:14:15").DaysInMonth() // 31

// Get day of the year
carbon.Parse("2020-08-05 13:14:15").DayOfYear() // 218
// Get week of the year
carbon.Parse("2020-08-05 13:14:15").WeekOfYear() // 32
// Get day of the month
carbon.Parse("2020-08-05 13:14:15").DayOfMonth() // 5
// Get week of the month
carbon.Parse("2020-08-05 13:14:15").WeekOfMonth() // 1
// Get day of the week
carbon.Parse("2020-08-05 13:14:15").DayOfWeek() // 3

// Get current year, month, day, hour, minute and second
carbon.Parse("2020-08-05 13:14:15").DateTime() // 2020, 8, 5, 13, 14, 15
// Get current year, month, day, hour, minute, second and millisecond
carbon.Parse("2020-08-05 13:14:15").DateTimeMilli() // 2020, 8, 5, 13, 14, 15, 999
// Get current year, month, day, hour, minute, second and microsecond
carbon.Parse("2020-08-05 13:14:15").DateTimeMicro() // 2020, 8, 5, 13, 14, 15, 999999
// Get current year, month, day, hour, minute, second and nanosecond
carbon.Parse("2020-08-05 13:14:15").DateTimeNano() // 2020, 8, 5, 13, 14, 15, 999999999

// Get current year, month and day
carbon.Parse("2020-08-05 13:14:15.999999999").Date() // 2020, 8, 5
// Get current year, month, day and millisecond
carbon.Parse("2020-08-05 13:14:15.999999999").DateMilli() // 2020, 8, 5, 999
// Get current year, month, day and microsecond
carbon.Parse("2020-08-05 13:14:15.999999999").DateMicro() // 2020, 8, 5, 999999
// Get current year, month, day and nanosecond
carbon.Parse("2020-08-05 13:14:15.999999999").DateNano() // 2020, 8, 5, 999999999

// Get current hour, minute and second
carbon.Parse("2020-08-05 13:14:15.999999999").Time() // 13, 14, 15
// Get current hour, minute, second and millisecond
carbon.Parse("2020-08-05 13:14:15.999999999").TimeMilli() // 13, 14, 15, 999
// Get current hour, minute, second and microsecond
carbon.Parse("2020-08-05 13:14:15.999999999").TimeMicro() // 13, 14, 15, 999999
// Get current hour, minute, second and nanosecond
carbon.Parse("2020-08-05 13:14:15.999999999").TimeNano() // 13, 14, 15, 999999999

// Get current century
carbon.Parse("2020-08-05 13:14:15").Century() // 21
// Get current decade
carbon.Parse("2019-08-05 13:14:15").Decade() // 10
carbon.Parse("2021-08-05 13:14:15").Decade() // 20
// Get current year
carbon.Parse("2020-08-05 13:14:15").Year() // 2020
// Get current quarter
carbon.Parse("2020-08-05 13:14:15").Quarter() // 3
// Get current month
carbon.Parse("2020-08-05 13:14:15").Month() // 8
// Get current week(start from 0)
carbon.Parse("2020-08-02 13:14:15").Week() // 0
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Sunday).Week() // 0
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Monday).Week() // 6
// Get current day
carbon.Parse("2020-08-05 13:14:15").Day() // 5
// Get current hour
carbon.Parse("2020-08-05 13:14:15").Hour() // 13
// Get current minute
carbon.Parse("2020-08-05 13:14:15").Minute() // 14
// Get current second
carbon.Parse("2020-08-05 13:14:15").Second() // 15
// Get current millisecond
carbon.Parse("2020-08-05 13:14:15.999").Millisecond() // 999
// Get current microsecond
carbon.Parse("2020-08-05 13:14:15.999").Microsecond() // 999000
// Get current nanosecond
carbon.Parse("2020-08-05 13:14:15.999").Nanosecond() // 999000000

// Get timestamp with second precision
carbon.Parse("2020-08-05 13:14:15").Timestamp() // 1596633255
// Get timestamp with millisecond precision
carbon.Parse("2020-08-05 13:14:15.999").TimestampMilli() // 1596633255999
// Get timestamp with microsecond precision
carbon.Parse("2020-08-05 13:14:15.999999").TimestampMicro() // 1596633255999999
// Get timestamp with nanosecond precision
carbon.Parse("2020-08-05 13:14:15.999999999").TimestampNano() // 1596633255999999999

// Get timezone location
carbon.SetTimezone(carbon.PRC).Timezone() // PRC
carbon.SetTimezone(carbon.Tokyo).Timezone() // Asia/Tokyo

// Get timezone name
carbon.SetTimezone(carbon.PRC).ZoneName() // CST
carbon.SetTimezone(carbon.Tokyo).ZoneName() // JST

// Get timezone offset seconds from the UTC timezone
carbon.SetTimezone(carbon.PRC).ZoneOffset() // 28800
carbon.SetTimezone(carbon.Tokyo).ZoneOffset() // 32400

// Get locale name
carbon.Now().SetLocale("en").Locale() // en
carbon.Now().SetLocale("zh-CN").Locale() // zh-CN

// Get constellation name
carbon.Now().Constellation() // Leo
carbon.Now().SetLocale("en").Constellation() // Leo
carbon.Now().SetLocale("zh-CN").Constellation() // 狮子座

// Get season name
carbon.Now().Season() // Summer
carbon.Now().SetLocale("en").Season() // Summer
carbon.Now().SetLocale("zh-CN").Season() // 夏季

// Get start day of the week
carbon.SetWeekStartsAt(carbon.Sunday).WeekStartsAt() // Sunday
carbon.SetWeekStartsAt(carbon.Monday).WeekStartsAt() // Monday

// Get end day of the week
carbon.SetWeekStartsAt(carbon.Sunday).WeekEndsAt() // Saturday
carbon.SetWeekStartsAt(carbon.Monday).WeekEndsAt() // Sunday

// Get current layout
carbon.Parse("now").CurrentLayout() // "2006-01-02 15:04:05"
carbon.ParseByLayout("2020-08-05", DateLayout).CurrentLayout() // "2006-01-02"

// Get current age
carbon.Parse("2002-01-01 13:14:15").Age() // 17
carbon.Parse("2002-12-31 13:14:15").Age() // 18
```