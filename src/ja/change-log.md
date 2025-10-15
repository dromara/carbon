# 変更履歴

## [v2.6.13](https://github.com/dromara/carbon/compare/v2.6.12...v2.6.13) (2025-10-15)

- `language.go` の `SetLocale` メソッドで `sync.Once` を使用して言語ファイルが一度だけ読み込まれることを保証し、`sync.Map` を使用してスレッドセーフなキャッシュを実装
- `helper.go` の `format2layout` メソッドでエスケープ文字処理の境界チェックを追加し、範囲外アクセスによる `panic` を防止

## [v2.6.12](https://github.com/dromara/carbon/compare/v2.6.11...v2.6.12) (2025-09-16)

- `golang` 環境依存を `1.21` から `1.18` に下げる
- `testify` テストフレームワークを `v1.10.0` から `v1.11.1` にアップグレード
- `type_carbon.go` ファイルの `UnmarshalJSON` メソッドで `isEmpty` フラグを設定して空値を示す
- `sync.Map` を使用して高性能な並行キャッシュを実装
- 潜在的な競合状態とnullポインタ逆参照の問題を修正し、並行安全性を向上

## [v2.6.11](https://github.com/dromara/carbon/compare/v2.6.10...v2.6.11) (2025-07-18)

- `Sleep` を構造体メソッドからグローバルメソッドに変更
- `ペルシア暦` をリファクタリングし、ベンチマークテストファイルを追加
- `中国旧暦` をリファクタリングし、ベンチマークテストファイルを追加
- `ユリウス日/修正ユリウス日` をリファクタリングし、ベンチマークテストファイルを追加
- `ヘブライ暦` サポートを追加し、ユニットテストとベンチマークテストファイルを追加
- 全体のパフォーマンステストレポートファイルを追加

## [v2.6.10](https://github.com/dromara/carbon/compare/v2.6.9...v2.6.10) (2025-07-07)

- `日本語` 翻訳ファイルを `jp.json` から `ja.json` に変更し、ドキュメントを `README.jp.md` から `README.ja.md` に名前変更して [ISO639-1](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes) 標準に準拠
- 非推奨の `ParseWithLayouts` メソッドを削除し、`ParseByLayouts` メソッドに置き換え
- 非推奨の `ParseWithFormats` メソッドを削除し、`ParseByFormats` メソッドに置き換え
- 非推奨の `CleanTestNow` メソッドを削除し、`ClearTestNow` メソッドに置き換え
- `ParseByLayout` と `ParseByFormat` メソッドからタイムスタンプ文字列の解析サポートを削除、タイムスタンプの解析には `CreateFromTimestamp`、`CreateFromTimestampMilli`、`CreateFromTimestampMicro`、`CreateFromTimestampNano` メソッドを使用
- `helper.go` の `getAbsValue` メソッドを最適化し、条件判断をビット演算に置き換え
- `frozen.go` ファイルの時間凍結関連メソッドを最適化し、原子操作を使用してロック競合を減らし、メモリ割り当てを最適化
- ベンチマークテストファイルを最適化し、`シリアルテスト`、`パラレルテスト`、`コンカレントテスト`をカバー
- 韓国語ドキュメント `README.ko.md` を追加
- `Sleep` メソッドと関連する`ユニットテスト`、`ベンチマークテスト`、`サンプルファイル`を追加
- `MaxYear`、`MinYear`、`MaxMonth`、`MinMonth`、`MaxDay`、`MinDay` などの数値定数を追加し、これらの定数でハードコードを置き換え

## [v2.6.9](https://github.com/dromara/carbon/compare/v2.6.8...v2.6.9) (2025-06-28)

- `gorm` の `GormDataType` インタフェースの実装の削除

## [v2.6.8](https://github.com/dromara/carbon/compare/v2.6.7...v2.6.8) (2025-06-12)

