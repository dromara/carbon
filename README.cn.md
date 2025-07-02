<p align="center" style="margin-bottom: -10px"><a href="https://carbon.go-pkg.com/zh" target="_blank"><img src="https://carbon.go-pkg.com/logo.svg?v=2.6.x" width="15%" alt="carbon" /></a></p>

[![Carbon Release](https://img.shields.io/github/release/dromara/carbon.svg)](https://github.com/dromara/carbon/releases)
[![Go Test](https://github.com/dromara/carbon/actions/workflows/test.yml/badge.svg)](https://github.com/dromara/carbon/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/dromara/carbon/v2)](https://goreportcard.com/report/github.com/dromara/carbon/v2)
[![Go Coverage](https://codecov.io/gh/dromara/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/dromara/carbon)
[![Carbon Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/dromara/carbon/v2)
[![Awesome](https://awesome.re/badge-flat2.svg)](https://github.com/avelino/awesome-go#date-and-time)
[![HelloGitHub](https://api.hellogithub.com/v1/widgets/recommend.svg?rid=0eddd8c3469549b7b246f85a83d1c42e&claim_uid=kKBvMpyxSgLhmJO&theme=small)](https://hellogithub.com/en/repository/dromara/carbon)
[![License](https://img.shields.io/github/license/dromara/carbon)](https://github.com/dromara/carbon/blob/master/LICENSE)

简体中文 | [English](README.md) | [日本語](README.ja.md)

## 项目简介

`Carbon` 是一个轻量级、语义化、对开发者友好的 `golang` 时间处理库，不依赖于 `任何` 第三方库， `100%` 单元测试覆盖率，已被 [awesome-go](https://github.com/yinggaozhen/awesome-go-cn#日期和时间 "awesome-go-cn") 和 [hello-github](https://hellogithub.com/repository/dromara/carbon "hello-github") 收录，并获得
`gitee` 2024 年最有价值项目（`GVP`）和 `gitcode` 2024 年度开源摘星计划 (`G-Star`) 项目

<img src="https://carbon.go-pkg.com/gvp.jpg?v=2.6.x" width="100%" alt="gvp"/>
<img src="https://carbon.go-pkg.com/gstar.jpg?v=2.6.x" width="100%" alt="g-star"/>

## 仓库地址

[github.com/dromara/carbon](https://github.com/dromara/carbon "github.com/dromara/carbon")

[gitee.com/dromara/carbon](https://gitee.com/dromara/carbon "gitee.com/dromara/carbon")

[gitcode.com/dromara/carbon](https://gitcode.com/dromara/carbon "gitcode.com/dromara/carbon")

## 快速开始

### 安装使用

> go version >= 1.21

```go
// 使用 github 库
go get -u github.com/dromara/carbon/v2
import "github.com/dromara/carbon/v2"

// 使用 gitee 库
go get -u gitee.com/dromara/carbon/v2
import "gitee.com/dromara/carbon/v2"

// 使用 gitcode 库
go get -u gitcode.com/dromara/carbon/v2
import "gitcode.com/dromara/carbon/v2"
```

`Carbon` 已经捐赠给了 [dromara](https://dromara.org/ "dromara") 开源组织，仓库地址发生了改变，如果之前用的路径是
`golang-module/carbon`，请在 `go.mod` 里将原地址更换为新路径，或执行如下命令

```go
go mod edit -replace github.com/golang-module/carbon/v2 = github.com/dromara/carbon/v2
```

### 用法示例

默认时区是 `UTC`, 语言环境是 `英语`，一周开始日期是 `周一`，周末是 `周六`和 `周日`。

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
carbon.Parse("2020-07-05 13:14:15").SetLocale("zh-CN").DiffForHumans() // 1 月前

carbon.ClearTestNow()
carbon.IsTestNow() // false
```
更多示例请查看 <a href="https://carbon.go-pkg.com/zh" target="_blank">官方文档</a>

## 参考项目

* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)
* [nodatime/nodatime](https://github.com/nodatime/nodatime)
* [jinzhu/now](https://github.com/jinzhu/now)
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)
* [jodaOrg/joda-time](https://github.com/jodaOrg/joda-time)
* [arrow-py/arrow](https://github.com/arrow-py/arrow)
* [moment/moment](https://github.com/moment/moment)
* [iamkun/dayjs](https://github.com/iamkun/dayjs)

## 贡献者

感谢以下所有为 `Carbon` 做出贡献的人：

<a href="https://github.com/dromara/carbon/graphs/contributors"><img src="https://contrib.rocks/image?repo=dromara/carbon&max=100&columns=16"/></a>

## 赞助

`Carbon` 是一个非商业开源项目, 如果你想支持 `Carbon`,
你可以为开发者 [购买一杯咖啡](https://www.gouguoyin.com/zanzhu.html)

## 致谢

`Carbon`已获取免费的 JetBrains 开源许可证，在此表示感谢

<a href="https://www.jetbrains.com" target="_blank"><img src="https://carbon.go-pkg.com/jetbrains.svg?v=2.6.x" height="50" alt="JetBrains"/></a>

## 开源协议

`Carbon` 遵循 `MIT` 开源协议, 请参阅 [LICENSE](./LICENSE) 查看详细信息。
