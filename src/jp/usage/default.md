#### グローバル・デフォルトの設定
デフォルトのタイムゾーンは` UTC `、ロケールは`英語`、週の開始日は`月曜日`、週末は`土曜日`、`日曜日`。
現在時刻が `2020-08-05 13:14:15.999999999 +0900 JST` であると仮定します。

## 単一のデフォルト値の設定
```go
carbon.SetLayout(carbon.DateTimeLayout)
carbon.SetTimezone(carbon.UTC)
carbon.SetLocale("en")
carbon.SetWeekStartsAt(carbon.Monday)
carbon.SetWeekendDays([]carbon.Weekday{carbon.Saturday, carbon.Sunday,})
```

## 複数のデフォルトを一度に設定
```go
carbon.SetDefault(carbon.Default{
  Layout: carbon.DateTimeLayout,
  Timezone: carbon.UTC,
  Locale: "en",
  WeekStartsAt: carbon.Monday,
  WeekendDays: []carbon.Weekday{carbon.Saturday, carbon.Sunday,},
})
```

## デフォルト値をリセット
```go
carbon.ResetDefault()
```
これは、1つのテストの変更が他のテストに影響を与えないようにするために、テストで特に便利です。