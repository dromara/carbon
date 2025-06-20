---
head:
  - - meta
    - name: description
      content: Yesterday、Today、Tomorrow | A simple, semantic and developer-friendly time package for golang
  - - meta
    - name: keywords
      content: carbon, go-carbon, yesterday, today, tomorrow
---

# Yesterday、Today、Tomorrow
Yesterday was the day before today, and tomorrow is the day after today, equivalent to

```go
carbon.Yesterday() = carbon.Now().SubDay()
carbon.Tomorrow() = carbon.Now().AddDay()
```
Assuming the current time is `2020-08-05 13:14:15.999999999 +0000 UTC`

## Yesterday
```go
// Return datetime of yesterday
fmt.Printf("%s", carbon.Yesterday()) // 2020-08-04 13:14:15
carbon.Yesterday().String() // 2020-08-04 13:14:15
carbon.Yesterday().ToString() // 2020-08-04 13:14:15.999999999 +0000 UTC
carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// Return date of yesterday
carbon.Yesterday().ToDateString() // 2020-08-04
// Return time of yesterday
carbon.Yesterday().ToTimeString() // 13:14:15
// Return datetime of yesterday in a given timezone
carbon.Yesterday(carbon.NewYork).ToDateTimeString() // 2020-08-04 13:14:15
// Return timestamp with second precision of yesterday
carbon.Yesterday().Timestamp() // 1596546855
// Return timestamp with millisecond precision of yesterday
carbon.Yesterday().TimestampMilli() // 1596546855999
// Return timestamp with microsecond precision of yesterday
carbon.Yesterday().TimestampMicro() // 1596546855999999
// Return timestamp with nanosecond precision of yesterday
carbon.Yesterday().TimestampNano() // 1596546855999999999
```

## Today
```go
// Return datetime of today
fmt.Printf("%s", carbon.Now()) // 2020-08-05 13:14:15
carbon.Now().String() // 2020-08-05 13:14:15
carbon.Now().ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC
carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
// Return date of today
carbon.Now().ToDateString() // 2020-08-05
// Return time of today
carbon.Now().ToTimeString() // 13:14:15
// Return datetime of today in a given timezone
carbon.Now(carbon.NewYork).ToDateTimeString() // 2020-08-05 13:14:15
// Return timestamp with second precision of today
carbon.Now().Timestamp() // 1596633255
// Return timestamp with millisecond precision of today
carbon.Now().TimestampMilli() // 1596633255999
// Return timestamp with microsecond precision of today
carbon.Now().TimestampMicro() // 1596633255999999
// Return timestamp with nanosecond precision of today
carbon.Now().TimestampNano() // 1596633255999999999
```

## Tomorrow
```go
// Return datetime of tomorrow
fmt.Printf("%s", carbon.Tomorrow()) // 2020-08-06 13:14:15
carbon.Tomorrow().String() // 2020-08-06 13:14:15
carbon.Tomorrow().ToString() // 22020-08-06 13:14:15.999999999 +0000 UTC
carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// Return date of tomorrow
carbon.Tomorrow().ToDateString() // 2020-08-06
// Return time of tomorrow
carbon.Tomorrow().ToTimeString() // 13:14:15
// Return datetime of tomorrow in a given timezone
carbon.Tomorrow(carbon.NewYork).ToDateTimeString() // 2020-08-06 13:14:15
// Return timestamp with second precision of tomorrow
carbon.Tomorrow().Timestamp() // 1596719655
// Return timestamp with millisecond precision of tomorrow
carbon.Tomorrow().TimestampMilli() // 1596719655999
// Return timestamp with microsecond precision of tomorrow
carbon.Tomorrow().TimestampMicro() // 1596719655999999
// Return timestamp with nanosecond precision of tomorrow
carbon.Tomorrow().TimestampNano() // 1596719655999999999
```