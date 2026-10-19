---
head:
  - - meta
    - name: description
      content: "Carbon Change Log, detailed records of all updates from historical versions, including new features, performance optimizations, bug fixes, API changes, dependency upgrades and compatibility notes"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, change log, version history, changelog, feature updates, bug fixes, performance optimization, API changes, release notes, version updates"
---

# Change Log

## [v2.7.0](https://github.com/dromara/carbon/compare/v2.6.18...v2.7.0) (2026-10-19)

### ✨ Features

- **Add `AddWeekdays`, `AddWeekday`, `SubWeekdays`, `SubWeekday` methods to skip weekend days when counting working days**: moves over days while skipping the weekend days configured on the instance, so deadlines can be counted in working days; the four names mirror `AddDays` / `AddDay` / `SubDays` / `SubDay`, and weekend days are read from the instance so `SetWeekendDays` keeps working; when all seven days are marked as weekend days, each weekday span is capped to a whole week to prevent infinite loops (by [@darakanoit](https://github.com/darakanoit) in [#366](https://github.com/dromara/carbon/pull/366))
- **Add `IsSameWeek` method to check whether two dates fall in the same week**: completes the `IsSame*` family (`IsSameCentury`, `IsSameDecade`, `IsSameYear`, `IsSameQuarter`, `IsSameWeek`, `IsSameMonth`, `IsSameDay`, `IsSameHour`, `IsSameMinute`, `IsSameSecond`); honours `SetWeekStartsAt` so two dates that share a week under one setting may not share it under another; the comparison is based on the date the week starts on, correctly handling year-boundary-crossing weeks (by [@darakanoit](https://github.com/darakanoit) in [#368](https://github.com/dromara/carbon/pull/368))
- **Add `IsBirthday` method to compare month and day for birthdays**: reports whether two dates fall on the same month and day, ignoring the year; without an argument it compares with now; February 29 only matches February 29, so it is never a birthday in a common year, matching PHP Carbon behaviour (by [@darakanoit](https://github.com/darakanoit) in [#364](https://github.com/dromara/carbon/pull/364))
- **Add `Asia/Jakarta` timezone constant support** (by [@mirzaq19](https://github.com/mirzaq19) in [#374](https://github.com/dromara/carbon/pull/374))

### 🐛 Fixes

- **Fix `StartOfSecond` and `EndOfSecond` selecting the wrong occurrence during daylight-saving rollbacks**: during a repeated daylight-saving hour, `time.Date` reconstruction could select the other occurrence of the same local time; now copies the instance and adjusts the existing timestamp's subsecond precision instead, preserving the selected instant, location and configuration (by [@mikamikasuki](https://github.com/mikamikasuki) in [#377](https://github.com/dromara/carbon/pull/377))
- **Fix `StartOfMinute` and `EndOfMinute` selecting the wrong occurrence during daylight-saving rollbacks**: same bug as #377 one level up; now copies the instance and subtracts the local seconds, keeping the instant, location and offset (by [@dualfroz](https://github.com/dualfroz) in [#378](https://github.com/dromara/carbon/pull/378))
- **Fix `StartOfHour` and `EndOfHour` selecting the wrong occurrence during daylight-saving rollbacks**: preserves the receiver's occurrence when the reconstructed hour boundary has a valid counterpart with the receiver's offset; keeps the existing calendar reconstruction when no such counterpart exists, including Lord Howe's partial-hour transitions (by [@dajiaohuang](https://github.com/dajiaohuang) in [#382](https://github.com/dromara/carbon/pull/382))
- **Fix `DiffInSeconds` ignoring the sub-second part and counting inaccurately across second boundaries**: subtracting whole-second timestamps ignored the fractional part, so a gap of 0.2s across a second boundary counted as a full second; now adjusts by one when the nanosecond difference has the opposite sign, counting whole seconds elapsed and truncated toward zero (by [@youdie006](https://github.com/youdie006) in [#379](https://github.com/dromara/carbon/pull/379))
- **Fix `DiffInDays` and `DiffInWeeks` computing incorrect results across DST transitions**: dividing elapsed seconds gave wrong results when crossing a DST change; now counts calendar days the way `DiffInMonths` counts months (`AddDate` plus an adjustment), and `DiffInWeeks` uses it; only intervals that cross a DST transition change (by [@youdie006](https://github.com/youdie006) in [#376](https://github.com/dromara/carbon/pull/376))
- **Fix `DiffInYears` and `DiffInMonths` producing incorrect estimates when comparing dates across different time zones**: the first estimate from each date's own `Year()` / `Month()` fields could be off when the two dates are in different locations; now reads the later date in the earlier date's location before taking the year and month fields, matching the frame `DiffInDays` uses (by [@ybai08](https://github.com/ybai08) in [#384](https://github.com/dromara/carbon/pull/384))
- **Fix `IsSameDecade` not comparing the century**: `Decade()` is the decade within the century, so `Parse("1995-08-05").IsSameDecade(Parse("2095-08-05"))` returned true; now also compares `Century()`, matching the buckets `StartOfDecade` uses (by [@youdie006](https://github.com/youdie006) in [#375](https://github.com/dromara/carbon/pull/375))
- **Fix `FromStdTime` returning the wrong lunar date just after midnight when the zone's UTC offset differs from 1900**: elapsed time from 1900-01-31 in the time's own zone leaked the two instants' offsets into the day count; now counts the days between the wall-clock dates in UTC, matching `persian.FromStdTime` and `hebrew.FromStdTime` (by [@youdie006](https://github.com/youdie006) in [#380](https://github.com/dromara/carbon/pull/380))
- **Fix race condition in frozen-clock mode where concurrent `Sleep` calls lost clock advances**: concurrent `Sleep` calls on a frozen clock raced with each other and with `Now`, losing clock advances; now uses the existing frozen-clock mutex around the check and update (by [@yuefdev](https://github.com/yuefdev) in [#383](https://github.com/dromara/carbon/pull/383))

### 🔧 Dependencies

- Upgrade `github/codeql-action` from `v4.38.1` to `v4.38.3` (by [@dependabot\[bot\]](https://github.com/dependabot) in [#372](https://github.com/dromara/carbon/pull/372) [#381](https://github.com/dromara/carbon/pull/381))

## [v2.6.18](https://github.com/dromara/carbon/compare/v2.6.17...v2.6.18) (2026-09-28)

### ✨ Features

- **Add `Kazakh` localised language support (`lang/kk.json`)**: month names, short month names, weekday names and short weekday names are taken verbatim from the Kazakh locale of PHP Carbon, which follows CLDR; relative time phrases follow Kazakh grammar where nouns do not inflect for number after a numeral; seasons and constellations are translated directly using traditional Kazakh names (by [@darakanoit](https://github.com/darakanoit) in [#355](https://github.com/dromara/carbon/pull/355))

### 🐛 Fixes

- **Fix possible out-of-bounds access when the `Format` method handles a trailing escape character**: added boundary checks for escape character handling in `format2layout` method of `helper.go` to prevent `panic` caused by out-of-bounds access (by [@gouguoyin](https://github.com/gouguoyin) in [#348](https://github.com/dromara/carbon/pull/348))
- **Fix incorrect weekday names returned by `ToWeekString`, `ToShortWeekString`, and `Format` when the week starts on a day other than Monday**: `ToWeekString` / `ToShortWeekString` indexed the name resource by `DayOfWeek() % DaysPerWeek`, but the resources are ordered absolutely (Sunday first); now indexes by the absolute weekday `int(StdTime().Weekday())`, matching the resource order and mirroring what `Format("l")` already does (by [@sneha4175](https://github.com/sneha4175) in [#371](https://github.com/dromara/carbon/pull/371))
- **Fix failure to clear the existing value when scanning `NULL` from a database into a field that already contains a time value**: scanning a SQL `NULL` into an already populated Carbon value returned success but left the previous date in place; now resets the destination in `Carbon.Scan` and the layout, format, and timestamp scanner wrappers (by [@vitalivo](https://github.com/vitalivo) in [#360](https://github.com/dromara/carbon/pull/360))

### ⚡ Performance

- **Optimize language instance and resource caching and sharing through copy-on-write to reduce memory allocation**: `SetLocale` copies the whole cached locale map into every `Language` instance, so every Carbon instance allocates a fresh 20-entry map; now `SetLocale` shares the cached map and marks it as shared, `SetResources` copies before its first write, and everything that only reads resources keeps reading the shared map; per `NewCarbon()`, allocation drops from 1600 B / 13 allocs to 232 B / 4 allocs (by [@darakanoit](https://github.com/darakanoit) in [#350](https://github.com/dromara/carbon/pull/350))
- **Optimize localized resource lookups for months, weekdays, seasons, and constellations to avoid splitting strings on every call**: six methods read a "|" separated language resource and call `strings.Split` on every call just to index one element; now a small helper locates the item in place using `strings.Count` and `strings.IndexByte`, zero allocations per call; `Constellation` -61.54%, `ToShortWeekString` -47.35%, `ToWeekString` -44.13%, `Season` -40.65%, `ToMonthString` -37.66%, `ToShortMonthString` -35.88% (by [@darakanoit](https://github.com/darakanoit) in [#352](https://github.com/dromara/carbon/pull/352))
- **Share one language instance per locale**: `NewCarbon()` built a fresh `Language` for every instance; now `defaultLanguage(locale)` caches one ready-to-use `Language` per locale in a `sync.Map`; constructing a Carbon goes from 4 allocations to 1 (by [@darakanoit](https://github.com/darakanoit) in [#354](https://github.com/dromara/carbon/pull/354))
- **Optimize lunar date conversion by precomputing tables for days per year and cumulative days**: both lunar conversions walked the year table instead of indexing it, costing O(years since 1900); now derives the number of days in each supported year and the running total once at package initialisation; 2020 date to lunar 1009 ns -> 282 ns (-72%), back to gregorian 738 ns -> 23 ns (-97%) (by [@darakanoit](https://github.com/darakanoit) in [#357](https://github.com/dromara/carbon/pull/357))

### 🔧 Dependencies

- Upgrade `github.com/stretchr/testify` from `v1.11.1` to `v1.12.1` (by [@dependabot\[bot\]](https://github.com/dependabot) in [#341](https://github.com/dromara/carbon/pull/341) [#344](https://github.com/dromara/carbon/pull/344))
- Upgrade `github/codeql-action` from `v4.37.4` to `v4.38.1` (by [@dependabot\[bot\]](https://github.com/dependabot) in [#340](https://github.com/dromara/carbon/pull/340) [#342](https://github.com/dromara/carbon/pull/342) [#345](https://github.com/dromara/carbon/pull/345) [#346](https://github.com/dromara/carbon/pull/346) [#362](https://github.com/dromara/carbon/pull/362) [#369](https://github.com/dromara/carbon/pull/369))

### 🚀 CI

- **Improve the `CI` test workflow to run the full test suite and race detection on PRs, and generate and upload coverage only on push**: test commands now target all packages via `./...` instead of only the root package; coverage file generation and Codecov upload are gated to `push` events; added a dedicated test step using `-race` across all packages (by [@gouguoyin](https://github.com/gouguoyin) in [#361](https://github.com/dromara/carbon/pull/361))

## [v2.6.17](https://github.com/dromara/carbon/compare/v2.6.16...v2.6.17) (2026-08-09)

### 🐛 Fixes

- Fix spelling error in the `Islamic` weekday name, changing `Jumaat` to `Jumat` #329
- Fix calculation error in conversion between the `Gregorian calendar` and the `Hebrew calendar` #338

### 🔧 Dependencies

- Upgrade `github/codeql-action` from `v4` to `v4.37.4` #339
- Upgrade `actions/setup-go` from `v6` to `v7` #337
- Upgrade `actions/checkout` from `v6` to `v7` #336
- Upgrade `codecov/codecov-action` from `v5` to `v6` #330
- Upgrade `codecov/codecov-action` from `v6` to `v7` #335

### 📝 Docs

- Unify `README` documents using the `NRG` standard template #331
- Clean up redundant advertising external links in `README` and remove the obsolete `Go Report Card` status badge

## [v2.6.16](https://github.com/dromara/carbon/compare/v2.6.15...v2.6.16) (2026-01-28)

### 🐛 Fixes

- Fix the bug in the parsing of `u`, `v`, and `x` symbols in the `Format` method #328

### 🔧 Dependencies

- Upgrade `actions/checkout` from `v5` to `v6`

## [v2.6.15](https://github.com/dromara/carbon/compare/v2.6.14...v2.6.15) (2025-11-15)

### ✨ Features

- Add `HttpLayout`, `HttpFormat` constants and `ToHttpString` method
- Add `Mongolian` localised language support(lang/mn.json)
- Add `Afrikaans` localised language support(lang/af.json)

### 🐛 Fixes

- Remove the `Abs` method used in the `DiffAbsInDuration` method of the `difference.go` file, which is only available from `go1.19`
- Fix duplicate `Zondag` error in the `weeks` field of the `Dutch` localised language file (lang/nl.json)

## [v2.6.14](https://github.com/dromara/carbon/compare/v2.6.13...v2.6.14) (2025-10-28)

### ✨ Features

- Add `Greek` localised language support(lang/el.json)
- Add `Finnish` localised language support(lang/fi.json)
- Add `Myanmar` localised language support(lang/my.json)

### 🐛 Fixes

- Optimize time addition and subtraction series methods in `traveler.go` file to ensure immutability by copying instances to avoid modifying the original instance

## [v2.6.13](https://github.com/dromara/carbon/compare/v2.6.12...v2.6.13) (2025-10-15)

### 🐛 Fixes

- Add boundary checks for escape character handling in `format2layout` method of `helper.go` to prevent `panic` caused by out-of-bounds access

### ⚡ Performance

- Use `sync.Once` in `SetLocale` method of `language.go` to ensure language files are loaded only once, and use `sync.Map` for thread-safe caching

## [v2.6.12](https://github.com/dromara/carbon/compare/v2.6.11...v2.6.12) (2025-09-16)

### 🐛 Fixes

- Set `isEmpty` flag in `UnmarshalJSON` method of `type_carbon.go` file to indicate empty values
- Fix potential race conditions and null pointer dereference issues, improving concurrency safety

### ⚡ Performance

- Use `sync.Map` to implement high-performance concurrent caching

### 🔧 Dependencies

- Lower `golang` environment dependency from `1.21` to `1.18`
- Upgrade `testify` testing framework from `v1.10.0` to `v1.11.1`

## [v2.6.11](https://github.com/dromara/carbon/compare/v2.6.10...v2.6.11) (2025-07-18)

### ✨ Features

- Add `Hebrew` calendar support and add unit tests and benchmark test files

### ♻️ Refactor & Changes

- Change `Sleep` from struct method to global method
- Refactor `Persian` calendar and add benchmark test files
- Refactor `Chinese Lunar` calendar and add benchmark test files
- Refactor `Julian/Modified Julian` calendar and add benchmark test files

### 📝 Docs

- Add overall performance test report file

## [v2.6.10](https://github.com/dromara/carbon/compare/v2.6.9...v2.6.10) (2025-07-07)

### ✨ Features

- Add `Sleep` method and related `unit tests`, `benchmark tests`, and `example file`
- Add number constants such as `MaxYear`, `MinYear`, `MaxMonth`, `MinMonth`, `MaxDay`, `MinDay` etc. and replace hard coded with these constants

### ⚡ Performance

- Optimize `getAbsValue` method in `helper.go` to replace conditional judgments with `bitwise operation`
- Optimize the methods related to time freezing in `frozen.go`, reduce lock contention using `atomic operation` and optimize memory allocation
- Optimize benchmark test files to cover `serial` testing, `parallel` testing and `concurrent` testing

### ♻️ Refactor & Changes

- Change `Japanese` translation file from `jp.json` to `ja.json` and rename document from `README.jp.md` to `README.ja.md` to comply with the [ISO639-1](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes) standard
- Remove deprecated `ParseWithLayouts` method and replace with `ParseByLayouts` method
- Remove deprecated `ParseWithFormats` method and replace with `ParseByFormats` method
- Remove deprecated `CleanTestNow` method and replace with `ClearTestNow` method
- Remove parsing support for timestamp strings from `ParseByLayout` and `ParseByFormat` methods, use` CreateFromTimestamp`, `CreateFromTimestampMilli`, `CreateFromTimestampMicro`, `CreateFromTimestampNano` to parse timestamp strings

### 📝 Docs

- Add Korean readme document `README.ko.md`

## [v2.6.9](https://github.com/dromara/carbon/compare/v2.6.8...v2.6.9) (2025-06-28)

### ♻️ Refactor & Changes

- Remove the implementation of the `GormDataType` interface for `gorm`

## [v2.6.8](https://github.com/dromara/carbon/compare/v2.6.7...v2.6.8) (2025-06-12)

### 🐛 Fixes

- Parsing when layouts or formats is empty, returns a error

### ♻️ Refactor & Changes

- Rename  `DateTimeType` to `dateTimeType`, `DateTimeXXXType` to `dateTimeXXXType`  in `type_builtin.go`
- Rename  `DateType` to `dateType`, `DateXXXType` to `dateXXXType` in `type_builtin.go`
- Rename  `TimeType` to `timeType`, `TimeXXXType` to `timeXXXType`  in `type_builtin.go`

### 🔧 Dependencies

- Upgrade `gorm.io/gorm` from `1.21.1` to `1.30.0` in `tests`
- Upgrade `gorm.io/driver/mysql` from `1.5.7` to `1.6.0` in `tests`
- Upgrade `gorm.io/driver/postgres` from `1.5.7` to `1.6.0` in `tests`
- Upgrade `gorm.io/driver/sqlite` from `1.5.7` to `1.6.0` in `tests`

### 📝 Docs

- Simplify readme file and migrate overview and example usage to the [official doc website](https://carbon.go-pkg.com/)
- Add [HelloGitHub](https://hellogithub.com/repository/dromara/carbon) badge link

## [v2.6.7](https://github.com/dromara/carbon/compare/v2.6.6...v2.6.7) (2025-05-26)

### ✨ Features

- Change the second parameter of the `Closest`/`Farthest` methods to an optional parameter
- Add `ZeroValue`/`EpochValue` methods
- Add  `DataTyper` interface, and enable builtin types to implement the `DataTyper` interface

### ♻️ Refactor & Changes

- `String` method removes the check for null value carbon
- Rename  `type_interface.go` file to `interfaces.go`

## [v2.6.6](https://github.com/dromara/carbon/compare/v2.6.5...v2.6.6) (2025-05-19)

### ✨ Features

- Add  `curd` integration testing for `xorm`, currently covering `MySQL`/`Postgres`/`SQLite`

### 🐛 Fixes

- Fix bug locale file cannot be found on `Windows` os
- Fix bug lose of values of `layout`, `weekStartsAt`, `weekendDays` and `lang` when create new `Carbon` instance
- Fix bug `StartOfWeek` and `EndOfWeek` methods change original `Carbon` instance Unexpectedly

### 🚀 CI

- Add unit testing for the `Windows` os in `ci` action

## [v2.6.5](https://github.com/dromara/carbon/compare/v2.6.4...v2.6.5) (2025-05-14)

### ✨ Features

- `SetLanguage` method of  `Carbon` struct add check for invalid  `Language` struct
- `Parse` method of  `Carbon` struct add support for `Postgres`/`SQLite` time format string
- `Parse`/`ParseByLayout`/`ParseByFormat` methods of  `Carbon` struct change the return value from `nil` to `empty` carbon when parsing an `empty` string
- `Carbon` struct add  `IsEmpty` method to determine if it is an  `empty` carbon
- `Carbon` struct add `ClearTestNow` method to replace `CleanTestNow`, `CleanTestNow` will be removed in the future
- `Carbon` struct add `ParseByLayouts` method to replace `ParseWithLayouts`, `ParseWithLayouts` will be removed in the future
- `Carbon` struct add `ParseByFormats` method to replace `ParseWithFormats`, `ParseWithFormats` will be removed in the future
- `SetResources` method of  `Language` struct add check invalid resources
- Add  `curd` integration testing for `gorm`, currently covering `MySQL`/`Postgres`/`SQLite`

### ♻️ Refactor & Changes

-  `Carbon` struct remove  `GormDataType` method and change `Value`/`MarshalJSON` methods from `pointer` receiver to `value` receiver
-  `LayoutType[T]` struct remove  `GormDataType` method and change `Value`/`MarshalJSON` methods from `pointer` receiver to `value` receiver
-  `FormatType[T]` struct remove  `GormDataType` method and change `Value`/`MarshalJSON` methods from `pointer` receiver to `value` receiver
-  `TimestampType[T]` struct remove  `GormDataType` method and change `Value`/`MarshalJSON` methods from `pointer` receiver to `value` receiver
- Replace  `github.com/stretchr/testify/assert` with `github.com/stretchr/testify/suite` for unit testing

### 🔧 Dependencies

- Upgrade `go` version from `1.18` to `1.21`

## [v2.6.4](https://github.com/dromara/carbon/compare/v2.6.3...v2.6.4) (2025-04-28)

### ✨ Features

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

### 🐛 Fixes

- Fix panic when database field type was `nil`

### ⚡ Performance

- Add `b.ResetTimer()` in  benchmark test files

### ♻️ Refactor & Changes

- Split  `database_types.go` into` type_carbon.go`, `type_layout.go`, `type_format.go`,  `type_timestamp.go`
- Rename  `LayoutFactory` interface to `LayoutTyper` and `SetLayout` method to `Layout`
- Rename  `FormatFactory` interface to `FormatTyper` and `SetFormat` method to `Format`
- Rename  `TimestampFactory` interface to `TimestampTyper` and `SetPrecision` method to `Precision`

## [v2.6.3](https://github.com/dromara/carbon/compare/v2.6.2...v2.6.3) (2025-04-21)

### ✨ Features

- Add `IsEpoch` method to report whether is a unix epoch time(1970-01-01 00:00:00 +0000 UTC).
- Add `WeekEndsAt` method to get end day of the week.
- Add `SetWeekendDays` method to set weekend days of the week.
- Add `DefaultWeekStartsAt` global variable to store  weekend days of the week.

### 🐛 Fixes

-  Fix bug with the same results in different countries using `IsWeekend` and `IsWeekday` methods.
-  Fix bug `carbon.Parse("").StdTime()` cause a nil pointer panic.  #294

### ⚡ Performance

- Add benchmark test files `xxx_bench_test.go`

### ♻️ Refactor & Changes

-  Change errors  from `private` method  to `public` method.
-  Change the default start date of the week from `Sunday` to `Monday`
-  Change the year of `MinValue` from `-9998` to `1`
-  Change `weeksPerLongYear` constant to `WeeksPerLongYear`

## [v2.6.2](https://github.com/dromara/carbon/compare/v2.6.1...v2.6.2) (2025-04-08)

### ✨ Features

-  Add format symbol `o` to get time zone offset
-  Add `TimestampLayout`, `TimestampMilliLayout`, `TimestampMicroLayout` and `TimestampNanoLayout` constants
-  Add `TimestampFormat`, `TimestampMilliFormat`, `TimestampMicroFormat` and `TimestampNanoFormat` constants
- Add `DateTimeMilli`, `DateTimeMicro`, `DateTimeNano` field types
- Add `DateMilli`、`DateMicro`、`DateNano` field types
- Add `TimeMilli`、`TimeMicro`、`TimeNano` field types
-  Add `MaxDuration`, `MinDuration` methods

### 🐛 Fixes

- Fix bug of incorrect judgment of `IsLeapMonth` in lunar calendar
- Fix bug of `AtomFormat` and `AtomLayout` return inconsistent value
- Fix bug of `RFC3339Format` and `RFC3339Layout` return inconsistent value
-  Fix missing timezone bug in `IsDST` method
- Fix missing timezone bug `StartOfXXX`,  `EndOfXXX` methods
- Fix missing timezone bug when converting other calendars to `Gregorian` calendar

### ♻️ Refactor & Changes

-  Remove  `hour`, `minute`,`second` parameter in `CreateFromLunar`,`CreateFromPersian ` methods
- Change some format symbol definitions involves symbols such as  `U`, `V`, `X`,`S`,`T` `Z`,`u`,`v`,`x`,`z`
-  `time.Local` is not updated when setting global time zone

## [v2.6.1](https://github.com/dromara/carbon/compare/v2.6.0...v2.6.1) (2025-03-27)

### ✨ Features

- Add `ParseWithLayouts` and `ParseWithFormats` methods

### 🐛 Fixes

- Fixed `updated_at` field is automatically updated and invalid when `gorm` updates data

### ♻️ Refactor & Changes

- Rename `formatFactory` interface to `FormatFactory`, `layoutFactory` interface to `LayoutFactory`, `timestampFactory` interface to `TimestampFactory`, and add type constraints
- Change the return value of the `GormDataType` method of the `LayoutType`, `FormatType`, `TimestampType`  `struct`  to `time`
- Change `DateTime`, `Date`, `Time` types from `struct` to `string`
- Change `Timestamp`、`TimestampMilli`、`TimestampMicro`,`TimestampNano` types from `struct` to `int64`
- Move built-in database field types to new file `types.go`

## [v2.6.0](https://github.com/dromara/carbon/compare/v2.5.4...v2.6.0) (2025-03-25)

### ✨ Features

-  Add `ZoneName` method to get the time zone name
- Add `HasError` method to check if there is an error
- Add `IsNil` method to check if it is `nil`
- Add `Copy` method to deep copy `carbon`
- Add `WeekStartsAt` method to get start day of the week
- Add example file `xxx_example.go`
- Add a new `constant.go` file and migrate constants from the `carbon.go` file to this file

### ♻️ Refactor & Changes

- `golang` minimum version dependency upgraded to `1.18`
- `carbon`, `julian`, `lunar`, `persian` changed from value passing to pointer passing
- Change default global timezone from `Local` to `UTC`
- Rename `Offset` method to `ZoneOffset` method
- Rename `IsSetTestNow` method to `IsTestNow` method
- Rename `UnSetTestNow` method to `CleanTestNow` method
- Remove `Location` method, replaced by `Timezone` method
- Change the judgment logic of `IsValid` and `IsInvalid` methods, `zero time` is no longer considered invalid time
- Update `time.Local` when setting the global default timezone
- Refactor `database.go` and remove `carbon.DateTime`, `carbon. DateTimeMilli`, `carbon.DateTimeMicro`,  `carbon.DateTimeNano`, `carbon. Date`, `carbon.DateMilli`, `carbon.DateMicro`,  `carbon.DateNano`, `carbon.Time`, `carbon.TimeMilli`,  `carbon.TimeMicro`, `carbon.TimeNano`, `carbon.Timestamp`, `carbon.TimestampMilli`, `carbon.TimestampMicro`, `carbon.TimestampNano` field types, use generic fields instead to implement custom output format when `MarshalJSON/UnmarshalJSON`

Refer to <a href="https://github.com/dromara/carbon/releases" target="_blank" rel="noreferrer">releases</a> for change logs of earlier versions
