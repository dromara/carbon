---
head:
  - - meta
    - name: description
      content: "Carbon テスト機能の詳細説明、SetTestNow/IsTestNow/ClearTestNow/Sleep メソッドを提供、ユニットテスト用の固定時刻の凍結をサポート、時間関連の処理をシミュレーション時刻に基づいて実行、時間依存コードのテストに便利"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, テスト時間, 凍結時間, シミュレーション時間, SetTestNow, IsTestNow, ClearTestNow, Sleep, ユニットテスト, 時間テスト, テストツール"
---

# テスト
現在の時刻を固定し、任意の時刻を現在の時刻に設定することができます。これにより、実際のシステム時刻ではなく、このシミュレーション時刻に基づいて後続の操作が実行され、ユニットテストが容易になります

## テスト時間の設定
```go
now := carbon.Parse("2020-08-05")
carbon.SetTestNow(now)

carbon.Now().ToDateString() // 2020-08-05
carbon.Yesterday().ToDateString() // 2020-08-04
carbon.Tomorrow().ToDateString() // 2020-08-06
carbon.Now().DiffForHumans() // just now
carbon.Yesterday().DiffForHumans() // 1 day ago
carbon.Tomorrow().DiffForHumans() // 1 day from now
carbon.Parse("2020-10-05").DiffForHumans() // 2 months from now
now.DiffForHumans(carbon.Parse("2020-10-05")) // 2 months before
```

## テスト時間かどうか
```go
carbon.IsTestNow() 
```

## テスト時間のクリア
```go
carbon.ClearTestNow()
```

## スリープ
時間を凍結した状態では、`Sleep` は実時間を待たずに凍結された現在時刻を直接進めます。凍結していない場合は `time.Sleep` と同じ動作になります。

```go
carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15"))
carbon.Sleep(2 * time.Hour)

carbon.Now().ToDateTimeString() // 2020-08-05 15:14:15
```