---
head:
  - - meta
    - name: description
      content: 시간 판단|가벼우면서도 의미론적이고 개발자 친화적인 golang 시간 처리 라이브러리
  - - meta
    - name: keywords
      content: carbon, go-carbon, 포함 판단, 여부 판단, 수학 판단, 거리 판단
---

# 시간 판단

## 포함 판단
```go
// 오류가 있는지
carbon.NewCarbon().HasError() // false
carbon.ZeroValue().HasError() // false
carbon.EpochValue().HasError() // false
carbon.CreateFromTimestamp(0).HasError() // false
carbon.Parse("").HasError() // false
carbon.Parse("0").HasError() // true
carbon.Parse("xxx").HasError() // true
carbon.Parse("2020-08-05").HasError() // false
```

## 여부 판단
```go
// 영값 시간인지(0001-01-01 00:00:00 +0000 UTC)
carbon.NewCarbon().IsZero() // true
carbon.ZeroValue().IsZero() // true
carbon.EpochValue().IsZero() // false
carbon.CreateFromTimestamp(0).IsZero() // false
carbon.Parse("").IsZero() // false
carbon.Parse("0").IsZero() // false
carbon.Parse("xxx").IsZero() // false
carbon.Parse("2020-08-05").IsZero() // false

// 빈값인지
carbon.NewCarbon().IsEmpty() // false
carbon.ZeroValue().IsEmpty() // false
carbon.EpochValue().IsEmpty() // false
carbon.CreateFromTimestamp(0).IsEmpty() // false
carbon.Parse("").IsEmpty() // true
carbon.Parse("0").IsEmpty() // false
carbon.Parse("xxx").IsEmpty() // false
carbon.Parse("2020-08-05").IsEmpty() // false

// UNIX 에포크 시간인지(1970-01-01 00:00:00 +0000 UTC)
carbon.NewCarbon().IsEpoch() // false
carbon.ZeroValue().IsEpoch() // false
carbon.EpochValue().IsEpoch() // true
carbon.CreateFromTimestamp(0).IsEpoch() // true
carbon.Parse("").IsEpoch() // false
carbon.Parse("0").IsEpoch() // false
carbon.Parse("xxx").IsEpoch() // false
carbon.Parse("2020-08-05").IsEpoch() // false

// 유효한 시간인지
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

// 무효한 시간인지
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

// 서머타임인지
carbon.Parse("").IsDST() // false
carbon.Parse("0").IsDST() // false
carbon.Parse("xxx").IsDST() // false
carbon.Parse("0000-00-00 00:00:00").IsDST() // false
carbon.Parse("0000-00-00").IsDST() // false
carbon.Parse("00:00:00").IsDST() // false
carbon.Parse("2023-01-01", "Australia/Brisbane").IsDST() // false
carbon.Parse("2023-01-01", "Australia/Sydney").IsDST() // true

// 오전인지
carbon.Parse("2020-08-05 00:00:00").IsAM() // true
carbon.Parse("2020-08-05 08:00:00").IsAM() // true
carbon.Parse("2020-08-05 12:00:00").IsAM() // false
carbon.Parse("2020-08-05 13:00:00").IsAM() // false
// 오후인지
carbon.Parse("2020-08-05 00:00:00").IsPM() // false
carbon.Parse("2020-08-05 08:00:00").IsPM() // false
carbon.Parse("2020-08-05 12:00:00").IsPM() // true
carbon.Parse("2020-08-05 13:00:00").IsPM() // true

// 현재 시간인지
carbon.Now().IsNow() // true
// 미래 시간인지
carbon.Tomorrow().IsFuture() // true
// 과거 시간인지
carbon.Yesterday().IsPast() // true

// 윤년인지
carbon.Parse("2020-08-05 13:14:15").IsLeapYear() // true
// 장년인지
carbon.Parse("2020-08-05 13:14:15").IsLongYear() // true

// 1월인지
carbon.Parse("2020-08-05 13:14:15").IsJanuary() // false
// 2월인지
carbon.Parse("2020-08-05 13:14:15").IsFebruary() // false
// 3월인지
carbon.Parse("2020-08-05 13:14:15").IsMarch() // false
// 4월인지
carbon.Parse("2020-08-05 13:14:15").IsApril() // false
// 5월인지
carbon.Parse("2020-08-05 13:14:15").IsMay() // false
// 6월인지
carbon.Parse("2020-08-05 13:14:15").IsJune() // false
// 7월인지
carbon.Parse("2020-08-05 13:14:15").IsJuly() // false
// 8월인지
carbon.Parse("2020-08-05 13:14:15").IsAugust() // false
// 9월인지
carbon.Parse("2020-08-05 13:14:15").IsSeptember() // true
// 10월인지
carbon.Parse("2020-08-05 13:14:15").IsOctober() // false
// 11월인지
carbon.Parse("2020-08-05 13:14:15").IsNovember() // false
// 12월인지
carbon.Parse("2020-08-05 13:14:15").IsDecember() // false

// 월요일인지
carbon.Parse("2020-08-05 13:14:15").IsMonday() // false
// 화요일인지
carbon.Parse("2020-08-05 13:14:15").IsTuesday() // true
// 수요일인지
carbon.Parse("2020-08-05 13:14:15").IsWednesday() // false
// 목요일인지
carbon.Parse("2020-08-05 13:14:15").IsThursday() // false
// 금요일인지
carbon.Parse("2020-08-05 13:14:15").IsFriday() // false
// 토요일인지
carbon.Parse("2020-08-05 13:14:15").IsSaturday() // false
// 일요일인지
carbon.Parse("2020-08-05 13:14:15").IsSunday() // false

// 평일인지
carbon.Parse("2020-08-05 13:14:15").IsWeekday() // false
// 주말인지
carbon.Parse("2020-08-05 13:14:15").IsWeekend() // true

// 어제인지
carbon.Parse("2020-08-04 13:14:15").IsYesterday() // true
carbon.Parse("2020-08-04 00:00:00").IsYesterday() // true
carbon.Parse("2020-08-04").IsYesterday() // true
// 오늘인지
carbon.Parse("2020-08-05 13:14:15").IsToday() // true
carbon.Parse("2020-08-05 00:00:00").IsToday() // true
carbon.Parse("2020-08-05").IsToday() // true
// 내일인지
carbon.Parse("2020-08-06 13:14:15").IsTomorrow() // true
carbon.Parse("2020-08-06 00:00:00").IsTomorrow() // true
carbon.Parse("2020-08-06").IsTomorrow() // true

// 같은 세기인지
carbon.Parse("2020-08-05 13:14:15").IsSameCentury(carbon.Parse("3020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameCentury(carbon.Parse("2099-08-05 13:14:15")) // true
// 같은 연대인지
carbon.Parse("2020-08-05 13:14:15").IsSameDecade(carbon.Parse("2030-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameDecade(carbon.Parse("2120-08-05 13:14:15")) // true
// 같은 연도인지
carbon.Parse("2020-08-05 00:00:00").IsSameYear(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameYear(carbon.Parse("2020-12-31 13:14:15")) // true
// 같은 분기인지
carbon.Parse("2020-08-05 00:00:00").IsSameQuarter(carbon.Parse("2020-09-05 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameQuarter(carbon.Parse("2021-01-31 13:14:15")) // true
// 같은 월인지
carbon.Parse("2020-01-01 00:00:00").IsSameMonth(carbon.Parse("2021-01-31 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameMonth(carbon.Parse("2020-01-31 13:14:15")) // true
// 같은 날인지
carbon.Parse("2020-08-05 13:14:15").IsSameDay(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 00:00:00").IsSameDay(carbon.Parse("2020-08-05 13:14:15")) // true
// 같은 시간인지
carbon.Parse("2020-08-05 13:14:15").IsSameHour(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:00:00").IsSameHour(carbon.Parse("2020-08-05 13:14:15")) // true
// 같은 분인지
carbon.Parse("2020-08-05 13:14:15").IsSameMinute(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:00").IsSameMinute(carbon.Parse("2020-08-05 13:14:15")) // true
// 같은 초인지
carbon.Parse("2020-08-05 13:14:15").IsSameSecond(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameSecond(carbon.Parse("2020-08-05 13:14:15")) // true
```

## 수학 판단
```go
// 큰지
carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-05 13:14:15")) // false

// 작은지
carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-05 13:14:15")) // false

// 같은지
carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:00")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:00")) // false

// 다른지
carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("!=", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<>", carbon.Parse("2020-08-05 13:14:15")) // false

// 크거나 같은지
carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-05 13:14:15")) // true

// 작거나 같은지
carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-05 13:14:15")) // true
```

## 거리 판단
```go
// 두 시간 사이에 있는지(이 두 시간은 포함하지 않음)
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 두 시간 사이에 있는지(시작 시간 포함)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 두 시간 사이에 있는지(종료 시간 포함)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 두 시간 사이에 있는지(이 두 시간 모두 포함)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
``` 