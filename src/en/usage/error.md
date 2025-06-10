# Error handling
If more than one error, only the first error is returned

```go
c1 := carbon.Parse("xxx")
if c1.HasError() {
  // Error handle...
  log.Fatal(c1.Error)
}
// Output
failed to parse "xxx" as carbon

c2 := carbon.Parse("2020-08-05").SetTimezone("xxx")
if c2.HasError() {
  // Error handle...
  log.Fatal(c2.Error)
}
// Output
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezones

c3 := carbon.Parse("xxx").SetTimezone("xxx")
if c3.HasError() {
  // Error handle...
  log.Fatal(c3.Error)
}
// Output
failed to parse "xxx" as carbon
```