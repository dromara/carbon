---
head:
  - - meta
    - name: description
      content: 時間設定|軽量、セマンティック、開発者フレンドリーな golang 時間処理ライブラリ
  - - meta
    - name: keywords
      content: carbon, go-carbon, タイムゾーン設定, 位置設定, 地域設定, 年月日時分秒設定, 年設定, 月設定, 週開始日設定, 週末日設定, 日設定, 時設定, 分設定, 秒設定, ミリ秒設定, マイクロ秒設定, ナノ秒設定
---

# 時間設定

```go
// タイムゾーン設定
carbon.Parse("2020-08-05 13:14:15").SetTimezone(carbon.UTC).ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-08-05 13:14:15").SetTimezone(carbon.PRC).ToString() // 2020-08-05 21:14:15 +0800 CST
carbon.Parse("2020-08-05 13:14:15").SetTimezone(carbon.Tokyo).ToString() // 2020-08-05 22:14:15 +0900 JST

// 位置設定
utc, _ := time.LoadLocation(carbon.UTC)
carbon.Parse("2020-08-05 13:14:15").SetLocation(utc).ToString() // 2020-08-05 13:14:15 +0000 UTC
prc, _ := time.LoadLocation(carbon.PRC)
carbon.Parse("2020-08-05 13:14:15").SetLocation(prc).ToString() // 2020-08-05 21:14:15 +0800 CST
tokyo, _ := time.LoadLocation(carbon.Tokyo)
carbon.Parse("2020-08-05 13:14:15").SetLocation(tokyo).ToString() // 2020-08-05 22:14:15 +0900 JST

// 地域設定
carbon.Parse("2020-07-05 13:14:15").SetLocale("en").DiffForHumans() // 1 month ago
carbon.Parse("2020-07-05 13:14:15").SetLocale("zh-CN").DiffForHumans() // 1 月前

// 年、月、日、時、分、秒設定
carbon.Parse("2020-01-01").SetDateTime(2019, 2, 2, 13, 14, 15).ToString() // 2019-02-02 13:14:15 +0000 UTC
carbon.Parse("2020-01-01").SetDateTime(2019, 2, 31, 13, 14, 15).ToString() // 2019-03-03 13:14:15 +0000 UTC
// 年、月、日、時、分、秒、ミリ秒設定
carbon.Parse("2020-01-01").SetDateTimeMilli(2019, 2, 2, 13, 14, 15, 999).ToString() // 2019-02-02 13:14:15.999 +0000 UTC
carbon.Parse("2020-01-01").SetDateTimeMilli(2019, 2, 31, 13, 14, 15, 999).ToString() // 2019-03-03 13:14:15.999 +0000 UTC
// 年、月、日、時、分、秒、マイクロ秒設定
carbon.Parse("2020-01-01").SetDateTimeMicro(2019, 2, 2, 13, 14, 15, 999999).ToString() // 2019-02-02 13:14:15.999999 +0000 UTC
carbon.Parse("2020-01-01").SetDateTimeMicro(2019, 2, 31, 13, 14, 15, 999999).ToString() // 2019-03-03 13:14:15.999999 +0000 UTC
// 年、月、日、時、分、秒、ナノ秒設定
carbon.Parse("2020-01-01").SetDateTimeNano(2019, 2, 2, 13, 14, 15, 999999999).ToString() // 2019-02-02 13:14:15.999999999 +0000 UTC
carbon.Parse("2020-01-01").SetDateTimeNano(2019, 2, 31, 13, 14, 15, 999999999).ToString() // 2019-03-03 13:14:15.999999999 +0000 UTC

// 年、月、日設定
carbon.Parse("2020-01-01").SetDate(2019, 2, 2).ToString() // 2019-02-02 00:00:00 +0000 UTC
carbon.Parse("2020-01-01").SetDate(2019, 2, 31).ToString() // 2019-03-03 00:00:00 +0000 UTC
// 年、月、日、ミリ秒設定
carbon.Parse("2020-01-01").SetDateMilli(2019, 2, 2, 999).ToString() // 2019-02-02 00:00:00.999 +0000 UTC
carbon.Parse("2020-01-01").SetDateMilli(2019, 2, 31, 999).ToString() // 2019-03-03 00:00:00.999 +0000 UTC
// 年、月、日、マイクロ秒設定
carbon.Parse("2020-01-01").SetDateMicro(2019, 2, 2, 999999).ToString() // 2019-02-02 00:00:00.999999 +0000 UTC
carbon.Parse("2020-01-01").SetDateMicro(2019, 2, 31, 999999).ToString() // 2019-03-03 00:00:00.999999 +0000 UTC
// 年、月、日、ナノ秒設定
carbon.Parse("2020-01-01").SetDateNano(2019, 2, 2, 999999999).ToString() // 2019-02-02 00:00:00.999999999 +0000 UTC
carbon.Parse("2020-01-01").SetDateNano(2019, 2, 31, 999999999).ToString() // 2019-03-03 00:00:00.999999999 +0000 UTC

// 時、分、秒設定
carbon.Parse("2020-01-01").SetTime(13, 14, 15).ToString() // 2020-01-01 13:14:15 +0000 UTC
carbon.Parse("2020-01-01").SetTime(13, 14, 90).ToString() // 2020-01-01 13:15:30 +0000 UTC
// 時、分、秒、ミリ秒設定
carbon.Parse("2020-01-01").SetTimeMilli(13, 14, 15, 999).ToString() // 2020-01-01 13:14:15.999 +0000 UTC
carbon.Parse("2020-01-01").SetTimeMilli(13, 14, 90, 999).ToString() // 2020-01-01 13:15:30.999 +0000 UTC
// 時、分、秒、マイクロ秒設定
carbon.Parse("2020-01-01").SetTimeMicro(13, 14, 15, 999999).ToString() // 2020-01-01 13:14:15.999999 +0000 UTC
carbon.Parse("2020-01-01").SetTimeMicro(13, 14, 90, 999999).ToString() // 2020-01-01 13:15:30.999999 +0000 UTC
// 時、分、秒、ナノ秒設定
carbon.Parse("2020-01-01").SetTimeNano(13, 14, 15, 999999999).ToString() // 2020-01-01 13:14:15.999999999 +0000 UTC
carbon.Parse("2020-01-01").SetTimeNano(13, 14, 90, 999999999).ToString() // 2020-01-01 13:15:30.999999999 +0000 UTC

// 年設定
carbon.Parse("2020-02-29").SetYear(2021).ToDateString() // 2021-03-01
// 年設定（月オーバーフローなし）
carbon.Parse("2020-02-29").SetYearNoOverflow(2021).ToDateString() // 2021-02-28

// 月設定
carbon.Parse("2020-01-31").SetMonth(2).ToDateString() // 2020-03-02
// 月設定（月オーバーフローなし）
carbon.Parse("2020-01-31").SetMonthNoOverflow(2).ToDateString() // 2020-02-29

// 週の開始日設定
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Sunday).Week() // 0
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Monday).Week() // 6

// 週の週末日設定
wd := []carbon.Weekday{
  carbon.Saturday, carbon.Sunday,
}
carbon.Parse("2025-04-11").SetWeekendDays(wd).IsWeekend() // false
carbon.Parse("2025-04-12").SetWeekendDays(wd).IsWeekend() // true
carbon.Parse("2025-04-13").SetWeekendDays(wd).IsWeekend() // true

// 日設定
carbon.Parse("2019-08-05").SetDay(31).ToDateString() // 2020-08-31
carbon.Parse("2020-02-01").SetDay(31).ToDateString() // 2020-03-02

// 時設定
carbon.Parse("2020-08-05 13:14:15").SetHour(10).ToDateTimeString() // 2020-08-05 10:14:15
carbon.Parse("2020-08-05 13:14:15").SetHour(24).ToDateTimeString() // 2020-08-06 00:14:15

// 分設定
carbon.Parse("2020-08-05 13:14:15").SetMinute(10).ToDateTimeString() // 2020-08-05 13:10:15
carbon.Parse("2020-08-05 13:14:15").SetMinute(60).ToDateTimeString() // 2020-08-05 14:00:15

// 秒設定
carbon.Parse("2020-08-05 13:14:15").SetSecond(10).ToDateTimeString() // 2020-08-05 13:14:10
carbon.Parse("2020-08-05 13:14:15").SetSecond(60).ToDateTimeString() // 2020-08-05 13:15:00

// ミリ秒設定
carbon.Parse("2020-08-05 13:14:15").SetMillisecond(100).Millisecond() // 100
carbon.Parse("2020-08-05 13:14:15").SetMillisecond(999).Millisecond() // 999

// マイクロ秒設定
carbon.Parse("2020-08-05 13:14:15").SetMicrosecond(100000).Microsecond() // 100000
carbon.Parse("2020-08-05 13:14:15").SetMicrosecond(999999).Microsecond() // 999999

// ナノ秒設定
carbon.Parse("2020-08-05 13:14:15").SetNanosecond(100000000).Nanosecond() // 100000000
carbon.Parse("2020-08-05 13:14:15").SetNanosecond(999999999).Nanosecond() // 999999999
```