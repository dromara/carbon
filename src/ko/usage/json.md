---
head:
  - - meta
    - name: description
      content: JSON 직렬화|가벼우면서도 의미론적이고 개발자 친화적인 golang 시간 처리 라이브러리
---

# JSON 직렬화

```go
// JSON 마샬링
type User struct {
    Name      string          `json:"name"`
    Age       int             `json:"age"`
    Birthday  carbon.DateTime `json:"birthday"`
    CreatedAt carbon.DateTime `json:"created_at"`
}

user := User{
    Name:      "홍길동",
    Age:       25,
    Birthday:  carbon.Parse("1995-01-01 00:00:00"),
    CreatedAt: carbon.Now(),
}

// JSON으로 마샬링
bytes, err := json.Marshal(user)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(bytes))
// {"name":"홍길동","age":25,"birthday":"1995-01-01 00:00:00","created_at":"2020-08-05 13:14:15"}

// JSON에서 언마샬링
var newUser User
err = json.Unmarshal(bytes, &newUser)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", newUser)
// {Name:홍길동 Age:25 Birthday:1995-01-01 00:00:00 CreatedAt:2020-08-05 13:14:15}

// 사용자 정의 형식으로 JSON 직렬화
type CustomUser struct {
    Name      string                `json:"name"`
    Birthday  carbon.DateTimeLayout `json:"birthday"`
    CreatedAt carbon.DateLayout     `json:"created_at"`
}

customUser := CustomUser{
    Name:      "김철수",
    Birthday:  carbon.Parse("1990-05-15 10:30:45"),
    CreatedAt: carbon.Now(),
}

bytes, _ = json.Marshal(customUser)
fmt.Println(string(bytes))
// {"name":"김철수","birthday":"1990-05-15 10:30:45","created_at":"2020-08-05"}
``` 