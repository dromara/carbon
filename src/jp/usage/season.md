# 季節
気象区分によると、`3-5`月は`春で`、`6-8`月は`夏で`、`9-11`月は`秋で`、`12-2`月は`冬です`
## 季節名前
```go
carbon.Parse("2020-08-05 13:14:15").Season() // Summer
carbon.Parse("2020-08-05 13:14:15").SetLocale("jp").Season() // 夏
```

## 季節境界
```go
// この季節の開始日
carbon.Parse("2020-08-05 13:14:15").StartOfSeason().ToDateTimeString() // 2020-06-01 00:00:00
// この季節の最終日
carbon.Parse("2020-08-05 13:14:15").EndOfSeason().ToDateTimeString() // 2020-08-31 23:59:59
```

## 季節判断
```go
// 春かどうか
carbon.Parse("2020-08-05 13:14:15").IsSpring() // false
// 夏かどうか
carbon.Parse("2020-08-05 13:14:15").IsSummer() // true
// 秋かどうか
carbon.Parse("2020-08-05 13:14:15").IsAutumn() // false
// 冬かどうか
carbon.Parse("2020-08-05 13:14:15").IsWinter() // false
```