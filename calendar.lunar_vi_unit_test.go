package carbon

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViLunar_Animal(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		0:  {"", ""},
		1:  {"0", ""},
		2:  {"0000-00-00", ""},
		3:  {"00:00:00", ""},
		4:  {"0000-00-00 00:00:00", ""},
		5:  {"1581-01-01 00:00:00", "Thìn"},
		6:  {"1582-01-01 00:00:00", "Tỵ"},
		7:  {"1583-01-01 00:00:00", "Ngọ"},
		8:  {"2020-05-01", "Tý"},
		9:  {"2020-08-05", "Tý"},
		10: {"2021-07-07", "Sửu"},
		11: {"2010-08-05", "Dần"},
		12: {"2011-08-05", "Mão"},
		13: {"2012-08-05", "Thìn"},
		14: {"2013-08-05", "Tỵ"},
		15: {"2014-08-05", "Ngọ"},
		16: {"2015-08-05", "Mùi"},
		17: {"2016-08-05", "Thân"},
		18: {"2017-08-05", "Dậu"},
		19: {"2018-08-05", "Tuất"},
		20: {"2019-08-05", "Hợi"},
		21: {"2020-04-23", "Tý"},
		22: {"2020-05-23", "Tý"}, // special boundary
		23: {"2020-06-21", "Tý"}, // special boundary
		24: {"2020-08-05", "Tý"},
		25: {"2021-05-12", "Sửu"},
		26: {"2021-08-05", "Sửu"},
		27: {"2200-08-05", "Tý"},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().Animal(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_Festival(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected []VLunarFestival
	}{
		0: {"", []VLunarFestival{}},
		1: {"0", []VLunarFestival{}},
		2: {"0000-00-00", []VLunarFestival{}},
		3: {"00:00:00", []VLunarFestival{}},
		4: {"0000-00-00 00:00:00", []VLunarFestival{}},

		5:  {"2020-04-23", []VLunarFestival{}},
		6:  {"2024-02-10", []VLunarFestival{{Day: 1, Month: 1, Name: "Tết Nguyên Đán"}}},
		7:  {"2024-02-24", []VLunarFestival{{Day: 15, Month: 1, Name: "Rằm tháng Giêng"}}},
		8:  {"2024-04-18", []VLunarFestival{{Day: 10, Month: 3, Name: "Giỗ Tổ Hùng Vương"}}},
		9:  {"2024-05-22", []VLunarFestival{{Day: 15, Month: 4, Name: "Phật Đản"}}},
		10: {"2024-06-10", []VLunarFestival{{Day: 5, Month: 5, Name: "Lễ Đoan Ngọ"}}},
		11: {"2024-08-18", []VLunarFestival{{Day: 15, Month: 7, Name: "Vu Lan"}}},
		12: {"2024-09-17", []VLunarFestival{{Day: 15, Month: 8, Name: "Tết Trung Thu"}}},
		13: {"2025-01-22", []VLunarFestival{{Day: 23, Month: 12, Name: "Ông Táo chầu trời"}}},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().Festivals(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_DateTime(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input                                  string
		year, month, day, hour, minute, second int
	}{
		0: {"", 0, 0, 0, 0, 0, 0},
		1: {"0", 0, 0, 0, 0, 0, 0},
		2: {"0000-00-00", 0, 0, 0, 0, 0, 0},
		3: {"00:00:00", 0, 0, 0, 0, 0, 0},
		4: {"0000-00-00 00:00:00", 0, 0, 0, 0, 0, 0},

		5: {"2020-08-05 13:14:15", 2020, 6, 16, 13, 14, 15},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		year, month, day, hour, minute, second := c.VLunar().DateTime()
		assert.Equal(test.year, year, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.month, month, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.day, day, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.hour, hour, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.minute, minute, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.second, second, "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_Date(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input            string
		year, month, day int
	}{
		0: {"", 0, 0, 0},
		1: {"0", 0, 0, 0},
		2: {"0000-00-00", 0, 0, 0},
		3: {"00:00:00", 0, 0, 0},
		4: {"0000-00-00 00:00:00", 0, 0, 0},

		5: {"2020-08-05 13:14:15", 2020, 6, 16},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		year, month, day := c.VLunar().Date()
		assert.Equal(test.year, year, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.month, month, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.day, day, "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_Time(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input                string
		hour, minute, second int
	}{
		0: {"", 0, 0, 0},
		1: {"0", 0, 0, 0},
		2: {"0000-00-00", 0, 0, 0},
		3: {"00:00:00", 0, 0, 0},
		4: {"0000-00-00 00:00:00", 0, 0, 0},

		5: {"2020-08-05 13:14:15", 13, 14, 15},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		hour, minute, second := c.VLunar().Time()
		assert.Equal(test.hour, hour, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.minute, minute, "Current test index is "+strconv.Itoa(index))
		assert.Equal(test.second, second, "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_Year(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected int
	}{
		0: {"", 0},
		1: {"0", 0},
		2: {"0000-00-00", 0},
		3: {"00:00:00", 0},
		4: {"0000-00-00 00:00:00", 0},

		5:  {"2020-04-23", 2020},
		6:  {"2020-05-01", 2020},
		7:  {"2020-08-05", 2020},
		8:  {"2021-01-01", 2020},
		9:  {"2021-05-12", 2021},
		10: {"2021-07-07", 2021},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().Year(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_Month(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected int
	}{
		0: {"", 0},
		1: {"0", 0},
		2: {"0000-00-00", 0},
		3: {"00:00:00", 0},
		4: {"0000-00-00 00:00:00", 0},

		5:  {"2021-01-05", 11},
		6:  {"2021-02-05", 12},
		7:  {"2021-03-05", 1},
		8:  {"2021-04-05", 2},
		9:  {"2021-05-05", 3},
		10: {"2021-06-05", 4},
		11: {"2021-07-05", 5},
		12: {"2021-08-05", 6},
		13: {"2021-09-05", 7},
		14: {"2021-10-05", 8},
		15: {"2021-11-05", 10},
		16: {"2021-12-05", 11},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().Month(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_Day(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected int
	}{
		0: {"", 0},
		1: {"0", 0},
		2: {"0000-00-00", 0},
		3: {"00:00:00", 0},
		4: {"0000-00-00 00:00:00", 0},

		5:  {"2020-08-01", 12},
		6:  {"2020-08-02", 13},
		7:  {"2020-08-03", 14},
		8:  {"2020-08-04", 15},
		9:  {"2020-08-05", 16},
		10: {"2020-08-06", 17},
		11: {"2020-08-07", 18},
		12: {"2020-08-08", 19},
		13: {"2020-08-09", 20},
		14: {"2020-08-10", 21},
		15: {"2020-08-11", 22},
		16: {"2020-08-12", 23},
		17: {"2020-08-13", 24},
		18: {"2020-08-14", 25},
		19: {"2020-08-15", 26},
		20: {"2020-08-16", 27},
		21: {"2020-08-17", 28},
		22: {"2020-08-18", 29},
		23: {"2020-08-19", 1},
		24: {"2020-08-20", 2},
		25: {"2020-08-21", 3},
		26: {"2020-08-22", 4},
		27: {"2020-08-23", 5},
		28: {"2020-08-24", 6},
		29: {"2020-08-25", 7},
		30: {"2020-08-26", 8},
		31: {"2020-08-27", 9},
		32: {"2020-08-28", 10},
		33: {"2020-08-29", 11},
		34: {"2020-08-30", 12},
		35: {"2020-08-31", 13},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().Day(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_LeapMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected int
	}{
		0: {"", 0},
		1: {"0", 0},
		2: {"0000-00-00", 0},
		3: {"00:00:00", 0},
		4: {"0000-00-00 00:00:00", 0},

		5: {"2020-04-23", 4},
		6: {"2020-05-01", 4},
		7: {"2020-08-05", 4},
		8: {"2021-07-07", 0},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().LeapMonth(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_LuckHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		0: {"", ""},
		1: {"0", ""},
		2: {"0000-00-00", ""},
		3: {"00:00:00", ""},
		4: {"0000-00-00 00:00:00", ""},

		5:  {"2023-06-20", "101100110100"},
		6:  {"2023-03-23", "001011001101"},
		7:  {"2030-02-12", "110011010010"},
		8:  {"2004-11-30", "001101001011"},
		9:  {"1997-01-19", "101100110100"},
		10: {"2024-02-10", "001011001101"},
		11: {"2024-02-24", "110100101100"},
		12: {"2024-04-18", "110100101100"},
		13: {"2024-05-22", "001011001101"},
		14: {"2024-06-10", "010010110011"},
		15: {"2024-08-18", "110011010010"},
		16: {"2024-09-17", "110011010010"},
		17: {"2025-01-22", "101100110100"},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().LuckyHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_LuckHourDetails(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected []LuckyHourDetail
	}{
		0: {"", []LuckyHourDetail{}},
		1: {"0", []LuckyHourDetail{}},
		2: {"0000-00-00", []LuckyHourDetail{}},
		3: {"00:00:00", []LuckyHourDetail{}},
		4: {"0000-00-00 00:00:00", []LuckyHourDetail{}},

		5: {"2023-06-20", []LuckyHourDetail{
			{
				Chi:  "Tý",
				From: 23,
				To:   1,
			},
			{
				Chi:  "Dần",
				From: 3,
				To:   5,
			},
			{
				Chi:  "Mão",
				From: 5,
				To:   7,
			},
			{
				Chi:  "Ngọ",
				From: 11,
				To:   13,
			},
			{
				Chi:  "Mùi",
				From: 13,
				To:   15,
			},
			{
				Chi:  "Dậu",
				From: 17,
				To:   19,
			},
		}},
		6: {"2023-03-23", []LuckyHourDetail{
			{
				Chi:  "Dần",
				From: 3,
				To:   5,
			},
			{
				Chi:  "Thìn",
				From: 7,
				To:   9,
			},
			{
				Chi:  "Tỵ",
				From: 9,
				To:   11,
			},
			{
				Chi:  "Thân",
				From: 15,
				To:   17,
			},
			{
				Chi:  "Dậu",
				From: 17,
				To:   19,
			},
			{
				Chi:  "Hợi",
				From: 21,
				To:   23,
			},
		}},
		7: {"2030-02-12", []LuckyHourDetail{
			{
				Chi:  "Tý",
				From: 23,
				To:   1,
			},
			{
				Chi:  "Sửu",
				From: 1,
				To:   3,
			},
			{
				Chi:  "Thìn",
				From: 7,
				To:   9,
			},
			{
				Chi:  "Tỵ",
				From: 9,
				To:   11,
			},
			{
				Chi:  "Mùi",
				From: 13,
				To:   15,
			},
			{
				Chi:  "Tuất",
				From: 19,
				To:   21,
			},
		}},
		8: {"2004-11-30", []LuckyHourDetail{
			{
				Chi:  "Dần",
				From: 3,
				To:   5,
			},
			{
				Chi:  "Mão",
				From: 5,
				To:   7,
			},
			{
				Chi:  "Tỵ",
				From: 9,
				To:   11,
			},
			{
				Chi:  "Thân",
				From: 15,
				To:   17,
			},
			{
				Chi:  "Tuất",
				From: 19,
				To:   21,
			},
			{
				Chi:  "Hợi",
				From: 21,
				To:   23,
			},
		}},
		9: {"1997-01-19", []LuckyHourDetail{
			{
				Chi:  "Tý",
				From: 23,
				To:   1,
			},
			{
				Chi:  "Dần",
				From: 3,
				To:   5,
			},
			{
				Chi:  "Mão",
				From: 5,
				To:   7,
			},
			{
				Chi:  "Ngọ",
				From: 11,
				To:   13,
			},
			{
				Chi:  "Mùi",
				From: 13,
				To:   15,
			},
			{
				Chi:  "Dậu",
				From: 17,
				To:   19,
			},
		}},
		10: {"2024-02-10", []LuckyHourDetail{
			{
				Chi:  "Dần",
				From: 3,
				To:   5,
			},
			{
				Chi:  "Thìn",
				From: 7,
				To:   9,
			},
			{
				Chi:  "Tỵ",
				From: 9,
				To:   11,
			},
			{
				Chi:  "Thân",
				From: 15,
				To:   17,
			},
			{
				Chi:  "Dậu",
				From: 17,
				To:   19,
			},
			{
				Chi:  "Hợi",
				From: 21,
				To:   23,
			},
		}},
		11: {"2024-02-24", []LuckyHourDetail{
			{
				Chi:  "Tý",
				From: 23,
				To:   1,
			},
			{
				Chi:  "Sửu",
				From: 1,
				To:   3,
			},
			{
				Chi:  "Mão",
				From: 5,
				To:   7,
			},
			{
				Chi:  "Ngọ",
				From: 11,
				To:   13,
			},
			{
				Chi:  "Thân",
				From: 15,
				To:   17,
			},
			{
				Chi:  "Dậu",
				From: 17,
				To:   19,
			},
		}},
		12: {"2024-04-18", []LuckyHourDetail{
			{
				Chi:  "Tý",
				From: 23,
				To:   1,
			},
			{
				Chi:  "Sửu",
				From: 1,
				To:   3,
			},
			{
				Chi:  "Mão",
				From: 5,
				To:   7,
			},
			{
				Chi:  "Ngọ",
				From: 11,
				To:   13,
			},
			{
				Chi:  "Thân",
				From: 15,
				To:   17,
			},
			{
				Chi:  "Dậu",
				From: 17,
				To:   19,
			},
		}},
		13: {"2024-05-22", []LuckyHourDetail{
			{
				Chi:  "Dần",
				From: 3,
				To:   5,
			},
			{
				Chi:  "Thìn",
				From: 7,
				To:   9,
			},
			{
				Chi:  "Tỵ",
				From: 9,
				To:   11,
			},
			{
				Chi:  "Thân",
				From: 15,
				To:   17,
			},
			{
				Chi:  "Dậu",
				From: 17,
				To:   19,
			},
			{
				Chi:  "Hợi",
				From: 21,
				To:   23,
			},
		}},
		14: {"2024-06-10", []LuckyHourDetail{
			{
				Chi:  "Sửu",
				From: 1,
				To:   3,
			},
			{
				Chi:  "Thìn",
				From: 7,
				To:   9,
			},
			{
				Chi:  "Ngọ",
				From: 11,
				To:   13,
			},
			{
				Chi:  "Mùi",
				From: 13,
				To:   15,
			},
			{
				Chi:  "Tuất",
				From: 19,
				To:   21,
			},
			{
				Chi:  "Hợi",
				From: 21,
				To:   23,
			},
		}},
		15: {"2024-08-18", []LuckyHourDetail{
			{
				Chi:  "Tý",
				From: 23,
				To:   1,
			},
			{
				Chi:  "Sửu",
				From: 1,
				To:   3,
			},
			{
				Chi:  "Thìn",
				From: 7,
				To:   9,
			},
			{
				Chi:  "Tỵ",
				From: 9,
				To:   11,
			},
			{
				Chi:  "Mùi",
				From: 13,
				To:   15,
			},
			{
				Chi:  "Tuất",
				From: 19,
				To:   21,
			},
		}},
		16: {"2024-09-17", []LuckyHourDetail{
			{
				Chi:  "Tý",
				From: 23,
				To:   1,
			},
			{
				Chi:  "Sửu",
				From: 1,
				To:   3,
			},
			{
				Chi:  "Thìn",
				From: 7,
				To:   9,
			},
			{
				Chi:  "Tỵ",
				From: 9,
				To:   11,
			},
			{
				Chi:  "Mùi",
				From: 13,
				To:   15,
			},
			{
				Chi:  "Tuất",
				From: 19,
				To:   21,
			},
		}},
		17: {"2025-01-22", []LuckyHourDetail{
			{
				Chi:  "Tý",
				From: 23,
				To:   1,
			},
			{
				Chi:  "Dần",
				From: 3,
				To:   5,
			},
			{
				Chi:  "Mão",
				From: 5,
				To:   7,
			},
			{
				Chi:  "Ngọ",
				From: 11,
				To:   13,
			},
			{
				Chi:  "Mùi",
				From: 13,
				To:   15,
			},
			{
				Chi:  "Dậu",
				From: 17,
				To:   19,
			},
		}},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.True(reflect.DeepEqual(test.expected, c.VLunar().LuckyHourDetails()), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_SolarTerm(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected SolarTerm
	}{
		0: {"", SolarTerm{}},
		1: {"0", SolarTerm{}},
		2: {"0000-00-00", SolarTerm{}},
		3: {"00:00:00", SolarTerm{}},
		4: {"0000-00-00 00:00:00", SolarTerm{}},

		5:  {"2023-06-20", SolarTerm{Longitude: 75, Name: "Mang chủng"}},
		6:  {"2023-03-23", SolarTerm{Longitude: 0, Name: "Xuân phân"}},
		7:  {"2030-02-12", SolarTerm{Longitude: 315, Name: "Lập xuân"}},
		8:  {"2004-11-30", SolarTerm{Longitude: 240, Name: "Tiểu tuyết"}},
		9:  {"1997-01-19", SolarTerm{Longitude: 285, Name: "Tiểu hàn"}},
		10: {"2024-02-10", SolarTerm{Longitude: 315, Name: "Lập xuân"}},
		11: {"2024-02-24", SolarTerm{Longitude: 330, Name: "Vũ Thủy"}},
		12: {"2024-04-18", SolarTerm{Longitude: 15, Name: "Thanh minh"}},
		13: {"2024-05-22", SolarTerm{Longitude: 60, Name: "Tiểu mãn"}},
		14: {"2024-06-10", SolarTerm{Longitude: 75, Name: "Mang chủng"}},
		15: {"2024-08-18", SolarTerm{Longitude: 135, Name: "Lập thu"}},
		16: {"2024-09-17", SolarTerm{Longitude: 165, Name: "Bạch lộ"}},
		17: {"2025-01-22", SolarTerm{Longitude: 300, Name: "Đại hàn"}},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().SolarTerm(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_ToYearString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		0: {"", ""},
		1: {"0", ""},
		2: {"0000-00-00", ""},
		3: {"00:00:00", ""},
		4: {"0000-00-00 00:00:00", ""},

		5:  {"2020-04-23", "2020"},
		6:  {"2020-05-01", "2020"},
		7:  {"2020-08-05", "2020"},
		8:  {"2021-01-01", "2020"},
		9:  {"2021-05-12", "2021"},
		10: {"2021-07-07", "2021"},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().ToYearString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_ToMonthString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		0: {"", ""},
		1: {"0", ""},
		2: {"0000-00-00", ""},
		3: {"00:00:00", ""},
		4: {"0000-00-00 00:00:00", ""},

		5:  {"2020-01-01", "12"},
		6:  {"2020-02-01", "1"},
		7:  {"2020-03-01", "2"},
		8:  {"2020-04-01", "3"},
		9:  {"2020-04-23", "4"},
		10: {"2020-05-01", "4"},
		11: {"2020-06-01", "4"},
		12: {"2020-07-01", "5"},
		13: {"2020-07-07", "5"},
		14: {"2020-08-01", "6"},
		15: {"2020-09-01", "7"},
		16: {"2020-10-01", "8"},
		17: {"2020-11-01", "9"},
		18: {"2020-12-01", "10"},
		19: {"2021-01-01", "11"},
		20: {"2021-02-01", "12"},
		21: {"2021-05-12", "4"},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().ToMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_ToDayString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		0: {"", ""},
		1: {"0", ""},
		2: {"0000-00-00", ""},
		3: {"00:00:00", ""},
		4: {"0000-00-00 00:00:00", ""},

		5:  {"2020-01-01", "7"},
		6:  {"2020-02-01", "8"},
		7:  {"2020-03-01", "8"},
		8:  {"2020-04-01", "9"},
		9:  {"2020-04-23", "1"},
		10: {"2020-05-01", "9"},
		11: {"2020-06-01", "10"},
		12: {"2020-07-01", "11"},
		13: {"2020-08-01", "12"},
		14: {"2020-09-01", "14"},
		15: {"2020-10-01", "15"},
		16: {"2020-11-01", "16"},
		17: {"2020-12-01", "17"},
		18: {"2021-01-03", "21"},
		19: {"2021-01-05", "23"},
		20: {"2021-04-11", "30"},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().ToDayString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		0: {"", ""},
		1: {"0", ""},
		2: {"0000-00-00", ""},
		3: {"00:00:00", ""},
		4: {"0000-00-00 00:00:00", ""},

		5:  {"2020-01-01", "2019-12-07 00:00:00"},
		6:  {"2020-02-01", "2020-01-08 00:00:00"},
		7:  {"2020-03-01", "2020-02-08 00:00:00"},
		8:  {"2020-04-01", "2020-03-09 00:00:00"},
		9:  {"2020-04-23", "2020-04-01 00:00:00"},
		10: {"2020-05-01", "2020-04-09 00:00:00"},
		11: {"2020-06-01", "2020-04-10 00:00:00"},
		12: {"2020-07-01", "2020-05-11 00:00:00"},
		13: {"2020-08-01", "2020-06-12 00:00:00"},
		14: {"2020-09-01", "2020-07-14 00:00:00"},
		15: {"2020-10-01", "2020-08-15 00:00:00"},
		16: {"2020-11-01", "2020-09-16 00:00:00"},
		17: {"2020-12-01", "2020-10-17 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().String(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_ToDateString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		0: {"", ""},
		1: {"0", ""},
		2: {"0000-00-00", ""},
		3: {"00:00:00", ""},
		4: {"0000-00-00 00:00:00", ""},

		5:  {"2020-01-01", "Ngày 7 tháng 12 năm 2019"},
		6:  {"2020-02-01", "Ngày 8 tháng 1 năm 2020"},
		7:  {"2020-03-01", "Ngày 8 tháng 2 năm 2020"},
		8:  {"2020-04-01", "Ngày 9 tháng 3 năm 2020"},
		9:  {"2020-04-23", "Ngày 1 tháng 4 năm 2020"},
		10: {"2020-05-01", "Ngày 9 tháng 4 năm 2020"},
		11: {"2020-06-01", "Ngày 10 tháng 4 năm 2020"},
		12: {"2020-07-01", "Ngày 11 tháng 5 năm 2020"},
		13: {"2020-08-01", "Ngày 12 tháng 6 năm 2020"},
		14: {"2020-09-01", "Ngày 14 tháng 7 năm 2020"},
		15: {"2020-10-01", "Ngày 15 tháng 8 năm 2020"},
		16: {"2020-11-01", "Ngày 16 tháng 9 năm 2020"},
		17: {"2020-12-01", "Ngày 17 tháng 10 năm 2020"},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsLeapYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-04-23", true},
		6: {"2020-05-01", true},
		7: {"2020-08-05", true},
		8: {"2021-01-01", false},
		9: {"2021-07-07", false},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsLeapYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsLeapMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-04-23", true},
		6: {"2020-05-01", true},
		7: {"2020-08-05", false},
		8: {"2021-01-01", false},
		9: {"2021-07-07", false},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsLeapMonth(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsRatYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", true},
		6: {"2020-08-05", true},
		7: {"2021-01-01", true},
		8: {"2021-07-07", false},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsRatYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsOxYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", true},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsOxYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsTigerYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2022-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsTigerYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsRabbitYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2023-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsCatYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsDragonYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2024-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsDragonYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsSnakeYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2025-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsSnakeYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsHorseYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2026-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsHorseYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsGoatYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2027-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsGoatYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsMonkeyYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2028-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsMonkeyYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsRoosterYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2029-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsRoosterYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsDogYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2030-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsDogYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsPigYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-05-01", false},
		6: {"2020-08-05", false},
		7: {"2021-01-01", false},
		8: {"2021-07-07", false},
		9: {"2031-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input, HoChiMinh)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsPigYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_DoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected string
	}{
		0: {"", ""},
		1: {"0", ""},
		2: {"0000-00-00", ""},
		3: {"00:00:00", ""},
		4: {"0000-00-00 00:00:00", ""},

		5:  {"2020-01-05 23:23:45", "Tý"},
		6:  {"2020-01-05 00:59:45", "Tý"},
		7:  {"2020-02-05 01:00:00", "Sửu"},
		8:  {"2020-02-05 03:40:00", "Dần"},
		9:  {"2020-02-05 05:00:00", "Mão"},
		10: {"2020-02-05 07:30:00", "Thìn"},
		11: {"2020-02-05 09:00:00", "Tỵ"},
		12: {"2020-02-05 11:00:00", "Ngọ"},
		13: {"2020-02-05 13:00:00", "Mùi"},
		14: {"2020-02-05 15:00:00", "Thân"},
		15: {"2020-02-05 14:59:00", "Mùi"},
		16: {"2020-02-05 17:00:00", "Dậu"},
		17: {"2020-02-05 19:00:00", "Tuất"},
		18: {"2020-02-05 21:00:00", "Hợi"},
	}

	for index, test := range tests {
		c := SetTimezone(HoChiMinh).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().DoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsFirstDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 00:00:00", true},
		6: {"2020-04-19 00:59:59", true},
		7: {"2020-08-05 23:00:00", true},
		8: {"2020-08-05 01:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(HoChiMinh).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsFirstDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsSecondDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 01:00:00", true},
		6: {"2020-04-19 02:59:59", true},
		7: {"2020-08-05 03:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(HoChiMinh).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsSecondDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsThirdDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 03:00:00", true},
		6: {"2020-04-19 04:59:59", true},
		7: {"2020-08-05 05:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(HoChiMinh).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsThirdDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}
func TestViLunar_IsFourthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 05:00:00", true},
		6: {"2020-04-19 06:59:59", true},
		7: {"2020-08-05 07:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(HoChiMinh).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsFourthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsFifthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 07:00:00", true},
		6: {"2020-04-19 08:59:59", true},
		7: {"2020-08-05 09:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(HoChiMinh).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsFifthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsSixthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 09:00:00", true},
		6: {"2020-04-19 10:59:59", true},
		7: {"2020-08-05 11:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(HoChiMinh).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsSixthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsSeventhDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 11:00:00", true},
		6: {"2020-04-19 12:59:59", true},
		7: {"2020-08-05 13:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(HoChiMinh).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsSeventhDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsEighthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 13:00:00", true},
		6: {"2020-04-19 14:59:59", true},
		7: {"2020-08-05 15:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(HoChiMinh).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsEighthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsNinthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 15:00:00", true},
		6: {"2020-04-19 16:59:59", true},
		7: {"2020-08-05 17:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(HoChiMinh).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsNinthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsTenthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 17:00:00", true},
		6: {"2020-04-19 18:59:59", true},
		7: {"2020-08-05 19:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(HoChiMinh).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsTenthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsEleventhDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 19:00:00", true},
		6: {"2020-04-19 20:59:59", true},
		7: {"2020-08-05 21:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(HoChiMinh).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsEleventhDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestViLunar_IsTwelfthDoubleHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		expected bool
	}{
		0: {"", false},
		1: {"0", false},
		2: {"0000-00-00", false},
		3: {"00:00:00", false},
		4: {"0000-00-00 00:00:00", false},

		5: {"2020-03-21 21:00:00", true},
		6: {"2020-04-19 22:59:59", true},
		7: {"2020-08-05 23:00:00", false},
	}

	for index, test := range tests {
		c := SetTimezone(HoChiMinh).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.VLunar().IsTwelfthDoubleHour(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_ViLunar(t *testing.T) {
	c := CreateFromDate(1840, 1, 1, "xxx").VLunar()
	assert.NotNil(t, c.Error, "It should catch an exception in VLunar()")
}
