package carbon_test

import (
	"encoding/json"
	"fmt"

	"github.com/dromara/carbon/v2"
)

func ExampleCarbon_MarshalJSON() {
	type User struct {
		Carbon1 carbon.Carbon  `json:"carbon1"`
		Carbon2 *carbon.Carbon `json:"carbon2"`
	}

	user := User{
		Carbon1: *carbon.Parse("2020-08-05 13:14:15.999999999"),
		Carbon2: carbon.Parse("2020-08-05 13:14:15.999999999"),
	}

	data, _ := json.Marshal(&user)
	fmt.Printf("%s", data)

	// Output:
	// {"carbon1":"2020-08-05 13:14:15","carbon2":"2020-08-05 13:14:15"}
}

func ExampleCarbon_UnmarshalJSON() {
	type User struct {
		Carbon1 carbon.Carbon  `json:"carbon1"`
		Carbon2 *carbon.Carbon `json:"carbon2"`
	}

	var user User

	value := `{"carbon1":"2020-08-05 13:14:15","carbon2":"2020-08-05 13:14:15"}`
	_ = json.Unmarshal([]byte(value), &user)

	fmt.Println("user.Carbon1:", user.Carbon1.String())
	fmt.Println("user.Carbon2:", user.Carbon2.String())

	// Output:
	// user.Carbon1: 2020-08-05 13:14:15
	// user.Carbon2: 2020-08-05 13:14:15
}

func ExampleCarbon_String() {
	fmt.Println(carbon.Parse("2020-08-05 13:14:15").String())

	// Output:
	// 2020-08-05 13:14:15
}

type W3CType string

func (W3CType) Layout() string {
	return carbon.W3cLayout
}

func ExampleLayoutType_MarshalJSON() {
	// type W3CType string
	//
	// func (W3CType) Layout() string {
	//	return carbon.W3cLayout
	// }

	type User struct {
		Date     carbon.Date                `json:"date"`
		Time     carbon.Time                `json:"time"`
		DateTime carbon.DateTime            `json:"date_time"`
		Customer carbon.LayoutType[W3CType] `json:"customer"`

		CreatedAt *carbon.DateTime `json:"created_at"`
		UpdatedAt *carbon.DateTime `json:"updated_at"`
	}

	c := carbon.Parse("2020-08-05 13:14:15.999999999")

	user := User{
		Date:     *carbon.NewDate(c),
		Time:     *carbon.NewTime(c),
		DateTime: *carbon.NewDateTime(c),
		Customer: *carbon.NewLayoutType[W3CType](c),

		CreatedAt: carbon.NewDateTime(c),
		UpdatedAt: carbon.NewDateTime(c),
	}

	data, _ := json.Marshal(&user)
	fmt.Printf("%s", data)

	// Output:
	// {"date":"2020-08-05","time":"13:14:15","date_time":"2020-08-05 13:14:15","customer":"2020-08-05T13:14:15Z","created_at":"2020-08-05 13:14:15","updated_at":"2020-08-05 13:14:15"}
}

func ExampleLayoutType_UnmarshalJSON() {
	// type W3CType string
	//
	// func (W3CType) Layout() string {
	//	return carbon.W3cLayout
	// }
	type User struct {
		Date     carbon.Date                `json:"date"`
		Time     carbon.Time                `json:"time"`
		DateTime carbon.DateTime            `json:"date_time"`
		Customer carbon.LayoutType[W3CType] `json:"customer"`

		CreatedAt *carbon.DateTime `json:"created_at"`
		UpdatedAt *carbon.DateTime `json:"updated_at"`
	}

	var user User

	value := `{"date":"2020-08-05","time":"13:14:15","date_time":"2020-08-05 13:14:15","customer":"2020-08-05T13:14:15Z","created_at":"2020-08-05 13:14:15","updated_at":"2020-08-05 13:14:15"}`
	_ = json.Unmarshal([]byte(value), &user)

	fmt.Println("user.Date:", user.Date.String())
	fmt.Println("user.Time:", user.Time.String())
	fmt.Println("user.DateTime:", user.DateTime.String())
	fmt.Println("user.Customer:", user.Customer.String())
	fmt.Println("user.CreatedAt:", user.CreatedAt.String())
	fmt.Println("user.UpdatedAt:", user.UpdatedAt.String())

	// Output:
	// user.Date: 2020-08-05
	// user.Time: 13:14:15
	// user.DateTime: 2020-08-05 13:14:15
	// user.Customer: 2020-08-05T13:14:15Z
	// user.CreatedAt: 2020-08-05 13:14:15
	// user.UpdatedAt: 2020-08-05 13:14:15
}

