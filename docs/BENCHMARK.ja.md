# Carbon パフォーマンステストレポート

## テスト範囲

このレポートは、現在のコードに対する代表的な Go benchmark の結果を記録します。順次実行、同時実行、並列実行を対象とし、`ns/op`、`B/op`、`allocs/op` は Go 標準ベンチマークの出力値です。

## テスト環境

- **オペレーティングシステム**: macOS 26.6.2
- **Go バージョン**: `go version go1.27.0 darwin/arm64z`
- **CPU**: `Apple M3 Pro`
- **アーキテクチャ**: `arm64`
- **ツール**: `go test -bench -benchmem`
- **コマンド**: `go test -run '^$' -bench=. -benchmem -benchtime=300ms -count=1`
- **実行モード**: 順次、同時、並列

> Go 1.27.0 環境で、現在のコード、Apple M3 Pro、上記の benchmark パラメータを使用して測定しました。

## 代表的な結果

### コア操作

| Benchmark | 順次 ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkNewCarbon` | 4.500 | 11 | 0 |
| `BenchmarkCopy` | 0.03491 | 0 | 0 |
| `BenchmarkCreateFromStdTime` | 4.588 | 11 | 0 |
| `BenchmarkCarbon_Constellation` | 8.711 | 0 | 0 |
| `BenchmarkCarbon_StartOfYear` | 3.531 | 11 | 0 |
| `BenchmarkCarbon_IsValid` | 0.04650 | 0 | 0 |

### 出力と解析

| Benchmark | 順次 ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkCarbon_GoString` | 6.894 | 7 | 0 |
| `BenchmarkCarbon_ToString` | 15.26 | 4 | 0 |
| `BenchmarkCarbon_Layout` | 9.348 | 2 | 0 |
| `BenchmarkCarbon_Format` | 19.15 | 13 | 0 |
| `BenchmarkParse` | 15.28 | 11 | 0 |
| `BenchmarkParseByLayout` | 14.97 | 11 | 0 |
| `BenchmarkParseByFormat` | 16.03 | 11 | 0 |
| `BenchmarkParseByLayouts` | 30.85 | 28 | 0 |
| `BenchmarkParseByFormats` | 33.52 | 27 | 0 |

### ヘルパー関数

| Benchmark | ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkFormat2Layout` | 531.8 | 0 | 0 |
| `BenchmarkFormat2Layout_Simple` | 59.05 | 0 | 0 |
| `BenchmarkFormat2Layout_Complex` | 92.01 | 0 | 0 |
| `BenchmarkFormat2Layout_Memory` | 763.9 | 320 | 2 |
| `BenchmarkParseTimezone` | 117.1 | 0 | 0 |
| `BenchmarkParseDuration` | 2451 | 2081 | 72 |

`parseDuration` は現在ゼロアロケーションの経路ではありません。旧レポートの `2-15 ns/op` と `0 allocs/op` という記述は使用しないでください。

### カレンダー変換

| Benchmark | 順次 ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkCarbon_Julian` | 5.590 | 4 | 0 |
| `BenchmarkCreateFromJulian` | 9.945 | 18 | 0 |
| `BenchmarkCarbon_Lunar` | 23.75 | 4 | 0 |
| `BenchmarkCreateFromLunar` | 7.157 | 20 | 0 |
| `BenchmarkCarbon_Persian` | 20.84 | 8 | 0 |
| `BenchmarkCreateFromPersian` | 8.244 | 20 | 0 |
| `BenchmarkCarbon_Hebrew` | 134.9 | 4 | 0 |
| `BenchmarkCreateFromHebrew` | 15.91 | 20 | 0 |

### ロケール設定

| Benchmark | 順次 ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkLanguage_SetLocale` | 6.999 | 11 | 0 |
| `BenchmarkCarbon_SetLocale` | 1.146 | 0 | 0 |

同時実行と並列実行の結果は goroutine のスケジューリング、CPU 負荷、`GOMAXPROCS` の影響を受けるため、順次実行の結果と加算できません。代表的な並列結果は、`BenchmarkParse` が `40.47 ns/op`、`BenchmarkParseByFormats` が `92.85 ns/op`、`BenchmarkCarbon_Constellation` が `106.6 ns/op` です。

## 結論

- コアの日付操作、比較操作、星座計算の多くは、順次実行で `30 ns/op` 未満かつゼロアロケーションです。
- 単一レイアウトの解析は約 `15-16 ns/op`、複数レイアウトと複数フォーマットの解析は約 `33-35 ns/op` です。
- 代表的な結果では、`parseDuration` と複雑なフォーマット変換が最も多くアロケーションする経路です。最適化は実測値を基準にしてください。
- Hebrew 変換の読み取り経路は Julian、Lunar、Persian より明らかに遅く、アルゴリズムとキャッシュの挙動を個別に評価する必要があります。
- 結果は CPU、Go バージョン、OS、benchmark パラメータに依存します。バージョンを比較する際は条件を固定し、複数回実行してください。
