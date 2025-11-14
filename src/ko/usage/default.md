---
head:
  - - meta
    - name: description
      content: "SetLayout/SetTimezone/SetLocale/SetWeekStartsAt/SetWeekendDays 단일 설정과 SetDefault 일괄 설정 제공, ResetDefault 로 기본값 재설정 지원, 테스트 분리에 유용"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, 단일 기본값 설정, 여러 기본값 설정, 기본값 재설정"
---

# 전역 기본값 설정
기본 시간대는 `UTC`이고, 언어 환경은 `영어`이며, 주 시작일은 `월요일`이고, 주말은 `토요일`과 `일요일`입니다.

## 단일 기본값 설정
```go
carbon.SetLayout(carbon.DateTimeLayout)
carbon.SetTimezone(carbon.UTC)
carbon.SetLocale("en")
carbon.SetWeekStartsAt(carbon.Monday)
carbon.SetWeekendDays([]carbon.Weekday{carbon.Saturday, carbon.Sunday,})
```

## 여러 기본값 설정
```go
carbon.SetDefault(carbon.Default{
  Layout: carbon.DateTimeLayout,
  Timezone: carbon.UTC,
  Locale: "en",
  WeekStartsAt: carbon.Monday,
  WeekendDays: []carbon.Weekday{carbon.Saturday, carbon.Sunday,},
})
```

## 기본값 재설정
```go
carbon.ResetDefault()
```
이는 테스트에서 특히 유용하며, 하나의 테스트에서의 변경사항이 다른 테스트에 영향을 주지 않도록 보장합니다. 