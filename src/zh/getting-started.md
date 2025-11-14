---
head:
  - - meta
    - name: description
      content: Carbon 快速开始指南，提供安装配置说明（支持 GitHub/Gitee/Gitcode 仓库）、导入方法、系统要求（Go >= 1.19）和仓库迁移指南（从 golang-module/carbon 迁移到 dromara/carbon）
  - - meta
    - name: keywords
      content: carbon, go-carbon, 快速开始, 安装指南, 配置说明, Go 安装, go get, 仓库地址, GitHub, Gitee, Gitcode, dromara, 迁移指南, 系统要求
---

# 快速开始
> go version >= 1.18

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

`carbon` 已经捐赠给了 [dromara](https://dromara.org/ "dromara") 开源组织，仓库地址发生了改变，如果之前用的路径是
`golang-module/carbon`，请在 `go.mod` 里将原地址更换为新路径，或执行如下命令

```go
go mod edit -replace github.com/golang-module/carbon/v2 = github.com/dromara/carbon/v2
```