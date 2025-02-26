package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_SetTestNow(t *testing.T) {
	carbon := NewCarbon()
	carbon.SetTestNow(Parse("2020-08-05"))

	assert.Equal(t, "2020-08-04", carbon.Yesterday().ToDateString(), "It should be equal to 2020-08-04")
	assert.Equal(t, "2020-08-05", carbon.Now().ToDateString(), "It should be equal to 2020-08-05")
	assert.Equal(t, "2020-08-06", carbon.Tomorrow().ToDateString(), "It should be equal to 2020-08-06")
	assert.Equal(t, 31, carbon.Parse("1989-08-05").Age(), "It should be equal to 31")
	assert.Equal(t, int64(1), carbon.Parse("2020-07-05").DiffInMonths(), "It should be equal to 1")
	assert.Equal(t, int64(1), carbon.Parse("2020-07-05").DiffAbsInMonths(), "It should be equal to 1")
	assert.Equal(t, int64(4), carbon.Parse("2020-07-05").DiffInWeeks(), "It should be equal to 4")
	assert.Equal(t, int64(4), carbon.Parse("2020-07-05").DiffAbsInWeeks(), "It should be equal to 4")
	assert.Equal(t, int64(31), carbon.Parse("2020-07-05").DiffInDays(), "It should be equal to 31")
	assert.Equal(t, int64(31), carbon.Parse("2020-07-05").DiffAbsInDays(), "It should be equal to 31")
	assert.Equal(t, int64(-13), carbon.Parse("2020-08-05 13:14:15").DiffInHours(), "It should be equal to -13")
	assert.Equal(t, int64(13), carbon.Parse("2020-08-05 13:14:15").DiffAbsInHours(), "It should be equal to 13")
	assert.Equal(t, int64(-794), carbon.Parse("2020-08-05 13:14:15").DiffInMinutes(), "It should be equal to -794")
	assert.Equal(t, int64(794), carbon.Parse("2020-08-05 13:14:15").DiffAbsInMinutes(), "It should be equal to 794")
	assert.Equal(t, int64(-47655), carbon.Parse("2020-08-05 13:14:15").DiffInSeconds(), "It should be equal to -47655")
	assert.Equal(t, int64(47655), carbon.Parse("2020-08-05 13:14:15").DiffAbsInSeconds(), "It should be equal to 47655")
	assert.Equal(t, "-13 hours", carbon.Parse("2020-08-05 13:14:15").DiffInString(), "It should be -13 hours")
	assert.Equal(t, "13 hours", carbon.Parse("2020-08-05 13:14:15").DiffAbsInString(), "It should be 13 hours")
	assert.Equal(t, "-13h0m0s", carbon.Parse("2020-08-05 13:00:00").DiffInDuration().String(), "It should be -13h0m0s")
	assert.Equal(t, "13h0m0s", carbon.Parse("2020-08-05 13:00:00").DiffAbsInDuration().String(), "It should be 13h0m0s")
	assert.Equal(t, "13 hours from now", carbon.Parse("2020-08-05 13:14:15").DiffForHumans(), "It should be 13 hours from now")
}

func TestCarbon_UnSetTestNow(t *testing.T) {
	carbon := NewCarbon()
	carbon.SetTestNow(Parse("2020-08-05"))
	carbon.UnSetTestNow()

	assert.Equal(t, Now().ToDateString(), carbon.Now().ToDateString())
}

func TestCarbon_IsSetTestNow(t *testing.T) {
	carbon := NewCarbon()
	assert.Equal(t, false, Now().IsSetTestNow(), "It should be equal to false")

	carbon.SetTestNow(Parse("2020-08-05"))
	assert.Equal(t, true, carbon.IsSetTestNow(), "It should be equal to true")
}
