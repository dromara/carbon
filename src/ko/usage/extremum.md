---
head:
  - - meta
    - name: description
      content: Carbon 은 Max/Min 으로 최대/최소 시간 조회, Closest/Farthest 로 가장 가까운/먼 시간 조회, MaxValue/MinValue 로 경계값 조회, MaxDuration/MinDuration 으로 Duration 극값 조회 기능 제공
  - - meta
    - name: keywords
      content: carbon, go-carbon, 극값 판단, 극값 경계
---

# 시간 극값

## 극값 판단
```go
c1 := carbon.Parse("2020-08-01")
c2 := carbon.Parse("2020-08-05")
c3 := carbon.Parse("2020-08-06")

// 최대 Carbon 인스턴스 반환
carbon.Max(c1, c2, c3) // c3
// 최소 Carbon 인스턴스 반환
carbon.Min(c1, c2, c3) // c1

c := carbon.Parse("2020-07-01")
// 가장 가까운 Carbon 인스턴스 반환
c.Closest(c1, c2, c3) // c1
// 가장 먼 Carbon 인스턴스 반환
c.Farthest(c1, c2, c3) // c3
```

## 극값 경계
```go
// 제로값 Carbon 반환
carbon.ZeroValue().ToString() // 0001-01-01 00:00:00 +0000 UTC
// Linux 에포크값 Carbon 반환
carbon.EpochValue().ToString() // 1970-01-01 00:00:00 +0000 UTC

// Carbon의 최대값 반환
carbon.MaxValue().ToString() // 9999-12-31 23:59:59.999999999 +0000 UTC
// Carbon의 최소값 반환
carbon.MinValue().ToString() // 0001-01-01 00:00:00 +0000 UTC

// Duration의 최대값 반환
carbon.MaxDuration().Seconds() // 9.223372036854776e+09
// Duration의 최소값 반환
carbon.MinDuration().Seconds() // -9.223372036854776e+09
``` 