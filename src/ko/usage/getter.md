---
head:
  - - meta
    - name: description
      content: "Carbon 시간 조회 메서드 대전, 년월일시분초 구성 요소, 연내/월내/주내 순서, 다정밀도 타임스탬프, 타임존 정보, 현지화된 별자리와 계절, 주 시작/종료 날짜, 레이아웃 템플릿 및 나이 계산 등 30+ 가지 조회 메서드 제공"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, 시간 조회, getter 메서드, 년월일, 시분초, 타임스탬프, 타임존 정보, 현지화, 별자리, 계절, 나이 계산"
---

# 시간 조회

```go
// 이번 연도 총 일수 조회
carbon.Parse("2019-08-05 13:14:15").DaysInYear() // 365
carbon.Parse("2020-08-05 13:14:15").DaysInYear() // 366
// 이번 달 총 일수 조회
carbon.Parse("2020-02-01 13:14:15").DaysInMonth() // 29
carbon.Parse("2020-04-01 13:14:15").DaysInMonth() // 30
carbon.Parse("2020-08-01 13:14:15").DaysInMonth() // 31

// 이번 연도 몇 번째 날인지 조회
carbon.Parse("2020-08-05 13:14:15").DayOfYear() // 218
// 이번 연도 몇 번째 주인지 조회
carbon.Parse("2019-12-31 13:14:15").WeekOfYear() // 1
carbon.Parse("2020-08-05 13:14:15").WeekOfYear() // 32
// 이번 달 몇 번째 날인지 조회
carbon.Parse("2020-08-05 13:14:15").DayOfMonth() // 5
// 이번 달 몇 번째 주인지 조회
carbon.Parse("2020-08-05 13:14:15").WeekOfMonth() // 1
// 이번 주 몇 번째 날인지 조회
carbon.Parse("2020-08-05 13:14:15").DayOfWeek() // 3

// 현재 년월일시분초 조회
carbon.Parse("2020-08-05 13:14:15").DateTime() // 2020,8,5,13,14,15
// 현재 년월일시분초밀리초 조회
carbon.Parse("2020-08-05 13:14:15").DateTimeMilli() // 2020,8,5,13,14,15,999
// 현재 년월일시분초마이크로초 조회
carbon.Parse("2020-08-05 13:14:15").DateTimeMicro() // 2020,8,5,13,14,15,999999
// 현재 년월일시분초나노초 조회
carbon.Parse("2020-08-05 13:14:15").DateTimeNano() // 2020,8,5,13,14,15,999999999

// 현재 년월일 조회
carbon.Parse("2020-08-05 13:14:15.999999999").Date() // 2020,8,5
// 현재 년월일밀리초 조회
carbon.Parse("2020-08-05 13:14:15.999999999").DateMilli() // 2020,8,5,999
// 현재 년월일마이크로초 조회
carbon.Parse("2020-08-05 13:14:15.999999999").DateMicro() // 2020,8,5,999999
// 현재 년월일나노초 조회
carbon.Parse("2020-08-05 13:14:15.999999999").DateNano() // 2020,8,5,999999999

// 현재 시분초 조회
carbon.Parse("2020-08-05 13:14:15.999999999").Time() // 13,14,15
// 현재 시분초밀리초 조회
carbon.Parse("2020-08-05 13:14:15.999999999").TimeMilli() // 13,14,15,999
// 현재 시분초마이크로초 조회
carbon.Parse("2020-08-05 13:14:15.999999999").TimeMicro() // 13,14,15,999999
// 현재 시분초나노초 조회
carbon.Parse("2020-08-05 13:14:15.999999999").TimeNano() // 13,14,15,999999999

// 현재 세기 조회
carbon.Parse("2020-08-05 13:14:15").Century() // 21
// 현재 연대 조회
carbon.Parse("2019-08-05 13:14:15").Decade() // 10
carbon.Parse("2021-08-05 13:14:15").Decade() // 20
// 현재 연도 조회
carbon.Parse("2020-08-05 13:14:15").Year() // 2020
// 현재 분기 조회
carbon.Parse("2020-08-05 13:14:15").Quarter() // 3
// 현재 월 조회
carbon.Parse("2020-08-05 13:14:15").Month() // 8
// 현재 주 조회(0부터 시작)
carbon.Parse("2020-08-02 13:14:15").Week() // 0
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Sunday).Week() // 0
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Monday).Week() // 6
// 현재 일 조회
carbon.Parse("2020-08-05 13:14:15").Day() // 5
// 현재 시간 조회
carbon.Parse("2020-08-05 13:14:15").Hour() // 13
// 현재 분 조회
carbon.Parse("2020-08-05 13:14:15").Minute() // 14
// 현재 초 조회
carbon.Parse("2020-08-05 13:14:15").Second() // 15
// 현재 밀리초 조회
carbon.Parse("2020-08-05 13:14:15.999").Millisecond() // 999
// 현재 마이크로초 조회
carbon.Parse("2020-08-05 13:14:15.999").Microsecond() // 999000
// 현재 나노초 조회
carbon.Parse("2020-08-05 13:14:15.999").Nanosecond() // 999000000

// 초 정밀도 타임스탬프 조회
carbon.Parse("2020-08-05 13:14:15").Timestamp() // 1596633255
// 밀리초 정밀도 타임스탬프 조회
carbon.Parse("2020-08-05 13:14:15.999").TimestampMilli() // 1596633255999
// 마이크로초 정밀도 타임스탬프 조회
carbon.Parse("2020-08-05 13:14:15.999999").TimestampMicro() // 1596633255999999
// 나노초 정밀도 타임스탬프 조회
carbon.Parse("2020-08-05 13:14:15.999999999").TimestampNano() // 1596633255999999999

// 시간대 위치 조회
carbon.SetTimezone(carbon.PRC).Timezone() // PRC
carbon.SetTimezone(carbon.Tokyo).Timezone() // Asia/Tokyo

// 시간대 이름 조회
carbon.SetTimezone(carbon.PRC).ZoneName() // CST
carbon.SetTimezone(carbon.Tokyo).ZoneName() // JST

// 시간대 오프셋 조회(초 단위)
carbon.SetTimezone(carbon.PRC).ZoneOffset() // 28800
carbon.SetTimezone(carbon.Tokyo).ZoneOffset() // 32400

// 현재 지역 조회
carbon.Now().Locale() // zh-CN
carbon.Now().SetLocale("ko").Locale() // ko

// 현재 별자리 조회
carbon.Now().Constellation() // 사자자리
carbon.Now().SetLocale("en").Constellation() // Leo
carbon.Now().SetLocale("ko-KR").Constellation() // 사자자리

// 현재 계절 조회
carbon.Now().Season() // 여름
carbon.Now().SetLocale("en").Season() // Summer
carbon.Now().SetLocale("ko-KR").Season() // 여름

// 주 시작 날짜 조회
carbon.SetWeekStartsAt(carbon.Sunday).WeekStartsAt() // Sunday
carbon.SetWeekStartsAt(carbon.Monday).WeekStartsAt() // Monday

// 주 종료 날짜 조회
carbon.SetWeekStartsAt(carbon.Sunday).WeekEndsAt() // Saturday
carbon.SetWeekStartsAt(carbon.Monday).WeekEndsAt() // Sunday

// 현재 레이아웃 템플릿 조회
carbon.Parse("now").CurrentLayout() // "2006-01-02 15:04:05"
carbon.ParseByLayout("2020-08-05", DateLayout).CurrentLayout() // "2006-01-02"

// 나이 조회
carbon.Parse("2002-01-01 13:14:15").Age() // 17
carbon.Parse("2002-12-31 13:14:15").Age() // 18
``` 