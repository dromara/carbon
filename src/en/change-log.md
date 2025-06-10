# 更新日志

## [v2.6.7](https://github.com/dromara/carbon/compare/v2.6.6...v2.6.7) (2025-05-26)

- `String` method removes the check for null value carbon
- Rename  `type_interface.go` file to `interfaces.go`
- Change the second parameter of the `Closest`/`Farthest` methods to an optional parameter
- Add `ZeroValue`/`EpochValue` methods
- Add  `DataTyper` interface, and enable builtin types to implement the `DataTyper` interface

## [v2.6.6](https://github.com/dromara/carbon/compare/v2.6.5...v2.6.6) (2025-05-19)

- Fix bug locale file cannot be found on `Windows` os
- Fix bug lose of values of `layout`, `weekStartsAt`, `weekendDays` and `lang` when create new `Carbon` instance #303
- Fix bug `StartOfWeek` and `EndOfWeek` methods change original `Carbon` instance Unexpectedly #304
- Add  `curd` integration testing for `xorm`, currently covering `MySQL`/`Postgres`/`SQLite`
- Add unit testing for the `Windows` os in `ci` action

## [v2.6.5](https://github.com/dromara/carbon/compare/v2.6.4...v2.6.5) (2025-05-14)

- Upgrade `go` version from `1.18` to `1.21`
- `SetLanguage` method of  `Carbon` struct add check for invalid  `Language` struct
- `Parse` method of  `Carbon` struct add support for `Postgres`/`SQLite` time format string
- `Parse`/`ParseByLayout`/`ParseByFormat` methods of  `Carbon` struct change the return value from `nil` to `empty` carbon when parsing an `empty` string
- `Carbon` struct add  `IsEmpty` method to determine if it is an  `empty` carbon
- `Carbon` struct add `ClearTestNow` method to replace `CleanTestNow`, `CleanTestNow` will be removed in the future
- `Carbon` struct add `ParseByLayouts` method to replace `ParseWithLayouts`, `ParseWithLayouts` will be removed in the future
- `Carbon` struct add `ParseByFormats` method to replace `ParseWithFormats`, `ParseWithFormats` will be removed in the future
-  `Carbon` struct remove  `GormDataType` method and change `Value`/`MarshalJSON` methods from `pointer` receiver to `value` receiver
-  `LayoutType[T]` struct remove  `GormDataType` method and change `Value`/`MarshalJSON` methods from `pointer` receiver to `value` receiver
-  `FormatType[T]` struct remove  `GormDataType` method and change `Value`/`MarshalJSON` methods from `pointer` receiver to `value` receiver
-  `TimestampType[T]` struct remove  `GormDataType` method and change `Value`/`MarshalJSON` methods from `pointer` receiver to `value` receiver
- `SetResources` method of  `Language` struct add check invalid resources
- Add  `curd` integration testing for `gorm`, currently covering `MySQL`/`Postgres`/`SQLite`
- Replace  `github.com/stretchr/testify/assert` with `github.com/stretchr/testify/suite` for unit testing

## [v2.6.4](https://github.com/dromara/carbon/compare/v2.6.3...v2.6.4) (2025-04-28)

- Fix panic when database field type was `nil`
- Split  `database_types.go` into` type_carbon.go`, `type_layout.go`, `type_format.go`,  `type_timestamp.go`
- Rename  `LayoutFactory` interface to `LayoutTyper` and `SetLayout` method to `Layout`
- Rename  `FormatFactory` interface to `FormatTyper` and `SeFormat` method to `Format`
- Rename  `TimestampFactory` interface to `TimestampTyper` and `SePrecision` method to `Precision`
- Add `b.ResetTimer()` in  benchmark test files
- Add `Copy` method for `Language` struct
- Add `carbon.Timestamp` type alias and `carbon.NewTimestamp` method
- Add `carbon.TimestampMilli` type alias and `carbon.NewTimestampMilli` method
- Add `carbon.TimestampMicro` type alias and `carbon.NewTimestampMicro` method
- Add `carbon.TimestampNano` type alias and `carbon.NewTimestampNano` method
- Add `carbon.DateTime` type alias and `carbon.NewDateTime` method
- Add `carbon.DateTimeMicro` type alias and `carbon.NewDateTimeMicro` method
- Add `carbon.DateTimeMilli` type alias and `carbon.NewDateTimeMilli` method
- Add `carbon.DateTimeNano` type alias and `carbon.NewDateTimeNano` method
- Add `carbon.Date` type alias and `carbon.NewDate` method
- Add `carbon.DateMilli` type alias and `carbon.NewDateMilli` method
- Add `carbon.DateMicro` type alias and `carbon.NewDateMicro` method
- Add `carbon.DateNano` type alias and `carbon.NewDateNano` method
- Add `carbon.Time` type alias and `carbon.NewTime` method
- Add `carbon.TimeMilli` type alias and `carbon.NewTimeMilli` method
- Add `carbon.TimeMicro` type alias and `carbon.NewTimeMicro` method
- Add `carbon.TimeNano` type alias and `carbon.NewTimeNano` method

