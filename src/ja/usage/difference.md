---
head:
  - - meta
    - name: description
      content: 時間差 | DiffInYears/Months/Weeks/Days/Hours/Minutes/Seconds の相対差、DiffAbsInXXX の絶対差、DiffForHumans のヒューマナイズ差を提供し、DiffInString/Duration をサポート
  - - meta
    - name: keywords
      content: carbon, go-carbon, 相対差, 絶対差, ヒューマナイズ差
---

# 時間差

現在時刻が `2020-08-05 13:14:15.999999999 +0000 UTC` であると仮定します
```go
carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15.999999999 +0000 UTC"))
```

## 相対差
```go
carbon.Parse("2021-08-05 13:14:15").DiffInYears(carbon.Parse("2020-08-05 13:14:15")) // -1
carbon.Parse("2020-08-05 13:14:15").DiffInMonths(carbon.Parse("2020-07-05 13:14:15")) // -1
carbon.Parse("2020-08-05 13:14:15").DiffInWeeks(carbon.Parse("2020-07-28 13:14:15")) // -1
carbon.Parse("2020-08-05 13:14:15").DiffInDays(carbon.Parse("2020-08-04 13:14:15")) // -1
carbon.Parse("2020-08-05 13:14:15").DiffInHours(carbon.Parse("2020-08-05 12:14:15")) // -1
carbon.Parse("2020-08-05 13:14:15").DiffInMinutes(carbon.Parse("2020-08-05 13:13:15")) // -1
carbon.Parse("2020-08-05 13:14:15").DiffInSeconds(carbon.Parse("2020-08-05 13:14:14")) // -1

carbon.Now().DiffInString() // just now
carbon.Now().AddYearsNoOverflow(1).DiffInString() // -1 year
carbon.Now().SubYearsNoOverflow(1).DiffInString() // 1 year

now := carbon.Now()
now.DiffInDuration(now).String() // 0s
now.Copy().AddHour().DiffInDuration(now).String() // 1h0m0s
now.Copy().SubHour().DiffInDuration(now).String() // -1h0m0s
```

## 絶対差
```go
carbon.Parse("2021-08-05 13:14:15").DiffAbsInYears(carbon.Parse("2020-08-05 13:14:15")) // 1
carbon.Parse("2020-08-05 13:14:15").DiffAbsInMonths(carbon.Parse("2020-07-05 13:14:15")) // 1
carbon.Parse("2020-08-05 13:14:15").DiffAbsInWeeks(carbon.Parse("2020-07-28 13:14:15")) // 1
carbon.Parse("2020-08-05 13:14:15").DiffAbsInDays(carbon.Parse("2020-08-04 13:14:15")) // 1
carbon.Parse("2020-08-05 13:14:15").DiffAbsInHours(carbon.Parse("2020-08-05 12:14:15")) // 1
carbon.Parse("2020-08-05 13:14:15").DiffAbsInMinutes(carbon.Parse("2020-08-05 13:13:15")) // 1
carbon.Parse("2020-08-05 13:14:15").DiffAbsInSeconds(carbon.Parse("2020-08-05 13:14:14")) // 1

carbon.Now().DiffAbsInString(carbon.Now()) // just now
carbon.Now().AddYearsNoOverflow(1).DiffAbsInString(carbon.Now()) // 1 year
carbon.Now().SubYearsNoOverflow(1).DiffAbsInString(carbon.Now()) // 1 year

now.DiffAbsInDuration(now).String() // 0s
now.Copy().AddHour().DiffAbsInDuration(now).String() // 1h0m0s
now.Copy().SubHour().DiffAbsInDuration(now).String() // 1h0m0s
```

## ヒューマナイズ差
```go
carbon.Parse("2020-08-05 13:14:15").DiffForHumans() // 現在
carbon.Parse("2019-08-05 13:14:15").DiffForHumans() // 1 年前
carbon.Parse("2018-08-05 13:14:15").DiffForHumans() // 2 年前
carbon.Parse("2021-08-05 13:14:15").DiffForHumans() // 1 年後
carbon.Parse("2022-08-05 13:14:15").DiffForHumans() // 2 年後

carbon.SetLocale("zh-CN")
carbon.Parse("2020-08-05 13:14:15").DiffForHumans() // 刚刚
carbon.Parse("2019-08-05 13:14:15").DiffForHumans() // 1 年前
carbon.Parse("2018-08-05 13:14:15").DiffForHumans() // 2 年前
carbon.Parse("2021-08-05 13:14:15").DiffForHumans() // 1 年後
carbon.Parse("2022-08-05 13:14:15").DiffForHumans() // 2 年後
```