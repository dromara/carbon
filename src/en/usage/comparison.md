---
head:
  - - meta
    - name: description
      content: Comparison | Provides checks for error/valid/zero/epoch/AM-PM/yesterday-today-tomorrow/same-period, supports maths compares (>, <, =, !=, >=, <=) and Between variants (exclusive/inclusive start/end)
  - - meta
    - name: keywords
      content: carbon, go-carbon, has compare, is compare, maths compare, between compare
---

# Comparison

## Has compare
```go
// Whether has error
carbon.NewCarbon().HasError() // false
carbon.ZeroValue().HasError() // false
carbon.EpochValue().HasError() // false
carbon.CreateFromTimestamp(0).HasError() // false
carbon.Parse("").HasError() // false
carbon.Parse("0").HasError() // true
carbon.Parse("xxx").HasError() // true
carbon.Parse("2020-08-05").HasError() // false
```

## Is compare
```go
// Whether is a zero time(0001-01-01 00:00:00 +0000 UTC)
carbon.NewCarbon().IsZero() // true
carbon.ZeroValue().IsZero() // true
carbon.EpochValue().IsZero() // false
carbon.CreateFromTimestamp(0).IsZero() // false
carbon.Parse("").IsZero() // false
carbon.Parse("0").IsZero() // false
carbon.Parse("xxx").IsZero() // false
carbon.Parse("2020-08-05").IsZero() // false

// Whether is a empty time
carbon.NewCarbon().IsEmpty() // false
carbon.ZeroValue().IsEmpty() // false
carbon.EpochValue().IsEmpty() // false
carbon.CreateFromTimestamp(0).IsEmpty() // false
carbon.Parse("").IsEmpty() // true
carbon.Parse("0").IsEmpty() // false
carbon.Parse("xxx").IsEmpty() // false
carbon.Parse("2020-08-05").IsEmpty() // false

// Whether is a unix epoch time(1970-01-01 00:00:00 +0000 UTC).
carbon.NewCarbon().IsEpoch() // false
carbon.ZeroValue().IsEpoch() // false
carbon.EpochValue().IsEpoch() // true
carbon.CreateFromTimestamp(0).IsEpoch() // true
carbon.Parse("").IsEpoch() // false
carbon.Parse("0").IsEpoch() // false
carbon.Parse("xxx").IsEpoch() // false
carbon.Parse("2020-08-05").IsEpoch() // false

// Whether is a valid time
carbon.NewCarbon().IsValid() // true
carbon.ZeroValue().IsValid() // true
carbon.EpochValue().IsEpoch() // true
carbon.CreateFromTimestamp(0).IsValid() // true
carbon.Parse("").IsValid() // false
carbon.Parse("0").IsValid() // false
carbon.Parse("xxx").IsValid() // false
carbon.Parse("0000-00-00 00:00:00").IsValid() // false
carbon.Parse("0000-00-00").IsValid() // false
carbon.Parse("00:00:00").IsValid() // false
carbon.Parse("2020-08-05 00:00:00").IsValid() // true
carbon.Parse("2020-08-05").IsValid() // true

// Whether is an invalid time
carbon.NewCarbon().IsInvalid() // false
carbon.ZeroValue().IsInvalid() // false
carbon.EpochValue().IsInvalid() // false
carbon.CreateFromTimestamp(0).IsInvalid() // false
carbon.Parse("").IsInvalid() // false
carbon.Parse("0").IsInvalid() // true
carbon.Parse("xxx").IsInvalid() // true
carbon.Parse("0000-00-00 00:00:00").IsInvalid() // true
carbon.Parse("0000-00-00").IsInvalid() // true
carbon.Parse("00:00:00").IsInvalid() // true
carbon.Parse("2020-08-05 00:00:00").IsInvalid() // false
carbon.Parse("2020-08-05").IsInvalid() // false

// Whether is a daylight saving time
carbon.Parse("").IsDST() // false
carbon.Parse("0").IsDST() // false
carbon.Parse("xxx").IsDST() // false
carbon.Parse("0000-00-00 00:00:00").IsDST() // false
carbon.Parse("0000-00-00").IsDST() // false
carbon.Parse("00:00:00").IsDST() // false
carbon.Parse("2023-01-01", "Australia/Brisbane").IsDST() // false
carbon.Parse("2023-01-01", "Australia/Sydney").IsDST() // true

// Whether is a forenoon time
carbon.Parse("2020-08-05 00:00:00").IsAM() // true
carbon.Parse("2020-08-05 08:00:00").IsAM() // true
carbon.Parse("2020-08-05 12:00:00").IsAM() // false
carbon.Parse("2020-08-05 13:00:00").IsAM() // false
// Whether is an afternoon time 
carbon.Parse("2020-08-05 00:00:00").IsPM() // false
carbon.Parse("2020-08-05 08:00:00").IsPM() // false
carbon.Parse("2020-08-05 12:00:00").IsPM() // true
carbon.Parse("2020-08-05 13:00:00").IsPM() // true

// Whether is a now time
carbon.Now().IsNow() // true
// Whether is a future time
carbon.Tomorrow().IsFuture() // true
// Whether is a pass time
carbon.Yesterday().IsPast() // true

// Whether is a leap year
carbon.Parse("2020-08-05 13:14:15").IsLeapYear() // true
// Whether is a long year
carbon.Parse("2020-08-05 13:14:15").IsLongYear() // true

// Whether is January
carbon.Parse("2020-08-05 13:14:15").IsJanuary() // false
// Whether is February
carbon.Parse("2020-08-05 13:14:15").IsFebruary() // false
// Whether is March
carbon.Parse("2020-08-05 13:14:15").IsMarch() // false
// Whether is April
carbon.Parse("2020-08-05 13:14:15").IsApril()  // false
// Whether is May
carbon.Parse("2020-08-05 13:14:15").IsMay() // false
// Whether is June
carbon.Parse("2020-08-05 13:14:15").IsJune() // false
// Whether is July
carbon.Parse("2020-08-05 13:14:15").IsJuly() // false
// Whether is August
carbon.Parse("2020-08-05 13:14:15").IsAugust() // false
// Whether is September
carbon.Parse("2020-08-05 13:14:15").IsSeptember() // true
// Whether is October
carbon.Parse("2020-08-05 13:14:15").IsOctober() // false
// Whether is November
carbon.Parse("2020-08-05 13:14:15").IsNovember() // false
// Whether is December
carbon.Parse("2020-08-05 13:14:15").IsDecember() // false

// Whether is Monday
carbon.Parse("2020-08-05 13:14:15").IsMonday() // false
// Whether is Tuesday
carbon.Parse("2020-08-05 13:14:15").IsTuesday() // true
// Whether is Wednesday
carbon.Parse("2020-08-05 13:14:15").IsWednesday() // false
// Whether is Thursday
carbon.Parse("2020-08-05 13:14:15").IsThursday() // false
// Whether is Friday
carbon.Parse("2020-08-05 13:14:15").IsFriday() // false
// Whether is Saturday
carbon.Parse("2020-08-05 13:14:15").IsSaturday() // false
// Whether is Sunday
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

// Whether is the same century
carbon.Parse("2020-08-05 13:14:15").IsSameCentury(carbon.Parse("3020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameCentury(carbon.Parse("2099-08-05 13:14:15")) // true
// Whether is the same decade
carbon.Parse("2020-08-05 13:14:15").IsSameDecade(carbon.Parse("2030-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameDecade(carbon.Parse("2120-08-05 13:14:15")) // true
// Whether is the same year
carbon.Parse("2020-08-05 00:00:00").IsSameYear(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameYear(carbon.Parse("2020-12-31 13:14:15")) // true
// Whether is the same quarter
carbon.Parse("2020-08-05 00:00:00").IsSameQuarter(carbon.Parse("2020-09-05 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameQuarter(carbon.Parse("2021-01-31 13:14:15")) // true
// Whether is the same month
carbon.Parse("2020-01-01 00:00:00").IsSameMonth(carbon.Parse("2021-01-31 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameMonth(carbon.Parse("2020-01-31 13:14:15")) // true
// Whether is the same day
carbon.Parse("2020-08-05 13:14:15").IsSameDay(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 00:00:00").IsSameDay(carbon.Parse("2020-08-05 13:14:15")) // true
// Whether is the same hour
carbon.Parse("2020-08-05 13:14:15").IsSameHour(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:00:00").IsSameHour(carbon.Parse("2020-08-05 13:14:15")) // true
// Whether is the same minute
carbon.Parse("2020-08-05 13:14:15").IsSameMinute(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:00").IsSameMinute(carbon.Parse("2020-08-05 13:14:15")) // true
// Whether is the same second
carbon.Parse("2020-08-05 13:14:15").IsSameSecond(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameSecond(carbon.Parse("2020-08-05 13:14:15")) // true
```

## Maths compare
```go
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
```

## Between compare
```go
// Whether between two Carbon instances, excluded the start and end Carbon instance
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// Whether between two Carbon instances, included the start Carbon instance
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// Whether between two Carbon instances, included the end Carbon instance
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// Whether between two Carbon instances, included the start and end Carbon instance
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
```