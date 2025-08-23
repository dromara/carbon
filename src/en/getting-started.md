---
head:
  - - meta
    - name: description
      content: Installation | A simple, semantic and developer-friendly time package for golang
---

# Installation

> go version >= 1.21

```go
// Via github
go get -u github.com/dromara/carbon/v2
import "github.com/dromara/carbon/v2"

// Via gitee
go get -u gitee.com/dromara/carbon/v2
import "gitee.com/dromara/carbon/v2"

// Via gitcode
go get -u gitcode.com/dromara/carbon/v2
import "gitee.com/dromara/gitcode/v2"
```

`Carbon` was donated to the [dromara](https://dromara.org/ "dromara") organization, the repository URL has changed. If the previous repository used was `golang-module/carbon`, please replace the original repository with the new repository in `go.mod`, or execute the following command:

```go
go mod edit -replace github.com/golang-module/carbon/v2=github.com/dromara/carbon/v2
```