---
head:
  - - meta
    - name: description
      content: "Carbon エラー処理機能の詳細説明、HasError によるエラー検出と Error によるエラー取得メソッドを提供、複数のエラーは最初のエラーのみを返す、時間解析やタイムゾーン設定などの一般的なエラーシナリオをカバー、完全なエラー処理例を含む"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, エラー処理, HasError, Error, エラー検出, エラー取得, 時間解析エラー, タイムゾーンエラー, 例外処理"
---

# エラー処理
複数のエラーが発生した場合、最初のエラーのみが返されます。<a href="https://github.com/dromara/carbon/blob/master/errors.go" target="_blank" rel="noreferrer">errors.go</a> を参照して、発生する可能性のあるすべてのエラーを確認してください

```go
c1 := carbon.Parse("xxx")
if c1.HasError() {
  // エラー処理...
  log.Fatal(c1.Error)
}
// 出力
failed to parse "xxx" as carbon

c2 := carbon.Parse("2020-08-05").SetTimezone("xxx")
if c2.HasError() {
  // エラー処理...
  log.Fatal(c2.Error)
}
// 出力
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezones

c3 := carbon.Parse("xxx").SetTimezone("xxx")
if c3.HasError() {
  // エラー処理...
  log.Fatal(c3.Error)
}
// 出力
failed to parse "xxx" as carbon
```
