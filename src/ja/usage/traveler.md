---
head:
  - - meta
    - name: description
      content: 時間旅行 | 世紀/年代/年/四半期/月/週/日/時/分/秒/ミリ秒/マイクロ秒/ナノ秒の加減とオーバーフローなしバリアントを提供し、Duration に基づく増減をサポート
  - - meta
    - name: keywords
      content: carbon, go-carbon, 世紀旅行, 年代旅行, 年旅行, 四半期旅行, 月旅行, 週旅行, 日旅行, 時間旅行, 分旅行, 秒旅行, ミリ秒旅行, マイクロ秒旅行, ナノ秒旅行
---

# タイムトラベル

## 世紀旅行
```go
// 3世紀追加
carbon.Parse("2020-02-29 13:14:15").AddCenturies(3).ToDateTimeString() // 2320-02-29 13:14:15
// 3世紀追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddCenturiesNoOverflow(3).ToDateTimeString() // 2320-02-29 13:14:15
// 1世紀追加
carbon.Parse("2020-02-29 13:14:15").AddCentury().ToDateTimeString() // 2120-02-29 13:14:15
// 1世紀追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddCenturyNoOverflow().ToDateTimeString() // 2120-02-29 13:14:15
// 3世紀減らす
carbon.Parse("2020-02-29 13:14:15").SubCenturies(3).ToDateTimeString() // 1720-02-29 13:14:15
// 3世紀減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubCenturiesNoOverflow(3).ToDateTimeString() // 1720-02-29 13:14:15
// 1世紀減らす
carbon.Parse("2020-02-29 13:14:15").SubCentury().ToDateTimeString() // 1920-02-29 13:14:15
// 1世紀減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubCenturyNoOverflow().ToDateTimeString() // 1920-02-29 13:14:15
```

## 年代旅行
```go
// 3年代追加
carbon.Parse("2020-02-29 13:14:15").AddDecades(3).ToDateTimeString() // 2050-03-01 13:14:15
// 3年代追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddDecadesNoOverflow(3).ToDateTimeString() // 2050-02-28 13:14:15
// 1年代追加
carbon.Parse("2020-02-29 13:14:15").AddDecade().ToDateTimeString() // 2030-03-01 13:14:15
// 1年代追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddDecadeNoOverflow().ToDateTimeString() // 2030-02-28 13:14:15
// 3年代減らす
carbon.Parse("2020-02-29 13:14:15").SubDecades(3).ToDateTimeString() // 1990-03-01 13:14:15
// 3年代減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubDecadesNoOverflow(3).ToDateTimeString() // 1990-02-28 13:14:15
// 1年代減らす
carbon.Parse("2020-02-29 13:14:15").SubDecade().ToDateTimeString() // 2010-03-01 13:14:15
// 1年代減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubDecadeNoOverflow().ToDateTimeString() // 2010-02-28 13:14:15
```

## 年旅行
```go
// 3年追加
carbon.Parse("2020-02-29 13:14:15").AddYears(3).ToDateTimeString() // 2023-03-01 13:14:15
// 3年追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddYearsNoOverflow(3).ToDateTimeString() // 2023-02-28 13:14:15
// 1年追加
carbon.Parse("2020-02-29 13:14:15").AddYear().ToDateTimeString() // 2021-03-01 13:14:15
// 1年追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddYearNoOverflow().ToDateTimeString() // 2021-02-28 13:14:15
// 3年減らす
carbon.Parse("2020-02-29 13:14:15").SubYears(3).ToDateTimeString() // 2017-03-01 13:14:15
// 3年減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubYearsNoOverflow(3).ToDateTimeString() // 2017-02-28 13:14:15
// 1年減らす
carbon.Parse("2020-02-29 13:14:15").SubYear().ToDateTimeString() // 2019-03-01 13:14:15
// 1年減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubYearNoOverflow().ToDateTimeString() // 2019-02-28 13:14:15
```

