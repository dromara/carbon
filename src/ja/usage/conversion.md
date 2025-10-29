---
head:
  - - meta
    - name: description
      content: carbon と time.Time の相互変換 | CreateFromStdTime/NewCarbon で time.Time から carbon へ、StdTime で標準 time.Time へ出力、タイムゾーン指定に対応
  - - meta
    - name: keywords
      content: carbon, go-carbon, 変換, time.Time
---

# `carbon`、`time.Time` 間の相互変換

## 標準 `time.Time` を `carbon` に変換

```go
carbon.NewCarbon(time.Now())
carbon.NewCarbon(time.Now().In(time.Local))
```
または
```go
carbon.CreateFromStdTime(time.Now())
carbon.CreateFromStdTime(time.Now().In(time.Local))
```

## `carbon` を標準 `time.Time` に変換

```go
carbon.Now().StdTime()
carbon.Now(carbon.Local).StdTime()
```