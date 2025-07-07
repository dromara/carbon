---
head:
  - - meta
    - name: description
      content: 어제, 오늘, 내일|가벼우면서도 의미론적이고 개발자 친화적인 golang 시간 처리 라이브러리
  - - meta
    - name: keywords
      content: carbon, go-carbon, 어제, 오늘, 내일
---

# 어제, 오늘, 내일
어제는 오늘의 전날이고, 내일은 오늘의 다음날입니다. 다음과 같이 동일합니다:

```go
carbon.Yesterday() = carbon.Now().SubDay()
carbon.Tomorrow() = carbon.Now().AddDay()
```
현재 시간이 `2020-08-05 13:14:15.999999999 +0000 UTC`라고 가정합니다

## 어제
```go
// 어제 이 시각
fmt.Printf("%s", carbon.Yesterday()) // 2020-08-04 13:14:15
carbon.Yesterday().String() // 2020-08-04 13:14:15
carbon.Yesterday().ToString() // 2020-08-04 13:14:15.999999999 +0000 UTC
carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// 어제 날짜
carbon.Yesterday().ToDateString() // 2020-08-04
// 어제 시간
carbon.Yesterday().ToTimeString() // 13:14:15
// 지정된 시간대의 어제 이 시각
carbon.Yesterday(carbon.NewYork).ToDateTimeString() // 2020-08-04 14:14:15
// 어제 초 정밀도 타임스탬프
carbon.Yesterday().Timestamp() // 1596546855
// 어제 밀리초 정밀도 타임스탬프
carbon.Yesterday().TimestampMilli() // 1596546855999
// 어제 마이크로초 정밀도 타임스탬프
carbon.Yesterday().TimestampMicro() // 1596546855999999
// 어제 나노초 정밀도 타임스탬프
carbon.Yesterday().TimestampNano() // 1596546855999999999
```

## 오늘
```go
// 오늘 이 시각
fmt.Printf("%s", carbon.Now()) // 2020-08-05 13:14:15
carbon.Now().String() // 2020-08-05 13:14:15
carbon.Now().ToString() // 2020-08-05 13:14:15.999999999 +0000 UTC
carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
// 오늘 날짜
carbon.Now().ToDateString() // 2020-08-05
// 오늘 시간
carbon.Now().ToTimeString() // 13:14:15
// 지정된 시간대의 오늘 이 시각
carbon.Now(carbon.NewYork).ToDateTimeString() // 2020-08-05 14:14:15
// 오늘 초 정밀도 타임스탬프
carbon.Now().Timestamp() // 1596633255
// 오늘 밀리초 정밀도 타임스탬프
carbon.Now().TimestampMilli() // 1596633255999
// 오늘 마이크로초 정밀도 타임스탬프
carbon.Now().TimestampMicro() // 1596633255999999
// 오늘 나노초 정밀도 타임스탬프
carbon.Now().TimestampNano() // 1596633255999999999
```

## 내일
```go
// 내일 이 시각
fmt.Printf("%s", carbon.Tomorrow()) // 2020-08-06 13:14:15
carbon.Tomorrow().String() // 2020-08-06 13:14:15
carbon.Tomorrow().ToString() // 2020-08-06 13:14:15.999999999 +0000 UTC
carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// 내일 날짜
carbon.Tomorrow().ToDateString() // 2020-08-06
// 내일 시간
carbon.Tomorrow().ToTimeString() // 13:14:15
// 지정된 시간대의 내일 이 시각
carbon.Tomorrow(carbon.NewYork).ToDateTimeString() // 2020-08-06 14:14:15
// 내일 초 정밀도 타임스탬프
carbon.Tomorrow().Timestamp() // 1596719655
// 내일 밀리초 정밀도 타임스탬프
carbon.Tomorrow().TimestampMilli() // 1596719655999
// 내일 마이크로초 정밀도 타임스탬프
carbon.Tomorrow().TimestampMicro() // 1596719655999999
// 내일 나노초 정밀도 타임스탬프
carbon.Tomorrow().TimestampNano() // 1596719655999999999
``` 