## 四半期旅行
```go
// 3四半期追加
carbon.Parse("2019-05-31 13:14:15").AddQuarters(3).ToDateTimeString() // 2020-03-02 13:14:15
// 3四半期追加(オーバーフローなし)
carbon.Parse("2019-05-31 13:14:15").AddQuartersNoOverflow(3).ToDateTimeString() // 2020-02-29 13:14:15
// 1四半期追加
carbon.Parse("2019-11-30 13:14:15").AddQuarter().ToDateTimeString() // 2020-03-01 13:14:15
// 1四半期追加(オーバーフローなし)
carbon.Parse("2019-11-30 13:14:15").AddQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
// 3四半期減らす
carbon.Parse("2019-08-31 13:14:15").SubQuarters(3).ToDateTimeString() // 2019-03-03 13:14:15
// 3四半期減らす(オーバーフローなし)
carbon.Parse("2019-08-31 13:14:15").SubQuartersNoOverflow(3).ToDateTimeString() // 2019-02-28 13:14:15
// 1四半期減らす
carbon.Parse("2020-05-31 13:14:15").SubQuarter().ToDateTimeString() // 2020-03-02 13:14:15
// 1四半期減らす(オーバーフローなし)
carbon.Parse("2020-05-31 13:14:15").SubQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
```

## 月旅行
```go
// 3ヶ月追加
carbon.Parse("2020-02-29 13:14:15").AddMonths(3).ToDateTimeString() // 2020-05-29 13:14:15
// 3ヶ月追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddMonthsNoOverflow(3).ToDateTimeString() // 2020-05-29 13:14:15
// 1ヶ月追加
carbon.Parse("2020-01-31 13:14:15").AddMonth().ToDateTimeString() // 2020-03-02 13:14:15
// 1ヶ月追加(オーバーフローなし)
carbon.Parse("2020-01-31 13:14:15").AddMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
// 3ヶ月減らす
carbon.Parse("2020-02-29 13:14:15").SubMonths(3).ToDateTimeString() // 2019-11-29 13:14:15
// 3ヶ月減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubMonthsNoOverflow(3).ToDateTimeString() // 2019-11-29 13:14:15
// 1ヶ月減らす
carbon.Parse("2020-03-31 13:14:15").SubMonth().ToDateTimeString() // 2020-03-02 13:14:15
// 1か月減らす(オーバーフローなし)
carbon.Parse("2020-03-31 13:14:15").SubMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
```

## 週旅行
```go
// 3週間追加
carbon.Parse("2020-02-29 13:14:15").AddWeeks(3).ToDateTimeString() // 2020-03-21 13:14:15
// 1週間追加
carbon.Parse("2020-02-29 13:14:15").AddWeek().ToDateTimeString() // 2020-03-07 13:14:15
// 3週間減らす
carbon.Parse("2020-02-29 13:14:15").SubWeeks(3).ToDateTimeString() // 2020-02-08 13:14:15
// 1週間減らす
carbon.Parse("2020-02-29 13:14:15").SubWeek().ToDateTimeString() // 2020-02-22 13:14:15
```

## 日旅行
```go
// 3日追加
carbon.Parse("2020-08-05 13:14:15").AddDays(3).ToDateTimeString() // 2020-08-08 13:14:15
// 1日追加
carbon.Parse("2020-08-05 13:14:15").AddDay().ToDateTimeString() // 2020-08-06 13:14:15
// 3日減らす
carbon.Parse("2020-08-05 13:14:15").SubDays(3).ToDateTimeString() // 2020-08-02 13:14:15
// 1日減らす
carbon.Parse("2020-08-05 13:14:15").SubDay().ToDateTimeString() // 2020-08-04 13:14:15
```

