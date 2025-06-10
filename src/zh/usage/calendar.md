# 日历

目前支持的日历有
- 公历/格里历
- 儒略日/简化儒略日
- 中国农历
- 波斯历(伊朗历)

## 儒略日/简化儒略日

### 将 `公历` 转换成 `儒略日`
```go
// 默认保留 6 位小数精度
carbon.Parse("2024-01-24 12:00:00").Julian().JD() // 2460334
carbon.Parse("2024-01-24 13:14:15").Julian().JD() // 2460334.051563

// 保留 4 位小数精度
carbon.Parse("2024-01-24 12:00:00").Julian().JD(4) // 2460334
carbon.Parse("2024-01-24 13:14:15").Julian().JD(4) // 2460334.0516
```

### 将 `公历` 转换成 `简化儒略日`
```go
// 默认保留 6 位小数精度
carbon.Parse("2024-01-24 12:00:00").Julian().MJD() // 60333.5
carbon.Parse("2024-01-24 13:14:15").Julian().MJD() // 60333.551563

// 保留 4 位小数精度
carbon.Parse("2024-01-24 12:00:00").Julian().MJD(4) // 60333.5
carbon.Parse("2024-01-24 13:14:15").Julian().MJD(4) // 60333.5516
```

### 将 `儒略日` 转换成 `简化儒略日`
```go
// 默认保留 6 位小数精度
carbon.CreateFromJulian(2460334).Julian().MJD() // 60333.5
carbon.CreateFromJulian(2460334.051563).Julian().MJD() // 60332.551563

// 保留 4 位小数精度
carbon.CreateFromJulian(2460334).Julian().MJD(4) // 60333.5
carbon.CreateFromJulian(2460334.051563).Julian().MJD(4) // 60332.5516
```

### 将 `简化儒略日` 转换成 `儒略日`
```go
// 默认保留 6 位小数精度
carbon.CreateFromJulian(60333.5).Julian().JD()() // 2460334
carbon.CreateFromJulian(60333.551563).Julian().JD()() // 2460333.051563

// 保留 4 位小数精度
carbon.CreateFromJulian(60333.5).Julian().JD(4) // 2460334
carbon.CreateFromJulian(60333.551563).Julian().JD(4) // 2460333.0516
```

### 将 `儒略日`/`简化儒略日` 转换成 `公历`
```go
// 将 儒略日 转换成 公历
carbon.CreateFromJulian(2460334).ToDateTimeString() // 2024-01-24 12:00:00
carbon.CreateFromJulian(2460334.051563).ToDateTimeString() // 2024-01-24 13:14:15

// 将 简化儒略日 转换成 公历
carbon.CreateFromJulian(60333.5).ToDateTimeString() // 2024-01-24 12:00:00
carbon.CreateFromJulian(60333.551563).ToDateTimeString() // 2024-01-24 13:14:15
```

## 波斯历/伊朗历

### 将 `公历` 转换成 `波斯历`
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

### 将 `波斯历` 转化成 `公历`
```go
carbon.CreateFromPersian(1, 1, 1).ToDateTimeString() // 2016-03-20 00:00:00
carbon.CreateFromPersian(622, 1, 1).ToDateTimeString() // 1243-03-21 00:00:00
carbon.CreateFromPersian(1395, 1, 1).ToDateTimeString() // 2016-03-20 00:00:00
carbon.CreateFromPersian(9377, 1, 1).ToDateTimeString() // 9998-03-19 00:00:00
```

### 日期判断
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



## 中国农历
> 目前仅支持公元 `1900` 年至 `2100` 年的 `200` 年时间跨度

### 将 `公历` 转换成 `农历`
```go
// 获取农历生肖
carbon.Parse("2020-08-05").Lunar().Animal() // 鼠
// 获取农历节日
carbon.Parse("2021-02-12").Lunar().Festival() // 春节

// 获取农历年份
carbon.Parse("2020-08-05").Lunar().Year() // 2020
// 获取农历月份
carbon.Parse("2020-08-05").Lunar().Month() // 6
// 获取农历闰月月份
carbon.Parse("2020-08-05").Lunar().LeapMonth() // 4
// 获取农历日期
carbon.Parse("2020-08-05").Lunar().Day() // 16
// 获取农历时辰
carbon.Parse("2020-08-05").Lunar().Hour() // 13
// 获取农历分钟
carbon.Parse("2020-08-05").Lunar().Minute() // 14
// 获取农历秒数
carbon.Parse("2020-08-05").Lunar().Second() // 15

// 获取农历日期时间字符串
carbon.Parse("2020-08-05").Lunar().String() // 2020-06-16
fmt.Printf("%s", carbon.Parse("2020-08-05").Lunar()) // 2020-06-16
// 获取农历年字符串
carbon.Parse("2020-08-05").Lunar().ToYearString() // 二零二零
// 获取农历月字符串
carbon.Parse("2020-08-05").Lunar().ToMonthString() // 六月
// 获取农历闰月字符串
carbon.Parse("2020-04-23").Lunar().ToMonthString() // 闰四月
// 获取农历周字符串
carbon.Parse("2020-04-23").Lunar().ToWeekString() // 周四
// 获取农历天字符串
carbon.Parse("2020-08-05").Lunar().ToDayString() // 十六
// 获取农历日期字符串
carbon.Parse("2020-08-05").Lunar().ToDateString() // 二零二零年六月十六

```

### 将 `农历` 转化成 `公历`
```go
// 将农历 二零二三年腊月十一 转化为 公历
carbon.CreateFromLunar(2023, 12, 11, false).ToDateTimeString() // 2024-01-21 00:00:00
// 将农历 二零二三年二月十一 转化为 公历
carbon.CreateFromLunar(2023, 2, 11, false).ToDateTimeString() // 2023-03-02 00:00:00
// 将农历 二零二三年闰二月十一 转化为 公历
carbon.CreateFromLunar(2023, 2, 11, true).ToDateTimeString() // 2023-04-01 00:00:00
```

### 日期判断
```go

// 是否是合法农历日期
carbon.Parse("0000-00-00").Lunar().IsValid() // false
carbon.Parse("2020-08-05").Lunar().IsValid() // true

// 是否是农历闰年
carbon.Parse("2020-08-05").Lunar().IsLeapYear() // true
// 是否是农历闰月
carbon.Parse("2020-08-05").Lunar().IsLeapMonth() // false

// 是否是鼠年
carbon.Parse("2020-08-05").Lunar().IsRatYear() // true
// 是否是牛年
carbon.Parse("2020-08-05").Lunar().IsOxYear() // false
// 是否是虎年
carbon.Parse("2020-08-05").Lunar().IsTigerYear() // false
// 是否是兔年
carbon.Parse("2020-08-05").Lunar().IsRabbitYear() // false
// 是否是龙年
carbon.Parse("2020-08-05").Lunar().IsDragonYear() // false
// 是否是蛇年
carbon.Parse("2020-08-05").Lunar().IsSnakeYear() // false
// 是否是马年
carbon.Parse("2020-08-05").Lunar().IsHorseYear() // false
// 是否是羊年
carbon.Parse("2020-08-05").Lunar().IsGoatYear() // false
// 是否是猴年
carbon.Parse("2020-08-05").Lunar().IsMonkeyYear() // false
// 是否是鸡年
carbon.Parse("2020-08-05").Lunar().IsRoosterYear() // false
// 是否是狗年
carbon.Parse("2020-08-05").Lunar().IsDogYear() // false
// 是否是猪年
carbon.Parse("2020-08-05").Lunar().IsPigYear() // false
```