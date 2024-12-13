package carbon

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCarbon_Scan(t *testing.T) {
	c := NewCarbon()

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestCarbon_Value(t *testing.T) {
	c := Now()
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestDateTime_Scan(t *testing.T) {
	c := NewDateTime(Now())

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestDateTime_Value(t *testing.T) {
	c := NewDateTime(Now())
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestDateTimeMilli_Scan(t *testing.T) {
	c := NewDateTimeMilli(Now())

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestDateTimeMilli_Value(t *testing.T) {
	c := NewDateTimeMilli(Now())
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestDateTimeMicro_Scan(t *testing.T) {
	c := NewDateTimeMicro(Now())

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestDateTimeMicro_Value(t *testing.T) {
	c := NewDateTimeMicro(Now())
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestDateTimeNano_Scan(t *testing.T) {
	c := NewDateTimeNano(Now())

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestDateTimeNano_Value(t *testing.T) {
	c := NewDateTimeNano(Now())
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestDate_Scan(t *testing.T) {
	c := NewDate(Now())

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestDate_Value(t *testing.T) {
	c := NewDate(Now())
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestDateMilli_Scan(t *testing.T) {
	c := NewDateMilli(Now())

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestDateMilli_Value(t *testing.T) {
	c := NewDateMilli(Now())
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestDateMicro_Scan(t *testing.T) {
	c := NewDateMicro(Now())

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestDateMicro_Value(t *testing.T) {
	c := NewDateMicro(Now())
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestDateNano_Scan(t *testing.T) {
	c := NewDateNano(Now())

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestDateNano_Value(t *testing.T) {
	c := NewDateNano(Now())
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestTime_Scan(t *testing.T) {
	c := NewTime(Now())

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestTime_Value(t *testing.T) {
	c := NewTime(Now())
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestTimeMilli_Scan(t *testing.T) {
	c := NewTimeMilli(Now())

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestTimeMilli_Value(t *testing.T) {
	c := NewTimeMilli(Now())
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestTimeMicro_Scan(t *testing.T) {
	c := NewTimeMicro(Now())

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestTimeMicro_Value(t *testing.T) {
	c := NewTimeMicro(Now())
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestTimeNano_Scan(t *testing.T) {
	c := NewTimeNano(Now())

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestTimeNano_Value(t *testing.T) {
	c := NewTimeNano(Now())
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestTimestamp_Scan(t *testing.T) {
	c := NewTimestamp(Now())

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestTimestamp_Value(t *testing.T) {
	c := NewTimestamp(Now())
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestTimestampMilli_Scan(t *testing.T) {
	c := NewTimestampMilli(Now())

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestTimestampMilli_Value(t *testing.T) {
	c := NewTimestampMilli(Now())
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestTimestampMicro_Scan(t *testing.T) {
	c := NewTimestampMicro(Now())

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestTimestampMicro_Value(t *testing.T) {
	c := NewTimestampMicro(Now())
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestTimestampNano_Scan(t *testing.T) {
	c := NewTimestampNano(Now())

	e1 := c.Scan(Now().ToDateString())
	require.NoError(t, e1)

	e2 := c.Scan([]byte(Now().ToDateString()))
	require.NoError(t, e2)

	e3 := c.Scan(time.Now())
	require.NoError(t, e3)
}

func TestTimestampNano_Value(t *testing.T) {
	c := NewTimestampNano(Now())
	v, err := c.Value()
	require.NoError(t, err)
	assert.Equal(t, v, c.time)
}

func TestError_Scan(t *testing.T) {
	str := "xxxx"

	c1 := NewCarbon()
	e1 := c1.Scan(str)
	assert.Equal(t, e1, failedScanError(str))

	c2 := NewDateTime(Now())
	e2 := c2.Scan(str)
	assert.Equal(t, e2, failedScanError(str))

	c3 := NewDateTimeMilli(Now())
	e3 := c3.Scan(str)
	assert.Equal(t, e3, failedScanError(str))

	c4 := NewDateTimeMicro(Now())
	e4 := c4.Scan(str)
	assert.Equal(t, e4, failedScanError(str))

	c5 := NewDateTimeNano(Now())
	e5 := c5.Scan(str)
	assert.Equal(t, e5, failedScanError(str))

	c6 := NewDate(Now())
	e6 := c6.Scan(str)
	assert.Equal(t, e6, failedScanError(str))

	c7 := NewDateMilli(Now())
	e7 := c7.Scan(str)
	assert.Equal(t, e7, failedScanError(str))

	c8 := NewDateMicro(Now())
	e8 := c8.Scan(str)
	assert.Equal(t, e8, failedScanError(str))

	c9 := NewDateNano(Now())
	e9 := c9.Scan(str)
	assert.Equal(t, e9, failedScanError(str))

	c10 := NewTime(Now())
	e10 := c10.Scan(str)
	assert.Equal(t, e10, failedScanError(str))

	c11 := NewTimeMilli(Now())
	e11 := c11.Scan(str)
	assert.Equal(t, e11, failedScanError(str))

	c12 := NewTimeMicro(Now())
	e12 := c12.Scan(str)
	assert.Equal(t, e12, failedScanError(str))

	c13 := NewTimeNano(Now())
	e13 := c13.Scan(str)
	assert.Equal(t, e13, failedScanError(str))

	c14 := NewTimestamp(Now())
	e14 := c14.Scan(str)
	assert.Equal(t, e14, failedScanError(str))

	c15 := NewTimestampMilli(Now())
	e15 := c15.Scan(str)
	assert.Equal(t, e15, failedScanError(str))

	c16 := NewTimestampMicro(Now())
	e16 := c16.Scan(str)
	assert.Equal(t, e16, failedScanError(str))

	c17 := NewTimestampNano(Now())
	e17 := c17.Scan(str)
	assert.Equal(t, e17, failedScanError(str))
}

func TestError_Value(t *testing.T) {
	c1 := Parse("")
	v1, e1 := c1.Value()
	require.NoError(t, e1)
	assert.Nil(t, v1)

	c2 := NewDateTime(c1)
	v2, e2 := c2.Value()
	require.NoError(t, e2)
	assert.Nil(t, v2)

	c3 := NewDateTimeMilli(c1)
	v3, e3 := c3.Value()
	require.NoError(t, e3)
	assert.Nil(t, v3)

	c4 := NewDateTimeMicro(c1)
	v4, e4 := c4.Value()
	require.NoError(t, e4)
	assert.Nil(t, v4)

	c5 := NewDateTimeNano(c1)
	v5, e5 := c5.Value()
	require.NoError(t, e5)
	assert.Nil(t, v5)

	c6 := NewDate(c1)
	v6, e6 := c6.Value()
	require.NoError(t, e6)
	assert.Nil(t, v6)

	c7 := NewDateMilli(c1)
	v7, e7 := c7.Value()
	require.NoError(t, e7)
	assert.Nil(t, v7)

	c8 := NewDateMicro(c1)
	v8, e8 := c8.Value()
	require.NoError(t, e8)
	assert.Nil(t, v8)

	c9 := NewDateNano(c1)
	v9, e9 := c9.Value()
	require.NoError(t, e9)
	assert.Nil(t, v9)

	c10 := NewTime(c1)
	v10, e10 := c10.Value()
	require.NoError(t, e10)
	assert.Nil(t, v10)

	c11 := NewTimeMilli(c1)
	v11, e11 := c11.Value()
	require.NoError(t, e11)
	assert.Nil(t, v11)

	c12 := NewTimeMicro(c1)
	v12, e12 := c12.Value()
	require.NoError(t, e12)
	assert.Nil(t, v12)

	c13 := NewTimeNano(c1)
	v13, e13 := c13.Value()
	require.NoError(t, e13)
	assert.Nil(t, v13)

	c14 := NewTimestamp(c1)
	v14, e14 := c14.Value()
	require.NoError(t, e14)
	assert.Nil(t, v14)

	c15 := NewTimestampMilli(c1)
	v15, e15 := c15.Value()
	require.NoError(t, e15)
	assert.Nil(t, v15)

	c16 := NewTimestampMicro(c1)
	v16, e16 := c16.Value()
	require.NoError(t, e16)
	assert.Nil(t, v16)

	c17 := NewTimestampNano(c1)
	v17, e17 := c17.Value()
	require.NoError(t, e17)
	assert.Nil(t, v17)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`

	Birthday1 Carbon `json:"birthday1"`

	Birthday2 DateTime      `json:"birthday2"`
	Birthday3 DateTimeMilli `json:"birthday3"`
	Birthday4 DateTimeMicro `json:"birthday4"`
	Birthday5 DateTimeNano  `json:"birthday5"`

	Birthday6 Date      `json:"birthday6"`
	Birthday7 DateMilli `json:"birthday7"`
	Birthday8 DateMicro `json:"birthday8"`
	Birthday9 DateNano  `json:"birthday9"`

	Birthday10 Time      `json:"birthday10"`
	Birthday11 TimeMilli `json:"birthday11"`
	Birthday12 TimeMicro `json:"birthday12"`
	Birthday13 TimeNano  `json:"birthday13"`

	Birthday14 Timestamp      `json:"birthday14"`
	Birthday15 TimestampMilli `json:"birthday15"`
	Birthday16 TimestampMicro `json:"birthday16"`
	Birthday17 TimestampNano  `json:"birthday17"`
}

func TestCarbon_MarshalJSON(t *testing.T) {
	var c = Parse("2020-08-05 13:14:15.999999999", PRC)
	var person = Person{
		Name: "gouguoyin",
		Age:  18,

		Birthday1: c,
		Birthday2: NewDateTime(c),
		Birthday3: NewDateTimeMilli(c),
		Birthday4: NewDateTimeMicro(c),
		Birthday5: NewDateTimeNano(c),

		Birthday6: NewDate(c),
		Birthday7: NewDateMilli(c),
		Birthday8: NewDateMicro(c),
		Birthday9: NewDateNano(c),

		Birthday10: NewTime(c),
		Birthday11: NewTimeMilli(c),
		Birthday12: NewTimeMicro(c),
		Birthday13: NewTimeNano(c),

		Birthday14: NewTimestamp(c),
		Birthday15: NewTimestampMilli(c),
		Birthday16: NewTimestampMicro(c),
		Birthday17: NewTimestampNano(c),
	}

	data, marshalErr := json.Marshal(&person)
	require.NoError(t, marshalErr)
	fmt.Printf("json encode:\n%s\n", data)

	tests := []struct {
		name   string
		want   string
		actual string
	}{
		{
			name:   "birthday1",
			actual: person.Birthday1.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "birthday2",
			actual: person.Birthday2.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "Birthday3",
			actual: person.Birthday3.String(),
			want:   "2020-08-05 13:14:15.999",
		},
		{
			name:   "Birthday4",
			actual: person.Birthday4.String(),
			want:   "2020-08-05 13:14:15.999999",
		},
		{
			name:   "Birthday5",
			actual: person.Birthday5.String(),
			want:   "2020-08-05 13:14:15.999999999",
		},

		{
			name:   "Birthday6",
			actual: person.Birthday6.String(),
			want:   "2020-08-05",
		},
		{
			name:   "Birthday7",
			actual: person.Birthday7.String(),
			want:   "2020-08-05.999",
		},
		{
			name:   "Birthday8",
			actual: person.Birthday8.String(),
			want:   "2020-08-05.999999",
		},
		{
			name:   "Birthday9",
			actual: person.Birthday9.String(),
			want:   "2020-08-05.999999999",
		},

		{
			name:   "Birthday10",
			actual: person.Birthday10.String(),
			want:   "13:14:15",
		},
		{
			name:   "Birthday11",
			actual: person.Birthday11.String(),
			want:   "13:14:15.999",
		},
		{
			name:   "Birthday12",
			actual: person.Birthday12.String(),
			want:   "13:14:15.999999",
		},
		{
			name:   "Birthday13",
			actual: person.Birthday13.String(),
			want:   "13:14:15.999999999",
		},

		{
			name:   "Birthday14",
			actual: person.Birthday14.String(),
			want:   "1596604455",
		},
		{
			name:   "Birthday15",
			actual: person.Birthday15.String(),
			want:   "1596604455999",
		},
		{
			name:   "Birthday16",
			actual: person.Birthday16.String(),
			want:   "1596604455999999",
		},
		{
			name:   "Birthday17",
			actual: person.Birthday17.String(),
			want:   "1596604455999999999",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "json.Marshal()")
		})
	}
}

func TestCarbon_UnmarshalJSON(t *testing.T) {
	jsonStr := `{
		"name":"gouguoyin",
		"age":18,
        "birthday1":"2020-08-05 13:14:15",
		"Birthday2":"2020-08-05 13:14:15",
		"Birthday3":"2020-08-05 13:14:15.999",
		"Birthday4":"2020-08-05 13:14:15.999999",
		"Birthday5":"2020-08-05 13:14:15.999999999",
		"Birthday6":"2020-08-05",
		"Birthday7":"2020-08-05.999",
		"Birthday8":"2020-08-05.999999",
		"Birthday9":"2020-08-05.999999999",
		"Birthday10":"13:14:15",
		"Birthday11":"13:14:15.999",
		"Birthday12":"13:14:15.999999",
		"Birthday13":"13:14:15.999999999",
		"Birthday14":1596604455,
		"Birthday15":1596604455999,
		"Birthday16":1596604455999999,
		"Birthday17":1596604455999999999
	}`

	var person Person
	unmarshalErr := json.Unmarshal([]byte(jsonStr), &person)
	require.NoError(t, unmarshalErr)
	fmt.Printf("json decode:\n%+v\n", person)

	tests := []struct {
		name   string
		want   string
		actual string
	}{
		{
			name:   "birthday1",
			actual: person.Birthday1.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "Birthday2",
			actual: person.Birthday2.String(),
			want:   "2020-08-05 13:14:15",
		},
		{
			name:   "Birthday3",
			actual: person.Birthday3.String(),
			want:   "2020-08-05 13:14:15.999",
		},
		{
			name:   "Birthday4",
			actual: person.Birthday4.String(),
			want:   "2020-08-05 13:14:15.999999",
		},
		{
			name:   "Birthday5",
			actual: person.Birthday5.String(),
			want:   "2020-08-05 13:14:15.999999999",
		},
		{
			name:   "Birthday6",
			actual: person.Birthday6.String(),
			want:   "2020-08-05",
		},
		{
			name:   "Birthday7",
			actual: person.Birthday7.String(),
			want:   "2020-08-05.999",
		},
		{
			name:   "Birthday8",
			actual: person.Birthday8.String(),
			want:   "2020-08-05.999999",
		},
		{
			name:   "Birthday9",
			actual: person.Birthday9.String(),
			want:   "2020-08-05.999999999",
		},

		{
			name:   "Birthday10",
			actual: person.Birthday10.String(),
			want:   "13:14:15",
		},
		{
			name:   "Birthday11",
			actual: person.Birthday11.String(),
			want:   "13:14:15.999",
		},
		{
			name:   "Birthday12",
			actual: person.Birthday12.String(),
			want:   "13:14:15.999999",
		},
		{
			name:   "Birthday13",
			actual: person.Birthday13.String(),
			want:   "13:14:15.999999999",
		},

		{
			name:   "Birthday14",
			actual: person.Birthday14.String(),
			want:   "1596604455",
		},
		{
			name:   "Birthday15",
			actual: person.Birthday15.String(),
			want:   "1596604455999",
		},
		{
			name:   "Birthday16",
			actual: person.Birthday16.String(),
			want:   "1596604455999999",
		},
		{
			name:   "Birthday17",
			actual: person.Birthday17.String(),
			want:   "1596604455999999999",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "json.Unmarshal()")
		})
	}
}

func TestCarbon_TimestampToInt64(t *testing.T) {
	jsonStr := `{
		"Birthday14":1596604455,
		"Birthday15":1596604455999,
		"Birthday16":1596604455999999,
		"Birthday17":1596604455999999999
	}`
	var person Person
	unmarshalErr := json.Unmarshal([]byte(jsonStr), &person)
	require.NoError(t, unmarshalErr)
	fmt.Printf("json decode:\n%+v\n", person)

	tests := []struct {
		name   string
		want   int64
		actual int64
	}{
		{
			name:   "Birthday14",
			actual: person.Birthday14.Int64(),
			want:   1596604455,
		},
		{
			name:   "Birthday15",
			actual: person.Birthday15.Int64(),
			want:   1596604455999,
		},
		{
			name:   "Birthday16",
			actual: person.Birthday16.Int64(),
			want:   1596604455999999,
		},
		{
			name:   "Birthday17",
			actual: person.Birthday17.Int64(),
			want:   1596604455999999999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.actual, "json.Unmarshal()")
		})
	}
}

// https://github.com/dromara/carbon/issues/225
func TestCarbon_Issue225(t *testing.T) {
	emptyStr := `{
		"birthday1":"",
		"birthday2":"",
		"birthday3":"",
		"birthday4":"",
		"birthday5":"",
		"birthday6":"",
		"birthday7":"",
		"birthday8":"",
		"birthday9":"",
		"birthday10":"",
		"birthday11":"",
		"birthday12":"",
		"birthday13":"",
		"birthday14":"",
		"birthday15":"",
		"birthday16":"",
		"birthday17":""
	}`
	nullStr := `{
		"birthday1":null,
		"birthday2":null,
		"birthday3":null,
		"birthday4":null,
		"birthday5":null,
		"birthday6":null,
		"birthday7":null,
		"birthday8":null,
		"birthday9":null,
		"birthday10":null,
		"birthday11":null,
		"birthday12":null,
		"birthday13":null,
		"birthday14":null,
		"birthday15":null,
		"birthday16":null,
		"birthday17":null
	}`

	var person Person
	emptyErr := json.Unmarshal([]byte(emptyStr), &person)
	require.NoError(t, emptyErr)

	nullErr := json.Unmarshal([]byte(nullStr), &person)
	require.NoError(t, nullErr)
}

// https://github.com/dromara/carbon/issues/240
func TestCarbon_Issue240(t *testing.T) {
	jsonStr := `{
		"birthday1":"",
		"birthday2":null
	}`

	var person Person
	emptyErr := json.Unmarshal([]byte(jsonStr), &person)
	require.NoError(t, emptyErr)
	assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", person.Birthday1.StdTime().String())
	assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", person.Birthday2.StdTime().String())
	assert.True(t, person.Birthday1.IsZero())
	assert.True(t, person.Birthday2.IsZero())
	assert.False(t, person.Birthday1.IsValid())
	assert.False(t, person.Birthday2.IsValid())
}

// https://github.com/dromara/carbon/issues/243
func TestCarbon_Issue243(t *testing.T) {
	prepareTest(t)

	SetDefault(Default{
		Layout:       DateTimeLayout,
		Timezone:     PRC,
		Locale:       "en",
		WeekStartsAt: Sunday,
	})

	type Project struct {
		StartDate DateTime `gorm:"column:start_date" json:"startDate"`
		EndDate   DateTime `gorm:"column:end_date" json:"endDate"`
	}

	project := new(Project)
	jsonStr := `{"startDate":"2024-10-01 00:00:00","endDate":"2024-10-31 23:59:59"}`
	err := json.Unmarshal([]byte(jsonStr), &project)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, "PRC", project.StartDate.Location())
	assert.Equal(t, "PRC", project.EndDate.Location())
}
