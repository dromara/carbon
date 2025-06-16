---
head:
  - - meta
    - name: description
      content: 季节|一个轻量级、语义化、对开发者友好的 golang 时间处理库
  - - meta
    - name: keywords
      content: carbon, go-carbon, 季节, 春季, 夏季, 秋季, 冬季, 季节边界, 季节判断
---

# 季节
按照气象划分，即 `3-5`月为 `春季`，`6-8`月为 `夏季`，`9-11`月为 `秋季`，12-2月为 `冬季`

## 季节名称
```go
carbon.Parse("2020-08-05 13:14:15").Season() // Summer
carbon.Parse("2020-08-05 13:14:15").SetLocale("zh-CN").Season() // 夏季
```

## 季节边界
```go
// 本季节开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfSeason().ToDateTimeString() // 2020-06-01 00:00:00
// 本季节结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfSeason().ToDateTimeString() // 2020-08-31 23:59:59
```

## 季节判断
```go
// 是否是春季
carbon.Parse("2020-08-05 13:14:15").IsSpring() // false
// 是否是夏季
carbon.Parse("2020-08-05 13:14:15").IsSummer() // true
// 是否是秋季
carbon.Parse("2020-08-05 13:14:15").IsAutumn() // false
// 是否是冬季
carbon.Parse("2020-08-05 13:14:15").IsWinter() // false
```