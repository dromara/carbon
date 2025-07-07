# 업데이트 로그
## [v2.6.10](https://github.com/dromara/carbon/compare/v2.6.7...v2.6.8) (2025-07-07)

- `한국어` 번역 파일을 `jp.json`에서 `ja.json`으로 변경하고, 설명 문서를 `README.jp.md`에서 `README.ja.md`로 변경하여 [ISO639-1](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes) 표준에 준수
- 더 이상 사용되지 않는 `ParseWithLayouts` 메서드를 제거하고 `ParseByLayouts` 메서드로 대체
- 더 이상 사용되지 않는 `ParseWithFormats` 메서드를 제거하고 `ParseByFormats` 메서드로 대체
- 더 이상 사용되지 않는 `CleanTestNow` 메서드를 제거하고 `ClearTestNow` 메서드로 대체
- `ParseByLayout`과 `ParseByFormat` 메서드의 `타임스탬프` 문자열 파싱 지원을 제거하고, `타임스탬프` 파싱에는 `CreateFromTimestamp`, `CreateFromTimestampMilli`, `CreateFromTimestampMicro`, `CreateFromTimestampNano` 메서드를 사용
- `helper.go`의 `getAbsValue` 메서드를 최적화하고 조건 판단을 `비트 연산`으로 대체
- `frozen.go` 파일의 시간 동결 관련 메서드를 최적화하고 `원자 연산`으로 락 경합을 줄이고 메모리 할당을 최적화
- 벤치마크 테스트 파일을 최적화하고 `순차 테스트`, `병렬 테스트`, `동시 테스트`를 커버
- 한국어 문서 `README.ko.md` 추가
- `Sleep` 메서드와 관련 `단위 테스트`, `벤치마크 테스트`, `예제 파일`을 추가
- `MaxYear`, `MinYear`, `MaxMonth`, `MinMonth`, `MaxDay`, `MinDay` 등의 숫자 상수를 추가하고 이러한 상수를 사용하여 하드코딩을 대체

## [v2.6.9](https://github.com/dromara/carbon/compare/v2.6.7...v2.6.8) (2025-06-28)

- `gorm`의 `GormDataType` 인터페이스 구현을 제거

## [v2.6.8](https://github.com/dromara/carbon/compare/v2.6.7...v2.6.8) (2025-06-12)

