---
head:
  - - meta
    - name: description
      content: Testing | A simple, semantic and developer-friendly time package for golang
  - - meta
    - name: keywords
      content: carbon, go-carbon, test time, freeze time, SetTestNow, IsTestNow, ClearTestNow
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