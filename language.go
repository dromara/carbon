package carbon

import (
	"embed"
	"encoding/json"
	"strconv"
	"strings"
	"sync"
)

//go:embed lang
var fs embed.FS

// Language defines a Language struct.
// 定义 Language 结构体
type Language struct {
	dir       string
	locale    string
	resources map[string]string
	Error     error
	rw        *sync.RWMutex
}

// NewLanguage returns a new Language instance.
// 初始化 Language 结构体
func NewLanguage() *Language {
	return &Language{
		dir:       "lang/",
		locale:    DefaultLocale,
		resources: make(map[string]string),
		rw:        new(sync.RWMutex),
	}
}

// SetLocale sets language locale.
// 设置区域
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
	fileName := lang.dir + locale + ".json"
	bytes, err := fs.ReadFile(fileName)
	if err != nil {
		lang.Error = ErrNotExistLocale(fileName)
		return lang
	}
	_ = json.Unmarshal(bytes, &lang.resources)
	return lang
}

// SetResources sets language resources.
// 设置资源
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
		if _, ok := lang.resources[k]; ok {
			lang.resources[k] = v
		} else {
			lang.Error = ErrInvalidResourcesError()
		}
	}

	if len(lang.resources) == 0 {
		lang.resources = resources
	}

	return lang
}

// returns a translated string.
// 翻译转换
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
