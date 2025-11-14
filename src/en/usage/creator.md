---
head:
  - - meta
    - name: description
      content: "Provides methods to create instances from timestamps (second/millisecond/microsecond/nanosecond) and from datetime/date/time: CreateFromTimestamp/DateTime/Date/Time and their precision variants"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, create Carbon instance from timestamp, create Carbon instance from date"
---

# Creator

## Create from timestamp
```go
// Create a Carbon instance from a given timestamp with second precision
carbon.CreateFromTimestamp(-1).ToString() // 1969-12-31 23:59:59 +0000 UTC
carbon.CreateFromTimestamp(0).ToString() // 1970-01-01 00:00:00 +0000 UTC
carbon.CreateFromTimestamp(1).ToString() // 1970-01-01 00:00:01 +0000 UTC
carbon.CreateFromTimestamp(1596633255).ToString() // 2020-08-05 13:14:15 +0000 UTC
// Create a Carbon instance from a given timestamp with millisecond precision
carbon.CreateFromTimestampMilli(1596604455999).ToString() // 2020-08-05 13:14:15.999 +0000 UTC
// Create a Carbon instance from a given timestamp with microsecond precision
carbon.CreateFromTimestampMicro(1596604455999999).ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
// Create a Carbon instance from a given timestamp with nanosecond precision
carbon.CreateFromTimestampNano(1596604455999999999).ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC
```

## Create from datetime
```go
// Create a Carbon instance from a given date and time
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToString() // 2020-08-05 13:14:15 +0000 UTC
// Create a Carbon instance from a given date and time with millisecond
carbon.CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString() // 2020-08-05 13:14:15.999 +0000 UTC
// Create a Carbon instance from a given date and time with microsecond
carbon.CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
// Create a Carbon instance from a given date and time with nanosecond
carbon.CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC

// Create a Carbon instance from a given year, month and day
carbon.CreateFromDate(2020, 8, 5).ToString() // 2020-08-05 00:00:00 +0000 UTC
// Create a Carbon instance from a given year, month and day with millisecond
carbon.CreateFromDateMilli(2020, 8, 5, 999).ToString() // 2020-08-05 00:00:00.999 +0000 UTC
// Create a Carbon instance from a given year, month and day with microsecond
carbon.CreateFromDateMicro(2020, 8, 5, 999999).ToString() // 2020-08-05 00:00:00.999999 +0000 UTC
// Create a Carbon instance from a given year, month and day with nanosecond
carbon.CreateFromDateNano(2020, 8, 5, 999999999).ToString() // 2020-08-05 00:00:00.999999999 +0000 UTC

// Create a Carbon instance from a given hour, minute and second (year, month, day default to current)
carbon.CreateFromTime(13, 14, 15).ToString() // 2020-08-05 13:14:15 +0000 UTC
// Create a Carbon instance from a given hour, minute and second with millisecond (year, month, day default to current)
carbon.CreateFromTimeMilli(13, 14, 15, 999).ToString() // 2020-08-05 13:14:15.999 +0000 UTC
// Create a Carbon instance from a given hour, minute and second with microsecond (year, month, day default to current)
carbon.CreateFromTimeMicro(13, 14, 15, 999999).ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
// Create a Carbon instance from a given hour, minute and second with nanosecond (year, month, day default to current)
carbon.CreateFromTimeNano(13, 14, 15, 999999999).ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC
```