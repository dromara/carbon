# Carbon

[![Carbon Release](https://img.shields.io/github/release/dromara/carbon.svg)](https://github.com/dromara/carbon/releases)
[![Go Test](https://github.com/dromara/carbon/actions/workflows/test.yml/badge.svg)](https://github.com/dromara/carbon/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/dromara/carbon/v2)](https://goreportcard.com/report/github.com/dromara/carbon/v2)
[![Go Coverage](https://codecov.io/gh/dromara/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/dromara/carbon)
[![Carbon Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/dromara/carbon/v2)
[![License](https://img.shields.io/github/license/dromara/carbon)](https://github.com/dromara/carbon/blob/master/LICENSE)

日本語 | [English](README.md) | [简体中文](README.cn.md)

#### イントロ

軽量、セマンティック、開発者に優しい `golang` 時間処理ライブラリ

Carbon は [awesome-go](https://github.com/avelino/awesome-go#date-and-time "awesome-go") に含まれています

#### リポジトリ

[github.com/dromara/carbon](https://github.com/dromara/carbon "github.com/dromara/carbon")

[gitee.com/dromara/carbon](https://gitee.com/dromara/carbon "gitee.com/dromara/carbon")

[gitcode.com/dromara/carbon](https://gitcode.com/dromara/carbon "gitcode.com/dromara/carbon")

#### インストール

##### go version >= 1.17

```go
// github から使う
go get -u github.com/dromara/carbon/v2
import "github.com/dromara/carbon/v2"

// gitee から使う
go get -u gitee.com/dromara/carbon/v2
import "gitee.com/dromara/carbon/v2"

// gitcode から使う
go get -u gitcode.com/dromara/carbon/v2
import "gitcode.com/dromara/carbon/v2"
```

`Carbon` は [dromara](https://dromara.org/ "dromara") 組織に寄付されたためリポジトリのURLが変更されました。以前のリポジトリ `golang-module/carbon` を使用している場合は`go.mod`で新しいリポジトリURLに変更するか下記コマンドを実行します

```go
go mod edit -replace github.com/golang-module/carbon/v2=github.com/dromara/carbon/v2
```

#### 使い方と例

> 現在時刻が 2020-08-05 13:14:15.999999999 +0800 CST であると仮定します。

##### グローバルのデフォルト値設定

```go
carbon.SetDefault(carbon.Default{
  Layout: carbon.DateTimeLayout,
  Timezone: carbon.Japan,
  WeekStartsAt: carbon.Sunday,
  Locale: "jp", // 値の範囲： langディレクトリ下の翻訳ファイル名でファイル拡張子は含まない
})
```

##### Carbon と time.Time 間の変換

```go
// time.Time を Carbon に変換
carbon.NewCarbon(time.Now()) または carbon.CreateFromStdTime(time.Now())
// Carbon を time.Time に変換
carbon.Now().StdTime()
```

##### 昨日、現在、明日

```go
// 現在時刻
fmt.Printf("%s", carbon.Now()) // 2020-08-05 13:14:15
carbon.Now().String() // 2020-08-05 13:14:15
carbon.Now().ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
// 現在の日付
carbon.Now().ToDateString() // 2020-08-05
// 現在の時間
carbon.Now().ToTimeString() // 13:14:15
// タイムゾーン指定の現在
carbon.Now(Carbon.NewYork).ToDateTimeString() // 2020-08-05 14:14:15
// 現在の秒タイムスタンプ
carbon.Now().Timestamp() // 1596604455
// 現在のミリ秒タイムスタンプ
carbon.Now().TimestampMilli() // 1596604455999
// 現在のマイクロ秒タイムスタンプ
carbon.Now().TimestampMicro() // 1596604455999999
// 現在のナノ秒タイムスタンプ
carbon.Now().TimestampNano() // 1596604455999999999

// 昨日の現在時刻
fmt.Printf("%s", carbon.Yesterday()) // 2020-08-04 13:14:15
carbon.Yesterday().String() // 2020-08-04 13:14:15
carbon.Yesterday().ToString() // 2020-08-04 13:14:15 +0800 CST
carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// 昨日の日付
carbon.Yesterday().ToDateString() // 2020-08-04
// 昨日の時間
carbon.Yesterday().ToTimeString() // 13:14:15
// 日付指定の昨日
carbon.Parse("2021-01-28 13:14:15").Yesterday().ToDateTimeString() // 2021-01-27 13:14:15
// タイムゾーン指定の昨日
carbon.Yesterday(Carbon.NewYork).ToDateTimeString() // 2020-08-04 14:14:15
// 昨日の秒タイムスタンプ
carbon.Yesterday().Timestamp() // 1596518055
// 昨日のミリ秒タイムスタンプ
carbon.Yesterday().TimestampMilli() // 1596518055999
// 昨日のマイクロ秒タイムスタンプ
carbon.Yesterday().TimestampMicro() // 1596518055999999
// 昨日のナノ秒タイムスタンプ
carbon.Yesterday().TimestampNano() // 1596518055999999999

// 明日の現在時刻
fmt.Printf("%s", carbon.Tomorrow()) // 2020-08-06 13:14:15
carbon.Tomorrow().String() // 2020-08-06 13:14:15
carbon.Tomorrow().ToString() // 2020-08-06 13:14:15 +0800 CST
carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// 明日の日付
carbon.Tomorrow().ToDateString() // 2020-08-06
// 明日の時間
carbon.Tomorrow().ToTimeString() // 13:14:15
// 日付指定の明日
carbon.Parse("2021-01-28 13:14:15").Tomorrow().ToDateTimeString() // 2021-01-29 13:14:15
// タイムゾーン指定の明日
carbon.Tomorrow(Carbon.NewYork).ToDateTimeString() // 2020-08-06 14:14:15
// 明日の秒タイムスタンプ
carbon.Tomorrow().Timestamp() // 1596690855
// 明日のミリ秒タイムスタンプ
carbon.Tomorrow().TimestampMilli() // 1596690855999
// 明日のマイクロ秒タイムスタンプ
carbon.Tomorrow().TimestampMicro() // 1596690855999999
// 明日のナノ秒タイムスタンプ
carbon.Tomorrow().TimestampNano() // 1596690855999999999
```

##### Carbon インスタンスを作成する

```go
// 秒タイムスタンプから Carbon インスタンスを作成します
carbon.CreateFromTimestamp(-1).ToString() // 1970-01-01 07:59:59 +0800 CST
carbon.CreateFromTimestamp(0).ToString() // 1970-01-01 08:00:00 +0800 CST
carbon.CreateFromTimestamp(1).ToString() // 1970-01-01 08:00:01 +0800 CST
carbon.CreateFromTimestamp(1649735755).ToString() // 2022-04-12 11:55:55 +0800 CST
// ミリ秒のタイムスタンプから Carbon インスタンスを作成します
carbon.CreateFromTimestampMilli(1649735755981).ToString() // 2022-04-12 11:55:55.981 +0800 CST
// マイクロ秒タイムスタンプから Carbon インスタンスを作成します
carbon.CreateFromTimestampMicro(1649735755981566).ToString() // 2022-04-12 11:55:55.981566 +0800 CST
// ナノタイムスタンプから Carbon インスタンスを作成します
carbon.CreateFromTimestampNano(1649735755981566000).ToString() // 2022-04-12 11:55:55.981566 +0800 CST

// 日付と時間から Carbon インスタンスを作成します
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
// 日付と時間、ミリ秒から Carbon インスタンスを作成します
carbon.CreateFromDateTimeMilli(2020, 1, 1, 13, 14, 15, 999).ToString() // 2020-01-01 13:14:15.999 +0800 CST
// 日付と時間、マイクロ秒から Carbon インスタンスを作成します
carbon.CreateFromDateTimeMicro(2020, 1, 1, 13, 14, 15, 999999).ToString() // 2020-01-01 13:14:15.999999 +0800 CST
// 日付と時間、ナノ秒から Carbon インスタンスを作成します
carbon.CreateFromDateTimeNano(2020, 1, 1, 13, 14, 15, 999999999).ToString() // 2020-01-01 13:14:15.999999999 +0800 CST

// 日付から Carbon インスタンスを作成します
carbon.CreateFromDate(2020, 8, 5).ToString() // 2020-08-05 00:00:00 +0800 CST
// 日付とミリ秒から Carbon インスタンスを作成します
carbon.CreateFromDateMilli(2020, 8, 5, 999).ToString() // 2020-08-05 00:00:00.999 +0800 CST
// 日付とマイクロ秒から Carbon インスタンスを作成します
carbon.CreateFromDateMicro(2020, 8, 5, 999999).ToString() // 2020-08-05 00:00:00.999999 +0800 CST
// 日付とナノ秒から Carbon インスタンスを作成します
carbon.CreateFromDateNano(2020, 8, 5, 999999999).ToString() // 2020-08-05 00:00:00.999999999 +0800 CST

// 時間から Carbon インスタンスを作成します(日付のデフォルトは現在の年月日です)
carbon.CreateFromTime(13, 14, 15).ToString() // 2020-08-05 13:14:15 +0800 CST
// 時間、ミリ秒から Carbon インスタンスを作成します(年月日のデフォルトは現在の年月日です)
carbon.CreateFromTimeMilli(13, 14, 15, 999).ToString() // 2020-08-05 13:14:15.999 +0800 CST
// 時間、マイクロ秒から Carbon インスタンスを作成します(年月日のデフォルトは現在の年月日です)
carbon.CreateFromTimeMicro(13, 14, 15, 999999).ToString() // 2020-08-05 13:14:15.999999 +0800 CST
// 時間、ナノ秒から Carbon インスタンスを作成します(年月日のデフォルトは現在の年月日です)
carbon.CreateFromTimeNano(13, 14, 15, 999999999).ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
```

##### 時間文字列を Carbon インスタンスにパース

```go
carbon.Parse("").ToDateTimeString() // 空の文字列
carbon.Parse("0").ToDateTimeString() // 空の文字列
carbon.Parse("00:00:00").ToDateTimeString() // 空の文字列
carbon.Parse("0000-00-00").ToDateTimeString() // 空の文字列
carbon.Parse("0000-00-00 00:00:00").ToDateTimeString() // 空の文字列

carbon.Parse("now").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("yesterday").ToString() // 2020-08-04 13:14:15 +0800 CST
carbon.Parse("tomorrow").ToString() // 2020-08-06 13:14:15 +0800 CST

carbon.Parse("2020").ToString() // 2020-01-01 00:00:00 +0800 CST
carbon.Parse("2020-8").ToString() // 2020-08-01 00:00:00 +0800 CST
carbon.Parse("2020-08").ToString() // 2020-08-01 00:00:00 +0800 CST
carbon.Parse("2020-8-5").ToString() // 2020-08-05 00:00:00 +0800 CST
carbon.Parse("2020-8-05").ToString() // 2020-08-05 00:00:00 +0800 CST
carbon.Parse("2020-08-05").ToString() // 2020-08-05 00:00:00 +0800 CST
carbon.Parse("2020-08-05.999").ToString() // 2020-08-05 00:00:00.999 +0800 CST
carbon.Parse("2020-08-05.999999").ToString() // 2020-08-05 00:00:00.999999 +0800 CST
carbon.Parse("2020-08-05.999999999").ToString() // 2020-08-05 00:00:00.999999999 +0800 CST

carbon.Parse("2020-8-5 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-8-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-5 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-05 13:14:15.999").ToString() // 2020-08-05 13:14:15.999 +0800 CST
carbon.Parse("2020-08-05 13:14:15.999999").ToString() // 2020-08-05 13:14:15.999999 +0800 CST
carbon.Parse("2020-08-05 13:14:15.999999999").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

carbon.Parse("2020-8-5T13:14:15+08:00").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-8-05T13:14:15+08:00").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-05T13:14:15+08:00").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-05T13:14:15.999+08:00").ToString() // 2020-08-05 13:14:15.999 +0800 CST
carbon.Parse("2020-08-05T13:14:15.999999+08:00").ToString() // 2020-08-05 13:14:15.999999 +0800 CST
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

carbon.Parse("20200805").ToString() // 2020-08-05 00:00:00 +0800 CST
carbon.Parse("20200805131415").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("20200805131415.999").ToString() // 2020-08-05 13:14:15.999 +0800 CST
carbon.Parse("20200805131415.999999").ToString() // 2020-08-05 13:14:15.999999 +0800 CST
carbon.Parse("20200805131415.999999999").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
carbon.Parse("20200805131415.999+08:00").ToString() // 2020-08-05 13:14:15.999 +0800 CST
carbon.Parse("20200805131415.999999+08:00").ToString() // 2020-08-05 13:14:15.999999 +0800 CST
carbon.Parse("20200805131415.999999999+08:00").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
```

##### フォーマットして文字列を Carbon インスタンスにパース

```go
carbon.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToDateTimeString() // 2020-08-05 13:14:15
```

##### レイアウトし、文字列を Carbon インスタンスにパース

```go
carbon.ParseByLayout("2020|08|05 13|14|15", "2006|01|02 15|04|05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("2020-08-05 13:14:15", "2006-01-02 15:04:05", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### 時間凍結

```go
now := carbon.Parse("2020-08-05")
carbon.SetTestNow(now)

carbon.IsTestNow() // true
carbon.Now().ToDateString() // 2020-08-05
carbon.Yesterday().ToDateString() // 2020-08-04
carbon.Tomorrow().ToDateString() // 2020-08-05
carbon.Now().DiffForHumans() // just now
carbon.Yesterday().DiffForHumans() // 1 day ago
carbon.Tomorrow().DiffForHumans() // 1 day from now
carbon.Parse("2020-10-05").DiffForHumans() // 2 months from now
now.DiffForHumans(carbon.Parse("2020-10-05")) // 2 months before

carbon.CleanTestNow()
carbon.IsTestNow() // false
```

##### 境界

```go
// 世紀始め
carbon.Parse("2020-08-05 13:14:15").StartOfCentury().ToDateTimeString() // 2000-01-01 00:00:00
// 世紀終わり
carbon.Parse("2020-08-05 13:14:15").EndOfCentury().ToDateTimeString() // 2999-12-31 23:59:59

// 十年紀始め
carbon.Parse("2020-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
carbon.Parse("2021-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
carbon.Parse("2029-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
// 十年紀終わり
carbon.Parse("2020-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
carbon.Parse("2021-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
carbon.Parse("2029-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59

// 今年始め
carbon.Parse("2020-08-05 13:14:15").StartOfYear().ToDateTimeString() // 2020-01-01 00:00:00
// 今年終わり
carbon.Parse("2020-08-05 13:14:15").EndOfYear().ToDateTimeString() // 2020-12-31 23:59:59

// 四半期始め
carbon.Parse("2020-08-05 13:14:15").StartOfQuarter().ToDateTimeString() // 2020-07-01 00:00:00
// 四半期終わり
carbon.Parse("2020-08-05 13:14:15").EndOfQuarter().ToDateTimeString() // 2020-09-30 23:59:59

// 月始め
carbon.Parse("2020-08-05 13:14:15").StartOfMonth().ToDateTimeString() // 2020-08-01 00:00:00
// 月終わり
carbon.Parse("2020-08-05 13:14:15").EndOfMonth().ToDateTimeString() // 2020-08-31 23:59:59

// 周始め
carbon.Parse("2020-08-05 13:14:15").StartOfWeek().ToDateTimeString() // 2020-08-02 00:00:00
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Sunday).StartOfWeek().ToDateTimeString() // 2020-08-02 00:00:00
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Monday).StartOfWeek().ToDateTimeString() // 2020-08-03 00:00:00
// 周終わり
carbon.Parse("2020-08-05 13:14:15").EndOfWeek().ToDateTimeString() // 2020-08-08 23:59:59
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Sunday).EndOfWeek().ToDateTimeString() // 2020-08-08 23:59:59
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Monday).EndOfWeek().ToDateTimeString() // 2020-08-09 23:59:59

// 日始め
carbon.Parse("2020-08-05 13:14:15").StartOfDay().ToDateTimeString() // 2020-08-05 00:00:00
// 日終わり
carbon.Parse("2020-08-05 13:14:15").EndOfDay().ToDateTimeString() // 2020-08-05 23:59:59

// 時間始め
carbon.Parse("2020-08-05 13:14:15").StartOfHour().ToDateTimeString() // 2020-08-05 13:00:00
// 時間終わり
carbon.Parse("2020-08-05 13:14:15").EndOfHour().ToDateTimeString() // 2020-08-05 13:59:59

// 分始め
carbon.Parse("2020-08-05 13:14:15").StartOfMinute().ToDateTimeString() // 2020-08-05 13:14:00
// 分終わり
carbon.Parse("2020-08-05 13:14:15").EndOfMinute().ToDateTimeString() // 2020-08-05 13:14:59

// 秒始め
carbon.Parse("2020-08-05 13:14:15").StartOfSecond().ToString() // 2020-08-05 13:14:15 +0800 CST
// 秒終わり
carbon.Parse("2020-08-05 13:14:15").EndOfSecond().ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
```

##### 大きな時間更新

```go
// 3世紀追加
carbon.Parse("2020-02-29 13:14:15").AddCenturies(3).ToDateTimeString() // 2320-02-29 13:14:15
// 3世紀追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddCenturiesNoOverflow(3).ToDateTimeString() // 2320-02-29 13:14:15
// 1世紀追加
carbon.Parse("2020-02-29 13:14:15").AddCentury().ToDateTimeString() // 2120-02-29 13:14:15
// 1世紀追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddCenturyNoOverflow().ToDateTimeString() // 2120-02-29 13:14:15
// 3世紀減らす
carbon.Parse("2020-02-29 13:14:15").SubCenturies(3).ToDateTimeString() // 1720-02-29 13:14:15
// 3世紀減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubCenturiesNoOverflow(3).ToDateTimeString() // 1720-02-29 13:14:15
// 1世紀減らす
carbon.Parse("2020-02-29 13:14:15").SubCentury().ToDateTimeString() // 1920-02-29 13:14:15
// 1世紀減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubCenturyNoOverflow().ToDateTimeString() // 1920-02-29 13:14:15

// 30年紀追加
carbon.Parse("2020-02-29 13:14:15").AddDecades(3).ToDateTimeString() // 2050-03-01 13:14:15
// 30年紀追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddDecadesNoOverflow(3).ToDateTimeString() // 2050-02-28 13:14:15
// 10年紀追加
carbon.Parse("2020-02-29 13:14:15").AddDecade().ToDateTimeString() // 2030-03-01 13:14:15
// 10年紀追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddDecadeNoOverflow().ToDateTimeString() // 2030-02-28 13:14:15
// 30年紀減らす
carbon.Parse("2020-02-29 13:14:15").SubDecades(3).ToDateTimeString() // 1990-03-01 13:14:15
// 30年紀減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubDecadesNoOverflow(3).ToDateTimeString() // 1990-02-28 13:14:15
// 10年紀減らす
carbon.Parse("2020-02-29 13:14:15").SubDecade().ToDateTimeString() // 2010-03-01 13:14:15
// 10年紀減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubDecadeNoOverflow().ToDateTimeString() // 2010-02-28 13:14:15

// 3年追加
carbon.Parse("2020-02-29 13:14:15").AddYears(3).ToDateTimeString() // 2023-03-01 13:14:15
// 3年追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddYearsNoOverflow(3).ToDateTimeString() // 2023-02-28 13:14:15
// 1年追加
carbon.Parse("2020-02-29 13:14:15").AddYear().ToDateTimeString() // 2021-03-01 13:14:15
// 1年追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddYearNoOverflow().ToDateTimeString() // 2021-02-28 13:14:15
// 3年減らす
carbon.Parse("2020-02-29 13:14:15").SubYears(3).ToDateTimeString() // 2017-03-01 13:14:15
// 3年減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubYearsNoOverflow(3).ToDateTimeString() // 2017-02-28 13:14:15
// 1年減らす
carbon.Parse("2020-02-29 13:14:15").SubYear().ToDateTimeString() // 2019-03-01 13:14:15
// 1年減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubYearNoOverflow().ToDateTimeString() // 2019-02-28 13:14:15

// 3四半期追加
carbon.Parse("2019-05-31 13:14:15").AddQuarters(3).ToDateTimeString() // 2020-03-02 13:14:15
// 3四半期追加(オーバーフローなし)
carbon.Parse("2019-05-31 13:14:15").AddQuartersNoOverflow(3).ToDateTimeString() // 2020-02-29 13:14:15
// 1四半期追加
carbon.Parse("2019-11-30 13:14:15").AddQuarter().ToDateTimeString() // 2020-03-01 13:14:15
// 1四半期追加(オーバーフローなし)
carbon.Parse("2019-11-30 13:14:15").AddQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
// 3四半期減らす
carbon.Parse("2019-08-31 13:14:15").SubQuarters(3).ToDateTimeString() // 2019-03-03 13:14:15
// 3四半期減らす(オーバーフローなし)
carbon.Parse("2019-08-31 13:14:15").SubQuartersNoOverflow(3).ToDateTimeString() // 2019-02-28 13:14:15
// 1四半期減らす
carbon.Parse("2020-05-31 13:14:15").SubQuarter().ToDateTimeString() // 2020-03-02 13:14:15
// 1四半期減らす(オーバーフローなし)
carbon.Parse("2020-05-31 13:14:15").SubQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// 3ヶ月追加
carbon.Parse("2020-02-29 13:14:15").AddMonths(3).ToDateTimeString() // 2020-05-29 13:14:15
// 3ヶ月追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddMonthsNoOverflow(3).ToDateTimeString() // 2020-05-29 13:14:15
// 1ヶ月追加
carbon.Parse("2020-01-31 13:14:15").AddMonth().ToDateTimeString() // 2020-03-02 13:14:15
// 1ヶ月追加(オーバーフローなし)
carbon.Parse("2020-01-31 13:14:15").AddMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
// 3ヶ月減らす
carbon.Parse("2020-02-29 13:14:15").SubMonths(3).ToDateTimeString() // 2019-11-29 13:14:15
// 3ヶ月減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubMonthsNoOverflow(3).ToDateTimeString() // 2019-11-29 13:14:15
// 1ヶ月減らす
carbon.Parse("2020-03-31 13:14:15").SubMonth().ToDateTimeString() // 2020-03-02 13:14:15
// 1か月減らす(オーバーフローなし)
carbon.Parse("2020-03-31 13:14:15").SubMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// 3週間追加
carbon.Parse("2020-02-29 13:14:15").AddWeeks(3).ToDateTimeString() // 2020-03-21 13:14:15
// 1週間追加
carbon.Parse("2020-02-29 13:14:15").AddWeek().ToDateTimeString() // 2020-03-07 13:14:15
// 3週間減らす
carbon.Parse("2020-02-29 13:14:15").SubWeeks(3).ToDateTimeString() // 2020-02-08 13:14:15
// 1週間減らす
carbon.Parse("2020-02-29 13:14:15").SubWeek().ToDateTimeString() // 2020-02-22 13:14:15

// 3日追加
carbon.Parse("2020-08-05 13:14:15").AddDays(3).ToDateTimeString() // 2020-08-08 13:14:15
// 1日追加
carbon.Parse("2020-08-05 13:14:15").AddDay().ToDateTimeString() // 2020-08-05 13:14:15
// 3日減らす
carbon.Parse("2020-08-05 13:14:15").SubDays(3).ToDateTimeString() // 2020-08-02 13:14:15
// 1日減らす
carbon.Parse("2020-08-05 13:14:15").SubDay().ToDateTimeString() // 2020-08-04 13:14:15

// 3時間追加
carbon.Parse("2020-08-05 13:14:15").AddHours(3).ToDateTimeString() // 2020-08-05 16:14:15
// 2時間半追加
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5h").ToDateTimeString() // 2020-08-05 15:44:15
carbon.Parse("2020-08-05 13:14:15").AddDuration("2h30m").ToDateTimeString() // 2020-08-05 15:44:15
// 1時間追加
carbon.Parse("2020-08-05 13:14:15").AddHour().ToDateTimeString() // 2020-08-05 14:14:15
// 3時間減らす
carbon.Parse("2020-08-05 13:14:15").SubHours(3).ToDateTimeString() // 2020-08-05 10:14:15
// 2時間半減らす
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5h").ToDateTimeString() // 2020-08-05 10:44:15
carbon.Parse("2020-08-05 13:14:15").SubDuration("2h30m").ToDateTimeString() // 2020-08-05 10:44:15
// 1時間減らす
carbon.Parse("2020-08-05 13:14:15").SubHour().ToDateTimeString() // 2020-08-05 12:14:15

// 3分追加
carbon.Parse("2020-08-05 13:14:15").AddMinutes(3).ToDateTimeString() // 2020-08-05 13:17:15
// 2分半追加
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5m").ToDateTimeString() // 2020-08-05 13:16:45
carbon.Parse("2020-08-05 13:14:15").AddDuration("2m30s").ToDateTimeString() // 2020-08-05 13:16:45
// 1分追加
carbon.Parse("2020-08-05 13:14:15").AddMinute().ToDateTimeString() // 2020-08-05 13:15:15
// 3分減らす
carbon.Parse("2020-08-05 13:14:15").SubMinutes(3).ToDateTimeString() // 2020-08-05 13:11:15
// 2分半減らす
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5m").ToDateTimeString() // 2020-08-05 13:11:45
carbon.Parse("2020-08-05 13:14:15").SubDuration("2m30s").ToDateTimeString() // 2020-08-05 13:11:45
// 1分減らす
carbon.Parse("2020-08-05 13:14:15").SubMinute().ToDateTimeString() // 2020-08-05 13:13:15

// 3秒追加
carbon.Parse("2020-08-05 13:14:15").AddSeconds(3).ToDateTimeString() // 2020-08-05 13:14:18
// 2秒半追加
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:17
// 1秒追加
carbon.Parse("2020-08-05 13:14:15").AddSecond().ToDateTimeString() // 2020-08-05 13:14:16
// 3秒減らす
carbon.Parse("2020-08-05 13:14:15").SubSeconds(3).ToDateTimeString() // 2020-08-05 13:14:12
// 2秒半減らす
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:12
// 1秒減らす
carbon.Parse("2020-08-05 13:14:15").SubSecond().ToDateTimeString() // 2020-08-05 13:14:14

// 3ミリ秒追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddMilliseconds(3).ToString() // 2020-08-05 13:14:15.225222222 +0800 CST
// 1ミリ秒追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddMillisecond().ToString() // 2020-08-05 13:14:15.223222222 +0800 CST
// 3ミリ秒減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubMilliseconds(3).ToString() // 2020-08-05 13:14:15.219222222 +0800 CST
// 1ミリ秒減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubMillisecond().ToString() // 2020-08-05 13:14:15.221222222 +0800 CST

// 3マイクロ秒追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddMicroseconds(3).ToString() // 2020-08-05 13:14:15.222225222 +0800 CST
// １マイクロ秒追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddMicrosecond().ToString() // 2020-08-05 13:14:15.222223222 +0800 CST
// 3マイクロ秒減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubMicroseconds(3).ToString() // 2020-08-05 13:14:15.222219222 +0800 CST
// １マイクロ秒減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubMicrosecond().ToString() // 2020-08-05 13:14:15.222221222 +0800 CST

// 3ナノ秒追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddNanoseconds(3).ToString() // 2020-08-05 13:14:15.222222225 +0800 CST
// 1ナノ秒追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddNanosecond().ToString() // 2020-08-05 13:14:15.222222223 +0800 CST
// 3ナノ秒減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubNanoseconds(3).ToString() // 2020-08-05 13:14:15.222222219 +0800 CST
// 1ナノ秒減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubNanosecond().ToString() // 2020-08-05 13:14:15.222222221 +0800 CST
```

##### 時間差分

```go
// 年差
carbon.Parse("2021-08-05 13:14:15").DiffInYears(carbon.Parse("2020-08-05 13:14:15")) // -1
// 年差（絶対値）
carbon.Parse("2021-08-05 13:14:15").DiffAbsInYears(carbon.Parse("2020-08-05 13:14:15")) // 1

// 月差
carbon.Parse("2020-08-05 13:14:15").DiffInMonths(carbon.Parse("2020-07-05 13:14:15")) // -1
// 月差（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInMonths(carbon.Parse("2020-07-05 13:14:15")) // 1

// 週差
carbon.Parse("2020-08-05 13:14:15").DiffInWeeks(carbon.Parse("2020-07-28 13:14:15")) // -1
// 週差（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInWeeks(carbon.Parse("2020-07-28 13:14:15")) // 1

// 日差
carbon.Parse("2020-08-05 13:14:15").DiffInDays(carbon.Parse("2020-08-04 13:14:15")) // -1
// 日差（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInDays(carbon.Parse("2020-08-04 13:14:15")) // 1

// 時間差
carbon.Parse("2020-08-05 13:14:15").DiffInHours(carbon.Parse("2020-08-05 12:14:15")) // -1
// 時間差（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInHours(carbon.Parse("2020-08-05 12:14:15")) // 1

// 分差
carbon.Parse("2020-08-05 13:14:15").DiffInMinutes(carbon.Parse("2020-08-05 13:13:15")) // -1
// 分差（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInMinutes(carbon.Parse("2020-08-05 13:13:15")) // 1

// 秒差
carbon.Parse("2020-08-05 13:14:15").DiffInSeconds(carbon.Parse("2020-08-05 13:14:14")) // -1
// 秒差（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInSeconds(carbon.Parse("2020-08-05 13:14:14")) // 1

// 時間差を文字列で返す
carbon.Now().DiffInString() // just now
carbon.Now().AddYearsNoOverflow(1).DiffInString() // -1 year
carbon.Now().SubYearsNoOverflow(1).DiffInString() // 1 year
// 時間差を文字列で返す（絶対値）
carbon.Now().DiffAbsInString(carbon.Now()) // just now
carbon.Now().AddYearsNoOverflow(1).DiffAbsInString(carbon.Now()) // 1 year
carbon.Now().SubYearsNoOverflow(1).DiffAbsInString(carbon.Now()) // 1 year

// 継続時間差
now := carbon.Now()
now.DiffInDuration(now).String() // 0s
now.AddHour().DiffInDuration(now).String() // 1h0m0s
now.SubHour().DiffInDuration(now).String() // -1h0m0s
// 継続時間差（絶対値）
now.DiffAbsInDuration(now).String() // 0s
now.AddHour().DiffAbsInDuration(now).String() // 1h0m0s
now.SubHour().DiffAbsInDuration(now).String() // 1h0m0s

// 人が読みやすいフォーマットで時間差を取得
carbon.Parse("2020-08-05 13:14:15").DiffForHumans() // just now
carbon.Parse("2019-08-05 13:14:15").DiffForHumans() // 1 year ago
carbon.Parse("2018-08-05 13:14:15").DiffForHumans() // 2 years ago
carbon.Parse("2021-08-05 13:14:15").DiffForHumans() // 1 year from now
carbon.Parse("2022-08-05 13:14:15").DiffForHumans() // 2 years from now

carbon.Parse("2020-08-05 13:14:15").DiffForHumans(carbon.Now()) // 1 year before
carbon.Parse("2019-08-05 13:14:15").DiffForHumans(carbon.Now()) // 2 years before
carbon.Parse("2018-08-05 13:14:15").DiffForHumans(carbon.Now()) // 1 year after
carbon.Parse("2022-08-05 13:14:15").DiffForHumans(carbon.Now()) // 2 years after
```

##### 極值

```go
c0 := carbon.Parse("2023-04-01")
c1 := carbon.Parse("2023-03-28")
c2 := carbon.Parse("2023-04-16")
// 最近のCarbonインスタンスを返す
c0.Closest(c1, c2) // c1
// 遠いCarbonインスタンスを返す
c0.Farthest(c1, c2) // c2

yesterday := carbon.Yesterday()
today     := carbon.Now()
tomorrow  := carbon.Tomorrow()
// 最大の Carbon インスタンスを返します
carbon.Max(yesterday, today, tomorrow) // tomorrow
// 最小の Carbon インスタンスを返します
carbon.Min(yesterday, today, tomorrow) // yesterday

// Carbonの最大値を戻す
carbon.MaxCarbon().ToString() // 9999-12-31 23:59:59.999999999 +0000 UTC
// Carbonの最小値を戻す
carbon.MinCarbon().ToString() // -9998-01-01 00:00:00 +0000 UTC
```

##### 比較

```go
// 夏時間かどうか
carbon.Parse("").IsDST() // false
carbon.Parse("0").IsDST() // false
carbon.Parse("0000-00-00 00:00:00").IsDST() // false
carbon.Parse("0000-00-00").IsDST() // false
carbon.Parse("00:00:00").IsDST() // false
carbon.Parse("2023-01-01", "Australia/Brisbane").IsDST() // false
carbon.Parse("2023-01-01", "Australia/Sydney").IsDST() // true

// ゼロ値の時間かどうか(0001-01-01 00:00:00 +0000 UTC)
carbon.Parse("").IsZero() // true
carbon.Parse("0").IsZero() // true
carbon.Parse("0000-00-00 00:00:00").IsZero() // true
carbon.Parse("0000-00-00").IsZero() // true
carbon.Parse("00:00:00").IsZero() // true
carbon.Parse("2020-08-05 00:00:00").IsZero() // false
carbon.Parse("2020-08-05").IsZero() // false
carbon.Parse("2020-08-05").SetTimezone("xxx").IsZero() // false

// 有効な時間かどうか
carbon.Parse("").IsValid() // false
carbon.Parse("0").IsValid() // false
carbon.Parse("0000-00-00 00:00:00").IsValid() // false
carbon.Parse("0000-00-00").IsValid() // false
carbon.Parse("00:00:00").IsValid() // false
carbon.Parse("2020-08-05 00:00:00").IsValid() // true
carbon.Parse("2020-08-05").IsValid() // true
carbon.Parse("2020-08-05").SetTimezone("xxx").IsValid() // false

// 無効な時間かどうか
carbon.Parse("").IsInvalid() // true
carbon.Parse("0").IsInvalid() // true
carbon.Parse("0000-00-00 00:00:00").IsInvalid() // true
carbon.Parse("0000-00-00").IsInvalid() // true
carbon.Parse("00:00:00").IsInvalid() // true
carbon.Parse("2020-08-05 00:00:00").IsInvalid() // false
carbon.Parse("2020-08-05").IsInvalid() // false
carbon.Parse("2020-08-05").SetTimezone("xxx").IsInvalid() // true

// 午前かどうか
carbon.Parse("2020-08-05 00:00:00").IsAM() // true
carbon.Parse("2020-08-05 08:00:00").IsAM() // true
carbon.Parse("2020-08-05 12:00:00").IsAM() // false
carbon.Parse("2020-08-05 13:00:00").IsAM() // false
// 午後かどうか
carbon.Parse("2020-08-05 00:00:00").IsPM() // false
carbon.Parse("2020-08-05 08:00:00").IsPM() // false
carbon.Parse("2020-08-05 12:00:00").IsPM() // true
carbon.Parse("2020-08-05 13:00:00").IsPM() // true

// 現在かどうか
carbon.Now().IsNow() // true
// 未来かどうか
carbon.Tomorrow().IsFuture() // true
// 過去かどうか
carbon.Yesterday().IsPast() // true

// 閏年かどうか
carbon.Parse("2020-08-05 13:14:15").IsLeapYear() // true
// ISO8601で定められたLong Yearかどうか
carbon.Parse("2020-08-05 13:14:15").IsLongYear() // true

// 1月かどうか
carbon.Parse("2020-08-05 13:14:15").IsJanuary() // false
// 2月かどうか
carbon.Parse("2020-08-05 13:14:15").IsFebruary() // false
// 3月かどうか
carbon.Parse("2020-08-05 13:14:15").IsMarch() // false
// 4月かどうか
carbon.Parse("2020-08-05 13:14:15").IsApril()  // false
// 5月かどうか
carbon.Parse("2020-08-05 13:14:15").IsMay() // false
// 6月かどうか
carbon.Parse("2020-08-05 13:14:15").IsJune() // false
// 7月かどうか
carbon.Parse("2020-08-05 13:14:15").IsJuly() // false
// 8月かどうか
carbon.Parse("2020-08-05 13:14:15").IsAugust() // false
// 9月かどうか
carbon.Parse("2020-08-05 13:14:15").IsSeptember() // true
// 10月かどうか
carbon.Parse("2020-08-05 13:14:15").IsOctober() // false
// 11月かどうか
carbon.Parse("2020-08-05 13:14:15").IsNovember() // false
// 12月かどうか
carbon.Parse("2020-08-05 13:14:15").IsDecember() // false

// 月曜日かどうか
carbon.Parse("2020-08-05 13:14:15").IsMonday() // false
// 火曜日かどうか
carbon.Parse("2020-08-05 13:14:15").IsTuesday() // true
// 水曜日かどうか
carbon.Parse("2020-08-05 13:14:15").IsWednesday() // false
// 木曜日かどうか
carbon.Parse("2020-08-05 13:14:15").IsThursday() // false
// 金曜日かどうか
carbon.Parse("2020-08-05 13:14:15").IsFriday() // false
// 土曜日かどうか
carbon.Parse("2020-08-05 13:14:15").IsSaturday() // false
// 日曜日かどうか
carbon.Parse("2020-08-05 13:14:15").IsSunday() // false

// 平日かどうか
carbon.Parse("2020-08-05 13:14:15").IsWeekday() // false
// 週末かどうか
carbon.Parse("2020-08-05 13:14:15").IsWeekend() // true

// 昨日かどうか
carbon.Parse("2020-08-04 13:14:15").IsYesterday() // true
carbon.Parse("2020-08-04 00:00:00").IsYesterday() // true
carbon.Parse("2020-08-04").IsYesterday() // true
// 今日かどうか
carbon.Parse("2020-08-05 13:14:15").IsToday() // true
carbon.Parse("2020-08-05 00:00:00").IsToday() // true
carbon.Parse("2020-08-05").IsToday() // true
// 明日かどうか
carbon.Parse("2020-08-06 13:14:15").IsTomorrow() // true
carbon.Parse("2020-08-06 00:00:00").IsTomorrow() // true
carbon.Parse("2020-08-06").IsTomorrow() // true

// 同世紀かどうか
carbon.Parse("2020-08-05 13:14:15").IsSameCentury(carbon.Parse("3020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameCentury(carbon.Parse("2099-08-05 13:14:15")) // true
// 同十年紀かどうか
carbon.Parse("2020-08-05 13:14:15").IsSameDecade(carbon.Parse("2030-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameDecade(carbon.Parse("2120-08-05 13:14:15")) // true
// 同年かどうか
carbon.Parse("2020-08-05 00:00:00").IsSameYear(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameYear(carbon.Parse("2020-12-31 13:14:15")) // true
// 同四半期かどうか
carbon.Parse("2020-08-05 00:00:00").IsSameQuarter(carbon.Parse("2020-09-05 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameQuarter(carbon.Parse("2021-01-31 13:14:15")) // true
// 同月かどうか
carbon.Parse("2020-01-01 00:00:00").IsSameMonth(carbon.Parse("2021-01-31 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameMonth(carbon.Parse("2020-01-31 13:14:15")) // true
// 同日かどうか
carbon.Parse("2020-08-05 13:14:15").IsSameDay(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 00:00:00").IsSameDay(carbon.Parse("2020-08-05 13:14:15")) // true
// 同時間かどうか
carbon.Parse("2020-08-05 13:14:15").IsSameHour(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:00:00").IsSameHour(carbon.Parse("2020-08-05 13:14:15")) // true
// 同分かどうか
carbon.Parse("2020-08-05 13:14:15").IsSameMinute(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:00").IsSameMinute(carbon.Parse("2020-08-05 13:14:15")) // true
// 同秒かどうか
carbon.Parse("2020-08-05 13:14:15").IsSameSecond(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameSecond(carbon.Parse("2020-08-05 13:14:15")) // true

// 超過かどうか
carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-05 13:14:15")) // false

// 未満かどうか
carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-05 13:14:15")) // false

// 等しいかどうか
carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:00")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:00")) // false

// と等しくないかどうか
carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("!=", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<>", carbon.Parse("2020-08-05 13:14:15")) // false

// 以上かどうか
carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-05 13:14:15")) // true

// 以下かどうか
carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-05 13:14:15")) // true

//　二つの Carbon インスタンスの間に含まれているか(開始時間、終了時間を含まない)
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

//　二つの Carbon インスタンスの間に含まれているか(開始時間を含む)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

//　二つの Carbon インスタンスの間に含まれているか(終了時間を含む)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

//　二つの Carbon インスタンスの間に含まれているか(開始時間、終了時間を含む)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
```

> 長年の定義については、読んでください https://en.wikipedia.org/wiki/ISO_8601#Week_dates

##### 設定

```go
// タイムゾーンを設定
carbon.SetTimezone(carbon.PRC).Now().ToDateTimeString() // 2020-08-05 13:14:15
carbon.SetTimezone(carbon.Tokyo).Now().ToDateTimeString() // 2020-08-05 14:14:15
carbon.SetTimezone(carbon.Tokyo).Now().SetTimezone(carbon.PRC).ToDateTimeString() // 2020-08-05 12:14:15

// リージョンを設定
utc, _ := time.LoadLocation(carbon.UTC)
carbon.SetLocation(utc).Now().ToDateTimeString() // 2022-06-28 09:25:38
tokyo, _ := time.LoadLocation(carbon.Tokyo)
carbon.SetLocation(tokyo).Now().ToDateTimeString() // 2022-06-28 18:25:38

// ロケールを設定
carbon.Parse("2020-07-05 13:14:15").SetLocale("en").DiffForHumans()) // 1 month ago
carbon.Parse("2020-07-05 13:14:15").SetLocale("jp").DiffForHumans() // 1 ヶ月前

// 年月日時分秒を設定する
carbon.Parse("2020-01-01").SetDateTime(2019, 2, 2, 13, 14, 15).ToString() // 2019-02-02 13:14:15 +0800 CST
carbon.Parse("2020-01-01").SetDateTime(2019, 2, 31, 13, 14, 15).ToString() // 2019-03-03 13:14:15 +0800 CST
// 年月日時分秒ミリ秒を設定する
carbon.Parse("2020-01-01").SetDateTimeMilli(2019, 2, 2, 13, 14, 15, 999).ToString() // 2019-02-02 13:14:15.999 +0800 CST
carbon.Parse("2020-01-01").SetDateTimeMilli(2019, 2, 31, 13, 14, 15, 999).ToString() // 2019-03-03 13:14:15.999 +0800 CST
// 年月日時分秒微秒を設定する
carbon.Parse("2020-01-01").SetDateTimeMicro(2019, 2, 2, 13, 14, 15, 999999).ToString() // 2019-02-02 13:14:15.999999 +0800 CST
carbon.Parse("2020-01-01").SetDateTimeMicro(2019, 2, 31, 13, 14, 15, 999999).ToString() // 2019-03-03 13:14:15.999999 +0800 CST
// 年月日時分秒ナノ秒を設定する
carbon.Parse("2020-01-01").SetDateTimeNano(2019, 2, 2, 13, 14, 15, 999999999).ToString() // 2019-02-02 13:14:15.999999999 +0800 CST
carbon.Parse("2020-01-01").SetDateTimeNano(2019, 2, 31, 13, 14, 15, 999999999).ToString() // 2019-03-03 13:14:15.999999999 +0800 CST

// 年月日を設定する
carbon.Parse("2020-01-01").SetDate(2019, 2, 2).ToString() // 2019-02-02 00:00:00 +0800 CST
carbon.Parse("2020-01-01").SetDate(2019, 2, 31).ToString() // 2019-03-03 00:00:00 +0800 CST
// 年月日ミリ秒を設定する
carbon.Parse("2020-01-01").SetDateMilli(2019, 2, 2, 999).ToString() // 2019-02-02 00:00:00.999 +0800 CST
carbon.Parse("2020-01-01").SetDateMilli(2019, 2, 31, 999).ToString() // 2019-03-03 00:00:00.999 +0800 CST
// 年月日微秒を設定する
carbon.Parse("2020-01-01").SetDateMicro(2019, 2, 2, 999999).ToString() // 2019-02-02 00:00:00.999999 +0800 CST
carbon.Parse("2020-01-01").SetDateMicro(2019, 2, 31, 999999).ToString() // 2019-03-03 00:00:00.999999 +0800 CST
// 年月日ナノ秒を設定する
carbon.Parse("2020-01-01").SetDateNano(2019, 2, 2, 999999999).ToString() // 2019-02-02 00:00:00.999999999 +0800 CST
carbon.Parse("2020-01-01").SetDateNano(2019, 2, 31, 999999999).ToString() // 2019-03-03 00:00:00.999999999 +0800 CST

// 時分秒を設定する
carbon.Parse("2020-01-01").SetTime(13, 14, 15).ToString() // 2020-01-01 13:14:15 +0800 CST
carbon.Parse("2020-01-01").SetTime(13, 14, 90).ToString() // 2020-01-01 13:15:30 +0800 CST
// 時分秒ミリ秒を設定する
carbon.Parse("2020-01-01").SetTimeMilli(13, 14, 15, 999).ToString() // 2020-01-01 13:14:15.999 +0800 CST
carbon.Parse("2020-01-01").SetTimeMilli(13, 14, 90, 999).ToString() // 2020-01-01 13:15:30.999 +0800 CST
// 時分秒微秒を設定する
carbon.Parse("2020-01-01").SetTimeMicro(13, 14, 15, 999999).ToString() // 2020-01-01 13:14:15.999999 +0800 CST
carbon.Parse("2020-01-01").SetTimeMicro(13, 14, 90, 999999).ToString() // 2020-01-01 13:15:30.999999 +0800 CST
// 時分秒ナノ秒を設定する
carbon.Parse("2020-01-01").SetTimeNano(13, 14, 15, 999999999).ToString() // 2020-01-01 13:14:15.999999999 +0800 CST
carbon.Parse("2020-01-01").SetTimeNano(13, 14, 90, 999999999).ToString() // 2020-01-01 13:15:30.999999999 +0800 CST

// 年を設定する
carbon.Parse("2020-02-29").SetYear(2021).ToDateString() // 2021-03-01
// 年を設定する(オーバーフローなし)
carbon.Parse("2020-02-29").SetYearNoOverflow(2021).ToDateString() // 2021-02-28

// 月を設定する
carbon.Parse("2020-01-31").SetMonth(2).ToDateString() // 2020-03-02
// 月を設定する(オーバーフローなし)
carbon.Parse("2020-01-31").SetMonthNoOverflow(2).ToDateString() // 2020-02-29

// 週の開始日を設定する
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Sunday).Week() // 0
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Monday).Week() // 6

// 日数を設定する
carbon.Parse("2019-08-05").SetDay(31).ToDateString() // 2020-08-31
carbon.Parse("2020-02-01").SetDay(31).ToDateString() // 2020-03-02

// 時間を設定する
carbon.Parse("2020-08-05 13:14:15").SetHour(10).ToDateTimeString() // 2020-08-05 10:14:15
carbon.Parse("2020-08-05 13:14:15").SetHour(24).ToDateTimeString() // 2020-08-06 00:14:15

// 分を設定する
carbon.Parse("2020-08-05 13:14:15").SetMinute(10).ToDateTimeString() // 2020-08-05 13:10:15
carbon.Parse("2020-08-05 13:14:15").SetMinute(60).ToDateTimeString() // 2020-08-05 14:00:15

// 秒を設定する
carbon.Parse("2020-08-05 13:14:15").SetSecond(10).ToDateTimeString() // 2020-08-05 13:14:10
carbon.Parse("2020-08-05 13:14:15").SetSecond(60).ToDateTimeString() // 2020-08-05 13:15:00

// ミリ秒を設定
carbon.Parse("2020-08-05 13:14:15").SetMillisecond(100).Millisecond() // 100
carbon.Parse("2020-08-05 13:14:15").SetMillisecond(999).Millisecond() // 999

// 微妙に設定
carbon.Parse("2020-08-05 13:14:15").SetMicrosecond(100000).Microsecond() // 100000
carbon.Parse("2020-08-05 13:14:15").SetMicrosecond(999999).Microsecond() // 999999

// ナノ秒を設定する
carbon.Parse("2020-08-05 13:14:15").SetNanosecond(100000000).Nanosecond() // 100000000
carbon.Parse("2020-08-05 13:14:15").SetNanosecond(999999999).Nanosecond() // 999999999
```

##### 取得

```go
// 年の総日数を取得
carbon.Parse("2019-08-05 13:14:15").DaysInYear() // 365
carbon.Parse("2020-08-05 13:14:15").DaysInYear() // 366
// 月の総日数を取得
carbon.Parse("2020-02-01 13:14:15").DaysInMonth() // 29
carbon.Parse("2020-04-01 13:14:15").DaysInMonth() // 30
carbon.Parse("2020-08-01 13:14:15").DaysInMonth() // 31

// 年間積算日を取得
carbon.Parse("2020-08-05 13:14:15").DayOfYear() // 218
// 本年の何週目かを取得
carbon.Parse("2019-12-31 13:14:15").WeekOfYear() // 1
carbon.Parse("2020-08-05 13:14:15").WeekOfYear() // 32
// 今月の何日目（1から）かを取得
carbon.Parse("2020-08-05 13:14:15").DayOfMonth() // 5
// 今月の何週目かを取得
carbon.Parse("2020-08-05 13:14:15").WeekOfMonth() // 1
// 今週の何日目かを取得(1が月曜日)
carbon.Parse("2020-08-05 13:14:15").DayOfWeek() // 3

// 現在の年月日時分秒を取得
carbon.Parse("2020-08-05 13:14:15").DateTime() // 2020,8,5,13,14,15
// 現在の年月日時分秒ミリ秒を取得
carbon.Parse("2020-08-05 13:14:15").DateTimeMilli() // 2020,8,5,13,14,15,999
// 現在の年月日時分秒マイクロ秒を取得
carbon.Parse("2020-08-05 13:14:15").DateTimeMicro() // 2020,8,5,13,14,15,999999
// 現在の年月日時分秒ナノ秒を取得
carbon.Parse("2020-08-05 13:14:15").DateTimeNano() // 2020,8,5,13,14,15,999999999

// 現在の年月日を取得
carbon.Parse("2020-08-05 13:14:15.999999999").Date() // 2020,8,5
// 現在の年月日ミリ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").DateMilli() // 2020,8,5,999
// 現在の年月日マイクロ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").DateMicro() // 2020,8,5,999999
// 現在の年月日ナノ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").DateNano() // 2020,8,5,999999999

// 現在の時分秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").Time() // 13,14,15
// 現在の時分秒ミリ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").TimeMilli() // 13,14,15,999
// 現在の時分秒マイクロ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").TimeMicro() // 13,14,15,999999
// 現在の時分秒ナノ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").TimeNano() // 13,14,15,999999999

// 現在の世紀を取得
carbon.Parse("2020-08-05 13:14:15").Century() // 21
// 現在の十年紀を取得
carbon.Parse("2019-08-05 13:14:15").Decade() // 10
carbon.Parse("2021-08-05 13:14:15").Decade() // 20
// 現在の年を取得
carbon.Parse("2020-08-05 13:14:15").Year() // 2020
// 現在の四半期を取得
carbon.Parse("2020-08-05 13:14:15").Quarter() // 3
// 現在の月を取得
carbon.Parse("2020-08-05 13:14:15").Month() // 8
// 現在の週を取得(0から開始)
carbon.Parse("2020-08-02 13:14:15").Week() // 0
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Sunday).Week() // 0
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Monday).Week() // 6
// 現在の日数を取得
carbon.Parse("2020-08-05 13:14:15").Day() // 5
// 現在の時間を取得
carbon.Parse("2020-08-05 13:14:15").Hour() // 13
// 現在の分を取得
carbon.Parse("2020-08-05 13:14:15").Minute() // 14
// 現在の秒を取得
carbon.Parse("2020-08-05 13:14:15").Second() // 15
// 現在のミリ秒を取得
carbon.Parse("2020-08-05 13:14:15.999").Millisecond() // 999
// 現在のマイクロ秒を取得
carbon.Parse("2020-08-05 13:14:15.999").Microsecond() // 999000
// 現在のナノ秒を取得
carbon.Parse("2020-08-05 13:14:15.999").Nanosecond() // 999000000

// 秒タイムスタンプを取得
carbon.Parse("2020-08-05 13:14:15").Timestamp() // 1596604455
// ミリ秒のタイムスタンプを取得
carbon.Parse("2020-08-05 13:14:15").TimestampMilli() // 1596604455000
// マイクロ秒タイムスタンプを取得
carbon.Parse("2020-08-05 13:14:15").TimestampMicro() // 1596604455000000
// ナノ秒タイムスタンプを取得
carbon.Parse("2020-08-05 13:14:15").TimestampNano() // 1596604455000000000

// タイムゾーンロケーションの取得
carbon.SetTimezone(carbon.PRC).Timezone() // PRC
carbon.SetTimezone(carbon.Tokyo).Timezone() // Asia/Tokyo

// タイムゾーン名の取得
carbon.SetTimezone(carbon.PRC).ZoneName() // CST
carbon.SetTimezone(carbon.Tokyo).ZoneName() // JST

// UTCタイムゾーンオフセットの秒を取得
carbon.SetTimezone(carbon.PRC).ZoneOffset() // 28800
carbon.SetTimezone(carbon.Tokyo).ZoneOffset() // 32400

// ロケール名を取得
carbon.Now().Locale() // en
carbon.Now().SetLocale("zh-CN").Locale() // zh-CN

// 星座を取得
carbon.Now().Constellation() // Leo
carbon.Now().SetLocale("en").Constellation() // Leo
carbon.Now().SetLocale("jp").Constellation() // しし座

// 季節を取得
carbon.Now().Season() // Summer
carbon.Now().SetLocale("en").Season() // Summer
carbon.Now().SetLocale("jp").Season() // 夏

// 週の開始日の取得
carbon.SetWeekStartsAt(carbon.Sunday).WeekStartsAt() // Sunday
carbon.SetWeekStartsAt(carbon.Monday).WeekStartsAt() // Monday

// 現在のレイアウトテンプレートの取得
carbon.Parse("now").CurrentLayout() // "2006-01-02 15:04:05"
carbon.ParseByLayout("2020-08-05", DateLayout).CurrentLayout() // "2006-01-02"

// 年齢を取得
carbon.Parse("2002-01-01 13:14:15").Age() // 17
carbon.Parse("2002-12-31 13:14:15").Age() // 18
```

##### 出力

```go
// datetimeを文字列出力
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
// ミリ秒を含むdatetimeを文字列出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeMilliString() // 2020-08-05 13:14:15.999
// マイクロ秒を含むdatetimeを文字列出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeMicroString() // 2020-08-05 13:14:15.999999
// ナノ秒を含むdatetimeを文字列出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeNanoString() // 2020-08-05 13:14:15.999999999

// datetimeを略語形式の文字列出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeString() // 20200805131415
// ミリ秒を含むdatetimeを略語形式の文字列出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeMilliString() // 20200805131415.999
// マイクロ秒を含むdatetimeを略語形式の文字列出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeMicroString() // 20200805131415.999999
// ナノ秒を含むdatetimeを略語形式の文字列出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeNanoString() // 20200805131415.999999999

// dateを文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateString() // 2020-08-05
// ミリ秒を含むdateを文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateMilliString() // 2020-08-05.999
// マイクロ秒を含むdateを文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateMicroString() // 2020-08-05.999999
// ナノ秒を含むdateを文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateNanoString() // 2020-08-05.999999999

// dateを略語形式の文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateString() // 20200805
// ミリ秒を含むdateを略語形式の文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateMilliString() // 20200805.999
// マイクロ秒を含むdateを略語形式の文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateMicroString() // 20200805.999999
// ナノ秒を含むdateを略語形式の文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateNanoString() // 20200805.999999999

// 時間を文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeString() // 13:14:15
// ミリ秒を含む時間を文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeMilliString() // 13:14:15.999
// マイクロ秒を含む時間を文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeMicroString() // 13:14:15.999999
// ナノ秒を含む時間を文字列出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeNanoString() // 13:14:15.999999999

// 時間を略語形式の出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeString() // 131415
// ミリ秒を含む時間を略語形式の出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeMilliString() // 131415.999
// マイクロ秒を含む時間を略語形式の出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeMicroString() // 131415.999999
// ナノ秒を含む時間を略語形式の出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeNanoString() // 131415.999999999

// Ansic フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToAnsicString() // Wed Aug  5 13:14:15 2020
// Atom フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToAtomString() // 2020-08-05T13:14:15+08:00
// UnixDate フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToUnixDateString() // Wed Aug  5 13:14:15 CST 2020
// RubyDate フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRubyDateString() // Wed Aug 05 13:14:15 +0800 2020
// Kitchen フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToKitchenString() // 1:14PM
// Cookie フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToCookieString() // Wednesday, 05-Aug-2020 13:14:15 CST
// DayDateTime フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString() // Wed, Aug 5, 2020 1:14 PM
// RSS フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRssString() // Wed, 05 Aug 2020 13:14:15 +0800
// W3C フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToW3cString() // 2020-08-05T13:14:15+08:00

// ISO8601 フォーマット文字列の出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601String() // 2020-08-05T13:14:15+08:00
// ISO8601Milli フォーマット文字列の出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601MilliString() // 2020-08-05T13:14:15.999+08:00
// ISO8601Micro フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601MicroString() // 2020-08-05T13:14:15.999999+08:00
// ISO8601Nano フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601NanoString() // 2020-08-05T13:14:15.999999999+08:00
// ISO8601Zulu フォーマット文字列の出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluString() // 2020-08-05T13:14:15Z
// ISO8601ZuluMilli フォーマット文字列の出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluMilliString() // 2020-08-05T13:14:15.999Z
// ISO8601ZuluMicro フォーマット文字列の出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluMicroString() // 2020-08-05T13:14:15.999999Z
// ISO8601ZuluNano フォーマット文字列の出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluNanoString() // 2020-08-05T13:14:15.999999999Z

// RFC822 フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRfc822String() // 05 Aug 20 13:14 CST
// RFC822Z フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRfc822zString() // 05 Aug 20 13:14 +0800
// RFC850 フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRfc850String() // Wednesday, 05-Aug-20 13:14:15 CST
// RFC1036 フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRfc1036String() // Wed, 05 Aug 20 13:14:15 +0800
// RFC1123 フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRfc1123String() // Wed, 05 Aug 2020 13:14:15 CST
// RFC1123Z フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRfc1123zString() // Wed, 05 Aug 2020 13:14:15 +0800
// RFC2822 フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRfc2822String() // Wed, 05 Aug 2020 13:14:15 +0800
// RFC7231 フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToRfc7231String() // Wed, 05 Aug 2020 13:14:15 GMT

// RFC3339 フォーマット文字列の出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339String() // 2020-08-05T13:14:15+08:00
// RFC3339 フォーマット文字列の出力(ミリ秒を含む)
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339MilliString() // 2020-08-05T13:14:15.999+08:00
// RFC3339 フォーマット文字列の出力(マイクロ秒を含む)
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339MicroString() // 2020-08-05T13:14:15.999999+08:00
// RFC3339 フォーマット文字列の出力(ナノ秒を含む)
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339NanoString() // 2020-08-05T13:14:15.999999999+08:00

// 日付時間文字列の出力
fmt.Printf("%s", carbon.Parse("2020-08-05 13:14:15")) // 2020-08-05 13:14:15

// "2006-01-02 15:04:05.999999999 -0700 MST" フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15.999999 +0800 CST

// "Jan 2, 2006" フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToFormattedDateString() // Aug 5, 2020
// "Mon, Jan 2, 2006" フォーマット文字列の出力
carbon.Parse("2020-08-05 13:14:15").ToFormattedDayDateString() // Wed, Aug 5, 2020

// レイアウトを指定する文字列の出力
carbon.Parse("2020-08-05 13:14:15").Layout(carbon.ISO8601Layout) // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").Layout("20060102150405") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Layout("2006年01月02日 15时04分05秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Layout("It is 2006-01-02 15:04:05") // It is 2020-08-05 13:14:15

// 指定されたフォーマットの文字列の出力
carbon.Parse("2020-08-05 13:14:15").Format("YmdHis") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Format("Y年m月d日 H时i分s秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Format("l jS \\o\\f F Y h:i:s A") // Wednesday 5th of August 2020 01:14:15 PM
carbon.Parse("2020-08-05 13:14:15").Format("\\I\\t \\i\\s Y-m-d H:i:s") // It is 2020-08-05 13:14:15
```

> もっとフォーマットした出力記号は付録を見てください <a href="#format-sign-table">書式設定記号表</a>

##### 星座

```go
// 星座を取得
carbon.Parse("2020-08-05 13:14:15").Constellation() // Leo

// おひつじ座かどうか
carbon.Parse("2020-08-05 13:14:15").IsAries() // false
// おうし座かどうか
carbon.Parse("2020-08-05 13:14:15").IsTaurus() // false
// ふたご座かどうか
carbon.Parse("2020-08-05 13:14:15").IsGemini() // false
// かに座かどうか
carbon.Parse("2020-08-05 13:14:15").IsCancer() // false
// しし座かどうか
carbon.Parse("2020-08-05 13:14:15").IsLeo() // true
// おとめ座かどうか
carbon.Parse("2020-08-05 13:14:15").IsVirgo() // false
// てんびん座かどうか
carbon.Parse("2020-08-05 13:14:15").IsLibra() // false
// さそり座かどうか
carbon.Parse("2020-08-05 13:14:15").IsScorpio() // false
// いて座かどうか
carbon.Parse("2020-08-05 13:14:15").IsSagittarius() // false
// やぎ座かどうか
carbon.Parse("2020-08-05 13:14:15").IsCapricorn() // false
// みずがめ座かどうか
carbon.Parse("2020-08-05 13:14:15").IsAquarius() // false
// うお座かどうか
carbon.Parse("2020-08-05 13:14:15").IsPisces() // false
```

##### 季節

> 気象区分によると、3-5月は春で、6-8月は夏で、9-11月は秋で、12-2月は冬です

```go
// 季節を取得
carbon.Parse("2020-08-05 13:14:15").Season() // Summer

// この季節の開始日
carbon.Parse("2020-08-05 13:14:15").StartOfSeason().ToDateTimeString() // 2020-06-01 00:00:00
// この季節の最終日
carbon.Parse("2020-08-05 13:14:15").EndOfSeason().ToDateTimeString() // 2020-08-31 23:59:59

// 春かどうか
carbon.Parse("2020-08-05 13:14:15").IsSpring() // false
// 夏かどうか
carbon.Parse("2020-08-05 13:14:15").IsSummer() // true
// 秋かどうか
carbon.Parse("2020-08-05 13:14:15").IsAutumn() // false
// 冬かどうか
carbon.Parse("2020-08-05 13:14:15").IsWinter() // false
```

##### JSON

```go
type Person struct {
  Name string `json:"name"`
  Age int `json:"age"`
  Birthday0 carbon.Carbon `json:"birthday0"`
  Birthday1 carbon.DateTime `json:"birthday1"`
  Birthday2 carbon.DateTimeMilli `json:"birthday2"`
  Birthday3 carbon.DateTimeMicro `json:"birthday3"`
  Birthday4 carbon.DateTimeNano `json:"birthday4"`
  GraduatedAt1 carbon.Date `json:"graduated_at1"`
  GraduatedAt2 carbon.DateMilli `json:"graduated_at2"`
  GraduatedAt3 carbon.DateMicro `json:"graduated_at3"`
  GraduatedAt4 carbon.DateNano `json:"graduated_at4"`
  OperatedAt1 carbon.Time `json:"operated_at1"`
  OperatedAt2 carbon.TimeMilli `json:"operated_at2"`
  OperatedAt3 carbon.TimeMicro `json:"operated_at3"`
  OperatedAt4 carbon.TimeNano `json:"operated_at4"`
  CreatedAt1 carbon.Timestamp `json:"created_at1"`
  CreatedAt2 carbon.TimestampMilli `json:"created_at2"`
  CreatedAt3 carbon.TimestampMicro `json:"created_at3"`
  CreatedAt4 carbon.TimestampNano `json:"created_at4"`
}

person := Person {
    Name:        "gouguoyin",
    Age:          18,
    Birthday0:    carbon.Now().SubYears(18),
    Birthday1:    carbon.NewDateTime(carbon.Now().SubYears(18)),
    Birthday2:    carbon.NewDateTimeMilli(carbon.Now().SubYears(18)),
    Birthday3:    carbon.NewDateTimeMicro(carbon.Now().SubYears(18)),
    Birthday4:    carbon.NewDateTimeNano(carbon.Now().SubYears(18)),
    GraduatedAt1: carbon.NewDate(carbon.Now()),
    GraduatedAt2: carbon.NewDateMilli(carbon.Now()),
    GraduatedAt3: carbon.NewDateMicro(carbon.Now()),
    GraduatedAt4: carbon.NewDateNano(carbon.Now()),
    OperatedAt1:  carbon.NewTime(carbon.Now()),
    OperatedAt2:  carbon.NewTimeMilli(carbon.Now()),
    OperatedAt3:  carbon.NewTimeMicro(carbon.Now()),
    OperatedAt4:  carbon.NewTimeNano(carbon.Now()),
    CreatedAt1:   carbon.NewTimestamp(carbon.Now()),
    CreatedAt2:   carbon.NewTimestampMilli(carbon.Now()),
    CreatedAt3:   carbon.NewTimestampMicro(carbon.Now()),
    CreatedAt4:   carbon.NewTimestampNano(carbon.Now()),
}

data, err := json.Marshal(&person)
if err != nil {
  // エラー処理...
  log.Fatal(err)
}
fmt.Printf("%s", data)
// 出力
{
  "name": "gouguoyin",
  "age": 18,
  "birthday0": "2003-07-16 13:14:15",
  "birthday1": "2003-07-16 13:14:15",
  "birthday2": "2003-07-16 13:14:15.999",
  "birthday3": "2003-07-16 13:14:15.999999",
  "birthday4": "2003-07-16 13:14:15.999999999",
  "graduated_at1": "2020-08-05",
  "graduated_at2": "2020-08-05.999",
  "graduated_at3": "2020-08-05.999999",
  "graduated_at4": "2020-08-05.999999999",
  "operated_at1": "13:14:15",
  "operated_at2": "13:14:15.999",
  "operated_at3": "13:14:15.999999",
  "operated_at4": "13:14:15.999999999",
  "created_at1": 1596604455,
  "created_at2": 1596604455999,
  "created_at3": 1596604455999999,
  "created_at4": 1596604455999999999
}

err := json.Unmarshal([]byte(data), &person)
if err != nil {
  // エラー処理...
  log.Fatal(err)
}

person.Birthday0.String() // 2003-07-16 13:14:15
person.Birthday1.String() // 2003-07-16 13:14:15
person.Birthday2.String() // 2003-07-16 13:14:15.999
person.Birthday3.String() // 2003-07-16 13:14:15.999999
person.Birthday4.String() // 2003-07-16 13:14:15.999999999

person.GraduatedAt1.String() // 2020-08-05
person.GraduatedAt2.String() // 2020-08-05.999
person.GraduatedAt3.String() // 2020-08-05.999999
person.GraduatedAt4.String() // 2020-08-05.999999999

person.OperatedAt1.String() // 13:14:15
person.OperatedAt2.String() // 13:14:15.999
person.OperatedAt3.String() // 13:14:15.999999
person.OperatedAt4.String() // 13:14:15.999999999

person.CreatedAt1.String() // "1596604455"
person.CreatedAt2.String() // "1596604455999"
person.CreatedAt3.String() // "1596604455999999"
person.CreatedAt4.String() // "1596604455999999999"

person.CreatedAt1.Int64() // 1596604455
person.CreatedAt2.Int64() // 1596604455999
person.CreatedAt3.Int64() // 1596604455999999
person.CreatedAt4.Int64() // 1596604455999999999
```

##### カレンダー

現在サポートされているカレンダー

* [儒略の日/簡略化儒略の日](./calendar/julian/README.jp.md "儒略日/简化儒略日")
* [中国の旧暦](./calendar/lunar/README.jp.md "中国の旧暦")
* [ペルシャ暦/イラン暦](./calendar/persian/README.jp.md "ペルシャ暦/イラン暦")

##### i18n

現在サポートされている言語

* [简体中国語(zh-CN)](./lang/zh-CN.json "简体中国語")：[gouguoyin](https://github.com/gouguoyin "gouguoyin") に翻訳
* [繁体中国語(zh-TW)](./lang/zh-TW.json "繁体中国語")：[gouguoyin](https://github.com/gouguoyin "gouguoyin") に翻訳
* [英語(en)](./lang/en.json "英語")：[gouguoyin](https://github.com/gouguoyin "gouguoyin") に翻訳
* [日本語(jp)](./lang/jp.json "日本語")：[gouguoyin](https://github.com/gouguoyin "gouguoyin") に翻訳
* [韓国語(kr)](./lang/kr.json "韓国語")：[benzammour](https://github.com/benzammour "benzammour") に翻訳
* [ドイツ語(de)](./lang/de.json "ドイツ語")：[benzammour](https://github.com/benzammour "benzammour") に翻訳
* [スペイン語(es)](./lang/es.json "スペイン語")：[hgisinger](https://github.com/hgisinger "hgisinger") に翻訳
* [トルコ語(tr)](./lang/tr.json "トルコ語")：[emresenyuva](https://github.com/emresenyuva "emresenyuva") に翻訳
* [ポルトガル語(pt)](./lang/pt.json "ポルトガル語")：[felipear89](https://github.com/felipear89 "felipear89") に翻訳
* [ロシア語(ru)](./lang/ru.json "ロシア語")：[zemlyak](https://github.com/zemlyak "zemlyak") に翻訳
* [ウクライナ語(uk)](./lang/uk.json "ウクライナ語")：[open-git](https://github.com/open-git "open-git") に翻訳
* [ルーマニア語(ro)](./lang/ro.json "ルーマニア語"): [DrOctavius](https://github.com/DrOctavius "DrOctavius") に翻訳
* [インドネシア語(id)](./lang/id.json "インドネシア語"): [justpoypoy](https://github.com/justpoypoy "justpoypoy") に翻訳
* [イタリア語(it)](./lang/it.json "イタリア語"): [nicoloHevelop](https://github.com/hollowaykeanho "nicoloHevelop") に翻訳
* [マレーシアバハマ語(ms-MY)](./lang/ms-MY.json "マレーシアバハマ語"): [hollowaykeanho](https://github.com/hollowaykeanho "hollowaykeanho") に翻訳
* [フランス語(fr)](./lang/fr.json "フランス語"): [hollowaykeanho](https://github.com/hollowaykeanho "hollowaykeanho") に翻訳
* [タイ語(th)](./lang/th.json "タイ語"): [izcream](https://github.com/izcream "izcream") に翻訳
* [スウェーデン語(se)](./lang/se.json "スウェーデン語"):  [jwanglof](https://github.com/jwanglof "jwanglof") に翻訳
* [ペルシア語(fa)](./lang/fa.json "ペルシア語"):  [Iranian](https://github.com/Iranian "Iranian") に翻訳
* [ポーランド語(nl)](./lang/nl.json "ポーランド語"): [RemcoE33](https://github.com/RemcoE33 "RemcoE33") に翻訳
* [ベトナム語(vi)](./lang/vi.json "ベトナム語"): [culy247](https://github.com/culy247 "culy247") に翻訳
* [ヒンディー語(hi)](./lang/hi.json "ヒンディー語"): [chauhan17nitin](https://github.com/chauhan17nitin "chauhan17nitin") に翻訳
* [ポーランド語(pl)](./lang/pl.json "ポーランド語"): [gouguoyin](https://github.com/gouguoyin "gouguoyin") に翻訳
* [ブルガリア語(bg)](./lang/bg.json "ブルガリア語"): [yuksbg](https://github.com/yuksbg "yuksbg") に翻訳
* [アラビア語(ar)](./lang/ar.json "アラビア語"): [zumoshi](https://github.com/zumoshi "zumoshi") に翻訳
* [ハンガリー語(hu)](./lang/hu.json "ハンガリー語"): [kenlas](https://github.com/kenlas "kenlas") に翻訳
* [デンマーク語(dk)](./lang/dk.json "デンマーク語"): [Munk91](https://github.com/Munk91 "Munk91") に翻訳
* [ノルウェー語(dk)](./lang/nb.json "ノルウェー語"): [bendikrb](https://github.com/bendikrb "bendikrb") に翻訳

現在サポートされている方法

* `Constellation()`：星座を取得，例えば `おひつじ座`
* `Season()`：シーズンを取得，例えば `春`
* `DiffForHumans()`：人が読みやすい形の時間差を取得します，例えば `1時間前`
* `ToMonthString()`：月文字列の出力，例えば `いちがつ`
* `ToShortMonthString()`：略語月文字列の出力，例えば `1がつ`
* `ToWeekString()`：週文字列の出力，例えば `日曜日`
* `ToShortWeekString()`：略語週文字列の出力，例えば `日`

###### エリアの設定

```go
lang := carbon.NewLanguage()
lang.SetLocale("jp")

c := carbon.SetLanguage(lang)
if c.Error != nil {
  // エラー処理
  log.Fatal(c.Error)
}

c.Now().AddHours(1).DiffForHumans() // 1 時間後
c.Now().AddHours(1).ToMonthString() // はちがつ
c.Now().AddHours(1).ToShortMonthString() // 8がつ
c.Now().AddHours(1).ToWeekString() // 日曜日
c.Now().AddHours(1).ToShortWeekString() // 日
c.Now().AddHours(1).Constellation() // しし座
c.Now().AddHours(1).Season() // 夏
```

###### 翻訳リソースの一部を書き換える(残りはまだ指定された `locale` ファイルの内容によって翻訳されます)

```go
lang := carbon.NewLanguage()

resources := map[string]string {
  "hour": "%dh",
}
lang.SetLocale("en").SetResources(resources)

c := carbon.SetLanguage(lang)
if c.Error != nil {
  // エラー処理
  log.Fatal(c.Error)
}

c.Now().AddYears(1).DiffForHumans() // 1 year from now
c.Now().AddHours(1).DiffForHumans() // 1h from now
c.Now().ToMonthString() // August
c.Now().ToShortMonthString() // Aug
c.Now().ToWeekString() // Tuesday
c.Now().ToShortWeekString() // Tue
c.Now().Constellation() // Leo
c.Now().Season() // Summer
```

###### すべての翻訳リソースを書き換える

```go
lang := carbon.NewLanguage()
resources := map[string]string {
  "months": "january|february|march|april|may|june|july|august|september|october|november|december",
  "short_months": "jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec",
  "weeks": "sunday|monday|tuesday|wednesday|thursday|friday|saturday",
  "short_weeks": "sun|mon|tue|wed|thu|fri|sat",
  "seasons": "spring|summer|autumn|winter",
  "constellations": "aries|taurus|gemini|cancer|leo|virgo|libra|scorpio|sagittarius|capricornus|aquarius|pisce",
  "year": "1 yr|%d yrs",
  "month": "1 mo|%d mos",
  "week": "%dw",
  "day": "%dd",
  "hour": "%dh",
  "minute": "%dm",
  "second": "%ds",
  "now": "just now",
  "ago": "%s ago",
  "from_now": "in %s",
  "before": "%s before",
  "after": "%s after",
}
lang.SetResources(resources)

c := carbon.SetLanguage(lang)
c.Now().AddYears(1).DiffForHumans() // in 1 yr
c.Now().AddHours(1).DiffForHumans() // in 1h
c.Now().ToMonthString() // august
c.Now().ToShortMonthString() // aug
c.Now().ToWeekString() // tuesday
c.Now().ToShortWeekString() // tue
c.Now().Constellation() // leo
c.Now().Season() // summer
```


##### エラー処理

```go
c := carbon.SetTimezone("xxx").Parse("2020-08-05")
if c.Error != nil {
  // エラー処理...
  log.Fatal(c.Error)
}
// 出力
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezones
```

#### 付録

##### <a id="format-sign-table">書式設定記号表</a>

|     記号            | 説明            | 長さ | 範囲            |  例            |
| :------------: | :------------: |:------:|:----------------:| :------------: |
| d | 月の日(2桁でパディング) | 2 | 01-31 | 02 |
| D | 略語の曜日 | 3 | Mon-Sun | Mon |
| j | 月の日(パディングしない) | - | 1-31 | 2 |
| S | 何日目の英語の略語の接尾語，普通はjと協力して使います | 2 | st/nd/rd/th | th |
| l | 曜日 | - | Monday-Sunday | Monday |
| F | 月 | - | January-December | January |
| m | 月(2桁でパディング) | 2 | 01-12 | 01 |
| M | 略語の月 | 3 | Jan-Dec | Jan |
| n | 月(パディングしない) | - | 1-12 | 1 |
| Y | 年 | 4 | 0000-9999 | 2006 |
| y | 年(下2桁) | 2 | 00-99 | 06 |
| a | 小文字の午前と午後 | 2 | am/pm | pm |
| A | 大文字の午前と午後 | 2 | AM/PM | PM |
| g | 時間, 12時間のフォーマット | - | 1-12 | 3 |
| G | 時間, 24時間のフォーマット | - | 0-23 | 15 |
| h | 時間, 12時間のフォーマット | 2 | 00-11 | 03 |
| H | 時間, 24時間のフォーマット | 2 | 00-23 | 15 |
| i | 分 | 2 | 01-59 | 04 |
| s | 秒 | 2 | 01-59 | 05 |
| O | グリニッジとの時間差の時間数 | - | - | -0700 |
| P | グリニッジと時間の差の時間数, 時間と分の間にコロンあり | - | - | -07:00 |
| T | タイムゾーンの略語 | - | - | MST |
| W | ISO8601 フォーマットの数字は年の中の第何週(2桁でパディング) | 2 | 1-52 | 01 |
| N | ISO8601 フォーマットの数字は曜日(2桁でパディング) | 2 | 01-07 | 02 |
| L | うるう年かどうか, うるう年が1であれば, 0 | 1 | 0-1 | 0 |
| U | 秒タイムスタンプ | - | - | 1596604455 |
| V | ミリ秒のタイムスタンプ | - | - | 1596604455666 |
| X | マイクロ秒タイムスタンプ | - | - | 1596604455666666 |
| Z | ナノ秒タイムスタンプ | - | - | 1596604455666666666 |
| v | ミリ秒 | - | 1-999 | 999 |
| u | マイクロ秒| - | 1-999999 | 999999 |
| x | ナノ秒 | - | 1-999999999 | 999999999 |
| w | 数字表示の曜日 | 1 | 0-6 | 1 |
| t | 月の総日数 | 2 | 28-31 | 31 |
| z | 年間積算日 | - | 1-365 | 2 |
| e | リージョン | - | - | America/New_York |
| Q | 四半期 | 1 | 1-4 | 1 |
| C | 世紀 | - | 0-99 | 21 |

#### FAQ

1、v1とv2のバージョンの違いは何ですか？
> APIのv1バージョンとv2バージョンに違いはありませんが `language.go`
> での翻訳リソースファイルの実装は異なります。v1は、サードパーティの拡張ライブラリ [packr](https://github.com/gobuffalo/packr) によって実装されています，v2は、 `golang1.16`
> の後に組み込みの標準ライブラリ [embed](https://pkg.go.dev/embed) によって実装されています。goバージョンが1.16より大きい場合は、v2バージョンを使用することをお勧めします。それ以外の場合は、v1バージョンを使用する必要があります。

#### 参考

* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)
* [nodatime/nodatime](https://github.com/nodatime/nodatime)
* [jinzhu/now](https://github.com/jinzhu/now)
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)
* [jodaOrg/joda-time](https://github.com/jodaOrg/joda-time)
* [arrow-py/arrow](https://github.com/arrow-py/arrow)
* [moment/moment](https://github.com/moment/moment)
* [iamkun/dayjs](https://github.com/iamkun/dayjs)

#### コントリビューター
`Carbon` に貢献してくれた以下のすべてに感謝します：

<a href="https://github.com/dromara/carbon/graphs/contributors"><img src="https://contrib.rocks/image?repo=dromara/carbon&max=100&columns=16"/></a>

#### スポンサー

`Carbon` は非営利のオープンソースプロジェクトです，`Carbon` をサポートしたい場合は、開発者のために [コーヒーを1杯購入](https://opencollective.com/go-carbon) できます

#### 謝辞

`Carbon` は無料の JetBrains オープンソースライセンスを取得しました，これに感謝します

<a href="https://www.jetbrains.com"><img src="https://foruda.gitee.com/images/1704325523163241662/1bf21f86_544375.png" height="100" alt="JetBrains"/></a>
