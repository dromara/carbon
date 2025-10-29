---
head:
  - - meta
    - name: description
      content: 시간 출력 | 날짜/시간/날짜시간 및 그 축약형에 대한 다중 정밀도 출력(밀리초/마이크로초/나노초)을 제공하고, ISO8601/RFC/Ansic/Kitchen/Cookie 등 표준 포맷을 지원하며, Layout/Format 기반 사용자 지정 서식도 지원합니다
---

# 시간 출력

```go
// 날짜 시간 문자열 출력
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeString() // 2020-08-05 13:14:15
// 날짜 시간 문자열 출력, 밀리초 포함
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeMilliString() // 2020-08-05 13:14:15.999
// 날짜 시간 문자열 출력, 마이크로초 포함
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeMicroString() // 2020-08-05 13:14:15.999999
// 날짜 시간 문자열 출력, 나노초 포함
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeNanoString() // 2020-08-05 13:14:15.999999999

// 축약형 날짜 시간 문자열 출력
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeString() // 20200805131415
// 축약형 날짜 시간 문자열 출력, 밀리초 포함
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeMilliString() // 20200805131415.999
// 축약형 날짜 시간 문자열 출력, 마이크로초 포함
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeMicroString() // 20200805131415.999999
// 축약형 날짜 시간 문자열 출력, 나노초 포함
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeNanoString() // 20200805131415.999999999

// 날짜 문자열 출력
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateString() // 2020-08-05
// 날짜 문자열 출력, 밀리초 포함
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateMilliString() // 2020-08-05.999
// 날짜 문자열 출력, 마이크로초 포함
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateMicroString() // 2020-08-05.999999
// 날짜 문자열 출력, 나노초 포함
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateNanoString() // 2020-08-05.999999999

// 축약형 날짜 문자열 출력
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateString() // 20200805
// 축약형 날짜 문자열 출력, 밀리초 포함
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateMilliString() // 20200805.999
// 축약형 날짜 문자열 출력, 마이크로초 포함
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateMicroString() // 20200805.999999
// 축약형 날짜 문자열 출력, 나노초 포함
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateNanoString() // 20200805.999999999

// 시간 문자열 출력
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeString() // 13:14:15
// 시간 문자열 출력, 밀리초 포함
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeMilliString() // 13:14:15.999
// 시간 문자열 출력, 마이크로초 포함
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeMicroString() // 13:14:15.999999
// 시간 문자열 출력, 나노초 포함
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeNanoString() // 13:14:15.999999999

// 축약형 시간 문자열 출력
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeString() // 131415
// 축약형 시간 문자열 출력, 밀리초 포함
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeMilliString() // 131415.999
// 축약형 시간 문자열 출력, 마이크로초 포함
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeMicroString() // 131415.999999
// 축약형 시간 문자열 출력, 나노초 포함
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeNanoString() // 131415.999999999

// Ansic 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToAnsicString() // Wed Aug  5 13:14:15 2020
// Atom 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToAtomString() // 2020-08-05T13:14:15+00:00
// UnixDate 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToUnixDateString() // Wed Aug  5 13:14:15 UTC 2020
// RubyDate 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToRubyDateString() // Wed Aug 05 13:14:15 +0000 2020
// Kitchen 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToKitchenString() // 1:14PM
// Cookie 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToCookieString() // Wednesday, 05-Aug-2020 13:14:15 UTC
// DayDateTime 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString() // Wed, Aug 5, 2020 1:14 PM
// RSS 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToRssString() // Wed, 05 Aug 2020 13:14:15 +0000
// W3C 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToW3cString() // 2020-08-05T13:14:15+00:00

// ISO8601 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601String() // 2020-08-05T13:14:15+00:00
// ISO8601Milli 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601MilliString() // 2020-08-05T13:14:15.999+00:00
// ISO8601Micro 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601MicroString() // 2020-08-05T13:14:15.999999+00:00
// ISO8601Nano 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601NanoString() // 2020-08-05T13:14:15.999999999+00:00
// ISO8601Zulu 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluString() // 2020-08-05T13:14:15Z
// ISO8601ZuluMilli 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluMilliString() // 2020-08-05T13:14:15.999Z
// ISO8601ZuluMicro 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluMicroString() // 2020-08-05T13:14:15.999999Z
// ISO8601ZuluNano 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluNanoString() // 2020-08-05T13:14:15.999999999Z

// RFC822 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToRfc822String() // 05 Aug 20 13:14 UTC
// RFC822Z 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToRfc822zString() // 05 Aug 20 13:14 +0000
// RFC850 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToRfc850String() // Wednesday, 05-Aug-20 13:14:15 UTC
// RFC1036 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToRfc1036String() // Wed, 05 Aug 20 13:14:15 +0000
// RFC1123 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToRfc1123String() // Wed, 05 Aug 2020 13:14:15 UTC
// RFC1123Z 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToRfc1123zString() // Wed, 05 Aug 2020 13:14:15 +0000
// RFC2822 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToRfc2822String() // Wed, 05 Aug 2020 13:14:15 +0000
// RFC7231 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToRfc7231String() // Wed, 05 Aug 2020 13:14:15 UTC

// RFC3339 형식 문자열 출력
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339String() // 2020-08-05T13:14:15+00:00
// RFC3339Milli 형식 문자열 출력
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339MilliString() // 2020-08-05T13:14:15.999+00:00
// RFC3339Micro 형식 문자열 출력
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339MicroString() // 2020-08-05T13:14:15.999999+00:00
// RFC3339Nano 형식 문자열 출력
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339NanoString() // 2020-08-05T13:14:15.999999999+00:00

// 날짜 시간 문자열 출력
fmt.Printf("%s", carbon.Parse("2020-08-05 13:14:15")) // 2020-08-05 13:14:15

// "2006-01-02 15:04:05.999999999 -0700 MST" 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15.999999 +0000 UTC

// "Jan 2, 2006" 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToFormattedDateString() // Aug 5, 2020
// "Mon, Jan 2, 2006" 형식 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToFormattedDayDateString() // Wed, Aug 5, 2020

// 지정된 레이아웃의 문자열 출력
carbon.Parse("2020-08-05 13:14:15").Layout(carbon.ISO8601Layout) // 2020-08-05T13:14:15+00:00
carbon.Parse("2020-08-05 13:14:15").Layout("20060102150405") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Layout("2006년01월02일 15시04분05초") // 2020년08월05일 13시14분15초
carbon.Parse("2020-08-05 13:14:15").Layout("It is 2006-01-02 15:04:05") // It is 2020-08-05 13:14:15

// 지정된 형식의 문자열 출력(사용하는 문자가 형식 기호와 충돌할 때는 \ 기호로 해당 문자를 이스케이프하세요)
carbon.Parse("2020-08-05 13:14:15").Format("YmdHis") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Format("Y년m월d일 H시i분s초") // 2020년08월05일 13시14분15초
carbon.Parse("2020-08-05 13:14:15").Format("l jK \\o\\f F Y h:i:s A") // Wednesday 5th of August 2020 01:14:15 PM
carbon.Parse("2020-08-05 13:14:15").Format("\\I\\t \\i\\s Y-m-d H:i:s") // It is 2020-08-05 13:14:15
```

더 많은 형식 출력 기호는 <a href="/ko/appendix/format-signs.html" target="_blank" rel="noreferrer">형식 기호</a>를 참조하세요 