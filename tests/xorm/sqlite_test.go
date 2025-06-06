package xorm

import (
	"encoding/json"
	"testing"

	"github.com/dromara/carbon/v2"
	"github.com/stretchr/testify/suite"
)

type SQLiteSuite struct {
	suite.Suite
}

func TestSQLiteSuite(t *testing.T) {
	suite.Run(t, new(SQLiteSuite))
}

func (s *SQLiteSuite) SetupSuite() {
	carbon.SetTimezone(carbon.PRC)
	carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15.111111111"))
	db = connect(driverSQLite)
	if err = db.Sync(&SQLiteModel1{}); err != nil {
		panic(err)
	}
	if err = db.Sync(&SQLiteModel2{}); err != nil {
		panic(err)
	}
}

func (s *SQLiteSuite) TearDownSuite() {
	carbon.ClearTestNow()
	if _, err = db.Unscoped().Where("1 = 1").Delete(SQLiteModel1{}); err != nil {
		panic(err)
	}
	if _, err = db.Unscoped().Where("1 = 1").Delete(SQLiteModel1{}); err != nil {
		panic(err)
	}
}

func (s *SQLiteSuite) TestCurd1() {
	s.Run("unset carbon", func() {
		var model1 SQLiteModel1

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 SQLiteModel1
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":null,"date":null,"time":null,"date_time":null,"rfc3339_layout":null,"iso8601_format":null,"timestamp":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("nil carbon", func() {
		var model1 SQLiteModel1

		var c *carbon.Carbon
		c = nil

		model1.Date = *carbon.NewDate(c)
		model1.Time = *carbon.NewTime(c)
		model1.DateTime = *carbon.NewDateTime(c)

		model1.RFC3339Layout = *carbon.NewLayoutType[RFC3339Layout](c)
		model1.ISO8601Format = *carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp = *carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 SQLiteModel1
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":null,"date":null,"time":null,"date_time":null,"rfc3339_layout":null,"iso8601_format":null,"timestamp":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("zero carbon", func() {
		var model1 SQLiteModel1

		c := carbon.NewCarbon()

		model1.Carbon = *c

		model1.Date = *carbon.NewDate(c)
		model1.Time = *carbon.NewTime(c)
		model1.DateTime = *carbon.NewDateTime(c)

		model1.RFC3339Layout = *carbon.NewLayoutType[RFC3339Layout](c)
		model1.ISO8601Format = *carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp = *carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 SQLiteModel1
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":null,"date":null,"time":null,"date_time":null,"rfc3339_layout":null,"iso8601_format":null,"timestamp":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("empty carbon", func() {
		var model1 SQLiteModel1

		c := carbon.Parse("")

		model1.Carbon = *c

		model1.Date = *carbon.NewDate(c)
		model1.Time = *carbon.NewTime(c)
		model1.DateTime = *carbon.NewDateTime(c)

		model1.RFC3339Layout = *carbon.NewLayoutType[RFC3339Layout](c)
		model1.ISO8601Format = *carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp = *carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 SQLiteModel1
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":null,"date":null,"time":null,"date_time":null,"rfc3339_layout":null,"iso8601_format":null,"timestamp":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("valid carbon", func() {
		var model1 SQLiteModel1

		c := carbon.Now()

		model1.Carbon = *c

		model1.Date = *carbon.NewDate(c)
		model1.Time = *carbon.NewTime(c)
		model1.DateTime = *carbon.NewDateTime(c)

		model1.RFC3339Layout = *carbon.NewLayoutType[RFC3339Layout](c)
		model1.ISO8601Format = *carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp = *carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 SQLiteModel1
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":"2020-08-05 13:14:15","date":"2020-08-05","time":"13:14:15","date_time":"2020-08-05 13:14:15","rfc3339_layout":"2020-08-05T13:14:15+08:00","iso8601_format":"2020-08-05T13:14:15+08:00","timestamp":1596604455}`, string(data1))

		c = c.Copy().AddDay()

		model2.Carbon = *c

		model2.Date = *carbon.NewDate(c)
		model2.Time = *carbon.NewTime(c)
		model2.DateTime = *carbon.NewDateTime(c)

		model2.RFC3339Layout = *carbon.NewLayoutType[RFC3339Layout](c)
		model2.ISO8601Format = *carbon.NewFormatType[ISO8601Format](c)

		model2.Timestamp = *carbon.NewTimestamp(c)

		// update
		_, err = db.Update(model2)
		s.Nil(err)

		data2, err2 := json.Marshal(&model2)
		s.Nil(err2)
		s.Equal(`{"carbon":"2020-08-06 13:14:15","date":"2020-08-06","time":"13:14:15","date_time":"2020-08-06 13:14:15","rfc3339_layout":"2020-08-06T13:14:15+08:00","iso8601_format":"2020-08-06T13:14:15+08:00","timestamp":1596690855}`, string(data2))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})
}

