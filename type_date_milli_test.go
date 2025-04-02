package carbon_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dromara/carbon/v2"
)

func TestDateMilli_Scan(t *testing.T) {
	c := carbon.NewDateMilli(carbon.Now())

	t.Run("[]byte type", func(t *testing.T) {
		assert.Nil(t, c.Scan([]byte(carbon.Now().ToDateMilliString())))
	})

	t.Run("string type", func(t *testing.T) {
		assert.Nil(t, c.Scan(carbon.Now().ToDateMilliString()))
	})

	t.Run("int64 type", func(t *testing.T) {
		assert.Nil(t, c.Scan(carbon.Now().Timestamp()))
	})

	t.Run("time type", func(t *testing.T) {
		assert.Nil(t, c.Scan(carbon.Now().StdTime()))
	})

	t.Run("unsupported type", func(t *testing.T) {
		assert.Error(t, c.Scan(nil))
	})
}

func TestDateMilli_Value(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := carbon.NewCarbon()
		v, err := carbon.NewDateMilli(c).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := carbon.Parse("xxx")
		v, err := carbon.NewDateMilli(c).Value()
		assert.Nil(t, v)
		assert.Error(t, err)
	})

	t.Run("valid time", func(t *testing.T) {
		c := carbon.Parse("2020-08-05")
		v, err := carbon.NewDateMilli(c).Value()
		assert.Equal(t, c.StdTime(), v)
		assert.Nil(t, err)
	})
}

func TestDateMilli_MarshalJSON(t *testing.T) {
	type User struct {
		DateMilli carbon.DateMilli `json:"date_milli"`
	}

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := carbon.NewCarbon()
		user.DateMilli = carbon.NewDateMilli(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date_milli":""}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("xxx")
		user.DateMilli = carbon.NewDateMilli(c)

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("2020-08-05 13:14:15.999999999")
		user.DateMilli = carbon.NewDateMilli(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date_milli":"2020-08-05.999"}`, string(data))
	})
}

func TestDateMilli_UnmarshalJSON(t *testing.T) {
	type User struct {
		DateMilli carbon.DateMilli `json:"date_milli"`
	}

	t.Run("empty value", func(t *testing.T) {
		var user User

		value := `{"date_milli":""}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.DateMilli.String())
	})

	t.Run("null value", func(t *testing.T) {
		var user User

		value := `{"date_milli":"null"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.DateMilli.String())
	})

	t.Run("zero value", func(t *testing.T) {
		var user User

		value := `{"date_milli":"0"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.DateMilli.String())
	})

	t.Run("valid value", func(t *testing.T) {
		var user User

		value := `{"date_milli":"2020-08-05.999"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Equal(t, "2020-08-05.999", user.DateMilli.String())
	})
}

func TestDateMilli_String(t *testing.T) {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	assert.Equal(t, "2020-08-05.999", carbon.NewDateMilli(c).String())
}

func TestDateMilli_GormDataType(t *testing.T) {
	assert.Equal(t, "time", carbon.NewDateMilli(carbon.Now()).GormDataType())
}