- 解析時にレイアウトテンプレートまたはフォーマットテンプレートが空の場合は、エラーを返します
- `tests` で `gorm.io/gorm` を `1.21.1` から `1.30.0` にアップグレードする
- `tests` で `gorm.io/driver/mysql` を `1.5.7` から `1.6.0` にアップグレード
- `tests` で `gorm.io/driver/postgres` を `1.5.7` から `1.6.0` にアップグレード
- `tests` で `gorm.io/driver/sqlite` を `1.5.7` から `1.6.0` にアップグレード
- `type_builtin.go` では` DateTimeType` を `dateTimeType` に、`DateTimeXXXXXType` を `dateTimeXXXXType` に、それぞれ名前を変更します
- `type_builtin.go` では` DateType` を `dateType` に、`DateXXXXXType` を `dateXXXType` に、それぞれ名前を変更します
- `type_builtin.go` では` TimeType` を `timeType` に、`TimeXXXXType` を `timeXXXType` に、それぞれ名前を変更します
- `README` ファイルをシンプル化し、詳細な使用方法と使用例を［公式サイト］に移行（https://carbon.go-pkg.com)
- [HelloGitHub]（https://hellogithub.com/repository/dromara/carbon）バッジリンク

## [v2.6.7](https://github.com/dromara/carbon/compare/v2.6.6...v2.6.7) (2025-05-26)

- `String` メソッドは、null値の炭素のチェックを取り除く
- ファイル名を `type_interface.go` に `interfaces.go` に変更
- オプションパラメータに `Closest`/ `Farthest` メソッドの2番目のパラメータを変更する
- `ZeroValue`/`EpochValue` メソッドを追加
- `DataTyper` インターフェースを追加し、`DataTyper` インターフェースを実装するために内蔵型を有効にします

## [v2.6.6](https://github.com/dromara/carbon/compare/v2.6.5...v2.6.6) (2025-05-19)

- 修正バグローカルファイルは `Windows` OSで見つからない
- 新しい `Carbon` インスタンス #303 を作成すると `layout`, `weekStartsAt`, `weekendDays` および `lang` の値の失いを修正します
- 修正バグ `StartOfWeek` と `EndOfWeek` メソッドは元の `Carbon` インスタンスを予期しないように変更します
- 現在「MySQL」/「Postgres」/「SQLite」をカバーしている「xorm」のための「curd」統合テストを追加
- `ci` アクションで `Windows` OS のユニットテストを追加

## [v2.6.5](https://github.com/dromara/carbon/compare/v2.6.4...v2.6.5) (2025-05-14)

- `go` バージョンを `1.18` から `1.21` にアップグレード
- `Carbon` 構造の `SetLanguage` メソッドは無効な `Language` 構造のチェックを追加します
- `Carbon` 構造の `Parse` メソッドは `Postgres` / `SQLite` タイムフォーマット文字列のサポートを追加します
- `Parse`/`ParseByLayout`/`ParseByFormat` のメソッドは `empty` 文字列を解析すると `nil` から `empty` carbon に戻り値を変更します
- `Carbon` 構造は `IsEmpty` メソッドを追加して `Empty` カーボンであるかどうかを判断します
- `Carbon` struct `CleanTestNow` を置き換える `ClearTestNow` メソッドを追加します。 `CleanTestNow` は将来削除されます。
- `Carbon` struct add `ParseByLayouts` method to replace `ParseWithLayouts`, `ParseWithLayouts` will be removed in the future.
- `Carbon` struct add `ParseByFormats` method to replace `ParseWithFormats` 、 `ParseWithFormats` は将来削除されます
- `Carbon` struct `GormDataType` メソッドを削除し `Value` / `MarshalJSON` メソッドを `pointer` 受信機から `value` 受信機に変更します
- `LayoutType[T]` struct `GormDataType` メソッドを削除し `Value` / `MarshalJSON` メソッドを `pointer` 受信機から `value` 受信機に変更します
- `FormatType[T]` struct `GormDataType` メソッドを削除し `Value` / `MarshalJSON` メソッドを `pointer` 受信機から `value` 受信機に変更します
- `TimestampType[T]` struct `GormDataType` メソッドを削除し `Value` / `MarshalJSON` メソッドを `pointer` 受信機から `value` 受信機に変更します
- `Language` 構造の `SetResources` メソッドは無効なリソースのチェックを追加します
- 現在 `MySQL` / `Postgres` / `SQLite` をカバーしている `gorm` の `curd` 統合テストを追加
- ユニットテストのために `github.com/stretchr/testify/assert` を `github.com/stretchr/testify/suite` に置き換えます

## [v2.6.4](https://github.com/dromara/carbon/compare/v2.6.3...v2.6.4) (2025-04-28)

