package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_SetDefault(t *testing.T) {
	prepareTest(t)

	SetDefault(Default{
		Layout:       DateTimeLayout,
		Timezone:     PRC,
		Locale:       "en",
		WeekStartsAt: Sunday,
	})
	assert.Equal(t, DateTimeLayout, defaultLayout)
	assert.Equal(t, PRC, defaultTimezone)
	assert.Equal(t, "en", defaultLocale)
	assert.Equal(t, "Sunday", defaultWeekStartsAt)
}
