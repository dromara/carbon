# 时间解析
この一連のメソッドは `タイムスタンプ` 文字列の解析をサポートしていません。タイムスタンプを解析するには、`CreateFromTimestamp` や `CreateFromTimestampXXX` などのメソッドを使用してください

## デフォルトの `レイアウトテンプレート` による解析
デフォルトの `レイアウトテンプレート` を使用して時間文字列を `Carbon` インスタンスに解析する
```go
carbon.Parse("").ToDateTimeString() // 空の文字列
carbon.Parse("00:00:00").ToDateTimeString() // 空の文字列
carbon.Parse("0000-00-00").ToDateTimeString() // 空の文字列
carbon.Parse("0000-00-00 00:00:00").ToDateTimeString() // 空の文字列

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

carbon.Parse("2022-03-08T03:01:14-07:00").ToString() // 2022-03-08 19:01:14 +0000 UTC
carbon.Parse("2022-03-08T10:01:14Z").ToString() // 2022-03-08 19:01:14 +0000 UTC
```

## 1つの `レイアウトテンプレート` による解析
確認された `レイアウトテンプレート` によって時間文字列を `Carbon` インスタンスに解析する
```go
carbon.ParseByLayout("2020|08|05 13|14|15", "2006|01|02 15|04|05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("2020-08-05 13:14:15", "2006-01-02 15:04:05", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

## 1つの `フォーマットテンプレート` による解析
時間文字列を確認 `フォーマットテンプレート` によって `Carbon` インスタンスに解析する
> 注：使用している文字がフォームテンプレートと競合している場合は、エスケープ文字 "\\" を使用して文字をエスケープします
```go
carbon.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("2020-08-05 13:14:15", "Y-m-d H:i:s", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

## 複数の `レイアウトテンプレート` による解析
複数のファジィな `レイアウトテンプレート `によって時間文字列を `Carbon` インスタンスに解析する
```go
c := carbon.ParseByLayouts("2020|08|05 13|14|15", []string{"2006|01|02 15|04|05", "2006|1|2 3|4|5"})
c.ToDateTimeString() // 2020-08-05 13:14:15
c.CurrentLayout() // 2006|01|02 15|04|05
```

## 複数の `フォーマットテンプレート` による解析
複数のファジィな `フォーマットテンプレート`` によって時間文字列を `Carbon` インスタンスに解析する
```go
c := carbon.ParseByFormats("2020|08|05 13|14|15", []string{"Y|m|d H|i|s", "y|m|d h|i|s"})
c.ToDateTimeString() // 2020-08-05 13:14:15
c.CurrentLayout() // 2006|01|02 15|04|05
```