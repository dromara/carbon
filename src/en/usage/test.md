---
head:
  - - meta
    - name: description
      content: Carbon testing feature detailed explanation, provides SetTestNow/IsTestNow/ClearTestNow methods, supports freezing fixed time for unit tests, making time-related operations run based on simulated time, convenient for testing time-dependent code
  - - meta
    - name: keywords
      content: carbon, go-carbon, test time, freeze time, simulated time, SetTestNow, IsTestNow, ClearTestNow, unit test, time testing, testing tool
---

# Testing
Supports freezing a fixed time, setting any time as the `current time`, so that subsequent operations run based on this simulated time rather than the real system time, which is convenient for unit testing

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

## Is test time
```go
carbon.IsTestNow() 
```

## Clear test now
```go
carbon.ClearTestNow()
```