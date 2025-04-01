# ペルシア暦（イラン暦）

日本語 | [English](README.md) | [简体中文](README.cn.md)

#### 使い方の例

##### `西暦` を `ペルシャ暦` に変換

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

##### ペルシャ暦を西暦に変換する

```go
carbon.CreateFromPersian(1, 1, 1).ToDateTimeString() // 2016-03-20 00:00:00
carbon.CreateFromPersian(622, 1, 1).ToDateTimeString() // 1243-03-21 00:00:00
carbon.CreateFromPersian(1395, 1, 1).ToDateTimeString() // 2016-03-20 00:00:00
carbon.CreateFromPersian(9377, 1, 1).ToDateTimeString() // 9998-03-19 00:00:00
```

##### 日付判断
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