<p align="center" style="margin-bottom: -10px"><a href="https://carbon.go-pkg.com/ja" target="_blank"><img src="./logo.svg" width="15%" alt="carbon" /></a></p>

[![Carbon Release](https://img.shields.io/github/release/dromara/carbon.svg)](https://github.com/dromara/carbon/releases)
[![Go Test](https://github.com/dromara/carbon/actions/workflows/test.yml/badge.svg)](https://github.com/dromara/carbon/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/dromara/carbon/v2)](https://goreportcard.com/report/github.com/dromara/carbon/v2)
[![Go Coverage](https://codecov.io/gh/dromara/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/dromara/carbon)
[![Carbon Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/dromara/carbon/v2)
[![Awesome](https://awesome.re/badge-flat2.svg)](https://github.com/avelino/awesome-go#date-and-time)
[![HelloGitHub](https://api.hellogithub.com/v1/widgets/recommend.svg?rid=0eddd8c3469549b7b246f85a83d1c42e&claim_uid=kKBvMpyxSgLhmJO&theme=small)](https://hellogithub.com/en/repository/dromara/carbon)
[![License](https://img.shields.io/github/license/dromara/carbon)](https://github.com/dromara/carbon/blob/master/LICENSE)

日本語 | [English](README.md) | [简体中文](README.cn.md)

## イントロ

`Carbon` は軽量、セマンティック、開発者に優しい `golang` 時間処理ライブラリ, `いかなる`第三者ライブラリにも依存せず、`100%`ユニットテストカバレッジ率は、[awesome-go](https://github.com/avelino/awesome-go#date-and-time "awesome-go") と [hello-github](https://hellogithub.com/en/repository/dromara/carbon "hello-github") 収録

## リポジトリ

[github.com/dromara/carbon](https://github.com/dromara/carbon "github.com/dromara/carbon")

[gitee.com/dromara/carbon](https://gitee.com/dromara/carbon "gitee.com/dromara/carbon")

[gitcode.com/dromara/carbon](https://gitcode.com/dromara/carbon "gitcode.com/dromara/carbon")

## クイックスタート

### インストール
> go version >= 1.21

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

### 使い方と例
デフォルトのタイムゾーンは` UTC `、ロケールは`英語`、週の開始日は`月曜日`、週末は`土曜日`、`日曜日`。

```go
carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15.999999999"))
carbon.IsTestNow() // true

carbon.Now().ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC
carbon.Yesterday().ToString() // 2020-08-04 13:14:15.999999999 +0000 UTC
carbon.Tomorrow().ToString() // 2020-08-06 13:14:15.999999999 +0000 UTC

carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2022-03-08T03:01:14-07:00").ToString() // 2022-03-08 10:01:14 +0000 UTC

carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToString() // 2020-08-05 13:14:15 +0000 UTC

carbon.CreateFromDate(2020, 8, 5).ToString() // 2020-08-05 00:00:00 +0000 UTC
carbon.CreateFromTime(13, 14, 15).ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.CreateFromTimestamp(1596633255).ToString() // 2020-08-05 13:14:15 +0000 UTC

carbon.Parse("2020-07-05 13:14:15").DiffForHumans() // 1 month before
carbon.Parse("2020-07-05 13:14:15").SetLocale("jp").DiffForHumans() // 1 ヶ月前

carbon.ClearTestNow()
carbon.IsTestNow() // false
```
詳細については <a href="https://carbon.go-pkg.com/ja" target="_blank">公式ドキュメント</a>

## リファレンス

* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)
* [nodatime/nodatime](https://github.com/nodatime/nodatime)
* [jinzhu/now](https://github.com/jinzhu/now)
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)
* [jodaOrg/joda-time](https://github.com/jodaOrg/joda-time)
* [arrow-py/arrow](https://github.com/arrow-py/arrow)
* [moment/moment](https://github.com/moment/moment)
* [iamkun/dayjs](https://github.com/iamkun/dayjs)

## コントリビューター
`Carbon` に貢献してくれた以下のすべてに感謝します：

<a href="https://github.com/dromara/carbon/graphs/contributors"><img src="https://contrib.rocks/image?repo=dromara/carbon&max=100&columns=16"/></a>

## スポンサー

`Carbon` は非営利のオープンソースプロジェクトです，`Carbon` をサポートしたい場合は、開発者のために [コーヒーを1杯購入](https://opencollective.com/go-carbon) できます

## 謝辞

`Carbon` は無料の JetBrains オープンソースライセンスを取得しました，これに感謝します

<a href="https://www.jetbrains.com"><img src="https://foruda.gitee.com/images/1704325523163241662/1bf21f86_544375.png" height="100" alt="JetBrains"/></a>

## オープンソースプロトコル

`Carbon` は `MIT` オープンソースプロトコルに準拠しており、詳細は[LICENSE]（./LICENSE）を参照してください
