---
head:
  - - meta
    - name: description
      content: 계절|가벼우면서도 의미론적이고 개발자 친화적인 golang 시간 처리 라이브러리
  - - meta
    - name: keywords
      content: carbon, go-carbon, 계절, 봄, 여름, 가을, 겨울, 계절 경계, 계절 판단
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