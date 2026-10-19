---
head:
  - - meta
    - name: description
      content: "Carbon 测试功能详解，提供 SetTestNow/IsTestNow/ClearTestNow/Sleep 方法，支持冻结固定时间用于单元测试，使时间相关操作基于模拟时间运行，便于测试时间依赖代码"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, 测试时间, 冻结时间, 模拟时间, SetTestNow, IsTestNow, ClearTestNow, Sleep, 单元测试, 时间测试, 测试工具"
---

# 测试
支持冻结固定时间，将任意时间设置为`当前时间`，使后续操作基于此模拟时间运行，而非真实系统时间，便于单元测试

## 设置测试时间
```go
now := carbon.Parse("2020-08-05")
carbon.SetTestNow(now)

carbon.Now().ToDateString() // 2020-08-05
carbon.Yesterday().ToDateString() // 2020-08-04
carbon.Tomorrow().ToDateString() // 2020-08-06
carbon.Now().DiffForHumans() // just now
carbon.Yesterday().DiffForHumans() // 1 day ago
carbon.Tomorrow().DiffForHumans() // 1 day from now
carbon.Parse("2020-10-05").DiffForHumans() // 2 months from now
now.DiffForHumans(carbon.Parse("2020-10-05")) // 2 months before
```

## 是否是测试时间
```go
carbon.IsTestNow() 
```

## 清除测试时间
```go
carbon.ClearTestNow()
```

## 睡眠
在冻结时间状态下，`Sleep` 不等待真实时间，而是直接推进冻结的当前时间；未冻结时行为与 `time.Sleep` 一致。

```go
carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15"))
carbon.Sleep(2 * time.Hour)

carbon.Now().ToDateTimeString() // 2020-08-05 15:14:15
```