func ExampleLayoutType_String() {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")

	fmt.Println("carbon.Date:", carbon.NewDate(c).String())
	fmt.Println("carbon.Time:", carbon.NewTime(c).String())
	fmt.Println("carbon.DateTime:", carbon.NewDateTime(c).String())
	fmt.Println("Customer:", carbon.NewLayoutType[W3CType](c).String())

	// Output:
	// carbon.Date: 2020-08-05
	// carbon.Time: 13:14:15
	// carbon.DateTime: 2020-08-05 13:14:15
	// Customer: 2020-08-05T13:14:15Z
}

type RFC3339Type string

func (RFC3339Type) Format() string {
	return carbon.RFC3339Format
}

func ExampleFormatType_MarshalJSON() {
	// type RFC3339Type string
	//
	// func (RFC3339Type) Format() string {
	//	return carbon.RFC3339Format
	// }
	type User struct {
		Customer1 carbon.FormatType[RFC3339Type]  `json:"customer1"`
		Customer2 *carbon.FormatType[RFC3339Type] `json:"customer2"`
	}

	c := carbon.Parse("2020-08-05 13:14:15.999999999")

	user := User{
		Customer1: *carbon.NewFormatType[RFC3339Type](c),
		Customer2: carbon.NewFormatType[RFC3339Type](c),
	}

	data, _ := json.Marshal(&user)
	fmt.Printf("%s", data)

	// Output:
	// {"customer1":"2020-08-05T13:14:15Z","customer2":"2020-08-05T13:14:15Z"}
}

func ExampleFormatType_UnmarshalJSON() {
	// type RFC3339Type string
	//
	// func (RFC3339Type) Format() string {
	//	return carbon.RFC3339Format
	// }

	type User struct {
		Customer1 carbon.FormatType[RFC3339Type]  `json:"customer1"`
		Customer2 *carbon.FormatType[RFC3339Type] `json:"customer2"`
	}

	var user User

	value := `{"customer1":"2020-08-05T13:14:15Z","customer2":"2020-08-05T13:14:15Z"}`
	_ = json.Unmarshal([]byte(value), &user)

	fmt.Println("user.Customer1:", user.Customer1.String())
	fmt.Println("user.Customer2:", user.Customer2.String())

	// Output:
	// user.Customer1: 2020-08-05T13:14:15Z
	// user.Customer2: 2020-08-05T13:14:15Z
}

func ExampleFormatType_String() {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	fmt.Println(carbon.NewFormatType[RFC3339Type](c).String())

	// Output:
	// 2020-08-05T13:14:15Z
}

func ExampleTimestamp_MarshalJSON() {
	type User struct {
		Timestamp1 carbon.Timestamp  `json:"timestamp1"`
		Timestamp2 *carbon.Timestamp `json:"timestamp2"`
	}

	c := carbon.Parse("2020-08-05 13:14:15.999999999")

	user := User{
		Timestamp1: *carbon.NewTimestamp(c),
		Timestamp2: carbon.NewTimestamp(c),
	}

	data, _ := json.Marshal(&user)
	fmt.Printf("%s", data)

	// Output:
	// {"timestamp1":1596633255,"timestamp2":1596633255}
}

