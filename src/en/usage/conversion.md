# Convert between `carbon` and `time.Time`

## Convert standard `time.Time` to `carbon`
```go
carbon.NewCarbon(time.Now())
carbon.NewCarbon(time.Now().In(time.Local))
```
or
```go
carbon.CreateFromStdTime(time.Now())
carbon.CreateFromStdTime(time.Now().In(time.Local))
```

## Convert `carbon` to standard `time.Time`
```go
carbon.Now().StdTime()
carbon.Now(carbon.Local).StdTime()
```