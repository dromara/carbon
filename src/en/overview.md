---
head:
  - - meta
    - name: description
      content: Overview | A simple, semantic and developer-friendly golang time processing library
---

# Overview

`Carbon` is a lightweight, semantic, and developer-friendly `golang` time package that doesn't depend on `any` third-party package, has `100%` unit test coverage, and has been officially used by [docker](https://github.com/docker/docker-language-server/blob/main/go.mod#L10 "docker") as well as included by [awesome-go](https://github.com/yinggaozhen/awesome-go-cn#日期和时间 "awesome-go-cn") and [hello-github](https://hellogithub.com/repository/dromara/carbon "hello-github").

<a href="https://github.com/docker/docker-language-server/blob/main/go.mod#L10" target="_blank"><img src="/docker.jpg" width="100%" alt="docker"/></a>

## Feature
- Lightweight & zero dependency：pure `go` implementation, doesn't depend on `any` third-party package, `100%` unit test coverage
- Semantic API: provide concise and elegant chain calls with high readability, ensuring code scalability and reusability
- Powerful time manipulation：support time resolution, time addition and subtraction, time setting, time boundary, time judgment, etc
- Rich time retrieve：retrieve each parts of time (such as year, month, day, hour, minute, second, etc.) and timestamp
- Diversified output format：output time strings of different precision and formats as needed, including ISO8601, RFC series, and custom formats
- Test friendly：support setting test time, freezes current time, facilitates unit testing
- Time zone handling：support setting and obtaining time zones, as well as converting between different time zones
- I18N：support over 30 localized translation languages and allow customization of translation resources
- Error handling：provide error checking mechanism for handling errors such as time resolution

## Repository

[github.com/dromara/carbon](https://github.com/dromara/carbon "github.com/dromara/carbon")

[gitee.com/dromara/carbon](https://gitee.com/dromara/carbon "gitee.com/dromara/carbon")

[gitcode.com/dromara/carbon](https://gitcode.com/dromara/carbon "gitcode.com/dromara/carbon")