func ExampleTimestamp_UnmarshalJSON() {
	type User struct {
		Timestamp1 carbon.Timestamp  `json:"timestamp1"`
		Timestamp2 *carbon.Timestamp `json:"timestamp2"`
	}

	var user User

	value := `{"timestamp1":1596633255,"timestamp2":1596633255}`
	_ = json.Unmarshal([]byte(value), &user)

	fmt.Printf("user.Timestamp1: string(%s)\n", user.Timestamp1.String())
	fmt.Printf("user.Timestamp1: int64(%d)\n", user.Timestamp1.Int64())

	fmt.Printf("user.Timestamp2: string(%s)\n", user.Timestamp2.String())
	fmt.Printf("user.Timestamp2: int64(%d)\n", user.Timestamp2.Int64())

	// Output:
	// user.Timestamp1: string(1596633255)
	// user.Timestamp1: int64(1596633255)
	// user.Timestamp2: string(1596633255)
	// user.Timestamp2: int64(1596633255)
}

func ExampleTimestamp_String() {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	fmt.Println(carbon.NewTimestamp(c).String())

	// Output:
	// 1596633255
}

func ExampleTimestamp_Int64() {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	fmt.Println(carbon.NewTimestamp(c).Int64())

	// Output:
	// 1596633255
}

func ExampleTimestampMilli_MarshalJSON() {
	type User struct {
		TimestampMilli1 carbon.TimestampMilli  `json:"timestamp_milli1"`
		TimestampMilli2 *carbon.TimestampMilli `json:"timestamp_milli2"`
	}

	c := carbon.Parse("2020-08-05 13:14:15.999999999")

	user := User{
		TimestampMilli1: *carbon.NewTimestampMilli(c),
		TimestampMilli2: carbon.NewTimestampMilli(c),
	}

	data, _ := json.Marshal(&user)
	fmt.Printf("%s", data)

	// Output:
	// {"timestamp_milli1":1596633255999,"timestamp_milli2":1596633255999}
}

func ExampleTimestampMilli_UnmarshalJSON() {
	type User struct {
		TimestampMilli1 carbon.TimestampMilli  `json:"timestamp_milli1"`
		TimestampMilli2 *carbon.TimestampMilli `json:"timestamp_milli2"`
	}

	var user User

	value := `{"timestamp_milli1":1596633255999,"timestamp_milli2":1596633255999}`
	_ = json.Unmarshal([]byte(value), &user)

	fmt.Printf("user.TimestampMilli1: string(%s)\n", user.TimestampMilli1.String())
	fmt.Printf("user.TimestampMilli1: int64(%d)\n", user.TimestampMilli1.Int64())

	fmt.Printf("user.TimestampMilli2: string(%s)\n", user.TimestampMilli2.String())
	fmt.Printf("user.TimestampMilli2: int64(%d)\n", user.TimestampMilli2.Int64())

	// Output:
	// user.TimestampMilli1: string(1596633255999)
	// user.TimestampMilli1: int64(1596633255999)
	// user.TimestampMilli2: string(1596633255999)
	// user.TimestampMilli2: int64(1596633255999)
}

func ExampleTimestampMilli_String() {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	fmt.Println(carbon.NewTimestampMilli(c).String())

	// Output:
	// 1596633255999
}

func ExampleTimestampMilli_Int64() {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	fmt.Println(carbon.NewTimestampMilli(c).Int64())

	// Output:
	// 1596633255999
}

func ExampleTimestampMicro_MarshalJSON() {
	type User struct {
		TimestampMicro1 carbon.TimestampMicro  `json:"timestamp_micro1"`
		TimestampMicro2 *carbon.TimestampMicro `json:"timestamp_micro2"`
	}

	c := carbon.Parse("2020-08-05 13:14:15.999999999")

	user := User{
		TimestampMicro1: *carbon.NewTimestampMicro(c),
		TimestampMicro2: carbon.NewTimestampMicro(c),
	}

	data, _ := json.Marshal(&user)
	fmt.Printf("%s", data)

	// Output:
	// {"timestamp_micro1":1596633255999999,"timestamp_micro2":1596633255999999}
}

