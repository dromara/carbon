---
head:
  - - meta
    - name: description
      content: "Carbon 変更履歴、歴史バージョンからのすべての更新内容を詳細に記録、新機能、パフォーマンス最適化、Bug修正、API変更、依存関係のアップグレード、互換性の説明を含む"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, 変更履歴, バージョン履歴, 変更記録, 機能更新, Bug修正, パフォーマンス最適化, API変更, リリースノート, バージョン更新"
---

# 変更履歴

## [v2.7.0](https://github.com/dromara/carbon/compare/v2.6.18...v2.7.0) (2026-10-19)

### ✨ 新機能

- **`AddWeekdays`、`AddWeekday`、`SubWeekdays`、`SubWeekday` メソッドを追加し、営業日計算時に週末日をスキップ可能に**：インスタンスに設定された週末日をスキップして日付を移動するため、営業日単位で期限を計算できます。4 つの名前は `AddDays` / `AddDay` / `SubDays` / `SubDay` と揃えており、週末日はインスタンスから読み取るため `SetWeekendDays` も引き続き機能します。7 日すべてが週末日に設定されている場合は、無限ループを防ぐため各営業日のスパンを 1 週間に制限します（by [@darakanoit](https://github.com/darakanoit) in [#366](https://github.com/dromara/carbon/pull/366)）
- **`IsSameWeek` メソッドを追加し、2 つの日付が同じ週に属するか判定**：`IsSame*` ファミリー（`IsSameCentury`、`IsSameDecade`、`IsSameYear`、`IsSameQuarter`、`IsSameWeek`、`IsSameMonth`、`IsSameDay`、`IsSameHour`、`IsSameMinute`、`IsSameSecond`）を補完します。`SetWeekStartsAt` を尊重するため、ある設定で同じ週となる 2 つの日付が、別の設定では同じ週にならない場合があります。比較は週の開始日に基づき、年をまたぐ境界週も正しく処理します（by [@darakanoit](https://github.com/darakanoit) in [#368](https://github.com/dromara/carbon/pull/368)）
- **`IsBirthday` メソッドを追加し、月日を比較して誕生日か判定**：年を無視して同じ月日かどうかを返します。引数を省略した場合は現在時刻と比較します。2 月 29 日は 2 月 29 日のみに一致するため平年では決して誕生日にならず、PHP Carbon の挙動と一致します（by [@darakanoit](https://github.com/darakanoit) in [#364](https://github.com/dromara/carbon/pull/364)）
- **`Asia/Jakarta` タイムゾーン定数のサポートを追加**（by [@mirzaq19](https://github.com/mirzaq19) in [#374](https://github.com/dromara/carbon/pull/374)）

### 🐛 修正

- **夏時間の巻き戻し時に `StartOfSecond` と `EndOfSecond` が誤った時刻の出現を選択する問題を修正**：夏時間で繰り返される時刻において、`time.Date` による再構築が同じローカル時刻のもう一方の出現を選択する可能性がありました。現在はインスタンスをコピーして既存タイムスタンプのサブ秒精度を調整し、選択された瞬間・ロケーション・設定を保持します（by [@mikamikasuki](https://github.com/mikamikasuki) in [#377](https://github.com/dromara/carbon/pull/377)）
- **夏時間の巻き戻し時に `StartOfMinute` と `EndOfMinute` が誤った時刻の出現を選択する問題を修正**：#377 と同種のバグで、1 段階上の単位での問題です。現在はインスタンスをコピーしてローカルの秒数を減算し、瞬間・ロケーション・オフセットを保持します（by [@dualfroz](https://github.com/dualfroz) in [#378](https://github.com/dromara/carbon/pull/378)）
- **夏時間の巻き戻し時に `StartOfHour` と `EndOfHour` が誤った時刻の出現を選択する問題を修正**：再構築された時刻境界にレシーバーのオフセットと一致する有効な対応出現がある場合はレシーバーの出現を保持します。対応する出現が存在しない場合（Lord Howe の部分的時間移行など）は既存のカレンダー再構築動作を維持します（by [@dajiaohuang](https://github.com/dajiaohuang) in [#382](https://github.com/dromara/carbon/pull/382)）
- **`DiffInSeconds` がサブ秒部分を無視し、秒の境界をまたぐと不正確にカウントする問題を修正**：整数秒のタイムスタンプを減算する際に小数部分が無視されるため、秒の境界をまたぐ 0.2 秒の差が 1 秒とカウントされていました。現在はナノ秒差の符号が逆の場合に ±1 調整し、経過した整数秒をゼロ方向へ切り捨ててカウントします（by [@youdie006](https://github.com/youdie006) in [#379](https://github.com/dromara/carbon/pull/379)）
- **`DiffInDays` と `DiffInWeeks` が夏時間の切り替えをまたぐと誤った結果を計算する問題を修正**：経過秒数を除算する方式では、夏時間の変更をまたぐと誤った結果になっていました。現在は `DiffInMonths` が月を数える方法（`AddDate` と調整）で暦日を数え、`DiffInWeeks` はそれを使用します。夏時間の切り替えをまたぐ区間のみ結果が変わります（by [@youdie006](https://github.com/youdie006) in [#376](https://github.com/dromara/carbon/pull/376)）
- **`DiffInYears` と `DiffInMonths` が異なるタイムゾーン間の比較で誤った推定を行う問題を修正**：それぞれの日付の `Year()` / `Month()` フィールドからの初回推定は、2 つの日付が異なるロケーションにある場合にずれる可能性がありました。現在は先に早い日付のロケーションで遅い日付を読み取ってから年と月のフィールドを取得し、`DiffInDays` と同じ枠組みに合わせています（by [@ybai08](https://github.com/ybai08) in [#384](https://github.com/dromara/carbon/pull/384)）
- **`IsSameDecade` が世紀を比較しておらず、世紀をまたぐ日付を誤判定する問題を修正**：`Decade()` は世紀内の十年であるため、`Parse("1995-08-05").IsSameDecade(Parse("2095-08-05"))` が true を返していました。現在は `Century()` も比較し、`StartOfDecade` が使用する区分と一致させています（by [@youdie006](https://github.com/youdie006) in [#375](https://github.com/dromara/carbon/pull/375)）
- **`FromStdTime` が UTC オフセットの相違により深夜直後に誤った旧暦日付を返す問題を修正**：時刻の属するタイムゾーンで 1900-01-31 からの経過時間を計算していたため、2 つの瞬間のオフセットが日数計算に漏れていました。現在は UTC で壁時計の日付間の日数を計算し、`persian.FromStdTime` と `hebrew.FromStdTime` と一致させています（by [@youdie006](https://github.com/youdie006) in [#380](https://github.com/dromara/carbon/pull/380)）
- **凍結クロックモードで `Sleep` を並行呼び出しするとクロックの進みが失われる競合状態を修正**：凍結クロック上の並行 `Sleep` 呼び出しが相互に、また `Now` と競合してクロックの進みが失われていました。現在は既存の凍結クロック用ミューテックスでチェックと更新を囲んでいます（by [@yuefdev](https://github.com/yuefdev) in [#383](https://github.com/dromara/carbon/pull/383)）

### 🔧 依存関係の更新

- `github/codeql-action` を `v4.38.1` から `v4.38.3` にアップグレード（by [@dependabot\[bot\]](https://github.com/dependabot) in [#372](https://github.com/dromara/carbon/pull/372) [#381](https://github.com/dromara/carbon/pull/381)）

## [v2.6.18](https://github.com/dromara/carbon/compare/v2.6.17...v2.6.18) (2026-09-28)

### ✨ 新機能

- **`カザフ語` のローカライズ言語サポートを追加（`lang/kk.json`）**：月名、短い月名、曜日名、短い曜日名は CLDR に準拠した PHP Carbon のカザフ語ロケールからそのまま取得しています。相対時間の表現は、数詞の後の名詞が数変化しないカザフ語の文法に従います。季節と星座は伝統的なカザフ語名で直接翻訳しています（by [@darakanoit](https://github.com/darakanoit) in [#355](https://github.com/dromara/carbon/pull/355)）

### 🐛 修正

- **`Format` メソッドで末尾のエスケープ文字を処理する際に範囲外アクセスが発生する可能性がある問題を修正**：`helper.go` の `format2layout` メソッドでエスケープ文字処理に境界チェックを追加し、範囲外アクセスによる `panic` を防止します（by [@gouguoyin](https://github.com/gouguoyin) in [#348](https://github.com/dromara/carbon/pull/348)）
- **週の開始日を月曜日以外に設定した場合、`ToWeekString`、`ToShortWeekString`、`Format` が誤った曜日名を返す問題を修正**：`ToWeekString` / `ToShortWeekString` は `DayOfWeek() % DaysPerWeek` で名称リソースを参照していましたが、リソースは絶対順（日曜日が先頭）で並んでいました。現在は絶対曜日 `int(StdTime().Weekday())` で参照し、リソースの順序と一致させ、`Format("l")` の動作に合わせています（by [@sneha4175](https://github.com/sneha4175) in [#371](https://github.com/dromara/carbon/pull/371)）
- **既存の時刻値を持つデータベースフィールドに `NULL` をスキャンした際、既存値がクリアされない問題を修正**：既に値が入っている Carbon に SQL の `NULL` をスキャンすると、成功を返しながら以前の日付が残っていました。現在は `Carbon.Scan` およびレイアウト・フォーマット・タイムスタンプの各スキャナラッパーで出力先をリセットします（by [@vitalivo](https://github.com/vitalivo) in [#360](https://github.com/dromara/carbon/pull/360)）

### ⚡ パフォーマンス

- **コピーオンライトにより言語インスタンスとリソースのキャッシュおよび共有を最適化し、メモリ割り当てを削減**：`SetLocale` はキャッシュされたロケールマップ全体をすべての `Language` インスタンスにコピーしていたため、Carbon インスタンスごとに 20 エントリのマップを新規割り当てしていました。現在は `SetLocale` がキャッシュマップを共有して共有済みとマークし、`SetResources` は初回書き込み前にコピーし、リソースを読むだけの処理は共有マップを読み続けます。`NewCarbon()` あたりの割り当ては 1600 B / 13 回から 232 B / 4 回に減少します（by [@darakanoit](https://github.com/darakanoit) in [#350](https://github.com/dromara/carbon/pull/350)）
- **月、曜日、季節、星座のローカライズリソース検索を最適化し、呼び出しごとの文字列分割を回避**：6 つのメソッドは呼び出しごとに「|」区切りの言語リソースを読み、1 要素を参照するためだけに `strings.Split` を実行していました。現在は小さなヘルパーが `strings.Count` と `strings.IndexByte` でその場で要素を特定し、呼び出しあたりの割り当てはゼロです。`Constellation` -61.54%、`ToShortWeekString` -47.35%、`ToWeekString` -44.13%、`Season` -40.65%、`ToMonthString` -37.66%、`ToShortMonthString` -35.88%（by [@darakanoit](https://github.com/darakanoit) in [#352](https://github.com/dromara/carbon/pull/352)）
- **ロケールごとに 1 つの言語インスタンスを共有**：`NewCarbon()` はインスタンスごとに新しい `Language` を構築していました。現在は `defaultLanguage(locale)` がロケールごとに即使用可能な `Language` を 1 つ `sync.Map` にキャッシュします。Carbon の構築時の割り当ては 4 回から 1 回に減少します（by [@darakanoit](https://github.com/darakanoit) in [#354](https://github.com/dromara/carbon/pull/354)）
- **年間日数と累積日数のテーブルを事前計算して旧暦の日付変換を最適化**：旧暦変換は年テーブルをインデックス参照せず走査していたため、O(1900 年からの年数) のコストがかかっていました。現在は対応する各年の日数と累計をパッケージ初期化時に一度だけ算出します。2020 年の日付から旧暦への変換は 1009 ns -> 282 ns（-72%）、グレゴリオ暦への逆変換は 738 ns -> 23 ns（-97%）になります（by [@darakanoit](https://github.com/darakanoit) in [#357](https://github.com/dromara/carbon/pull/357)）

### 🔧 依存関係の更新

- `github.com/stretchr/testify` を `v1.11.1` から `v1.12.1` にアップグレード（by [@dependabot\[bot\]](https://github.com/dependabot) in [#341](https://github.com/dromara/carbon/pull/341) [#344](https://github.com/dromara/carbon/pull/344)）
- `github/codeql-action` を `v4.37.4` から `v4.38.1` にアップグレード（by [@dependabot\[bot\]](https://github.com/dependabot) in [#340](https://github.com/dromara/carbon/pull/340) [#342](https://github.com/dromara/carbon/pull/342) [#345](https://github.com/dromara/carbon/pull/345) [#346](https://github.com/dromara/carbon/pull/346) [#362](https://github.com/dromara/carbon/pull/362) [#369](https://github.com/dromara/carbon/pull/369)）

### 🚀 CI

- **`CI` テストワークフローを改善し、`PR` で完全なテストスイートと競合検出を実行し、`push` 時のみカバレッジを生成およびアップロード**：テストコマンドはルートパッケージのみではなく `./...` で全パッケージを対象とするようになりました。カバレッジファイルの生成と Codecov へのアップロードは `push` イベントに限定し、全パッケージで `-race` を使用する専用テストステップを追加しました（by [@gouguoyin](https://github.com/gouguoyin) in [#361](https://github.com/dromara/carbon/pull/361)）

## [v2.6.17](https://github.com/dromara/carbon/compare/v2.6.16...v2.6.17) (2026-08-09)

### 🐛 修正

- `イスラム` の曜日名のスペルミスを修正し、`Jumaat` を `Jumat` に変更
- `グレゴリオ暦` と `ヘブライ暦` の間の変換における計算エラーを修正

### 🔧 依存関係の更新

- `github/codeql-action` を `v4` から `v4.37.4` にアップグレード
- `actions/setup-go` を `v6` から `v7` にアップグレード
- `actions/checkout` を `v6` から `v7` にアップグレード
- `codecov/codecov-action` を `v5` から `v6` にアップグレード
- `codecov/codecov-action` を `v6` から `v7` にアップグレード

### 📝 ドキュメント

- `README` ドキュメントを `NRG` 標準テンプレートに統一
- `README` 内の冗長な広告外部リンクを整理し、廃止された `Go Report Card` ステータスバッジを削除

## [v2.6.16](https://github.com/dromara/carbon/compare/v2.6.15...v2.6.16) (2026-01-28)

### 🐛 修正

- `Format` メソッドで `u`、`v`、`x` 記号の解析が誤っていた bug を修正

### 🔧 依存関係の更新

- `actions/checkout` を `v5` から `v6` にアップグレード

## [v2.6.15](https://github.com/dromara/carbon/compare/v2.6.14...v2.6.15) (2025-11-15)

### ✨ 新機能

- `HttpLayout`、`HttpFormat` 定数および `ToHttpString` メソッドを追加
- `モンゴル語` のローカライズ言語サポートを追加（lang/mn.json）
- `アフリカーンス語` のローカライズ言語サポートを追加（lang/af.json）

### 🐛 修正

- `difference.go` ファイルの `DiffAbsInDuration` メソッドで使用されている `go1.19` 以降でのみ利用可能な `Abs` メソッドを削除
- `オランダ語` ローカライズ言語ファイル（lang/nl.json）の `weeks` フィールドで重複している `Zondag` エラーを修正

## [v2.6.14](https://github.com/dromara/carbon/compare/v2.6.13...v2.6.14) (2025-10-28)

### ✨ 新機能

- `ギリシャ語` のローカライズ言語サポートを追加（lang/el.json）
- `フィンランド語` のローカライズ言語サポートを追加（lang/fi.json）
- `ミャンマー語` のローカライズ言語サポートを追加（lang/my.json）

### 🐛 修正

- `traveler.go` ファイルの時間増減シリーズメソッドを最適化し、インスタンスをコピーすることで不変性を保証し、元のインスタンスの変更を防ぎます

## [v2.6.13](https://github.com/dromara/carbon/compare/v2.6.12...v2.6.13) (2025-10-15)

### 🐛 修正

- `helper.go` の `format2layout` メソッドでエスケープ文字処理の境界チェックを追加し、範囲外アクセスによる `panic` を防止

### ⚡ パフォーマンス

- `language.go` の `SetLocale` メソッドで `sync.Once` を使用して言語ファイルが一度だけ読み込まれることを保証し、`sync.Map` を使用してスレッドセーフなキャッシュを実装

## [v2.6.12](https://github.com/dromara/carbon/compare/v2.6.11...v2.6.12) (2025-09-16)

### 🐛 修正

- `type_carbon.go` ファイルの `UnmarshalJSON` メソッドで `isEmpty` フラグを設定して空値を示す
- 潜在的な競合状態と null ポインタ逆参照の問題を修正し、並行安全性を向上

### ⚡ パフォーマンス

- `sync.Map` を使用して高性能な並行キャッシュを実装

### 🔧 依存関係の更新

- `golang` 環境依存を `1.21` から `1.18` に下げる
- `testify` テストフレームワークを `v1.10.0` から `v1.11.1` にアップグレード

## [v2.6.11](https://github.com/dromara/carbon/compare/v2.6.10...v2.6.11) (2025-07-18)

### ✨ 新機能

- `ヘブライ暦` サポートを追加し、ユニットテストとベンチマークテストファイルを追加

### ♻️ リファクタリングと変更

- `Sleep` を構造体メソッドからグローバルメソッドに変更
- `ペルシア暦` をリファクタリングし、ベンチマークテストファイルを追加
- `中国旧暦` をリファクタリングし、ベンチマークテストファイルを追加
- `ユリウス日/修正ユリウス日` をリファクタリングし、ベンチマークテストファイルを追加

### 📝 ドキュメント

- 全体のパフォーマンステストレポートファイルを追加

## [v2.6.10](https://github.com/dromara/carbon/compare/v2.6.9...v2.6.10) (2025-07-07)

### ✨ 新機能

- `Sleep` メソッドと関連する`ユニットテスト`、`ベンチマークテスト`、`サンプルファイル`を追加
- `MaxYear`、`MinYear`、`MaxMonth`、`MinMonth`、`MaxDay`、`MinDay` などの数値定数を追加し、これらの定数でハードコードを置き換え

### ⚡ パフォーマンス

- `helper.go` の `getAbsValue` メソッドを最適化し、条件判断をビット演算に置き換え
- `frozen.go` ファイルの時間凍結関連メソッドを最適化し、原子操作を使用してロック競合を減らし、メモリ割り当てを最適化
- ベンチマークテストファイルを最適化し、`シリアルテスト`、`パラレルテスト`、`コンカレントテスト`をカバー

### ♻️ リファクタリングと変更

- `日本語` 翻訳ファイルを `jp.json` から `ja.json` に変更し、ドキュメントを `README.jp.md` から `README.ja.md` に名前変更して [ISO639-1](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes) 標準に準拠
- 非推奨の `ParseWithLayouts` メソッドを削除し、`ParseByLayouts` メソッドに置き換え
- 非推奨の `ParseWithFormats` メソッドを削除し、`ParseByFormats` メソッドに置き換え
- 非推奨の `CleanTestNow` メソッドを削除し、`ClearTestNow` メソッドに置き換え
- `ParseByLayout` と `ParseByFormat` メソッドからタイムスタンプ文字列の解析サポートを削除、タイムスタンプの解析には `CreateFromTimestamp`、`CreateFromTimestampMilli`、`CreateFromTimestampMicro`、`CreateFromTimestampNano` メソッドを使用

### 📝 ドキュメント

- 韓国語ドキュメント `README.ko.md` を追加

## [v2.6.9](https://github.com/dromara/carbon/compare/v2.6.8...v2.6.9) (2025-06-28)

### ♻️ リファクタリングと変更

- `gorm` の `GormDataType` インタフェースの実装の削除

## [v2.6.8](https://github.com/dromara/carbon/compare/v2.6.7...v2.6.8) (2025-06-12)

### 🐛 修正

- 解析時にレイアウトテンプレートまたはフォーマットテンプレートが空の場合は、エラーを返します

### ♻️ リファクタリングと変更

- `type_builtin.go` では `DateTimeType` を `dateTimeType` に、`DateTimeXXXType` を `dateTimeXXXType` に、それぞれ名前を変更します
- `type_builtin.go` では `DateType` を `dateType` に、`DateXXXType` を `dateXXXType` に、それぞれ名前を変更します
- `type_builtin.go` では `TimeType` を `timeType` に、`TimeXXXType` を `timeXXXType` に、それぞれ名前を変更します

### 🔧 依存関係の更新

- `tests` で `gorm.io/gorm` を `1.21.1` から `1.30.0` にアップグレード
- `tests` で `gorm.io/driver/mysql` を `1.5.7` から `1.6.0` にアップグレード
- `tests` で `gorm.io/driver/postgres` を `1.5.7` から `1.6.0` にアップグレード
- `tests` で `gorm.io/driver/sqlite` を `1.5.7` から `1.6.0` にアップグレード

### 📝 ドキュメント

- `README` ファイルをシンプル化し、詳細な使用方法と使用例を[公式サイト](https://carbon.go-pkg.com)に移行
- [HelloGitHub](https://hellogithub.com/repository/dromara/carbon) バッジリンクを追加

## [v2.6.7](https://github.com/dromara/carbon/compare/v2.6.6...v2.6.7) (2025-05-26)

### ✨ 新機能

- `Closest`/ `Farthest` メソッドの 2 番目のパラメータをオプションパラメータに変更
- `ZeroValue`/`EpochValue` メソッドを追加
- `DataTyper` インターフェースを追加し、内蔵型が `DataTyper` インターフェースを実装するようにします

### ♻️ リファクタリングと変更

- `String` メソッドから null 値 carbon のチェックを削除
- ファイル名を `type_interface.go` から `interfaces.go` に変更

## [v2.6.6](https://github.com/dromara/carbon/compare/v2.6.5...v2.6.6) (2025-05-19)

### ✨ 新機能

- `xorm` の `curd` 統合テストを追加し、現在 `MySQL`/`Postgres`/`SQLite` をカバー

### 🐛 修正

- `Windows` OS で言語ファイルが見つからないバグを修正
- 新しい `Carbon` インスタンスを作成すると `layout`、`weekStartsAt`、`weekendDays` および `lang` の値が失われるエラーを修正
- `StartOfWeek` と `EndOfWeek` メソッドが元の `Carbon` インスタンスを予期せず変更するバグを修正

### 🚀 CI

- `ci` アクションで `Windows` OS のユニットテストを追加

## [v2.6.5](https://github.com/dromara/carbon/compare/v2.6.4...v2.6.5) (2025-05-14)

### ✨ 新機能

- `Carbon` 構造体の `SetLanguage` メソッドに無効な `Language` 構造体のチェックを追加します
- `Carbon` 構造体の `Parse` メソッドに `Postgres` / `SQLite` タイムフォーマット文字列のサポートを追加します
- `Parse`/`ParseByLayout`/`ParseByFormat` メソッドは `空` 文字列を解析すると戻り値を `nil` から空の carbon に変更します
- `Carbon` 構造体に `IsEmpty` メソッドを追加して空の carbon であるかを判断します
- `Carbon` 構造体に `CleanTestNow` を置き換える `ClearTestNow` メソッドを追加します。`CleanTestNow` は将来削除されます
- `Carbon` 構造体に `ParseWithLayouts` を置き換える `ParseByLayouts` メソッドを追加します。`ParseWithLayouts` は将来削除されます
- `Carbon` 構造体に `ParseWithFormats` を置き換える `ParseByFormats` メソッドを追加します。`ParseWithFormats` は将来削除されます
- `Language` 構造体の `SetResources` メソッドに無効なリソースのチェックを追加します
- `gorm` の `curd` 統合テストを追加し、現在 `MySQL`/`Postgres`/`SQLite` をカバー

### ♻️ リファクタリングと変更

- `Carbon` 構造体から `GormDataType` メソッドを削除し、`Value` / `MarshalJSON` メソッドを `ポインタ` 受信機から `値` 受信機に変更します
- `LayoutType[T]` 構造体から `GormDataType` メソッドを削除し、`Value` / `MarshalJSON` メソッドを `ポインタ` 受信機から `値` 受信機に変更します
- `FormatType[T]` 構造体から `GormDataType` メソッドを削除し、`Value` / `MarshalJSON` メソッドを `ポインタ` 受信機から `値` 受信機に変更します
- `TimestampType[T]` 構造体から `GormDataType` メソッドを削除し、`Value` / `MarshalJSON` メソッドを `ポインタ` 受信機から `値` 受信機に変更します
- ユニットテストのために `github.com/stretchr/testify/assert` を `github.com/stretchr/testify/suite` に置き換えます

### 🔧 依存関係の更新

- `go` バージョンを `1.18` から `1.21` にアップグレード

## [v2.6.4](https://github.com/dromara/carbon/compare/v2.6.3...v2.6.4) (2025-04-28)

### ✨ 新機能

- `Language` 構造体に `Copy` メソッドを追加
- `carbon.Timestamp` 型エイリアスと `carbon.NewTimestamp` メソッドを追加
- `carbon.TimestampMilli` 型エイリアスと `carbon.NewTimestampMilli` メソッドを追加
- `carbon.TimestampMicro` 型エイリアスと `carbon.NewTimestampMicro` メソッドを追加
- `carbon.TimestampNano` 型エイリアスと `carbon.NewTimestampNano` メソッドを追加
- `carbon.DateTime` 型エイリアスと `carbon.NewDateTime` メソッドを追加
- `carbon.DateTimeMilli` 型エイリアスと `carbon.NewDateTimeMilli` メソッドを追加
- `carbon.DateTimeMicro` 型エイリアスと `carbon.NewDateTimeMicro` メソッドを追加
- `carbon.DateTimeNano` 型エイリアスと `carbon.NewDateTimeNano` メソッドを追加
- `carbon.Date` 型エイリアスと `carbon.NewDate` メソッドを追加
- `carbon.DateMilli` 型エイリアスと `carbon.NewDateMilli` メソッドを追加
- `carbon.DateMicro` 型エイリアスと `carbon.NewDateMicro` メソッドを追加
- `carbon.DateNano` 型エイリアスと `carbon.NewDateNano` メソッドを追加
- `carbon.Time` 型エイリアスと `carbon.NewTime` メソッドを追加
- `carbon.TimeMilli` 型エイリアスと `carbon.NewTimeMilli` メソッドを追加
- `carbon.TimeMicro` 型エイリアスと `carbon.NewTimeMicro` メソッドを追加
- `carbon.TimeNano` 型エイリアスと `carbon.NewTimeNano` メソッドを追加

### 🐛 修正

- データベースフィールドタイプが `nil` だったときのパニックを修正

### ⚡ パフォーマンス

- ベンチマークテストファイルに `b.ResetTimer()` を追加

### ♻️ リファクタリングと変更

- `database_types.go` を `type_carbon.go`、`type_layout.go`、`type_format.go`、`type_timestamp.go` に分割
- `LayoutFactory` インターフェイスを `LayoutTyper` に、`SetLayout` メソッドを `Layout` に変更
- `FormatFactory` インターフェイスを `FormatTyper` に、`SetFormat` メソッドを `Format` に変更
- `TimestampFactory` インターフェイスを `TimestampTyper` に、`SetPrecision` メソッドを `Precision` に変更

## [v2.6.3](https://github.com/dromara/carbon/compare/v2.6.2...v2.6.3) (2025-04-21)

### ✨ 新機能

- UNIX エポック時間（1970-01-01 00:00:00 +0000 UTC）であるかを判定する `IsEpoch` メソッドを追加
- 週の終了日を取得する `WeekEndsAt` メソッドを追加
- 週の週末日を設定する `SetWeekendDays` メソッドを追加
- デフォルトの週末日を保持する `DefaultWeekStartsAt` グローバル変数を追加

### 🐛 修正

- `IsWeekend` と `IsWeekday` メソッドが異なる国で同じ結果を返すバグを修正
- `carbon.Parse("").StdTime()` が nil ポインタパニックを引き起こすバグを修正 #294

### ⚡ パフォーマンス

- ベンチマークテストファイル `xxx_bench_test.go` を追加

### ♻️ リファクタリングと変更

- エラーメソッドを `private` メソッドから `public` メソッドに変更
- 週のデフォルトの開始日を `日曜日` から `月曜日` に変更
- `MinValue` の年を `-9998` から `1` に変更
- 定数 `weeksPerLongYear` を `WeeksPerLongYear` に変更

## [v2.6.2](https://github.com/dromara/carbon/compare/v2.6.1...v2.6.2) (2025-04-08)

### ✨ 新機能

- タイムゾーンオフセットを取得するフォーマットシンボル `o` を追加
- `TimestampLayout`、`TimestampMilliLayout`、`TimestampMicroLayout`、`TimestampNanoLayout` 定数を追加
- `TimestampFormat`、`TimestampMilliFormat`、`TimestampMicroFormat`、`TimestampNanoFormat` 定数を追加
- `DateTimeMilli`、`DateTimeMicro`、`DateTimeNano` フィールドタイプを追加
- `DateMilli`、`DateMicro`、`DateNano` フィールドタイプを追加
- `TimeMilli`、`TimeMicro`、`TimeNano` フィールドタイプを追加
- `MaxDuration`、`MinDuration` メソッドを追加

### 🐛 修正

- 旧暦の `IsLeapMonth` の不正判断のバグを修正
- `AtomFormat` と `AtomLayout` が一致しない値を返すバグを修正
- `RFC3339Format` と `RFC3339Layout` が一致しない値を返すバグを修正
- `IsDST` メソッドのタイムゾーンが失われるバグを修正
- `StartOfXXX`、`EndOfXXX` の一部メソッドでタイムゾーンが失われるバグを修正
- 他のカレンダーを `グレゴリオ暦` に変換するときにタイムゾーンが失われるバグを修正

### ♻️ リファクタリングと変更

- `CreateFromLunar`、`CreateFromPersian` メソッドから `hour`、`minute`、`second` パラメータを削除
- 一部のフォーマットシンボル定義を変更、対象となるシンボルは `U`, `V`, `X`, `S`, `T`, `Z`, `u`, `v`, `x`, `z`
- グローバルデフォルトタイムゾーンを設定する際に `time.Local` を同期更新しないように変更

## [v2.6.1](https://github.com/dromara/carbon/compare/v2.6.0...v2.6.1) (2025-03-27)

### ✨ 新機能

- `ParseWithLayouts` と `ParseWithFormats` メソッドを追加

### 🐛 修正

- `gorm` がデータを更新する際に `updated_at` フィールドの自動更新が無効になるバグを修正

### ♻️ リファクタリングと変更

- `formatFactory` インターフェイスを `FormatFactory` に、`layoutFactory` インターフェイスを `LayoutFactory` に、`timestampFactory` インターフェイスを `TimestampFactory` に名前変更し、タイプ制約を追加
- `LayoutType`、`FormatType`、`TimestampType` 構造体の `GormDataType` メソッドの返り値を `time` に変更
- `DateTime`、`Date`、`Time` タイプを `struct` から `string` に変更
- `Timestamp`、`TimestampMilli`、`TimestampMicro`、`TimestampNano` タイプを `struct` から `int64` に変更
- 内蔵データベースフィールドタイプを新しいファイル `types.go` に移動

## [v2.6.0](https://github.com/dromara/carbon/compare/v2.5.4...v2.6.0) (2025-03-25)

### ✨ 新機能

- タイムゾーン名を取得する `ZoneName` メソッドを追加
- エラーがあるかどうかを確認する `HasError` メソッドを追加
- `nil` であるかどうかを確認する `IsNil` メソッドを追加
- `carbon` を深くコピーする `Copy` メソッドを追加
- 週の開始日を取得する `WeekStartsAt` メソッドを追加
- サンプルファイル `xxx_example.go` を追加
- 新しい `constant.go` ファイルを追加し、`carbon.go` ファイルから定数をこのファイルに移行

### ♻️ リファクタリングと変更

- `golang` 最低バージョン依存性を `1.18` にアップグレード
- `carbon`、`julian`、`lunar`、`persian` を値渡しからポインタ渡しに変更
- デフォルトのグローバルタイムゾーンを `Local` から `UTC` に変更
- `Offset` メソッドを `ZoneOffset` メソッドに変更
- `IsSetTestNow` メソッドを `IsTestNow` メソッドに変更
- `UnSetTestNow` メソッドを `CleanTestNow` メソッドに変更
- `Location` メソッドを削除し、`Timezone` メソッドに置き換え
- `IsValid` と `IsInvalid` メソッドの判断論理を変更し、`ゼロタイム` はもはや無効な時刻とみなされません
- グローバルデフォルトタイムゾーンを設定する際に `time.Local` を同期更新
- `database.go` をリファクタリングし、`carbon.DateTime`、`carbon.DateTimeMilli`、`carbon.DateTimeMicro`、`carbon.DateTimeNano`、`carbon.Date`、`carbon.DateMilli`、`carbon.DateMicro`、`carbon.DateNano`、`carbon.Time`、`carbon.TimeMilli`、`carbon.TimeMicro`、`carbon.TimeNano`、`carbon.Timestamp`、`carbon.TimestampMilli`、`carbon.TimestampMicro`、`carbon.TimestampNano` フィールドタイプを削除し、`MarshalJSON/UnmarshalJSON` 時のカスタム出力フォーマットを実装するために汎用フィールドで置き換え

旧バージョンの更新ログについては <a href="https://github.com/dromara/carbon/releases" target="_blank" rel="noreferrer">releases</a> を参照してください
