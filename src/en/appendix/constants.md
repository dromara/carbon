---
head:
  - - meta
    - name: description
      content: "Carbon Library Built-in Constants Reference Manual, including version, timezone, month, season, constellation, weekday, numeric constants and complete layout template and format constants, covering international standard formats such as RFC and ISO8601"
  - - meta
    - name: keywords
      content: "version constants, timezone constants, month constants, season constants, constellation constants, weekday constants, numeric constants, layout template constants, RFC format, ISO8601 format, time formatting"
---

# Constants

## Version constants
```go
carbon.Version // 2.6.15
```

## Timezone constants
```go
carbon.Local // Local
carbon.UTC   // UTC

carbon.CET  // CET
carbon.EET  // EET
carbon.EST  // EST
carbon.GMT  // GMT
carbon.MET  // MET
carbon.MST  // MST
carbon.UCT  // MST
carbon.WET  // WET
carbon.Zulu // Zulu

carbon.Cuba      // Cuba
carbon.Egypt     // Egypt
carbon.Eire      // Eire
carbon.Greenwich // Greenwich
carbon.Iceland   // Iceland
carbon.Iran      // Iran
carbon.Israel    // Israel
carbon.Jamaica   // Jamaica
carbon.Japan     // Japan
carbon.Libya     // Libya
carbon.Poland    // Poland
carbon.Portugal  // Portugal
carbon.PRC       // PRC
carbon.Singapore // Singapore
carbon.Turkey    // Turkey

carbon.Shanghai   // Asia/Shanghai
carbon.Chongqing  // Asia/Chongqing
carbon.Harbin     // Asia/Harbin
carbon.Urumqi     // Asia/Urumqi
carbon.HongKong   // Asia/Hong_Kong
carbon.Macao      // Asia/Macao
carbon.Taipei     // Asia/Taipei
carbon.Tokyo      // Asia/Tokyo
carbon.HoChiMinh  // Asia/Ho_Chi_Minh
carbon.Hanoi      // Asia/Hanoi
carbon.Saigon     // Asia/Saigon
carbon.Seoul      // Asia/Seoul
carbon.Pyongyang  // Asia/Pyongyang
carbon.Bangkok    // Asia/Bangkok
carbon.Dubai      // Asia/Dubai
carbon.Qatar      // Asia/Qatar
carbon.Bangalore  // Asia/Bangalore
carbon.Kolkata    // Asia/Kolkata
carbon.Mumbai     // Asia/Mumbai
carbon.MexicoCity // America/Mexico_City
carbon.NewYork    // America/New_York
carbon.LosAngeles // America/Los_Angeles
carbon.Chicago    // America/Chicago
carbon.SaoPaulo   // America/Sao_Paulo
carbon.Moscow     // Europe/Moscow
carbon.London     // Europe/London
carbon.Berlin     // Europe/Berlin
carbon.Paris      // Europe/Paris
carbon.Rome       // Europe/Rome
carbon.Sydney     // Australia/Sydney
carbon.Melbourne  // Australia/Melbourne
carbon.Darwin     // Australia/Darwin
```

## Month constants
```go
carbon.January   // time.January
carbon.February  // time.February
carbon.March     // time.March
carbon.April     // time.April
carbon.May       // time.May
carbon.June      // time.June
carbon.July      // time.July
carbon.August    // time.August
carbon.September // time.September
carbon.October   // time.October
carbon.November  // time.November
carbon.December  // time.December
```

## Season constants
```go
carbon.Spring // Spring
carbon.Summer // Summer
carbon.Autumn // Autumn
carbon.Winter // Winter
```

## Constellation constants
```go
carbon.Aries       // Aries
carbon.Taurus      // Taurus
carbon.Gemini      // Gemini
carbon.Cancer      // Cancer
carbon.Leo         // Leo
carbon.Virgo       // Virgo
carbon.Libra       // Libra
carbon.Scorpio     // Scorpio
carbon.Sagittarius // Sagittarius
carbon.Capricorn   // Capricorn
carbon.Aquarius    // Aquarius
carbon.Pisces      // Pisces
```

## Week constants
```go
carbon.Monday    // time.Monday
carbon.Tuesday   // time.Tuesday
carbon.Wednesday // time.Wednesday
carbon.Thursday  // time.Thursday
carbon.Friday    // time.Friday
carbon.Saturday  // time.Saturday
carbon.Sunday    // time.Sunday
```

