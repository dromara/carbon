---
head:
  - - meta
    - name: description
      content: carbon과 time.Time 간 상호 변환|가벼우면서도 의미론적이고 개발자 친화적인 golang 시간 처리 라이브러리
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