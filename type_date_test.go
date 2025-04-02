package carbon_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dromara/carbon/v2"
)

func TestDate_Scan(t *testing.T) {
	c := carbon.NewDate(carbon.Now())

	t.Run("[]byte type", func(t *testing.T) {
		assert.Nil(t, c.Scan([]byte(carbon.Now().ToDateString())))
	})

	t.Run("string type", func(t *testing.T) {
		assert.Nil(t, c.Scan(carbon.Now().ToDateString()))
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

func TestDate_Value(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := carbon.NewCarbon()
		v, err := carbon.NewDate(c).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := carbon.Parse("xxx")
		v, err := carbon.NewDate(c).Value()
		assert.Nil(t, v)
		assert.Error(t, err)
	})

	t.Run("valid time", func(t *testing.T) {
		c := carbon.Parse("2020-08-05")
		v, err := carbon.NewDate(c).Value()
		assert.Equal(t, c.StdTime(), v)
		assert.Nil(t, err)
	})
}

func TestDate_MarshalJSON(t *testing.T) {
	type User struct {
		Date carbon.Date `json:"date"`
	}

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := carbon.NewCarbon()
		user.Date = carbon.NewDate(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":""}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("xxx")
		user.Date = carbon.NewDate(c)

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("2020-08-05 13:14:15.999999999")
		user.Date = carbon.NewDate(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"2020-08-05"}`, string(data))
	})
}

func TestDate_UnmarshalJSON(t *testing.T) {
	type User struct {
		Date carbon.Date `json:"date"`
	}

	t.Run("empty value", func(t *testing.T) {
		var user User

		value := `{"date":""}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.Date.String())
	})

	t.Run("null value", func(t *testing.T) {
		var user User

		value := `{"date":"null"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.Date.String())
	})

	t.Run("zero value", func(t *testing.T) {
		var user User

		value := `{"date":"0"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.Date.String())
	})

	t.Run("valid value", func(t *testing.T) {
		var user User

		value := `{"date":"2020-08-05"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Equal(t, "2020-08-05", user.Date.String())
	})
}

func TestDate_String(t *testing.T) {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	assert.Equal(t, "2020-08-05", carbon.NewDate(c).String())
}

func TestDate_GormDataType(t *testing.T) {
	assert.Equal(t, "time", carbon.NewDate(carbon.Now()).GormDataType())
}
