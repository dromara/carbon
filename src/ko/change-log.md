---
head:
  - - meta
    - name: description
      content: "Carbon 업데이트 로그, 역사 버전부터의 모든 업데이트 내용을 상세히 기록, 신기능, 성능 최적화, Bug 수정, API 변경, 의존성 업그레이드 및 호환성 설명 포함"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, 업데이트 로그, 버전 역사, 변경 기록, 기능 업데이트, Bug 수정, 성능 최적화, API 변경, 릴리스 노트, 버전 업데이트"
---

# 업데이트 로그

## [v2.7.0](https://github.com/dromara/carbon/compare/v2.6.18...v2.7.0) (2026-10-19)

### ✨ 새로운 기능

- **`AddWeekdays`, `AddWeekday`, `SubWeekdays`, `SubWeekday` 메서드를 추가하여 영업일 계산 시 주말 건너뛰기 지원**: 인스턴스에 설정된 주말을 건너뛰며 날짜를 이동하므로 영업일 기준으로 기한을 계산할 수 있습니다. 네 가지 이름은 `AddDays` / `AddDay` / `SubDays` / `SubDay` 와 일관되며, 주말은 인스턴스에서 읽으므로 `SetWeekendDays` 도 계속 동작합니다. 7일 모두 주말로 설정된 경우 무한 루프를 방지하기 위해 각 영업일 범위를 한 주로 제한합니다 (by [@darakanoit](https://github.com/darakanoit) in [#366](https://github.com/dromara/carbon/pull/366))
- **`IsSameWeek` 메서드를 추가하여 두 날짜가 같은 주에 속하는지 판단**: `IsSame*` 패밀리(`IsSameCentury`, `IsSameDecade`, `IsSameYear`, `IsSameQuarter`, `IsSameWeek`, `IsSameMonth`, `IsSameDay`, `IsSameHour`, `IsSameMinute`, `IsSameSecond`)를 완성합니다. `SetWeekStartsAt` 설정을 존중하므로 한 설정에서 같은 주인 두 날짜가 다른 설정에서는 같은 주가 아닐 수 있습니다. 비교는 주 시작일을 기준으로 하며 연도를 넘는 경계 주도 올바르게 처리합니다 (by [@darakanoit](https://github.com/darakanoit) in [#368](https://github.com/dromara/carbon/pull/368))
- **`IsBirthday` 메서드를 추가하여 월과 일을 비교해 생일 여부 판단**: 연도를 무시하고 같은 월일인지 여부를 반환합니다. 인수가 없으면 현재 시간과 비교합니다. 2월 29일은 2월 29일만 일치하므로 평년에는 생일이 될 수 없으며, PHP Carbon 동작과 일치합니다 (by [@darakanoit](https://github.com/darakanoit) in [#364](https://github.com/dromara/carbon/pull/364))
- **`Asia/Jakarta` 시간대 상수 지원 추가** (by [@mirzaq19](https://github.com/mirzaq19) in [#374](https://github.com/dromara/carbon/pull/374))

### 🐛 버그 수정

- **서머타임 해제 시 `StartOfSecond` 와 `EndOfSecond` 가 잘못된 시각 인스턴스를 선택하는 문제 수정**: 서머타임으로 반복되는 시간대에서 `time.Date` 재구성이 동일한 지역 시각의 다른 인스턴스를 선택할 수 있었습니다. 이제 인스턴스를 복사하고 기존 타임스탬프의 서브초 정밀도를 조정하여 선택된 시점, 위치, 설정을 유지합니다 (by [@mikamikasuki](https://github.com/mikamikasuki) in [#377](https://github.com/dromara/carbon/pull/377))
- **서머타임 해제 시 `StartOfMinute` 와 `EndOfMinute` 가 잘못된 시각 인스턴스를 선택하는 문제 수정**: #377 과 동일한 종류의 버그로 한 단계 위 단위의 문제입니다. 이제 인스턴스를 복사하고 지역 초를 빼서 시점, 위치, 오프셋을 유지합니다 (by [@dualfroz](https://github.com/dualfroz) in [#378](https://github.com/dromara/carbon/pull/378))
- **서머타임 해제 시 `StartOfHour` 와 `EndOfHour` 가 잘못된 시각 인스턴스를 선택하는 문제 수정**: 재구성된 시간 경계에 수신자의 오프셋과 일치하는 유효한 대응 인스턴스가 있으면 수신자의 인스턴스를 유지합니다. 대응 인스턴스가 없는 경우(Lord Howe의 부분 시간 전환 등)에는 기존 달력 재구성 동작을 유지합니다 (by [@dajiaohuang](https://github.com/dajiaohuang) in [#382](https://github.com/dromara/carbon/pull/382))
- **`DiffInSeconds` 가 서브초 부분을 무시하여 초 경계를 넘을 때 부정확하게 계산하는 문제 수정**: 정수 초 타임스탬프를 빼면서 소수 부분이 무시되어 초 경계를 넘는 0.2초 간격이 1초로 계산되었습니다. 이제 나노초 차이의 부호가 반대일 때 ±1로 조정하고, 경과한 정수 초를 0 방향으로 절삭하여 계산합니다 (by [@youdie006](https://github.com/youdie006) in [#379](https://github.com/dromara/carbon/pull/379))
- **`DiffInDays` 와 `DiffInWeeks` 가 서머타임 전환을 넘을 때 잘못된 결과를 계산하는 문제 수정**: 경과 초를 나누는 방식은 서머타임 변경을 넘을 때 잘못된 결과를 주었습니다. 이제 `DiffInMonths` 가 월을 세는 방식(`AddDate` 와 보정)으로 달력 일수를 세고 `DiffInWeeks` 는 이를 사용합니다. 서머타임 전환을 넘는 구간만 결과가 달라집니다 (by [@youdie006](https://github.com/youdie006) in [#376](https://github.com/dromara/carbon/pull/376))
- **`DiffInYears` 와 `DiffInMonths` 가 서로 다른 시간대의 날짜를 비교할 때 잘못된 추정치를 산출하는 문제 수정**: 각 날짜의 `Year()` / `Month()` 필드로 처음 추정할 때 두 날짜가 서로 다른 위치에 있으면 오차가 발생할 수 있었습니다. 이제 먼저 이른 날짜의 위치에서 늦은 날짜를 읽은 뒤 연과 월 필드를 가져와 `DiffInDays` 가 사용하는 틀과 일치시킵니다 (by [@ybai08](https://github.com/ybai08) in [#384](https://github.com/dromara/carbon/pull/384))
- **`IsSameDecade` 가 세기를 비교하지 않아 세기를 넘는 날짜를 잘못 판정하는 문제 수정**: `Decade()` 는 세기 내의 10년 단위이므로 `Parse("1995-08-05").IsSameDecade(Parse("2095-08-05"))` 가 true를 반환했습니다. 이제 `Century()` 도 비교하여 `StartOfDecade` 가 사용하는 구간과 일치시킵니다 (by [@youdie006](https://github.com/youdie006) in [#375](https://github.com/dromara/carbon/pull/375))
- **`FromStdTime` 이 UTC 오프셋 차이로 자정 직후에 잘못된 음력 날짜를 반환하는 문제 수정**: 시간이 속한 시간대에서 1900-01-31 부터의 경과 시간을 계산하여 두 시점의 오프셋이 일수 계산에 누출되었습니다. 이제 UTC에서 벽시계 날짜 사이의 일수를 계산하여 `persian.FromStdTime` 및 `hebrew.FromStdTime` 과 일치시킵니다 (by [@youdie006](https://github.com/youdie006) in [#380](https://github.com/dromara/carbon/pull/380))
- **동결 시계 모드에서 `Sleep` 을 동시 호출하면 시계 진행이 유실되는 경합 상태 수정**: 동결 시계에서의 동시 `Sleep` 호출이 서로 그리고 `Now` 와 경합하여 시계 진행이 유실되었습니다. 이제 기존 동결 시계 뮤텍스로 검사와 갱신을 감쌉니다 (by [@yuefdev](https://github.com/yuefdev) in [#383](https://github.com/dromara/carbon/pull/383))

### 🔧 의존성 업그레이드

- `github/codeql-action` 을 `v4.38.1` 에서 `v4.38.3` 으로 업그레이드 (by [@dependabot\[bot\]](https://github.com/dependabot) in [#372](https://github.com/dromara/carbon/pull/372) [#381](https://github.com/dromara/carbon/pull/381))

## [v2.6.18](https://github.com/dromara/carbon/compare/v2.6.17...v2.6.18) (2026-09-28)

### ✨ 새로운 기능

- **`카자흐어` 로컬라이즈 언어 지원 추가(`lang/kk.json`)**: 월 이름, 짧은 월 이름, 요일 이름, 짧은 요일 이름은 CLDR 을 따르는 PHP Carbon 의 카자흐어 로케일에서 그대로 가져왔습니다. 상대 시간 표현은 수사 뒤의 명사가 수 변화하지 않는 카자흐어 문법을 따릅니다. 계절과 별자리는 전통적인 카자흐어 이름으로 직접 번역했습니다 (by [@darakanoit](https://github.com/darakanoit) in [#355](https://github.com/dromara/carbon/pull/355))

### 🐛 버그 수정

- **`Format` 메서드에서 마지막 이스케이프 문자를 처리할 때 범위를 벗어난 접근이 발생할 수 있는 문제 수정**: `helper.go` 의 `format2layout` 메서드에서 이스케이프 문자 처리에 경계 검사를 추가하여 범위를 벗어난 접근으로 인한 `panic` 을 방지합니다 (by [@gouguoyin](https://github.com/gouguoyin) in [#348](https://github.com/dromara/carbon/pull/348))
- **주 시작일을 월요일이 아닌 날로 설정하면 `ToWeekString`, `ToShortWeekString`, `Format` 이 잘못된 요일 이름을 반환하는 문제 수정**: `ToWeekString` / `ToShortWeekString` 은 `DayOfWeek() % DaysPerWeek` 로 이름 리소스를 참조했지만 리소스는 절대 순서(일요일 우선)로 정렬되어 있었습니다. 이제 절대 요일 `int(StdTime().Weekday())` 로 참조하여 리소스 순서와 일치시키고 `Format("l")` 동작과 같게 했습니다 (by [@sneha4175](https://github.com/sneha4175) in [#371](https://github.com/dromara/carbon/pull/371))
- **기존 시간 값이 있는 데이터베이스 필드에 `NULL` 을 스캔할 때 기존 값이 지워지지 않는 문제 수정**: 이미 값이 있는 Carbon 에 SQL `NULL` 을 스캔하면 성공을 반환하면서 이전 날짜가 남아 있었습니다. 이제 `Carbon.Scan` 과 레이아웃, 포맷, 타임스탬프 스캐너 래퍼에서 대상을 초기화합니다 (by [@vitalivo](https://github.com/vitalivo) in [#360](https://github.com/dromara/carbon/pull/360))

### ⚡ 성능 최적화

- **카피 온 라이트를 통해 언어 인스턴스와 리소스의 캐싱 및 공유를 최적화하여 메모리 할당 감소**: `SetLocale` 은 캐시된 로케일 맵 전체를 모든 `Language` 인스턴스에 복사하여 Carbon 인스턴스마다 20개 항목 맵을 새로 할당했습니다. 이제 `SetLocale` 은 캐시 맵을 공유하고 공유 상태로 표시하며, `SetResources` 는 첫 쓰기 전에 복사하고, 리소스를 읽기만 하는 처리는 공유 맵을 계속 읽습니다. `NewCarbon()` 당 할당은 1600 B / 13 회에서 232 B / 4 회로 감소합니다 (by [@darakanoit](https://github.com/darakanoit) in [#350](https://github.com/dromara/carbon/pull/350))
- **월, 요일, 계절, 별자리 로컬라이즈 리소스 조회를 최적화하여 호출할 때마다 문자열을 분할하는 작업 방지**: 6개 메서드는 호출할 때마다 "|" 로 구분된 언어 리소스를 읽고 한 요소를 참조하기 위해 `strings.Split` 을 실행했습니다. 이제 작은 헬퍼가 `strings.Count` 와 `strings.IndexByte` 로 제자리에서 요소를 찾아 호출당 할당이 0입니다. `Constellation` -61.54%, `ToShortWeekString` -47.35%, `ToWeekString` -44.13%, `Season` -40.65%, `ToMonthString` -37.66%, `ToShortMonthString` -35.88% (by [@darakanoit](https://github.com/darakanoit) in [#352](https://github.com/dromara/carbon/pull/352))
- **로케일마다 언어 인스턴스 하나를 공유**: `NewCarbon()` 은 인스턴스마다 새로운 `Language` 를 생성했습니다. 이제 `defaultLanguage(locale)` 이 로케일마다 바로 사용 가능한 `Language` 하나를 `sync.Map` 에 캐시합니다. Carbon 생성 시 할당은 4회에서 1회로 감소합니다 (by [@darakanoit](https://github.com/darakanoit) in [#354](https://github.com/dromara/carbon/pull/354))
- **연도별 일수와 누적 일수 테이블을 미리 계산하여 음력 날짜 변환 최적화**: 음력 변환은 연도 테이블을 인덱스로 참조하지 않고 순회하여 O(1900년 이후 연수) 비용이 들었습니다. 이제 지원되는 각 연도의 일수와 누적 합계를 패키지 초기화 시 한 번만 계산합니다. 2020년 날짜를 음력으로 변환할 때 1009 ns -> 282 ns(-72%), 그레고리력으로 되돌릴 때 738 ns -> 23 ns(-97%) 입니다 (by [@darakanoit](https://github.com/darakanoit) in [#357](https://github.com/dromara/carbon/pull/357))

### 🔧 의존성 업그레이드

- `github.com/stretchr/testify` 를 `v1.11.1` 에서 `v1.12.1` 로 업그레이드 (by [@dependabot\[bot\]](https://github.com/dependabot) in [#341](https://github.com/dromara/carbon/pull/341) [#344](https://github.com/dromara/carbon/pull/344))
- `github/codeql-action` 을 `v4.37.4` 에서 `v4.38.1` 로 업그레이드 (by [@dependabot\[bot\]](https://github.com/dependabot) in [#340](https://github.com/dromara/carbon/pull/340) [#342](https://github.com/dromara/carbon/pull/342) [#345](https://github.com/dromara/carbon/pull/345) [#346](https://github.com/dromara/carbon/pull/346) [#362](https://github.com/dromara/carbon/pull/362) [#369](https://github.com/dromara/carbon/pull/369))

### 🚀 CI

- **`CI` 테스트 워크플로를 개선하여 `PR` 에서 전체 테스트 스위트와 경합 상태 검사를 실행하고, `push` 에서만 커버리지를 생성 및 업로드**: 테스트 명령은 루트 패키지만이 아니라 `./...` 로 모든 패키지를 대상으로 합니다. 커버리지 파일 생성과 Codecov 업로드는 `push` 이벤트로 제한하고, 모든 패키지에서 `-race` 를 사용하는 전용 테스트 단계를 추가했습니다 (by [@gouguoyin](https://github.com/gouguoyin) in [#361](https://github.com/dromara/carbon/pull/361))

## [v2.6.17](https://github.com/dromara/carbon/compare/v2.6.16...v2.6.17) (2026-08-09)

### 🐛 버그 수정

- `이슬람` 요일 이름의 철자 오류를 수정하고 `Jumaat` 을 `Jumat` 으로 변경
- `그레고리력` 과 `히브리력` 간 변환에서 발생하는 계산 오류 수정

### 🔧 의존성 업그레이드

- `github/codeql-action` 을 `v4` 에서 `v4.37.4` 로 업그레이드
- `actions/setup-go` 를 `v6` 에서 `v7` 로 업그레이드
- `actions/checkout` 을 `v6` 에서 `v7` 로 업그레이드
- `codecov/codecov-action` 을 `v5` 에서 `v6` 으로 업그레이드
- `codecov/codecov-action` 을 `v6` 에서 `v7` 으로 업그레이드

### 📝 문서

- `README` 문서를 `NRG` 표준 템플릿으로 통일
- `README` 의 불필요한 광고 외부 링크를 정리하고 더 이상 사용되지 않는 `Go Report Card` 상태 배지 제거

## [v2.6.16](https://github.com/dromara/carbon/compare/v2.6.15...v2.6.16) (2026-01-28)

### 🐛 버그 수정

- `Format` 메서드에서 `u`, `v`, `x` 기호 파싱 오류 bug 수정

### 🔧 의존성 업그레이드

- `actions/checkout` 을 `v5` 에서 `v6` 로 업그레이드

## [v2.6.15](https://github.com/dromara/carbon/compare/v2.6.14...v2.6.15) (2025-11-15)

### ✨ 새로운 기능

- `HttpLayout`, `HttpFormat` 상수 및 `ToHttpString` 메서드 추가
- `몽골어` 로컬라이즈 언어 지원 추가(lang/mn.json)
- `아프리칸스어` 로컬라이즈 언어 지원 추가(lang/af.json)

### 🐛 버그 수정

- `difference.go` 파일의 `DiffAbsInDuration` 메서드에서 사용된 `go1.19` 부터만 사용 가능한 `Abs` 메서드 제거
- `네덜란드어` 로컬라이즈 언어 파일(lang/nl.json)의 `weeks` 필드에서 중복된 `Zondag` 오류 수정

## [v2.6.14](https://github.com/dromara/carbon/compare/v2.6.13...v2.6.14) (2025-10-28)

### ✨ 새로운 기능

- `그리스어` 로컬라이즈 언어 지원 추가(lang/el.json)
- `핀란드어` 로컬라이즈 언어 지원 추가(lang/fi.json)
- `미얀마어` 로컬라이즈 언어 지원 추가(lang/my.json)

### 🐛 버그 수정

- `traveler.go` 파일의 시간 증감 시리즈 메서드를 최적화하여 인스턴스를 복사함으로써 불변성을 보장하여 원본 인스턴스의 수정을 방지

## [v2.6.13](https://github.com/dromara/carbon/compare/v2.6.12...v2.6.13) (2025-10-15)

### 🐛 버그 수정

- `helper.go` 의 `format2layout` 메서드에서 이스케이프 문자 처리에 대한 경계 검사를 추가하여 범위를 벗어난 접근으로 인한 `panic` 방지

### ⚡ 성능 최적화

- `language.go` 의 `SetLocale` 메서드에서 `sync.Once` 를 사용하여 언어 파일이 한 번만 로드되도록 보장하고, `sync.Map` 을 사용하여 스레드 안전한 캐싱을 구현

## [v2.6.12](https://github.com/dromara/carbon/compare/v2.6.11...v2.6.12) (2025-09-16)

### 🐛 버그 수정

- `type_carbon.go` 파일의 `UnmarshalJSON` 메서드에서 `isEmpty` 플래그를 설정하여 빈 값을 나타냄
- 잠재적인 경합 상태와 null 포인터 역참조 문제를 수정하여 동시성 안전성 향상

### ⚡ 성능 최적화

- `sync.Map` 을 사용하여 고성능 동시 캐싱을 구현

### 🔧 의존성 업그레이드

- `golang` 환경 의존성을 `1.21` 에서 `1.18` 로 낮춤
- `testify` 테스트 프레임워크를 `v1.10.0` 에서 `v1.11.1` 로 업그레이드

## [v2.6.11](https://github.com/dromara/carbon/compare/v2.6.10...v2.6.11) (2025-07-18)

### ✨ 새로운 기능

- `히브리력(유대력)` 지원 추가 및 단위 테스트와 벤치마크 테스트 파일 추가

### ♻️ 리팩터링 및 변경

- `Sleep` 을 구조체 메서드에서 전역 메서드로 변경
- `페르시아력` 리팩토링 및 벤치마크 테스트 파일 추가
- `중국 음력` 리팩토링 및 벤치마크 테스트 파일 추가
- `율리우스일/수정율리우스일` 리팩토링 및 벤치마크 테스트 파일 추가

### 📝 문서

- 전체 성능 테스트 리포트 파일 추가

## [v2.6.10](https://github.com/dromara/carbon/compare/v2.6.9...v2.6.10) (2025-07-07)

### ✨ 새로운 기능

- `Sleep` 메서드와 관련 `단위 테스트`, `벤치마크 테스트`, `예제 파일` 을 추가
- `MaxYear`, `MinYear`, `MaxMonth`, `MinMonth`, `MaxDay`, `MinDay` 등의 숫자 상수를 추가하고 이러한 상수를 사용하여 하드코딩을 대체

### ⚡ 성능 최적화

- `helper.go` 의 `getAbsValue` 메서드를 최적화하고 조건 판단을 `비트 연산` 으로 대체
- `frozen.go` 파일의 시간 동결 관련 메서드를 최적화하고 `원자 연산` 으로 락 경합을 줄이고 메모리 할당을 최적화
- 벤치마크 테스트 파일을 최적화하고 `순차 테스트`, `병렬 테스트`, `동시 테스트` 를 커버

### ♻️ 리팩터링 및 변경

- `일본어` 번역 파일을 `jp.json` 에서 `ja.json` 으로 변경하고, 설명 문서를 `README.jp.md` 에서 `README.ja.md` 로 변경하여 [ISO639-1](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes) 표준에 준수
- 더 이상 사용되지 않는 `ParseWithLayouts` 메서드를 제거하고 `ParseByLayouts` 메서드로 대체
- 더 이상 사용되지 않는 `ParseWithFormats` 메서드를 제거하고 `ParseByFormats` 메서드로 대체
- 더 이상 사용되지 않는 `CleanTestNow` 메서드를 제거하고 `ClearTestNow` 메서드로 대체
- `ParseByLayout` 과 `ParseByFormat` 메서드의 `타임스탬프` 문자열 파싱 지원을 제거하고, `타임스탬프` 파싱에는 `CreateFromTimestamp`, `CreateFromTimestampMilli`, `CreateFromTimestampMicro`, `CreateFromTimestampNano` 메서드를 사용

### 📝 문서

- 한국어 문서 `README.ko.md` 추가

## [v2.6.9](https://github.com/dromara/carbon/compare/v2.6.8...v2.6.9) (2025-06-28)

### ♻️ 리팩터링 및 변경

- `gorm` 의 `GormDataType` 인터페이스 구현을 제거

## [v2.6.8](https://github.com/dromara/carbon/compare/v2.6.7...v2.6.8) (2025-06-12)

### 🐛 버그 수정

- 파싱 시 레이아웃 템플릿 또는 형식 템플릿이 비어있으면 오류를 반환

### ♻️ 리팩터링 및 변경

- `type_builtin.go` 에서 `DateTimeType` 을 `dateTimeType` 으로, `DateTimeXXXType` 을 `dateTimeXXXType` 으로 이름 변경
- `type_builtin.go` 에서 `DateType` 을 `dateType` 으로, `DateXXXType` 을 `dateXXXType` 으로 이름 변경
- `type_builtin.go` 에서 `TimeType` 을 `timeType` 으로, `TimeXXXType` 을 `timeXXXType` 으로 이름 변경

### 🔧 의존성 업그레이드

- `tests` 에서 `gorm.io/gorm` 을 `1.21.1` 에서 `1.30.0` 으로 업그레이드
- `tests` 에서 `gorm.io/driver/mysql` 을 `1.5.7` 에서 `1.6.0` 으로 업그레이드
- `tests` 에서 `gorm.io/driver/postgres` 를 `1.5.7` 에서 `1.6.0` 으로 업그레이드
- `tests` 에서 `gorm.io/driver/sqlite` 를 `1.5.7` 에서 `1.6.0` 으로 업그레이드

### 📝 문서

- README 파일을 간소화하고 상세한 사용 설명과 예제 사용법을 [공식 사이트](https://carbon.go-pkg.com)로 이전
- [HelloGitHub](https://hellogithub.com/repository/dromara/carbon) 배지 링크 추가

## [v2.6.7](https://github.com/dromara/carbon/compare/v2.6.6...v2.6.7) (2025-05-26)

### ✨ 새로운 기능

- `Closest`/`Farthest` 메서드의 두 번째 매개변수를 선택적 매개변수로 변경
- `ZeroValue`/`EpochValue` 메서드 추가
- `DataTyper` 인터페이스를 추가하고 내장 타입이 `DataTyper` 인터페이스를 구현하도록 활성화

### ♻️ 리팩터링 및 변경

- `String` 메서드에서 null 값 carbon의 검사를 제거
- 파일명을 `type_interface.go` 에서 `interfaces.go` 로 변경

## [v2.6.6](https://github.com/dromara/carbon/compare/v2.6.5...v2.6.6) (2025-05-19)

### ✨ 새로운 기능

- `xorm` 의 `curd` 통합 테스트를 추가하고 현재 `MySQL`/`Postgres`/`SQLite` 를 커버

### 🐛 버그 수정

- `Windows` OS에서 언어 파일을 찾을 수 없는 버그 수정
- 새로운 `Carbon` 인스턴스를 생성할 때 `layout`, `weekStartsAt`, `weekendDays` 및 `lang` 값이 손실되는 오류 수정
- `StartOfWeek` 와 `EndOfWeek` 메서드가 원본 `Carbon` 인스턴스를 예상치 못하게 변경하는 버그 수정

### 🚀 CI

- `ci` 액션에서 `Windows` OS 단위 테스트 추가

## [v2.6.5](https://github.com/dromara/carbon/compare/v2.6.4...v2.6.5) (2025-05-14)

### ✨ 새로운 기능

- `Carbon` 구조체의 `SetLanguage` 메서드에 잘못된 `Language` 구조체 검사를 추가
- `Carbon` 구조체의 `Parse` 메서드에 `Postgres`/`SQLite` 시간 형식 문자열 파싱 지원을 추가
- `Parse`/`ParseByLayout`/`ParseByFormat` 메서드가 `빈` 문자열을 파싱할 때 반환값을 `nil` 에서 빈 carbon으로 변경
- `Carbon` 구조체에 `IsEmpty` 메서드를 추가하여 빈 carbon인지 판단
- `Carbon` 구조체에 `CleanTestNow` 를 대체하는 `ClearTestNow` 메서드를 추가합니다. `CleanTestNow` 는 향후 제거됩니다.
- `Carbon` 구조체에 `ParseWithLayouts` 를 대체하는 `ParseByLayouts` 메서드를 추가합니다. `ParseWithLayouts` 는 향후 제거됩니다.
- `Carbon` 구조체에 `ParseWithFormats` 를 대체하는 `ParseByFormats` 메서드를 추가합니다. `ParseWithFormats` 는 향후 제거됩니다.
- `Language` 구조체의 `SetResources` 메서드에 잘못된 리소스 검사를 추가
- `gorm` 의 `curd` 통합 테스트를 추가하고 현재 `MySQL`/`Postgres`/`SQLite` 를 커버

### ♻️ 리팩터링 및 변경

- `Carbon` 구조체에서 `GormDataType` 메서드를 제거하고 `Value`/`MarshalJSON` 메서드를 `포인터` 수신자에서 `값` 수신자로 변경
- `LayoutType[T]` 구조체에서 `GormDataType` 메서드를 제거하고 `Value`/`MarshalJSON` 메서드를 `포인터` 수신자에서 `값` 수신자로 변경
- `FormatType[T]` 구조체에서 `GormDataType` 메서드를 제거하고 `Value`/`MarshalJSON` 메서드를 `포인터` 수신자에서 `값` 수신자로 변경
- `TimestampType[T]` 구조체에서 `GormDataType` 메서드를 제거하고 `Value`/`MarshalJSON` 메서드를 `포인터` 수신자에서 `값` 수신자로 변경
- 단위 테스트를 위해 `github.com/stretchr/testify/assert` 를 `github.com/stretchr/testify/suite` 로 대체

### 🔧 의존성 업그레이드

- `go` 버전을 `1.18` 에서 `1.21` 로 업그레이드

## [v2.6.4](https://github.com/dromara/carbon/compare/v2.6.3...v2.6.4) (2025-04-28)

### ✨ 새로운 기능

- `Language` 구조체에 `Copy` 메서드 추가
- `carbon.Timestamp` 타입 별칭과 `carbon.NewTimestamp` 메서드 추가
- `carbon.TimestampMilli` 타입 별칭과 `carbon.NewTimestampMilli` 메서드 추가
- `carbon.TimestampMicro` 타입 별칭과 `carbon.NewTimestampMicro` 메서드 추가
- `carbon.TimestampNano` 타입 별칭과 `carbon.NewTimestampNano` 메서드 추가
- `carbon.DateTime` 타입 별칭과 `carbon.NewDateTime` 메서드 추가
- `carbon.DateTimeMilli` 타입 별칭과 `carbon.NewDateTimeMilli` 메서드 추가
- `carbon.DateTimeMicro` 타입 별칭과 `carbon.NewDateTimeMicro` 메서드 추가
- `carbon.DateTimeNano` 타입 별칭과 `carbon.NewDateTimeNano` 메서드 추가
- `carbon.Date` 타입 별칭과 `carbon.NewDate` 메서드 추가
- `carbon.DateMilli` 타입 별칭과 `carbon.NewDateMilli` 메서드 추가
- `carbon.DateMicro` 타입 별칭과 `carbon.NewDateMicro` 메서드 추가
- `carbon.DateNano` 타입 별칭과 `carbon.NewDateNano` 메서드 추가
- `carbon.Time` 타입 별칭과 `carbon.NewTime` 메서드 추가
- `carbon.TimeMilli` 타입 별칭과 `carbon.NewTimeMilli` 메서드 추가
- `carbon.TimeMicro` 타입 별칭과 `carbon.NewTimeMicro` 메서드 추가
- `carbon.TimeNano` 타입 별칭과 `carbon.NewTimeNano` 메서드 추가

### 🐛 버그 수정

- 데이터베이스 필드 타입이 `nil` 일 때 패닉을 수정

### ⚡ 성능 최적화

- 벤치마크 테스트 파일에 `b.ResetTimer()` 추가

### ♻️ 리팩터링 및 변경

- `database_types.go` 를 `type_carbon.go`, `type_layout.go`, `type_format.go`, `type_timestamp.go` 로 분할
- `LayoutFactory` 인터페이스를 `LayoutTyper` 로, `SetLayout` 메서드를 `Layout` 으로 변경
- `FormatFactory` 인터페이스를 `FormatTyper` 로, `SetFormat` 메서드를 `Format` 으로 변경
- `TimestampFactory` 인터페이스를 `TimestampTyper` 로, `SetPrecision` 메서드를 `Precision` 으로 변경

## [v2.6.3](https://github.com/dromara/carbon/compare/v2.6.2...v2.6.3) (2025-04-21)

### ✨ 새로운 기능

- UNIX 에포크 시간(1970-01-01 00:00:00 +0000 UTC)인지 판단하는 `IsEpoch` 메서드 추가
- 주의 끝나는 날을 가져오는 `WeekEndsAt` 메서드 추가
- 주의 주말 날짜를 설정하는 `SetWeekendDays` 메서드 추가
- 기본 주 시작일을 저장하는 `DefaultWeekStartsAt` 전역 변수 추가

### 🐛 버그 수정

- `IsWeekend`, `IsWeekday` 메서드가 다른 국가에서 일관된 결과를 반환하는 버그 수정
- `StdTime` 메서드의 null 포인터로 인한 예외 수정 #294

### ⚡ 성능 최적화

- 벤치마크 테스트 파일 `xxx_bench_test.go` 추가

### ♻️ 리팩터링 및 변경

- 오류 메서드를 `private` 메서드에서 `public` 메서드로 변경
- 주의 기본 시작일을 `일요일` 에서 `월요일` 로 변경
- `MinValue` 메서드의 연도를 `-9998` 에서 `1` 로 변경
- `weeksPerLongYear` 상수를 `WeeksPerLongYear` 로 변경

## [v2.6.2](https://github.com/dromara/carbon/compare/v2.6.1...v2.6.2) (2025-04-08)

### ✨ 새로운 기능

- 시간대 오프셋을 가져오기 위한 형식 기호 `o` 추가
- `TimestampLayout`, `TimestampMilliLayout`, `TimestampMicroLayout`, `TimestampNanoLayout` 상수 추가
- `TimestampFormat`, `TimestampMilliFormat`, `TimestampMicroFormat`, `TimestampNanoFormat` 상수 추가
- `DateTimeMilli`, `DateTimeMicro`, `DateTimeNano` 필드 타입 추가
- `DateMilli`, `DateMicro`, `DateNano` 필드 타입 추가
- `TimeMilli`, `TimeMicro`, `TimeNano` 필드 타입 추가
- `MaxDuration`, `MinDuration` 메서드 추가

### 🐛 버그 수정

- 음력에서 `IsLeapMonth` 판단 오류 버그 수정
- `AtomFormat` 과 `AtomLayout` 형식이 일치하지 않는 값을 반환하는 버그 수정
- `RFC3339Format` 과 `RFC3339Layout` 형식이 일치하지 않는 값을 반환하는 버그 수정
- `IsDST` 메서드의 시간대 버그 수정
- `StartOfXXX`, `EndOfXXX` 메서드에서 시간대가 누락되는 버그 수정
- 다른 달력을 `그레고리력` 으로 변환할 때 시간대가 누락되는 버그 수정

### ♻️ 리팩터링 및 변경

- `CreateFromLunar`, `CreateFromPersian` 메서드에서 hour, minute, second 매개변수를 제거
- 일부 형식 기호 정의를 변경하고, 관련된 기호는 `U`, `V`, `X`, `S`, `T`, `Z`, `u`, `v`, `x`, `z` 를 포함
- 전역 기본 시간대를 설정할 때 `time.Local` 이 더 이상 동기화되지 않음

## [v2.6.1](https://github.com/dromara/carbon/compare/v2.6.0...v2.6.1) (2025-03-27)

### ✨ 새로운 기능

- `ParseWithLayouts` 와 `ParseWithFormats` 메서드 추가

### 🐛 버그 수정

- `gorm` 이 데이터를 업데이트할 때 `updated_at` 필드가 자동으로 업데이트되지 않는 버그 수정

### ♻️ 리팩터링 및 변경

- `formatFactory` 인터페이스를 `FormatFactory` 로, `layoutFactory` 인터페이스를 `LayoutFactory` 로, `timestampFactory` 인터페이스를 `TimestampFactory` 로 이름 변경하고 타입 제약 추가
- `LayoutType`, `FormatType`, `TimestampType` 구조체 메서드의 반환값을 `time` 으로 변경
- `DateTime`, `Date`, `Time` 타입을 `struct` 에서 `string` 으로 변경
- `Timestamp`, `TimestampMilli`, `TimestampMicro`, `TimestampNano` 타입을 `struct` 에서 `int64` 로 변경
- 내장 데이터베이스 필드 타입을 새 파일 `types.go` 로 이동

## [v2.6.0](https://github.com/dromara/carbon/compare/v2.5.4...v2.6.0) (2025-03-25)

### ✨ 새로운 기능

- 시간대 이름을 가져오는 `ZoneName` 메서드 추가
- 오류가 있는지 확인하는 `HasError` 메서드 추가
- `nil` 인지 확인하는 `IsNil` 메서드 추가
- `carbon` 에 대한 깊은 복사를 위한 `Copy` 메서드 추가
- 주 시작일을 가져오는 `WeekStartsAt` 메서드 추가
- 예제 파일 `xxx_example.go` 추가
- 새로운 `constant.go` 파일을 추가하고 `carbon.go` 파일에서 상수를 이 파일로 이전

### ♻️ 리팩터링 및 변경

- `golang` 최소 버전 의존성을 `1.18` 로 업그레이드
- `carbon`, `julian`, `lunar`, `persian` 이 값 전달에서 포인터 전달로 변경
- 기본 전역 시간대를 `Local` 에서 `UTC` 로 변경
- `Offset` 메서드를 `ZoneOffset` 로 이름 변경
- `IsSetTestNow` 메서드를 `IsTestNow` 로 이름 변경
- `UnSetTestNow` 메서드를 `CleanTestNow` 로 이름 변경
- `Location` 메서드를 제거하고 `Timezone` 메서드로 대체
- `IsValid` 와 `IsInvalid` 메서드의 판단 로직을 변경하고, `zero time` 은 더 이상 유효하지 않은 시간으로 간주되지 않음
- 전역 기본 시간대를 설정할 때 `time.Local` 을 동기화하여 업데이트
- `database.go` 를 리팩토링하고 `carbon.DateTime`, `carbon.DateTimeMilli`, `carbon.DateTimeMicro`, `carbon.DateTimeNano`, `carbon.Date`, `carbon.DateMilli`, `carbon.DateMicro`, `carbon.DateNano`, `carbon.Time`, `carbon.TimeMilli`, `carbon.TimeMicro`, `carbon.TimeNano`, `carbon.Timestamp`, `carbon.TimestampMilli`, `carbon.TimestampMicro`, `carbon.TimestampNano` 필드 타입을 제거하고, `MarshalJSON/UnmarshalJSON` 에서 사용자 정의 출력 형식을 구현하기 위해 제네릭 필드를 사용

이전 버전의 업데이트 로그는 <a href="https://github.com/dromara/carbon/releases" target="_blank" rel="noreferrer">releases</a>를 참조하세요
