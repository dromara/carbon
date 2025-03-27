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

var (
	// invalid locale error
	// 无效的区域错误
	invalidLocaleError = func(locale string) error {
		return fmt.Errorf("invalid locale file %q, please make sure the json file exists and is valid", locale)
	}
)

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
	lang.rw.Lock()
	defer lang.rw.Unlock()

	if len(lang.resources) != 0 {
		return lang
	}
	lang.locale = locale
	fileName := lang.dir + locale + ".json"
	bytes, err := fs.ReadFile(fileName)
	if err != nil {
		lang.Error = invalidLocaleError(fileName)
		return lang
	}
	_ = json.Unmarshal(bytes, &lang.resources)
	return lang
}

// SetResources sets language resources.
// 设置资源
func (lang *Language) SetResources(resources map[string]string) *Language {
	lang.rw.Lock()
	defer lang.rw.Unlock()

	if len(lang.resources) == 0 {
		lang.resources = resources
		return lang
	}
	for k, v := range resources {
		if _, ok := lang.resources[k]; ok {
			lang.resources[k] = v
		}
	}
	return lang
}

// returns a translated string.
// 翻译转换
func (lang *Language) translate(unit string, value int64) string {
	if len(lang.resources) == 0 {
		lang.SetLocale(DefaultLocale)
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
