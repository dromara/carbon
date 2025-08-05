package carbon

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type carbonTypeModel struct {
	Carbon1 Carbon  `json:"carbon1"`
	Carbon2 *Carbon `json:"carbon2"`
}

type CarbonTypeSuite struct {
	suite.Suite
}

func TestCarbonTypeSuite(t *testing.T) {
	suite.Run(t, new(CarbonTypeSuite))
}

func (s *CarbonTypeSuite) TestCarbonType_Scan() {
	c := Now()

	s.Run("[]byte type", func() {
		s.Nil(c.Scan([]byte(c.ToDateString())))
	})

	s.Run("string type", func() {
		s.Nil(c.Scan(c.ToDateString()))
	})

	s.Run("time type", func() {
		tt := time.Now()
		s.Nil(c.Scan(tt))
	})

	s.Run("*time type", func() {
		tt := time.Now()
		s.Nil(c.Scan(&tt))
	})

	s.Run("nil type", func() {
		s.Nil(c.Scan(nil))
	})

	s.Run("unsupported type", func() {
		s.Error(c.Scan(true))
		s.Error(c.Scan(int64(0)))
		s.Error(c.Scan(func() {}))
		s.Error(c.Scan(float64(0)))
		s.Error(c.Scan(map[string]string{}))
	})
}

func (s *CarbonTypeSuite) TestCarbonType_Value() {
	s.Run("zero carbon", func() {
		v, e := NewCarbon().Value()
		s.Nil(v)
		s.Nil(e)
	})

	s.Run("error carbon", func() {
		v, e := Parse("xxx").Value()
		s.Nil(v)
		s.Error(e)
	})

	s.Run("empty carbon", func() {
		v, e := Parse("").Value()
		s.Empty(v)
		s.Nil(e)
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05")
		v, e := c.Value()
		s.Equal(c.StdTime(), v)
		s.Nil(e)
	})
}

func (s *CarbonTypeSuite) TestCarbonType_MarshalJSON() {
	var model carbonTypeModel

	s.Run("unset carbon", func() {
		v, e := json.Marshal(&model)
		s.Nil(e)
		s.Equal(`{"carbon1":null,"carbon2":null}`, string(v))
	})

	s.Run("nil carbon", func() {
		model.Carbon2 = nil

		v, e := json.Marshal(&model)
		s.Nil(e)
		s.Equal(`{"carbon1":null,"carbon2":null}`, string(v))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		model.Carbon1 = *c
		model.Carbon2 = c

		v, e := json.Marshal(&model)
		s.Nil(e)
		s.Equal(`{"carbon1":null,"carbon2":null}`, string(v))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		model.Carbon1 = *c
		model.Carbon2 = c

		v, e := json.Marshal(&model)
		s.Nil(e)
		s.Equal(`{"carbon1":null,"carbon2":null}`, string(v))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		model.Carbon1 = *c
		model.Carbon2 = c

		v, e := json.Marshal(&model)
		s.Error(e)
		s.Empty(string(v))
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15.999999999")
		model.Carbon1 = *c
		model.Carbon2 = c

		v, e := json.Marshal(&model)
		s.Nil(e)
		s.Equal(`{"carbon1":"2020-08-05 13:14:15","carbon2":"2020-08-05 13:14:15"}`, string(v))
	})
}

