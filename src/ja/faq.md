# よくある質問

1、v2.5.x と v2.6.x のバージョンの違いは何ですか?
>  `v2.5.x` および以下のバージョンは `値` 転送であり、`v2.6.x` および以上のバージョンは `ポインタ` 転送であり、`v2.6.x` および以上のバージョンを使用することを強くお勧めします。

2、window システムの下でバイナリファイルタイムゾーンエラーを展開する

>  window システムにgolang環境がインストールされていない場合は、配備時に `GOROOT/lib/time/zoneinfo.zip：no such file or directory` 異常、原因は window システムにはタイムゾーンファイルが組み込まれていないので、手動でダウンロードして `zoneinfo.zip` パスを指定するだけで、例えば `go/lib/time/zoneinfo.zip`

```go
os.Setenv("ZONEINFO", "./go/lib/time/zoneinfo.zip")
```

3、docker コンテナ配置バイナリファイルタイムゾーンエラー

> `docker` コンテナに `golang` 環境がインストールされていない場合は、配備時に `open/usr/local/go/lib/time/zoneinfo.zip：no such file or directory` 異常、`zoneinfo.zip`をコンテナにコピーするだけ、すなわち `Dockerfile` に追加

```go
COPY ./zoneinfo.zip /usr/local/go /lib/time/zoneinfo.zip
```