---
head:
  - - meta
    - name: description
      content: 달력|가벼우면서도 의미론적이고 개발자 친화적인 golang 시간 처리 라이브러리
  - - meta
    - name: keywords
      content: carbon, go-carbon, 달력, 공력, 그레고리력, 율리우스일, 수정율리우스일, 중국음력, 페르시아력, 이란력, 히브리력, 유대력
---

# 달력

현재 지원되는 달력은

- [율리우스일/수정율리우스일](#율리우스일-수정율리우스일)
- [중국음력](#중국음력)
- [페르시아력(이란력)](#페르시아력-이란력)
- [히브리력(유대력)](#히브리력-유대력)


## 율리우스일/수정율리우스일

### `공력`을 `율리우스일`로 변환
```go
// 기본적으로 6자리 소수점 정밀도 유지
carbon.Parse("2024-01-24 12:00:00").Julian().JD() // 2460334
carbon.Parse("2024-01-24 13:14:15").Julian().JD() // 2460334.051563

// 4자리 소수점 정밀도 유지
carbon.Parse("2024-01-24 12:00:00").Julian().JD(4) // 2460334
carbon.Parse("2024-01-24 13:14:15").Julian().JD(4) // 2460334.0516
```

### `공력`을 `수정율리우스일`로 변환
```go
// 기본적으로 6자리 소수점 정밀도 유지
carbon.Parse("2024-01-24 12:00:00").Julian().MJD() // 60333.5
carbon.Parse("2024-01-24 13:14:15").Julian().MJD() // 60333.551563

// 4자리 소수점 정밀도 유지
carbon.Parse("2024-01-24 12:00:00").Julian().MJD(4) // 60333.5
carbon.Parse("2024-01-24 13:14:15").Julian().MJD(4) // 60333.5516
```

### `율리우스일`을 `수정율리우스일`로 변환
```go
// 기본적으로 6자리 소수점 정밀도 유지
carbon.CreateFromJulian(2460334).Julian().MJD() // 60333.5
carbon.CreateFromJulian(2460334.051563).Julian().MJD() // 60332.551563

// 4자리 소수점 정밀도 유지
carbon.CreateFromJulian(2460334).Julian().MJD(4) // 60333.5
carbon.CreateFromJulian(2460334.051563).Julian().MJD(4) // 60332.5516
```

### `수정율리우스일`을 `율리우스일`로 변환
```go
// 기본적으로 6자리 소수점 정밀도 유지
carbon.CreateFromJulian(60333.5).Julian().JD() // 2460334
carbon.CreateFromJulian(60333.551563).Julian().JD() // 2460333.051563

// 4자리 소수점 정밀도 유지
carbon.CreateFromJulian(60333.5).Julian().JD(4) // 2460334
carbon.CreateFromJulian(60333.551563).Julian().JD(4) // 2460333.0516
```

### `율리우스일`/`수정율리우스일`을 `공력`으로 변환
```go
// 율리우스일을 공력으로 변환
carbon.CreateFromJulian(2460334).ToDateTimeString() // 2024-01-24 12:00:00
carbon.CreateFromJulian(2460334.051563).ToDateTimeString() // 2024-01-24 13:14:15

// 수정율리우스일을 공력으로 변환
carbon.CreateFromJulian(60333.5).ToDateTimeString() // 2024-01-24 12:00:00
carbon.CreateFromJulian(60333.551563).ToDateTimeString() // 2024-01-24 13:14:15
```

## 중국음력
> 현재는 서기 `1900`년부터 `2100`년까지의 `200`년 시간 범위만 지원

### `공력`을 `음력`으로 변환
```go
// 음력 띠 동물 가져오기
carbon.Parse("2020-08-05").Lunar().Animal() // 鼠
// 음력 축제 가져오기
carbon.Parse("2021-02-12").Lunar().Festival() // 春节

// 음력 년도 가져오기
carbon.Parse("2020-08-05").Lunar().Year() // 2020
// 음력 월 가져오기
carbon.Parse("2020-08-05").Lunar().Month() // 6
// 음력 윤월 가져오기
carbon.Parse("2020-08-05").Lunar().LeapMonth() // 4
// 음력 일 가져오기
carbon.Parse("2020-08-05").Lunar().Day() // 16

// 음력 날짜 시간 문자열 가져오기
carbon.Parse("2020-08-05").Lunar().String() // 2020-06-16
fmt.Printf("%s", carbon.Parse("2020-08-05").Lunar()) // 2020-06-16
// 음력 년도 문자열 가져오기
carbon.Parse("2020-08-05").Lunar().ToYearString() // 二零二零
// 음력 월 문자열 가져오기
carbon.Parse("2020-08-05").Lunar().ToMonthString() // 六月
// 음력 윤월 문자열 가져오기
carbon.Parse("2020-04-23").Lunar().ToMonthString() // 闰四月
// 음력 주 문자열 가져오기
carbon.Parse("2020-04-23").Lunar().ToWeekString() // 周四
// 음력 일 문자열 가져오기
carbon.Parse("2020-08-05").Lunar().ToDayString() // 十六
// 음력 날짜 문자열 가져오기
carbon.Parse("2020-08-05").Lunar().ToDateString() // 二零二零年六月十六
```

### `음력`을 `공력`으로 변환
```go
// 음력 2023년 12월 11일을 공력으로 변환
carbon.CreateFromLunar(2023, 12, 11, false).ToDateString() // 2024-01-21
// 음력 2023년 2월 11일을 공력으로 변환
carbon.CreateFromLunar(2023, 2, 11, false).ToDateString() // 2023-03-02
// 음력 2023년 윤2월 11일을 공력으로 변환
carbon.CreateFromLunar(2023, 2, 11, true).ToDateString() // 2023-04-01
```

### 날짜 판단
```go
// 유효한 음력 날짜인지
carbon.Parse("0000-00-00").Lunar().IsValid() // false
carbon.Parse("2020-08-05").Lunar().IsValid() // true

// 음력 윤년인지
carbon.Parse("2020-08-05").Lunar().IsLeapYear() // true
// 음력 윤월인지
carbon.Parse("2020-08-05").Lunar().IsLeapMonth() // false

// 쥐띠해인지
carbon.Parse("2020-08-05").Lunar().IsRatYear() // true
// 소띠해인지
carbon.Parse("2020-08-05").Lunar().IsOxYear() // false
// 호랑이띠해인지
carbon.Parse("2020-08-05").Lunar().IsTigerYear() // false
// 토끼띠해인지
carbon.Parse("2020-08-05").Lunar().IsRabbitYear() // false
// 용띠해인지
carbon.Parse("2020-08-05").Lunar().IsDragonYear() // false
// 뱀띠해인지
carbon.Parse("2020-08-05").Lunar().IsSnakeYear() // false
// 말띠해인지
carbon.Parse("2020-08-05").Lunar().IsHorseYear() // false
// 양띠해인지
carbon.Parse("2020-08-05").Lunar().IsGoatYear() // false
// 원숭이띠해인지
carbon.Parse("2020-08-05").Lunar().IsMonkeyYear() // false
// 닭띠해인지
carbon.Parse("2020-08-05").Lunar().IsRoosterYear() // false
// 개띠해인지
carbon.Parse("2020-08-05").Lunar().IsDogYear() // false
// 돼지띠해인지
carbon.Parse("2020-08-05").Lunar().IsPigYear() // false
```

## 페르시아력(이란력)

### `공력`을 `페르시아력`으로 변환
```go
// 페르시아력 년도 가져오기
carbon.Parse("2020-08-05").Persian().Year() // 1399
// 페르시아력 월 가져오기
carbon.Parse("2020-08-05").Persian().Month() // 5
// 페르시아력 일 가져오기
carbon.Parse("2020-08-05").Persian().Day() // 15

// 페르시아력 날짜 시간 문자열 가져오기
carbon.Parse("2020-08-05").Persian().String() // 1399-05-15
fmt.Printf("%s", carbon.Parse("2020-08-05").Persian()) // 1399-05-15

// 페르시아력 월 문자열 가져오기
carbon.Parse("2020-08-05").Persian().ToMonthString() // Mordad
carbon.Parse("2020-08-05").Persian().ToMonthString("en") // Mordad
carbon.Parse("2020-08-05").Persian().ToMonthString("fa") // مرداد

// 페르시아력 주 문자열 가져오기
carbon.Parse("2020-08-05").Persian().ToWeekString() // Chaharshanbeh
carbon.Parse("2020-08-05").Persian().ToWeekString("en") // Chaharshanbeh
carbon.Parse("2020-08-05").Persian().ToWeekString("fa") // چهارشنبه
```

### `페르시아력`을 `공력`으로 변환
```go
carbon.CreateFromPersian(1, 1, 1).ToDateString() // 2016-03-20
carbon.CreateFromPersian(622, 1, 1).ToDateString() // 1243-03-21
carbon.CreateFromPersian(1395, 1, 1).ToDateString() // 2016-03-20
carbon.CreateFromPersian(9377, 1, 1).ToDateString() // 9998-03-19
```

### 날짜 판단
```go
// 유효한 페르시아력 날짜인지
carbon.CreateFromPersian(1, 1, 1).IsValid() // true
carbon.CreateFromPersian(622, 1, 1).IsValid() // true
carbon.CreateFromPersian(9377, 1, 1).IsValid() // true
carbon.CreateFromPersian(0, 0, 0, 0).IsValid() // false
carbon.CreateFromPersian(2024, 0, 1).IsValid() // false
carbon.CreateFromPersian(2024, 1, 0).IsValid() // false

// 페르시아력 윤년인지
carbon.CreateFromPersian(1395, 1, 1).IsLeapYear() // true
carbon.CreateFromPersian(9377, 1, 1).IsLeapYear() // true
carbon.CreateFromPersian(622, 1, 1).IsLeapYear() // false
carbon.CreateFromPersian(9999, 1, 1).IsLeapYear() // false
```

## 히브리력(유대력)

### `공력`을 `히브리력`으로 변환
```go
// 히브리력 년도 가져오기
carbon.Parse("2024-01-01").Hebrew().Year() // 5784
// 히브리력 월 가져오기
carbon.Parse("2024-01-01").Hebrew().Month() // 10
// 히브리력 일 가져오기
carbon.Parse("2024-01-01").Hebrew().Day() // 20

// 히브리력 날짜 시간 문자열 가져오기
carbon.Parse("2024-01-01").Hebrew().String() // 5784-10-20
fmt.Printf("%s", carbon.Parse("2024-01-01").Hebrew()) // 5784-10-20

// 히브리력 월 문자열 가져오기
carbon.Parse("2020-08-05").Hebrew().ToMonthString() // Av
carbon.Parse("2020-08-05").Hebrew().ToMonthString("en") // Av
carbon.Parse("2020-08-05").Hebrew().ToMonthString("he") // אב

// 히브리력 주 문자열 가져오기
carbon.Parse("2020-08-05").Hebrew().ToWeekString() // Wednesday
carbon.Parse("2020-08-05").Hebrew().ToWeekString("en") // Wednesday
carbon.Parse("2020-08-05").Hebrew().ToWeekString("he") // רביעי
```

### `히브리력`을 `공력`으로 변환
```go
carbon.CreateFromHebrew(5784, 10, 20).ToDateString() // 2023-12-17
carbon.CreateFromHebrew(5784, 5, 1).ToDateString() // 2024-07-21
carbon.CreateFromHebrew(5786, 7, 10).ToDateString() // 2025-09-18
```

### 날짜 판단
```go
// 유효한 히브리력 날짜인지
carbon.CreateFromHebrew(5780, 14, 1).IsValid() // false
carbon.CreateFromHebrew(5780, 10, 30).IsValid() // true

// 히브리력 윤년인지
carbon.CreateFromHebrew(5784, 1, 1).IsLeapYear() // true
carbon.CreateFromHebrew(5788, 1, 1).IsLeapYear() // false
``` 