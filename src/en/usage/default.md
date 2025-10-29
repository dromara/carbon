---
head:
  - - meta
    - name: description
      content: Set default | Provide single setters SetLayout/SetTimezone/SetLocale/SetWeekStartsAt/SetWeekendDays and batch setter SetDefault; support ResetDefault to restore defaults, useful for test isolation
  - - meta
    - name: keywords
      content: carbon, go-carbon, set single default, set multiple default, reset default
---

# Set global default
Default timezone is `UTC`, language locale is `English`, start day of the week is `Monday` and weekend days of the week are `Saturday` and `Sunday`.

## Set single default values
```go
carbon.SetLayout(carbon.DateTimeLayout)
carbon.SetTimezone(carbon.UTC)
carbon.SetLocale("en")
carbon.SetWeekStartsAt(carbon.Monday)
carbon.SetWeekendDays([]carbon.Weekday{carbon.Saturday, carbon.Sunday,})
```

## Set multiple default values
```go
carbon.SetDefault(carbon.Default{
  Layout: carbon.DateTimeLayout,
  Timezone: carbon.UTC,
  Locale: "en",
  WeekStartsAt: carbon.Monday,
  WeekendDays: []carbon.Weekday{carbon.Saturday, carbon.Sunday,},
})
```

## Reset default values
```go
carbon.ResetDefault()
```
This is particularly useful in testing to ensure that changes in one test don't affect others.


