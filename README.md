# Carbon
[![Build Status](https://github.com/golang-module/carbon/workflows/Go/badge.svg)](https://github.com/golang-module/carbon/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-module/carbon)](https://goreportcard.com/report/github.com/golang-module/carbon)
[![codecov](https://codecov.io/gh/golang-module/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/golang-module/carbon)
[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/golang-module/carbon)

English | [Chinese](README.cn.md)

#### Introduction
A simple, semantic and developer-friendly golang package for datetime

If you think it is helpful, please give me a star

github:[github.com/golang-module/carbon](https://github.com/golang-module/carbon "github.com/golang-module/carbon")

gitee:[gitee.com/go-package/carbon](https://gitee.com/go-package/carbon "gitee.com/go-package/carbon")

#### Installation
```go
// By github
go get -u github.com/golang-module/carbon

import (
    "github.com/golang-module/carbon"
)

// By gitee
go get -u gitee.com/go-package/carbon

import (
    "gitee.com/go-package/carbon"
)               
```

#### Usage and example
> The default timezone is Local, assuming the current time is 2020-08-05 13:14:15

##### Yesterday, today and tomorrow
```go
// Datetime of today
fmt.Sprintf("%s", carbon.
w()) // 2020-08-05 13:14:15
carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
// Date of today
carbon.Now().ToDateString() // 2020-08-05
// Time of today
carbon.Now().ToTimeString() // 13:14:15
// Timestamp with second of today
carbon.Now().ToTimestamp() // 1596604455
carbon.Now().ToTimestampWithSecond() // 1596604455
// Timestamp with millisecond of today
carbon.Now().ToTimestampWithMillisecond() // 1596604455000
// Timestamp with microsecond of today
carbon.Now().ToTimestampWithMicrosecond() // 1596604455000000
// Timestamp with nanosecond of today
carbon.Now().ToTimestampWithNanosecond() // 1596604455000000000
// Datetime of today in given timezone
carbon.Now(Carbon.NewYork).ToDateTimeString() // 2020-08-05 14:14:15
carbon.SetTimezone(Carbon.NewYork).Now().ToDateTimeString() // 2020-08-05 14:14:15

// Datetime of yesterday 
fmt.Sprintf("%s", carbon.Yesterday()) // 2020-08-04 13:14:15
carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// Date of yesterday
carbon.Yesterday().ToDateString() // 2020-08-04
// Time of yesterday
carbon.Yesterday().ToTimeString() // 13:14:15
// Timestamp with second of yesterday
carbon.Yesterday().ToTimestamp() // 1596518055
carbon.Yesterday().ToTimestampWithSecond() // 1596518055
// Timestamp with millisecond of yesterday
carbon.Yesterday().ToTimestampWithMillisecond() // 1596518055000
// Timestamp with microsecond of yesterday
carbon.Yesterday().ToTimestampWithMicrosecond() // 1596518055000000
// Timestamp with nanosecond of yesterday
carbon.Yesterday().ToTimestampWithNanosecond() // 1596518055000000000
// Datetime of yesterday in given day
carbon.Parse("2021-01-28 13:14:15").Yesterday().ToDateTimeString() // 2021-01-27 13:14:15
// Datetime of yesterday in given timezone
carbon.Yesterday(Carbon.NewYork).ToDateTimeString() // 2020-08-04 14:14:15
carbon.SetTimezone(Carbon.NewYork).Yesterday().ToDateTimeString() // 2020-08-04 14:14:15

// Datetime of tomorrow
fmt.Sprintf("%s", carbon.Tomorrow()) // 2020-08-06 13:14:15
carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// Date of tomorrow
carbon.Tomorrow().ToDateString() // 2020-08-06
// Time of tomorrow
carbon.Tomorrow().ToTimeString() // 13:14:15
// Timestamp with second of tomorrow
carbon.Tomorrow().ToTimestamp() // 1596690855
carbon.Tomorrow().ToTimestampWithSecond() // 1596690855
// Timestamp with millisecond of tomorrow
carbon.Tomorrow().ToTimestampWithMillisecond() // 1596690855000
// Timestamp with microsecond of tomorrow
carbon.Tomorrow().ToTimestampWithMicrosecond() // 1596690855000000
// Timestamp with nanosecond of tomorrow
carbon.Tomorrow().ToTimestampWithNanosecond() // 1596690855000000000
// Datetime of tomorrow in given day
carbon.Parse("2021-01-28 13:14:15").Tomorrow().ToDateTimeString() // 2021-01-29 13:14:15
// Datetime of tomorrow in given timezone
carbon.Tomorrow(Carbon.NewYork).ToDateTimeString() // 2020-08-06 14:14:15
carbon.SetTimezone(Carbon.NewYork).Tomorrow().ToDateTimeString() // 2020-08-06 14:14:15
```

##### Create carbon instance
```go
// Create carbon instance from timestamp with second
carbon.CreateFromTimestamp(-1).ToDateTimeString() // 1970-01-01 07:59:59
carbon.CreateFromTimestamp(-1, carbon.Tokyo).ToDateTimeString() // 1970-01-01 08:59:59
carbon.CreateFromTimestamp(0).ToDateTimeString() // 1970-01-01 08:00:00
carbon.CreateFromTimestamp(0, carbon.Tokyo).ToDateTimeString() // 1970-01-01 09:00:00
carbon.CreateFromTimestamp(1596604455).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromTimestamp(1596604455, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// Create carbon instance from timestamp with millisecond
carbon.CreateFromTimestamp(1596604455000).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromTimestamp(1596604455000, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// Create carbon instance from timestamp with microsecond
carbon.CreateFromTimestamp(1596604455000000).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromTimestamp(1596604455000000, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// Create carbon instance from timestamp with nanosecond
carbon.CreateFromTimestamp(1596604455000000000).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromTimestamp(1596604455000000000, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15

// Create carbon instance from year, month, day, hour, minute and second
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// Create carbon instance from year, month and day
carbon.CreateFromDate(2020, 8, 5).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromDate(2020, 8, 5, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// Create carbon instance from hour, minute and second
carbon.CreateFromTime(13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromTime(13, 14, 15, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### Parse as carbon by standard string
```go
carbon.Parse("").ToDateTimeString() // empty string
carbon.Parse("0").ToDateTimeString() // empty string
carbon.Parse("0000-00-00 00:00:00").ToDateTimeString() // empty string
carbon.Parse("0000-00-00").ToDateTimeString() // empty string
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05").ToDateTimeString() // 2020-08-05 00:00:00
carbon.Parse("20200805131415").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("20200805").ToDateTimeString() // 2020-08-05 00:00:00
carbon.Parse("2020-08-05T13:14:15+08:00").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### Parse as carbon by format string
```go
carbon.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("2020-08-05 13:14:15", "Y-m-d H:i:s", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### Parse as carbon by layout string
```go
carbon.ParseByLayout("2020|08|05 13|14|15", "2006|01|02 15|04|05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("2020-08-05 13:14:15", "2006-01-02 15:04:05", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### Convert between carbon and time.Time
```go
// Time.time convert to Carbon
carbon.Time2Carbon(time.Now())
// Carbon convert to Time.time
carbon.Now().Carbon2Time()
```

##### Start and end
```go
// Start of the century
carbon.Parse("2020-08-05 13:14:15").StartOfCentury().ToDateTimeString() // 2000-01-01 00:00:00
// End of the century
carbon.Parse("2020-08-05 13:14:15").EndOfCentury().ToDateTimeString() // 2999-12-31 23:59:59

// Start of the decade
carbon.Parse("2020-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
carbon.Parse("2021-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
carbon.Parse("2029-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
// End of the decade
carbon.Parse("2020-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
carbon.Parse("2021-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
carbon.Parse("2029-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59

// Start of the year
carbon.Parse("2020-08-05 13:14:15").StartOfYear().ToDateTimeString() // 2020-01-01 00:00:00
// End of the year
carbon.Parse("2020-08-05 13:14:15").EndOfYear().ToDateTimeString() // 2020-12-31 23:59:59

// Start of the quarter
carbon.Parse("2020-08-05 13:14:15").StartOfQuarter().ToDateTimeString() // 2020-07-01 00:00:00
// End of the quarter
carbon.Parse("2020-08-05 13:14:15").EndOfQuarter().ToDateTimeString() // 2020-09-30 23:59:59

// Start of the month
carbon.Parse("2020-08-05 13:14:15").StartOfMonth().ToStartTimeString() // 2020-08-01 00:00:00
// End of the month
carbon.Parse("2020-08-05 13:14:15").EndOfMonth().ToDateTimeString() // 2020-08-31 23:59:59

// Start of the week
carbon.Parse("2020-08-05 13:14:15").StartOfWeek(time.Sunday).ToDateTimeString() // 2020-08-02 00:00:00
carbon.Parse("2020-08-05 13:14:15").StartOfWeek(time.Monday).ToDateTimeString() // 2020-08-03 00:00:00
// End of the week
carbon.Parse("2020-08-05 13:14:15").EndOfWeek(time.Sunday).ToDateTimeString() // 2020-08-08 23:59:59
carbon.Parse("2020-08-05 13:14:15").EndOfWeek(time.Monday).ToDateTimeString() // 2020-08-09 23:59:59

// Start of the day
carbon.Parse("2020-08-05 13:14:15").StartOfDay().ToDateTimeString() // 2020-08-05 00:00:00
// End of the day
carbon.Parse("2020-08-05 13:14:15").EndOfDay().ToDateTimeString() // 2020-08-05 23:59:59

// Start of the hour
carbon.Parse("2020-08-05 13:14:15").StartOfHour().ToDateTimeString() // 2020-08-05 13:00:00
// End of the hour
carbon.Parse("2020-08-05 13:14:15").EndOfHour().ToDateTimeString() // 2020-08-05 13:59:59

// Start of the minute
carbon.Parse("2020-08-05 13:14:15").StartOfMinute().ToDateTimeString() // 2020-08-05 13:14:00
// End of the minute
carbon.Parse("2020-08-05 13:14:15").EndOfMinute().ToDateTimeString() // 2020-08-05 13:14:59

// Start of the second
carbon.Parse("2020-08-05 13:14:15").StartOfSecond().Format("Y-m-d H:i:s.u") // 2020-08-05 13:14:15.0
// End of the second
carbon.Parse("2020-08-05 13:14:15").EndOfSecond().Format("Y-m-d H:i:s.u") // 2020-08-05 13:14:15.999
```

##### Addition and subtraction
```go
// Add three centuries
carbon.Parse("2020-02-29 13:14:15").AddCenturies(3).ToDateTimeString() // 2320-02-29 13:14:15
// Add three centuries with no overflow
carbon.Parse("2020-02-29 13:14:15").AddCenturiesNoOverflow(3).ToDateTimeString() // 2320-02-29 13:14:15

// Add one century
carbon.Parse("2020-02-29 13:14:15").AddCentury().ToDateTimeString() // 2120-02-29 13:14:15
// Add one century with no overflow
carbon.Parse("2020-02-29 13:14:15").AddCenturyNoOverflow().ToDateTimeString() // 2120-02-29 13:14:15

// Subtract three centuries
carbon.Parse("2020-02-29 13:14:15").SubCenturies(3).ToDateTimeString() // 1720-02-29 13:14:15
// Subtract three centuries with no overflow
carbon.Parse("2020-02-29 13:14:15").SubCenturiesNoOverflow(3).ToDateTimeString() // 1720-02-29 13:14:15

// Subtract one century
carbon.Parse("2020-02-29 13:14:15").SubCentury().ToDateTimeString() // 1920-02-29 13:14:15
// Subtract one century with no overflow
carbon.Parse("2020-02-29 13:14:15").SubCenturyNoOverflow().ToDateTimeString() // 1920-02-20 13:14:15

// Add three years
carbon.Parse("2020-02-29 13:14:15").AddYears(3).ToDateTimeString() // 2023-03-01 13:14:15
// Add three years with no overflow
carbon.Parse("2020-02-29 13:14:15").AddYearsNoOverflow(3).ToDateTimeString() // 2023-02-28 13:14:15

// Add one year
carbon.Parse("2020-02-29 13:14:15").AddYear().ToDateTimeString() // 2021-03-01 13:14:15
// Add one year with no overflow
carbon.Parse("2020-02-29 13:14:15").AddYearNoOverflow().ToDateTimeString() // 2021-02-28 13:14:15

// Subtract three years
carbon.Parse("2020-02-29 13:14:15").SubYears(3).ToDateTimeString() // 2017-03-01 13:14:15
// Subtract three years with no overflow
carbon.Parse("2020-02-29 13:14:15").SubYearsNoOverflow(3).ToDateTimeString() // 2017-02-28 13:14:15

// Subtract one year
carbon.Parse("2020-02-29 13:14:15").SubYear().ToDateTimeString() // 2019-03-01 13:14:15
// Subtract one year with no overflow
carbon.Parse("2020-02-29 13:14:15").SubYearNoOverflow().ToDateTimeString() // 2019-02-28 13:14:15

// Add three quarters
carbon.Parse("2019-08-31 13:14:15").AddQuarters(3).ToDateTimeString() // 2019-03-02 13:14:15
// Add three quarters with no overflow
carbon.Parse("2019-08-31 13:14:15").AddQuartersNoOverflow(3).ToDateTimeString() // 2019-02-29 13:14:15

// Add one quarter
carbon.Parse("2019-11-30 13:14:15").AddQuarter().ToDateTimeString() // 2020-03-01 13:14:15
// Add one quarter with no overflow
carbon.Parse("2019-11-30 13:14:15").AddQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// Subtract three quarters
carbon.Parse("2019-08-31 13:14:15").SubQuarters(3).ToDateTimeString() // 2019-03-03 13:14:15
// Subtract three quarters with no overflow
carbon.Parse("2019-08-31 13:14:15").SubQuartersNoOverflow(3).ToDateTimeString() // 2019-02-28 13:14:15

// Subtract one quarter
carbon.Parse("2020-05-31 13:14:15").SubQuarter().ToDateTimeString() // 2020-03-02 13:14:15
// Subtract one quarter with no overflow
carbon.Parse("2020-05-31 13:14:15").SubQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// Add three months
carbon.Parse("2020-02-29 13:14:15").AddMonths(3).ToDateTimeString() // 2020-05-29 13:14:15
// Add three months with no overflow
carbon.Parse("2020-02-29 13:14:15").AddMonthsNoOverflow(3).ToDateTimeString() // 2020-05-29 13:14:15

// Add one month
carbon.Parse("2020-01-31 13:14:15").AddMonth().ToDateTimeString() // 2020-03-02 13:14:15
// Add one month with no overflow
carbon.Parse("2020-01-31 13:14:15").AddMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// Subtract three months
carbon.Parse("2020-02-29 13:14:15").SubMonths(3).ToDateTimeString() // 2019-11-29 13:14:15
// Subtract three months with no overflow
carbon.Parse("2020-02-29 13:14:15").SubMonthsNoOverflow(3).ToDateTimeString() // 2019-11-29 13:14:15

// Subtract one month
carbon.Parse("2020-03-31 13:14:15").SubMonth().ToDateTimeString() // 2020-03-02 13:14:15
// Subtract one month with no overflow
carbon.Parse("2020-03-31 13:14:15").SubMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// Add three weeks
carbon.Parse("2020-02-29 13:14:15").AddWeeks(3).ToDateTimeString() // 2020-03-21 13:14:15
// Add one week
carbon.Parse("2020-02-29 13:14:15").AddWeek().ToDateTimeString() // 2020-03-07 13:14:15

// Subtract three weeks
carbon.Parse("2020-02-29 13:14:15").SubWeeks(3).ToDateTimeString() // 2020-02-08 13:14:15
// Subtract three week
carbon.Parse("2020-02-29 13:14:15").SubWeek().ToDateTimeString() // 2020-02-22 13:14:15

// Add three days
carbon.Parse("2020-08-05 13:14:15").AddDays(3).ToDateTimeString() // 2020-08-08 13:14:15
// Add one day
carbon.Parse("2020-08-05 13:14:15").AddDay().ToDateTimeString() // 2020-08-05 13:14:15

// Subtract three days
carbon.Parse("2020-08-05 13:14:15").SubDays(3).ToDateTimeString() // 2020-08-02 13:14:15
// Subtract one day
carbon.Parse("2020-08-05 13:14:15").SubDay().ToDateTimeString() // 2020-08-04 13:14:15

// Add three hours
carbon.Parse("2020-08-05 13:14:15").AddHours(3).ToDateTimeString() // 2020-08-05 16:14:15
// Add two and a half hours
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5h").ToDateTimeString() // 2020-08-05 15:44:15
carbon.Parse("2020-08-05 13:14:15").AddDuration("2h30m").ToDateTimeString() // 2020-08-05 15:44:15
// Add one hour
carbon.Parse("2020-08-05 13:14:15").AddHour().ToDateTimeString() // 2020-08-05 14:14:15

// Subtract three hours
carbon.Parse("2020-08-05 13:14:15").SubHours(3).ToDateTimeString() // 2020-08-05 10:14:15
// Subtract two and a half hours
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5h").ToDateTimeString() // 2020-08-05 10:44:15
carbon.Parse("2020-08-05 13:14:15").SubDuration("2h30m").ToDateTimeString() // 2020-08-05 10:44:15
// Subtract one hour
carbon.Parse("2020-08-05 13:14:15").SubHour().ToDateTimeString() // 2020-08-05 12:14:15

// Add three minutes
carbon.Parse("2020-08-05 13:14:15").AddMinutes(3).ToDateTimeString() // 2020-08-05 13:17:15
// Add two and a half minutes
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5m").ToDateTimeString() // 2020-08-05 13:16:45
carbon.Parse("2020-08-05 13:14:15").AddDuration("2m30s").ToDateTimeString() // 2020-08-05 13:16:45
// Add one minute
carbon.Parse("2020-08-05 13:14:15").AddMinute().ToDateTimeString() // 2020-08-05 13:15:15

// Subtract three minutes
carbon.Parse("2020-08-05 13:14:15").SubMinutes(3).ToDateTimeString() // 2020-08-05 13:11:15
// Subtract two and a half minutes
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5m").ToDateTimeString() // 2020-08-05 13:11:45
// Subtract one minute
carbon.Parse("2020-08-05 13:14:15").SubMinute().ToDateTimeString() // 2020-08-05 13:13:15

// Add three seconds
carbon.Parse("2020-08-05 13:14:15").AddSeconds(3).ToDateTimeString() // 2020-08-05 13:14:18
// Add two and a half seconds
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:17
// Add one second
carbon.Parse("2020-08-05 13:14:15").AddSecond().ToDateTimeString() // 2020-08-05 13:14:16

// Subtract three seconds
carbon.Parse("2020-08-05 13:14:15").SubSeconds(3).ToDateTimeString() // 2020-08-05 13:14:12
// Subtract two and a half seconds
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:12
// Subtract one second
carbon.Parse("2020-08-05 13:14:15").SubSecond().ToDateTimeString() // 2020-08-05 13:14:14
```

##### Difference
```go
// Difference in weeks
carbon.Parse("2020-08-05 13:14:15").DiffInWeeks(carbon.Parse("2020-07-28 13:14:15")) // -1
// Difference in weeks with absolute value
carbon.Parse("2020-08-05 13:14:15").DiffInWeeksWithAbs(carbon.Parse("2020-07-28 13:14:15")) // 1

// Difference in days
carbon.Parse("2020-08-05 13:14:15").DiffInDays(carbon.Parse("2020-08-04 13:14:15")) // -1
// Difference in days with absolute value
carbon.Parse("2020-08-05 13:14:15").DiffInDaysWithAbs(carbon.Parse("2020-08-04 13:14:15")) // 1

// Difference in hours
carbon.Parse("2020-08-05 13:14:15").DiffInHours(carbon.Parse("2020-08-05 12:14:15")) // -1
// Difference in hours with absolute value
carbon.Parse("2020-08-05 13:14:15").DiffInHoursWithAbs(carbon.Parse("2020-08-05 12:14:15")) // 1

// Difference in minutes
carbon.Parse("2020-08-05 13:14:15").DiffInMinutes(carbon.Parse("2020-08-05 13:13:15")) // -1
// Difference in minutes with absolute value
carbon.Parse("2020-08-05 13:14:15").DiffInMinutesWithAbs(carbon.Parse("2020-08-05 13:13:15")) // 1

// Difference in seconds
carbon.Parse("2020-08-05 13:14:15").DiffInSeconds(carbon.Parse("2020-08-05 13:14:14")) // -1
// Difference in seconds with absolute value
carbon.Parse("2020-08-05 13:14:15").DiffInSecondsWithAbs(carbon.Parse("2020-08-05 13:14:14")) // 1

// Difference for humans from now time
carbon.Parse("2020-08-05 13:14:15").DiffForHumans()) // just now
carbon.Parse("2019-08-05 13:14:15").DiffForHumans() // 1 year ago
carbon.Parse("2018-08-05 13:14:15").DiffForHumans() // 2 years ago
carbon.Parse("2021-08-05 13:14:15").DiffForHumans() // 1 year from now
carbon.Parse("2022-08-05 13:14:15").DiffForHumans() // 2 years from now
// Difference for humans from another time
carbon.Parse("2020-08-05 13:14:15").DiffForHumans(carbon.Now()) // 1 year before
carbon.Parse("2019-08-05 13:14:15").DiffForHumans(carbon.Now()) // 2 years before
carbon.Parse("2018-08-05 13:14:15").DiffForHumans(carbon.Now()) // 1 year after
carbon.Parse("2022-08-05 13:14:15").DiffForHumans(carbon.Now()) // 2 years after
```

##### Comparison
```go
// Whether is zero time
carbon.Parse("").IsZero() // true
carbon.Parse("0").IsZero() // true
carbon.Parse("0000-00-00 00:00:00").IsZero() // true
carbon.Parse("0000-00-00").IsZero() // true
carbon.Parse("00:00:00").IsZero() // true
carbon.Parse("2020-08-05 00:00:00").IsZero() // false
carbon.Parse("2020-08-05").IsZero() // false
carbon.Parse("2020-08-05").SetTimezone("xxx").IsZero() // false

// Whether is invalid time
carbon.Parse("").IsInvalid() // true
carbon.Parse("0").IsInvalid() // true
carbon.Parse("0000-00-00 00:00:00").IsInvalid() // true
carbon.Parse("0000-00-00").IsInvalid() // true
carbon.Parse("00:00:00").IsInvalid() // true
carbon.Parse("2020-08-05 00:00:00").IsInvalid() // false
carbon.Parse("2020-08-05").IsInvalid() // false
carbon.Parse("2020-08-05").SetTimezone("xxx").IsInvalid() // true

// Whether is now time
carbon.Now().IsNow() // true
// Whether is future time
carbon.Tomorrow().IsFuture() // true
// Whether is pass time
carbon.Yesterday().IsPast() // true

// Whether is leap year
carbon.Parse("2020-08-05 13:14:15").IsLeapYear() // true
// Whether is long year
carbon.Parse("2020-08-05 13:14:15").IsLongYear() // true

// Whether is january 
carbon.Parse("2020-08-05 13:14:15").IsJanuary() // false
// Whether is february
carbon.Parse("2020-08-05 13:14:15").IsFebruary() // false
// Whether is march
carbon.Parse("2020-08-05 13:14:15").IsMarch() // false
// Whether is april
carbon.Parse("2020-08-05 13:14:15").IsApril()  // false
// Whether is may
carbon.Parse("2020-08-05 13:14:15").IsMay() // false
// Whether is june
carbon.Parse("2020-08-05 13:14:15").IsJune() // false
// Whether is july
carbon.Parse("2020-08-05 13:14:15").IsJuly() // false
// Whether is august
carbon.Parse("2020-08-05 13:14:15").IsAugust() // false
// Whether is september
carbon.Parse("2020-08-05 13:14:15").IsSeptember() // true
// Whether is october
carbon.Parse("2020-08-05 13:14:15").IsOctober() // false
// Whether is november
carbon.Parse("2020-08-05 13:14:15").IsNovember() // false
// Whether is december
carbon.Parse("2020-08-05 13:14:15").IsDecember() // false

// Whether is monday
carbon.Parse("2020-08-05 13:14:15").IsMonday() // false
// Whether is tuesday
carbon.Parse("2020-08-05 13:14:15").IsTuesday() // true
// Whether is wednesday
carbon.Parse("2020-08-05 13:14:15").IsWednesday() // false
// Whether is thursday
carbon.Parse("2020-08-05 13:14:15").IsThursday()  // false
// Whether is friday
carbon.Parse("2020-08-05 13:14:15").IsFriday() // false
// Whether is saturday
carbon.Parse("2020-08-05 13:14:15").IsSaturday() // false
// Whether is sunday
carbon.Parse("2020-08-05 13:14:15").IsSunday() // false
// Whether is weekday
carbon.Parse("2020-08-05 13:14:15").IsWeekday() // false
// Whether is weekend
carbon.Parse("2020-08-05 13:14:15").IsWeekend() // true

// Whether is yesterday
carbon.Parse("2020-08-04 13:14:15").IsYesterday() // true
carbon.Parse("2020-08-04 00:00:00").IsYesterday() // true
carbon.Parse("2020-08-04").IsYesterday() // true
// Whether is today
carbon.Parse("2020-08-05 13:14:15").IsToday() // true
carbon.Parse("2020-08-05 00:00:00").IsToday() // true
carbon.Parse("2020-08-05").IsToday() // true
// Whether is tomorrow
carbon.Parse("2020-08-06 13:14:15").IsTomorrow() // true
carbon.Parse("2020-08-06 00:00:00").IsTomorrow() // true
carbon.Parse("2020-08-06").IsTomorrow() // true

// Whether greater than
carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-05 13:14:15")) // false

// Whether less than
carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-05 13:14:15")) // false

// Whether equal
carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:00")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:00")) // false

// Whether not equal
carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("!=", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<>", carbon.Parse("2020-08-05 13:14:15")) // false

// Whether greater than or equal
carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-05 13:14:15")) // true

// Whether less than or equal
carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-05 13:14:15")) // true

// Whether between
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// Whether between included start
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// Whether between included end 
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// Whether between included both
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
```
> For the definition of long year, please refer to https://en.wikipedia.org/wiki/ISO_8601#Week_dates

##### Output   
```go
// Output timestamp with second, ToTimestamp() is short for ToTimestampWithSecond()
carbon.Parse("2020-08-05 13:14:15").ToTimestamp() // 1596604455
carbon.Parse("2020-08-05 13:14:15").ToTimestampWithSecond() // 1596604455
// Output timestamp with millisecond
carbon.Parse("2020-08-05 13:14:15").ToTimestampWithMillisecond() // 1596604455000
// Output timestamp with microsecond
carbon.Parse("2020-08-05 13:14:15").ToTimestampWithMicrosecond() // 1596604455000000
// Output timestamp with nanosecond
carbon.Parse("2020-08-05 13:14:15").ToTimestampWithNanosecond() // 1596604455000000000

// Output datetime format string
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString(carbon.Tokyo) // 2020-08-05 14:14:15
// Output short datetime format string
carbon.Parse("2020-08-05 13:14:15").ToShortDateTimeString() // 20200805131415
carbon.Parse("2020-08-05 13:14:15").ToShortDateTimeString(carbon.Tokyo) // 20200805141415

// Output date format string
carbon.Parse("2020-08-05 13:14:15").ToDateString() // 2020-08-05
carbon.Parse("2020-08-05 13:14:15").ToDateString(carbon.Tokyo) // 2020-08-05
// Output short date format string
carbon.Parse("2020-08-05 13:14:15").ToShortDateString() // 20200805
carbon.Parse("2020-08-05 13:14:15").ToShortDateString(carbon.Tokyo) // 20200805

// Output time format string
carbon.Parse("2020-08-05 13:14:15").ToTimeString() // 13:14:15
carbon.Parse("2020-08-05 13:14:15").ToTimeString(carbon.Tokyo) // 14:14:15
// Output short time format string
carbon.Parse("2020-08-05 13:14:15").ToShortTimeString() // 131415
carbon.Parse("2020-08-05 13:14:15").ToShortTimeString(carbon.Tokyo) // 141415

// Output Ansic format string
carbon.Parse("2020-08-05 13:14:15").ToAnsicString() // Wed Aug  5 13:14:15 2020
carbon.Parse("2020-08-05 13:14:15").ToAnsicString(carbon.Tokyo) // Wed Aug  5 14:14:15 2020
// Output Atom format string
carbon.Parse("2020-08-05 13:14:15").ToAtomString() // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").ToAtomString(carbon.Tokyo) // 2020-08-05T14:14:15+08:00
// Output UnixDate format string
carbon.Parse("2020-08-05 13:14:15").ToUnixDateString() // Wed Aug  5 13:14:15 CST 2020
carbon.Parse("2020-08-05 13:14:15").ToUnixDateString(carbon.Tokyo) // Wed Aug  5 14:14:15 JST 2020
// Output RubyDate format string
carbon.Parse("2020-08-05 13:14:15").ToRubyDateString() // Wed Aug 05 13:14:15 +0800 2020
carbon.Parse("2020-08-05 13:14:15").ToRubyDateString(carbon.Tokyo) // Wed Aug 05 14:14:15 +0900 2020
// Output Kitchen format string
carbon.Parse("2020-08-05 13:14:15").ToKitchenString() // 1:14PM
carbon.Parse("2020-08-05 13:14:15").ToKitchenString(carbon.Tokyo) // 2:14PM
// Output Cookie format string
carbon.Parse("2020-08-05 13:14:15").ToCookieString() // Wednesday, 05-Aug-2020 13:14:15 CST
carbon.Parse("2020-08-05 13:14:15").ToCookieString(carbon.Tokyo) // Wednesday, 05-Aug-2020 14:14:15 JST
// Output DayDateTime format string
carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString() // Wed, Aug 5, 2020 1:14 PM
carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString(carbon.Tokyo) // Wed, Aug 5, 2020 2:14 PM
// Output RSS format string
carbon.Parse("2020-08-05 13:14:15").ToRssString() // Wed, 05 Aug 2020 13:14:15 +0800
carbon.Parse("2020-08-05 13:14:15").ToRssString(carbon.Tokyo) // Wed, 05 Aug 2020 14:14:15 +0900
// Output W3C format string
carbon.Parse("2020-08-05 13:14:15").ToW3cString() // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").ToW3cString(carbon.Tokyo) // 2020-08-05T14:14:15+09:00

// Output ISO8601 format string
carbon.Parse("2020-08-05 13:14:15").ToIso8601String() // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").ToIso8601String(carbon.Tokyo) // 2020-08-05T14:14:15+09:00
// Output RFC822 format string
carbon.Parse("2020-08-05 13:14:15").ToRfc822String() // 05 Aug 20 13:14 CST
carbon.Parse("2020-08-05 13:14:15").ToRfc822String(carbon.Tokyo) // 05 Aug 20 14:14 JST
// Output RFC822Z format string
carbon.Parse("2020-08-05 13:14:15").ToRfc822zString() // 05 Aug 20 13:14 +0800
carbon.Parse("2020-08-05 13:14:15").ToRfc822zString(carbon.Tokyo) // 05 Aug 20 14:14 +0900
// Output RFC850 format string
carbon.Parse("2020-08-05 13:14:15").ToRfc850String() // Wednesday, 05-Aug-20 13:14:15 CST
carbon.Parse("2020-08-05 13:14:15").ToRfc850String(carbon.Tokyo) // Wednesday, 05-Aug-20 14:14:15 JST
// Output RFC1036 format string
carbon.Parse("2020-08-05 13:14:15").ToRfc1036String() // Wed, 05 Aug 20 13:14:15 +0800
carbon.Parse("2020-08-05 13:14:15").ToRfc1036String(carbon.Tokyo) // Wed, 05 Aug 20 14:14:15 +0900
// Output RFC1123 format string
carbon.Parse("2020-08-05 13:14:15").ToRfc1123String() // Wed, 05 Aug 2020 13:14:15 CST
carbon.Parse("2020-08-05 13:14:15").ToRfc1123String(carbon.Tokyo) // Wed, 05 Aug 2020 14:14:15 JST
// Output RFC2822 format string
carbon.Parse("2020-08-05 13:14:15").ToRfc2822String() // Wed, 05 Aug 2020 13:14:15 +0800
carbon.Parse("2020-08-05 13:14:15").ToRfc2822String(carbon.Tokyo) // Wed, 05 Aug 2020 14:14:15 +0900
// Output RFC3339 format string
carbon.Parse("2020-08-05 13:14:15").ToRfc3339String() // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").ToRfc3339String(carbon.Tokyo) // 2020-08-05T14:14:15+09:00
// Output RFC7231 format string
carbon.Parse("2020-08-05 13:14:15").ToRfc7231String() // Wed, 05 Aug 2020 13:14:15 GMT
carbon.Parse("2020-08-05 13:14:15").ToRfc7231String(carbon.Tokyo) // Wed, 05 Aug 2020 14:14:15 GMT

// Output string
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-05 13:14:15").ToString(carbon.Tokyo) // 2020-08-05 14:14:15 +0900 JST

// Output string by layout, Layout() is short for ToLayoutString()
carbon.Parse("2020-08-05 13:14:15").Layout("20060102150405") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Layout("2006年01月02日 15时04分05秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Layout("It is 2006-01-02 15:04:05") // It is 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15").Layout("2006-01-02 15:04:05", carbon.Tokyo) // 2020-08-05 14:14:15

// Output string by format, Format() is short for ToFormatString()
carbon.Parse("2020-08-05 13:14:15").Format("YmdHis") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Format("Y年m月d日 H时i分s秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Format("l jS \\o\\f F Y h:i:s A") // Wednesday 5th of August 2020 01:14:15 PM
carbon.Parse("2020-08-05 13:14:15").Format("\\I\\t \\i\\s Y-m-d H:i:s") // It is 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15").Format("Y-m-d H:i:s", carbon.Tokyo) // 2020-08-05 14:14:15
```
> For more supported format signs, please see the <a href="#format-sign-table">Format sign table</a>

##### Getter
```go
// Get total days of the year
carbon.Parse("2019-08-05 13:14:15").DaysInYear() // 365
carbon.Parse("2020-08-05 13:14:15").DaysInYear() // 366
// Get total days of the month
carbon.Parse("2020-02-01 13:14:15").DaysInMonth() // 29
carbon.Parse("2020-04-01 13:14:15").DaysInMonth() // 30
carbon.Parse("2020-08-01 13:14:15").DaysInMonth() // 31

// Get day of the year (start with 1)
carbon.Parse("2020-08-05 13:14:15").DayOfYear() // 218
// Get week of the year (start with 1)
carbon.Parse("2020-08-05 13:14:15").WeekOfYear() // 32
// Get day of the month (start with 1)
carbon.Parse("2020-08-05 13:14:15").DayOfMonth() // 5
// Get week of the month (start with 1)
carbon.Parse("2020-08-05 13:14:15").WeekOfMonth() // 1
// Get day of the week (start with 1)
carbon.Parse("2020-08-05 13:14:15").DayOfWeek() // 3

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
// Get current week(start with 0)
carbon.Parse("2020-08-05 13:14:15").Week() // 3
carbon.Parse("2020-08-05 13:14:15").Week() // 3
// Get current day
carbon.Parse("2020-08-05 13:14:15").Day() // 5
// Get current hour
carbon.Parse("2020-08-05 13:14:15").Hour() // 13
// Get current minute
carbon.Parse("2020-08-05 13:14:15").Minute() // 14
// Get current second
carbon.Parse("2020-08-05 13:14:15").Second() // 15
// Get current millisecond
carbon.Parse("2020-08-05 13:14:15").Millisecond() // 1596604455000
// Get current microsecond
carbon.Parse("2020-08-05 13:14:15").Microsecond() // 1596604455000000
// Get current nanosecond
carbon.Parse("2020-08-05 13:14:15").Nanosecond() // 1596604455000000000

// Get timezone name
carbon.SetTimezone(carbon.PRC).Timezone() // CST
carbon.SetTimezone(carbon.Tokyo).Timezone() // JST

// Get location name
carbon.SetTimezone(carbon.PRC).Location() // PRC
carbon.SetTimezone(carbon.Tokyo).Location() // Asia/Tokyo

// Get offset seconds from the UTC timezone
carbon.SetTimezone(carbon.PRC).Offset() // 28800
carbon.SetTimezone(carbon.Tokyo).Offset() // 32400

// Get locale name
carbon.Now().SetLocale("en").Locale() // en
carbon.Now().SetLocale("zh-CN").Locale() // zh-CN

// Get constellation name
carbon.Now().Constellation() // Leo
carbon.Now().SetLocale("en").Constellation() // Leo
carbon.Now().SetLocale("zh-CN").Constellation() // 狮子座

//Get season name
carbon.Now().Season() // Summer
carbon.Now().SetLocale("en").Season() // Summer
carbon.Now().SetLocale("zh-CN").Season() // 夏季

// Get current age
carbon.Parse("2002-01-01 13:14:15").Age() // 17
carbon.Parse("2002-12-31 13:14:15").Age() // 18
```
> If you have any doubt about 'WeekOfYear()' and 'WeekOfMonth()', please refer to [ISO8601 standard](https://baike.baidu.com/item/ISO%208601/3910715#3)

##### Setter
```go
// Set timezone
carbon.SetTimezone(carbon.PRC).Now().ToDateTimeString() // 2020-08-05 13:14:15
carbon.SetTimezone(carbon.Tokyo).Now().ToDateTimeString() // 2020-08-05 14:14:15
carbon.SetTimezone(carbon.Tokyo).Now().SetTimezone(carbon.PRC).ToDateTimeString() // 2020-08-05 12:14:15

// Set locale
carbon.Parse("2020-07-05 13:14:15").SetLocale("en").DiffForHumans()) // 1 month before
carbon.Parse("2020-07-05 13:14:15").SetLocale("zh-CN").DiffForHumans()) // 1 月前

// Set year
carbon.Parse("2019-08-05").SetYear(2020).ToDateString() // 2020-08-05
carbon.Parse("2020-02-29").SetYear(2019).ToDateString() // 2019-03-01

// Set month
carbon.Parse("2020-01-31").SetMonth(2).ToDateString() // 2020-03-02
carbon.Parse("2020-08-05").SetMonth(2).ToDateString() // 2020-02-05

// Set day
carbon.Parse("2019-08-05").SetDay(31).ToDateString() // 2020-08-31
carbon.Parse("2020-02-01").SetDay(31).ToDateString() // 2020-03-02

// Set hour
carbon.Parse("2020-08-05 13:14:15").SetHour(10).ToDateTimeString() // 2020-08-05 10:14:15
carbon.Parse("2020-08-05 13:14:15").SetHour(24).ToDateTimeString() // 2020-08-06 00:14:15

// Set minute
carbon.Parse("2020-08-05 13:14:15").SetMinute(10).ToDateTimeString() // 2020-08-05 13:10:15
carbon.Parse("2020-08-05 13:14:15").SetMinute(60).ToDateTimeString() // 2020-08-05 14:00:15

// Set second
carbon.Parse("2020-08-05 13:14:15").SetSecond(10).ToDateTimeString() // 2020-08-05 13:14:10
carbon.Parse("2020-08-05 13:14:15").SetSecond(60).ToDateTimeString() // 2020-08-05 13:15:00
```

##### Constellation
```go
// Get constellation name
carbon.Parse("2020-08-05 13:14:15").Constellation() // Leo

// Whether is aries
carbon.Parse("2020-08-05 13:14:15").IsAries() // false
// Whether is taurus
carbon.Parse("2020-08-05 13:14:15").IsTaurus() // false
// Whether is gemini
carbon.Parse("2020-08-05 13:14:15").IsGemini() // false
// Whether is cancer
carbon.Parse("2020-08-05 13:14:15").IsCancer() // false
// Whether is leo
carbon.Parse("2020-08-05 13:14:15").IsLeo() // true
// Whether is virgo
carbon.Parse("2020-08-05 13:14:15").IsVirgo() // false
// Whether is libra
carbon.Parse("2020-08-05 13:14:15").IsLibra() // false
// Whether is scorpio
carbon.Parse("2020-08-05 13:14:15").IsScorpio() // false
// Whether is sagittarius
carbon.Parse("2020-08-05 13:14:15").IsSagittarius() // false
// Whether is capricorn
carbon.Parse("2020-08-05 13:14:15").IsCapricorn() // false
// Whether is aquarius
carbon.Parse("2020-08-05 13:14:15").IsAquarius() // false
// Whether is pisces
carbon.Parse("2020-08-05 13:14:15").IsPisces() // false
```

##### Season
> According to the meteorological division, March to May is spring, June to August is summer, September to November is autumn, and December to February is winter
```go
// Get season name
carbon.Parse("2020-08-05 13:14:15").Season() // Summer

// Start of the season
carbon.Parse("2020-08-05 13:14:15").StartOfSeason().ToDateTimeString() // 2020-06-01 00:00:00
// End of the season
carbon.Parse("2020-08-05 13:14:15").EndOfSeason().ToDateTimeString() // 2020-08-31 23:59:59

// Whether is spring
carbon.Parse("2020-08-05 13:14:15").IsSpring() // false
// Whether is summer
carbon.Parse("2020-08-05 13:14:15").IsSummer() // true
// Whether is autumn
carbon.Parse("2020-08-05 13:14:15").IsAutumn() // false
// Whether is winter
carbon.Parse("2020-08-05 13:14:15").IsWinter() // false
```

##### Lunar
> Currently only `200` years from `1900` to `2100` are supported
```go
// Get year of animal
carbon.Parse("2020-08-05 13:14:15").Lunar().Animal() // 鼠

// Get festival of lunar
carbon.Parse("2021-02-12 13:14:15").Lunar().Festival() // 春节

// Get year of lunar
carbon.Parse("2020-08-05 13:14:15").Lunar().Year() // 2020
// Get month of lunar
carbon.Parse("2020-08-05 13:14:15").Lunar().Month() // 6
// Get leap month of lunar
carbon.Parse("2020-08-05 13:14:15").Lunar().LeapMonth() // 4
// Get day of lunar
carbon.Parse("2020-08-05 13:14:15").Lunar().Day() // 16

// Get year in chinese
carbon.Parse("2020-08-05 13:14:15").Lunar().ToYearString() // 二零二零
// Get month in chinese
carbon.Parse("2020-08-05 13:14:15").Lunar().ToMonthString() // 六
// Get day in chinese
carbon.Parse("2020-08-05 13:14:15").Lunar().ToDayString() // 十六
// Get full string in chinese
fmt.Sprintf("%s", carbon.Parse("2020-08-05 13:14:15").Lunar()) // 二零二零年六月十六
carbon.Parse("2020-08-05 13:14:15").Lunar().ToString() // 二零二零年六月十六

// Whether is leap year
carbon.Parse("2020-08-05 13:14:15").Lunar().IsLeapYear() // true
// Whether is leap month
carbon.Parse("2020-08-05 13:14:15").Lunar().IsLeapMonth() // false

// Whether is year of the rat
carbon.Parse("2020-08-05 13:14:15").Lunar().IsRatYear() // true
// Whether is year of the ox
carbon.Parse("2020-08-05 13:14:15").Lunar().IsOxYear() // false
// Whether is year of the tiger
carbon.Parse("2020-08-05 13:14:15").Lunar().IsTigerYear() // false
// Whether is year of the rabbit
carbon.Parse("2020-08-05 13:14:15").Lunar().IsRabbitYear() // false
// Whether is year of the dragon
carbon.Parse("2020-08-05 13:14:15").Lunar().IsDragonYear() // false
// Whether is year of the snake
carbon.Parse("2020-08-05 13:14:15").Lunar().IsSnakeYear() // false
// Whether is year of the horse
carbon.Parse("2020-08-05 13:14:15").Lunar().IsHorseYear() // false
// Whether is year of the goat
carbon.Parse("2020-08-05 13:14:15").Lunar().IsGoatYear() // false
// Whether is year of the monkey
carbon.Parse("2020-08-05 13:14:15").Lunar().IsMonkeyYear() // false
// Whether is year of the rooster
carbon.Parse("2020-08-05 13:14:15").Lunar().IsRoosterYear() // false
// Whether is year of the dog
carbon.Parse("2020-08-05 13:14:15").Lunar().IsDogYear() // false
// Whether is year of the dig
carbon.Parse("2020-08-05 13:14:15").Lunar().IsPigYear() // false
```

##### JSON handling

###### Define model
```go
type Person struct {
    ID  int64  `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
    Birthday carbon.ToDateTimeString `json:"birthday"`
    GraduatedAt carbon.ToDateString `json:"graduated_at"`
    CreatedAt carbon.ToTimeString `json:"created_at"`
    UpdatedAt carbon.ToTimestamp `json:"updated_at"`
    DateTime1 carbon.ToTimestampWithSecond `json:"date_time1"`
    DateTime2 carbon.ToTimestampWithMillisecond `json:"date_time2"`
    DateTime3 carbon.ToTimestampWithMicrosecond `json:"date_time3"`
    DateTime4 carbon.ToTimestampWithNanosecond `json:"date_time4"`
}
```

###### Instantiate model
```go
person := Person {
	ID:          1,
	Name:        "gouguoyin",
	Age:         18,
	Birthday:    ToDateTimeString{Now().SubYears(18)},
	GraduatedAt: ToDateString{Parse("2020-08-05 13:14:15")},
	CreatedAt:   ToTimeString{Parse("2021-08-05 13:14:15")},
	UpdatedAt:   ToTimestamp{Parse("2022-08-05 13:14:15")},
	DateTime1:   ToTimestampWithSecond{Parse("2023-08-05 13:14:15")},
	DateTime2:   ToTimestampWithMillisecond{Parse("2024-08-05 13:14:15")},
	DateTime3:   ToTimestampWithMicrosecond{Parse("2025-08-05 13:14:15")},
	DateTime4:   ToTimestampWithNanosecond{Parse("2025-08-05 13:14:15")},
}
```

###### JSON encode
```go
data, err := json.Marshal(&person)
if err != nil {
	t.Fatal(err)
}
fmt.Printf("%s",data)
// output
{
	"id":1,
	"name":"gouguoyin",
	"age":18,
	"birthday":"2003-07-16 16:22:02",
	"graduated_at":"2020-08-05",
	"created_at":"13:14:15",
	"updated_at":1659676455,
	"date_time1":1691212455,
	"date_time2":1722834855000,
	"date_time3":1754370855000000,
	"date_time4":1754370855000000000
}
```

###### JSON decode
```go
str := `{
	"id":1,
	"name":"gouguoyin",
	"age":18,
	"birthday":"2003-07-16 16:22:02",
	"graduated_at":"2020-08-05",
	"created_at":"13:14:15",
	"updated_at":1659676455,
	"date_time1":1691212455,
	"date_time2":1722834855000,
	"date_time3":1754370855000000,
	"date_time4":1754370855000000000
}`
person := new(Person)
err := json.Unmarshal([]byte(str), &person)
if err != nil {
	t.Fatal(err)
}
fmt.Printf("%+v", *person)
// output
{ID:1 Name:gouguoyin Age:18 Birthday:2003-07-16 16:22:02 GraduatedAt:2020-08-05 00:00:00 CreatedAt:0000-01-01 13:14:15 UpdatedAt:2022-08-05 13:14:15 DTime1:2023-08-05 13:14:15 DateTime2:2024-08-05 13:14:15 DateTime3:2025-08-05 13:14:15 DateTime4:2025-08-05 13:14:15}
```

##### I18n
> If you need to use i18n, please copy Lang directory to project directory first

The following languages are supported
* [simplified Chinese(zh-CN)](./lang/zh-CN.json "simplified Chinese")
* [traditional Chinese(zh-TW)](./lang/zh-TW.json "traditional Chinese")
* [English(en)](./lang/en.json "English")
* [Japanese(jp)](./lang/jp.json "Japanese")
* [Korean(kr)](./lang/kr.json "Korean")

The following methods are supported
* `DiffForHumans()`：to string difference in human friendly readable format
* `ToMonthString()`：to string of month
* `ToShortMonthString()`：to string of short month
* `ToWeekString()`：to string of week
* `ToShortWeekString()`：to string of short week
* `Constellation()`：get constellation name
* `Season()`：get season name

###### Set locale
```go
lang := NewLanguage()
if err := lang.SetLocale("zh-CN");err != nil {
    // Error handle...
    log.Fatal(err)
}

c := carbon.SetLanguage(lang)
c.Now().AddHours(1).DiffForHumans() // 1 小时后
c.Now().AddHours(1).ToMonthString() // 八月
c.Now().AddHours(1).ToShortMonthString() // 8月
c.Now().AddHours(1).ToWeekString() // 星期二
c.Now().AddHours(1).ToShortWeekString() // 周二
c.Now().AddHours(1).Constellation() // 狮子座
c.Now().AddHours(1).Season() // 夏季
```

###### Set dir
```go
lang := NewLanguage()
if err := lang.SetDir("lang");err != nil {
    // Error handle...
    log.Fatal(err)
}

c := carbon.SetLanguage(lang)
Now().AddHours(1).DiffForHumans() // 1 hour from now
Now().AddHours(1).ToMonthString() // August
Now().AddHours(1).ToShortMonthString() // Aug
Now().AddHours(1).ToWeekString() // Tuesday
Now().AddHours(1).ToShortWeekString() // Tue
Now().AddHours(1).Constellation() // Leo
Now().AddHours(1).Season() // Summer
```

###### Set some resources(the rest still translate from the given locale)
```go
lang := NewLanguage()

if err := lang.SetLocale("en");err != nil {
	// Error handle...
    log.Fatal(err)
}

resources := map[string]string {
    "hour":"%dh",
}
lang.SetResources(resources)

c := carbon.SetLanguage(lang)
c.Now().AddYears(1).DiffForHumans() // 1 year from now
c.Now().AddHours(1).DiffForHumans() // 1h from now
c.Now().ToMonthString() // August
c.Now().ToShortMonthString() // Aug
c.Now().ToWeekString() // Tuesday
c.Now().ToShortWeekString() // Tue
c.Now().Constellation() // Leo
c.Now().Season() // Summer
```

###### Set all resources
```go
lang := NewLanguage()
resources := map[string]string {
    "seasons": "Spring|Summer|Autumn|Winter",
    "months": "January|February|March|April|May|June|July|August|September|October|November|December",
    "months_short": "Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec",
    "weeks": "Sunday|Monday|Tuesday|Wednesday|Thursday|Friday|Saturday",
    "weeks_short": "Sun|Mon|Tue|Wed|Thu|Fri|Sat",
    "constellations": "Aries|Taurus|Gemini|Cancer|Leo|Virgo|Libra|Scorpio|Sagittarius|Capricornus|Aquarius|Pisce",
    "year":"1 yr|%d yrs",
    "month":"1 mo|%d mos",
    "week":"%dw",
    "day":"%dd",
    "hour":"%dh",
    "minute":"%dm",
    "second":"%ds",
    "now": "just now",
    "ago":"%s ago",
    "from_now":"in %s",
    "before":"%s before",
    "after":"%s after",
}
lang.SetResources(resources)

c := carbon.SetLanguage(lang)
c.Now().AddYears(1).DiffForHumans() // in 1 yr
c.Now().AddHours(1).DiffForHumans() // in 1h
c.Now().ToMonthString() // August
c.Now().ToShortMonthString() // Aug
c.Now().ToWeekString() // Tuesday
c.Now().ToShortWeekString() // Tue
c.Now().Constellation() // Leo
c.Now().Season() // Summer
```

##### Error handling
> If more than one error occurs, only the first error is returned

###### Scene one
```go
c := carbon.SetTimezone(PRC).Parse("xxx")
if c.Error != nil {
    // Error handle...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// Output
cannot parse "xxx" to carbon, please make sure the value is valid
```

###### Scene two
```go
c := carbon.SetTimezone("xxx").Parse("2020-08-05")
if c.Error != nil {
    // Error handle...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// Output
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezones
```
###### Scene three
```go
c := carbon.SetTimezone("xxx").Parse("12345678")
if c.Error != nil {
    // Error handle...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// Output
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezones
```

#### Appendix
##### <a id="format-sign-table">Format sign table</a>

| sign | desc | length | range | example |
| :------------: | :------------: | :------------: | :------------: | :------------: |
| d | Day of the month, 2 digits with leading zeros | 2 | 01-31 | 05 |
| D | A textual representation of a day, three letters | 3 | Mon-Sun | Wed |
| j | Day of the month without leading zeros | - |1-31 | 5 |
| S | English ordinal suffix for the day of the month, 2 characters. Eg: st, nd, rd or th. Works well with j | 2 | st/nd/rd/th | th |
| l | A full textual representation of the day of the week | - | Monday-Sunday | Wednesday |
| F | A full textual representation of a month | - | January-December | August |
| m | Numeric representation of a month, with leading zeros | 2 | 01-12 | 08 |
| M | A short textual representation of a month, three letters | 3 | Jan-Dec | Aug |
| n | Numeric representation of a month, without leading zeros | - | 1-12 | 8 |
| y | A two digit representation of a year | 2 | 00-99 | 20 |
| Y | A full numeric representation of a year, 4 digits | 4 | 0000-9999 | 2020 |
| a | Lowercase morning or afternoon sign | 2 | am/pm | pm |
| A | Uppercase morning or afternoon sign | 2 | AM/PM | PM |
| g | 12-hour format of an hour without leading zeros | - | 1-12 | 1 |
| G | 24-hour format of an hour without leading zeros | - | 0-23 | 15 |
| h | 12-hour format of an hour with leading zeros | 2 | 00-11 | 03 |
| H | 24-hour format of an hour with leading zeros | 2 | 00-23 | 15 |
| i | Minutes with leading zeros | 2 | 01-59 | 14 |
| s | Seconds with leading zeros | 2 | 01-59 | 15 |
| c | ISO 8601 date | - | - | 2020-08-05T15:19:21+00:00 |
| r | RFC 2822 date | - | - | Thu, 21 Dec 2020 16:01:07 +0200 |
| O | Difference to Greenwich time (GMT) without colon between hours and minutes | - | - | +0200 |
| P | Difference to Greenwich time (GMT) with colon between hours and minutes | - | - | +02:00 |
| T | Timezone abbreviation | - | - | EST |
| W | ISO-8601 numeric representation of the week of the year | - | 1-52 | 42 |
| N | ISO-8601 numeric representation of the day of the week | 1 | 1-7 | 6 |
| L | Whether it's a leap year | 1 | 0-1 | 1 |
| U | Seconds since the Unix Epoch (January 1 1970 00:00:00 GMT) | 10 | - | 1611818268 |
| u | Millisecond| 3 | - | 999 |
| w | Numeric representation of the day of the week | 1 | 0-6 | 6 |
| t | Total days of the month | 2 | 28-31 | 30 |
| z | Day of the year | - | 0-365 | 15 |
| e | Location | - | - | America/New_York |
| Q | Quarter | 1 | 1-4 | 1 |
| C | Century | - | 0-99 | 21 |

#### Reference
* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)
* [jinzhu/now](https://github.com/jinzhu/now/)
* [araddon/dateparse](https://github.com/araddon/dateparse)
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)
* [kofoworola/godate](https://github.com/kofoworola/godate)
* [arrow-py/arrow](https://github.com/arrow-py/arrow)
* [overtrue/chinese-calendar](https://github.com/overtrue/chinese-calendar)
* [moment/moment](https://github.com/moment/moment)
* [iamkun/dayjs](https://github.com/iamkun/dayjs)

#### Links
* [changelog](https://github.com/golang-module/carbon/wiki/%E6%9B%B4%E6%96%B0%E6%97%A5%E5%BF%97)
* [segmentfault](https://segmentfault.com/t/go-carbon)
* [oschina](https://www.oschina.net/p/go-carbon)
