---
head:
  - - meta
    - name: description
      content: Carbon 계절 기능 상세 설명, Season으로 계절 이름 조회(지역화 지원), StartOfSeason/EndOfSeason로 계절 경계 조회, IsSpring/IsSummer/IsAutumn/IsWinter 계절 판단, 기상 구분에 따른(3-5 월 봄, 6-8 월 여름, 9-11 월 가을, 12-2 월 겨울)
  - - meta
    - name: keywords
      content: carbon, go-carbon, 계절, 봄, 여름, 가을, 겨울, 계절 경계, 계절 판단, Season, StartOfSeason, EndOfSeason, 지역화
---

# 계절
기상학적 구분에 따라 `3-5월`은 `봄`, `6-8월`은 `여름`, `9-11월`은 `가을`, 12-2월은 `겨울`입니다

## 계절 명칭
```go
carbon.Parse("2020-08-05 13:14:15").Season() // Summer
carbon.Parse("2020-08-05 13:14:15").SetLocale("ko").Season() // 여름
```

## 계절 경계
```go
// 이번 계절 시작 시간
carbon.Parse("2020-08-05 13:14:15").StartOfSeason().ToDateTimeString() // 2020-06-01 00:00:00
// 이번 계절 종료 시간
carbon.Parse("2020-08-05 13:14:15").EndOfSeason().ToDateTimeString() // 2020-08-31 23:59:59
```

## 계절 판단
```go
// 봄인지
carbon.Parse("2020-08-05 13:14:15").IsSpring() // false
// 여름인지
carbon.Parse("2020-08-05 13:14:15").IsSummer() // true
// 가을인지
carbon.Parse("2020-08-05 13:14:15").IsAutumn() // false
// 겨울인지
carbon.Parse("2020-08-05 13:14:15").IsWinter() // false
``` 