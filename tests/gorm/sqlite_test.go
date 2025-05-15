package gorm

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
	carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15"))
	db = connect(driverSQLite)
	if err := db.AutoMigrate(&SQLiteModel1{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&SQLiteModel2{}); err != nil {
		panic(err)
	}
}

func (s *SQLiteSuite) TearDownSuite() {
	carbon.ClearTestNow()
	db.Unscoped().Where("1 = 1").Delete(&SQLiteModel1{})
	db.Unscoped().Where("1 = 1").Delete(&SQLiteModel2{})
}

func (s *SQLiteSuite) TestCurd1() {
	s.Run("unset carbon", func() {
		var model1 SQLiteModel1

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 SQLiteModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":null,"date":null,"time":null,"date_time":null,"rfc3339_layout":null,"iso8601_format":null,"timestamp":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("zero carbon", func() {
		var model1 SQLiteModel1

		c := carbon.NewCarbon()
		model1.Carbon = c

		model1.Date = carbon.NewDate(c)
		model1.Time = carbon.NewTime(c)
		model1.DateTime = carbon.NewDateTime(c)

		model1.RFC3339Layout = carbon.NewLayoutType[RFC3339Layout](c)
		model1.ISO8601Format = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp = carbon.NewTimestamp(c)

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 SQLiteModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":null,"date":null,"time":null,"date_time":null,"rfc3339_layout":null,"iso8601_format":null,"timestamp":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("empty carbon", func() {
		var model1 SQLiteModel1

		c := carbon.Parse("")
		model1.Carbon = c

		model1.Date = carbon.NewDate(c)
		model1.Time = carbon.NewTime(c)
		model1.DateTime = carbon.NewDateTime(c)

		model1.RFC3339Layout = carbon.NewLayoutType[RFC3339Layout](c)
		model1.ISO8601Format = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp = carbon.NewTimestamp(c)

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 SQLiteModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":null,"date":null,"time":null,"date_time":null,"rfc3339_layout":null,"iso8601_format":null,"timestamp":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("valid carbon", func() {
		var model1 SQLiteModel1

		c := carbon.Now()
		model1.Carbon = c

		model1.Date = carbon.NewDate(c)
		model1.Time = carbon.NewTime(c)
		model1.DateTime = carbon.NewDateTime(c)

		model1.RFC3339Layout = carbon.NewLayoutType[RFC3339Layout](c)
		model1.ISO8601Format = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp = carbon.NewTimestamp(c)

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 SQLiteModel1
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":"2020-08-05 13:14:15","date":"2020-08-05","time":"13:14:15","date_time":"2020-08-05 13:14:15","rfc3339_layout":"2020-08-05T13:14:15+08:00","iso8601_format":"2020-08-05T13:14:15+08:00","timestamp":1596604455}`, string(data1))

		c = c.AddDay()

		model2.Carbon = c

		model2.Date = carbon.NewDate(c)
		model2.Time = carbon.NewTime(c)
		model2.DateTime = carbon.NewDateTime(c)

		model2.RFC3339Layout = carbon.NewLayoutType[RFC3339Layout](c)
		model2.ISO8601Format = carbon.NewFormatType[ISO8601Format](c)

		model2.Timestamp = carbon.NewTimestamp(c)

		// update
		db.Save(&model2)

		data2, err2 := json.Marshal(&model2)
		s.Nil(err2)
		s.Equal(`{"carbon":"2020-08-06 13:14:15","date":"2020-08-06","time":"13:14:15","date_time":"2020-08-06 13:14:15","rfc3339_layout":"2020-08-06T13:14:15+08:00","iso8601_format":"2020-08-06T13:14:15+08:00","timestamp":1596690855}`, string(data2))

		// delete
		db.Delete(&model2)
	})
}

func (s *SQLiteSuite) TestCurd2() {
	s.Run("unset carbon", func() {
		var model1 SQLiteModel2

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 SQLiteModel2
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":null,"date":null,"time":null,"date_time":null,"rfc3339_layout":null,"iso8601_format":null,"timestamp":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("zero carbon", func() {
		var model1 SQLiteModel2

		c := carbon.NewCarbon()
		model1.Carbon = &c

		date := carbon.NewDate(c)
		model1.Date = &date

		time := carbon.NewTime(c)
		model1.Time = &time

		datetime := carbon.NewDateTime(c)
		model1.DateTime = &datetime

		rfc3339Layout := carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout = &rfc3339Layout

		iso8601Format := carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format = &iso8601Format

		timestamp := carbon.NewTimestamp(c)
		model1.Timestamp = &timestamp

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 SQLiteModel2
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":null,"date":null,"time":null,"date_time":null,"rfc3339_layout":null,"iso8601_format":null,"timestamp":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("empty carbon", func() {
		var model1 SQLiteModel2

		c := carbon.Parse("")
		model1.Carbon = &c

		date := carbon.NewDate(c)
		model1.Date = &date

		time := carbon.NewTime(c)
		model1.Time = &time

		datetime := carbon.NewDateTime(c)
		model1.DateTime = &datetime

		rfc3339Layout := carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout = &rfc3339Layout

		iso8601Format := carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format = &iso8601Format

		timestamp := carbon.NewTimestamp(c)
		model1.Timestamp = &timestamp

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 SQLiteModel2
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":null,"date":null,"time":null,"date_time":null,"rfc3339_layout":null,"iso8601_format":null,"timestamp":null}`, string(data1))

		// delete
		db.Delete(&model2)
	})

	s.Run("valid carbon", func() {
		var model1 SQLiteModel2

		c := carbon.Now()
		model1.Carbon = &c

		date := carbon.NewDate(c)
		model1.Date = &date

		time := carbon.NewTime(c)
		model1.Time = &time

		datetime := carbon.NewDateTime(c)
		model1.DateTime = &datetime

		rfc3339Layout := carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout = &rfc3339Layout

		iso8601Format := carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format = &iso8601Format

		timestamp := carbon.NewTimestamp(c)
		model1.Timestamp = &timestamp

		// create
		if err = db.Create(&model1).Error; err != nil {
			panic(err)
		}

		// read
		var model2 SQLiteModel2
		db.Last(&model2)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon":"2020-08-05 13:14:15","date":"2020-08-05","time":"13:14:15","date_time":"2020-08-05 13:14:15","rfc3339_layout":"2020-08-05T13:14:15+08:00","iso8601_format":"2020-08-05T13:14:15+08:00","timestamp":1596604455}`, string(data1))

		c = c.AddDay()

		model2.Carbon = &c

		date = carbon.NewDate(c)
		model2.Date = &date

		time = carbon.NewTime(c)
		model2.Time = &time

		datetime = carbon.NewDateTime(c)
		model2.DateTime = &datetime

		rfc3339Layout = carbon.NewLayoutType[RFC3339Layout](c)
		model2.RFC3339Layout = &rfc3339Layout

		iso8601Format = carbon.NewFormatType[ISO8601Format](c)
		model2.ISO8601Format = &iso8601Format

		timestamp = carbon.NewTimestamp(c)
		model2.Timestamp = &timestamp

		// update
		db.Save(&model2)

		data2, err2 := json.Marshal(&model2)
		s.Nil(err2)
		s.Equal(`{"carbon":"2020-08-06 13:14:15","date":"2020-08-06","time":"13:14:15","date_time":"2020-08-06 13:14:15","rfc3339_layout":"2020-08-06T13:14:15+08:00","iso8601_format":"2020-08-06T13:14:15+08:00","timestamp":1596690855}`, string(data2))

		// delete
		db.Delete(&model2)
	})
}
