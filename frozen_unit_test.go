package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetTestNow(t *testing.T) {
	now := Parse("2020-08-05")

	SetTestNow(now)
	assert.Equal(t, "2020-08-05", Now().ToDateString())
	assert.Equal(t, "2020-08-04", Yesterday().ToDateString())
	assert.Equal(t, "2020-08-06", Tomorrow().ToDateString())
	assert.Equal(t, "just now", Now().DiffForHumans())
	assert.Equal(t, "1 day ago", Yesterday().DiffForHumans())
	assert.Equal(t, "1 day from now", Tomorrow().DiffForHumans())
	assert.Equal(t, "2 months from now", Parse("2020-10-05").DiffForHumans())
	assert.Equal(t, "2 months before", now.DiffForHumans(Parse("2020-10-05")))
	assert.True(t, IsTestNow())
	assert.True(t, IsSetTestNow())

	CleanTestNow()
	UnSetTestNow()
	assert.Equal(t, time.Now().In(time.UTC).Format(DateLayout), Now(UTC).ToDateString())
	assert.Equal(t, time.Now().In(time.UTC).Add(time.Hour*-24).Format(DateLayout), Yesterday(UTC).ToDateString())
	assert.Equal(t, time.Now().In(time.UTC).Add(time.Hour*24).Format(DateLayout), Tomorrow(UTC).ToDateString())
	assert.False(t, IsTestNow())
	assert.False(t, IsSetTestNow())
}
