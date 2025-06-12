# エラー処理
複数のエラーが発生した場合、最初のエラーのみが返されます, アクセス <a href="https://github.com/dromara/carbon/blob/master/errors.go" target="_blank" rel="noreferrer">errors.go</a> 発生する可能性のあるすべてのエラーを表示

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
