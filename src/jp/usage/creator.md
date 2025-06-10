#### `Carbon` インスタンスの作成

## タイムスタンプからの `Carbon` インスタンスの作成
```go
// 秒タイムスタンプから Carbon インスタンスを作成します
carbon.CreateFromTimestamp(-1).ToString() // 1970-01-01 08:59:59 +0900 JST
carbon.CreateFromTimestamp(0).ToString() // 1970-01-01 09:00:00 +0900 JST
carbon.CreateFromTimestamp(1).ToString() // 1970-01-01 09:00:01 +0900 JST
carbon.CreateFromTimestamp(1596600855).ToString() // 2020-08-05 13:14:15 +0900 JST
// ミリ秒のタイムスタンプから Carbon インスタンスを作成します
carbon.CreateFromTimestampMilli(1596600855999).ToString() // 2020-08-05 13:14:15.999 +0900 JST
// マイクロ秒タイムスタンプから Carbon インスタンスを作成します
carbon.CreateFromTimestampMicro(1596600855999999).ToString() // 2020-08-05 13:14:15.999999 +0900 JST
// ナノタイムスタンプから Carbon インスタンスを作成します
carbon.CreateFromTimestampNano(1596600855999999999).ToString() // 2020-08-05 13:14:15.999999999 +0900 JST
```

## 日付からの `Carbon` インスタンスの作成
```go
// 年、月、日、時、分、秒からCarbonインスタンスを作成する
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
// 年、月、日、時、分、秒、ミリ秒から Carbon インスタンスを作成します
carbon.CreateFromDateTimeMilli(2020, 1, 1, 13, 14, 15, 999).ToString() // 2020-01-01 13:14:15.999 +0900 JST
// 年、月、日、時、分、秒、マイクロ秒から Carbon インスタンスを作成します
carbon.CreateFromDateTimeMicro(2020, 1, 1, 13, 14, 15, 999999).ToString() // 2020-01-01 13:14:15.999999 +0900 JST
// 年、月、日、時、分、秒、ナノ秒からCarbonインスタンスを作成する
carbon.CreateFromDateTimeNano(2020, 1, 1, 13, 14, 15, 999999999).ToString() // 2020-01-01 13:14:15.999999999 +0900 JST

// 年、月、日からCarbonインスタンスを作成する
carbon.CreateFromDate(2020, 8, 5).ToString() // 2020-08-05 00:00:00 +0900 JST
// 年、月、日、ミリ秒からCarbonインスタンスを作成する
carbon.CreateFromDateMilli(2020, 8, 5, 999).ToString() // 2020-08-05 00:00:00.999 +0900 JST
// 年、月、日、マイクロ秒からCarbonインスタンスを作成する
carbon.CreateFromDateMicro(2020, 8, 5, 999999).ToString() // 2020-08-05 00:00:00.999999 +0900 JST
// 年、月、日、ナノ秒からCarbonインスタンスを作成する
carbon.CreateFromDateNano(2020, 8, 5, 999999999).ToString() // 2020-08-05 00:00:00.999999999 +0900 JST

// 時間、分、秒から Carbon インスタンスを作成します (年、月、日はデフォルトで現在の年、月、日になります)
carbon.CreateFromTime(13, 14, 15).ToString() // 2020-08-05 13:14:15 +0900 JST
// 時間、分、秒、ミリ秒から Carbon インスタンスを作成します (年、月、日はデフォルトで現在の年、月、日になります)
carbon.CreateFromTimeMilli(13, 14, 15, 999).ToString() // 2020-08-05 13:14:15.999 +0900 JST
// 時間、分、秒、マイクロ秒から Carbon インスタンスを作成します (年、月、日はデフォルトで現在の年、月、日になります)
carbon.CreateFromTimeMicro(13, 14, 15, 999999).ToString() // 2020-08-05 13:14:15.999999 +0900 JST
// 時間、分、秒、ナノ秒から Carbon インスタンスを作成します (年、月、日はデフォルトで現在の年、月、日になります)
carbon.CreateFromTimeNano(13, 14, 15, 999999999).ToString() // 2020-08-05 13:14:15.999999999 +0900 JST
```