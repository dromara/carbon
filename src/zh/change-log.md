---
head:
  - - meta
    - name: description
      content: "Carbon 更新日志，详细记录历史版本的所有更新内容，包括新功能、性能优化、Bug 修复、API 变更、依赖升级和兼容性说明"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, 更新日志, 版本历史, 变更记录, 功能更新, Bug修复, 性能优化, API变更, 发布记录, 版本更新"
---

# 更新日志

## [v2.7.0](https://github.com/dromara/carbon/compare/v2.6.18...v2.7.0) (2026-10-19)

### ✨ 新功能

- **新增 `AddWeekdays`、`AddWeekday`、`SubWeekdays`、`SubWeekday` 方法，支持跳过周末日进行工作日计算**：按实例配置的周末日跳过非工作日进行日期推移，方便按工作日计算截止日期；四个方法命名与 `AddDays` / `AddDay` / `SubDays` / `SubDay` 保持一致，周末日从实例读取因此 `SetWeekendDays` 仍生效；当七天全部被标记为周末日时，每个工作日跨度被限制为整周以防止死循环（by [@darakanoit](https://github.com/darakanoit) in [#366](https://github.com/dromara/carbon/pull/366)）
- **新增 `IsSameWeek` 方法，判断两个日期是否在同一周**：补全 `IsSame*` 家族（`IsSameCentury`、`IsSameDecade`、`IsSameYear`、`IsSameQuarter`、`IsSameWeek`、`IsSameMonth`、`IsSameDay`、`IsSameHour`、`IsSameMinute`、`IsSameSecond`）；尊重 `SetWeekStartsAt` 设置，因此同一组日期在不同设置下可能不属于同一周；比较基于周起始日期，正确处理跨年份的边界周（by [@darakanoit](https://github.com/darakanoit) in [#368](https://github.com/dromara/carbon/pull/368)）
- **新增 `IsBirthday` 方法，比较月日判断是否生日**：判断两个日期是否落在同一月同日，忽略年份；无参数时与当前时间比较；2月29日仅匹配2月29日，平年永远不会是生日，与 PHP Carbon 行为一致（by [@darakanoit](https://github.com/darakanoit) in [#364](https://github.com/dromara/carbon/pull/364)）
- **新增 `Asia/Jakarta` 时区常量支持**（by [@mirzaq19](https://github.com/mirzaq19) in [#374](https://github.com/dromara/carbon/pull/374)）

### 🐛 修复

- **修复夏令时回退期间 `StartOfSecond` 和 `EndOfSecond` 方法选择错误时间出现的问题**：在重复的夏令时小时内，`time.Date` 重建可能选择同一本地时间的另一次出现；现在改为复制实例并调整现有时间戳的亚秒精度，保留选中的时刻、位置和配置（by [@mikamikasuki](https://github.com/mikamikasuki) in [#377](https://github.com/dromara/carbon/pull/377)）
- **修复夏令时回退期间 `StartOfMinute` 和 `EndOfMinute` 方法选择错误时间出现的问题**：与 #377 同一类 bug，只是上升了一级；现在复制实例并减去本地秒数，保留时刻、位置和偏移（by [@dualfroz](https://github.com/dualfroz) in [#378](https://github.com/dromara/carbon/pull/378)）
- **修复夏令时回退期间 `StartOfHour` 和 `EndOfHour` 方法选择错误时间出现的问题**：当重建的小时边界存在与接收器偏移一致的有效对应出现时，保留接收器的出现；不存在对应出现时（如 Lord Howe 的部分小时间隔回退）保持现有日历重建行为（by [@dajiaohuang](https://github.com/dajiaohuang) in [#382](https://github.com/dromara/carbon/pull/382)）
- **修复 `DiffInSeconds` 方法忽略亚秒部分导致跨秒边界计时不准确的问题**：减去整秒时间戳时忽略了小数部分，导致跨越秒边界的 0.2 秒间隙被计为整秒；现在当纳秒差符号相反时调整 ±1，按已流逝的整秒计数并朝零截断（by [@youdie006](https://github.com/youdie006) in [#379](https://github.com/dromara/carbon/pull/379)）
- **修复 `DiffInDays` 和 `DiffInWeeks` 方法跨越夏令时转换时计算错误的问题**：除以流逝秒数在跨越夏令时转换时给出错误结果；现在改为按日历日计数（与 `DiffInMonths` 按月计数方式一致，`AddDate` 加调整），`DiffInWeeks` 基于它计算；仅跨越夏令时转换的区间结果会变化（by [@youdie006](https://github.com/youdie006) in [#376](https://github.com/dromara/carbon/pull/376)）
- **修复 `DiffInYears` 和 `DiffInMonths` 方法在不同时区之间比较时估算错误的问题**：从各自日期的 `Year()` / `Month()` 字段进行首次估算时，两个日期处于不同位置可能导致偏差；现在先以较早日期所在位置读取较晚日期，再取年月字段，与 `DiffInDays` 使用的框架一致（by [@ybai08](https://github.com/ybai08) in [#384](https://github.com/dromara/carbon/pull/384)）
- **修复 `IsSameDecade` 方法未比较世纪导致跨世纪日期误判的问题**：`Decade()` 是世纪内的十年，因此 `Parse("1995-08-05").IsSameDecade(Parse("2095-08-05"))` 曾返回 true；现在也比较 `Century()`，与 `StartOfDecade` 使用的分桶一致（by [@youdie006](https://github.com/youdie006) in [#375](https://github.com/dromara/carbon/pull/375)）
- **修复 `FromStdTime` 方法在凌晨时刻因 UTC 偏移差异返回错误农历日期的问题**：以时间所在时区计算从 1900-01-31 起流逝的时间，两个时刻的偏移会泄漏到日计数中；现在改为在 UTC 中计算挂钟日期之间的天数，与 `persian.FromStdTime` 和 `hebrew.FromStdTime` 一致（by [@youdie006](https://github.com/youdie006) in [#380](https://github.com/dromara/carbon/pull/380)）
- **修复冻结时钟模式下并发调用 `Sleep` 方法导致时钟推进丢失的竞态问题**：冻结时钟上的并发 `Sleep` 调用彼此以及与 `Now` 发生竞态，导致时钟推进丢失；现在使用现有的冻结时钟互斥锁包裹检查和更新（by [@yuefdev](https://github.com/yuefdev) in [#383](https://github.com/dromara/carbon/pull/383)）

### 🔧 依赖升级

- 将 `github/codeql-action` 从 `v4.38.1` 升级到 `v4.38.3`（by [@dependabot\[bot\]](https://github.com/dependabot) in [#372](https://github.com/dromara/carbon/pull/372) [#381](https://github.com/dromara/carbon/pull/381)）

## [v2.6.18](https://github.com/dromara/carbon/compare/v2.6.17...v2.6.18) (2026-09-28)

### ✨ 新功能

- **新增对 `哈萨克语` 的本地化语言支持（`lang/kk.json`）**：月份名称、短月份名称、星期名称和短星期名称取自 PHP Carbon 的哈萨克语本地化文件（遵循 CLDR）；相对时间短语遵循哈萨克语语法，名词在数字后不变格；季节和星座使用传统哈萨克名称直译（by [@darakanoit](https://github.com/darakanoit) in [#355](https://github.com/dromara/carbon/pull/355)）

### 🐛 修复

- **修复 `Format` 方法处理末尾转义字符时可能发生越界访问的 bug**：在 `helper.go` 的 `format2layout` 方法中为转义字符处理添加边界检查，防止越界访问导致的 `panic`（by [@gouguoyin](https://github.com/gouguoyin) in [#348](https://github.com/dromara/carbon/pull/348)）
- **修复设置非周一为一周开始日期时，`ToWeekString`、`ToShortWeekString` 和 `Format` 方法返回星期名称错误的 bug**：`ToWeekString` / `ToShortWeekString` 曾通过 `DayOfWeek() % DaysPerWeek` 索引名称资源，但资源按绝对顺序（周日优先）排列；现在改用 `int(StdTime().Weekday())` 按绝对星期索引，与 `Format("l")` 行为一致（by [@sneha4175](https://github.com/sneha4175) in [#371](https://github.com/dromara/carbon/pull/371)）
- **修复向已有时间值的数据库字段扫描 `NULL` 时未清空原值的 bug**：向已包含时间值的 Carbon 扫描 SQL `NULL` 时返回成功但保留原日期；现在在 `Carbon.Scan` 及布局、格式和时间戳扫描器包装器中重置目标值（by [@vitalivo](https://github.com/vitalivo) in [#360](https://github.com/dromara/carbon/pull/360)）

### ⚡ 性能优化

- **优化语言实例和语言资源的缓存与共享机制，通过写时复制减少内存分配**：`SetLocale` 此前将完整的缓存语言映射复制到每个 `Language` 实例，导致每个 Carbon 实例分配全新的 20 项映射；现在 `SetLocale` 共享缓存映射并标记为共享，`SetResources` 在首次写入前复制，只读资源的方法继续使用共享映射；每个 `NewCarbon()` 的分配从 1600 B / 13 次降至 232 B / 4 次（by [@darakanoit](https://github.com/darakanoit) in [#350](https://github.com/dromara/carbon/pull/350)）
- **优化月份、星期、季节和星座本地化资源查询，避免每次调用分割字符串**：六个方法每次调用都读取 "|" 分隔的语言资源并使用 `strings.Split` 分割，仅为了索引一个元素；现在使用小型辅助函数通过 `strings.Count` 和 `strings.IndexByte` 原地定位元素，每次调用零分配；`Constellation` -61.54%，`ToShortWeekString` -47.35%，`ToWeekString` -44.13%，`Season` -40.65%，`ToMonthString` -37.66%，`ToShortMonthString` -35.88%（by [@darakanoit](https://github.com/darakanoit) in [#352](https://github.com/dromara/carbon/pull/352)）
- **每个 locale 共享一个语言实例**：`NewCarbon()` 此前为每个实例构建全新的 `Language`；现在 `defaultLanguage(locale)` 在 `sync.Map` 中缓存每个 locale 一个即用型 `Language`；构造 Carbon 的分配从 4 次降至 1 次（by [@darakanoit](https://github.com/darakanoit) in [#354](https://github.com/dromara/carbon/pull/354)）
- **优化农历日期转换，预计算年度天数和累计天数表**：农历转换此前遍历年份表而非索引，时间复杂度为 O(1900年以来的年数)；现在在每个支持的年份和累计总数在包初始化时推导一次；2020 日期转农历 1009 ns -> 282 ns（-72%），转回公历 738 ns -> 23 ns（-97%）（by [@darakanoit](https://github.com/darakanoit) in [#357](https://github.com/dromara/carbon/pull/357)）

### 🔧 依赖升级

- 将 `github.com/stretchr/testify` 从 `v1.11.1` 升级到 `v1.12.1`（by [@dependabot\[bot\]](https://github.com/dependabot) in [#341](https://github.com/dromara/carbon/pull/341) [#344](https://github.com/dromara/carbon/pull/344)）
- 将 `github/codeql-action` 从 `v4.37.4` 升级到 `v4.38.1`（by [@dependabot\[bot\]](https://github.com/dependabot) in [#340](https://github.com/dromara/carbon/pull/340) [#342](https://github.com/dromara/carbon/pull/342) [#345](https://github.com/dromara/carbon/pull/345) [#346](https://github.com/dromara/carbon/pull/346) [#362](https://github.com/dromara/carbon/pull/362) [#369](https://github.com/dromara/carbon/pull/369)）

### 🚀 CI

- **完善 `CI` 测试流程，在 `PR` 上运行完整测试和竞态检测，并仅在 `push` 时生成和上传覆盖率**：测试命令现在通过 `./...` 覆盖所有包，而非仅根包；覆盖率文件生成和 Codecov 上传限制为 `push` 事件；新增使用 `-race` 的全包竞态检测步骤（by [@gouguoyin](https://github.com/gouguoyin) in [#361](https://github.com/dromara/carbon/pull/361)）

## [v2.6.17](https://github.com/dromara/carbon/compare/v2.6.16...v2.6.17) (2026-08-09)

### 🐛 修复

- 修正`伊斯兰`星期名称拼写错误，将 `Jumaat` 修改为 `Jumat`
- 修复`公历`与`希伯来历`之间的换算计算异常问题

### 🔧 依赖升级

- 将 `github/codeql-action` 从 `v4` 升级至 `v4.37.4`
- 将 `actions/setup-go` 从 `v6` 升级至 `v7`
- 将 `actions/checkout` 从 `v6` 升级至 `v7`
- 将 `codecov/codecov-action` 从 `v5` 升级至 `v6`
- 将 `codecov/codecov-action` 从 `v6` 升级至 `v7`

### 📝 文档

- 将 `README` 文档统一替换为 `NRG` 标准化模板
- 清理 `README` 内冗余广告外链，移除已废弃的 `Go Report Card` 状态徽章

## [v2.6.16](https://github.com/dromara/carbon/compare/v2.6.15...v2.6.16) (2026-01-28)

### 🐛 修复

- 修复 `Format` 方法中`u`、`v`、`x`符号解析错误的 bug

### 🔧 依赖升级

- 将 `actions/checkout` 从 `v5` 升级到 `v6`

## [v2.6.15](https://github.com/dromara/carbon/compare/v2.6.14...v2.6.15) (2025-11-15)

### ✨ 新功能

- 增加 `HttpLayout`, `HttpFormat` 常量以及 `ToHttpString` 方法
- 增加对 `蒙古语` 的本地化语言支持(lang/mn.json)
- 增加对 `南非语` 的本地化语言支持(lang/af.json)

### 🐛 修复

- 移除 `difference.go` 文件的 `DiffAbsInDuration` 方法里使用的从 `go1.19` 才有的 `Abs` 方法
- 修复 `荷兰语` 本地化语言文件(lang/nl.json)里 `weeks` 字段里重复的 `Zondag` 错误

## [v2.6.14](https://github.com/dromara/carbon/compare/v2.6.13...v2.6.14) (2025-10-28)

### ✨ 新功能

- 增加对 `希腊语` 的本地化语言支持(lang/el.json)
- 增加对 `芬兰语` 的本地化语言支持(lang/fi.json)
- 增加对 `缅甸语` 的本地化语言支持(lang/my.json)

### 🐛 修复

- 优化 `traveler.go` 文件里时间增减系列方法通过复制实例来保证不可变性，以避免修改原始实例

## [v2.6.13](https://github.com/dromara/carbon/compare/v2.6.12...v2.6.13) (2025-10-15)

### 🐛 修复

- 在 `helper.go` 的 `format2layout` 方法中为转义字符处理添加边界检查，防止越界访问导致的 `panic`

### ⚡ 性能优化

- 在 `language.go` 的 `SetLocale` 方法中使用 `sync.Once` 确保语言文件只加载一次，使用 `sync.Map` 进行线程安全的缓存

## [v2.6.12](https://github.com/dromara/carbon/compare/v2.6.11...v2.6.12) (2025-09-16)

### 🐛 修复

- 在 `type_carbon.go` 文件 `UnmarshalJSON` 方法中设置 `isEmpty` 标志以表示空值
- 修复潜在的竞态条件和空指针解引用问题，提高并发安全性

### ⚡ 性能优化

- 使用 `sync.Map` 实现高性能并发缓存

### 🔧 依赖升级

- 将 `golang` 环境依赖从 `1.21` 降低到 `1.18`
- 将 `testify` 测试框架从 `v1.10.0` 升级到 `v1.11.1`

## [v2.6.11](https://github.com/dromara/carbon/compare/v2.6.10...v2.6.11) (2025-07-18)

### ✨ 新功能

- 新增 `希伯来历` 支持并添加单元测试和基准测试文件

### ♻️ 重构与变更

- 将 `Sleep` 由结构体方法更改成全局方法
- 重构 `波斯历` 并添加基准测试文件
- 重构 `中国农历` 并添加基准测试文件
- 重构 `儒略日/简化儒略日` 并添加基准测试文件

### 📝 文档

- 新增整体性能测试报告文件

## [v2.6.10](https://github.com/dromara/carbon/compare/v2.6.9...v2.6.10) (2025-07-07)

### ✨ 新功能

- 新增 `Sleep` 方法及相关`单元测试`、`基准测试`和`示例文件`
- 新增数字常量，如 `MaxYear`, `MinYear`, `MaxMonth`, `MinMonth`, `MaxDay`, `MinDay` 等，并使用这些常量替换硬编码

### ⚡ 性能优化

- 优化 `helper.go` 里 `getAbsValue` 方法，用`位操作`替换条件判断
- 优化 `frozen.go` 文件里时间冻结相关方法，用`原子操作`减少锁竞争，优化内存分配
- 优化基准测试文件，覆盖`串行测试`、`并行测试`和`并发测试`

### ♻️ 重构与变更

- 将`日语`翻译文件从 `jp.json` 改成 `ja.json`，说明文档从 `README.jp.md` 更名为 `README.ja.md`，以符合 [ISO639-1](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes) 标准
- 移除已弃用的 `ParseWithLayouts` 方法，用 `ParseByLayouts` 方法替代
- 移除已弃用的 `ParseWithFormats` 方法，用 `ParseByFormats` 方法替代
- 移除已弃用的 `CleanTestNow` 方法，用 `ClearTestNow` 方法替代
- 移除 `ParseByLayout` 和 `ParseByFormat` 方法对`时间戳`字符串的解析支持，解析`时间戳`请使用 `CreateFromTimestamp`, `CreateFromTimestampMilli`, `CreateFromTimestampMicro`, `CreateFromTimestampNano` 方法

### 📝 文档

- 新增韩语文档 `README.ko.md`

## [v2.6.9](https://github.com/dromara/carbon/compare/v2.6.8...v2.6.9) (2025-06-28)

### ♻️ 重构与变更

- 移除对 `gorm` 的 `GormDataType` 接口的实现

## [v2.6.8](https://github.com/dromara/carbon/compare/v2.6.7...v2.6.8) (2025-06-12)

### 🐛 修复

- 解析时如果布局模板或格式模板为空时，返回错误

### ♻️ 重构与变更

- 在 `type_builtin.go` 中将 `DateTimeType` 重命名为`dateTimeType`，将 `DateTimeXXXType` 重命名为 `dateTimeXXXType`
- 在 `type_builtin.go` 中将 `DateType` 重命名为`dateType`，将 `DateXXXType` 重命名为 `dateXXXType`
- 在 `type_builtin.go` 中将 `TimeType` 重命名为`timeType`，将 `TimeXXXType` 重命名为 `timeXXXType`

### 🔧 依赖升级

- 在 `tests` 中将 `gorm.io/gorm` 从 `1.21.1` 升级到 `1.30.0`
- 在 `tests` 中将 `gorm.io/driver/mysql` 从 `1.5.7` 升级到 `1.6.0`
- 在 `tests` 中将 `gorm.io/driver/postgres` 从 `1.5.7` 升级到 `1.6.0`
- 在 `tests` 中将 `gorm.io/driver/sqlite` 从 `1.5.7` 升级到 `1.6.0`

### 📝 文档

- 简化 READEME 文件，将详细使用说明和示例用法迁移到 [官方网站](https://carbon.go-pkg.com)
- 添加 [HelloGitHub](https://hellogithub.com/repository/dromara/carbon) 徽章链接

## [v2.6.7](https://github.com/dromara/carbon/compare/v2.6.6...v2.6.7) (2025-05-26)

### ✨ 新功能

- 将`Closest`/`Farthest` 方法第 2 个参数改成可选参数
- 新增 `ZeroValue`/`EpochValue` 方法
- 新增 `DataTyper` 接口和 `DataType`方法并让内置类型实现 `DataTyper` 接口

### ♻️ 重构与变更

- `String` 方法去掉对空值的检查
- 将 `type_interface.go` 更名为 `interfaces.go`

## [v2.6.6](https://github.com/dromara/carbon/compare/v2.6.5...v2.6.6) (2025-05-19)

### ✨ 新功能

- 新增对 `xorm` 的 `curd` 集成测试，目前已覆盖 `MySQL`/`Postgres`/`SQLite`

### 🐛 修复

- 修复在 `window` 平台无法找到语言文件的 bug
- 修复在创建新的 `Carbon` 实例时丢失`layout`、`weekStartsAt`、`weekendDays` 和 `lang` 值的错误
- 修复 `StartOfWeek` 和 `EndOfWeek` 方法意外更改原始 `Carbon` 实例的错误

### 🚀 CI

- 在 `ci` 中新增 `window` 系统的单元测试

## [v2.6.5](https://github.com/dromara/carbon/compare/v2.6.4...v2.6.5) (2025-05-14)

### ✨ 新功能

- `Carbon` 结构体的 `SetLanguage` 方法增加对非法 `Language` 结构体的判断
- `Carbon` 结构体的`Parse` 方法增加对 `Postgres`/`SQLite` 时间格式字符串的解析支持
- `Carbon` 结构体的`Parse`/`ParseByLayout`/`ParseByFormat`方法解析 `空字符串` 时返回值从 `nil` 更改成空 `carbon`
- `Carbon` 结构体新增 `IsEmpty` 方法用于判断是否为空 `carbon`
- `Carbon` 结构体新增 `ClearTestNow` 方法替代 `CleanTestNow`, `CleanTestNow` 方法未来将移除
- `Carbon` 结构体新增 `ParseByLayouts` 方法替代 `ParseWithLayouts`, `ParseWithLayouts` 方法未来将移除
- `Carbon` 结构体新增 `ParseByFormats` 方法替代 `ParseWithLayouts`, `ParseWithFormats` 方法未来将移除
- `Language` 结构体的 `SetResources` 方法增加对非法资源的判断
- 新增对 `gorm` 的 `curd` 集成测试，目前已覆盖 `MySQL`/`Postgres`/`SQLite`

### ♻️ 重构与变更

- `Carbon` 结构体移除 `GormDataType `方法, 并将 `Value`/`MarshalJSON `方法从`指针`接收者改成`值`接收者
- `LayoutType[T]` 结构体移除 `GormDataType `方法, 并将 `Value`/`MarshalJSON `方法从`指针`接收者改成`值`接收者
- `FormatType[T]` 结构体移除 `GormDataType `方法, 并将 `Value`/`MarshalJSON `方法从`指针`接收者改成`值`接收者
- `TimestampType[T]` 结构体移除 `GormDataType `方法, 并将 `Value`/`MarshalJSON `方法从`指针`接收者改成`值`接收者
- 使用 `github.com/stretchr/testify/assert` 替换 `github.com/stretchr/testify/suite` 进行单元测试

### 🔧 依赖升级

- 将 `Go` 最低版本要求从 `1.18` 提升到 `1.21`

## [v2.6.4](https://github.com/dromara/carbon/compare/v2.6.3...v2.6.4) (2025-04-28)

### ✨ 新功能

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

### 🐛 修复

- 修复数据库字段类型为 `nil` 时抛出异常的bug

### ⚡ 性能优化

- 性能测试文件增加 `b.ResetTimer()`

### ♻️ 重构与变更

- 将 `database_types.go` 拆分成 `type_carbon.go`, `type_layout.go`, `type_format.go`, `type_timestamp.go`
- 将 `LayoutFactory`  接口重命名为 `LayoutTyper` 和  `SetLayout` 方法重命名为 `Layout`
- 将 `FormatFactory`  接口重命名为 `FormatTyper` 和 `SetFormat` 方法重命名为 `Format`
- 将 `TimestampFactory`  接口重命名为 `TimestampTyper` 和 `SetPrecision` 方法重命名为 `Precision`

## [v2.6.3](https://github.com/dromara/carbon/compare/v2.6.2...v2.6.3) (2025-04-21)

### ✨ 新功能

- 新增 `IsEpoch` 方法用于判断是否是 UNIX 纪元时间(1970-01-01 00:00:00 +0000 UTC)
- 新增 `WeekEndsAt` 方法用于获取一周的结束日期
- 新增 `SetWeekendDays` 方法用于设置一周周末日期
- 新增 `DefaultWeekStartsAt` 全局变量用于存储默认一周休息日

### 🐛 修复

- 修复 `IsWeekend`, `IsWeekday` 方法不同国家返回结果一致的 bug
- 修复 `StdTime` 方法空指针引起的异常 #294

### ⚡ 性能优化

- 新增性能测试文件 `xxx_bench_test.go`

### ♻️ 重构与变更

- 将错误方法由 `私有` 方法改成 `公开` 方法
- 将一周默认开始日期从 `周日` 改成 `周一`
- 将 `MinValue` 方法的年份从 `-9998` 更改为 `1`
- 将 `weeksPerLongYear` 常量更名为 `WeeksPerLongYear`

## [v2.6.2](https://github.com/dromara/carbon/compare/v2.6.1...v2.6.2) (2025-04-08)

### ✨ 新功能

- 新增格式符号`o` 来获取时区偏移量
- 新增 `TimestampLayout`、`TimestampMilliLayout`、`TimestampMicroLayout` 和 `TimestampNanoLayout` 常量
- 新增 `TimestampFormat`、`TimestampMilliFormat`、`TimestampMicroFormat` 和 `TimestampNanoFormat` 常量
- 新增 `DateTimeMilli`、`DateTimeMicro`、`DateTimeNano` 字段类型
- 新增 `DateMilli`、`DateMicro`、`DateNano` 字段类型
- 新增 `TimeMilli`、`TimeMicro`、`TimeNano` 字段类型
- 新增 `MaxDuration`、`MinDuration` 方法

### 🐛 修复

- 修复农历中 `IsLeapMonth` 判断错误的 bug
- 修复 `AtomFormat` 和 `AtomLayout` 格式返回值不一致的 bug
- 修复 `RFC3339Format` 和 `RFC3339Layout` 格式返回值不一致的 bug
- 修复 `IsDST` 方法丢失时区的 bug
- 修复 `StartOfXXX`、`EndOfXXX` 部分方法丢失时区的 bug
- 修复其他日历转化为公历时缺失时区的 bug

### ♻️ 重构与变更

- `CreateFromLunar`, `CreateFromPersian` 方法去掉 hour, minute, second 参数
- 更改部分格式符号定义，涉及到的符号有 `U`, `V`, `X`,`S`,`T` `Z`,`u`,`v`,`x`,`z`
- 设置全局默认时区时不再同步更新 `time.Local`

## [v2.6.1](https://github.com/dromara/carbon/compare/v2.6.0...v2.6.1) (2025-03-27)

### ✨ 新功能

- 新增 `ParseWithLayouts`  和 `ParseWithFormats` 方法

### 🐛 修复

- 修复 `gorm` 更新数据时 `updated_at` 字段自动更新无效的 bug

### ♻️ 重构与变更

- 将 `layoutFactory` 接口更名为 `LayoutFactory`, `formatFactory` 接口更名为 `FormatFactory`, `timestampFactory` 接口更名为 `TimestampFactory`, 并添加类型约束
- 将 `LayoutType` ,`FormatType`, `TimestampType` 结构体`GormDataType` 方法的返回值更改为 `time`
- 将 `DateTime`、`Date`、`Time` 类型从 `struct` 更改为 `string`
- 将`Timestamp`、`TimestampMilli`、`TimestampMicro`,`TimestampNano` 类型从 `struct` 更改为 `int64`
- 将内置数据库字段类型移动到新文件 `types.go`

## [v2.6.0](https://github.com/dromara/carbon/compare/v2.5.4...v2.6.0) (2025-03-25)

### ✨ 新功能

- 新增 `ZoneName` 方法获取时区名称
- 新增 `HasError` 方法判断是否有错误
- 新增 `IsNil` 方法判断是否是 `nil`
- 新增 `Copy` 方法对 `carbon` 进行深度复制
- 新增 `WeekStartsAt` 方法获取周起始日期
- 新增示例文件 `xxx_example.go`
- 新增`constant.go` 文件，将常量从 `carbon.go` 文件迁移到此文件

### ♻️ 重构与变更

- `golang` 最低版本依赖升级到 `1.18`
- `carbon`, `julian`, `lunar`, `persian` 从值传递改成指针传递
- 默认全局时区从 `Local` 更改为 `UTC`
- `Offset` 方法更名为 `ZoneOffset`
- `IsSetTestNow` 方法更名为 `IsTestNow`
- `UnSetTestNow ` 方法更名为 `CleanTestNow`
- 移除 `Location` 方法，由 `Timezone` 方法替代
- 更改 `IsValid` 和 `IsInvalid` 方法判断逻辑，`zero time` 不再视为无效时间
- 设置全局默认时区时同步更新 `time.Local`
- 重构 `database.go`，移除 `carbon.DateTime`、`carbon. DateTimeMilli `、 `carbon.DateTimeMicro`、`carbon.DateTimeNano`、 `carbon. Date`、`carbon.DateMilli`、 `carbon.DateMicro`、 `carbon.DateNano`、 `carbon.Time`、 `carbon.TimeMilli`、 `carbon.TimeMicro`、  `carbon.TimeNano`、`carbon.Timestamp` 、`carbon.TimestampMilli ` 、`carbon.TimestampMicro`、`carbon.TimestampNano` 字段类型, 使用泛型字段替代以实现 `MarshalJSON/UnmarshalJSON` 时自定义输出格式

有关更早版本的更新日志，请参阅 <a href="https://github.com/dromara/carbon/releases" target="_blank" rel="noreferrer">releases</a>
