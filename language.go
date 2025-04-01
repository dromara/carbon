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
	// empty locale error
	// 空区域错误
	emptyLocaleError = func() error {
		return fmt.Errorf("locale cannot be empty")
	}

	// invalid locale error
	// 无效的区域错误
	invalidLocaleError = func(locale string) error {
		return fmt.Errorf("invalid locale file %q, please make sure the json file exists and is valid", locale)
	}

	// invalid resources error
	// 无效的资源错误
	invalidResourcesError = func() error {
		return fmt.Errorf("invalid resources, please make sure the resources exists and is valid")
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
	if lang == nil || lang.Error != nil {
		return lang
	}

	if locale == "" {
		lang.Error = emptyLocaleError()
		return lang
	}

	lang.rw.Lock()
	defer lang.rw.Unlock()

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
	if lang == nil || lang.Error != nil {
		return lang
	}

	if resources == nil {
		lang.Error = invalidResourcesError()
		return lang
	}

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
