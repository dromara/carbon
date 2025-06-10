# JSON

## Builtin type

```go
type User struct {
  Date      *carbon.Date      `json:"date"`
  DateMilli *carbon.DateMilli `json:"date_milli"`
  DateMicro *carbon.DateMicro `json:"date_micro"`
  DateNano  *carbon.DateNano  `json:"date_nano"`
  
  Time      *carbon.Time      `json:"time"`
  TimeMilli *carbon.TimeMilli `json:"time_milli"`
  TimeMicro *carbon.TimeMicro `json:"time_micro"`
  TimeNano  *carbon.TimeNano  `json:"time_nano"`
  
  DateTime      *carbon.DateTime      `json:"date_time"`
  DateTimeMilli *carbon.DateTimeMilli `json:"date_time_milli"`
  DateTimeMicro *carbon.DateTimeMicro `json:"date_time_micro"`
  DateTimeNano  *carbon.DateTimeNano  `json:"date_time_nano"`
  
  Timestamp      *carbon.Timestamp      `json:"timestamp"`
  TimestampMilli *carbon.TimestampMilli `json:"timestamp_milli"`
  TimestampMicro *carbon.TimestampMicro `json:"timestamp_micro"`
  TimestampNano  *carbon.TimestampNano  `json:"timestamp_nano"`

  CreatedAt *carbon.DateTime `json:"created_at"`
  UpdatedAt *carbon.DateTime `json:"updated_at"`
  DeletedAt *carbon.Timestamp `json:"deleted_at"`
}

var user User

c := carbon.Parse("2020-08-05 13:14:15.999999999")

user.Date      = carbon.NewDate(c)
user.DateMilli = carbon.NewDateMilli(c)
user.DateMicro = carbon.NewDateMicro(c)
user.DateNano  = carbon.NewDateNano(c)

user.Time      = carbon.NewTime(c)
user.TimeMilli = carbon.NewTimeMilli(c)
user.TimeMicro = carbon.NewTimeMicro(c)
user.TimeNano  = carbon.NewTimeNano(c)

user.DateTime      = carbon.NewDateTime(c)
user.DateTimeMilli = carbon.NewDateTimeMilli(c)
user.DateTimeMicro = carbon.NewDateTimeMicro(c)
user.DateTimeNano  = carbon.NewDateTimeNano(c)

user.Timestamp      = carbon.NewTimestamp(c)
user.TimestampMilli = carbon.NewTimestampMilli(c)
user.TimestampMicro = carbon.NewTimestampMicro(c)
user.TimestampNano  = carbon.NewTimestampNano(c)

user.CreatedAt = carbon.NewDateTime(c)
user.UpdatedAt = carbon.NewDateTime(c)
user.DeletedAt = carbon.NewTimestamp(c)

data, err := json.Marshal(&user)
if err != nil {
  // Error handle...
  log.Fatal(err)
}
fmt.Printf("%s\n", data)
// Output
{
  "date": "2020-08-05",
  "date_milli": "2020-08-05.999",
  "date_micro": "2020-08-05.999999",
  "date_nano": "2020-08-05.999999999",
  "time": "13:14:15",
  "time_milli": "13:14:15.999",
  "time_micro": "13:14:15.999999",
  "time_nano": "13:14:15.999999999",
  "date_time": "2020-08-05 13:14:15",
  "date_time_milli": "2020-08-05 13:14:15.999",
  "date_time_micro": "2020-08-05 13:14:15.999999",
  "date_time_nano": "2020-08-05 13:14:15.999999999",
  "timestamp": 1596633255,
  "timestamp_milli": 1596633255999,
  "timestamp_micro": 1596633255999999,
  "timestamp_nano": 1596633255999999999,
  "created_at": "2020-08-05 13:14:15",
  "updated_at": "2020-08-05 13:14:15",
  "deleted_at": 1596633255
}

var person User
err := json.Unmarshal(data, &person)
if err != nil {
  // Error handle...
  log.Fatal(err)
}

fmt.Printf("person: %+v\n", person)
// Output
person: {Date:2020-08-05 DateMilli:2020-08-05.999 DateMicro:2020-08-05.999999 DateNano:2020-08-05.999999999 Time:13:14:15 TimeMilli:13:14:15.999 TimeMicro:13:14:15.999999 TimeNano:13:14:15.999999999 DateTime:2020-08-05 13:14:15 DateTimeMilli:2020-08-05 13:14:15.999 DateTimeMicro:2020-08-05 13:14:15.999999 DateTimeNano:2020-08-05 13:14:15.999999999 Timestamp:1596633255 TimestampMilli:1596633255999 TimestampMicro:1596633255999999 TimestampNano:1596633255999999999 CreatedAt:2020-08-05 13:14:15 UpdatedAt:2020-08-05 13:14:15 DeletedAt:1596633255}
```

## Customize type
```go
type RFC3339Type string
// implement "carbon.LayoutTyper" interface
func (RFC3339Type) Layout() string {
  return carbon.RFC3339Layout
}
// implement "carbon.DataTyper" interface (not required, default data type is datetime)
func (RFC3339Type) DataType() string {
  return "datetime"
}

type ISO8601Type string
// implement "carbon.FormatTyper" interface
func (ISO8601Type) Format() string {
  return carbon.ISO8601Format
}
// implement "carbon.DataTyper" interface (not required, default data type is datetime)
func (RFC3339Type) DataType() string {
  return "datetime"
}

type User struct {
  Customer1 *carbon.LayoutType[RFC3339Type] `json:"customer1"`
  Customer2 *carbon.FormatType[ISO8601Type] `json:"customer2"`
}

var user User

c := carbon.Parse("2020-08-05 13:14:15")

user.Customer1 = carbon.NewLayoutType[RFC3339Type](c)
user.Customer2 = carbon.NewFormatType[ISO8601Type](c)

data, err := json.Marshal(&user)
if err != nil {
  // Error handle...
  log.Fatal(err)
}
fmt.Printf("%s\n", data)
// Output
{"customer1":"2020-08-05T13:14:15Z", "customer2":"2020-08-05T13:14:15+00:00"}

var person User
err := json.Unmarshal(data, &person)
if err != nil {
  // Error handle...
  log.Fatal(err)
}

fmt.Printf("person: %+v\n", person)
// Output
person: {Customer1:2020-08-05T13:14:15Z Customer2:2020-08-05T13:14:15+00:00}
```