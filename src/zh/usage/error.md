# 错误处理
如果出现多个错误，则只返回第一个错误

```go
c1 := carbon.Parse("xxx")
if c1.HasError() {
  // 错误处理...
  log.Fatal(c1.Error)
}
// 输出
failed to parse "xxx" as carbon

c2 := carbon.Parse("2020-08-05").SetTimezone("xxx")
if c2.HasError() {
  // 错误处理...
  log.Fatal(c2.Error)
}
// 输出
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezones

c3 := carbon.Parse("xxx").SetTimezone("xxx")
if c3.HasError() {
  // 错误处理...
  log.Fatal(c3.Error)
}
// 输出
failed to parse "xxx" as carbon
```
