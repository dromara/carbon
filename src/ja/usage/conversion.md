# `carbon`、`time.Time` の間で相互に回転する

## `time.Time` を `carbon` に変換

```go
carbon.NewCarbon(time.Now())
carbon.NewCarbon(time.Now().In(time.Local))
```
または
```go
carbon.CreateFromStdTime(time.Now())
carbon.CreateFromStdTime(time.Now().In(time.Local))
```

## `carbon` を `time.Time` に変換

```go
carbon.Now().StdTime()
carbon.Now(carbon.Local).StdTime()
```