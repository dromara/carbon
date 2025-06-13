# Testing
The testing methods allow you to set a `Carbon` instance to be returned when a `now` instance is created. The provided instance will be used when retrieving any relative time from `Carbon` (now, today, yesterday, next month, etc.)

## Set freeze time as test now
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

## Is test now
```go
carbon.IsTestNow() 
```

## Clear test now
```go
carbon.ClearTestNow()
```