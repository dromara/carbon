---
head:
  - - meta
    - name: description
      content: グローバル・デフォルトの設定|軽量で、意味的に、開発者に優しい golang 時間処理ライブラリ
  - - meta
    - name: keywords
      content: carbon, go-carbon, 単一のデフォルト値設定, 複数のデフォルト値設定, デフォルト値リセット
---

# グローバル・デフォルトの設定
デフォルトのタイムゾーンは`UTC`、ロケールは`英語`、週の開始日は`月曜日`、週末は`土曜日`と`日曜日`。

## 単一のデフォルト値の設定
```go
carbon.SetLayout(carbon.DateTimeLayout)
carbon.SetTimezone(carbon.UTC)
carbon.SetLocale("en")
carbon.SetWeekStartsAt(carbon.Monday)
carbon.SetWeekendDays([]carbon.Weekday{carbon.Saturday, carbon.Sunday,})
```

## 複数のデフォルト値の設定
```go
carbon.SetDefault(carbon.Default{
  Layout: carbon.DateTimeLayout,
  Timezone: carbon.UTC,
  Locale: "en",
  WeekStartsAt: carbon.Monday,
  WeekendDays: []carbon.Weekday{carbon.Saturday, carbon.Sunday,},
})
```

## デフォルト値のリセット
```go
carbon.ResetDefault()
```
これは、1つのテストの変更が他のテストに影響を与えないようにするために、テストで特に便利です。