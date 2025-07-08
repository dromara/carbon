---
head:
  - - meta
    - name: description
      content: 时间解析|一个轻量级、语义化、对开发者友好的 golang 时间处理库
  - - meta
    - name: keywords
      content: carbon, go-carbon, 默认模板, 布局模板, 格式模板
---

# 时间解析
该系列方法不支持 `时间戳` 字符串解析，解析时间戳请使用 `CreateFromTimestamp`, `CreateFromTimestampMilli`, `CreateFromTimestampMicro`, `CreateFromTimestampNano` 方法

## 通过默认 `布局模板` 解析
通过默认 `布局模板` 将时间字符串解析成 `Carbon` 实例
```go
carbon.Parse("").ToDateTimeString() // 空字符串
carbon.Parse("00:00:00").ToDateTimeString() // 空字符串
carbon.Parse("0000-00-00").ToDateTimeString() // 空字符串
carbon.Parse("0000-00-00 00:00:00").ToDateTimeString() // 空字符串

carbon.Parse("now").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("yesterday").ToString() // 2020-08-04 13:14:15 +0000 UTC
carbon.Parse("tomorrow").ToString() // 2020-08-06 13:14:15 +0000 UTC

carbon.Parse("2020").ToString() // 2020-01-01 00:00:00 +0000 UTC
carbon.Parse("2020-8").ToString() // 2020-08-01 00:00:00 +0000 UTC
carbon.Parse("2020-08").ToString() // 2020-08-01 00:00:00 +0000 UTC
carbon.Parse("2020-8-5").ToString() // 2020-08-05 00:00:00 +0000 UTC
carbon.Parse("2020-8-05").ToString() // 2020-08-05 00:00:00 +0000 UTC
carbon.Parse("2020-08-05").ToString() // 2020-08-05 00:00:00 +0000 UTC
carbon.Parse("2020-08-05.999").ToString() // 2020-08-05 00:00:00.999 +0000 UTC
carbon.Parse("2020-08-05.999999").ToString() // 2020-08-05 00:00:00.999999 +0000 UTC
carbon.Parse("2020-08-05.999999999").ToString() // 2020-08-05 00:00:00.999999999 +0000 UTC

carbon.Parse("2020-8-5 13:14:15").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-8-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-08-5 13:14:15").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-08-05 13:14:15.999").ToString() // 2020-08-05 13:14:15.999 +0000 UTC
carbon.Parse("2020-08-05 13:14:15.999999").ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
carbon.Parse("2020-08-05 13:14:15.999999999").ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC

carbon.Parse("2020-8-5T13:14:15+08:00").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-8-05T13:14:15+08:00").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-08-05T13:14:15+08:00").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-08-05T13:14:15.999+08:00").ToString() // 2020-08-05 13:14:15.999 +0000 UTC
carbon.Parse("2020-08-05T13:14:15.999999+08:00").ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC

carbon.Parse("20200805").ToString() // 2020-08-05 00:00:00 +0000 UTC
carbon.Parse("20200805131415").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("20200805131415.999").ToString() // 2020-08-05 13:14:15.999 +0000 UTC
carbon.Parse("20200805131415.999999").ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
carbon.Parse("20200805131415.999999999").ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC
carbon.Parse("20200805131415.999+08:00").ToString() // 2020-08-05 13:14:15.999 +0000 UTC
carbon.Parse("20200805131415.999999+08:00").ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
carbon.Parse("20200805131415.999999999+08:00").ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC

carbon.Parse("2022-03-08T03:01:14-07:00").ToString() // 2022-03-08 18:01:14 +0000 UTC
carbon.Parse("2022-03-08T10:01:14Z").ToString() // 2022-03-08 18:01:14 +0000 UTC
```

## 通过一个 `布局模板` 解析
通过一个确认的 `布局模板` 将时间字符串解析成 `Carbon` 实例
```go
carbon.ParseByLayout("2020|08|05 13|14|15", "2006|01|02 15|04|05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("2020-08-05 13:14:15", "2006-01-02 15:04:05", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

## 通过一个 `格式模板` 解析
通过一个确认 `格式模板` 将时间字符串解析成 `Carbon` 实例
> 注：如果使用的字母与格式模板冲突时，请使用转义符 "\\" 转义该字母
```go
carbon.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("2020-08-05 13:14:15", "Y-m-d H:i:s", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

## 通过多个 `布局模板` 解析
通过多个模糊的 `布局模板` 将时间字符串解析成 `Carbon` 实例
```go
c := carbon.ParseByLayouts("2020|08|05 13|14|15", []string{"2006|01|02 15|04|05", "2006|1|2 3|4|5"})
c.ToDateTimeString() // 2020-08-05 13:14:15
c.CurrentLayout() // 2006|01|02 15|04|05
```

## 通过多个 `格式模板` 解析
通过多个模糊的 `格式模板` 将时间字符串解析成 `Carbon` 实例
```go
c := carbon.ParseByFormats("2020|08|05 13|14|15", []string{"Y|m|d H|i|s", "y|m|d h|i|s"})
c.ToDateTimeString() // 2020-08-05 13:14:15
c.CurrentLayout() // 2006|01|02 15|04|05
```