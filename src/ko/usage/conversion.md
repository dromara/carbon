---
head:
  - - meta
    - name: description
      content: CreateFromStdTime/NewCarbon 으로 time.Time 을 carbon 으로 변환, StdTime 으로 표준 time.Time 내보내기, 타임존 지정 지원
  - - meta
    - name: keywords
      content: carbon, go-carbon, 변환, time.Time
---

# `carbon`과 `time.Time` 간 상호 변환

## 표준 `time.Time`을 `carbon`으로 변환

```go
carbon.NewCarbon(time.Now())
carbon.NewCarbon(time.Now().In(time.Local))
```
또는
```go
carbon.CreateFromStdTime(time.Now())
carbon.CreateFromStdTime(time.Now().In(time.Local))
```

## `carbon`을 표준 `time.Time`으로 변환

```go
carbon.Now().StdTime()
carbon.Now(carbon.Local).StdTime()
``` 