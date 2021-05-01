package carbon

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Parse 将标准格式时间字符串解析成 Carbon 实例
func (c Carbon) Parse(value string) Carbon {
	if c.Error != nil {
		return c
	}
	layout := DateTimeFormat
	if value == "" || value == "0" || value == "0000-00-00 00:00:00" || value == "0000-00-00" || value == "00:00:00" {
		return c
	}
	if len(value) == 10 && strings.Count(value, "-") == 2 {
		layout = DateFormat
	}
	if strings.Index(value, "T") == 10 {
		layout = RFC3339Format
	}
	if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		switch len(value) {
		case 8:
			layout = ShortDateFormat
		case 14:
			layout = ShortDateTimeFormat
		}
	}
	return c.ParseByLayout(value, layout)
}

// Parse 将标准格式时间字符串解析成 Carbon 实例(默认时区)
func Parse(value string) Carbon {
	return SetTimezone(Local).Parse(value)
}

// ParseByFormat 将特殊格式时间字符串解析成 Carbon 实例
func (c Carbon) ParseByFormat(value string, format string) Carbon {
	if c.Error != nil {
		return c
	}
	layout := format2layout(format)
	return c.ParseByLayout(value, layout)
}

// ParseByFormat 将特殊格式时间字符串解析成 Carbon 实例(默认时区)
func ParseByFormat(value string, format string) Carbon {
	return SetTimezone(Local).ParseByFormat(value, format)
}

// ParseByLayout 将布局时间字符串解析成 Carbon 实例
func (c Carbon) ParseByLayout(value string, layout string) Carbon {
	if c.Loc == nil {
		c.Loc = time.Local
	}
	tt, err := time.ParseInLocation(layout, value, c.Loc)
	c.Time = tt
	if err != nil {
		c.Error = errors.New("the value \"" + value + "\" can't parse string as time")
	}
	return c
}

// ParseByLayout 将布局时间字符串解析成 Carbon 实例(默认时区)
func ParseByLayout(value string, layout string) Carbon {
	return SetTimezone(Local).ParseByLayout(value, layout)
}

// ParseByInput 将布局时间字符串解析成 Carbon 实例
func (c Carbon) ParseByInput(value string, input string) Carbon {
	if c.Loc == nil {
		c.Loc = time.Local
	}
	year, month, day, hour, minute, second := 0, 0, 0, 0, 0, 0
	value = strings.Replace(value, " ", "", -1)
	input = strings.Replace(input, " ", "", -1)
	input = strings.Replace(input, "Y", "YYYY", 1)
	input = strings.Replace(input, "m", "mm", 1)
	input = strings.Replace(input, "d", "dd", 1)
	input = strings.Replace(input, "H", "HH", 1)
	input = strings.Replace(input, "i", "ii", 1)
	input = strings.Replace(input, "s", "ss", 1)
	for i := range value {
		if value[i] != input[i] {
			// fmt.Println(int(value[i]))
			if int(input[i]) == 89 {
				year = year*10 + (int(value[i]) - 48)
			} else if int(input[i]) == 109 {
				month = month*10 + (int(value[i]) - 48)
			} else if int(input[i]) == 100 {
				day = day*10 + (int(value[i]) - 48)
			} else if int(input[i]) == 72 {
				hour = hour*10 + (int(value[i]) - 48)
			} else if int(input[i]) == 105 {
				minute = minute*10 + (int(value[i]) - 48)
			} else if int(input[i]) == 115 {
				second = second*10 + (int(value[i]) - 48)
			}
		}
	}
	value2 := strconv.Itoa(year) + "|" + strconv.Itoa(month) + "|" + strconv.Itoa(day) + " " + strconv.Itoa(hour) + "|" + strconv.Itoa(minute) + "|" + strconv.Itoa(second)
	tt, err := time.ParseInLocation("2006|01|02 15|04|05", value2, c.Loc)
	c.Time = tt
	if err != nil {
		c.Error = errors.New("the value \"" + value + "\" can't parse string as time")
	}
	return c
}

// ParseByInput 将输入时间字符串解析成 Carbon 实例(默认时区)
func ParseByInput(value string, input string) Carbon {
	return SetTimezone(Local).ParseByInput(value, input)
}