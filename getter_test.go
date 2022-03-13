package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_DaysInYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-08-05", 366},
		{"2021-08-05", 365},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DaysInYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DaysInMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-01-05", 31},
		{"2020-02-05", 29},
		{"2020-03-05", 31},
		{"2020-04-05", 30},
		{"2020-05-05", 31},
		{"2021-06-05", 30},
		{"2021-07-05", 31},
		{"2021-08-05", 31},
		{"2021-09-05", 30},
		{"2021-10-05", 31},
		{"2021-11-05", 30},
		{"2021-12-05", 31},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DaysInMonth(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_MonthOfYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-01-05", 1},
		{"2020-02-05", 2},
		{"2020-03-05", 3},
		{"2020-04-05", 4},
		{"2020-05-05", 5},
		{"2021-06-05", 6},
		{"2021-07-05", 7},
		{"2021-08-05", 8},
		{"2021-09-05", 9},
		{"2021-10-05", 10},
		{"2021-11-05", 11},
		{"2021-12-05", 12},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.MonthOfYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DayOfYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-01-01", 1},
		{"2020-01-31", 31},
		{"2020-08-05", 218},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DayOfYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DayOfMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-01-01", 1},
		{"2020-01-31", 31},
		{"2020-08-05", 5},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DayOfMonth(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DayOfWeek(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-08-03", 1},
		{"2020-08-04", 2},
		{"2020-08-05", 3},
		{"2020-08-06", 4},
		{"2020-08-07", 5},
		{"2020-08-08", 6},
		{"2020-08-09", 7},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DayOfWeek(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_WeekOfYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2021-01-01", 53},
		{"2021-02-01", 5},
		{"2021-03-01", 9},
		{"2021-04-01", 13},
		{"2021-05-01", 17},
		{"2021-06-01", 22},
		{"2021-07-01", 26},
		{"2021-08-01", 30},
		{"2021-09-01", 35},
		{"2021-10-01", 39},
		{"2021-11-01", 44},
		{"2021-12-01", 48},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.WeekOfYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_WeekOfMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2021-07-01", 1},
		{"2021-07-02", 1},
		{"2021-07-03", 1},
		{"2021-07-04", 1},
		{"2021-07-05", 2},
		{"2021-07-06", 2},
		{"2021-07-07", 2},
		{"2021-07-08", 2},
		{"2021-07-09", 2},
		{"2021-07-10", 2},
		{"2021-07-11", 2},
		{"2021-07-12", 3},
		{"2021-07-13", 3},
		{"2021-07-14", 3},
		{"2021-07-15", 3},
		{"2021-07-16", 3},
		{"2021-07-17", 3},
		{"2021-07-18", 3},
		{"2021-07-19", 4},
		{"2021-07-20", 4},
		{"2021-07-21", 4},
		{"2021-07-22", 4},
		{"2021-07-23", 4},
		{"2021-07-24", 4},
		{"2021-07-25", 4},
		{"2021-07-26", 5},
		{"2021-07-27", 5},
		{"2021-07-28", 5},
		{"2021-07-29", 5},
		{"2021-07-30", 5},
		{"2021-07-31", 5},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.WeekOfMonth(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Century(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-08-05", 21},
		{"2020-08-05 13:14:15", 21},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Century(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Decade(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2019-08-05", 10},
		{"2019-08-05 13:14:15", 10},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Decade(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Date(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input         string // 输入值
		expectedYear  int    // 期望值
		expectedMonth int    // 期望值
		expectedDay   int    // 期望值
	}{
		{"", 0, 0, 0},
		{"0", 0, 0, 0},
		{"0000-00-00", 0, 0, 0},
		{"00:00:00", 0, 0, 0},
		{"0000-00-00 00:00:00", 0, 0, 0},

		{"2020-08-05", 2020, 8, 5},
		{"2020-08-05 13:14:15", 2020, 8, 5},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		year, month, day := c.Date()
		assert.Equal(test.expectedYear, year, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.expectedMonth, int(month), "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.expectedDay, day, "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Year(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-08-05", 2020},
		{"2020-08-05 13:14:15", 2020},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Year(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Quarter(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-01-05", 1},
		{"2020-04-05", 2},
		{"2020-08-05", 3},
		{"2020-11-05", 4},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Quarter(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Month(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-08-05 13:14:15", 8},
		{"2020-08-05", 8},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Month(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Week(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", -1},
		{"0", -1},
		{"0000-00-00", -1},
		{"00:00:00", -1},
		{"0000-00-00 00:00:00", -1},

		{"2020-08-03", 1},
		{"2020-08-04", 2},
		{"2020-08-05", 3},
		{"2020-08-06", 4},
		{"2020-08-07", 5},
		{"2020-08-08", 6},
		{"2020-08-09", 0},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Week(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Day(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-08-05 13:14:15", 5},
		{"2020-08-05", 5},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Day(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Clock(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input          string // 输入值
		expectedHour   int    // 期望值
		expectedMinute int    // 期望值
		expectedSecond int    // 期望值
	}{
		{"", 0, 0, 0},
		{"0", 0, 0, 0},
		{"0000-00-00", 0, 0, 0},
		{"00:00:00", 0, 0, 0},
		{"0000-00-00 00:00:00", 0, 0, 0},

		{"2020-08-05 13:14:15", 13, 14, 15},
		{"2020-08-05", 0, 0, 0},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		hour, minute, second := c.Clock()
		assert.Equal(test.expectedHour, hour, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.expectedMinute, minute, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.expectedSecond, second, "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Hour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-08-05 13:14:15", 13},
		{"2020-08-05", 0},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Hour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Minute(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-08-05 13:14:15", 14},
		{"2020-08-05", 0},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Minute(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Second(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-08-05 13:14:15", 15},
		{"2020-08-05", 0},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Second(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Millisecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-08-05 13:14:15.999", 999},
		{"2020-08-05", 0},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Millisecond(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Microsecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-08-05 13:14:15", 0},
		{"2020-08-05", 0},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Microsecond(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Nanosecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-08-05 13:14:15", 0},
		{"2020-08-05", 0},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Nanosecond(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Timestamp(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int64  // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-01-01 13:14:15", 1577855655},
		{"2020-01-31 13:14:15", 1580447655},
		{"2020-02-01 13:14:15", 1580534055},
		{"2020-02-28 13:14:15", 1582866855},
		{"2020-02-29 13:14:15", 1582953255},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Timestamp(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_TimestampWithSecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int64  // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-01-01 13:14:15", 1577855655},
		{"2020-01-31 13:14:15", 1580447655},
		{"2020-02-01 13:14:15", 1580534055},
		{"2020-02-28 13:14:15", 1582866855},
		{"2020-02-29 13:14:15", 1582953255},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.TimestampWithSecond(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_TimestampWithMillisecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int64  // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-01-01 13:14:15", 1577855655000},
		{"2020-01-31 13:14:15", 1580447655000},
		{"2020-02-01 13:14:15", 1580534055000},
		{"2020-02-28 13:14:15", 1582866855000},
		{"2020-02-29 13:14:15", 1582953255000},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.TimestampWithMillisecond(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_TimestampWithMicrosecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int64  // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-01-01 13:14:15", 1577855655000000},
		{"2020-01-31 13:14:15", 1580447655000000},
		{"2020-02-01 13:14:15", 1580534055000000},
		{"2020-02-28 13:14:15", 1582866855000000},
		{"2020-02-29 13:14:15", 1582953255000000},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.TimestampWithMicrosecond(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_TimestampWithNanosecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int64  // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{"2020-01-01 13:14:15", 1577855655000000000},
		{"2020-01-31 13:14:15", 1580447655000000000},
		{"2020-02-01 13:14:15", 1580534055000000000},
		{"2020-02-28 13:14:15", 1582866855000000000},
		{"2020-02-29 13:14:15", 1582953255000000000},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.TimestampWithNanosecond(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Timezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{PRC, "CST"},
		{Tokyo, "JST"},
	}

	for index, test := range tests {
		c := Now(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Timezone(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Location(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{PRC, PRC},
		{Tokyo, Tokyo},
	}

	for index, test := range tests {
		c := Now(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Location(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Offset(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{PRC, 28800},
		{Tokyo, 32400},
	}

	for index, test := range tests {
		c := Now(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Offset(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Locale(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"en", "en"},
		{"zh-CN", "zh-CN"},
	}

	for index, test := range tests {
		c := SetLocale(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Locale(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Age(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected int    // 期望值
	}{
		{"", 0},
		{"0", 0},
		{"0000-00-00", 0},
		{"00:00:00", 0},
		{"0000-00-00 00:00:00", 0},

		{Now().AddYears(18).ToDateTimeString(), 0},
		{Now().SubYears(18).ToDateTimeString(), 18},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Age(), "Current test index is "+strconv.Itoa(index))
	}
}