- データベースフィールドタイプが `nil' だったときのパニックを修正
- 分割 `database_types.go` に `type_carbon.go`, `type_layout.go`, `type_format.go`, `type_timestamp.go`
- `LayoutFactory` インターフェイスを `LayoutTyper` に、`SetLayout` メソッドを `Layout` に変更
- `FormatFactory` インターフェイスを `FormatTyper` に、`SetFormat` メソッドを `Format` に変更
- `TimestampFactory` インターフェイスを `TimestampTyper` に、`SetPrecision` メソッドを `Precision` に変更
- ベンチマークテストファイルに `b.ResetTimer()` を追加する
- `Language` 構造の `Copy` メソッドを追加
- `carbon.Timestamp` 型エイリアスと `carbon.NewTimestamp` メソッドを追加
- `carbon.TimestampMilli` 型エイリアスと `carbon.NewTimestampMilli` メソッドを追加
- `carbon.TimestampMicro` 型エイリアスと `carbon.NewTimestampMicro` メソッドを追加
- `carbon.TimestampNano` 型エイリアスと `carbon.NewTimestampNano` メソッドを追加
- 炭素を追加します。DateTime タイプアリアスと carbonNewDateTime メソッド
- 炭素を追加します。DateTimeMicro タイプアリアスと `carbon.NewDateTimeMicro` メソッド
- 炭素を追加します。DateTimeMilli 型の別名と `carbon.NewDateTimeMilli` メソッド
- 炭素を追加します。DateTimeNano `タイプアリアスと `炭素.NewDateTimeNano` メソッド
- 炭素を追加します。日付のタイプの別名と炭素。NewDate メソッド
- 炭素を追加します。DateMilli 型の別名と `carbon.NewDateMilli` メソッド
- 炭素を追加します。DateMicro タイプの別名と `carbon.NewDateMicro` メソッド
- 炭素を追加します。DateNano のタイプアリアスと `carbon.NewDateNano` メソッド
- 炭素を追加します。タイムタイプの別名と炭素。NewTime メソッド
- 炭素を追加します。TimeMilli`タイプの別名と`炭素。NewTimeMilli` メソッド
- 炭素を追加します。タイムマイクロのタイプ別名とカーボン。NewTimeMicro メソッド
- 炭素を追加します。TimeNano`タイプ別名と`炭素。NewTimeNanoメソッド

## [v2.6.3](https://github.com/dromara/carbon/compare/v2.6.2...v2.6.3) (2025-04-21)

- 異なる国で同じ結果のバグを修正 `IsWeekend` と `IsWeekday` メソッドを使用します。
- バグ`カーボンを修正します。解析("").StdTime() はポインターのパニックを引き起こします。  #294
  エラーを「private」メソッドから「public」メソッドに変更します
- 週のデフォルトの開始日を `日曜日` から `月曜日` に変更する
- MinValue の年を `-9998` から `1` に変更する
- 定数 `weeksPerLongYear` を `WeeksPerLongYear` に変更する
- ベンチマークテストファイル `xxx_bench_test.go` を追加
- ユニックス時代であるかどうかを報告するために `IsEpoch` メソッドを追加します(1970-01-01 00:00:00 +0000 UTC)
- 週末の日を得るために `WeekEndsAt` メソッドを追加します
- 週末の日を設定するために `SetWeekendDays` メソッドを追加します
- 週の週末の日を格納するために `DefaultWeekStartsAt` グローバル変数を追加します

## [v2.6.2](https://github.com/dromara/carbon/compare/v2.6.1...v2.6.2) (2025-04-08)

