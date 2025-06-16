---
head:
  - - meta
    - name: description
      content: 设置全局默认值|一个轻量级、语义化、对开发者友好的 golang 时间处理库
  - - meta
    - name: keywords
      content: carbon, go-carbon, 设置单个默认值, 设置多个默认值, 重置默认值
---

# 设置全局默认值
默认时区是 `UTC`, 语言环境是 `英语`，一周开始日期是 `周一`，周末是 `周六`和 `周日`。

## 设置单个默认值
```go
carbon.SetLayout(carbon.DateTimeLayout)
carbon.SetTimezone(carbon.UTC)
carbon.SetLocale("en")
carbon.SetWeekStartsAt(carbon.Monday)
carbon.SetWeekendDays([]carbon.Weekday{carbon.Saturday, carbon.Sunday,})
```

## 设置多个默认值
```go
carbon.SetDefault(carbon.Default{
  Layout: carbon.DateTimeLayout,
  Timezone: carbon.UTC,
  Locale: "en",
  WeekStartsAt: carbon.Monday,
  WeekendDays: []carbon.Weekday{carbon.Saturday, carbon.Sunday,},
})
```

## 重置默认值
```go
carbon.ResetDefault()
```
这在测试中特别有用，确保一个测试中的更改不会影响其他测试。