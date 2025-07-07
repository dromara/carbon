---
head:
  - - meta
    - name: description
      content: 달력|가벼우면서도 의미론적이고 개발자 친화적인 golang 시간 처리 라이브러리
---

# 달력

```go
// 달력 판단
carbon.Parse("2020-01-01").IsJanuary() // true
carbon.Parse("2020-02-01").IsFebruary() // true
carbon.Parse("2020-08-05").IsAugust() // true
carbon.Parse("2020-12-01").IsDecember() // true

// 요일 판단
carbon.Parse("2020-08-02").IsMonday() // true
carbon.Parse("2020-08-03").IsTuesday() // true
carbon.Parse("2020-08-04").IsWednesday() // true
carbon.Parse("2020-08-05").IsThursday() // true
carbon.Parse("2020-08-06").IsFriday() // true
carbon.Parse("2020-08-07").IsSaturday() // true
carbon.Parse("2020-08-08").IsSunday() // true

// 특별한 날 판단
carbon.Parse("2020-01-01").IsNewYear() // true
carbon.Parse("2020-02-14").IsValentinesDay() // true
carbon.Parse("2020-04-01").IsAprilFoolsDay() // true
carbon.Parse("2020-12-25").IsChristmas() // true

// 윤년 판단
carbon.Parse("2020-08-05").IsLeapYear() // true
carbon.Parse("2021-08-05").IsLeapYear() // false

// 평일/주말 판단
carbon.Parse("2020-08-05").IsWeekday() // true
carbon.Parse("2020-08-08").IsWeekend() // true

// 시간 범위 판단
carbon.Parse("2020-08-05 13:14:15").IsAM() // false
carbon.Parse("2020-08-05 13:14:15").IsPM() // true
carbon.Parse("2020-08-05 00:00:00").IsMidnight() // true
carbon.Parse("2020-08-05 12:00:00").IsNoon() // true

// 날짜 유효성 검사
carbon.IsValid("2020-08-05") // true
carbon.IsValid("2020-13-05") // false
carbon.IsValid("2020-08-32") // false
``` 