func (s *SQLiteSuite) TestCurd2() {
	s.Run("unset carbon", func() {
		var model1 SQLiteModel2

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 SQLiteModel2
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":null,"date":null,"time":null,"date_time":null,"rfc3339_layout":null,"iso8601_format":null,"timestamp":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("nil carbon", func() {
		var model1 SQLiteModel2

		var c *carbon.Carbon
		c = nil

		model1.Carbon = c

		model1.Date = carbon.NewDate(c)
		model1.Time = carbon.NewTime(c)
		model1.DateTime = carbon.NewDateTime(c)

		model1.RFC3339Layout = carbon.NewLayoutType[RFC3339Layout](c)
		model1.ISO8601Format = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp = carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 SQLiteModel2
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":null,"date":null,"time":null,"date_time":null,"rfc3339_layout":null,"iso8601_format":null,"timestamp":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("zero carbon", func() {
		var model1 SQLiteModel2

		c := carbon.NewCarbon()

		model1.Carbon = c

		model1.Date = carbon.NewDate(c)
		model1.Time = carbon.NewTime(c)
		model1.DateTime = carbon.NewDateTime(c)

		model1.RFC3339Layout = carbon.NewLayoutType[RFC3339Layout](c)
		model1.ISO8601Format = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp = carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 SQLiteModel2
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":null,"date":null,"time":null,"date_time":null,"rfc3339_layout":null,"iso8601_format":null,"timestamp":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("empty carbon", func() {
		var model1 SQLiteModel2

		c := carbon.Parse("")

		model1.Carbon = c

		model1.Date = carbon.NewDate(c)
		model1.Time = carbon.NewTime(c)
		model1.DateTime = carbon.NewDateTime(c)

		model1.RFC3339Layout = carbon.NewLayoutType[RFC3339Layout](c)
		model1.ISO8601Format = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp = carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 SQLiteModel2
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":null,"date":null,"time":null,"date_time":null,"rfc3339_layout":null,"iso8601_format":null,"timestamp":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("valid carbon", func() {
		var model1 SQLiteModel2

		c := carbon.Now()

		model1.Carbon = c

		model1.Date = carbon.NewDate(c)
		model1.Time = carbon.NewTime(c)
		model1.DateTime = carbon.NewDateTime(c)

		model1.RFC3339Layout = carbon.NewLayoutType[RFC3339Layout](c)
		model1.ISO8601Format = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp = carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 SQLiteModel2
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":"2020-08-05 13:14:15","date":"2020-08-05","time":"13:14:15","date_time":"2020-08-05 13:14:15","rfc3339_layout":"2020-08-05T13:14:15+08:00","iso8601_format":"2020-08-05T13:14:15+08:00","timestamp":1596604455}`, string(data1))

		c = c.Copy().AddDay()

		model2.Carbon = c

		model2.Date = carbon.NewDate(c)
		model2.Time = carbon.NewTime(c)
		model2.DateTime = carbon.NewDateTime(c)

		model2.RFC3339Layout = carbon.NewLayoutType[RFC3339Layout](c)
		model2.ISO8601Format = carbon.NewFormatType[ISO8601Format](c)

		model2.Timestamp = carbon.NewTimestamp(c)

		// update
		_, err = db.Update(model2)
		s.Nil(err)

		data2, err2 := json.Marshal(&model2)
		s.Nil(err2)
		s.Equal(`{"carbon":"2020-08-06 13:14:15","date":"2020-08-06","time":"13:14:15","date_time":"2020-08-06 13:14:15","rfc3339_layout":"2020-08-06T13:14:15+08:00","iso8601_format":"2020-08-06T13:14:15+08:00","timestamp":1596690855}`, string(data2))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})
}
