# Carbon 성능 테스트 보고서

## 테스트 범위

이 보고서는 현재 코드의 대표적인 Go benchmark 결과를 기록합니다. 순차, 동시, 병렬 실행을 다루며 `ns/op`, `B/op`, `allocs/op` 값은 Go 표준 benchmark 출력에서 가져왔습니다.

## 테스트 환경

- **운영체제**: macOS 26.6.2
- **Go 버전**: `go version go1.27.0 darwin/arm64z`
- **CPU**: `Apple M3 Pro`
- **아키텍처**: `arm64`
- **도구**: `go test -bench -benchmem`
- **명령어**: `go test -run '^$' -bench=. -benchmem -benchtime=300ms -count=1`
- **실행 모드**: 순차, 동시, 병렬

> Go 1.27.0 환경에서 현재 코드, Apple M3 Pro 및 위 benchmark 매개변수로 측정했습니다.

## 대표 결과

### 핵심 작업

| Benchmark | 순차 ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkNewCarbon` | 4.500 | 11 | 0 |
| `BenchmarkCopy` | 0.03491 | 0 | 0 |
| `BenchmarkCreateFromStdTime` | 4.588 | 11 | 0 |
| `BenchmarkCarbon_Constellation` | 8.711 | 0 | 0 |
| `BenchmarkCarbon_StartOfYear` | 3.531 | 11 | 0 |
| `BenchmarkCarbon_IsValid` | 0.04650 | 0 | 0 |

### 출력 및 파싱

| Benchmark | 순차 ns/op | B/op | allocs/op |
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

### 헬퍼 함수

| Benchmark | ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkFormat2Layout` | 531.8 | 0 | 0 |
| `BenchmarkFormat2Layout_Simple` | 59.05 | 0 | 0 |
| `BenchmarkFormat2Layout_Complex` | 92.01 | 0 | 0 |
| `BenchmarkFormat2Layout_Memory` | 763.9 | 320 | 2 |
| `BenchmarkParseTimezone` | 117.1 | 0 | 0 |
| `BenchmarkParseDuration` | 2451 | 2081 | 72 |

`parseDuration`는 현재 제로 할당 경로가 아닙니다. 이전 보고서의 `2-15 ns/op`, `0 allocs/op` 수치는 더 이상 사용하면 안 됩니다.

### 달력 변환

| Benchmark | 순차 ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkCarbon_Julian` | 5.590 | 4 | 0 |
| `BenchmarkCreateFromJulian` | 9.945 | 18 | 0 |
| `BenchmarkCarbon_Lunar` | 23.75 | 4 | 0 |
| `BenchmarkCreateFromLunar` | 7.157 | 20 | 0 |
| `BenchmarkCarbon_Persian` | 20.84 | 8 | 0 |
| `BenchmarkCreateFromPersian` | 8.244 | 20 | 0 |
| `BenchmarkCarbon_Hebrew` | 134.9 | 4 | 0 |
| `BenchmarkCreateFromHebrew` | 15.91 | 20 | 0 |

### 로케일 설정

| Benchmark | 순차 ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| `BenchmarkLanguage_SetLocale` | 6.999 | 11 | 0 |
| `BenchmarkCarbon_SetLocale` | 1.146 | 0 | 0 |

동시 및 병렬 결과는 goroutine 스케줄링, CPU 부하, `GOMAXPROCS`의 영향을 받으므로 순차 결과와 합산하면 안 됩니다. 대표적인 병렬 결과는 `BenchmarkParse` `40.47 ns/op`, `BenchmarkParseByFormats` `92.85 ns/op`, `BenchmarkCarbon_Constellation` `106.6 ns/op`입니다.

## 결론

- 핵심 날짜 작업, 비교 작업, 별자리 계산의 대부분은 순차 실행에서 `30 ns/op` 미만이며 할당이 없습니다.
- 단일 레이아웃 파싱은 약 `15-16 ns/op`, 다중 레이아웃 및 다중 포맷 파싱은 약 `33-35 ns/op`입니다.
- 대표 결과에서 `parseDuration`과 복잡한 포맷 변환이 가장 많은 할당을 발생시키는 경로입니다. 최적화는 실측값을 기준으로 진행해야 합니다.
- Hebrew 변환 읽기 경로는 Julian, Lunar, Persian보다 뚜렷하게 느리므로 알고리즘과 캐시 동작을 별도로 평가해야 합니다.
- 결과는 CPU, Go 버전, 운영체제, benchmark 매개변수에 의존합니다. 버전 비교 시 조건을 고정하고 여러 번 실행하십시오.
