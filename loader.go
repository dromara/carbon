package carbon

import (
	"fmt"
)

var (
	langDir = "lang/"

	// invalidLocaleError returns an invalid locale error.
	invalidLocaleError = func(locale string) error {
		return fmt.Errorf("invalid locale %q, please see the directory %q for all valid locales", locale, langDir)
	}

	// invalidJsonFileError returns an invalid json file error.
	invalidJsonFileError = func(file string) error {
		return fmt.Errorf("invalid json file %q, please make sure the json file is valid", langDir+file)
	}
)

var loadLocale func(locale string) (map[string]string, error) = nil
