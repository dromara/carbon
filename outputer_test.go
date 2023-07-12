package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05 13:14:15"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, Parse(test.input).String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05 13:14:15 +0800 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToMonthString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-05", "January"},
		{"2020-02-05", "February"},
		{"2020-03-05", "March"},
		{"2020-04-05", "April"},
		{"2020-05-05", "May"},
		{"2020-06-05", "June"},
		{"2020-07-05", "July"},
		{"2020-08-05", "August"},
		{"2020-09-05", "September"},
		{"2020-10-05", "October"},
		{"2020-11-05", "November"},
		{"2020-12-05", "December"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToShortMonthString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-05", "Jan"},
		{"2020-02-05", "Feb"},
		{"2020-03-05", "Mar"},
		{"2020-04-05", "Apr"},
		{"2020-05-05", "May"},
		{"2020-06-05", "Jun"},
		{"2020-07-05", "Jul"},
		{"2020-08-05", "Aug"},
		{"2020-09-05", "Sep"},
		{"2020-10-05", "Oct"},
		{"2020-11-05", "Nov"},
		{"2020-12-05", "Dec"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortMonthString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortMonthString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToWeekString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-01", "Saturday"},
		{"2020-08-02", "Sunday"},
		{"2020-08-03", "Monday"},
		{"2020-08-04", "Tuesday"},
		{"2020-08-05", "Wednesday"},
		{"2020-08-06", "Thursday"},
		{"2020-08-07", "Friday"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToWeekString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToWeekString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToShortWeekString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-01", "Sat"},
		{"2020-08-02", "Sun"},
		{"2020-08-03", "Mon"},
		{"2020-08-04", "Tue"},
		{"2020-08-05", "Wed"},
		{"2020-08-06", "Thu"},
		{"2020-08-07", "Fri"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortWeekString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortWeekString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToDayDateTimeString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, Aug 5, 2020 1:14 PM"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDayDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDayDateTimeString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToDateTimeString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05 13:14:15"},
		{"2020-08-05", "2020-08-05 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToDateTimeMilliString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05T13:14:15.999999999+08:00", "2020-08-05 13:14:15.999"},
		{"2020-08-05", "2020-08-05 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeMilliString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeMilliString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToDateTimeMicroString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05T13:14:15.999999999+08:00", "2020-08-05 13:14:15.999999"},
		{"2020-08-05", "2020-08-05 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeMicroString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeMicroString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToDateTimeNanoString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05T13:14:15.999999999+08:00", "2020-08-05 13:14:15.999999999"},
		{"2020-08-05", "2020-08-05 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeNanoString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeNanoString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToShortDateTimeString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05T13:14:15.999999999+08:00", "20200805131415"},
		{"2020-08-05", "20200805000000"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortDateTimeString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToShortDateTimeMilliString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05T13:14:15.999999999+08:00", "20200805131415.999"},
		{"2020-08-05", "20200805000000"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortDateTimeMilliString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortDateTimeMilliString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToShortDateTimeMicroString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05T13:14:15.999999999+08:00", "20200805131415.999999"},
		{"2020-08-05", "20200805000000"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortDateTimeMicroString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortDateTimeMicroString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToShortDateTimeNanoString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05T13:14:15.999999999+08:00", "20200805131415.999999999"},
		{"2020-08-05", "20200805000000"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortDateTimeNanoString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortDateTimeNanoString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToDateString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05"},
		{"2020-08-05", "2020-08-05"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToDateMilliString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15.999999999", "2020-08-05.999"},
		{"2020-08-05", "2020-08-05"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateMilliString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateMilliString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToDateMicroString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15.999999999", "2020-08-05.999999"},
		{"2020-08-05", "2020-08-05"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateMicroString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateMicroString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToDateNanoString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15.999999999", "2020-08-05.999999999"},
		{"2020-08-05", "2020-08-05"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateNanoString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateNanoString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToShortDateString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "20200805"},
		{"2020-08-05", "20200805"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortDateString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortDateString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToShortDateMilliString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15.999999999", "20200805.999"},
		{"2020-08-05", "20200805"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortDateMilliString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortDateMilliString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToShortDateNanoString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15.999999999", "20200805.999999999"},
		{"2020-08-05", "20200805"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortDateNanoString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortDateNanoString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToShortDateMicroString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15.999999999", "20200805.999999"},
		{"2020-08-05", "20200805"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortDateMicroString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortDateMicroString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToTimeString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "13:14:15"},
		{"2020-08-05", "00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToTimeString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToTimeString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToTimeMilliString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15.999999999", "13:14:15.999"},
		{"2020-08-05", "00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToTimeMilliString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToTimeMilliString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToTimeMicroString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15.999999999", "13:14:15.999999"},
		{"2020-08-05", "00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToTimeMicroString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToTimeMicroString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToTimeNanoString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15.999999999", "13:14:15.999999999"},
		{"2020-08-05", "00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToTimeNanoString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToTimeNanoString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToShortTimeString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "131415"},
		{"2020-08-05", "000000"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortTimeString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortTimeString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToShortTimeMilliString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15.999999999", "131415.999"},
		{"2020-08-05", "000000"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortTimeMilliString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortTimeMilliString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToShortTimeMicroString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15.999999999", "131415.999999"},
		{"2020-08-05", "000000"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortTimeMicroString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortTimeMicroString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToShortTimeNanoString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15.999999999", "131415.999999999"},
		{"2020-08-05", "000000"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortTimeNanoString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToShortTimeNanoString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToAtomString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05T13:14:15+08:00"},
		{"2020-08-05", "2020-08-05T00:00:00+08:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToAtomString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToAtomString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToANSICString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed Aug  5 13:14:15 2020"},
		{"2020-08-05", "Wed Aug  5 00:00:00 2020"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToANSICString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToANSICString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToCookieString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wednesday, 05-Aug-2020 13:14:15 CST"},
		{"2020-08-05", "Wednesday, 05-Aug-2020 00:00:00 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToCookieString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToCookieString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToRssString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 +0800"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 +0800"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRssString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRssString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToW3cString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05T13:14:15+08:00"},
		{"2020-08-05", "2020-08-05T00:00:00+08:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToW3cString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToW3cString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToUnixDateString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed Aug  5 13:14:15 CST 2020"},
		{"2020-08-05", "Wed Aug  5 00:00:00 CST 2020"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToUnixDateString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToUnixDateString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToRubyDateString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed Aug 05 13:14:15 +0800 2020"},
		{"2020-08-05", "Wed Aug 05 00:00:00 +0800 2020"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRubyDateString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRubyDateString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToKitchenString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "1:14PM"},
		{"2020-08-05", "12:00AM"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, test.expected, c.ToKitchenString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, test.expected, c.ToKitchenString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToIso8601String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05T13:14:15+08:00"},
		{"2020-08-05", "2020-08-05T00:00:00+08:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToIso8601String(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToIso8601String(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToIso8601MilliString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15.999999999", "2020-08-05T13:14:15.999+08:00"},
		{"2020-08-05", "2020-08-05T00:00:00+08:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToIso8601MilliString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToIso8601MilliString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToIso8601MicroString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15.999999999", "2020-08-05T13:14:15.999999+08:00"},
		{"2020-08-05", "2020-08-05T00:00:00+08:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToIso8601MicroString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToIso8601MicroString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToIso8601NanoString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15.999999999", "2020-08-05T13:14:15.999999999+08:00"},
		{"2020-08-05", "2020-08-05T00:00:00+08:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToIso8601NanoString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToIso8601NanoString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToRfc822String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "05 Aug 20 13:14 CST"},
		{"2020-08-05", "05 Aug 20 00:00 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc822String(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc822String(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToRfc822zString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "05 Aug 20 13:14 +0800"},
		{"2020-08-05", "05 Aug 20 00:00 +0800"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc822zString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc822zString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToRfc850String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wednesday, 05-Aug-20 13:14:15 CST"},
		{"2020-08-05", "Wednesday, 05-Aug-20 00:00:00 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc850String(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc850String(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToRfc1036String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 20 13:14:15 +0800"},
		{"2020-08-05", "Wed, 05 Aug 20 00:00:00 +0800"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc1036String(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc1036String(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToRfc1123String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 CST"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc1123String(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc1123String(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToRfc1123zString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 +0800"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 +0800"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc1123zString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc1123zString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToRfc2822String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 +0800"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 +0800"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc2822String(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc2822String(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToRfc3339String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05T13:14:15+08:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc3339String(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc3339String(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToRfc3339MilliString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05T13:14:15.999999999+08:00", "2020-08-05T13:14:15.999+08:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc3339MilliString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc3339MilliString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToRfc3339MicroString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05T13:14:15.999999999+08:00", "2020-08-05T13:14:15.999999+08:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc3339MicroString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc3339MicroString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToRfc3339NanoString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05T13:14:15.999999999+08:00", "2020-08-05T13:14:15.999999999+08:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc3339NanoString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc3339NanoString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ToRfc7231String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "Wed, 05 Aug 2020 13:14:15 CST"},
		{"2020-08-05", "Wed, 05 Aug 2020 00:00:00 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc7231String(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToRfc7231String(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Layout(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected string // 期望值
	}{
		{"", "2006年01月02日", ""},
		{"0", "2006年01月02日", ""},
		{"0000-00-00", "2006年01月02日", ""},
		{"00:00:00", "2006年01月02日", ""},
		{"0000-00-00 00:00:00", "Y年m月d日", ""},

		{"2020-08-05 13:14:15", "2006年01月02日", "2020年08月05日"},
		{"2020-08-05 13:14:15", "Mon, 02 Jan 2006 15:04:05 GMT", "Wed, 05 Aug 2020 13:14:15 GMT"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Layout(test.param), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToLayoutString(test.param), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Layout(test.param, PRC), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToLayoutString(test.param, PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Format(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		format   string // 参数值
		locale   string // 参数值
		expected string // 期望值
	}{
		{"", "Y年m月d日", "en", ""},
		{"0", "Y年m月d日", "en", ""},
		{"0000-00-00", "Y年m月d日", "en", ""},
		{"00:00:00", "Y年m月d日", "en", ""},
		{"0000-00-00 00:00:00", "Y年m月d日", "en", ""},

		{"2020-08-05 01:14:15", "D", "de", "Mi"},
		{"2020-08-05 01:14:15", "l", "ru", "Среда"},
		{"2020-08-05 01:14:15", "F", "jp", "はちがつ"},
		{"2020-08-05 01:14:15", "M", "zh-CN", "8月"},

		{"2020-08-05 13:14:15", "Y年m月d日", "en", "2020年08月05日"},
		{"2020-08-05 01:14:15", "j", "en", "5"},
		{"2020-08-05 01:14:15", "W", "en", "32"},
		{"2020-08-05 01:14:15", "M", "en", "Aug"},
		{"2020-08-05 01:14:15", "F", "en", "August"},
		{"2020-08-05 01:14:15", "N", "en", "03"},
		{"2020-08-05 01:14:15", "L", "en", "1"},
		{"2020-08-05 01:14:15", "L", "en", "1"},
		{"2021-08-05 01:14:15", "L", "en", "0"},
		{"2020-08-05 01:14:15", "G", "en", "1"},
		{"2020-08-05 13:14:15", "U", "en", "1596604455"},
		{"2020-08-05 13:14:15.999", "u", "en", "999"},
		{"2020-08-05 13:14:15", "w", "en", "2"},
		{"2020-08-05 13:14:15", "t", "en", "31"},
		{"2020-08-05 13:14:15", "z", "en", "217"},
		{"2020-08-05 13:14:15", "e", "en", "PRC"},
		{"2020-08-05 13:14:15", "Q", "en", "3"},
		{"2020-08-05 13:14:15", "C", "en", "21"},
		{"2020-08-05 13:14:15", "jS", "en", "5th"},
		{"2020-08-22 13:14:15", "jS", "en", "22nd"},
		{"2020-08-23 13:14:15", "jS", "en", "23rd"},
		{"2020-08-31 13:14:15", "jS", "en", "31st"},
		{"2020-08-31 13:14:15", "I\\t \\i\\s Y-m-d H:i:s", "en", "It is 2020-08-31 13:14:15"},
		{"2020-08-05 13:14:15", "上次上报时间:Y-m-d H:i:s，请每日按时打卡", "en", "上次上报时间:2020-08-05 13:14:15，请每日按时打卡"},
		{"2020-08-05 13:14:15", "l jS of F Y h:i:s A", "en", "Wednesday 5th of August 2020 01:14:15 PM"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.SetLocale(test.locale).Format(test.format), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.SetLocale(test.locale).ToFormatString(test.format), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.SetLocale(test.locale).ToFormatString(test.format, PRC), "Current test index is "+strconv.Itoa(index))
	}
}
