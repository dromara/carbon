# Carbon  #

[![Carbon Release](https://img.shields.io/github/release/golang-module/carbon.svg)](https://github.com/golang-module/carbon/releases)
[![Go Test](https://github.com/golang-module/carbon/actions/workflows/test.yml/badge.svg)](https://github.com/golang-module/carbon/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-module/carbon)](https://goreportcard.com/report/github.com/golang-module/carbon)
[![Go Coverage](https://codecov.io/gh/golang-module/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/golang-module/carbon)
[![Carbon Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/golang-module/carbon)
![License](https://img.shields.io/github/license/golang-module/carbon)

简体中文 | [English](README.md) | [日本語](README.jp.md)

一个轻量级、语义化、对开发者友好的 golang 时间处理库，支持链式调用

Carbon 已被 [awesome-go](https://github.com/avelino/awesome-go#date-and-time "awesome-go") 收录, 如果您觉得不错，请给个 star 吧

[github.com/golang-module/carbon](https://github.com/golang-module/carbon "github.com/golang-module/carbon")

[gitee.com/go-package/carbon](https://gitee.com/go-package/carbon "gitee.com/go-package/carbon")

#### 安装使用

##### Golang 版本小于1.16

```go
// 使用 github 库
go get -u github.com/golang-module/carbon

import (
    "github.com/golang-module/carbon"
)

// 使用 gitee 库
go get -u gitee.com/go-package/carbon

import (
    "gitee.com/go-package/carbon"
)
```

##### Golang 版本大于等于1.16

```go
// 使用 github 库
go get -u github.com/golang-module/carbon/v2

import (
    "github.com/golang-module/carbon/v2"
)

// 使用 gitee 库
go get -u gitee.com/go-package/carbon/v2

import (
    "gitee.com/go-package/carbon/v2"
)
```

#### 用法示例

> 默认时区为 Local，即服务器所在时区，假设当前时间为 2020-08-05 13:14:15

##### 昨天、今天、明天

```go
// 今天此刻
fmt.Sprintf("%s", carbon.Now()) // 2020-08-05 13:14:15
carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
// 今天日期
carbon.Now().ToDateString() // 2020-08-05
// 今天时间
carbon.Now().ToTimeString() // 13:14:15
// 指定时区的今天此刻
carbon.Now(Carbon.NewYork).ToDateTimeString() // 2020-08-05 14:14:15
carbon.SetTimezone(Carbon.NewYork).Now().ToDateTimeString() // 2020-08-05 14:14:15
// 今天秒级时间戳
carbon.Now().Timestamp() // 1596604455
carbon.Now().TimestampWithSecond() // 1596604455
// 今天毫秒级时间戳
carbon.Now().TimestampWithMillisecond() // 1596604455000
// 今天微秒级时间戳
carbon.Now().TimestampWithMicrosecond() // 1596604455000000
// 今天纳秒级时间戳
carbon.Now().TimestampWithNanosecond() // 1596604455000000000

// 昨天此刻
fmt.Sprintf("%s", carbon.Yesterday()) // 2020-08-04 13:14:15
carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// 昨天日期
carbon.Yesterday().ToDateString() // 2020-08-04
// 昨天时间
carbon.Yesterday().ToTimeString() // 13:14:15
// 指定日期的昨天此刻
carbon.Parse("2021-01-28 13:14:15").Yesterday().ToDateTimeString() // 2021-01-27 13:14:15
// 指定时区的昨天此刻
carbon.Yesterday(Carbon.NewYork).ToDateTimeString() // 2020-08-04 14:14:15
carbon.SetTimezone(Carbon.NewYork).Yesterday().ToDateTimeString() // 2020-08-04 14:14:15
// 昨天秒级时间戳
carbon.Yesterday().Timestamp() // 1596518055
carbon.Yesterday().TimestampWithSecond() // 1596518055
// 昨天毫秒级时间戳
carbon.Yesterday().TimestampWithMillisecond() // 1596518055000
// 昨天微秒级时间戳
carbon.Yesterday().TimestampWithMicrosecond() // 1596518055000000
// 昨天纳秒级时间戳
carbon.Yesterday().TimestampWithNanosecond() // 1596518055000000000

// 明天此刻
fmt.Sprintf("%s", carbon.Tomorrow()) // 2020-08-06 13:14:15
carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// 明天日期
carbon.Tomorrow().ToDateString() // 2020-08-06
// 明天时间
carbon.Tomorrow().ToTimeString() // 13:14:15
// 指定日期的明天此刻
carbon.Parse("2021-01-28 13:14:15").Tomorrow().ToDateTimeString() // 2021-01-29 13:14:15
// 指定时区的明天此刻
carbon.Tomorrow(Carbon.NewYork).ToDateTimeString() // 2020-08-06 14:14:15
carbon.SetTimezone(Carbon.NewYork).Tomorrow().ToDateTimeString() // 2020-08-06 14:14:15
// 明天秒级时间戳
carbon.Tomorrow().Timestamp() // 1596690855
carbon.Tomorrow().TimestampWithSecond() // 1596690855
// 明天毫秒级时间戳
carbon.Tomorrow().TimestampWithMillisecond() // 1596690855000
// 明天微秒级时间戳
carbon.Tomorrow().TimestampWithMicrosecond() // 1596690855000000
// 明天纳秒级时间戳
carbon.Tomorrow().TimestampWithNanosecond() // 1596690855000000000
```

##### 创建 Carbon 实例

```go
// 从秒级时间戳创建 Carbon 实例
carbon.CreateFromTimestamp(-1).ToDateTimeString() // 1970-01-01 07:59:59
carbon.CreateFromTimestamp(-1, carbon.Tokyo).ToDateTimeString() // 1970-01-01 08:59:59
carbon.CreateFromTimestamp(0).ToDateTimeString() // 1970-01-01 08:00:00
carbon.CreateFromTimestamp(0, carbon.Tokyo).ToDateTimeString() // 1970-01-01 09:00:00
carbon.CreateFromTimestamp(1596604455).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromTimestamp(1596604455, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// 从毫秒级时间戳创建 Carbon 实例
carbon.CreateFromTimestamp(1596604455000).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromTimestamp(1596604455000, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// 从微秒级时间戳创建 Carbon 实例
carbon.CreateFromTimestamp(1596604455000000).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromTimestamp(1596604455000000, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// 从纳级时间戳创建 Carbon 实例
carbon.CreateFromTimestamp(1596604455000000000).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromTimestamp(1596604455000000000, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15

// 从年月日时分秒创建 Carbon 实例
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// 从年月日创建 Carbon 实例(时分秒默认为当前时分秒)
carbon.CreateFromDate(2020, 8, 5).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromDate(2020, 8, 5, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// 从时分秒创建 Carbon 实例(年月日默认为当前年月日)
carbon.CreateFromTime(13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
carbon.CreateFromTime(13, 14, 15, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### 将标准格式时间字符串解析成 Carbon 实例

```go
carbon.Parse("").ToDateTimeString() // 空字符串
carbon.Parse("0").ToDateTimeString() // 空字符串
carbon.Parse("0000-00-00 00:00:00").ToDateTimeString() // 空字符串
carbon.Parse("0000-00-00").ToDateTimeString() // 空字符串
carbon.Parse("00:00:00").ToDateTimeString() // 空字符串

carbon.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05").ToDateTimeString() // 2020-08-05 00:00:00
carbon.Parse("20200805131415").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("20200805").ToDateTimeString() // 2020-08-05 00:00:00
carbon.Parse("2020-08-05T13:14:15+08:00").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### 通过格式化字符将字符串解析成 Carbon 实例

> 如果使用的字母与格式化字符冲突时，请使用转义符转义该字母

```go
carbon.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("2020-08-05 13:14:15", "Y-m-d H:i:s", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### 通过布局字符将字符串解析成 Carbon 实例

```go
carbon.ParseByLayout("2020|08|05 13|14|15", "2006|01|02 15|04|05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("2020-08-05 13:14:15", "2006-01-02 15:04:05", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### Carbon 和 time.Time 互转

```go
// 将 time.Time 转换成 Carbon
carbon.Time2Carbon(time.Now())
// 将 Carbon 转换成 time.Time
carbon.Now().Carbon2Time()
```

##### 开始时间、结束时间

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
carbon.Parse("2020-08-05 13:14:15").StartOfSecond().Format("Y-m-d H:i:s.u") // 2020-08-05 13:14:15.0
// 本秒结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfSecond().Format("Y-m-d H:i:s.u") // 2020-08-05 13:14:15.999
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
carbon.Parse("2020-02-29 13:14:15").Decades(3).ToDateTimeString() // 2050-03-01 13:14:15
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
carbon.Parse("2019-08-31 13:14:15").AddQuarters(3).ToDateTimeString() // 2019-03-02 13:14:15
// 三个季度后(月份不溢出)
carbon.Parse("2019-08-31 13:14:15").AddQuartersNoOverflow(3).ToDateTimeString() // 2019-02-29 13:14:15
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
```

##### 时间差

```go
// 相差多少年
carbon.Parse("2021-08-05 13:14:15").DiffInYears(carbon.Parse("2020-08-05 13:14:15")) // -1
// 相差多少年（绝对值）
carbon.Parse("2021-08-05 13:14:15").DiffInYearsWithAbs(carbon.Parse("2020-08-05 13:14:15")) // 1

// 相差多少月
carbon.Parse("2020-08-05 13:14:15").DiffInMonths(carbon.Parse("2020-07-05 13:14:15")) // -1
// 相差多少月（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffInMonthsWithAbs(carbon.Parse("2020-07-05 13:14:15")) // 1

// 相差多少周
carbon.Parse("2020-08-05 13:14:15").DiffInWeeks(carbon.Parse("2020-07-28 13:14:15")) // -1
// 相差多少周（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffInWeeksWithAbs(carbon.Parse("2020-07-28 13:14:15")) // 1

// 相差多少天
carbon.Parse("2020-08-05 13:14:15").DiffInDays(carbon.Parse("2020-08-04 13:14:15")) // -1
// 相差多少天（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffInDaysWithAbs(carbon.Parse("2020-08-04 13:14:15")) // 1

// 相差多少小时
carbon.Parse("2020-08-05 13:14:15").DiffInHours(carbon.Parse("2020-08-05 12:14:15")) // -1
// 相差多少小时（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffInHoursWithAbs(carbon.Parse("2020-08-05 12:14:15")) // 1

// 相差多少分
carbon.Parse("2020-08-05 13:14:15").DiffInMinutes(carbon.Parse("2020-08-05 13:13:15")) // -1
// 相差多少分（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffInMinutesWithAbs(carbon.Parse("2020-08-05 13:13:15")) // 1

// 相差多少秒
carbon.Parse("2020-08-05 13:14:15").DiffInSeconds(carbon.Parse("2020-08-05 13:14:14")) // -1
// 相差多少秒（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffInSecondsWithAbs(carbon.Parse("2020-08-05 13:14:14")) // 1

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

##### 时间判断

```go
// 是否是零值时间
carbon.Parse("").IsZero() // true
carbon.Parse("0").IsZero() // true
carbon.Parse("0000-00-00 00:00:00").IsZero() // true
carbon.Parse("0000-00-00").IsZero() // true
carbon.Parse("00:00:00").IsZero() // true
carbon.Parse("2020-08-05 00:00:00").IsZero() // false
carbon.Parse("2020-08-05").IsZero() // false
carbon.Parse("2020-08-05").SetTimezone("xxx").IsZero() // false

// 是否是无效时间
carbon.Parse("").IsInvalid() // true
carbon.Parse("0").IsInvalid() // true
carbon.Parse("0000-00-00 00:00:00").IsInvalid() // true
carbon.Parse("0000-00-00").IsInvalid() // true
carbon.Parse("00:00:00").IsInvalid() // true
carbon.Parse("2020-08-05 00:00:00").IsInvalid() // false
carbon.Parse("2020-08-05").IsInvalid() // false
carbon.Parse("2020-08-05").SetTimezone("xxx").IsInvalid() // true

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
carbon.Parse("2020-08-05 13:14:15").IsApril()  // false
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
carbon.SetTimezone(carbon.PRC).Now().ToDateTimeString() // 2020-08-05 13:14:15
carbon.SetTimezone(carbon.Tokyo).Now().ToDateTimeString() // 2020-08-05 14:14:15
carbon.SetTimezone(carbon.Tokyo).Now().SetTimezone(carbon.PRC).ToDateTimeString() // 2020-08-05 12:14:15

// 设置区域
carbon.Parse("2020-07-05 13:14:15").SetLocale("en").DiffForHumans() // 1 month ago
carbon.Parse("2020-07-05 13:14:15").SetLocale("zh-CN").DiffForHumans() // 1 月前

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
carbon.Parse("2020-08-05 13:14:15").Millisecond() // 1596604455000
// 获取当前微秒
carbon.Parse("2020-08-05 13:14:15").Microsecond() // 1596604455000000
// 获取当前纳秒
carbon.Parse("2020-08-05 13:14:15").Nanosecond() // 1596604455000000000

// 获取秒级时间戳, Timestamp() 是TimestampWithSecond()的简写
carbon.Parse("2020-08-05 13:14:15").Timestamp() // 1596604455
carbon.Parse("2020-08-05 13:14:15").TimestampWithSecond() // 1596604455
// 获取毫秒级时间戳
carbon.Parse("2020-08-05 13:14:15").TimestampWithMillisecond() // 1596604455000
// 获取微秒级时间戳
carbon.Parse("2020-08-05 13:14:15").TimestampWithMicrosecond() // 1596604455000000
// 获取纳秒级时间戳
carbon.Parse("2020-08-05 13:14:15").TimestampWithNanosecond() // 1596604455000000000

// 获取时区
carbon.SetTimezone(carbon.PRC).Timezone() // CST
carbon.SetTimezone(carbon.Tokyo).Timezone() // JST

// 获取位置
carbon.SetTimezone(carbon.PRC).Location() // PRC
carbon.SetTimezone(carbon.Tokyo).Location() // Asia/Tokyo

// 获取距离UTC时区的偏移量，单位秒
carbon.SetTimezone(carbon.PRC).Offset() // 28800
carbon.SetTimezone(carbon.Tokyo).Offset() // 32400

// 获取当前区域
carbon.Now().Locale() // en
carbon.Now().SetLocale("zh-CN").Locale() // zh-CN

// 获取当前星座
carbon.Now().Constellation() // Leo
carbon.Now().SetLocale("en").Constellation() // Leo
carbon.Now().SetLocale("zh-CN").Constellation() // 狮子座

// 获取当前季节
carbon.Now().Season() // Summer
carbon.Now().SetLocale("en").Season() // Summer
carbon.Now().SetLocale("zh-CN").Season() // 夏季

// 获取年龄
carbon.Parse("2002-01-01 13:14:15").Age() // 17
carbon.Parse("2002-12-31 13:14:15").Age() // 18

```

##### 时间输出

```go
// 输出日期时间字符串
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString(carbon.Tokyo) // 2020-08-05 14:14:15
// 输出简写日期时间字符串
carbon.Parse("2020-08-05 13:14:15").ToShortDateTimeString() // 20200805131415
carbon.Parse("2020-08-05 13:14:15").ToShortDateTimeString(carbon.Tokyo) // 20200805141415

// 输出日期字符串
carbon.Parse("2020-08-05 13:14:15").ToDateString() // 2020-08-05
carbon.Parse("2020-08-05 13:14:15").ToDateString(carbon.Tokyo) // 2020-08-05
// 输出简写日期字符串
carbon.Parse("2020-08-05 13:14:15").ToShortDateString() // 20200805
carbon.Parse("2020-08-05 13:14:15").ToShortDateString(carbon.Tokyo) // 20200805

// 输出时间字符串
carbon.Parse("2020-08-05 13:14:15").ToTimeString() // 13:14:15
carbon.Parse("2020-08-05 13:14:15").ToTimeString(carbon.Tokyo) // 14:14:15
// 输出简写时间字符串
carbon.Parse("2020-08-05 13:14:15").ToShortTimeString() // 131415
carbon.Parse("2020-08-05 13:14:15").ToShortTimeString(carbon.Tokyo) // 141415

// 输出 Ansic 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToAnsicString() // Wed Aug  5 13:14:15 2020
carbon.Parse("2020-08-05 13:14:15").ToAnsicString(carbon.Tokyo) // Wed Aug  5 14:14:15 2020
// 输出 Atom 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToAtomString() // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").ToAtomString(carbon.Tokyo) // 2020-08-05T14:14:15+08:00
// 输出 UnixDate 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToUnixDateString() // Wed Aug  5 13:14:15 CST 2020
carbon.Parse("2020-08-05 13:14:15").ToUnixDateString(carbon.Tokyo) // Wed Aug  5 14:14:15 JST 2020
// 输出 RubyDate 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRubyDateString() // Wed Aug 05 13:14:15 +0800 2020
carbon.Parse("2020-08-05 13:14:15").ToRubyDateString(carbon.Tokyo) // Wed Aug 05 14:14:15 +0900 2020
// 输出 Kitchen 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToKitchenString() // 1:14PM
carbon.Parse("2020-08-05 13:14:15").ToKitchenString(carbon.Tokyo) // 2:14PM
// 输出 Cookie 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToCookieString() // Wednesday, 05-Aug-2020 13:14:15 CST
carbon.Parse("2020-08-05 13:14:15").ToCookieString(carbon.Tokyo) // Wednesday, 05-Aug-2020 14:14:15 JST
// 输出 DayDateTime 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString() // Wed, Aug 5, 2020 1:14 PM
carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString(carbon.Tokyo) // Wed, Aug 5, 2020 2:14 PM
// 输出 RSS 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRssString() // Wed, 05 Aug 2020 13:14:15 +0800
carbon.Parse("2020-08-05 13:14:15").ToRssString(carbon.Tokyo) // Wed, 05 Aug 2020 14:14:15 +0900
// 输出 W3C 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToW3cString() // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").ToW3cString(carbon.Tokyo) // 2020-08-05T14:14:15+09:00

// 输出 ISO8601 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToIso8601String() // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").ToIso8601String(carbon.Tokyo) // 2020-08-05T14:14:15+09:00
// 输出 RFC822 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc822String() // 05 Aug 20 13:14 CST
carbon.Parse("2020-08-05 13:14:15").ToRfc822String(carbon.Tokyo) // 05 Aug 20 14:14 JST
// 输出 RFC822Z 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc822zString() // 05 Aug 20 13:14 +0800
carbon.Parse("2020-08-05 13:14:15").ToRfc822zString(carbon.Tokyo) // 05 Aug 20 14:14 +0900
// 输出 RFC850 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc850String() // Wednesday, 05-Aug-20 13:14:15 CST
carbon.Parse("2020-08-05 13:14:15").ToRfc850String(carbon.Tokyo) // Wednesday, 05-Aug-20 14:14:15 JST
// 输出 RFC1036 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc1036String() // Wed, 05 Aug 20 13:14:15 +0800
carbon.Parse("2020-08-05 13:14:15").ToRfc1036String(carbon.Tokyo) // Wed, 05 Aug 20 14:14:15 +0900
// 输出 RFC1123 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc1123String() // Wed, 05 Aug 2020 13:14:15 CST
carbon.Parse("2020-08-05 13:14:15").ToRfc1123String(carbon.Tokyo) // Wed, 05 Aug 2020 14:14:15 JST
// 输出 RFC1123Z 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc1123zString() // Wed, 05 Aug 2020 13:14:15 +0800
carbon.Parse("2020-08-05 13:14:15").ToRfc1123zString(carbon.Tokyo) // Wed, 05 Aug 2020 14:14:15 0800
// 输出 RFC2822 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc2822String() // Wed, 05 Aug 2020 13:14:15 +0800
carbon.Parse("2020-08-05 13:14:15").ToRfc2822String(carbon.Tokyo) // Wed, 05 Aug 2020 14:14:15 +0900
// 输出 RFC3339 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc3339String() // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").ToRfc3339String(carbon.Tokyo) // 2020-08-05T14:14:15+09:00
// 输出 RFC7231 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc7231String() // Wed, 05 Aug 2020 13:14:15 GMT
carbon.Parse("2020-08-05 13:14:15").ToRfc7231String(carbon.Tokyo) // Wed, 05 Aug 2020 14:14:15 GMT

// 输出日期时间字符串
fmt.Sprintf("%s", carbon.Parse("2020-08-05 13:14:15")) // 2020-08-05 13:14:15
fmt.Sprintf("%s", carbon.Parse("2020-08-05 13:14:15", carbon.Tokyo)) // 2020-08-05 13:14:15

// 输出"2006-01-02 15:04:05.999999999 -0700 MST"格式字符串
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-05 13:14:15").ToString(carbon.Tokyo) // 2020-08-05 14:14:15 +0900 JST

// 输出指定布局的字符串,Layout()是ToLayoutString()的简写
carbon.Parse("2020-08-05 13:14:15").Layout("20060102150405") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Layout("2006年01月02日 15时04分05秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Layout("It is 2006-01-02 15:04:05") // It is 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15").Layout("2006-01-02 15:04:05", carbon.Tokyo) // 2020-08-05 14:14:15

// 输出指定格式的字符串,Format()是ToFormatString()的简写(如果使用的字母与格式化字符冲突时，请使用\符号转义该字符)
carbon.Parse("2020-08-05 13:14:15").Format("YmdHis") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Format("Y年m月d日 H时i分s秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Format("l jS \\o\\f F Y h:i:s A") // Wednesday 5th of August 2020 01:14:15 PM
carbon.Parse("2020-08-05 13:14:15").Format("\\I\\t \\i\\s Y-m-d H:i:s") // It is 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15").Format("Y-m-d H:i:s", carbon.Tokyo) // 2020-08-05 14:14:15
```

> 更多格式化输出符号请查看附录 <a href="#format-sign-table">格式化符号表</a>

##### 星座

```go
// 获取星座
carbon.Parse("2020-08-05 13:14:15").Constellation() // Leo

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
carbon.Parse("2020-08-05 13:14:15").Season() // Summer

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

##### 农历

> 目前仅支持公元`1900`年至`2100`年的`200`年时间跨度

```go
// 获取农历生肖
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().Animal() // 鼠

// 获取农历节日
carbon.Parse("2021-02-12 13:14:15", carbon.PRC).Lunar().Festival() // 春节

// 获取农历年年份
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().Year() // 2020
// 获取农历月月份
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().Month() // 6
// 获取农历闰月月份
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().LeapMonth() // 4
// 获取农历日日期
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().Day() // 16
// 获取农历 YYYY-MM-DD 格式字符串
fmt.Sprintf("%s", carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar()) // 2020-06-16

// 获取农历年字符串
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().ToYearString() // 二零二零
// 获取农历月字符串
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().ToMonthString() // 六
// 获取农历天字符串
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().ToDayString() // 十六
// 获取农历日期字符串
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().ToDateString() // 二零二零年六月十六

// 是否是农历闰年
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsLeapYear() // true
// 是否是农历闰月
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsLeapMonth() // false

// 是否是鼠年
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsRatYear() // true
// 是否是牛年
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsOxYear() // false
// 是否是虎年
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsTigerYear() // false
// 是否是兔年
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsRabbitYear() // false
// 是否是龙年
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsDragonYear() // false
// 是否是蛇年
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsSnakeYear() // false
// 是否是马年
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsHorseYear() // false
// 是否是羊年
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsGoatYear() // false
// 是否是猴年
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsMonkeyYear() // false
// 是否是鸡年
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsRoosterYear() // false
// 是否是狗年
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsDogYear() // false
// 是否是猪年
carbon.Parse("2020-08-05 13:14:15", carbon.PRC).Lunar().IsPigYear() // false
```

##### JSON 支持

###### 定义模型

```go
type Person struct {
    ID  int64  `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
    Birthday carbon.DateTime `json:"birthday"`
    GraduatedAt carbon.Date `json:"graduated_at"`
    CreatedAt carbon.Time `json:"created_at"`
    UpdatedAt carbon.Timestamp `json:"updated_at"`
    DateTime1 carbon.TimestampWithSecond `json:"date_time1"`
    DateTime2 carbon.TimestampWithMillisecond `json:"date_time2"`
    DateTime3 carbon.TimestampWithMicrosecond `json:"date_time3"`
    DateTime4 carbon.TimestampWithNanosecond `json:"date_time4"`
}
```

###### 实例化模型

```go
person := Person {
    ID:          1,
    Name:        "gouguoyin",
    Age:         18,
    Birthday:    carbon.DateTime{carbon.Now().SubYears(18)},
    GraduatedAt: carbon.Date{carbon.Parse("2020-08-05 13:14:15")},
    CreatedAt:   carbon.Time{carbon.Parse("2021-08-05 13:14:15")},
    UpdatedAt:   carbon.Timestamp{carbon.Parse("2022-08-05 13:14:15")},
    DateTime1:   carbon.TimestampWithSecond{carbon.Parse("2023-08-05 13:14:15")},
    DateTime2:   carbon.TimestampWithMillisecond{carbon.Parse("2024-08-05 13:14:15")},
    DateTime3:   carbon.TimestampWithMicrosecond{carbon.Parse("2025-08-05 13:14:15")},
    DateTime4:   carbon.TimestampWithNanosecond{carbon.Parse("2025-08-05 13:14:15")},
}
```

###### JSON 编码

```go
data, err := json.Marshal(&person)
if err != nil {
    // 错误处理
    log.Fatal(err)
}
fmt.Printf("%s", data)
// 输出
{
    "id": 1,
    "name": "gouguoyin",
    "age": 18,
    "birthday": "2003-07-16 16:22:02",
    "graduated_at": "2020-08-05",
    "created_at": "13:14:15",
    "updated_at": 1659676455,
    "date_time1": 1691212455,
    "date_time2": 1722834855000,
    "date_time3": 1754370855000000,
    "date_time4": 1754370855000000000
}
```

###### JSON 解码

```go
jsonString := `{
	"id": 1,
	"name": "gouguoyin",
	"age": 18,
	"birthday": "2003-07-16 16:22:02",
	"graduated_at": "2020-08-05",
	"updated_at": 1659676455,
	"date_time1": 1691212455,
	"date_time2": 1722834855000,
	"date_time3": 1754370855000000,
	"date_time4": 1754370855000000000
}`
person := new(Person)
err := json.Unmarshal([]byte(jsonString), &person)
if err != nil {
    // 错误处理
    log.Fatal(err)
}
fmt.Printf("%+v", *person)
// 输出
{ID:1 Name:gouguoyin Age:18 Birthday:2003-07-16 16:22:02 GraduatedAt:2020-08-05 00:00:00 UpdatedAt:2022-08-05 13:14:15 DateTime1:2023-08-05 13:14:15 DateTime2:2024-08-05 13:14:15 DateTime3:2025-08-05 13:14:15 DateTime4:2025-08-05 13:14:15}
```

##### 国际化支持

目前支持的语言有

* [简体中文(zh-CN)](./lang/zh-CN.json "简体中文")
* [繁体中文(zh-TW)](./lang/zh-TW.json "繁体中文")
* [英语(en)](./lang/en.json "英语")
* [日语(jp)](./lang/jp.json "日语")
* [韩语(kr)](./lang/kr.json "韩语")
* [西班牙语(es)](./lang/es.json "西班牙语"): 由 [hgisinger](https://github.com/hgisinger "hgisinger") 翻译
* [德语(de)](./lang/de.json "德语"): 由 [benzammour](https://github.com/benzammour "benzammour") 翻译
* [土耳其语(tr)](./lang/tr.json "土耳其语"): 由 [emresenyuva](https://github.com/emresenyuva "emresenyuva") 翻译
* [葡萄牙语(pt)](./lang/pt.json "葡萄牙语"): 由 [felipear89](https://github.com/felipear89 "felipear89") 翻译
* [俄罗斯语(ru)](./lang/ru.json "俄罗斯语"): 由 [zemlyak](https://github.com/zemlyak "zemlyak") 翻译

目前支持的方法有

* `Constellation()`：获取星座
* `Season()`：获取季节
* `DiffForHumans()`：获取对人类友好的可读格式时间差
* `ToMonthString()`：输出完整月份字符串
* `ToShortMonthString()`：输出缩写月份字符串
* `ToWeekString()`：输出完整星期字符串
* `ToShortWeekString()`：输出缩写星期字符串

###### 设置区域

```go
lang := NewLanguage()
if err := lang.SetLocale("zh-CN"); err != nil {
    // 错误处理
    log.Fatal(err)
}

c := carbon.SetLanguage(lang)
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
lang := NewLanguage()

if err := lang.SetLocale("en"); err != nil {
    // 错误处理
    log.Fatal(err)
}

resources := map[string]string {
    "hour": "%dh",
}
lang.SetResources(resources)

c := carbon.SetLanguage(lang)
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
lang := NewLanguage()
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

> 如果有多个错误发生，只返回第一个错误，前一个错误排除后才返回下一个错误

###### 场景一

```go
c := carbon.SetTimezone(PRC).Parse("xxx")
if c.Error != nil {
    // 错误处理...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// 输出
cannot parse "xxx" as carbon, please make sure the value is valid
```

###### 场景二

```go
c := carbon.SetTimezone("xxx").Parse("2020-08-05")
if c.Error != nil {
    // 错误处理...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// 输出
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezones
```

###### 场景三

```go
c := carbon.SetTimezone("xxx").Parse("12345678")
if c.Error != nil {
    // 错误处理...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// 输出
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezone
```

#### 附录

##### <a id="format-sign-table">格式化符号表</a>

| 符号 | 描述 |  长度 | 范围 | 示例 |
| :------------: | :------------: | :------------: | :------------: | :------------: |
| d | 月份中的第几天，有前导零 | 2 | 01-31 | 02 |
| D | 缩写单词表示的周几 | 3 | Mon-Sun | Mon |
| j | 月份中的第几天，没有前导零 | - |1-31 | 2 |
| S | 第几天的英文缩写后缀，一般和j配合使用 | 2 | st/nd/rd/th | th |
| l | 完整单词表示的周几 | - | Monday-Sunday | Monday |
| F | 完整单词表示的月份 | - | January-December | January |
| m | 数字表示的月份，有前导零 | 2 | 01-12 | 01 |
| M | 缩写单词表示的月份 | 3 | Jan-Dec | Jan |
| n | 数字表示的月份，没有前导零 | - | 1-12 | 1 |
| Y | 4 位数字完整表示的年份 | 4 | 0000-9999 | 2006 |
| y | 2 位数字表示的年份 | 2 | 00-99 | 06 |
| a | 小写的上午和下午标识 | 2 | am/pm | pm |
| A | 大写的上午和下午标识 | 2 | AM/PM | PM |
| g | 小时，12 小时格式 | - | 1-12 | 3 |
| G | 小时，24 小时格式 | - | 0-23 | 15 |
| h | 小时，12 小时格式 | 2 | 00-11 | 03 |
| H | 小时，24 小时格式 | 2 | 00-23 | 15 |
| i | 分钟 | 2 | 01-59 | 04 |
| s | 秒数 | 2 | 01-59 | 05 |
| c | ISO8601 格式日期 | - | - | 2006-01-02T15:04:05-07:00 |
| r | RFC2822 格式日期 | - | - | Mon, 02 Jan 2006 15:04:05 -0700 |
| O | 与格林威治时间相差的小时数 | - | - | -0700 |
| P | 与格林威治时间相差的小时数，小时和分钟之间有冒号分隔 | - | - | +07:00 |
| T | 时区缩写 | - | - | MST |
| W | ISO8601 格式数字表示的年份中的第几周 | - | 1-52 | 1 |
| N | ISO8601 格式数字表示的星期中的第几天 | 1 | 1-7 | 1 |
| L | 是否为闰年，如果是闰年为 1，否则为 0 | 1 | 0-1 | 0 |
| U | 秒级时间戳 | 10 | - | 1611818268 |
| u | 毫秒 | 3 | 000-999 | 999 |
| w | 数字表示的周几 | 1 | 0-6 | 1 |
| t | 月份中的总天数 | 2 | 28-31 | 31 |
| z | 年份中的第几天 | - | 0-365 | 2 |
| e | 当前位置 | - | - | America/New_York |
| Q | 当前季节 | 1 | 1-4 | 1 |
| C | 当前世纪数 | - | 0-99 | 21 |

#### 参考项目

* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)
* [jinzhu/now](https://github.com/jinzhu/now)
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)
* [arrow-py/arrow](https://github.com/arrow-py/arrow)
* [moment/moment](https://github.com/moment/moment)
* [iamkun/dayjs](https://github.com/iamkun/dayjs)
