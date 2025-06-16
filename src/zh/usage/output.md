---
head:
  - - meta
    - name: description
      content: 时间输出|一个轻量级、语义化、对开发者友好的 golang 时间处理库
---

# 时间输出

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

更多格式输出符号请查看 <a href="/zh/appendix/format-signs" target="_blank" rel="noreferrer">格式化符号</a>
