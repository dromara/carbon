# Carbon 性能测试报告

## 测试说明

本报告记录当前代码的代表性 Go benchmark 结果。测试覆盖顺序、并发和并行执行；`ns/op`、`B/op` 和 `allocs/op` 均来自 Go 标准 benchmark 输出。

## 测试环境

- **操作系统**: macOS 26.6.2
- **Go 版本**: `go version go1.27.0 darwin/arm64z`
- **CPU**: `Apple M3 Pro`
- **架构**: `arm64`
- **测试工具**: `go test -bench -benchmem`
- **测试命令**: `go test -run '^$' -bench=. -benchmem -benchtime=300ms -count=1`
- **测试模式**: 顺序、并发、并行

> 测试已在 Go 1.27.0 环境下完成，结果基于当前代码、Apple M3 Pro 和上述 benchmark 参数。

## 代表性结果

### 核心操作

| Benchmark | 顺序 ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkNewCarbon` | 4.500 | 11 | 0 |
| `BenchmarkCopy` | 0.03491 | 0 | 0 |
| `BenchmarkCreateFromStdTime` | 4.588 | 11 | 0 |
| `BenchmarkCarbon_Constellation` | 8.711 | 0 | 0 |
| `BenchmarkCarbon_StartOfYear` | 3.531 | 11 | 0 |
| `BenchmarkCarbon_IsValid` | 0.04650 | 0 | 0 |

### 输出和解析

| Benchmark | 顺序 ns/op | B/op | allocs/op |
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

### 辅助函数

| Benchmark | ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkFormat2Layout` | 531.8 | 0 | 0 |
| `BenchmarkFormat2Layout_Simple` | 59.05 | 0 | 0 |
| `BenchmarkFormat2Layout_Complex` | 92.01 | 0 | 0 |
| `BenchmarkFormat2Layout_Memory` | 763.9 | 320 | 2 |
| `BenchmarkParseTimezone` | 117.1 | 0 | 0 |
| `BenchmarkParseDuration` | 2451 | 2081 | 72 |

`parseDuration` 当前不是零分配路径，不能继续使用旧报告中的 `2-15 ns/op` 和 `0 allocs/op` 结论。

### 日历转换

| Benchmark | 顺序 ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkCarbon_Julian` | 5.590 | 4 | 0 |
| `BenchmarkCreateFromJulian` | 9.945 | 18 | 0 |
| `BenchmarkCarbon_Lunar` | 23.75 | 4 | 0 |
| `BenchmarkCreateFromLunar` | 7.157 | 20 | 0 |
| `BenchmarkCarbon_Persian` | 20.84 | 8 | 0 |
| `BenchmarkCreateFromPersian` | 8.244 | 20 | 0 |
| `BenchmarkCarbon_Hebrew` | 134.9 | 4 | 0 |
| `BenchmarkCreateFromHebrew` | 15.91 | 20 | 0 |

### 语言设置

| Benchmark | 顺序 ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkLanguage_SetLocale` | 6.999 | 11 | 0 |
| `BenchmarkCarbon_SetLocale` | 1.146 | 0 | 0 |

并发和并行结果会受到 goroutine 调度、CPU 负载和 `GOMAXPROCS` 影响，不能直接与顺序结果相加。代表性并行结果包括：`BenchmarkParse` 为 `40.47 ns/op`、`BenchmarkParseByFormats` 为 `92.85 ns/op`、`BenchmarkCarbon_Constellation` 为 `106.6 ns/op`。

## 结论

- 当前核心日期操作、比较操作和星座计算的顺序执行大多低于 `30 ns/op`，并保持零分配。
- 解析路径中，单布局解析约为 `15-16 ns/op`；多布局和多格式解析约为 `33-35 ns/op`。
- `parseDuration` 和复杂格式转换是当前代表性结果中分配最多的路径，应以实际 benchmark 为准继续优化。
- 日历转换中 Hebrew 读取路径明显高于 Julian、Lunar 和 Persian，后续优化应单独评估其算法和缓存行为。
- 报告中的数字依赖 CPU、Go 版本、操作系统和 benchmark 参数；比较不同版本时应固定这些条件，并至少重复运行多次。
