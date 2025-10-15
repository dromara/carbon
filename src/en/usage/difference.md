---
head:
  - - meta
    - name: description
      content: Difference | A simple, semantic and developer-friendly time package for golang
  - - meta
    - name: keywords
      content: carbon, go-carbon, relative difference, absolute difference, humanized difference
---

# Difference

Assuming the current time is `2020-08-05 13:14:15.999999999 +0000 UTC`
```go
carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15.999999999 +0000 UTC"))
```

## Relative difference
```go
// Difference in years
carbon.Parse("2021-08-05 13:14:15").DiffInYears(carbon.Parse("2020-08-05 13:14:15")) // -1

// Difference in months
carbon.Parse("2020-08-05 13:14:15").DiffInMonths(carbon.Parse("2020-07-05 13:14:15")) // -1

// Difference in weeks
carbon.Parse("2020-08-05 13:14:15").DiffInWeeks(carbon.Parse("2020-07-28 13:14:15")) // -1

// Difference in days
carbon.Parse("2020-08-05 13:14:15").DiffInDays(carbon.Parse("2020-08-04 13:14:15")) // -1

// Difference in hours
carbon.Parse("2020-08-05 13:14:15").DiffInHours(carbon.Parse("2020-08-05 12:14:15")) // -1

// Difference in minutes
carbon.Parse("2020-08-05 13:14:15").DiffInMinutes(carbon.Parse("2020-08-05 13:13:15")) // -1

// Difference in seconds
carbon.Parse("2020-08-05 13:14:15").DiffInSeconds(carbon.Parse("2020-08-05 13:14:14")) // -1

// Difference in string
carbon.Now().DiffInString() // just now
carbon.Now().AddYearsNoOverflow(1).DiffInString() // -1 year
carbon.Now().SubYearsNoOverflow(1).DiffInString() // 1 year

// Difference in duration
now := carbon.Now()
now.DiffInDuration(now).String() // 0s
now.Copy().AddHour().DiffInDuration(now).String() // 1h0m0s
now.Copy().SubHour().DiffInDuration(now).String() // -1h0m0s
```

## Absolute difference
```go
// Difference in years with absolute value
carbon.Parse("2021-08-05 13:14:15").DiffAbsInYears(carbon.Parse("2020-08-05 13:14:15")) // 1

// Difference in months with absolute value
carbon.Parse("2020-08-05 13:14:15").DiffAbsInMonths(carbon.Parse("2020-07-05 13:14:15")) // 1

// Difference in weeks with absolute value
carbon.Parse("2020-08-05 13:14:15").DiffAbsInWeeks(carbon.Parse("2020-07-28 13:14:15")) // 1

// Difference in days with absolute value
carbon.Parse("2020-08-05 13:14:15").DiffAbsInDays(carbon.Parse("2020-08-04 13:14:15")) // 1

// Difference in hours with absolute value
carbon.Parse("2020-08-05 13:14:15").DiffAbsInHours(carbon.Parse("2020-08-05 12:14:15")) // 1

// Difference in minutes with absolute value
carbon.Parse("2020-08-05 13:14:15").DiffAbsInMinutes(carbon.Parse("2020-08-05 13:13:15")) // 1

// Difference in seconds with absolute value
carbon.Parse("2020-08-05 13:14:15").DiffAbsInSeconds(carbon.Parse("2020-08-05 13:14:14")) // 1

// Difference in string with absolute value
carbon.Now().DiffAbsInString(carbon.Now()) // just now
carbon.Now().AddYearsNoOverflow(1).DiffAbsInString(carbon.Now()) // 1 year
carbon.Now().SubYearsNoOverflow(1).DiffAbsInString(carbon.Now()) // 1 year

// Difference in duration with absolute value
now.DiffAbsInDuration(now).String() // 0s
now.Copy().AddHour().DiffAbsInDuration(now).String() // 1h0m0s
now.Copy().SubHour().DiffAbsInDuration(now).String() // 1h0m0s
```

## Humanized difference
```go
carbon.Parse("2020-08-05 13:14:15").DiffForHumans() // just now
carbon.Parse("2019-08-05 13:14:15").DiffForHumans() // 1 year ago
carbon.Parse("2018-08-05 13:14:15").DiffForHumans() // 2 years ago
carbon.Parse("2021-08-05 13:14:15").DiffForHumans() // 1 year from now
carbon.Parse("2022-08-05 13:14:15").DiffForHumans() // 2 years from now

time := carbon.Parse("2020-08-05 13:14:15")
carbon.Parse("2020-08-05 13:14:15").DiffForHumans(time) // just now
carbon.Parse("2019-08-05 13:14:15").DiffForHumans(time) // 1 year before
carbon.Parse("2018-08-05 13:14:15").DiffForHumans(time) // 2 years before
carbon.Parse("2021-08-05 13:14:15").DiffForHumans(time) // 1 year after
carbon.Parse("2022-08-05 13:14:15").DiffForHumans(time) // 2 years after

carbon.SetLocale("zh-CN")
carbon.Parse("2020-08-05 13:14:15").DiffForHumans() // 刚刚
carbon.Parse("2019-08-05 13:14:15").DiffForHumans() // 1 年前
carbon.Parse("2018-08-05 13:14:15").DiffForHumans() // 2 年前
carbon.Parse("2021-08-05 13:14:15").DiffForHumans() // 1 年后
carbon.Parse("2022-08-05 13:14:15").DiffForHumans() // 2 年后
```