## Number constants
```go
carbon.EpochYear          // 1970
carbon.YearsPerMillennium // 1000
carbon.YearsPerCentury    // 100
carbon.YearsPerDecade     // 10
carbon.QuartersPerYear    // 4
carbon.MonthsPerYear      // 12
carbon.MonthsPerQuarter   // 3
carbon.WeeksPerNormalYear // 52
carbon.WeeksPerLongYear   // 53
carbon.WeeksPerMonth      // 4
carbon.DaysPerLeapYear    // 366
carbon.DaysPerNormalYear  // 365
carbon.DaysPerWeek        // 7
carbon.HoursPerWeek       // 168
carbon.HoursPerDay        // 24
carbon.MinutesPerDay      // 1440
carbon.MinutesPerHour     // 60
carbon.SecondsPerWeek     // 604800
carbon.SecondsPerDay      // 86400
carbon.SecondsPerHour     // 3600
carbon.SecondsPerMinute   // 60

carbon.EpochYear     // 1970
carbon.MaxYear       // 9999
carbon.MinYear       // 1
carbon.MaxMonth      // 12
carbon.MinMonth      // 1
carbon.MaxDay        // 31
carbon.MinDay        // 1
carbon.MaxHour       // 23
carbon.MinHour       // 0
carbon.MaxMinute     // 59
carbon.MinMinute     // 0
carbon.MaxSecond     // 59
carbon.MinSecond     // 0
carbon.MaxNanosecond // 999999999
carbon.MinNanosecond // 0
```

## Layout constants
```go
carbon.AtomLayout     // 2006-01-02T15:04:05Z07:00"
carbon.ANSICLayout    // Mon Jan _2 15:04:05 2006
carbon.CookieLayout   // Monday, 02-Jan-2006 15:04:05 MST
carbon.KitchenLayout  // 3:04PM
carbon.RssLayout      // Mon, 02 Jan 2006 15:04:05 -0700
carbon.RubyDateLayout // Mon Jan 02 15:04:05 -0700 2006
carbon.UnixDateLayout // Mon Jan _2 15:04:05 MST 2006
carbon.W3cLayout      // 2006-01-02T15:04:05Z07:00
carbon.HttpLayout     // Mon, 02 Jan 2006 15:04:05 GMT

carbon.RFC1036Layout      // Mon, 02 Jan 06 15:04:05 -0700
carbon.RFC1123Layout      // Mon, 02 Jan 2006 15:04:05 MST
carbon.RFC1123ZLayout     // Mon, 02 Jan 2006 15:04:05 -0700
carbon.RFC2822Layout      // Mon, 02 Jan 2006 15:04:05 -0700
carbon.RFC3339Layout      // 2006-01-02T15:04:05Z07:00
carbon.RFC3339MilliLayout // 2006-01-02T15:04:05.999Z07:00
carbon.RFC3339MicroLayout // 2006-01-02T15:04:05.999999Z07:00
carbon.RFC3339NanoLayout  // 2006-01-02T15:04:05.999999999Z07:00
carbon.RFC7231Layout      // Mon, 02 Jan 2006 15:04:05 MST
carbon.RFC822Layout       // 02 Jan 06 15:04 MST
carbon.RFC822ZLayout      // 02 Jan 06 15:04 -0700
carbon.RFC850Layout       // Monday, 02-Jan-06 15:04:05 MST

carbon.ISO8601Layout      // 2006-01-02T15:04:05-07:00
carbon.ISO8601MilliLayout // 2006-01-02T15:04:05.999-07:00
carbon.ISO8601MicroLayout // 2006-01-02T15:04:05.999999-07:00
carbon.ISO8601NanoLayout  // 2006-01-02T15:04:05.999999999-07:00

carbon.ISO8601ZuluLayout      // 2006-01-02T15:04:05Z
carbon.ISO8601ZuluMilliLayout // 2006-01-02T15:04:05.999Z
carbon.ISO8601ZuluMicroLayout // 2006-01-02T15:04:05.999999Z
carbon.ISO8601ZuluNanoLayout  // 2006-01-02T15:04:05.999999999Z

carbon.FormattedDateLayout    // Jan 2, 2006
carbon.FormattedDayDateLayout // Mon, Jan 2, 2006

carbon.DayDateTimeLayout        // Mon, Jan 2, 2006 3:04 PM
carbon.DateTimeLayout           // 2006-01-02 15:04:05
carbon.DateTimeMilliLayout      // 2006-01-02 15:04:05.999
carbon.DateTimeMicroLayout      // 2006-01-02 15:04:05.999999
carbon.DateTimeNanoLayout       // 2006-01-02 15:04:05.999999999
carbon.ShortDateTimeLayout      // 20060102150405
carbon.ShortDateTimeMilliLayout // 20060102150405.999
carbon.ShortDateTimeMicroLayout // 20060102150405.999999
carbon.ShortDateTimeNanoLayout  // 20060102150405.999999999

carbon.DateLayout           // 2006-01-02
carbon.DateMilliLayout      // 2006-01-02.999
carbon.DateMicroLayout      // 2006-01-02.999999
carbon.DateNanoLayout       // 2006-01-02.999999999
carbon.ShortDateLayout      // 20060102
carbon.ShortDateMilliLayout // 20060102.999
carbon.ShortDateMicroLayout // 20060102.999999
carbon.ShortDateNanoLayout  // 20060102.999999999

carbon.TimeLayout           // 15:04:05
carbon.TimeMilliLayout      // 15:04:05.999
carbon.TimeMicroLayout      // 15:04:05.999999
carbon.TimeNanoLayout       // 15:04:05.999999999
carbon.ShortTimeLayout      // 150405
carbon.ShortTimeMilliLayout // 150405.999
carbon.ShortTimeMicroLayout // 150405.999999
carbon.ShortTimeNanoLayout  // 150405.999999999

carbon.TimestampLayout      // unix
carbon.TimestampMilliLayout // unixMilli
carbon.TimestampMicroLayout // unixMicro
carbon.TimestampNanoLayout  // unixNano
```

