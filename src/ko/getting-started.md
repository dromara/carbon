---
head:
  - - meta
    - name: description
      content: 빠른 시작|가벼우면서도 의미론적이고 개발자 친화적인 golang 시간 처리 라이브러리
---

# 빠른 시작
> go version >= 1.19

```go
// github 라이브러리 사용
go get -u github.com/dromara/carbon/v2
import "github.com/dromara/carbon/v2"

// gitee 라이브러리 사용
go get -u gitee.com/dromara/carbon/v2
import "gitee.com/dromara/carbon/v2"

// gitcode 라이브러리 사용
go get -u gitcode.com/dromara/carbon/v2
import "gitcode.com/dromara/carbon/v2"
```

`carbon`은 [dromara](https://dromara.org/ "dromara") 오픈소스 조직에 기부되었으며, 저장소 주소가 변경되었습니다. 이전에 사용하던 경로가
`golang-module/carbon`인 경우, `go.mod`에서 기존 주소를 새 경로로 변경하거나 다음 명령을 실행하세요

```go
go mod edit -replace github.com/golang-module/carbon/v2 = github.com/dromara/carbon/v2
``` 