# Carbon  #

[![Carbon Release](https://img.shields.io/github/release/dromara/carbon.svg)](https://github.com/dromara/carbon/releases)
[![Go Test](https://github.com/dromara/carbon/actions/workflows/test.yml/badge.svg)](https://github.com/dromara/carbon/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/dromara/carbon/v2)](https://goreportcard.com/report/github.com/dromara/carbon/v2)
[![Go Coverage](https://codecov.io/gh/dromara/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/dromara/carbon)
[![Carbon Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/dromara/carbon/v2)
[![License](https://img.shields.io/github/license/dromara/carbon)](https://github.com/dromara/carbon/blob/master/LICENSE)

简体中文 | [English](README.md) | [日本語](README.jp.md)

#### 项目简介

一个轻量级、语义化、对开发者友好的 `golang`
时间处理库，已被 [awesome-go](https://github.com/yinggaozhen/awesome-go-cn#日期和时间 "awesome-go-cn") 收录，并获得
`gitee` 2024 年最有价值项目（`GVP`）和 `gitcode` 2024 年度开源摘星计划 (`G-Star`) 项目

#### 仓库地址

[github.com/dromara/carbon](https://github.com/dromara/carbon "github.com/dromara/carbon")

[gitee.com/dromara/carbon](https://gitee.com/dromara/carbon "gitee.com/dromara/carbon")

[gitcode.com/dromara/carbon](https://gitcode.com/dromara/carbon "gitcode.com/dromara/carbon")

#### 安装使用

##### go version >= 1.18

```go
// 使用 github 库
go get -u github.com/dromara/carbon/v2
import "github.com/dromara/carbon/v2"

// 使用 gitee 库
go get -u gitee.com/dromara/carbon/v2
import "gitee.com/dromara/carbon/v2"

// 使用 gitcode 库
go get -u gitcode.com/dromara/carbon/v2
import "gitcode.com/dromara/carbon/v2"
```

`Carbon` 已经捐赠给了 [dromara](https://dromara.org/ "dromara") 组织，仓库地址发生了改变，如果之前用的路径是
`golang-module/carbon`，请在 `go.mod` 里将原地址更换为新路径，或执行如下命令

```go
go mod edit -replace github.com/golang-module/carbon/v2 = github.com/dromara/carbon/v2
```

#### 用法示例

> 假设当前时间为 2020-08-05 13:14:15.999999999 +0800 CST

##### 设置全局默认值

```go
carbon.SetLayout(carbon.DateTimeLayout)
carbon.SetTimezone(carbon.PRC)
carbon.SetWeekStartsAt(carbon.Sunday)
carbon.SetLocale("zh-CN")

或

carbon.SetDefault(carbon.Default{
	Layout: carbon.DateTimeLayout,
	Timezone: carbon.PRC,
	WeekStartsAt: carbon.Sunday,
	Locale: "zh-CN",
})
```

##### `Carbon` 和 `time.Time` 互转

```go
// 将标准 time.Time 转换成 Carbon
carbon.NewCarbon(time.Now())
// 将 Carbon 转换成标准 time.Time
carbon.Now().StdTime()

或

// 将标准 time.Time 转换成 Carbon
carbon.CreateFromStdTime(time.Now())
// 将 Carbon 转换成标准 time.Time
carbon.Now().StdTime()
```

##### 昨天、今天、明天

```go
// 今天此刻
fmt.Printf("%s", carbon.Now()) // 2020-08-05 13:14:15
carbon.Now().String() // 2020-08-05 13:14:15
carbon.Now().ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
// 今天日期
carbon.Now().ToDateString() // 2020-08-05
// 今天时间
carbon.Now().ToTimeString() // 13:14:15
// 指定时区的今天此刻
carbon.Now(carbon.NewYork).ToDateTimeString() // 2020-08-05 14:14:15
// 今天秒级时间戳
carbon.Now().Timestamp() // 1596604455
// 今天毫秒级时间戳
carbon.Now().TimestampMilli() // 1596604455999
// 今天微秒级时间戳
carbon.Now().TimestampMicro() // 1596604455999999
// 今天纳秒级时间戳
carbon.Now().TimestampNano() // 1596604455999999999

// 昨天此刻
fmt.Printf("%s", carbon.Yesterday()) // 2020-08-04 13:14:15
carbon.Yesterday().String() // 2020-08-04 13:14:15
carbon.Yesterday().ToString() // 2020-08-04 13:14:15.999999999 +0800 CST
carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// 昨天日期
carbon.Yesterday().ToDateString() // 2020-08-04
// 昨天时间
carbon.Yesterday().ToTimeString() // 13:14:15
// 指定日期的昨天此刻
carbon.Parse("2021-01-28 13:14:15").Yesterday().ToDateTimeString() // 2021-01-27 13:14:15
// 指定时区的昨天此刻
carbon.Yesterday(carbon.NewYork).ToDateTimeString() // 2020-08-04 14:14:15
// 昨天秒级时间戳
carbon.Yesterday().Timestamp() // 1596518055
// 昨天毫秒级时间戳
carbon.Yesterday().TimestampMilli() // 1596518055999
// 昨天微秒级时间戳
carbon.Yesterday().TimestampMicro() // 1596518055999999
// 昨天纳秒级时间戳
carbon.Yesterday().TimestampNano() // 1596518055999999999

// 明天此刻
fmt.Printf("%s", carbon.Tomorrow()) // 2020-08-06 13:14:15
carbon.Tomorrow().String() // 2020-08-06 13:14:15
carbon.Tomorrow().ToString() // 2020-08-06 13:14:15.999999999 +0800 CST
carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// 明天日期
carbon.Tomorrow().ToDateString() // 2020-08-06
// 明天时间
carbon.Tomorrow().ToTimeString() // 13:14:15
// 指定日期的明天此刻
carbon.Parse("2021-01-28 13:14:15").Tomorrow().ToDateTimeString() // 2021-01-29 13:14:15
// 指定时区的明天此刻
carbon.Tomorrow(carbon.NewYork).ToDateTimeString() // 2020-08-06 14:14:15
// 明天秒级时间戳
carbon.Tomorrow().Timestamp() // 1596690855
// 明天毫秒级时间戳
carbon.Tomorrow().TimestampMilli() //1596690855999
// 明天微秒级时间戳
carbon.Tomorrow().TimestampMicro() // 1596690855999999
// 明天纳秒级时间戳
carbon.Tomorrow().TimestampNano() // 1596690855999999999
```

##### 创建 `Carbon` 实例

```go
// 从秒级时间戳创建 Carbon 实例
carbon.CreateFromTimestamp(-1).ToString() // 1970-01-01 07:59:59 +0800 CST
carbon.CreateFromTimestamp(0).ToString() // 1970-01-01 08:00:00 +0800 CST
carbon.CreateFromTimestamp(1).ToString() // 1970-01-01 08:00:01 +0800 CST
carbon.CreateFromTimestamp(1596604455).ToString() // 2020-08-05 13:14:15 +0800 CST
// 从毫秒级时间戳创建 Carbon 实例
carbon.CreateFromTimestampMilli(1596604455999).ToString() // 2020-08-05 13:14:15.999 +0800 CST
// 从微秒级时间戳创建 Carbon 实例
carbon.CreateFromTimestampMicro(1596604455999999).ToString() // 2020-08-05 13:14:15.999999 +0800 CST
// 从纳秒级时间戳创建 Carbon 实例
carbon.CreateFromTimestampNano(1596604455999999999).ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

// 从年、月、日、时、分、秒创建 Carbon 实例
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToString() // 2020-08-05 13:14:15 +0800 CST
// 从年、月、日、时、分、秒、毫秒创建 Carbon 实例
carbon.CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString() // 2020-08-05 13:14:15.999 +0800 CST
// 从年、月、日、时、分、秒、微秒创建 Carbon 实例
carbon.CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString() // 2020-08-05 13:14:15.999999 +0800 CST
// 从年、月、日、时、分、秒、纳秒创建 Carbon 实例
carbon.CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

// 从年、月、日创建 Carbon 实例
carbon.CreateFromDate(2020, 8, 5).ToString() // 2020-08-05 00:00:00 +0800 CST
// 从年、月、日、毫秒创建 Carbon 实例
carbon.CreateFromDateMilli(2020, 8, 5, 999).ToString() // 2020-08-05 00:00:00.999 +0800 CST
// 从年、月、日、微秒创建 Carbon 实例
carbon.CreateFromDateMicro(2020, 8, 5, 999999).ToString() // 2020-08-05 00:00:00.999999 +0800 CST
// 从年、月、日、纳秒创建 Carbon 实例
carbon.CreateFromDateNano(2020, 8, 5, 999999999).ToString() // 2020-08-05 00:00:00.999999999 +0800 CST

// 从时、分、秒创建 Carbon 实例(年月日默认为当前年月日)
carbon.CreateFromTime(13, 14, 15).ToString() // 2020-08-05 13:14:15 +0800 CST
// 从时、分、秒、毫秒创建 Carbon 实例(年月日默认为当前年月日)
carbon.CreateFromTimeMilli(13, 14, 15, 999).ToString() // 2020-08-05 13:14:15.999 +0800 CST
// 从时、分、秒、微秒创建 Carbon 实例(年月日默认为当前年月日)
carbon.CreateFromTimeMicro(13, 14, 15, 999999).ToString() // 2020-08-05 13:14:15.999999 +0800 CST
// 从时、分、秒、纳秒创建 Carbon 实例(年月日默认为当前年月日)
carbon.CreateFromTimeNano(13, 14, 15, 999999999).ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
```

##### 将 `时间字符串` 解析成 `Carbon` 实例

```go
carbon.Parse("").ToDateTimeString() // 空字符串
carbon.Parse("0").ToDateTimeString() // 空字符串
carbon.Parse("xxx").ToDateTimeString() // 空字符串
carbon.Parse("00:00:00").ToDateTimeString() // 空字符串
carbon.Parse("0000-00-00").ToDateTimeString() // 空字符串
carbon.Parse("0000-00-00 00:00:00").ToDateTimeString() // 空字符串

carbon.Parse("now").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("yesterday").ToString() // 2020-08-04 13:14:15 +0800 CST
carbon.Parse("tomorrow").ToString() // 2020-08-06 13:14:15 +0800 CST

carbon.Parse("2020").ToString() // 2020-01-01 00:00:00 +0800 CST
carbon.Parse("2020-8").ToString() // 2020-08-01 00:00:00 +0800 CST
carbon.Parse("2020-08").ToString() // 2020-08-01 00:00:00 +0800 CST
carbon.Parse("2020-8-5").ToString() // 2020-08-05 00:00:00 +0800 CST
carbon.Parse("2020-8-05").ToString() // 2020-08-05 00:00:00 +0800 CST
carbon.Parse("2020-08-05").ToString() // 2020-08-05 00:00:00 +0800 CST
carbon.Parse("2020-08-05.999").ToString() // 2020-08-05 00:00:00.999 +0800 CST
carbon.Parse("2020-08-05.999999").ToString() // 2020-08-05 00:00:00.999999 +0800 CST
carbon.Parse("2020-08-05.999999999").ToString() // 2020-08-05 00:00:00.999999999 +0800 CST

carbon.Parse("2020-8-5 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-8-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-5 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-05 13:14:15.999").ToString() // 2020-08-05 13:14:15.999 +0800 CST
carbon.Parse("2020-08-05 13:14:15.999999").ToString() // 2020-08-05 13:14:15.999999 +0800 CST
carbon.Parse("2020-08-05 13:14:15.999999999").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

carbon.Parse("2020-8-5T13:14:15+08:00").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-8-05T13:14:15+08:00").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-05T13:14:15+08:00").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-05T13:14:15.999+08:00").ToString() // 2020-08-05 13:14:15.999 +0800 CST
carbon.Parse("2020-08-05T13:14:15.999999+08:00").ToString() // 2020-08-05 13:14:15.999999 +0800 CST
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

carbon.Parse("20200805").ToString() // 2020-08-05 00:00:00 +0800 CST
carbon.Parse("20200805131415").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("20200805131415.999").ToString() // 2020-08-05 13:14:15.999 +0800 CST
carbon.Parse("20200805131415.999999").ToString() // 2020-08-05 13:14:15.999999 +0800 CST
carbon.Parse("20200805131415.999999999").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
carbon.Parse("20200805131415.999+08:00").ToString() // 2020-08-05 13:14:15.999 +0800 CST
carbon.Parse("20200805131415.999999+08:00").ToString() // 2020-08-05 13:14:15.999999 +0800 CST
carbon.Parse("20200805131415.999999999+08:00").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

```

##### 通过 `布局模板` 将时间字符串解析成 `Carbon` 实例

```go
carbon.ParseByLayout("2020|08|05 13|14|15", "2006|01|02 15|04|05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("2020-08-05 13:14:15", "2006-01-02 15:04:05", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### 通过自定义 `布局模板` 将时间字符串解析成 `Carbon` 实例

```go
carbon.ParseWithLayouts("2020|08|05 13|14|15", []string{"2006|01|02 15|04|05", "2006|1|2 3|4|5"}).ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseWithLayouts("2020|08|05 13|14|15", []string{"2006|01|02 15|04|05", "2006|1|2 3|4|5"}).CurrentLayout() // 2006|01|02 15|04|05
```

##### 通过 `格式模板` 将时间字符串解析成 `Carbon` 实例

> 如果使用的字母与格式模板冲突时，请使用转义符转义该字母

```go
carbon.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("2020-08-05 13:14:15", "Y-m-d H:i:s", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### 通过自定义 `格式模板` 将时间字符串解析成 `Carbon` 实例

```go
carbon.ParseWithFormats("2020|08|05 13|14|15", []string{"Y|m|d H|i|s", "y|m|d h|i|s"}).ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseWithFormats("2020|08|05 13|14|15", []string{"Y|m|d H|i|s", "y|m|d h|i|s"}).CurrentLayout() // 2006|01|02 15|04|05
```

##### 时间冻结

```go
now := carbon.Parse("2020-08-05")
carbon.SetTestNow(now)

carbon.IsTestNow() // true
carbon.Now().ToDateString() // 2020-08-05
carbon.Yesterday().ToDateString() // 2020-08-04
carbon.Tomorrow().ToDateString() // 2020-08-05
carbon.Now().DiffForHumans() // just now
carbon.Yesterday().DiffForHumans() // 1 day ago
carbon.Tomorrow().DiffForHumans() // 1 day from now
carbon.Parse("2020-10-05").DiffForHumans() // 2 months from now
now.DiffForHumans(carbon.Parse("2020-10-05")) // 2 months before

carbon.CleanTestNow()
carbon.IsTestNow() // false
```

##### 时间边界

```go
// 本世纪开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfCentury().ToDateTimeString() // 2000-01-01 00:00:00
// 本世纪结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfCentury().ToDateTimeString() // 2999-12-31 23:59:59

// 本年代开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
carbon.Parse("2021-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
carbon.Parse("2029-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
// 本年代结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
carbon.Parse("2021-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
carbon.Parse("2029-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59

// 本年开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfYear().ToDateTimeString() // 2020-01-01 00:00:00
// 本年结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfYear().ToDateTimeString() // 2020-12-31 23:59:59

// 本季度开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfQuarter().ToDateTimeString() // 2020-07-01 00:00:00
// 本季度结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfQuarter().ToDateTimeString() // 2020-09-30 23:59:59

// 本月开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfMonth().ToDateTimeString() // 2020-08-01 00:00:00
// 本月结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfMonth().ToDateTimeString() // 2020-08-31 23:59:59

// 本周开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfWeek().ToDateTimeString() // 2020-08-02 00:00:00
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Sunday).StartOfWeek().ToDateTimeString() // 2020-08-02 00:00:00
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Monday).StartOfWeek().ToDateTimeString() // 2020-08-03 00:00:00
// 本周结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfWeek().ToDateTimeString() // 2020-08-08 23:59:59
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Sunday).EndOfWeek().ToDateTimeString() // 2020-08-08 23:59:59
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Monday).EndOfWeek().ToDateTimeString() // 2020-08-09 23:59:59

// 本日开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfDay().ToDateTimeString() // 2020-08-05 00:00:00
// 本日结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfDay().ToDateTimeString() // 2020-08-05 23:59:59

// 本小时开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfHour().ToDateTimeString() // 2020-08-05 13:00:00
// 本小时结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfHour().ToDateTimeString() // 2020-08-05 13:59:59

// 本分钟开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfMinute().ToDateTimeString() // 2020-08-05 13:14:00
// 本分钟结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfMinute().ToDateTimeString() // 2020-08-05 13:14:59

// 本秒开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfSecond().ToString() // 2020-08-05 13:14:15 +0800 CST
// 本秒结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfSecond().ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
```

##### 时间旅行

```go
// 三个世纪后
carbon.Parse("2020-02-29 13:14:15").AddCenturies(3).ToDateTimeString() // 2320-02-29 13:14:15
// 三个世纪后(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").AddCenturiesNoOverflow(3).ToDateTimeString() // 2320-02-29 13:14:15
// 一个世纪后
carbon.Parse("2020-02-29 13:14:15").AddCentury().ToDateTimeString() // 2120-02-29 13:14:15
// 一个世纪后(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").AddCenturyNoOverflow().ToDateTimeString() // 2120-02-29 13:14:15
// 三个世纪前
carbon.Parse("2020-02-29 13:14:15").SubCenturies(3).ToDateTimeString() // 1720-02-29 13:14:15
// 三个世纪前(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").SubCenturiesNoOverflow(3).ToDateTimeString() // 1720-02-29 13:14:15
// 一个世纪前
carbon.Parse("2020-02-29 13:14:15").SubCentury().ToDateTimeString() // 1920-02-29 13:14:15
// 一世纪前(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").SubCenturyNoOverflow().ToDateTimeString() // 1920-02-29 13:14:15

// 三个年代后
carbon.Parse("2020-02-29 13:14:15").AddDecades(3).ToDateTimeString() // 2050-03-01 13:14:15
// 三个年代后(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").AddDecadesNoOverflow(3).ToDateTimeString() // 2050-02-28 13:14:15
// 一个年代后
carbon.Parse("2020-02-29 13:14:15").AddDecade().ToDateTimeString() // 2030-03-01 13:14:15
// 一个年代后(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").AddDecadeNoOverflow().ToDateTimeString() // 2030-02-28 13:14:15
// 三个年代前
carbon.Parse("2020-02-29 13:14:15").SubDecades(3).ToDateTimeString() // 1990-03-01 13:14:15
// 三个年代前(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").SubDecadesNoOverflow(3).ToDateTimeString() // 1990-02-28 13:14:15
// 一个年代前
carbon.Parse("2020-02-29 13:14:15").SubDecade().ToDateTimeString() // 2010-03-01 13:14:15
// 一个年代前(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").SubDecadeNoOverflow().ToDateTimeString() // 2010-02-28 13:14:15

// 三年后
carbon.Parse("2020-02-29 13:14:15").AddYears(3).ToDateTimeString() // 2023-03-01 13:14:15
// 三年后(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").AddYearsNoOverflow(3).ToDateTimeString() // 2023-02-28 13:14:15
// 一年后
carbon.Parse("2020-02-29 13:14:15").AddYear().ToDateTimeString() // 2021-03-01 13:14:15
// 一年后(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").AddYearNoOverflow().ToDateTimeString() // 2021-02-28 13:14:15
// 三年前
carbon.Parse("2020-02-29 13:14:15").SubYears(3).ToDateTimeString() // 2017-03-01 13:14:15
// 三年前(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").SubYearsNoOverflow(3).ToDateTimeString() // 2017-02-28 13:14:15
// 一年前
carbon.Parse("2020-02-29 13:14:15").SubYear().ToDateTimeString() // 2019-03-01 13:14:15
// 一年前(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").SubYearNoOverflow().ToDateTimeString() // 2019-02-28 13:14:15

// 三个季度后
carbon.Parse("2019-05-31 13:14:15").AddQuarters(3).ToDateTimeString() // 2020-03-02 13:14:15
// 三个季度后(月份不溢出)
carbon.Parse("2019-05-31 13:14:15").AddQuartersNoOverflow(3).ToDateTimeString() // 2020-02-29 13:14:15
// 一个季度后
carbon.Parse("2019-11-30 13:14:15").AddQuarter().ToDateTimeString() // 2020-03-01 13:14:15
// 一个季度后(月份不溢出)
carbon.Parse("2019-11-30 13:14:15").AddQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
// 三个季度前
carbon.Parse("2019-08-31 13:14:15").SubQuarters(3).ToDateTimeString() // 2019-03-03 13:14:15
// 三个季度前(月份不溢出)
carbon.Parse("2019-08-31 13:14:15").SubQuartersNoOverflow(3).ToDateTimeString() // 2019-02-28 13:14:15
// 一个季度前
carbon.Parse("2020-05-31 13:14:15").SubQuarter().ToDateTimeString() // 2020-03-02 13:14:15
// 一个季度前(月份不溢出)
carbon.Parse("2020-05-31 13:14:15").SubQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// 三个月后
carbon.Parse("2020-02-29 13:14:15").AddMonths(3).ToDateTimeString() // 2020-05-29 13:14:15
// 三个月后(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").AddMonthsNoOverflow(3).ToDateTimeString() // 2020-05-29 13:14:15
// 一个月后
carbon.Parse("2020-01-31 13:14:15").AddMonth().ToDateTimeString() // 2020-03-02 13:14:15
// 一个月后(月份不溢出)
carbon.Parse("2020-01-31 13:14:15").AddMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
// 三个月前
carbon.Parse("2020-02-29 13:14:15").SubMonths(3).ToDateTimeString() // 2019-11-29 13:14:15
// 三个月前(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").SubMonthsNoOverflow(3).ToDateTimeString() // 2019-11-29 13:14:15
// 一个月前
carbon.Parse("2020-03-31 13:14:15").SubMonth().ToDateTimeString() // 2020-03-02 13:14:15
// 一个月前(月份不溢出)
carbon.Parse("2020-03-31 13:14:15").SubMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// 三周后
carbon.Parse("2020-02-29 13:14:15").AddWeeks(3).ToDateTimeString() // 2020-03-21 13:14:15
// 一周后
carbon.Parse("2020-02-29 13:14:15").AddWeek().ToDateTimeString() // 2020-03-07 13:14:15
// 三周前
carbon.Parse("2020-02-29 13:14:15").SubWeeks(3).ToDateTimeString() // 2020-02-08 13:14:15
// 一周前
carbon.Parse("2020-02-29 13:14:15").SubWeek().ToDateTimeString() // 2020-02-22 13:14:15

// 三天后
carbon.Parse("2020-08-05 13:14:15").AddDays(3).ToDateTimeString() // 2020-08-08 13:14:15
// 一天后
carbon.Parse("2020-08-05 13:14:15").AddDay().ToDateTimeString() // 2020-08-05 13:14:15
// 三天前
carbon.Parse("2020-08-05 13:14:15").SubDays(3).ToDateTimeString() // 2020-08-02 13:14:15
// 一天前
carbon.Parse("2020-08-05 13:14:15").SubDay().ToDateTimeString() // 2020-08-04 13:14:15

// 三小时后
carbon.Parse("2020-08-05 13:14:15").AddHours(3).ToDateTimeString() // 2020-08-05 16:14:15
// 二小时半后
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5h").ToDateTimeString() // 2020-08-05 15:44:15
carbon.Parse("2020-08-05 13:14:15").AddDuration("2h30m").ToDateTimeString() // 2020-08-05 15:44:15
// 一小时后
carbon.Parse("2020-08-05 13:14:15").AddHour().ToDateTimeString() // 2020-08-05 14:14:15
// 三小时前
carbon.Parse("2020-08-05 13:14:15").SubHours(3).ToDateTimeString() // 2020-08-05 10:14:15
// 二小时半前
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5h").ToDateTimeString() // 2020-08-05 10:44:15
carbon.Parse("2020-08-05 13:14:15").SubDuration("2h30m").ToDateTimeString() // 2020-08-05 10:44:15
// 一小时前
carbon.Parse("2020-08-05 13:14:15").SubHour().ToDateTimeString() // 2020-08-05 12:14:15

// 三分钟后
carbon.Parse("2020-08-05 13:14:15").AddMinutes(3).ToDateTimeString() // 2020-08-05 13:17:15
// 二分钟半后
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5m").ToDateTimeString() // 2020-08-05 13:16:45
carbon.Parse("2020-08-05 13:14:15").AddDuration("2m30s").ToDateTimeString() // 2020-08-05 13:16:45
// 一分钟后
carbon.Parse("2020-08-05 13:14:15").AddMinute().ToDateTimeString() // 2020-08-05 13:15:15
// 三分钟前
carbon.Parse("2020-08-05 13:14:15").SubMinutes(3).ToDateTimeString() // 2020-08-05 13:11:15
// 二分钟半前
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5m").ToDateTimeString() // 2020-08-05 13:11:45
carbon.Parse("2020-08-05 13:14:15").SubDuration("2m30s").ToDateTimeString() // 2020-08-05 13:11:45
// 一分钟前
carbon.Parse("2020-08-05 13:14:15").SubMinute().ToDateTimeString() // 2020-08-05 13:13:15

// 三秒钟后
carbon.Parse("2020-08-05 13:14:15").AddSeconds(3).ToDateTimeString() // 2020-08-05 13:14:18
// 二秒钟半后
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:17
// 一秒钟后
carbon.Parse("2020-08-05 13:14:15").AddSecond().ToDateTimeString() // 2020-08-05 13:14:16
// 三秒钟前
carbon.Parse("2020-08-05 13:14:15").SubSeconds(3).ToDateTimeString() // 2020-08-05 13:14:12
// 二秒钟半前
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:12
// 一秒钟前
carbon.Parse("2020-08-05 13:14:15").SubSecond().ToDateTimeString() // 2020-08-05 13:14:14

// 三毫秒后
carbon.Parse("2020-08-05 13:14:15.222222222").AddMilliseconds(3).ToString() // 2020-08-05 13:14:15.225222222 +0800 CST
// 一毫秒后
carbon.Parse("2020-08-05 13:14:15.222222222").AddMillisecond().ToString() // 2020-08-05 13:14:15.223222222 +0800 CST
// 三毫秒前
carbon.Parse("2020-08-05 13:14:15.222222222").SubMilliseconds(3).ToString() // 2020-08-05 13:14:15.219222222 +0800 CST
// 一毫秒前
carbon.Parse("2020-08-05 13:14:15.222222222").SubMillisecond().ToString() // 2020-08-05 13:14:15.221222222 +0800 CST

// 三微秒后
carbon.Parse("2020-08-05 13:14:15.222222222").AddMicroseconds(3).ToString() // 2020-08-05 13:14:15.222225222 +0800 CST
// 一微秒后
carbon.Parse("2020-08-05 13:14:15.222222222").AddMicrosecond().ToString() // 2020-08-05 13:14:15.222223222 +0800 CST
// 三微秒前
carbon.Parse("2020-08-05 13:14:15.222222222").SubMicroseconds(3).ToString() // 2020-08-05 13:14:15.222219222 +0800 CST
// 一微秒前
carbon.Parse("2020-08-05 13:14:15.222222222").SubMicrosecond().ToString() // 2020-08-05 13:14:15.222221222 +0800 CST

// 三纳秒后
carbon.Parse("2020-08-05 13:14:15.222222222").AddNanoseconds(3).ToString() // 2020-08-05 13:14:15.222222225 +0800 CST
// 一纳秒后
carbon.Parse("2020-08-05 13:14:15.222222222").AddNanosecond().ToString() // 2020-08-05 13:14:15.222222223 +0800 CST
// 三纳秒前
carbon.Parse("2020-08-05 13:14:15.222222222").SubNanoseconds(3).ToString() // 2020-08-05 13:14:15.222222219 +0800 CST
// 一纳秒前
carbon.Parse("2020-08-05 13:14:15.222222222").SubNanosecond().ToString() // 2020-08-05 13:14:15.222222221 +0800 CST
```

##### 时间差值

```go
// 相差多少年
carbon.Parse("2021-08-05 13:14:15").DiffInYears(carbon.Parse("2020-08-05 13:14:15")) // -1
// 相差多少年（绝对值）
carbon.Parse("2021-08-05 13:14:15").DiffAbsInYears(carbon.Parse("2020-08-05 13:14:15")) // 1

// 相差多少月
carbon.Parse("2020-08-05 13:14:15").DiffInMonths(carbon.Parse("2020-07-05 13:14:15")) // -1
// 相差多少月（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInMonths(carbon.Parse("2020-07-05 13:14:15")) // 1

// 相差多少周
carbon.Parse("2020-08-05 13:14:15").DiffInWeeks(carbon.Parse("2020-07-28 13:14:15")) // -1
// 相差多少周（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInWeeks(carbon.Parse("2020-07-28 13:14:15")) // 1

// 相差多少天
carbon.Parse("2020-08-05 13:14:15").DiffInDays(carbon.Parse("2020-08-04 13:14:15")) // -1
// 相差多少天（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInDays(carbon.Parse("2020-08-04 13:14:15")) // 1

// 相差多少小时
carbon.Parse("2020-08-05 13:14:15").DiffInHours(carbon.Parse("2020-08-05 12:14:15")) // -1
// 相差多少小时（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInHours(carbon.Parse("2020-08-05 12:14:15")) // 1

// 相差多少分
carbon.Parse("2020-08-05 13:14:15").DiffInMinutes(carbon.Parse("2020-08-05 13:13:15")) // -1
// 相差多少分（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInMinutes(carbon.Parse("2020-08-05 13:13:15")) // 1

// 相差多少秒
carbon.Parse("2020-08-05 13:14:15").DiffInSeconds(carbon.Parse("2020-08-05 13:14:14")) // -1
// 相差多少秒（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInSeconds(carbon.Parse("2020-08-05 13:14:14")) // 1

// 相差字符串
carbon.Now().DiffInString() // just now
carbon.Now().AddYearsNoOverflow(1).DiffInString() // -1 year
carbon.Now().SubYearsNoOverflow(1).DiffInString() // 1 year
// 相差字符串（绝对值）
carbon.Now().DiffAbsInString(carbon.Now()) // just now
carbon.Now().AddYearsNoOverflow(1).DiffAbsInString(carbon.Now()) // 1 year
carbon.Now().SubYearsNoOverflow(1).DiffAbsInString(carbon.Now()) // 1 year

// 相差时长
now := carbon.Now()
now.DiffInDuration(now).String() // 0s
now.AddHour().DiffInDuration(now).String() // 1h0m0s
now.SubHour().DiffInDuration(now).String() // -1h0m0s
// 相差时长（绝对值）
now.DiffAbsInDuration(now).String() // 0s
now.AddHour().DiffAbsInDuration(now).String() // 1h0m0s
now.SubHour().DiffAbsInDuration(now).String() // 1h0m0s

// 对人类友好的可读格式时间差
carbon.Parse("2020-08-05 13:14:15").DiffForHumans() // just now
carbon.Parse("2019-08-05 13:14:15").DiffForHumans() // 1 year ago
carbon.Parse("2018-08-05 13:14:15").DiffForHumans() // 2 years ago
carbon.Parse("2021-08-05 13:14:15").DiffForHumans() // 1 year from now
carbon.Parse("2022-08-05 13:14:15").DiffForHumans() // 2 years from now

carbon.Parse("2020-08-05 13:14:15").DiffForHumans(carbon.Now()) // 1 year before
carbon.Parse("2019-08-05 13:14:15").DiffForHumans(carbon.Now()) // 2 years before
carbon.Parse("2018-08-05 13:14:15").DiffForHumans(carbon.Now()) // 1 year after
carbon.Parse("2022-08-05 13:14:15").DiffForHumans(carbon.Now()) // 2 years after
```

##### 时间极值

```go
c0 := carbon.Parse("2023-04-01")
c1 := carbon.Parse("2023-03-28")
c2 := carbon.Parse("2023-04-16")
// 返回最近的 Carbon 实例
c0.Closest(c1, c2) // c1
// 返回最远的 Carbon 实例
c0.Farthest(c1, c2) // c2

yesterday := carbon.Yesterday()
today := carbon.Now()
tomorrow := carbon.Tomorrow()
// 返回最大的 Carbon 实例
carbon.Max(yesterday, today, tomorrow) // tomorrow
// 返回最小的 Carbon 实例
carbon.Min(yesterday, today, tomorrow) // yesterday

// 返回 Carbon 的最大值
carbon.MaxValue().ToString() // 9999-12-31 23:59:59.999999999 +0000 UTC
// 返回 Carbon 的最小值
carbon.MinValue().ToString() // -9998-01-01 00:00:00 +0000 UTC

// 返回 Duration 的最大值
carbon.MaxDuration().Seconds() // 9.223372036854776e+09
// 返回 Duration 的最小值
carbon.MinDuration().Seconds() // -9.223372036854776e+09
```

##### 时间判断

```go
// 是否有错误
carbon.Parse("0001-01-01 00:00:00 +0000 UTC").HasError() // false
carbon.NewCarbon().HasError() // false
carbon.Parse("").HasError() // false
carbon.Parse("0").HasError() // true
carbon.Parse("xxx").HasError() // true
carbon.Parse("2020-08-05").IsNil() // false

// 是否是 nil
carbon.Parse("0001-01-01 00:00:00 +0000 UTC").IsNil() // false
carbon.NewCarbon().IsNil() // false
carbon.Parse("").IsNil() // true
carbon.Parse("0").IsNil() // false
carbon.Parse("xxx").IsNil() // false
carbon.Parse("2020-08-05").IsNil() // false

// 是否是零值时间(0001-01-01 00:00:00 +0000 UTC)
carbon.Parse("0001-01-01 00:00:00 +0000 UTC").IsZero() // true
carbon.NewCarbon().IsZero() // true
carbon.Parse("").IsZero() // false
carbon.Parse("0").IsZero() // false
carbon.Parse("xxx").IsZero() // false
carbon.Parse("0000-00-00 00:00:00").IsZero() // false
carbon.Parse("0000-00-00").IsZero() // false
carbon.Parse("00:00:00").IsZero() // false
carbon.Parse("2020-08-05 00:00:00").IsZero() // false
carbon.Parse("2020-08-05").IsZero() // false
carbon.Parse("2020-08-05").SetTimezone("xxx").IsZero() // false

// 是否是有效时间
carbon.Parse("0001-01-01 00:00:00 +0000 UTC").IsValid() // true
carbon.NewCarbon().IsValid() // true
carbon.Parse("").IsValid() // false
carbon.Parse("0").IsValid() // false
carbon.Parse("xxx").IsValid() // false
carbon.Parse("0000-00-00 00:00:00").IsValid() // false
carbon.Parse("0000-00-00").IsValid() // false
carbon.Parse("00:00:00").IsValid() // false
carbon.Parse("2020-08-05 00:00:00").IsValid() // true
carbon.Parse("2020-08-05").IsValid() // true
carbon.Parse("2020-08-05").SetTimezone("xxx").IsValid() // false

// 是否是无效时间
carbon.Parse("0001-01-01 00:00:00 +0000 UTC").IsValid() // false
carbon.NewCarbon().IsValid() // false
carbon.Parse("").IsInvalid() // true
carbon.Parse("0").IsInvalid() // true
carbon.Parse("xxx").IsInvalid() // true
carbon.Parse("0000-00-00 00:00:00").IsInvalid() // true
carbon.Parse("0000-00-00").IsInvalid() // true
carbon.Parse("00:00:00").IsInvalid() // true
carbon.Parse("2020-08-05 00:00:00").IsInvalid() // false
carbon.Parse("2020-08-05").IsInvalid() // false
carbon.Parse("2020-08-05").SetTimezone("xxx").IsInvalid() // true

// 是否是夏令时
carbon.Parse("").IsDST() // false
carbon.Parse("0").IsDST() // false
carbon.Parse("xxx").IsDST() // false
carbon.Parse("0000-00-00 00:00:00").IsDST() // false
carbon.Parse("0000-00-00").IsDST() // false
carbon.Parse("00:00:00").IsDST() // false
carbon.Parse("2023-01-01", "Australia/Brisbane").IsDST() // false
carbon.Parse("2023-01-01", "Australia/Sydney").IsDST() // true

// 是否是上午
carbon.Parse("2020-08-05 00:00:00").IsAM() // true
carbon.Parse("2020-08-05 08:00:00").IsAM() // true
carbon.Parse("2020-08-05 12:00:00").IsAM() // false
carbon.Parse("2020-08-05 13:00:00").IsAM() // false
// 是否是下午
carbon.Parse("2020-08-05 00:00:00").IsPM() // false
carbon.Parse("2020-08-05 08:00:00").IsPM() // false
carbon.Parse("2020-08-05 12:00:00").IsPM() // true
carbon.Parse("2020-08-05 13:00:00").IsPM() // true

// 是否是当前时间
carbon.Now().IsNow() // true
// 是否是未来时间
carbon.Tomorrow().IsFuture() // true
// 是否是过去时间
carbon.Yesterday().IsPast() // true

// 是否是闰年
carbon.Parse("2020-08-05 13:14:15").IsLeapYear() // true
// 是否是长年
carbon.Parse("2020-08-05 13:14:15").IsLongYear() // true

// 是否是一月
carbon.Parse("2020-08-05 13:14:15").IsJanuary() // false
// 是否是二月
carbon.Parse("2020-08-05 13:14:15").IsFebruary() // false
// 是否是三月
carbon.Parse("2020-08-05 13:14:15").IsMarch() // false
// 是否是四月
carbon.Parse("2020-08-05 13:14:15").IsApril() // false
// 是否是五月
carbon.Parse("2020-08-05 13:14:15").IsMay() // false
// 是否是六月
carbon.Parse("2020-08-05 13:14:15").IsJune() // false
// 是否是七月
carbon.Parse("2020-08-05 13:14:15").IsJuly() // false
// 是否是八月
carbon.Parse("2020-08-05 13:14:15").IsAugust() // false
// 是否是九月
carbon.Parse("2020-08-05 13:14:15").IsSeptember() // true
// 是否是十月
carbon.Parse("2020-08-05 13:14:15").IsOctober() // false
// 是否是十一月
carbon.Parse("2020-08-05 13:14:15").IsNovember() // false
// 是否是十二月
carbon.Parse("2020-08-05 13:14:15").IsDecember() // false

// 是否是周一
carbon.Parse("2020-08-05 13:14:15").IsMonday() // false
// 是否是周二
carbon.Parse("2020-08-05 13:14:15").IsTuesday() // true
// 是否是周三
carbon.Parse("2020-08-05 13:14:15").IsWednesday() // false
// 是否是周四
carbon.Parse("2020-08-05 13:14:15").IsThursday() // false
// 是否是周五
carbon.Parse("2020-08-05 13:14:15").IsFriday() // false
// 是否是周六
carbon.Parse("2020-08-05 13:14:15").IsSaturday() // false
// 是否是周日
carbon.Parse("2020-08-05 13:14:15").IsSunday() // false

// 是否是工作日
carbon.Parse("2020-08-05 13:14:15").IsWeekday() // false
// 是否是周末
carbon.Parse("2020-08-05 13:14:15").IsWeekend() // true

// 是否是昨天
carbon.Parse("2020-08-04 13:14:15").IsYesterday() // true
carbon.Parse("2020-08-04 00:00:00").IsYesterday() // true
carbon.Parse("2020-08-04").IsYesterday() // true
// 是否是今天
carbon.Parse("2020-08-05 13:14:15").IsToday() // true
carbon.Parse("2020-08-05 00:00:00").IsToday() // true
carbon.Parse("2020-08-05").IsToday() // true
// 是否是明天
carbon.Parse("2020-08-06 13:14:15").IsTomorrow() // true
carbon.Parse("2020-08-06 00:00:00").IsTomorrow() // true
carbon.Parse("2020-08-06").IsTomorrow() // true

// 是否是同一世纪
carbon.Parse("2020-08-05 13:14:15").IsSameCentury(carbon.Parse("3020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameCentury(carbon.Parse("2099-08-05 13:14:15")) // true
// 是否是同一年代
carbon.Parse("2020-08-05 13:14:15").IsSameDecade(carbon.Parse("2030-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameDecade(carbon.Parse("2120-08-05 13:14:15")) // true
// 是否是同一年
carbon.Parse("2020-08-05 00:00:00").IsSameYear(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameYear(carbon.Parse("2020-12-31 13:14:15")) // true
// 是否是同一季节
carbon.Parse("2020-08-05 00:00:00").IsSameQuarter(carbon.Parse("2020-09-05 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameQuarter(carbon.Parse("2021-01-31 13:14:15")) // true
// 是否是同一月
carbon.Parse("2020-01-01 00:00:00").IsSameMonth(carbon.Parse("2021-01-31 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameMonth(carbon.Parse("2020-01-31 13:14:15")) // true
// 是否是同一天
carbon.Parse("2020-08-05 13:14:15").IsSameDay(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 00:00:00").IsSameDay(carbon.Parse("2020-08-05 13:14:15")) // true
// 是否是同一小时
carbon.Parse("2020-08-05 13:14:15").IsSameHour(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:00:00").IsSameHour(carbon.Parse("2020-08-05 13:14:15")) // true
// 是否是同一分钟
carbon.Parse("2020-08-05 13:14:15").IsSameMinute(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:00").IsSameMinute(carbon.Parse("2020-08-05 13:14:15")) // true
// 是否是同一秒
carbon.Parse("2020-08-05 13:14:15").IsSameSecond(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameSecond(carbon.Parse("2020-08-05 13:14:15")) // true

// 是否大于
carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-05 13:14:15")) // false

// 是否小于
carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-05 13:14:15")) // false

// 是否等于
carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:00")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:00")) // false

// 是否不等于
carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("!=", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<>", carbon.Parse("2020-08-05 13:14:15")) // false

// 是否大于等于
carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-05 13:14:15")) // true

// 是否小于等于
carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-05 13:14:15")) // true

// 是否在两个时间之间(不包括这两个时间)
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 是否在两个时间之间(包括开始时间)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 是否在两个时间之间(包括结束时间)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 是否在两个时间之间(包括这两个时间)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
```

> 关于长年(LongYear)的定义, 请查看 https://en.wikipedia.org/wiki/ISO_8601#Week_dates

##### 时间设置

```go
// 设置时区
carbon.Parse("2020-08-05 13:14:15").SetTimezone(carbon.UTC).ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-08-05 13:14:15").SetTimezone(carbon.PRC).ToString() // 2020-08-05 21:14:15 +0800 CST
carbon.Parse("2020-08-05 13:14:15").SetTimezone(carbon.Tokyo).ToString() // 2020-08-05 22:14:15 +0900 JST

// 设置位置
utc, _ := time.LoadLocation(carbon.UTC)
carbon.Parse("2020-08-05 13:14:15").SetLocation(utc).ToString() // 2020-08-05 13:14:15 +0000 UTC
prc, _ := time.LoadLocation(carbon.PRC)
carbon.Parse("2020-08-05 13:14:15").SetLocation(prc).ToString() // 2020-08-05 21:14:15 +0800 CST
tokyo, _ := time.LoadLocation(carbon.Tokyo)
carbon.Parse("2020-08-05 13:14:15").SetLocation(tokyo).ToString() // 2020-08-05 22:14:15 +0900 JST

// 设置区域
carbon.Parse("2020-07-05 13:14:15").SetLocale("en").DiffForHumans() // 1 month ago
carbon.Parse("2020-07-05 13:14:15").SetLocale("zh-CN").DiffForHumans() // 1 月前

// 设置年、月、日、时、分、秒
carbon.Parse("2020-01-01").SetDateTime(2019, 2, 2, 13, 14, 15).ToString() // 2019-02-02 13:14:15 +0800 CST
carbon.Parse("2020-01-01").SetDateTime(2019, 2, 31, 13, 14, 15).ToString() // 2019-03-03 13:14:15 +0800 CST
// 设置年、月、日、时、分、秒、毫秒
carbon.Parse("2020-01-01").SetDateTimeMilli(2019, 2, 2, 13, 14, 15, 999).ToString() // 2019-02-02 13:14:15.999 +0800 CST
carbon.Parse("2020-01-01").SetDateTimeMilli(2019, 2, 31, 13, 14, 15, 999).ToString() // 2019-03-03 13:14:15.999 +0800 CST
// 设置年、月、日、时、分、秒、微秒
carbon.Parse("2020-01-01").SetDateTimeMicro(2019, 2, 2, 13, 14, 15, 999999).ToString() // 2019-02-02 13:14:15.999999 +0800 CST
carbon.Parse("2020-01-01").SetDateTimeMicro(2019, 2, 31, 13, 14, 15, 999999).ToString() // 2019-03-03 13:14:15.999999 +0800 CST
// 设置年、月、日、时、分、秒、纳秒
carbon.Parse("2020-01-01").SetDateTimeNano(2019, 2, 2, 13, 14, 15, 999999999).ToString() // 2019-02-02 13:14:15.999999999 +0800 CST
carbon.Parse("2020-01-01").SetDateTimeNano(2019, 2, 31, 13, 14, 15, 999999999).ToString() // 2019-03-03 13:14:15.999999999 +0800 CST

// 设置年、月、日
carbon.Parse("2020-01-01").SetDate(2019, 2, 2).ToString() // 2019-02-02 00:00:00 +0800 CST
carbon.Parse("2020-01-01").SetDate(2019, 2, 31).ToString() // 2019-03-03 00:00:00 +0800 CST
// 设置年、月、日、毫秒
carbon.Parse("2020-01-01").SetDateMilli(2019, 2, 2, 999).ToString() // 2019-02-02 00:00:00.999 +0800 CST
carbon.Parse("2020-01-01").SetDateMilli(2019, 2, 31, 999).ToString() // 2019-03-03 00:00:00.999 +0800 CST
// 设置年、月、日、微秒
carbon.Parse("2020-01-01").SetDateMicro(2019, 2, 2, 999999).ToString() // 2019-02-02 00:00:00.999999 +0800 CST
carbon.Parse("2020-01-01").SetDateMicro(2019, 2, 31, 999999).ToString() // 2019-03-03 00:00:00.999999 +0800 CST
// 设置年、月、日、纳秒
carbon.Parse("2020-01-01").SetDateNano(2019, 2, 2, 999999999).ToString() // 2019-02-02 00:00:00.999999999 +0800 CST
carbon.Parse("2020-01-01").SetDateNano(2019, 2, 31, 999999999).ToString() // 2019-03-03 00:00:00.999999999 +0800 CST

// 设置时、分、秒
carbon.Parse("2020-01-01").SetTime(13, 14, 15).ToString() // 2020-01-01 13:14:15 +0800 CST
carbon.Parse("2020-01-01").SetTime(13, 14, 90).ToString() // 2020-01-01 13:15:30 +0800 CST
// 设置时、分、秒、毫秒
carbon.Parse("2020-01-01").SetTimeMilli(13, 14, 15, 999).ToString() // 2020-01-01 13:14:15.999 +0800 CST
carbon.Parse("2020-01-01").SetTimeMilli(13, 14, 90, 999).ToString() // 2020-01-01 13:15:30.999 +0800 CST
// 设置时、分、秒、微秒
carbon.Parse("2020-01-01").SetTimeMicro(13, 14, 15, 999999).ToString() // 2020-01-01 13:14:15.999999 +0800 CST
carbon.Parse("2020-01-01").SetTimeMicro(13, 14, 90, 999999).ToString() // 2020-01-01 13:15:30.999999 +0800 CST
// 设置时、分、秒、纳秒
carbon.Parse("2020-01-01").SetTimeNano(13, 14, 15, 999999999).ToString() // 2020-01-01 13:14:15.999999999 +0800 CST
carbon.Parse("2020-01-01").SetTimeNano(13, 14, 90, 999999999).ToString() // 2020-01-01 13:15:30.999999999 +0800 CST

// 设置年份
carbon.Parse("2020-02-29").SetYear(2021).ToDateString() // 2021-03-01
// 设置年份(月份不溢出)
carbon.Parse("2020-02-29").SetYearNoOverflow(2021).ToDateString() // 2021-02-28

// 设置月份
carbon.Parse("2020-01-31").SetMonth(2).ToDateString() // 2020-03-02
// 设置月份(月份不溢出)
carbon.Parse("2020-01-31").SetMonthNoOverflow(2).ToDateString() // 2020-02-29

// 设置一周的开始日期
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Monday).Week() // 6
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Sunday).Week() // 0

// 设置日期
carbon.Parse("2019-08-05").SetDay(31).ToDateString() // 2020-08-31
carbon.Parse("2020-02-01").SetDay(31).ToDateString() // 2020-03-02

// 设置小时
carbon.Parse("2020-08-05 13:14:15").SetHour(10).ToDateTimeString() // 2020-08-05 10:14:15
carbon.Parse("2020-08-05 13:14:15").SetHour(24).ToDateTimeString() // 2020-08-06 00:14:15

// 设置分钟
carbon.Parse("2020-08-05 13:14:15").SetMinute(10).ToDateTimeString() // 2020-08-05 13:10:15
carbon.Parse("2020-08-05 13:14:15").SetMinute(60).ToDateTimeString() // 2020-08-05 14:00:15

// 设置秒
carbon.Parse("2020-08-05 13:14:15").SetSecond(10).ToDateTimeString() // 2020-08-05 13:14:10
carbon.Parse("2020-08-05 13:14:15").SetSecond(60).ToDateTimeString() // 2020-08-05 13:15:00

// 设置毫秒
carbon.Parse("2020-08-05 13:14:15").SetMillisecond(100).Millisecond() // 100
carbon.Parse("2020-08-05 13:14:15").SetMillisecond(999).Millisecond() // 999

// 设置微妙
carbon.Parse("2020-08-05 13:14:15").SetMicrosecond(100000).Microsecond() // 100000
carbon.Parse("2020-08-05 13:14:15").SetMicrosecond(999999).Microsecond() // 999999

// 设置纳秒
carbon.Parse("2020-08-05 13:14:15").SetNanosecond(100000000).Nanosecond() // 100000000
carbon.Parse("2020-08-05 13:14:15").SetNanosecond(999999999).Nanosecond() // 999999999
```

##### 时间获取

```go
// 获取本年总天数
carbon.Parse("2019-08-05 13:14:15").DaysInYear() // 365
carbon.Parse("2020-08-05 13:14:15").DaysInYear() // 366
// 获取本月总天数
carbon.Parse("2020-02-01 13:14:15").DaysInMonth() // 29
carbon.Parse("2020-04-01 13:14:15").DaysInMonth() // 30
carbon.Parse("2020-08-01 13:14:15").DaysInMonth() // 31

// 获取本年第几天
carbon.Parse("2020-08-05 13:14:15").DayOfYear() // 218
// 获取本年第几周
carbon.Parse("2019-12-31 13:14:15").WeekOfYear() // 1
carbon.Parse("2020-08-05 13:14:15").WeekOfYear() // 32
// 获取本月第几天
carbon.Parse("2020-08-05 13:14:15").DayOfMonth() // 5
// 获取本月第几周
carbon.Parse("2020-08-05 13:14:15").WeekOfMonth() // 1
// 获取本周第几天
carbon.Parse("2020-08-05 13:14:15").DayOfWeek() // 3

// 获取当前年月日时分秒
carbon.Parse("2020-08-05 13:14:15").DateTime() // 2020,8,5,13,14,15
// 获取当前年月日时分秒毫秒
carbon.Parse("2020-08-05 13:14:15").DateTimeMilli() // 2020,8,5,13,14,15,999
// 获取当前年月日时分秒微秒
carbon.Parse("2020-08-05 13:14:15").DateTimeMicro() // 2020,8,5,13,14,15,999999
// 获取当前年月日时分秒纳秒
carbon.Parse("2020-08-05 13:14:15").DateTimeNano() // 2020,8,5,13,14,15,999999999

// 获取当前年月日
carbon.Parse("2020-08-05 13:14:15.999999999").Date() // 2020,8,5
// 获取当前年月日毫秒
carbon.Parse("2020-08-05 13:14:15.999999999").DateMilli() // 2020,8,5,999
// 获取当前年月日微秒
carbon.Parse("2020-08-05 13:14:15.999999999").DateMicro() // 2020,8,5,999999
// 获取当前年月日纳秒
carbon.Parse("2020-08-05 13:14:15.999999999").DateNano() // 2020,8,5,999999999

// 获取当前时分秒
carbon.Parse("2020-08-05 13:14:15.999999999").Time() // 13,14,15
// 获取当前时分秒毫秒
carbon.Parse("2020-08-05 13:14:15.999999999").TimeMilli() // 13,14,15,999
// 获取当前时分秒微秒
carbon.Parse("2020-08-05 13:14:15.999999999").TimeMicro() // 13,14,15,999999
// 获取当前时分秒纳秒
carbon.Parse("2020-08-05 13:14:15.999999999").TimeNano() // 13,14,15,999999999

// 获取当前世纪
carbon.Parse("2020-08-05 13:14:15").Century() // 21
// 获取当前年代
carbon.Parse("2019-08-05 13:14:15").Decade() // 10
carbon.Parse("2021-08-05 13:14:15").Decade() // 20
// 获取当前年份
carbon.Parse("2020-08-05 13:14:15").Year() // 2020
// 获取当前季度
carbon.Parse("2020-08-05 13:14:15").Quarter() // 3
// 获取当前月份
carbon.Parse("2020-08-05 13:14:15").Month() // 8
// 获取当前周(从0开始)
carbon.Parse("2020-08-02 13:14:15").Week() // 0
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Sunday).Week() // 0
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Monday).Week() // 6
// 获取当前天数
carbon.Parse("2020-08-05 13:14:15").Day() // 5
// 获取当前小时
carbon.Parse("2020-08-05 13:14:15").Hour() // 13
// 获取当前分钟
carbon.Parse("2020-08-05 13:14:15").Minute() // 14
// 获取当前秒钟
carbon.Parse("2020-08-05 13:14:15").Second() // 15
// 获取当前毫秒
carbon.Parse("2020-08-05 13:14:15.999").Millisecond() // 999
// 获取当前微秒
carbon.Parse("2020-08-05 13:14:15.999").Microsecond() // 999000
// 获取当前纳秒
carbon.Parse("2020-08-05 13:14:15.999").Nanosecond() // 999000000

// 获取秒级时间戳
carbon.Parse("2020-08-05 13:14:15").Timestamp() // 1596604455
// 获取毫秒级时间戳
carbon.Parse("2020-08-05 13:14:15").TimestampMilli() // 1596604455000
// 获取微秒级时间戳
carbon.Parse("2020-08-05 13:14:15").TimestampMicro() // 1596604455000000
// 获取纳秒级时间戳
carbon.Parse("2020-08-05 13:14:15").TimestampNano() // 1596604455000000000

// 获取时区位置
carbon.SetTimezone(carbon.PRC).Timezone() // PRC
carbon.SetTimezone(carbon.Tokyo).Timezone() // Asia/Tokyo

// 获取时区名称
carbon.SetTimezone(carbon.PRC).ZoneName() // CST
carbon.SetTimezone(carbon.Tokyo).ZoneName() // JST

// 获取时区偏移量，单位秒
carbon.SetTimezone(carbon.PRC).ZoneOffset() // 28800
carbon.SetTimezone(carbon.Tokyo).ZoneOffset() // 32400

// 获取当前区域
carbon.Now().Locale() // zh-CN
carbon.Now().SetLocale("en").Locale() // en

// 获取当前星座
carbon.Now().Constellation() // 狮子座
carbon.Now().SetLocale("en").Constellation() // Leo
carbon.Now().SetLocale("zh-CN").Constellation() // 狮子座

// 获取当前季节
carbon.Now().Season() // 夏季
carbon.Now().SetLocale("en").Season() // Summer
carbon.Now().SetLocale("zh-CN").Season() // 夏季

// 获取一周的开始日期
carbon.SetWeekStartsAt(carbon.Sunday).WeekStartsAt() // Sunday
carbon.SetWeekStartsAt(carbon.Monday).WeekStartsAt() // Monday

// 获取当前布局模板
carbon.Parse("now").CurrentLayout() // "2006-01-02 15:04:05"
carbon.ParseByLayout("2020-08-05", DateLayout).CurrentLayout() // "2006-01-02"

// 获取年龄
carbon.Parse("2002-01-01 13:14:15").Age() // 17
carbon.Parse("2002-12-31 13:14:15").Age() // 18
```

##### 时间输出

```go
// 输出日期时间字符串
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeString() // 2020-08-05 13:14:15
// 输出日期时间字符串，包含毫秒
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeMilliString() // 2020-08-05 13:14:15.999
// 输出日期时间字符串，包含微秒
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeMicroString() // 2020-08-05 13:14:15.999999
// 输出日期时间字符串，包含纳秒
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeNanoString() // 2020-08-05 13:14:15.999999999

// 输出简写日期时间字符串
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeString() // 20200805131415
// 输出简写日期时间字符串，包含毫秒
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeMilliString() // 20200805131415.999
// 输出简写日期时间字符串，包含微秒
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeMicroString() // 20200805131415.999999
// 输出简写日期时间字符串，包含纳秒
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeNanoString() // 20200805131415.999999999

// 输出日期字符串
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateString() // 2020-08-05
// 输出日期字符串，包含毫秒
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateMilliString() // 2020-08-05.999
// 输出日期字符串，包含微秒
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateMicroString() // 2020-08-05.999999
// 输出日期字符串，包含纳秒
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateNanoString() // 2020-08-05.999999999

// 输出简写日期字符串
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateString() // 20200805
// 输出简写日期字符串，包含毫秒
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateMilliString() // 20200805.999
// 输出简写日期字符串，包含微秒
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateMicroString() // 20200805.999999
// 输出简写日期字符串，包含纳秒
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateNanoString() // 20200805.999999999

// 输出时间字符串
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeString() // 13:14:15
// 输出时间字符串，包含毫秒
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeMilliString() // 13:14:15.999
// 输出时间字符串，包含微秒
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeMicroString() // 13:14:15.999999
// 输出时间字符串，包含纳秒
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeNanoString() // 13:14:15.999999999

// 输出简写时间字符串
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeString() // 131415
// 输出简写时间字符串，包含毫秒
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeMilliString() // 131415.999
// 输出简写时间字符串，包含微秒
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeMicroString() // 131415.999999
// 输出简写时间字符串，包含纳秒
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeNanoString() // 131415.999999999

// 输出 Ansic 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToAnsicString() // Wed Aug  5 13:14:15 2020
// 输出 Atom 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToAtomString() // 2020-08-05T13:14:15+08:00
// 输出 UnixDate 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToUnixDateString() // Wed Aug  5 13:14:15 CST 2020
// 输出 RubyDate 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRubyDateString() // Wed Aug 05 13:14:15 +0800 2020
// 输出 Kitchen 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToKitchenString() // 1:14PM
// 输出 Cookie 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToCookieString() // Wednesday, 05-Aug-2020 13:14:15 CST
// 输出 DayDateTime 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString() // Wed, Aug 5, 2020 1:14 PM
// 输出 RSS 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRssString() // Wed, 05 Aug 2020 13:14:15 +0800
// 输出 W3C 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToW3cString() // 2020-08-05T13:14:15+08:00

// 输出 ISO8601 格式字符串
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601String() // 2020-08-05T13:14:15+08:00
// 输出 ISO8601Milli 格式字符串
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601MilliString() // 2020-08-05T13:14:15.999+08:00
// 输出 ISO8601Micro 格式字符串
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601MicroString() // 2020-08-05T13:14:15.999999+08:00
// 输出 ISO8601Nano 格式字符串
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601NanoString() // 2020-08-05T13:14:15.999999999+08:00
// 输出 ISO8601Zulu 格式字符串
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluString() // 2020-08-05T13:14:15Z
// 输出 ISO8601ZuluMilli 格式字符串
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluMilliString() // 2020-08-05T13:14:15.999Z
// 输出 ISO8601ZuluMicro 格式字符串
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluMicroString() // 2020-08-05T13:14:15.999999Z
// 输出 ISO8601ZuluNano 格式字符串
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluNanoString() // 2020-08-05T13:14:15.999999999Z

// 输出 RFC822 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc822String() // 05 Aug 20 13:14 CST
// 输出 RFC822Z 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc822zString() // 05 Aug 20 13:14 +0800
// 输出 RFC850 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc850String() // Wednesday, 05-Aug-20 13:14:15 CST
// 输出 RFC1036 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc1036String() // Wed, 05 Aug 20 13:14:15 +0800
// 输出 RFC1123 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc1123String() // Wed, 05 Aug 2020 13:14:15 CST
// 输出 RFC1123Z 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc1123zString() // Wed, 05 Aug 2020 13:14:15 +0800
// 输出 RFC2822 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc2822String() // Wed, 05 Aug 2020 13:14:15 +0800
// 输出 RFC7231 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc7231String() // Wed, 05 Aug 2020 13:14:15 CST

// 输出 RFC3339 格式字符串
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339String() // 2020-08-05T13:14:15+08:00
// 输出 RFC3339Milli 格式字符串
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339MilliString() // 2020-08-05T13:14:15.999+08:00
// 输出 RFC3339Micro 格式字符串
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339MicroString() // 2020-08-05T13:14:15.999999+08:00
// 输出 RFC3339Nano 格式字符串
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339NanoString() // 2020-08-05T13:14:15.999999999+08:00

// 输出日期时间字符串
fmt.Printf("%s", carbon.Parse("2020-08-05 13:14:15")) // 2020-08-05 13:14:15

// 输出"2006-01-02 15:04:05.999999999 -0700 MST"格式字符串
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15.999999 +0800 CST

// 输出 "Jan 2, 2006" 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToFormattedDateString() // Aug 5, 2020
// 输出 "Mon, Jan 2, 2006" 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToFormattedDayDateString() // Wed, Aug 5, 2020

// 输出指定布局的字符串
carbon.Parse("2020-08-05 13:14:15").Layout(carbon.ISO8601Layout) // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").Layout("20060102150405") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Layout("2006年01月02日 15时04分05秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Layout("It is 2006-01-02 15:04:05") // It is 2020-08-05 13:14:15

// 输出指定格式的字符串(如果使用的字母与格式符号冲突时，请使用\符号转义该字符)
carbon.Parse("2020-08-05 13:14:15").Format("YmdHis") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Format("Y年m月d日 H时i分s秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Format("l jK \\o\\f F Y h:i:s A") // Wednesday 5th of August 2020 01:14:15 PM
carbon.Parse("2020-08-05 13:14:15").Format("\\I\\t \\i\\s Y-m-d H:i:s") // It is 2020-08-05 13:14:15
```

> 更多格式输出符号请查看附录 <a href="#format-sign-table">格式化符号表</a>

##### 星座

```go
// 获取星座
carbon.Parse("2020-08-05 13:14:15").Constellation() // 狮子座

// 是否是白羊座
carbon.Parse("2020-08-05 13:14:15").IsAries() // false
// 是否是金牛座
carbon.Parse("2020-08-05 13:14:15").IsTaurus() // false
// 是否是双子座
carbon.Parse("2020-08-05 13:14:15").IsGemini() // false
// 是否是巨蟹座
carbon.Parse("2020-08-05 13:14:15").IsCancer() // false
// 是否是狮子座
carbon.Parse("2020-08-05 13:14:15").IsLeo() // true
// 是否是处女座
carbon.Parse("2020-08-05 13:14:15").IsVirgo() // false
// 是否是天秤座
carbon.Parse("2020-08-05 13:14:15").IsLibra() // false
// 是否是天蝎座
carbon.Parse("2020-08-05 13:14:15").IsScorpio() // false
// 是否是射手座
carbon.Parse("2020-08-05 13:14:15").IsSagittarius() // false
// 是否是摩羯座
carbon.Parse("2020-08-05 13:14:15").IsCapricorn() // false
// 是否是水瓶座
carbon.Parse("2020-08-05 13:14:15").IsAquarius() // false
// 是否是双鱼座
carbon.Parse("2020-08-05 13:14:15").IsPisces() // false
```

##### 季节

> 按照气象划分，即3-5月为春季，6-8月为夏季，9-11月为秋季，12-2月为冬季

```go
// 获取季节
carbon.Parse("2020-08-05 13:14:15").Season() // 夏季

// 本季节开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfSeason().ToDateTimeString() // 2020-06-01 00:00:00
// 本季节结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfSeason().ToDateTimeString() // 2020-08-31 23:59:59

// 是否是春季
carbon.Parse("2020-08-05 13:14:15").IsSpring() // false
// 是否是夏季
carbon.Parse("2020-08-05 13:14:15").IsSummer() // true
// 是否是秋季
carbon.Parse("2020-08-05 13:14:15").IsAutumn() // false
// 是否是冬季
carbon.Parse("2020-08-05 13:14:15").IsWinter() // false
```

##### JSON

###### 内置字段类型

```go
type User struct {
	Date      carbon.LayoutType[carbon.Date]      `json:"date"`
	DateMilli carbon.LayoutType[carbon.DateMilli] `json:"date_milli"`
	DateMicro carbon.LayoutType[carbon.DateMicro] `json:"date_micro"`
	DateNano  carbon.LayoutType[carbon.DateNano]  `json:"date_nano"`
	
	Time      carbon.LayoutType[carbon.Time]      `json:"time"`
	TimeMilli carbon.LayoutType[carbon.TimeMilli] `json:"time_milli"`
	TimeMicro carbon.LayoutType[carbon.TimeMicro] `json:"time_micro"`
	TimeNano  carbon.LayoutType[carbon.TimeNano]  `json:"time_nano"`
	
	DateTime      carbon.FormatType[carbon.DateTime]      `json:"date_time"`
	DateTimeMilli carbon.FormatType[carbon.DateTimeMilli] `json:"date_time_milli"`
	DateTimeMicro carbon.FormatType[carbon.DateTimeMicro] `json:"date_time_micro"`
	DateTimeNano  carbon.FormatType[carbon.DateTimeNano]  `json:"date_time_nano"`
	
	Timestamp      carbon.TimestampType[carbon.Timestamp]      `json:"timestamp"`
	TimestampMilli carbon.TimestampType[carbon.TimestampMilli] `json:"timestamp_milli"`
	TimestampMicro carbon.TimestampType[carbon.TimestampMicro] `json:"timestamp_micro"`
	TimestampNano  carbon.TimestampType[carbon.TimestampNano]  `json:"timestamp_nano"`
}

var user User

c := carbon.Parse("2020-08-05 13:14:15.999999999")

user.Date = carbon.NewFormatType[carbon.Date](c)
user.DateMilli = carbon.NewLayoutType[carbon.DateMilli](c)
user.DateMicro = carbon.NewLayoutType[carbon.DateMicro](c)
user.DateNano = carbon.NewLayoutType[carbon.DateNano](c)

user.Time = carbon.NewLayoutType[carbon.Time](c)
user.TimeMilli = carbon.NewLayoutType[carbon.TimeMilli](c)
user.TimeMicro = carbon.NewLayoutType[carbon.TimeMicro](c)
user.TimeNano = carbon.NewLayoutType[carbon.TimeNano](c)

user.DateTime = carbon.NewFormatType[carbon.DateTime](c)
user.DateTimeMilli = carbon.NewFormatType[carbon.DateTimeMilli](c)
user.DateTimeMicro = carbon.NewFormatType[carbon.DateTimeMicro](c)
user.DateTimeNano = carbon.NewFormatType[carbon.DateTimeNano](c)

user.Timestamp = carbon.NewTimestampType[carbon.Timestamp](c)
user.TimestampMilli = carbon.NewTimestampType[carbon.TimestampMilli](c)
user.TimestampMicro = carbon.NewTimestampType[carbon.TimestampMicro](c)
user.TimestampNano = carbon.NewTimestampType[carbon.TimestampNano](c)

data, err := json.Marshal(&user)
if err != nil {
	// 错误处理
	log.Fatal(err)
}
fmt.Printf("%s\n", data)
// 输出
{
	"date": "2020-08-05",
	"date_milli": "2020-08-05.999",
	"date_micro": "2020-08-05.999999",
	"date_nano": "2020-08-05.999999999",
	"time": "13:14:15",
	"time_milli": "13:14:15.999",
	"time_micro": "13:14:15.999999",
	"time_nano": "13:14:15.999999999",
	"date_time": "2020-08-05 13:14:15",
	"date_time_milli": "2020-08-05 13:14:15.999",
	"date_time_micro": "2020-08-05 13:14:15.999999",
	"date_time_nano": "2020-08-05 13:14:15.999999999",
	"timestamp": 1596633255,
	"timestamp_milli": 1596633255999,
	"timestamp_micro": 1596633255999999,
	"timestamp_nano": 1596633255999999999
}

var person User
err := json.Unmarshal(data, &person)
if err != nil {
	// 错误处理
	log.Fatal(err)
}

fmt.Printf("person: %+v\n", person)
// 输出
person: {Date:2020-08-05 DateMilli:2020-08-05.999 DateMicro:2020-08-05.999999 DateNano:2020-08-05.999999999 Time:13:14:15 TimeMilli:13:14:15.999 TimeMicro:13:14:15.999999 TimeNano:13:14:15.999999999 DateTime:2020-08-05 13:14:15 DateTimeMilli:2020-08-05 13:14:15.999 DateTimeMicro:2020-08-05 13:14:15.999999 DateTimeNano:2020-08-05 13:14:15.999999999 Timestamp:1596633255 TimestampMilli:1596633255999 TimestampMicro:1596633255999999 TimestampNano:1596633255999999999}
```

###### 自定义字段类型

```go
type RFC3339Layout string
	func (t CustomerLayout) SetLayout() string {
	return carbon.RFC3339Layout
}

type ISO8601Format string
	func (t CustomerFormat) SetFormat() string {
	return carbon.ISO8601Format
}

type User struct {
	Customer1 carbon.LayoutType[RFC3339Layout] `json:"customer1"`
	Customer2 carbon.FormatType[ISO8601Format] `json:"customer2"`
}

var user User

c := carbon.Parse("2020-08-05 13:14:15")

user.Customer1 = carbon.NewLayoutType[RFC3339Layout](c)
user.Customer2 = carbon.NewFormatType[ISO8601Format](c)

data, err := json.Marshal(&user)
if err != nil {
	// 错误处理
	log.Fatal(err)
}
fmt.Printf("%s\n", data)
// 输出
{"customer1":"2020-08-05T13:14:15Z", "customer2":"2020-08-05T13:14:15+00:00"}

var person User
err := json.Unmarshal(data, &person)
if err != nil {
	// 错误处理
	log.Fatal(err)
}

fmt.Printf("person: %+v\n", person)
// 输出
person: {Customer1:2020-08-05T13:14:15Z Customer2:2020-08-05T13:14:15+00:00}
```

##### 日历

目前支持的日历有

* [儒略日/简化儒略日](./calendar/julian/README.cn.md "儒略日/简化儒略日")
* [中国农历](./calendar/lunar/README.cn.md "中国农历")
* [波斯历/伊朗历](./calendar/persian/README.cn.md "波斯历/伊朗历")

##### 国际化

目前支持的语言有

* [简体中文(zh-CN)](./lang/zh-CN.json "简体中文"): 由 [gouguoyin](https://github.com/gouguoyin "gouguoyin") 翻译
* [繁体中文(zh-TW)](./lang/zh-TW.json "繁体中文"): 由 [gouguoyin](https://github.com/gouguoyin "gouguoyin") 翻译
* [英语(en)](./lang/en.json "英语"): 由 [gouguoyin](https://github.com/gouguoyin "gouguoyin") 翻译
* [日语(jp)](./lang/jp.json "日语"): 由 [gouguoyin](https://github.com/gouguoyin "gouguoyin") 翻译
* [韩语(kr)](./lang/kr.json "韩语"): 由 [nannul](https://github.com/nannul "nannul") 翻译
* [德语(de)](./lang/de.json "德语"): 由 [benzammour](https://github.com/benzammour "benzammour") 翻译
* [西班牙语(es)](./lang/es.json "西班牙语"): 由 [hgisinger](https://github.com/hgisinger "hgisinger") 翻译
* [土耳其语(tr)](./lang/tr.json "土耳其语"): 由 [emresenyuva](https://github.com/emresenyuva "emresenyuva") 翻译
* [葡萄牙语(pt)](./lang/pt.json "葡萄牙语"): 由 [felipear89](https://github.com/felipear89 "felipear89") 翻译
* [俄罗斯语(ru)](./lang/ru.json "俄罗斯语"): 由 [zemlyak](https://github.com/zemlyak "zemlyak") 翻译
* [乌克兰语(uk)](./lang/uk.json "乌克兰语"): 由 [open-git](https://github.com/open-git "open-git") 翻译
* [罗马尼亚语(ro)](./lang/ro.json "罗马尼亚语"): 由 [DrOctavius](https://github.com/DrOctavius "DrOctavius") 翻译
* [印度尼西亚语(id)](./lang/id.json "印度尼西亚语"): 由 [justpoypoy](https://github.com/justpoypoy "justpoypoy") 翻译
* [意大利语(it)](./lang/it.json "意大利语"): 由 [nicoloHevelop](https://github.com/justpoypoy "nicoloHevelop") 翻译
* [马来西亚巴哈马语(ms-MY)](./lang/ms-MY.json "马来西亚巴哈马语"):
	由 [hollowaykeanho](https://github.com/hollowaykeanho "hollowaykeanho") 翻译
* [法语(fr)](./lang/fr.json "法语"): 由 [hollowaykeanho](https://github.com/hollowaykeanho "hollowaykeanho") 翻译
* [泰语(th)](./lang/th.json "泰语"): 由 [izcream](https://github.com/izcream "izcream") 翻译
* [瑞典语(se)](./lang/se.json "瑞典语"): 由 [jwanglof](https://github.com/jwanglof "jwanglof") 翻译
* [波斯语(fa)](./lang/fa.json "波斯语"): 由 [erfanMomeniii](https://github.com/ErfanMomeniii "ErfanMomeniii") 翻译
* [波兰语(nl)](./lang/nl.json "波兰语"): 由 [RemcoE33](https://github.com/RemcoE33 "RemcoE33") 翻译
* [越南语(vi)](./lang/vi.json "越南语"): 由 [culy247](https://github.com/culy247 "culy247") 翻译
* [印地语(hi)](./lang/hi.json "印地语"): 由 [chauhan17nitin](https://github.com/chauhan17nitin "chauhan17nitin") 翻译
* [波兰语(pl)](./lang/pl.json "波兰语"): 由 [gouguoyin](https://github.com/gouguoyin "gouguoyin") 翻译
* [保加利亚语(bg)](./lang/bg.json "保加利亚语"): 由 [yuksbg](https://github.com/yuksbg "yuksbg") 翻译
* [阿拉伯语(ar)](./lang/bg.json "阿拉伯语"): 由 [zumoshi](https://github.com/zumoshi "zumoshi") 翻译
* [匈牙利语(hu)](./lang/hu.json "匈牙利语"): 由 [kenlas](https://github.com/kenlas "kenlas") 翻译
* [丹麦语(dk)](./lang/dk.json "丹麦语"): 由 [Munk91](https://github.com/Munk91 "Munk91") 翻译
* [挪威语(nb)](./lang/nb.json "挪威语"): 由 [bendikrb](https://github.com/bendikrb "bendikrb") 翻译

目前支持的方法有

* `Constellation()`：获取星座，如 `白羊座`
* `Season()`：获取季节，如 `夏季`
* `DiffForHumans()`：获取对人类友好的可读格式时间差，如 `一小时前`
* `ToMonthString()`：输出完整月份字符串，如 `一月`
* `ToShortMonthString()`：输出缩写月份字符串，如 `1月`
* `ToWeekString()`：输出完整星期字符串，如 `星期一`
* `ToShortWeekString()`：输出缩写星期字符串，如 `周一`

###### 设置区域

```go
lang := carbon.NewLanguage()
lang.SetLocale("zh-CN")

c := carbon.SetLanguage(lang)
if c.Error != nil {
	// 错误处理
	log.Fatal(c.Error)
}

c.Now().AddHours(1).DiffForHumans() // 1 小时后
c.Now().AddHours(1).ToMonthString() // 八月
c.Now().AddHours(1).ToShortMonthString() // 8月
c.Now().AddHours(1).ToWeekString() // 星期二
c.Now().AddHours(1).ToShortWeekString() // 周二
c.Now().AddHours(1).Constellation() // 狮子座
c.Now().AddHours(1).Season() // 夏季
```

###### 重写部分翻译资源(其余仍然按照指定的 `locale` 文件内容翻译)

```go
lang := carbon.NewLanguage()

resources := map[string]string {
	"hour": "%dh",
}
lang.SetLocale("en").SetResources(resources)

c := carbon.SetLanguage(lang)
if c.Error != nil {
	// 错误处理
	log.Fatal(c.Error)
}

c.Now().AddYears(1).DiffForHumans() // 1 year from now
c.Now().AddHours(1).DiffForHumans() // 1h from now
c.Now().ToMonthString() // August
c.Now().ToShortMonthString() // Aug
c.Now().ToWeekString() // Tuesday
c.Now().ToShortWeekString() // Tue
c.Now().Constellation() // Leo
c.Now().Season() // Summer
```

###### 重写全部翻译资源(无需指定 `locale`)

```go
lang := carbon.NewLanguage()
resources := map[string]string {
	"months": "january|february|march|april|may|june|july|august|september|october|november|december",
	"short_months": "jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec",
	"weeks": "sunday|monday|tuesday|wednesday|thursday|friday|saturday",
	"short_weeks": "sun|mon|tue|wed|thu|fri|sat",
	"seasons": "spring|summer|autumn|winter",
	"constellations": "aries|taurus|gemini|cancer|leo|virgo|libra|scorpio|sagittarius|capricornus|aquarius|pisce",
	"year": "1 yr|%d yrs",
	"month": "1 mo|%d mos",
	"week": "%dw",
	"day": "%dd",
	"hour": "%dh",
	"minute": "%dm",
	"second": "%ds",
	"now": "just now",
	"ago": "%s ago",
	"from_now": "in %s",
	"before": "%s before",
	"after": "%s after",
}
lang.SetResources(resources)

c := carbon.SetLanguage(lang)
c.Now().AddYears(1).DiffForHumans() // in 1 yr
c.Now().AddHours(1).DiffForHumans() // in 1h
c.Now().ToMonthString() // august
c.Now().ToShortMonthString() // aug
c.Now().ToWeekString() // tuesday
c.Now().ToShortWeekString() // tue
c.Now().Constellation() // leo
c.Now().Season() // summer
```

##### 错误处理

```go
c := carbon.Parse("2020-08-05").SetTimezone("xxx")
if c.HasError() {
	// 错误处理
	log.Fatal(c.Error)
}
// 输出
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezones
```

#### 附录

##### <a id="format-sign-table">格式符号表</a>

| 符号 |             描述             | 长度 |        范围        |         示例          |
|:--:|:--------------------------:|:--:|:----------------:|:-------------------:|
| d  |        月份中的第几天，有前导零        | 2  |      01-31       |         02          |
| D  |         缩写单词表示的周几          | 3  |     Mon-Sun      |         Mon         |
| j  |       月份中的第几天，没有前导零        | -  |       1-31       |          2          |
| K  |    第几天的英文缩写后缀，一般和j配合使用     | 2  |   st/nd/rd/th    |         th          |
| l  |         完整单词表示的周几          | -  |  Monday-Sunday   |       Monday        |
| F  |         完整单词表示的月份          | -  | January-December |       January       |
| m  |        数字表示的月份，有前导零        | 2  |      01-12       |         01          |
| M  |         缩写单词表示的月份          | 3  |     Jan-Dec      |         Jan         |
| n  |       数字表示的月份，没有前导零        | -  |       1-12       |          1          |
| Y  |        4 位数字完整表示的年份        | 4  |    0000-9999     |        2006         |
| y  |         2 位数字表示的年份         | 2  |      00-99       |         06          |
| a  |         小写的上午和下午标识         | 2  |      am/pm       |         pm          |
| A  |         大写的上午和下午标识         | 2  |      AM/PM       |         PM          |
| g  |         小时，12 小时格式         | -  |       1-12       |          3          |
| G  |         小时，24 小时格式         | -  |       0-23       |         15          |
| h  |         小时，12 小时格式         | 2  |      00-11       |         03          |
| H  |         小时，24 小时格式         | 2  |      00-23       |         15          |
| i  |             分钟             | 2  |      01-59       |         04          |
| s  |             秒数             | 2  |      01-59       |         05          |
| O  |       与格林威治时间相差的小时数        | -  |        -         |        -0700        |
| P  | 与格林威治时间相差的小时数，小时和分钟之间有冒号分隔 | -  |        -         |       -07:00        |
| Z  |            时区名字            | -  |        -         |         CST         |
| W  |   ISO8601 格式数字表示的年份中的第几周   | 2  |      01-52       |         01          |
| N  |   ISO8601 格式数字表示的星期中的第几天   | 2  |      01-07       |         02          |
| L  |    是否为闰年，如果是闰年为 1，否则为 0    | 1  |       0-1        |          0          |
| S  |           秒级时间戳            | -  |        -         |     1596604455      |
| U  |           毫级时间戳            | -  |        -         |    1596604455666    |
| V  |           微级时间戳            | -  |        -         |  1596604455666666   |
| X  |           纳级时间戳            | -  |        -         | 1596604455666666666 |
| u  |             毫秒             | -  |      1-999       |         999         |
| v  |             微秒             | -  |     1-999999     |       999999        |
| x  |             纳秒             | -  |   1-999999999    |      999999999      |
| w  |          数字表示的周几           | 1  |       0-6        |          1          |
| t  |          月份中的总天数           | 2  |      28-31       |         31          |
| z  |            时区位置            | -  |        -         |    Asia/Shanghai    |
| o  |           时区偏移量            | -  |        -         |        28800        |
| q  |            当前季节            | 1  |       1-4        |          1          |
| c  |           当前世纪数            | -  |       0-99       |         21          |

#### 常见问题

1、v2.5.x 和 v2.6.x 版本有什么区别？
> v2.5.x 及以下版本是值传递， v2.6.x 及以上版本是指针传递，并且使用了泛型实现了 JSON 编码输出格式自定义。两个版本都会长期维护，但是强烈建议使用 v2.6.x 版本。

2、window 系统部署时时区报错

> `window` 系统如果没有安装 `golang` 环境，部署时会报 `GOROOT/lib/time/zoneinfo.zip: no such file or directory` 异常，原因是由于
> `window` 系统没有内置时区文件，只需要手动下载并指定 `zoneinfo.zip` 路径即可，如 `go/lib/time/zoneinfo.zip`

```go
os.Setenv("ZONEINFO", "./go/lib/time/zoneinfo.zip")
```

3、docker 容器部署时时区报错

> `docker` 容器如果没有安装 `golang` 环境，部署时会报 `open /usr/local/go/lib/time/zoneinfo.zip: no such file or directory`
> 异常，只需要把 `zoneinfo.zip` 复制到容器中即可，即在 `Dockerfile` 中加入

```go
COPY ./zoneinfo.zip /usr/local/go /lib/time/zoneinfo.zip
```

#### 参考项目

* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)
* [nodatime/nodatime](https://github.com/nodatime/nodatime)
* [jinzhu/now](https://github.com/jinzhu/now)
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)
* [jodaOrg/joda-time](https://github.com/jodaOrg/joda-time)
* [arrow-py/arrow](https://github.com/arrow-py/arrow)
* [moment/moment](https://github.com/moment/moment)
* [iamkun/dayjs](https://github.com/iamkun/dayjs)

#### 贡献者

感谢以下所有为 `Carbon` 做出贡献的人：

<a href="https://github.com/dromara/carbon/graphs/contributors"><img src="https://contrib.rocks/image?repo=dromara/carbon&max=100&columns=16"/></a>

#### 赞助

`Carbon` 是一个非商业开源项目, 如果你想支持 `Carbon`,
你可以为开发者 [购买一杯咖啡](https://www.gouguoyin.com/zanzhu.html)

#### 致谢

`Carbon`已获取免费的 JetBrains 开源许可证，在此表示感谢

<a href="https://www.jetbrains.com"><img src="https://foruda.gitee.com/images/1704325523163241662/1bf21f86_544375.png" height="100" alt="JetBrains"/></a>
