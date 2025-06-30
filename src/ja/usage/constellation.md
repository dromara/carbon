# 星座

## 星座名前
```go
carbon.Parse("2020-08-05 13:14:15").Constellation() // Leo
carbon.Parse("2020-08-05 13:14:15").SetLocale("ja").Constellation() // しし座
```

## 星座判定
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