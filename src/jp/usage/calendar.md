# カレンダ＃カレンダ

現在サポートされているカレンダーは
- 西暦/グレゴリオ暦
- ユリウスの日/簡略化ユリウスの日
- 中国の旧暦
- ペルシャ暦/イラン暦

## ユリウスの日/簡略化ユリウスの日

### `西暦` を `儒略日` に変換する
```go
// デフォルトの保持 6 ビット小数点精度
carbon.Parse("2024-01-24 12:00:00").Julian().JD() // 2460334
carbon.Parse("2024-01-24 13:14:15").Julian().JD() // 2460334.051563

// 4 ビット小数点精度の保持
carbon.Parse("2024-01-24 12:00:00").Julian().JD(4) // 2460334
carbon.Parse("2024-01-24 13:14:15").Julian().JD(4) // 2460334.0516
```

### `西暦` を `簡略儒略日` に変換する
```go
// デフォルトの保持 6 ビット小数点精度
carbon.Parse("2024-01-24 12:00:00").Julian().MJD() // 60333.5
carbon.Parse("2024-01-24 13:14:15").Julian().MJD() // 60333.551563

// 4 ビット小数点精度の保持
carbon.Parse("2024-01-24 12:00:00").Julian().MJD(4) // 60333.5
carbon.Parse("2024-01-24 13:14:15").Julian().MJD(4) // 60333.5516
```

### `儒略日` を `簡略儒略日` に変換する
```go
// デフォルトの保持 6 ビット小数点精度
carbon.CreateFromJulian(2460334).Julian().MJD() // 60333.5
carbon.CreateFromJulian(2460334.051563).Julian().MJD() // 60332.551563

// 4 ビット小数点精度の保持
carbon.CreateFromJulian(2460334).Julian().MJD(4) // 60333.5
carbon.CreateFromJulian(2460334.051563).Julian().MJD(4) // 60332.5516
```

### `簡略儒略日` を `儒略日` に変換する
```go
// デフォルトの保持 6 ビット小数点精度
carbon.CreateFromJulian(60333.5).Julian().JD()() // 2460334
carbon.CreateFromJulian(60333.551563).Julian().JD()() // 2460333.051563

// 4 ビット小数点精度の保持
carbon.CreateFromJulian(60333.5).Julian().JD(4) // 2460334
carbon.CreateFromJulian(60333.551563).Julian().JD(4) // 2460333.0516
```

### `儒略日`/`簡略儒略日` を `公历` に変換する
```go
// 儒略日 を 公历 に変換する
carbon.CreateFromJulian(2460334).ToDateTimeString() // 2024-01-24 12:00:00
carbon.CreateFromJulian(2460334.051563).ToDateTimeString() // 2024-01-24 13:14:15

// 簡略儒略日 を 公历 に変換する
carbon.CreateFromJulian(60333.5).ToDateTimeString() // 2024-01-24 12:00:00
carbon.CreateFromJulian(60333.551563).ToDateTimeString() // 2024-01-24 13:14:15
```

## 中国の旧暦
> 現在は西暦` 1900 `年から` 2100 `年までの` 200 `年の時間スパンのみをサポートしている

### `西暦`を`旧暦`に変換する

```go
// 旧暦の干支を手に入れる
carbon.Parse("2020-08-05").Lunar().Animal() // 鼠
// 旧暦の祝日を取得する
carbon.Parse("2021-02-12").Lunar().Festival() // 春节

// 旧正月の取得
carbon.Parse("2020-08-05").Lunar().Year() // 2020
// 旧暦月の取得
carbon.Parse("2020-08-05").Lunar().Month() // 6
// 旧暦うるう月の取得
carbon.Parse("2020-08-05").Lunar().LeapMonth() // 4
// 旧暦日の取得
carbon.Parse("2020-08-05").Lunar().Day() // 16
// 旧暦時刻の取得
carbon.Parse("2020-08-05").Lunar().Hour() // 13
// 旧暦分の取得
carbon.Parse("2020-08-05").Lunar().Minute() // 14
// 旧暦の取得秒数
carbon.Parse("2020-08-05").Lunar().Second() // 15

// 旧暦日時文字列の取得
carbon.Parse("2020-08-05").Lunar().String() // 2020-06-16
fmt.Printf("%s", carbon.Parse("2020-08-05").Lunar()) // 2020-06-16
// 旧正月文字列の取得
carbon.Parse("2020-08-05").Lunar().ToYearString() // 二零二零
// 旧暦月文字列の取得
carbon.Parse("2020-08-05").Lunar().ToMonthString() // 六月
// 旧暦うるう月文字列の取得
carbon.Parse("2020-04-23").Lunar().ToMonthString() // 闰四月
// 旧暦週文字列の取得
carbon.Parse("2020-04-23").Lunar().ToWeekString() // 周四
// 旧暦日文字列の取得
carbon.Parse("2020-08-05").Lunar().ToDayString() // 十六
// 旧暦日付文字列の取得
carbon.Parse("2020-08-05").Lunar().ToDateString() // 二零二零年六月十六
```

### `旧暦`を`西暦`に変換する
```go
// 2023 年の旧暦 12 月 11 日をグレゴリオ暦に変換します
carbon.CreateFromLunar(2023, 12, 11, 0, 0, 0, false).ToDateTimeString() // 2024-01-21 00:00:00
// 旧暦の 2023 年 2 月 11 日をグレゴリオ暦に変換します
carbon.CreateFromLunar(2023, 2, 11, 0, 0, 0, false).ToDateTimeString() // 2023-03-02 00:00:00
// 旧暦 2023 年、閏 2 月 11 日をグレゴリオ暦に変換します
carbon.CreateFromLunar(2023, 2, 11, 0, 0, 0, true).ToDateTimeString() // 2023-04-01 00:00:00
```

