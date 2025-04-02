package carbon_test

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dromara/carbon/v2"
)

func TestTimestamp_Scan(t *testing.T) {
	t.Run("[]byte type", func(t *testing.T) {
		ts := carbon.NewTimestamp(carbon.Now())
		assert.Nil(t, ts.Scan([]byte(strconv.Itoa(int(ts.Timestamp())))))
		assert.Error(t, ts.Scan([]byte("xxx")))
	})

	t.Run("string type", func(t *testing.T) {
		ts := carbon.NewTimestamp(carbon.Now())
		assert.Nil(t, ts.Scan(strconv.Itoa(int(ts.Timestamp()))))
		assert.Error(t, ts.Scan("xxx"))
	})

	t.Run("int64 type", func(t *testing.T) {
		ts := carbon.NewTimestamp(carbon.Now())
		assert.Nil(t, ts.Scan(carbon.Now().Timestamp()))
	})

	t.Run("time type", func(t *testing.T) {
		ts := carbon.NewTimestamp(carbon.Now())
		assert.Nil(t, ts.Scan(carbon.Now().StdTime()))
	})

	t.Run("unsupported type", func(t *testing.T) {
		ts := carbon.NewTimestamp(carbon.Now())
		assert.Error(t, ts.Scan(nil))
	})
}

func TestTimestamp_Value(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := carbon.NewCarbon()

		v, e := carbon.NewTimestamp(c).Value()
		assert.Nil(t, v)
		assert.Nil(t, e)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := carbon.Parse("xxx")

		v, e := carbon.NewTimestamp(c).Value()
		assert.Nil(t, v)
		assert.Error(t, e)
	})

	t.Run("valid time", func(t *testing.T) {
		c := carbon.Parse("2020-08-05")

		v, e := carbon.NewTimestamp(c).Value()
		assert.Equal(t, c.Timestamp(), v)
		assert.Nil(t, e)
	})
}

func TestTimestamp_MarshalJSON(t *testing.T) {
	type User struct {
		Timestamp carbon.Timestamp `json:"timestamp"`
	}

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := carbon.NewCarbon()
		user.Timestamp = carbon.NewTimestamp(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"timestamp":0}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("xxx")
		user.Timestamp = carbon.NewTimestamp(c)

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("2020-08-05 13:14:15.999999999")
		user.Timestamp = carbon.NewTimestamp(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"timestamp":1596633255}`, string(data))
	})
}

func TestTimestamp_UnmarshalJSON(t *testing.T) {
	type User struct {
		Timestamp carbon.Timestamp `json:"timestamp"`
	}

	t.Run("empty value", func(t *testing.T) {
		var user User

		value := `{"timestamp":""}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Equal(t, "0", user.Timestamp.String())
		assert.Zero(t, user.Timestamp.Int64())
	})

	t.Run("null value", func(t *testing.T) {
		var user User

		value := `{"timestamp":"null"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Equal(t, "0", user.Timestamp.String())
		assert.Zero(t, user.Timestamp.Int64())
	})

	t.Run("zero value", func(t *testing.T) {
		var user User

		value := `{"timestamp":"0"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Equal(t, "0", user.Timestamp.String())
		assert.Zero(t, user.Timestamp.Int64())
	})

	t.Run("invalid value", func(t *testing.T) {
		var user User

		value := `{"timestamp":"xxx"}`
		assert.Error(t, json.Unmarshal([]byte(value), &user))

		assert.Equal(t, "0", user.Timestamp.String())

		assert.Zero(t, user.Timestamp.Int64())
	})

	t.Run("valid value", func(t *testing.T) {
		var user User

		value := `{"timestamp":1596633255}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Equal(t, "1596633255", user.Timestamp.String())
		assert.Equal(t, int64(1596633255), user.Timestamp.Int64())
	})
}

func TestTimestamp_String(t *testing.T) {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	assert.Equal(t, "1596633255", carbon.NewTimestamp(c).String())
}

func TestTimestamp_Int64(t *testing.T) {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	assert.Equal(t, int64(1596633255), carbon.NewTimestamp(c).Int64())
}

func TestTimestamp_GormDataType(t *testing.T) {
	assert.Equal(t, "time", carbon.NewTimestamp(carbon.Now()).GormDataType())
}
