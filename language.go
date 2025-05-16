package carbon

import (
	"embed"
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"sync"
)

//go:embed lang
var fs embed.FS

var validResourcesKey = []string{
	"months", "short_months", "weeks", "short_weeks", "seasons", "constellations",
	"year", "month", "week", "day", "hour", "minute", "second",
	"now", "ago", "from_now", "before", "after",
}

// Language defines a Language struct.
type Language struct {
	dir       string
	locale    string
	resources map[string]string
	Error     error
	rw        *sync.RWMutex
}

// NewLanguage returns a new Language instance.
func NewLanguage() *Language {
	return &Language{
		dir:       "lang",
		locale:    DefaultLocale,
		resources: make(map[string]string),
		rw:        new(sync.RWMutex),
	}
}

// SetLocale sets language locale.
func (lang *Language) SetLocale(locale string) *Language {
	if lang == nil || lang.Error != nil {
		return lang
	}

	if locale == "" {
		lang.Error = ErrEmptyLocale()
		return lang
	}

	lang.rw.Lock()
	defer lang.rw.Unlock()

	lang.locale = locale
	fileName := fmt.Sprintf("%s/%s.json", lang.dir, locale)
	bytes, err := fs.ReadFile(fileName)
	if err != nil {
		lang.Error = ErrNotExistLocale(fileName)
		return lang
	}
	_ = json.Unmarshal(bytes, &lang.resources)
	return lang
}

// SetResources sets language resources.
func (lang *Language) SetResources(resources map[string]string) *Language {
	if lang == nil || lang.Error != nil {
		return lang
	}

	if len(resources) == 0 {
		lang.Error = ErrEmptyResources()
		return lang
	}

	lang.rw.Lock()
	defer lang.rw.Unlock()

	if len(lang.resources) == 0 {
		lang.resources = resources
	}
	for i := range resources {
		if !slices.Contains(validResourcesKey, i) {
			lang.Error = ErrInvalidResourcesError(resources)
			return lang
		}
		if _, ok := lang.resources[i]; ok {
			lang.resources[i] = resources[i]
		}
	}
	return lang
}

// returns a translated string.
func (lang *Language) translate(unit string, value int64) string {
	if lang == nil || lang.resources == nil {
		return ""
	}

	lang.rw.Lock()
	defer lang.rw.Unlock()

	if len(lang.resources) == 0 {
		lang.rw.Unlock()
		lang.SetLocale(DefaultLocale)
		lang.rw.Lock()
	}
	slice := strings.Split(lang.resources[unit], "|")
	number := getAbsValue(value)
	if len(slice) == 1 {
		return strings.Replace(slice[0], "%d", strconv.FormatInt(value, 10), 1)
	}
	if int64(len(slice)) <= number {
		return strings.Replace(slice[len(slice)-1], "%d", strconv.FormatInt(value, 10), 1)
	}
	if !strings.Contains(slice[number-1], "%d") && value < 0 {
		return "-" + slice[number-1]
	}
	return strings.Replace(slice[number-1], "%d", strconv.FormatInt(value, 10), 1)
}
