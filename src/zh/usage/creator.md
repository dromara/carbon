---
head:
  - - meta
    - name: description
      content: 创建 Carbon 实例|一个轻量级、语义化、对开发者友好的 golang 时间处理库,
  - - meta
    - name: keywords
      content: 从时间戳创建Carbon实例, 从日期创建Carbon实例
---

# 创建 `Carbon` 实例

## 从时间戳创建 `Carbon` 实例
```go
// 从秒精度时间戳创建 Carbon 实例
carbon.CreateFromTimestamp(-1).ToString() // 1970-01-01 07:59:59 +0800 CST
carbon.CreateFromTimestamp(0).ToString() // 1970-01-01 08:00:00 +0800 CST
carbon.CreateFromTimestamp(1).ToString() // 1970-01-01 08:00:01 +0800 CST
carbon.CreateFromTimestamp(1596604455).ToString() // 2020-08-05 13:14:15 +0800 CST
// 从毫秒精度时间戳创建 Carbon 实例
carbon.CreateFromTimestampMilli(1596604455999).ToString() // 2020-08-05 13:14:15.999 +0800 CST
// 从微秒精度时间戳创建 Carbon 实例
carbon.CreateFromTimestampMicro(1596604455999999).ToString() // 2020-08-05 13:14:15.999999 +0800 CST
// 从纳秒精度时间戳创建 Carbon 实例
carbon.CreateFromTimestampNano(1596604455999999999).ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
```

## 从日期创建 `Carbon` 实例
```go
// 从年、月、日、时、分、秒创建 Carbon 实例
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToString() // 2020-08-05 13:14:15 +0800 CST
// 从年、月、日、时、分、秒、毫秒创建 Carbon 实例
carbon.CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString() // 2020-08-05 13:14:15.999 +0800 CST
// 从年、月、日、时、分、秒、微秒创建 Carbon 实例
carbon.CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString() // 2020-08-05 13:14:15.999999 +0800 CST
// 从年、月、日、时、分、秒、纳秒创建 Carbon 实例
carbon.CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

// 从年、月、日创建 Carbon 实例
carbon.CreateFromDate(2020, 8, 5).ToString() // 2020-08-05 00:00:00 +0800 CST
// 从年、月、日、毫秒创建 Carbon 实例
carbon.CreateFromDateMilli(2020, 8, 5, 999).ToString() // 2020-08-05 00:00:00.999 +0800 CST
// 从年、月、日、微秒创建 Carbon 实例
carbon.CreateFromDateMicro(2020, 8, 5, 999999).ToString() // 2020-08-05 00:00:00.999999 +0800 CST
// 从年、月、日、纳秒创建 Carbon 实例
carbon.CreateFromDateNano(2020, 8, 5, 999999999).ToString() // 2020-08-05 00:00:00.999999999 +0800 CST

// 从时、分、秒创建 Carbon 实例(年月日默认为当前年月日)
carbon.CreateFromTime(13, 14, 15).ToString() // 2020-08-05 13:14:15 +0800 CST
// 从时、分、秒、毫秒创建 Carbon 实例(年月日默认为当前年月日)
carbon.CreateFromTimeMilli(13, 14, 15, 999).ToString() // 2020-08-05 13:14:15.999 +0800 CST
// 从时、分、秒、微秒创建 Carbon 实例(年月日默认为当前年月日)
carbon.CreateFromTimeMicro(13, 14, 15, 999999).ToString() // 2020-08-05 13:14:15.999999 +0800 CST
// 从时、分、秒、纳秒创建 Carbon 实例(年月日默认为当前年月日)
carbon.CreateFromTimeNano(13, 14, 15, 999999999).ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
```