- 파싱 시 레이아웃 템플릿 또는 형식 템플릿이 비어있으면 오류를 반환
- `tests`에서 `gorm.io/gorm`을 `1.21.1`에서 `1.30.0`으로 업그레이드
- `tests`에서 `gorm.io/driver/mysql`을 `1.5.7`에서 `1.6.0`으로 업그레이드
- `tests`에서 `gorm.io/driver/postgres`를 `1.5.7`에서 `1.6.0`으로 업그레이드
- `tests`에서 `gorm.io/driver/sqlite`를 `1.5.7`에서 `1.6.0`으로 업그레이드
- `type_builtin.go`에서 `DateTimeType`을 `dateTimeType`으로, `DateTimeXXXType`을 `dateTimeXXXType`으로 이름 변경
- `type_builtin.go`에서 `DateType`을 `dateType`으로, `DateXXXType`을 `dateXXXType`으로 이름 변경
- `type_builtin.go`에서 `TimeType`을 `timeType`으로, `TimeXXXType`을 `timeXXXType`으로 이름 변경
- README 파일을 간소화하고 상세한 사용 설명과 예제 사용법을 [공식 사이트](https://carbon.go-pkg.com)로 이전
- [HelloGitHub](https://hellogithub.com/repository/dromara/carbon) 배지 링크 추가

## [v2.6.7](https://github.com/dromara/carbon/compare/v2.6.6...v2.6.7) (2025-05-26)

- `String` 메서드에서 null 값 carbon의 검사를 제거
- 파일명을 `type_interface.go`에서 `interfaces.go`로 변경
- `Closest`/`Farthest` 메서드의 두 번째 매개변수를 선택적 매개변수로 변경
- `ZeroValue`/`EpochValue` 메서드 추가
- `DataTyper` 인터페이스를 추가하고 내장 타입이 `DataTyper` 인터페이스를 구현하도록 활성화

## [v2.6.6](https://github.com/dromara/carbon/compare/v2.6.5...v2.6.6) (2025-05-19)

- `Windows` OS에서 언어 파일을 찾을 수 없는 버그 수정
- 새로운 `Carbon` 인스턴스를 생성할 때 `layout`, `weekStartsAt`, `weekendDays` 및 `lang` 값이 손실되는 오류 수정
- `StartOfWeek`와 `EndOfWeek` 메서드가 원본 `Carbon` 인스턴스를 예상치 못하게 변경하는 버그 수정
- `xorm`의 `curd` 통합 테스트를 추가하고 현재 `MySQL`/`Postgres`/`SQLite`를 커버
- `ci` 액션에서 `Windows` OS 단위 테스트 추가

## [v2.6.5](https://github.com/dromara/carbon/compare/v2.6.4...v2.6.5) (2025-05-14)

- `go` 버전을 `1.18`에서 `1.21`로 업그레이드
- `Carbon` 구조체의 `SetLanguage` 메서드에 잘못된 `Language` 구조체 검사를 추가
- `Carbon` 구조체의 `Parse` 메서드에 `Postgres`/`SQLite` 시간 형식 문자열 파싱 지원을 추가
- `Parse`/`ParseByLayout`/`ParseByFormat` 메서드가 `빈` 문자열을 파싱할 때 반환값을 `nil`에서 빈 carbon으로 변경
- `Carbon` 구조체에 `IsEmpty` 메서드를 추가하여 빈 carbon인지 판단
- `Carbon` 구조체에 `CleanTestNow`를 대체하는 `ClearTestNow` 메서드를 추가합니다. `CleanTestNow`는 향후 제거됩니다.
- `Carbon` 구조체에 `ParseWithLayouts`를 대체하는 `ParseByLayouts` 메서드를 추가합니다. `ParseWithLayouts`는 향후 제거됩니다.
- `Carbon` 구조체에 `ParseWithFormats`를 대체하는 `ParseByFormats` 메서드를 추가합니다. `ParseWithFormats`는 향후 제거됩니다.
- `Carbon` 구조체에서 `GormDataType` 메서드를 제거하고 `Value`/`MarshalJSON` 메서드를 `포인터` 수신자에서 `값` 수신자로 변경
- `LayoutType[T]` 구조체에서 `GormDataType` 메서드를 제거하고 `Value`/`MarshalJSON` 메서드를 `포인터` 수신자에서 `값` 수신자로 변경
- `FormatType[T]` 구조체에서 `GormDataType` 메서드를 제거하고 `Value`/`MarshalJSON` 메서드를 `포인터` 수신자에서 `값` 수신자로 변경
- `TimestampType[T]` 구조체에서 `GormDataType` 메서드를 제거하고 `Value`/`MarshalJSON` 메서드를 `포인터` 수신자에서 `값` 수신자로 변경
- `Language` 구조체의 `SetResources` 메서드에 잘못된 리소스 검사를 추가
- `gorm`의 `curd` 통합 테스트를 추가하고 현재 `MySQL`/`Postgres`/`SQLite`를 커버
- 단위 테스트를 위해 `github.com/stretchr/testify/assert`를 `github.com/stretchr/testify/suite`로 대체

## [v2.6.4](https://github.com/dromara/carbon/compare/v2.6.3...v2.6.4) (2025-04-28)

- 데이터베이스 필드 타입이 `nil`일 때 패닉을 수정
- `database_types.go`를 `type_carbon.go`, `type_layout.go`, `type_format.go`, `type_timestamp.go`로 분할
- `LayoutFactory` 인터페이스를 `LayoutTyper`로, `SetLayout` 메서드를 `Layout`으로 변경
- `FormatFactory` 인터페이스를 `FormatTyper`로, `SetFormat` 메서드를 `Format`으로 변경
- `TimestampFactory` 인터페이스를 `TimestampTyper`로, `SetPrecision` 메서드를 `Precision`으로 변경
- 벤치마크 테스트 파일에 `b.ResetTimer()` 추가
- `Language` 구조체에 `Copy` 메서드 추가
- `carbon.Timestamp` 타입 별칭과 `carbon.NewTimestamp` 메서드 추가
- `carbon.TimestampMilli` 타입 별칭과 `carbon.NewTimestampMilli` 메서드 추가
- `carbon.TimestampMicro` 타입 별칭과 `carbon.NewTimestampMicro` 메서드 추가
- `carbon.TimestampNano` 타입 별칭과 `carbon.NewTimestampNano` 메서드 추가
- `carbon.DateTime` 타입 별칭과 `carbon.NewDateTime` 메서드 추가
- `carbon.DateTimeMicro` 타입 별칭과 `carbon.NewDateTimeMicro` 메서드 추가
- `carbon.DateTimeMilli` 타입 별칭과 `carbon.NewDateTimeMilli` 메서드 추가
- `carbon.DateTimeNano` 타입 별칭과 `carbon.NewDateTimeNano` 메서드 추가
- `carbon.Date` 타입 별칭과 `carbon.NewDate` 메서드 추가
- `carbon.DateMilli` 타입 별칭과 `carbon.NewDateMilli` 메서드 추가
- `carbon.DateMicro` 타입 별칭과 `carbon.NewDateMicro` 메서드 추가
- `carbon.DateNano` 타입 별칭과 `carbon.NewDateNano` 메서드 추가
- `carbon.Time` 타입 별칭과 `carbon.NewTime` 메서드 추가
- `carbon.TimeMilli` 타입 별칭과 `carbon.NewTimeMilli` 메서드 추가
- `carbon.TimeMicro` 타입 별칭과 `carbon.NewTimeMicro` 메서드 추가
- `carbon.TimeNano` 타입 별칭과 `carbon.NewTimeNano` 메서드 추가

## [v2.6.3](https://github.com/dromara/carbon/compare/v2.6.2...v2.6.3) (2025-04-21)

- `IsWeekend`, `IsWeekday` 메서드가 다른 국가에서 일관된 결과를 반환하는 버그 수정
- `StdTime` 메서드의 null 포인터로 인한 예외 수정 #294
- 오류 메서드를 `private` 메서드에서 `public` 메서드로 변경
- 주의 기본 시작일을 `일요일`에서 `월요일`로 변경
- `MinValue` 메서드의 연도를 `-9998`에서 `1`로 변경
- `weeksPerLongYear` 상수를 `WeeksPerLongYear`로 변경
- 벤치마크 테스트 파일 `xxx_bench_test.go` 추가
- UNIX 에포크 시간(1970-01-01 00:00:00 +0000 UTC)인지 판단하는 `IsEpoch` 메서드 추가
- 주의 끝나는 날을 가져오는 `WeekEndsAt` 메서드 추가
- 주의 주말 날짜를 설정하는 `SetWeekendDays` 메서드 추가
- 기본 주 시작일을 저장하는 `DefaultWeekStartsAt` 전역 변수 추가

## [v2.6.2](https://github.com/dromara/carbon/compare/v2.6.1...v2.6.2) (2025-04-08)

- `CreateFromLunar`, `CreateFromPersian` 메서드에서 hour, minute, second 매개변수를 제거
- 일부 형식 기호 정의를 변경하고, 관련된 기호는 `U`, `V`, `X`, `S`, `T`, `Z`, `u`, `v`, `x`, `z`를 포함
- 음력에서 `IsLeapMonth` 판단 오류 버그 수정
- `AtomFormat`과 `AtomLayout` 형식이 일치하지 않는 값을 반환하는 버그 수정
- `RFC3339Format`과 `RFC3339Layout` 형식이 일치하지 않는 값을 반환하는 버그 수정
- 전역 기본 시간대를 설정할 때 `time.Local`이 더 이상 동기화되지 않음
- 시간대 오프셋을 가져오기 위한 형식 기호 `o` 추가
- `TimestampLayout`, `TimestampMilliLayout`, `TimestampMicroLayout`, `TimestampNanoLayout` 상수 추가
- `TimestampFormat`, `TimestampMilliFormat`, `TimestampMicroFormat`, `TimestampNanoFormat` 상수 추가
- `DateTimeMilli`, `DateTimeMicro`, `DateTimeNano` 필드 타입 추가
- `DateMilli`, `DateMicro`, `DateNano` 필드 타입 추가
- `TimeMilli`, `TimeMicro`, `TimeNano` 필드 타입 추가
- `IsDST` 메서드의 시간대 버그 수정
- `StartOfXXX`, `EndOfXXX` 메서드에서 시간대가 누락되는 버그 수정
- 다른 달력을 `그레고리력`으로 변환할 때 시간대가 누락되는 버그 수정
- 기본 시간대를 설정할 때 `time.Local`이 더 이상 동기화되지 않음
- `MaxDuration`, `MinDuration` 메서드 추가

## [v2.6.1](https://github.com/dromara/carbon/compare/v2.6.0...v2.6.1) (2025-03-27)

- `ParseWithLayouts`와 `ParseWithFormats` 메서드 추가
- `formatFactory` 인터페이스를 `FormatFactory`로 이름 변경하고 타입 제약 추가
- `LayoutType`, `FormatType`, `TimestampType` 구조체 메서드의 반환값을 `time`으로 변경
- `DateTime`, `Date`, `Time` 타입을 `struct`에서 `string`으로 변경
- `Timestamp`, `TimestampMilli`, `TimestampMicro`, `TimestampNano` 타입을 `struct`에서 `int64`로 변경
- 내장 데이터베이스 필드 타입을 새 파일 `types.go`로 이동
- `gorm`이 데이터를 업데이트할 때 `updated_at` 필드가 자동으로 업데이트되지 않는 버그 수정

## [v2.6.0](https://github.com/dromara/carbon/compare/v2.5.4...v2.6.0) (2025-03-25)

- `golang` 최소 버전 의존성을 `1.18`로 업그레이드
- `carbon`, `julian`, `lunar`, `persian`이 값 전달에서 포인터 전달로 변경
- 시간대 이름을 가져오는 `ZoneName` 메서드 추가
- 오류가 있는지 확인하는 `HasError` 메서드 추가
- `nil`인지 확인하는 `IsNil` 메서드 추가
- `carbon`에 대한 깊은 복사를 위한 `Copy` 메서드 추가
- 주 시작일을 가져오는 `WeekStartsAt` 메서드 추가
- 예제 파일 `xxx_example.go` 추가
- 새로운 `constant.go` 파일을 추가하고 `carbon.go` 파일에서 상수를 이 파일로 이전
- 기본 전역 시간대를 `Local`에서 `UTC`로 변경
- `Offset` 메서드를 `ZoneOffset`로 이름 변경
- `IsSetTestNow` 메서드를 `IsTestNow`로 이름 변경
- `UnSetTestNow` 메서드를 `CleanTestNow`로 이름 변경
- `Location` 메서드를 제거하고 `Timezone` 메서드로 대체
- `IsValid`와 `IsInvalid` 메서드의 판단 로직을 변경하고, `zero time`은 더 이상 유효하지 않은 시간으로 간주되지 않음
- 전역 기본 시간대를 설정할 때 `time.Local`을 동기화하여 업데이트
- `database.go`를 리팩토링하고 `carbon.DateTime`, `carbon.DateTimeMilli`, `carbon.DateTimeMicro`, `carbon.DateTimeNano`, `carbon.Date`, `carbon.DateMilli`, `carbon.DateMicro`, `carbon.DateNano`, `carbon.Time`, `carbon.TimeMilli`, `carbon.TimeMicro`, `carbon.TimeNano`, `carbon.Timestamp`, `carbon.TimestampMilli`, `carbon.TimestampMicro`, `carbon.TimestampNano` 필드 타입을 제거하고, `MarshalJSON/UnmarshalJSON`에서 사용자 정의 출력 형식을 구현하기 위해 제네릭 필드를 사용

이전 버전의 업데이트 로그는 <a href="https://github.com/dromara/carbon/releases" target="_blank" rel="noreferrer">releases</a>를 참조하세요 