- CreateFromLunar、CreateFromPersian メソッドの `hour`、 `minute`、 `second` パラメータを削除します。
- フォーマットシンボルの定義を変更するには `U`, `V`, `X`, `S`, `T` `Z`, `u`, `v`, `x`, `z` などのシンボルが含まれています
- 月カレンダーの `IsLeapMonth` の不正判断のバグを修正
- 修正 `AtomFormat` と `AtomLayout` のバグが一致しない値を返します
- 修正 `RFC3339Format` と `RFC3339Layout` のバグが一致しない値を返します
- 時間ローカルはグローバルタイムゾーンを設定すると更新されません
- タイムゾーンオフセットを取得するためにフォーマットシンボル `o` を追加します
- `TimestampLayout`, `TimestampMilliLayout`, `TimestampMicroLayout`, `TimestampNanoLayout` 定数を追加
- 追加 `TimestampFormat`, `TimestampMilliFormat`, `TimestampMicroFormat` および `TimestampNanoFormat` 定数
- `DateTimeMilli`, `DateTimeMicro`, `DateTimeNano` フィールドタイプを追加
- `DateMilli`、`DateMicro`、`DateNano` フィールドタイプを追加
- `TimeMilli`、`TimeMicro`、`TimeNano` フィールドタイプを追加
- `IsDST` メソッドのタイムゾーンバグを修正
- タイムゾーンバグ `StartOfXXX` 、 `EndOfXXX` メソッドが欠けていることを修正します
- 他のカレンダーを `グレゴリアン` カレンダーに変換するときに欠けているタイムゾーンのバグを修正
- 時間デフォルトタイムゾーンを設定するときにローカルは更新されません
- `MaxDuration`、`MinDuration` メソッドを追加

## [v2.6.1](https://github.com/dromara/carbon/compare/v2.6.0...v2.6.1) (2025-03-27)

- 追加 `ParseWithLayouts` と `ParseWithFormats` メソッド
- `formatFactory` インターフェイスを `FormatFactory` に、`layoutFactory` インターフェイスを `LayoutFactory` に、`timestampFactory` インターフェイスを `TimestampFactory` に名前変更し、タイプ制約を追加
- LayoutType、FormatType、TimestampType、struct メソッドの返り値を time に変更します。
- `DateTime`, `Date`, `Time` タイプを `struct` から `string` に変更する
- `TimestampMilli`、`TimestampMicro`、`TimestampNano` タイプを `struct` から `int64` に変更する
- 内蔵データベースフィールドタイプを新しいファイル `types.go` に移動します
- 修正された `updated_at` フィールドは `gorm` がデータを更新すると自動的に更新され、無効です

## [v2.6.0](https://github.com/dromara/carbon/compare/v2.5.4...v2.6.0) (2025-03-25)

- `golang` 最低バージョン依存性を `1.18` にアップグレードしました
- `carbon`, `julian`, `lunar`, `persian` は値のパスからポインターのパスに変更された
- タイムゾーン名を取得するために `ZoneName` メソッドを追加します
- エラーがあるかどうかを確認するために `HasError` メソッドを追加します
- それが `nil` であるかどうかを確認するために `IsNil` メソッドを追加します
- 深いコピー `炭素` に `コピー` メソッドを追加
- 週の開始日を得るために `WeekStartsAt` メソッドを追加します
- 例ファイル `xxx_example.go` を追加する
- 新しい `constant.go` ファイルを追加し、 `carbon.go` ファイルから定数をこのファイルに移行
- デフォルトのグローバルタイムゾーンを `Local` から `UTC` に変更する
- オフセットメソッドをZoneOffsetメソッドに変更
- IsSetTestNow メソッドをIsTestNow メソッドに変更
- UnSetTestNow メソッドをCleanTestNow メソッドに変更
- `Location` メソッドを削除し、`Timezone` メソッドに置き換えます
- `IsValid` と `IsInvalid` メソッドの判断論理を変更し、`ゼロタイム` はもはや無効タイムとみなされません
- 更新 `時間。ローカル` グローバルデフォルトタイムゾーンを設定するとき
- リファクター `database.go` および削除 `炭素。日付時間、炭素。DateTimeMilli、カーボン。DateTimeMicro、カーボン。DateTimeNanoは、カーボンです。日付、炭素。DateMilli、カーボン。デートマイクロ、カーボン。デートナノ、カーボン。時間、炭素。TimeMilli、カーボン。タイムマイクロ、カーボン。タイムナノ、カーボン。タイムスタンプ、カーボン。タイムスタンプミリ、カーボン。タイムスタンプマイクロ、カーボン。TimestampNano`フィールドタイプは、`MarshalJSON/UnmarshalJSON`でカスタム出力フォーマットを実装するために代わりに一般的なフィールドを使用します

旧バージョンの更新ログについては <a href="https://github.com/dromara/carbon/releases" target="_blank" rel="noreferrer">releases</a> を参照してください
