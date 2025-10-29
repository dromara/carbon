---
head:
  - - meta
    - name: description
      content: Error handling | Provide HasError/Error for error detection and retrieval; when multiple errors occur, only the first is returned; covers common scenarios like parsing and timezone setting
  - - meta
    - name: keywords
      content: carbon, go-carbon, error
---

# Error handling
If more than one error, only the first error is returned, refer to <a href="https://github.com/dromara/carbon/blob/master/errors.go" target="_blank" rel="noreferrer">errors.go</a> for all possible errors

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