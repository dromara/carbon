---
head:
  - - meta
    - name: description
      content: 시간 차이|가벼우면서도 의미론적이고 개발자 친화적인 golang 시간 처리 라이브러리
---

# 시간 차이

```go
// 시간 차이 계산
c1 := carbon.Parse("2020-08-05 13:14:15")
c2 := carbon.Parse("2020-08-06 13:14:15")

// 절대 차이
c1.DiffInYears(c2) // 0
c1.DiffInMonths(c2) // 0
c1.DiffInWeeks(c2) // 0
c1.DiffInDays(c2) // 1
c1.DiffInHours(c2) // 24
c1.DiffInMinutes(c2) // 1440
c1.DiffInSeconds(c2) // 86400

// 부호 있는 차이
c1.DiffInYearsWithAbs(c2) // 0
c1.DiffInMonthsWithAbs(c2) // 0
c1.DiffInWeeksWithAbs(c2) // 0
c1.DiffInDaysWithAbs(c2) // -1
c1.DiffInHoursWithAbs(c2) // -24
c1.DiffInMinutesWithAbs(c2) // -1440
c1.DiffInSecondsWithAbs(c2) // -86400

// 인간 친화적 차이
c1.DiffForHumans() // 1 day ago
c1.DiffForHumans(c2) // 1 day before
c1.SetLocale("ko").DiffForHumans() // 1일 전
c1.SetLocale("ko").DiffForHumans(c2) // 1일 전
``` 