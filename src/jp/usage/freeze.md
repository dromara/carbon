# 時間凍結
テスト時間の設定をサポートし、テストがフレンドリーで、ユニットテストに便利である
## 凍結時間の設定
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

## 凍結時間かどうかを判断する
```go
carbon.IsTestNow() 
```

## パージ凍結時間
```go
carbon.ClearTestNow()
```