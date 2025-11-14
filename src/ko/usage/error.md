---
head:
  - - meta
    - name: description
      content: "Carbon 오류 처리 기능 상세 설명, HasError 오류 감지 및 Error 오류 조회 메서드 제공, 여러 오류는 첫 번째 오류만 반환, 시간 파싱 및 타임존 설정 등 일반적인 오류 시나리오 포함, 완전한 오류 처리 예제 포함"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, 오류 처리, HasError, Error, 오류 감지, 오류 조회, 시간 파싱 오류, 타임존 오류, 예외 처리"
---

# 오류 처리
여러 오류가 발생하면 첫 번째 오류만 반환됩니다. <a href="https://github.com/dromara/carbon/blob/master/errors.go" target="_blank" rel="noreferrer">errors.go</a>를 방문하여 발생할 수 있는 모든 오류를 확인하세요

```go
c1 := carbon.Parse("xxx")
if c1.HasError() {
  // 오류 처리...
  log.Fatal(c1.Error)
}
// 출력
failed to parse "xxx" as carbon

c2 := carbon.Parse("2020-08-05").SetTimezone("xxx")
if c2.HasError() {
  // 오류 처리...
  log.Fatal(c2.Error)
}
// 출력
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezones

c3 := carbon.Parse("xxx").SetTimezone("xxx")
if c3.HasError() {
  // 오류 처리...
  log.Fatal(c3.Error)
}
// 출력
failed to parse "xxx" as carbon
``` 