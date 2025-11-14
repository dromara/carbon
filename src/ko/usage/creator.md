---
head:
  - - meta
    - name: description
      content: "타임스탬프(초/밀리초/마이크로초/나노초)와 날짜시간/날짜/시간에서 인스턴스를 생성하는 메서드 제공: CreateFromTimestamp/DateTime/Date/Time 및 정밀도 변형"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, 타임스탬프에서 Carbon 인스턴스 생성, 날짜에서 Carbon 인스턴스 생성"
---

# `Carbon` 인스턴스 생성

## 타임스탬프에서 `Carbon` 인스턴스 생성
```go
// 초 정밀도 타임스탬프에서 Carbon 인스턴스 생성
carbon.CreateFromTimestamp(-1).ToString() // 1969-12-31 23:59:59 +0000 UTC
carbon.CreateFromTimestamp(0).ToString() // 1970-01-01 00:00:00 +0000 UTC
carbon.CreateFromTimestamp(1).ToString() // 1970-01-01 00:00:01 +0000 UTC
carbon.CreateFromTimestamp(1596633255).ToString() // 2020-08-05 13:14:15 +0000 UTC
// 밀리초 정밀도 타임스탬프에서 Carbon 인스턴스 생성
carbon.CreateFromTimestampMilli(1596604455999).ToString() // 2020-08-05 13:14:15.999 +0000 UTC
// 마이크로초 정밀도 타임스탬프에서 Carbon 인스턴스 생성
carbon.CreateFromTimestampMicro(1596604455999999).ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
// 나노초 정밀도 타임스탬프에서 Carbon 인스턴스 생성
carbon.CreateFromTimestampNano(1596604455999999999).ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC
```

## 날짜에서 `Carbon` 인스턴스 생성
```go
// 년, 월, 일, 시, 분, 초에서 Carbon 인스턴스 생성
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToString() // 2020-08-05 13:14:15 +0000 UTC
// 년, 월, 일, 시, 분, 초, 밀리초에서 Carbon 인스턴스 생성
carbon.CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString() // 2020-08-05 13:14:15.999 +0000 UTC
// 년, 월, 일, 시, 분, 초, 마이크로초에서 Carbon 인스턴스 생성
carbon.CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
// 년, 월, 일, 시, 분, 초, 나노초에서 Carbon 인스턴스 생성
carbon.CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC

// 년, 월, 일에서 Carbon 인스턴스 생성
carbon.CreateFromDate(2020, 8, 5).ToString() // 2020-08-05 00:00:00 +0000 UTC
// 년, 월, 일, 밀리초에서 Carbon 인스턴스 생성
carbon.CreateFromDateMilli(2020, 8, 5, 999).ToString() // 2020-08-05 00:00:00.999 +0000 UTC
// 년, 월, 일, 마이크로초에서 Carbon 인스턴스 생성
carbon.CreateFromDateMicro(2020, 8, 5, 999999).ToString() // 2020-08-05 00:00:00.999999 +0000 UTC
// 년, 월, 일, 나노초에서 Carbon 인스턴스 생성
carbon.CreateFromDateNano(2020, 8, 5, 999999999).ToString() // 2020-08-05 00:00:00.999999999 +0000 UTC

// 시, 분, 초에서 Carbon 인스턴스 생성(년월일은 현재 년월일로 기본값)
carbon.CreateFromTime(13, 14, 15).ToString() // 2020-08-05 13:14:15 +0000 UTC
// 시, 분, 초, 밀리초에서 Carbon 인스턴스 생성(년월일은 현재 년월일로 기본값)
carbon.CreateFromTimeMilli(13, 14, 15, 999).ToString() // 2020-08-05 13:14:15.999 +0000 UTC
// 시, 분, 초, 마이크로초에서 Carbon 인스턴스 생성(년월일은 현재 년월일로 기본값)
carbon.CreateFromTimeMicro(13, 14, 15, 999999).ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
// 시, 분, 초, 나노초에서 Carbon 인스턴스 생성(년월일은 현재 년월일로 기본값)
carbon.CreateFromTimeNano(13, 14, 15, 999999999).ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC
``` 