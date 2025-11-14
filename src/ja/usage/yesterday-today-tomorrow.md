---
head:
  - - meta
    - name: description
      content: "Carbon 昨日今日明日機能の詳細説明、Yesterday/Now/Tomorrow の3つの便利なメソッドを提供して相対時間を取得、日時フォーマット、多精度タイムスタンプ（秒/ミリ秒/マイクロ秒/ナノ秒）と指定タイムゾーンをサポート、Now().SubDay() と Now().AddDay() と同等"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, 昨日, 今日, 明日, Yesterday, Now, Tomorrow, 相対時間, タイムスタンプ, タイムゾーン設定, 日付フォーマット"
---

# 昨日、今日、明日
昨日は今日の前日、明日は今日の翌日、同等である

```go
carbon.Yesterday() = carbon.Now().SubDay()
carbon.Tomorrow() = carbon.Now().AddDay()
```
現在時刻が `2020-08-05 13:14:15.999999999 +0000 UTC` であると仮定します

## 昨日
```go
// 昨日の今頃
fmt.Printf("%s", carbon.Yesterday()) // 2020-08-04 13:14:15
carbon.Yesterday().String() // 2020-08-04 13:14:15
carbon.Yesterday().ToString() // 2020-08-04 13:14:15.999999999 +0000 UTC
carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// 昨日の日付
carbon.Yesterday().ToDateString() // 2020-08-04
// 昨日の時間
carbon.Yesterday().ToTimeString() // 13:14:15
// 指定されたタイムゾーンの昨日の今
carbon.Yesterday(carbon.NewYork).ToDateTimeString() // 2020-08-04 14:14:15
// 昨日の秒精度タイムスタンプ
carbon.Yesterday().Timestamp() // 1596546855
// 昨日のミリ秒精度タイムスタンプ
carbon.Yesterday().TimestampMilli() // 1596546855999
// 昨日のマイクロ秒精度タイムスタンプ
carbon.Yesterday().TimestampMicro() // 1596546855999999
// 昨日のナノ秒精度タイムスタンプ
carbon.Yesterday().TimestampNano() // 1596546855999999999
```

## 今日
```go
// 今日の今頃
fmt.Printf("%s", carbon.Now()) // 2020-08-05 13:14:15
carbon.Now().String() // 2020-08-05 13:14:15
carbon.Now().ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC
carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
// 今日の日付
carbon.Now().ToDateString() // 2020-08-05
// 今日の時間
carbon.Now().ToTimeString() // 13:14:15
// 指定されたタイムゾーンの今
carbon.Now(carbon.NewYork).ToDateTimeString() // 2020-08-05 14:14:15
// 今日の秒精度タイムスタンプ
carbon.Now().Timestamp() // 1596633255
// 今日のミリ秒精度タイムスタンプ
carbon.Now().TimestampMilli() // 1596633255999
// 今日のマイクロ秒精度タイムスタンプ
carbon.Now().TimestampMicro() // 1596633255999999
// 今日のナノ秒精度タイムスタンプ
carbon.Now().TimestampNano() // 1596633255999999999
```

## 明日
```go
// 明日の今頃
fmt.Printf("%s", carbon.Tomorrow()) // 2020-08-06 13:14:15
carbon.Tomorrow().String() // 2020-08-06 13:14:15
carbon.Tomorrow().ToString() // 2020-08-06 13:14:15.999999999 +0000 UTC
carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// 明日の日付
carbon.Tomorrow().ToDateString() // 2020-08-06
// 明日の時間
carbon.Tomorrow().ToTimeString() // 13:14:15
// 指定されたタイムゾーンの明日の今頃
carbon.Tomorrow(carbon.NewYork).ToDateTimeString() // 2020-08-06 14:14:15
// 明日の秒精度タイムスタンプ
carbon.Tomorrow().Timestamp() // 1596719655
// 明日のミリ秒精度タイムスタンプ
carbon.Tomorrow().TimestampMilli() // 1596719655999
// 明日のマイクロ秒精度タイムスタンプ
carbon.Tomorrow().TimestampMicro() // 1596719655999999
// 明日のナノ秒精度タイムスタンプ
carbon.Tomorrow().TimestampNano() // 1596719655999999999
```