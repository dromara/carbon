package carbon

import (
	"embed"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

//go:embed lang
var fs embed.FS

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

// Copy returns a new copy of the current Language instance
func (lang *Language) Copy() *Language {
	if lang == nil {
		return nil
	}
	newLang := &Language{
		dir:    lang.dir,
		locale: lang.locale,
		Error:  lang.Error,
		rw:     new(sync.RWMutex),
	}
	if lang.resources == nil {
		return newLang
	}

	lang.rw.RLock()
	resources := lang.resources
	lang.rw.RUnlock()

	newLang.resources = make(map[string]string, len(resources))
	for k, v := range resources {
		newLang.resources[k] = v
	}
	return newLang
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

	fileName := fmt.Sprintf("%s/%s.json", lang.dir, locale)
	bs, err := fs.ReadFile(fileName)
	if err != nil {
		lang.Error = fmt.Errorf("%w: %w", ErrNotExistLocale(fileName), err)
		return lang
	}

	var resources map[string]string
	_ = json.Unmarshal(bs, &resources)

	lang.rw.Lock()
	defer lang.rw.Unlock()

	lang.locale = locale
	lang.resources = resources
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

	for k, v := range resources {
		lang.resources[k] = v
	}
	return lang
}

// returns a translated string.
func (lang *Language) translate(unit string, value int64) string {
	if lang == nil || lang.resources == nil {
		return ""
	}

	lang.rw.RLock()
	resources := lang.resources
	lang.rw.RUnlock()

	// If resources is empty, set default locale and retry
	if len(resources) == 0 {
		lang.SetLocale(DefaultLocale)
		lang.rw.RLock()
		resources = lang.resources
		lang.rw.RUnlock()
	}
	resource, exists := resources[unit]
	if !exists {
		return ""
	}
	slice := strings.Split(resource, "|")
	number := getAbsValue(value)
	str := strconv.FormatInt(value, 10)
	if len(slice) == 1 {
		return strings.Replace(slice[0], "%d", str, 1)
	}
	if int64(len(slice)) <= number {
		return strings.Replace(slice[len(slice)-1], "%d", str, 1)
	}
	if !strings.Contains(slice[number-1], "%d") && value < 0 {
		return "-" + slice[number-1]
	}
	return strings.Replace(slice[number-1], "%d", str, 1)
}
