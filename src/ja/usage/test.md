# テスト
現在の時刻を固定し、任意の時刻を現在の時刻に設定することができます。これにより、実際のシステム時刻ではなく、このシミュレーション時刻に基づいて後続の操作が実行され、セルテストが容易になります

## テスト時間の設定
```go
now := carbon.Parse("2020-08-05")
carbon.SetTestNow(now)

carbon.Now().ToDateString() // 2020-08-05
carbon.Yesterday().ToDateString() // 2020-08-04
carbon.Tomorrow().ToDateString() // 2020-08-05
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

## パージテスト時間
```go
carbon.ClearTestNow()
```