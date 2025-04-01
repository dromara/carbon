# 波斯历(伊朗历)

简体中文 | [English](README.md) | [日本語](README.jp.md)

#### 用法示例

##### 将 `公历` 转换成 `波斯历`

```go
// 获取波斯历年份
carbon.Parse("2020-08-05").Persian().Year() // 1399
// 获取波斯历月份
carbon.Parse("2020-08-05").Persian().Month() // 5
// 获取波斯历日期
carbon.Parse("2020-08-05").Persian().Day() // 15
// 获取波斯历小时
carbon.Parse("2020-08-05").Persian().Hour() // 13
// 获取波斯历分钟
carbon.Parse("2020-08-05").Persian().Minute() // 14
// 获取波斯历秒数
carbon.Parse("2020-08-05").Persian().Second() // 15

// 获取波斯历日期时间字符串
carbon.Parse("2020-08-05").Persian().String() // 1399-05-15
fmt.Printf("%s", carbon.Parse("2020-08-05").Persian()) // 1399-05-15

// 获取波斯历月字符串
carbon.Parse("2020-08-05").Persian().ToMonthString() // Mordad
carbon.Parse("2020-08-05").Persian().ToMonthString("en") // Mordad
carbon.Parse("2020-08-05").Persian().ToMonthString("fa") // مرداد

// 获取简写波斯历月字符串
carbon.Parse("2020-08-05").Persian().ToShortMonthString() // Mor
carbon.Parse("2020-08-05").Persian().ToShortMonthString("en") // Mor
carbon.Parse("2020-08-05").Persian().ToShortMonthString("fa") // مرد

// 获取波斯历周字符串
carbon.Parse("2020-08-05").Persian().ToWeekString() // Chaharshanbeh
carbon.Parse("2020-08-05").Persian().ToWeekString("en") // Chaharshanbeh
carbon.Parse("2020-08-05").Persian().ToWeekString("fa") // چهارشنبه

// 获取简写波斯历周字符串
carbon.Parse("2020-08-05").Persian().ToShortWeekString() // Cha
carbon.Parse("2020-08-05").Persian().ToShortWeekString("en") // Cha
carbon.Parse("2020-08-05").Persian().ToShortWeekString("fa") // د

```

##### 将 `波斯历` 转化成 `公历`

```go
carbon.CreateFromPersian(1, 1, 1).ToDateTimeString() // 2016-03-20 00:00:00
carbon.CreateFromPersian(622, 1, 1).ToDateTimeString() // 1243-03-21 00:00:00
carbon.CreateFromPersian(1395, 1, 1).ToDateTimeString() // 2016-03-20 00:00:00
carbon.CreateFromPersian(9377, 1, 1).ToDateTimeString() // 9998-03-19 00:00:00
```

##### 日期判断
```go
// 是否是合法的波斯历日期
carbon.CreateFromPersian(1, 1, 1).IsValid() // true
carbon.CreateFromPersian(622, 1, 1).IsValid() // true
carbon.CreateFromPersian(9377, 1, 1).IsValid() // true
carbon.CreateFromPersian(0, 0, 0, 0).IsValid() // false
carbon.CreateFromPersian(2024, 0, 1).IsValid() // false
carbon.CreateFromPersian(2024, 1, 0).IsValid() // false

// 是否是波斯历闰年
carbon.CreateFromPersian(1395, 1, 1).IsLeapYear() // true
carbon.CreateFromPersian(9377, 1, 1).IsLeapYear() // true
carbon.CreateFromPersian(622, 1, 1).IsLeapYear() // false
carbon.CreateFromPersian(9999, 1, 1).IsLeapYear() // false

```