func (s *CarbonTypeSuite) TestCarbonType_UnmarshalJSON() {
	var model carbonTypeModel

	s.Run("empty value", func() {
		value := `{"carbon1":"","carbon2":""}`
		s.Nil(json.Unmarshal([]byte(value), &model))

		s.Empty(model.Carbon1.String())
		s.Empty(model.Carbon2.String())
		s.False(model.Carbon1.IsValid())
		s.False(model.Carbon2.IsValid())
	})

	s.Run("null value", func() {
		value1 := `{"carbon1":null,"carbon2":null}`
		s.Nil(json.Unmarshal([]byte(value1), &model))
		s.Empty(model.Carbon1.String())
		s.Empty(model.Carbon2.String())

		value2 := `{"carbon1":"null","carbon2":"null"}`
		s.Nil(json.Unmarshal([]byte(value2), &model))
		s.Empty(model.Carbon1.String())
		s.Empty(model.Carbon2.String())
	})

	s.Run("valid value", func() {
		value := `{"carbon1":"2020-08-05 13:14:15","carbon2":"2020-08-05 13:14:15"}`
		s.Nil(json.Unmarshal([]byte(value), &model))

		s.Equal("2020-08-05 13:14:15", model.Carbon1.String())
		s.Equal("2020-08-05 13:14:15", model.Carbon2.String())
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

type BuiltinTypeSuite struct {
	suite.Suite
}

func TestBuiltinTypeSuite(t *testing.T) {
	suite.Run(t, new(BuiltinTypeSuite))
}

func (s *BuiltinTypeSuite) TestBuiltinType_Scan() {
	c := Now()

	s.Run("[]byte type", func() {
		s.Nil(NewDateTime(c).Scan([]byte(c.ToDateString())))

		ts1 := NewTimestamp(c)
		s.Error(ts1.Scan([]byte("xxx")))

		ts2 := NewTimestampMilli(c)
		s.Error(ts2.Scan([]byte("xxx")))

		ts3 := NewTimestampMicro(c)
		s.Error(ts3.Scan([]byte("xxx")))

		ts4 := NewTimestampNano(c)
		s.Error(ts4.Scan([]byte("xxx")))
	})

	s.Run("string type", func() {
		s.Nil(NewDateTime(c).Scan(c.ToDateString()))

		ts1 := NewTimestamp(c)
		s.Error(ts1.Scan("xxx"))

		ts2 := NewTimestampMilli(c)
		s.Error(ts2.Scan("xxx"))

		ts3 := NewTimestampMicro(c)
		s.Error(ts3.Scan("xxx"))

		ts4 := NewTimestampNano(c)
		s.Error(ts4.Scan("xxx"))
	})

	s.Run("time type", func() {
		tt := time.Now()
		s.Nil(NewDateTime(c).Scan(tt))

		ts1 := NewTimestamp(c)
		s.Nil(ts1.Scan(tt))

		ts2 := NewTimestampMilli(c)
		s.Nil(ts2.Scan(tt))

		ts3 := NewTimestampMicro(c)
		s.Nil(ts3.Scan(tt))

		ts4 := NewTimestampNano(c)
		s.Nil(ts4.Scan(tt))
	})

	s.Run("*time type", func() {
		tt := time.Now()
		s.Nil(NewDateTime(c).Scan(&tt))

		ts1 := NewTimestamp(c)
		s.Nil(ts1.Scan(&tt))

		ts2 := NewTimestampMilli(c)
		s.Nil(ts2.Scan(&tt))

		ts3 := NewTimestampMicro(c)
		s.Nil(ts3.Scan(&tt))

		ts4 := NewTimestampNano(c)
		s.Nil(ts4.Scan(&tt))
	})

	s.Run("nil type", func() {
		s.Nil(NewDateTime(c).Scan(nil))

		ts1 := NewTimestamp(c)
		s.Nil(ts1.Scan(nil))

		ts2 := NewTimestampMilli(c)
		s.Nil(ts2.Scan(nil))

		ts3 := NewTimestampMicro(c)
		s.Nil(ts3.Scan(nil))

		ts4 := NewTimestampNano(c)
		s.Nil(ts4.Scan(nil))
	})

	s.Run("unsupported type", func() {
		s.Error(NewDateTime(c).Scan(true))
		s.Error(NewDateTime(c).Scan(func() {}))
		s.Error(NewDateTime(c).Scan(int64(0)))
		s.Error(NewDateTime(c).Scan(float64(0)))
		s.Error(NewDateTime(c).Scan(map[string]string{}))

		s.Error(NewTimestamp(c).Scan(true))
		s.Error(NewTimestamp(c).Scan(func() {}))
		s.Error(NewTimestamp(c).Scan(int64(0)))
		s.Error(NewTimestamp(c).Scan(float64(0)))
		s.Error(NewTimestamp(c).Scan(map[string]string{}))
	})
}

func (s *BuiltinTypeSuite) TestBuiltinType_Value() {
	s.Run("nil carbon", func() {
		v1, e1 := NewDateTime(nil).Value()
		s.Nil(v1)
		s.Nil(e1)

		v2, e2 := NewTimestamp(nil).Value()
		s.Nil(v2)
		s.Nil(e2)
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()

		v1, e1 := NewDateTime(c).Value()
		s.Nil(v1)
		s.Nil(e1)

		v2, e2 := NewTimestamp(c).Value()
		s.Nil(v2)
		s.Nil(e2)
	})

	s.Run("empty carbon", func() {
		c := Parse("")

		v1, e1 := NewDateTime(c).Value()
		s.Empty(v1)
		s.Nil(e1)

		v2, e2 := NewTimestamp(c).Value()
		s.Empty(v2)
		s.Nil(e2)
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")

		v1, e1 := NewDateTime(c).Value()
		s.Nil(v1)
		s.Error(e1)

		v2, e2 := NewTimestamp(c).Value()
		s.Nil(v2)
		s.Error(e2)
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05")

		v1, e1 := NewDateTime(c).Value()
		s.Equal(c.StdTime(), v1)
		s.Nil(e1)

		v2, e2 := NewTimestamp(c).Value()
		s.Equal(c.StdTime(), v2)
		s.Nil(e2)

		v3, e3 := NewTimestampMilli(c).Value()
		s.Equal(c.StdTime(), v3)
		s.Nil(e3)

		v4, e4 := NewTimestampMicro(c).Value()
		s.Equal(c.StdTime(), v4)
		s.Nil(e4)

		v5, e5 := NewTimestampNano(c).Value()
		s.Equal(c.StdTime(), v5)
		s.Nil(e5)
	})
}

func (s *BuiltinTypeSuite) TestBuiltinType_MarshalJSON() {
	var model builtinTypeModel

	s.Run("unset carbon", func() {
		v, e := json.Marshal(&model)
		s.Nil(e)
		s.Equal(`{"date":null,"date_milli":null,"date_micro":null,"date_nano":null,"time":null,"time_milli":null,"time_micro":null,"time_nano":null,"date_time":null,"date_time_milli":null,"date_time_micro":null,"date_time_nano":null,"created_at":null,"updated_at":null,"timestamp":null,"timestamp_milli":null,"timestamp_micro":null,"timestamp_nano":null,"deleted_at":null}`, string(v))
	})

	s.Run("nil carbon", func() {
		model.Date = *NewDate(nil)
		model.DateMilli = *NewDateMilli(nil)
		model.DateMicro = *NewDateMicro(nil)
		model.DateNano = *NewDateNano(nil)

		model.Time = *NewTime(nil)
		model.TimeMilli = *NewTimeMilli(nil)
		model.TimeMicro = *NewTimeMicro(nil)
		model.TimeNano = *NewTimeNano(nil)

		model.DateTime = *NewDateTime(nil)
		model.DateTimeMilli = *NewDateTimeMilli(nil)
		model.DateTimeMicro = *NewDateTimeMicro(nil)
		model.DateTimeNano = *NewDateTimeNano(nil)

		model.Timestamp = *NewTimestamp(nil)
		model.TimestampMilli = *NewTimestampMilli(nil)
		model.TimestampMicro = *NewTimestampMicro(nil)
		model.TimestampNano = *NewTimestampNano(nil)

		model.CreatedAt = NewDateTime(nil)
		model.UpdatedAt = NewDateTime(nil)
		model.DeletedAt = NewTimestamp(nil)

		v, e := json.Marshal(&model)
		s.Nil(e)
		s.Equal(`{"date":null,"date_milli":null,"date_micro":null,"date_nano":null,"time":null,"time_milli":null,"time_micro":null,"time_nano":null,"date_time":null,"date_time_milli":null,"date_time_micro":null,"date_time_nano":null,"created_at":null,"updated_at":null,"timestamp":null,"timestamp_milli":null,"timestamp_micro":null,"timestamp_nano":null,"deleted_at":null}`, string(v))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()

		model.Date = *NewDate(c)
		model.DateMilli = *NewDateMilli(c)
		model.DateMicro = *NewDateMicro(c)
		model.DateNano = *NewDateNano(c)

		model.Time = *NewTime(c)
		model.TimeMilli = *NewTimeMilli(c)
		model.TimeMicro = *NewTimeMicro(c)
		model.TimeNano = *NewTimeNano(c)

		model.DateTime = *NewDateTime(c)
		model.DateTimeMilli = *NewDateTimeMilli(c)
		model.DateTimeMicro = *NewDateTimeMicro(c)
		model.DateTimeNano = *NewDateTimeNano(c)

		model.Timestamp = *NewTimestamp(c)
		model.TimestampMilli = *NewTimestampMilli(c)
		model.TimestampMicro = *NewTimestampMicro(c)
		model.TimestampNano = *NewTimestampNano(c)

		model.CreatedAt = NewDateTime(c)
		model.UpdatedAt = NewDateTime(c)
		model.DeletedAt = NewTimestamp(c)

		v, e := json.Marshal(&model)
		s.Nil(e)
		s.Equal(`{"date":null,"date_milli":null,"date_micro":null,"date_nano":null,"time":null,"time_milli":null,"time_micro":null,"time_nano":null,"date_time":null,"date_time_milli":null,"date_time_micro":null,"date_time_nano":null,"created_at":null,"updated_at":null,"timestamp":null,"timestamp_milli":null,"timestamp_micro":null,"timestamp_nano":null,"deleted_at":null}`, string(v))
	})

	s.Run("empty carbon", func() {
		c := Parse("")

		model.Date = *NewDate(c)
		model.DateMilli = *NewDateMilli(c)
		model.DateMicro = *NewDateMicro(c)
		model.DateNano = *NewDateNano(c)

		model.Time = *NewTime(c)
		model.TimeMilli = *NewTimeMilli(c)
		model.TimeMicro = *NewTimeMicro(c)
		model.TimeNano = *NewTimeNano(c)

		model.DateTime = *NewDateTime(c)
		model.DateTimeMilli = *NewDateTimeMilli(c)
		model.DateTimeMicro = *NewDateTimeMicro(c)
		model.DateTimeNano = *NewDateTimeNano(c)

		model.Timestamp = *NewTimestamp(c)
		model.TimestampMilli = *NewTimestampMilli(c)
		model.TimestampMicro = *NewTimestampMicro(c)
		model.TimestampNano = *NewTimestampNano(c)

		model.CreatedAt = NewDateTime(c)
		model.UpdatedAt = NewDateTime(c)
		model.DeletedAt = NewTimestamp(c)

		v, e := json.Marshal(&model)
		s.Nil(e)
		s.Equal(`{"date":null,"date_milli":null,"date_micro":null,"date_nano":null,"time":null,"time_milli":null,"time_micro":null,"time_nano":null,"date_time":null,"date_time_milli":null,"date_time_micro":null,"date_time_nano":null,"created_at":null,"updated_at":null,"timestamp":null,"timestamp_milli":null,"timestamp_micro":null,"timestamp_nano":null,"deleted_at":null}`, string(v))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")

		var model1 builtinTypeModel

		model1.Date = *NewDate(c)
		model1.DateMilli = *NewDateMilli(c)
		model1.DateMicro = *NewDateMicro(c)
		model1.DateNano = *NewDateNano(c)

		model1.Time = *NewTime(c)
		model1.TimeMilli = *NewTimeMilli(c)
		model1.TimeMicro = *NewTimeMicro(c)
		model1.TimeNano = *NewTimeNano(c)

		model1.DateTime = *NewDateTime(c)
		model1.DateTimeMilli = *NewDateTimeMilli(c)
		model1.DateTimeMicro = *NewDateTimeMicro(c)
		model1.DateTimeNano = *NewDateTimeNano(c)

		model1.DateTime = *NewDateTime(c)
		model1.DateTimeMilli = *NewDateTimeMilli(c)
		model1.DateTimeMicro = *NewDateTimeMicro(c)
		model1.DateTimeNano = *NewDateTimeNano(c)

		model1.CreatedAt = NewDateTime(c)
		model1.UpdatedAt = NewDateTime(c)

		v1, e1 := json.Marshal(&model1)
		s.Error(e1)
		s.Empty(string(v1))

		var model2 builtinTypeModel

		model2.Timestamp = *NewTimestamp(c)
		model2.TimestampMilli = *NewTimestampMilli(c)
		model2.TimestampMicro = *NewTimestampMicro(c)
		model2.TimestampNano = *NewTimestampNano(c)

		model2.DeletedAt = NewTimestamp(c)

		v2, e2 := json.Marshal(&model2)
		s.Error(e2)
		s.Empty(string(v2))
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15.999999999")

		model.Date = *NewDate(c)
		model.DateMilli = *NewDateMilli(c)
		model.DateMicro = *NewDateMicro(c)
		model.DateNano = *NewDateNano(c)

		model.Time = *NewTime(c)
		model.TimeMilli = *NewTimeMilli(c)
		model.TimeMicro = *NewTimeMicro(c)
		model.TimeNano = *NewTimeNano(c)

		model.DateTime = *NewDateTime(c)
		model.DateTimeMilli = *NewDateTimeMilli(c)
		model.DateTimeMicro = *NewDateTimeMicro(c)
		model.DateTimeNano = *NewDateTimeNano(c)

		model.Timestamp = *NewTimestamp(c)
		model.TimestampMilli = *NewTimestampMilli(c)
		model.TimestampMicro = *NewTimestampMicro(c)
		model.TimestampNano = *NewTimestampNano(c)

		model.CreatedAt = NewDateTime(c)
		model.UpdatedAt = NewDateTime(c)
		model.DeletedAt = NewTimestamp(c)

		v, e := json.Marshal(&model)
		s.Nil(e)
		s.Equal(`{"date":"2020-08-05","date_milli":"2020-08-05.999","date_micro":"2020-08-05.999999","date_nano":"2020-08-05.999999999","time":"13:14:15","time_milli":"13:14:15.999","time_micro":"13:14:15.999999","time_nano":"13:14:15.999999999","date_time":"2020-08-05 13:14:15","date_time_milli":"2020-08-05 13:14:15.999","date_time_micro":"2020-08-05 13:14:15.999999","date_time_nano":"2020-08-05 13:14:15.999999999","created_at":"2020-08-05 13:14:15","updated_at":"2020-08-05 13:14:15","timestamp":1596633255,"timestamp_milli":1596633255999,"timestamp_micro":1596633255999999,"timestamp_nano":1596633255999999999,"deleted_at":1596633255}`, string(v))
	})
}

func (s *BuiltinTypeSuite) TestBuiltinType_UnmarshalJSON() {
	var model builtinTypeModel

	s.Run("empty value", func() {
		value := `{"date":"","date_milli":"","date_micro":"","date_nano":"","time":"","time_milli":"","time_micro":"","time_nano":"","date_time":"","date_time_milli":"","date_time_micro":"","date_time_nano":"","created_at":"","updated_at":"","timestamp":"","timestamp_milli":"","timestamp_micro":"","timestamp_nano":"","deleted_at":""}`
		s.Nil(json.Unmarshal([]byte(value), &model))

		s.Empty(model.Date.String())
		s.Empty(model.DateMilli.String())
		s.Empty(model.DateMicro.String())
		s.Empty(model.DateNano.String())

		s.Empty(model.Time.String())
		s.Empty(model.TimeMilli.String())
		s.Empty(model.TimeMicro.String())
		s.Empty(model.TimeNano.String())

		s.Empty(model.DateTime.String())
		s.Empty(model.DateTimeMilli.String())
		s.Empty(model.DateTimeMicro.String())
		s.Empty(model.DateTimeNano.String())

		s.Equal("0", model.Timestamp.String())
		s.Equal("0", model.TimestampMilli.String())
		s.Equal("0", model.TimestampMicro.String())
		s.Equal("0", model.TimestampNano.String())

		s.Zero(model.Timestamp.Int64())
		s.Zero(model.TimestampMilli.Int64())
		s.Zero(model.TimestampMicro.Int64())
		s.Zero(model.TimestampNano.Int64())

		s.Empty(model.CreatedAt.String())
		s.Empty(model.UpdatedAt.String())
		s.Equal("0", model.DeletedAt.String())
		s.Equal(int64(0), model.DeletedAt.Int64())
	})

	s.Run("null value", func() {
		value1 := `{"date":null,"date_milli":null,"date_micro":null,"date_nano":null,"time":null,"time_milli":null,"time_micro":null,"time_nano":null,"date_time":null,"date_time_milli":null,"date_time_micro":null,"date_time_nano":null,"created_at":null,"updated_at":null,"timestamp":null,"timestamp_milli":null,"timestamp_micro":null,"timestamp_nano":null,"deleted_at":null}`
		s.Nil(json.Unmarshal([]byte(value1), &model))

		s.Empty(model.Date.String())
		s.Empty(model.DateMilli.String())
		s.Empty(model.DateMicro.String())
		s.Empty(model.DateNano.String())

		s.Empty(model.Time.String())
		s.Empty(model.TimeMilli.String())
		s.Empty(model.TimeMicro.String())
		s.Empty(model.TimeNano.String())

		s.Empty(model.DateTime.String())
		s.Empty(model.DateTimeMilli.String())
		s.Empty(model.DateTimeMicro.String())
		s.Empty(model.DateTimeNano.String())

		s.Equal("0", model.Timestamp.String())
		s.Equal("0", model.TimestampMilli.String())
		s.Equal("0", model.TimestampMicro.String())
		s.Equal("0", model.TimestampNano.String())

		s.Zero(model.Timestamp.Int64())
		s.Zero(model.TimestampMilli.Int64())
		s.Zero(model.TimestampMicro.Int64())
		s.Zero(model.TimestampNano.Int64())

		s.Empty(model.CreatedAt.String())
		s.Empty(model.UpdatedAt.String())
		s.Equal("0", model.DeletedAt.String())
		s.Equal(int64(0), model.DeletedAt.Int64())

		value2 := `{"date":"null","date_milli":"null","date_micro":"null","date_nano":"null","time":"null","time_milli":"null","time_micro":"null","time_nano":"null","date_time":"null","date_time_milli":"null","date_time_micro":"null","date_time_nano":"null","created_at":"null","updated_at":"null","timestamp":"null","timestamp_milli":"null","timestamp_micro":"null","timestamp_nano":"null","deleted_at":"null"}`
		s.Nil(json.Unmarshal([]byte(value2), &model))

		s.Empty(model.Date.String())
		s.Empty(model.DateMilli.String())
		s.Empty(model.DateMicro.String())
		s.Empty(model.DateNano.String())

		s.Empty(model.Time.String())
		s.Empty(model.TimeMilli.String())
		s.Empty(model.TimeMicro.String())
		s.Empty(model.TimeNano.String())

		s.Empty(model.DateTime.String())
		s.Empty(model.DateTimeMilli.String())
		s.Empty(model.DateTimeMicro.String())
		s.Empty(model.DateTimeNano.String())

		s.Equal("0", model.Timestamp.String())
		s.Equal("0", model.TimestampMilli.String())
		s.Equal("0", model.TimestampMicro.String())
		s.Equal("0", model.TimestampNano.String())

		s.Zero(model.Timestamp.Int64())
		s.Zero(model.TimestampMilli.Int64())
		s.Zero(model.TimestampMicro.Int64())
		s.Zero(model.TimestampNano.Int64())

		s.Empty(model.CreatedAt.String())
		s.Empty(model.UpdatedAt.String())
		s.Equal("0", model.DeletedAt.String())
		s.Equal(int64(0), model.DeletedAt.Int64())
	})

	s.Run("error value", func() {
		var model1 builtinTypeModel

		value1 := `{"date":"xxx","date_milli":"xxx","date_micro":"xxx","date_nano":"xxx","time":"xxx","time_milli":"xxx","time_micro":"xxx","time_nano":"xxx","date_time":"xxx","date_time_milli":"xxx","date_time_micro":"xxx","date_time_nano":"xxx","created_at":"xxx","updated_at":"xxx","timestamp":"xxx"}`
		s.Error(json.Unmarshal([]byte(value1), &model1))

		var model2 builtinTypeModel

		value2 := `{"timestamp":"xxx","timestamp_milli":"xxx","timestamp_micro":"xxx","timestamp_nano":"xxx","deleted_at":"xxx"}`
		s.Error(json.Unmarshal([]byte(value2), &model2))

		s.Empty(model.Date.String())
		s.Empty(model.DateMilli.String())
		s.Empty(model.DateMicro.String())
		s.Empty(model.DateNano.String())

		s.Empty(model.Time.String())
		s.Empty(model.TimeMilli.String())
		s.Empty(model.TimeMicro.String())
		s.Empty(model.TimeNano.String())

		s.Empty(model.DateTime.String())
		s.Empty(model.DateTimeMilli.String())
		s.Empty(model.DateTimeMicro.String())
		s.Empty(model.DateTimeNano.String())

		s.Equal("0", model.Timestamp.String())
		s.Equal("0", model.TimestampMilli.String())
		s.Equal("0", model.TimestampMicro.String())
		s.Equal("0", model.TimestampNano.String())

		s.Zero(model.Timestamp.Int64())
		s.Zero(model.TimestampMilli.Int64())
		s.Zero(model.TimestampMicro.Int64())
		s.Zero(model.TimestampNano.Int64())

		s.Empty(model.CreatedAt.String())
		s.Empty(model.UpdatedAt.String())
		s.Equal("0", model.DeletedAt.String())
		s.Equal(int64(0), model.DeletedAt.Int64())
	})

	s.Run("valid value", func() {
		value := `{"date":"2020-08-05","date_milli":"2020-08-05.999","date_micro":"2020-08-05.999999","date_nano":"2020-08-05.999999999","time":"13:14:15","time_milli":"13:14:15.999","time_micro":"13:14:15.999999","time_nano":"13:14:15.999999999","date_time":"2020-08-05 13:14:15","date_time_milli":"2020-08-05 13:14:15.999","date_time_micro":"2020-08-05 13:14:15.999999","date_time_nano":"2020-08-05 13:14:15.999999999","created_at":"2020-08-05 13:14:15","updated_at":"2020-08-05 13:14:15","timestamp":1596633255,"timestamp_milli":1596633255999,"timestamp_micro":1596633255999999,"timestamp_nano":1596633255999999999,"deleted_at":1596633255}`
		s.Nil(json.Unmarshal([]byte(value), &model))

		s.Equal("2020-08-05", model.Date.String())
		s.Equal("2020-08-05.999", model.DateMilli.String())
		s.Equal("2020-08-05.999999", model.DateMicro.String())
		s.Equal("2020-08-05.999999999", model.DateNano.String())

		s.Equal("13:14:15", model.Time.String())
		s.Equal("13:14:15.999", model.TimeMilli.String())
		s.Equal("13:14:15.999999", model.TimeMicro.String())
		s.Equal("13:14:15.999999999", model.TimeNano.String())

		s.Equal("2020-08-05 13:14:15", model.DateTime.String())
		s.Equal("2020-08-05 13:14:15.999", model.DateTimeMilli.String())
		s.Equal("2020-08-05 13:14:15.999999", model.DateTimeMicro.String())
		s.Equal("2020-08-05 13:14:15.999999999", model.DateTimeNano.String())

		s.Equal("1596633255", model.Timestamp.String())
		s.Equal("1596633255999", model.TimestampMilli.String())
		s.Equal("1596633255999999", model.TimestampMicro.String())
		s.Equal("1596633255999999999", model.TimestampNano.String())

		s.Equal(int64(1596633255), model.Timestamp.Int64())
		s.Equal(int64(1596633255999), model.TimestampMilli.Int64())
		s.Equal(int64(1596633255999999), model.TimestampMicro.Int64())
		s.Equal(int64(1596633255999999999), model.TimestampNano.Int64())

		s.Equal("2020-08-05 13:14:15", model.CreatedAt.String())
		s.Equal("2020-08-05 13:14:15", model.UpdatedAt.String())
		s.Equal("1596633255", model.DeletedAt.String())
		s.Equal(int64(1596633255), model.DeletedAt.Int64())
	})
}

type rfc3339Type string

func (t rfc3339Type) Layout() string {
	return RFC3339Layout
}

type w3cType string

func (w3cType) Layout() string {
	return W3cLayout
}

type iso8601Type string

func (t iso8601Type) Format() string {
	return ISO8601Format
}

type rssType string

func (rssType) Format() string {
	return RssFormat
}

type CustomerTypeModel struct {
	Customer1 LayoutType[rfc3339Type] `json:"customer1"`
	Customer2 LayoutType[w3cType]     `json:"customer2"`
	Customer3 FormatType[iso8601Type] `json:"customer3"`
	Customer4 FormatType[rssType]     `json:"customer4"`

	CreatedAt *FormatType[iso8601Type] `json:"created_at"`
	UpdatedAt *LayoutType[rfc3339Type] `json:"updated_at"`
}

type CustomerTypeSuite struct {
	suite.Suite
}

func TestCustomerTypeSuite(t *testing.T) {
	suite.Run(t, new(CustomerTypeSuite))
}

func (s *CustomerTypeSuite) TestCustomerType_Scan() {
	c := Now()

	t1 := NewLayoutType[rfc3339Type](c)
	t2 := NewLayoutType[w3cType](c)
	t3 := NewFormatType[iso8601Type](c)
	t4 := NewFormatType[rssType](c)

	s.Run("[]byte type", func() {
		s.Nil(t1.Scan([]byte(c.ToDateString())))
		s.Nil(t2.Scan([]byte(c.ToDateString())))
		s.Nil(t3.Scan([]byte(c.ToDateString())))
		s.Nil(t4.Scan([]byte(c.ToDateString())))
	})

	s.Run("string type", func() {
		s.Nil(t1.Scan(c.ToDateString()))
		s.Nil(t2.Scan(c.ToDateString()))
		s.Nil(t3.Scan(c.ToDateString()))
		s.Nil(t4.Scan(c.ToDateString()))
	})

	s.Run("time type", func() {
		tt := time.Now()
		s.Nil(t1.Scan(tt))
		s.Nil(t2.Scan(tt))
		s.Nil(t3.Scan(tt))
		s.Nil(t4.Scan(tt))
	})

	s.Run("*time type", func() {
		tt := time.Now()
		s.Nil(t1.Scan(&tt))
		s.Nil(t2.Scan(&tt))
		s.Nil(t3.Scan(&tt))
		s.Nil(t4.Scan(&tt))
	})

	s.Run("nil type", func() {
		s.Nil(t1.Scan(nil))
		s.Nil(t2.Scan(nil))
		s.Nil(t3.Scan(nil))
		s.Nil(t4.Scan(nil))
	})

	s.Run("unsupported type", func() {
		s.Error(t1.Scan(true))
		s.Error(t1.Scan(func() {}))
		s.Error(t1.Scan(int64(0)))
		s.Error(t1.Scan(float64(0)))
		s.Error(t1.Scan(map[string]string{}))

		s.Error(t2.Scan(true))
		s.Error(t2.Scan(func() {}))
		s.Error(t2.Scan(int64(0)))
		s.Error(t2.Scan(float64(0)))
		s.Error(t2.Scan(map[string]string{}))

		s.Error(t3.Scan(true))
		s.Error(t3.Scan(func() {}))
		s.Error(t3.Scan(int64(0)))
		s.Error(t3.Scan(float64(0)))
		s.Error(t3.Scan(map[string]string{}))

		s.Error(t4.Scan(true))
		s.Error(t4.Scan(func() {}))
		s.Error(t4.Scan(int64(0)))
		s.Error(t4.Scan(float64(0)))
		s.Error(t4.Scan(map[string]string{}))
	})
}

func (s *CustomerTypeSuite) TestCustomerType_Value() {
	s.Run("nil carbon", func() {
		t1, e1 := NewLayoutType[rfc3339Type](nil).Value()
		s.Nil(t1)
		s.Nil(e1)

		t2, e2 := NewLayoutType[w3cType](nil).Value()
		s.Nil(t2)
		s.Nil(e2)

		t3, e3 := NewFormatType[iso8601Type](nil).Value()
		s.Nil(t3)
		s.Nil(e3)

		t4, e4 := NewFormatType[rssType](nil).Value()
		s.Nil(t4)
		s.Nil(e4)
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()

		t1, e1 := NewLayoutType[rfc3339Type](c).Value()
		s.Nil(t1)
		s.Nil(e1)

		t2, e2 := NewLayoutType[w3cType](c).Value()
		s.Nil(t2)
		s.Nil(e2)

		t3, e3 := NewFormatType[iso8601Type](c).Value()
		s.Nil(t3)
		s.Nil(e3)

		t4, e4 := NewFormatType[rssType](c).Value()
		s.Nil(t4)
		s.Nil(e4)
	})

	s.Run("empty carbon", func() {
		c := Parse("")

		t1, e1 := NewLayoutType[rfc3339Type](c).Value()
		s.Nil(t1)
		s.Nil(e1)

		t2, e2 := NewLayoutType[w3cType](c).Value()
		s.Nil(t2)
		s.Nil(e2)

		t3, e3 := NewFormatType[iso8601Type](c).Value()
		s.Nil(t3)
		s.Nil(e3)

		t4, e4 := NewFormatType[rssType](c).Value()
		s.Nil(t4)
		s.Nil(e4)
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")

		t1, e1 := NewLayoutType[rfc3339Type](c).Value()
		s.Nil(t1)
		s.Error(e1)

		t2, e2 := NewLayoutType[w3cType](c).Value()
		s.Nil(t2)
		s.Error(e2)

		t3, e3 := NewFormatType[iso8601Type](c).Value()
		s.Nil(t3)
		s.Error(e3)

		t4, e4 := NewFormatType[rssType](c).Value()
		s.Nil(t4)
		s.Error(e4)
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05")

		t1, e1 := NewLayoutType[rfc3339Type](c).Value()
		s.Equal(c.StdTime(), t1)
		s.Nil(e1)

		t2, e2 := NewLayoutType[w3cType](c).Value()
		s.Equal(c.StdTime(), t2)
		s.Nil(e2)

		t3, e3 := NewFormatType[iso8601Type](c).Value()
		s.Equal(c.StdTime(), t3)
		s.Nil(e3)

		t4, e4 := NewFormatType[rssType](c).Value()
		s.Equal(c.StdTime(), t4)
		s.Nil(e4)
	})
}

func (s *CustomerTypeSuite) TestCustomerType_MarshalJSON() {
	var model CustomerTypeModel

	s.Run("nil carbon", func() {
		model.Customer1 = *NewLayoutType[rfc3339Type](nil)
		model.Customer2 = *NewLayoutType[w3cType](nil)
		model.Customer3 = *NewFormatType[iso8601Type](nil)
		model.Customer4 = *NewFormatType[rssType](nil)

		model.CreatedAt = NewFormatType[iso8601Type](nil)
		model.UpdatedAt = NewLayoutType[rfc3339Type](nil)

		v, e := json.Marshal(&model)
		s.Nil(e)
		s.Equal(`{"customer1":null,"customer2":null,"customer3":null,"customer4":null,"created_at":null,"updated_at":null}`, string(v))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()

		model.Customer1 = *NewLayoutType[rfc3339Type](c)
		model.Customer2 = *NewLayoutType[w3cType](c)
		model.Customer3 = *NewFormatType[iso8601Type](c)
		model.Customer4 = *NewFormatType[rssType](c)

		model.CreatedAt = NewFormatType[iso8601Type](c)
		model.UpdatedAt = NewLayoutType[rfc3339Type](c)

		v, e := json.Marshal(&model)
		s.Nil(e)
		s.Equal(`{"customer1":null,"customer2":null,"customer3":null,"customer4":null,"created_at":null,"updated_at":null}`, string(v))
	})

	s.Run("empty carbon", func() {
		c := Parse("")

		model.Customer1 = *NewLayoutType[rfc3339Type](c)
		model.Customer2 = *NewLayoutType[w3cType](c)
		model.Customer3 = *NewFormatType[iso8601Type](c)
		model.Customer4 = *NewFormatType[rssType](c)

		model.CreatedAt = NewFormatType[iso8601Type](c)
		model.UpdatedAt = NewLayoutType[rfc3339Type](c)

		v, e := json.Marshal(&model)
		s.Nil(e)
		s.Equal(`{"customer1":null,"customer2":null,"customer3":null,"customer4":null,"created_at":null,"updated_at":null}`, string(v))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")

		var model1 CustomerTypeModel
		model1.Customer1 = *NewLayoutType[rfc3339Type](c)
		model1.Customer2 = *NewLayoutType[w3cType](c)
		v1, e1 := json.Marshal(&model1)
		s.Error(e1)
		s.Empty(string(v1))

		var model2 CustomerTypeModel
		model2.Customer3 = *NewFormatType[iso8601Type](c)
		model2.Customer4 = *NewFormatType[rssType](c)
		v2, e2 := json.Marshal(&model2)
		s.Error(e2)
		s.Empty(string(v2))

		var model3 CustomerTypeModel
		model3.CreatedAt = NewFormatType[iso8601Type](c)
		model3.UpdatedAt = NewLayoutType[rfc3339Type](c)
		v3, e3 := json.Marshal(&model3)
		s.Error(e3)
		s.Empty(string(v3))

	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15.999999999")

		model.Customer1 = *NewLayoutType[rfc3339Type](c)
		model.Customer2 = *NewLayoutType[w3cType](c)
		model.Customer3 = *NewFormatType[iso8601Type](c)
		model.Customer4 = *NewFormatType[rssType](c)

		model.CreatedAt = NewFormatType[iso8601Type](c)
		model.UpdatedAt = NewLayoutType[rfc3339Type](c)

		v, e := json.Marshal(&model)
		s.Nil(e)
		s.Equal(`{"customer1":"2020-08-05T13:14:15Z","customer2":"2020-08-05T13:14:15Z","customer3":"2020-08-05T13:14:15+00:00","customer4":"Wed, 05 Aug 2020 13:14:15 +0000","created_at":"2020-08-05T13:14:15+00:00","updated_at":"2020-08-05T13:14:15Z"}`, string(v))
	})
}

func (s *CustomerTypeSuite) TestCustomerType_UnmarshalJSON() {
	var model CustomerTypeModel

	s.Run("empty value", func() {
		value := `{"customer1":"","customer2":"","customer3":"","customer4":"","created_at":"","updated_at":""}`
		s.Nil(json.Unmarshal([]byte(value), &model))

		s.Empty(model.Customer1.String())
		s.Empty(model.Customer2.String())
		s.Empty(model.Customer3.String())
		s.Empty(model.Customer4.String())
		s.Empty(model.CreatedAt.String())
		s.Empty(model.UpdatedAt.String())
	})

	s.Run("null value", func() {
		value1 := `{"customer1":null,"customer2":null,"customer3":null,"customer4":null,"created_at":null,"updated_at":null}`
		s.Nil(json.Unmarshal([]byte(value1), &model))

		s.Empty(model.Customer1.String())
		s.Empty(model.Customer2.String())
		s.Empty(model.Customer3.String())
		s.Empty(model.Customer4.String())
		s.Empty(model.CreatedAt.String())
		s.Empty(model.UpdatedAt.String())

		value2 := `{"customer1":"null","customer2":"null","customer3":"null","customer4":"null","created_at":"null","updated_at":"null"}`
		s.Nil(json.Unmarshal([]byte(value2), &model))

		s.Empty(model.Customer1.String())
		s.Empty(model.Customer2.String())
		s.Empty(model.CreatedAt.String())
		s.Empty(model.UpdatedAt.String())
	})

	s.Run("error value", func() {
		value := `{"customer1":"xxx","customer2":"xxx","customer3":"xxx","customer4":"xxx","created_at":"xxx","updated_at":"xxx"}`
		s.Error(json.Unmarshal([]byte(value), &model))

		s.Empty(model.Customer1.String())
		s.Empty(model.Customer2.String())
		s.Empty(model.Customer3.String())
		s.Empty(model.Customer4.String())
		s.Empty(model.CreatedAt.String())
		s.Empty(model.UpdatedAt.String())
	})

	s.Run("valid value", func() {
		value := `{"customer1":"2020-08-05T13:14:15Z","customer2":"2020-08-05T13:14:15Z","customer3":"2020-08-05T13:14:15+00:00","customer4":"Wed, 05 Aug 2020 13:14:15 +0000","created_at":"2020-08-05T13:14:15+00:00","updated_at":"2020-08-05T13:14:15Z"}`
		s.Nil(json.Unmarshal([]byte(value), &model))

		s.Equal("2020-08-05T13:14:15Z", model.Customer1.String())
		s.Equal("2020-08-05T13:14:15Z", model.Customer2.String())
		s.Equal("2020-08-05T13:14:15+00:00", model.Customer3.String())
		s.Equal("Wed, 05 Aug 2020 13:14:15 +0000", model.Customer4.String())
		s.Equal("2020-08-05T13:14:15+00:00", model.CreatedAt.String())
		s.Equal("2020-08-05T13:14:15Z", model.UpdatedAt.String())
	})
}
