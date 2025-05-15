package xorm

import (
	"encoding/json"
	"testing"

	"github.com/dromara/carbon/v2"
	"github.com/stretchr/testify/suite"
)

type MySQLSuite struct {
	suite.Suite
}

func TestMySQLSuite(t *testing.T) {
	suite.Run(t, new(MySQLSuite))
}

func (s *MySQLSuite) SetupSuite() {
	carbon.SetTimezone(carbon.PRC)
	carbon.SetTestNow(carbon.Parse("2020-08-05 13:14:15"))
	db = connect(driverMySQL)
	if err = db.Sync(&MySQLModel1{}); err != nil {
		panic(err)
	}
	if err = db.Sync(&MySQLModel2{}); err != nil {
		panic(err)
	}
}

func (s *MySQLSuite) TearDownSuite() {
	carbon.ClearTestNow()
	if _, err = db.Unscoped().Where("1 = 1").Delete(MySQLModel1{}); err != nil {
		panic(err)
	}
	if _, err = db.Unscoped().Where("1 = 1").Delete(MySQLModel2{}); err != nil {
		panic(err)
	}
}

func (s *MySQLSuite) TestCurd1() {
	s.Run("unset carbon", func() {
		var model1 MySQLModel1

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 MySQLModel1
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"carbon3":null,"date1":null,"date2":null,"date3":null,"time1":null,"time2":null,"time3":null,"date_time1":null,"date_time2":null,"date_time3":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"rfc3339_layout3":null,"iso8601_format1":null,"iso8601_format2":null,"iso8601_format3":null,"timestamp1":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("zero carbon", func() {
		var model1 MySQLModel1

		c := carbon.NewCarbon()
		model1.Carbon1 = c
		model1.Carbon2 = c
		model1.Carbon3 = c

		model1.Date1 = carbon.NewDate(c)
		model1.Date2 = carbon.NewDate(c)
		model1.Date3 = carbon.NewDate(c)

		model1.Time1 = carbon.NewTime(c)
		model1.Time2 = carbon.NewTime(c)
		model1.Time3 = carbon.NewTime(c)

		model1.DateTime1 = carbon.NewDateTime(c)
		model1.DateTime2 = carbon.NewDateTime(c)
		model1.DateTime3 = carbon.NewDateTime(c)

		model1.RFC3339Layout1 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout3 = carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format3 = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 MySQLModel1
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"carbon3":null,"date1":null,"date2":null,"date3":null,"time1":null,"time2":null,"time3":null,"date_time1":null,"date_time2":null,"date_time3":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"rfc3339_layout3":null,"iso8601_format1":null,"iso8601_format2":null,"iso8601_format3":null,"timestamp1":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("empty carbon", func() {
		var model1 MySQLModel1

		c := carbon.Parse("")
		model1.Carbon1 = c
		model1.Carbon2 = c
		model1.Carbon3 = c

		model1.Date1 = carbon.NewDate(c)
		model1.Date2 = carbon.NewDate(c)
		model1.Date3 = carbon.NewDate(c)

		model1.Time1 = carbon.NewTime(c)
		model1.Time2 = carbon.NewTime(c)
		model1.Time3 = carbon.NewTime(c)

		model1.DateTime1 = carbon.NewDateTime(c)
		model1.DateTime2 = carbon.NewDateTime(c)
		model1.DateTime3 = carbon.NewDateTime(c)

		model1.RFC3339Layout1 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout3 = carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format3 = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 MySQLModel1
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"carbon3":null,"date1":null,"date2":null,"date3":null,"time1":null,"time2":null,"time3":null,"date_time1":null,"date_time2":null,"date_time3":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"rfc3339_layout3":null,"iso8601_format1":null,"iso8601_format2":null,"iso8601_format3":null,"timestamp1":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("valid carbon", func() {
		var model1 MySQLModel1

		c := carbon.Now()
		model1.Carbon1 = c
		model1.Carbon2 = c
		model1.Carbon3 = c

		model1.Date1 = carbon.NewDate(c)
		model1.Date2 = carbon.NewDate(c)
		model1.Date3 = carbon.NewDate(c)

		model1.Time1 = carbon.NewTime(c)
		model1.Time2 = carbon.NewTime(c)
		model1.Time3 = carbon.NewTime(c)

		model1.DateTime1 = carbon.NewDateTime(c)
		model1.DateTime2 = carbon.NewDateTime(c)
		model1.DateTime3 = carbon.NewDateTime(c)

		model1.RFC3339Layout1 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout2 = carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout3 = carbon.NewLayoutType[RFC3339Layout](c)

		model1.ISO8601Format1 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format2 = carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format3 = carbon.NewFormatType[ISO8601Format](c)

		model1.Timestamp1 = carbon.NewTimestamp(c)

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 MySQLModel1
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":"2020-08-05 13:14:15","carbon2":"2020-08-05 13:14:15","carbon3":"2020-08-05 13:14:15","date1":"2020-08-05","date2":"2020-08-05","date3":"2020-08-05","time1":"13:14:15","time2":"13:14:15","time3":"13:14:15","date_time1":"2020-08-05 13:14:15","date_time2":"2020-08-05 13:14:15","date_time3":"2020-08-05 13:14:15","rfc3339_layout1":"2020-08-05T13:14:15+08:00","rfc3339_layout2":"2020-08-05T13:14:15+08:00","rfc3339_layout3":"2020-08-05T13:14:15+08:00","iso8601_format1":"2020-08-05T13:14:15+08:00","iso8601_format2":"2020-08-05T13:14:15+08:00","iso8601_format3":"2020-08-05T13:14:15+08:00","timestamp1":1596604455}`, string(data1))

		c = c.AddDay()

		model2.Carbon1 = c
		model2.Carbon2 = c
		model2.Carbon3 = c

		model2.Date1 = carbon.NewDate(c)
		model2.Date2 = carbon.NewDate(c)
		model2.Date3 = carbon.NewDate(c)

		model2.Time1 = carbon.NewTime(c)
		model2.Time2 = carbon.NewTime(c)
		model2.Time3 = carbon.NewTime(c)

		model2.DateTime1 = carbon.NewDateTime(c)
		model2.DateTime2 = carbon.NewDateTime(c)
		model2.DateTime3 = carbon.NewDateTime(c)

		model2.RFC3339Layout1 = carbon.NewLayoutType[RFC3339Layout](c)
		model2.RFC3339Layout2 = carbon.NewLayoutType[RFC3339Layout](c)
		model2.RFC3339Layout3 = carbon.NewLayoutType[RFC3339Layout](c)

		model2.ISO8601Format1 = carbon.NewFormatType[ISO8601Format](c)
		model2.ISO8601Format2 = carbon.NewFormatType[ISO8601Format](c)
		model2.ISO8601Format3 = carbon.NewFormatType[ISO8601Format](c)

		model2.Timestamp1 = carbon.NewTimestamp(c)

		// update
		_, err = db.Update(model2)
		s.Nil(err)

		data2, err2 := json.Marshal(&model2)
		s.Nil(err2)
		s.Equal(`{"carbon1":"2020-08-06 13:14:15","carbon2":"2020-08-06 13:14:15","carbon3":"2020-08-06 13:14:15","date1":"2020-08-06","date2":"2020-08-06","date3":"2020-08-06","time1":"13:14:15","time2":"13:14:15","time3":"13:14:15","date_time1":"2020-08-06 13:14:15","date_time2":"2020-08-06 13:14:15","date_time3":"2020-08-06 13:14:15","rfc3339_layout1":"2020-08-06T13:14:15+08:00","rfc3339_layout2":"2020-08-06T13:14:15+08:00","rfc3339_layout3":"2020-08-06T13:14:15+08:00","iso8601_format1":"2020-08-06T13:14:15+08:00","iso8601_format2":"2020-08-06T13:14:15+08:00","iso8601_format3":"2020-08-06T13:14:15+08:00","timestamp1":1596690855}`, string(data2))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})
}

func (s *MySQLSuite) TestCurd2() {
	s.Run("unset carbon", func() {
		var model1 MySQLModel2

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 MySQLModel2
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"carbon3":null,"date1":null,"date2":null,"date3":null,"time1":null,"time2":null,"time3":null,"date_time1":null,"date_time2":null,"date_time3":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"rfc3339_layout3":null,"iso8601_format1":null,"iso8601_format2":null,"iso8601_format3":null,"timestamp1":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("zero carbon", func() {
		var model1 MySQLModel2

		c := carbon.NewCarbon()
		model1.Carbon1 = &c
		model1.Carbon2 = &c
		model1.Carbon3 = &c

		date := carbon.NewDate(c)
		model1.Date1 = &date
		model1.Date2 = &date
		model1.Date3 = &date

		time := carbon.NewTime(c)
		model1.Time1 = &time
		model1.Time2 = &time
		model1.Time3 = &time

		dateTime := carbon.NewDateTime(c)
		model1.DateTime1 = &dateTime
		model1.DateTime2 = &dateTime
		model1.DateTime3 = &dateTime

		rfc3339Layout := carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout1 = &rfc3339Layout
		model1.RFC3339Layout2 = &rfc3339Layout
		model1.RFC3339Layout3 = &rfc3339Layout

		iso8601Format := carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format1 = &iso8601Format
		model1.ISO8601Format2 = &iso8601Format
		model1.ISO8601Format3 = &iso8601Format

		timestamp := carbon.NewTimestamp(c)
		model1.Timestamp1 = &timestamp

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 MySQLModel2
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"carbon3":null,"date1":null,"date2":null,"date3":null,"time1":null,"time2":null,"time3":null,"date_time1":null,"date_time2":null,"date_time3":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"rfc3339_layout3":null,"iso8601_format1":null,"iso8601_format2":null,"iso8601_format3":null,"timestamp1":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("empty carbon", func() {
		var model1 MySQLModel2

		c := carbon.Parse("")
		model1.Carbon1 = &c
		model1.Carbon2 = &c
		model1.Carbon3 = &c

		date := carbon.NewDate(c)
		model1.Date1 = &date
		model1.Date2 = &date
		model1.Date3 = &date

		time := carbon.NewTime(c)
		model1.Time1 = &time
		model1.Time2 = &time
		model1.Time3 = &time

		dateTime := carbon.NewDateTime(c)
		model1.DateTime1 = &dateTime
		model1.DateTime2 = &dateTime
		model1.DateTime3 = &dateTime

		rfc3339Layout := carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout1 = &rfc3339Layout
		model1.RFC3339Layout2 = &rfc3339Layout
		model1.RFC3339Layout3 = &rfc3339Layout

		iso8601Format := carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format1 = &iso8601Format
		model1.ISO8601Format2 = &iso8601Format
		model1.ISO8601Format3 = &iso8601Format

		timestamp := carbon.NewTimestamp(c)
		model1.Timestamp1 = &timestamp

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 MySQLModel1
		_, err = db.Delete(model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":null,"carbon2":null,"carbon3":null,"date1":null,"date2":null,"date3":null,"time1":null,"time2":null,"time3":null,"date_time1":null,"date_time2":null,"date_time3":null,"rfc3339_layout1":null,"rfc3339_layout2":null,"rfc3339_layout3":null,"iso8601_format1":null,"iso8601_format2":null,"iso8601_format3":null,"timestamp1":null}`, string(data1))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})

	s.Run("valid carbon", func() {
		var model1 MySQLModel2

		c := carbon.Now()
		model1.Carbon1 = &c
		model1.Carbon2 = &c
		model1.Carbon3 = &c

		date := carbon.NewDate(c)
		model1.Date1 = &date
		model1.Date2 = &date
		model1.Date3 = &date

		time := carbon.NewTime(c)
		model1.Time1 = &time
		model1.Time2 = &time
		model1.Time3 = &time

		dateTime := carbon.NewDateTime(c)
		model1.DateTime1 = &dateTime
		model1.DateTime2 = &dateTime
		model1.DateTime3 = &dateTime

		rfc3339Layout := carbon.NewLayoutType[RFC3339Layout](c)
		model1.RFC3339Layout1 = &rfc3339Layout
		model1.RFC3339Layout2 = &rfc3339Layout
		model1.RFC3339Layout3 = &rfc3339Layout

		iso8601Format := carbon.NewFormatType[ISO8601Format](c)
		model1.ISO8601Format1 = &iso8601Format
		model1.ISO8601Format2 = &iso8601Format
		model1.ISO8601Format3 = &iso8601Format

		timestamp := carbon.NewTimestamp(c)
		model1.Timestamp1 = &timestamp

		// create
		_, err = db.Insert(model1)
		s.Nil(err)

		// read
		var model2 MySQLModel2
		_, err = db.Desc("id").Get(&model2)
		s.Nil(err)

		data1, err1 := json.Marshal(&model2)
		s.Nil(err1)
		s.Equal(`{"carbon1":"2020-08-05 13:14:15","carbon2":"2020-08-05 13:14:15","carbon3":"2020-08-05 13:14:15","date1":"2020-08-05","date2":"2020-08-05","date3":"2020-08-05","time1":"13:14:15","time2":"13:14:15","time3":"13:14:15","date_time1":"2020-08-05 13:14:15","date_time2":"2020-08-05 13:14:15","date_time3":"2020-08-05 13:14:15","rfc3339_layout1":"2020-08-05T13:14:15+08:00","rfc3339_layout2":"2020-08-05T13:14:15+08:00","rfc3339_layout3":"2020-08-05T13:14:15+08:00","iso8601_format1":"2020-08-05T13:14:15+08:00","iso8601_format2":"2020-08-05T13:14:15+08:00","iso8601_format3":"2020-08-05T13:14:15+08:00","timestamp1":1596604455}`, string(data1))

		c = c.AddDay()

		model2.Carbon1 = &c
		model2.Carbon2 = &c
		model2.Carbon3 = &c

		date = carbon.NewDate(c)
		model2.Date1 = &date
		model2.Date2 = &date
		model2.Date3 = &date

		time = carbon.NewTime(c)
		model2.Time1 = &time
		model2.Time2 = &time
		model2.Time3 = &time

		dateTime = carbon.NewDateTime(c)
		model2.DateTime1 = &dateTime
		model2.DateTime2 = &dateTime
		model2.DateTime3 = &dateTime

		rfc3339Layout = carbon.NewLayoutType[RFC3339Layout](c)
		model2.RFC3339Layout1 = &rfc3339Layout
		model2.RFC3339Layout2 = &rfc3339Layout
		model2.RFC3339Layout3 = &rfc3339Layout

		iso8601Format = carbon.NewFormatType[ISO8601Format](c)
		model2.ISO8601Format1 = &iso8601Format
		model2.ISO8601Format2 = &iso8601Format
		model2.ISO8601Format3 = &iso8601Format

		timestamp = carbon.NewTimestamp(c)
		model2.Timestamp1 = &timestamp

		// update
		_, err = db.Update(model2)
		s.Nil(err)

		data2, err2 := json.Marshal(&model2)
		s.Nil(err2)
		s.Equal(`{"carbon1":"2020-08-06 13:14:15","carbon2":"2020-08-06 13:14:15","carbon3":"2020-08-06 13:14:15","date1":"2020-08-06","date2":"2020-08-06","date3":"2020-08-06","time1":"13:14:15","time2":"13:14:15","time3":"13:14:15","date_time1":"2020-08-06 13:14:15","date_time2":"2020-08-06 13:14:15","date_time3":"2020-08-06 13:14:15","rfc3339_layout1":"2020-08-06T13:14:15+08:00","rfc3339_layout2":"2020-08-06T13:14:15+08:00","rfc3339_layout3":"2020-08-06T13:14:15+08:00","iso8601_format1":"2020-08-06T13:14:15+08:00","iso8601_format2":"2020-08-06T13:14:15+08:00","iso8601_format3":"2020-08-06T13:14:15+08:00","timestamp1":1596690855}`, string(data2))

		// delete
		_, err = db.Delete(model2)
		s.Nil(err)
	})
}
