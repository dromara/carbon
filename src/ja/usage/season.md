---
head:
  - - meta
    - name: description
      content: 季節|軽量、セマンティック、開発者フレンドリーな golang 時間処理ライブラリ
  - - meta
    - name: keywords
      content: carbon, go-carbon, 季節, 春季, 夏季, 秋季, 冬季, 季節境界, 季節判断
---

# 季節
気象区分によると、`3-5`月は`春`、`6-8`月は`夏`、`9-11`月は`秋`、`12-2`月は`冬`
## 季節名称
```go
carbon.Parse("2020-08-05 13:14:15").Season() // Summer
carbon.Parse("2020-08-05 13:14:15").SetLocale("ja").Season() // 夏
```

## 季節境界
```go
// 本季節開始時間
carbon.Parse("2020-08-05 13:14:15").StartOfSeason().ToDateTimeString() // 2020-06-01 00:00:00
// 本季節終了時間
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