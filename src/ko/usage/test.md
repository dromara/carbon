---
head:
  - - meta
    - name: description
      content: 테스트 | SetTestNow/IsTestNow/ClearTestNow 제공, 고정 시간을 동결해 단위 테스트를 지원하고 시간 관련 동작을 시뮬레이션 시간 기준으로 실행
  - - meta
    - name: keywords
      content: carbon, go-carbon, 테스트 시간, 동결 시간
---

# 테스트
고정된 시간을 동결하여 임의의 시간을 `현재 시간`으로 설정하고, 후속 작업이 실제 시스템 시간이 아닌 이 시뮬레이션 시간을 기반으로 실행되도록 하여 단위 테스트에 편리합니다

## 테스트 시간 설정
```go
now := carbon.Parse("2020-08-05")
carbon.SetTestNow(now)

carbon.Now().ToDateString() // 2020-08-05
carbon.Yesterday().ToDateString() // 2020-08-04
carbon.Tomorrow().ToDateString() // 2020-08-05
carbon.Now().DiffForHumans() // just now
carbon.Yesterday().DiffForHumans() // 1 day ago
carbon.Tomorrow().DiffForHumans() // 1 day from now
carbon.Parse("2020-10-05").DiffForHumans() // 2 months from now
now.DiffForHumans(carbon.Parse("2020-10-05")) // 2 months before
```

## 테스트 시간인지
```go
carbon.IsTestNow() 
```

## 테스트 시간 지우기
```go
carbon.ClearTestNow()
``` 