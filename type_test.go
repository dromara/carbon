package carbon

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type carbonTypeModel struct {
	Carbon1 Carbon  `json:"carbon1"`
	Carbon2 *Carbon `json:"carbon2"`
}

func TestCarbonType_Scan(t *testing.T) {
	c := Now()

	t.Run("[]byte type", func(t *testing.T) {
		assert.Nil(t, c.Scan([]byte(c.ToDateString())))
	})

	t.Run("string type", func(t *testing.T) {
		assert.Nil(t, c.Scan(c.ToDateString()))
	})

	t.Run("int64 type", func(t *testing.T) {
		assert.Nil(t, c.Scan(c.Timestamp()))
	})

	t.Run("time type", func(t *testing.T) {
		tt := time.Now()
		assert.Nil(t, c.Scan(tt))
	})

	t.Run("*time type", func(t *testing.T) {
		tt := time.Now()
		assert.Nil(t, c.Scan(&tt))
	})

	t.Run("nil type", func(t *testing.T) {
		assert.Nil(t, c.Scan(nil))
	})

	t.Run("unsupported type", func(t *testing.T) {
		assert.Error(t, c.Scan(true))
		assert.Error(t, c.Scan(func() {}))
		assert.Error(t, c.Scan(float64(0)))
		assert.Error(t, c.Scan(map[string]string{}))
	})
}