func ExampleTimestampMicro_UnmarshalJSON() {
	type User struct {
		TimestampMicro1 carbon.TimestampMicro  `json:"timestamp_micro1"`
		TimestampMicro2 *carbon.TimestampMicro `json:"timestamp_micro2"`
	}

	var user User

	value := `{"timestamp_micro1":1596633255999999,"timestamp_micro2":1596633255999999}`
	_ = json.Unmarshal([]byte(value), &user)

	fmt.Printf("user.TimestampMicro1: string(%s)\n", user.TimestampMicro1.String())
	fmt.Printf("user.TimestampMicro1: int64(%d)\n", user.TimestampMicro1.Int64())

	fmt.Printf("user.TimestampMicro2: string(%s)\n", user.TimestampMicro2.String())
	fmt.Printf("user.TimestampMicro2: int64(%d)\n", user.TimestampMicro2.Int64())

	// Output:
	// user.TimestampMicro1: string(1596633255999999)
	// user.TimestampMicro1: int64(1596633255999999)
	// user.TimestampMicro2: string(1596633255999999)
	// user.TimestampMicro2: int64(1596633255999999)
}

func ExampleTimestampMicro_String() {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	fmt.Println(carbon.NewTimestampMicro(c).String())

	// Output:
	// 1596633255999999
}

func ExampleTimestampMicro_Int64() {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	fmt.Println(carbon.NewTimestampMicro(c).Int64())

	// Output:
	// 1596633255999999
}

func ExampleTimestampNano_MarshalJSON() {
	type User struct {
		TimestampNano1 carbon.TimestampNano  `json:"timestamp_nano1"`
		TimestampNano2 *carbon.TimestampNano `json:"timestamp_nano2"`
	}

	c := carbon.Parse("2020-08-05 13:14:15.999999999")

	user := User{
		TimestampNano1: *carbon.NewTimestampNano(c),
		TimestampNano2: carbon.NewTimestampNano(c),
	}

	data, _ := json.Marshal(&user)
	fmt.Printf("%s", data)

	// Output:
	// {"timestamp_nano1":1596633255999999999,"timestamp_nano2":1596633255999999999}
}

func ExampleTimestampNano_UnmarshalJSON() {
	type User struct {
		TimestampNano1 carbon.TimestampNano  `json:"timestamp_nano1"`
		TimestampNano2 *carbon.TimestampNano `json:"timestamp_nano2"`
	}

	var user User

	value := `{"timestamp_nano1":1596633255999999999,"timestamp_nano2":1596633255999999999}`
	_ = json.Unmarshal([]byte(value), &user)

	fmt.Printf("user.TimestampNano1: string(%s)\n", user.TimestampNano1.String())
	fmt.Printf("user.TimestampNano1: int64(%d)\n", user.TimestampNano1.Int64())

	fmt.Printf("user.TimestampNano2: string(%s)\n", user.TimestampNano2.String())
	fmt.Printf("user.TimestampNano2: int64(%d)\n", user.TimestampNano2.Int64())

	// Output:
	// user.TimestampNano1: string(1596633255999999999)
	// user.TimestampNano1: int64(1596633255999999999)
	// user.TimestampNano2: string(1596633255999999999)
	// user.TimestampNano2: int64(1596633255999999999)
}

func ExampleTimestampNano_String() {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	fmt.Println(carbon.NewTimestampNano(c).String())

	// Output:
	// 1596633255999999999
}

func ExampleTimestampNano_Int64() {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	fmt.Println(carbon.NewTimestampNano(c).Int64())

	// Output:
	// 1596633255999999999
}
