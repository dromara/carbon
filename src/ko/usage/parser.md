---
head:
  - - meta
    - name: description
      content: 시간 파싱|가벼우면서도 의미론적이고 개발자 친화적인 golang 시간 처리 라이브러리
  - - meta
    - name: keywords
      content: carbon, go-carbon, 기본 템플릿, 레이아웃 템플릿, 형식 템플릿
---

# 시간 파싱
이 시리즈 메서드는 `타임스탬프` 문자열 파싱을 지원하지 않습니다. 타임스탬프를 파싱하려면 `CreateFromTimestamp` 또는 `CreateFromTimestampXXX` 등의 메서드를 사용하세요

## 기본 `레이아웃 템플릿`을 통한 파싱
기본 `레이아웃 템플릿`을 통해 시간 문자열을 `Carbon` 인스턴스로 파싱합니다
```go
carbon.Parse("").ToDateTimeString() // 빈 문자열
carbon.Parse("00:00:00").ToDateTimeString() // 빈 문자열
carbon.Parse("0000-00-00").ToDateTimeString() // 빈 문자열
carbon.Parse("0000-00-00 00:00:00").ToDateTimeString() // 빈 문자열

carbon.Parse("now").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("yesterday").ToString() // 2020-08-04 13:14:15 +0000 UTC
carbon.Parse("tomorrow").ToString() // 2020-08-06 13:14:15 +0000 UTC

carbon.Parse("2020").ToString() // 2020-01-01 00:00:00 +0000 UTC
carbon.Parse("2020-8").ToString() // 2020-08-01 00:00:00 +0000 UTC
carbon.Parse("2020-08").ToString() // 2020-08-01 00:00:00 +0000 UTC
carbon.Parse("2020-8-5").ToString() // 2020-08-05 00:00:00 +0000 UTC
carbon.Parse("2020-8-05").ToString() // 2020-08-05 00:00:00 +0000 UTC
carbon.Parse("2020-08-05").ToString() // 2020-08-05 00:00:00 +0000 UTC
carbon.Parse("2020-08-05.999").ToString() // 2020-08-05 00:00:00.999 +0000 UTC
carbon.Parse("2020-08-05.999999").ToString() // 2020-08-05 00:00:00.999999 +0000 UTC
carbon.Parse("2020-08-05.999999999").ToString() // 2020-08-05 00:00:00.999999999 +0000 UTC

carbon.Parse("2020-8-5 13:14:15").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-8-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-08-5 13:14:15").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-08-05 13:14:15.999").ToString() // 2020-08-05 13:14:15.999 +0000 UTC
carbon.Parse("2020-08-05 13:14:15.999999").ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
carbon.Parse("2020-08-05 13:14:15.999999999").ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC

carbon.Parse("2020-8-5T13:14:15+08:00").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-8-05T13:14:15+08:00").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-08-05T13:14:15+08:00").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("2020-08-05T13:14:15.999+08:00").ToString() // 2020-08-05 13:14:15.999 +0000 UTC
carbon.Parse("2020-08-05T13:14:15.999999+08:00").ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC

carbon.Parse("20200805").ToString() // 2020-08-05 00:00:00 +0000 UTC
carbon.Parse("20200805131415").ToString() // 2020-08-05 13:14:15 +0000 UTC
carbon.Parse("20200805131415.999").ToString() // 2020-08-05 13:14:15.999 +0000 UTC
carbon.Parse("20200805131415.999999").ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
carbon.Parse("20200805131415.999999999").ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC
carbon.Parse("20200805131415.999+08:00").ToString() // 2020-08-05 13:14:15.999 +0000 UTC
carbon.Parse("20200805131415.999999+08:00").ToString() // 2020-08-05 13:14:15.999999 +0000 UTC
carbon.Parse("20200805131415.999999999+08:00").ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC

carbon.Parse("2022-03-08T03:01:14-07:00").ToString() // 2022-03-08 18:01:14 +0000 UTC
carbon.Parse("2022-03-08T10:01:14Z").ToString() // 2022-03-08 18:01:14 +0000 UTC
```

## 하나의 `레이아웃 템플릿`을 통한 파싱
확정된 `레이아웃 템플릿`을 통해 시간 문자열을 `Carbon` 인스턴스로 파싱합니다
```go
carbon.ParseByLayout("2020|08|05 13|14|15", "2006|01|02 15|04|05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("오늘은 2020년08월05일13시14분15초", "오늘은 2006년01월02일15시04분05초").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("2020-08-05 13:14:15", "2006-01-02 15:04:05", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

## 하나의 `형식 템플릿`을 통한 파싱
확정된 `형식 템플릿`을 통해 시간 문자열을 `Carbon` 인스턴스로 파싱합니다
> 주의: 사용하는 문자가 형식 템플릿과 충돌할 때는 이스케이프 문자 "\\"를 사용하여 해당 문자를 이스케이프하세요
```go
carbon.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("오늘은 2020년08월05일13시14분15초", "오늘은 Y년m월d일H시i분s초").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("2020-08-05 13:14:15", "Y-m-d H:i:s", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

## 여러 `레이아웃 템플릿`을 통한 파싱
여러 모호한 `레이아웃 템플릿`을 통해 시간 문자열을 `Carbon` 인스턴스로 파싱합니다
```go
c := carbon.ParseByLayouts("2020|08|05 13|14|15", []string{"2006|01|02 15|04|05", "2006|1|2 3|4|5"})
c.ToDateTimeString() // 2020-08-05 13:14:15
c.CurrentLayout() // 2006|01|02 15|04|05
```

## 여러 `형식 템플릿`을 통한 파싱
여러 모호한 `형식 템플릿`을 통해 시간 문자열을 `Carbon` 인스턴스로 파싱합니다
```go
c := carbon.ParseByFormats("2020|08|05 13|14|15", []string{"Y|m|d H|i|s", "y|m|d h|i|s"})
c.ToDateTimeString() // 2020-08-05 13:14:15
c.CurrentLayout() // 2006|01|02 15|04|05
``` 