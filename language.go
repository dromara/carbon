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

// localeCache caches parsed locale resources to avoid repeated file loading and JSON parsing
var localeCache sync.Map

// cachedResources holds the cached resources for each language.
type cachedResources struct {
	once      sync.Once
	resources map[string]string
	err       error
}

// Language defines a Language struct.
type Language struct {
	dir       string
	locale    string
	resources map[string]string
	shared    bool
	Error     error
	rw        *sync.RWMutex
}

// NewLanguage returns a new Language instance.
func NewLanguage() *Language {
	return &Language{
		dir:    "lang",
		locale: DefaultLocale,
		rw:     new(sync.RWMutex),
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
	resources, shared := lang.resources, lang.shared
	lang.rw.RUnlock()

	// The cached resources are immutable, so a shared copy can stay shared.
	if shared {
		newLang.resources, newLang.shared = resources, true
		return newLang
	}

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

	// Early return if locale hasn't changed and resources are already loaded
	lang.rw.RLock()
	if lang.locale == locale && lang.resources != nil && len(lang.resources) > 0 {
		lang.rw.RUnlock()
		return lang
	}
	lang.rw.RUnlock()

	fileName := lang.dir + "/" + locale + ".json"
	load, ok := localeCache.Load(fileName)
	if !ok {
		load, _ = localeCache.LoadOrStore(fileName, new(cachedResources))
	}
	entry := load.(*cachedResources)

	entry.once.Do(func() {
		bs, err := fs.ReadFile(fileName)
		if err != nil {
			entry.err = fmt.Errorf("%w: %w", ErrNotExistLocale(fileName), err)
			return
		}

		var resources map[string]string
		_ = json.Unmarshal(bs, &resources)
		entry.resources = resources
	})

	if entry.err != nil {
		lang.Error = entry.err
		return lang
	}

	// Share the cached resources instead of copying them, and mark them shared so
	// that SetResources copies before it writes. This keeps the cache unmodified.
	lang.rw.Lock()
	lang.locale = locale
	lang.resources = entry.resources
	lang.shared = true
	lang.rw.Unlock()

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

	// Copy on write: the resources may be shared with the cache or another instance.
	if lang.shared || lang.resources == nil {
		newResources := make(map[string]string, len(lang.resources)+len(resources))
		for k, v := range lang.resources {
			newResources[k] = v
		}
		lang.resources = newResources
		lang.shared = false
	}
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
	if resources == nil || len(resources) == 0 {
		return ""
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
