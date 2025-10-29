---
head:
  - - meta
    - name: description
      content: Boundary | Provides start and end times for century/decade/year/quarter/month/week/day/hour/minute/second, supports configuring the week start day
  - - meta
    - name: keywords
      content: carbon, go-carbon, century boundary, decade boundary, year boundary, quarter boundary, month boundary, week boundary, day boundary, hour boundary, minute boundary, second boundary
---

# Boundary

## Century boundary
```go
// Start of the century
carbon.Parse("2020-08-05 13:14:15").StartOfCentury().ToDateTimeString() // 2000-01-01 00:00:00
// End of the century
carbon.Parse("2020-08-05 13:14:15").EndOfCentury().ToDateTimeString() // 2999-12-31 23:59:59
```

## Decade boundary
```go
// Start of the decade
carbon.Parse("2020-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
carbon.Parse("2021-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
carbon.Parse("2029-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
// End of the decade
carbon.Parse("2020-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
carbon.Parse("2021-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
carbon.Parse("2029-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
```

## Year boundary
```go
// Start of the year
carbon.Parse("2020-08-05 13:14:15").StartOfYear().ToDateTimeString() // 2020-01-01 00:00:00
// End of the year
carbon.Parse("2020-08-05 13:14:15").EndOfYear().ToDateTimeString() // 2020-12-31 23:59:59
```

## Quarter boundary
```go
// Start of the quarter
carbon.Parse("2020-08-05 13:14:15").StartOfQuarter().ToDateTimeString() // 2020-07-01 00:00:00
// End of the quarter
carbon.Parse("2020-08-05 13:14:15").EndOfQuarter().ToDateTimeString() // 2020-09-30 23:59:59
```

## Month boundary
```go
// Start of the month
carbon.Parse("2020-08-05 13:14:15").StartOfMonth().ToDateTimeString() // 2020-08-01 00:00:00
// End of the month
carbon.Parse("2020-08-05 13:14:15").EndOfMonth().ToDateTimeString() // 2020-08-31 23:59:59
```

## Week boundary
```go
// Start of the week
carbon.Parse("2020-08-05 13:14:15").StartOfWeek().ToDateTimeString() // 2020-08-02 00:00:00
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Sunday).StartOfWeek().ToDateTimeString() // 2020-08-02 00:00:00
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Monday).StartOfWeek().ToDateTimeString() // 2020-08-03 00:00:00
// End of the week
carbon.Parse("2020-08-05 13:14:15").EndOfWeek().ToDateTimeString() // 2020-08-08 23:59:59
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Sunday).EndOfWeek().ToDateTimeString() // 2020-08-08 23:59:59
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Monday).EndOfWeek().ToDateTimeString() // 2020-08-09 23:59:59
```

## Day boundary
```go
// Start of the day
carbon.Parse("2020-08-05 13:14:15").StartOfDay().ToDateTimeString() // 2020-08-05 00:00:00
// End of the day
carbon.Parse("2020-08-05 13:14:15").EndOfDay().ToDateTimeString() // 2020-08-05 23:59:59
```

## Hour boundary
```go
// Start of the hour
carbon.Parse("2020-08-05 13:14:15").StartOfHour().ToDateTimeString() // 2020-08-05 13:00:00
// End of the hour
carbon.Parse("2020-08-05 13:14:15").EndOfHour().ToDateTimeString() // 2020-08-05 13:59:59
```

## Minute boundary
```go
// Start of the minute
carbon.Parse("2020-08-05 13:14:15").StartOfMinute().ToDateTimeString() // 2020-08-05 13:14:00
// End of the minute
carbon.Parse("2020-08-05 13:14:15").EndOfMinute().ToDateTimeString() // 2020-08-05 13:14:59
```

## Second boundary
```go
// Start of the second
carbon.Parse("2020-08-05 13:14:15").StartOfSecond().ToString() // 2020-08-05 13:14:15 +0000 UTC
// End of the second
carbon.Parse("2020-08-05 13:14:15").EndOfSecond().ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC
```