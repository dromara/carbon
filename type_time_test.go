package carbon_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dromara/carbon/v2"
)

func TestTime_Scan(t *testing.T) {
	c := carbon.NewTime(carbon.Now())

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
		e := c.Scan(nil)
		assert.Error(t, e)
	})
}

func TestTime_Value(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := carbon.NewCarbon()
		v, err := carbon.NewTime(c).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := carbon.Parse("xxx")
		v, err := carbon.NewTime(c).Value()
		assert.Nil(t, v)
		assert.Error(t, err)
	})

	t.Run("valid time", func(t *testing.T) {
		c := carbon.Parse("2020-08-05")
		v, err := carbon.NewTime(c).Value()
		assert.Equal(t, c.StdTime(), v)
		assert.Nil(t, err)
	})
}

func TestTime_MarshalJSON(t *testing.T) {
	type User struct {
		Time carbon.Time `json:"time"`
	}

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := carbon.NewCarbon()
		user.Time = carbon.NewTime(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"time":""}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("xxx")
		user.Time = carbon.NewTime(c)

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("2020-08-05 13:14:15.999999999")
		user.Time = carbon.NewTime(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"time":"13:14:15"}`, string(data))
	})
}

func TestTime_UnmarshalJSON(t *testing.T) {
	type User struct {
		Time carbon.Time `json:"time"`
	}

	t.Run("empty value", func(t *testing.T) {
		var user User

		value := `{"time":""}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.Time.String())
	})

	t.Run("null value", func(t *testing.T) {
		var user User

		value := `{"time":"null"}`

		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.Time.String())
	})

	t.Run("zero value", func(t *testing.T) {
		var user User

		value := `{"time":"0"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Empty(t, user.Time.String())
	})

	t.Run("valid value", func(t *testing.T) {
		var user User

		value := `{"time":"13:14:15"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Equal(t, "13:14:15", user.Time.String())
	})
}

func TestTime_String(t *testing.T) {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	assert.Equal(t, "13:14:15", carbon.NewTime(c).String())
}

func TestTime_GormDataType(t *testing.T) {
	assert.Equal(t, "time", carbon.NewTime(carbon.Now()).GormDataType())
}
