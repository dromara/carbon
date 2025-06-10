# 更新日志

## [v2.6.7](https://github.com/dromara/carbon/compare/v2.6.6...v2.6.7) (2025-05-26)

-  `String` 方法去掉对空值的检查
- 将 `type_interface.go` 更名为 `interfaces.go`
- 将`Closest`/`Farthest` 方法第 2 个参数改成可选参数
- 新增 `ZeroValue`/`EpochValue` 方法
- 新增 `DataTyper` 接口和 `DataType`方法并让内置类型实现 `DataTyper` 接口

## [v2.6.6](https://github.com/dromara/carbon/compare/v2.6.5...v2.6.6) (2025-05-19)

- 修复在 `window` 平台无法找到语言文件的 bug
- 修复在创建新的 `Carbon` 实例时丢失`layout`、`weekStartsAt`、`weekendDays` 和 `lang 值的错误 #303
- 修复 `StartOfWeek` 和 `EndOfWeek`方法意外更改原始 `Carbon`实例的错误 #304
- 新增对 `xorm` 的 `curd` 集成测试，目前已覆盖 `MySQL`/`Postgres`/`SQLite`
- 在  `ci` 中新增 `window` 系统的单元测试

## [v2.6.5](https://github.com/dromara/carbon/compare/v2.6.4...v2.6.5) (2025-05-14)

- 将 `Go` 最低版本要求从 `1.18` 提升到 `1.21`
- `Carbon` 结构体的 `SetLanguage` 方法增加对非法 `Language` 结构体的判断
- `Carbon` 结构体的`Parse` 方法增加对 `Postgres`/`SQLite` 时间格式字符串的解析支持
- `Carbon` 结构体的`Parse`/`ParseByLayout`/`ParseByFormat`方法解析 `空字符串` 时返回值从 `nil` 更改成空 `carbon`
- `Carbon` 结构体新增 `IsEmpty` 方法用于判断是否为空 `carbon`
- `Carbon` 结构体新增 `ClearTestNow` 方法替代 `CleanTestNow`, `CleanTestNow` 方法未来将移除
- `Carbon` 结构体新增 `ParseByLayouts` 方法替代 `ParseWithLayouts`, `ParseWithLayouts` 方法未来将移除
- `Carbon` 结构体新增 `ParseByFormats` 方法替代 `ParseWithLayouts`, `ParseWithFormats` 方法未来将移除
- `Carbon` 结构体移除 `GormDataType `方法, 并将 `Value`/`MarshalJSON `方法从`指针`接收者改成`值`接收者
- `LayoutType[T]` 结构体移除 `GormDataType `方法, 并将 `Value`/`MarshalJSON `方法从`指针`接收者改成`值`接收者
- `FormatType[T]` 结构体移除 `GormDataType `方法, 并将 `Value`/`MarshalJSON `方法从`指针`接收者改成`值`接收者
- `TimestampType[T]` 结构体移除 `GormDataType `方法, 并将 `Value`/`MarshalJSON `方法从`指针`接收者改成`值`接收者
- `Language` 结构体的 `SetResources` 方法增加对非法资源的判断
- 新增对 `gorm` 的 `curd` 集成测试，目前已覆盖 `MySQL`/`Postgres`/`SQLite`
- 使用 `github.com/stretchr/testify/assert` 替换 `github.com/stretchr/testify/suite` 进行单元测试

## [v2.6.4](https://github.com/dromara/carbon/compare/v2.6.3...v2.6.4) (2025-04-28)

-  修复数据库字段类型为 `nil` 时抛出异常的bug
- 将 `database_types.go` 拆分成 `type_carbon.go`, `type_layout.go`, `type_format.go`, `type_timestamp.go`
- 将 `LayoutFactory`  接口重命名为 `LayoutTyper` 和  `SetLayout` 方法重命名为 `Layout`
- 将 `FormatFactory`  接口重命名为 `FormatTyper` 和 `SetFormat` 方法重命名为 `Format`
- 将 `TimestampFactory`  接口重命名为 `TimestampTyper` 和 `SetPrecision` 方法重命名为 `Precision`
- 性能测试文件增加 `b.ResetTimer()`
- `Language` 结构体新增 `Copy` 方法
- 新增 `carbon.Timestamp` 类型别名和 `carbon.NewTimestamp` 方法
- 新增 `carbon.TimestampMilli` 类型别名和 `carbon.NewTimestampMilli` 方法
- 新增 `carbon.TimestampMicro` 类型别名和 `carbon.NewTimestampMicro` 方法
- 新增 `carbon.TimestampNano` 类型别名和 `carbon.NewTimestampNano` 方法
- 新增 `carbon.DateTime` 类型别名和 `carbon.NewDateTime` 方法
- 新增 `carbon.DateTimeMicro` 类型别名和 `carbon.NewDateTimeMicro` 方法
- 新增 `carbon.DateTimeMilli` 类型别名和 `carbon.NewDateTimeMilli` 方法
- 新增 `carbon.DateTimeNano` 类型别名和 `carbon.NewDateTimeNano` 方法
- 新增 `carbon.Date` 类型别名和 `carbon.NewDate` 方法
- 新增 `carbon.DateMilli` 类型别名和 `carbon.NewDateMilli` 方法
- 新增 `carbon.DateMicro` 类型别名和 `carbon.NewDateMicro` 方法
- 新增 `carbon.DateNano` 类型别名和 `carbon.NewDateNano` 方法
- 新增 `carbon.Time` 类型别名和 `carbon.NewTime` 方法
- 新增 `carbon.TimeMilli` 类型别名和 `carbon.NewTimeMilli` 方法
- 新增 `carbon.TimeMicro` 类型别名和 `carbon.NewTimeMicro` 方法
- 新增 `carbon.TimeNano` 类型别名和 `carbon.NewTimeNano` 方法

## [v2.6.3](https://github.com/dromara/carbon/compare/v2.6.2...v2.6.3) (2025-04-21)

-  修复 `IsWeekend`, `IsWeekday` 方法不同国家返回结果一致的 bug
-  修复 `StdTime` 方法空指针引起的异常 #294
-  将错误方法由 `私有` 方法改成 `公开` 方法
-  将一周默认开始日期从 `周日` 改成 `周一`
-  将 `MinValue` 方法的年份从 `-9998` 更改为 `1`
-  将 `weeksPerLongYear` 常量更名为 `WeeksPerLongYear`
-  新增性能测试文件 `xxx_bench_test.go`
-  新增 `IsEpoch` 方法用于判断是否是 UNIX 纪元时间(1970-01-01 00:00:00 +0000 UTC)
-  新增 `WeekEndsAt` 方法用于获取一周的结束日期
-  新增 `SetWeekendDays` 方法用于设置一周周末日期
-  新增 `DefaultWeekStartsAt` 全局变量用于存储默认一周休息日

## [v2.6.2](https://github.com/dromara/carbon/compare/v2.6.1...v2.6.2) (2025-04-08)

-  `CreateFromLunar`, `CreateFromPersian` 方法去掉 hour, minute, second 参数
- 更改部分格式符号定义，涉及到的符号有 `U`, `V`, `X`,`S`,`T` `Z`,`u`,`v`,`x`,`z`
- 修复农历中 `IsLeapMonth` 判断错误的 bug
- 修复 `AtomFormat` 和 `AtomLayout` 格式返回值不一致的 bug
- 修复 `RFC3339Format` 和 `RFC3339Layout` 格式返回值不一致的 bug
- 设置全局默认时区时不再同步更新 `time.Local`
- 新增格式符号`o` 来获取时区偏移量
- 新增 `TimestampLayout`、`TimestampMilliLayout`、`TimestampMicroLayout` 和 `TimestampNanoLayout` 常量
- 新增 `TimestampFormat`、`TimestampMilliFormat`、`TimestampMicroFormat` 和 `TimestampNanoFormat` 常量
- 新增 `DateTimeMilli`、`DateTimeMicro`、`DateTimeNano` 字段类型
- 新增 `DateMilli`、`DateMicro`、`DateNano` 字段类型
- 新增 `TimeMilli`、`TimeMicro`、`TimeNano` 字段类型
-  修复 `IsDST` 方法丢失时区的 bug
-  修复 `StartOfXXX`、`EndOfXXX` 部分方法丢失时区的 bug
-  修复其他日历转化为公历时缺失时区的 bug
-  设置默认时区时不再同步更新  `time.Local`
-  新增 `MaxDuration`、`MinDuration` 方法

## [v2.6.1](https://github.com/dromara/carbon/compare/v2.6.0...v2.6.1) (2025-03-27)

-  新增 `ParseWithLayouts`  和 `ParseWithFormats` 方法
-  将 `formatFactory` 接口更名为 `FormatFactory`, `formatFactory` 接口更名为 `FormatFactory`, `formatFactory` 接口更名为 `FormatFactory`, 并添加类型约束
-  将 `LayoutType` ,`FormatType`, `TimestampType` 结构体`GormDataType` 方法的返回值更改为 `time`
-  将 `DateTime`、`Date`、`Time` 类型从 `struct` 更改为 `string`
-  将`Timestamp`、`TimestampMilli`、`TimestampMicro`,`TimestampNano` 类型从 `struct` 更改为 `int64`
-  将内置数据库字段类型移动到新文件 `types.go`
-  修复 `gorm` 更新数据时 `updated_at` 字段自动更新无效的 bug

## [v2.6.0](https://github.com/dromara/carbon/compare/v2.5.4...v2.6.0) (2025-03-25)

- `golang` 最低版本依赖升级到 `1.18`
- `carbon`, `julian`, `lunar`, `persian` 从值传递改成指针传递
-  新增 `ZoneName` 方法获取时区名称
-  新增 `HasError` 方法判断是否有错误
-  新增 `IsNil` 方法判断是否是 `nil`
-  新增 `Copy` 方法对 `carbon` 进行深度复制
-  新增 `WeekStartsAt` 方法获取周起始日期
-  新增示例文件 `xxx_example.go`
-  新增`constant.go` 文件，将常量从 `carbon.go` 文件迁移到此文件
-  默认全局时区从 `Local` 更改为 `UTC`
-  `Offset` 方法更名为 `ZoneOffset`
-  `IsSetTestNow` 方法更名为 `IsTestNow`
-  `UnSetTestNow ` 方法更名为 `CleanTestNow`
-  移除 `Location` 方法，由 `Timezone` 方法替代
-  更改 `IsValid` 和 `IsInvalid` 方法判断逻辑，`zero time` 不再视为无效时间
- 设置全局默认时区时同步更新 `time.Local`
- 重构 `database.go`，移除 `carbon.DateTime`、`carbon. DateTimeMilli `、 `carbon.DateTimeMicro`、`carbon.DateTimeNano`、 `carbon. Date`、`carbon.DateMilli`、 `carbon.DateMicro`、 `carbon.DateNano`、 `carbon.Time`、 `carbon.TimeMilli`、 `carbon.TimeMicro`、  `carbon.TimeNano`、`carbon.Timestamp` 、`carbon.TimestampMilli ` 、`carbon.TimestampMicro`、`carbon.TimestampNano` 字段类型, 使用泛型字段替代以实现 `MarshalJSON/UnmarshalJSON` 时自定义输出格式
