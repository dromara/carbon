---
head:
  - - meta
    - name: description
      content: 내장 상수|가벼우면서도 의미론적이고 개발자 친화적인 golang 시간 처리 라이브러리
  - - meta
    - name: keywords
      content: 버전 상수, 시간대 상수, 월 상수, 계절 상수, 별자리 상수, 요일 상수, 숫자 상수, 레이아웃 템플릿 상수
---

# 내장 상수

## 버전 상수
```go
carbon.Version // 2.6.15
```

## 시간대 상수
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

## 월 상수
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

## 계절 상수
```go
carbon.Spring // Spring
carbon.Summer // Summer
carbon.Autumn // Autumn
carbon.Winter // Winter
```

## 별자리 상수
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

## 요일 상수
```go
carbon.Monday    // time.Monday
carbon.Tuesday   // time.Tuesday
carbon.Wednesday // time.Wednesday
carbon.Thursday  // time.Thursday
carbon.Friday    // time.Friday
carbon.Saturday  // time.Saturday
carbon.Sunday    // time.Sunday
```

## 숫자 상수
```go
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

## 레이아웃 상수
```go
carbon.AtomLayout     // 2006-01-02T15:04:05Z07:00"
carbon.ANSICLayout    // Mon Jan _2 15:04:05 2006
carbon.CookieLayout   // Monday, 02-Jan-2006 15:04:05 MST
carbon.KitchenLayout  // 3:04PM
carbon.RssLayout      // Mon, 02 Jan 2006 15:04:05 -0700
carbon.RubyDateLayout // Mon Jan 02 15:04:05 -0700 2006
carbon.UnixDateLayout // Mon Jan _2 15:04:05 MST 2006
carbon.W3cLayout      // 2006-01-02T15:04:05Z07:00

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
``` 