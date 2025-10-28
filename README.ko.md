<p align="center" style="margin-bottom: -10px"><a href="https://carbon.go-pkg.com/ko" target="_blank"><img src="https://carbon.go-pkg.com/logo.svg?v=2.6.x" width="15%" alt="carbon" /></a></p>

[![Carbon Release](https://img.shields.io/github/release/dromara/carbon.svg)](https://github.com/dromara/carbon/releases)
[![Go Test](https://github.com/dromara/carbon/actions/workflows/test.yml/badge.svg)](https://github.com/dromara/carbon/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/dromara/carbon/v2)](https://goreportcard.com/report/github.com/dromara/carbon/v2)
[![Go Coverage](https://codecov.io/gh/dromara/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/dromara/carbon)
[![Carbon Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/dromara/carbon/v2)
[![Awesome](https://awesome.re/badge-flat2.svg)](https://github.com/avelino/awesome-go#date-and-time)
[![HelloGitHub](https://api.hellogithub.com/v1/widgets/recommend.svg?rid=0eddd8c3469549b7b246f85a83d1c42e&claim_uid=kKBvMpyxSgLhmJO&theme=small)](https://hellogithub.com/en/repository/dromara/carbon)
[![License](https://img.shields.io/github/license/dromara/carbon)](https://github.com/dromara/carbon/blob/master/LICENSE)

한국어 | [English](README.md) | [简体中文](README.zh.md) | [日本語](README.ja.md)

## 소개

`Carbon`은 `golang`을 위한 간단하고, 의미론적이며, 개발자 친화적인 시간 패키지입니다. `100%` 단위 테스트 커버리지를 제공하며, `어떤` 서드파티 패키지에도 의존하지 않으며, [awesome-go](https://github.com/avelino/awesome-go#date-and-time "awesome-go")와 [hello-github](https://hellogithub.com/en/repository/dromara/carbon "hello-github")에 포함되어 있습니다.

## 저장소

[github.com/dromara/carbon](https://github.com/dromara/carbon "github.com/dromara/carbon")

[gitee.com/dromara/carbon](https://gitee.com/dromara/carbon "gitee.com/dromara/carbon")

[gitcode.com/dromara/carbon](https://gitcode.com/dromara/carbon "gitcode.com/dromara/carbon")

## 빠른 시작

### 설치
> go version >= 1.18

```go
// GitHub를 통해
go get -u github.com/dromara/carbon/v2
import "github.com/dromara/carbon/v2"

// Gitee를 통해
go get -u gitee.com/dromara/carbon/v2
import "gitee.com/dromara/carbon/v2"

// GitCode를 통해
go get -u gitcode.com/dromara/carbon/v2
import "gitee.com/dromara/gitcode/v2"
```

`Carbon`은 [dromara](https://dromara.org/ "dromara") 조직에 기부되었으며, 저장소 URL이 변경되었습니다. 이전에 사용하던 저장소가 `golang-module/carbon`이었다면, `go.mod`에서 원래 저장소를 새 저장소로 교체하거나 다음 명령을 실행하세요:

```go
go mod edit -replace github.com/golang-module/carbon/v2 = github.com/dromara/carbon/v2
```

### 사용 예시
기본 시간대는 `UTC`이고, 언어 로케일은 `English`이며, 주의 시작일은 `Monday`이고 주말은 `Saturday`와 `Sunday`입니다.

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

더 많은 사용 예시는 <a href="https://carbon.go-pkg.com/ko" target="_blank">공식 문서</a>를 참조하세요. 성능 테스트 보고서는 [분석 보고서](docs/BENCHMARK.ko.md)를 참조하세요.

## 참고 자료

* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)
* [nodatime/nodatime](https://github.com/nodatime/nodatime)
* [jinzhu/now](https://github.com/jinzhu/now)
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)
* [jodaOrg/joda-time](https://github.com/jodaOrg/joda-time)
* [arrow-py/arrow](https://github.com/arrow-py/arrow)
* [moment/moment](https://github.com/moment/moment)
* [iamkun/dayjs](https://github.com/iamkun/dayjs)

## 기여자
`Carbon`에 기여한 모든 분들께 감사드립니다:

<a href="https://github.com/dromara/carbon/graphs/contributors"><img src="https://contrib.rocks/image?repo=dromara/carbon&max=100&columns=16" /></a>

[Carbon에 새로운 로컬라이즈드 언어 지원을 추가하는 방법](docs/CONTRIBUTING.ko.md)

## 스폰서

`Carbon`은 비상업적 오픈소스 프로젝트입니다. `Carbon`을 지원하고 싶으시다면 개발자에게 [커피 한 잔을 사주세요](https://carbon.go-pkg.com/ko/sponsor.html).

## 감사의 말

`Carbon`은 JetBrains 오픈소스 라이선스의 무료 GoLand로 개발되었습니다. 여기서 감사의 말을 전하고 싶습니다.

<a href="https://www.jetbrains.com" target="_blank"><img src="https://carbon.go-pkg.com/jetbrains.svg?v=2.6.x" height="50" alt="JetBrains"/></a>

## 라이선스

`Carbon`은 `MIT` 라이선스 하에 제공됩니다. 자세한 내용은 [LICENSE](./LICENSE) 파일을 참조하세요.
