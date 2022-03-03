package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_CreateFromTimestamp(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		timestamp int64  // 输入参数
		expected  string // 期望值
	}{
		{-1, "1970-01-01 07:59:59"},
		{0, "1970-01-01 08:00:00"},
		{1, "1970-01-01 08:00:01"},
		{1577855655, "2020-01-01 13:14:15"},
		{1604074084682, "2020-10-31 00:08:04"},
		{1604074196366540, "2020-10-31 00:09:56"},
		{1604074298500312000, "2020-10-31 00:11:38"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromTimestamp(test.timestamp)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromTimestamp(test.timestamp, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_CreateFromTimestamp(t *testing.T) {
	timestamp, timezone := int64(1577855655), "xxx"
	c1 := SetTimezone(timezone).CreateFromTimestamp(timestamp)
	assert.NotNil(t, c1.Error, "It should catch an exception in CreateFromTimestamp()")

	c2 := CreateFromTimestamp(timestamp, timezone)
	assert.NotNil(t, c2.Error, "It should catch an exception in CreateFromTimestamp()")
}

func TestCarbon_CreateFromDateTime(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		year, month, day, hour, minute, second, nsecond int    // 输入参数
		expected                                        string // 期望值
	}{
		{2020, 1, 1, 13, 14, 15, 0, "2020-01-01 13:14:15"},
		{2020, 1, 31, 13, 14, 15, 0, "2020-01-31 13:14:15"},
		{2020, 2, 1, 13, 14, 15, 0, "2020-02-01 13:14:15"},
		{2020, 2, 28, 13, 14, 15, 0, "2020-02-28 13:14:15"},
		{2020, 2, 29, 13, 14, 15, 0, "2020-02-29 13:14:15"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromDateTime(test.year, test.month, test.day, test.hour, test.minute, test.second, test.nsecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromDateTime(test.year, test.month, test.day, test.hour, test.minute, test.second, test.nsecond, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_CreateFromDateTime(t *testing.T) {
	year, month, day, hour, minute, second, nsecond, timezone := 2020, 1, 1, 13, 14, 15, 0, "xxx"
	c1 := SetTimezone(timezone).CreateFromDateTime(year, month, day, hour, minute, second, nsecond)
	assert.NotNil(t, c1.Error, "It should catch an exception in CreateFromDateTime()")

	c2 := CreateFromDateTime(year, month, day, hour, minute, second, nsecond, timezone)
	assert.NotNil(t, c2.Error, "It should catch an exception in CreateFromDateTime()")
}

func TestCarbon_CreateFromDate(t *testing.T) {
	assert := assert.New(t)

	clock := Now(PRC).ToTimeString()
	tests := []struct {
		year, month, day int    // 输入参数
		expected         string // 期望值
	}{
		{2020, 1, 1, "2020-01-01 " + clock},
		{2020, 1, 31, "2020-01-31 " + clock},
		{2020, 2, 1, "2020-02-01 " + clock},
		{2020, 2, 28, "2020-02-28 " + clock},
		{2020, 2, 29, "2020-02-29 " + clock},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromDate(test.year, test.month, test.day)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromDate(test.year, test.month, test.day, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_CreateFromDate(t *testing.T) {
	year, month, day, timezone := 13, 14, 15, "xxx"
	c1 := SetTimezone(timezone).CreateFromDate(year, month, day)
	assert.NotNil(t, c1.Error, "It should catch an exception in CreateFromDate()")

	c2 := CreateFromDate(year, month, day, timezone)
	assert.NotNil(t, c2.Error, "It should catch an exception in CreateFromDate()")
}

func TestCarbon_CreateFromTime(t *testing.T) {
	assert := assert.New(t)

	date := Now(PRC).ToDateString()
	tests := []struct {
		hour, minute, second int    // 输入参数
		expected             string // 期望值
	}{
		{0, 0, 0, date + " 00:00:00"},
		{00, 00, 15, date + " 00:00:15"},
		{00, 14, 15, date + " 00:14:15"},
		{13, 14, 15, date + " 13:14:15"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromTime(test.hour, test.minute, test.second)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromTime(test.hour, test.minute, test.second, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_CreateFromTime(t *testing.T) {
	hour, minute, second, timezone := 13, 14, 15, "xxx"
	c1 := SetTimezone(timezone).CreateFromTime(hour, minute, second)
	assert.NotNil(t, c1.Error, "It should catch an exception in CreateFromTime()")

	c2 := CreateFromTime(hour, minute, second, timezone)
	assert.NotNil(t, c2.Error, "It should catch an exception in CreateFromTime()")
}