func TestCarbonType_Value(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		v, err := Parse("").Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("zero time", func(t *testing.T) {
		v, err := NewCarbon().Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("invalid time", func(t *testing.T) {
		v, err := Parse("xxx").Value()
		assert.Nil(t, v)
		assert.Error(t, err)
	})

	t.Run("valid time", func(t *testing.T) {
		c := Parse("2020-08-05")
		v, err := c.Value()
		assert.Equal(t, c.StdTime(), v)
		assert.Nil(t, err)
	})
}

func TestCarbonType_MarshalJSON(t *testing.T) {
	var model carbonTypeModel

	t.Run("nil time", func(t *testing.T) {
		model.Carbon2 = nil

		data, err := json.Marshal(&model)
		assert.NoError(t, err)
		assert.Equal(t, `{"carbon1":null,"carbon2":null}`, string(data))
	})

	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon()
		model.Carbon1 = c
		model.Carbon2 = &c

		data, err := json.Marshal(&model)
		assert.NoError(t, err)
		assert.Equal(t, `{"carbon1":null,"carbon2":null}`, string(data))
	})

	t.Run("empty time", func(t *testing.T) {
		c := Parse("")
		model.Carbon1 = c
		model.Carbon2 = &c

		data, err := json.Marshal(&model)
		assert.NoError(t, err)
		assert.Equal(t, `{"carbon1":"","carbon2":""}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		c := Parse("xxx")
		model.Carbon1 = c
		model.Carbon2 = &c

		data, err := json.Marshal(&model)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		c := Parse("2020-08-05 13:14:15.999999999")
		model.Carbon1 = c
		model.Carbon2 = &c

		data, err := json.Marshal(&model)
		assert.NoError(t, err)
		assert.Equal(t, `{"carbon1":"2020-08-05 13:14:15","carbon2":"2020-08-05 13:14:15"}`, string(data))
	})
}

func TestCarbonType_UnmarshalJSON(t *testing.T) {
	var model carbonTypeModel

	t.Run("empty value", func(t *testing.T) {
		value := `{"carbon1":"","carbon2":""}`
		assert.NoError(t, json.Unmarshal([]byte(value), &model))

		assert.Empty(t, model.Carbon1.String())
		assert.Empty(t, model.Carbon2.String())
	})

	t.Run("null value", func(t *testing.T) {
		value := `{"carbon1":"null","carbon2":"null"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &model))

		assert.Empty(t, model.Carbon1.String())
		assert.Empty(t, model.Carbon2.String())
	})

	t.Run("valid value", func(t *testing.T) {
		value := `{"carbon1":"2020-08-05 13:14:15","carbon2":"2020-08-05 13:14:15"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &model))

		assert.Equal(t, "2020-08-05 13:14:15", model.Carbon1.String())
		assert.Equal(t, "2020-08-05 13:14:15", model.Carbon2.String())
	})
}

func TestCarbonType_String(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		c := Parse("")
		assert.Empty(t, c.String())
	})

	t.Run("zero time", func(t *testing.T) {
		assert.Empty(t, NewCarbon().String())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").String())
		assert.Empty(t, Parse("0").String())
		assert.Empty(t, Parse("xxx").String())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15", Parse("2020-08-05 13:14:15").String())
		assert.Equal(t, "2020-08-05", Parse("2020-08-05 13:14:15").SetLayout(DateLayout).String())
	})
}

type builtinTypeModel struct {
	Date      Date      `json:"date"`
	DateMilli DateMilli `json:"date_milli"`
	DateMicro DateMicro `json:"date_micro"`
	DateNano  DateNano  `json:"date_nano"`

	Time      Time      `json:"time"`
	TimeMilli TimeMilli `json:"time_milli"`
	TimeMicro TimeMicro `json:"time_micro"`
	TimeNano  TimeNano  `json:"time_nano"`

	DateTime      DateTime      `json:"date_time"`
	DateTimeMilli DateTimeMilli `json:"date_time_milli"`
	DateTimeMicro DateTimeMicro `json:"date_time_micro"`
	DateTimeNano  DateTimeNano  `json:"date_time_nano"`

	CreatedAt *DateTime `json:"created_at"`
	UpdatedAt *DateTime `json:"updated_at"`

	Timestamp      Timestamp      `json:"timestamp"`
	TimestampMilli TimestampMilli `json:"timestamp_milli"`
	TimestampMicro TimestampMicro `json:"timestamp_micro"`
	TimestampNano  TimestampNano  `json:"timestamp_nano"`

	DeletedAt *Timestamp `json:"deleted_at"`
}

func TestBuiltinType_Scan(t *testing.T) {
	c := Now()

	t.Run("[]byte type", func(t *testing.T) {
		dt := NewDateTime(c)
		assert.Nil(t, dt.Scan([]byte(c.ToDateString())))

		ts1 := NewTimestamp(c)
		assert.Error(t, ts1.Scan([]byte("xxx")))
		assert.Nil(t, ts1.Scan([]byte(strconv.Itoa(int(ts1.Timestamp())))))

		ts2 := NewTimestampMilli(c)
		assert.Error(t, ts2.Scan([]byte("xxx")))
		assert.Nil(t, ts2.Scan([]byte(strconv.Itoa(int(ts2.TimestampMilli())))))

		ts3 := NewTimestampMicro(c)
		assert.Error(t, ts3.Scan([]byte("xxx")))
		assert.Nil(t, ts3.Scan([]byte(strconv.Itoa(int(ts3.TimestampMicro())))))

		ts4 := NewTimestampNano(c)
		assert.Error(t, ts4.Scan([]byte("xxx")))
		assert.Nil(t, ts4.Scan([]byte(strconv.Itoa(int(ts3.TimestampNano())))))
	})

	t.Run("string type", func(t *testing.T) {
		dt := NewDateTime(c)
		assert.Nil(t, dt.Scan(c.ToDateString()))

		ts1 := NewTimestamp(c)
		assert.Error(t, ts1.Scan("xxx"))
		assert.Nil(t, ts1.Scan(strconv.Itoa(int(ts1.Timestamp()))))

		ts2 := NewTimestampMilli(c)
		assert.Error(t, ts2.Scan("xxx"))
		assert.Nil(t, ts2.Scan(strconv.Itoa(int(ts2.TimestampMilli()))))

		ts3 := NewTimestampMicro(c)
		assert.Error(t, ts3.Scan("xxx"))
		assert.Nil(t, ts3.Scan(strconv.Itoa(int(ts3.TimestampMicro()))))

		ts4 := NewTimestampNano(c)
		assert.Error(t, ts4.Scan("xxx"))
		assert.Nil(t, ts4.Scan(strconv.Itoa(int(ts4.TimestampNano()))))
	})

	t.Run("int64 type", func(t *testing.T) {
		dt := NewDateTime(c)
		assert.Nil(t, dt.Scan(c.Timestamp()))

		ts1 := NewTimestamp(c)
		assert.Nil(t, ts1.Scan(c.Timestamp()))

		ts2 := NewTimestampMilli(c)
		assert.Nil(t, ts2.Scan(c.TimestampMilli()))

		ts3 := NewTimestampMicro(c)
		assert.Nil(t, ts3.Scan(c.TimestampMicro()))

		ts4 := NewTimestampNano(c)
		assert.Nil(t, ts4.Scan(c.TimestampNano()))
	})

	t.Run("time type", func(t *testing.T) {
		tt := time.Now()

		dt := NewDateTime(c)
		assert.Nil(t, dt.Scan(tt))

		ts1 := NewTimestamp(c)
		assert.Nil(t, ts1.Scan(tt))

		ts2 := NewTimestampMilli(c)
		assert.Nil(t, ts2.Scan(tt))

		ts3 := NewTimestampMicro(c)
		assert.Nil(t, ts3.Scan(tt))

		ts4 := NewTimestampNano(c)
		assert.Nil(t, ts4.Scan(tt))
	})

	t.Run("*time type", func(t *testing.T) {
		tt := time.Now()

		dt := NewDateTime(c)
		assert.Nil(t, dt.Scan(&tt))

		ts1 := NewTimestamp(c)
		assert.Nil(t, ts1.Scan(&tt))

		ts2 := NewTimestampMilli(c)
		assert.Nil(t, ts2.Scan(&tt))

		ts3 := NewTimestampMicro(c)
		assert.Nil(t, ts3.Scan(&tt))

		ts4 := NewTimestampNano(c)
		assert.Nil(t, ts4.Scan(&tt))
	})

	t.Run("nil type", func(t *testing.T) {
		dt := NewDateTime(c)
		assert.Nil(t, dt.Scan(nil))

		ts1 := NewTimestamp(c)
		assert.Nil(t, ts1.Scan(nil))

		ts2 := NewTimestampMilli(c)
		assert.Nil(t, ts2.Scan(nil))

		ts3 := NewTimestampMicro(c)
		assert.Nil(t, ts3.Scan(nil))

		ts4 := NewTimestampNano(c)
		assert.Nil(t, ts4.Scan(nil))
	})

	t.Run("unsupported type", func(t *testing.T) {
		dt := NewDateTime(c)
		assert.Error(t, dt.Scan(true))
		assert.Error(t, dt.Scan(func() {}))
		assert.Error(t, dt.Scan(float64(0)))
		assert.Error(t, dt.Scan(map[string]string{}))

		ts := NewTimestamp(c)
		assert.Error(t, ts.Scan(true))
		assert.Error(t, ts.Scan(func() {}))
		assert.Error(t, ts.Scan(float64(0)))
		assert.Error(t, ts.Scan(map[string]string{}))
	})
}

func TestBuiltinType_Value(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon()

		v1, e1 := NewDateTime(c).Value()
		assert.Nil(t, v1)
		assert.Nil(t, e1)

		v2, e2 := NewTimestamp(c).Value()
		assert.Nil(t, v2)
		assert.Nil(t, e2)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := Parse("xxx")

		v1, e1 := NewDateTime(c).Value()
		assert.Nil(t, v1)
		assert.Error(t, e1)

		v2, e2 := NewTimestamp(c).Value()
		assert.Nil(t, v2)
		assert.Error(t, e2)
	})

	t.Run("valid time", func(t *testing.T) {
		c := Parse("2020-08-05")

		v1, e1 := NewDateTime(c).Value()
		assert.Equal(t, c.StdTime(), v1)
		assert.Nil(t, e1)

		v2, e2 := NewTimestamp(c).Value()
		assert.Equal(t, c.Timestamp(), v2)
		assert.Nil(t, e2)

		v3, e3 := NewTimestampMilli(c).Value()
		assert.Equal(t, c.TimestampMilli(), v3)
		assert.Nil(t, e3)

		v4, e4 := NewTimestampMicro(c).Value()
		assert.Equal(t, c.TimestampMicro(), v4)
		assert.Nil(t, e4)

		v5, e5 := NewTimestampNano(c).Value()
		assert.Equal(t, c.TimestampNano(), v5)
		assert.Nil(t, e5)
	})
}

func TestBuiltinType_MarshalJSON(t *testing.T) {
	var model builtinTypeModel

	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon()

		model.Date = NewDate(c)
		model.DateMilli = NewDateMilli(c)
		model.DateMicro = NewDateMicro(c)
		model.DateNano = NewDateNano(c)

		model.Time = NewTime(c)
		model.TimeMilli = NewTimeMilli(c)
		model.TimeMicro = NewTimeMicro(c)
		model.TimeNano = NewTimeNano(c)

		model.DateTime = NewDateTime(c)
		model.DateTimeMilli = NewDateTimeMilli(c)
		model.DateTimeMicro = NewDateTimeMicro(c)
		model.DateTimeNano = NewDateTimeNano(c)

		model.Timestamp = NewTimestamp(c)
		model.TimestampMilli = NewTimestampMilli(c)
		model.TimestampMicro = NewTimestampMicro(c)
		model.TimestampNano = NewTimestampNano(c)

		createdAt := NewDateTime(c)
		model.CreatedAt = &createdAt
		updatedAt := NewDateTime(c)
		model.UpdatedAt = &updatedAt
		deletedAt := NewTimestamp(c)
		model.DeletedAt = &deletedAt

		data, err := json.Marshal(&model)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":null,"date_milli":null,"date_micro":null,"date_nano":null,"time":null,"time_milli":null,"time_micro":null,"time_nano":null,"date_time":null,"date_time_milli":null,"date_time_micro":null,"date_time_nano":null,"created_at":null,"updated_at":null,"timestamp":null,"timestamp_milli":null,"timestamp_micro":null,"timestamp_nano":null,"deleted_at":null}`, string(data))
	})

	t.Run("empty time", func(t *testing.T) {
		c := Parse("")

		model.Date = NewDate(c)
		model.DateMilli = NewDateMilli(c)
		model.DateMicro = NewDateMicro(c)
		model.DateNano = NewDateNano(c)

		model.Time = NewTime(c)
		model.TimeMilli = NewTimeMilli(c)
		model.TimeMicro = NewTimeMicro(c)
		model.TimeNano = NewTimeNano(c)

		model.DateTime = NewDateTime(c)
		model.DateTimeMilli = NewDateTimeMilli(c)
		model.DateTimeMicro = NewDateTimeMicro(c)
		model.DateTimeNano = NewDateTimeNano(c)

		model.Timestamp = NewTimestamp(c)
		model.TimestampMilli = NewTimestampMilli(c)
		model.TimestampMicro = NewTimestampMicro(c)
		model.TimestampNano = NewTimestampNano(c)

		createdAt := NewDateTime(c)
		model.CreatedAt = &createdAt
		updatedAt := NewDateTime(c)
		model.UpdatedAt = &updatedAt
		deletedAt := NewTimestamp(c)
		model.DeletedAt = &deletedAt

		data, err := json.Marshal(&model)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"","date_milli":"","date_micro":"","date_nano":"","time":"","time_milli":"","time_micro":"","time_nano":"","date_time":"","date_time_milli":"","date_time_micro":"","date_time_nano":"","created_at":"","updated_at":"","timestamp":0,"timestamp_milli":0,"timestamp_micro":0,"timestamp_nano":0,"deleted_at":0}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		c := Parse("xxx")

		var model1 builtinTypeModel

		model1.Date = NewDate(c)
		model1.DateMilli = NewDateMilli(c)
		model1.DateMicro = NewDateMicro(c)
		model1.DateNano = NewDateNano(c)

		model1.Time = NewTime(c)
		model1.TimeMilli = NewTimeMilli(c)
		model1.TimeMicro = NewTimeMicro(c)
		model1.TimeNano = NewTimeNano(c)

		model1.DateTime = NewDateTime(c)
		model1.DateTimeMilli = NewDateTimeMilli(c)
		model1.DateTimeMicro = NewDateTimeMicro(c)
		model1.DateTimeNano = NewDateTimeNano(c)

		model1.DateTime = NewDateTime(c)
		model1.DateTimeMilli = NewDateTimeMilli(c)
		model1.DateTimeMicro = NewDateTimeMicro(c)
		model1.DateTimeNano = NewDateTimeNano(c)

		createdAt := NewDateTime(c)
		model1.CreatedAt = &createdAt
		updatedAt := NewDateTime(c)
		model1.UpdatedAt = &updatedAt

		data1, err1 := json.Marshal(&model1)
		assert.Error(t, err1)
		assert.Empty(t, string(data1))

		var model2 builtinTypeModel

		model2.Timestamp = NewTimestamp(c)
		model2.TimestampMilli = NewTimestampMilli(c)
		model2.TimestampMicro = NewTimestampMicro(c)
		model2.TimestampNano = NewTimestampNano(c)

		deletedAt := NewTimestamp(c)
		model2.DeletedAt = &deletedAt

		data2, err2 := json.Marshal(&model2)
		assert.Error(t, err2)
		assert.Empty(t, string(data2))
	})

	t.Run("valid time", func(t *testing.T) {
		c := Parse("2020-08-05 13:14:15.999999999")

		model.Date = NewDate(c)
		model.DateMilli = NewDateMilli(c)
		model.DateMicro = NewDateMicro(c)
		model.DateNano = NewDateNano(c)

		model.Time = NewTime(c)
		model.TimeMilli = NewTimeMilli(c)
		model.TimeMicro = NewTimeMicro(c)
		model.TimeNano = NewTimeNano(c)

		model.DateTime = NewDateTime(c)
		model.DateTimeMilli = NewDateTimeMilli(c)
		model.DateTimeMicro = NewDateTimeMicro(c)
		model.DateTimeNano = NewDateTimeNano(c)

		model.Timestamp = NewTimestamp(c)
		model.TimestampMilli = NewTimestampMilli(c)
		model.TimestampMicro = NewTimestampMicro(c)
		model.TimestampNano = NewTimestampNano(c)

		createdAt := NewDateTime(c)
		model.CreatedAt = &createdAt
		updatedAt := NewDateTime(c)
		model.UpdatedAt = &updatedAt
		deletedAt := NewTimestamp(c)
		model.DeletedAt = &deletedAt

		data, err := json.Marshal(&model)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"2020-08-05","date_milli":"2020-08-05.999","date_micro":"2020-08-05.999999","date_nano":"2020-08-05.999999999","time":"13:14:15","time_milli":"13:14:15.999","time_micro":"13:14:15.999999","time_nano":"13:14:15.999999999","date_time":"2020-08-05 13:14:15","date_time_milli":"2020-08-05 13:14:15.999","date_time_micro":"2020-08-05 13:14:15.999999","date_time_nano":"2020-08-05 13:14:15.999999999","created_at":"2020-08-05 13:14:15","updated_at":"2020-08-05 13:14:15","timestamp":1596633255,"timestamp_milli":1596633255999,"timestamp_micro":1596633255999999,"timestamp_nano":1596633255999999999,"deleted_at":1596633255}`, string(data))
	})
}

func TestBuiltinType_UnmarshalJSON(t *testing.T) {
	var model builtinTypeModel

	t.Run("empty value", func(t *testing.T) {
		value := `{"date":"","date_milli":"","date_micro":"","date_nano":"","time":"","time_milli":"","time_micro":"","time_nano":"","date_time":"","date_time_milli":"","date_time_micro":"","date_time_nano":"","created_at":"","updated_at":"","timestamp":"","timestamp_milli":"","timestamp_micro":"","timestamp_nano":"","deleted_at":""}`
		assert.NoError(t, json.Unmarshal([]byte(value), &model))

		assert.Empty(t, model.Date.String())
		assert.Empty(t, model.DateMilli.String())
		assert.Empty(t, model.DateMicro.String())
		assert.Empty(t, model.DateNano.String())

		assert.Empty(t, model.Time.String())
		assert.Empty(t, model.TimeMilli.String())
		assert.Empty(t, model.TimeMicro.String())
		assert.Empty(t, model.TimeNano.String())

		assert.Empty(t, model.DateTime.String())
		assert.Empty(t, model.DateTimeMilli.String())
		assert.Empty(t, model.DateTimeMicro.String())
		assert.Empty(t, model.DateTimeNano.String())

		assert.Equal(t, "0", model.Timestamp.String())
		assert.Equal(t, "0", model.TimestampMilli.String())
		assert.Equal(t, "0", model.TimestampMicro.String())
		assert.Equal(t, "0", model.TimestampNano.String())

		assert.Zero(t, model.Timestamp.Int64())
		assert.Zero(t, model.TimestampMilli.Int64())
		assert.Zero(t, model.TimestampMicro.Int64())
		assert.Zero(t, model.TimestampNano.Int64())

		assert.Empty(t, model.CreatedAt.String())
		assert.Empty(t, model.UpdatedAt.String())
		assert.Equal(t, "0", model.DeletedAt.String())
		assert.Zero(t, model.DeletedAt.Int64())
	})

	t.Run("null value", func(t *testing.T) {
		value := `{"date":"null","date_milli":"null","date_micro":"null","date_nano":"null","time":"null","time_milli":"null","time_micro":"null","time_nano":"null","date_time":"null","date_time_milli":"null","date_time_micro":"null","date_time_nano":"null","created_at":"null","updated_at":"null","timestamp":"null","timestamp_milli":"null","timestamp_micro":"null","timestamp_nano":"null","deleted_at":"null"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &model))

		assert.Empty(t, model.Date.String())
		assert.Empty(t, model.DateMilli.String())
		assert.Empty(t, model.DateMicro.String())
		assert.Empty(t, model.DateNano.String())

		assert.Empty(t, model.Time.String())
		assert.Empty(t, model.TimeMilli.String())
		assert.Empty(t, model.TimeMicro.String())
		assert.Empty(t, model.TimeNano.String())

		assert.Empty(t, model.DateTime.String())
		assert.Empty(t, model.DateTimeMilli.String())
		assert.Empty(t, model.DateTimeMicro.String())
		assert.Empty(t, model.DateTimeNano.String())

		assert.Equal(t, "0", model.Timestamp.String())
		assert.Equal(t, "0", model.TimestampMilli.String())
		assert.Equal(t, "0", model.TimestampMicro.String())
		assert.Equal(t, "0", model.TimestampNano.String())

		assert.Zero(t, model.Timestamp.Int64())
		assert.Zero(t, model.TimestampMilli.Int64())
		assert.Zero(t, model.TimestampMicro.Int64())
		assert.Zero(t, model.TimestampNano.Int64())

		assert.Empty(t, model.CreatedAt.String())
		assert.Empty(t, model.UpdatedAt.String())
		assert.Equal(t, "0", model.DeletedAt.String())
		assert.Zero(t, model.DeletedAt.Int64())
	})

	t.Run("zero value", func(t *testing.T) {
		value := `{"date":"0","date_milli":"0","date_micro":"0","date_nano":"0","time":"0","time_milli":"0","time_micro":"0","time_nano":"0","date_time":"0","date_time_milli":"0","date_time_micro":"0","date_time_nano":"0","created_at":"0","updated_at":"0","timestamp":"0","timestamp_milli":"0","timestamp_micro":"0","timestamp_nano":"0","deleted_at":"0"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &model))

		assert.Empty(t, model.Date.String())
		assert.Empty(t, model.DateMilli.String())
		assert.Empty(t, model.DateMicro.String())
		assert.Empty(t, model.DateNano.String())

		assert.Empty(t, model.Time.String())
		assert.Empty(t, model.TimeMilli.String())
		assert.Empty(t, model.TimeMicro.String())
		assert.Empty(t, model.TimeNano.String())

		assert.Empty(t, model.DateTime.String())
		assert.Empty(t, model.DateTimeMilli.String())
		assert.Empty(t, model.DateTimeMicro.String())
		assert.Empty(t, model.DateTimeNano.String())

		assert.Equal(t, "0", model.Timestamp.String())
		assert.Equal(t, "0", model.TimestampMilli.String())
		assert.Equal(t, "0", model.TimestampMicro.String())
		assert.Equal(t, "0", model.TimestampNano.String())

		assert.Zero(t, model.Timestamp.Int64())
		assert.Zero(t, model.TimestampMilli.Int64())
		assert.Zero(t, model.TimestampMicro.Int64())
		assert.Zero(t, model.TimestampNano.Int64())

		assert.Empty(t, model.CreatedAt.String())
		assert.Empty(t, model.UpdatedAt.String())
		assert.Equal(t, "0", model.DeletedAt.String())
		assert.Zero(t, model.DeletedAt.Int64())
	})

	t.Run("invalid value", func(t *testing.T) {
		var model1 builtinTypeModel

		value1 := `{"date":"xxx","date_milli":"xxx","date_micro":"xxx","date_nano":"xxx","time":"xxx","time_milli":"xxx","time_micro":"xxx","time_nano":"xxx","date_time":"xxx","date_time_milli":"xxx","date_time_micro":"xxx","date_time_nano":"xxx","created_at":"xxx","updated_at":"xxx","timestamp":"xxx"}`
		assert.Error(t, json.Unmarshal([]byte(value1), &model1))

		var model2 builtinTypeModel

		value2 := `{"timestamp":"xxx","timestamp_milli":"xxx","timestamp_micro":"xxx","timestamp_nano":"xxx","deleted_at":"xxx"}`
		assert.Error(t, json.Unmarshal([]byte(value2), &model2))

		assert.Empty(t, model.Date.String())
		assert.Empty(t, model.DateMilli.String())
		assert.Empty(t, model.DateMicro.String())
		assert.Empty(t, model.DateNano.String())

		assert.Empty(t, model.Time.String())
		assert.Empty(t, model.TimeMilli.String())
		assert.Empty(t, model.TimeMicro.String())
		assert.Empty(t, model.TimeNano.String())

		assert.Empty(t, model.DateTime.String())
		assert.Empty(t, model.DateTimeMilli.String())
		assert.Empty(t, model.DateTimeMicro.String())
		assert.Empty(t, model.DateTimeNano.String())

		assert.Equal(t, "0", model.Timestamp.String())
		assert.Equal(t, "0", model.TimestampMilli.String())
		assert.Equal(t, "0", model.TimestampMicro.String())
		assert.Equal(t, "0", model.TimestampNano.String())

		assert.Zero(t, model.Timestamp.Int64())
		assert.Zero(t, model.TimestampMilli.Int64())
		assert.Zero(t, model.TimestampMicro.Int64())
		assert.Zero(t, model.TimestampNano.Int64())

		assert.Empty(t, model.CreatedAt.String())
		assert.Empty(t, model.UpdatedAt.String())
		assert.Equal(t, "0", model.DeletedAt.String())
		assert.Zero(t, model.DeletedAt.Int64())
	})

	t.Run("valid value", func(t *testing.T) {
		value := `{"date":"2020-08-05","date_milli":"2020-08-05.999","date_micro":"2020-08-05.999999","date_nano":"2020-08-05.999999999","time":"13:14:15","time_milli":"13:14:15.999","time_micro":"13:14:15.999999","time_nano":"13:14:15.999999999","date_time":"2020-08-05 13:14:15","date_time_milli":"2020-08-05 13:14:15.999","date_time_micro":"2020-08-05 13:14:15.999999","date_time_nano":"2020-08-05 13:14:15.999999999","created_at":"2020-08-05 13:14:15","updated_at":"2020-08-05 13:14:15","timestamp":1596633255,"timestamp_milli":1596633255999,"timestamp_micro":1596633255999999,"timestamp_nano":1596633255999999999,"deleted_at":1596633255}`
		assert.NoError(t, json.Unmarshal([]byte(value), &model))

		assert.Equal(t, "2020-08-05", model.Date.String())
		assert.Equal(t, "2020-08-05.999", model.DateMilli.String())
		assert.Equal(t, "2020-08-05.999999", model.DateMicro.String())
		assert.Equal(t, "2020-08-05.999999999", model.DateNano.String())

		assert.Equal(t, "13:14:15", model.Time.String())
		assert.Equal(t, "13:14:15.999", model.TimeMilli.String())
		assert.Equal(t, "13:14:15.999999", model.TimeMicro.String())
		assert.Equal(t, "13:14:15.999999999", model.TimeNano.String())

		assert.Equal(t, "2020-08-05 13:14:15", model.DateTime.String())
		assert.Equal(t, "2020-08-05 13:14:15.999", model.DateTimeMilli.String())
		assert.Equal(t, "2020-08-05 13:14:15.999999", model.DateTimeMicro.String())
		assert.Equal(t, "2020-08-05 13:14:15.999999999", model.DateTimeNano.String())

		assert.Equal(t, "1596633255", model.Timestamp.String())
		assert.Equal(t, "1596633255999", model.TimestampMilli.String())
		assert.Equal(t, "1596633255999999", model.TimestampMicro.String())
		assert.Equal(t, "1596633255999999999", model.TimestampNano.String())

		assert.Equal(t, int64(1596633255), model.Timestamp.Int64())
		assert.Equal(t, int64(1596633255999), model.TimestampMilli.Int64())
		assert.Equal(t, int64(1596633255999999), model.TimestampMicro.Int64())
		assert.Equal(t, int64(1596633255999999999), model.TimestampNano.Int64())

		assert.Equal(t, "2020-08-05 13:14:15", model.CreatedAt.String())
		assert.Equal(t, "2020-08-05 13:14:15", model.UpdatedAt.String())
		assert.Equal(t, "1596633255", model.DeletedAt.String())
		assert.Equal(t, int64(1596633255), model.DeletedAt.Int64())
	})
}

type iso8601Type string

func (t iso8601Type) Format() string {
	return ISO8601Format
}

type rfc3339Type string

func (t rfc3339Type) Layout() string {
	return RFC3339Layout
}

type CustomerTypeModel struct {
	Customer1 FormatType[iso8601Type] `json:"customer1"`
	Customer2 LayoutType[rfc3339Type] `json:"customer2"`

	CreatedAt *FormatType[iso8601Type] `json:"created_at"`
	UpdatedAt *LayoutType[rfc3339Type] `json:"updated_at"`
}

func TestCustomerType_Scan(t *testing.T) {
	t1 := NewFormatType[iso8601Type](Now())
	t2 := NewLayoutType[rfc3339Type](Now())

	t.Run("[]byte type", func(t *testing.T) {
		assert.Nil(t, t1.Scan([]byte(Now().ToDateString())))
		assert.Nil(t, t2.Scan([]byte(Now().ToDateString())))
	})

	t.Run("string type", func(t *testing.T) {
		assert.Nil(t, t1.Scan(Now().ToDateString()))
		assert.Nil(t, t2.Scan(Now().ToDateString()))
	})

	t.Run("int64 type", func(t *testing.T) {
		assert.Nil(t, t1.Scan(Now().Timestamp()))
		assert.Nil(t, t2.Scan(Now().Timestamp()))
	})

	t.Run("time type", func(t *testing.T) {
		tt := time.Now()

		assert.Nil(t, t1.Scan(tt))
		assert.Nil(t, t2.Scan(tt))
	})

	t.Run("*time type", func(t *testing.T) {
		tt := time.Now()

		assert.Nil(t, t1.Scan(&tt))
		assert.Nil(t, t2.Scan(&tt))
	})

	t.Run("nil type", func(t *testing.T) {
		assert.Nil(t, t1.Scan(nil))
		assert.Nil(t, t2.Scan(nil))
	})

	t.Run("unsupported type", func(t *testing.T) {
		assert.Error(t, t1.Scan(true))
		assert.Error(t, t1.Scan(func() {}))
		assert.Error(t, t1.Scan(float64(0)))
		assert.Error(t, t1.Scan(map[string]string{}))

		assert.Error(t, t2.Scan(true))
		assert.Error(t, t2.Scan(func() {}))
		assert.Error(t, t2.Scan(float64(0)))
		assert.Error(t, t2.Scan(map[string]string{}))
	})
}

func TestCustomerType_Value(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon()

		t1, e1 := NewFormatType[iso8601Type](c).Value()
		assert.Nil(t, t1)
		assert.Nil(t, e1)

		t2, e2 := NewLayoutType[rfc3339Type](c).Value()
		assert.Nil(t, t2)
		assert.Nil(t, e2)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := Parse("xxx")

		t1, e1 := NewFormatType[iso8601Type](c).Value()
		assert.Nil(t, t1)
		assert.Error(t, e1)

		t2, e2 := NewLayoutType[rfc3339Type](c).Value()
		assert.Nil(t, t2)
		assert.Error(t, e2)
	})

	t.Run("valid time", func(t *testing.T) {
		c := Parse("2020-08-05")

		t1, e1 := NewFormatType[iso8601Type](c).Value()
		assert.Equal(t, c.StdTime(), t1)
		assert.Nil(t, e1)

		t2, e2 := NewLayoutType[rfc3339Type](c).Value()
		assert.Equal(t, c.StdTime(), t2)
		assert.Nil(t, e2)
	})
}

func TestCustomerType_MarshalJSON(t *testing.T) {
	var model CustomerTypeModel

	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon()

		model.Customer1 = NewFormatType[iso8601Type](c)
		model.Customer2 = NewLayoutType[rfc3339Type](c)

		createdAt := NewFormatType[iso8601Type](c)
		model.CreatedAt = &createdAt
		updatedAt := NewLayoutType[rfc3339Type](c)
		model.UpdatedAt = &updatedAt

		data, err := json.Marshal(&model)
		assert.NoError(t, err)
		assert.Equal(t, `{"customer1":null,"customer2":null,"created_at":null,"updated_at":null}`, string(data))
	})

	t.Run("empty time", func(t *testing.T) {
		c := Parse("")

		model.Customer1 = NewFormatType[iso8601Type](c)
		model.Customer2 = NewLayoutType[rfc3339Type](c)

		createdAt := NewFormatType[iso8601Type](c)
		model.CreatedAt = &createdAt
		updatedAt := NewLayoutType[rfc3339Type](c)
		model.UpdatedAt = &updatedAt

		data, err := json.Marshal(&model)
		assert.NoError(t, err)
		assert.Equal(t, `{"customer1":"","customer2":"","created_at":"","updated_at":""}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		c := Parse("xxx")

		model.Customer1 = NewFormatType[iso8601Type](c)
		model.Customer2 = NewLayoutType[rfc3339Type](c)

		createdAt := NewFormatType[iso8601Type](c)
		model.CreatedAt = &createdAt
		updatedAt := NewLayoutType[rfc3339Type](c)
		model.UpdatedAt = &updatedAt

		data, err := json.Marshal(&model)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		c := Parse("2020-08-05 13:14:15.999999999")

		model.Customer1 = NewFormatType[iso8601Type](c)
		model.Customer2 = NewLayoutType[rfc3339Type](c)

		createdAt := NewFormatType[iso8601Type](c)
		model.CreatedAt = &createdAt
		updatedAt := NewLayoutType[rfc3339Type](c)
		model.UpdatedAt = &updatedAt

		data, err := json.Marshal(&model)
		assert.NoError(t, err)
		assert.Equal(t, `{"customer1":"2020-08-05T13:14:15+00:00","customer2":"2020-08-05T13:14:15Z","created_at":"2020-08-05T13:14:15+00:00","updated_at":"2020-08-05T13:14:15Z"}`, string(data))
	})
}

