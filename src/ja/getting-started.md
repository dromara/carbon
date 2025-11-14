---
head:
  - - meta
    - name: description
      content: Carbon クイックスタートガイド、インストールと設定の説明（GitHub/Gitee/Gitcode リポジトリをサポート）、インポート方法、システム要件（Go >= 1.18）およびリポジトリ移行ガイド（golang-module/carbon から dromara/carbon へ）を提供
  - - meta
    - name: keywords
      content: carbon, go-carbon, クイックスタート, インストールガイド, 設定説明, Go インストール, go get, リポジトリアドレス, GitHub, Gitee, Gitcode, dromara, 移行ガイド, システム要件
---

# クイックスタート
> go version >= 1.18

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
go mod edit -replace github.com/golang-module/carbon/v2 = github.com/dromara/carbon/v2
```