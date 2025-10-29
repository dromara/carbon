---
head:
  - - meta
    - name: description
      content: Carbon インスタンスの作成 | タイムスタンプ(秒/ミリ秒/マイクロ秒/ナノ秒)および 日時/日付/時間 から作成：CreateFromTimestamp/DateTime/Date/Time と各精度バリアントを提供
  - - meta
    - name: keywords
      content: carbon, go-carbon, タイムスタンプからCarbonインスタンス作成, 日付からCarbonインスタンス作成
---

# `Carbon` インスタンスの作成

## タイムスタンプからの `Carbon` インスタンスの作成
```go
// 秒タイムスタンプから Carbon インスタンスを作成します
carbon.CreateFromTimestamp(-1).ToString() // 1969-12-31 23:59:59 +0000 UTC
carbon.CreateFromTimestamp(0).ToString() // 1970-01-01 00:00:00 +0000 UTC
carbon.CreateFromTimestamp(1).ToString() // 1970-01-01 00:00:01 +0000 UTC
carbon.CreateFromTimestamp(1596633255).ToString() // 2020-08-05 13:14:15 +0000 UTC
// ミリ秒のタイムスタンプから Carbon インスタンスを作成します
carbon.CreateFromTimestampMilli(1596604455999).ToString() // 2020-08-05 13:14:15.999 +0000 UTC
// マイクロ秒タイムスタンプから Carbon インスタンスを作成します
carbon.CreateFromTimestampMicro(1596604455999999).ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
// ナノタイムスタンプから Carbon インスタンスを作成します
carbon.CreateFromTimestampNano(1596604455999999999).ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC
```

## 日付からの `Carbon` インスタンスの作成
```go
// 年、月、日、時、分、秒からCarbonインスタンスを作成する
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToString() // 2020-08-05 13:14:15 +0000 UTC
// 年、月、日、時、分、秒、ミリ秒から Carbon インスタンスを作成します
carbon.CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString() // 2020-08-05 13:14:15.999 +0000 UTC
// 年、月、日、時、分、秒、マイクロ秒から Carbon インスタンスを作成します
carbon.CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
// 年、月、日、時、分、秒、ナノ秒からCarbonインスタンスを作成する
carbon.CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC

// 年、月、日からCarbonインスタンスを作成する
carbon.CreateFromDate(2020, 8, 5).ToString() // 2020-08-05 00:00:00 +0000 UTC
// 年、月、日、ミリ秒からCarbonインスタンスを作成する
carbon.CreateFromDateMilli(2020, 8, 5, 999).ToString() // 2020-08-05 00:00:00.999 +0000 UTC
// 年、月、日、マイクロ秒からCarbonインスタンスを作成する
carbon.CreateFromDateMicro(2020, 8, 5, 999999).ToString() // 2020-08-05 00:00:00.999999 +0000 UTC
// 年、月、日、ナノ秒からCarbonインスタンスを作成する
carbon.CreateFromDateNano(2020, 8, 5, 999999999).ToString() // 2020-08-05 00:00:00.999999999 +0000 UTC

// 時間、分、秒から Carbon インスタンスを作成します (年、月、日はデフォルトで現在の年、月、日になります)
carbon.CreateFromTime(13, 14, 15).ToString() // 2020-08-05 13:14:15 +0000 UTC
// 時間、分、秒、ミリ秒から Carbon インスタンスを作成します (年、月、日はデフォルトで現在の年、月、日になります)
carbon.CreateFromTimeMilli(13, 14, 15, 999).ToString() // 2020-08-05 13:14:15.999 +0000 UTC
// 時間、分、秒、マイクロ秒から Carbon インスタンスを作成します (年、月、日はデフォルトで現在の年、月、日になります)
carbon.CreateFromTimeMicro(13, 14, 15, 999999).ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
// 時間、分、秒、ナノ秒から Carbon インスタンスを作成します (年、月、日はデフォルトで現在の年、月、日になります)
carbon.CreateFromTimeNano(13, 14, 15, 999999999).ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC
```