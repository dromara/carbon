<p align="center" style="margin-bottom: -10px"><a href="https://carbon.go-pkg.com/zh" target="_blank"><img src="https://carbon.go-pkg.com/logo.svg?v=2.6.x" width="15%" alt="carbon" /></a></p>

[![Carbon Release](https://img.shields.io/github/release/dromara/carbon.svg)](https://github.com/dromara/carbon/releases)
[![Go Test](https://github.com/dromara/carbon/actions/workflows/test.yml/badge.svg)](https://github.com/dromara/carbon/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/dromara/carbon/v2)](https://goreportcard.com/report/github.com/dromara/carbon/v2)
[![Go Coverage](https://codecov.io/gh/dromara/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/dromara/carbon)
[![Carbon Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/dromara/carbon/v2)
[![Awesome](https://awesome.re/badge-flat2.svg)](https://github.com/avelino/awesome-go#date-and-time)
[![HelloGitHub](https://api.hellogithub.com/v1/widgets/recommend.svg?rid=0eddd8c3469549b7b246f85a83d1c42e&claim_uid=kKBvMpyxSgLhmJO&theme=small)](https://hellogithub.com/en/repository/dromara/carbon)
[![License](https://img.shields.io/github/license/dromara/carbon)](https://github.com/dromara/carbon/blob/master/LICENSE)

English | [简体中文](README.cn.md) | [日本語](README.ja.md) | [한국어](README.ko.md)

## Introduction

`Carbon` is a simple, semantic and developer-friendly time package for `golang`, `100%` unit test coverage, doesn't depend on `any` third-party package and has been included by [awesome-go](https://github.com/avelino/awesome-go#date-and-time "awesome-go") and [hello-github](https://hellogithub.com/en/repository/dromara/carbon "hello-github")

## Repository

[github.com/dromara/carbon](https://github.com/dromara/carbon "github.com/dromara/carbon")

[gitee.com/dromara/carbon](https://gitee.com/dromara/carbon "gitee.com/dromara/carbon")

[gitcode.com/dromara/carbon](https://gitcode.com/dromara/carbon "gitcode.com/dromara/carbon")

## Quick Start

### Installation
> go version >= 1.21

```go
// By github
go get -u github.com/dromara/carbon/v2
import "github.com/dromara/carbon/v2"

// By gitee
go get -u gitee.com/dromara/carbon/v2
import "gitee.com/dromara/carbon/v2"

// By gitcode
go get -u gitcode.com/dromara/carbon/v2
import "gitee.com/dromara/gitcode/v2"
```

`Carbon` was donated to the [dromara](https://dromara.org/ "dromara") organization, the repository URL has changed. If
the previous repository used was `golang-module/carbon`, please replace the original repository with the new repository
in `go.mod`, or execute the following command:

```go
go mod edit -replace github.com/golang-module/carbon/v2 = github.com/dromara/carbon/v2
```

### Example Usage
Default timezone is `UTC`, language locale is `English`, start day of the week is `Monday` and weekend days of the week
are `Saturday` and `Sunday`.

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

For more usage examples, please refer to <a href="https://carbon.go-pkg.com" target="_blank">official document</a>. 

For performance test reports, please refer to [analysis report](test_report.en.md)

## References

* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)
* [nodatime/nodatime](https://github.com/nodatime/nodatime)
* [jinzhu/now](https://github.com/jinzhu/now)
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)
* [jodaOrg/joda-time](https://github.com/jodaOrg/joda-time)
* [arrow-py/arrow](https://github.com/arrow-py/arrow)
* [moment/moment](https://github.com/moment/moment)
* [iamkun/dayjs](https://github.com/iamkun/dayjs)

## Contributors
Thanks to all the following who contributed to `Carbon`:

<a href="https://github.com/dromara/carbon/graphs/contributors"><img src="https://contrib.rocks/image?repo=dromara/carbon&max=100&columns=16" /></a>

## Sponsors

`Carbon` is a non-commercial open source project. If you want to support `Carbon`, you can [buy a cup of coffee](https://carbon.go-pkg.com/sponsor.html) for developer.

## Thanks

`Carbon` had been being developed with GoLand under the free JetBrains Open Source license, I would like to express my
thanks here.

<a href="https://www.jetbrains.com" target="_blank"><img src="https://carbon.go-pkg.com/jetbrains.svg?v=2.6.x" height="50" alt="JetBrains"/></a>

## License

`Carbon` is licensed under the `MIT` License, see the [LICENSE](./LICENSE) file for details.
