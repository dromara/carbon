---
head:
  - - meta
    - name: description
      content: carbon 与 time.Time 互转|提供 CreateFromStdTime/NewCarbon 将 time.Time 转为 carbon，提供 StdTime 导出标准 time.Time，支持指定时区
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