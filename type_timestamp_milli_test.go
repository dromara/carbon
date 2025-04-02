package carbon_test

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dromara/carbon/v2"
)

func TestTimestampMilli_Scan(t *testing.T) {
	t.Run("[]byte type", func(t *testing.T) {
		ts := carbon.NewTimestampMilli(carbon.Now())
		assert.Nil(t, ts.Scan([]byte(strconv.Itoa(int(ts.Timestamp())))))
		assert.Error(t, ts.Scan([]byte("xxx")))
	})

	t.Run("string type", func(t *testing.T) {
		ts := carbon.NewTimestampMilli(carbon.Now())
		assert.Nil(t, ts.Scan(strconv.Itoa(int(ts.Timestamp()))))
		assert.Error(t, ts.Scan("xxx"))
	})

	t.Run("int64 type", func(t *testing.T) {
		ts := carbon.NewTimestampMilli(carbon.Now())
		assert.Nil(t, ts.Scan(carbon.Now().Timestamp()))
	})

	t.Run("time type", func(t *testing.T) {
		ts := carbon.NewTimestampMilli(carbon.Now())
		assert.Nil(t, ts.Scan(carbon.Now().StdTime()))
	})

	t.Run("unsupported type", func(t *testing.T) {
		ts := carbon.NewTimestampMilli(carbon.Now())
		assert.Error(t, ts.Scan(nil))
	})
}

func TestTimestampMilli_Value(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		c := carbon.NewCarbon()

		v, e := carbon.NewTimestampMilli(c).Value()
		assert.Nil(t, v)
		assert.Nil(t, e)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := carbon.Parse("xxx")

		v, e := carbon.NewTimestampMilli(c).Value()
		assert.Nil(t, v)
		assert.Error(t, e)
	})

	t.Run("valid time", func(t *testing.T) {
		c := carbon.Parse("2020-08-05")

		v, e := carbon.NewTimestampMilli(c).Value()
		assert.Equal(t, c.TimestampMilli(), v)
		assert.Nil(t, e)
	})
}

func TestTimestampMilli_MarshalJSON(t *testing.T) {
	type User struct {
		TimestampMilli carbon.TimestampMilli `json:"timestamp_milli"`
	}

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := carbon.NewCarbon()
		user.TimestampMilli = carbon.NewTimestampMilli(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"timestamp_milli":0}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("xxx")
		user.TimestampMilli = carbon.NewTimestampMilli(c)

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("2020-08-05 13:14:15.999999999")
		user.TimestampMilli = carbon.NewTimestampMilli(c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"timestamp_milli":1596633255999}`, string(data))
	})
}

func TestTimestampMilli_UnmarshalJSON(t *testing.T) {
	type User struct {
		TimestampMilli carbon.TimestampMilli `json:"timestamp_milli"`
	}

	t.Run("empty value", func(t *testing.T) {
		var user User

		value := `{"timestamp_milli":""}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))
		assert.Equal(t, "0", user.TimestampMilli.String())
		assert.Zero(t, user.TimestampMilli.Int64())
	})

	t.Run("null value", func(t *testing.T) {
		var user User

		value := `{"timestamp_milli":"null"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Equal(t, "0", user.TimestampMilli.String())
		assert.Zero(t, user.TimestampMilli.Int64())
	})

	t.Run("zero value", func(t *testing.T) {
		var user User

		value := `{"timestamp_milli":"0"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Equal(t, "0", user.TimestampMilli.String())
		assert.Zero(t, user.TimestampMilli.Int64())
	})

	t.Run("invalid value", func(t *testing.T) {
		var user User

		value := `{"timestamp_milli":"xxx"}`
		assert.Error(t, json.Unmarshal([]byte(value), &user))

		assert.Equal(t, "0", user.TimestampMilli.String())

		assert.Zero(t, user.TimestampMilli.Int64())
	})

	t.Run("valid value", func(t *testing.T) {
		var user User

		value := `{"timestamp_milli":1596633255999}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Equal(t, "1596633255999", user.TimestampMilli.String())
		assert.Equal(t, int64(1596633255999), user.TimestampMilli.Int64())
	})
}

func TestTimestampMilli_String(t *testing.T) {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	assert.Equal(t, "1596633255999", carbon.NewTimestampMilli(c).String())
}

func TestTimestampMilli_Int64(t *testing.T) {
	c := carbon.Parse("2020-08-05 13:14:15.999999999")
	assert.Equal(t, int64(1596633255999), carbon.NewTimestampMilli(c).Int64())
}

func TestTimestampMilli_GormDataType(t *testing.T) {
	assert.Equal(t, "time", carbon.NewTimestampMilli(carbon.Now()).GormDataType())
}