func TestCustomerType_UnmarshalJSON(t *testing.T) {
	var model CustomerTypeModel

	t.Run("empty value", func(t *testing.T) {
		value := `{"customer1":"","customer2":"","created_at":"","updated_at":""}`
		assert.NoError(t, json.Unmarshal([]byte(value), &model))

		assert.Empty(t, model.Customer1.String())
		assert.Empty(t, model.Customer2.String())
		assert.Empty(t, model.CreatedAt.String())
		assert.Empty(t, model.UpdatedAt.String())
	})

	t.Run("null value", func(t *testing.T) {
		value := `{"customer1":"null","customer2":"null","created_at":"null","updated_at":"null"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &model))

		assert.Empty(t, model.Customer1.String())
		assert.Empty(t, model.Customer2.String())
		assert.Empty(t, model.CreatedAt.String())
		assert.Empty(t, model.UpdatedAt.String())
	})

	t.Run("invalid value", func(t *testing.T) {
		value := `{"customer1":"xxx","customer2":"xxx","created_at":"xxx","updated_at":"xxx"}`
		assert.Error(t, json.Unmarshal([]byte(value), &model))

		assert.Empty(t, model.Customer1.String())
		assert.Empty(t, model.Customer2.String())
		assert.Empty(t, model.CreatedAt.String())
		assert.Empty(t, model.UpdatedAt.String())
	})

	t.Run("valid value", func(t *testing.T) {
		value := `{"customer1":"2020-08-05T13:14:15+00:00","customer2":"2020-08-05T13:14:15Z","created_at":"2020-08-05T13:14:15+00:00","updated_at":"2020-08-05T13:14:15Z"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &model))

		assert.Equal(t, "2020-08-05T13:14:15+00:00", model.Customer1.String())
		assert.Equal(t, "2020-08-05T13:14:15Z", model.Customer2.String())
		assert.Equal(t, "2020-08-05T13:14:15+00:00", model.CreatedAt.String())
		assert.Equal(t, "2020-08-05T13:14:15Z", model.UpdatedAt.String())
	})
}
