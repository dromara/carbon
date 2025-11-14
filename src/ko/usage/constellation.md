---
head:
  - - meta
    - name: description
      content: "Constellation 으로 별자리 이름 조회(로컬라이즈 지원), IsAries/IsTaurus/IsGemini 등 별자리 판단 메서드 제공"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, 별자리, 양자리, 황소자리, 쌍둥이자리, 게자리, 사자자리, 처녀자리, 천칭자리, 전갈자리, 궁수자리, 염소자리, 물병자리, 물고기자리"
---

# 별자리

## 별자리 이름
```go
carbon.Parse("2020-08-05 13:14:15").Constellation() // Leo
carbon.Parse("2020-08-05 13:14:15").SetLocale("zh-CN").Constellation() // 狮子座
```

## 별자리 판단
```go
// 양자리인지
carbon.Parse("2020-08-05 13:14:15").IsAries() // false
// 황소자리인지
carbon.Parse("2020-08-05 13:14:15").IsTaurus() // false
// 쌍둥이자리인지
carbon.Parse("2020-08-05 13:14:15").IsGemini() // false
// 게자리인지
carbon.Parse("2020-08-05 13:14:15").IsCancer() // false
// 사자자리인지
carbon.Parse("2020-08-05 13:14:15").IsLeo() // true
// 처녀자리인지
carbon.Parse("2020-08-05 13:14:15").IsVirgo() // false
// 천칭자리인지
carbon.Parse("2020-08-05 13:14:15").IsLibra() // false
// 전갈자리인지
carbon.Parse("2020-08-05 13:14:15").IsScorpio() // false
// 궁수자리인지
carbon.Parse("2020-08-05 13:14:15").IsSagittarius() // false
// 염소자리인지
carbon.Parse("2020-08-05 13:14:15").IsCapricorn() // false
// 물병자리인지
carbon.Parse("2020-08-05 13:14:15").IsAquarius() // false
// 물고기자리인지
carbon.Parse("2020-08-05 13:14:15").IsPisces() // false
``` 