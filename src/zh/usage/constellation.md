---
head:
  - - meta
    - name: description
      content: 提供 Constellation 获取星座名称（支持本地化），并提供 IsAries/IsTaurus/IsGemini 等星座判断方法
  - - meta
    - name: keywords
      content: carbon, go-carbon, 星座, 白羊座, 金牛座, 双子座, 巨蟹座, 狮子座, 处女座, 天秤座, 天蝎座, 射手座, 摩羯座, 水瓶座, 双鱼座
---

# 星座

## 星座名称
```go
carbon.Parse("2020-08-05 13:14:15").Constellation() // Leo
carbon.Parse("2020-08-05 13:14:15").SetLocale("zh-CN").Constellation() // 狮子座
```

## 星座判断
```go
// 是否是白羊座
carbon.Parse("2020-08-05 13:14:15").IsAries() // false
// 是否是金牛座
carbon.Parse("2020-08-05 13:14:15").IsTaurus() // false
// 是否是双子座
carbon.Parse("2020-08-05 13:14:15").IsGemini() // false
// 是否是巨蟹座
carbon.Parse("2020-08-05 13:14:15").IsCancer() // false
// 是否是狮子座
carbon.Parse("2020-08-05 13:14:15").IsLeo() // true
// 是否是处女座
carbon.Parse("2020-08-05 13:14:15").IsVirgo() // false
// 是否是天秤座
carbon.Parse("2020-08-05 13:14:15").IsLibra() // false
// 是否是天蝎座
carbon.Parse("2020-08-05 13:14:15").IsScorpio() // false
// 是否是射手座
carbon.Parse("2020-08-05 13:14:15").IsSagittarius() // false
// 是否是摩羯座
carbon.Parse("2020-08-05 13:14:15").IsCapricorn() // false
// 是否是水瓶座
carbon.Parse("2020-08-05 13:14:15").IsAquarius() // false
// 是否是双鱼座
carbon.Parse("2020-08-05 13:14:15").IsPisces() // false
```