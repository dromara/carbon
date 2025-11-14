---
head:
  - - meta
    - name: description
      content: ユリウス日/修正ユリウス日、中国旧暦、ペルシャ暦、ヘブライ暦をサポートし、グレゴリオ暦との相互変換、フォーマット（月/週/日文字列）、妥当性および閏年判定を提供
  - - meta
    - name: keywords
      content: carbon, go-carbon, カレンダー, グレゴリオ暦, ユリウス日, 修正ユリウス日, 中国旧暦, ペルシャ暦, イラン暦, ヘブライ暦, ユダヤ暦
---

# 日历

現在サポートされているカレンダーは

- [ユリウス日/簡略化ユリウス日](#ユリウス日-簡略化ユリウス日)
- [中国旧暦](#中国旧暦)
- [ペルシャ暦（イラン暦）](#ペルシャ暦-イラン暦)
- [ヘブライ暦（ユダヤ暦）](#ヘブライ暦-ユダヤ暦)

## ユリウス日/簡略化ユリウス日

### `西暦` を `ユリウス日` に変換する
```go
// デフォルトで6桁の小数精度を保持
carbon.Parse("2024-01-24 12:00:00").Julian().JD() // 2460334
carbon.Parse("2024-01-24 13:14:15").Julian().JD() // 2460334.051563

// 4桁の小数精度を保持
carbon.Parse("2024-01-24 12:00:00").Julian().JD(4) // 2460334
carbon.Parse("2024-01-24 13:14:15").Julian().JD(4) // 2460334.0516
```

### `西暦` を `簡略化ユリウス日` に変換する
```go
// デフォルトで6桁の小数精度を保持
carbon.Parse("2024-01-24 12:00:00").Julian().MJD() // 60333.5
carbon.Parse("2024-01-24 13:14:15").Julian().MJD() // 60333.551563

// 4桁の小数精度を保持
carbon.Parse("2024-01-24 12:00:00").Julian().MJD(4) // 60333.5
carbon.Parse("2024-01-24 13:14:15").Julian().MJD(4) // 60333.5516
```

### `ユリウス日` を `簡略化ユリウス日` に変換する
```go
// デフォルトで6桁の小数精度を保持
carbon.CreateFromJulian(2460334).Julian().MJD() // 60333.5
carbon.CreateFromJulian(2460334.051563).Julian().MJD() // 60332.551563

// 4桁の小数精度を保持
carbon.CreateFromJulian(2460334).Julian().MJD(4) // 60333.5
carbon.CreateFromJulian(2460334.051563).Julian().MJD(4) // 60332.5516
```

### `簡略化ユリウス日` を `ユリウス日` に変換する
```go
// デフォルトで6桁の小数精度を保持
carbon.CreateFromJulian(60333.5).Julian().JD() // 2460334
carbon.CreateFromJulian(60333.551563).Julian().JD() // 2460333.051563

// 4桁の小数精度を保持
carbon.CreateFromJulian(60333.5).Julian().JD(4) // 2460334
carbon.CreateFromJulian(60333.551563).Julian().JD(4) // 2460333.0516
```

### `ユリウス日`/`簡略化ユリウス日` を `西暦` に変換する
```go
// ユリウス日を西暦に変換する
carbon.CreateFromJulian(2460334).ToDateTimeString() // 2024-01-24 12:00:00
carbon.CreateFromJulian(2460334.051563).ToDateTimeString() // 2024-01-24 13:14:15

// 簡略化ユリウス日を西暦に変換する
carbon.CreateFromJulian(60333.5).ToDateTimeString() // 2024-01-24 12:00:00
carbon.CreateFromJulian(60333.551563).ToDateTimeString() // 2024-01-24 13:14:15
```

## 中国旧暦
> 現在は西暦`1900`年から`2100`年までの`200`年の時間スパンのみをサポートしている

### `西暦`を`旧暦`に変換する
```go
// 旧暦の干支を取得する
carbon.Parse("2020-08-05").Lunar().Animal() // 鼠
// 旧暦の祝日を取得する
carbon.Parse("2021-02-12").Lunar().Festival() // 春节

// 旧暦年を取得
carbon.Parse("2020-08-05").Lunar().Year() // 2020
// 旧暦月を取得
carbon.Parse("2020-08-05").Lunar().Month() // 6
// 旧暦閏月を取得
carbon.Parse("2020-08-05").Lunar().LeapMonth() // 4
// 旧暦日を取得
carbon.Parse("2020-08-05").Lunar().Day() // 16

// 旧暦日時文字列を取得
carbon.Parse("2020-08-05").Lunar().String() // 2020-06-16
fmt.Printf("%s", carbon.Parse("2020-08-05").Lunar()) // 2020-06-16
// 旧暦年文字列を取得
carbon.Parse("2020-08-05").Lunar().ToYearString() // 二零二零
// 旧暦月文字列を取得
carbon.Parse("2020-08-05").Lunar().ToMonthString() // 六月
// 旧暦閏月文字列を取得
carbon.Parse("2020-04-23").Lunar().ToMonthString() // 闰四月
// 旧暦週文字列を取得
carbon.Parse("2020-04-23").Lunar().ToWeekString() // 周四
// 旧暦日文字列を取得
carbon.Parse("2020-08-05").Lunar().ToDayString() // 十六
// 旧暦日付文字列を取得
carbon.Parse("2020-08-05").Lunar().ToDateString() // 二零二零年六月十六
```

### `旧暦`を`西暦`に変換する
```go
// 旧暦2023年12月11日を西暦に変換する
carbon.CreateFromLunar(2023, 12, 11, false).ToDateString() // 2024-01-21
// 旧暦2023年2月11日を西暦に変換する
carbon.CreateFromLunar(2023, 2, 11, false).ToDateString() // 2023-03-02
// 旧暦2023年閏2月11日を西暦に変換する
carbon.CreateFromLunar(2023, 2, 11, true).ToDateString() // 2023-04-01
```

### 日付判定
```go
// 有効な旧暦日付かどうか
carbon.Parse("0000-00-00").Lunar().IsValid() // false
carbon.Parse("2020-08-05").Lunar().IsValid() // true

// 旧暦閏年かどうか
carbon.Parse("2020-08-05").Lunar().IsLeapYear() // true
// 旧暦閏月かどうか
carbon.Parse("2020-08-05").Lunar().IsLeapMonth() // false

// 鼠年かどうか
carbon.Parse("2020-08-05").Lunar().IsRatYear() // true
// 牛年かどうか
carbon.Parse("2020-08-05").Lunar().IsOxYear() // false
// 虎年かどうか
carbon.Parse("2020-08-05").Lunar().IsTigerYear() // false
// 兔年かどうか
carbon.Parse("2020-08-05").Lunar().IsRabbitYear() // false
// 龍年かどうか
carbon.Parse("2020-08-05").Lunar().IsDragonYear() // false
// 蛇年かどうか
carbon.Parse("2020-08-05").Lunar().IsSnakeYear() // false
// 馬年かどうか
carbon.Parse("2020-08-05").Lunar().IsHorseYear() // false
// 羊年かどうか
carbon.Parse("2020-08-05").Lunar().IsGoatYear() // false
// 猴年かどうか
carbon.Parse("2020-08-05").Lunar().IsMonkeyYear() // false
// 鶏年かどうか
carbon.Parse("2020-08-05").Lunar().IsRoosterYear() // false
// 狗年かどうか
carbon.Parse("2020-08-05").Lunar().IsDogYear() // false
// 猪年かどうか
carbon.Parse("2020-08-05").Lunar().IsPigYear() // false
```

## ペルシャ暦（イラン暦）

### `西暦` を `ペルシャ暦` に変換する
```go
// ペルシャ暦年を取得
carbon.Parse("2020-08-05").Persian().Year() // 1399
// ペルシャ暦月を取得
carbon.Parse("2020-08-05").Persian().Month() // 5
// ペルシャ暦日を取得
carbon.Parse("2020-08-05").Persian().Day() // 15

// ペルシャ暦日時文字列を取得
carbon.Parse("2020-08-05").Persian().String() // 1399-05-15
fmt.Printf("%s", carbon.Parse("2020-08-05").Persian()) // 1399-05-15

// ペルシャ暦月文字列を取得
carbon.Parse("2020-08-05").Persian().ToMonthString() // Mordad
carbon.Parse("2020-08-05").Persian().ToMonthString("en") // Mordad
carbon.Parse("2020-08-05").Persian().ToMonthString("fa") // مرداد

// ペルシャ暦週文字列を取得
carbon.Parse("2020-08-05").Persian().ToWeekString() // Chaharshanbeh
carbon.Parse("2020-08-05").Persian().ToWeekString("en") // Chaharshanbeh
carbon.Parse("2020-08-05").Persian().ToWeekString("fa") // چهارشنبه
```

### `ペルシャ暦` を `西暦` に変換する
```go
carbon.CreateFromPersian(1, 1, 1).ToDateString() // 2016-03-20
carbon.CreateFromPersian(622, 1, 1).ToDateString() // 1243-03-21
carbon.CreateFromPersian(1395, 1, 1).ToDateString() // 2016-03-20
carbon.CreateFromPersian(9377, 1, 1).ToDateString() // 9998-03-19
```

### 日付判定
```go
// 有効なペルシャ暦日付かどうか
carbon.CreateFromPersian(1, 1, 1).IsValid() // true
carbon.CreateFromPersian(622, 1, 1).IsValid() // true
carbon.CreateFromPersian(9377, 1, 1).IsValid() // true
carbon.CreateFromPersian(0, 0, 0, 0).IsValid() // false
carbon.CreateFromPersian(2024, 0, 1).IsValid() // false
carbon.CreateFromPersian(2024, 1, 0).IsValid() // false

// ペルシャ暦閏年かどうか
carbon.CreateFromPersian(1395, 1, 1).IsLeapYear() // true
carbon.CreateFromPersian(9377, 1, 1).IsLeapYear() // true
carbon.CreateFromPersian(622, 1, 1).IsLeapYear() // false
carbon.CreateFromPersian(9999, 1, 1).IsLeapYear() // false
```

## ヘブライ暦（ユダヤ暦）

### `西暦` を `ヘブライ暦` に変換する
```go
// ヘブライ暦年を取得
carbon.Parse("2024-01-01").Hebrew().Year() // 5784
// ヘブライ暦月を取得
carbon.Parse("2024-01-01").Hebrew().Month() // 10
// ヘブライ暦日を取得
carbon.Parse("2024-01-01").Hebrew().Day() // 20

// ヘブライ暦日時文字列を取得
carbon.Parse("2024-01-01").Hebrew().String() // 5784-10-20
fmt.Printf("%s", carbon.Parse("2024-01-01").Hebrew()) // 5784-10-20

// ヘブライ暦月文字列を取得
carbon.Parse("2020-08-05").Hebrew().ToMonthString() // Av
carbon.Parse("2020-08-05").Hebrew().ToMonthString("en") // Av
carbon.Parse("2020-08-05").Hebrew().ToMonthString("he") // אב

// ヘブライ暦週文字列を取得
carbon.Parse("2020-08-05").Hebrew().ToWeekString() // Wednesday
carbon.Parse("2020-08-05").Hebrew().ToWeekString("en") // Wednesday
carbon.Parse("2020-08-05").Hebrew().ToWeekString("he") // רביעי
```

### `ヘブライ暦` を `西暦` に変換する
```go
carbon.CreateFromHebrew(5784, 10, 20).ToDateString() // 2023-12-17
carbon.CreateFromHebrew(5784, 5, 1).ToDateString() // 2024-07-21
carbon.CreateFromHebrew(5786, 7, 10).ToDateString() // 2025-09-18
```

### 日付判定
```go
// 有効なヘブライ暦日付かどうか
carbon.CreateFromHebrew(5780, 14, 1).IsValid() // false
carbon.CreateFromHebrew(5780, 10, 30).IsValid() // true

// ヘブライ暦閏年かどうか
carbon.CreateFromHebrew(5784, 1, 1).IsLeapYear() // true
carbon.CreateFromHebrew(5788, 1, 1).IsLeapYear() // false
```