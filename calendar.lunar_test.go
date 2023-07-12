package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLunar_Animal(t *testing.T) {
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

		{"2020-05-01", "鼠"},
		{"2020-08-05", "鼠"},
		{"2021-07-07", "牛"},
		{"2010-08-05", "虎"},
		{"2011-08-05", "兔"},
		{"2012-08-05", "龙"},
		{"2013-08-05", "蛇"},
		{"2014-08-05", "马"},
		{"2015-08-05", "羊"},
		{"2016-08-05", "猴"},
		{"2017-08-05", "鸡"},
		{"2018-08-05", "狗"},
		{"2019-08-05", "猪"},
		{"2020-04-23", "鼠"},
		{"2020-05-23", "鼠"}, // 特殊边界值
		{"2020-06-21", "鼠"}, // 特殊边界值
		{"2020-08-05", "鼠"},
		{"2021-05-12", "牛"},
		{"2021-08-05", "牛"},
		{"2200-08-05", ""},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().Animal(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_Festival(t *testing.T) {
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

		{"2020-04-23", ""},
		{"2021-02-12", "春节"},
		{"2021-02-26", "元宵节"},
		{"2021-05-12", ""},
		{"2021-06-14", "端午节"},
		{"2021-08-14", "七夕节"},
		{"2021-08-22", "中元节"},
		{"2021-09-21", "中秋节"},
		{"2021-10-14", "重阳节"},
		{"2021-10-14", "重阳节"},
		{"2021-11-05", "寒衣节"},
		{"2021-11-19", "下元节"},
		{"2022-01-10", "腊八节"},
		{"2022-01-25", "小年"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().Festival(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_DateTime(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input                                  string // 输入值
		year, month, day, hour, minute, second int    // 期望值
	}{
		{"", 0, 0, 0, 0, 0, 0},
		{"0", 0, 0, 0, 0, 0, 0},
		{"0000-00-00", 0, 0, 0, 0, 0, 0},
		{"00:00:00", 0, 0, 0, 0, 0, 0},
		{"0000-00-00 00:00:00", 0, 0, 0, 0, 0, 0},

		{"2020-08-05 13:14:15", 2020, 6, 16, 13, 14, 15},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		year, month, day, hour, minute, second := c.Lunar().DateTime()
		assert.Equal(test.year, year, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.month, month, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.day, day, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.hour, hour, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.minute, minute, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.second, second, "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_Date(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input            string // 输入值
		year, month, day int    // 期望值
	}{
		{"", 0, 0, 0},
		{"0", 0, 0, 0},
		{"0000-00-00", 0, 0, 0},
		{"00:00:00", 0, 0, 0},
		{"0000-00-00 00:00:00", 0, 0, 0},

		{"2020-08-05 13:14:15", 2020, 6, 16},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		year, month, day := c.Lunar().Date()
		assert.Equal(test.year, year, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.month, month, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.day, day, "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_Time(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input                string // 输入值
		hour, minute, second int    // 期望值
	}{
		{"", 0, 0, 0},
		{"0", 0, 0, 0},
		{"0000-00-00", 0, 0, 0},
		{"00:00:00", 0, 0, 0},
		{"0000-00-00 00:00:00", 0, 0, 0},

		{"2020-08-05 13:14:15", 13, 14, 15},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		hour, minute, second := c.Lunar().Time()
		assert.Equal(test.hour, hour, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.minute, minute, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.second, second, "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_Year(t *testing.T) {
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

		{"2020-04-23", 2020},
		{"2020-05-01", 2020},
		{"2020-08-05", 2020},
		{"2021-01-01", 2020},
		{"2021-05-12", 2021},
		{"2021-07-07", 2021},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().Year(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_Month(t *testing.T) {
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

		{"2021-01-05", 11},
		{"2021-02-05", 12},
		{"2021-03-05", 1},
		{"2021-04-05", 2},
		{"2021-05-05", 3},
		{"2021-06-05", 4},
		{"2021-07-05", 5},
		{"2021-08-05", 6},
		{"2021-09-05", 7},
		{"2021-10-05", 8},
		{"2021-11-05", 10},
		{"2021-12-05", 11},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().Month(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_Day(t *testing.T) {
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

		{"2020-08-01", 12},
		{"2020-08-02", 13},
		{"2020-08-03", 14},
		{"2020-08-04", 15},
		{"2020-08-05", 16},
		{"2020-08-06", 17},
		{"2020-08-07", 18},
		{"2020-08-08", 19},
		{"2020-08-09", 20},
		{"2020-08-10", 21},
		{"2020-08-11", 22},
		{"2020-08-12", 23},
		{"2020-08-13", 24},
		{"2020-08-14", 25},
		{"2020-08-15", 26},
		{"2020-08-16", 27},
		{"2020-08-17", 28},
		{"2020-08-18", 29},
		{"2020-08-19", 1},
		{"2020-08-20", 2},
		{"2020-08-21", 3},
		{"2020-08-22", 4},
		{"2020-08-23", 5},
		{"2020-08-24", 6},
		{"2020-08-25", 7},
		{"2020-08-26", 8},
		{"2020-08-27", 9},
		{"2020-08-28", 10},
		{"2020-08-29", 11},
		{"2020-08-30", 12},
		{"2020-08-31", 13},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().Day(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_LeapMonth(t *testing.T) {
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

		{"2020-04-23", 4},
		{"2020-05-01", 4},
		{"2020-08-05", 4},
		{"2021-07-07", 0},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().LeapMonth(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_ToYearString(t *testing.T) {
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

		{"2020-04-23", "二零二零"},
		{"2020-05-01", "二零二零"},
		{"2020-08-05", "二零二零"},
		{"2021-01-01", "二零二零"},
		{"2021-05-12", "二零二一"},
		{"2021-07-07", "二零二一"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().ToYearString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_ToMonthString(t *testing.T) {
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

		{"2020-01-01", "腊月"},
		{"2020-02-01", "正月"},
		{"2020-03-01", "二月"},
		{"2020-04-01", "三月"},
		{"2020-04-23", "四月"},
		{"2020-05-01", "四月"},
		{"2020-06-01", "四月"},
		{"2020-07-01", "五月"},
		{"2020-07-07", "五月"},
		{"2020-08-01", "六月"},
		{"2020-09-01", "七月"},
		{"2020-10-01", "八月"},
		{"2020-11-01", "九月"},
		{"2020-12-01", "十月"},
		{"2021-01-01", "十一月"},
		{"2021-02-01", "腊月"},
		{"2021-05-12", "四月"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().ToMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_ToDayString(t *testing.T) {
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

		{"2020-01-01", "初七"},
		{"2020-02-01", "初八"},
		{"2020-03-01", "初八"},
		{"2020-04-01", "初九"},
		{"2020-04-23", "初一"},
		{"2020-05-01", "初九"},
		{"2020-06-01", "初十"},
		{"2020-07-01", "十一"},
		{"2020-08-01", "十二"},
		{"2020-09-01", "十四"},
		{"2020-10-01", "十五"},
		{"2020-11-01", "十六"},
		{"2020-12-01", "十七"},
		{"2021-01-03", "二十"},
		{"2021-01-05", "廿二"},
		{"2021-04-11", "三十"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().ToDayString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_String(t *testing.T) {
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

		{"2020-01-01", "2019-12-07 00:00:00"},
		{"2020-02-01", "2020-01-08 00:00:00"},
		{"2020-03-01", "2020-02-08 00:00:00"},
		{"2020-04-01", "2020-03-09 00:00:00"},
		{"2020-04-23", "2020-04-01 00:00:00"},
		{"2020-05-01", "2020-04-09 00:00:00"},
		{"2020-06-01", "2020-04-10 00:00:00"},
		{"2020-07-01", "2020-05-11 00:00:00"},
		{"2020-08-01", "2020-06-12 00:00:00"},
		{"2020-09-01", "2020-07-14 00:00:00"},
		{"2020-10-01", "2020-08-15 00:00:00"},
		{"2020-11-01", "2020-09-16 00:00:00"},
		{"2020-12-01", "2020-10-17 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_ToDateString(t *testing.T) {
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

		{"2020-01-01", "二零一九年腊月初七"},
		{"2020-02-01", "二零二零年正月初八"},
		{"2020-03-01", "二零二零年二月初八"},
		{"2020-04-01", "二零二零年三月初九"},
		{"2020-04-23", "二零二零年四月初一"},
		{"2020-05-01", "二零二零年四月初九"},
		{"2020-06-01", "二零二零年四月初十"},
		{"2020-07-01", "二零二零年五月十一"},
		{"2020-08-01", "二零二零年六月十二"},
		{"2020-09-01", "二零二零年七月十四"},
		{"2020-10-01", "二零二零年八月十五"},
		{"2020-11-01", "二零二零年九月十六"},
		{"2020-12-01", "二零二零年十月十七"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsLeapYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-04-23", true},
		{"2020-05-01", true},
		{"2020-08-05", true},
		{"2021-01-01", true},
		{"2021-07-07", false},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsLeapYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsLeapMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-04-23", true},
		{"2020-05-01", true},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsLeapMonth(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsRatYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-05-01", true},
		{"2020-08-05", true},
		{"2021-01-01", true},
		{"2021-07-07", false},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsRatYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsOxYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", true},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsOxYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsTigerYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2022-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsTigerYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsRabbitYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2023-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsRabbitYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsDragonYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2024-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsDragonYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsSnakeYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2025-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsSnakeYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsHorseYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2026-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsHorseYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsGoatYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2027-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsGoatYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsMonkeyYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2028-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsMonkeyYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsRoosterYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2029-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsRoosterYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsDogYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2030-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsDogYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsPigYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-05-01", false},
		{"2020-08-05", false},
		{"2021-01-01", false},
		{"2021-07-07", false},
		{"2031-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsPigYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_DoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值1
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-05 23:23:45", "子时"},
		{"2020-01-05 00:59:45", "子时"},
		{"2020-02-05 01:00:00", "丑时"},
		{"2020-02-05 03:40:00", "寅时"},
		{"2020-02-05 05:00:00", "卯时"},
		{"2020-02-05 07:30:00", "辰时"},
		{"2020-02-05 09:00:00", "巳时"},
		{"2020-02-05 11:00:00", "午时"},
		{"2020-02-05 13:00:00", "未时"},
		{"2020-02-05 15:00:00", "申时"},
		{"2020-02-05 14:59:00", "未时"},
		{"2020-02-05 17:00:00", "酉时"},
		{"2020-02-05 19:00:00", "戌时"},
		{"2020-02-05 21:00:00", "亥时"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().DoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsFirstDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 00:00:00", true},
		{"2020-04-19 00:59:59", true},
		{"2020-08-05 23:00:00", true},
		{"2020-08-05 01:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsFirstDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsSecondDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 01:00:00", true},
		{"2020-04-19 02:59:59", true},
		{"2020-08-05 03:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsSecondDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsThirdDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 03:00:00", true},
		{"2020-04-19 04:59:59", true},
		{"2020-08-05 05:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsThirdDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}
func TestLunar_IsFourthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 05:00:00", true},
		{"2020-04-19 06:59:59", true},
		{"2020-08-05 07:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsFourthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsFifthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 07:00:00", true},
		{"2020-04-19 08:59:59", true},
		{"2020-08-05 09:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsFifthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsSixthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 09:00:00", true},
		{"2020-04-19 10:59:59", true},
		{"2020-08-05 11:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsSixthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsSeventhDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 11:00:00", true},
		{"2020-04-19 12:59:59", true},
		{"2020-08-05 13:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsSeventhDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsEighthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 13:00:00", true},
		{"2020-04-19 14:59:59", true},
		{"2020-08-05 15:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsEighthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsNinthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 15:00:00", true},
		{"2020-04-19 16:59:59", true},
		{"2020-08-05 17:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsNinthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsTenthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 17:00:00", true},
		{"2020-04-19 18:59:59", true},
		{"2020-08-05 19:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsTenthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsEleventhDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 19:00:00", true},
		{"2020-04-19 20:59:59", true},
		{"2020-08-05 21:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsEleventhDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLunar_IsTwelfthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-03-21 21:00:00", true},
		{"2020-04-19 22:59:59", true},
		{"2020-08-05 23:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Lunar().IsTwelfthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_Lunar(t *testing.T) {
	year, month, day, timezone := 1840, 1, 1, "xxx"
	c := CreateFromDate(year, month, day, timezone).Lunar()
	assert.NotNil(t, c.Error, "It should catch an exception in Lunar()")
}