### 日付判断
```go
// 合法的なペルシャ暦の日付かどうか
carbon.Parse("0000-00-00").Lunar().IsValid() // false
carbon.Parse("2020-08-05").Lunar().IsValid() // true

// 旧暦うるう年かどうか
carbon.Parse("2020-08-05").Lunar().IsLeapYear() // true
// 旧暦うるう月かどうか
carbon.Parse("2020-08-05").Lunar().IsLeapMonth() // false

// ねずみ年かどうか
carbon.Parse("2020-08-05").Lunar().IsRatYear() // true
// 牛年かどうか
carbon.Parse("2020-08-05").Lunar().IsOxYear() // false
// 寅年かどうか
carbon.Parse("2020-08-05").Lunar().IsTigerYear() // false
// うさぎ年かどうか
carbon.Parse("2020-08-05").Lunar().IsRabbitYear() // false
// 龍年かどうか
carbon.Parse("2020-08-05").Lunar().IsDragonYear() // false
// 蛇の年かどうか
carbon.Parse("2020-08-05").Lunar().IsSnakeYear() // false
// 馬年かどうか
carbon.Parse("2020-08-05").Lunar().IsHorseYear() // false
// 羊年かどうか
carbon.Parse("2020-08-05").Lunar().IsGoatYear() // false
// 申年かどうか
carbon.Parse("2020-08-05").Lunar().IsMonkeyYear() // false
// 鶏の年かどうか
carbon.Parse("2020-08-05").Lunar().IsRoosterYear() // false
// 犬年かどうか
carbon.Parse("2020-08-05").Lunar().IsDogYear() // false
// 豚年かどうか
carbon.Parse("2020-08-05").Lunar().IsPigYear() // false
```

## ペルシャ暦/イラン暦

### `西暦` を `ペルシャ暦` に変換
```go
// ペルシャ暦の取得
carbon.Parse("2020-08-05").Persian().Year() // 1399
// ペルシャ暦月の取得
carbon.Parse("2020-08-05").Persian().Month() // 5
// ペルシャ暦の取得日
carbon.Parse("2020-08-05").Persian().Day() // 15
// ペルシャ暦時間の取得
carbon.Parse("2020-08-05").Persian().Hour() // 13
// ペルシャ暦分の取得
carbon.Parse("2020-08-05").Persian().Minute() // 14
// ペルシャ暦秒の取得
carbon.Parse("2020-08-05").Persian().Second() // 15

// ペルシャ暦日時文字列の取得
carbon.Parse("2020-08-05").Persian().String() // 1399-05-15
fmt.Printf("%s", carbon.Parse("2020-08-05").Persian()) // 1399-05-15

// ペルシア暦月文字列の取得
carbon.Parse("2020-08-05").Persian().ToMonthString() // Mordad
carbon.Parse("2020-08-05").Persian().ToMonthString("en") // Mordad
carbon.Parse("2020-08-05").Persian().ToMonthString("fa") // مرداد

// 略語ペルシャ暦文字列の取得
carbon.Parse("2020-08-05").Persian().ToShortMonthString() // Mor
carbon.Parse("2020-08-05").Persian().ToShortMonthString("en") // Mor
carbon.Parse("2020-08-05").Persian().ToShortMonthString("fa") // مرد

// ペルシャ暦週文字列の取得
carbon.Parse("2020-08-05").Persian().ToWeekString() // Chaharshanbeh
carbon.Parse("2020-08-05").Persian().ToWeekString("en") // Chaharshanbeh
carbon.Parse("2020-08-05").Persian().ToWeekString("fa") // چهارشنبه

// 略語ペルシャ暦週文字列の取得
carbon.Parse("2020-08-05").Persian().ToShortWeekString() // Cha
carbon.Parse("2020-08-05").Persian().ToShortWeekString("en") // Cha
carbon.Parse("2020-08-05").Persian().ToShortWeekString("fa") // د
```

### `ペルシャ暦` を `西暦` に変換する
```go
carbon.CreateFromPersian(1, 1, 1).ToDateTimeString() // 2016-03-20 00:00:00
carbon.CreateFromPersian(622, 1, 1).ToDateTimeString() // 1243-03-21 00:00:00
carbon.CreateFromPersian(1395, 1, 1).ToDateTimeString() // 2016-03-20 00:00:00
carbon.CreateFromPersian(9377, 1, 1).ToDateTimeString() // 9998-03-19 00:00:00
```

### 日付判断
```go
// 合法的なペルシャ暦の日付かどうか
carbon.CreateFromPersian(1, 1, 1).IsValid() // true
carbon.CreateFromPersian(622, 1, 1).IsValid() // true
carbon.CreateFromPersian(9377, 1, 1).IsValid() // true
carbon.CreateFromPersian(0, 0, 0, 0).IsValid() // false
carbon.CreateFromPersian(2024, 0, 1).IsValid() // false
carbon.CreateFromPersian(2024, 1, 0).IsValid() // false

// ペルシア暦閏年かどうか
carbon.CreateFromPersian(1395, 1, 1).IsLeapYear() // true
carbon.CreateFromPersian(9377, 1, 1).IsLeapYear() // true
carbon.CreateFromPersian(622, 1, 1).IsLeapYear() // false
carbon.CreateFromPersian(9999, 1, 1).IsLeapYear() // false
```