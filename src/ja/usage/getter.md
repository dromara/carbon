---
head:
  - - meta
    - name: description
      content: 時間取得|軽量で、意味的に、開発者に優しい golang 時間処理ライブラリ
  - - meta
    - name: keywords
      content: carbon, go-carbon, 時間取得
---

# 時間取得

```go
// 年の総日数を取得
carbon.Parse("2019-08-05 13:14:15").DaysInYear() // 365
carbon.Parse("2020-08-05 13:14:15").DaysInYear() // 366
// 月の総日数を取得
carbon.Parse("2020-02-01 13:14:15").DaysInMonth() // 29
carbon.Parse("2020-04-01 13:14:15").DaysInMonth() // 30
carbon.Parse("2020-08-01 13:14:15").DaysInMonth() // 31

// 年間積算日を取得
carbon.Parse("2020-08-05 13:14:15").DayOfYear() // 218
// 本年の何週目かを取得
carbon.Parse("2019-12-31 13:14:15").WeekOfYear() // 1
carbon.Parse("2020-08-05 13:14:15").WeekOfYear() // 32
// 今月の何日目（1から）かを取得
carbon.Parse("2020-08-05 13:14:15").DayOfMonth() // 5
// 今月の何週目かを取得
carbon.Parse("2020-08-05 13:14:15").WeekOfMonth() // 1
// 今週の何日目かを取得(1が月曜日)
carbon.Parse("2020-08-05 13:14:15").DayOfWeek() // 3

// 現在の年月日時分秒を取得
carbon.Parse("2020-08-05 13:14:15").DateTime() // 2020,8,5,13,14,15
// 現在の年月日時分秒ミリ秒を取得
carbon.Parse("2020-08-05 13:14:15").DateTimeMilli() // 2020,8,5,13,14,15,999
// 現在の年月日時分秒マイクロ秒を取得
carbon.Parse("2020-08-05 13:14:15").DateTimeMicro() // 2020,8,5,13,14,15,999999
// 現在の年月日時分秒ナノ秒を取得
carbon.Parse("2020-08-05 13:14:15").DateTimeNano() // 2020,8,5,13,14,15,999999999

// 現在の年月日を取得
carbon.Parse("2020-08-05 13:14:15.999999999").Date() // 2020,8,5
// 現在の年月日ミリ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").DateMilli() // 2020,8,5,999
// 現在の年月日マイクロ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").DateMicro() // 2020,8,5,999999
// 現在の年月日ナノ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").DateNano() // 2020,8,5,999999999

// 現在の時分秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").Time() // 13,14,15
// 現在の時分秒ミリ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").TimeMilli() // 13,14,15,999
// 現在の時分秒マイクロ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").TimeMicro() // 13,14,15,999999
// 現在の時分秒ナノ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").TimeNano() // 13,14,15,999999999

// 現在の世紀を取得
carbon.Parse("2020-08-05 13:14:15").Century() // 21
// 現在の十年紀を取得
carbon.Parse("2019-08-05 13:14:15").Decade() // 10
carbon.Parse("2021-08-05 13:14:15").Decade() // 20
// 現在の年を取得
carbon.Parse("2020-08-05 13:14:15").Year() // 2020
// 現在の四半期を取得
carbon.Parse("2020-08-05 13:14:15").Quarter() // 3
// 現在の月を取得
carbon.Parse("2020-08-05 13:14:15").Month() // 8
// 現在の週を取得(0から開始)
carbon.Parse("2020-08-02 13:14:15").Week() // 0
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Sunday).Week() // 0
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Monday).Week() // 6
// 現在の日数を取得
carbon.Parse("2020-08-05 13:14:15").Day() // 5
// 現在の時間を取得
carbon.Parse("2020-08-05 13:14:15").Hour() // 13
// 現在の分を取得
carbon.Parse("2020-08-05 13:14:15").Minute() // 14
// 現在の秒を取得
carbon.Parse("2020-08-05 13:14:15").Second() // 15
// 現在のミリ秒を取得
carbon.Parse("2020-08-05 13:14:15.999").Millisecond() // 999
// 現在のマイクロ秒を取得
carbon.Parse("2020-08-05 13:14:15.999").Microsecond() // 999000
// 現在のナノ秒を取得
carbon.Parse("2020-08-05 13:14:15.999").Nanosecond() // 999000000

// 秒タイムスタンプを取得
carbon.Parse("2020-08-05 13:14:15").Timestamp() // 1596633255
// ミリ秒のタイムスタンプを取得
carbon.Parse("2020-08-05 13:14:15.999").TimestampMilli() // 1596633255999
// マイクロ秒タイムスタンプを取得
carbon.Parse("2020-08-05 13:14:15.999999").TimestampMicro() // 1596633255999999
// ナノ秒タイムスタンプを取得
carbon.Parse("2020-08-05 13:14:15.999999999").TimestampNano() // 1596633255999999999

// タイムゾーンロケーションの取得
carbon.SetTimezone(carbon.PRC).Timezone() // PRC
carbon.SetTimezone(carbon.Tokyo).Timezone() // Asia/Tokyo

// タイムゾーン名の取得
carbon.SetTimezone(carbon.PRC).ZoneName() // CST
carbon.SetTimezone(carbon.Tokyo).ZoneName() // JST

// UTCタイムゾーンオフセットの秒を取得
carbon.SetTimezone(carbon.PRC).ZoneOffset() // 28800
carbon.SetTimezone(carbon.Tokyo).ZoneOffset() // 32400

// ロケール名を取得
carbon.Now().Locale() // zh-CN
carbon.Now().SetLocale("en").Locale() // en

// 星座を取得
carbon.Now().Constellation() // 狮子座
carbon.Now().SetLocale("en").Constellation() // Leo
carbon.Now().SetLocale("ja").Constellation() // しし座

// 季節を取得
carbon.Now().Season() // 夏季
carbon.Now().SetLocale("en").Season() // Summer
carbon.Now().SetLocale("ja").Season() // 夏

// 週の開始日の取得
carbon.SetWeekStartsAt(carbon.Sunday).WeekStartsAt() // Sunday
carbon.SetWeekStartsAt(carbon.Monday).WeekStartsAt() // Monday

// 週の終了日の取得
carbon.SetWeekStartsAt(carbon.Sunday).WeekEndsAt() // Saturday
carbon.SetWeekStartsAt(carbon.Monday).WeekEndsAt() // Sunday

// 現在のレイアウトテンプレートの取得
carbon.Parse("now").CurrentLayout() // "2006-01-02 15:04:05"
carbon.ParseByLayout("2020-08-05", DateLayout).CurrentLayout() // "2006-01-02"

// 年齢を取得
carbon.Parse("2002-01-01 13:14:15").Age() // 17
carbon.Parse("2002-12-31 13:14:15").Age() // 18
```