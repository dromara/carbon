---
head:
  - - meta
    - name: description
      content: 错误处理|一个轻量级、语义化、对开发者友好的 golang 时间处理库,
  - - meta
    - name: keywords
      content: 错误
---

# 错误处理
如果出现多个错误，则只返回第一个错误，访问 <a href="https://github.com/dromara/carbon/blob/master/errors.go" target="_blank" rel="noreferrer">errors.go</a> 查看所有可能出现的错误

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