## Format constants
```go
carbon.AtomFormat     // Y-m-d\\TH:i:sR
carbon.ANSICFormat    // D M  j H:i:s Y
carbon.CookieFormat   // l, d-M-Y H:i:s Z
carbon.KitchenFormat  // g:iA
carbon.RssFormat      // D, d M Y H:i:s O
carbon.RubyDateFormat // D M d H:i:s O Y
carbon.UnixDateFormat // D M  j H:i:s Z Y
carbon.W3cFormat      // Y-m-d\\TH:i:sR
carbon.HttpFormat     // D, d M Y H:i:s GMT

carbon.RFC1036Format      // D, d M y H:i:s O
carbon.RFC1123Format      // D, d M Y H:i:s Z
carbon.RFC1123ZFormat     // D, d M Y H:i:s O
carbon.RFC2822Format      // D, d M Y H:i:s O
carbon.RFC3339Format      // Y-m-d\\TH:i:sR
carbon.RFC3339MilliFormat // Y-m-d\\TH:i:s.uR
carbon.RFC3339MicroFormat // Y-m-d\\TH:i:s.vR
carbon.RFC3339NanoFormat  // Y-m-d\\TH:i:s.xR
carbon.RFC7231Format      // D, d M Y H:i:s Z
carbon.RFC822Format       // d M y H:i Z
carbon.RFC822ZFormat      // d M y H:i O
carbon.RFC850Format       // l, d-M-y H:i:s Z

carbon.ISO8601Format      // Y-m-d\\TH:i:sP
carbon.ISO8601MilliFormat // Y-m-d\\TH:i:s.uP
carbon.ISO8601MicroFormat // Y-m-d\\TH:i:s.vP
carbon.ISO8601NanoFormat  // Y-m-d\\TH:i:s.xP

carbon.ISO8601ZuluFormat      // Y-m-d\\TH:i:s\\Z
carbon.ISO8601ZuluMilliFormat // Y-m-d\\TH:i:s.u\\Z
carbon.ISO8601ZuluMicroFormat // Y-m-d\\TH:i:s.v\\Z
carbon.ISO8601ZuluNanoFormat  // Y-m-d\\TH:i:s.x\\Z

carbon.FormattedDateFormat    // M j, Y
carbon.FormattedDayDateFormat // D, M j, Y

carbon.DayDateTimeFormat        // D, M j, Y g:i A
carbon.DateTimeFormat           // Y-m-d H:i:s
carbon.DateTimeMilliFormat      // Y-m-d H:i:s.u
carbon.DateTimeMicroFormat      // Y-m-d H:i:s.v
carbon.DateTimeNanoFormat       // Y-m-d H:i:s.x
carbon.ShortDateTimeFormat      // YmdHis
carbon.ShortDateTimeMilliFormat // YmdHis.u
carbon.ShortDateTimeMicroFormat // YmdHis.v
carbon.ShortDateTimeNanoFormat  // YmdHis.x

carbon.DateFormat           // Y-m-d
carbon.DateMilliFormat      // Y-m-d.u
carbon.DateMicroFormat      // Y-m-d.v
carbon.DateNanoFormat       // Y-m-d.x
carbon.ShortDateFormat      // Ymd
carbon.ShortDateMilliFormat // Ymd.u
carbon.ShortDateMicroFormat // Ymd.v
carbon.ShortDateNanoFormat  // Ymd.x

carbon.TimeFormat           // H:i:s
carbon.TimeMilliFormat      // H:i:s.u
carbon.TimeMicroFormat      // H:i:s.v
carbon.TimeNanoFormat       // H:i:s.x
carbon.ShortTimeFormat      // His
carbon.ShortTimeMilliFormat // His.u
carbon.ShortTimeMicroFormat // His.v
carbon.ShortTimeNanoFormat  // His.x

carbon.TimestampFormat      // S
carbon.TimestampMilliFormat // U
carbon.TimestampMicroFormat // V
carbon.TimestampNanoFormat  // X
```

