# Carbon  #
中文 | [英文](./README.en.md)

carbon 是一个轻量级、语义化、对开发者友好的Golang时间处理库，支持链式调用和gorm、xorm等主流orm

如果您觉得不错，请给个star吧

github:[github.com/golang-module/carbon](https://github.com/golang-module/carbon "github.com/golang-module/carbon")

gitee:[gitee.com/go-package/carbon](https://gitee.com/go-package/carbon "gitee.com/go-package/carbon")

#### 安装使用
```go
// 使用github库
go get -u github.com/golang-module/carbon

import (
    "github.com/golang-module/carbon"
)

// 使用gitee库
go get -u gitee.com/go-package/carbon

import (
    "gitee.com/go-package/carbon"
)
```

#### 用法示例
> 默认时区为Local，即服务器所在时区，假设当前时间为2020-08-05 13:14:15

##### 昨天、今天、明天
```go
// 今天
carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
// 今天日期
carbon.Now().ToDateString() // 2020-08-05
// 今天时间
carbon.Now().ToTimeString() // 13:14:15
// 今天时间戳
carbon.Now().ToTimestamp() // 1596604455

// 昨天
carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// 昨天日期
carbon.Yesterday().ToDateString() // 2020-08-04
// 昨天时间
carbon.Yesterday().ToTimeString() // 13:14:15
// 昨天时间戳
carbon.Yesterday().ToTimestamp() // 1596518055

// 明天
carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// 明天日期
carbon.Tomorrow().ToDateString() // 2020-08-06
// 明天时间
carbon.Tomorrow().ToTimeString() // 13:14:15
// 明天时间戳
carbon.Tomorrow().ToTimestamp() // 1596690855
```

##### 创建Carbon实例
```go
// 从秒级时间戳创建Carbon实例
carbon.CreateFromTimestamp(1596604455).ToDateTimeString() // 2020-08-05 13:14:15
// 从毫秒级时间戳创建Carbon实例
carbon.CreateFromTimestamp(1596604455000).ToDateTimeString() // 2020-08-05 13:14:15
// 从微秒级时间戳创建Carbon实例
carbon.CreateFromTimestamp(1596604455000000).ToDateTimeString() // 2020-08-05 13:14:15
// 从纳级时间戳创建Carbon实例
carbon.CreateFromTimestamp(1596604455000000000).ToDateTimeString() // 2020-08-05 13:14:15

// 从年月日时分秒创建Carbon实例
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
// 从年月日创建Carbon实例(时分秒默认为当前时分秒)
carbon.CreateFromDate(2020, 8, 5).ToDateTimeString() // 2020-08-05 13:14:15
// 从时分秒创建Carbon实例(年月日默认为当前年月日)
carbon.CreateFromTime(13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
// 从原生time.Time创建Carbon实例
carbon.CreateFromGoTime(time.Now()).ToTimestamp() // 1596604455
```

##### 解析标准格式时间字符串
```go
carbon.Parse("").ToDateTimeString() // 空字符串
carbon.Parse("0000-00-00 00:00:00").ToDateTimeString() // 空字符串
carbon.Parse("0000-00-00").ToDateTimeString() // 空字符串
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05").ToDateTimeString() // 2020-08-05 00:00:00
carbon.Parse("20200805131415").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("20200805").ToDateTimeString() // 2020-08-05 00:00:00
carbon.Parse("2020-08-05T13:14:15+08:00").ToDateTimeString() // 2020-08-05 00:00:00
```

##### 解析自定义格式时间字符串
```go
carbon.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString // 2020-08-05 13:14:15
carbon.ParseByFormat("2020%08%05% 13%14%15", "Y%m%d% h%i%s").ToDateTimeString // 2020-08-05 13:14:15
carbon.ParseByFormat("2020年08月05日 13时14分15秒", "Y年m月d日 H时i分s秒").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("2020年08月05日", "Y年m月d日").ToDateTimeString() // 2020-08-05 00:00:00
carbon.ParseByFormat("13时14分15秒", "H时i分s秒").ToDateTimeString() // 2020-08-05 13:14:15
```

##### 解析持续时间字符串(基于当前时间)

支持正负整数/浮点数和符号ns(纳秒)、us(微妙)、ms(毫秒)、s(秒)、m(分钟)、h(小时)的组合

```go
// 十小时后
carbon.ParseByDuration("10h").ToDateTimeString() // 2020-08-06 23:14:15
// 十小时半前
carbon.ParseByDuration("-10.5h").ToDateTimeString() // 2020-08-05 02:44:15
// 十分钟后
carbon.ParseByDuration("10m").ToDateTimeString() // 2020-08-05 13:24:15
// 十分钟半前
carbon.ParseByDuration("-10.5m").ToDateTimeString() // 2020-08-05 13:03:45
// 十秒后
carbon.ParseByDuration("10s").ToDateTimeString() // 2020-08-05 13:14:25
// 十秒半前
carbon.ParseByDuration("-10.5s").ToDateTimeString() // 2020-08-05 13:14:04
```

##### 时间设置
```go
// 设置时区
carbon.SetTimezone(carbon.PRC).Now().ToDateTimeString() // 2020-08-05 13:14:15
carbon.SetTimezone(carbon.Tokyo).Now().ToDateTimeString() // 2020-08-05 14:14:15
carbon.SetTimezone(carbon.Tokyo).SetTimezone(carbon.PRC).Now().ToDateTimeString() // 2020-08-05 13:14:15

// 设置年
carbon.Parse("2019-08-05").SetYear(2020).ToDateString() // 2020-08-05
carbon.Parse("2020-02-29").SetYear(2019).ToDateString() // 2019-03-01

// 设置月
carbon.Parse("2020-01-30").SetMonth(2).ToDateString() // 2020-03-01
carbon.Parse("2020-01-31").SetMonth(2).ToDateString() // 2020-03-02
carbon.Parse("2020-08-05").SetMonth(2).ToDateString() // 2020-02-05

// 设置日
carbon.Parse("2019-08-05").SetDay(31).ToDateString() // 2020-08-31
carbon.Parse("2020-02-01").SetDay(31).ToDateString() // 2020-03-02

// 设置时
carbon.Parse("2020-08-05 13:14:15").SetHour(10).ToDateTimeString() // 2020-08-05 10:14:15
carbon.Parse("2020-08-05 13:14:15").SetHour(24).ToDateTimeString() // 2020-08-06 00:14:15

// 设置分
carbon.Parse("2020-08-05 13:14:15").SetMinute(10).ToDateTimeString() // 2020-08-05 13:10:15
carbon.Parse("2020-08-05 13:14:15").SetMinute(60).ToDateTimeString() // 2020-08-05 14:00:15

// 设置秒
carbon.Parse("2020-08-05 13:14:15").SetSecond(10).ToDateTimeString() // 2020-08-05 13:14:10
carbon.Parse("2020-08-05 13:14:15").SetSecond(60).ToDateTimeString() // 2020-08-05 13:15:00
```
>更多时区常量请查看[const.go](./const.go)文件

##### 开始时间、结束时间
```go
// 本年开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfYear().ToDateTimeString() // 2020-01-01 00:00:00
// 本年结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfYear().ToDateTimeString() // 2020-12-31 23:59:59

// 本月开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfMonth().ToDateTimeString() // 2020-08-01 00:00:00
// 本月结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfMonth().ToDateTimeString() // 2020-08-31 23:59:59

// 本周开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfWeek().ToDateTimeString() // 2020-08-03 00:00:00
// 本周结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfWeek().ToDateTimeString() // 2020-08-09 23:59:59

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
```

##### 时间旅行
```go
// 三年后
carbon.Parse("2020-02-29 13:14:15").AddYears(3).ToDateTimeString() // 2023-03-01 13:14:15
carbon.Parse("2020-02-29 13:14:15").NextYears(3).ToDateTimeString() // 2023-02-28 13:14:15

// 一年后
carbon.Parse("2020-02-29 13:14:15").AddYear().ToDateTimeString() // 2021-03-01 13:14:15
carbon.Parse("2020-02-29 13:14:15").NextYear().ToDateTimeString() // 2021-02-28 13:14:15

// 三年前
carbon.Parse("2020-02-29 13:14:15").SubYears(3).ToDateTimeString() // 2017-03-01 13:14:15
carbon.Parse("2020-02-29 13:14:15").PreYears(3).ToDateTimeString() // 2017-02-28 13:14:15

// 一年前
carbon.Parse("2020-02-29 13:14:15").SubYear().ToDateTimeString() // 2019-03-01 13:14:15
carbon.Parse("2020-02-29 13:14:15").PreYear().ToDateTimeString() // 2019-02-28 13:14:15

// 三季度后
carbon.Parse("2019-08-31 13:14:15").AddQuarters(3).ToDateTimeString() // 2019-03-02 13:14:15
carbon.Parse("2019-08-31 13:14:15").NextQuarters(3).ToDateTimeString() // 2019-02-29 13:14:15

// 一季度后
carbon.Parse("2019-11-30 13:14:15").AddQuarter().ToDateTimeString() // 2020-03-01 13:14:15
carbon.Parse("2019-11-30 13:14:15").NextQuarter().ToDateTimeString() // 2020-02-29 13:14:15

// 三季度前
carbon.Parse("2019-08-31 13:14:15").SubQuarters(3).ToDateTimeString() // 2019-03-03 13:14:15
carbon.Parse("2019-08-31 13:14:15").PreQuarters(3).ToDateTimeString() // 2019-02-28 13:14:15

// 一季度前
carbon.Parse("2020-05-31 13:14:15").SubQuarter().ToDateTimeString() // 2020-03-02 13:14:15
carbon.Parse("2020-05-31 13:14:15").PreQuarter().ToDateTimeString() // 2020-02-29 13:14:15

// 三月后
carbon.Parse("2020-02-29 13:14:15").AddMonths(3).ToDateTimeString() // 2020-05-29 13:14:15
carbon.Parse("2020-02-29 13:14:15").NextMonths(3).ToDateTimeString() // 2020-05-29 13:14:15

// 一月后
carbon.Parse("2020-01-31 13:14:15").AddMonth().ToDateTimeString() // 2020-03-02 13:14:15
carbon.Parse("2020-01-31 13:14:15").NextMonth().ToDateTimeString() // 2020-02-29 13:14:15

// 三月前
carbon.Parse("2020-02-29 13:14:15").SubMonths(3).ToDateTimeString() // 2019-11-29 13:14:15
carbon.Parse("2020-02-29 13:14:15").PreMonths(3).ToDateTimeString() // 2019-11-29 13:14:15

// 一月前
carbon.Parse("2020-03-31 13:14:15").SubMonth().ToDateTimeString() // 2020-03-02 13:14:15
carbon.Parse("2020-03-31 13:14:15").PreMonth().ToDateTimeString() // 2020-02-29 13:14:15

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
// 一小时后
carbon.Parse("2020-08-05 13:14:15").AddHour().ToDateTimeString() // 2020-08-05 14:14:15

// 三小时前
carbon.Parse("2020-08-05 13:14:15").SubHours(3).ToDateTimeString() // 2020-08-05 10:14:15
// 二小时半前
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5h").ToDateTimeString() // 2020-08-05 10:44:15
// 一小时前
carbon.Parse("2020-08-05 13:14:15").SubHour().ToDateTimeString() // 2020-08-05 12:14:15

// 三分钟后
carbon.Parse("2020-08-05 13:14:15").AddMinutes(3).ToDateTimeString() // 2020-08-05 13:17:15
// 二分钟半后
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5m").ToDateTimeString() // 2020-08-05 13:16:45
// 一分钟后
carbon.Parse("2020-08-05 13:14:15").AddMinute().ToDateTimeString() // 2020-08-05 13:15:15

// 三分钟前
carbon.Parse("2020-08-05 13:14:15").SubMinutes(3).ToDateTimeString() // 2020-08-05 13:11:15
// 二分钟半前
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5m").ToDateTimeString() // 2020-08-05 13:11:45
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
```

##### 时间输出
```go
// 输出秒级时间戳
carbon.Parse("2020-08-05 13:14:15").ToTimestamp() // 1596604455
carbon.Parse("2020-08-05 13:14:15").ToTimestampWithSecond() // 1596604455
// 输出毫秒级时间戳
carbon.Parse("2020-08-05 13:14:15").ToTimestampWithMillisecond() // 1596604455000
// 输出微秒级时间戳
carbon.Parse("2020-08-05 13:14:15").ToTimestampWithMicrosecond() // 1596604455000000
// 输出纳秒级时间戳
carbon.Parse("2020-08-05 13:14:15").ToTimestampWithNanosecond() // 1596604455000000000

// 输出字符串
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
// 输出格式化字符串
carbon.Parse("2020-08-05 13:14:15").Format("YmdHis") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").ToFormatString("YmdHis") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").ToFormatString("Y年m月d H时i分s秒") // 2020年08月05日 13时14分15秒
// 输出日期时间字符串
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
// 输出日期字符串
carbon.Parse("2020-08-05 13:14:15").ToDateString() // 2020-08-05
// 输出时间字符串
carbon.Parse("2020-08-05 13:14:15").ToTimeString() // 13:14:15

// 输出Ansic格式字符串
carbon.Parse("2020-08-05 13:14:15").ToAnsicString() // Wed Aug  5 13:14:15 2020
// 输出Atom字符串
carbon.Parse("2020-08-05 13:14:15").ToAtomString() // Wed Aug  5 13:14:15 2020
// 输出UnixDate格式字符串
carbon.Parse("2020-08-05 13:14:15").ToUnixDateString() // Wed Aug  5 13:14:15 CST 2020
// 输出RubyDate格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRubyDateString() // Wed Aug 05 13:14:15 +0800 2020
// 输出Kitchen格式字符串
carbon.Parse("2020-08-05 13:14:15").ToKitchenString() // 1:14PM
// 输出Cookie格式字符串
carbon.Parse("2020-08-05 13:14:15").ToCookieString() // Wednesday, 05-Aug-2020 13:14:15 CST
// 输出DayDateTime格式字符串
carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString() // Wed, Aug 5, 2020 1:14 PM
// 输出RSS格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRssString() // Wed, 05 Aug 2020 13:14:15 +0800
// 输出W3C格式字符串
carbon.Parse("2020-08-05 13:14:15").ToW3cString() // 2020-08-05T13:14:15+08:00

// 输出RFC822格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRFC822String() // 05 Aug 20 13:14 CST
// 输出RFC822Z格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRFC822zString() // 05 Aug 20 13:14 +0800
// 输出RFC850格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRFC850String() // Wednesday, 05-Aug-20 13:14:15 CST
// 输出RFC1036格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRFC1036String() // Wed, 05 Aug 20 13:14:15 +0800
// 输出RFC1123格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRFC1123String() // Wed, 05 Aug 2020 13:14:15 CST
// 输出RFC2822格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRFC2822String() // Wed, 05 Aug 2020 13:14:15 +0800
// 输出RFC3339格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRFC3339String() // 2020-08-05T13:14:15+08:00
// 输出RFC7231格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRFC7231String() // Wed, 05 Aug 2020 05:14:15 GMT

```
>更多格式化输出符号请查看附录 <a href="#格式化符号表">格式化符号表</a>

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
carbon.Parse("2020-08-05 13:14:15").WeekOfYear() // 32
// 获取本月第几天
carbon.Parse("2020-08-05 13:14:15").DayOfMonth() // 5
// 获取本月第几周
carbon.Parse("2020-08-05 13:14:15").WeekOfMonth() // 1
// 获取本周第几天
carbon.Parse("2020-08-05 13:14:15").DayOfWeek() // 3

// 获取当前年
carbon.Parse("2020-08-05 13:14:15").Year() // 2020
// 获取当前季度
carbon.Parse("2020-08-05 13:14:15").Quarter() // 3
// 获取当前月
carbon.Parse("2020-08-05 13:14:15").Month() // 8
// 获取当前日
carbon.Parse("2020-08-05 13:14:15").Day() // 5
// 获取当前时
carbon.Parse("2020-08-05 13:14:15").Hour() // 13
// 获取当前分
carbon.Parse("2020-08-05 13:14:15").Minute() // 14
// 获取当前秒
carbon.Parse("2020-08-05 13:14:15").Second() // 15

// 获取时区
carbon.SetTimezone(carbon.PRC).Timezone() // PRC
carbon.SetTimezone(carbon.Tokyo).Timezone() // Asia/Tokyo

// 获取年龄
carbon.Parse("2002-01-01 13:14:15").Age() // 17
carbon.Parse("2002-12-31 13:14:15").Age() // 18

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

// 是否是当前时间
carbon.Parse(carbon.Now().ToDateTimeString()).IsNow() // true
// 是否是未来时间
carbon.Parse("2020-08-06 13:14:15").IsFuture() // true
// 是否是过去时间
carbon.Parse("2020-08-04 13:14:15").IsPast() // true

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
carbon.Parse("2020-08-05 13:14:15").IsThursday()  // false
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
```

##### 农历支持
```go
// 获取生肖年
carbon.Parse("2020-08-05 13:14:15").ToAnimalYear() // 鼠
// 获取农历年
carbon.Parse("2020-08-05 13:14:15").ToLunarYear() // 庚子

// 是否是鼠年
carbon.Parse("2020-08-05 13:14:15").IsYearOfRat() // true
// 是否是牛年
carbon.Parse("2020-08-05 13:14:15").IsYearOfOx() // false
// 是否是虎年
carbon.Parse("2020-08-05 13:14:15").IsYearOfTiger() // false
// 是否是兔年
carbon.Parse("2020-08-05 13:14:15").IsYearOfRabbit() // false
// 是否是龙年
carbon.Parse("2020-08-05 13:14:15").IsYearOfDragon() // false
// 是否是蛇年
carbon.Parse("2020-08-05 13:14:15").IsYearOfSnake() // false
// 是否是马年
carbon.Parse("2020-08-05 13:14:15").IsYearOfHorse() // false
// 是否是羊年
carbon.Parse("2020-08-05 13:14:15").IsYearOfGoat() // false
// 是否是猴年
carbon.Parse("2020-08-05 13:14:15").IsYearOfMonkey() // false
// 是否是鸡年
carbon.Parse("2020-08-05 13:14:15").IsYearOfRooster() // false
// 是否是狗年
carbon.Parse("2020-08-05 13:14:15").IsYearOfDog() // false
// 是否是猪年
carbon.Parse("2020-08-05 13:14:15").IsYearOfPig() // false
```

##### 数据库支持
假设数据表为users，字段有id(int)、name(varchar)、age(int)、graduated_at(date)、birthday(date)、created_at(datetime)、updated_at(datetime)、deleted_at(datetime)

###### 定义模型
```go
type UserModel struct {
    ID  int64  `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
    Birthday carbon.Carbon `json:"birthday"`
    CreatedAt carbon.ToDateTimeString `json:"created_at"`
    DeletedAt carbon.ToTimestamp `json:"deleted_at"`
    GraduatedAt carbon.ToDateString `json:"graduated_at"`
    UpdatedAt carbon.ToTimeString `json:"updated_at"`
}
```

###### 实例化模型
```go
user := UserModel {
    Name: "勾国印",
    Age: 18,
    Birthday: carbon.Now().SubYears(18),
    CreatedAt: carbon.ToDateTimeString{carbon.Now()},
    DeletedAt: carbon.ToTimestamp{carbon.Parse("2020-08-05 13:14:15")},
    GraduatedAt: carbon.ToDateString{carbon.Parse("2012-09-09")},
    UpdatedAt: carbon.ToTimeString{carbon.Now()},
}
```

###### 输出模型字段
```go
user.ID // 18
user.Name // 勾国印
user.Birthday.ToDateString() // 2012-08-05
user.CreatedAt.ToDateTimeString() // 2012-08-05 13:14:15
user.DeletedAt.ToTimestamp() // 1596604455
user.GraduatedAt.AddDay().ToDateString() // 2012-09-10
user.UpdatedAt.ToTimeString() // 13:14:15
```

###### JSON输出模型
```go
data, _ := json.Marshal(&user)
fmt.Print(string(data))
// 输出
{
    "id": 42,
    "name": "勾国印",
    "age": 18,
    "birthday": "2012-08-05",
    "created_at": "2020-08-05 13:14:15",
    "deleted_at": 1596604455
    "graduated_at": "2012-09-09",
    "updated_at": "13:14:15",
}
```

###### 输出自定义格式
```go
// 定义输出格式
type ToRssString struct {
    carbon.Carbon
}

// 定义模型
type UserModel struct {
    Birthday ToRssString `json:"birthday"`
}

// 实例化模型
user := UserModel {
    Birthday: ToRssString{carbon.Now()},
}

// 重写MarshalJSON方法
func (c ToRssString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, c.ToRssString())), nil
}

// json.Marshal(&user)输出
{
    "birthday": "Wed, 05 Aug 2020 13:14:15 +0800",
}
```

#### 附录
##### <a id="格式化符号表">格式化符号表</a>
| 符号 | 描述 | 类型 | 长度 | 范围 | 示例 |
| :------------: | :------------: | :------------: | :------------: | :------------: | :------------: |
| Y | 年份 | 数字 | 4 | 0000-9999 | 2020 |
| y | 年份 | 数字 | 2 | 00-99 | 20 |
| M | 月份 | 字母 | 3 | Jan-Dec | Aug |
| m | 月份 | 数字 | 2 | 01-12 | 08 |
| F | 月份 | 字母 | - | January-December | August |
| n | 月份 | 数字 | 1/2 | 1-12 | 8 |
| l | 周几 | 字母 | - | Monday-Sunday | Wednesday |
| D | 周几 | 字母 | 3 | Mon-Sun | Wed |
| d | 天数 | 数字 | 2 | 01-31 | 05 |
| j | 天数 | 数字 | 1/2 |1-31 | 5 |
| H | 小时 | 数字 | 2 | 00-23 | 15 |
| h | 小时 | 数字 | 2 | 00-11 | 03 |
| i | 分钟 | 数字 | 2 | 01-59 | 14 |
| s | 秒钟 | 数字 | 2 | 01-59 | 15 |
| P | 上下午 | 字母 | 2 | AM/PM | PM |
| p | 上下午 | 字母 | 2 | am/pm | pm |

#### 参考项目
* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)
* [uniplaces/carbon](https://github.com/uniplaces/carbon)
* [jinzhu/now](https://github.com/jinzhu/now)
* [araddon/dateparse](https://github.com/araddon/dateparse)

#### 更新日志
##### 2020-11-02
* 新增测试覆盖率报告文件coverage.html
* CreateFromTimestamp()方法支持秒、毫秒、微秒、纳秒级时间戳
* 新增ToTimestampWithSecond()方法获取秒级时间戳，等价于ToTimestamp()
* 新增ToTimestampWithMillisecond()方法获取毫秒级时间戳
* 新增ToTimestampWithMicrosecond()方法获取微秒级时间戳
* 新增ToTimestampWithNanosecond()方法获取微秒级时间戳
##### 2020-10-22
* 新增SetYear()方法设置年
* 新增SetMonth()方法设置月
* 新增SetDay()方法设置日
* 新增SetHour()方法设置时
* 新增SetMinute()方法设置分
* 新增SetSecond方法设置秒
* 新增DiffInWeeks()方法计算相差多少周
* 新增DiffAbsInWeeks()方法计算相差多少周(绝对值)
* 新增DiffInDays()方法计算相差多少天
* 新增DiffAbsInDays()方法计算相差多少天(绝对值)
* 新增DiffInHours()方法计算相差多少小时
* 新增DiffAbsInHours()方法计算相差多少小时(绝对值)
* 新增DiffInMinutes()方法计算相差多少分钟
* 新增DiffAbsInMinutes()方法计算相差多少分钟(绝对值)
* 新增DiffInSeconds()方法计算相差多少秒
* 新增DiffAbsInSeconds()方法计算相差多少秒(绝对值)

##### 2020-10-16
* 新增Timezone()方法获取时区名
* 新增Age()方法获取年龄
* 新增Year()方法获取当前年
* 新增Month()方法获取当前月
* 新增Day()方法获取当前日
* 新增Hour()方法获取当前小时
* 新增Minute()方法获取当前分钟数
* 新增Second()方法获取当前秒数

##### 2020-10-12
* 完善单元测试，代码覆盖率100%
* 统一异常处理
* 新增英文版README.MD说明文档
* 统一输出函数命名规则，将Format() 改为 ToFormatString()
* 新增CreateFromGoTime(t time.Time)方法从原生time.Time创建Carbon实例

##### 2020-10-08
* 完善单元测试
* 完善优化对ORM的多场景支持
* 优化代码组织结构，将不可继承的最终方法统一放到final.go文件里
* 废弃New()初始化函数，无需初始化即可直接使用
* 新增多种时间格式输出，如Cookie、W3C、RSS、RFC7231
* 新增ParseByDuration()方法解析持续时间字符串(相对于今天)，支持正负整数/浮点数和符号ns(纳秒)、us(微妙)、ms(毫秒)、s(秒)、m(分钟)、h(小时)的组合
* 新增NextYears()、NextYear()、PreYears()、PreYear()方法防止出现添加/减少指定年时出现跨月的现象
* 新增NextMonths()、NextMonth()、PreMonths()、PreMonth()方法防止出现添加/减少指定月后出现跨月的现象
* 新增DaysInYear()方法获取本年的总天数
* 新增DaysInMonth()方法获取本月的总天数
* 新增MonthOfYear()方法获取本年的第几月
* 新增DayOfYear()方法获取本年的第几天
* 新增DayOfMonth()方法获取本月的第几天
* 新增DayOfWeek()方法获取本周的第几天
* 新增WeekOfYear()方法获取本年的第几周
* 新增WeekOfMonth()方法获取本月的第几周
* 新增IsZero()方法判断是否是零值时间
* 新增IsFuture()方法判断是否是未来时间
* 新增IsPast()方法判断是否是过去时间
* 新增IsNow()方法判断是否是当前时间

##### 2020-10-01
* 完善单元测试
* 修复 AddHours() 传入参数小于1天时变成浮点数的BUG
* 修复 AddHour() 浮点数BUG
* 修复 SubHours() 传入参数小于1天时变成浮点数的BUG
* 修复 SubHour() 浮点数BUG
* 修复 AddMinutes() 传入参数小于1天时变成浮点数的BUG
* 修复 AddMinute() 浮点数BUG
* 修复 SubMinutes() 传入参数小于1天时变成浮点数的BUG
* 修复 SubMinute() 浮点数BUG
* 修复 AddSeconds() 传入参数小于1天时变成浮点数的BUG
* 修复 AddSecond() 浮点数BUG
* 修复 SubSeconds() 传入参数小于1天时变成浮点数的BUG
* 修复 SubSecond() 浮点数BUG
* 修复orm中时间字段类型设置为carbon.ToDateTimeString时报错的BUG
* 改名解析自定义时间格式方法ParseByCustom() 为 ParseByFormat()
* 新增 ParseByDuration() 方法将持续时间字符串转化成时间实例
* 新增 ToAnimalYear() 方法获取生肖年
* 新增 ToLunarYear() 方法获取农历年
* 新增 IsYearOfRat() 方法判断是否是鼠年
* 新增 IsYearOfOx() 方法判断是否是牛年
* 新增 IsYearOfTiger() 方法判断是否是虎年
* 新增 IsYearOfRabbit() 方法判断是否是兔年
* 新增 IsYearOfDragon() 方法判断是否是龙年
* 新增 IsYearOfSnake() 方法判断是否是蛇年
* 新增 IsYearOfHorse() 方法判断是否是马年
* 新增 IsYearOfGoat() 方法判断是否是羊年
* 新增 IsYearOfMonkey() 方法判断是否是猴年
* 新增 IsYearOfRooster() 方法判断是否是鸡年
* 新增 IsYearOfDog() 方法判断是否是狗年
* 新增 IsYearOfPig() 方法判断是否是猪年

##### 2020-09-14
* 完善单元测试
* 时区常量移到const.go文件里
* 新增StartOfYear()方法获取当年开始时间
* 新增EndOfYear()方法获取当年结束时间
* 新增StartOfMonth()方法获取当月开始时间
* 新增EndOfMonth()方法获取当月结束时间
* 新增StartOfDay()方法获取当天开始时间
* 新增EndOfDay()方法获取当天结束时间
* 新增StartOfYesterday()方法获取昨天开始时间
* 新增EndOfYesterday()方法获取昨天结束时间
* 新增StartOfToday()方法获取今天开始时间
* 新增EndOfToday()方法获取今天结束时间
* 新增StartOfTomorrow()方法获取明天开始时间
* 新增EndOfTomorrow方法获取明天结束时间
* 新增ToDateStartString方法转换成日期开始时间
* 新增ToDateEndString方法转换成日期结束时间
* 新增ToTimeStartString方法转换成小时开始时间
* 新增ToTimeEndString方法转换成小时结束时间

##### 2020-09-10
* 优化代码组织结构，将私有方法统一放到private.go文件里
* 新增IsToday方法判断是否是今天
* 新增IsYesterday方法判断是否是昨天
* 新增IsTomorrow方法判断是否是明天
* 新增IsStartOfToday方法判断是否是今天开始时间
* 新增IsEndOfToday方法判断是否是今天结束时间
* 新增IsStartOfTomorrow方法判断是否是明天开始时间
* 新增IsEndOfTomorrow方法判断是否是明天结束时间
* 新增IsStartOfYesterday方法判断是否是昨天开始时间
* 新增IsEndOfYesterday方法判断是否是昨天结束时间

##### 2020-09-02
* 修复数据库中时间类型字段值为null或0000-00-00 00:00:00时，json格式化后为0001-01-01 00:00:00的BUG
* 完善单元测试
* 优化代码组织结构，精简代码
* 新增对xorm结构体的json输出时间格式化支持，支持输出多种标准时间格式
* 新增AddWeeks(N)方法获取N周后时间
* 新增AddWeek()方法获取1周后时间
* 新增SubWeeks(N)方法获取N周前时间
* 新增SubWeek()方法获取1周前时间

##### 2020-08-21
* 修复readme.md错误描述
* 完善单元测试
* 新增对gorm结构体的json输出时间格式化支持，支持输出多种标准时间格式
* 新增IsJanuary()方法判断是否是一月
* 新增IsFebruary()方法判断是否是二月
* 新增IsMarch()方法判断是否是三月
* 新增IsApril()方法判断是否是四月
* 新增IsMay()方法判断是否是五月
* 新增IsJune()方法判断是否是六月
* 新增IsJuly()方法判断是否是七月
* 新增IsAugust()方法判断是否是八月
* 新增IsSeptember()方法判断是否是九月
* 新增IsOctober()方法判断是否是十月
* 新增IsNovember()方法判断是否是十一月
* 新增IsDecember()方法判断是否是十二月
 
##### 2020-08-05
* 修复已知BUG
* 添加单元测试
* 新增FirstOfYear()方法获取年初第一天
* 新增LastOfYear()方法获取年末最后一天
* 新增FirstOfMonth()方法获取月初第一天
* 新增LastOfMonth()方法获取月末最后一天
* 新增IsFirstOfYaer()方法判断是否年初第一天
* 新增IsLastOfYear()方法判断是否年末最后一天
* 新增IsFirstOfMonth()方法判断是否月初第一天
* 新增IsLastOfMonth()方法判断是否月末最后一天