## 時間旅行
```go
// 3時間追加
carbon.Parse("2020-08-05 13:14:15").AddHours(3).ToDateTimeString() // 2020-08-05 16:14:15
// 2時間半追加
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5h").ToDateTimeString() // 2020-08-05 15:44:15
carbon.Parse("2020-08-05 13:14:15").AddDuration("2h30m").ToDateTimeString() // 2020-08-05 15:44:15
// 1時間追加
carbon.Parse("2020-08-05 13:14:15").AddHour().ToDateTimeString() // 2020-08-05 14:14:15
// 3時間減らす
carbon.Parse("2020-08-05 13:14:15").SubHours(3).ToDateTimeString() // 2020-08-05 10:14:15
// 2時間半減らす
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5h").ToDateTimeString() // 2020-08-05 10:44:15
carbon.Parse("2020-08-05 13:14:15").SubDuration("2h30m").ToDateTimeString() // 2020-08-05 10:44:15
// 1時間減らす
carbon.Parse("2020-08-05 13:14:15").SubHour().ToDateTimeString() // 2020-08-05 12:14:15
```

## 分旅行
```go
// 3分追加
carbon.Parse("2020-08-05 13:14:15").AddMinutes(3).ToDateTimeString() // 2020-08-05 13:17:15
// 2分半追加
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5m").ToDateTimeString() // 2020-08-05 13:16:45
carbon.Parse("2020-08-05 13:14:15").AddDuration("2m30s").ToDateTimeString() // 2020-08-05 13:16:45
// 1分追加
carbon.Parse("2020-08-05 13:14:15").AddMinute().ToDateTimeString() // 2020-08-05 13:15:15
// 3分減らす
carbon.Parse("2020-08-05 13:14:15").SubMinutes(3).ToDateTimeString() // 2020-08-05 13:11:15
// 2分半減らす
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5m").ToDateTimeString() // 2020-08-05 13:11:45
carbon.Parse("2020-08-05 13:14:15").SubDuration("2m30s").ToDateTimeString() // 2020-08-05 13:11:45
// 1分減らす
carbon.Parse("2020-08-05 13:14:15").SubMinute().ToDateTimeString() // 2020-08-05 13:13:15
```

## 秒旅行
```go
// 3秒追加
carbon.Parse("2020-08-05 13:14:15").AddSeconds(3).ToDateTimeString() // 2020-08-05 13:14:18
// 2秒半追加
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:17
// 1秒追加
carbon.Parse("2020-08-05 13:14:15").AddSecond().ToDateTimeString() // 2020-08-05 13:14:16
// 3秒減らす
carbon.Parse("2020-08-05 13:14:15").SubSeconds(3).ToDateTimeString() // 2020-08-05 13:14:12
// 2秒半減らす
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:12
// 1秒減らす
carbon.Parse("2020-08-05 13:14:15").SubSecond().ToDateTimeString() // 2020-08-05 13:14:14
```

## ミリ秒旅行
```go
// 3ミリ秒追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddMilliseconds(3).ToString() // 2020-08-05 13:14:15.225222222 +0000 UTC
// 1ミリ秒追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddMillisecond().ToString() // 2020-08-05 13:14:15.223222222 +0000 UTC
// 3ミリ秒減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubMilliseconds(3).ToString() // 2020-08-05 13:14:15.219222222 +0000 UTC
// 1ミリ秒減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubMillisecond().ToString() // 2020-08-05 13:14:15.221222222 +0000 UTC
```

## マイクロ秒旅行
```go
// 3マイクロ秒追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddMicroseconds(3).ToString() // 2020-08-05 13:14:15.222225222 +0000 UTC
// 1マイクロ秒追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddMicrosecond().ToString() // 2020-08-05 13:14:15.222223222 +0000 UTC
// 3マイクロ秒減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubMicroseconds(3).ToString() // 2020-08-05 13:14:15.222219222 +0000 UTC
// 1マイクロ秒減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubMicrosecond().ToString() // 2020-08-05 13:14:15.222221222 +0000 UTC
```

## ナノ秒旅行
```go
// 3ナノ秒追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddNanoseconds(3).ToString() // 2020-08-05 13:14:15.222222225 +0000 UTC
// 1ナノ秒追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddNanosecond().ToString() // 2020-08-05 13:14:15.222222223 +0000 UTC
// 3ナノ秒減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubNanoseconds(3).ToString() // 2020-08-05 13:14:15.222222219 +0000 UTC
// 1ナノ秒減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubNanosecond().ToString() // 2020-08-05 13:14:15.222222221 +0000 UTC
```