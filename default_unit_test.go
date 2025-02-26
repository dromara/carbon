package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_SetDefault(t *testing.T) {
	defer SetDefault(getDefault())

	SetDefault(Default{
		Layout:       DateTimeLayout,
		Timezone:     UTC,
		Locale:       "en",
		WeekStartsAt: Sunday,
	})

	assert.Equal(t, DateTimeLayout, defaultLayout)
	assert.Equal(t, UTC, defaultTimezone)
	assert.Equal(t, "en", defaultLocale)
	assert.Equal(t, "Sunday", defaultWeekStartsAt)
}
