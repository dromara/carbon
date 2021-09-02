// +build go1.16

package carbon

import (
	"embed"
	"encoding/json"
	"io/ioutil"
)

//go:embed lang
var langFiles embed.FS

func init() {
	loadLocale = func(locale string) (map[string]string, error) {
		resources := make(map[string]string)
		fileName := langDir + locale + ".json"
		f, err := langFiles.Open(fileName)
		if err != nil {
			return resources, invalidLocaleError(locale)
		}
		defer f.Close()

		bytes, err := ioutil.ReadAll(f)
		if err != nil {
			return resources, err
		}

		if err := json.Unmarshal(bytes, &resources); err != nil {
			return resources, invalidJsonFileError(fileName)
		}
		return resources, nil
	}
}
