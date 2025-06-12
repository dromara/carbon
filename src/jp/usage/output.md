# 時間出力

```go
// datetimeを文字列出力
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
// ミリ秒を含むdatetimeを文字列出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeMilliString() // 2020-08-05 13:14:15.999
// マイクロ秒を含むdatetimeを文字列出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeMicroString() // 2020-08-05 13:14:15.999999
// ナノ秒を含むdatetimeを文字列出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeNanoString() // 2020-08-05 13:14:15.999999999

// datetimeを略語形式の文字列出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeString() // 20200805131415
// ミリ秒を含むdatetimeを略語形式の文字列出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeMilliString() // 20200805131415.999
// マイクロ秒を含むdatetimeを略語形式の文字列出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeMicroString() // 20200805131415.999999
// ナノ秒を含むdatetimeを略語形式の文字列出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeNanoString() // 20200805131415.999999999

// dateを文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateString() // 2020-08-05
// ミリ秒を含むdateを文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateMilliString() // 2020-08-05.999
// マイクロ秒を含むdateを文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateMicroString() // 2020-08-05.999999
// ナノ秒を含むdateを文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateNanoString() // 2020-08-05.999999999

// dateを略語形式の文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateString() // 20200805
// ミリ秒を含むdateを略語形式の文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateMilliString() // 20200805.999
// マイクロ秒を含むdateを略語形式の文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateMicroString() // 20200805.999999
// ナノ秒を含むdateを略語形式の文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateNanoString() // 20200805.999999999

// 時間を文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeString() // 13:14:15
// ミリ秒を含む時間を文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeMilliString() // 13:14:15.999
// マイクロ秒を含む時間を文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeMicroString() // 13:14:15.999999
// ナノ秒を含む時間を文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeNanoString() // 13:14:15.999999999

// 時間を略語形式の出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeString() // 131415
// ミリ秒を含む時間を略語形式の出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeMilliString() // 131415.999
// マイクロ秒を含む時間を略語形式の出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeMicroString() // 131415.999999
// ナノ秒を含む時間を略語形式の出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeNanoString() // 131415.999999999

// Ansic フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToAnsicString() // Wed Aug  5 13:14:15 2020
// Atom フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToAtomString() // 2020-08-05T13:14:15+08:00
// UnixDate フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToUnixDateString() // Wed Aug  5 13:14:15 JST 2020
// RubyDate フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRubyDateString() // Wed Aug 05 13:14:15 +0900 2020
// Kitchen フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToKitchenString() // 1:14PM
// Cookie フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToCookieString() // Wednesday, 05-Aug-2020 13:14:15 JST
// DayDateTime フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString() // Wed, Aug 5, 2020 1:14 PM
// RSS フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRssString() // Wed, 05 Aug 2020 13:14:15 +0900
// W3C フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToW3cString() // 2020-08-05T13:14:15+09:00

// ISO8601 フォーマット文字列の出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601String() // 2020-08-05T13:14:15+09:00
// ISO8601Milli フォーマット文字列の出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601MilliString() // 2020-08-05T13:14:15.999+09:00
// ISO8601Micro フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601MicroString() // 2020-08-05T13:14:15.999999+09:00
// ISO8601Nano フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601NanoString() // 2020-08-05T13:14:15.999999999+09:00
// ISO8601Zulu フォーマット文字列の出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluString() // 2020-08-05T13:14:15Z
// ISO8601ZuluMilli フォーマット文字列の出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluMilliString() // 2020-08-05T13:14:15.999Z
// ISO8601ZuluMicro フォーマット文字列の出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluMicroString() // 2020-08-05T13:14:15.999999Z
// ISO8601ZuluNano フォーマット文字列の出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluNanoString() // 2020-08-05T13:14:15.999999999Z

// RFC822 フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRfc822String() // 05 Aug 20 13:14 JST
// RFC822Z フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRfc822zString() // 05 Aug 20 13:14 +0900
// RFC850 フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRfc850String() // Wednesday, 05-Aug-20 13:14:15 JST
// RFC1036 フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRfc1036String() // Wed, 05 Aug 20 13:14:15 +0900
// RFC1123 フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRfc1123String() // Wed, 05 Aug 2020 13:14:15 JST
// RFC1123Z フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRfc1123zString() // Wed, 05 Aug 2020 13:14:15 +0900
// RFC2822 フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRfc2822String() // Wed, 05 Aug 2020 13:14:15 +0900
// RFC7231 フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRfc7231String() // Wed, 05 Aug 2020 13:14:15 JST

// RFC3339 フォーマット文字列の出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339String() // 2020-08-05T13:14:15+09:00
// RFC3339 フォーマット文字列の出力(ミリ秒を含む)
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339MilliString() // 2020-08-05T13:14:15.999+09:00
// RFC3339 フォーマット文字列の出力(マイクロ秒を含む)
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339MicroString() // 2020-08-05T13:14:15.999999+09:00
// RFC3339 フォーマット文字列の出力(ナノ秒を含む)
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339NanoString() // 2020-08-05T13:14:15.999999999+09:00

// 日付時間文字列の出力
fmt.Printf("%s", carbon.Parse("2020-08-05 13:14:15")) // 2020-08-05 13:14:15

// "2006-01-02 15:04:05.999999999 -0700 MST" フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0900 JST

// "Jan 2, 2006" フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToFormattedDateString() // Aug 5, 2020
// "Mon, Jan 2, 2006" フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToFormattedDayDateString() // Wed, Aug 5, 2020

// レイアウトを指定する文字列の出力
carbon.Parse("2020-08-05 13:14:15").Layout(carbon.ISO8601Layout) // 2020-08-05T13:14:15+09:00
carbon.Parse("2020-08-05 13:14:15").Layout("20060102150405") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Layout("2006年01月02日 15时04分05秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Layout("It is 2006-01-02 15:04:05") // It is 2020-08-05 13:14:15

// 指定されたフォーマットの文字列の出力
carbon.Parse("2020-08-05 13:14:15").Format("YmdHis") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Format("Y年m月d日 H时i分s秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Format("l jK \\o\\f F Y h:i:s A") // Wednesday 5th of August 2020 01:14:15 PM
carbon.Parse("2020-08-05 13:14:15").Format("\\I\\t \\i\\s Y-m-d H:i:s") // It is 2020-08-05 13:14:15
```

もっとフォーマットした出力記号は付録を見てください <a href="/jp/appendix/format-signs" target="_blank" rel="noreferrer">書式設定記号</a>