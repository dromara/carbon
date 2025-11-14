---
head:
  - - meta
    - name: description
      content: エラー/有効性/ゼロ値/エポック/AM-PM/昨日-今日-明日/同一期間などの判定を提供し、大小/等号比較と Between（開始/終了境界の選択可能）などの距離判定をサポート
  - - meta
    - name: keywords
      content: carbon, go-carbon, 含有判定, 判定, 数学判定, 距離判定
---

# 時間比較

## 含有判定
```go
// エラーがありますか
carbon.NewCarbon().HasError() // false
carbon.ZeroValue().HasError() // false
carbon.EpochValue().HasError() // false
carbon.CreateFromTimestamp(0).HasError() // false
carbon.Parse("").HasError() // false
carbon.Parse("0").HasError() // true
carbon.Parse("xxx").HasError() // true
carbon.Parse("2020-08-05").HasError() // false
```

## 判定
```go
// ゼロ値の時間かどうか(0001-01-01 00:00:00 +0000 UTC)
carbon.NewCarbon().IsZero() // true
carbon.ZeroValue().IsZero() // true
carbon.EpochValue().IsZero() // false
carbon.CreateFromTimestamp(0).IsZero() // false
carbon.Parse("").IsZero() // false
carbon.Parse("0").IsZero() // false
carbon.Parse("xxx").IsZero() // false
carbon.Parse("2020-08-05").IsZero() // false

// 空値かどうか
carbon.NewCarbon().IsEmpty() // false
carbon.ZeroValue().IsEmpty() // false
carbon.EpochValue().IsEmpty() // false
carbon.CreateFromTimestamp(0).IsEmpty() // false
carbon.Parse("").IsEmpty() // true
carbon.Parse("0").IsEmpty() // false
carbon.Parse("xxx").IsEmpty() // false
carbon.Parse("2020-08-05").IsEmpty() // false

// UNIX 紀元時間かどうか（1970-01-01 00:00:00 +0000 UTC）
carbon.NewCarbon().IsEpoch() // false
carbon.ZeroValue().IsEpoch() // false
carbon.EpochValue().IsEpoch() // true
carbon.CreateFromTimestamp(0).IsEpoch() // true
carbon.Parse("").IsEpoch() // false
carbon.Parse("0").IsEpoch() // false
carbon.Parse("xxx").IsEpoch() // false
carbon.Parse("2020-08-05").IsEpoch() // false

// 有効な時間かどうか
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

// 無効な時間かどうか
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

// 夏時間かどうか
carbon.Parse("").IsDST() // false
carbon.Parse("0").IsDST() // false
carbon.Parse("xxx").IsDST() // false
carbon.Parse("0000-00-00 00:00:00").IsDST() // false
carbon.Parse("0000-00-00").IsDST() // false
carbon.Parse("00:00:00").IsDST() // false
carbon.Parse("2023-01-01", "Australia/Brisbane").IsDST() // false
carbon.Parse("2023-01-01", "Australia/Sydney").IsDST() // true

// 午前かどうか
carbon.Parse("2020-08-05 00:00:00").IsAM() // true
carbon.Parse("2020-08-05 08:00:00").IsAM() // true
carbon.Parse("2020-08-05 12:00:00").IsAM() // false
carbon.Parse("2020-08-05 13:00:00").IsAM() // false
// 午後かどうか
carbon.Parse("2020-08-05 00:00:00").IsPM() // false
carbon.Parse("2020-08-05 08:00:00").IsPM() // false
carbon.Parse("2020-08-05 12:00:00").IsPM() // true
carbon.Parse("2020-08-05 13:00:00").IsPM() // true

// 現在かどうか
carbon.Now().IsNow() // true
// 未来かどうか
carbon.Tomorrow().IsFuture() // true
// 過去かどうか
carbon.Yesterday().IsPast() // true

// 閏年かどうか
carbon.Parse("2020-08-05 13:14:15").IsLeapYear() // true
// ISO8601で定められたLong Yearかどうか
carbon.Parse("2020-08-05 13:14:15").IsLongYear() // true

// 1月かどうか
carbon.Parse("2020-08-05 13:14:15").IsJanuary() // false
// 2月かどうか
carbon.Parse("2020-08-05 13:14:15").IsFebruary() // false
// 3月かどうか
carbon.Parse("2020-08-05 13:14:15").IsMarch() // false
// 4月かどうか
carbon.Parse("2020-08-05 13:14:15").IsApril()  // false
// 5月かどうか
carbon.Parse("2020-08-05 13:14:15").IsMay() // false
// 6月かどうか
carbon.Parse("2020-08-05 13:14:15").IsJune() // false
// 7月かどうか
carbon.Parse("2020-08-05 13:14:15").IsJuly() // false
// 8月かどうか
carbon.Parse("2020-08-05 13:14:15").IsAugust() // false
// 9月かどうか
carbon.Parse("2020-08-05 13:14:15").IsSeptember() // true
// 10月かどうか
carbon.Parse("2020-08-05 13:14:15").IsOctober() // false
// 11月かどうか
carbon.Parse("2020-08-05 13:14:15").IsNovember() // false
// 12月かどうか
carbon.Parse("2020-08-05 13:14:15").IsDecember() // false

// 月曜日かどうか
carbon.Parse("2020-08-05 13:14:15").IsMonday() // false
// 火曜日かどうか
carbon.Parse("2020-08-05 13:14:15").IsTuesday() // true
// 水曜日かどうか
carbon.Parse("2020-08-05 13:14:15").IsWednesday() // false
// 木曜日かどうか
carbon.Parse("2020-08-05 13:14:15").IsThursday() // false
// 金曜日かどうか
carbon.Parse("2020-08-05 13:14:15").IsFriday() // false
// 土曜日かどうか
carbon.Parse("2020-08-05 13:14:15").IsSaturday() // false
// 日曜日かどうか
carbon.Parse("2020-08-05 13:14:15").IsSunday() // false

// 平日かどうか
carbon.Parse("2020-08-05 13:14:15").IsWeekday() // false
// 週末かどうか
carbon.Parse("2020-08-05 13:14:15").IsWeekend() // true

// 昨日かどうか
carbon.Parse("2020-08-04 13:14:15").IsYesterday() // true
carbon.Parse("2020-08-04 00:00:00").IsYesterday() // true
carbon.Parse("2020-08-04").IsYesterday() // true
// 今日かどうか
carbon.Parse("2020-08-05 13:14:15").IsToday() // true
carbon.Parse("2020-08-05 00:00:00").IsToday() // true
carbon.Parse("2020-08-05").IsToday() // true
// 明日かどうか
carbon.Parse("2020-08-06 13:14:15").IsTomorrow() // true
carbon.Parse("2020-08-06 00:00:00").IsTomorrow() // true
carbon.Parse("2020-08-06").IsTomorrow() // true

// 同世紀かどうか
carbon.Parse("2020-08-05 13:14:15").IsSameCentury(carbon.Parse("3020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameCentury(carbon.Parse("2099-08-05 13:14:15")) // true
// 同十年紀かどうか
carbon.Parse("2020-08-05 13:14:15").IsSameDecade(carbon.Parse("2030-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameDecade(carbon.Parse("2120-08-05 13:14:15")) // true
// 同年かどうか
carbon.Parse("2020-08-05 00:00:00").IsSameYear(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameYear(carbon.Parse("2020-12-31 13:14:15")) // true
// 同四半期かどうか
carbon.Parse("2020-08-05 00:00:00").IsSameQuarter(carbon.Parse("2020-09-05 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameQuarter(carbon.Parse("2021-01-31 13:14:15")) // true
// 同月かどうか
carbon.Parse("2020-01-01 00:00:00").IsSameMonth(carbon.Parse("2021-01-31 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameMonth(carbon.Parse("2020-01-31 13:14:15")) // true
// 同日かどうか
carbon.Parse("2020-08-05 13:14:15").IsSameDay(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 00:00:00").IsSameDay(carbon.Parse("2020-08-05 13:14:15")) // true
// 同時間かどうか
carbon.Parse("2020-08-05 13:14:15").IsSameHour(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:00:00").IsSameHour(carbon.Parse("2020-08-05 13:14:15")) // true
// 同分かどうか
carbon.Parse("2020-08-05 13:14:15").IsSameMinute(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:00").IsSameMinute(carbon.Parse("2020-08-05 13:14:15")) // true
// 同秒かどうか
carbon.Parse("2020-08-05 13:14:15").IsSameSecond(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameSecond(carbon.Parse("2020-08-05 13:14:15")) // true
```

## 数学判定
```go
// 超過かどうか
carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-05 13:14:15")) // false

// 未満かどうか
carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-05 13:14:15")) // false

// 等しいかどうか
carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:00")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:00")) // false

// 等しくないかどうか
carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("!=", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<>", carbon.Parse("2020-08-05 13:14:15")) // false

// 以上かどうか
carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-05 13:14:15")) // true

// 以下かどうか
carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-05 13:14:15")) // true
```

## 距離判定
```go
// 二つの Carbon インスタンスの間に含まれているか(開始時間、終了時間を含まない)
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 二つの Carbon インスタンスの間に含まれているか(開始時間を含む)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 二つの Carbon インスタンスの間に含まれているか(終了時間を含む)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 二つの Carbon インスタンスの間に含まれているか(開始時間、終了時間を含む)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
```