package carbon

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguage_SetLocale(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected error  // 期望值
	}{
		{"en", nil},
		{"zh-CN", nil},
	}

	for index, test := range tests {
		assert.ErrorIs(test.expected, NewLanguage().SetLocale(test.input), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLangError_SetLocale(t *testing.T) {
	locale, lang := "xxx", NewLanguage()
	expected := fmt.Errorf("invalid locale %q, please see the directory %q for all valid locales", locale, langDir)
	actual := lang.SetLocale(locale)
	assert.Equal(t, expected, actual, "It should catch an exception in SetLocale()")
}

func TestLanguage_SetResources1(t *testing.T) {
	assert := assert.New(t)

	lang := NewLanguage()
	resources := map[string]string{
		"seasons":  "spring|summer|autumn|winter",
		"year":     "1 yr|%d yrs",
		"month":    "1 mo|%d mos",
		"week":     "%dw",
		"day":      "%dd",
		"hour":     "%dh",
		"minute":   "%dm",
		"second":   "%ds",
		"now":      "just now",
		"ago":      "%s ago",
		"from_now": "in %s",
		"before":   "%s before",
		"after":    "%s after",
	}
	lang.SetResources(resources)

	tests := []struct {
		input1   string // 输入值
		input2   string // 输入值
		expected string // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "just now"},
		{"2020-08-05 13:14:15", "2021-08-05 13:14:15", "1 yr before"},
		{"2020-08-05 13:14:15", "2019-08-05 13:14:15", "1 yr after"},
		{"2020-08-05 13:14:15", "2030-08-05 13:14:15", "10 yrs before"},
		{"2020-08-05 13:14:15", "2010-08-05 13:14:15", "10 yrs after"},

		{"2020-08-05 13:14:15", "2020-09-05 13:14:15", "1 mo before"},
		{"2020-08-05 13:14:15", "2020-07-05 13:14:15", "1 mo after"},
		{"2020-08-05 13:14:15", "2021-06-05 13:14:15", "10 mos before"},
		{"2020-08-05 13:14:15", "2019-10-05 13:14:15", "10 mos after"},

		{"2020-08-05 13:14:15", "2020-08-06 13:14:15", "1d before"},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "1d after"},
		{"2020-08-05 13:14:15", "2020-08-15 13:14:15", "1w before"},
		{"2020-08-05 13:14:15", "2020-07-26 13:14:15", "1w after"},

		{"2020-08-05 13:14:15", "2020-08-05 14:14:15", "1h before"},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:15", "1h after"},
		{"2020-08-05 13:14:15", "2020-08-05 23:14:15", "10h before"},
		{"2020-08-05 13:14:15", "2020-08-05 03:14:15", "10h after"},

		{"2020-08-05 13:14:15", "2020-08-05 13:15:15", "1m before"},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:15", "1m after"},
		{"2020-08-05 13:14:15", "2020-08-05 13:24:15", "10m before"},
		{"2020-08-05 13:14:15", "2020-08-05 13:04:15", "10m after"},

		{"2020-08-05 13:14:15", "2020-08-05 13:14:16", "1s before"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:14", "1s after"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:25", "10s before"},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:05", "10s after"},
	}

	for index, test := range tests {
		c1 := Parse(test.input1)
		c2 := Parse(test.input2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.SetLanguage(lang).DiffForHumans(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLanguage_SetResources2(t *testing.T) {
	assert := assert.New(t)

	lang := NewLanguage()
	resources := map[string]string{
		"xxx": "xxx",
	}
	lang.SetResources(resources)
	lang.SetLocale("en")

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},

		{"0000-00-00 00:00:00", ""},
		{"2021-08-05 13:14:15", ""},
	}

	for index, test := range tests {
		assert.Equal(test.expected, Parse(test.input).SetLanguage(lang).DiffForHumans(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		assert.Equal(test.expected, Parse(test.input).SetLanguage(lang).Constellation(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		assert.Equal(test.expected, Parse(test.input).SetLanguage(lang).Season(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		assert.Equal(test.expected, Parse(test.input).SetLanguage(lang).ToWeekString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		assert.Equal(test.expected, Parse(test.input).SetLanguage(lang).ToShortWeekString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		assert.Equal(test.expected, Parse(test.input).SetLanguage(lang).ToMonthString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		assert.Equal(test.expected, Parse(test.input).SetLanguage(lang).ToShortMonthString(), "Current test index is "+strconv.Itoa(index))
	}
}
