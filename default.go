package carbon

var (
	// DefaultLayout default layout
	// 默认布局模板
	DefaultLayout = DateTimeLayout

	// DefaultTimezone default timezone
	// 默认时区
	DefaultTimezone = UTC

	// DefaultWeekStartsAt default week start date
	// 默认一周开始日期
	DefaultWeekStartsAt = Sunday

	// DefaultLocale default language locale
	// 默认语言区域
	DefaultLocale = "en"
)

// Default defines a Default struct.
// 定义 Default 结构体
type Default struct {
	Layout       string
	Timezone     string
	WeekStartsAt string
	Locale       string
}

// SetDefault sets default.
// 设置全局默认值
func SetDefault(d Default) {
	if d.Layout != "" {
		DefaultLayout = d.Layout
	}
	if d.Timezone != "" {
		DefaultTimezone = d.Timezone
	}
	if d.WeekStartsAt != "" {
		DefaultWeekStartsAt = d.WeekStartsAt
	}
	if d.Locale != "" {
		DefaultLocale = d.Locale
	}
}

// ResetDefault resets default.
// 重置全局默认值
func ResetDefault() {
	DefaultLayout = DateTimeLayout
	DefaultTimezone = UTC
	DefaultWeekStartsAt = Sunday
	DefaultLocale = "en"
}
