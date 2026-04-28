<!--@nrg.languages=en,cn,ja,ko-->
<!--@nrg.defaultLanguage=en-->
<p align="center" style="margin-bottom: -10px"><a href="https://carbon.go-pkg.com/zh" target="_blank"><img src="https://carbon.go-pkg.com/logo.svg?v=2.6.x" width="15%" alt="carbon" /></a></p><!--en-->
<!--en-->
[![Carbon Release](https://img.shields.io/github/release/dromara/carbon.svg)](https://github.com/dromara/carbon/releases)<!--en-->
[![Go Test](https://github.com/dromara/carbon/actions/workflows/test.yml/badge.svg)](https://github.com/dromara/carbon/actions)<!--en-->
[![Go Report Card](https://goreportcard.com/badge/github.com/dromara/carbon/v2)](https://goreportcard.com/report/github.com/dromara/carbon/v2)<!--en-->
[![Go Coverage](https://codecov.io/gh/dromara/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/dromara/carbon)<!--en-->
[![Carbon Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/dromara/carbon/v2)<!--en-->
[![Awesome](https://awesome.re/badge-flat2.svg)](https://github.com/avelino/awesome-go#date-and-time)<!--en-->
[![HelloGitHub](https://api.hellogithub.com/v1/widgets/recommend.svg?rid=0eddd8c3469549b7b246f85a83d1c42e&claim_uid=kKBvMpyxSgLhmJO&theme=small)](https://hellogithub.com/en/repository/dromara/carbon)<!--en-->
[![License](https://img.shields.io/github/license/dromara/carbon)](https://github.com/dromara/carbon/blob/master/LICENSE)<!--en-->
<!--en-->
English | [简体中文](README.cn.md) | [日本語](README.ja.md) | [한국어](README.ko.md)<!--en-->
<!--en-->
## Introduction<!--en-->
<!--en-->
`Carbon` is a lightweight, semantic, and developer-friendly `golang` time package that doesn't depend on `any` third-party package, has `100%` unit test coverage, and has been used by [docker](https://github.com/docker/docker-language-server/blob/main/go.mod#L10 "docker") organization as well as included by [awesome-go](https://github.com/yinggaozhen/awesome-go-cn#日期和时间 "awesome-go-cn") and [hello-github](https://hellogithub.com/repository/dromara/carbon "hello-github").<!--en-->
<!--en-->
<a href="https://github.com/docker/docker-language-server/blob/main/go.mod#L10" target="_blank"><img src="https://carbon.go-pkg.com/docker.jpg" width="100%" alt="docker"/></a><!--en-->
<!--en-->
## Repository<!--en-->
<!--en-->
[github.com/dromara/carbon](https://github.com/dromara/carbon "github.com/dromara/carbon")<!--en-->
<!--en-->
[gitee.com/dromara/carbon](https://gitee.com/dromara/carbon "gitee.com/dromara/carbon")<!--en-->
<!--en-->
[gitcode.com/dromara/carbon](https://gitcode.com/dromara/carbon "gitcode.com/dromara/carbon")<!--en-->
<!--en-->
## Quick Start<!--en-->
<!--en-->
### Installation<!--en-->
> go version >= 1.18<!--en-->
<!--en-->
```go<!--en-->
// Via github<!--en-->
go get -u github.com/dromara/carbon/v2<!--en-->
import "github.com/dromara/carbon/v2"<!--en-->
<!--en-->
// Via gitee<!--en-->
go get -u gitee.com/dromara/carbon/v2<!--en-->
import "gitee.com/dromara/carbon/v2"<!--en-->
<!--en-->
// Via gitcode<!--en-->
go get -u gitcode.com/dromara/carbon/v2<!--en-->
import "gitcode.com/dromara/gitcode/v2"<!--en-->
```<!--en-->
<!--en-->
`Carbon` was donated to the [dromara](https://dromara.org/ "dromara") organization, the repository URL has changed. If<!--en-->
the previous repository used was `golang-module/carbon`, please replace the original repository with the new repository<!--en-->
in `go.mod`, or execute the following command:<!--en-->
<!--en-->
```go<!--en-->
go mod edit -replace github.com/golang-module/carbon/v2 = github.com/dromara/carbon/v2<!--en-->
```<!--en-->
<!--en-->
### Example Usage<!--en-->
Default timezone is `UTC`, language locale is `English`, start day of the week is `Monday` and weekend days of the week<!--en-->
are `Saturday` and `Sunday`.<!--en-->
<!--en-->
```go<!--en-->
carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15.999999999"))<!--en-->
carbon.IsTestNow() // true<!--en-->
<!--en-->
carbon.Now().ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC<!--en-->
carbon.Yesterday().ToString() // 2020-08-04 13:14:15.999999999 +0000 UTC<!--en-->
carbon.Tomorrow().ToString() // 2020-08-06 13:14:15.999999999 +0000 UTC<!--en-->
<!--en-->
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0000 UTC<!--en-->
carbon.Parse("2022-03-08T03:01:14-07:00").ToString() // 2022-03-08 10:01:14 +0000 UTC<!--en-->
<!--en-->
carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToString() // 2020-08-05 13:14:15 +0000 UTC<!--en-->
carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToString() // 2020-08-05 13:14:15 +0000 UTC<!--en-->
<!--en-->
carbon.CreateFromDate(2020, 8, 5).ToString() // 2020-08-05 00:00:00 +0000 UTC<!--en-->
carbon.CreateFromTime(13, 14, 15).ToString() // 2020-08-05 13:14:15 +0000 UTC<!--en-->
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToString() // 2020-08-05 13:14:15 +0000 UTC<!--en-->
carbon.CreateFromTimestamp(1596633255).ToString() // 2020-08-05 13:14:15 +0000 UTC<!--en-->
<!--en-->
carbon.Parse("2020-07-05 13:14:15").DiffForHumans() // 1 month before<!--en-->
carbon.Parse("2020-07-05 13:14:15").SetLocale("zh-CN").DiffForHumans() // 1 月前<!--en-->
<!--en-->
carbon.ClearTestNow()<!--en-->
carbon.IsTestNow() // false<!--en-->
```<!--en-->
<!--en-->
For more usage examples, please refer to <a href="https://carbon.go-pkg.com" target="_blank">official document</a>. <!--en-->
<!--en-->
For performance test reports, please refer to [benchmark report](docs/BENCHMARK.en.md)<!--en-->
<!--en-->
## References<!--en-->
<!--en-->
* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)<!--en-->
* [nodatime/nodatime](https://github.com/nodatime/nodatime)<!--en-->
* [jinzhu/now](https://github.com/jinzhu/now)<!--en-->
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)<!--en-->
* [jodaOrg/joda-time](https://github.com/jodaOrg/joda-time)<!--en-->
* [arrow-py/arrow](https://github.com/arrow-py/arrow)<!--en-->
* [moment/moment](https://github.com/moment/moment)<!--en-->
* [iamkun/dayjs](https://github.com/iamkun/dayjs)<!--en-->
<!--en-->
## Contributors<!--en-->
Thanks to all the following who contributed to `Carbon`:<!--en-->
<!--en-->
<a href="https://github.com/dromara/carbon/graphs/contributors"><img src="https://contrib.rocks/image?repo=dromara/carbon&max=100&columns=16" /></a><!--en-->
<!--en-->
## Translators<!--en-->
Ask for help to translate `Carbon` in other localized languages<!--en-->
<!--en-->
[How to add new localized language support to carbon](https://carbon.go-pkg.com/appendix/contribution-guide.html)<!--en-->
<!--en-->
## Sponsors<!--en-->
<!--en-->
`Carbon` is a non-commercial open source project. If you want to support `Carbon`, you can [buy a cup of coffee](https://carbon.go-pkg.com/sponsor.html) for developer.<!--en-->
<!--en-->
## Thanks<!--en-->
<!--en-->
`Carbon` had been being developed with GoLand under the free JetBrains Open Source license, I would like to express my<!--en-->
thanks here.<!--en-->
<!--en-->
<a href="https://www.jetbrains.com" target="_blank"><img src="https://carbon.go-pkg.com/jetbrains.svg?v=2.6.x" height="50" alt="JetBrains"/></a><!--en-->
<!--en-->
## License<!--en-->
<!--en-->
`Carbon` is licensed under the `MIT` License, see the [LICENSE](./LICENSE) file for details.<!--en-->
<p align="center" style="margin-bottom: -10px"><a href="https://carbon.go-pkg.com/zh" target="_blank"><img src="https://carbon.go-pkg.com/logo.svg?v=2.6.x" width="15%" alt="carbon" /></a></p><!--cn-->
<!--cn-->
[![Carbon Release](https://img.shields.io/github/release/dromara/carbon.svg)](https://github.com/dromara/carbon/releases)<!--cn-->
[![Go Test](https://github.com/dromara/carbon/actions/workflows/test.yml/badge.svg)](https://github.com/dromara/carbon/actions)<!--cn-->
[![Go Report Card](https://goreportcard.com/badge/github.com/dromara/carbon/v2)](https://goreportcard.com/report/github.com/dromara/carbon/v2)<!--cn-->
[![Go Coverage](https://codecov.io/gh/dromara/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/dromara/carbon)<!--cn-->
[![Carbon Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/dromara/carbon/v2)<!--cn-->
[![Awesome](https://awesome.re/badge-flat2.svg)](https://github.com/avelino/awesome-go#date-and-time)<!--cn-->
[![HelloGitHub](https://api.hellogithub.com/v1/widgets/recommend.svg?rid=0eddd8c3469549b7b246f85a83d1c42e&claim_uid=kKBvMpyxSgLhmJO&theme=small)](https://hellogithub.com/en/repository/dromara/carbon)<!--cn-->
[![License](https://img.shields.io/github/license/dromara/carbon)](https://github.com/dromara/carbon/blob/master/LICENSE)<!--cn-->
<!--cn-->
简体中文 | [English](README.md) | [日本語](README.ja.md) | [한국어](README.ko.md)<!--cn-->
<!--cn-->
## 项目简介<!--cn-->
<!--cn-->
`Carbon` 是一个轻量级、语义化、对开发者友好的 `golang` 时间处理库，不依赖于 `任何` 第三方库， `100%` 单元测试覆盖率，已被 [docker](https://github.com/docker/docker-language-server/blob/main/go.mod#L10 "docker") 组织使用以及被 [awesome-go](https://github.com/yinggaozhen/awesome-go-cn#日期和时间 "awesome-go-cn") 和 [hello-github](https://hellogithub.com/repository/dromara/carbon "hello-github") 收录，并获得<!--cn-->
`gitee` 2024 年最有价值项目（`GVP`）和 `gitcode` 2024 年度开源摘星计划 (`G-Star`) 项目<!--cn-->
<!--cn-->
<a href="https://github.com/docker/docker-language-server/blob/main/go.mod#L10" target="_blank"><img src="https://carbon.go-pkg.com/docker.jpg" width="100%" alt="docker"/></a><!--cn-->
<a href="https://gitee.com/dromara/carbon" target="_blank"><img src="https://carbon.go-pkg.com/gvp.jpg" width="100%" alt="gvp"/></a><!--cn-->
<a href="https://gitcode.com/dromara/carbon" target="_blank"><img src="https://carbon.go-pkg.com/gstar.jpg" width="100%" alt="g-star"/></a><!--cn-->
<!--cn-->
## 仓库地址<!--cn-->
<!--cn-->
[github.com/dromara/carbon](https://github.com/dromara/carbon "github.com/dromara/carbon")<!--cn-->
<!--cn-->
[gitee.com/dromara/carbon](https://gitee.com/dromara/carbon "gitee.com/dromara/carbon")<!--cn-->
<!--cn-->
[gitcode.com/dromara/carbon](https://gitcode.com/dromara/carbon "gitcode.com/dromara/carbon")<!--cn-->
<!--cn-->
## 快速开始<!--cn-->
<!--cn-->
### 安装使用<!--cn-->
<!--cn-->
> go version >= 1.18<!--cn-->
<!--cn-->
```go<!--cn-->
// 使用 github 库<!--cn-->
go get -u github.com/dromara/carbon/v2<!--cn-->
import "github.com/dromara/carbon/v2"<!--cn-->
<!--cn-->
// 使用 gitee 库<!--cn-->
go get -u gitee.com/dromara/carbon/v2<!--cn-->
import "gitee.com/dromara/carbon/v2"<!--cn-->
<!--cn-->
// 使用 gitcode 库<!--cn-->
go get -u gitcode.com/dromara/carbon/v2<!--cn-->
import "gitcode.com/dromara/carbon/v2"<!--cn-->
```<!--cn-->
<!--cn-->
`Carbon` 已经捐赠给了 [dromara](https://dromara.org/ "dromara") 开源组织，仓库地址发生了改变，如果之前用的路径是<!--cn-->
`golang-module/carbon`，请在 `go.mod` 里将原地址更换为新路径，或执行如下命令<!--cn-->
<!--cn-->
```go<!--cn-->
go mod edit -replace github.com/golang-module/carbon/v2 = github.com/dromara/carbon/v2<!--cn-->
```<!--cn-->
<!--cn-->
### 用法示例<!--cn-->
<!--cn-->
默认时区是 `UTC`, 语言环境是 `英语`，一周开始日期是 `周一`，周末是 `周六`和 `周日`。<!--cn-->
<!--cn-->
```go<!--cn-->
carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15.999999999"))<!--cn-->
carbon.IsTestNow() // true<!--cn-->
<!--cn-->
carbon.Now().ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC<!--cn-->
carbon.Yesterday().ToString() // 2020-08-04 13:14:15.999999999 +0000 UTC<!--cn-->
carbon.Tomorrow().ToString() // 2020-08-06 13:14:15.999999999 +0000 UTC<!--cn-->
<!--cn-->
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0000 UTC<!--cn-->
carbon.Parse("2022-03-08T03:01:14-07:00").ToString() // 2022-03-08 10:01:14 +0000 UTC<!--cn-->
<!--cn-->
carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToString() // 2020-08-05 13:14:15 +0000 UTC<!--cn-->
carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToString() // 2020-08-05 13:14:15 +0000 UTC<!--cn-->
<!--cn-->
carbon.CreateFromDate(2020, 8, 5).ToString() // 2020-08-05 00:00:00 +0000 UTC<!--cn-->
carbon.CreateFromTime(13, 14, 15).ToString() // 2020-08-05 13:14:15 +0000 UTC<!--cn-->
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToString() // 2020-08-05 13:14:15 +0000 UTC<!--cn-->
carbon.CreateFromTimestamp(1596633255).ToString() // 2020-08-05 13:14:15 +0000 UTC<!--cn-->
<!--cn-->
carbon.Parse("2020-07-05 13:14:15").DiffForHumans() // 1 month before<!--cn-->
carbon.Parse("2020-07-05 13:14:15").SetLocale("zh-CN").DiffForHumans() // 1 月前<!--cn-->
<!--cn-->
carbon.ClearTestNow()<!--cn-->
carbon.IsTestNow() // false<!--cn-->
```<!--cn-->
更多用法示例请查看 <a href="https://carbon.go-pkg.com/zh" target="_blank">官方文档</a>，性能测试报告请查看 [分析报告](docs/BENCHMARK.cn.md)<!--cn-->
<!--cn-->
## 参考项目<!--cn-->
<!--cn-->
* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)<!--cn-->
* [nodatime/nodatime](https://github.com/nodatime/nodatime)<!--cn-->
* [jinzhu/now](https://github.com/jinzhu/now)<!--cn-->
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)<!--cn-->
* [jodaOrg/joda-time](https://github.com/jodaOrg/joda-time)<!--cn-->
* [arrow-py/arrow](https://github.com/arrow-py/arrow)<!--cn-->
* [moment/moment](https://github.com/moment/moment)<!--cn-->
* [iamkun/dayjs](https://github.com/iamkun/dayjs)<!--cn-->
<!--cn-->
## 贡献者<!--cn-->
<!--cn-->
感谢以下所有为 `Carbon` 做出贡献的人：<!--cn-->
<!--cn-->
<a href="https://github.com/dromara/carbon/graphs/contributors"><img src="https://contrib.rocks/image?repo=dromara/carbon&max=100&columns=16"/></a><!--cn-->
<!--cn-->
## 翻译者<!--cn-->
欢迎帮助将 `Carbon` 翻译为更多本地化语言<!--cn-->
<!--cn-->
[如何为 carbon 添加新的本地化语言支持](https://carbon.go-pkg.com/zh/appendix/contribution-guide.html)<!--cn-->
<!--cn-->
## 赞助<!--cn-->
`Carbon` 是一个非商业开源项目, 如果你想支持 `Carbon`,<!--cn-->
你可以为开发者 [购买一杯咖啡](https://carbon.go-pkg.com/zh/sponsor.html) 或 点击以下赞助广告<!--cn-->
<!--cn-->
- **Easysearch**  <!--cn-->
  企业级的分布式搜索型数据库，ES 国产化的首选替代方案<!--cn-->
 <a href="https://easysearch.cn/" target="_blank"><img src="https://easysearch.cn/img/header/logo.svg" height="50" alt="Easysearch"/></a><!--cn-->
<!--cn-->
- **雨云**  <!--cn-->
  KVM高配版，4核 8G 300M，仅需68元/月起，半年付八折，年付七折，七天无理由退订（送5折券）  <!--cn-->
  [点击购买](https://www.rainyun.com/gopkg_?s=github)<!--cn-->
<!--cn-->
- **林枫云**  <!--cn-->
  9950X高防，4核 8G 10M，仅需168元/月，续费同价  <!--cn-->
  [点击购买](https://www.dkdun.cn/aff/NVHPHCEF)<!--cn-->
<!--cn-->
- **莱卡云**  <!--cn-->
  浙江电信云，2核 4G 10M，仅需38元/月，续费同价  <!--cn-->
  [点击购买](https://www.lcayun.com/aff/MBVNVNFX)<!--cn-->
## 致谢<!--cn-->
<!--cn-->
`Carbon`已获取免费的 JetBrains 开源许可证，在此表示感谢<!--cn-->
<!--cn-->
<a href="https://www.jetbrains.com" target="_blank"><img src="https://carbon.go-pkg.com/jetbrains.svg?v=2.6.x" height="50" alt="JetBrains"/></a><!--cn-->
<!--cn-->
## 开源协议<!--cn-->
<!--cn-->
`Carbon` 遵循 `MIT` 开源协议, 请参阅 [LICENSE](./LICENSE) 查看详细信息。<!--cn-->
<p align="center" style="margin-bottom: -10px"><a href="https://carbon.go-pkg.com/zh" target="_blank"><img src="https://carbon.go-pkg.com/logo.svg?v=2.6.x" width="15%" alt="carbon" /></a></p><!--ja-->
<!--ja-->
[![Carbon Release](https://img.shields.io/github/release/dromara/carbon.svg)](https://github.com/dromara/carbon/releases)<!--ja-->
[![Go Test](https://github.com/dromara/carbon/actions/workflows/test.yml/badge.svg)](https://github.com/dromara/carbon/actions)<!--ja-->
[![Go Report Card](https://goreportcard.com/badge/github.com/dromara/carbon/v2)](https://goreportcard.com/report/github.com/dromara/carbon/v2)<!--ja-->
[![Go Coverage](https://codecov.io/gh/dromara/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/dromara/carbon)<!--ja-->
[![Carbon Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/dromara/carbon/v2)<!--ja-->
[![Awesome](https://awesome.re/badge-flat2.svg)](https://github.com/avelino/awesome-go#date-and-time)<!--ja-->
[![HelloGitHub](https://api.hellogithub.com/v1/widgets/recommend.svg?rid=0eddd8c3469549b7b246f85a83d1c42e&claim_uid=kKBvMpyxSgLhmJO&theme=small)](https://hellogithub.com/en/repository/dromara/carbon)<!--ja-->
[![License](https://img.shields.io/github/license/dromara/carbon)](https://github.com/dromara/carbon/blob/master/LICENSE)<!--ja-->
<!--ja-->
日本語 | [English](README.md) | [简体中文](README.cn.md) | [한국어](README.ko.md)<!--ja-->
<!--ja-->
## イントロ<!--ja-->
<!--ja-->
`Carbon` は軽量でセマンティックで開発者に優しい `golang` 時間処理ライブラリで、サードパーティ製ライブラリに依存せず、ユニットテストのカバー率は `100%` で、[docker](https://github.com/docker/docker-language-server/blob/main/go.mod#L10 "docker") に公式採用され、[awesome-go](https://github.com/yinggaozhen/awesome-go-cn#日期和时间 "awesome-go-cn") と [hello-github](https://hellogithub.com/repository/dromara/carbon "hello-github") にも収録されています。<!--ja-->
<!--ja-->
<a href="https://github.com/docker/docker-language-server/blob/main/go.mod#L10" target="_blank"><img src="https://carbon.go-pkg.com/docker.jpg" width="100%" alt="docker"/></a><!--ja-->
<!--ja-->
## リポジトリ<!--ja-->
<!--ja-->
[github.com/dromara/carbon](https://github.com/dromara/carbon "github.com/dromara/carbon")<!--ja-->
<!--ja-->
[gitee.com/dromara/carbon](https://gitee.com/dromara/carbon "gitee.com/dromara/carbon")<!--ja-->
<!--ja-->
[gitcode.com/dromara/carbon](https://gitcode.com/dromara/carbon "gitcode.com/dromara/carbon")<!--ja-->
<!--ja-->
## クイックスタート<!--ja-->
<!--ja-->
### インストール<!--ja-->
> go version >= 1.18<!--ja-->
<!--ja-->
```go<!--ja-->
// github から使う<!--ja-->
go get -u github.com/dromara/carbon/v2<!--ja-->
import "github.com/dromara/carbon/v2"<!--ja-->
<!--ja-->
// gitee から使う<!--ja-->
go get -u gitee.com/dromara/carbon/v2<!--ja-->
import "gitee.com/dromara/carbon/v2"<!--ja-->
<!--ja-->
// gitcode から使う<!--ja-->
go get -u gitcode.com/dromara/carbon/v2<!--ja-->
import "gitcode.com/dromara/gitcode/v2"<!--ja-->
```<!--ja-->
<!--ja-->
`Carbon` は [dromara](https://dromara.org/ "dromara") 組織に寄付されたためリポジトリのURLが変更されました。以前のリポジトリ `golang-module/carbon` を使用している場合は`go.mod`で新しいリポジトリURLに変更するか下記コマンドを実行します<!--ja-->
<!--ja-->
```go<!--ja-->
go mod edit -replace github.com/golang-module/carbon/v2 = github.com/dromara/carbon/v2<!--ja-->
```<!--ja-->
<!--ja-->
### 使い方と例<!--ja-->
デフォルトのタイムゾーンは` UTC `、ロケールは`英語`、週の開始日は`月曜日`、週末は`土曜日`、`日曜日`。<!--ja-->
<!--ja-->
```go<!--ja-->
carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15.999999999"))<!--ja-->
carbon.IsTestNow() // true<!--ja-->
<!--ja-->
carbon.Now().ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC<!--ja-->
carbon.Yesterday().ToString() // 2020-08-04 13:14:15.999999999 +0000 UTC<!--ja-->
carbon.Tomorrow().ToString() // 2020-08-06 13:14:15.999999999 +0000 UTC<!--ja-->
<!--ja-->
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0000 UTC<!--ja-->
carbon.Parse("2022-03-08T03:01:14-07:00").ToString() // 2022-03-08 10:01:14 +0000 UTC<!--ja-->
<!--ja-->
carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToString() // 2020-08-05 13:14:15 +0000 UTC<!--ja-->
carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToString() // 2020-08-05 13:14:15 +0000 UTC<!--ja-->
<!--ja-->
carbon.CreateFromDate(2020, 8, 5).ToString() // 2020-08-05 00:00:00 +0000 UTC<!--ja-->
carbon.CreateFromTime(13, 14, 15).ToString() // 2020-08-05 13:14:15 +0000 UTC<!--ja-->
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToString() // 2020-08-05 13:14:15 +0000 UTC<!--ja-->
carbon.CreateFromTimestamp(1596633255).ToString() // 2020-08-05 13:14:15 +0000 UTC<!--ja-->
<!--ja-->
carbon.Parse("2020-07-05 13:14:15").DiffForHumans() // 1 month before<!--ja-->
carbon.Parse("2020-07-05 13:14:15").SetLocale("zh-CN").DiffForHumans() // 1 月前<!--ja-->
<!--ja-->
carbon.ClearTestNow()<!--ja-->
carbon.IsTestNow() // false<!--ja-->
```<!--ja-->
より多くの使用例については、<a href="https://carbon.go-pkg.com" target="_blank">公式ドキュメント</a>をご覧ください。<!--ja-->
<!--ja-->
性能テストレポートについては、[ベンチマークレポート](docs/BENCHMARK.ja.md)をご参照ください<!--ja-->
<!--ja-->
## リファレンス<!--ja-->
<!--ja-->
* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)<!--ja-->
* [nodatime/nodatime](https://github.com/nodatime/nodatime)<!--ja-->
* [jinzhu/now](https://github.com/jinzhu/now)<!--ja-->
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)<!--ja-->
* [jodaOrg/joda-time](https://github.com/jodaOrg/joda-time)<!--ja-->
* [arrow-py/arrow](https://github.com/arrow-py/arrow)<!--ja-->
* [moment/moment](https://github.com/moment/moment)<!--ja-->
* [iamkun/dayjs](https://github.com/iamkun/dayjs)<!--ja-->
<!--ja-->
## コントリビューター<!--ja-->
`Carbon` に貢献してくれた以下のすべてに感謝します：<!--ja-->
<!--ja-->
<a href="https://github.com/dromara/carbon/graphs/contributors"><img src="https://contrib.rocks/image?repo=dromara/carbon&max=100&columns=16"/></a><!--ja-->
<!--ja-->
## 翻訳者<!--ja-->
`Carbon` を他のローカライズ言語に翻訳してくださる方を募集しています<!--ja-->
<!--ja-->
[Carbon に新しいローカライズ言語サポートを追加する方法](https://carbon.go-pkg.com/ja/appendix/contribution-guide.html)<!--ja-->
<!--ja-->
## スポンサー<!--ja-->
<!--ja-->
`Carbon` は非営利のオープンソースプロジェクトです，`Carbon` をサポートしたい場合は、開発者のために [コーヒーを1杯購入](https://carbon.go-pkg.com/ja/sponsor.html) できます<!--ja-->
<!--ja-->
## 謝辞<!--ja-->
<!--ja-->
`Carbon` は無料の JetBrains オープンソースライセンスを取得しました，これに感謝します<!--ja-->
<!--ja-->
<a href="https://www.jetbrains.com" target="_blank"><img src="https://carbon.go-pkg.com/jetbrains.svg?v=2.6.x" height="50" alt="JetBrains"/></a><!--ja-->
<!--ja-->
## オープンソースプロトコル<!--ja-->
<!--ja-->
`Carbon` は `MIT` オープンソースプロトコルに準拠しており、詳細は [LICENSE](./LICENSE) を参照してください<!--ja-->
<p align="center" style="margin-bottom: -10px"><a href="https://carbon.go-pkg.com/ko" target="_blank"><img src="https://carbon.go-pkg.com/logo.svg?v=2.6.x" width="15%" alt="carbon" /></a></p><!--ko-->
<!--ko-->
[![Carbon Release](https://img.shields.io/github/release/dromara/carbon.svg)](https://github.com/dromara/carbon/releases)<!--ko-->
[![Go Test](https://github.com/dromara/carbon/actions/workflows/test.yml/badge.svg)](https://github.com/dromara/carbon/actions)<!--ko-->
[![Go Report Card](https://goreportcard.com/badge/github.com/dromara/carbon/v2)](https://goreportcard.com/report/github.com/dromara/carbon/v2)<!--ko-->
[![Go Coverage](https://codecov.io/gh/dromara/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/dromara/carbon)<!--ko-->
[![Carbon Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/dromara/carbon/v2)<!--ko-->
[![Awesome](https://awesome.re/badge-flat2.svg)](https://github.com/avelino/awesome-go#date-and-time)<!--ko-->
[![HelloGitHub](https://api.hellogithub.com/v1/widgets/recommend.svg?rid=0eddd8c3469549b7b246f85a83d1c42e&claim_uid=kKBvMpyxSgLhmJO&theme=small)](https://hellogithub.com/en/repository/dromara/carbon)<!--ko-->
[![License](https://img.shields.io/github/license/dromara/carbon)](https://github.com/dromara/carbon/blob/master/LICENSE)<!--ko-->
<!--ko-->
한국어 | [English](README.md) | [简体中文](README.cn.md) | [日本語](README.ja.md)<!--ko-->
<!--ko-->
## 소개<!--ko-->
<!--ko-->
`Carbon` 은 가벼우면서도 의미론적이고 개발자 친화적인 `golang` 시간 처리 라이브러리로, `어떤` 서드파티 라이브러리에도 의존하지 않으며, `100%` 단위 테스트 커버리지를 가지고 있으며, [docker](https://github.com/docker/docker-language-server/blob/main/go.mod#L10 "docker")에 공식적으로 사용되고 있고 [awesome-go](https://github.com/yinggaozhen/awesome-go-cn#日期和时间 "awesome-go-cn") 와 [hello-github](https://hellogithub.com/repository/dromara/carbon "hello-github")에도 수록되어 있습니다.<!--ko-->
<!--ko-->
<a href="https://github.com/docker/docker-language-server/blob/main/go.mod#L10" target="_blank"><img src="https://carbon.go-pkg.com/docker.jpg" width="100%" alt="docker"/></a><!--ko-->
<!--ko-->
## 저장소<!--ko-->
<!--ko-->
[github.com/dromara/carbon](https://github.com/dromara/carbon "github.com/dromara/carbon")<!--ko-->
<!--ko-->
[gitee.com/dromara/carbon](https://gitee.com/dromara/carbon "gitee.com/dromara/carbon")<!--ko-->
<!--ko-->
[gitcode.com/dromara/carbon](https://gitcode.com/dromara/carbon "gitcode.com/dromara/carbon")<!--ko-->
<!--ko-->
## 빠른 시작<!--ko-->
<!--ko-->
### 설치<!--ko-->
> go version >= 1.18<!--ko-->
<!--ko-->
```go<!--ko-->
// GitHub를 통해<!--ko-->
go get -u github.com/dromara/carbon/v2<!--ko-->
import "github.com/dromara/carbon/v2"<!--ko-->
<!--ko-->
// Gitee를 통해<!--ko-->
go get -u gitee.com/dromara/carbon/v2<!--ko-->
import "gitee.com/dromara/carbon/v2"<!--ko-->
<!--ko-->
// GitCode를 통해<!--ko-->
go get -u gitcode.com/dromara/carbon/v2<!--ko-->
import "gitcode.com/dromara/gitcode/v2"<!--ko-->
```<!--ko-->
<!--ko-->
`Carbon`은 [dromara](https://dromara.org/ "dromara") 조직에 기부되었으며, 저장소 URL이 변경되었습니다. 이전에 사용하던 저장소가 `golang-module/carbon`이었다면, `go.mod`에서 원래 저장소를 새 저장소로 교체하거나 다음 명령을 실행하세요:<!--ko-->
<!--ko-->
```go<!--ko-->
go mod edit -replace github.com/golang-module/carbon/v2 = github.com/dromara/carbon/v2<!--ko-->
```<!--ko-->
<!--ko-->
### 사용 예시<!--ko-->
기본 시간대는 `UTC`이고, 언어 로케일은 `English`이며, 주의 시작일은 `Monday`이고 주말은 `Saturday`와 `Sunday`입니다.<!--ko-->
<!--ko-->
```go<!--ko-->
carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15.999999999"))<!--ko-->
carbon.IsTestNow() // true<!--ko-->
<!--ko-->
carbon.Now().ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC<!--ko-->
carbon.Yesterday().ToString() // 2020-08-04 13:14:15.999999999 +0000 UTC<!--ko-->
carbon.Tomorrow().ToString() // 2020-08-06 13:14:15.999999999 +0000 UTC<!--ko-->
<!--ko-->
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0000 UTC<!--ko-->
carbon.Parse("2022-03-08T03:01:14-07:00").ToString() // 2022-03-08 10:01:14 +0000 UTC<!--ko-->
<!--ko-->
carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToString() // 2020-08-05 13:14:15 +0000 UTC<!--ko-->
carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToString() // 2020-08-05 13:14:15 +0000 UTC<!--ko-->
<!--ko-->
carbon.CreateFromDate(2020, 8, 5).ToString() // 2020-08-05 00:00:00 +0000 UTC<!--ko-->
carbon.CreateFromTime(13, 14, 15).ToString() // 2020-08-05 13:14:15 +0000 UTC<!--ko-->
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToString() // 2020-08-05 13:14:15 +0000 UTC<!--ko-->
carbon.CreateFromTimestamp(1596633255).ToString() // 2020-08-05 13:14:15 +0000 UTC<!--ko-->
<!--ko-->
carbon.Parse("2020-07-05 13:14:15").DiffForHumans() // 1 month before<!--ko-->
carbon.Parse("2020-07-05 13:14:15").SetLocale("zh-CN").DiffForHumans() // 1 月前<!--ko-->
<!--ko-->
carbon.ClearTestNow()<!--ko-->
carbon.IsTestNow() // false<!--ko-->
```<!--ko-->
<!--ko-->
더 많은 사용 예시는 <a href="https://carbon.go-pkg.com" target="_blank">공식 문서</a>를 참조하세요.<!--ko-->
<!--ko-->
성능 테스트 보고서는 [벤치마크 보고서](docs/BENCHMARK.ko.md)를 참조하세요.<!--ko-->
<!--ko-->
## 참고 자료<!--ko-->
<!--ko-->
* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)<!--ko-->
* [nodatime/nodatime](https://github.com/nodatime/nodatime)<!--ko-->
* [jinzhu/now](https://github.com/jinzhu/now)<!--ko-->
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)<!--ko-->
* [jodaOrg/joda-time](https://github.com/jodaOrg/joda-time)<!--ko-->
* [arrow-py/arrow](https://github.com/arrow-py/arrow)<!--ko-->
* [moment/moment](https://github.com/moment/moment)<!--ko-->
* [iamkun/dayjs](https://github.com/iamkun/dayjs)<!--ko-->
<!--ko-->
## 기여자<!--ko-->
`Carbon`에 기여한 모든 분들께 감사드립니다:<!--ko-->
<!--ko-->
<a href="https://github.com/dromara/carbon/graphs/contributors"><img src="https://contrib.rocks/image?repo=dromara/carbon&max=100&columns=16" /></a><!--ko-->
<!--ko-->
## 번역자<!--ko-->
`Carbon` 을 더 많은 로컬라이즈드 언어로 번역하는 데 도움을 요청드립니다<!--ko-->
<!--ko-->
[Carbon에 새로운 로컬라이즈드 언어 지원을 추가하는 방법](https://carbon.go-pkg.com/ko/appendix/contribution-guide.html)<!--ko-->
<!--ko-->
## 스폰서<!--ko-->
<!--ko-->
`Carbon`은 비상업적 오픈소스 프로젝트입니다. `Carbon`을 지원하고 싶으시다면 개발자에게 [커피 한 잔을 사주세요](https://carbon.go-pkg.com/ko/sponsor.html).<!--ko-->
<!--ko-->
## 감사의 말<!--ko-->
<!--ko-->
`Carbon`은 JetBrains 오픈소스 라이선스의 무료 GoLand로 개발되었습니다. 여기서 감사의 말을 전하고 싶습니다.<!--ko-->
<!--ko-->
<a href="https://www.jetbrains.com" target="_blank"><img src="https://carbon.go-pkg.com/jetbrains.svg?v=2.6.x" height="50" alt="JetBrains"/></a><!--ko-->
<!--ko-->
## 라이선스<!--ko-->
<!--ko-->
`Carbon`은 `MIT` 라이선스 하에 제공됩니다. 자세한 내용은 [LICENSE](./LICENSE) 파일을 참조하세요.<!--ko-->
