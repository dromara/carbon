package carbon_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dromara/carbon/v2"
)

func TestTimeMilli_Scan(t *testing.T) {
	c := carbon.NewTimeMilli(carbon.Now())

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

func TestTimeMilli_Value(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := carbon.NewCarbon()
		v, err := carbon.NewTimeMilli(c).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := carbon.Parse("xxx")
		v, err := carbon.NewTimeMilli(c).Value()
		assert.Nil(t, v)
		assert.Error(t, err)
	})

	t.Run("valid time", func(t *testing.T) {
		c := carbon.Parse("2020-08-05")
		v, err := carbon.NewTimeMilli(c).Value()
		assert.Equal(t, c.StdTime(), v)
		assert.Nil(t, err)
	})
}

func TestTimeMilli_MarshalJSON(t *testing.T) {
	type User struct {
		TimeMilli carbon.TimeMilli `json:"time_milli"`
	}

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := carbon.NewCarbon()
		user.TimeMilli = carbon.NewTimeMilli(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"time_milli":""}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("xxx")
		user.TimeMilli = carbon.NewTimeMilli(c)

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("2020-08-05 13:14:15.999999999")
		user.TimeMilli = carbon.NewTimeMilli(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"time_milli":"13:14:15.999"}`, string(data))
	})
}

func TestTimeMilli_UnmarshalJSON(t *testing.T) {
	type User struct {
		TimeMilli carbon.TimeMilli `json:"time_milli"`
	}

	t.Run("empty value", func(t *testing.T) {
		var user User

		value := `{"time_milli":""}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.TimeMilli.String())
	})

	t.Run("null value", func(t *testing.T) {
		var user User

		value := `{"time_milli":"null"}`

		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.TimeMilli.String())
	})

	t.Run("zero value", func(t *testing.T) {
		var user User

		value := `{"time_milli":"0"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Empty(t, user.TimeMilli.String())
	})

	t.Run("valid value", func(t *testing.T) {
		var user User

		value := `{"time_milli":"13:14:15.999"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Equal(t, "13:14:15.999", user.TimeMilli.String())
	})
}

func TestTimeMilli_String(t *testing.T) {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	assert.Equal(t, "13:14:15.999", carbon.NewTimeMilli(c).String())
}

func TestTimeMilli_GormDataType(t *testing.T) {
	assert.Equal(t, "time", carbon.NewTimeMilli(carbon.Now()).GormDataType())
}
