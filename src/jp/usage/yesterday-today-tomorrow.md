# 昨日、今日、明日
昨日、今日、明日は時間旅行の特別な例であり、同等である
```go
carbon.Yesterday() = carbon.Now().SubDay()
carbon.Tomorrow() = carbon.Now().AddDay()
```

## 昨日
```go
// 昨日の今頃
fmt.Printf("%s", carbon.Yesterday()) // 2020-08-04 13:14:15
carbon.Yesterday().String() // 2020-08-04 13:14:15
carbon.Yesterday().ToString() // 2020-08-04 13:14:15.999999999 +0800 CST
carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// 昨日の日付
carbon.Yesterday().ToDateString() // 2020-08-04
// 昨日の時間
carbon.Yesterday().ToTimeString() // 13:14:15
// 指定されたタイムゾーンの昨日の今
carbon.Yesterday(carbon.NewYork).ToDateTimeString() // 2020-08-04 14:14:15
// 昨日の秒精度タイムスタンプ
carbon.Yesterday().Timestamp() // 1596518055
// 昨日のミリ秒精度タイムスタンプ
carbon.Yesterday().TimestampMilli() // 1596518055999
// 昨日のマイクロ秒精度タイムスタンプ
carbon.Yesterday().TimestampMicro() // 1596518055999999
// 昨日のナノ秒精度タイムスタンプ
carbon.Yesterday().TimestampNano() // 1596518055999999999
```

## 今日
```go
// 今日は今
fmt.Printf("%s", carbon.Now()) // 2020-08-05 13:14:15
carbon.Now().String() // 2020-08-05 13:14:15
carbon.Now().ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
// 今日の日付
carbon.Now().ToDateString() // 2020-08-05
// 今日の時間
carbon.Now().ToTimeString() // 13:14:15
// 指定されたタイムゾーンの今
carbon.Now(carbon.NewYork).ToDateTimeString() // 2020-08-05 14:14:15
// 今日の秒精度タイムスタンプ
carbon.Now().Timestamp() // 1596604455
// 今日のミリ秒精度タイムスタンプ
carbon.Now().TimestampMilli() // 1596604455999
// 今日のマイクロ秒精度タイムスタンプ
carbon.Now().TimestampMicro() // 1596604455999999
// 今日のナノ秒精度タイムスタンプ
carbon.Now().TimestampNano() // 1596604455999999999
```

## 明日
```go
// 明天此刻
fmt.Printf("%s", carbon.Tomorrow()) // 2020-08-06 13:14:15
carbon.Tomorrow().String() // 2020-08-06 13:14:15
carbon.Tomorrow().ToString() // 2020-08-06 13:14:15.999999999 +0800 CST
carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// 明天日期
carbon.Tomorrow().ToDateString() // 2020-08-06
// 明天时间
carbon.Tomorrow().ToTimeString() // 13:14:15
// 指定时区的明天此刻
carbon.Tomorrow(carbon.NewYork).ToDateTimeString() // 2020-08-06 14:14:15
// 明天秒精度时间戳
carbon.Tomorrow().Timestamp() // 1596690855
// 明天毫秒精度时间戳
carbon.Tomorrow().TimestampMilli() //1596690855999
// 明天微秒精度时间戳
carbon.Tomorrow().TimestampMicro() // 1596690855999999
// 明天纳秒精度时间戳
carbon.Tomorrow().TimestampNano() // 1596690855999999999
```