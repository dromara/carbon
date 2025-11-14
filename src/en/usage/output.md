---
head:
  - - meta
    - name: description
      content: Carbon time output methods comprehensive guide, provides 40+ time formatting methods, supports multi-precision output (millisecond/microsecond/nanosecond) for datetime and short forms, covers ISO8601, RFC series, Ansic, Kitchen, Cookie and other standard formats, supports Layout and Format custom formatting
  - - meta
    - name: keywords
      content: carbon, go-carbon, time output, formatting, multi-precision output, ISO8601, RFC format, standard format, Layout, Format, custom format, time string
---

# Output

```go
// Output datetime format string
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeString() // 2020-08-05 13:14:15
// Output datetime with millisecond format string
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeMilliString() // 2020-08-05 13:14:15.999
// Output datetime with microsecond format string
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeMicroString() // 2020-08-05 13:14:15.999999
// Output datetime with nanosecond format string
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeNanoString() // 2020-08-05 13:14:15.999999999

// Output short datetime format string
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeString() // 20200805131415
// Output short datetime with millisecond format string
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeMilliString() // 20200805131415.999
// Output short datetime with microsecond format string
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeMicroString() // 20200805131415.999999
// Output short datetime with nanosecond format string
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeNanoString() // 20200805131415.999999999

// Output date format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateString() // 2020-08-05
// Output date with millisecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateMilliString() // 2020-08-05.999
// Output date with microsecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateMicroString() // 2020-08-05.999999
// Output date with nanosecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateNanoString() // 2020-08-05.999999999

// Output short date format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateString() // 20200805
// Output short date with millisecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateMilliString() // 20200805.999
// Output short date with microsecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateMicroString() // 20200805.999999
// Output short date with nanosecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateNanoString() // 20200805.999999999

// Output time format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeString() // 13:14:15
// Output time with millisecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeMilliString() // 13:14:15.999
// Output time with microsecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeMicroString() // 13:14:15.999999
// Output time with nanosecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeNanoString() // 13:14:15.999999999

// Output short time format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeString() // 131415
// Output short time with millisecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeMilliString() // 131415.999
// Output short time with microsecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeMicroString() // 131415.999999
// Output short time with nanosecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeNanoString() // 131415.999999999

// Output Ansic format string
carbon.Parse("2020-08-05 13:14:15").ToAnsicString() // Wed Aug  5 13:14:15 2020
// Output Atom format string
carbon.Parse("2020-08-05 13:14:15").ToAtomString() // 2020-08-05T13:14:15+00:00
// Output Unix date format string
carbon.Parse("2020-08-05 13:14:15").ToUnixDateString() // Wed Aug  5 13:14:15 UTC 2020
// Output Ruby date format string
carbon.Parse("2020-08-05 13:14:15").ToRubyDateString() // Wed Aug 05 13:14:15 +0000 2020
// Output Kitchen format string
carbon.Parse("2020-08-05 13:14:15").ToKitchenString() // 1:14PM
// Output Cookie format string
carbon.Parse("2020-08-05 13:14:15").ToCookieString() // Wednesday, 05-Aug-2020 13:14:15 UTC
// Output DayDateTime format string
carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString() // Wed, Aug 5, 2020 1:14 PM
// Output RSS format string
carbon.Parse("2020-08-05 13:14:15").ToRssString() // Wed, 05 Aug 2020 13:14:15 +0000
// Output W3C format string
carbon.Parse("2020-08-05 13:14:15").ToW3cString() // 2020-08-05T13:14:15+00:00

// Output ISO8601 format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601String() // 2020-08-05T13:14:15+00:00
// Output ISO8601 with millisecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601MilliString() // 2020-08-05T13:14:15.999+00:00
// Output ISO8601 with microsecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601MicroString() // 2020-08-05T13:14:15.999999+00:00
// Output ISO8601 with nanosecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601NanoString() // 2020-08-05T13:14:15.999999999+00:00
// Output ISO8601Zulu format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluString() // 2020-08-05T13:14:15Z
// Output ISO8601Zulu with millisecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluMilliString() // 2020-08-05T13:14:15.999Z
// Output ISO8601Zulu with microsecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluMicroString() // 2020-08-05T13:14:15.999999Z
// Output ISO8601Zulu with nanosecond format string
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluNanoString() // 2020-08-05T13:14:15.999999999Z

// Output RFC822 format string
carbon.Parse("2020-08-05 13:14:15").ToRfc822String() // 05 Aug 20 13:14 UTC
// Output RFC822Z format string
carbon.Parse("2020-08-05 13:14:15").ToRfc822zString() // 05 Aug 20 13:14 +0000
// Output RFC850 format string
carbon.Parse("2020-08-05 13:14:15").ToRfc850String() // Wednesday, 05-Aug-20 13:14:15 UTC
// Output RFC1036 format string
carbon.Parse("2020-08-05 13:14:15").ToRfc1036String() // Wed, 05 Aug 20 13:14:15 +0000
// Output RFC1123 format string
carbon.Parse("2020-08-05 13:14:15").ToRfc1123String() // Wed, 05 Aug 2020 13:14:15 UTC
// Output RFC1123Z format string
carbon.Parse("2020-08-05 13:14:15").ToRfc1123zString() // Wed, 05 Aug 2020 13:14:15 +0000
// Output RFC2822 format string
carbon.Parse("2020-08-05 13:14:15").ToRfc2822String() // Wed, 05 Aug 2020 13:14:15 +0000
// Output RFC7231 format string
carbon.Parse("2020-08-05 13:14:15").ToRfc7231String() // Wed, 05 Aug 2020 13:14:15 UTC

// Output RFC3339 format string
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339String() // 2020-08-05T13:14:15+00:00
// Output RFC3339 with millisecond format string
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339MilliString() // 2020-08-05T13:14:15.999+00:00
// Output RFC3339 with microsecond format string
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339MicroString() // 2020-08-05T13:14:15.999999+00:00
// Output RFC3339 with nanosecond format string
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339NanoString() // 2020-08-05T13:14:15.999999999+00:00

// Output datetime format string
fmt.Printf("%s", carbon.Parse("2020-08-05 13:14:15")) // 2020-08-05 13:14:15

// Output "2006-01-02 15:04:05.999999999 -0700 MST" format string
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0000 UTC

// Output "Jan 2, 2006" format string
carbon.Parse("2020-08-05 13:14:15").ToFormattedDateString() // Aug 5, 2020
// Output "Mon, Jan 2, 2006" format string
carbon.Parse("2020-08-05 13:14:15").ToFormattedDayDateString() // Wed, Aug 5, 2020

// Output string by layout
carbon.Parse("2020-08-05 13:14:15").Layout(carbon.ISO8601Layout) // 2020-08-05T13:14:15+00:00
carbon.Parse("2020-08-05 13:14:15").Layout("20060102150405") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Layout("2006年01月02日 15时04分05秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Layout("It is 2006-01-02 15:04:05") // It is 2020-08-05 13:14:15

// Output string by format
carbon.Parse("2020-08-05 13:14:15").Format("YmdHis") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Format("Y年m月d日 H时i分s秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Format("l jK \\o\\f F Y h:i:s A") // Wednesday 5th of August 2020 01:14:15 PM
carbon.Parse("2020-08-05 13:14:15").Format("\\I\\t \\i\\s Y-m-d H:i:s") // It is 2020-08-05 13:14:15
```
Refer to <a href="/appendix/format-signs.html" target="_blank" rel="noreferrer">Format signs</a> for more supported format signs
