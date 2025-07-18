# よくある質問

1. v2.5.x と v2.6.x のバージョンの違いは何ですか?
>  `v2.5.x` および以下のバージョンは `値` 転送であり、`v2.6.x` および以上のバージョンは `ポインタ` 転送であり、`v2.6.x` および以上のバージョンを使用することを強くお勧めします。

2. Windows システムでのデプロイ時のタイムゾーンエラー

> Windows システムに golang 環境がインストールされていない場合は、デプロイ時に `GOROOT/lib/time/zoneinfo.zip: no such file or directory` エラーが発生します。原因は Windows システムにはタイムゾーンファイルが組み込まれていないためです。手動でダウンロードして `zoneinfo.zip` パスを指定するだけで解決できます。例：`go/lib/time/zoneinfo.zip`

```go
os.Setenv("ZONEINFO", "./go/lib/time/zoneinfo.zip")
```

3. Docker コンテナでのデプロイ時のタイムゾーンエラー

> Docker コンテナに golang 環境がインストールされていない場合は、デプロイ時に `open /usr/local/go/lib/time/zoneinfo.zip: no such file or directory` エラーが発生します。`zoneinfo.zip` をコンテナにコピーするだけで解決できます。つまり、`Dockerfile` に以下を追加します

```dockerfile
COPY ./zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
```