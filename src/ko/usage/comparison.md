---
head:
  - - meta
    - name: description
      content: 시간 비교|가벼우면서도 의미론적이고 개발자 친화적인 golang 시간 처리 라이브러리
---

# 시간 비교

```go
// 시간 비교
c1 := carbon.Parse("2020-08-05 13:14:15")
c2 := carbon.Parse("2020-08-05 13:14:15")
c3 := carbon.Parse("2020-08-06 13:14:15")

// 같은지 비교
c1.Eq(c2) // true
c1.Equal(c2) // true
c1.Eq(c3) // false

// 다른지 비교
c1.Ne(c2) // false
c1.NotEqual(c2) // false
c1.Ne(c3) // true

// 큰지 비교
c1.Gt(c2) // false
c1.GreaterThan(c2) // false
c1.Gt(c3) // false
c3.Gt(c1) // true

// 크거나 같은지 비교
c1.Gte(c2) // true
c1.GreaterThanOrEqual(c2) // true
c1.Gte(c3) // false

// 작은지 비교
c1.Lt(c2) // false
c1.LessThan(c2) // false
c1.Lt(c3) // true

// 작거나 같은지 비교
c1.Lte(c2) // true
c1.LessThanOrEqual(c2) // true
c1.Lte(c3) // true

// 범위 내에 있는지 비교
c1.Between(c2, c3) // true
c1.BetweenIncludedStart(c2, c3) // true
c1.BetweenIncludedEnd(c2, c3) // true
c1.BetweenIncludedBoth(c2, c3) // true
``` 