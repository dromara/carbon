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

	assert.Equal(t, DateTimeLayout, DefaultLayout)
	assert.Equal(t, UTC, DefaultTimezone)
	assert.Equal(t, "en", DefaultLocale)
	assert.Equal(t, "Sunday", DefaultWeekStartsAt)
}
