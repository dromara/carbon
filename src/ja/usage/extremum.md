---
head:
  - - meta
    - name: description
      content: Carbon は Max/Min で最大/最小時間を取得、Closest/Farthest で最も近い/遠い時間を取得、MaxValue/MinValue で境界値を取得、MaxDuration/MinDuration で Duration 極値を取得する機能を提供
  - - meta
    - name: keywords
      content: carbon, go-carbon, 極値判定, 極値境界
---

# 時間極値

## 極値判定
```go
c1 := carbon.Parse("2020-08-01")
c2 := carbon.Parse("2020-08-05")
c3 := carbon.Parse("2020-08-06")

// 最大の Carbon インスタンスを返します
carbon.Max(c1, c2, c3) // c3
// 最小の Carbon インスタンスを返します
carbon.Min(c1, c2, c3) // c1

c := carbon.Parse("2020-07-01")
// 最も近い Carbon インスタンスを返します
c.Closest(c1, c2, c3) // c1
// 最も遠い Carbon インスタンスを返します
c.Farthest(c1, c2, c3) // c3
```

## 極値境界
```go
// ゼロ値 Carbon を返します
carbon.ZeroValue().ToString() // 0001-01-01 00:00:00 +0000 UTC
// UNIX 紀元値 Carbon を返します
carbon.EpochValue().ToString() // 1970-01-01 00:00:00 +0000 UTC

// Carbon の最大値を返します
carbon.MaxValue().ToString() // 9999-12-31 23:59:59.999999999 +0000 UTC
// Carbon の最小値を返します
carbon.MinValue().ToString() // 0001-01-01 00:00:00 +0000 UTC

// 期間の最大値を返します
carbon.MaxDuration().Seconds() // 9.223372036854776e+09
// Duration の最小値を返します
carbon.MinDuration().Seconds() // -9.223372036854776e+09
```