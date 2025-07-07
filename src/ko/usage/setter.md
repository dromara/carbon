---
head:
  - - meta
    - name: description
      content: 시간 설정|가벼우면서도 의미론적이고 개발자 친화적인 golang 시간 처리 라이브러리
---

# 시간 설정

```go
// 시간대 설정
carbon.Parse("2020-08-05 13:14:15").SetTimezone(carbon.PRC).ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15").SetTimezone(carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
carbon.Parse("2020-08-05 13:14:15").SetTimezone(carbon.NewYork).ToDateTimeString() // 2020-08-05 01:14:15

// 지역 설정
carbon.Parse("2020-08-05 13:14:15").SetLocale("ko").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15").SetLocale("en").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15").SetLocale("ja").ToDateTimeString() // 2020-08-05 13:14:15

// 연도 설정
carbon.Parse("2020-08-05 13:14:15").SetYear(2021).ToDateTimeString() // 2021-08-05 13:14:15
// 월 설정
carbon.Parse("2020-08-05 13:14:15").SetMonth(12).ToDateTimeString() // 2020-12-05 13:14:15
// 일 설정
carbon.Parse("2020-08-05 13:14:15").SetDay(31).ToDateTimeString() // 2020-08-31 13:14:15
// 시간 설정
carbon.Parse("2020-08-05 13:14:15").SetHour(10).ToDateTimeString() // 2020-08-05 10:14:15
// 분 설정
carbon.Parse("2020-08-05 13:14:15").SetMinute(30).ToDateTimeString() // 2020-08-05 13:30:15
// 초 설정
carbon.Parse("2020-08-05 13:14:15").SetSecond(59).ToDateTimeString() // 2020-08-05 13:14:59

// 주 시작일 설정
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Sunday).WeekStartsAt() // Sunday
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Monday).WeekStartsAt() // Monday
``` 