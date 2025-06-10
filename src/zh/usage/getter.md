# 时间获取

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

// 获取秒精度时间戳
carbon.Parse("2020-08-05 13:14:15").Timestamp() // 1596633255
// 获取毫秒精度时间戳
carbon.Parse("2020-08-05 13:14:15.999").TimestampMilli() // 1596633255999
// 获取微秒精度时间戳
carbon.Parse("2020-08-05 13:14:15.999999").TimestampMicro() // 1596633255999999
// 获取纳秒精度时间戳
carbon.Parse("2020-08-05 13:14:15.999999999").TimestampNano() // 1596633255999999999

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

// 获取一周的结束日期
carbon.SetWeekStartsAt(carbon.Sunday).WeekEndsAt() // Saturday
carbon.SetWeekStartsAt(carbon.Monday).WeekEndsAt() // Sunday

// 获取当前布局模板
carbon.Parse("now").CurrentLayout() // "2006-01-02 15:04:05"
carbon.ParseByLayout("2020-08-05", DateLayout).CurrentLayout() // "2006-01-02"

// 获取年龄
carbon.Parse("2002-01-01 13:14:15").Age() // 17
carbon.Parse("2002-12-31 13:14:15").Age() // 18
```