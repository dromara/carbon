---
head:
  - - meta
    - name: description
      content: "提供 Max/Min 获取最大/最小时间、Closest/Farthest 获取最近/最远时间、MaxValue/MinValue 获取边界值、MaxDuration/MinDuration 获取 Duration 极值"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, 极值判断, 极值边界"
---

# 时间极值

## 极值判断
```go
c1 := carbon.Parse("2020-08-01")
c2 := carbon.Parse("2020-08-05")
c3 := carbon.Parse("2020-08-06")

// 返回最大的 Carbon 实例
carbon.Max(c1, c2, c3) // c3
// 返回最小的 Carbon 实例
carbon.Min(c1, c2, c3) // c1

c := carbon.Parse("2020-07-01")
// 返回最近的 Carbon 实例
c.Closest(c1, c2, c3) // c1
// 返回最远的 Carbon 实例
c.Farthest(c1, c2, c3) // c3
```

## 极值边界
```go
// 返回零值 Carbon
carbon.ZeroValue().ToString() // 0001-01-01 00:00:00 +0000 UTC
// 返回 linux 纪元值 Carbon
carbon.EpochValue().ToString() // 1970-01-01 00:00:00 +0000 UTC

// 返回 Carbon 的最大值
carbon.MaxValue().ToString() // 9999-12-31 23:59:59.999999999 +0000 UTC
// 返回 Carbon 的最小值
carbon.MinValue().ToString() // 0001-01-01 00:00:00 +0000 UTC

// 返回 Duration 的最大值
carbon.MaxDuration().Seconds() // 9.223372036854776e+09
// 返回 Duration 的最小值
carbon.MinDuration().Seconds() // -9.223372036854776e+09
```