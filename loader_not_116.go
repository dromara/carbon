// +build !go1.16

package carbon

import (
	"encoding/json"
	"io/ioutil"
)

func init() {
	loadLocale = func(locale string) (map[string]string, error) {
		resources := make(map[string]string)
		fileName := langDir + locale + ".json"
		bytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			return resources, invalidLocaleError(locale)
		}
		if err := json.Unmarshal(bytes, &resources); err != nil {
			return resources, invalidJsonFileError(fileName)
		}
		return resources, nil
	}
}
