package carbon

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name         string         `json:"name"`
	Age          int            `json:"age"`
	Birthday1    DateTime       `json:"birthday1"`
	Birthday2    DateTimeMilli  `json:"birthday2"`
	Birthday3    DateTimeMicro  `json:"birthday3"`
	Birthday4    DateTimeNano   `json:"birthday4"`
	GraduatedAt1 Date           `json:"graduated_at1"`
	GraduatedAt2 DateMilli      `json:"graduated_at2"`
	GraduatedAt3 DateMicro      `json:"graduated_at3"`
	GraduatedAt4 DateNano       `json:"graduated_at4"`
	CreatedAt1   Timestamp      `json:"created_at1"`
	CreatedAt2   TimestampMilli `json:"created_at2"`
	CreatedAt3   TimestampMicro `json:"created_at3"`
	CreatedAt4   TimestampNano  `json:"created_at4"`
	Time1        Time           `json:"time1"`
	Time2        TimeMilli      `json:"time2"`
	Time3        TimeMicro      `json:"time3"`
	Time4        TimeNano       `json:"time4"`
}

var person Person

func TestCarbon_MarshalJSON(t *testing.T) {
	person = Person{
		Name:         "gouguoyin",
		Age:          18,
		Birthday1:    DateTime{Now().SubYears(18)},
		Birthday2:    DateTimeMilli{Now().SubYears(18)},
		Birthday3:    DateTimeMicro{Now().SubYears(18)},
		Birthday4:    DateTimeNano{Now().SubYears(18)},
		GraduatedAt1: Date{Parse("2020-08-05 13:14:15")},
		GraduatedAt2: DateMilli{Parse("2020-08-05 13:14:15.999")},
		GraduatedAt3: DateMicro{Parse("2020-08-05 13:14:15.999999")},
		GraduatedAt4: DateNano{Parse("2020-08-05 13:14:15.999999999")},
		CreatedAt1:   Timestamp{Parse("2023-08-05 13:14:15")},
		CreatedAt2:   TimestampMilli{Parse("2024-08-05 13:14:15.999")},
		CreatedAt3:   TimestampMicro{Parse("2025-08-05 13:14:15.999999")},
		CreatedAt4:   TimestampNano{Parse("2025-08-05 13:14:15.999999999")},
		Time1:        Time{Parse("15:04:05")},
		Time2:        TimeMilli{Parse("15:04:05.999")},
		Time3:        TimeMicro{Parse("15:04:05.999999")},
		Time4:        TimeNano{Parse("15:04:05.999999999")},
	}
	data, err := json.Marshal(&person)
	assert.Nil(t, err)
	fmt.Printf("Person output by json:\n%s\n", data)
}

func TestCarbon_UnmarshalJSON(t *testing.T) {
	str := `{
		"name": "gouguoyin",
		"age": 18,
		"birthday1": "2003-07-16 16:22:02",
		"birthday2": "2003-07-16 16:22:02.999",
		"birthday3": "2003-07-16 16:22:02.999999",
		"birthday4": "2003-07-16 16:22:02.999999999",
		"graduated_at1": "2020-08-05",
		"graduated_at2": "2020-08-05.999",
		"graduated_at3": "2020-08-05.999999",
		"graduated_at4": "2020-08-05.999999999",
		"created_at1": 1596604455,
		"created_at2": 1596604455999,
		"created_at3": 1596604455999999,
		"created_at4": 1596604455999999999,
		"time1": "15:04:05",
		"time2": "15:04:05.999",
		"time3": "15:04:05.999999",
		"time4": "15:04:05.999999999",
	}`

	err := json.Unmarshal([]byte(str), &person)
	assert.Nil(t, err)
	fmt.Printf("Json string parse to person:\n%+v\n", person)
}

func TestError_Json(t *testing.T) {
	str := `{
		"name": "",
		"age": 0,
		"birthday1": "",
		"birthday2": "",
		"birthday3": "",
		"birthday4": "",
		"graduated_at1": "xxx",
		"graduated_at2": "xxx",
		"graduated_at3": "xxx",
		"graduated_at4": "xxx",
		"created_at1": 0,
		"created_at2": 0,
		"created_at3": 0,
		"created_at4": 0,
		"time1": "",
		"time2": "",
		"time3": "",
		"time4": "",
	}`
	err := json.Unmarshal([]byte(str), &person)
	assert.NotNil(t, err)
	fmt.Printf("Json string parse to person:\n%+v\n", person)
}
