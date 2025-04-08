package carbon

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarbon_Scan(t *testing.T) {
	c := Now()

	t.Run("[]byte type", func(t *testing.T) {
		assert.Nil(t, c.Scan([]byte(Now().ToDateMicroString())))
	})

	t.Run("string type", func(t *testing.T) {
		assert.Nil(t, c.Scan(Now().ToDateMicroString()))
	})

	t.Run("int64 type", func(t *testing.T) {
		assert.Nil(t, c.Scan(Now().Timestamp()))
	})

	t.Run("time type", func(t *testing.T) {
		assert.Nil(t, c.Scan(Now().StdTime()))
	})

	t.Run("unsupported type", func(t *testing.T) {
		assert.Error(t, c.Scan(nil))
	})
}

func TestCarbon_Value(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		v, err := NewCarbon().Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("invalid time", func(t *testing.T) {
		v, err := Parse("xxx").Value()
		assert.Nil(t, v)
		assert.Error(t, err)
	})

	t.Run("valid time", func(t *testing.T) {
		c := Parse("2020-08-05")
		v, err := c.Value()
		assert.Equal(t, c.StdTime(), v)
		assert.Nil(t, err)
	})
}

func TestCarbon_MarshalJSON(t *testing.T) {
	type User struct {
		Carbon Carbon `json:"carbon"`
	}

	t.Run("zero time", func(t *testing.T) {
		var user User

		user.Carbon = NewCarbon()

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"carbon":""}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		user.Carbon = Parse("xxx")

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		user.Carbon = Parse("2020-08-05 13:14:15.999999999")

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"carbon":"2020-08-05 13:14:15"}`, string(data))
	})
}

func TestCarbon_UnmarshalJSON(t *testing.T) {
	type User struct {
		Carbon Carbon `json:"carbon"`
	}

	t.Run("empty value", func(t *testing.T) {
		var user User

		value := `{"carbon":""}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.Carbon.String())
	})

	t.Run("null value", func(t *testing.T) {
		var user User

		value := `{"carbon":"null"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.Carbon.String())
	})

	t.Run("zero value", func(t *testing.T) {
		var user User

		value := `{"carbon":"0"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Empty(t, user.Carbon.String())
	})

	t.Run("valid value", func(t *testing.T) {
		var user User

		value := `{"carbon":"2020-08-05 13:14:15"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Equal(t, "2020-08-05 13:14:15", user.Carbon.String())
	})
}

func TestCarbon_String(t *testing.T) {
	c := Parse("2020-08-05 13:14:15.999999999")
	assert.Equal(t, "2020-08-05 13:14:15", c.String())
}

func TestCarbon_GormDataType(t *testing.T) {
	assert.Equal(t, "time", Now().GormDataType())
}
