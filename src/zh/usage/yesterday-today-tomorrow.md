---
head:
  - - meta
    - name: description
      content: "Carbon 昨天今天明天功能详解，提供 Yesterday/Now/Tomorrow 三个便捷方法获取相对时间，支持日期时间格式化、多精度时间戳（秒/毫秒/微秒/纳秒）和指定时区，等价于 Now().SubDay() 和 Now().AddDay()"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, 昨天, 今天, 明天, Yesterday, Now, Tomorrow, 相对时间, 时间戳, 时区设置, 日期格式化"
---

# 昨天、今天、明天
昨天是今天的前一天，明天是今天的后一天，等价于

```go
carbon.Yesterday() = carbon.Now().SubDay()
carbon.Tomorrow() = carbon.Now().AddDay()
```
假设当前时间为 `2020-08-05 13:14:15.999999999 +0000 UTC`

## 昨天
```go
// 昨天此刻
fmt.Printf("%s", carbon.Yesterday()) // 2020-08-04 13:14:15
carbon.Yesterday().String() // 2020-08-04 13:14:15
carbon.Yesterday().ToString() // 2020-08-04 13:14:15.999999999 +0000 UTC
carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// 昨天日期
carbon.Yesterday().ToDateString() // 2020-08-04
// 昨天时间
carbon.Yesterday().ToTimeString() // 13:14:15
// 指定时区的昨天此刻
carbon.Yesterday(carbon.NewYork).ToDateTimeString() // 2020-08-04 14:14:15
// 昨天秒精度时间戳
carbon.Yesterday().Timestamp() // 1596546855
// 昨天毫秒精度时间戳
carbon.Yesterday().TimestampMilli() // 1596546855999
// 昨天微秒精度时间戳
carbon.Yesterday().TimestampMicro() // 1596546855999999
// 昨天纳秒精度时间戳
carbon.Yesterday().TimestampNano() // 1596546855999999999
```

## 今天
```go
// 今天此刻
fmt.Printf("%s", carbon.Now()) // 2020-08-05 13:14:15
carbon.Now().String() // 2020-08-05 13:14:15
carbon.Now().ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC
carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
// 今天日期
carbon.Now().ToDateString() // 2020-08-05
// 今天时间
carbon.Now().ToTimeString() // 13:14:15
// 指定时区的今天此刻
carbon.Now(carbon.NewYork).ToDateTimeString() // 2020-08-05 14:14:15
// 今天秒精度时间戳
carbon.Now().Timestamp() // 1596633255
// 今天毫秒精度时间戳
carbon.Now().TimestampMilli() // 1596633255999
// 今天微秒精度时间戳
carbon.Now().TimestampMicro() // 1596633255999999
// 今天纳秒精度时间戳
carbon.Now().TimestampNano() // 1596633255999999999
```

## 明天
```go
// 明天此刻
fmt.Printf("%s", carbon.Tomorrow()) // 2020-08-06 13:14:15
carbon.Tomorrow().String() // 2020-08-06 13:14:15
carbon.Tomorrow().ToString() // 2020-08-06 13:14:15.999999999 +0000 UTC
carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// 明天日期
carbon.Tomorrow().ToDateString() // 2020-08-06
// 明天时间
carbon.Tomorrow().ToTimeString() // 13:14:15
// 指定时区的明天此刻
carbon.Tomorrow(carbon.NewYork).ToDateTimeString() // 2020-08-06 14:14:15
// 明天秒精度时间戳
carbon.Tomorrow().Timestamp() // 1596719655
// 明天毫秒精度时间戳
carbon.Tomorrow().TimestampMilli() // 1596719655999
// 明天微秒精度时间戳
carbon.Tomorrow().TimestampMicro() // 1596719655999999
// 明天纳秒精度时间戳
carbon.Tomorrow().TimestampNano() // 1596719655999999999
```