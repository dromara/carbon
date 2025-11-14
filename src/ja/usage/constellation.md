---
head:
  - - meta
    - name: description
      content: Constellation で星座名を取得（ローカライズ対応）、IsAries/IsTaurus/IsGemini などの星座判定メソッドを提供
  - - meta
    - name: keywords
      content: carbon, go-carbon, 星座, おひつじ座, おうし座, ふたご座, かに座, しし座, おとめ座, てんびん座, さそり座, いて座, やぎ座, みずがめ座, うお座
---

# 星座

## 星座名称
```go
carbon.Parse("2020-08-05 13:14:15").Constellation() // Leo
carbon.Parse("2020-08-05 13:14:15").SetLocale("zh-CN").Constellation() // 狮子座
```

## 星座判断
```go
// おひつじ座かどうか
carbon.Parse("2020-08-05 13:14:15").IsAries() // false
// おうし座かどうか
carbon.Parse("2020-08-05 13:14:15").IsTaurus() // false
// ふたご座かどうか
carbon.Parse("2020-08-05 13:14:15").IsGemini() // false
// かに座かどうか
carbon.Parse("2020-08-05 13:14:15").IsCancer() // false
// しし座かどうか
carbon.Parse("2020-08-05 13:14:15").IsLeo() // true
// おとめ座かどうか
carbon.Parse("2020-08-05 13:14:15").IsVirgo() // false
// てんびん座かどうか
carbon.Parse("2020-08-05 13:14:15").IsLibra() // false
// さそり座かどうか
carbon.Parse("2020-08-05 13:14:15").IsScorpio() // false
// いて座かどうか
carbon.Parse("2020-08-05 13:14:15").IsSagittarius() // false
// やぎ座かどうか
carbon.Parse("2020-08-05 13:14:15").IsCapricorn() // false
// みずがめ座かどうか
carbon.Parse("2020-08-05 13:14:15").IsAquarius() // false
// うお座かどうか
carbon.Parse("2020-08-05 13:14:15").IsPisces() // false
```