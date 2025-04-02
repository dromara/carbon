package carbon_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dromara/carbon/v2"
)

func TestTimeMicro_Scan(t *testing.T) {
	c := carbon.NewTimeMicro(carbon.Now())

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

func TestTimeMicro_Value(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := carbon.NewCarbon()
		v, err := carbon.NewTimeMicro(c).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := carbon.Parse("xxx")
		v, err := carbon.NewTimeMicro(c).Value()
		assert.Nil(t, v)
		assert.Error(t, err)
	})

	t.Run("valid time", func(t *testing.T) {
		c := carbon.Parse("2020-08-05")
		v, err := carbon.NewTimeMicro(c).Value()
		assert.Equal(t, c.StdTime(), v)
		assert.Nil(t, err)
	})
}

func TestTimeMicro_MarshalJSON(t *testing.T) {
	type User struct {
		TimeMicro carbon.TimeMicro `json:"time_micro"`
	}

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := carbon.NewCarbon()
		user.TimeMicro = carbon.NewTimeMicro(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"time_micro":""}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("xxx")
		user.TimeMicro = carbon.NewTimeMicro(c)

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("2020-08-05 13:14:15.999999999999")
		user.TimeMicro = carbon.NewTimeMicro(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"time_micro":"13:14:15.999999"}`, string(data))
	})
}

func TestTimeMicro_UnmarshalJSON(t *testing.T) {
	type User struct {
		TimeMicro carbon.TimeMicro `json:"time_micro"`
	}

	t.Run("empty value", func(t *testing.T) {
		var user User

		value := `{"time_micro":""}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.TimeMicro.String())
	})

	t.Run("null value", func(t *testing.T) {
		var user User

		value := `{"time_micro":"null"}`

		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.TimeMicro.String())
	})

	t.Run("zero value", func(t *testing.T) {
		var user User

		value := `{"time_micro":"0"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Empty(t, user.TimeMicro.String())
	})

	t.Run("valid value", func(t *testing.T) {
		var user User

		value := `{"time_micro":"13:14:15.999999"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Equal(t, "13:14:15.999999", user.TimeMicro.String())
	})
}

func TestTimeMicro_String(t *testing.T) {
	c := carbon.Parse("2020-08-05 13:14:15.999999999999")
	assert.Equal(t, "13:14:15.999999", carbon.NewTimeMicro(c).String())
}

func TestTimeMicro_GormDataType(t *testing.T) {
	assert.Equal(t, "time", carbon.NewTimeMicro(carbon.Now()).GormDataType())
}
