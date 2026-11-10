# Carbon Performance Test Report

## Test Scope

This report records representative Go benchmark results for the current code. It covers sequential, concurrent, and parallel execution. The `ns/op`, `B/op`, and `allocs/op` values come from Go's standard benchmark output.

## Test Environment

- **Operating system**: macOS 26.6.2
- **Go version**: `go version go1.27.0 darwin/arm64z`
- **CPU**: `Apple M3 Pro`
- **Architecture**: `arm64`
- **Tool**: `go test -bench -benchmem`
- **Command**: `go test -run '^$' -bench=. -benchmem -benchtime=300ms -count=1`
- **Modes**: sequential, concurrent, parallel

> The benchmark was run with Go 1.27.0 on the current code, Apple M3 Pro, and the parameters above.

## Representative Results

### Core Operations

| Benchmark | Sequential ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkNewCarbon` | 4.500 | 11 | 0 |
| `BenchmarkCopy` | 0.03491 | 0 | 0 |
| `BenchmarkCreateFromStdTime` | 4.588 | 11 | 0 |
| `BenchmarkCarbon_Constellation` | 8.711 | 0 | 0 |
| `BenchmarkCarbon_StartOfYear` | 3.531 | 11 | 0 |
| `BenchmarkCarbon_IsValid` | 0.04650 | 0 | 0 |

### Output and Parsing

| Benchmark | Sequential ns/op | B/op | allocs/op |
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

### Helper Functions

| Benchmark | ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkFormat2Layout` | 531.8 | 0 | 0 |
| `BenchmarkFormat2Layout_Simple` | 59.05 | 0 | 0 |
| `BenchmarkFormat2Layout_Complex` | 92.01 | 0 | 0 |
| `BenchmarkFormat2Layout_Memory` | 763.9 | 320 | 2 |
| `BenchmarkParseTimezone` | 117.1 | 0 | 0 |
| `BenchmarkParseDuration` | 2451 | 2081 | 72 |

`parseDuration` is currently not a zero-allocation path. The old `2-15 ns/op` and `0 allocs/op` claims should no longer be used.

### Calendar Conversion

| Benchmark | Sequential ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkCarbon_Julian` | 5.590 | 4 | 0 |
| `BenchmarkCreateFromJulian` | 9.945 | 18 | 0 |
| `BenchmarkCarbon_Lunar` | 23.75 | 4 | 0 |
| `BenchmarkCreateFromLunar` | 7.157 | 20 | 0 |
| `BenchmarkCarbon_Persian` | 20.84 | 8 | 0 |
| `BenchmarkCreateFromPersian` | 8.244 | 20 | 0 |
| `BenchmarkCarbon_Hebrew` | 134.9 | 4 | 0 |
| `BenchmarkCreateFromHebrew` | 15.91 | 20 | 0 |

### Locale Setting

| Benchmark | Sequential ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkLanguage_SetLocale` | 6.999 | 11 | 0 |
| `BenchmarkCarbon_SetLocale` | 1.146 | 0 | 0 |

Concurrent and parallel results are affected by goroutine scheduling, CPU load, and `GOMAXPROCS`; they must not be added to sequential results. Representative parallel results are `40.47 ns/op` for `BenchmarkParse`, `92.85 ns/op` for `BenchmarkParseByFormats`, and `106.6 ns/op` for `BenchmarkCarbon_Constellation`.

## Conclusions

- Most core date operations, comparisons, and constellation calculations stay below `30 ns/op` sequentially with zero allocations.
- Single-layout parsing is about `15-16 ns/op`; multi-layout and multi-format parsing are about `33-35 ns/op`.
- `parseDuration` and complex format conversion are the highest-allocation paths in the representative results and should be optimized from measured data.
- Hebrew conversion is noticeably slower than Julian, Lunar, and Persian reads and should be evaluated separately for algorithm and caching behavior.
- Results depend on the CPU, Go version, operating system, and benchmark parameters. Keep those conditions fixed and repeat runs before comparing versions.
