---
head:
  - - meta
    - name: description
      content: 冻结测试|一个轻量级、语义化、对开发者友好的 golang 时间处理库
  - - meta
    - name: keywords
      content: carbon, go-carbon, 测试时间, 冻结时间
---

# 测试
支持冻结固定时间，将任意时间设置为`当前时间`，使后续操作基于此模拟时间运行，而非真实系统时间，便于单元测试

## 设置测试时间
```go
now := carbon.Parse("2020-08-05")
carbon.SetTestNow(now)

carbon.Now().ToDateString() // 2020-08-05
carbon.Yesterday().ToDateString() // 2020-08-04
carbon.Tomorrow().ToDateString() // 2020-08-05
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