## [v2.6.3](https://github.com/dromara/carbon/compare/v2.6.2...v2.6.3) (2025-04-21)

-  Fix bug with the same results in different countries using `IsWeekend` and `IsWeekday` methods.
-  Fix bug `carbon.Parse("").StdTime()` cause a nil pointer panic.  #294
-  Change errors  from `private` method  to `public` method.
-  Change the default start date of the week from `Sunday` to `Monday`
-  Change the year of `MinValue` from `-9998` to `1`
-  Change `weeksPerLongYear` constant to `WeeksPerLongYear`
- Add benchmark test files `xxx_bench_test.go`
-  Add `IsEpoch` method to report whether is a unix epoch time(1970-01-01 00:00:00 +0000 UTC).
-  Add `WeekEndsAt` method to get end day of the week.
-  Add `SetWeekendDays` method to set weekend days of the week.
-  Add `DefaultWeekStartsAt` global variable to store  weekend days of the week.

## [v2.6.2](https://github.com/dromara/carbon/compare/v2.6.1...v2.6.2) (2025-04-08)

-  Remove  `hour`, `Minute`,`second` parameter in `CreateFromLunar`,`CreateFromPersian ` methods
- Change some format symbol definitions involves symbols such as  `U`, `V`, `X`,`S`,`T` `Z`,`u`,`v`,`x`,`z`
- Fix bug of incorrect judgment of `IsLeapMonth` in lunar calendar
- Fix bug of `AtomFormat` and `AtomLayout` return inconsistent value
- Fix bug of `RFC3339Format` and `RFC3339Layout` return inconsistent value
-  `time.Local` is not updated when setting global time zone
-  Add format symbol `o` to get time zone offset
-  Add `TimestampLayout`, `TimestampMilliLayout`, `TimestampMicroLayout` and `TimestampNanoLayout` constants
-  Add `TimestampFormat`, `TimestampMilliFormat`, `TimestampMicroFormat` and `TimestampNanoFormat` constants
- Add `DateTimeMilli`, `DateTimeMicro`, `DateTimeNano` field types
- Add `DateMilli`、`DateMicro`、`DateNano` field types
- Add `TimeMilli`、`TimeMicro`、`TimeNano` field types
-  Fix missing timezone bug in `IsDST` method
- Fix missing timezone bug `StartOfXXX`,  `EndOfXXX` methods
- Fix missing timezone bug when converting other calendars to `Gregorian` calendar
-  `time.Local` is not updated when setting default timezone
-  Add `MaxDuration`, `MinDuration` methods

## [v2.6.1](https://github.com/dromara/carbon/compare/v2.6.0...v2.6.1) (2025-03-27)

- Add `ParseWithLayouts` and `ParseWithFormats` methods
- Rename `formatFactory` interface to `FormatFactory`, `formatFactory` interface to `FormatFactory`, `formatFactory` interface to `FormatFactory`, and add type constraints
- Change the return value of the `GormDataType` method of the `LayoutType`, `FormatType`, `TimestampType`  `struct`  to `time`
- Change `DateTime`, `Date`, `Time` types from `struct` to `string`
- Change `Timestamp`、`TimestampMilli`、`TimestampMicro`,`TimestampNano` types from `struct` to `int64`
- Move built-in database field types to new file `types.go`
- Fixed `updated_at` field is automatically updated and invalid when `gorm` updates data

## [v2.6.0](https://github.com/dromara/carbon/compare/v2.5.4...v2.6.0) (2025-03-25)

- `golang` minimum version dependency upgraded to `1.18`
- `carbon`, `julian`, `lunar`, `persian` changed from value passing to pointer passing
-  Add `ZoneName` method to get the time zone name
- Add `HasError` method to check if there is an error
- Add `IsNil` method to check if it is `nil`
- Add `Copy` method to deep copy `carbon`
- Add `WeekStartsAt` method to get start day of the week
- Add example file `xxx_example.go`
- Add a new `constant.go` file and migrate constants from the `carbon.go` file to this file
- Change default global timezone from `Local` to `UTC`
- Rename `Offset` method to `ZoneOffset` method
- Rename `IsSetTestNow` method to `IsTestNow` method
- Rename `UnSetTestNow` method to `CleanTestNow` method
- Remove `Location` method, replaced by `Timezone` method
- Change the judgment logic of `IsValid` and `IsInvalid` methods, `zero time` is no longer considered invalid time
- Update `time.Local` when setting the global default timezone
- Refactor `database.go` and remove `carbon.DateTime`, `carbon. DateTimeMilli`, `carbon.DateTimeMicro`,  `carbon.DateTimeNano`, `carbon. Date`, `carbon.DateMilli`, `carbon.DateMicro`,  `carbon.DateNano`, `carbon.Time`, `carbon.TimeMilli`,  `carbon.TimeMicro`, `carbon.TimeNano`, `carbon.Timestamp`, `carbon.TimestampMilli`, `carbon.TimestampMicro`, `carbon.TimestampNano` field types, use generic fields instead to implement custom output format when `MarshalJSON/UnmarshalJSON`