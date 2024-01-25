package carbon

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Lunar(t *testing.T) {
	l := Parse("2024-01-18").Lunar()
	assert.Nil(t, l.Error)
	assert.Equal(t, "2023-12-08 00:00:00", l.String())
}

func TestCreateFromLunar(t *testing.T) {
	l := CreateFromLunar(2023, 12, 11, 0, 0, 0, false)
	assert.Nil(t, l.Error)
	assert.Equal(t, "2024-01-21 00:00:00", l.String())
}

func TestError_Lunar(t *testing.T) {
	l := Parse("xxx").Lunar()
	assert.NotNil(t, l.Error)
}

func TestName(t *testing.T) {
	r := CreateFromLunar(2023, 12, 11, 0, 0, 0, false).ToDateTimeString()
	fmt.Println(r)
}
