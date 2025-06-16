---
head:
  - - meta
    - name: description
      content: Set default | A simple, semantic and developer-friendly time package for golang
  - - meta
    - name: keywords
      content: carbon, go-carbon, default
---

# Set globally default
Default timezone is `UTC`, language locale is `English`, start day of the week is `Monday` and weekend days of the week are `Saturday` and `Sunday`.

## Set single default value
```go
carbon.SetLayout(carbon.DateTimeLayout)
carbon.SetTimezone(carbon.UTC)
carbon.SetLocale("en")
carbon.SetWeekStartsAt(carbon.Monday)
carbon.SetWeekendDays([]carbon.Weekday{carbon.Saturday, carbon.Sunday,})
```

## Set multiple default value
```go
carbon.SetDefault(carbon.Default{
  Layout: carbon.DateTimeLayout,
  Timezone: carbon.UTC,
  Locale: "en",
  WeekStartsAt: carbon.Monday,
  WeekendDays: []carbon.Weekday{carbon.Saturday, carbon.Sunday,},
})
```

## Resetting default value
```go
carbon.ResetDefault()
```
This is particularly useful in testing to ensure that changes in one test don't affect others.


