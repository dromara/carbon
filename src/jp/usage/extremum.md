# 時間極值

## 極値判定
```go
c1 := carbon.Parse("2020-08-01")
c2 := carbon.Parse("2020-08-05")
c3 := carbon.Parse("2020-08-06")

// 最大の Carbon インスタンスを返します
carbon.Max(c1, c2, c3) // c3
// 最小の Carbon インスタンスを返します
carbon.Min(c1, c2, c3) // c1

// 最近のCarbonインスタンスを返す
c1.Closest(c2, c3) // c2
// 遠いCarbonインスタンスを返す
c1.Farthest(c2, c3) // c3
```

## 極値境界
```go
// ゼロ値 Carbon を戻す
carbon.ZeroValue().ToString() // 0001-01-01 00:00:00 +0000 UTC
// linux 紀元値 Carbon を戻す
carbon.EpochValue().ToString() // 1970-01-01 00:00:00 +0000 UTC

// Carbonの最大値を戻す
carbon.MaxValue().ToString() // 9999-12-31 23:59:59.999999999 +0000 UTC
// Carbonの最小値を戻す
carbon.MinValue().ToString() // 0001-01-01 00:00:00 +0000 UTC

// 期間の最大値を返します
carbon.MaxDuration().Seconds() // 9.223372036854776e+09
// 最小の持続時間値を返します
carbon.MinDuration().Seconds() // -9.223372036854776e+09
```