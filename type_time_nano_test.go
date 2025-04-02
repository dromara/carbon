package carbon_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dromara/carbon/v2"
)

func TestTimeNano_Scan(t *testing.T) {
	c := carbon.NewTimeNano(carbon.Now())

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

func TestTimeNano_Value(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := carbon.NewCarbon()
		v, err := carbon.NewTimeNano(c).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := carbon.Parse("xxx")
		v, err := carbon.NewTimeNano(c).Value()
		assert.Nil(t, v)
		assert.Error(t, err)
	})

	t.Run("valid time", func(t *testing.T) {
		c := carbon.Parse("2020-08-05")
		v, err := carbon.NewTimeNano(c).Value()
		assert.Equal(t, c.StdTime(), v)
		assert.Nil(t, err)
	})
}

func TestTimeNano_MarshalJSON(t *testing.T) {
	type User struct {
		TimeNano carbon.TimeNano `json:"time_nano"`
	}

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := carbon.NewCarbon()
		user.TimeNano = carbon.NewTimeNano(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"time_nano":""}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("xxx")
		user.TimeNano = carbon.NewTimeNano(c)

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("2020-08-05 13:14:15.999999999999999")
		user.TimeNano = carbon.NewTimeNano(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"time_nano":"13:14:15.999999999"}`, string(data))
	})
}

func TestTimeNano_UnmarshalJSON(t *testing.T) {
	type User struct {
		TimeNano carbon.TimeNano `json:"time_nano"`
	}

	t.Run("empty value", func(t *testing.T) {
		var user User

		value := `{"time_nano":""}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.TimeNano.String())
	})

	t.Run("null value", func(t *testing.T) {
		var user User

		value := `{"time_nano":"null"}`

		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.TimeNano.String())
	})

	t.Run("zero value", func(t *testing.T) {
		var user User

		value := `{"time_nano":"0"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Empty(t, user.TimeNano.String())
	})

	t.Run("valid value", func(t *testing.T) {
		var user User

		value := `{"time_nano":"13:14:15.999999999"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Equal(t, "13:14:15.999999999", user.TimeNano.String())
	})
}

func TestTimeNano_String(t *testing.T) {
	c := carbon.Parse("2020-08-05 13:14:15.999999999999999")
	assert.Equal(t, "13:14:15.999999999", carbon.NewTimeNano(c).String())
}

func TestTimeNano_GormDataType(t *testing.T) {
	assert.Equal(t, "time", carbon.NewTimeNano(carbon.Now()).GormDataType())
}
