---
head:
  - - meta
    - name: description
      content: 시간 출력|가벼우면서도 의미론적이고 개발자 친화적인 golang 시간 처리 라이브러리
---

# 시간 출력

```go
// 문자열 출력
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15").ToDateString() // 2020-08-05
carbon.Parse("2020-08-05 13:14:15").ToTimeString() // 13:14:15
carbon.Parse("2020-08-05 13:14:15").ToFormattedDateString() // Aug 5, 2020
carbon.Parse("2020-08-05 13:14:15").ToFormattedTimeString() // 1:14 PM

// 형식 지정 출력
carbon.Parse("2020-08-05 13:14:15").Format("Y-m-d H:i:s") // 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15").Format("Y년 m월 d일 H시 i분 s초") // 2020년 08월 05일 13시 14분 15초
carbon.Parse("2020-08-05 13:14:15").Format("Y/m/d") // 2020/08/05
carbon.Parse("2020-08-05 13:14:15").Format("H:i:s") // 13:14:15

// 레이아웃 출력
carbon.Parse("2020-08-05 13:14:15").Layout(carbon.DateTimeLayout) // 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15").Layout(carbon.DateLayout) // 2020-08-05
carbon.Parse("2020-08-05 13:14:15").Layout(carbon.TimeLayout) // 13:14:15
carbon.Parse("2020-08-05 13:14:15").Layout("2006-01-02 15:04:05") // 2020-08-05 13:14:15

// 인간 친화적 시간 차이 출력
carbon.Parse("2020-08-05 13:14:15").DiffForHumans() // 1 second ago
carbon.Parse("2020-08-05 13:14:15").DiffForHumans(carbon.Now()) // 1 second before
carbon.Parse("2020-08-05 13:14:15").SetLocale("ko").DiffForHumans() // 1초 전
``` 