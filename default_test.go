package carbon_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dromara/carbon/v2"
)

func TestSetDefault(t *testing.T) {
	defer carbon.ResetDefault()

	carbon.SetDefault(carbon.Default{
		Layout:       carbon.DateTimeLayout,
		Timezone:     carbon.PRC,
		Locale:       "zh-CN",
		WeekStartsAt: carbon.Monday,
		WeekendDays: []carbon.Weekday{
			carbon.Saturday, carbon.Sunday,
		},
	})

	assert.Equal(t, carbon.DateTimeLayout, carbon.DefaultLayout)
	assert.Equal(t, carbon.PRC, carbon.DefaultTimezone)
	assert.Equal(t, "zh-CN", carbon.DefaultLocale)
	assert.Equal(t, carbon.Monday, carbon.DefaultWeekStartsAt)
	assert.Equal(t, []carbon.Weekday{
		carbon.Saturday, carbon.Sunday,
	}, carbon.DefaultWeekendDays)
}
