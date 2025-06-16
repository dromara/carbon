---
head:
  - - meta
    - name: description
      content: carbon 、 time.Time 之间互转|一个轻量级、语义化、对开发者友好的 golang 时间处理库
  - - meta
    - name: keywords
      content: carbon, go-carbon, 转换, time.Time
---

# `carbon` 、 `time.Time` 之间互转

## 将标准 `time.Time` 转换成 `carbon`

```go
carbon.NewCarbon(time.Now())
carbon.NewCarbon(time.Now().In(time.Local))
```
或
```go
carbon.CreateFromStdTime(time.Now())
carbon.CreateFromStdTime(time.Now().In(time.Local))
```

## 将 `carbon` 转换成标准 `time.Time`

```go
carbon.Now().StdTime()
carbon.Now(carbon.Local).StdTime()
```