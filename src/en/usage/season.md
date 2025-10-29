---
head:
  - - meta
    - name: description
      content: Season | Provides Season to get season names (with localization), supports StartOfSeason/EndOfSeason to get season boundaries, and season checks such as IsSpring/IsSummer/IsAutumn/IsWinter
  - - meta
    - name: keywords
      content: carbon, go-carbon, season, spring, summer, autumn, winter, season boundary, season judgment
---

# Season
According to the meteorological division method, `March` to `May` is `spring`, `June` to `August` is `summer`, `September` to `November` is `autumn`, and `December` to `February` is `winter`

## Season name
```go
carbon.Parse("2020-08-05 13:14:15").Season() // Summer
carbon.Parse("2020-08-05 13:14:15").SetLocale("zh-CN").Season() // 夏季
```

## Season boundary
```go
// Start of the season
carbon.Parse("2020-08-05 13:14:15").StartOfSeason().ToDateTimeString() // 2020-06-01 00:00:00
// End of the season
carbon.Parse("2020-08-05 13:14:15").EndOfSeason().ToDateTimeString() // 2020-08-31 23:59:59
```

## Season judgment
```go
// Is spring
carbon.Parse("2020-08-05 13:14:15").IsSpring() // false
// Is summer
carbon.Parse("2020-08-05 13:14:15").IsSummer() // true
// Is autumn
carbon.Parse("2020-08-05 13:14:15").IsAutumn() // false
// Is winter
carbon.Parse("2020-08-05 13:14:15").IsWinter() // false
```